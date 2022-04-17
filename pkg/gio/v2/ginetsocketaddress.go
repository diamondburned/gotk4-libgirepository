// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for ginetsocketaddress.go.
var GTypeInetSocketAddress = externglib.Type(C.g_inet_socket_address_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeInetSocketAddress, F: marshalInetSocketAddress},
	})
}

// InetSocketAddressOverrider contains methods that are overridable.
type InetSocketAddressOverrider interface {
	externglib.Objector
}

// InetSocketAddress: IPv4 or IPv6 socket address; that is, the combination of a
// Address and a port number.
type InetSocketAddress struct {
	_ [0]func() // equal guard
	SocketAddress
}

var (
	_ SocketAddresser = (*InetSocketAddress)(nil)
)

func classInitInetSocketAddresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapInetSocketAddress(obj *externglib.Object) *InetSocketAddress {
	return &InetSocketAddress{
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
	}
}

func marshalInetSocketAddress(p uintptr) (interface{}, error) {
	return wrapInetSocketAddress(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewInetSocketAddress creates a new SocketAddress for address and port.
//
// The function takes the following parameters:
//
//    - address: Address.
//    - port number.
//
// The function returns the following values:
//
//    - inetSocketAddress: new SocketAddress.
//
func NewInetSocketAddress(address *InetAddress, port uint16) *InetSocketAddress {
	var _arg1 *C.GInetAddress   // out
	var _arg2 C.guint16         // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(externglib.InternObject(address).Native()))
	_arg2 = C.guint16(port)

	_cret = C.g_inet_socket_address_new(_arg1, _arg2)
	runtime.KeepAlive(address)
	runtime.KeepAlive(port)

	var _inetSocketAddress *InetSocketAddress // out

	_inetSocketAddress = wrapInetSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetSocketAddress
}

// NewInetSocketAddressFromString creates a new SocketAddress for address and
// port.
//
// If address is an IPv6 address, it can also contain a scope ID (separated from
// the address by a %).
//
// The function takes the following parameters:
//
//    - address: string form of an IP address.
//    - port number.
//
// The function returns the following values:
//
//    - inetSocketAddress (optional): new SocketAddress, or NULL if address
//      cannot be parsed.
//
func NewInetSocketAddressFromString(address string, port uint) *InetSocketAddress {
	var _arg1 *C.char           // out
	var _arg2 C.guint           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(address)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(port)

	_cret = C.g_inet_socket_address_new_from_string(_arg1, _arg2)
	runtime.KeepAlive(address)
	runtime.KeepAlive(port)

	var _inetSocketAddress *InetSocketAddress // out

	if _cret != nil {
		_inetSocketAddress = wrapInetSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _inetSocketAddress
}

// Address gets address's Address.
//
// The function returns the following values:
//
//    - inetAddress for address, which must be g_object_ref()'d if it will be
//      stored.
//
func (address *InetSocketAddress) Address() *InetAddress {
	var _arg0 *C.GInetSocketAddress // out
	var _cret *C.GInetAddress       // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(externglib.InternObject(address).Native()))

	_cret = C.g_inet_socket_address_get_address(_arg0)
	runtime.KeepAlive(address)

	var _inetAddress *InetAddress // out

	_inetAddress = wrapInetAddress(externglib.Take(unsafe.Pointer(_cret)))

	return _inetAddress
}

// Flowinfo gets the sin6_flowinfo field from address, which must be an IPv6
// address.
//
// The function returns the following values:
//
//    - guint32: flowinfo field.
//
func (address *InetSocketAddress) Flowinfo() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(externglib.InternObject(address).Native()))

	_cret = C.g_inet_socket_address_get_flowinfo(_arg0)
	runtime.KeepAlive(address)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Port gets address's port.
//
// The function returns the following values:
//
//    - guint16: port for address.
//
func (address *InetSocketAddress) Port() uint16 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint16             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(externglib.InternObject(address).Native()))

	_cret = C.g_inet_socket_address_get_port(_arg0)
	runtime.KeepAlive(address)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// ScopeID gets the sin6_scope_id field from address, which must be an IPv6
// address.
//
// The function returns the following values:
//
//    - guint32: scope id field.
//
func (address *InetSocketAddress) ScopeID() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(externglib.InternObject(address).Native()))

	_cret = C.g_inet_socket_address_get_scope_id(_arg0)
	runtime.KeepAlive(address)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}
