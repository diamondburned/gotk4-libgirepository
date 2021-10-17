// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_context_get_type()), F: marshalContexter},
	})
}

// Itemize breaks a piece of text into segments with consistent directional
// level and font.
//
// Each byte of text will be contained in exactly one of the items in the
// returned list; the generated list of items will be in logical order (the
// start offsets of the items are ascending).
//
// cached_iter should be an iterator over attrs currently positioned at a range
// before or containing start_index; cached_iter will be advanced to the range
// covering the position just after start_index + length. (i.e. if itemizing in
// a loop, just keep passing in the same cached_iter).
//
// The function takes the following parameters:
//
//    - context: structure holding information that affects the itemization
//    process.
//    - text to itemize. Must be valid UTF-8.
//    - startIndex: first byte in text to process.
//    - length: number of bytes (not characters) to process after start_index.
//    This must be >= 0.
//    - attrs: set of attributes that apply to text.
//    - cachedIter: cached attribute iterator, or NULL.
//
func Itemize(context *Context, text string, startIndex, length int, attrs *AttrList, cachedIter *AttrIterator) []Item {
	var _arg1 *C.PangoContext      // out
	var _arg2 *C.char              // out
	var _arg3 C.int                // out
	var _arg4 C.int                // out
	var _arg5 *C.PangoAttrList     // out
	var _arg6 *C.PangoAttrIterator // out
	var _cret *C.GList             // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(startIndex)
	_arg4 = C.int(length)
	_arg5 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	if cachedIter != nil {
		_arg6 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(cachedIter)))
	}

	_cret = C.pango_itemize(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startIndex)
	runtime.KeepAlive(length)
	runtime.KeepAlive(attrs)
	runtime.KeepAlive(cachedIter)

	var _list []Item // out

	_list = make([]Item, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.PangoItem)(v)
		var dst Item // out
		dst = *(*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(&dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_item_free((*C.PangoItem)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// ItemizeWithBaseDir: like pango_itemize(), but with an explicitly specified
// base direction.
//
// The base direction is used when computing bidirectional levels. (see
// pango.Context.SetBaseDir()). itemize gets the base direction from the
// PangoContext.
//
// The function takes the following parameters:
//
//    - context: structure holding information that affects the itemization
//    process.
//    - baseDir: base direction to use for bidirectional processing.
//    - text to itemize.
//    - startIndex: first byte in text to process.
//    - length: number of bytes (not characters) to process after start_index.
//    This must be >= 0.
//    - attrs: set of attributes that apply to text.
//    - cachedIter: cached attribute iterator, or NULL.
//
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex, length int, attrs *AttrList, cachedIter *AttrIterator) []Item {
	var _arg1 *C.PangoContext      // out
	var _arg2 C.PangoDirection     // out
	var _arg3 *C.char              // out
	var _arg4 C.int                // out
	var _arg5 C.int                // out
	var _arg6 *C.PangoAttrList     // out
	var _arg7 *C.PangoAttrIterator // out
	var _cret *C.GList             // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.PangoDirection(baseDir)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.int(startIndex)
	_arg5 = C.int(length)
	_arg6 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	if cachedIter != nil {
		_arg7 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(cachedIter)))
	}

	_cret = C.pango_itemize_with_base_dir(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(baseDir)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startIndex)
	runtime.KeepAlive(length)
	runtime.KeepAlive(attrs)
	runtime.KeepAlive(cachedIter)

	var _list []Item // out

	_list = make([]Item, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.PangoItem)(v)
		var dst Item // out
		dst = *(*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(&dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_item_free((*C.PangoItem)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// Context: PangoContext stores global information used to control the
// itemization process.
//
// The information stored by `PangoContext includes the fontmap used to look up
// fonts, and default values such as the default language, default gravity, or
// default font.
//
// To obtain a PangoContext, use pango.FontMap.CreateContext().
type Context struct {
	*externglib.Object
}

func wrapContext(obj *externglib.Object) *Context {
	return &Context{
		Object: obj,
	}
}

func marshalContexter(p uintptr) (interface{}, error) {
	return wrapContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewContext creates a new PangoContext initialized to default values.
//
// This function is not particularly useful as it should always be followed by a
// pango.Context.SetFontMap() call, and the function
// pango.FontMap.CreateContext() does these two steps together and hence users
// are recommended to use that.
//
// If you are using Pango as part of a higher-level system, that system may have
// it's own way of create a PangoContext. For instance, the GTK toolkit has,
// among others, gtk_widget_get_pango_context(). Use those instead.
func NewContext() *Context {
	var _cret *C.PangoContext // in

	_cret = C.pango_context_new()

	var _context *Context // out

	_context = wrapContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _context
}

// Changed forces a change in the context, which will cause any PangoLayout
// using this context to re-layout.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if they
// have attached extra data to the context and such data is changed.
func (context *Context) Changed() {
	var _arg0 *C.PangoContext // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_context_changed(_arg0)
	runtime.KeepAlive(context)
}

// BaseDir retrieves the base direction for the context.
//
// See pango.Context.SetBaseDir().
func (context *Context) BaseDir() Direction {
	var _arg0 *C.PangoContext  // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_base_dir(_arg0)
	runtime.KeepAlive(context)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}

// BaseGravity retrieves the base gravity for the context.
//
// See pango.Context.SetBaseGravity().
func (context *Context) BaseGravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_base_gravity(_arg0)
	runtime.KeepAlive(context)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// FontDescription: retrieve the default font description for the context.
func (context *Context) FontDescription() *FontDescription {
	var _arg0 *C.PangoContext         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_font_description(_arg0)
	runtime.KeepAlive(context)

	var _fontDescription *FontDescription // out

	_fontDescription = (*FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _fontDescription
}

// FontMap gets the PangoFontMap used to look up fonts for this context.
func (context *Context) FontMap() FontMapper {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoFontMap // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_font_map(_arg0)
	runtime.KeepAlive(context)

	var _fontMap FontMapper // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontMapper is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(FontMapper)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not pango.FontMapper")
		}
		_fontMap = rv
	}

	return _fontMap
}

// Gravity retrieves the gravity for the context.
//
// This is similar to pango.Context.GetBaseGravity(), except for when the base
// gravity is PANGO_GRAVITY_AUTO for which pango.Gravity.GetForMatrix is used to
// return the gravity from the current context matrix.
func (context *Context) Gravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_gravity(_arg0)
	runtime.KeepAlive(context)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// GravityHint retrieves the gravity hint for the context.
//
// See pango.Context.SetGravityHint() for details.
func (context *Context) GravityHint() GravityHint {
	var _arg0 *C.PangoContext    // out
	var _cret C.PangoGravityHint // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_gravity_hint(_arg0)
	runtime.KeepAlive(context)

	var _gravityHint GravityHint // out

	_gravityHint = GravityHint(_cret)

	return _gravityHint
}

// Language retrieves the global language tag for the context.
func (context *Context) Language() *Language {
	var _arg0 *C.PangoContext  // out
	var _cret *C.PangoLanguage // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_language(_arg0)
	runtime.KeepAlive(context)

	var _language *Language // out

	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_language)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _language
}

// Matrix gets the transformation matrix that will be applied when rendering
// with this context.
//
// See pango.Context.SetMatrix().
func (context *Context) Matrix() *Matrix {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoMatrix  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_matrix(_arg0)
	runtime.KeepAlive(context)

	var _matrix *Matrix // out

	if _cret != nil {
		_matrix = (*Matrix)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _matrix
}

// Metrics: get overall metric information for a particular font description.
//
// Since the metrics may be substantially different for different scripts, a
// language tag can be provided to indicate that the metrics should be retrieved
// that correspond to the script(s) used by that language.
//
// The PangoFontDescription is interpreted in the same way as by itemize, and
// the family name may be a comma separated list of names. If characters from
// multiple of these families would be used to render the string, then the
// returned fonts would be a composite of the metrics for the fonts loaded for
// the individual families.
//
// The function takes the following parameters:
//
//    - desc: PangoFontDescription structure. NULL means that the font
//    description from the context will be used.
//    - language tag used to determine which script to get the metrics for.
//    NULL means that the language tag from the context will be used. If no
//    language tag is set on the context, metrics for the default language (as
//    determined by pango.Language.GetDefault will be returned.
//
func (context *Context) Metrics(desc *FontDescription, language *Language) *FontMetrics {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontMetrics     // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if desc != nil {
		_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	}
	if language != nil {
		_arg2 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	}

	_cret = C.pango_context_get_metrics(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)
	runtime.KeepAlive(language)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fontMetrics)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_font_metrics_unref((*C.PangoFontMetrics)(intern.C))
		},
	)

	return _fontMetrics
}

