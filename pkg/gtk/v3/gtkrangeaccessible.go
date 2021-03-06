// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeRangeAccessible returns the GType for the type RangeAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRangeAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_range_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalRangeAccessible)
	return gtype
}

// RangeAccessibleOverrider contains methods that are overridable.
type RangeAccessibleOverrider interface {
}

type RangeAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Value
}

var (
	_ coreglib.Objector = (*RangeAccessible)(nil)
)

func classInitRangeAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRangeAccessible(obj *coreglib.Object) *RangeAccessible {
	return &RangeAccessible{
		WidgetAccessible: WidgetAccessible{
			Accessible: Accessible{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Component: atk.Component{
				Object: obj,
			},
		},
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalRangeAccessible(p uintptr) (interface{}, error) {
	return wrapRangeAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
