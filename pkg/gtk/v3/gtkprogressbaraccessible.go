// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_progress_bar_accessible_get_type()), F: marshalProgressBarAccessibler},
	})
}

type ProgressBarAccessible struct {
	WidgetAccessible

	atk.Value
}

func wrapProgressBarAccessible(obj *externglib.Object) *ProgressBarAccessible {
	return &ProgressBarAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalProgressBarAccessibler(p uintptr) (interface{}, error) {
	return wrapProgressBarAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*ProgressBarAccessible) privateProgressBarAccessible() {}
