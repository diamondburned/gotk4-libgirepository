// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_dbus_object_proxy_get_type()), F: marshalDBusObjectProXY},
	})
}

// DBusObjectProXY: a BusObjectProxy is an object used to represent a remote
// object with one or more D-Bus interfaces. Normally, you don't instantiate a
// BusObjectProxy yourself - typically BusObjectManagerClient is used to obtain
// it.
type DBusObjectProXY interface {
	gextras.Objector
	DBusObject
}

// dBusObjectProXY implements the DBusObjectProXY class.
type dBusObjectProXY struct {
	gextras.Objector
	DBusObject
}

var _ DBusObjectProXY = (*dBusObjectProXY)(nil)

// WrapDBusObjectProXY wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectProXY(obj *externglib.Object) DBusObjectProXY {
	return dBusObjectProXY{
		Objector:   obj,
		DBusObject: WrapDBusObject(obj),
	}
}

func marshalDBusObjectProXY(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectProXY(obj), nil
}
