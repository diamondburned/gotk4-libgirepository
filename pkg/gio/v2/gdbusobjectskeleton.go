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
		{T: externglib.Type(C.g_dbus_object_skeleton_get_type()), F: marshalDBusObjectSkeleton},
	})
}

// DBusObjectSkeleton: a BusObjectSkeleton instance is essentially a group of
// D-Bus interfaces. The set of exported interfaces on the object may be dynamic
// and change at runtime.
//
// This type is intended to be used with BusObjectManager.
type DBusObjectSkeleton interface {
	gextras.Objector
	DBusObject

	// AddInterface adds @interface_ to @object.
	//
	// If @object already contains a BusInterfaceSkeleton with the same
	// interface name, it is removed before @interface_ is added.
	//
	// Note that @object takes its own reference on @interface_ and holds it
	// until removed.
	AddInterface(interface_ DBusInterfaceSkeleton)
	// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
	// interfaces belonging to @object. See that method for when flushing is
	// useful.
	Flush()
	// RemoveInterface removes @interface_ from @object.
	RemoveInterface(interface_ DBusInterfaceSkeleton)
	// RemoveInterfaceByName removes the BusInterface with @interface_name from
	// @object.
	//
	// If no D-Bus interface of the given interface exists, this function does
	// nothing.
	RemoveInterfaceByName(interfaceName string)
	// SetObjectPath sets the object path for @object.
	SetObjectPath(objectPath string)
}

// dBusObjectSkeleton implements the DBusObjectSkeleton interface.
type dBusObjectSkeleton struct {
	gextras.Objector
	DBusObject
}

var _ DBusObjectSkeleton = (*dBusObjectSkeleton)(nil)

// WrapDBusObjectSkeleton wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectSkeleton(obj *externglib.Object) DBusObjectSkeleton {
	return DBusObjectSkeleton{
		Objector:   obj,
		DBusObject: WrapDBusObject(obj),
	}
}

func marshalDBusObjectSkeleton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectSkeleton(obj), nil
}

// NewDBusObjectSkeleton constructs a class DBusObjectSkeleton.
func NewDBusObjectSkeleton(objectPath string) DBusObjectSkeleton {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GDBusObjectSkeleton

	cret = C.g_dbus_object_skeleton_new(_arg1)

	var _dBusObjectSkeleton DBusObjectSkeleton

	_dBusObjectSkeleton = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DBusObjectSkeleton)

	return _dBusObjectSkeleton
}

// AddInterface adds @interface_ to @object.
//
// If @object already contains a BusInterfaceSkeleton with the same
// interface name, it is removed before @interface_ is added.
//
// Note that @object takes its own reference on @interface_ and holds it
// until removed.
func (o dBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeleton) {
	var _arg0 *C.GDBusObjectSkeleton
	var _arg1 *C.GDBusInterfaceSkeleton

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_add_interface(_arg0, _arg1)
}

// Flush: this method simply calls g_dbus_interface_skeleton_flush() on all
// interfaces belonging to @object. See that method for when flushing is
// useful.
func (o dBusObjectSkeleton) Flush() {
	var _arg0 *C.GDBusObjectSkeleton

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(o.Native()))

	C.g_dbus_object_skeleton_flush(_arg0)
}

// RemoveInterface removes @interface_ from @object.
func (o dBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeleton) {
	var _arg0 *C.GDBusObjectSkeleton
	var _arg1 *C.GDBusInterfaceSkeleton

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(interface_.Native()))

	C.g_dbus_object_skeleton_remove_interface(_arg0, _arg1)
}

// RemoveInterfaceByName removes the BusInterface with @interface_name from
// @object.
//
// If no D-Bus interface of the given interface exists, this function does
// nothing.
func (o dBusObjectSkeleton) RemoveInterfaceByName(interfaceName string) {
	var _arg0 *C.GDBusObjectSkeleton
	var _arg1 *C.gchar

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_remove_interface_by_name(_arg0, _arg1)
}

// SetObjectPath sets the object path for @object.
func (o dBusObjectSkeleton) SetObjectPath(objectPath string) {
	var _arg0 *C.GDBusObjectSkeleton
	var _arg1 *C.gchar

	_arg0 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_dbus_object_skeleton_set_object_path(_arg0, _arg1)
}
