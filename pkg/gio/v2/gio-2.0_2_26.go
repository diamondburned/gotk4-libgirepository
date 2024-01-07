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
// extern void _gotk4_gio2_DBusConnection_ConnectClosed(gpointer, gboolean, GError*, guintptr);
// extern gboolean _gotk4_gio2_DBusServer_ConnectNewConnection(gpointer, void*, guintptr);
// extern gboolean _gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer(gpointer, void*, void*, guintptr);
// extern gboolean _gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeDBusAuthObserver     = coreglib.Type(girepository.MustFind("Gio", "DBusAuthObserver").RegisteredGType())
	GTypeDBusConnection       = coreglib.Type(girepository.MustFind("Gio", "DBusConnection").RegisteredGType())
	GTypeDBusMessage          = coreglib.Type(girepository.MustFind("Gio", "DBusMessage").RegisteredGType())
	GTypeDBusMethodInvocation = coreglib.Type(girepository.MustFind("Gio", "DBusMethodInvocation").RegisteredGType())
	GTypeDBusServer           = coreglib.Type(girepository.MustFind("Gio", "DBusServer").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDBusAuthObserver, F: marshalDBusAuthObserver},
		coreglib.TypeMarshaler{T: GTypeDBusConnection, F: marshalDBusConnection},
		coreglib.TypeMarshaler{T: GTypeDBusMessage, F: marshalDBusMessage},
		coreglib.TypeMarshaler{T: GTypeDBusMethodInvocation, F: marshalDBusMethodInvocation},
		coreglib.TypeMarshaler{T: GTypeDBusServer, F: marshalDBusServer},
	})
}

// DBusAuthObserver type provides a mechanism for participating in how a
// BusServer (or a BusConnection) authenticates remote peers. Simply instantiate
// a BusAuthObserver and connect to the signals you are interested in. Note that
// new signals may be added in the future
//
//
// Controlling Authentication Mechanisms
//
// By default, a BusServer or server-side BusConnection will allow any
// authentication mechanism to be used. If you only want to allow D-Bus
// connections with the EXTERNAL mechanism, which makes use of credentials
// passing and is the recommended mechanism for modern Unix platforms such as
// Linux and the BSD family, you would use a signal handler like this:
//
//    static gboolean
//    on_authorize_authenticated_peer (GDBusAuthObserver *observer,
//                                     GIOStream         *stream,
//                                     GCredentials      *credentials,
//                                     gpointer           user_data)
//    {
//      gboolean authorized;
//
//      authorized = FALSE;
//      if (credentials != NULL)
//        {
//          GCredentials *own_credentials;
//          own_credentials = g_credentials_new ();
//          if (g_credentials_is_same_user (credentials, own_credentials, NULL))
//            authorized = TRUE;
//          g_object_unref (own_credentials);
//        }
//
//      return authorized;
//    }.
type DBusAuthObserver struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusAuthObserver)(nil)
)

func wrapDBusAuthObserver(obj *coreglib.Object) *DBusAuthObserver {
	return &DBusAuthObserver{
		Object: obj,
	}
}

func marshalDBusAuthObserver(p uintptr) (interface{}, error) {
	return wrapDBusAuthObserver(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAllowMechanism is emitted to check if mechanism is allowed to be used.
func (v *DBusAuthObserver) ConnectAllowMechanism(f func(mechanism string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "allow-mechanism", false, unsafe.Pointer(C._gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism), f)
}

// ConnectAuthorizeAuthenticatedPeer is emitted to check if a peer that is
// successfully authenticated is authorized.
func (v *DBusAuthObserver) ConnectAuthorizeAuthenticatedPeer(f func(stream IOStreamer, credentials *Credentials) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "authorize-authenticated-peer", false, unsafe.Pointer(C._gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer), f)
}

// DBusConnection type is used for D-Bus connections to remote peers such as a
// message buses. It is a low-level API that offers a lot of flexibility. For
// instance, it lets you establish a connection over any transport that can by
// represented as a OStream.
//
// This class is rarely used directly in D-Bus clients. If you are writing a
// D-Bus client, it is often easier to use the g_bus_own_name(),
// g_bus_watch_name() or g_dbus_proxy_new_for_bus() APIs.
//
// As an exception to the usual GLib rule that a particular object must not be
// used by two threads at the same time, BusConnection's methods may be called
// from any thread. This is so that g_bus_get() and g_bus_get_sync() can safely
// return the same BusConnection when called from any thread.
//
// Most of the ways to obtain a BusConnection automatically initialize it (i.e.
// connect to D-Bus): for instance, g_dbus_connection_new() and g_bus_get(), and
// the synchronous versions of those methods, give you an initialized
// connection. Language bindings for GIO should use g_initable_new() or
// g_async_initable_new_async(), which also initialize the connection.
//
// If you construct an uninitialized BusConnection, such as via g_object_new(),
// you must initialize it via g_initable_init() or g_async_initable_init_async()
// before using its methods or properties. Calling methods or accessing
// properties on a BusConnection that has not completed initialization
// successfully is considered to be invalid, and leads to undefined behaviour.
// In particular, if initialization fails with a #GError, the only valid thing
// you can do with that BusConnection is to free it with g_object_unref().
//
//
// An example D-Bus server
//
// Here is an example for a D-Bus server: gdbus-example-server.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-server.c)
//
//
// An example for exporting a subtree
//
// Here is an example for exporting a subtree: gdbus-example-subtree.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-subtree.c)
//
//
// An example for file descriptor passing
//
// Here is an example for passing UNIX file descriptors: gdbus-unix-fd-client.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-unix-fd-client.c)
//
//
// An example for exporting a GObject
//
// Here is an example for exporting a #GObject: gdbus-example-export.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-export.c).
type DBusConnection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	AsyncInitable
	Initable
}

