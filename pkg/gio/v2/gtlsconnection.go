// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.g_tls_connection_get_type()), F: marshalTLSConnection},
	})
}

// TLSConnection is the base TLS connection class type, which wraps a OStream
// and provides TLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side TLS,
// respectively.
//
// For DTLS (Datagram TLS) support, see Connection.
type TLSConnection interface {
	IOStream

	// EmitAcceptCertificate: used by Connection implementations to emit the
	// Connection::accept-certificate signal.
	EmitAcceptCertificate(peerCert TLSCertificate, errors TLSCertificateFlags) bool
	// Certificate gets @conn's certificate, as set by
	// g_tls_connection_set_certificate().
	Certificate() TLSCertificate
	// ChannelBindingData: query the TLS backend for TLS channel binding data of
	// @type for @conn.
	//
	// This call retrieves TLS channel binding data as specified in RFC 5056
	// (https://tools.ietf.org/html/rfc5056), RFC 5929
	// (https://tools.ietf.org/html/rfc5929), and related RFCs. The binding data
	// is returned in @data. The @data is resized by the callee using Array
	// buffer management and will be freed when the @data is destroyed by
	// g_byte_array_unref(). If @data is nil, it will only check whether TLS
	// backend is able to fetch the data (e.g. whether @type is supported by the
	// TLS backend). It does not guarantee that the data will be available
	// though. That could happen if TLS connection does not support @type or the
	// binding data is not available yet due to additional negotiation or input
	// required.
	ChannelBindingData(typ TLSChannelBindingType) (data []byte, err error)
	// Database gets the certificate database that @conn uses to verify peer
	// certificates. See g_tls_connection_set_database().
	Database() TLSDatabase
	// Interaction: get the object that will be used to interact with the user.
	// It will be used for things like prompting the user for passwords. If nil
	// is returned, then no user interaction will occur for this connection.
	Interaction() TLSInteraction
	// NegotiatedProtocol gets the name of the application-layer protocol
	// negotiated during the handshake.
	//
	// If the peer did not use the ALPN extension, or did not advertise a
	// protocol that matched one of @conn's protocols, or the TLS backend does
	// not support ALPN, then this will be nil. See
	// g_tls_connection_set_advertised_protocols().
	NegotiatedProtocol() string
	// PeerCertificate gets @conn's peer's certificate after the handshake has
	// completed or failed. (It is not set during the emission of
	// Connection::accept-certificate.)
	PeerCertificate() TLSCertificate
	// PeerCertificateErrors gets the errors associated with validating @conn's
	// peer's certificate, after the handshake has completed or failed. (It is
	// not set during the emission of Connection::accept-certificate.)
	PeerCertificateErrors() TLSCertificateFlags
	// RehandshakeMode gets @conn rehandshaking mode. See
	// g_tls_connection_set_rehandshake_mode() for details.
	RehandshakeMode() TLSRehandshakeMode
	// RequireCloseNotify tests whether or not @conn expects a proper TLS close
	// notification when the connection is closed. See
	// g_tls_connection_set_require_close_notify() for details.
	RequireCloseNotify() bool
	// UseSystemCertdb gets whether @conn uses the system certificate database
	// to verify peer certificates. See
	// g_tls_connection_set_use_system_certdb().
	UseSystemCertdb() bool
	// Handshake attempts a TLS handshake on @conn.
	//
	// On the client side, it is never necessary to call this method; although
	// the connection needs to perform a handshake after connecting (or after
	// sending a "STARTTLS"-type command), Connection will handle this for you
	// automatically when you try to send or receive data on the connection. You
	// can call g_tls_connection_handshake() manually if you want to know
	// whether the initial handshake succeeded or failed (as opposed to just
	// immediately trying to use @conn to read or write, in which case, if it
	// fails, it may not be possible to tell if it failed before or after
	// completing the handshake), but beware that servers may reject client
	// authentication after the handshake has completed, so a successful
	// handshake does not indicate the connection will be usable.
	//
	// Likewise, on the server side, although a handshake is necessary at the
	// beginning of the communication, you do not need to call this function
	// explicitly unless you want clearer error reporting.
	//
	// Previously, calling g_tls_connection_handshake() after the initial
	// handshake would trigger a rehandshake; however, this usage was deprecated
	// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
	// TLS 1.3. Since GLib 2.64, calling this function after the initial
	// handshake will no longer do anything.
	//
	// When using a Connection created by Client, the Client performs the
	// initial handshake, so calling this function manually is not recommended.
	//
	// Connection::accept_certificate may be emitted during the handshake.
	Handshake(cancellable Cancellable) error
	// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
	// g_tls_connection_handshake() for more information.
	HandshakeAsync()
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_tls_connection_handshake() for more information.
	HandshakeFinish(result AsyncResult) error
	// SetAdvertisedProtocols sets the list of application-layer protocols to
	// advertise that the caller is willing to speak on this connection. The
	// Application-Layer Protocol Negotiation (ALPN) extension will be used to
	// negotiate a compatible protocol with the peer; use
	// g_tls_connection_get_negotiated_protocol() to find the negotiated
	// protocol after the handshake. Specifying nil for the the value of
	// @protocols will disable ALPN negotiation.
	//
	// See IANA TLS ALPN Protocol IDs
	// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
	// for a list of registered protocol IDs.
	SetAdvertisedProtocols(protocols []string)
	// SetCertificate: this sets the certificate that @conn will present to its
	// peer during the TLS handshake. For a ServerConnection, it is mandatory to
	// set this, and that will normally be done at construct time.
	//
	// For a ClientConnection, this is optional. If a handshake fails with
	// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
	// certificate, and if you try connecting again, you should call this method
	// first. You can call g_tls_client_connection_get_accepted_cas() on the
	// failed connection to get a list of Certificate Authorities that the
	// server will accept certificates from.
	//
	// (It is also possible that a server will allow the connection with or
	// without a certificate; in that case, if you don't provide a certificate,
	// you can tell that the server requested one by the fact that
	// g_tls_client_connection_get_accepted_cas() will return non-nil.)
	SetCertificate(certificate TLSCertificate)
	// SetDatabase sets the certificate database that is used to verify peer
	// certificates. This is set to the default database by default. See
	// g_tls_backend_get_default_database(). If set to nil, then peer
	// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA
	// error (meaning Connection::accept-certificate will always be emitted on
	// client-side connections, unless that bit is not set in
	// ClientConnection:validation-flags).
	SetDatabase(database TLSDatabase)
	// SetInteraction: set the object that will be used to interact with the
	// user. It will be used for things like prompting the user for passwords.
	//
	// The @interaction argument will normally be a derived subclass of
	// Interaction. nil can also be provided if no user interaction should occur
	// for this connection.
	SetInteraction(interaction TLSInteraction)
	// SetRehandshakeMode: since GLib 2.64, changing the rehandshake mode is no
	// longer supported and will have no effect. With TLS 1.3, rehandshaking has
	// been removed from the TLS protocol, replaced by separate post-handshake
	// authentication and rekey operations.
	SetRehandshakeMode(mode TLSRehandshakeMode)
	// SetRequireCloseNotify sets whether or not @conn expects a proper TLS
	// close notification before the connection is closed. If this is true (the
	// default), then @conn will expect to receive a TLS close notification from
	// its peer before the connection is closed, and will return a
	// G_TLS_ERROR_EOF error if the connection is closed without proper
	// notification (since this may indicate a network error, or
	// man-in-the-middle attack).
	//
	// In some protocols, the application will know whether or not the
	// connection was closed cleanly based on application-level data (because
	// the application-level data includes a length field, or is somehow
	// self-delimiting); in this case, the close notify is redundant and
	// sometimes omitted. (TLS 1.1 explicitly allows this; in TLS 1.0 it is
	// technically an error, but often done anyway.) You can use
	// g_tls_connection_set_require_close_notify() to tell @conn to allow an
	// "unannounced" connection close, in which case the close will show up as a
	// 0-length read, as in a non-TLS Connection, and it is up to the
	// application to check that the data has been fully received.
	//
	// Note that this only affects the behavior when the peer closes the
	// connection; when the application calls g_io_stream_close() itself on
	// @conn, this will send a close notification regardless of the setting of
	// this property. If you explicitly want to do an unclean close, you can
	// close @conn's Connection:base-io-stream rather than closing @conn itself,
	// but note that this may only be done when no other operations are pending
	// on @conn or the base I/O stream.
	SetRequireCloseNotify(requireCloseNotify bool)
	// SetUseSystemCertdb sets whether @conn uses the system certificate
	// database to verify peer certificates. This is true by default. If set to
	// false, then peer certificate validation will always set the
	// G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
	// Connection::accept-certificate will always be emitted on client-side
	// connections, unless that bit is not set in
	// ClientConnection:validation-flags).
	SetUseSystemCertdb(useSystemCertdb bool)
}

