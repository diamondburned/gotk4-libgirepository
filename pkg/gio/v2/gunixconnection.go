// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_unix_connection_get_type()), F: marshalUnixConnection},
	})
}

// UnixConnection: this is the subclass of Connection that is created for UNIX
// domain sockets.
//
// It contains functions to do some of the UNIX socket specific functionality
// like passing file descriptors.
//
// Note that `<gio/gunixconnection.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixConnection interface {
	SocketConnection
}

// unixConnection implements the UnixConnection class.
type unixConnection struct {
	SocketConnection
}

var _ UnixConnection = (*unixConnection)(nil)

// WrapUnixConnection wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixConnection(obj *externglib.Object) UnixConnection {
	return unixConnection{
		SocketConnection: WrapSocketConnection(obj),
	}
}

func marshalUnixConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixConnection(obj), nil
}
