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
// extern void _gotk4_gio2_DBusProxyClass_g_signal(void*, void*, void*, void*);
// extern void _gotk4_gio2_DBusProxy_ConnectGPropertiesChanged(gpointer, void*, void, guintptr);
// extern void _gotk4_gio2_DBusProxy_ConnectGSignal(gpointer, void*, void*, void*, guintptr);
import "C"

// glib.Type values for gdbusproxy.go.
var GTypeDBusProxy = coreglib.Type(C.g_dbus_proxy_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDBusProxy, F: marshalDBusProxy},
	})
}

// DBusProxyOverrider contains methods that are overridable.
type DBusProxyOverrider interface {
	// The function takes the following parameters:
	//
	//    - senderName
	//    - signalName
	//    - parameters
	//
	GSignal(senderName, signalName string, parameters *glib.Variant)
}

// DBusProxy is a base class used for proxies to access a D-Bus interface on a
// remote object. A BusProxy can be constructed for both well-known and unique
// names.
//
// By default, BusProxy will cache all properties (and listen to changes) of the
// remote object, and proxy all signals that get emitted. This behaviour can be
// changed by passing suitable BusProxyFlags when the proxy is created. If the
// proxy is for a well-known name, the property cache is flushed when the name
// owner vanishes and reloaded when a name owner appears.
//
// The unique name owner of the proxy's name is tracked and can be read from
// BusProxy:g-name-owner. Connect to the #GObject::notify signal to get notified
// of changes. Additionally, only signals and property changes emitted from the
// current name owner are considered and calls are always sent to the current
// name owner. This avoids a number of race conditions when the name is lost by
// one owner and claimed by another. However, if no name owner currently exists,
// then calls will be sent to the well-known name which may result in the
// message bus launching an owner (unless G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START
// is set).
//
// The generic BusProxy::g-properties-changed and BusProxy::g-signal signals are
// not very convenient to work with. Therefore, the recommended way of working
// with proxies is to subclass BusProxy, and have more natural properties and
// signals in your derived class. This [example][gdbus-example-gdbus-codegen]
// shows how this can easily be done using the [gdbus-codegen][gdbus-codegen]
// tool.
//
// A BusProxy instance can be used from multiple threads but note that all
// signals (e.g. BusProxy::g-signal, BusProxy::g-properties-changed and
// #GObject::notify) are emitted in the [thread-default main
// context][g-main-context-push-thread-default] of the thread where the instance
// was constructed.
//
// An example using a proxy for a well-known name can be found in
// gdbus-example-watch-proxy.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-watch-proxy.c).
type DBusProxy struct {
	_ [0]func() // equal guard
	*coreglib.Object

	AsyncInitable
	DBusInterface
	Initable
}

var (
	_ coreglib.Objector = (*DBusProxy)(nil)
)

func classInitDBusProxier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GDBusProxyClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GDBusProxyClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		GSignal(senderName, signalName string, parameters *glib.Variant)
	}); ok {
		pclass.g_signal = (*[0]byte)(C._gotk4_gio2_DBusProxyClass_g_signal)
	}
}

//export _gotk4_gio2_DBusProxyClass_g_signal
func _gotk4_gio2_DBusProxyClass_g_signal(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		GSignal(senderName, signalName string, parameters *glib.Variant)
	})

	var _senderName string        // out
	var _signalName string        // out
	var _parameters *glib.Variant // out

	_senderName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_signalName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_parameters = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_parameters)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	iface.GSignal(_senderName, _signalName, _parameters)
}

