// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_child_anchor_get_type()), F: marshalTextChildAnchor},
	})
}

// TextChildAnchor: a `GtkTextChildAnchor` is a spot in a `GtkTextBuffer` where
// child widgets can be “anchored”.
//
// The anchor can have multiple widgets anchored, to allow for multiple views.
type TextChildAnchor interface {
	gextras.Objector

	// Deleted determines whether a child anchor has been deleted from the
	// buffer.
	//
	// Keep in mind that the child anchor will be unreferenced when removed from
	// the buffer, so you need to hold your own reference (with g_object_ref())
	// if you plan to use this function — otherwise all deleted child anchors
	// will also be finalized.
	Deleted() bool
}

// textChildAnchor implements the TextChildAnchor class.
type textChildAnchor struct {
	gextras.Objector
}

var _ TextChildAnchor = (*textChildAnchor)(nil)

// WrapTextChildAnchor wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextChildAnchor(obj *externglib.Object) TextChildAnchor {
	return textChildAnchor{
		Objector: obj,
	}
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextChildAnchor(obj), nil
}

// NewTextChildAnchor constructs a class TextChildAnchor.
func NewTextChildAnchor() TextChildAnchor {
	var _cret C.GtkTextChildAnchor // in

	_cret = C.gtk_text_child_anchor_new()

	var _textChildAnchor TextChildAnchor // out

	_textChildAnchor = WrapTextChildAnchor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textChildAnchor
}

// Deleted determines whether a child anchor has been deleted from the
// buffer.
//
// Keep in mind that the child anchor will be unreferenced when removed from
// the buffer, so you need to hold your own reference (with g_object_ref())
// if you plan to use this function — otherwise all deleted child anchors
// will also be finalized.
func (a textChildAnchor) Deleted() bool {
	var _arg0 *C.GtkTextChildAnchor // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_text_child_anchor_get_deleted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
