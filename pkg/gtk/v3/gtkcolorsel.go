// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_color_selection_get_type()), F: marshalColorSelectioner},
	})
}

// ColorSelectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ColorSelectionOverrider interface {
	ColorChanged()
}

// ColorSelectioner describes ColorSelection's methods.
type ColorSelectioner interface {
	// CurrentAlpha returns the current alpha value.
	CurrentAlpha() uint16
	// CurrentColor sets color to be the current color in the GtkColorSelection
	// widget.
	CurrentColor() gdk.Color
	// CurrentRGBA sets rgba to be the current color in the GtkColorSelection
	// widget.
	CurrentRGBA() gdk.RGBA
	// HasOpacityControl determines whether the colorsel has an opacity control.
	HasOpacityControl() bool
	// HasPalette determines whether the color selector has a color palette.
	HasPalette() bool
	// PreviousAlpha returns the previous alpha value.
	PreviousAlpha() uint16
	// PreviousColor fills color in with the original color value.
	PreviousColor() gdk.Color
	// PreviousRGBA fills rgba in with the original color value.
	PreviousRGBA() gdk.RGBA
	// IsAdjusting gets the current state of the colorsel.
	IsAdjusting() bool
	// SetCurrentAlpha sets the current opacity to be alpha.
	SetCurrentAlpha(alpha uint16)
	// SetCurrentColor sets the current color to be color.
	SetCurrentColor(color *gdk.Color)
	// SetCurrentRGBA sets the current color to be rgba.
	SetCurrentRGBA(rgba *gdk.RGBA)
	// SetHasOpacityControl sets the colorsel to use or not use opacity.
	SetHasOpacityControl(hasOpacity bool)
	// SetHasPalette shows and hides the palette based upon the value of
	// has_palette.
	SetHasPalette(hasPalette bool)
	// SetPreviousAlpha sets the “previous” alpha to be alpha.
	SetPreviousAlpha(alpha uint16)
	// SetPreviousColor sets the “previous” color to be color.
	SetPreviousColor(color *gdk.Color)
	// SetPreviousRGBA sets the “previous” color to be rgba.
	SetPreviousRGBA(rgba *gdk.RGBA)
}

type ColorSelection struct {
	Box
}

var (
	_ ColorSelectioner = (*ColorSelection)(nil)
	_ gextras.Nativer  = (*ColorSelection)(nil)
)

