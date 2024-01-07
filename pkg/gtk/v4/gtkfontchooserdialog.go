// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
	GTypeFontChooserDialog = coreglib.Type(girepository.MustFind("Gtk", "FontChooserDialog").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFontChooserDialog, F: marshalFontChooserDialog},
	})
}

// FontChooserDialog: GtkFontChooserDialog widget is a dialog for selecting a
// font.
//
// !An example GtkFontChooserDialog (fontchooser.png)
//
// GtkFontChooserDialog implements the gtk.FontChooser interface and does not
// provide much API of its own.
//
// To create a GtkFontChooserDialog, use gtk.FontChooserDialog.New.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the GtkBuildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	FontChooser
}

var (
	_ coreglib.Objector = (*FontChooserDialog)(nil)
	_ Widgetter         = (*FontChooserDialog)(nil)
)

func wrapFontChooserDialog(obj *coreglib.Object) *FontChooserDialog {
	return &FontChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
				},
				Object: obj,
				Root: Root{
					NativeSurface: NativeSurface{
						Widget: Widget{
							InitiallyUnowned: coreglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
							Accessible: Accessible{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							ConstraintTarget: ConstraintTarget{
								Object: obj,
							},
						},
					},
				},
				ShortcutManager: ShortcutManager{
					Object: obj,
				},
			},
		},
		Object: obj,
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserDialog(p uintptr) (interface{}, error) {
	return wrapFontChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
