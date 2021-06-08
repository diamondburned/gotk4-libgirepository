// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoxft pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangoxft.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_xft_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_xft_font_map_get_type()), F: marshalFontMap},
	})
}

// SetDefaultSubstitute sets a function that will be called to do final
// configuration substitution on a Pattern before it is used to load the font.
// This function can be used to do things like set hinting and antialiasing
// options.
func SetDefaultSubstitute() {
	C.pango_xft_set_default_substitute(arg1, arg2, arg3, arg4, arg5)
}

// Font is an implementation of FcFont using the Xft library for rendering. It
// is used in conjunction with XftFontMap.
type Font interface {
	pangofc.Font

	// Glyph gets the glyph index for a given Unicode character for @font. If
	// you only want to determine whether the font has the glyph, use
	// pango_xft_font_has_char().
	//
	// Use pango_fc_font_get_glyph() instead.
	Glyph(wc uint32) uint
	// HasChar determines whether @font has a glyph for the codepoint @wc.
	//
	// Use pango_fc_font_has_char() instead.
	HasChar(wc uint32) bool
	// UnlockFace releases a font previously obtained with
	// pango_xft_font_lock_face().
	//
	// Use pango_fc_font_unlock_face() instead.
	UnlockFace()
}

// font implements the Font interface.
type font struct {
	pangofc.Font
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return Font{
		pangofc.Font: pangofc.WrapFont(obj),
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// Glyph gets the glyph index for a given Unicode character for @font. If
// you only want to determine whether the font has the glyph, use
// pango_xft_font_has_char().
//
// Use pango_fc_font_get_glyph() instead.
func (f font) Glyph(wc uint32) uint {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	arg1 = C.gunichar(wc)

	var cret C.guint
	var goret uint

	cret = C.pango_xft_font_get_glyph(arg0, arg1)

	goret = uint(cret)

	return goret
}

// HasChar determines whether @font has a glyph for the codepoint @wc.
//
// Use pango_fc_font_has_char() instead.
func (f font) HasChar(wc uint32) bool {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	arg1 = C.gunichar(wc)

	var cret C.gboolean
	var goret bool

	cret = C.pango_xft_font_has_char(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// UnlockFace releases a font previously obtained with
// pango_xft_font_lock_face().
//
// Use pango_fc_font_unlock_face() instead.
func (f font) UnlockFace() {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_xft_font_unlock_face(arg0)
}

// FontMap is an implementation of FcFontMap suitable for the Xft library as the
// renderer. It is used in to create fonts of type XftFont.
type FontMap interface {
	pangofc.FontMap
}

// fontMap implements the FontMap interface.
type fontMap struct {
	pangofc.FontMap
}

var _ FontMap = (*fontMap)(nil)

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return FontMap{
		pangofc.FontMap: pangofc.WrapFontMap(obj),
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}
