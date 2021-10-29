// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_arrow_accessible_get_type()), F: marshalArrowAccessibler},
	})
}

type ArrowAccessible struct {
	WidgetAccessible

	atk.Image
}

var (
	_ externglib.Objector = (*ArrowAccessible)(nil)
)

func wrapArrowAccessible(obj *externglib.Object) *ArrowAccessible {
	return &ArrowAccessible{
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
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalArrowAccessibler(p uintptr) (interface{}, error) {
	return wrapArrowAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