func wrapColorSelection(obj *externglib.Object) *ColorSelection {
	return &ColorSelection{
		Box: Box{
			Container: Container{
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
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalColorSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorSelection(obj), nil
}

// NewColorSelection creates a new GtkColorSelection.
func NewColorSelection() *ColorSelection {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_selection_new()

	var _colorSelection *ColorSelection // out

	_colorSelection = wrapColorSelection(externglib.Take(unsafe.Pointer(_cret)))

	return _colorSelection
}

// CurrentAlpha returns the current alpha value.
func (colorsel *ColorSelection) CurrentAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_get_current_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// CurrentColor sets color to be the current color in the GtkColorSelection
// widget.
//
// Deprecated: Use gtk_color_selection_get_current_rgba() instead.
func (colorsel *ColorSelection) CurrentColor() gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	C.gtk_color_selection_get_current_color(_arg0, &_arg1)

	var _color gdk.Color // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// CurrentRGBA sets rgba to be the current color in the GtkColorSelection
// widget.
func (colorsel *ColorSelection) CurrentRGBA() gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	C.gtk_color_selection_get_current_rgba(_arg0, &_arg1)

	var _rgba gdk.RGBA // out

	_rgba = *(*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rgba
}

// HasOpacityControl determines whether the colorsel has an opacity control.
func (colorsel *ColorSelection) HasOpacityControl() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_get_has_opacity_control(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasPalette determines whether the color selector has a color palette.
func (colorsel *ColorSelection) HasPalette() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_get_has_palette(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PreviousAlpha returns the previous alpha value.
func (colorsel *ColorSelection) PreviousAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_get_previous_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// PreviousColor fills color in with the original color value.
//
// Deprecated: Use gtk_color_selection_get_previous_rgba() instead.
func (colorsel *ColorSelection) PreviousColor() gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	C.gtk_color_selection_get_previous_color(_arg0, &_arg1)

	var _color gdk.Color // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// PreviousRGBA fills rgba in with the original color value.
func (colorsel *ColorSelection) PreviousRGBA() gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	C.gtk_color_selection_get_previous_rgba(_arg0, &_arg1)

	var _rgba gdk.RGBA // out

	_rgba = *(*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rgba
}

// IsAdjusting gets the current state of the colorsel.
func (colorsel *ColorSelection) IsAdjusting() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_is_adjusting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCurrentAlpha sets the current opacity to be alpha.
//
// The first time this is called, it will also set the original opacity to be
// alpha too.
func (colorsel *ColorSelection) SetCurrentAlpha(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_current_alpha(_arg0, _arg1)
}

// SetCurrentColor sets the current color to be color.
//
// The first time this is called, it will also set the original color to be
// color too.
//
// Deprecated: Use gtk_color_selection_set_current_rgba() instead.
func (colorsel *ColorSelection) SetCurrentColor(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_selection_set_current_color(_arg0, _arg1)
}

// SetCurrentRGBA sets the current color to be rgba.
//
// The first time this is called, it will also set the original color to be rgba
// too.
func (colorsel *ColorSelection) SetCurrentRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gtk_color_selection_set_current_rgba(_arg0, _arg1)
}

// SetHasOpacityControl sets the colorsel to use or not use opacity.
func (colorsel *ColorSelection) SetHasOpacityControl(hasOpacity bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	if hasOpacity {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_opacity_control(_arg0, _arg1)
}

// SetHasPalette shows and hides the palette based upon the value of
// has_palette.
func (colorsel *ColorSelection) SetHasPalette(hasPalette bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	if hasPalette {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_palette(_arg0, _arg1)
}

// SetPreviousAlpha sets the “previous” alpha to be alpha.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that alpha change.
func (colorsel *ColorSelection) SetPreviousAlpha(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_previous_alpha(_arg0, _arg1)
}

// SetPreviousColor sets the “previous” color to be color.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_color() will also set this color the first
// time it is called.
//
// Deprecated: Use gtk_color_selection_set_previous_rgba() instead.
func (colorsel *ColorSelection) SetPreviousColor(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_selection_set_previous_color(_arg0, _arg1)
}

// SetPreviousRGBA sets the “previous” color to be rgba.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_rgba() will also set this color the first
// time it is called.
func (colorsel *ColorSelection) SetPreviousRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(colorsel.Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gtk_color_selection_set_previous_rgba(_arg0, _arg1)
}

// ColorSelectionPaletteFromString parses a color palette string; the string is
// a colon-separated list of color names readable by gdk_color_parse().
func ColorSelectionPaletteFromString(str string) ([]gdk.Color, bool) {
	var _arg1 *C.gchar // out
	var _arg2 *C.GdkColor
	var _arg3 C.gint     // in
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))

	_cret = C.gtk_color_selection_palette_from_string(_arg1, &_arg2, &_arg3)

	var _colors []gdk.Color
	var _ok bool // out

	defer C.free(unsafe.Pointer(_arg2))
	_colors = make([]gdk.Color, _arg3)
	copy(_colors, unsafe.Slice((*gdk.Color)(unsafe.Pointer(_arg2)), _arg3))
	if _cret != 0 {
		_ok = true
	}

	return _colors, _ok
}

// ColorSelectionPaletteToString encodes a palette as a string, useful for
// persistent storage.
func ColorSelectionPaletteToString(colors []gdk.Color) string {
	var _arg1 *C.GdkColor
	var _arg2 C.gint
	var _cret *C.gchar // in

	_arg2 = (C.gint)(len(colors))
	if len(colors) > 0 {
		_arg1 = (*C.GdkColor)(unsafe.Pointer(&colors[0]))
	}

	_cret = C.gtk_color_selection_palette_to_string(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
