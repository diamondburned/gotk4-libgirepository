// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hscrollbar_get_type()), F: marshalHScrollbar},
	})
}

// HScrollbar: the HScrollbar widget is a widget arranged horizontally creating
// a scrollbar. See Scrollbar for details on scrollbars. Adjustment pointers may
// be added to handle the adjustment of the scrollbar or it may be left nil in
// which case one will be created for you. See Scrollbar for a description of
// what the fields in an adjustment represent for a scrollbar.
//
// GtkHScrollbar has been deprecated, use Scrollbar instead.
type HScrollbar interface {
	Scrollbar
	Buildable
	Orientable
}

// hScrollbar implements the HScrollbar class.
type hScrollbar struct {
	Scrollbar
	Buildable
	Orientable
}

var _ HScrollbar = (*hScrollbar)(nil)

// WrapHScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapHScrollbar(obj *externglib.Object) HScrollbar {
	return hScrollbar{
		Scrollbar:  WrapScrollbar(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHScrollbar(obj), nil
}

// NewHScrollbar constructs a class HScrollbar.
func NewHScrollbar(adjustment Adjustment) HScrollbar {
	var _arg1 *C.GtkAdjustment // out
	var _cret C.GtkHScrollbar  // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_hscrollbar_new(_arg1)

	var _hScrollbar HScrollbar // out

	_hScrollbar = WrapHScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _hScrollbar
}
