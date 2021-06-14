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
		{T: externglib.Type(C.gtk_paned_accessible_get_type()), F: marshalPanedAccessible},
	})
}

type PanedAccessible interface {
	ContainerAccessible
}

// panedAccessible implements the PanedAccessible class.
type panedAccessible struct {
	ContainerAccessible
}

var _ PanedAccessible = (*panedAccessible)(nil)

// WrapPanedAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapPanedAccessible(obj *externglib.Object) PanedAccessible {
	return panedAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalPanedAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPanedAccessible(obj), nil
}
