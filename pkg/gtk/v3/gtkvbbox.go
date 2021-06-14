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
		{T: externglib.Type(C.gtk_vbutton_box_get_type()), F: marshalVButtonBox},
	})
}

type VButtonBox interface {
	ButtonBox
	Buildable
	Orientable
}

// vButtonBox implements the VButtonBox class.
type vButtonBox struct {
	ButtonBox
	Buildable
	Orientable
}

var _ VButtonBox = (*vButtonBox)(nil)

// WrapVButtonBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapVButtonBox(obj *externglib.Object) VButtonBox {
	return vButtonBox{
		ButtonBox:  WrapButtonBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVButtonBox(obj), nil
}
