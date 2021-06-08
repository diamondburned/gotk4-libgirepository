// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_interface_get_type()), F: marshalDBusInterface},
	})
}

// DBusInterfaceOverrider contains methods that are overridable. This
// interface is a subset of the interface DBusInterface.
type DBusInterfaceOverrider interface {
	// DupObject gets the BusObject that @interface_ belongs to, if any.
	DupObject() DBusObject
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by @interface_.
	Info() *DBusInterfaceInfo
	// Object gets the BusObject that @interface_ belongs to, if any.
	//
	// It is not safe to use the returned object if @interface_ or the returned
	// object is being used from other threads. See
	// g_dbus_interface_dup_object() for a thread-safe alternative.
	Object() DBusObject
	// SetObject sets the BusObject for @interface_ to @object.
	//
	// Note that @interface_ will hold a weak reference to @object.
	SetObject(object DBusObject)
}

// DBusInterface: the BusInterface type is the base type for D-Bus interfaces
// both on the service side (see BusInterfaceSkeleton) and client side (see
// BusProxy).
type DBusInterface interface {
	gextras.Objector
	DBusInterfaceOverrider
}

// dBusInterface implements the DBusInterface interface.
type dBusInterface struct {
	gextras.Objector
}

var _ DBusInterface = (*dBusInterface)(nil)

// WrapDBusInterface wraps a GObject to a type that implements interface
// DBusInterface. It is primarily used internally.
func WrapDBusInterface(obj *externglib.Object) DBusInterface {
	return DBusInterface{
		Objector: obj,
	}
}

func marshalDBusInterface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusInterface(obj), nil
}

// DupObject gets the BusObject that @interface_ belongs to, if any.
func (i dBusInterface) DupObject() DBusObject {
	var arg0 *C.GDBusInterface

	arg0 = (*C.GDBusInterface)(unsafe.Pointer(i.Native()))

	cret := new(C.GDBusObject)
	var goret DBusObject

	cret = C.g_dbus_interface_dup_object(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObject)

	return goret
}

// Info gets D-Bus introspection information for the D-Bus interface
// implemented by @interface_.
func (i dBusInterface) Info() *DBusInterfaceInfo {
	var arg0 *C.GDBusInterface

	arg0 = (*C.GDBusInterface)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusInterfaceInfo
	var goret *DBusInterfaceInfo

	cret = C.g_dbus_interface_get_info(arg0)

	goret = WrapDBusInterfaceInfo(unsafe.Pointer(cret))

	return goret
}

// SetObject sets the BusObject for @interface_ to @object.
//
// Note that @interface_ will hold a weak reference to @object.
func (i dBusInterface) SetObject(object DBusObject) {
	var arg0 *C.GDBusInterface
	var arg1 *C.GDBusObject

	arg0 = (*C.GDBusInterface)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	C.g_dbus_interface_set_object(arg0, arg1)
}
