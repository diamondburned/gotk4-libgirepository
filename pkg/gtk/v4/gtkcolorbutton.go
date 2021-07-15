// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_button_get_type()), F: marshalColorButtoner},
	})
}

// ColorButtoner describes ColorButton's methods.
type ColorButtoner interface {
	// Modal gets whether the dialog is modal.
	Modal() bool
	// Title gets the title of the color chooser dialog.
	Title() string
	// SetModal sets whether the dialog should be modal.
	SetModal(modal bool)
	// SetTitle sets the title for the color chooser dialog.
	SetTitle(title string)
}

// ColorButton: GtkColorButton allows to open a color chooser dialog to change
// the color.
//
// !An example GtkColorButton (color-button.png)
//
// It is suitable widget for selecting a color in a preference dialog.
//
// CSS nodes
//
//    colorbutton
//    ╰── button.color
//        ╰── [content]
//
//
// GtkColorButton has a single CSS node with name colorbutton which contains a
// button node. To differentiate it from a plain GtkButton, it gets the .color
// style class.
type ColorButton struct {
	Widget

	ColorChooser
}

var (
	_ ColorButtoner   = (*ColorButton)(nil)
	_ gextras.Nativer = (*ColorButton)(nil)
)

func wrapColorButton(obj *externglib.Object) *ColorButton {
	return &ColorButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorButton(obj), nil
}

// NewColorButton creates a new color button.
//
// This returns a widget in the form of a small button containing a swatch
// representing the current selected color. When the button is clicked, a color
// chooser dialog will open, allowing the user to select a color. The swatch
// will be updated to reflect the new color when the user finishes.
func NewColorButton() *ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithRGBA creates a new color button showing the given color.
func NewColorButtonWithRGBA(rgba *gdk.RGBA) *ColorButton {
	var _arg1 *C.GdkRGBA   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	_cret = C.gtk_color_button_new_with_rgba(_arg1)

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ColorButton) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// Modal gets whether the dialog is modal.
func (button *ColorButton) Modal() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_color_button_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of the color chooser dialog.
func (button *ColorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_color_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetModal sets whether the dialog should be modal.
func (button *ColorButton) SetModal(modal bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_modal(_arg0, _arg1)
}

// SetTitle sets the title for the color chooser dialog.
func (button *ColorButton) SetTitle(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))

	C.gtk_color_button_set_title(_arg0, _arg1)
}
