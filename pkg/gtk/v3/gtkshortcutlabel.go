// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
	GTypeShortcutLabel = coreglib.Type(girepository.MustFind("Gtk", "ShortcutLabel").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutLabel, F: marshalShortcutLabel},
	})
}

// ShortcutLabel is a widget that represents a single keyboard shortcut or
// gesture in the user interface.
type ShortcutLabel struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ShortcutLabel)(nil)
	_ coreglib.Objector = (*ShortcutLabel)(nil)
)

func wrapShortcutLabel(obj *coreglib.Object) *ShortcutLabel {
	return &ShortcutLabel{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutLabel(p uintptr) (interface{}, error) {
	return wrapShortcutLabel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
