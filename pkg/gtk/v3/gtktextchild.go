// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeTextChildAnchor returns the GType for the type TextChildAnchor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTextChildAnchor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TextChildAnchor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTextChildAnchor)
	return gtype
}

// TextChildAnchorOverrider contains methods that are overridable.
type TextChildAnchorOverrider interface {
}

// TextChildAnchor is a spot in the buffer where child widgets can be “anchored”
// (inserted inline, as if they were characters). The anchor can have multiple
// widgets anchored, to allow for multiple views.
type TextChildAnchor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TextChildAnchor)(nil)
)

func classInitTextChildAnchorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTextChildAnchor(obj *coreglib.Object) *TextChildAnchor {
	return &TextChildAnchor{
		Object: obj,
	}
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	return wrapTextChildAnchor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextChildAnchor creates a new TextChildAnchor. Usually you would then
// insert it into a TextBuffer with gtk_text_buffer_insert_child_anchor(). To
// perform the creation and insertion in one step, use the convenience function
// gtk_text_buffer_create_child_anchor().
//
// The function returns the following values:
//
//    - textChildAnchor: new TextChildAnchor.
//
func NewTextChildAnchor() *TextChildAnchor {
	_gret := girepository.MustFind("Gtk", "TextChildAnchor").InvokeMethod("new_TextChildAnchor", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _textChildAnchor *TextChildAnchor // out

	_textChildAnchor = wrapTextChildAnchor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textChildAnchor
}

// Deleted determines whether a child anchor has been deleted from the buffer.
// Keep in mind that the child anchor will be unreferenced when removed from the
// buffer, so you need to hold your own reference (with g_object_ref()) if you
// plan to use this function — otherwise all deleted child anchors will also be
// finalized.
//
// The function returns the following values:
//
//    - ok: TRUE if the child anchor has been deleted from its buffer.
//
func (anchor *TextChildAnchor) Deleted() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(anchor).Native()))

	_gret := girepository.MustFind("Gtk", "TextChildAnchor").InvokeMethod("get_deleted", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(anchor)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Widgets gets a list of all widgets anchored at this child anchor. The
// returned list should be freed with g_list_free().
//
// The function returns the following values:
//
//    - list of widgets anchored at anchor.
//
func (anchor *TextChildAnchor) Widgets() []Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(anchor).Native()))

	_gret := girepository.MustFind("Gtk", "TextChildAnchor").InvokeMethod("get_widgets", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(anchor)

	var _list []Widgetter // out

	_list = make([]Widgetter, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst Widgetter // out
		{
			objptr := unsafe.Pointer(src)
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
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}
