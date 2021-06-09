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
		{T: externglib.Type(C.g_socket_address_get_type()), F: marshalSocketAddress},
	})
}

// SocketAddress is the equivalent of struct sockaddr in the BSD sockets API.
// This is an abstract class; use SocketAddress for internet sockets, or
// SocketAddress for UNIX domain sockets.
type SocketAddress interface {
	gextras.Objector
	SocketConnectable

	// Family gets the socket family type of @address.
	Family() SocketFamily
	// NativeSize gets the size of @address's native struct sockaddr. You can
	// use this to allocate memory to pass to g_socket_address_to_native().
	NativeSize() int
	// ToNative converts a Address to a native struct sockaddr, which can be
	// passed to low-level functions like connect() or bind().
	//
	// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
	// returned. If the address type is not known on the system then a
	// G_IO_ERROR_NOT_SUPPORTED error is returned.
	ToNative(dest interface{}, destlen uint) error
}

// socketAddress implements the SocketAddress interface.
type socketAddress struct {
	gextras.Objector
	SocketConnectable
}

var _ SocketAddress = (*socketAddress)(nil)

// WrapSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketAddress(obj *externglib.Object) SocketAddress {
	return SocketAddress{
		Objector:          obj,
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAddress(obj), nil
}

// NewSocketAddressFromNative constructs a class SocketAddress.
func NewSocketAddressFromNative(native interface{}, len uint) SocketAddress {
	var _arg1 C.gpointer
	var _arg2 C.gsize

	_arg1 = C.gpointer(native)
	_arg2 = C.gsize(len)

	var _cret C.GSocketAddress

	cret = C.g_socket_address_new_from_native(_arg1, _arg2)

	var _socketAddress SocketAddress

	_socketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SocketAddress)

	return _socketAddress
}

// Family gets the socket family type of @address.
func (a socketAddress) Family() SocketFamily {
	var _arg0 *C.GSocketAddress

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))

	var _cret C.GSocketFamily

	cret = C.g_socket_address_get_family(_arg0)

	var _socketFamily SocketFamily

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// NativeSize gets the size of @address's native struct sockaddr. You can
// use this to allocate memory to pass to g_socket_address_to_native().
func (a socketAddress) NativeSize() int {
	var _arg0 *C.GSocketAddress

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))

	var _cret C.gssize

	cret = C.g_socket_address_get_native_size(_arg0)

	var _gssize int

	_gssize = (int)(_cret)

	return _gssize
}

// ToNative converts a Address to a native struct sockaddr, which can be
// passed to low-level functions like connect() or bind().
//
// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
// returned. If the address type is not known on the system then a
// G_IO_ERROR_NOT_SUPPORTED error is returned.
func (a socketAddress) ToNative(dest interface{}, destlen uint) error {
	var _arg0 *C.GSocketAddress
	var _arg1 C.gpointer
	var _arg2 C.gsize

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(a.Native()))
	_arg1 = C.gpointer(dest)
	_arg2 = C.gsize(destlen)

	var _cerr *C.GError

	C.g_socket_address_to_native(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
