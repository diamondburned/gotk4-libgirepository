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
// #include <glib-object.h>
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeSocketClient returns the GType for the type SocketClient.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSocketClient() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SocketClient").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSocketClient)
	return gtype
}

// SocketClientOverrider contains methods that are overridable.
type SocketClientOverrider interface {
}

// SocketClient is a lightweight high-level utility class for connecting to a
// network host using a connection oriented socket type.
//
// You create a Client object, set any options you want, and then call a sync or
// async connect operation, which returns a Connection subclass on success.
//
// The type of the Connection object returned depends on the type of the
// underlying socket that is in use. For instance, for a TCP/IP connection it
// will be a Connection.
//
// As Client is a lightweight object, you don't need to cache it. You can just
// create a new one any time you need one.
type SocketClient struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SocketClient)(nil)
)

func classInitSocketClienter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSocketClient(obj *coreglib.Object) *SocketClient {
	return &SocketClient{
		Object: obj,
	}
}

func marshalSocketClient(p uintptr) (interface{}, error) {
	return wrapSocketClient(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSocketClient creates a new Client with the default options.
//
// The function returns the following values:
//
//    - socketClient: Client. Free the returned object with g_object_unref().
//
func NewSocketClient() *SocketClient {
	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("new_SocketClient", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _socketClient *SocketClient // out

	_socketClient = wrapSocketClient(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketClient
}

// AddApplicationProxy: enable proxy protocols to be handled by the application.
// When the indicated proxy protocol is returned by the Resolver, Client will
// consider this protocol as supported but will not try to find a #GProxy
// instance to handle handshaking. The application must check for this case by
// calling g_socket_connection_get_remote_address() on the returned Connection,
// and seeing if it's a Address of the appropriate type, to determine whether or
// not it needs to handle the proxy handshaking itself.
//
// This should be used for proxy protocols that are dialects of another protocol
// such as HTTP proxy. It also allows cohabitation of proxy protocols that are
// reused between protocols. A good example is HTTP. It can be used to proxy
// HTTP, FTP and Gopher and can also be use as generic socket proxy through the
// HTTP CONNECT method.
//
// When the proxy is detected as being an application proxy, TLS handshake will
// be skipped. This is required to let the application do the proxy specific
// handshake.
//
// The function takes the following parameters:
//
//    - protocol: proxy protocol.
//
func (client *SocketClient) AddApplicationProxy(protocol string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("add_application_proxy", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(protocol)
}

// ConnectSocketClient tries to resolve the connectable and make a network
// connection to it.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// The type of the Connection object returned depends on the type of the
// underlying socket that is used. For instance, for a TCP/IP connection it will
// be a Connection.
//
// The socket created will be the same family as the address that the
// connectable resolves to, unless family is set with
// g_socket_client_set_family() or indirectly via
// g_socket_client_set_local_address(). The socket type defaults to
// G_SOCKET_TYPE_STREAM but can be set with g_socket_client_set_socket_type().
//
// If a local address is specified with g_socket_client_set_local_address() the
// socket will be bound to this address before connecting.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - connectable specifying the remote address.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectSocketClient(ctx context.Context, connectable SocketConnectabler) (*SocketConnection, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectAsync: this is the asynchronous version of g_socket_client_connect().
//
// You may wish to prefer the asynchronous version even in synchronous command
// line programs because, since 2.60, it implements RFC 8305
// (https://tools.ietf.org/html/rfc8305) "Happy Eyeballs" recommendations to
// work around long connection timeouts in networks where IPv6 is broken by
// performing an IPv4 connection simultaneously without waiting for IPv6 to time
// out, which is not supported by the synchronous call. (This is not an API
// guarantee, and may change in the future.)
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - connectable specifying the remote address.
//    - callback (optional): ReadyCallback.
//
func (client *SocketClient) ConnectAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_async", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)
	runtime.KeepAlive(callback)
}

// ConnectFinish finishes an async connect operation. See
// g_socket_client_connect_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectFinish(result AsyncResulter) (*SocketConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToHost: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection to the named host.
//
// host_and_port may be in any of a number of recognized formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS lookup is
// performed). Quoting with [] is supported for all address types. A port
// override may be specified in the usual way with a colon. Ports may be given
// as decimal numbers or symbolic names (in which case an /etc/services lookup
// is performed).
//
// If no port override is given in host_and_port then default_port will be used
// as the port number to connect to.
//
// In general, host_and_port is expected to be provided by the user (allowing
// them to give the hostname, and a port override if necessary) and default_port
// is expected to be provided by the application.
//
// In the case that an IP address is given, a single connection attempt is made.
// In the case that a name is given, multiple connection attempts may be made,
// in turn and according to the number of address records in DNS, until a
// connection succeeds.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - hostAndPort: name and optionally port of the host to connect to.
//    - defaultPort: default port to connect to.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectToHost(ctx context.Context, hostAndPort string, defaultPort uint16) (*SocketConnection, error) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.guint16)(unsafe.Pointer(&_args[2])) = C.guint16(defaultPort)

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_host", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToHostAsync: this is the asynchronous version of
// g_socket_client_connect_to_host().
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_to_host_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - hostAndPort: name and optionally the port of the host to connect to.
//    - defaultPort: default port to connect to.
//    - callback (optional): ReadyCallback.
//
func (client *SocketClient) ConnectToHostAsync(ctx context.Context, hostAndPort string, defaultPort uint16, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.guint16)(unsafe.Pointer(&_args[2])) = C.guint16(defaultPort)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_host_async", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)
	runtime.KeepAlive(callback)
}

// ConnectToHostFinish finishes an async connect operation. See
// g_socket_client_connect_to_host_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectToHostFinish(result AsyncResulter) (*SocketConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_host_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToService attempts to create a TCP connection to a service.
//
// This call looks up the SRV record for service at domain for the "tcp"
// protocol. It then attempts to connect, in turn, to each of the hosts
// providing the service until either a connection succeeds or there are no
// hosts remaining.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - domain name.
//    - service: name of the service to connect to.
//
// The function returns the following values:
//
//    - socketConnection if successful, or NULL on error.
//
func (client *SocketClient) ConnectToService(ctx context.Context, domain, service string) (*SocketConnection, error) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_args[2]))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_service", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(service)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToServiceAsync: this is the asynchronous version of
// g_socket_client_connect_to_service().
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - domain name.
//    - service: name of the service to connect to.
//    - callback (optional): ReadyCallback.
//
func (client *SocketClient) ConnectToServiceAsync(ctx context.Context, domain, service string, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_args[2]))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_service_async", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(service)
	runtime.KeepAlive(callback)
}

