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

// GTypeScaleAccessible returns the GType for the type ScaleAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScaleAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_scale_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalScaleAccessible)
	return gtype
}

// ScaleAccessibleOverrider contains methods that are overridable.
type ScaleAccessibleOverrider interface {
}

type ScaleAccessible struct {
	_ [0]func() // equal guard
	RangeAccessible
}

var (
	_ coreglib.Objector = (*ScaleAccessible)(nil)
)

func classInitScaleAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapScaleAccessible(obj *coreglib.Object) *ScaleAccessible {
	return &ScaleAccessible{
		RangeAccessible: RangeAccessible{
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
		},
	}
}

func marshalScaleAccessible(p uintptr) (interface{}, error) {
	return wrapScaleAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
