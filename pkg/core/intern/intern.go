// Package intern implements value interning for Cgo sharing.
package intern

// #cgo pkg-config: gobject-2.0
// #include <glib-object.h>
// extern void goToggleNotify(gpointer, GObject*, gboolean);
import "C"

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"

	// Require a non-moving GC for heap pointers. Current GC is moving only by
	// the stack. See https://github.com/go4org/intern.
	_ "go4.org/unsafe/assume-no-moving-gc"
)

// Box contains possible interned values for each GObject.
type Box struct {
	Closures closure.Registry
	obj      unsafe.Pointer
}

// Object returns Box's C GObject pointer.
func (b *Box) Object() unsafe.Pointer { return b.obj }

// Hack to force an object on the heap.
var never bool
var sink interface{}

var (
	traceObjectFile *os.File
	objectProfile   *pprof.Profile
)

func init() {
	debug := os.Getenv("GOTK4_DEBUG")
	if debug == "" {
		return
	}

	for _, flag := range strings.Split(debug, ",") {
		switch flag {
		case "trace-objects":
			f, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("gotk4-%d-*", os.Getpid()))
			if err != nil {
				log.Panicln("cannot create temp traceObjects file:", err)
			}
			traceObjectFile = f
		case "profile-objects":
			objectProfile = pprof.NewProfile("gotk4-object-box")
		default:
			log.Panicf("unknown GOTK4_DEBUG flag %q", flag)
		}
	}
}

// newBox creates a zero-value instance of Box.
func newBox(obj unsafe.Pointer) *Box {
	box := &Box{obj: obj}

	if objectProfile != nil {
		objectProfile.Add(box.obj, 3)
	}

	if traceObjectFile != nil {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "%p: %s\n", obj, debug.Stack())
		traceObjectFile.Write(buf.Bytes())
	}

	// Force box on the heap. Objects on the stack can move, but not objects on
	// the heap. At least not for now; the assume-no-moving-gc import will
	// guard against that.
	if never {
		sink = box
	}

	return box
}

// shared contains shared closure data.
var shared = struct {
	mu sync.RWMutex
	// weak stores *Box while the object is in Go's heap. The finalizer will
	// move *Box to strong if the reference is toggled. This is only the case,
	// because the finalizer will not run otherwise.
	weak map[unsafe.Pointer]uintptr
	// strong stores *Box while the object is still referenced by C but not Go.
	strong map[unsafe.Pointer]*Box
}{
	weak:   make(map[unsafe.Pointer]uintptr, 1024),
	strong: make(map[unsafe.Pointer]*Box, 1024),
}

// ObjectClosure gets the FuncStack instance from the given GObject and GClosure
// pointers. The given unsafe.Pointers MUST be C pointers.
func ObjectClosure(gobject, gclosure unsafe.Pointer) *closure.FuncStack {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

	if box == nil {
		return nil
	}

	return box.Closures.Load(gclosure)
}

// RemoveClosure removes the given GClosure callback.
func RemoveClosure(gobject, gclosure unsafe.Pointer) {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

	if box != nil {
		box.Closures.Delete(gclosure)
	}

	// The closure missing here isn't very important. It can happen when the
	// function is called not because the user explicitly wanted to detach a
	// signal handler, but because the object is destroyed. In that case, Go
	// will have already handled it.
}

// Get gets the interned box for the given GObject C pointer. If the object is
// new or unknown, then a new box is made. If the intern box already exists for
// a given C pointer, then that box is weakly referenced and returned. The box
// will be reference-counted; the caller must use ShouldFree to unreference it.
func Get(gobject unsafe.Pointer, take bool) *Box {
	// If the registry does not exist, then we'll have to globally register it.
	// If the registry is currently strongly referenced, then we must move it to
	// a weak reference.

	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

	if box != nil {
		return box
	}

	shared.mu.Lock()

	box, _ = gets(gobject)
	if box != nil {
		shared.mu.Unlock()
		return box
	}

	box = newBox(gobject)

	// add_toggle_ref's documentation states:
	//
	//    Since a (normal) reference must be held to the object before
	//    calling g_object_add_toggle_ref(), the initial state of the
	//    reverse link is always strong.
	//
	shared.strong[gobject] = box
	runtime.SetFinalizer(box, finalizeBox)

	shared.mu.Unlock()

	// We should already have a strong reference. Sink the object in case. This
	// will force the reference to be truly strong.
	if C.g_object_is_floating(C.gpointer(gobject)) != C.FALSE {
		C.g_object_ref_sink(C.gpointer(gobject))
	}

	C.g_object_add_toggle_ref(
		(*C.GObject)(gobject),
		(*[0]byte)(C.goToggleNotify), nil,
	)

	// If we're "not taking," then we can assume our ownership over the object,
	// meaning the strong reference is now ours. That means we need to replace
	// it, not add.
	if !take {
		C.g_object_unref(C.gpointer(gobject))
	}

	// Undo the initial ref_sink.
	// C.g_object_unref(C.gpointer(gobject))

	return box
}

