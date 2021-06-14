// Code generated by girgen. DO NOT EDIT.

package pangofc

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangofc
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangofc-fontmap.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fc_font_map_get_type()), F: marshalFontMap},
	})
}

// FontMap: `PangoFcFontMap` is a base class for font map implementations using
// the Fontconfig and FreeType libraries.
//
// It is used in the Xft and FreeType backends shipped with Pango, but can also
// be used when creating new backends. Any backend deriving from this base class
// will take advantage of the wide range of shapers implemented using FreeType
// that come with Pango.
type FontMap interface {
	pango.FontMap

	// CacheClear: clear all cached information and fontsets for this font map.
	//
	// This should be called whenever there is a change in the output of the
	// default_substitute() virtual function of the font map, or if fontconfig
	// has been reinitialized to new configuration.
	CacheClear()
	// ConfigChanged informs font map that the fontconfig configuration (i.e.,
	// FcConfig object) used by this font map has changed.
	//
	// This currently calls [method@PangoFc.FontMap.cache_clear] which ensures
	// that list of fonts, etc will be regenerated using the updated
	// configuration.
	ConfigChanged()
	// Shutdown clears all cached information for the fontmap and marks all
	// fonts open for the fontmap as dead.
	//
	// See the shutdown() virtual function of `PangoFcFont`.
	//
	// This function might be used by a backend when the underlying windowing
	// system for the font map exits. This function is only intended to be
	// called only for backend implementations deriving from `PangoFcFontMap`.
	Shutdown()
	// SubstituteChanged: call this function any time the results of the default
	// substitution function set with
	// [method@PangoFc.FontMap.set_default_substitute] change.
	//
	// That is, if your substitution function will return different results for
	// the same input pattern, you must call this function.
	SubstituteChanged()
}

// fontMap implements the FontMap class.
type fontMap struct {
	pango.FontMap
}

var _ FontMap = (*fontMap)(nil)

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{
		pango.FontMap: pango.WrapFontMap(obj),
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// CacheClear: clear all cached information and fontsets for this font map.
//
// This should be called whenever there is a change in the output of the
// default_substitute() virtual function of the font map, or if fontconfig
// has been reinitialized to new configuration.
func (f fontMap) CacheClear() {
	var _arg0 *C.PangoFcFontMap // out

	_arg0 = (*C.PangoFcFontMap)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_map_cache_clear(_arg0)
}

// ConfigChanged informs font map that the fontconfig configuration (i.e.,
// FcConfig object) used by this font map has changed.
//
// This currently calls [method@PangoFc.FontMap.cache_clear] which ensures
// that list of fonts, etc will be regenerated using the updated
// configuration.
func (f fontMap) ConfigChanged() {
	var _arg0 *C.PangoFcFontMap // out

	_arg0 = (*C.PangoFcFontMap)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_map_config_changed(_arg0)
}

// Shutdown clears all cached information for the fontmap and marks all
// fonts open for the fontmap as dead.
//
// See the shutdown() virtual function of `PangoFcFont`.
//
// This function might be used by a backend when the underlying windowing
// system for the font map exits. This function is only intended to be
// called only for backend implementations deriving from `PangoFcFontMap`.
func (f fontMap) Shutdown() {
	var _arg0 *C.PangoFcFontMap // out

	_arg0 = (*C.PangoFcFontMap)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_map_shutdown(_arg0)
}

// SubstituteChanged: call this function any time the results of the default
// substitution function set with
// [method@PangoFc.FontMap.set_default_substitute] change.
//
// That is, if your substitution function will return different results for
// the same input pattern, you must call this function.
func (f fontMap) SubstituteChanged() {
	var _arg0 *C.PangoFcFontMap // out

	_arg0 = (*C.PangoFcFontMap)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_map_substitute_changed(_arg0)
}
