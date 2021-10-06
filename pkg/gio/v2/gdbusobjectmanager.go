// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_dbus_object_manager_get_type()), F: marshalDBusObjectManagerer},
	})
}

// DBusObjectManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectManagerOverrider interface {
	// Interface gets the interface proxy for interface_name at object_path, if
	// any.
	Interface(objectPath, interfaceName string) DBusInterfacer
	// GetObject gets the BusObjectProxy at object_path, if any.
	GetObject(objectPath string) DBusObjector
	// ObjectPath gets the object path that manager is for.
	ObjectPath() string
	// Objects gets all BusObject objects known to manager.
	Objects() []DBusObjector
	InterfaceAdded(object DBusObjector, interface_ DBusInterfacer)
	InterfaceRemoved(object DBusObjector, interface_ DBusInterfacer)
	ObjectAdded(object DBusObjector)
	ObjectRemoved(object DBusObjector)
}

// DBusObjectManager type is the base type for service- and client-side
// implementations of the standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
type DBusObjectManager struct {
	*externglib.Object
}

// DBusObjectManagerer describes DBusObjectManager's abstract methods.
type DBusObjectManagerer interface {
	externglib.Objector

	// Interface gets the interface proxy for interface_name at object_path, if
	// any.
	Interface(objectPath, interfaceName string) DBusInterfacer
	// GetObject gets the BusObjectProxy at object_path, if any.
	GetObject(objectPath string) DBusObjector
	// ObjectPath gets the object path that manager is for.
	ObjectPath() string
	// Objects gets all BusObject objects known to manager.
	Objects() []DBusObjector
}

var _ DBusObjectManagerer = (*DBusObjectManager)(nil)

func wrapDBusObjectManager(obj *externglib.Object) *DBusObjectManager {
	return &DBusObjectManager{
		Object: obj,
	}
}

func marshalDBusObjectManagerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectManager(obj), nil
}

// Interface gets the interface proxy for interface_name at object_path, if any.
//
// The function takes the following parameters:
//
//    - objectPath: object path to look up.
//    - interfaceName d-Bus interface name to look up.
//
func (manager *DBusObjectManager) Interface(objectPath, interfaceName string) DBusInterfacer {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _cret *C.GDBusInterface     // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_object_manager_get_interface(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(DBusInterfacer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.DBusInterfacer")
		}
		_dBusInterface = rv
	}

	return _dBusInterface
}

// GetObject gets the BusObjectProxy at object_path, if any.
//
// The function takes the following parameters:
//
//    - objectPath: object path to look up.
//
func (manager *DBusObjectManager) GetObject(objectPath string) DBusObjector {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusObject        // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_manager_get_object(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(objectPath)

	var _dBusObject DBusObjector // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.DBusObjector is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(DBusObjector)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.DBusObjector")
		}
		_dBusObject = rv
	}

	return _dBusObject
}

// ObjectPath gets the object path that manager is for.
func (manager *DBusObjectManager) ObjectPath() string {
	var _arg0 *C.GDBusObjectManager // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_get_object_path(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Objects gets all BusObject objects known to manager.
func (manager *DBusObjectManager) Objects() []DBusObjector {
	var _arg0 *C.GDBusObjectManager // out
	var _cret *C.GList              // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(manager.Native()))

	_cret = C.g_dbus_object_manager_get_objects(_arg0)
	runtime.KeepAlive(manager)

	var _list []DBusObjector // out

	_list = make([]DBusObjector, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusObject)(v)
		var dst DBusObjector // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.DBusObjector is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(DBusObjector)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.DBusObjector")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// ConnectInterfaceAdded: emitted when interface is added to object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceAdded(f func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle {
	return manager.Connect("interface-added", f)
}

// ConnectInterfaceRemoved: emitted when interface has been removed from object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (manager *DBusObjectManager) ConnectInterfaceRemoved(f func(object DBusObjector, iface DBusInterfacer)) externglib.SignalHandle {
	return manager.Connect("interface-removed", f)
}

// ConnectObjectAdded: emitted when object is added to manager.
func (manager *DBusObjectManager) ConnectObjectAdded(f func(object DBusObjector)) externglib.SignalHandle {
	return manager.Connect("object-added", f)
}

// ConnectObjectRemoved: emitted when object is removed from manager.
func (manager *DBusObjectManager) ConnectObjectRemoved(f func(object DBusObjector)) externglib.SignalHandle {
	return manager.Connect("object-removed", f)
}
