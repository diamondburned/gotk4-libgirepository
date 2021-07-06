// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_color_get_type()), F: marshalColor},
	})
}

// Color is used to describe a color, similar to the XColor struct used in the
// X11 drawing API.
//
// Deprecated: since version 3.14.
type Color struct {
	native C.GdkColor
}

// WrapColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColor(ptr unsafe.Pointer) *Color {
	return (*Color)(ptr)
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Color)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Pixel: for allocated colors, the pixel value used to draw this color on the
// screen. Not used anymore.
func (c *Color) Pixel() uint32 {
	var v uint32 // out
	v = uint32(c.pixel)
	return v
}

// Red: the red component of the color. This is a value between 0 and 65535,
// with 65535 indicating full intensity
func (c *Color) Red() uint16 {
	var v uint16 // out
	v = uint16(c.red)
	return v
}

// Green: the green component of the color
func (c *Color) Green() uint16 {
	var v uint16 // out
	v = uint16(c.green)
	return v
}

// Blue: the blue component of the color
func (c *Color) Blue() uint16 {
	var v uint16 // out
	v = uint16(c.blue)
	return v
}

// Copy makes a copy of a Color.
//
// The result must be freed using gdk_color_free().
//
// Deprecated: since version 3.14.
func (c *Color) Copy() *Color {
	var _arg0 *C.GdkColor // out
	var _cret *C.GdkColor // in

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c))

	_cret = C.gdk_color_copy(_arg0)

	var _ret *Color // out

	_ret = (*Color)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Color) {
		C.free(unsafe.Pointer(v))
	})

	return _ret
}

// Equal compares two colors.
//
// Deprecated: since version 3.14.
func (c *Color) Equal(colorb *Color) bool {
	var _arg0 *C.GdkColor // out
	var _arg1 *C.GdkColor // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(colorb))

	_cret = C.gdk_color_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Free frees a Color created with gdk_color_copy().
//
// Deprecated: since version 3.14.
func (c *Color) free() {
	var _arg0 *C.GdkColor // out

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c))

	C.gdk_color_free(_arg0)
}

// Hash: hash function suitable for using for a hash table that stores Colors.
//
// Deprecated: since version 3.14.
func (c *Color) Hash() uint {
	var _arg0 *C.GdkColor // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c))

	_cret = C.gdk_color_hash(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// String returns a textual specification of @color in the hexadecimal form
// “\#rrrrggggbbbb” where “r”, “g” and “b” are hex digits representing the red,
// green and blue components respectively.
//
// The returned string can be parsed by gdk_color_parse().
//
// Deprecated: since version 3.14.
func (c *Color) String() string {
	var _arg0 *C.GdkColor // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c))

	_cret = C.gdk_color_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
