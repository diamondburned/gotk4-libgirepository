// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_proxy_get_type()), F: marshalDBusProxier},
	})
}

// DBusProxyOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusProxyOverrider interface {
	GSignal(senderName string, signalName string, parameters *glib.Variant)
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
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-watch-proxy.c)
type DBusProxy struct {
	*externglib.Object

	AsyncInitable
	DBusInterface
	Initable
}

var _ gextras.Nativer = (*DBusProxy)(nil)

func wrapDBusProxy(obj *externglib.Object) *DBusProxy {
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

func marshalDBusProxier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusProxy(obj), nil
}

// NewDBusProxyFinish finishes creating a BusProxy.
func NewDBusProxyFinish(res AsyncResulter) (*DBusProxy, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GDBusProxy   // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_dbus_proxy_new_finish(_arg1, &_cerr)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// NewDBusProxyForBusFinish finishes creating a BusProxy.
func NewDBusProxyForBusFinish(res AsyncResulter) (*DBusProxy, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GDBusProxy   // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_dbus_proxy_new_for_bus_finish(_arg1, &_cerr)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// NewDBusProxyForBusSync: like g_dbus_proxy_new_sync() but takes a Type instead
// of a BusConnection.
//
// BusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxyForBusSync(ctx context.Context, busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string) (*DBusProxy, error) {
	var _arg7 *C.GCancellable       // out
	var _arg1 C.GBusType            // out
	var _arg2 C.GDBusProxyFlags     // out
	var _arg3 *C.GDBusInterfaceInfo // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.gchar              // out
	var _arg6 *C.gchar              // out
	var _cret *C.GDBusProxy         // in
	var _cerr *C.GError             // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)
	_arg2 = C.GDBusProxyFlags(flags)
	_arg3 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))

	_cret = C.g_dbus_proxy_new_for_bus_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, &_cerr)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// NewDBusProxySync creates a proxy for accessing interface_name on the remote
