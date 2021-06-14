// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_credentials_get_type()), F: marshalCredentials},
	})
}

// Credentials: the #GCredentials type is a reference-counted wrapper for native
// credentials. This information is typically used for identifying,
// authenticating and authorizing other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a `struct ucred` - see the unix(7)
// man page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a `struct xucred`. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a `struct cmsgcred`. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a `struct unpcbid`. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a `struct sockpeercred`. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a `ucred_t`. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials interface {
	gextras.Objector

	// SetNative copies the native credentials of type @native_type from @native
	// into @credentials.
	//
	// It is a programming error (which will cause a warning to be logged) to
	// use this method if there is no #GCredentials support for the OS or if
	// @native_type isn't supported by the OS.
	SetNative(nativeType CredentialsType, native interface{})
	// String creates a human-readable textual representation of @credentials
	// that can be used in logging and debug messages. The format of the
	// returned string may change in future GLib release.
	String() string
}

// credentials implements the Credentials class.
type credentials struct {
	gextras.Objector
}

var _ Credentials = (*credentials)(nil)

// WrapCredentials wraps a GObject to the right type. It is
// primarily used internally.
func WrapCredentials(obj *externglib.Object) Credentials {
	return credentials{
		Objector: obj,
	}
}

func marshalCredentials(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCredentials(obj), nil
}

// SetNative copies the native credentials of type @native_type from @native
// into @credentials.
//
// It is a programming error (which will cause a warning to be logged) to
// use this method if there is no #GCredentials support for the OS or if
// @native_type isn't supported by the OS.
func (c credentials) SetNative(nativeType CredentialsType, native interface{}) {
	var _arg0 *C.GCredentials    // out
	var _arg1 C.GCredentialsType // out
	var _arg2 C.gpointer         // out

	_arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GCredentialsType)(nativeType)
	_arg2 = C.gpointer(native)

	C.g_credentials_set_native(_arg0, _arg1, _arg2)
}

// String creates a human-readable textual representation of @credentials
// that can be used in logging and debug messages. The format of the
// returned string may change in future GLib release.
func (c credentials) String() string {
	var _arg0 *C.GCredentials // out

	_arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.g_credentials_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
