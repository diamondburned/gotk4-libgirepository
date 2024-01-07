// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
	GTypeConverter = coreglib.Type(girepository.MustFind("Gio", "Converter").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConverter, F: marshalConverter},
	})
}

// ConverterOverrider contains methods that are overridable.
type ConverterOverrider interface {
}

// Converter is implemented by objects that convert binary data in various ways.
// The conversion can be stateful and may fail at any place.
//
// Some example conversions are: character set conversion, compression,
// decompression and regular expression replace.
//
// Converter wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Converter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Converter)(nil)
)

// Converterer describes Converter's interface methods.
type Converterer interface {
	coreglib.Objector

	baseConverter() *Converter
}

var _ Converterer = (*Converter)(nil)

func ifaceInitConverterer(gifacePtr, data C.gpointer) {
}

func wrapConverter(obj *coreglib.Object) *Converter {
	return &Converter{
		Object: obj,
	}
}

func marshalConverter(p uintptr) (interface{}, error) {
	return wrapConverter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Converter) baseConverter() *Converter {
	return v
}

// BaseConverter returns the underlying base object.
func BaseConverter(obj Converterer) *Converter {
	return obj.baseConverter()
}

// ConverterIface provides an interface for converting data from one type to
// another type. The conversion can be stateful and may fail at any place.
//
// An instance of this type is always passed by reference.
type ConverterIface struct {
	*converterIface
}

// converterIface is the struct that's finalized.
type converterIface struct {
	native unsafe.Pointer
}

var GIRInfoConverterIface = girepository.MustFind("Gio", "ConverterIface")