// object at object_path owned by name at connection and synchronously loads
// D-Bus properties unless the G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is
// used.
//
// If the G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets
// up match rules for signals. Connect to the BusProxy::g-signal signal to
// handle signals from the remote object.
//
// If both G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES and
// G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS are set, this constructor is
// guaranteed to return immediately without blocking.
//
// If name is a well-known name and the G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and
// G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION flags aren't set and no
// name owner currently exists, the message bus will be requested to launch a
// name owner for the name.
//
// This is a synchronous failable constructor. See g_dbus_proxy_new() and
// g_dbus_proxy_new_finish() for the asynchronous version.
//
// BusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxySync(ctx context.Context, connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string) (*DBusProxy, error) {
	var _arg7 *C.GCancellable       // out
	var _arg1 *C.GDBusConnection    // out
	var _arg2 C.GDBusProxyFlags     // out
	var _arg3 *C.GDBusInterfaceInfo // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.gchar              // out
	var _arg6 *C.gchar              // out
	var _cret *C.GDBusProxy         // in
	var _cerr *C.GError             // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = C.GDBusProxyFlags(flags)
	_arg3 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))

	_cret = C.g_dbus_proxy_new_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, &_cerr)

	var _dBusProxy *DBusProxy // out
	var _goerr error          // out

	_dBusProxy = wrapDBusProxy(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// Call: asynchronously invokes the method_name method on proxy.
//
// If method_name contains any dots, then name is split into interface and
// method name parts. This allows using proxy for invoking methods on other
// interfaces.
//
// If the BusConnection associated with proxy is closed then the operation will
// fail with G_IO_ERROR_CLOSED. If cancellable is canceled, the operation will
// fail with G_IO_ERROR_CANCELLED. If parameters contains a value not compatible
// with the D-Bus protocol, the operation fails with
// G_IO_ERROR_INVALID_ARGUMENT.
//
// If the parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
//
//    g_dbus_proxy_call (proxy,
//                       "TwoStrings",
//                       g_variant_new ("(ss)",
//                                      "Thing One",
//                                      "Thing Two"),
//                       G_DBUS_CALL_FLAGS_NONE,
//                       -1,
//                       NULL,
//                       (GAsyncReadyCallback) two_strings_done,
//                       &data);
//
// If proxy has an expected interface (see BusProxy:g-interface-info) and
// method_name is referenced by it, then the return value is checked against the
// return type.
//
// This is an asynchronous method. When the operation is finished, callback will
// be invoked in the [thread-default main
// context][g-main-context-push-thread-default] of the thread you are calling
// this method from. You can then call g_dbus_proxy_call_finish() to get the
// result of the operation. See g_dbus_proxy_call_sync() for the synchronous
// version of this method.
//
// If callback is NULL then the D-Bus method call message will be sent with the
// G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
func (proxy *DBusProxy) Call(ctx context.Context, methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, callback AsyncReadyCallback) {
	var _arg0 *C.GDBusProxy         // out
	var _arg5 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 C.GDBusCallFlags      // out
	var _arg4 C.gint                // out
	var _arg6 C.GAsyncReadyCallback // out
	var _arg7 C.gpointer

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(methodName)))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameters)))
	_arg3 = C.GDBusCallFlags(flags)
	_arg4 = C.gint(timeoutMsec)
	_arg6 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg7 = C.gpointer(gbox.AssignOnce(callback))

	C.g_dbus_proxy_call(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// CallFinish finishes an operation started with g_dbus_proxy_call().
func (proxy *DBusProxy) CallFinish(res AsyncResulter) (*glib.Variant, error) {
	var _arg0 *C.GDBusProxy   // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GVariant     // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_dbus_proxy_call_finish(_arg0, _arg1, &_cerr)

	var _variant *glib.Variant // out
	var _goerr error           // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _variant, _goerr
}

// CallSync: synchronously invokes the method_name method on proxy.
//
// If method_name contains any dots, then name is split into interface and
// method name parts. This allows using proxy for invoking methods on other
// interfaces.
//
// If the BusConnection associated with proxy is disconnected then the operation
// will fail with G_IO_ERROR_CLOSED. If cancellable is canceled, the operation
// will fail with G_IO_ERROR_CANCELLED. If parameters contains a value not
// compatible with the D-Bus protocol, the operation fails with
// G_IO_ERROR_INVALID_ARGUMENT.
//
// If the parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
//
//    g_dbus_proxy_call_sync (proxy,
//                            "TwoStrings",
//                            g_variant_new ("(ss)",
//                                           "Thing One",
//                                           "Thing Two"),
//                            G_DBUS_CALL_FLAGS_NONE,
//                            -1,
//                            NULL,
//                            &error);
//
// The calling thread is blocked until a reply is received. See
// g_dbus_proxy_call() for the asynchronous version of this method.
//
// If proxy has an expected interface (see BusProxy:g-interface-info) and
// method_name is referenced by it, then the return value is checked against the
// return type.
func (proxy *DBusProxy) CallSync(ctx context.Context, methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int) (*glib.Variant, error) {
	var _arg0 *C.GDBusProxy    // out
	var _arg5 *C.GCancellable  // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariant      // out
	var _arg3 C.GDBusCallFlags // out
	var _arg4 C.gint           // out
	var _cret *C.GVariant      // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(methodName)))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameters)))
	_arg3 = C.GDBusCallFlags(flags)
	_arg4 = C.gint(timeoutMsec)

	_cret = C.g_dbus_proxy_call_sync(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)

	var _variant *glib.Variant // out
	var _goerr error           // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _variant, _goerr
}

// CallWithUnixFdList: like g_dbus_proxy_call() but also takes a FDList object.
//
// This method is only available on UNIX.
func (proxy *DBusProxy) CallWithUnixFdList(ctx context.Context, methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, fdList *UnixFDList, callback AsyncReadyCallback) {
	var _arg0 *C.GDBusProxy         // out
	var _arg6 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 C.GDBusCallFlags      // out
	var _arg4 C.gint                // out
	var _arg5 *C.GUnixFDList        // out
	var _arg7 C.GAsyncReadyCallback // out
	var _arg8 C.gpointer

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(methodName)))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameters)))
	_arg3 = C.GDBusCallFlags(flags)
	_arg4 = C.gint(timeoutMsec)
	_arg5 = (*C.GUnixFDList)(unsafe.Pointer(fdList.Native()))
	_arg7 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg8 = C.gpointer(gbox.AssignOnce(callback))

	C.g_dbus_proxy_call_with_unix_fd_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// CallWithUnixFdListFinish finishes an operation started with
