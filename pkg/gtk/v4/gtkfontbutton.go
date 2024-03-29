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
// extern void _gotk4_gtk4_FontButton_ConnectFontSet(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeFontButton = coreglib.Type(girepository.MustFind("Gtk", "FontButton").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFontButton, F: marshalFontButton},
	})
}

// FontButton: GtkFontButton allows to open a font chooser dialog to change the
// font.
//
// !An example GtkFontButton (font-button.png)
//
// It is suitable widget for selecting a font in a preference dialog.
//
// CSS nodes
//
//    fontbutton
//    ╰── button.font
//        ╰── [content]
//
//
// GtkFontButton has a single CSS node with name fontbutton which contains a
// button node with the .font style class.
type FontButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	FontChooser
}

var (
	_ Widgetter         = (*FontButton)(nil)
	_ coreglib.Objector = (*FontButton)(nil)
)

func wrapFontButton(obj *coreglib.Object) *FontButton {
	return &FontButton{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontButton(p uintptr) (interface{}, error) {
	return wrapFontButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectFontSet is emitted when the user selects a font.
//
// When handling this signal, use gtk.FontChooser.GetFont() to find out which
// font was just selected.
//
// Note that this signal is only emitted when the user changes the font. If you
// need to react to programmatic font changes as well, use the notify::font
// signal.
func (v *FontButton) ConnectFontSet(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "font-set", false, unsafe.Pointer(C._gotk4_gtk4_FontButton_ConnectFontSet), f)
}