// finalizeBox only delays its finalization until GLib notifies us a toggle. It
// does so for as long as an object is stored only in the Go heap. Once the
// object is also shared, the toggle notifier will strongly reference the Box.
func finalizeBox(box *Box) {
	shared.mu.Lock()

	if !freeBox(box) {
		// Delegate finalizing to the next cycle.
		runtime.SetFinalizer(box, finalizeBox)
		shared.mu.Unlock()
	} else {
		shared.mu.Unlock()
		// Unreference the object. This will potentially free the object as
		// well. The closures are definitely gone at this point.
		C.g_object_remove_toggle_ref(
			(*C.GObject)(unsafe.Pointer(box.obj)),
			(*[0]byte)(C.goToggleNotify), nil,
		)

		if objectProfile != nil {
			objectProfile.Remove(box.obj)
		}
	}
}

// freeBox must only be called during finalizing of a box. It's used to know if
// a box should be freed or not during finalization. If false is returned, then
// the object must not be freed yet.
//
//go:nocheckptr
func freeBox(box *Box) bool {
	_, ok := shared.strong[box.obj]
	if ok {
		// If the closures are strong-referenced, then they might still be
		// referenced from the C side, and those closures might access this
		// object. Don't free.
		return false
	}

	_, ok = shared.weak[box.obj]
	if ok {
		// If the closures are weak-referenced, then the object reference hasn't
		// been toggled yet. Since the object is going away and we're still
		// weakly referenced, we can wipe the closures away.
		delete(shared.weak, box.obj)

		// By setting *box to a zero-value of closures, we're nilling out the
		// maps, which will signal to Go that these cyclical objects can be
		// freed altogether.
		*box = Box{obj: box.obj}
	}

	// We can proceed to free the object.
	return true
}

// goToggleNotify is called by GLib on each toggle notification. It doesn't
// actually free anything and relies on Box's finalizer to free both the box and
// the C GObject.
//
//export goToggleNotify
func goToggleNotify(_ C.gpointer, obj *C.GObject, isLastInt C.gboolean) {
	gobject := unsafe.Pointer(obj)
	isLast := isLastInt != 0

	shared.mu.Lock()
	defer shared.mu.Unlock()

	if isLast {
		// delete(shared.sharing, gobject)
		makeWeak(gobject)
	} else {
		// shared.sharing[gobject] = struct{}{}
		makeStrong(gobject)
	}
}

//go:nocheckptr
func gets(gobject unsafe.Pointer) (b *Box, strong bool) {
	if strong, ok := shared.strong[gobject]; ok {
		return strong, true
	}

	if weak, ok := shared.weak[gobject]; ok {
		// If forObject is false, then that probably means this was called
		// inside goMarshal while the Go object is still alive, otherwise
		// toggleNotify would've moved it over. We don't have to worry about
		// this being freed as long as we acquire the mutex.
		return (*Box)(unsafe.Pointer(weak)), false
	}

	return nil, false
}

// makeStrong forces the Box instance associated with the given object to be
// strongly referenced.
func makeStrong(gobject unsafe.Pointer) {
	// TODO: double mutex check, similar to ShouldFree.

	box, strong := gets(gobject)
	if box == nil {
		return
	}

	if !strong {
		shared.strong[gobject] = box
		delete(shared.weak, gobject)
	}
}

// makeWeak forces the Box intsance associated with the given object to be
// weakly referenced.
func makeWeak(gobject unsafe.Pointer) {
	box, strong := gets(gobject)
	if box == nil {
		return
	}

	if strong {
		shared.weak[gobject] = uintptr(unsafe.Pointer(box))
		delete(shared.strong, gobject)
	}
}