// ConnectToServiceFinish finishes an async connect operation. See
// g_socket_client_connect_to_service_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectToServiceFinish(result AsyncResulter) (*SocketConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_service_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToURI: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection with a network URI.
//
// uri may be any valid URI containing an "authority" (hostname/port) component.
// If a port is not specified in the URI, default_port will be used. TLS will be
// negotiated if Client:tls is TRUE. (Client does not know to automatically
// assume TLS for certain URI schemes.)
//
// Using this rather than g_socket_client_connect() or
// g_socket_client_connect_to_host() allows Client to determine when to use
// application-specific proxy protocols.
//
// Upon a successful connection, a new Connection is constructed and returned.
// The caller owns this new object and must drop their reference to it when
// finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) NULL is returned and error (if non-NULL) is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: network URI.
//    - defaultPort: default port to connect to.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectToURI(ctx context.Context, uri string, defaultPort uint16) (*SocketConnection, error) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.guint16)(unsafe.Pointer(&_args[2])) = C.guint16(defaultPort)

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_uri", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(defaultPort)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// ConnectToURIAsync: this is the asynchronous version of
// g_socket_client_connect_to_uri().
//
// When the operation is finished callback will be called. You can then call
// g_socket_client_connect_to_uri_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: network uri.
//    - defaultPort: default port to connect to.
//    - callback (optional): ReadyCallback.
//
func (client *SocketClient) ConnectToURIAsync(ctx context.Context, uri string, defaultPort uint16, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.guint16)(unsafe.Pointer(&_args[2])) = C.guint16(defaultPort)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_uri_async", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(defaultPort)
	runtime.KeepAlive(callback)
}

// ConnectToURIFinish finishes an async connect operation. See
// g_socket_client_connect_to_uri_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketConnection on success, NULL on error.
//
func (client *SocketClient) ConnectToURIFinish(result AsyncResulter) (*SocketConnection, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("connect_to_uri_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)
	runtime.KeepAlive(result)

	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _socketConnection, _goerr
}

// EnableProxy gets the proxy enable state; see
// g_socket_client_set_enable_proxy().
//
// The function returns the following values:
//
//    - ok: whether proxying is enabled.
//
func (client *SocketClient) EnableProxy() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("get_enable_proxy", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// LocalAddress gets the local address of the socket client.
//
// See g_socket_client_set_local_address() for details.
//
// The function returns the following values:
//
//    - socketAddress (optional) or NULL. Do not free.
//
func (client *SocketClient) LocalAddress() SocketAddresser {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("get_local_address", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)

	var _socketAddress SocketAddresser // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(SocketAddresser)
				return ok
			})
			rv, ok := casted.(SocketAddresser)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
			}
			_socketAddress = rv
		}
	}

	return _socketAddress
}

