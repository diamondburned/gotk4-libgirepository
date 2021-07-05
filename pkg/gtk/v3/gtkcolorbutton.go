// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_color_button_get_type()), F: marshalColorButton},
	})
}

// ColorButton: the ColorButton is a button which displays the currently
// selected color and allows to open a color selection dialog to change the
// color. It is suitable widget for selecting a color in a preference dialog.
//
//
// CSS nodes
//
// GtkColorButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .color style class.
type ColorButton interface {
	Button

	// AsActionable casts the class to the Actionable interface.
	AsActionable() Actionable
	// AsActivatable casts the class to the Activatable interface.
	AsActivatable() Activatable
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsColorChooser casts the class to the ColorChooser interface.
	AsColorChooser() ColorChooser

	// Alpha returns the current alpha value.
	//
	// Deprecated: since version 3.4.
	Alpha() uint16
	// Color sets @color to be the current color in the ColorButton widget.
	//
	// Deprecated: since version 3.4.
	Color() gdk.Color
	// Title gets the title of the color selection dialog.
	Title() string
	// UseAlpha does the color selection dialog use the alpha channel ?
	//
	// Deprecated: since version 3.4.
	UseAlpha() bool
	// SetAlphaColorButton sets the current opacity to be @alpha.
	//
	// Deprecated: since version 3.4.
	SetAlphaColorButton(alpha uint16)
	// SetColorColorButton sets the current color to be @color.
	//
	// Deprecated: since version .
	SetColorColorButton(color *gdk.Color)
	// SetTitleColorButton sets the title for the color selection dialog.
	SetTitleColorButton(title string)
	// SetUseAlphaColorButton sets whether or not the color button should use
	// the alpha channel.
	//
	// Deprecated: since version 3.4.
	SetUseAlphaColorButton(useAlpha bool)
}

// colorButton implements the ColorButton class.
type colorButton struct {
	Button
}

// WrapColorButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorButton(obj *externglib.Object) ColorButton {
	return colorButton{
		Button: WrapButton(obj),
	}
}

func marshalColorButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorButton(obj), nil
}

// NewColorButton creates a new color button.
//
// This returns a widget in the form of a small button containing a swatch
// representing the current selected color. When the button is clicked, a
// color-selection dialog will open, allowing the user to select a color. The
// swatch will be updated to reflect the new color when the user finishes.
func NewColorButton() ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton ColorButton // out

	_colorButton = WrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithColor creates a new color button.
//
// Deprecated: since version 3.4.
func NewColorButtonWithColor(color *gdk.Color) ColorButton {
	var _arg1 *C.GdkColor  // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkColor)(unsafe.Pointer(color))

	_cret = C.gtk_color_button_new_with_color(_arg1)

	var _colorButton ColorButton // out

	_colorButton = WrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithRGBA creates a new color button.
func NewColorButtonWithRGBA(rgba *gdk.RGBA) ColorButton {
	var _arg1 *C.GdkRGBA   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	_cret = C.gtk_color_button_new_with_rgba(_arg1)

	var _colorButton ColorButton // out

	_colorButton = WrapColorButton(externglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

func (b colorButton) Alpha() uint16 {
	var _arg0 *C.GtkColorButton // out
	var _cret C.guint16         // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (b colorButton) Color() gdk.Color {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.GdkColor        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	C.gtk_color_button_get_color(_arg0, &_arg1)

	var _color gdk.Color // out

	{
		var refTmpIn *C.GdkColor
		var refTmpOut *gdk.Color

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Color)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (b colorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b colorButton) UseAlpha() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_use_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b colorButton) SetAlphaColorButton(alpha uint16) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.guint16         // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_button_set_alpha(_arg0, _arg1)
}

func (b colorButton) SetColorColorButton(color *gdk.Color) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.GdkColor       // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_color_button_set_color(_arg0, _arg1)
}

func (b colorButton) SetTitleColorButton(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_color_button_set_title(_arg0, _arg1)
}

func (b colorButton) SetUseAlphaColorButton(useAlpha bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_use_alpha(_arg0, _arg1)
}

func (c colorButton) AsActionable() Actionable {
	return WrapActionable(gextras.InternObject(c))
}

func (c colorButton) AsActivatable() Activatable {
	return WrapActivatable(gextras.InternObject(c))
}

func (c colorButton) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(c))
}

func (c colorButton) AsColorChooser() ColorChooser {
	return WrapColorChooser(gextras.InternObject(c))
}
