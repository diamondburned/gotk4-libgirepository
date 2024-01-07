// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeNetworkAddress = coreglib.Type(girepository.MustFind("Gio", "NetworkAddress").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNetworkAddress, F: marshalNetworkAddress},
	})
}

// NetworkAddressOverrides contains methods that are overridable.
type NetworkAddressOverrides struct {
}

func defaultNetworkAddressOverrides(v *NetworkAddress) NetworkAddressOverrides {
	return NetworkAddressOverrides{}
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
	_ [0]func() // equal guard
	*coreglib.Object

	SocketConnectable
}

var (
	_ coreglib.Objector = (*NetworkAddress)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NetworkAddress, *NetworkAddressClass, NetworkAddressOverrides](
		GTypeNetworkAddress,
		initNetworkAddressClass,
		wrapNetworkAddress,
		defaultNetworkAddressOverrides,
	)
}

func initNetworkAddressClass(gclass unsafe.Pointer, overrides NetworkAddressOverrides, classInitFunc func(*NetworkAddressClass)) {
	if classInitFunc != nil {
		class := (*NetworkAddressClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNetworkAddress(obj *coreglib.Object) *NetworkAddress {
	return &NetworkAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkAddress(p uintptr) (interface{}, error) {
	return wrapNetworkAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NetworkAddressClass: instance of this type is always passed by reference.
type NetworkAddressClass struct {
	*networkAddressClass
}

// networkAddressClass is the struct that's finalized.
type networkAddressClass struct {
	native unsafe.Pointer
}

var GIRInfoNetworkAddressClass = girepository.MustFind("Gio", "NetworkAddressClass")
