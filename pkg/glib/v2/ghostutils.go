// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// HostnameIsASCIIEncoded tests if @hostname contains segments with an
// ASCII-compatible encoding of an Internationalized Domain Name. If this
// returns true, you should decode the hostname with g_hostname_to_unicode()
// before displaying it to the user.
//
// Note that a hostname might contain a mix of encoded and unencoded segments,
// and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return true for a name.
func HostnameIsASCIIEncoded(hostname string) bool {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_hostname_is_ascii_encoded(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HostnameIsIpAddress tests if @hostname is the string form of an IPv4 or IPv6
// address. (Eg, "192.168.0.1".)
//
// Since 2.66, IPv6 addresses with a zone-id are accepted (RFC6874).
func HostnameIsIpAddress(hostname string) bool {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_hostname_is_ip_address(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HostnameIsNonASCII tests if @hostname contains Unicode characters. If this
// returns true, you need to encode the hostname with g_hostname_to_ascii()
// before using it in non-IDN-aware contexts.
//
// Note that a hostname might contain a mix of encoded and unencoded segments,
// and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return true for a name.
func HostnameIsNonASCII(hostname string) bool {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_hostname_is_non_ascii(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HostnameToASCII converts @hostname to its canonical ASCII form; an ASCII-only
// string containing no uppercase letters and not ending with a trailing dot.
func HostnameToASCII(hostname string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_hostname_to_ascii(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HostnameToUnicode converts @hostname to its canonical presentation form; a
// UTF-8 string in Unicode normalization form C, containing no uppercase
// letters, no forbidden characters, and no ASCII-encoded segments, and not
// ending with a trailing dot.
//
// Of course if @hostname is not an internationalized hostname, then the
// canonical presentation form will be entirely ASCII.
func HostnameToUnicode(hostname string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_hostname_to_unicode(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
