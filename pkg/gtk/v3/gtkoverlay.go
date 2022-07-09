// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_OverlayClass_get_child_position(void*, void*, void*);
// extern gboolean _gotk4_gtk3_Overlay_ConnectGetChildPosition(gpointer, void*, void*, guintptr);
import "C"

// GTypeOverlay returns the GType for the type Overlay.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeOverlay() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Overlay").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalOverlay)
	return gtype
}

// OverlayOverrider contains methods that are overridable.
type OverlayOverrider interface {
	// The function takes the following parameters:
	//
	//    - widget
	//    - allocation
	//
	// The function returns the following values:
	//
	ChildPosition(widget Widgetter, allocation *Allocation) bool
}

// Overlay is a container which contains a single main child, on top of which it
// can place “overlay” widgets. The position of each overlay widget is
// determined by its Widget:halign and Widget:valign properties. E.g. a widget
// with both alignments set to GTK_ALIGN_START will be placed at the top left
// corner of the GtkOverlay container, whereas an overlay with halign set to
// GTK_ALIGN_CENTER and valign set to GTK_ALIGN_END will be placed a the bottom
// edge of the GtkOverlay, horizontally centered. The position can be adjusted
// by setting the margin properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting to the
// Overlay::get-child-position signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The GtkOverlay implementation of the GtkBuildable interface supports placing
// a child as an overlay by specifying “overlay” as the “type” attribute of a
// <child> element.
//
//
// CSS nodes
//
// GtkOverlay has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Overlay)(nil)
)

func classInitOverlayer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "OverlayClass")

	if _, ok := goval.(interface {
		ChildPosition(widget Widgetter, allocation *Allocation) bool
	}); ok {
		o := pclass.StructFieldOffset("get_child_position")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_OverlayClass_get_child_position)
	}
}

//export _gotk4_gtk3_OverlayClass_get_child_position
func _gotk4_gtk3_OverlayClass_get_child_position(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		ChildPosition(widget Widgetter, allocation *Allocation) bool
	})

	var _widget Widgetter       // out
	var _allocation *Allocation // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}
	_allocation = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.ChildPosition(_widget, _allocation)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapOverlay(obj *coreglib.Object) *Overlay {
	return &Overlay{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	return wrapOverlay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Overlay_ConnectGetChildPosition
func _gotk4_gtk3_Overlay_ConnectGetChildPosition(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(widget Widgetter) (allocation *gdk.Rectangle, ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter) (allocation *gdk.Rectangle, ok bool))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	allocation, ok := f(_widget)

	*arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(allocation)))
	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectGetChildPosition signal is emitted to determine the position and size
// of any overlay child widgets. A handler for this signal should fill
// allocation with the desired position and size for widget, relative to the
// 'main' child of overlay.
//
// The default handler for this signal uses the widget's halign and valign
// properties to determine the position and gives the widget its natural size
// (except that an alignment of GTK_ALIGN_FILL will cause the overlay to be
// full-width/height). If the main child is a ScrolledWindow, the overlays are
// placed relative to its contents.
func (overlay *Overlay) ConnectGetChildPosition(f func(widget Widgetter) (allocation *gdk.Rectangle, ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(overlay, "get-child-position", false, unsafe.Pointer(C._gotk4_gtk3_Overlay_ConnectGetChildPosition), f)
}

// NewOverlay creates a new Overlay.
//
// The function returns the following values:
//
//    - overlay: new Overlay object.
//
func NewOverlay() *Overlay {
	_gret := girepository.MustFind("Gtk", "Overlay").InvokeMethod("new_Overlay", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _overlay *Overlay // out

	_overlay = wrapOverlay(coreglib.Take(unsafe.Pointer(_cret)))

	return _overlay
}

// AddOverlay adds widget to overlay.
//
// The widget will be stacked on top of the main widget added with
// gtk_container_add().
//
// The position at which widget is placed is determined from its Widget:halign
// and Widget:valign properties.
//
// The function takes the following parameters:
//
//    - widget to be added to the container.
//
func (overlay *Overlay) AddOverlay(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	girepository.MustFind("Gtk", "Overlay").InvokeMethod("add_overlay", _args[:], nil)

	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
}

// OverlayPassThrough: convenience function to get the value of the
// Overlay:pass-through child property for widget.
//
// The function takes the following parameters:
//
//    - widget: overlay child of Overlay.
//
// The function returns the following values:
//
//    - ok: whether the widget is a pass through child.
//
func (overlay *Overlay) OverlayPassThrough(widget Widgetter) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "Overlay").InvokeMethod("get_overlay_pass_through", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ReorderOverlay moves child to a new index in the list of overlay children.
// The list contains overlays in the order that these were added to overlay by
// default. See also Overlay:index.
//
// A widget’s index in the overlay children list determines which order the
// children are drawn if they overlap. The first child is drawn at the bottom.
// It also affects the default focus chain order.
//
// The function takes the following parameters:
//
//    - child: overlaid Widget to move.
//    - index_: new index for child in the list of overlay children of overlay,
//      starting from 0. If negative, indicates the end of the list.
//
func (overlay *Overlay) ReorderOverlay(child Widgetter, index_ int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(index_)

	girepository.MustFind("Gtk", "Overlay").InvokeMethod("reorder_overlay", _args[:], nil)

	runtime.KeepAlive(overlay)
	runtime.KeepAlive(child)
	runtime.KeepAlive(index_)
}

// SetOverlayPassThrough: convenience function to set the value of the
// Overlay:pass-through child property for widget.
//
// The function takes the following parameters:
//
//    - widget: overlay child of Overlay.
//    - passThrough: whether the child should pass the input through.
//
func (overlay *Overlay) SetOverlayPassThrough(widget Widgetter, passThrough bool) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if passThrough {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Overlay").InvokeMethod("set_overlay_pass_through", _args[:], nil)

	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(passThrough)
}
