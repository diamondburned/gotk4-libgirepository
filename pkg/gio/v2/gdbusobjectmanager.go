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
		{T: externglib.Type(C.g_dbus_object_manager_get_type()), F: marshalDBusObjectManager},
	})
}

// DBusObjectManagerOverrider contains methods that are overridable. This
// interface is a subset of the interface DBusObjectManager.
type DBusObjectManagerOverrider interface {
	// ObjectPath gets the object path that @manager is for.
	ObjectPath() string

	InterfaceAdded(object DBusObject, interface_ DBusInterface)

	InterfaceRemoved(object DBusObject, interface_ DBusInterface)

	ObjectAdded(object DBusObject)

	ObjectRemoved(object DBusObject)
}

// DBusObjectManager: the BusObjectManager type is the base type for service-
// and client-side implementations of the standardized
// org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
type DBusObjectManager interface {
	gextras.Objector
	DBusObjectManagerOverrider
}

// dBusObjectManager implements the DBusObjectManager interface.
type dBusObjectManager struct {
	gextras.Objector
}

var _ DBusObjectManager = (*dBusObjectManager)(nil)

// WrapDBusObjectManager wraps a GObject to a type that implements interface
// DBusObjectManager. It is primarily used internally.
func WrapDBusObjectManager(obj *externglib.Object) DBusObjectManager {
	return DBusObjectManager{
		Objector: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectManager(obj), nil
}

// ObjectPath gets the object path that @manager is for.
func (m dBusObjectManager) ObjectPath() string {
	var _arg0 *C.GDBusObjectManager // out

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))

	var _cret *C.gchar // in

	_cret = C.g_dbus_object_manager_get_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
