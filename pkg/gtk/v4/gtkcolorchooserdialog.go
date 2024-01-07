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
	GTypeColorChooserDialog = coreglib.Type(girepository.MustFind("Gtk", "ColorChooserDialog").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorChooserDialog, F: marshalColorChooserDialog},
	})
}

// ColorChooserDialog: dialog for choosing a color.
//
// !An example GtkColorChooserDialog (colorchooser.png)
//
// GtkColorChooserDialog implements the gtk.ColorChooser interface and does not
// provide much API of its own.
//
// To create a GtkColorChooserDialog, use gtk.ColorChooserDialog.New.
//
// To change the initially selected color, use gtk.ColorChooser.SetRGBA(). To
// get the selected color use gtk.ColorChooser.GetRGBA().
type ColorChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserDialog)(nil)
	_ Widgetter         = (*ColorChooserDialog)(nil)
)

func wrapColorChooserDialog(obj *coreglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
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
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	return wrapColorChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