// RoundGlyphPositions returns whether font rendering with this context should
// round glyph positions and widths.
func (context *Context) RoundGlyphPositions() bool {
	var _arg0 *C.PangoContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_round_glyph_positions(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Serial returns the current serial number of context.
//
// The serial number is initialized to an small number larger than zero when a
// new context is created and is increased whenever the context is changed using
// any of the setter functions, or the PangoFontMap it uses to find fonts has
// changed. The serial may wrap, but will never have the value 0. Since it can
// wrap, never compare it with "less than", always use "not equals".
//
// This can be used to automatically detect changes to a PangoContext, and is
// only useful when implementing objects that need update when their
// PangoContext changes, like PangoLayout.
func (context *Context) Serial() uint {
	var _arg0 *C.PangoContext // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_serial(_arg0)
	runtime.KeepAlive(context)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ListFamilies: list all families for a context.
func (context *Context) ListFamilies() []FontFamilier {
	var _arg0 *C.PangoContext     // out
	var _arg1 **C.PangoFontFamily // in
	var _arg2 C.int               // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_context_list_families(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _families []FontFamilier // out

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice(_arg1, _arg2)
		_families = make([]FontFamilier, _arg2)
		for i := 0; i < int(_arg2); i++ {
			{
				objptr := unsafe.Pointer(src[i])
				if objptr == nil {
					panic("object of type pango.FontFamilier is nil")
				}

				object := externglib.Take(objptr)
				rv, ok := (externglib.CastObject(object)).(FontFamilier)
				if !ok {
					panic("object of type " + object.TypeFromInstance().String() + " is not pango.FontFamilier")
				}
				_families[i] = rv
			}
		}
	}

	return _families
}

// LoadFont loads the font in one of the fontmaps in the context that is the
// closest match for desc.
//
// The function takes the following parameters:
//
//    - desc: PangoFontDescription describing the font to load.
//
func (context *Context) LoadFont(desc *FontDescription) Fonter {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	_cret = C.pango_context_load_font(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)

	var _font Fonter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Fonter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not pango.Fonter")
			}
			_font = rv
		}
	}

	return _font
}

