// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeDBusObjectManagerServer = coreglib.Type(girepository.MustFind("Gio", "DBusObjectManagerServer").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDBusObjectManagerServer, F: marshalDBusObjectManagerServer},
	})
}

// DBusObjectManagerServerOverrides contains methods that are overridable.
type DBusObjectManagerServerOverrides struct {
}

func defaultDBusObjectManagerServerOverrides(v *DBusObjectManagerServer) DBusObjectManagerServerOverrides {
	return DBusObjectManagerServerOverrides{}
}

// DBusObjectManagerServer is used to export BusObject instances using the
// standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface. For example, remote D-Bus clients can get all objects and
// properties in a single call. Additionally, any change in the object hierarchy
// is broadcast using signals. This means that D-Bus clients can keep caches up
// to date by only listening to D-Bus signals.
//
// The recommended path to export an object manager at is the path form of the
// well-known name of a D-Bus service, or below. For example, if a D-Bus service
// is available at the well-known name net.example.ExampleService1, the object
// manager should typically be exported at /net/example/ExampleService1, or
// below (to allow for multiple object managers in a service).
//
// It is supported, but not recommended, to export an object manager at the root
// path, /.
//
// See BusObjectManagerClient for the client-side code that is intended to be
// used with BusObjectManagerServer or any D-Bus object implementing the
// org.freedesktop.DBus.ObjectManager interface.
type DBusObjectManagerServer struct {
	_ [0]func() // equal guard
	*coreglib.Object

	DBusObjectManager
}

var (
	_ coreglib.Objector = (*DBusObjectManagerServer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DBusObjectManagerServer, *DBusObjectManagerServerClass, DBusObjectManagerServerOverrides](
		GTypeDBusObjectManagerServer,
		initDBusObjectManagerServerClass,
		wrapDBusObjectManagerServer,
		defaultDBusObjectManagerServerOverrides,
	)
}

func initDBusObjectManagerServerClass(gclass unsafe.Pointer, overrides DBusObjectManagerServerOverrides, classInitFunc func(*DBusObjectManagerServerClass)) {
	if classInitFunc != nil {
		class := (*DBusObjectManagerServerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDBusObjectManagerServer(obj *coreglib.Object) *DBusObjectManagerServer {
	return &DBusObjectManagerServer{
		Object: obj,
		DBusObjectManager: DBusObjectManager{
			Object: obj,
		},
	}
}

func marshalDBusObjectManagerServer(p uintptr) (interface{}, error) {
	return wrapDBusObjectManagerServer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusObjectManagerServerClass class structure for BusObjectManagerServer.
//
// An instance of this type is always passed by reference.
type DBusObjectManagerServerClass struct {
	*dBusObjectManagerServerClass
}

// dBusObjectManagerServerClass is the struct that's finalized.
type dBusObjectManagerServerClass struct {
	native unsafe.Pointer
}

var GIRInfoDBusObjectManagerServerClass = girepository.MustFind("Gio", "DBusObjectManagerServerClass")