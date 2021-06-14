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
		{T: externglib.Type(C.gtk_vseparator_get_type()), F: marshalVSeparator},
	})
}

// VSeparator: the VSeparator widget is a vertical separator, used to group the
// widgets within a window. It displays a vertical line with a shadow to make it
// appear sunken into the interface.
//
// GtkVSeparator has been deprecated, use Separator instead.
type VSeparator interface {
	Separator
	Buildable
	Orientable
}

// vSeparator implements the VSeparator class.
type vSeparator struct {
	Separator
	Buildable
	Orientable
}

var _ VSeparator = (*vSeparator)(nil)

// WrapVSeparator wraps a GObject to the right type. It is
// primarily used internally.
func WrapVSeparator(obj *externglib.Object) VSeparator {
	return vSeparator{
		Separator:  WrapSeparator(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVSeparator(obj), nil
}
