// Code generated by girgen. DO NOT EDIT.

package gdk

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
	GTypeRGBA = coreglib.Type(girepository.MustFind("Gdk", "RGBA").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRGBA, F: marshalRGBA},
	})
}

// RGBA is used to represent a (possibly translucent) color, in a way that is
// compatible with cairo’s notion of color.
//
// An instance of this type is always passed by reference.
type RGBA struct {
	*rgbA
}

// rgbA is the struct that's finalized.
type rgbA struct {
	native unsafe.Pointer
}

var GIRInfoRGBA = girepository.MustFind("Gdk", "RGBA")

func marshalRGBA(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &RGBA{&rgbA{(unsafe.Pointer)(b)}}, nil
}

// NewRGBA creates a new RGBA instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewRGBA(red, green, blue, alpha float64) RGBA {
	var f0 C.gdouble // out
	f0 = C.gdouble(red)
	var f1 C.gdouble // out
	f1 = C.gdouble(green)
	var f2 C.gdouble // out
	f2 = C.gdouble(blue)
	var f3 C.gdouble // out
	f3 = C.gdouble(alpha)

	size := GIRInfoRGBA.StructSize()
	native := make([]byte, size)
	gextras.Sink(&native[0])

	offset0 := GIRInfoRGBA.StructFieldOffset("red")
	valptr0 := (*C.gdouble)(unsafe.Add(unsafe.Pointer(&native[0]), offset0))
	*valptr0 = f0

	offset1 := GIRInfoRGBA.StructFieldOffset("green")
	valptr1 := (*C.gdouble)(unsafe.Add(unsafe.Pointer(&native[0]), offset1))
	*valptr1 = f1

	offset2 := GIRInfoRGBA.StructFieldOffset("blue")
	valptr2 := (*C.gdouble)(unsafe.Add(unsafe.Pointer(&native[0]), offset2))
	*valptr2 = f2

	offset3 := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr3 := (*C.gdouble)(unsafe.Add(unsafe.Pointer(&native[0]), offset3))
	*valptr3 = f3

	return *(*RGBA)(gextras.NewStructNative(unsafe.Pointer(&native[0])))
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Red() float64 {
	offset := GIRInfoRGBA.StructFieldOffset("red")
	valptr := (*float64)(unsafe.Add(r.native, offset))
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Green() float64 {
	offset := GIRInfoRGBA.StructFieldOffset("green")
	valptr := (*float64)(unsafe.Add(r.native, offset))
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Blue() float64 {
	offset := GIRInfoRGBA.StructFieldOffset("blue")
	valptr := (*float64)(unsafe.Add(r.native, offset))
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) Alpha() float64 {
	offset := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr := (*float64)(unsafe.Add(r.native, offset))
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetRed(red float64) {
	offset := GIRInfoRGBA.StructFieldOffset("red")
	valptr := (*C.gdouble)(unsafe.Add(r.native, offset))
	*valptr = C.gdouble(red)
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetGreen(green float64) {
	offset := GIRInfoRGBA.StructFieldOffset("green")
	valptr := (*C.gdouble)(unsafe.Add(r.native, offset))
	*valptr = C.gdouble(green)
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetBlue(blue float64) {
	offset := GIRInfoRGBA.StructFieldOffset("blue")
	valptr := (*C.gdouble)(unsafe.Add(r.native, offset))
	*valptr = C.gdouble(blue)
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) SetAlpha(alpha float64) {
	offset := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr := (*C.gdouble)(unsafe.Add(r.native, offset))
	*valptr = C.gdouble(alpha)
}
