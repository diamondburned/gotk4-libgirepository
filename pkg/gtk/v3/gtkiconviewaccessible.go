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
	GTypeIconViewAccessible = coreglib.Type(girepository.MustFind("Gtk", "IconViewAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIconViewAccessible, F: marshalIconViewAccessible},
	})
}

// IconViewAccessibleOverrides contains methods that are overridable.
type IconViewAccessibleOverrides struct {
}

func defaultIconViewAccessibleOverrides(v *IconViewAccessible) IconViewAccessibleOverrides {
	return IconViewAccessibleOverrides{}
}

type IconViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ coreglib.Objector = (*IconViewAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*IconViewAccessible, *IconViewAccessibleClass, IconViewAccessibleOverrides](
		GTypeIconViewAccessible,
		initIconViewAccessibleClass,
		wrapIconViewAccessible,
		defaultIconViewAccessibleOverrides,
	)
}

func initIconViewAccessibleClass(gclass unsafe.Pointer, overrides IconViewAccessibleOverrides, classInitFunc func(*IconViewAccessibleClass)) {
	if classInitFunc != nil {
		class := (*IconViewAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapIconViewAccessible(obj *coreglib.Object) *IconViewAccessible {
	return &IconViewAccessible{
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

func marshalIconViewAccessible(p uintptr) (interface{}, error) {
	return wrapIconViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// IconViewAccessibleClass: instance of this type is always passed by reference.
type IconViewAccessibleClass struct {
	*iconViewAccessibleClass
}

// iconViewAccessibleClass is the struct that's finalized.
type iconViewAccessibleClass struct {
	native unsafe.Pointer
}

var GIRInfoIconViewAccessibleClass = girepository.MustFind("Gtk", "IconViewAccessibleClass")
