// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeFlowBoxAccessible returns the GType for the type FlowBoxAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFlowBoxAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FlowBoxAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFlowBoxAccessible)
	return gtype
}

// FlowBoxAccessibleOverrider contains methods that are overridable.
type FlowBoxAccessibleOverrider interface {
}

type FlowBoxAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ coreglib.Objector = (*FlowBoxAccessible)(nil)
)

func classInitFlowBoxAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFlowBoxAccessible(obj *coreglib.Object) *FlowBoxAccessible {
	return &FlowBoxAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
	}
}

func marshalFlowBoxAccessible(p uintptr) (interface{}, error) {
	return wrapFlowBoxAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
