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

// GTypeRadioButtonAccessible returns the GType for the type RadioButtonAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRadioButtonAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_radio_button_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalRadioButtonAccessible)
	return gtype
}

// RadioButtonAccessibleOverrider contains methods that are overridable.
type RadioButtonAccessibleOverrider interface {
}

type RadioButtonAccessible struct {
	_ [0]func() // equal guard
	ToggleButtonAccessible
}

var (
	_ coreglib.Objector = (*RadioButtonAccessible)(nil)
)

func classInitRadioButtonAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRadioButtonAccessible(obj *coreglib.Object) *RadioButtonAccessible {
	return &RadioButtonAccessible{
		ToggleButtonAccessible: ToggleButtonAccessible{
			ButtonAccessible: ButtonAccessible{
				ContainerAccessible: ContainerAccessible{
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
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Image: atk.Image{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioButtonAccessible(p uintptr) (interface{}, error) {
	return wrapRadioButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
