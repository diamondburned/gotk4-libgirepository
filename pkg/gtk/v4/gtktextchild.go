// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_child_anchor_get_type()), F: marshalTextChildAnchor},
	})
}

// TextChildAnchor: a TextChildAnchor is a spot in the buffer where child
// widgets can be “anchored” (inserted inline, as if they were characters). The
// anchor can have multiple widgets anchored, to allow for multiple views.
type TextChildAnchor interface {
	gextras.Objector

	// Deleted determines whether a child anchor has been deleted from the
	// buffer. Keep in mind that the child anchor will be unreferenced when
	// removed from the buffer, so you need to hold your own reference (with
	// g_object_ref()) if you plan to use this function — otherwise all deleted
	// child anchors will also be finalized.
	Deleted() bool
	// Widgets gets a list of all widgets anchored at this child anchor.
	//
	// The order in which the widgets are returned is not defined.
	Widgets() []Widget
}

// textChildAnchor implements the TextChildAnchor interface.
type textChildAnchor struct {
	gextras.Objector
}

var _ TextChildAnchor = (*textChildAnchor)(nil)

// WrapTextChildAnchor wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextChildAnchor(obj *externglib.Object) TextChildAnchor {
	return TextChildAnchor{
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
	var _cret C.GtkTextChildAnchor

	cret = C.gtk_text_child_anchor_new()

	var _textChildAnchor TextChildAnchor

	_textChildAnchor = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TextChildAnchor)

	return _textChildAnchor
}

// Deleted determines whether a child anchor has been deleted from the
// buffer. Keep in mind that the child anchor will be unreferenced when
// removed from the buffer, so you need to hold your own reference (with
// g_object_ref()) if you plan to use this function — otherwise all deleted
// child anchors will also be finalized.
func (a textChildAnchor) Deleted() bool {
	var _arg0 *C.GtkTextChildAnchor

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_text_child_anchor_get_deleted(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Widgets gets a list of all widgets anchored at this child anchor.
//
// The order in which the widgets are returned is not defined.
func (a textChildAnchor) Widgets() []Widget {
	var _arg0 *C.GtkTextChildAnchor

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	var _cret **C.GtkWidget
	var _arg1 *C.guint

	cret = C.gtk_text_child_anchor_get_widgets(_arg0)

	var _widgets []Widget

	{
		var src []*C.GtkWidget
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg1))

		_widgets = make([]Widget, _arg1)
		for i := 0; i < uintptr(_arg1); i++ {
			_widgets = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)
		}
	}

	return _widgets
}
