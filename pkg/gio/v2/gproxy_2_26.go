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
import "C"

// GType values.
var (
	GTypeProxy = coreglib.Type(girepository.MustFind("Gio", "Proxy").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeProxy, F: marshalProxy},
	})
}

// PROXY_EXTENSION_POINT_NAME: extension point for proxy functionality. See
// [Extending GIO][extending-gio].
const PROXY_EXTENSION_POINT_NAME = "gio-proxy"

// ProxyOverrider contains methods that are overridable.
type ProxyOverrider interface {
}

// Proxy handles connecting to a remote host via a given type of proxy server.
// It is implemented by the 'gio-proxy' extension point. The extensions are
// named after their proxy protocol name. As an example, a SOCKS5 proxy
// implementation can be retrieved with the name 'socks5' using the function
// g_io_extension_point_get_extension_by_name().
//
// Proxy wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Proxy struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Proxy)(nil)
)

// Proxier describes Proxy's interface methods.
type Proxier interface {
	coreglib.Objector

	baseProxy() *Proxy
}

var _ Proxier = (*Proxy)(nil)

func ifaceInitProxier(gifacePtr, data C.gpointer) {
}

func wrapProxy(obj *coreglib.Object) *Proxy {
	return &Proxy{
		Object: obj,
	}
}

func marshalProxy(p uintptr) (interface{}, error) {
	return wrapProxy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Proxy) baseProxy() *Proxy {
	return v
}

// BaseProxy returns the underlying base object.
func BaseProxy(obj Proxier) *Proxy {
	return obj.baseProxy()
}

// ProxyInterface provides an interface for handling proxy connection and
// payload.
//
// An instance of this type is always passed by reference.
type ProxyInterface struct {
	*proxyInterface
}

// proxyInterface is the struct that's finalized.
type proxyInterface struct {
	native unsafe.Pointer
}

var GIRInfoProxyInterface = girepository.MustFind("Gio", "ProxyInterface")
