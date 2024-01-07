// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFontMap = coreglib.Type(girepository.MustFind("Pango", "FontMap").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFontMap, F: marshalFontMap},
	})
}

// FontMapOverrides contains methods that are overridable.
type FontMapOverrides struct {
}

func defaultFontMapOverrides(v *FontMap) FontMapOverrides {
	return FontMapOverrides{}
}

// FontMap: PangoFontMap represents the set of fonts available for a particular
// rendering system.
//
// This is a virtual object with implementations being specific to particular
// rendering systems.
type FontMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FontMap)(nil)
)

// FontMapper describes types inherited from class FontMap.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FontMapper interface {
	coreglib.Objector
	baseFontMap() *FontMap
}

var _ FontMapper = (*FontMap)(nil)

func init() {
	coreglib.RegisterClassInfo[*FontMap, *FontMapClass, FontMapOverrides](
		GTypeFontMap,
		initFontMapClass,
		wrapFontMap,
		defaultFontMapOverrides,
	)
}

func initFontMapClass(gclass unsafe.Pointer, overrides FontMapOverrides, classInitFunc func(*FontMapClass)) {
	if classInitFunc != nil {
		class := (*FontMapClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFontMap(obj *coreglib.Object) *FontMap {
	return &FontMap{
		Object: obj,
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

// FontMapClass structure holds the virtual functions for a particular FontMap
// implementation.
//
// An instance of this type is always passed by reference.
type FontMapClass struct {
	*fontMapClass
}

// fontMapClass is the struct that's finalized.
type fontMapClass struct {
	native unsafe.Pointer
}

var GIRInfoFontMapClass = girepository.MustFind("Pango", "FontMapClass")

// ShapeEngineType: type of rendering-system-dependent engines that can handle
// fonts of this fonts loaded with this fontmap.
func (f *FontMapClass) ShapeEngineType() string {
	offset := GIRInfoFontMapClass.StructFieldOffset("shape_engine_type")
	valptr := (*string)(unsafe.Add(f.native, offset))
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}
