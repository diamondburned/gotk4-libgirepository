// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_network_address_get_type()), F: marshalNetworkAddresser},
	})
}

// NetworkAddresser describes NetworkAddress's methods.
type NetworkAddresser interface {
	// Hostname gets @addr's hostname.
	Hostname() string
	// Port gets @addr's port number
	Port() uint16
	// Scheme gets @addr's scheme
	Scheme() string
}

// NetworkAddress provides an easy way to resolve a hostname and then attempt to
// connect to that host, handling the possibility of multiple IP addresses and
// multiple address families.
//
// The enumeration results of resolved addresses *may* be cached as long as this
// object is kept alive which may have unexpected results if alive for too long.
//
// See Connectable for an example of using the connectable interface.
type NetworkAddress struct {
	*externglib.Object

	SocketConnectable
}

var (
	_ NetworkAddresser = (*NetworkAddress)(nil)
	_ gextras.Nativer  = (*NetworkAddress)(nil)
)

func wrapNetworkAddress(obj *externglib.Object) NetworkAddresser {
	return &NetworkAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNetworkAddress(obj), nil
}

// NewNetworkAddress creates a new Connectable for connecting to the given
// @hostname and @port.
//
// Note that depending on the configuration of the machine, a @hostname of
// `localhost` may refer to the IPv4 loopback address only, or to both IPv4 and
// IPv6; use g_network_address_new_loopback() to create a Address that is
// guaranteed to resolve to both addresses.
func NewNetworkAddress(hostname string, port uint16) *NetworkAddress {
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _cret *C.GSocketConnectable // in

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(port)

	_cret = C.g_network_address_new(_arg1, _arg2)

	var _networkAddress *NetworkAddress // out

	_networkAddress = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NetworkAddress)

	return _networkAddress
}

// NewNetworkAddressLoopback creates a new Connectable for connecting to the
// local host over a loopback connection to the given @port. This is intended
// for use in connecting to local services which may be running on IPv4 or IPv6.
//
// The connectable will return IPv4 and IPv6 loopback addresses, regardless of
// how the host resolves `localhost`. By contrast, g_network_address_new() will
// often only return an IPv4 address when resolving `localhost`, and an IPv6
// address for `localhost6`.
//
// g_network_address_get_hostname() will always return `localhost` for a Address
// created with this constructor.
func NewNetworkAddressLoopback(port uint16) *NetworkAddress {
	var _arg1 C.guint16             // out
	var _cret *C.GSocketConnectable // in

	_arg1 = C.guint16(port)

	_cret = C.g_network_address_new_loopback(_arg1)

	var _networkAddress *NetworkAddress // out

	_networkAddress = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NetworkAddress)

	return _networkAddress
}

// Hostname gets @addr's hostname. This might be either UTF-8 or ASCII-encoded,
// depending on what @addr was created with.
func (addr *NetworkAddress) Hostname() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_hostname(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// Port gets @addr's port number
func (addr *NetworkAddress) Port() uint16 {
	var _arg0 *C.GNetworkAddress // out
	var _cret C.guint16          // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Scheme gets @addr's scheme
func (addr *NetworkAddress) Scheme() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_scheme(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}
