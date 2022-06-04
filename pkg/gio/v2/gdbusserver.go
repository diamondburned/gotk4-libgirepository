// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// ClientAddress gets a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses) string
// that can be used by clients to connect to server.
//
// The function returns the following values:
//
//    - utf8 d-Bus address string. Do not free, the string is owned by server.
//
func (server *DBusServer) ClientAddress() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_gret := girepository.MustFind("Gio", "DBusServer").InvokeMethod("get_client_address", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(server)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GUID gets the GUID for server.
//
// The function returns the following values:
//
//    - utf8 d-Bus GUID. Do not free this string, it is owned by server.
//
func (server *DBusServer) GUID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_gret := girepository.MustFind("Gio", "DBusServer").InvokeMethod("get_guid", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(server)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsActive gets whether server is active.
//
// The function returns the following values:
//
//    - ok: TRUE if server is active, FALSE otherwise.
//
func (server *DBusServer) IsActive() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_gret := girepository.MustFind("Gio", "DBusServer").InvokeMethod("is_active", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(server)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Start starts server.
func (server *DBusServer) Start() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	girepository.MustFind("Gio", "DBusServer").InvokeMethod("start", _args[:], nil)

	runtime.KeepAlive(server)
}

// Stop stops server.
func (server *DBusServer) Stop() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	girepository.MustFind("Gio", "DBusServer").InvokeMethod("stop", _args[:], nil)

	runtime.KeepAlive(server)
}
