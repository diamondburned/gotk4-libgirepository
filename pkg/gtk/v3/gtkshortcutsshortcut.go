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
	GTypeShortcutsShortcut = coreglib.Type(girepository.MustFind("Gtk", "ShortcutsShortcut").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutsShortcut, F: marshalShortcutsShortcut},
	})
}

// ShortcutsShortcut represents a single keyboard shortcut or gesture with a
// short text. This widget is only meant to be used with ShortcutsWindow.
type ShortcutsShortcut struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ShortcutsShortcut)(nil)
	_ coreglib.Objector = (*ShortcutsShortcut)(nil)
)

func wrapShortcutsShortcut(obj *coreglib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{
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

func marshalShortcutsShortcut(p uintptr) (interface{}, error) {
	return wrapShortcutsShortcut(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