// tlsConnection implements the TLSConnection interface.
type tlsConnection struct {
	IOStream
}

var _ TLSConnection = (*tlsConnection)(nil)

// WrapTLSConnection wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSConnection(obj *externglib.Object) TLSConnection {
	return TLSConnection{
		IOStream: WrapIOStream(obj),
	}
}

func marshalTLSConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSConnection(obj), nil
}

// EmitAcceptCertificate: used by Connection implementations to emit the
// Connection::accept-certificate signal.
func (c tlsConnection) EmitAcceptCertificate(peerCert TLSCertificate, errors TLSCertificateFlags) bool {
	var arg0 *C.GTlsConnection
	var arg1 *C.GTlsCertificate
	var arg2 C.GTlsCertificateFlags

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(peerCert.Native()))
	arg2 = (C.GTlsCertificateFlags)(errors)

	var cret C.gboolean
	var goret bool

	cret = C.g_tls_connection_emit_accept_certificate(arg0, arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// Certificate gets @conn's certificate, as set by
// g_tls_connection_set_certificate().
func (c tlsConnection) Certificate() TLSCertificate {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GTlsCertificate
	var goret TLSCertificate

	cret = C.g_tls_connection_get_certificate(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TLSCertificate)

	return goret
}

// ChannelBindingData: query the TLS backend for TLS channel binding data of
// @type for @conn.
//
// This call retrieves TLS channel binding data as specified in RFC 5056
// (https://tools.ietf.org/html/rfc5056), RFC 5929
// (https://tools.ietf.org/html/rfc5929), and related RFCs. The binding data
// is returned in @data. The @data is resized by the callee using Array
// buffer management and will be freed when the @data is destroyed by
// g_byte_array_unref(). If @data is nil, it will only check whether TLS
// backend is able to fetch the data (e.g. whether @type is supported by the
// TLS backend). It does not guarantee that the data will be available
// though. That could happen if TLS connection does not support @type or the
// binding data is not available yet due to additional negotiation or input
// required.
func (c tlsConnection) ChannelBindingData(typ TLSChannelBindingType) (data []byte, err error) {
	var arg0 *C.GTlsConnection
	var arg1 C.GTlsChannelBindingType

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (C.GTlsChannelBindingType)(typ)

	var arg2 *C.GByteArray
	var ret2 []byte
	var cerr *C.GError
	var goerr error

	C.g_tls_connection_get_channel_binding_data(arg0, arg1, arg2, &cerr)

	{
		var length int
		for p := arg2; *p != 0; p = (*C.GByteArray)(ptr.Add(unsafe.Pointer(p), C.sizeof_guint8)) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret2 = make([]byte, length)
		for i := uintptr(0); i < uintptr(length); i += C.sizeof_guint8 {
			src := (C.guint8)(ptr.Add(unsafe.Pointer(arg2), i))
			ret2[i] = byte(src)
		}
	}
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return ret2, goerr
}

// Database gets the certificate database that @conn uses to verify peer
// certificates. See g_tls_connection_set_database().
func (c tlsConnection) Database() TLSDatabase {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GTlsDatabase
	var goret TLSDatabase

	cret = C.g_tls_connection_get_database(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TLSDatabase)

	return goret
}

// Interaction: get the object that will be used to interact with the user.
// It will be used for things like prompting the user for passwords. If nil
// is returned, then no user interaction will occur for this connection.
func (c tlsConnection) Interaction() TLSInteraction {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GTlsInteraction
	var goret TLSInteraction

	cret = C.g_tls_connection_get_interaction(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TLSInteraction)

	return goret
}

// NegotiatedProtocol gets the name of the application-layer protocol
// negotiated during the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a
// protocol that matched one of @conn's protocols, or the TLS backend does
// not support ALPN, then this will be nil. See
// g_tls_connection_set_advertised_protocols().
func (c tlsConnection) NegotiatedProtocol() string {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret *C.gchar
	var goret string

	cret = C.g_tls_connection_get_negotiated_protocol(arg0)

	goret = C.GoString(cret)

	return goret
}

// PeerCertificate gets @conn's peer's certificate after the handshake has
// completed or failed. (It is not set during the emission of
// Connection::accept-certificate.)
func (c tlsConnection) PeerCertificate() TLSCertificate {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GTlsCertificate
	var goret TLSCertificate

	cret = C.g_tls_connection_get_peer_certificate(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TLSCertificate)

	return goret
}

// PeerCertificateErrors gets the errors associated with validating @conn's
// peer's certificate, after the handshake has completed or failed. (It is
// not set during the emission of Connection::accept-certificate.)
func (c tlsConnection) PeerCertificateErrors() TLSCertificateFlags {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret C.GTlsCertificateFlags
	var goret TLSCertificateFlags

	cret = C.g_tls_connection_get_peer_certificate_errors(arg0)

	goret = TLSCertificateFlags(cret)

	return goret
}

// RehandshakeMode gets @conn rehandshaking mode. See
// g_tls_connection_set_rehandshake_mode() for details.
func (c tlsConnection) RehandshakeMode() TLSRehandshakeMode {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret C.GTlsRehandshakeMode
	var goret TLSRehandshakeMode

	cret = C.g_tls_connection_get_rehandshake_mode(arg0)

	goret = TLSRehandshakeMode(cret)

	return goret
}

// RequireCloseNotify tests whether or not @conn expects a proper TLS close
// notification when the connection is closed. See
// g_tls_connection_set_require_close_notify() for details.
func (c tlsConnection) RequireCloseNotify() bool {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_tls_connection_get_require_close_notify(arg0)

	if cret {
		goret = true
	}

	return goret
}

// UseSystemCertdb gets whether @conn uses the system certificate database
// to verify peer certificates. See
// g_tls_connection_set_use_system_certdb().
func (c tlsConnection) UseSystemCertdb() bool {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_tls_connection_get_use_system_certdb(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Handshake attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method; although
// the connection needs to perform a handshake after connecting (or after
// sending a "STARTTLS"-type command), Connection will handle this for you
// automatically when you try to send or receive data on the connection. You
// can call g_tls_connection_handshake() manually if you want to know
// whether the initial handshake succeeded or failed (as opposed to just
// immediately trying to use @conn to read or write, in which case, if it
// fails, it may not be possible to tell if it failed before or after
// completing the handshake), but beware that servers may reject client
// authentication after the handshake has completed, so a successful
// handshake does not indicate the connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at the
// beginning of the communication, you do not need to call this function
// explicitly unless you want clearer error reporting.
//
// Previously, calling g_tls_connection_handshake() after the initial
// handshake would trigger a rehandshake; however, this usage was deprecated
// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
// TLS 1.3. Since GLib 2.64, calling this function after the initial
// handshake will no longer do anything.
//
// When using a Connection created by Client, the Client performs the
// initial handshake, so calling this function manually is not recommended.
//
// Connection::accept_certificate may be emitted during the handshake.
func (c tlsConnection) Handshake(cancellable Cancellable) error {
	var arg0 *C.GTlsConnection
	var arg1 *C.GCancellable

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError
	var goerr error

	C.g_tls_connection_handshake(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
// g_tls_connection_handshake() for more information.
func (c tlsConnection) HandshakeAsync() {
	var arg0 *C.GTlsConnection

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))

	C.g_tls_connection_handshake_async(arg0, arg1, arg2, arg3, arg4)
}

// HandshakeFinish: finish an asynchronous TLS handshake operation. See
// g_tls_connection_handshake() for more information.
func (c tlsConnection) HandshakeFinish(result AsyncResult) error {
	var arg0 *C.GTlsConnection
	var arg1 *C.GAsyncResult

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.g_tls_connection_handshake_finish(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// SetAdvertisedProtocols sets the list of application-layer protocols to
// advertise that the caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be used to
// negotiate a compatible protocol with the peer; use
// g_tls_connection_get_negotiated_protocol() to find the negotiated
// protocol after the handshake. Specifying nil for the the value of
// @protocols will disable ALPN negotiation.
//
// See IANA TLS ALPN Protocol IDs
// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
func (c tlsConnection) SetAdvertisedProtocols(protocols []string) {
	var arg0 *C.GTlsConnection
	var arg1 **C.gchar

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (**C.gchar)(C.malloc((len(protocols) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(protocols)))

		for i := range protocols {
			out[i] = (*C.gchar)(C.CString(protocols[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_tls_connection_set_advertised_protocols(arg0, arg1)
}

// SetCertificate: this sets the certificate that @conn will present to its
// peer during the TLS handshake. For a ServerConnection, it is mandatory to
// set this, and that will normally be done at construct time.
//
// For a ClientConnection, this is optional. If a handshake fails with
// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
// certificate, and if you try connecting again, you should call this method
// first. You can call g_tls_client_connection_get_accepted_cas() on the
// failed connection to get a list of Certificate Authorities that the
// server will accept certificates from.
//
// (It is also possible that a server will allow the connection with or
// without a certificate; in that case, if you don't provide a certificate,
// you can tell that the server requested one by the fact that
// g_tls_client_connection_get_accepted_cas() will return non-nil.)
func (c tlsConnection) SetCertificate(certificate TLSCertificate) {
	var arg0 *C.GTlsConnection
	var arg1 *C.GTlsCertificate

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	C.g_tls_connection_set_certificate(arg0, arg1)
}

// SetDatabase sets the certificate database that is used to verify peer
// certificates. This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to nil, then peer
// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA
// error (meaning Connection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// ClientConnection:validation-flags).
func (c tlsConnection) SetDatabase(database TLSDatabase) {
	var arg0 *C.GTlsConnection
	var arg1 *C.GTlsDatabase

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsDatabase)(unsafe.Pointer(database.Native()))

	C.g_tls_connection_set_database(arg0, arg1)
}

// SetInteraction: set the object that will be used to interact with the
// user. It will be used for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of
// Interaction. nil can also be provided if no user interaction should occur
// for this connection.
func (c tlsConnection) SetInteraction(interaction TLSInteraction) {
	var arg0 *C.GTlsConnection
	var arg1 *C.GTlsInteraction

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))

	C.g_tls_connection_set_interaction(arg0, arg1)
}

// SetRehandshakeMode: since GLib 2.64, changing the rehandshake mode is no
// longer supported and will have no effect. With TLS 1.3, rehandshaking has
// been removed from the TLS protocol, replaced by separate post-handshake
// authentication and rekey operations.
func (c tlsConnection) SetRehandshakeMode(mode TLSRehandshakeMode) {
	var arg0 *C.GTlsConnection
	var arg1 C.GTlsRehandshakeMode

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (C.GTlsRehandshakeMode)(mode)

	C.g_tls_connection_set_rehandshake_mode(arg0, arg1)
}

// SetRequireCloseNotify sets whether or not @conn expects a proper TLS
// close notification before the connection is closed. If this is true (the
// default), then @conn will expect to receive a TLS close notification from
// its peer before the connection is closed, and will return a
// G_TLS_ERROR_EOF error if the connection is closed without proper
// notification (since this may indicate a network error, or
// man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the
// connection was closed cleanly based on application-level data (because
// the application-level data includes a length field, or is somehow
// self-delimiting); in this case, the close notify is redundant and
// sometimes omitted. (TLS 1.1 explicitly allows this; in TLS 1.0 it is
// technically an error, but often done anyway.) You can use
// g_tls_connection_set_require_close_notify() to tell @conn to allow an
// "unannounced" connection close, in which case the close will show up as a
// 0-length read, as in a non-TLS Connection, and it is up to the
// application to check that the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the
// connection; when the application calls g_io_stream_close() itself on
// @conn, this will send a close notification regardless of the setting of
// this property. If you explicitly want to do an unclean close, you can
// close @conn's Connection:base-io-stream rather than closing @conn itself,
// but note that this may only be done when no other operations are pending
// on @conn or the base I/O stream.
func (c tlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	var arg0 *C.GTlsConnection
	var arg1 C.gboolean

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	if requireCloseNotify {
		arg1 = C.gboolean(1)
	}

	C.g_tls_connection_set_require_close_notify(arg0, arg1)
}

// SetUseSystemCertdb sets whether @conn uses the system certificate
// database to verify peer certificates. This is true by default. If set to
// false, then peer certificate validation will always set the
// G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// Connection::accept-certificate will always be emitted on client-side
// connections, unless that bit is not set in
// ClientConnection:validation-flags).
func (c tlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	var arg0 *C.GTlsConnection
	var arg1 C.gboolean

	arg0 = (*C.GTlsConnection)(unsafe.Pointer(c.Native()))
	if useSystemCertdb {
		arg1 = C.gboolean(1)
	}

	C.g_tls_connection_set_use_system_certdb(arg0, arg1)
}
