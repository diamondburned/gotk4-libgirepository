// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// RefString: a typedef for a reference-counted string. A pointer to a String
// can be treated like a standard `char*` array by all code, but can
// additionally have `g_ref_string_*()` methods called on it. `g_ref_string_*()`
// methods cannot be called on `char*` arrays not allocated using
// g_ref_string_new().
//
// If using String with autocleanups, g_autoptr() must be used rather than
// g_autofree(), so that the reference counting metadata is also freed.
type RefString byte

// RefStringAcquire acquires a reference on a string.
func RefStringAcquire(str string) string {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char

	cret = C.g_ref_string_acquire(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RefStringLength retrieves the length of @str.
func RefStringLength(str string) uint {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gsize

	cret = C.g_ref_string_length(_arg1)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// NewRefString creates a new reference counted string and copies the contents
// of @str into it.
func NewRefString(str string) string {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char

	cret = C.g_ref_string_new(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RefStringNewIntern creates a new reference counted string and copies the
// content of @str into it.
//
// If you call this function multiple times with the same @str, or with the same
// contents of @str, it will return a new reference, instead of creating a new
// string.
func RefStringNewIntern(str string) string {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char

	cret = C.g_ref_string_new_intern(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RefStringNewLen creates a new reference counted string and copies the
// contents of @str into it, up to @len bytes.
//
// Since this function does not stop at nul bytes, it is the caller's
// responsibility to ensure that @str has at least @len addressable bytes.
func RefStringNewLen(str string, len int) string {
	var _arg1 *C.char
	var _arg2 C.gssize

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(len)

	var _cret *C.char

	cret = C.g_ref_string_new_len(_arg1, _arg2)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RefStringRelease releases a reference on a string; if it was the last
// reference, the resources allocated by the string are freed as well.
func RefStringRelease(str string) {
	var _arg1 *C.char

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_ref_string_release(_arg1)
}