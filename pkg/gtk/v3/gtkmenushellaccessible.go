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
		{T: externglib.Type(C.gtk_menu_shell_accessible_get_type()), F: marshalMenuShellAccessible},
	})
}

type MenuShellAccessible interface {
	ContainerAccessible
}

// menuShellAccessible implements the MenuShellAccessible class.
type menuShellAccessible struct {
	ContainerAccessible
}

var _ MenuShellAccessible = (*menuShellAccessible)(nil)

// WrapMenuShellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuShellAccessible(obj *externglib.Object) MenuShellAccessible {
	return menuShellAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalMenuShellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuShellAccessible(obj), nil
}