// LoadFontset: load a set of fonts in the context that can be used to render a
// font matching desc.
//
// The function takes the following parameters:
//
//    - desc: PangoFontDescription describing the fonts to load.
//    - language: PangoLanguage the fonts will be used for.
//
func (context *Context) LoadFontset(desc *FontDescription, language *Language) Fontsetter {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	_arg2 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_context_load_fontset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)
	runtime.KeepAlive(language)

	var _fontset Fontsetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Fontsetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not pango.Fontsetter")
			}
			_fontset = rv
		}
	}

	return _fontset
}

// SetBaseDir sets the base direction for the context.
//
// The base direction is used in applying the Unicode bidirectional algorithm;
// if the direction is PANGO_DIRECTION_LTR or PANGO_DIRECTION_RTL, then the
// value will be used as the paragraph direction in the Unicode bidirectional
// algorithm. A value of PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL is
// used only for paragraphs that do not contain any strong characters
// themselves.
//
// The function takes the following parameters:
//
//    - direction: new base direction.
//
func (context *Context) SetBaseDir(direction Direction) {
	var _arg0 *C.PangoContext  // out
	var _arg1 C.PangoDirection // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoDirection(direction)

	C.pango_context_set_base_dir(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(direction)
}

// SetBaseGravity sets the base gravity for the context.
//
// The base gravity is used in laying vertical text out.
//
// The function takes the following parameters:
//
//    - gravity: new base gravity.
//
func (context *Context) SetBaseGravity(gravity Gravity) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.PangoGravity  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoGravity(gravity)

	C.pango_context_set_base_gravity(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(gravity)
}

// SetFontDescription: set the default font description for the context.
//
// The function takes the following parameters:
//
//    - desc: new pango font description.
//
func (context *Context) SetFontDescription(desc *FontDescription) {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	C.pango_context_set_font_description(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)
}

// SetFontMap sets the font map to be searched when fonts are looked-up in this
// context.
//
// This is only for internal use by Pango backends, a PangoContext obtained via
// one of the recommended methods should already have a suitable font map.
//
// The function takes the following parameters:
//
//    - fontMap: PangoFontMap to set.
//
func (context *Context) SetFontMap(fontMap FontMapper) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoFontMap // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontMap.Native()))

	C.pango_context_set_font_map(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(fontMap)
}

// SetGravityHint sets the gravity hint for the context.
//
// The gravity hint is used in laying vertical text out, and is only relevant if
// gravity of the context as returned by pango.Context.GetGravity() is set to
// PANGO_GRAVITY_EAST or PANGO_GRAVITY_WEST.
//
// The function takes the following parameters:
//
//    - hint: new gravity hint.
//
func (context *Context) SetGravityHint(hint GravityHint) {
	var _arg0 *C.PangoContext    // out
	var _arg1 C.PangoGravityHint // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoGravityHint(hint)

	C.pango_context_set_gravity_hint(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(hint)
}

// SetLanguage sets the global language tag for the context.
//
// The default language for the locale of the running process can be found using
// pango.Language.GetDefault.
//
// The function takes the following parameters:
//
//    - language: new language tag.
//
func (context *Context) SetLanguage(language *Language) {
	var _arg0 *C.PangoContext  // out
	var _arg1 *C.PangoLanguage // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	C.pango_context_set_language(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(language)
}

// SetMatrix sets the transformation matrix that will be applied when rendering
// with this context.
//
// Note that reported metrics are in the user space coordinates before the
// application of the matrix, not device-space coordinates after the application
// of the matrix. So, they don't scale with the matrix, though they may change
// slightly for different matrices, depending on how the text is fit to the
// pixel grid.
//
// The function takes the following parameters:
//
//    - matrix: PangoMatrix, or NULL to unset any existing matrix. (No matrix
//    set is the same as setting the identity matrix.).
//
func (context *Context) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoMatrix  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if matrix != nil {
		_arg1 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	C.pango_context_set_matrix(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(matrix)
}

// SetRoundGlyphPositions sets whether font rendering with this context should
// round glyph positions and widths to integral positions, in device units.
//
// This is useful when the renderer can't handle subpixel positioning of glyphs.
//
// The default value is to round glyph positions, to remain compatible with
// previous Pango behavior.
//
// The function takes the following parameters:
//
//    - roundPositions: whether to round glyph positions.
//
func (context *Context) SetRoundGlyphPositions(roundPositions bool) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if roundPositions {
		_arg1 = C.TRUE
	}

	C.pango_context_set_round_glyph_positions(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(roundPositions)
}