// g_dbus_proxy_call_with_unix_fd_list().
func (proxy *DBusProxy) CallWithUnixFdListFinish(res AsyncResulter) (*UnixFDList, *glib.Variant, error) {
	var _arg0 *C.GDBusProxy   // out
	var _arg1 *C.GUnixFDList  // in
	var _arg2 *C.GAsyncResult // out
	var _cret *C.GVariant     // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg2 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_dbus_proxy_call_with_unix_fd_list_finish(_arg0, &_arg1, _arg2, &_cerr)

	var _outFdList *UnixFDList // out
	var _variant *glib.Variant // out
	var _goerr error           // out

	_outFdList = wrapUnixFDList(externglib.AssumeOwnership(unsafe.Pointer(_arg1)))
	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outFdList, _variant, _goerr
}

// CallWithUnixFdListSync: like g_dbus_proxy_call_sync() but also takes and
// returns FDList objects.
//
// This method is only available on UNIX.
func (proxy *DBusProxy) CallWithUnixFdListSync(ctx context.Context, methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, fdList *UnixFDList) (*UnixFDList, *glib.Variant, error) {
	var _arg0 *C.GDBusProxy    // out
	var _arg7 *C.GCancellable  // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariant      // out
	var _arg3 C.GDBusCallFlags // out
	var _arg4 C.gint           // out
	var _arg5 *C.GUnixFDList   // out
	var _arg6 *C.GUnixFDList   // in
	var _cret *C.GVariant      // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(methodName)))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameters)))
	_arg3 = C.GDBusCallFlags(flags)
	_arg4 = C.gint(timeoutMsec)
	_arg5 = (*C.GUnixFDList)(unsafe.Pointer(fdList.Native()))

	_cret = C.g_dbus_proxy_call_with_unix_fd_list_sync(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_arg6, _arg7, &_cerr)

	var _outFdList *UnixFDList // out
	var _variant *glib.Variant // out
	var _goerr error           // out

	_outFdList = wrapUnixFDList(externglib.AssumeOwnership(unsafe.Pointer(_arg6)))
	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outFdList, _variant, _goerr
}

