// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_context_get_type()), F: marshalContext},
	})
}

// Context: `PangoContext` stores global information used to control the
// itemization process.
//
// The information stored by `PangoContext includes the fontmap used to look up
// fonts, and default values such as the default language, default gravity, or
// default font.
//
// To obtain a `PangoContext`, use [method@Pango.FontMap.create_context].
type Context interface {
	gextras.Objector

	// Changed forces a change in the context, which will cause any
	// `PangoLayout` using this context to re-layout.
	//
	// This function is only useful when implementing a new backend for Pango,
	// something applications won't do. Backends should call this function if
	// they have attached extra data to the context and such data is changed.
	Changed()
	// BaseDir retrieves the base direction for the context.
	//
	// See [method@Pango.Context.set_base_dir].
	BaseDir() Direction
	// BaseGravity retrieves the base gravity for the context.
	//
	// See [method@Pango.Context.set_base_gravity].
	BaseGravity() Gravity
	// FontDescription: retrieve the default font description for the context.
	FontDescription() *FontDescription
	// FontMap gets the `PangoFontMap` used to look up fonts for this context.
	FontMap() FontMap
	// Gravity retrieves the gravity for the context.
	//
	// This is similar to [method@Pango.Context.get_base_gravity], except for
	// when the base gravity is PANGO_GRAVITY_AUTO for which
	// [type_func@Pango.Gravity.get_for_matrix] is used to return the gravity
	// from the current context matrix.
	Gravity() Gravity
	// GravityHint retrieves the gravity hint for the context.
	//
	// See [method@Pango.Context.set_gravity_hint] for details.
	GravityHint() GravityHint
	// Language retrieves the global language tag for the context.
	Language() *Language
	// Matrix gets the transformation matrix that will be applied when rendering
	// with this context.
	//
	// See [method@Pango.Context.set_matrix].
	Matrix() *Matrix
	// Metrics: get overall metric information for a particular font
	// description.
	//
	// Since the metrics may be substantially different for different scripts, a
	// language tag can be provided to indicate that the metrics should be
	// retrieved that correspond to the script(s) used by that language.
	//
	// The `PangoFontDescription` is interpreted in the same way as by
	// [func@itemize], and the family name may be a comma separated list of
	// names. If characters from multiple of these families would be used to
	// render the string, then the returned fonts would be a composite of the
	// metrics for the fonts loaded for the individual families.
	Metrics(desc *FontDescription, language *Language) *FontMetrics
	// RoundGlyphPositions returns whether font rendering with this context
	// should round glyph positions and widths.
	RoundGlyphPositions() bool
	// Serial returns the current serial number of @context.
	//
	// The serial number is initialized to an small number larger than zero when
	// a new context is created and is increased whenever the context is changed
	// using any of the setter functions, or the `PangoFontMap` it uses to find
	// fonts has changed. The serial may wrap, but will never have the value 0.
	// Since it can wrap, never compare it with "less than", always use "not
	// equals".
	//
	// This can be used to automatically detect changes to a `PangoContext`, and
	// is only useful when implementing objects that need update when their
	// `PangoContext` changes, like `PangoLayout`.
	Serial() uint
	// ListFamilies: list all families for a context.
	ListFamilies() []FontFamily
	// LoadFont loads the font in one of the fontmaps in the context that is the
	// closest match for @desc.
	LoadFont(desc *FontDescription) Font
	// LoadFontset: load a set of fonts in the context that can be used to
	// render a font matching @desc.
	LoadFontset(desc *FontDescription, language *Language) Fontset
	// SetBaseDir sets the base direction for the context.
	//
	// The base direction is used in applying the Unicode bidirectional
	// algorithm; if the @direction is PANGO_DIRECTION_LTR or
	// PANGO_DIRECTION_RTL, then the value will be used as the paragraph
	// direction in the Unicode bidirectional algorithm. A value of
	// PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL is used only for
	// paragraphs that do not contain any strong characters themselves.
	SetBaseDir(direction Direction)
	// SetBaseGravity sets the base gravity for the context.
	//
	// The base gravity is used in laying vertical text out.
	SetBaseGravity(gravity Gravity)
	// SetFontDescription: set the default font description for the context
	SetFontDescription(desc *FontDescription)
	// SetFontMap sets the font map to be searched when fonts are looked-up in
	// this context.
	//
	// This is only for internal use by Pango backends, a `PangoContext`
	// obtained via one of the recommended methods should already have a
	// suitable font map.
	SetFontMap(fontMap FontMap)
	// SetGravityHint sets the gravity hint for the context.
	//
	// The gravity hint is used in laying vertical text out, and is only
	// relevant if gravity of the context as returned by
	// [method@Pango.Context.get_gravity] is set to PANGO_GRAVITY_EAST or
	// PANGO_GRAVITY_WEST.
	SetGravityHint(hint GravityHint)
	// SetLanguage sets the global language tag for the context.
	//
	// The default language for the locale of the running process can be found
	// using [type_func@Pango.Language.get_default].
	SetLanguage(language *Language)
	// SetMatrix sets the transformation matrix that will be applied when
	// rendering with this context.
	//
	// Note that reported metrics are in the user space coordinates before the
	// application of the matrix, not device-space coordinates after the
	// application of the matrix. So, they don't scale with the matrix, though
	// they may change slightly for different matrices, depending on how the
	// text is fit to the pixel grid.
	SetMatrix(matrix *Matrix)
	// SetRoundGlyphPositions sets whether font rendering with this context
	// should round glyph positions and widths to integral positions, in device
	// units.
	//
	// This is useful when the renderer can't handle subpixel positioning of
	// glyphs.
	//
	// The default value is to round glyph positions, to remain compatible with
	// previous Pango behavior.
	SetRoundGlyphPositions(roundPositions bool)
}

