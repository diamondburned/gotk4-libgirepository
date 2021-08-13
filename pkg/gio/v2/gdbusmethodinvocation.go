// Code generated by girgen. DO NOT EDIT.

package gio

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
import "C"

// DBUS_METHOD_INVOCATION_HANDLED: value returned by handlers of the signals
// generated by the gdbus-codegen tool to indicate that a method call has been
// handled by an implementation. It is equal to TRUE, but using this macro is
// sometimes more readable.
//
// In code that needs to be backwards-compatible with older GLib, use TRUE
// instead, often written like this:
//
//    g_dbus_method_invocation_return_error (invocation, ...);
//    return TRUE;    // handled
const DBUS_METHOD_INVOCATION_HANDLED = true

// DBUS_METHOD_INVOCATION_UNHANDLED: value returned by handlers of the signals
// generated by the gdbus-codegen tool to indicate that a method call has not
// been handled by an implementation. It is equal to FALSE, but using this macro
// is sometimes more readable.
//
// In code that needs to be backwards-compatible with older GLib, use FALSE
// instead.
const DBUS_METHOD_INVOCATION_UNHANDLED = false
