// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectRemoved(gpointer, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectObjectAdded(gpointer, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved(gpointer, void*, void*, guintptr);
// extern void _gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded(gpointer, void*, void*, guintptr);
import "C"

// GType values.
var (
	GTypeDBusObjectManager = coreglib.Type(girepository.MustFind("Gio", "DBusObjectManager").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDBusObjectManager, F: marshalDBusObjectManager},
	})
}

// DBusObjectManagerOverrider contains methods that are overridable.
type DBusObjectManagerOverrider interface {
}

// DBusObjectManager type is the base type for service- and client-side
// implementations of the standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
//
// DBusObjectManager wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DBusObjectManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusObjectManager)(nil)
)

// DBusObjectManagerer describes DBusObjectManager's interface methods.
type DBusObjectManagerer interface {
	coreglib.Objector

	baseDBusObjectManager() *DBusObjectManager
}

var _ DBusObjectManagerer = (*DBusObjectManager)(nil)

func ifaceInitDBusObjectManagerer(gifacePtr, data C.gpointer) {
}

func wrapDBusObjectManager(obj *coreglib.Object) *DBusObjectManager {
	return &DBusObjectManager{
		Object: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	return wrapDBusObjectManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *DBusObjectManager) baseDBusObjectManager() *DBusObjectManager {
	return v
}

// BaseDBusObjectManager returns the underlying base object.
func BaseDBusObjectManager(obj DBusObjectManagerer) *DBusObjectManager {
	return obj.baseDBusObjectManager()
}

// ConnectInterfaceAdded is emitted when interface is added to object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (v *DBusObjectManager) ConnectInterfaceAdded(f func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "interface-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceAdded), f)
}

// ConnectInterfaceRemoved is emitted when interface has been removed from
// object.
//
// This signal exists purely as a convenience to avoid having to connect signals
// to all objects managed by manager.
func (v *DBusObjectManager) ConnectInterfaceRemoved(f func(object DBusObjector, iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "interface-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectInterfaceRemoved), f)
}

// ConnectObjectAdded is emitted when object is added to manager.
func (v *DBusObjectManager) ConnectObjectAdded(f func(object DBusObjector)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "object-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectAdded), f)
}

// ConnectObjectRemoved is emitted when object is removed from manager.
func (v *DBusObjectManager) ConnectObjectRemoved(f func(object DBusObjector)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "object-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObjectManager_ConnectObjectRemoved), f)
}