// context implements the Context interface.
type context struct {
	*externglib.Object
}

var _ Context = (*context)(nil)

// WrapContext wraps a GObject to a type that implements
// interface Context. It is primarily used internally.
func WrapContext(obj *externglib.Object) Context {
	return context{obj}
}

func marshalContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContext(obj), nil
}

// NewContext creates a new `PangoContext` initialized to default values.
//
// This function is not particularly useful as it should always be followed by a
// [method@Pango.Context.set_font_map] call, and the function
// [method@Pango.FontMap.create_context] does these two steps together and hence
// users are recommended to use that.
//
// If you are using Pango as part of a higher-level system, that system may have
// it's own way of create a `PangoContext`. For instance, the GTK toolkit has,
// among others, `gtk_widget_get_pango_context()`. Use those instead.
func NewContext() Context {
	var _cret *C.PangoContext // in

	_cret = C.pango_context_new()

	var _context Context // out

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Context)

	return _context
}

func (c context) Changed() {
	var _arg0 *C.PangoContext // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	C.pango_context_changed(_arg0)
}

func (c context) BaseDir() Direction {
	var _arg0 *C.PangoContext  // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_base_dir(_arg0)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}

func (c context) BaseGravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_base_gravity(_arg0)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

func (c context) FontDescription() *FontDescription {
	var _arg0 *C.PangoContext         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_font_description(_arg0)

	var _fontDescription *FontDescription // out

	_fontDescription = (*FontDescription)(unsafe.Pointer(_cret))

	return _fontDescription
}

func (c context) FontMap() FontMap {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoFontMap // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_font_map(_arg0)

	var _fontMap FontMap // out

	_fontMap = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontMap)

	return _fontMap
}

func (c context) Gravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_gravity(_arg0)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

func (c context) GravityHint() GravityHint {
	var _arg0 *C.PangoContext    // out
	var _cret C.PangoGravityHint // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_gravity_hint(_arg0)

	var _gravityHint GravityHint // out

	_gravityHint = GravityHint(_cret)

	return _gravityHint
}

