// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
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

// GTypeRange returns the GType for the type Range.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRange() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "Range").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRange)
	return gtype
}

// Range are used on Value, in order to represent the full range of a given
// component (for example an slider or a range control), or to define each
// individual subrange this full range is splitted if available. See Value
// documentation for further details.
//
// An instance of this type is always passed by reference.
type Range struct {
	*_range
}

// _range is the struct that's finalized.
type _range struct {
	native unsafe.Pointer
}

func marshalRange(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Range{&_range{(unsafe.Pointer)(b)}}, nil
}

// NewRange constructs a struct Range.
func NewRange(lowerLimit float64, upperLimit float64, description string) *Range {
	var _args [3]girepository.Argument

	*(*C.gdouble)(unsafe.Pointer(&_args[0])) = C.gdouble(lowerLimit)
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(upperLimit)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(lowerLimit)
	runtime.KeepAlive(upperLimit)
	runtime.KeepAlive(description)

	var __range *Range // out

	__range = (*Range)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(__range)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("Atk", "Range").InvokeMethod("free", args[:], nil)
			}
		},
	)

	return __range
}

// Copy returns a new Range that is a exact copy of src.
//
// The function returns the following values:
//
//    - _range: new Range copy of src.
//
func (src *Range) Copy() *Range {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(src)

	var __range *Range // out

	__range = (*Range)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(__range)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("Atk", "Range").InvokeMethod("free", args[:], nil)
			}
		},
	)

	return __range
}

// Description returns the human readable description of range.
//
// The function returns the following values:
//
//    - utf8: human-readable description of range.
//
func (_range *Range) Description() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(_range)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LowerLimit returns the lower limit of range.
//
// The function returns the following values:
//
//    - gdouble: lower limit of range.
//
func (_range *Range) LowerLimit() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(_range)))

	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// UpperLimit returns the upper limit of range.
//
// The function returns the following values:
//
//    - gdouble: upper limit of range.
//
func (_range *Range) UpperLimit() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(_range)))

	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}
