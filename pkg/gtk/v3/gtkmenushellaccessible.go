// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeMenuShellAccessible = coreglib.Type(girepository.MustFind("Gtk", "MenuShellAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuShellAccessible, F: marshalMenuShellAccessible},
	})
}

// MenuShellAccessibleOverrides contains methods that are overridable.
type MenuShellAccessibleOverrides struct {
}

func defaultMenuShellAccessibleOverrides(v *MenuShellAccessible) MenuShellAccessibleOverrides {
	return MenuShellAccessibleOverrides{}
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
	coreglib.RegisterClassInfo[*MenuShellAccessible, *MenuShellAccessibleClass, MenuShellAccessibleOverrides](
		GTypeMenuShellAccessible,
		initMenuShellAccessibleClass,
		wrapMenuShellAccessible,
		defaultMenuShellAccessibleOverrides,
	)
}

func initMenuShellAccessibleClass(gclass unsafe.Pointer, overrides MenuShellAccessibleOverrides, classInitFunc func(*MenuShellAccessibleClass)) {
	if classInitFunc != nil {
		class := (*MenuShellAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
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
	native unsafe.Pointer
}

var GIRInfoMenuShellAccessibleClass = girepository.MustFind("Gtk", "MenuShellAccessibleClass")
