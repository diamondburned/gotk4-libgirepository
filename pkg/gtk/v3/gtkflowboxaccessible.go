// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeFlowBoxAccessible = coreglib.Type(C.gtk_flow_box_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFlowBoxAccessible, F: marshalFlowBoxAccessible},
	})
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

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeFlowBoxAccessible,
		GoType:    reflect.TypeOf((*FlowBoxAccessible)(nil)),
		InitClass: initClassFlowBoxAccessible,
	})
}

func initClassFlowBoxAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitFlowBoxAccessible(*FlowBoxAccessibleClass) }); ok {
		klass := (*FlowBoxAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitFlowBoxAccessible(klass)
	}
}

func wrapFlowBoxAccessible(obj *coreglib.Object) *FlowBoxAccessible {
	return &FlowBoxAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
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

// FlowBoxAccessibleClass: instance of this type is always passed by reference.
type FlowBoxAccessibleClass struct {
	*flowBoxAccessibleClass
}

// flowBoxAccessibleClass is the struct that's finalized.
type flowBoxAccessibleClass struct {
	native *C.GtkFlowBoxAccessibleClass
}

func (f *FlowBoxAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &f.native.parent_class
	var v *ContainerAccessibleClass // out
	v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
