// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.g_socket_client_get_type()), F: marshalSocketClient},
	})
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
type SocketClient interface {
	gextras.Objector

	// AddApplicationProXY: enable proxy protocols to be handled by the
	// application. When the indicated proxy protocol is returned by the
	// Resolver, Client will consider this protocol as supported but will not
	// try to find a #GProxy instance to handle handshaking. The application
	// must check for this case by calling
	// g_socket_connection_get_remote_address() on the returned Connection, and
	// seeing if it's a Address of the appropriate type, to determine whether or
	// not it needs to handle the proxy handshaking itself.
	//
	// This should be used for proxy protocols that are dialects of another
	// protocol such as HTTP proxy. It also allows cohabitation of proxy
	// protocols that are reused between protocols. A good example is HTTP. It
	// can be used to proxy HTTP, FTP and Gopher and can also be use as generic
	// socket proxy through the HTTP CONNECT method.
	//
	// When the proxy is detected as being an application proxy, TLS handshake
	// will be skipped. This is required to let the application do the proxy
	// specific handshake.
	AddApplicationProXY(protocol string)
	// Connect tries to resolve the @connectable and make a network connection
	// to it.
	//
	// Upon a successful connection, a new Connection is constructed and
	// returned. The caller owns this new object and must drop their reference
	// to it when finished with it.
	//
	// The type of the Connection object returned depends on the type of the
	// underlying socket that is used. For instance, for a TCP/IP connection it
	// will be a Connection.
	//
	// The socket created will be the same family as the address that the
	// @connectable resolves to, unless family is set with
	// g_socket_client_set_family() or indirectly via
	// g_socket_client_set_local_address(). The socket type defaults to
	// G_SOCKET_TYPE_STREAM but can be set with
	// g_socket_client_set_socket_type().
	//
	// If a local address is specified with g_socket_client_set_local_address()
	// the socket will be bound to this address before connecting.
	Connect(connectable SocketConnectable, cancellable Cancellable) (SocketConnection, error)
	// ConnectAsync: this is the asynchronous version of
	// g_socket_client_connect().
	//
	// You may wish to prefer the asynchronous version even in synchronous
	// command line programs because, since 2.60, it implements RFC 8305
	// (https://tools.ietf.org/html/rfc8305) "Happy Eyeballs" recommendations to
	// work around long connection timeouts in networks where IPv6 is broken by
	// performing an IPv4 connection simultaneously without waiting for IPv6 to
	// time out, which is not supported by the synchronous call. (This is not an
	// API guarantee, and may change in the future.)
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_client_connect_finish() to get the result of the operation.
	ConnectAsync()
	// ConnectFinish finishes an async connect operation. See
	// g_socket_client_connect_async()
	ConnectFinish(result AsyncResult) (SocketConnection, error)
	// ConnectToHost: this is a helper function for g_socket_client_connect().
	//
	// Attempts to create a TCP connection to the named host.
	//
	// @host_and_port may be in any of a number of recognized formats; an IPv6
	// address, an IPv4 address, or a domain name (in which case a DNS lookup is
	// performed). Quoting with [] is supported for all address types. A port
	// override may be specified in the usual way with a colon. Ports may be
	// given as decimal numbers or symbolic names (in which case an
	// /etc/services lookup is performed).
	//
	// If no port override is given in @host_and_port then @default_port will be
	// used as the port number to connect to.
	//
	// In general, @host_and_port is expected to be provided by the user
	// (allowing them to give the hostname, and a port override if necessary)
	// and @default_port is expected to be provided by the application.
	//
	// In the case that an IP address is given, a single connection attempt is
	// made. In the case that a name is given, multiple connection attempts may
	// be made, in turn and according to the number of address records in DNS,
	// until a connection succeeds.
	//
	// Upon a successful connection, a new Connection is constructed and
	// returned. The caller owns this new object and must drop their reference
	// to it when finished with it.
	//
	// In the event of any failure (DNS error, service not found, no hosts
	// connectable) nil is returned and @error (if non-nil) is set accordingly.
	ConnectToHost(hostAndPort string, defaultPort uint16, cancellable Cancellable) (SocketConnection, error)
	// ConnectToHostAsync: this is the asynchronous version of
	// g_socket_client_connect_to_host().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_client_connect_to_host_finish() to get the result of the
	// operation.
	ConnectToHostAsync()
	// ConnectToHostFinish finishes an async connect operation. See
	// g_socket_client_connect_to_host_async()
	ConnectToHostFinish(result AsyncResult) (SocketConnection, error)
	// ConnectToService attempts to create a TCP connection to a service.
	//
	// This call looks up the SRV record for @service at @domain for the "tcp"
	// protocol. It then attempts to connect, in turn, to each of the hosts
	// providing the service until either a connection succeeds or there are no
	// hosts remaining.
	//
	// Upon a successful connection, a new Connection is constructed and
	// returned. The caller owns this new object and must drop their reference
	// to it when finished with it.
	//
	// In the event of any failure (DNS error, service not found, no hosts
	// connectable) nil is returned and @error (if non-nil) is set accordingly.
	ConnectToService(domain string, service string, cancellable Cancellable) (SocketConnection, error)
	// ConnectToServiceAsync: this is the asynchronous version of
	// g_socket_client_connect_to_service().
	ConnectToServiceAsync()
	// ConnectToServiceFinish finishes an async connect operation. See
	// g_socket_client_connect_to_service_async()
	ConnectToServiceFinish(result AsyncResult) (SocketConnection, error)
	// ConnectToURI: this is a helper function for g_socket_client_connect().
	//
	// Attempts to create a TCP connection with a network URI.
	//
	// @uri may be any valid URI containing an "authority" (hostname/port)
	// component. If a port is not specified in the URI, @default_port will be
	// used. TLS will be negotiated if Client:tls is true. (Client does not know
	// to automatically assume TLS for certain URI schemes.)
	//
	// Using this rather than g_socket_client_connect() or
	// g_socket_client_connect_to_host() allows Client to determine when to use
	// application-specific proxy protocols.
	//
	// Upon a successful connection, a new Connection is constructed and
	// returned. The caller owns this new object and must drop their reference
	// to it when finished with it.
	//
	// In the event of any failure (DNS error, service not found, no hosts
	// connectable) nil is returned and @error (if non-nil) is set accordingly.
	ConnectToURI(uri string, defaultPort uint16, cancellable Cancellable) (SocketConnection, error)
	// ConnectToURIAsync: this is the asynchronous version of
	// g_socket_client_connect_to_uri().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_client_connect_to_uri_finish() to get the result of the
	// operation.
	ConnectToURIAsync()
	// ConnectToURIFinish finishes an async connect operation. See
	// g_socket_client_connect_to_uri_async()
	ConnectToURIFinish(result AsyncResult) (SocketConnection, error)
	// EnableProXY gets the proxy enable state; see
	// g_socket_client_set_enable_proxy()
	EnableProXY() bool
	// Family gets the socket family of the socket client.
	//
	// See g_socket_client_set_family() for details.
	Family() SocketFamily
	// LocalAddress gets the local address of the socket client.
	//
	// See g_socket_client_set_local_address() for details.
	LocalAddress() SocketAddress
	// Protocol gets the protocol name type of the socket client.
	//
	// See g_socket_client_set_protocol() for details.
	Protocol() SocketProtocol
	// ProXYResolver gets the Resolver being used by @client. Normally, this
	// will be the resolver returned by g_proxy_resolver_get_default(), but you
	// can override it with g_socket_client_set_proxy_resolver().
	ProXYResolver() ProXYResolver
	// SocketType gets the socket type of the socket client.
	//
	// See g_socket_client_set_socket_type() for details.
	SocketType() SocketType
	// Timeout gets the I/O timeout time for sockets created by @client.
	//
	// See g_socket_client_set_timeout() for details.
	Timeout() uint
	// TLS gets whether @client creates TLS connections. See
	// g_socket_client_set_tls() for details.
	TLS() bool
	// TLSValidationFlags gets the TLS validation flags used creating TLS
	// connections via @client.
	TLSValidationFlags() TLSCertificateFlags
	// SetEnableProXY sets whether or not @client attempts to make connections
	// via a proxy server. When enabled (the default), Client will use a
	// Resolver to determine if a proxy protocol such as SOCKS is needed, and
	// automatically do the necessary proxy negotiation.
	//
	// See also g_socket_client_set_proxy_resolver().
	SetEnableProXY(enable bool)
	// SetFamily sets the socket family of the socket client. If this is set to
	// something other than G_SOCKET_FAMILY_INVALID then the sockets created by
	// this object will be of the specified family.
	//
	// This might be useful for instance if you want to force the local
	// connection to be an ipv4 socket, even though the address might be an ipv6
	// mapped to ipv4 address.
	SetFamily(family SocketFamily)
	// SetLocalAddress sets the local address of the socket client. The sockets
	// created by this object will bound to the specified address (if not nil)
	// before connecting.
	//
	// This is useful if you want to ensure that the local side of the
	// connection is on a specific port, or on a specific interface.
	SetLocalAddress(address SocketAddress)
	// SetProtocol sets the protocol of the socket client. The sockets created
	// by this object will use of the specified protocol.
	//
	// If @protocol is G_SOCKET_PROTOCOL_DEFAULT that means to use the default
	// protocol for the socket family and type.
	SetProtocol(protocol SocketProtocol)
	// SetProXYResolver overrides the Resolver used by @client. You can call
	// this if you want to use specific proxies, rather than using the system
	// default proxy settings.
	//
	// Note that whether or not the proxy resolver is actually used depends on
	// the setting of Client:enable-proxy, which is not changed by this function
	// (but which is true by default)
	SetProXYResolver(proxyResolver ProXYResolver)
	// SetSocketType sets the socket type of the socket client. The sockets
	// created by this object will be of the specified type.
	//
	// It doesn't make sense to specify a type of G_SOCKET_TYPE_DATAGRAM, as
	// GSocketClient is used for connection oriented services.
	SetSocketType(typ SocketType)
	// SetTimeout sets the I/O timeout for sockets created by @client. @timeout
	// is a time in seconds, or 0 for no timeout (the default).
	//
	// The timeout value affects the initial connection attempt as well, so
	// setting this may cause calls to g_socket_client_connect(), etc, to fail
	// with G_IO_ERROR_TIMED_OUT.
	SetTimeout(timeout uint)
	// SetTLS sets whether @client creates TLS (aka SSL) connections. If @tls is
	// true, @client will wrap its connections in a ClientConnection and perform
	// a TLS handshake when connecting.
	//
	// Note that since Client must return a Connection, but ClientConnection is
	// not a Connection, this actually wraps the resulting ClientConnection in a
	// WrapperConnection when returning it. You can use
	// g_tcp_wrapper_connection_get_base_io_stream() on the return value to
	// extract the ClientConnection.
	//
	// If you need to modify the behavior of the TLS handshake (eg, by setting a
	// client-side certificate to use, or connecting to the
	// Connection::accept-certificate signal), you can connect to @client's
	// Client::event signal and wait for it to be emitted with
	// G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you a chance to see the
	// ClientConnection before the handshake starts.
	SetTLS(tls bool)
	// SetTLSValidationFlags sets the TLS validation flags used when creating
	// TLS connections via @client. The default value is
	// G_TLS_CERTIFICATE_VALIDATE_ALL.
	SetTLSValidationFlags(flags TLSCertificateFlags)
}