func wrapDBusProxy(obj *coreglib.Object) *DBusProxy {
	return &DBusProxy{
		Object: obj,
		AsyncInitable: AsyncInitable{
			Object: obj,
		},
		DBusInterface: DBusInterface{
			Object: obj,
		},
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusProxy(p uintptr) (interface{}, error) {
	return wrapDBusProxy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusProxy_ConnectGPropertiesChanged
func _gotk4_gio2_DBusProxy_ConnectGPropertiesChanged(arg0 C.gpointer, arg1 *C.void, arg2 C.void, arg3 C.guintptr) {
	var f func(changedProperties *glib.Variant, invalidatedProperties []string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(changedProperties *glib.Variant, invalidatedProperties []string))
	}

	var _changedProperties *glib.Variant // out
	var _invalidatedProperties []string  // out

	_changedProperties = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_variant_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_changedProperties)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	{
		var i int
		var z *C.void
		for p := arg2; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(arg2, i)
		_invalidatedProperties = make([]string, i)
		for i := range src {
			_invalidatedProperties[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	f(_changedProperties, _invalidatedProperties)
}

// ConnectGPropertiesChanged is emitted when one or more D-Bus properties on
// proxy changes. The local cache has already been updated when this signal
// fires. Note that both changed_properties and invalidated_properties are
// guaranteed to never be NULL (either may be empty though).
//
// If the proxy has the flag G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES set,
// then invalidated_properties will always be empty.
//
// This signal corresponds to the PropertiesChanged D-Bus signal on the
// org.freedesktop.DBus.Properties interface.
func (proxy *DBusProxy) ConnectGPropertiesChanged(f func(changedProperties *glib.Variant, invalidatedProperties []string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(proxy, "g-properties-changed", false, unsafe.Pointer(C._gotk4_gio2_DBusProxy_ConnectGPropertiesChanged), f)
}

//export _gotk4_gio2_DBusProxy_ConnectGSignal
func _gotk4_gio2_DBusProxy_ConnectGSignal(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.guintptr) {
	var f func(senderName, signalName string, parameters *glib.Variant)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(senderName, signalName string, parameters *glib.Variant))
	}

	var _senderName string        // out
	var _signalName string        // out
	var _parameters *glib.Variant // out

	if arg1 != nil {
		_senderName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	}
	_signalName = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_parameters = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_parameters)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	f(_senderName, _signalName, _parameters)
}

// ConnectGSignal is emitted when a signal from the remote object and interface
// that proxy is for, has been received.
func (proxy *DBusProxy) ConnectGSignal(f func(senderName, signalName string, parameters *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(proxy, "g-signal", false, unsafe.Pointer(C._gotk4_gio2_DBusProxy_ConnectGSignal), f)
}

// NewDBusProxyFinish finishes creating a BusProxy.
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback function passed to
//      g_dbus_proxy_new().
//
// The function returns the following values:
//
//    - dBusProxy or NULL if error is set. Free with g_object_unref().
//
func NewDBusProxyFinish(res AsyncResulter) (*DBusProxy, error) {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("new_DBusProxy_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(res)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusProxy, _goerr
}

// NewDBusProxyForBusFinish finishes creating a BusProxy.
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback function passed to
//      g_dbus_proxy_new_for_bus().
//
// The function returns the following values:
//
//    - dBusProxy or NULL if error is set. Free with g_object_unref().
//
func NewDBusProxyForBusFinish(res AsyncResulter) (*DBusProxy, error) {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("new_DBusProxy_for_bus_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(res)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusProxy, _goerr
}

// CallFinish finishes an operation started with g_dbus_proxy_call().
//
// The function takes the following parameters:
//
//    - res obtained from the ReadyCallback passed to g_dbus_proxy_call().
//
// The function returns the following values:
//
//    - variant: NULL if error is set. Otherwise a #GVariant tuple with return
//      values. Free with g_variant_unref().
//
func (proxy *DBusProxy) CallFinish(res AsyncResulter) (*glib.Variant, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("call_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(res)

	var _variant *glib.Variant // out
	var _goerr error           // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _variant, _goerr
}

// CachedProperty looks up the value for a property from the cache. This call
// does no blocking IO.
//
// If proxy has an expected interface (see BusProxy:g-interface-info) and
// property_name is referenced by it, then value is checked against the type of
// the property.
//
// The function takes the following parameters:
//
//    - propertyName: property name.
//
// The function returns the following values:
//
//    - variant (optional): reference to the #GVariant instance that holds the
//      value for property_name or NULL if the value is not in the cache. The
//      returned reference must be freed with g_variant_unref().
//
func (proxy *DBusProxy) CachedProperty(propertyName string) *glib.Variant {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_cached_property", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(propertyName)

	var _variant *glib.Variant // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _variant
}

// CachedPropertyNames gets the names of all cached properties on proxy.
//
// The function returns the following values:
//
//    - utf8s (optional): a NULL-terminated array of strings or NULL if proxy has
//      no cached properties. Free the returned array with g_strfreev().
//
func (proxy *DBusProxy) CachedPropertyNames() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_cached_property_names", _args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _utf8s []string // out

	if *(***C.void)(unsafe.Pointer(&_cret)) != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.void
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// Connection gets the connection proxy is for.
//
// The function returns the following values:
//
//    - dBusConnection owned by proxy. Do not free.
//
func (proxy *DBusProxy) Connection() *DBusConnection {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_connection", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// DefaultTimeout gets the timeout to use if -1 (specifying default timeout) is
// passed as timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
//
// The function returns the following values:
//
//    - gint: timeout to use for proxy.
//
func (proxy *DBusProxy) DefaultTimeout() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_default_timeout", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the interface
// that proxy conforms to. See the BusProxy:g-interface-info property for more
// details.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo (optional) or NULL. Do not unref the returned object,
//      it is owned by proxy.
//
func (proxy *DBusProxy) InterfaceInfo() *DBusInterfaceInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_interface_info", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_interface_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _dBusInterfaceInfo
}

// InterfaceName gets the D-Bus interface name proxy is for.
//
// The function returns the following values:
//
//    - utf8: string owned by proxy. Do not free.
//
func (proxy *DBusProxy) InterfaceName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_interface_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the name that proxy was constructed for.
//
// The function returns the following values:
//
//    - utf8: string owned by proxy. Do not free.
//
func (proxy *DBusProxy) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NameOwner: unique name that owns the name that proxy is for or NULL if no-one
// currently owns that name. You may connect to the #GObject::notify signal to
// track changes to the BusProxy:g-name-owner property.
//
// The function returns the following values:
//
//    - utf8 (optional): name owner or NULL if no name owner exists. Free with
//      g_free().
//
func (proxy *DBusProxy) NameOwner() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_name_owner", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// ObjectPath gets the object path proxy is for.
//
// The function returns the following values:
//
//    - utf8: string owned by proxy. Do not free.
//
func (proxy *DBusProxy) ObjectPath() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_gret := girepository.MustFind("Gio", "DBusProxy").InvokeMethod("get_object_path", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetCachedProperty: if value is not NULL, sets the cached value for the
// property with name property_name to the value in value.
//
// If value is NULL, then the cached value is removed from the property cache.
//
// If proxy has an expected interface (see BusProxy:g-interface-info) and
// property_name is referenced by it, then value is checked against the type of
// the property.
//
// If the value #GVariant is floating, it is consumed. This allows convenient
// 'inline' use of g_variant_new(), e.g.
//
//    g_dbus_proxy_set_cached_property (proxy,
//                                      "SomeProperty",
//                                      g_variant_new ("(si)",
//                                                    "A String",
//                                                    42));
//
// Normally you will not need to use this method since proxy is tracking changes
// using the org.freedesktop.DBus.Properties.PropertiesChanged D-Bus signal.
// However, for performance reasons an object may decide to not use this signal
// for some properties and instead use a proprietary out-of-band mechanism to
// transmit changes.
//
// As a concrete example, consider an object with a property
// ChatroomParticipants which is an array of strings. Instead of transmitting
// the same (long) array every time the property changes, it is more efficient
// to only transmit the delta using e.g. signals
// ChatroomParticipantJoined(String name) and ChatroomParticipantParted(String
// name).
//
// The function takes the following parameters:
//
//    - propertyName: property name.
//    - value (optional): value for the property or NULL to remove it from the
//      cache.
//
func (proxy *DBusProxy) SetCachedProperty(propertyName string, value *glib.Variant) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_args[1]))
	if value != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(value)))
	}

	girepository.MustFind("Gio", "DBusProxy").InvokeMethod("set_cached_property", _args[:], nil)

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(propertyName)
	runtime.KeepAlive(value)
}

// SetDefaultTimeout sets the timeout to use if -1 (specifying default timeout)
// is passed as timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
//
// The function takes the following parameters:
//
//    - timeoutMsec: timeout in milliseconds.
//
func (proxy *DBusProxy) SetDefaultTimeout(timeoutMsec int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(timeoutMsec)

	girepository.MustFind("Gio", "DBusProxy").InvokeMethod("set_default_timeout", _args[:], nil)

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(timeoutMsec)
}

// SetInterfaceInfo: ensure that interactions with proxy conform to the given
// interface. See the BusProxy:g-interface-info property for more details.
//
// The function takes the following parameters:
//
//    - info (optional): minimum interface this proxy conforms to or NULL to
//      unset.
//
func (proxy *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	if info != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))
	}

	girepository.MustFind("Gio", "DBusProxy").InvokeMethod("set_interface_info", _args[:], nil)

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(info)
}
