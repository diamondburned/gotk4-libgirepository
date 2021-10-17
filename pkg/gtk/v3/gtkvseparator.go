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
		{T: externglib.Type(C.gtk_vseparator_get_type()), F: marshalVSeparatorrer},
	})
}

// VSeparator widget is a vertical separator, used to group the widgets within a
// window. It displays a vertical line with a shadow to make it appear sunken
// into the interface.
//
// GtkVSeparator has been deprecated, use Separator instead.
type VSeparator struct {
	Separator
}

func wrapVSeparator(obj *externglib.Object) *VSeparator {
	return &VSeparator{
		Separator: Separator{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalVSeparatorrer(p uintptr) (interface{}, error) {
	return wrapVSeparator(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVSeparator creates a new VSeparator.
//
// Deprecated: Use gtk_separator_new() with GTK_ORIENTATION_VERTICAL instead.
func NewVSeparator() *VSeparator {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vseparator_new()

	var _vSeparator *VSeparator // out

	_vSeparator = wrapVSeparator(externglib.Take(unsafe.Pointer(_cret)))

	return _vSeparator
}

func (*VSeparator) privateVSeparator() {}
