// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern GDBusInterfaceInfo* _gotk4_gio2_DBusInterfaceSkeletonClass_get_info(void*);
// extern GVariant* _gotk4_gio2_DBusInterfaceSkeletonClass_get_properties(void*);
// extern gboolean _gotk4_gio2_DBusInterfaceSkeletonClass_g_authorize_method(void*, void*);
// extern gboolean _gotk4_gio2_DBusInterfaceSkeleton_ConnectGAuthorizeMethod(gpointer, void*, guintptr);
// extern void _gotk4_gio2_DBusInterfaceSkeletonClass_flush(void*);
import "C"

// GTypeDBusInterfaceSkeleton returns the GType for the type DBusInterfaceSkeleton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusInterfaceSkeleton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusInterfaceSkeleton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusInterfaceSkeleton)
	return gtype
}

// DBusInterfaceSkeletonOverrider contains methods that are overridable.
type DBusInterfaceSkeletonOverrider interface {
	// Flush: if interface_ has outstanding changes, request for these changes
	// to be emitted immediately.
	//
	// For example, an exported D-Bus interface may queue up property changes
	// and emit the org.freedesktop.DBus.Properties.PropertiesChanged signal
	// later (e.g. in an idle handler). This technique is useful for collapsing
	// multiple property changes into one.
	Flush()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	GAuthorizeMethod(invocation *DBusMethodInvocation) bool
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	//
	// The function returns the following values:
	//
	//    - dBusInterfaceInfo (never NULL). Do not free.
	//
	Info() *DBusInterfaceInfo
	// Properties gets all D-Bus properties for interface_.
	//
	// The function returns the following values:
	//
	//    - variant of type ['a{sv}'][G-VARIANT-TYPE-VARDICT:CAPS]. Free with
	//      g_variant_unref().
	//
	Properties() *glib.Variant
}

// DBusInterfaceSkeleton: abstract base class for D-Bus interfaces on the
// service side.
type DBusInterfaceSkeleton struct {
	_ [0]func() // equal guard
	*coreglib.Object

	DBusInterface
}

var (
	_ coreglib.Objector = (*DBusInterfaceSkeleton)(nil)
)

// DBusInterfaceSkeletonner describes types inherited from class DBusInterfaceSkeleton.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DBusInterfaceSkeletonner interface {
	coreglib.Objector
	baseDBusInterfaceSkeleton() *DBusInterfaceSkeleton
}

var _ DBusInterfaceSkeletonner = (*DBusInterfaceSkeleton)(nil)

func classInitDBusInterfaceSkeletonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "DBusInterfaceSkeletonClass")

	if _, ok := goval.(interface{ Flush() }); ok {
		o := pclass.StructFieldOffset("flush")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceSkeletonClass_flush)
	}

	if _, ok := goval.(interface {
		GAuthorizeMethod(invocation *DBusMethodInvocation) bool
	}); ok {
		o := pclass.StructFieldOffset("g_authorize_method")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceSkeletonClass_g_authorize_method)
	}

	if _, ok := goval.(interface{ Info() *DBusInterfaceInfo }); ok {
		o := pclass.StructFieldOffset("get_info")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceSkeletonClass_get_info)
	}

	if _, ok := goval.(interface{ Properties() *glib.Variant }); ok {
		o := pclass.StructFieldOffset("get_properties")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceSkeletonClass_get_properties)
	}
}

//export _gotk4_gio2_DBusInterfaceSkeletonClass_flush
func _gotk4_gio2_DBusInterfaceSkeletonClass_flush(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Flush() })

	iface.Flush()
}

//export _gotk4_gio2_DBusInterfaceSkeletonClass_g_authorize_method
func _gotk4_gio2_DBusInterfaceSkeletonClass_g_authorize_method(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		GAuthorizeMethod(invocation *DBusMethodInvocation) bool
	})

	var _invocation *DBusMethodInvocation // out

	_invocation = wrapDBusMethodInvocation(coreglib.Take(unsafe.Pointer(arg1)))

	ok := iface.GAuthorizeMethod(_invocation)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_DBusInterfaceSkeletonClass_get_info
