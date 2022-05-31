// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GIOStream* _gotk4_gio2_ProxyInterface_connect(void*, void*, void*, void*, GError**);
// extern GIOStream* _gotk4_gio2_ProxyInterface_connect_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_ProxyInterface_supports_hostname(void*);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// glib.Type values for gproxy.go.
var GTypeProxy = coreglib.Type(C.g_proxy_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeProxy, F: marshalProxy},
	})
}

// PROXY_EXTENSION_POINT_NAME: extension point for proxy functionality. See
// [Extending GIO][extending-gio].
const PROXY_EXTENSION_POINT_NAME = "gio-proxy"

// Proxy handles connecting to a remote host via a given type of proxy server.
// It is implemented by the 'gio-proxy' extension point. The extensions are
// named after their proxy protocol name. As an example, a SOCKS5 proxy
// implementation can be retrieved with the name 'socks5' using the function
// g_io_extension_point_get_extension_by_name().
//
// Proxy wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Proxy struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Proxy)(nil)
)

// Proxier describes Proxy's interface methods.
type Proxier interface {
	coreglib.Objector

	// ConnectProxy: given connection to communicate with a proxy (eg, a
	// Connection that is connected to the proxy server), this does the
	// necessary handshake to connect to proxy_address, and if required, wraps
	// the OStream to handle proxy payload.
	ConnectProxy(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress) (IOStreamer, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress, callback AsyncReadyCallback)
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(result AsyncResulter) (IOStreamer, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves.
	SupportsHostname() bool
}

var _ Proxier = (*Proxy)(nil)

func wrapProxy(obj *coreglib.Object) *Proxy {
	return &Proxy{
		Object: obj,
	}
}

func marshalProxy(p uintptr) (interface{}, error) {
	return wrapProxy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectProxy: given connection to communicate with a proxy (eg, a Connection
// that is connected to the proxy server), this does the necessary handshake to
// connect to proxy_address, and if required, wraps the OStream to handle proxy
// payload.
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - connection: OStream.
//    - proxyAddress: Address.
//
// The function returns the following values:
//
//    - ioStream that will replace connection. This might be the same as
//      connection, in which case a reference will be added.
//
func (proxy *Proxy) ConnectProxy(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress) (IOStreamer, error) {
	var _args [4]girepository.Argument
	var _arg0 *C.void // out
	var _arg3 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxyAddress).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2
	*(**C.void)(unsafe.Pointer(&_args[3])) = _arg3

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(proxyAddress)

	var _ioStream IOStreamer // out
	var _goerr error         // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStream, _goerr
}

// ConnectAsync asynchronous version of g_proxy_connect().
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - connection: OStream.
//    - proxyAddress: Address.
//    - callback (optional): ReadyCallback.
//
func (proxy *Proxy) ConnectAsync(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument
	var _arg0 *C.void    // out
	var _arg3 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 *C.void    // out
	var _arg4 C.gpointer // out
	var _arg5 C.gpointer

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxyAddress).Native()))
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2
	*(**C.void)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.gpointer)(unsafe.Pointer(&_args[4])) = _arg4

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(proxyAddress)
	runtime.KeepAlive(callback)
}

// ConnectFinish: see g_proxy_connect().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - ioStream: OStream.
//
func (proxy *Proxy) ConnectFinish(result AsyncResulter) (IOStreamer, error) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)
	runtime.KeepAlive(result)

	var _ioStream IOStreamer // out
	var _goerr error         // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStream, _goerr
}

// SupportsHostname: some proxy protocols expect to be passed a hostname, which
// they will resolve to an IP address themselves. Others, like SOCKS4, do not
// allow this. This function will return FALSE if proxy is implementing such a
// protocol. When FALSE is returned, the caller should resolve the destination
// hostname first, and then pass a Address containing the stringified IP address
// to g_proxy_connect() or g_proxy_connect_async().
//
// The function returns the following values:
//
//    - ok: TRUE if hostname resolution is supported.
//
func (proxy *Proxy) SupportsHostname() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ProxyGetDefaultForProtocol: find the gio-proxy extension point for a proxy
// implementation that supports the specified protocol.
//
// The function takes the following parameters:
//
//    - protocol: proxy protocol name (e.g. http, socks, etc).
//
// The function returns the following values:
//
//    - proxy (optional): return a #GProxy or NULL if protocol is not supported.
//
func ProxyGetDefaultForProtocol(protocol string) *Proxy {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg0))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "get_default_for_protocol").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(protocol)

	var _proxy *Proxy // out

	if _cret != nil {
		_proxy = wrapProxy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _proxy
}