// ProxyResolver gets the Resolver being used by client. Normally, this will be
// the resolver returned by g_proxy_resolver_get_default(), but you can override
// it with g_socket_client_set_proxy_resolver().
//
// The function returns the following values:
//
//    - proxyResolver being used by client.
//
func (client *SocketClient) ProxyResolver() *ProxyResolver {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("get_proxy_resolver", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)

	var _proxyResolver *ProxyResolver // out

	_proxyResolver = wrapProxyResolver(coreglib.Take(unsafe.Pointer(_cret)))

	return _proxyResolver
}

// Timeout gets the I/O timeout time for sockets created by client.
//
// See g_socket_client_set_timeout() for details.
//
// The function returns the following values:
//
//    - guint: timeout in seconds.
//
func (client *SocketClient) Timeout() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("get_timeout", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// TLS gets whether client creates TLS connections. See
// g_socket_client_set_tls() for details.
//
// The function returns the following values:
//
//    - ok: whether client uses TLS.
//
func (client *SocketClient) TLS() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_gret := girepository.MustFind("Gio", "SocketClient").InvokeMethod("get_tls", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(client)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetEnableProxy sets whether or not client attempts to make connections via a
// proxy server. When enabled (the default), Client will use a Resolver to
// determine if a proxy protocol such as SOCKS is needed, and automatically do
// the necessary proxy negotiation.
//
// See also g_socket_client_set_proxy_resolver().
//
// The function takes the following parameters:
//
//    - enable: whether to enable proxies.
//
func (client *SocketClient) SetEnableProxy(enable bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	if enable {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("set_enable_proxy", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(enable)
}

// SetLocalAddress sets the local address of the socket client. The sockets
// created by this object will bound to the specified address (if not NULL)
// before connecting.
//
// This is useful if you want to ensure that the local side of the connection is
// on a specific port, or on a specific interface.
//
// The function takes the following parameters:
//
//    - address (optional) or NULL.
//
func (client *SocketClient) SetLocalAddress(address SocketAddresser) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	if address != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("set_local_address", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(address)
}

// SetProxyResolver overrides the Resolver used by client. You can call this if
// you want to use specific proxies, rather than using the system default proxy
// settings.
//
// Note that whether or not the proxy resolver is actually used depends on the
// setting of Client:enable-proxy, which is not changed by this function (but
// which is TRUE by default).
//
// The function takes the following parameters:
//
//    - proxyResolver (optional) or NULL for the default.
//
func (client *SocketClient) SetProxyResolver(proxyResolver ProxyResolverer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	if proxyResolver != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxyResolver).Native()))
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("set_proxy_resolver", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(proxyResolver)
}

// SetTimeout sets the I/O timeout for sockets created by client. timeout is a
// time in seconds, or 0 for no timeout (the default).
//
// The timeout value affects the initial connection attempt as well, so setting
// this may cause calls to g_socket_client_connect(), etc, to fail with
// G_IO_ERROR_TIMED_OUT.
//
// The function takes the following parameters:
//
//    - timeout: timeout.
//
func (client *SocketClient) SetTimeout(timeout uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(timeout)

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("set_timeout", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(timeout)
}

// SetTLS sets whether client creates TLS (aka SSL) connections. If tls is TRUE,
// client will wrap its connections in a ClientConnection and perform a TLS
// handshake when connecting.
//
// Note that since Client must return a Connection, but ClientConnection is not
// a Connection, this actually wraps the resulting ClientConnection in a
// WrapperConnection when returning it. You can use
// g_tcp_wrapper_connection_get_base_io_stream() on the return value to extract
// the ClientConnection.
//
// If you need to modify the behavior of the TLS handshake (eg, by setting a
// client-side certificate to use, or connecting to the
// Connection::accept-certificate signal), you can connect to client's
// Client::event signal and wait for it to be emitted with
// G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you a chance to see the
// ClientConnection before the handshake starts.
//
// The function takes the following parameters:
//
//    - tls: whether to use TLS.
//
func (client *SocketClient) SetTLS(tls bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	if tls {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "SocketClient").InvokeMethod("set_tls", _args[:], nil)

	runtime.KeepAlive(client)
	runtime.KeepAlive(tls)
}