// socketClient implements the SocketClient interface.
type socketClient struct {
	gextras.Objector
}

var _ SocketClient = (*socketClient)(nil)

// WrapSocketClient wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketClient(obj *externglib.Object) SocketClient {
	return SocketClient{
		Objector: obj,
	}
}

func marshalSocketClient(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketClient(obj), nil
}

// NewSocketClient constructs a class SocketClient.
func NewSocketClient() SocketClient {
	var _cret C.GSocketClient

	cret = C.g_socket_client_new()

	var _socketClient SocketClient

	_socketClient = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketClient)

	return _socketClient
}

// AddApplicationProXY: enable proxy protocols to be handled by the
// application. When the indicated proxy protocol is returned by the
// Resolver, Client will consider this protocol as supported but will not
// try to find a #GProxy instance to handle handshaking. The application
// must check for this case by calling
// g_socket_connection_get_remote_address() on the returned Connection, and
// seeing if it's a Address of the appropriate type, to determine whether or
// not it needs to handle the proxy handshaking itself.
//
// This should be used for proxy protocols that are dialects of another
// protocol such as HTTP proxy. It also allows cohabitation of proxy
// protocols that are reused between protocols. A good example is HTTP. It
// can be used to proxy HTTP, FTP and Gopher and can also be use as generic
// socket proxy through the HTTP CONNECT method.
//
// When the proxy is detected as being an application proxy, TLS handshake
// will be skipped. This is required to let the application do the proxy
// specific handshake.
func (c socketClient) AddApplicationProXY(protocol string) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.gchar

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_socket_client_add_application_proxy(_arg0, _arg1)
}

