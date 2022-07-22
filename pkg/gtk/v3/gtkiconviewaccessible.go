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

// GType values.
var (
	GTypeIconViewAccessible = coreglib.Type(C.gtk_icon_view_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIconViewAccessible, F: marshalIconViewAccessible},
	})
}

// IconViewAccessibleOverrider contains methods that are overridable.
type IconViewAccessibleOverrider interface {
}

type IconViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Selection
}

var (
	_ coreglib.Objector = (*IconViewAccessible)(nil)
)

func initClassIconViewAccessible(gclass unsafe.Pointer, goval any) {
}

func wrapIconViewAccessible(obj *coreglib.Object) *IconViewAccessible {
	return &IconViewAccessible{
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

func marshalIconViewAccessible(p uintptr) (interface{}, error) {
	return wrapIconViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
