// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_ConverterIface_reset(void*);
import "C"

// GTypeConverter returns the GType for the type Converter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConverter() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "Converter").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConverter)
	return gtype
}

// ConverterOverrider contains methods that are overridable.
type ConverterOverrider interface {
	// Reset resets all internal state in the converter, making it behave as if
	// it was just created. If the converter has any internal state that would
	// produce output then that output is lost.
	Reset()
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

	// Reset resets all internal state in the converter, making it behave as if
	// it was just created.
	Reset()
}

var _ Converterer = (*Converter)(nil)

func ifaceInitConverterer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "ConverterIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("reset"))) = unsafe.Pointer(C._gotk4_gio2_ConverterIface_reset)
}

//export _gotk4_gio2_ConverterIface_reset
func _gotk4_gio2_ConverterIface_reset(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ConverterOverrider)

	iface.Reset()
}

func wrapConverter(obj *coreglib.Object) *Converter {
	return &Converter{
		Object: obj,
	}
}

func marshalConverter(p uintptr) (interface{}, error) {
	return wrapConverter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Reset resets all internal state in the converter, making it behave as if it
// was just created. If the converter has any internal state that would produce
// output then that output is lost.
func (converter *Converter) Reset() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(converter).Native()))

	_info := girepository.MustFind("Gio", "Converter")
	_info.InvokeIfaceMethod("reset", _args[:], nil)

	runtime.KeepAlive(converter)
}
