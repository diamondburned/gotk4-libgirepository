// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFontMap = coreglib.Type(girepository.MustFind("PangoCairo", "FontMap").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFontMap, F: marshalFontMap},
	})
}

// FontMap: PangoCairoFontMap is an interface exported by font maps for use with
// Cairo.
//
// The actual type of the font map will depend on the particular font technology
// Cairo was compiled to use.
//
// FontMap wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FontMap struct {
	_ [0]func() // equal guard
	pango.FontMap
}

var (
	_ pango.FontMapper = (*FontMap)(nil)
)

// FontMapper describes FontMap's interface methods.
type FontMapper interface {
	coreglib.Objector

	baseFontMap() *FontMap
}

var _ FontMapper = (*FontMap)(nil)

func wrapFontMap(obj *coreglib.Object) *FontMap {
	return &FontMap{
		FontMap: pango.FontMap{
			Object: obj,
		},
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	return wrapFontMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *FontMap) baseFontMap() *FontMap {
	return v
}

// BaseFontMap returns the underlying base object.
func BaseFontMap(obj FontMapper) *FontMap {
	return obj.baseFontMap()
}
