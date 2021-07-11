// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// Base64Encode: encode a sequence of binary data into its Base-64 stringified
// representation.
func Base64Encode(data []byte) string {
	var _arg1 *C.guchar
	var _arg2 C.gsize
	var _cret *C.gchar // in

	_arg2 = C.gsize(len(data))
	_arg1 = (*C.guchar)(unsafe.Pointer(&data[0]))

	_cret = C.g_base64_encode(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