func _gotk4_gio2_DBusInterfaceSkeletonClass_get_info(arg0 *C.void) (cret *C.GDBusInterfaceInfo) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Info() *DBusInterfaceInfo })

	dBusInterfaceInfo := iface.Info()

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(dBusInterfaceInfo)))

	return cret
}

//export _gotk4_gio2_DBusInterfaceSkeletonClass_get_properties
func _gotk4_gio2_DBusInterfaceSkeletonClass_get_properties(arg0 *C.void) (cret *C.GVariant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Properties() *glib.Variant })

	variant := iface.Properties()

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))

	return cret
}

func wrapDBusInterfaceSkeleton(obj *coreglib.Object) *DBusInterfaceSkeleton {
	return &DBusInterfaceSkeleton{
		Object: obj,
		DBusInterface: DBusInterface{
			Object: obj,
		},
	}
}

func marshalDBusInterfaceSkeleton(p uintptr) (interface{}, error) {
	return wrapDBusInterfaceSkeleton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (interface_ *DBusInterfaceSkeleton) baseDBusInterfaceSkeleton() *DBusInterfaceSkeleton {
	return interface_
}

// BaseDBusInterfaceSkeleton returns the underlying base object.
func BaseDBusInterfaceSkeleton(obj DBusInterfaceSkeletonner) *DBusInterfaceSkeleton {
	return obj.baseDBusInterfaceSkeleton()
}

//export _gotk4_gio2_DBusInterfaceSkeleton_ConnectGAuthorizeMethod
func _gotk4_gio2_DBusInterfaceSkeleton_ConnectGAuthorizeMethod(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) (cret C.gboolean) {
	var f func(invocation *DBusMethodInvocation) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(invocation *DBusMethodInvocation) (ok bool))
	}

	var _invocation *DBusMethodInvocation // out

	_invocation = wrapDBusMethodInvocation(coreglib.Take(unsafe.Pointer(arg1)))

	ok := f(_invocation)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectGAuthorizeMethod is emitted when a method is invoked by a remote
// caller and used to determine if the method call is authorized.
//
// Note that this signal is emitted in a thread dedicated to handling the method
// call so handlers are allowed to perform blocking IO. This means that it is
// appropriate to call e.g. polkit_authority_check_authorization_sync()
// (http://hal.freedesktop.org/docs/polkit/PolkitAuthority.html#polkit-authority-check-authorization-sync)
// with the POLKIT_CHECK_AUTHORIZATION_FLAGS_ALLOW_USER_INTERACTION
// (http://hal.freedesktop.org/docs/polkit/PolkitAuthority.htmlLKIT-CHECK-AUTHORIZATION-FLAGS-ALLOW-USER-INTERACTION:CAPS)
// flag set.
//
// If FALSE is returned then no further handlers are run and the signal handler
// must take a reference to invocation and finish handling the call (e.g. return
// an error via g_dbus_method_invocation_return_error()).
//
// Otherwise, if TRUE is returned, signal emission continues. If no handlers
// return FALSE, then the method is dispatched. If interface has an enclosing
// BusObjectSkeleton, then the BusObjectSkeleton::authorize-method signal
// handlers run before the handlers for this signal.
//
// The default class handler just returns TRUE.
//
// Please note that the common case is optimized: if no signals handlers are
// connected and the default class handler isn't overridden (for both interface
// and the enclosing BusObjectSkeleton, if any) and BusInterfaceSkeleton:g-flags
// does not have the
// G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD flags
// set, no dedicated thread is ever used and the call will be handled in the
// same thread as the object that interface belongs to was exported in.
func (interface_ *DBusInterfaceSkeleton) ConnectGAuthorizeMethod(f func(invocation *DBusMethodInvocation) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(interface_, "g-authorize-method", false, unsafe.Pointer(C._gotk4_gio2_DBusInterfaceSkeleton_ConnectGAuthorizeMethod), f)
}

