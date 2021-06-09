// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_button_get_type()), F: marshalFontButton},
	})
}

// FontButton: the FontButton is a button which displays the currently selected
// font an allows to open a font chooser dialog to change the font. It is
// suitable widget for selecting a font in a preference dialog.
//
//
// CSS nodes
//
// GtkFontButton has a single CSS node with name button and style class .font.
type FontButton interface {
	Button
	Actionable
	Activatable
	Buildable
	FontChooser

	// FontName retrieves the name of the currently selected font. This name
	// includes style and size information as well. If you want to render
	// something with the font, use this string with
	// pango_font_description_from_string() . If you’re interested in peeking
	// certain values (family name, style, size, weight) just query these
	// properties from the FontDescription object.
	FontName() string
	// ShowSize returns whether the font size will be shown in the label.
	ShowSize() bool
	// ShowStyle returns whether the name of the font style will be shown in the
	// label.
	ShowStyle() bool
	// Title retrieves the title of the font chooser dialog.
	Title() string
	// UseFont returns whether the selected font is used in the label.
	UseFont() bool
	// UseSize returns whether the selected size is used in the label.
	UseSize() bool
	// SetFontName sets or updates the currently-displayed font in font picker
	// dialog.
	SetFontName(fontname string) bool
	// SetShowSize: if @show_size is true, the font size will be displayed along
	// with the name of the selected font.
	SetShowSize(showSize bool)
	// SetShowStyle: if @show_style is true, the font style will be displayed
	// along with name of the selected font.
	SetShowStyle(showStyle bool)
	// SetTitle sets the title for the font chooser dialog.
	SetTitle(title string)
	// SetUseFont: if @use_font is true, the font name will be written using the
	// selected font.
	SetUseFont(useFont bool)
	// SetUseSize: if @use_size is true, the font name will be written using the
	// selected size.
	SetUseSize(useSize bool)
}

// fontButton implements the FontButton interface.
type fontButton struct {
	Button
	Actionable
	Activatable
	Buildable
	FontChooser
}

var _ FontButton = (*fontButton)(nil)

// WrapFontButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontButton(obj *externglib.Object) FontButton {
	return FontButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
		FontChooser: WrapFontChooser(obj),
	}
}

func marshalFontButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontButton(obj), nil
}

// NewFontButton constructs a class FontButton.
func NewFontButton() FontButton {
	var _cret C.GtkFontButton

	cret = C.gtk_font_button_new()

	var _fontButton FontButton

	_fontButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FontButton)

	return _fontButton
}

// NewFontButtonWithFont constructs a class FontButton.
func NewFontButtonWithFont(fontname string) FontButton {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkFontButton

	cret = C.gtk_font_button_new_with_font(_arg1)

	var _fontButton FontButton

	_fontButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FontButton)

	return _fontButton
}

// FontName retrieves the name of the currently selected font. This name
// includes style and size information as well. If you want to render
// something with the font, use this string with
// pango_font_description_from_string() . If you’re interested in peeking
// certain values (family name, style, size, weight) just query these
// properties from the FontDescription object.
func (f fontButton) FontName() string {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret *C.gchar

	cret = C.gtk_font_button_get_font_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ShowSize returns whether the font size will be shown in the label.
func (f fontButton) ShowSize() bool {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean

	cret = C.gtk_font_button_get_show_size(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowStyle returns whether the name of the font style will be shown in the
// label.
func (f fontButton) ShowStyle() bool {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean

	cret = C.gtk_font_button_get_show_style(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
func (f fontButton) Title() string {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret *C.gchar

	cret = C.gtk_font_button_get_title(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
func (f fontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean

	cret = C.gtk_font_button_get_use_font(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
func (f fontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean

	cret = C.gtk_font_button_get_use_size(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetFontName sets or updates the currently-displayed font in font picker
// dialog.
func (f fontButton) SetFontName(fontname string) bool {
	var _arg0 *C.GtkFontButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.gtk_font_button_set_font_name(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetShowSize: if @show_size is true, the font size will be displayed along
// with the name of the selected font.
func (f fontButton) SetShowSize(showSize bool) {
	var _arg0 *C.GtkFontButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if showSize {
		_arg1 = C.gboolean(1)
	}

	C.gtk_font_button_set_show_size(_arg0, _arg1)
}

// SetShowStyle: if @show_style is true, the font style will be displayed
// along with name of the selected font.
func (f fontButton) SetShowStyle(showStyle bool) {
	var _arg0 *C.GtkFontButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if showStyle {
		_arg1 = C.gboolean(1)
	}

	C.gtk_font_button_set_show_style(_arg0, _arg1)
}

// SetTitle sets the title for the font chooser dialog.
func (f fontButton) SetTitle(title string) {
	var _arg0 *C.GtkFontButton
	var _arg1 *C.gchar

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_button_set_title(_arg0, _arg1)
}

// SetUseFont: if @use_font is true, the font name will be written using the
// selected font.
func (f fontButton) SetUseFont(useFont bool) {
	var _arg0 *C.GtkFontButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if useFont {
		_arg1 = C.gboolean(1)
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
}

// SetUseSize: if @use_size is true, the font name will be written using the
// selected size.
func (f fontButton) SetUseSize(useSize bool) {
	var _arg0 *C.GtkFontButton
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(f.Native()))
	if useSize {
		_arg1 = C.gboolean(1)
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
}
