// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabaser},
	})
}

// TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT: purpose used to verify the client
// certificate in a TLS connection. Used by TLS servers.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT = "1.3.6.1.5.5.7.3.2"

// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER: purpose used to verify the server
// certificate in a TLS connection. This is the most common purpose in use. Used
// by TLS clients.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER = "1.3.6.1.5.5.7.3.1"

// TLSDatabaseOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSDatabaseOverrider interface {
	// CreateCertificateHandle: create a handle string for the certificate. The
	// database will only be able to create a handle for certificates that
	// originate from the database. In cases where the database cannot create a
	// handle for a certificate, NULL will be returned.
	//
	// This handle should be stable across various instances of the application,
	// and between applications. If a certificate is modified in the database,
	// then it is not guaranteed that this handle will continue to point to it.
	CreateCertificateHandle(certificate TLSCertificater) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	//
	// The handle should have been created by calling
	// g_tls_database_create_certificate_handle() on a Database object of the
	// same TLS backend. The handle is designed to remain valid across
	// instantiations of the database.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then NULL will be returned.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_for_handle_async() to perform the
	// lookup operation asynchronously.
	LookupCertificateForHandle(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	LookupCertificateForHandleAsync(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then NULL will be returned.
	LookupCertificateForHandleFinish(result AsyncResulter) (TLSCertificater, error)
	// LookupCertificateIssuer: look up the issuer of certificate in the
	// database.
	//
	// The Certificate:issuer property of certificate is not modified, and the
	// two certificates are not hooked into a chain.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
	// operation asynchronously.
	LookupCertificateIssuer(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// certificate in the database. See
	// g_tls_database_lookup_certificate_issuer() for more information.
	LookupCertificateIssuerAsync(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResulter) (TLSCertificater, error)
	// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
	// the database.
	//
	// This function can block, use
	// g_tls_database_lookup_certificates_issued_by_async() to perform the
	// lookup operation asynchronously.
	LookupCertificatesIssuedBy(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) ([]TLSCertificater, error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database. See
	// g_tls_database_lookup_certificates_issued_by() for more information.
	//
	// The database may choose to hold a reference to the issuer byte array for
	// the duration of of this asynchronous operation. The byte array should not
	// be modified during this time.
	LookupCertificatesIssuedByAsync(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
	// certificates. See g_tls_database_lookup_certificates_issued_by() for more
	// information.
	LookupCertificatesIssuedByFinish(result AsyncResulter) ([]TLSCertificater, error)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	//
	// chain is a chain of Certificate objects each pointing to the next
	// certificate in the chain by its Certificate:issuer property. The chain
	// may initially consist of one or more certificates. After the verification
	// process is complete, chain may be modified by adding missing
	// certificates, or removing extra certificates. If a certificate anchor was
	// found, then it is added to the chain.
	//
	// purpose describes the purpose (or usage) for which the certificate is
	// being used. Typically purpose will be set to
	// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
	// is being used to authenticate a server (and we are acting as the client).
	//
	// The identity is used to ensure the server certificate is valid for the
	// expected peer identity. If the identity does not match the certificate,
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
	// identity is NULL, that bit will never be set in the return value. The
	// peer identity may also be used to check for pinned certificates (trust
	// exceptions) in the database. These may override the normal verification
	// process on a host-by-host basis.
	//
	// Currently there are no flags, and G_TLS_DATABASE_VERIFY_NONE should be
	// used.
	//
	// If chain is found to be valid, then the return value will be 0. If chain
	// is found to be invalid, then the return value will indicate the problems
	// found. If the function is unable to determine whether chain is valid or
	// not (eg, because cancellable is triggered before it completes) then the
	// return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and error will be
	// set accordingly. error is not set when chain is successfully analyzed but
	// found to be invalid.
	//
	// This function can block, use g_tls_database_verify_chain_async() to
	// perform the verification operation asynchronously.
	VerifyChain(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags) (TLSCertificateFlags, error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	// See g_tls_database_verify_chain() for more information.
	VerifyChainAsync(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags, callback AsyncReadyCallback)
	// VerifyChainFinish: finish an asynchronous verify chain operation. See
	// g_tls_database_verify_chain() for more information.
	//
	// If chain is found to be valid, then the return value will be 0. If chain
	// is found to be invalid, then the return value will indicate the problems
	// found. If the function is unable to determine whether chain is valid or
	// not (eg, because cancellable is triggered before it completes) then the
	// return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and error will be
	// set accordingly. error is not set when chain is successfully analyzed but
	// found to be invalid.
	VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error)
}

// TLSDatabase is used to look up certificates and other information from a
// certificate or key store. It is an abstract base class which TLS library
// specific subtypes override.
//
// A Database may be accessed from multiple threads by the TLS backend. All
// implementations are required to be fully thread-safe.
//
// Most common client applications will not directly interact with Database. It
// is used internally by Connection.
type TLSDatabase struct {
	*externglib.Object
}

// TLSDatabaser describes TLSDatabase's abstract methods.
type TLSDatabaser interface {
	externglib.Objector

	// CreateCertificateHandle: create a handle string for the certificate.
	CreateCertificateHandle(certificate TLSCertificater) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	LookupCertificateForHandle(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database.
	LookupCertificateForHandleAsync(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle.
	LookupCertificateForHandleFinish(result AsyncResulter) (TLSCertificater, error)
	// LookupCertificateIssuer: look up the issuer of certificate in the
	// database.
	LookupCertificateIssuer(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// certificate in the database.
	LookupCertificateIssuerAsync(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation.
	LookupCertificateIssuerFinish(result AsyncResulter) (TLSCertificater, error)
	// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
	// the database.
	LookupCertificatesIssuedBy(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) ([]TLSCertificater, error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database.
	LookupCertificatesIssuedByAsync(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback)
	// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
	// certificates.
	LookupCertificatesIssuedByFinish(result AsyncResulter) ([]TLSCertificater, error)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	VerifyChain(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags) (TLSCertificateFlags, error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	VerifyChainAsync(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags, callback AsyncReadyCallback)
	// VerifyChainFinish: finish an asynchronous verify chain operation.
	VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error)
}

var _ TLSDatabaser = (*TLSDatabase)(nil)

func wrapTLSDatabase(obj *externglib.Object) *TLSDatabase {
	return &TLSDatabase{
		Object: obj,
	}
}

func marshalTLSDatabaser(p uintptr) (interface{}, error) {
	return wrapTLSDatabase(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CreateCertificateHandle: create a handle string for the certificate. The
// database will only be able to create a handle for certificates that originate
// from the database. In cases where the database cannot create a handle for a
// certificate, NULL will be returned.
//
// This handle should be stable across various instances of the application, and
// between applications. If a certificate is modified in the database, then it
// is not guaranteed that this handle will continue to point to it.
//
// The function takes the following parameters:
//
//    - certificate for which to create a handle.
//
func (self *TLSDatabase) CreateCertificateHandle(certificate TLSCertificater) string {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GTlsCertificate // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	_cret = C.g_tls_database_create_certificate_handle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(certificate)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// LookupCertificateForHandle: look up a certificate by its handle.
//
// The handle should have been created by calling
// g_tls_database_create_certificate_handle() on a Database object of the same
// TLS backend. The handle is designed to remain valid across instantiations of
// the database.
//
// If the handle is no longer valid, or does not point to a certificate in this
// database, then NULL will be returned.
//
// This function can block, use
// g_tls_database_lookup_certificate_for_handle_async() to perform the lookup
// operation asynchronously.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - handle: certificate handle.
//    - interaction: used to interact with the user if necessary.
//    - flags flags which affect the lookup.
//
func (self *TLSDatabase) LookupCertificateForHandle(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	defer C.free(unsafe.Pointer(_arg1))
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)

	_cret = C.g_tls_database_lookup_certificate_for_handle(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(handle)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)

	var _tlsCertificate TLSCertificater // out
	var _goerr error                    // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSCertificater)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// LookupCertificateForHandleAsync: asynchronously look up a certificate by its
// handle in the database. See g_tls_database_lookup_certificate_for_handle()
// for more information.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - handle: certificate handle.
//    - interaction: used to interact with the user if necessary.
//    - flags flags which affect the lookup.
//    - callback to call when the operation completes.
//
func (self *TLSDatabase) LookupCertificateForHandleAsync(ctx context.Context, handle string, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	defer C.free(unsafe.Pointer(_arg1))
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_database_lookup_certificate_for_handle_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(handle)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
// certificate by its handle. See g_tls_database_lookup_certificate_for_handle()
// for more information.
//
// If the handle is no longer valid, or does not point to a certificate in this
// database, then NULL will be returned.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (self *TLSDatabase) LookupCertificateForHandleFinish(result AsyncResulter) (TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _tlsCertificate TLSCertificater // out
	var _goerr error                    // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSCertificater is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(TLSCertificater)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
		}
		_tlsCertificate = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuer: look up the issuer of certificate in the database.
//
// The Certificate:issuer property of certificate is not modified, and the two
// certificates are not hooked into a chain.
//
// This function can block, use g_tls_database_lookup_certificate_issuer_async()
// to perform the lookup operation asynchronously.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - certificate: Certificate.
//    - interaction: used to interact with the user if necessary.
//    - flags which affect the lookup operation.
//
func (self *TLSDatabase) LookupCertificateIssuer(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) (TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)

	_cret = C.g_tls_database_lookup_certificate_issuer(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(certificate)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)

	var _tlsCertificate TLSCertificater // out
	var _goerr error                    // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSCertificater is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(TLSCertificater)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
		}
		_tlsCertificate = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuerAsync: asynchronously look up the issuer of
// certificate in the database. See g_tls_database_lookup_certificate_issuer()
// for more information.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - certificate: Certificate.
//    - interaction: used to interact with the user if necessary.
//    - flags which affect the lookup operation.
//    - callback to call when the operation completes.
//
func (self *TLSDatabase) LookupCertificateIssuerAsync(ctx context.Context, certificate TLSCertificater, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_database_lookup_certificate_issuer_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(certificate)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
// operation. See g_tls_database_lookup_certificate_issuer() for more
// information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (self *TLSDatabase) LookupCertificateIssuerFinish(result AsyncResulter) (TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _tlsCertificate TLSCertificater // out
	var _goerr error                    // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSCertificater is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(TLSCertificater)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
		}
		_tlsCertificate = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// LookupCertificatesIssuedBy: look up certificates issued by this issuer in the
// database.
//
// This function can block, use
// g_tls_database_lookup_certificates_issued_by_async() to perform the lookup
// operation asynchronously.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - issuerRawDn which holds the DER encoded issuer DN.
//    - interaction: used to interact with the user if necessary.
//    - flags flags which affect the lookup operation.
//
func (self *TLSDatabase) LookupCertificatesIssuedBy(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags) ([]TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.GByteArray             // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _cret *C.GList                  // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.g_byte_array_sized_new(C.guint(len(issuerRawDn)))
	if len(issuerRawDn) > 0 {
		_arg1 = C.g_byte_array_append(_arg1, (*C.guint8)(&issuerRawDn[0]), C.guint(len(issuerRawDn)))
	}
	defer C.g_byte_array_unref(_arg1)
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)

	_cret = C.g_tls_database_lookup_certificates_issued_by(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(issuerRawDn)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)

	var _list []TLSCertificater // out
	var _goerr error            // out

	_list = make([]TLSCertificater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GTlsCertificate)(v)
		var dst TLSCertificater // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.TLSCertificater is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSCertificater)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupCertificatesIssuedByAsync: asynchronously look up certificates issued
// by this issuer in the database. See
// g_tls_database_lookup_certificates_issued_by() for more information.
//
// The database may choose to hold a reference to the issuer byte array for the
// duration of of this asynchronous operation. The byte array should not be
// modified during this time.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - issuerRawDn which holds the DER encoded issuer DN.
//    - interaction: used to interact with the user if necessary.
//    - flags flags which affect the lookup operation.
//    - callback to call when the operation completes.
//
func (self *TLSDatabase) LookupCertificatesIssuedByAsync(ctx context.Context, issuerRawDn []byte, interaction *TLSInteraction, flags TLSDatabaseLookupFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg4 *C.GCancellable           // out
	var _arg1 *C.GByteArray             // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.g_byte_array_sized_new(C.guint(len(issuerRawDn)))
	if len(issuerRawDn) > 0 {
		_arg1 = C.g_byte_array_append(_arg1, (*C.guint8)(&issuerRawDn[0]), C.guint(len(issuerRawDn)))
	}
	defer C.g_byte_array_unref(_arg1)
	if interaction != nil {
		_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_database_lookup_certificates_issued_by_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(issuerRawDn)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
// certificates. See g_tls_database_lookup_certificates_issued_by() for more
// information.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (self *TLSDatabase) LookupCertificatesIssuedByFinish(result AsyncResulter) ([]TLSCertificater, error) {
	var _arg0 *C.GTlsDatabase // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_lookup_certificates_issued_by_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _list []TLSCertificater // out
	var _goerr error            // out

	_list = make([]TLSCertificater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GTlsCertificate)(v)
		var dst TLSCertificater // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.TLSCertificater is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(TLSCertificater)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSCertificater")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// VerifyChain determines the validity of a certificate chain after looking up
// and adding any missing certificates to the chain.
//
// chain is a chain of Certificate objects each pointing to the next certificate
// in the chain by its Certificate:issuer property. The chain may initially
// consist of one or more certificates. After the verification process is
// complete, chain may be modified by adding missing certificates, or removing
// extra certificates. If a certificate anchor was found, then it is added to
// the chain.
//
// purpose describes the purpose (or usage) for which the certificate is being
// used. Typically purpose will be set to
// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate is
// being used to authenticate a server (and we are acting as the client).
//
// The identity is used to ensure the server certificate is valid for the
// expected peer identity. If the identity does not match the certificate,
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If identity
// is NULL, that bit will never be set in the return value. The peer identity
// may also be used to check for pinned certificates (trust exceptions) in the
// database. These may override the normal verification process on a
// host-by-host basis.
//
// Currently there are no flags, and G_TLS_DATABASE_VERIFY_NONE should be used.
//
// If chain is found to be valid, then the return value will be 0. If chain is
// found to be invalid, then the return value will indicate the problems found.
// If the function is unable to determine whether chain is valid or not (eg,
// because cancellable is triggered before it completes) then the return value
// will be G_TLS_CERTIFICATE_GENERIC_ERROR and error will be set accordingly.
// error is not set when chain is successfully analyzed but found to be invalid.
//
// This function can block, use g_tls_database_verify_chain_async() to perform
// the verification operation asynchronously.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - chain: Certificate chain.
//    - purpose that this certificate chain will be used for.
//    - identity: expected peer identity.
//    - interaction: used to interact with the user if necessary.
//    - flags: additional verify flags.
//
func (self *TLSDatabase) VerifyChain(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg6 *C.GCancellable           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _cret C.GTlsCertificateFlags    // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	defer C.free(unsafe.Pointer(_arg2))
	if identity != nil {
		_arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	}
	if interaction != nil {
		_arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)

	_cret = C.g_tls_database_verify_chain(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(chain)
	runtime.KeepAlive(purpose)
	runtime.KeepAlive(identity)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificateFlags, _goerr
}

// VerifyChainAsync: asynchronously determines the validity of a certificate
// chain after looking up and adding any missing certificates to the chain. See
// g_tls_database_verify_chain() for more information.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - chain: Certificate chain.
//    - purpose that this certificate chain will be used for.
//    - identity: expected peer identity.
//    - interaction: used to interact with the user if necessary.
//    - flags: additional verify flags.
//    - callback to call when the operation completes.
//
func (self *TLSDatabase) VerifyChainAsync(ctx context.Context, chain TLSCertificater, purpose string, identity SocketConnectabler, interaction *TLSInteraction, flags TLSDatabaseVerifyFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg6 *C.GCancellable           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _arg7 C.GAsyncReadyCallback     // out
	var _arg8 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	defer C.free(unsafe.Pointer(_arg2))
	if identity != nil {
		_arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	}
	if interaction != nil {
		_arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	}
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)
	if callback != nil {
		_arg7 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg8 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_tls_database_verify_chain_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(chain)
	runtime.KeepAlive(purpose)
	runtime.KeepAlive(identity)
	runtime.KeepAlive(interaction)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// VerifyChainFinish: finish an asynchronous verify chain operation. See
// g_tls_database_verify_chain() for more information.
//
// If chain is found to be valid, then the return value will be 0. If chain is
// found to be invalid, then the return value will indicate the problems found.
// If the function is unable to determine whether chain is valid or not (eg,
// because cancellable is triggered before it completes) then the return value
// will be G_TLS_CERTIFICATE_GENERIC_ERROR and error will be set accordingly.
// error is not set when chain is successfully analyzed but found to be invalid.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (self *TLSDatabase) VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase        // out
	var _arg1 *C.GAsyncResult        // out
	var _cret C.GTlsCertificateFlags // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_database_verify_chain_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificateFlags, _goerr
}
