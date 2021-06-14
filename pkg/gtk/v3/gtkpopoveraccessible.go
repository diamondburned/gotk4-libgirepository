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
		{T: externglib.Type(C.gtk_popover_accessible_get_type()), F: marshalPopoverAccessible},
	})
}

type PopoverAccessible interface {
	ContainerAccessible
}

// popoverAccessible implements the PopoverAccessible class.
type popoverAccessible struct {
	ContainerAccessible
}

var _ PopoverAccessible = (*popoverAccessible)(nil)

// WrapPopoverAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopoverAccessible(obj *externglib.Object) PopoverAccessible {
	return popoverAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalPopoverAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopoverAccessible(obj), nil
}
