// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// UUIDStringIsValid parses the string @str and verify if it is a UUID.
//
// The function accepts the following syntax:
//
// - simple forms (e.g. `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`)
//
// Note that hyphens are required within the UUID string itself, as per the
// aforementioned RFC.
func UUIDStringIsValid(str string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_uuid_string_is_valid(arg1)

	if cret {
		goret = true
	}

	return goret
}

// UUIDStringRandom generates a random UUID (RFC 4122 version 4) as a string. It
// has the same randomness guarantees as #GRand, so must not be used for
// cryptographic purposes such as key generation, nonces, salts or one-time
// pads.
func UUIDStringRandom() string {
	cret := new(C.gchar)
	var goret string

	cret = C.g_uuid_string_random()

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}
