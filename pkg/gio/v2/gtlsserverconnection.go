// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_server_connection_get_type()), F: marshalTLSServerConnectioner},
	})
}

// TLSServerConnectioner describes TLSServerConnection's methods.
type TLSServerConnectioner interface {
	privateTLSServerConnection()
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
type TLSServerConnection struct {
	TLSConnection
}

var (
	_ TLSServerConnectioner = (*TLSServerConnection)(nil)
	_ gextras.Nativer       = (*TLSServerConnection)(nil)
)

func wrapTLSServerConnection(obj *externglib.Object) TLSServerConnectioner {
	return &TLSServerConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSServerConnectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSServerConnection(obj), nil
}

func (*TLSServerConnection) privateTLSServerConnection() {}

// NewTLSServerConnection creates a new ServerConnection wrapping
// @base_io_stream (which must have pollable input and output streams).
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the @base_io_stream after this
// function has returned.
func TlsServerConnectionNew(baseIoStream IOStreamer, certificate TLSCertificater) (*TLSServerConnection, error) {
	var _arg1 *C.GIOStream       // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GIOStream       // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer((baseIoStream).(gextras.Nativer).Native()))
	_arg2 = (*C.GTlsCertificate)(unsafe.Pointer((certificate).(gextras.Nativer).Native()))

	_cret = C.g_tls_server_connection_new(_arg1, _arg2, &_cerr)

	var _tlsServerConnection *TLSServerConnection // out
	var _goerr error                              // out

	_tlsServerConnection = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TLSServerConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsServerConnection, _goerr
}
