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
// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_font_map_get_type()), F: marshalFontMapper},
	})
}

// FontMapOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontMapOverrider interface {
	// Changed forces a change in the context, which will cause any PangoContext
	// using this fontmap to change.
	//
	// This function is only useful when implementing a new backend for Pango,
	// something applications won't do. Backends should call this function if
	// they have attached extra data to the context and such data is changed.
	Changed()
	// Family gets a font family by name.
	Family(name string) FontFamilier
	// Serial returns the current serial number of fontmap.
	//
	// The serial number is initialized to an small number larger than zero when
	// a new fontmap is created and is increased whenever the fontmap is
	// changed. It may wrap, but will never have the value 0. Since it can wrap,
	// never compare it with "less than", always use "not equals".
	//
	// The fontmap can only be changed using backend-specific API, like changing
	// fontmap resolution.
	//
	// This can be used to automatically detect changes to a PangoFontMap, like
	// in PangoContext.
	Serial() uint
	// ListFamilies: list all families for a fontmap.
	ListFamilies() []FontFamilier
	// LoadFont: load the font in the fontmap that is the closest match for
	// desc.
	LoadFont(context *Context, desc *FontDescription) Fonter
	// LoadFontset: load a set of fonts in the fontmap that can be used to
	// render a font matching desc.
	LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter
}

// FontMap: PangoFontMap represents the set of fonts available for a particular
// rendering system.
//
// This is a virtual object with implementations being specific to particular
// rendering systems.
type FontMap struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*FontMap)(nil)
)

// FontMapper describes types inherited from class FontMap.

// To get the original type, the caller must assert this to an interface or
// another type.
type FontMapper interface {
	externglib.Objector
	baseFontMap() *FontMap
}

var _ FontMapper = (*FontMap)(nil)

func wrapFontMap(obj *externglib.Object) *FontMap {
	return &FontMap{
		Object: obj,
	}
}

func marshalFontMapper(p uintptr) (interface{}, error) {
	return wrapFontMap(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Changed forces a change in the context, which will cause any PangoContext
// using this fontmap to change.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if they
// have attached extra data to the context and such data is changed.
func (fontmap *FontMap) Changed() {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	C.pango_font_map_changed(_arg0)
	runtime.KeepAlive(fontmap)
}

// CreateContext creates a PangoContext connected to fontmap.
//
// This is equivalent to pango.Context.New followed by
// pango.Context.SetFontMap().
//
// If you are using Pango as part of a higher-level system, that system may have
// it's own way of create a PangoContext. For instance, the GTK toolkit has,
// among others, gtk_widget_get_pango_context(). Use those instead.
func (fontmap *FontMap) CreateContext() *Context {
	var _arg0 *C.PangoFontMap // out
	var _cret *C.PangoContext // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	_cret = C.pango_font_map_create_context(_arg0)
	runtime.KeepAlive(fontmap)

	var _context *Context // out

	_context = wrapContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _context
}

// Family gets a font family by name.
//
// The function takes the following parameters:
//
//    - name: family name.
//
func (fontmap *FontMap) Family(name string) FontFamilier {
	var _arg0 *C.PangoFontMap    // out
	var _arg1 *C.char            // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_font_map_get_family(_arg0, _arg1)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(name)

	var _fontFamily FontFamilier // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(FontFamilier)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not pango.FontFamilier")
		}
		_fontFamily = rv
	}

	return _fontFamily
}

// Serial returns the current serial number of fontmap.
//
// The serial number is initialized to an small number larger than zero when a
// new fontmap is created and is increased whenever the fontmap is changed. It
// may wrap, but will never have the value 0. Since it can wrap, never compare
// it with "less than", always use "not equals".
//
// The fontmap can only be changed using backend-specific API, like changing
// fontmap resolution.
//
// This can be used to automatically detect changes to a PangoFontMap, like in
// PangoContext.
func (fontmap *FontMap) Serial() uint {
	var _arg0 *C.PangoFontMap // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	_cret = C.pango_font_map_get_serial(_arg0)
	runtime.KeepAlive(fontmap)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ListFamilies: list all families for a fontmap.
func (fontmap *FontMap) ListFamilies() []FontFamilier {
	var _arg0 *C.PangoFontMap     // out
	var _arg1 **C.PangoFontFamily // in
	var _arg2 C.int               // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	C.pango_font_map_list_families(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(fontmap)

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

// LoadFont: load the font in the fontmap that is the closest match for desc.
//
// The function takes the following parameters:
//
//    - context: PangoContext the font will be used with.
//    - desc: PangoFontDescription describing the font to load.
//
func (fontmap *FontMap) LoadFont(context *Context, desc *FontDescription) Fonter {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	_cret = C.pango_font_map_load_font(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fontmap)
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

// LoadFontset: load a set of fonts in the fontmap that can be used to render a
// font matching desc.
//
// The function takes the following parameters:
//
//    - context: PangoContext the font will be used with.
//    - desc: PangoFontDescription describing the font to load.
//    - language: PangoLanguage the fonts will be used for.
//
func (fontmap *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _arg3 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	_arg3 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_font_map_load_fontset(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fontmap)
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

func (fontmap *FontMap) baseFontMap() *FontMap {
	return fontmap
}

// BaseFontMap returns the underlying base object.
func BaseFontMap(obj FontMapper) *FontMap {
	return obj.baseFontMap()
}
