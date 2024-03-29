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
	GTypeLinkButtonAccessible = coreglib.Type(girepository.MustFind("Gtk", "LinkButtonAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLinkButtonAccessible, F: marshalLinkButtonAccessible},
	})
}

// LinkButtonAccessibleOverrides contains methods that are overridable.
type LinkButtonAccessibleOverrides struct {
}

func defaultLinkButtonAccessibleOverrides(v *LinkButtonAccessible) LinkButtonAccessibleOverrides {
	return LinkButtonAccessibleOverrides{}
}

type LinkButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible

	*coreglib.Object
	atk.HyperlinkImpl
}

var (
	_ coreglib.Objector = (*LinkButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*LinkButtonAccessible, *LinkButtonAccessibleClass, LinkButtonAccessibleOverrides](
		GTypeLinkButtonAccessible,
		initLinkButtonAccessibleClass,
		wrapLinkButtonAccessible,
		defaultLinkButtonAccessibleOverrides,
	)
}

func initLinkButtonAccessibleClass(gclass unsafe.Pointer, overrides LinkButtonAccessibleOverrides, classInitFunc func(*LinkButtonAccessibleClass)) {
	if classInitFunc != nil {
		class := (*LinkButtonAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLinkButtonAccessible(obj *coreglib.Object) *LinkButtonAccessible {
	return &LinkButtonAccessible{
		ButtonAccessible: ButtonAccessible{
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
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
		Object: obj,
		HyperlinkImpl: atk.HyperlinkImpl{
			Object: obj,
		},
	}
}

func marshalLinkButtonAccessible(p uintptr) (interface{}, error) {
	return wrapLinkButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LinkButtonAccessibleClass: instance of this type is always passed by
// reference.
type LinkButtonAccessibleClass struct {
	*linkButtonAccessibleClass
}

// linkButtonAccessibleClass is the struct that's finalized.
type linkButtonAccessibleClass struct {
	native unsafe.Pointer
}

var GIRInfoLinkButtonAccessibleClass = girepository.MustFind("Gtk", "LinkButtonAccessibleClass")
