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
	GTypePlugAccessible = coreglib.Type(girepository.MustFind("Gtk", "PlugAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePlugAccessible, F: marshalPlugAccessible},
	})
}

// PlugAccessibleOverrides contains methods that are overridable.
type PlugAccessibleOverrides struct {
}

func defaultPlugAccessibleOverrides(v *PlugAccessible) PlugAccessibleOverrides {
	return PlugAccessibleOverrides{}
}

type PlugAccessible struct {
	_ [0]func() // equal guard
	WindowAccessible
}

var (
	_ coreglib.Objector = (*PlugAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PlugAccessible, *PlugAccessibleClass, PlugAccessibleOverrides](
		GTypePlugAccessible,
		initPlugAccessibleClass,
		wrapPlugAccessible,
		defaultPlugAccessibleOverrides,
	)
}

func initPlugAccessibleClass(gclass unsafe.Pointer, overrides PlugAccessibleOverrides, classInitFunc func(*PlugAccessibleClass)) {
	if classInitFunc != nil {
		class := (*PlugAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPlugAccessible(obj *coreglib.Object) *PlugAccessible {
	return &PlugAccessible{
		WindowAccessible: WindowAccessible{
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
			Window: atk.Window{
				AtkObject: atk.AtkObject{
					Object: obj,
				},
			},
		},
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	return wrapPlugAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PlugAccessibleClass: instance of this type is always passed by reference.
type PlugAccessibleClass struct {
	*plugAccessibleClass
}

// plugAccessibleClass is the struct that's finalized.
type plugAccessibleClass struct {
	native unsafe.Pointer
}

var GIRInfoPlugAccessibleClass = girepository.MustFind("Gtk", "PlugAccessibleClass")
