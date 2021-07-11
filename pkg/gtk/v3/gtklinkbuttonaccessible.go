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
		{T: externglib.Type(C.gtk_link_button_accessible_get_type()), F: marshalLinkButtonAccessibler},
	})
}

// LinkButtonAccessibler describes LinkButtonAccessible's methods.
type LinkButtonAccessibler interface {
	privateLinkButtonAccessible()
}

//
type LinkButtonAccessible struct {
	ButtonAccessible

	atk.HyperlinkImpl
}

var (
	_ LinkButtonAccessibler = (*LinkButtonAccessible)(nil)
	_ gextras.Nativer       = (*LinkButtonAccessible)(nil)
)

func wrapLinkButtonAccessible(obj *externglib.Object) LinkButtonAccessibler {
	return &LinkButtonAccessible{
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
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
		HyperlinkImpl: atk.HyperlinkImpl{
			Object: obj,
		},
	}
}

func marshalLinkButtonAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLinkButtonAccessible(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *LinkButtonAccessible) Native() uintptr {
	return v.ButtonAccessible.ContainerAccessible.WidgetAccessible.Accessible.ObjectClass.Object.Native()
}

func (*LinkButtonAccessible) privateLinkButtonAccessible() {}