// Export exports interface_ at object_path on connection.
//
// This can be called multiple times to export the same interface_ onto multiple
// connections however the object_path provided must be the same for all
// connections.
//
// Use g_dbus_interface_skeleton_unexport() to unexport the object.
//
// The function takes the following parameters:
//
//    - connection to export interface_ on.
//    - objectPath: path to export the interface at.
//
func (interface_ *DBusInterfaceSkeleton) Export(connection *DBusConnection, objectPath string) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_args[2]))

	girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("export", _args[:], nil)

	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(objectPath)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Flush: if interface_ has outstanding changes, request for these changes to be
// emitted immediately.
//
// For example, an exported D-Bus interface may queue up property changes and
// emit the org.freedesktop.DBus.Properties.PropertiesChanged signal later (e.g.
// in an idle handler). This technique is useful for collapsing multiple
// property changes into one.
func (interface_ *DBusInterfaceSkeleton) Flush() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("flush", _args[:], nil)

	runtime.KeepAlive(interface_)
}

// Connection gets the first connection that interface_ is exported on, if any.
//
// The function returns the following values:
//
//    - dBusConnection (optional) or NULL if interface_ is not exported anywhere.
//      Do not free, the object belongs to interface_.
//
func (interface_ *DBusInterfaceSkeleton) Connection() *DBusConnection {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("get_connection", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _dBusConnection *DBusConnection // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusConnection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _dBusConnection
}

// Connections gets a list of the connections that interface_ is exported on.
//
// The function returns the following values:
//
//    - list of all the connections that interface_ is exported on. The returned
//      list should be freed with g_list_free() after each element has been freed
//      with g_object_unref().
//
func (interface_ *DBusInterfaceSkeleton) Connections() []*DBusConnection {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("get_connections", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _list []*DBusConnection // out

	_list = make([]*DBusConnection, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *DBusConnection // out
		dst = wrapDBusConnection(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by interface_.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo (never NULL). Do not free.
//
func (interface_ *DBusInterfaceSkeleton) Info() *DBusInterfaceInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("get_info", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _dBusInterfaceInfo
}

// ObjectPath gets the object path that interface_ is exported on, if any.
//
// The function returns the following values:
//
//    - utf8 (optional): string owned by interface_ or NULL if interface_ is not
//      exported anywhere. Do not free, the string belongs to interface_.
//
func (interface_ *DBusInterfaceSkeleton) ObjectPath() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("get_object_path", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Properties gets all D-Bus properties for interface_.
//
// The function returns the following values:
//
//    - variant of type ['a{sv}'][G-VARIANT-TYPE-VARDICT:CAPS]. Free with
//      g_variant_unref().
//
func (interface_ *DBusInterfaceSkeleton) Properties() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("get_properties", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _variant
}

// HasConnection checks if interface_ is exported on connection.
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//
// The function returns the following values:
//
//    - ok: TRUE if interface_ is exported on connection, FALSE otherwise.
//
func (interface_ *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))

	_gret := girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("has_connection", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Unexport stops exporting interface_ on all connections it is exported on.
//
// To unexport interface_ from only a single connection, use
// g_dbus_interface_skeleton_unexport_from_connection().
func (interface_ *DBusInterfaceSkeleton) Unexport() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("unexport", _args[:], nil)

	runtime.KeepAlive(interface_)
}

// UnexportFromConnection stops exporting interface_ on connection.
//
// To stop exporting on all connections the interface is exported on, use
// g_dbus_interface_skeleton_unexport().
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//
func (interface_ *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))

	girepository.MustFind("Gio", "DBusInterfaceSkeleton").InvokeMethod("unexport_from_connection", _args[:], nil)

	runtime.KeepAlive(interface_)
	runtime.KeepAlive(connection)
}
