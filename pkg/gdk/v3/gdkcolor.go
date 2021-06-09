// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_color_get_type()), F: marshalColor},
	})
}

// Color: a Color is used to describe a color, similar to the XColor struct used
// in the X11 drawing API.
type Color struct {
	native C.GdkColor
}

// WrapColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColor(ptr unsafe.Pointer) *Color {
	if ptr == nil {
		return nil
	}

	return (*Color)(ptr)
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColor(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Pixel gets the field inside the struct.
func (c *Color) Pixel() uint32 {
	var v uint32
	v = (uint32)(c.native.pixel)
	return v
}

// Red gets the field inside the struct.
func (c *Color) Red() uint16 {
	var v uint16
	v = (uint16)(c.native.red)
	return v
}

// Green gets the field inside the struct.
func (c *Color) Green() uint16 {
	var v uint16
	v = (uint16)(c.native.green)
	return v
}

// Blue gets the field inside the struct.
func (c *Color) Blue() uint16 {
	var v uint16
	v = (uint16)(c.native.blue)
	return v
}

// Copy makes a copy of a Color.
//
// The result must be freed using gdk_color_free().
func (c *Color) Copy() *Color {
	var _arg0 *C.GdkColor

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkColor

	cret = C.gdk_color_copy(_arg0)

	var _ret *Color

	_ret = WrapColor(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Color) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _ret
}

// Equal compares two colors.
func (c *Color) Equal(colorb *Color) bool {
	var _arg0 *C.GdkColor
	var _arg1 *C.GdkColor

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(colorb.Native()))

	var _cret C.gboolean

	cret = C.gdk_color_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees a Color created with gdk_color_copy().
func (c *Color) Free() {
	var _arg0 *C.GdkColor

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	C.gdk_color_free(_arg0)
}

// Hash: a hash function suitable for using for a hash table that stores Colors.
func (c *Color) Hash() uint {
	var _arg0 *C.GdkColor

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var _cret C.guint

	cret = C.gdk_color_hash(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// String returns a textual specification of @color in the hexadecimal form
// “\#rrrrggggbbbb” where “r”, “g” and “b” are hex digits representing the red,
// green and blue components respectively.
//
// The returned string can be parsed by gdk_color_parse().
func (c *Color) String() string {
	var _arg0 *C.GdkColor

	_arg0 = (*C.GdkColor)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar

	cret = C.gdk_color_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