// CachedProperty looks up the value for a property from the cache. This call
// does no blocking IO.
//
// If proxy has an expected interface (see BusProxy:g-interface-info) and
// property_name is referenced by it, then value is checked against the type of
// the property.
func (proxy *DBusProxy) CachedProperty(propertyName string) *glib.Variant {
	var _arg0 *C.GDBusProxy // out
	var _arg1 *C.gchar      // out
	var _cret *C.GVariant   // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))

	_cret = C.g_dbus_proxy_get_cached_property(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variant
}

// CachedPropertyNames gets the names of all cached properties on proxy.
func (proxy *DBusProxy) CachedPropertyNames() []string {
	var _arg0 *C.GDBusProxy // out
	var _cret **C.gchar     // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_cached_property_names(_arg0)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Connection gets the connection proxy is for.
func (proxy *DBusProxy) Connection() *DBusConnection {
	var _arg0 *C.GDBusProxy      // out
	var _cret *C.GDBusConnection // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_connection(_arg0)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(externglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}

// DefaultTimeout gets the timeout to use if -1 (specifying default timeout) is
// passed as timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (proxy *DBusProxy) DefaultTimeout() int {
	var _arg0 *C.GDBusProxy // out
	var _cret C.gint        // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_default_timeout(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Flags gets the flags that proxy was constructed with.
func (proxy *DBusProxy) Flags() DBusProxyFlags {
	var _arg0 *C.GDBusProxy     // out
	var _cret C.GDBusProxyFlags // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_flags(_arg0)

	var _dBusProxyFlags DBusProxyFlags // out

	_dBusProxyFlags = DBusProxyFlags(_cret)

	return _dBusProxyFlags
}

// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the interface
// that proxy conforms to. See the BusProxy:g-interface-info property for more
// details.
func (proxy *DBusProxy) InterfaceInfo() *DBusInterfaceInfo {
	var _arg0 *C.GDBusProxy         // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_interface_info(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _dBusInterfaceInfo
}

// InterfaceName gets the D-Bus interface name proxy is for.
func (proxy *DBusProxy) InterfaceName() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_interface_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the name that proxy was constructed for.
func (proxy *DBusProxy) Name() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NameOwner: unique name that owns the name that proxy is for or NULL if no-one
// currently owns that name. You may connect to the #GObject::notify signal to
// track changes to the BusProxy:g-name-owner property.
func (proxy *DBusProxy) NameOwner() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_name_owner(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ObjectPath gets the object path proxy is for.
func (proxy *DBusProxy) ObjectPath() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_dbus_proxy_get_object_path(_arg0)

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
func (proxy *DBusProxy) SetCachedProperty(propertyName string, value *glib.Variant) {
	var _arg0 *C.GDBusProxy // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GVariant   // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_dbus_proxy_set_cached_property(_arg0, _arg1, _arg2)
}

// SetDefaultTimeout sets the timeout to use if -1 (specifying default timeout)
// is passed as timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (proxy *DBusProxy) SetDefaultTimeout(timeoutMsec int) {
	var _arg0 *C.GDBusProxy // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = C.gint(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout(_arg0, _arg1)
}

// SetInterfaceInfo: ensure that interactions with proxy conform to the given
// interface. See the BusProxy:g-interface-info property for more details.
func (proxy *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	var _arg0 *C.GDBusProxy         // out
	var _arg1 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))

	C.g_dbus_proxy_set_interface_info(_arg0, _arg1)
}

// NewDBusProxy creates a proxy for accessing interface_name on the remote
// object at object_path owned by name at connection and asynchronously loads
// D-Bus properties unless the G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is
// used. Connect to the BusProxy::g-properties-changed signal to get notified
// about property changes.
//
// If the G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets
// up match rules for signals. Connect to the BusProxy::g-signal signal to
// handle signals from the remote object.
//
// If both G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES and
// G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS are set, this constructor is
// guaranteed to complete immediately without blocking.
//
// If name is a well-known name and the G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and
// G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION flags aren't set and no
// name owner currently exists, the message bus will be requested to launch a
// name owner for the name.
//
// This is a failable asynchronous constructor - when the proxy is ready,
// callback will be invoked and you can use g_dbus_proxy_new_finish() to get the
// result.
//
// See g_dbus_proxy_new_sync() and for a synchronous version of this
// constructor.
//
// BusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxy(ctx context.Context, connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, callback AsyncReadyCallback) {
	var _arg7 *C.GCancellable       // out
	var _arg1 *C.GDBusConnection    // out
	var _arg2 C.GDBusProxyFlags     // out
	var _arg3 *C.GDBusInterfaceInfo // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.gchar              // out
	var _arg6 *C.gchar              // out
	var _arg8 C.GAsyncReadyCallback // out
	var _arg9 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	_arg2 = C.GDBusProxyFlags(flags)
	_arg3 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	_arg8 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg9 = C.gpointer(gbox.AssignOnce(callback))

	C.g_dbus_proxy_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// NewDBusProxyForBus: like g_dbus_proxy_new() but takes a Type instead of a
// BusConnection.
//
// BusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxyForBus(ctx context.Context, busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, callback AsyncReadyCallback) {
	var _arg7 *C.GCancellable       // out
	var _arg1 C.GBusType            // out
	var _arg2 C.GDBusProxyFlags     // out
	var _arg3 *C.GDBusInterfaceInfo // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.gchar              // out
	var _arg6 *C.gchar              // out
	var _arg8 C.GAsyncReadyCallback // out
	var _arg9 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GBusType(busType)
	_arg2 = C.GDBusProxyFlags(flags)
	_arg3 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	_arg8 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg9 = C.gpointer(gbox.AssignOnce(callback))

	C.g_dbus_proxy_new_for_bus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}
