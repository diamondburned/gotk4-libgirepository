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
// extern GTlsDatabase* _gotk4_gio2_TlsBackendInterface_get_default_database(GTlsBackend*);
// extern gboolean _gotk4_gio2_TlsBackendInterface_supports_dtls(GTlsBackend*);
// extern gboolean _gotk4_gio2_TlsBackendInterface_supports_tls(GTlsBackend*);
import "C"

// glib.Type values for gtlsbackend.go.
var GTypeTLSBackend = externglib.Type(C.g_tls_backend_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTLSBackend, F: marshalTLSBackend},
	})
}

// TLS_BACKEND_EXTENSION_POINT_NAME: extension point for TLS functionality via
// Backend. See [Extending GIO][extending-gio].
const TLS_BACKEND_EXTENSION_POINT_NAME = "gio-tls-backend"

// TLSBackendOverrider contains methods that are overridable.
type TLSBackendOverrider interface {
	// DefaultDatabase gets the default Database used to verify TLS connections.
	//
	// The function returns the following values:
	//
	//    - tlsDatabase: default database, which should be unreffed when done.
	//
	DefaultDatabase() TLSDatabaser
	// SupportsDTLS checks if DTLS is supported. DTLS support may not be
	// available even if TLS support is available, and vice-versa.
	//
	// The function returns the following values:
	//
	//    - ok: whether DTLS is supported.
	//
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	//
	// The function returns the following values:
	//
	//    - ok: whether or not TLS is supported.
	//
	SupportsTLS() bool
}

// TLSBackend: TLS (Transport Layer Security, aka SSL) and DTLS backend.
//
// TLSBackend wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSBackend struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TLSBackend)(nil)
)

// TLSBackender describes TLSBackend's interface methods.
type TLSBackender interface {
	externglib.Objector

	// CertificateType gets the #GType of backend's Certificate implementation.
	CertificateType() externglib.Type
	// ClientConnectionType gets the #GType of backend's ClientConnection
	// implementation.
	ClientConnectionType() externglib.Type
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() TLSDatabaser
	// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
	// implementation.
	DTLSClientConnectionType() externglib.Type
	// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
	// implementation.
	DTLSServerConnectionType() externglib.Type
	// FileDatabaseType gets the #GType of backend's FileDatabase
	// implementation.
	FileDatabaseType() externglib.Type
	// ServerConnectionType gets the #GType of backend's ServerConnection
	// implementation.
	ServerConnectionType() externglib.Type
	// SetDefaultDatabase: set the default Database used to verify TLS
	// connections Any subsequent call to g_tls_backend_get_default_database()
	// will return the database set in this call.
	SetDefaultDatabase(database TLSDatabaser)
	// SupportsDTLS checks if DTLS is supported.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

var _ TLSBackender = (*TLSBackend)(nil)

func ifaceInitTLSBackender(gifacePtr, data C.gpointer) {
	iface := (*C.GTlsBackendInterface)(unsafe.Pointer(gifacePtr))
	iface.get_default_database = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_get_default_database)
	iface.supports_dtls = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_supports_dtls)
	iface.supports_tls = (*[0]byte)(C._gotk4_gio2_TlsBackendInterface_supports_tls)
}

//export _gotk4_gio2_TlsBackendInterface_get_default_database
func _gotk4_gio2_TlsBackendInterface_get_default_database(arg0 *C.GTlsBackend) (cret *C.GTlsDatabase) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	tlsDatabase := iface.DefaultDatabase()

	cret = (*C.GTlsDatabase)(unsafe.Pointer(externglib.InternObject(tlsDatabase).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(tlsDatabase).Native()))

	return cret
}

//export _gotk4_gio2_TlsBackendInterface_supports_dtls
func _gotk4_gio2_TlsBackendInterface_supports_dtls(arg0 *C.GTlsBackend) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	ok := iface.SupportsDTLS()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_TlsBackendInterface_supports_tls
func _gotk4_gio2_TlsBackendInterface_supports_tls(arg0 *C.GTlsBackend) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TLSBackendOverrider)

	ok := iface.SupportsTLS()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTLSBackend(obj *externglib.Object) *TLSBackend {
	return &TLSBackend{
		Object: obj,
	}
}

func marshalTLSBackend(p uintptr) (interface{}, error) {
	return wrapTLSBackend(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CertificateType gets the #GType of backend's Certificate implementation.
//
// The function returns the following values:
//
//    - gType of backend's Certificate implementation.
//
func (backend *TLSBackend) CertificateType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_certificate_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ClientConnectionType gets the #GType of backend's ClientConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend's ClientConnection implementation.
//
func (backend *TLSBackend) ClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DefaultDatabase gets the default Database used to verify TLS connections.
//
// The function returns the following values:
//
//    - tlsDatabase: default database, which should be unreffed when done.
//
func (backend *TLSBackend) DefaultDatabase() TLSDatabaser {
	var _arg0 *C.GTlsBackend  // out
	var _cret *C.GTlsDatabase // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_default_database(_arg0)
	runtime.KeepAlive(backend)

	var _tlsDatabase TLSDatabaser // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSDatabaser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TLSDatabaser)
			return ok
		})
		rv, ok := casted.(TLSDatabaser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSDatabaser")
		}
		_tlsDatabase = rv
	}

	return _tlsDatabase
}

// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend’s ClientConnection implementation, or G_TYPE_INVALID if
//      this backend doesn’t support DTLS.
//
func (backend *TLSBackend) DTLSClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_dtls_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend’s ServerConnection implementation, or G_TYPE_INVALID if
//      this backend doesn’t support DTLS.
//
func (backend *TLSBackend) DTLSServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_dtls_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// FileDatabaseType gets the #GType of backend's FileDatabase implementation.
//
// The function returns the following values:
//
//    - gType of backend's FileDatabase implementation.
//
func (backend *TLSBackend) FileDatabaseType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_file_database_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ServerConnectionType gets the #GType of backend's ServerConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend's ServerConnection implementation.
//
func (backend *TLSBackend) ServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
//
// The function takes the following parameters:
//
//    - database (optional): Database.
//
func (backend *TLSBackend) SetDefaultDatabase(database TLSDatabaser) {
	var _arg0 *C.GTlsBackend  // out
	var _arg1 *C.GTlsDatabase // out

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))
	if database != nil {
		_arg1 = (*C.GTlsDatabase)(unsafe.Pointer(externglib.InternObject(database).Native()))
	}

	C.g_tls_backend_set_default_database(_arg0, _arg1)
	runtime.KeepAlive(backend)
	runtime.KeepAlive(database)
}

// SupportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
//
// The function returns the following values:
//
//    - ok: whether DTLS is supported.
//
func (backend *TLSBackend) SupportsDTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_supports_dtls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsTLS checks if TLS is supported; if this returns FALSE for the default
// Backend, it means no "real" TLS backend is available.
//
// The function returns the following values:
//
//    - ok: whether or not TLS is supported.
//
func (backend *TLSBackend) SupportsTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(externglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_supports_tls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TLSBackendGetDefault gets the default Backend for the system.
//
// The function returns the following values:
//
//    - tlsBackend which will be a dummy object if no TLS backend is available.
//
func TLSBackendGetDefault() TLSBackender {
	var _cret *C.GTlsBackend // in

	_cret = C.g_tls_backend_get_default()

	var _tlsBackend TLSBackender // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSBackender is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TLSBackender)
			return ok
		})
		rv, ok := casted.(TLSBackender)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSBackender")
		}
		_tlsBackend = rv
	}

	return _tlsBackend
}
