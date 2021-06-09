// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
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
	// Widgets gets a list of all widgets anchored at this child anchor. The
	// returned list should be freed with g_list_free().
	Widgets() *glib.List
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

// Widgets gets a list of all widgets anchored at this child anchor. The
// returned list should be freed with g_list_free().
func (a textChildAnchor) Widgets() *glib.List {
	var _arg0 *C.GtkTextChildAnchor

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(a.Native()))

	var _cret *C.GList

	cret = C.gtk_text_child_anchor_get_widgets(_arg0)

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _list
}