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
		{T: externglib.Type(C.gtk_vscale_get_type()), F: marshalVScale},
	})
}

// VScale: the VScale widget is used to allow the user to select a value using a
// vertical slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkVScale has been deprecated, use Scale instead.
type VScale interface {
	Scale
	Buildable
	Orientable
}

// vScale implements the VScale class.
type vScale struct {
	Scale
	Buildable
	Orientable
}

var _ VScale = (*vScale)(nil)

// WrapVScale wraps a GObject to the right type. It is
// primarily used internally.
func WrapVScale(obj *externglib.Object) VScale {
	return vScale{
		Scale:      WrapScale(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVScale(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVScale(obj), nil
}

// NewVScale constructs a class VScale.
func NewVScale(adjustment Adjustment) VScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret C.GtkVScale      // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_vscale_new(_arg1)

	var _vScale VScale // out

	_vScale = WrapVScale(externglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}

// NewVScaleWithRange constructs a class VScale.
func NewVScaleWithRange(min float64, max float64, step float64) VScale {
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out
	var _arg3 C.gdouble   // out
	var _cret C.GtkVScale // in

	_arg1 = (C.gdouble)(min)
	_arg2 = (C.gdouble)(max)
	_arg3 = (C.gdouble)(step)

	_cret = C.gtk_vscale_new_with_range(_arg1, _arg2, _arg3)

	var _vScale VScale // out

	_vScale = WrapVScale(externglib.Take(unsafe.Pointer(_cret)))

	return _vScale
}