var (
	_ coreglib.Objector = (*DBusConnection)(nil)
)

func wrapDBusConnection(obj *coreglib.Object) *DBusConnection {
	return &DBusConnection{
		Object: obj,
		AsyncInitable: AsyncInitable{
			Object: obj,
		},
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusConnection(p uintptr) (interface{}, error) {
	return wrapDBusConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClosed is emitted when the connection is closed.
//
//
// The cause of this event can be
//
// - If g_dbus_connection_close() is called. In this case remote_peer_vanished
// is set to FALSE and error is NULL.
//
// - If the remote peer closes the connection. In this case remote_peer_vanished
// is set to TRUE and error is set.
//
// - If the remote peer sends invalid or malformed data. In this case
// remote_peer_vanished is set to FALSE and error is set.
//
// Upon receiving this signal, you should give up your reference to connection.
// You are guaranteed that this signal is emitted only once.
func (v *DBusConnection) ConnectClosed(f func(remotePeerVanished bool, err error)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "closed", false, unsafe.Pointer(C._gotk4_gio2_DBusConnection_ConnectClosed), f)
}

// DBusMessage: type for representing D-Bus messages that can be sent or
// received on a BusConnection.
type DBusMessage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusMessage)(nil)
)

func wrapDBusMessage(obj *coreglib.Object) *DBusMessage {
	return &DBusMessage{
		Object: obj,
	}
}

func marshalDBusMessage(p uintptr) (interface{}, error) {
	return wrapDBusMessage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusMethodInvocation instances of the BusMethodInvocation class are used when
// handling D-Bus method calls. It provides a way to asynchronously return
// results and errors.
//
// The normal way to obtain a BusMethodInvocation object is to receive it as an
// argument to the handle_method_call() function in a BusInterfaceVTable that
// was passed to g_dbus_connection_register_object().
type DBusMethodInvocation struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusMethodInvocation)(nil)
)

func wrapDBusMethodInvocation(obj *coreglib.Object) *DBusMethodInvocation {
	return &DBusMethodInvocation{
		Object: obj,
	}
}

func marshalDBusMethodInvocation(p uintptr) (interface{}, error) {
	return wrapDBusMethodInvocation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusServer is a helper for listening to and accepting D-Bus connections. This
// can be used to create a new D-Bus server, allowing two peers to use the D-Bus
// protocol for their own specialized communication. A server instance provided
// in this way will not perform message routing or implement the
// org.freedesktop.DBus interface.
//
// To just export an object on a well-known name on a message bus, such as the
// session or system bus, you should instead use g_bus_own_name().
//
// An example of peer-to-peer communication with GDBus can be found in
// gdbus-example-peer.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-peer.c).
//
// Note that a minimal BusServer will accept connections from any peer. In many
// use-cases it will be necessary to add a BusAuthObserver that only accepts
// connections that have successfully authenticated as the same user that is
// running the BusServer. Since GLib 2.68 this can be achieved more simply by
// passing the G_DBUS_SERVER_FLAGS_AUTHENTICATION_REQUIRE_SAME_USER flag to the
// server.
type DBusServer struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Initable
}

var (
	_ coreglib.Objector = (*DBusServer)(nil)
)

func wrapDBusServer(obj *coreglib.Object) *DBusServer {
	return &DBusServer{
		Object: obj,
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusServer(p uintptr) (interface{}, error) {
	return wrapDBusServer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectNewConnection is emitted when a new authenticated connection has been
// made. Use g_dbus_connection_get_peer_credentials() to figure out what
// identity (if any), was authenticated.
//
// If you want to accept the connection, take a reference to the connection
// object and return TRUE. When you are done with the connection call
// g_dbus_connection_close() and give up your reference. Note that the other
// peer may disconnect at any time - a typical thing to do when accepting a
// connection is to listen to the BusConnection::closed signal.
//
// If BusServer:flags contains G_DBUS_SERVER_FLAGS_RUN_IN_THREAD then the signal
// is emitted in a new thread dedicated to the connection. Otherwise the signal
// is emitted in the [thread-default main
// context][g-main-context-push-thread-default] of the thread that server was
// constructed in.
//
// You are guaranteed that signal handlers for this signal runs before incoming
// messages on connection are processed. This means that it's suitable to call
// g_dbus_connection_register_object() or similar from the signal handler.
func (v *DBusServer) ConnectNewConnection(f func(connection *DBusConnection) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "new-connection", false, unsafe.Pointer(C._gotk4_gio2_DBusServer_ConnectNewConnection), f)
}
