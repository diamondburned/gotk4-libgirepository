// Code generated by girgen. DO NOT EDIT.

package pangofc

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangofc
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangofc-fontmap.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fc_decoder_get_type()), F: marshalDecoder},
	})
}

// Decoder: `PangoFcDecoder` is a virtual base class that implementations will
// inherit from.
//
// It's the interface that is used to define a custom encoding for a font. These
// objects are created in your code from a function callback that was originally
// registered with [method@PangoFc.FontMap.add_decoder_find_func]. Pango
// requires information about the supported charset for a font as well as the
// individual character to glyph conversions. Pango gets that information via
// the #get_charset and #get_glyph callbacks into your object implementation.
type Decoder interface {
	gextras.Objector
}

// decoder implements the Decoder class.
type decoder struct {
	gextras.Objector
}

var _ Decoder = (*decoder)(nil)

// WrapDecoder wraps a GObject to the right type. It is
// primarily used internally.
func WrapDecoder(obj *externglib.Object) Decoder {
	return decoder{
		Objector: obj,
	}
}

func marshalDecoder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDecoder(obj), nil
}