func (c context) Language() *Language {
	var _arg0 *C.PangoContext  // out
	var _cret *C.PangoLanguage // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_language(_arg0)

	var _language *Language // out

	_language = (*Language)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_language, func(v *Language) {
		C.free(unsafe.Pointer(v))
	})

	return _language
}

func (c context) Matrix() *Matrix {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoMatrix  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_matrix(_arg0)

	var _matrix *Matrix // out

	_matrix = (*Matrix)(unsafe.Pointer(_cret))

	return _matrix
}

func (c context) Metrics(desc *FontDescription, language *Language) *FontMetrics {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontMetrics     // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc))
	_arg2 = (*C.PangoLanguage)(unsafe.Pointer(language))

	_cret = C.pango_context_get_metrics(_arg0, _arg1, _arg2)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(unsafe.Pointer(_cret))
	C.pango_font_metrics_ref(_cret)
	runtime.SetFinalizer(_fontMetrics, func(v *FontMetrics) {
		C.pango_font_metrics_unref((*C.PangoFontMetrics)(unsafe.Pointer(v)))
	})

	return _fontMetrics
}

func (c context) RoundGlyphPositions() bool {
	var _arg0 *C.PangoContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_round_glyph_positions(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c context) Serial() uint {
	var _arg0 *C.PangoContext // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	_cret = C.pango_context_get_serial(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (c context) ListFamilies() []FontFamily {
	var _arg0 *C.PangoContext // out
	var _arg1 **C.PangoFontFamily
	var _arg2 C.int // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	C.pango_context_list_families(_arg0, &_arg1, &_arg2)

	var _families []FontFamily

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice(_arg1, _arg2)
		_families = make([]FontFamily, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_families[i] = gextras.CastObject(externglib.Take(unsafe.Pointer(src[i]))).(FontFamily)
		}
	}

	return _families
}

func (c context) LoadFont(desc *FontDescription) Font {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc))

	_cret = C.pango_context_load_font(_arg0, _arg1)

	var _font Font // out

	_font = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Font)

	return _font
}

func (c context) LoadFontset(desc *FontDescription, language *Language) Fontset {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc))
	_arg2 = (*C.PangoLanguage)(unsafe.Pointer(language))

	_cret = C.pango_context_load_fontset(_arg0, _arg1, _arg2)

	var _fontset Fontset // out

	_fontset = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Fontset)

	return _fontset
}

func (c context) SetBaseDir(direction Direction) {
	var _arg0 *C.PangoContext  // out
	var _arg1 C.PangoDirection // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.PangoDirection(direction)

	C.pango_context_set_base_dir(_arg0, _arg1)
}

func (c context) SetBaseGravity(gravity Gravity) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.PangoGravity  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.PangoGravity(gravity)

	C.pango_context_set_base_gravity(_arg0, _arg1)
}

func (c context) SetFontDescription(desc *FontDescription) {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc))

	C.pango_context_set_font_description(_arg0, _arg1)
}

func (c context) SetFontMap(fontMap FontMap) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoFontMap // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontMap.Native()))

	C.pango_context_set_font_map(_arg0, _arg1)
}

func (c context) SetGravityHint(hint GravityHint) {
	var _arg0 *C.PangoContext    // out
	var _arg1 C.PangoGravityHint // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.PangoGravityHint(hint)

	C.pango_context_set_gravity_hint(_arg0, _arg1)
}

func (c context) SetLanguage(language *Language) {
	var _arg0 *C.PangoContext  // out
	var _arg1 *C.PangoLanguage // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoLanguage)(unsafe.Pointer(language))

	C.pango_context_set_language(_arg0, _arg1)
}

func (c context) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoMatrix  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	C.pango_context_set_matrix(_arg0, _arg1)
}

func (c context) SetRoundGlyphPositions(roundPositions bool) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	if roundPositions {
		_arg1 = C.TRUE
	}

	C.pango_context_set_round_glyph_positions(_arg0, _arg1)
}