// Connect tries to resolve the @connectable and make a network connection
// to it.
//
// Upon a successful connection, a new Connection is constructed and
// returned. The caller owns this new object and must drop their reference
// to it when finished with it.
//
// The type of the Connection object returned depends on the type of the
// underlying socket that is used. For instance, for a TCP/IP connection it
// will be a Connection.
//
// The socket created will be the same family as the address that the
// @connectable resolves to, unless family is set with
// g_socket_client_set_family() or indirectly via
// g_socket_client_set_local_address(). The socket type defaults to
// G_SOCKET_TYPE_STREAM but can be set with
// g_socket_client_set_socket_type().
//
// If a local address is specified with g_socket_client_set_local_address()
// the socket will be bound to this address before connecting.
func (c socketClient) Connect(connectable SocketConnectable, cancellable Cancellable) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GSocketConnectable
	var _arg2 *C.GCancellable

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect(_arg0, _arg1, _arg2, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectAsync: this is the asynchronous version of
// g_socket_client_connect().
//
// You may wish to prefer the asynchronous version even in synchronous
// command line programs because, since 2.60, it implements RFC 8305
// (https://tools.ietf.org/html/rfc8305) "Happy Eyeballs" recommendations to
// work around long connection timeouts in networks where IPv6 is broken by
// performing an IPv4 connection simultaneously without waiting for IPv6 to
// time out, which is not supported by the synchronous call. (This is not an
// API guarantee, and may change in the future.)
//
// When the operation is finished @callback will be called. You can then
// call g_socket_client_connect_finish() to get the result of the operation.
func (c socketClient) ConnectAsync() {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	C.g_socket_client_connect_async(_arg0)
}

// ConnectFinish finishes an async connect operation. See
// g_socket_client_connect_async()
func (c socketClient) ConnectFinish(result AsyncResult) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_finish(_arg0, _arg1, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToHost: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection to the named host.
//
// @host_and_port may be in any of a number of recognized formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS lookup is
// performed). Quoting with [] is supported for all address types. A port
// override may be specified in the usual way with a colon. Ports may be
// given as decimal numbers or symbolic names (in which case an
// /etc/services lookup is performed).
//
// If no port override is given in @host_and_port then @default_port will be
// used as the port number to connect to.
//
// In general, @host_and_port is expected to be provided by the user
// (allowing them to give the hostname, and a port override if necessary)
// and @default_port is expected to be provided by the application.
//
// In the case that an IP address is given, a single connection attempt is
// made. In the case that a name is given, multiple connection attempts may
// be made, in turn and according to the number of address records in DNS,
// until a connection succeeds.
//
// Upon a successful connection, a new Connection is constructed and
// returned. The caller owns this new object and must drop their reference
// to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) nil is returned and @error (if non-nil) is set accordingly.
func (c socketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable Cancellable) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.gchar
	var _arg2 C.guint16
	var _arg3 *C.GCancellable

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(hostAndPort))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_host(_arg0, _arg1, _arg2, _arg3, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToHostAsync: this is the asynchronous version of
// g_socket_client_connect_to_host().
//
// When the operation is finished @callback will be called. You can then
// call g_socket_client_connect_to_host_finish() to get the result of the
// operation.
func (c socketClient) ConnectToHostAsync() {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	C.g_socket_client_connect_to_host_async(_arg0)
}

// ConnectToHostFinish finishes an async connect operation. See
// g_socket_client_connect_to_host_async()
func (c socketClient) ConnectToHostFinish(result AsyncResult) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_host_finish(_arg0, _arg1, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToService attempts to create a TCP connection to a service.
//
// This call looks up the SRV record for @service at @domain for the "tcp"
// protocol. It then attempts to connect, in turn, to each of the hosts
// providing the service until either a connection succeeds or there are no
// hosts remaining.
//
// Upon a successful connection, a new Connection is constructed and
// returned. The caller owns this new object and must drop their reference
// to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) nil is returned and @error (if non-nil) is set accordingly.
func (c socketClient) ConnectToService(domain string, service string, cancellable Cancellable) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _arg3 *C.GCancellable

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_service(_arg0, _arg1, _arg2, _arg3, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToServiceAsync: this is the asynchronous version of
// g_socket_client_connect_to_service().
func (c socketClient) ConnectToServiceAsync() {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	C.g_socket_client_connect_to_service_async(_arg0)
}

// ConnectToServiceFinish finishes an async connect operation. See
// g_socket_client_connect_to_service_async()
func (c socketClient) ConnectToServiceFinish(result AsyncResult) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_service_finish(_arg0, _arg1, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToURI: this is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection with a network URI.
//
// @uri may be any valid URI containing an "authority" (hostname/port)
// component. If a port is not specified in the URI, @default_port will be
// used. TLS will be negotiated if Client:tls is true. (Client does not know
// to automatically assume TLS for certain URI schemes.)
//
// Using this rather than g_socket_client_connect() or
// g_socket_client_connect_to_host() allows Client to determine when to use
// application-specific proxy protocols.
//
// Upon a successful connection, a new Connection is constructed and
// returned. The caller owns this new object and must drop their reference
// to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) nil is returned and @error (if non-nil) is set accordingly.
func (c socketClient) ConnectToURI(uri string, defaultPort uint16, cancellable Cancellable) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.gchar
	var _arg2 C.guint16
	var _arg3 *C.GCancellable

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_uri(_arg0, _arg1, _arg2, _arg3, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// ConnectToURIAsync: this is the asynchronous version of
// g_socket_client_connect_to_uri().
//
// When the operation is finished @callback will be called. You can then
// call g_socket_client_connect_to_uri_finish() to get the result of the
// operation.
func (c socketClient) ConnectToURIAsync() {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	C.g_socket_client_connect_to_uri_async(_arg0)
}

// ConnectToURIFinish finishes an async connect operation. See
// g_socket_client_connect_to_uri_async()
func (c socketClient) ConnectToURIFinish(result AsyncResult) (SocketConnection, error) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GSocketConnection
	var _cerr *C.GError

	cret = C.g_socket_client_connect_to_uri_finish(_arg0, _arg1, _cerr)

	var _socketConnection SocketConnection
	var _goerr error

	_socketConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketConnection, _goerr
}

// EnableProXY gets the proxy enable state; see
// g_socket_client_set_enable_proxy()
func (c socketClient) EnableProXY() bool {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean

	cret = C.g_socket_client_get_enable_proxy(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Family gets the socket family of the socket client.
//
// See g_socket_client_set_family() for details.
func (c socketClient) Family() SocketFamily {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.GSocketFamily

	cret = C.g_socket_client_get_family(_arg0)

	var _socketFamily SocketFamily

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// LocalAddress gets the local address of the socket client.
//
// See g_socket_client_set_local_address() for details.
func (c socketClient) LocalAddress() SocketAddress {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret *C.GSocketAddress

	cret = C.g_socket_client_get_local_address(_arg0)

	var _socketAddress SocketAddress

	_socketAddress = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(SocketAddress)

	return _socketAddress
}

// Protocol gets the protocol name type of the socket client.
//
// See g_socket_client_set_protocol() for details.
func (c socketClient) Protocol() SocketProtocol {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.GSocketProtocol

	cret = C.g_socket_client_get_protocol(_arg0)

	var _socketProtocol SocketProtocol

	_socketProtocol = SocketProtocol(_cret)

	return _socketProtocol
}

// ProXYResolver gets the Resolver being used by @client. Normally, this
// will be the resolver returned by g_proxy_resolver_get_default(), but you
// can override it with g_socket_client_set_proxy_resolver().
func (c socketClient) ProXYResolver() ProXYResolver {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret *C.GProxyResolver

	cret = C.g_socket_client_get_proxy_resolver(_arg0)

	var _proxyResolver ProXYResolver

	_proxyResolver = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ProXYResolver)

	return _proxyResolver
}

// SocketType gets the socket type of the socket client.
//
// See g_socket_client_set_socket_type() for details.
func (c socketClient) SocketType() SocketType {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.GSocketType

	cret = C.g_socket_client_get_socket_type(_arg0)

	var _socketType SocketType

	_socketType = SocketType(_cret)

	return _socketType
}

// Timeout gets the I/O timeout time for sockets created by @client.
//
// See g_socket_client_set_timeout() for details.
func (c socketClient) Timeout() uint {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.guint

	cret = C.g_socket_client_get_timeout(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// TLS gets whether @client creates TLS connections. See
// g_socket_client_set_tls() for details.
func (c socketClient) TLS() bool {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean

	cret = C.g_socket_client_get_tls(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// TLSValidationFlags gets the TLS validation flags used creating TLS
// connections via @client.
func (c socketClient) TLSValidationFlags() TLSCertificateFlags {
	var _arg0 *C.GSocketClient

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))

	var _cret C.GTlsCertificateFlags

	cret = C.g_socket_client_get_tls_validation_flags(_arg0)

	var _tlsCertificateFlags TLSCertificateFlags

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// SetEnableProXY sets whether or not @client attempts to make connections
// via a proxy server. When enabled (the default), Client will use a
// Resolver to determine if a proxy protocol such as SOCKS is needed, and
// automatically do the necessary proxy negotiation.
//
// See also g_socket_client_set_proxy_resolver().
func (c socketClient) SetEnableProXY(enable bool) {
	var _arg0 *C.GSocketClient
	var _arg1 C.gboolean

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	if enable {
		_arg1 = C.gboolean(1)
	}

	C.g_socket_client_set_enable_proxy(_arg0, _arg1)
}

// SetFamily sets the socket family of the socket client. If this is set to
// something other than G_SOCKET_FAMILY_INVALID then the sockets created by
// this object will be of the specified family.
//
// This might be useful for instance if you want to force the local
// connection to be an ipv4 socket, even though the address might be an ipv6
// mapped to ipv4 address.
func (c socketClient) SetFamily(family SocketFamily) {
	var _arg0 *C.GSocketClient
	var _arg1 C.GSocketFamily

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GSocketFamily)(family)

	C.g_socket_client_set_family(_arg0, _arg1)
}

// SetLocalAddress sets the local address of the socket client. The sockets
// created by this object will bound to the specified address (if not nil)
// before connecting.
//
// This is useful if you want to ensure that the local side of the
// connection is on a specific port, or on a specific interface.
func (c socketClient) SetLocalAddress(address SocketAddress) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GSocketAddress

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))

	C.g_socket_client_set_local_address(_arg0, _arg1)
}

// SetProtocol sets the protocol of the socket client. The sockets created
// by this object will use of the specified protocol.
//
// If @protocol is G_SOCKET_PROTOCOL_DEFAULT that means to use the default
// protocol for the socket family and type.
func (c socketClient) SetProtocol(protocol SocketProtocol) {
	var _arg0 *C.GSocketClient
	var _arg1 C.GSocketProtocol

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GSocketProtocol)(protocol)

	C.g_socket_client_set_protocol(_arg0, _arg1)
}

// SetProXYResolver overrides the Resolver used by @client. You can call
// this if you want to use specific proxies, rather than using the system
// default proxy settings.
//
// Note that whether or not the proxy resolver is actually used depends on
// the setting of Client:enable-proxy, which is not changed by this function
// (but which is true by default)
func (c socketClient) SetProXYResolver(proxyResolver ProXYResolver) {
	var _arg0 *C.GSocketClient
	var _arg1 *C.GProxyResolver

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GProxyResolver)(unsafe.Pointer(proxyResolver.Native()))

	C.g_socket_client_set_proxy_resolver(_arg0, _arg1)
}

// SetSocketType sets the socket type of the socket client. The sockets
// created by this object will be of the specified type.
//
// It doesn't make sense to specify a type of G_SOCKET_TYPE_DATAGRAM, as
// GSocketClient is used for connection oriented services.
func (c socketClient) SetSocketType(typ SocketType) {
	var _arg0 *C.GSocketClient
	var _arg1 C.GSocketType

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GSocketType)(typ)

	C.g_socket_client_set_socket_type(_arg0, _arg1)
}

// SetTimeout sets the I/O timeout for sockets created by @client. @timeout
// is a time in seconds, or 0 for no timeout (the default).
//
// The timeout value affects the initial connection attempt as well, so
// setting this may cause calls to g_socket_client_connect(), etc, to fail
// with G_IO_ERROR_TIMED_OUT.
func (c socketClient) SetTimeout(timeout uint) {
	var _arg0 *C.GSocketClient
	var _arg1 C.guint

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint(timeout)

	C.g_socket_client_set_timeout(_arg0, _arg1)
}

// SetTLS sets whether @client creates TLS (aka SSL) connections. If @tls is
// true, @client will wrap its connections in a ClientConnection and perform
// a TLS handshake when connecting.
//
// Note that since Client must return a Connection, but ClientConnection is
// not a Connection, this actually wraps the resulting ClientConnection in a
// WrapperConnection when returning it. You can use
// g_tcp_wrapper_connection_get_base_io_stream() on the return value to
// extract the ClientConnection.
//
// If you need to modify the behavior of the TLS handshake (eg, by setting a
// client-side certificate to use, or connecting to the
// Connection::accept-certificate signal), you can connect to @client's
// Client::event signal and wait for it to be emitted with
// G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you a chance to see the
// ClientConnection before the handshake starts.
func (c socketClient) SetTLS(tls bool) {
	var _arg0 *C.GSocketClient
	var _arg1 C.gboolean

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	if tls {
		_arg1 = C.gboolean(1)
	}

	C.g_socket_client_set_tls(_arg0, _arg1)
}

// SetTLSValidationFlags sets the TLS validation flags used when creating
// TLS connections via @client. The default value is
// G_TLS_CERTIFICATE_VALIDATE_ALL.
func (c socketClient) SetTLSValidationFlags(flags TLSCertificateFlags) {
	var _arg0 *C.GSocketClient
	var _arg1 C.GTlsCertificateFlags

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GTlsCertificateFlags)(flags)

	C.g_socket_client_set_tls_validation_flags(_arg0, _arg1)
}