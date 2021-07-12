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
		{T: externglib.Type(C.g_dtls_server_connection_get_type()), F: marshalDTLSServerConnectioner},
	})
}

// DTLSServerConnectioner describes DTLSServerConnection's methods.
type DTLSServerConnectioner interface {
	privateDTLSServerConnection()
}

// DTLSServerConnection is the server-side subclass of Connection, representing
// a server-side DTLS connection.
type DTLSServerConnection struct {
	DTLSConnection
}

var (
	_ DTLSServerConnectioner = (*DTLSServerConnection)(nil)
	_ gextras.Nativer        = (*DTLSServerConnection)(nil)
)

func wrapDTLSServerConnection(obj *externglib.Object) DTLSServerConnectioner {
	return &DTLSServerConnection{
		DTLSConnection: DTLSConnection{
			DatagramBased: DatagramBased{
				Object: obj,
			},
		},
	}
}

func marshalDTLSServerConnectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDTLSServerConnection(obj), nil
}

func (*DTLSServerConnection) privateDTLSServerConnection() {}

// NewDTLSServerConnection creates a new ServerConnection wrapping @base_socket.
func DtlsServerConnectionNew(baseSocket DatagramBaseder, certificate TLSCertificater) (*DTLSServerConnection, error) {
	var _arg1 *C.GDatagramBased  // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GDatagramBased  // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GDatagramBased)(unsafe.Pointer((baseSocket).(gextras.Nativer).Native()))
	_arg2 = (*C.GTlsCertificate)(unsafe.Pointer((certificate).(gextras.Nativer).Native()))

	_cret = C.g_dtls_server_connection_new(_arg1, _arg2, &_cerr)

	var _dtlsServerConnection *DTLSServerConnection // out
	var _goerr error                                // out

	_dtlsServerConnection = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DTLSServerConnection)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dtlsServerConnection, _goerr
}
