// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GDBusInterface* _gotk4_gio2_DBusObjectIface_get_interface(void*, void*);
// extern GList* _gotk4_gio2_DBusObjectIface_get_interfaces(void*);
// extern gchar* _gotk4_gio2_DBusObjectIface_get_object_path(void*);
// extern void _gotk4_gio2_DBusObjectIface_interface_added(void*, void*);
// extern void _gotk4_gio2_DBusObjectIface_interface_removed(void*, void*);
// extern void _gotk4_gio2_DBusObject_ConnectInterfaceAdded(gpointer, void*, guintptr);
// extern void _gotk4_gio2_DBusObject_ConnectInterfaceRemoved(gpointer, void*, guintptr);
import "C"

// glib.Type values for gdbusobject.go.
var GTypeDBusObject = coreglib.Type(C.g_dbus_object_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDBusObject, F: marshalDBusObject},
	})
}

// DBusObjectOverrider contains methods that are overridable.
type DBusObjectOverrider interface {
	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	//
	// The function takes the following parameters:
	//
	//    - interfaceName d-Bus interface name.
	//
	// The function returns the following values:
	//
	//    - dBusInterface (optional): NULL if not found, otherwise a BusInterface
	//      that must be freed with g_object_unref().
	//
	Interface(interfaceName string) *DBusInterface
	// Interfaces gets the D-Bus interfaces associated with object.
	//
	// The function returns the following values:
	//
	//    - list of BusInterface instances. The returned list must be freed by
	//      g_list_free() after each element has been freed with
	//      g_object_unref().
	//
	Interfaces() []*DBusInterface
	// ObjectPath gets the object path for object.
	//
	// The function returns the following values:
	//
	//    - utf8: string owned by object. Do not free.
	//
	ObjectPath() string
	// The function takes the following parameters:
	//
	InterfaceAdded(interface_ DBusInterfacer)
	// The function takes the following parameters:
	//
	InterfaceRemoved(interface_ DBusInterfacer)
}

// DBusObject type is the base type for D-Bus objects on both the service side
// (see BusObjectSkeleton) and the client side (see BusObjectProxy). It is
// essentially just a container of interfaces.
//
// DBusObject wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DBusObject struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusObject)(nil)
)

// DBusObjector describes DBusObject's interface methods.
type DBusObjector interface {
	coreglib.Objector

	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) *DBusInterface
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []*DBusInterface
	// ObjectPath gets the object path for object.
	ObjectPath() string

	// Interface-added is emitted when interface is added to object.
	ConnectInterfaceAdded(func(iface DBusInterfacer)) coreglib.SignalHandle
	// Interface-removed is emitted when interface is removed from object.
	ConnectInterfaceRemoved(func(iface DBusInterfacer)) coreglib.SignalHandle
}

var _ DBusObjector = (*DBusObject)(nil)

func ifaceInitDBusObjector(gifacePtr, data C.gpointer) {
	iface := (*C.GDBusObjectIface)(unsafe.Pointer(gifacePtr))
	iface.get_interface = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_interface)
	iface.get_interfaces = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_interfaces)
	iface.get_object_path = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_get_object_path)
	iface.interface_added = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_interface_added)
	iface.interface_removed = (*[0]byte)(C._gotk4_gio2_DBusObjectIface_interface_removed)
}

//export _gotk4_gio2_DBusObjectIface_get_interface
func _gotk4_gio2_DBusObjectIface_get_interface(arg0 *C.void, arg1 *C.void) (cret *C.GDBusInterface) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interfaceName string // out

	_interfaceName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	dBusInterface := iface.Interface(_interfaceName)

	if dBusInterface != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(dBusInterface).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(dBusInterface).Native()))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectIface_get_interfaces
func _gotk4_gio2_DBusObjectIface_get_interfaces(arg0 *C.void) (cret *C.GList) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	list := iface.Interfaces()

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.void // out
		dst = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_gio2_DBusObjectIface_get_object_path
func _gotk4_gio2_DBusObjectIface_get_object_path(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	utf8 := iface.ObjectPath()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_DBusObjectIface_interface_added
func _gotk4_gio2_DBusObjectIface_interface_added(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceAdded(_interface_)
}

//export _gotk4_gio2_DBusObjectIface_interface_removed
func _gotk4_gio2_DBusObjectIface_interface_removed(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusObjectOverrider)

	var _interface_ DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_interface_ = rv
	}

	iface.InterfaceRemoved(_interface_)
}

func wrapDBusObject(obj *coreglib.Object) *DBusObject {
	return &DBusObject{
		Object: obj,
	}
}

func marshalDBusObject(p uintptr) (interface{}, error) {
	return wrapDBusObject(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusObject_ConnectInterfaceAdded
func _gotk4_gio2_DBusObject_ConnectInterfaceAdded(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(iface DBusInterfacer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iface DBusInterfacer))
	}

	var _iface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_iface)
}

// ConnectInterfaceAdded is emitted when interface is added to object.
func (object *DBusObject) ConnectInterfaceAdded(f func(iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(object, "interface-added", false, unsafe.Pointer(C._gotk4_gio2_DBusObject_ConnectInterfaceAdded), f)
}

//export _gotk4_gio2_DBusObject_ConnectInterfaceRemoved
func _gotk4_gio2_DBusObject_ConnectInterfaceRemoved(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(iface DBusInterfacer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iface DBusInterfacer))
	}

	var _iface DBusInterfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.DBusInterfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(DBusInterfacer)
			return ok
		})
		rv, ok := casted.(DBusInterfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusInterfacer")
		}
		_iface = rv
	}

	f(_iface)
}

// ConnectInterfaceRemoved is emitted when interface is removed from object.
func (object *DBusObject) ConnectInterfaceRemoved(f func(iface DBusInterfacer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(object, "interface-removed", false, unsafe.Pointer(C._gotk4_gio2_DBusObject_ConnectInterfaceRemoved), f)
}

// Interface gets the D-Bus interface with name interface_name associated with
// object, if any.
//
// The function takes the following parameters:
//
//    - interfaceName d-Bus interface name.
//
// The function returns the following values:
//
//    - dBusInterface (optional): NULL if not found, otherwise a BusInterface
//      that must be freed with g_object_unref().
//
func (object *DBusObject) Interface(interfaceName string) *DBusInterface {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface *DBusInterface // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusInterface = wrapDBusInterface(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _dBusInterface
}

// Interfaces gets the D-Bus interfaces associated with object.
//
// The function returns the following values:
//
//    - list of BusInterface instances. The returned list must be freed by
//      g_list_free() after each element has been freed with g_object_unref().
//
func (object *DBusObject) Interfaces() []*DBusInterface {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(object)

	var _list []*DBusInterface // out

	_list = make([]*DBusInterface, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *DBusInterface // out
		dst = wrapDBusInterface(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// ObjectPath gets the object path for object.
//
// The function returns the following values:
//
//    - utf8: string owned by object. Do not free.
//
func (object *DBusObject) ObjectPath() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(object)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
