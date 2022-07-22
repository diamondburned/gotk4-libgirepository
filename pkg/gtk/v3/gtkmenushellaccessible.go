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
	GTypeMenuShellAccessible = coreglib.Type(C.gtk_menu_shell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuShellAccessible, F: marshalMenuShellAccessible},
	})
}

// MenuShellAccessibleOverrider contains methods that are overridable.
type MenuShellAccessibleOverrider interface {
}

type MenuShellAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ coreglib.Objector = (*MenuShellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeMenuShellAccessible,
		GoType:    reflect.TypeOf((*MenuShellAccessible)(nil)),
		InitClass: initClassMenuShellAccessible,
	})
}

func initClassMenuShellAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitMenuShellAccessible(*MenuShellAccessibleClass)
	}); ok {
		klass := (*MenuShellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitMenuShellAccessible(klass)
	}
}

func wrapMenuShellAccessible(obj *coreglib.Object) *MenuShellAccessible {
	return &MenuShellAccessible{
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

func marshalMenuShellAccessible(p uintptr) (interface{}, error) {
	return wrapMenuShellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MenuShellAccessibleClass: instance of this type is always passed by
// reference.
type MenuShellAccessibleClass struct {
	*menuShellAccessibleClass
}

// menuShellAccessibleClass is the struct that's finalized.
type menuShellAccessibleClass struct {
	native *C.GtkMenuShellAccessibleClass
}

func (m *MenuShellAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &m.native.parent_class
	var v *ContainerAccessibleClass // out
	v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
