// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewDBusServerSync creates a new D-Bus server that listens on the first
// address in address that works.
//
// Once constructed, you can use g_dbus_server_get_client_address() to get a
// D-Bus address string that clients can use to connect.
//
// To have control over the available authentication mechanisms and the users
// that are authorized to connect, it is strongly recommended to provide a
// non-NULL BusAuthObserver.
//
// Connect to the BusServer::new-connection signal to handle incoming
// connections.
//
// The returned BusServer isn't active - you have to start it with
// g_dbus_server_start().
//
// BusServer is used in this [example][gdbus-peer-to-peer].
//
// This is a synchronous failable constructor. There is currently no
// asynchronous version.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - address d-Bus address.
//    - flags flags from the BusServerFlags enumeration.
//    - guid d-Bus GUID.
//    - observer (optional) or NULL.
//
// The function returns the following values:
//
//    - dBusServer or NULL if error is set. Free with g_object_unref().
//
func NewDBusServerSync(ctx context.Context, address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver) (*DBusServer, error) {
	var _arg5 *C.GCancellable      // out
	var _arg1 *C.gchar             // out
	var _arg2 C.GDBusServerFlags   // out
	var _arg3 *C.gchar             // out
	var _arg4 *C.GDBusAuthObserver // out
	var _cret *C.GDBusServer       // in
	var _cerr *C.GError            // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GDBusServerFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(guid)))
	defer C.free(unsafe.Pointer(_arg3))
	if observer != nil {
		_arg4 = (*C.GDBusAuthObserver)(unsafe.Pointer(coreglib.InternObject(observer).Native()))
	}

	_cret = C.g_dbus_server_new_sync(_arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(guid)
	runtime.KeepAlive(observer)

	var _dBusServer *DBusServer // out
	var _goerr error            // out

	_dBusServer = wrapDBusServer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusServer, _goerr
}

// ClientAddress gets a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses) string
// that can be used by clients to connect to server.
//
// The function returns the following values:
//
//    - utf8 d-Bus address string. Do not free, the string is owned by server.
//
func (server *DBusServer) ClientAddress() string {
	var _arg0 *C.GDBusServer // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_cret = C.g_dbus_server_get_client_address(_arg0)
	runtime.KeepAlive(server)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Flags gets the flags for server.
//
// The function returns the following values:
//
//    - dBusServerFlags: set of flags from the BusServerFlags enumeration.
//
func (server *DBusServer) Flags() DBusServerFlags {
	var _arg0 *C.GDBusServer     // out
	var _cret C.GDBusServerFlags // in

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_cret = C.g_dbus_server_get_flags(_arg0)
	runtime.KeepAlive(server)

	var _dBusServerFlags DBusServerFlags // out

	_dBusServerFlags = DBusServerFlags(_cret)

	return _dBusServerFlags
}

// GUID gets the GUID for server.
//
// The function returns the following values:
//
//    - utf8 d-Bus GUID. Do not free this string, it is owned by server.
//
func (server *DBusServer) GUID() string {
	var _arg0 *C.GDBusServer // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_cret = C.g_dbus_server_get_guid(_arg0)
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
	var _arg0 *C.GDBusServer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	_cret = C.g_dbus_server_is_active(_arg0)
	runtime.KeepAlive(server)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Start starts server.
func (server *DBusServer) Start() {
	var _arg0 *C.GDBusServer // out

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	C.g_dbus_server_start(_arg0)
	runtime.KeepAlive(server)
}

// Stop stops server.
func (server *DBusServer) Stop() {
	var _arg0 *C.GDBusServer // out

	_arg0 = (*C.GDBusServer)(unsafe.Pointer(coreglib.InternObject(server).Native()))

	C.g_dbus_server_stop(_arg0)
	runtime.KeepAlive(server)
}
