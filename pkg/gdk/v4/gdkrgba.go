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

// RGBA: GdkRGBA is used to represent a color, in a way that is compatible with
// cairo’s notion of color.
//
// GdkRGBA is a convenient way to pass colors around. It’s based on cairo’s way
// to deal with colors and mirrors its behavior. All values are in the range
// from 0.0 to 1.0 inclusive. So the color (0.0, 0.0, 0.0, 0.0) represents
// transparent black and (1.0, 1.0, 1.0, 1.0) is opaque white. Other values will
// be clamped to this range when drawing.
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
func NewRGBA(red, green, blue, alpha float32) RGBA {
	var f0 C.float // out
	f0 = C.float(red)
	var f1 C.float // out
	f1 = C.float(green)
	var f2 C.float // out
	f2 = C.float(blue)
	var f3 C.float // out
	f3 = C.float(alpha)

	size := GIRInfoRGBA.StructSize()
	native := make([]byte, size)
	gextras.Sink(&native[0])

	offset0 := GIRInfoRGBA.StructFieldOffset("red")
	valptr0 := (*C.float)(unsafe.Add(unsafe.Pointer(&native[0]), offset0))
	*valptr0 = f0

	offset1 := GIRInfoRGBA.StructFieldOffset("green")
	valptr1 := (*C.float)(unsafe.Add(unsafe.Pointer(&native[0]), offset1))
	*valptr1 = f1

	offset2 := GIRInfoRGBA.StructFieldOffset("blue")
	valptr2 := (*C.float)(unsafe.Add(unsafe.Pointer(&native[0]), offset2))
	*valptr2 = f2

	offset3 := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr3 := (*C.float)(unsafe.Add(unsafe.Pointer(&native[0]), offset3))
	*valptr3 = f3

	return *(*RGBA)(gextras.NewStructNative(unsafe.Pointer(&native[0])))
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Red() float32 {
	offset := GIRInfoRGBA.StructFieldOffset("red")
	valptr := (*float32)(unsafe.Add(r.native, offset))
	var _v float32 // out
	_v = float32(*valptr)
	return _v
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Green() float32 {
	offset := GIRInfoRGBA.StructFieldOffset("green")
	valptr := (*float32)(unsafe.Add(r.native, offset))
	var _v float32 // out
	_v = float32(*valptr)
	return _v
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Blue() float32 {
	offset := GIRInfoRGBA.StructFieldOffset("blue")
	valptr := (*float32)(unsafe.Add(r.native, offset))
	var _v float32 // out
	_v = float32(*valptr)
	return _v
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) Alpha() float32 {
	offset := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr := (*float32)(unsafe.Add(r.native, offset))
	var _v float32 // out
	_v = float32(*valptr)
	return _v
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetRed(red float32) {
	offset := GIRInfoRGBA.StructFieldOffset("red")
	valptr := (*C.float)(unsafe.Add(r.native, offset))
	*valptr = C.float(red)
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetGreen(green float32) {
	offset := GIRInfoRGBA.StructFieldOffset("green")
	valptr := (*C.float)(unsafe.Add(r.native, offset))
	*valptr = C.float(green)
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetBlue(blue float32) {
	offset := GIRInfoRGBA.StructFieldOffset("blue")
	valptr := (*C.float)(unsafe.Add(r.native, offset))
	*valptr = C.float(blue)
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) SetAlpha(alpha float32) {
	offset := GIRInfoRGBA.StructFieldOffset("alpha")
	valptr := (*C.float)(unsafe.Add(r.native, offset))
	*valptr = C.float(alpha)
}
