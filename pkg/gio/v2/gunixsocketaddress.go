// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_unix_socket_address_get_type()), F: marshalUnixSocketAddress},
	})
}

// UnixSocketAddress: support for UNIX-domain (also known as local) sockets.
//
// UNIX domain sockets are generally visible in the filesystem. However, some
// systems support abstract socket names which are not visible in the filesystem
// and not affected by the filesystem permissions, visibility, etc. Currently
// this is only supported under Linux. If you attempt to use abstract sockets on
// other systems, function calls may return G_IO_ERROR_NOT_SUPPORTED errors. You
// can use g_unix_socket_address_abstract_names_supported() to see if abstract
// names are supported.
//
// Note that `<gio/gunixsocketaddress.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixSocketAddress interface {
	SocketAddress
	SocketConnectable

	// AddressType gets @address's type.
	AddressType() UnixSocketAddressType
	// IsAbstract tests if @address is abstract.
	IsAbstract() bool
	// Path gets @address's path, or for abstract sockets the "name".
	//
	// Guaranteed to be zero-terminated, but an abstract socket may contain
	// embedded zeros, and thus you should use
	// g_unix_socket_address_get_path_len() to get the true length of this
	// string.
	Path() string
	// PathLen gets the length of @address's path.
	//
	// For details, see g_unix_socket_address_get_path().
	PathLen() uint
}

// unixSocketAddress implements the UnixSocketAddress interface.
type unixSocketAddress struct {
	SocketAddress
	SocketConnectable
}

var _ UnixSocketAddress = (*unixSocketAddress)(nil)

// WrapUnixSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixSocketAddress(obj *externglib.Object) UnixSocketAddress {
	return UnixSocketAddress{
		SocketAddress:     WrapSocketAddress(obj),
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalUnixSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixSocketAddress(obj), nil
}

// NewUnixSocketAddress constructs a class UnixSocketAddress.
func NewUnixSocketAddress(path string) UnixSocketAddress {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.GUnixSocketAddress)
	var goret UnixSocketAddress

	cret = C.g_unix_socket_address_new(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixSocketAddress)

	return goret
}

// NewUnixSocketAddressAbstract constructs a class UnixSocketAddress.
func NewUnixSocketAddressAbstract() UnixSocketAddress {
	cret := new(C.GUnixSocketAddress)
	var goret UnixSocketAddress

	cret = C.g_unix_socket_address_new_abstract(arg1, arg2)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixSocketAddress)

	return goret
}

// AddressType gets @address's type.
func (a unixSocketAddress) AddressType() UnixSocketAddressType {
	var arg0 *C.GUnixSocketAddress

	arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(a.Native()))

	var cret C.GUnixSocketAddressType
	var goret UnixSocketAddressType

	cret = C.g_unix_socket_address_get_address_type(arg0)

	goret = UnixSocketAddressType(cret)

	return goret
}

// IsAbstract tests if @address is abstract.
func (a unixSocketAddress) IsAbstract() bool {
	var arg0 *C.GUnixSocketAddress

	arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_unix_socket_address_get_is_abstract(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Path gets @address's path, or for abstract sockets the "name".
//
// Guaranteed to be zero-terminated, but an abstract socket may contain
// embedded zeros, and thus you should use
// g_unix_socket_address_get_path_len() to get the true length of this
// string.
func (a unixSocketAddress) Path() string {
	var arg0 *C.GUnixSocketAddress

	arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(a.Native()))

	var cret *C.char
	var goret string

	cret = C.g_unix_socket_address_get_path(arg0)

	goret = C.GoString(cret)

	return goret
}

// PathLen gets the length of @address's path.
//
// For details, see g_unix_socket_address_get_path().
func (a unixSocketAddress) PathLen() uint {
	var arg0 *C.GUnixSocketAddress

	arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(a.Native()))

	var cret C.gsize
	var goret uint

	cret = C.g_unix_socket_address_get_path_len(arg0)

	goret = uint(cret)

	return goret
}
