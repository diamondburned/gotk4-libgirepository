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
		{T: externglib.Type(C.gtk_arrow_accessible_get_type()), F: marshalArrowAccessible},
	})
}

type ArrowAccessible interface {
	WidgetAccessible
}

// arrowAccessible implements the ArrowAccessible class.
type arrowAccessible struct {
	WidgetAccessible
}

var _ ArrowAccessible = (*arrowAccessible)(nil)

// WrapArrowAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapArrowAccessible(obj *externglib.Object) ArrowAccessible {
	return arrowAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalArrowAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapArrowAccessible(obj), nil
}
