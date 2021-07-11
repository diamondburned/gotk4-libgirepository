// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_child_accessible_get_type()), F: marshalFlowBoxChildAccessibler},
	})
}

// FlowBoxChildAccessibler describes FlowBoxChildAccessible's methods.
type FlowBoxChildAccessibler interface {
	privateFlowBoxChildAccessible()
}

//
type FlowBoxChildAccessible struct {
	ContainerAccessible
}

var (
	_ FlowBoxChildAccessibler = (*FlowBoxChildAccessible)(nil)
	_ gextras.Nativer         = (*FlowBoxChildAccessible)(nil)
)

func wrapFlowBoxChildAccessible(obj *externglib.Object) FlowBoxChildAccessibler {
	return &FlowBoxChildAccessible{
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
	}
}

func marshalFlowBoxChildAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFlowBoxChildAccessible(obj), nil
}

func (*FlowBoxChildAccessible) privateFlowBoxChildAccessible() {}
