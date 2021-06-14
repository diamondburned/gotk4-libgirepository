// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangocairo
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangocairo.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_cairo_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_cairo_font_map_get_type()), F: marshalFontMap},
	})
}

// ContextGetResolution gets the resolution for the context. See
// [func@PangoCairo.context_set_resolution]
func ContextGetResolution(context pango.Context) float64 {
	var _arg1 *C.PangoContext // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	var _cret C.double // in

	_cret = C.pango_cairo_context_get_resolution(_arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// ContextSetFontOptions sets the font options used when rendering text with
// this context.
//
// These options override any options that [func@update_context] derives from
// the target surface.
func ContextSetFontOptions(context pango.Context, options *cairo.FontOptions) {
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.cairo_font_options_t // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_font_options_t)(unsafe.Pointer(options.Native()))

	C.pango_cairo_context_set_font_options(_arg1, _arg2)
}

// ContextSetResolution sets the resolution for the context.
//
// This is a scale factor between points specified in a `PangoFontDescription`
// and Cairo units. The default value is 96, meaning that a 10 point font will
// be 13 units high. (10 * 96. / 72. = 13.3).
func ContextSetResolution(context pango.Context, dpi float64) {
	var _arg1 *C.PangoContext // out
	var _arg2 C.double        // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(dpi)

	C.pango_cairo_context_set_resolution(_arg1, _arg2)
}

// ErrorUnderlinePath: add a squiggly line to the current path in the specified
// cairo context that approximately covers the given rectangle in the style of
// an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_error_underline_path(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// GlyphStringPath adds the glyphs in @glyphs to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be at the
// current point of the cairo context.
func GlyphStringPath(cr *cairo.Context, font pango.Font, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg3 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_cairo_glyph_string_path(_arg1, _arg2, _arg3)
}

// LayoutPath adds the text in a `PangoLayout` to the current path in the
// specified cairo context.
//
// The top-left corner of the `PangoLayout` will be at the current point of the
// cairo context.
func LayoutPath(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_layout_path(_arg1, _arg2)
}

// ShowErrorUnderline: draw a squiggly line in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used to
// indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_show_error_underline(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// ShowGlyphItem draws the glyphs in @glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// [func@show_glyph_string].
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that @text is the start of the text for layout, which is then indexed by
// `glyph_item->item->offset`.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.char           // out
	var _arg3 *C.PangoGlyphItem // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem.Native()))

	C.pango_cairo_show_glyph_item(_arg1, _arg2, _arg3)
}

// ShowGlyphString draws the glyphs in @glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
func ShowGlyphString(cr *cairo.Context, font pango.Font, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg3 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_cairo_show_glyph_string(_arg1, _arg2, _arg3)
}

// ShowLayout draws a `PangoLayout` in the specified cairo context.
//
// The top-left corner of the `PangoLayout` will be drawn at the current point
// of the cairo context.
func ShowLayout(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_show_layout(_arg1, _arg2)
}

// UpdateContext updates a `PangoContext` previously created for use with Cairo
// to match the current transformation and target surface of a Cairo context.
//
// If any layouts have been created for the context, it's necessary to call
// [method@Pango.Layout.context_changed] on those layouts.
func UpdateContext(cr *cairo.Context, context pango.Context) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.PangoContext // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_cairo_update_context(_arg1, _arg2)
}

// UpdateLayout updates the private `PangoContext` of a `PangoLayout` created
// with [func@create_layout] to match the current transformation and target
// surface of a Cairo context.
func UpdateLayout(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_update_layout(_arg1, _arg2)
}

// Font: `PangoCairoFont` is an interface exported by fonts for use with Cairo.
//
// The actual type of the font will depend on the particular font technology
// Cairo was compiled to use.
type Font interface {
	pango.Font
}

// font implements the Font interface.
type font struct {
	pango.Font
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to a type that implements interface
// Font. It is primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return Font{
		pango.Font: pango.WrapFont(obj),
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// FontMap: `PangoCairoFontMap` is an interface exported by font maps for use
// with Cairo.
//
// The actual type of the font map will depend on the particular font technology
// Cairo was compiled to use.
type FontMap interface {
	pango.FontMap

	// Resolution gets the resolution for the fontmap.
	//
	// See [method@PangoCairo.FontMap.set_resolution].
	Resolution() float64
	// SetDefault sets a default `PangoCairoFontMap` to use with Cairo.
	//
	// This can be used to change the Cairo font backend that the default
	// fontmap uses for example. The old default font map is unreffed and the
	// new font map referenced.
	//
	// Note that since Pango 1.32.6, the default fontmap is per-thread. This
	// function only changes the default fontmap for the current thread. Default
	// fontmaps of existing threads are not changed. Default fontmaps of any new
	// threads will still be created using [type_func@PangoCairo.FontMap.new].
	//
	// A value of nil for @fontmap will cause the current default font map to be
	// released and a new default font map to be created on demand, using
	// [type_func@PangoCairo.FontMap.new].
	SetDefault()
	// SetResolution sets the resolution for the fontmap.
	//
	// This is a scale factor between points specified in a
	// `PangoFontDescription` and Cairo units. The default value is 96, meaning
	// that a 10 point font will be 13 units high. (10 * 96. / 72. = 13.3).
	SetResolution(dpi float64)
}

// fontMap implements the FontMap interface.
type fontMap struct {
	pango.FontMap
}

var _ FontMap = (*fontMap)(nil)

// WrapFontMap wraps a GObject to a type that implements interface
// FontMap. It is primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return FontMap{
		pango.FontMap: pango.WrapFontMap(obj),
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// Resolution gets the resolution for the fontmap.
//
// See [method@PangoCairo.FontMap.set_resolution].
func (f fontMap) Resolution() float64 {
	var _arg0 *C.PangoCairoFontMap // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(f.Native()))

	var _cret C.double // in

	_cret = C.pango_cairo_font_map_get_resolution(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// SetDefault sets a default `PangoCairoFontMap` to use with Cairo.
//
// This can be used to change the Cairo font backend that the default
// fontmap uses for example. The old default font map is unreffed and the
// new font map referenced.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. This
// function only changes the default fontmap for the current thread. Default
// fontmaps of existing threads are not changed. Default fontmaps of any new
// threads will still be created using [type_func@PangoCairo.FontMap.new].
//
// A value of nil for @fontmap will cause the current default font map to be
// released and a new default font map to be created on demand, using
// [type_func@PangoCairo.FontMap.new].
func (f fontMap) SetDefault() {
	var _arg0 *C.PangoCairoFontMap // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(f.Native()))

	C.pango_cairo_font_map_set_default(_arg0)
}

// SetResolution sets the resolution for the fontmap.
//
// This is a scale factor between points specified in a
// `PangoFontDescription` and Cairo units. The default value is 96, meaning
// that a 10 point font will be 13 units high. (10 * 96. / 72. = 13.3).
func (f fontMap) SetResolution(dpi float64) {
	var _arg0 *C.PangoCairoFontMap // out
	var _arg1 C.double             // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = C.double(dpi)

	C.pango_cairo_font_map_set_resolution(_arg0, _arg1)
}
