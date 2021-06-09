// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabase},
	})
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
type TLSDatabase interface {
	gextras.Objector

	// CreateCertificateHandle: create a handle string for the certificate. The
	// database will only be able to create a handle for certificates that
	// originate from the database. In cases where the database cannot create a
	// handle for a certificate, nil will be returned.
	//
	// This handle should be stable across various instances of the application,
	// and between applications. If a certificate is modified in the database,
	// then it is not guaranteed that this handle will continue to point to it.
	CreateCertificateHandle(certificate TLSCertificate) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	//
	// The handle should have been created by calling
	// g_tls_database_create_certificate_handle() on a Database object of the
	// same TLS backend. The handle is designed to remain valid across
	// instantiations of the database.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_for_handle_async() to perform the
	// lookup operation asynchronously.
	LookupCertificateForHandle(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	LookupCertificateForHandleAsync()
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	LookupCertificateForHandleFinish(result AsyncResult) (TLSCertificate, error)
	// LookupCertificateIssuer: look up the issuer of @certificate in the
	// database.
	//
	// The Certificate:issuer property of @certificate is not modified, and the
	// two certificates are not hooked into a chain.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
	// operation asynchronously.
	LookupCertificateIssuer(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// @certificate in the database. See
	// g_tls_database_lookup_certificate_issuer() for more information.
	LookupCertificateIssuerAsync()
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResult) (TLSCertificate, error)
	// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
	// the database.
	//
	// This function can block, use
	// g_tls_database_lookup_certificates_issued_by_async() to perform the
	// lookup operation asynchronously.
	LookupCertificatesIssuedBy(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (*glib.List, error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database. See
	// g_tls_database_lookup_certificates_issued_by() for more information.
	//
	// The database may choose to hold a reference to the issuer byte array for
	// the duration of of this asynchronous operation. The byte array should not
	// be modified during this time.
	LookupCertificatesIssuedByAsync()
	// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
	// certificates. See g_tls_database_lookup_certificates_issued_by() for more
	// information.
	LookupCertificatesIssuedByFinish(result AsyncResult) (*glib.List, error)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	//
	// @chain is a chain of Certificate objects each pointing to the next
	// certificate in the chain by its Certificate:issuer property. The chain
	// may initially consist of one or more certificates. After the verification
	// process is complete, @chain may be modified by adding missing
	// certificates, or removing extra certificates. If a certificate anchor was
	// found, then it is added to the @chain.
	//
	// @purpose describes the purpose (or usage) for which the certificate is
	// being used. Typically @purpose will be set to
	// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
	// is being used to authenticate a server (and we are acting as the client).
	//
	// The @identity is used to ensure the server certificate is valid for the
	// expected peer identity. If the identity does not match the certificate,
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
	// @identity is nil, that bit will never be set in the return value. The
	// peer identity may also be used to check for pinned certificates (trust
	// exceptions) in the database. These may override the normal verification
	// process on a host-by-host basis.
	//
	// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be
	// used.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	//
	// This function can block, use g_tls_database_verify_chain_async() to
	// perform the verification operation asynchronously.
	VerifyChain(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (TLSCertificateFlags, error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	// See g_tls_database_verify_chain() for more information.
	VerifyChainAsync()
	// VerifyChainFinish: finish an asynchronous verify chain operation. See
	// g_tls_database_verify_chain() for more information.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	VerifyChainFinish(result AsyncResult) (TLSCertificateFlags, error)
}

// tlsDatabase implements the TLSDatabase interface.
type tlsDatabase struct {
	gextras.Objector
}

var _ TLSDatabase = (*tlsDatabase)(nil)

// WrapTLSDatabase wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSDatabase(obj *externglib.Object) TLSDatabase {
	return TLSDatabase{
		Objector: obj,
	}
}

func marshalTLSDatabase(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSDatabase(obj), nil
}

// CreateCertificateHandle: create a handle string for the certificate. The
// database will only be able to create a handle for certificates that
// originate from the database. In cases where the database cannot create a
// handle for a certificate, nil will be returned.
//
// This handle should be stable across various instances of the application,
// and between applications. If a certificate is modified in the database,
// then it is not guaranteed that this handle will continue to point to it.
func (s tlsDatabase) CreateCertificateHandle(certificate TLSCertificate) string {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GTlsCertificate

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	var _cret *C.gchar

	cret = C.g_tls_database_create_certificate_handle(_arg0, _arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LookupCertificateForHandle: look up a certificate by its handle.
//
// The handle should have been created by calling
// g_tls_database_create_certificate_handle() on a Database object of the
// same TLS backend. The handle is designed to remain valid across
// instantiations of the database.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then nil will be returned.
//
// This function can block, use
// g_tls_database_lookup_certificate_for_handle_async() to perform the
// lookup operation asynchronously.
func (s tlsDatabase) LookupCertificateForHandle(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.gchar
	var _arg2 *C.GTlsInteraction
	var _arg3 C.GTlsDatabaseLookupFlags
	var _arg4 *C.GCancellable

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(handle))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GTlsCertificate
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificate_for_handle(_arg0, _arg1, _arg2, _arg3, _arg4, _cerr)

	var _tlsCertificate TLSCertificate
	var _goerr error

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateForHandleAsync: asynchronously look up a certificate by
// its handle in the database. See
// g_tls_database_lookup_certificate_for_handle() for more information.
func (s tlsDatabase) LookupCertificateForHandleAsync() {
	var _arg0 *C.GTlsDatabase

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificate_for_handle_async(_arg0)
}

// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
// certificate by its handle. See
// g_tls_database_lookup_certificate_for_handle() for more information.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then nil will be returned.
func (s tlsDatabase) LookupCertificateForHandleFinish(result AsyncResult) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GTlsCertificate
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificate_for_handle_finish(_arg0, _arg1, _cerr)

	var _tlsCertificate TLSCertificate
	var _goerr error

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuer: look up the issuer of @certificate in the
// database.
//
// The Certificate:issuer property of @certificate is not modified, and the
// two certificates are not hooked into a chain.
//
// This function can block, use
// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
// operation asynchronously.
func (s tlsDatabase) LookupCertificateIssuer(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GTlsCertificate
	var _arg2 *C.GTlsInteraction
	var _arg3 C.GTlsDatabaseLookupFlags
	var _arg4 *C.GCancellable

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GTlsCertificate
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificate_issuer(_arg0, _arg1, _arg2, _arg3, _arg4, _cerr)

	var _tlsCertificate TLSCertificate
	var _goerr error

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuerAsync: asynchronously look up the issuer of
// @certificate in the database. See
// g_tls_database_lookup_certificate_issuer() for more information.
func (s tlsDatabase) LookupCertificateIssuerAsync() {
	var _arg0 *C.GTlsDatabase

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificate_issuer_async(_arg0)
}

// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
// operation. See g_tls_database_lookup_certificate_issuer() for more
// information.
func (s tlsDatabase) LookupCertificateIssuerFinish(result AsyncResult) (TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GTlsCertificate
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificate_issuer_finish(_arg0, _arg1, _cerr)

	var _tlsCertificate TLSCertificate
	var _goerr error

	_tlsCertificate = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
// the database.
//
// This function can block, use
// g_tls_database_lookup_certificates_issued_by_async() to perform the
// lookup operation asynchronously.
func (s tlsDatabase) LookupCertificatesIssuedBy(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (*glib.List, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GByteArray
	var _arg2 *C.GTlsInteraction
	var _arg3 C.GTlsDatabaseLookupFlags
	var _arg4 *C.GCancellable

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GByteArray)(C.malloc((len(issuerRawDn) + 1) * C.sizeof_guint8))
	defer C.free(unsafe.Pointer(_arg1))

	{
		var out []C.guint8
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(_arg1), int(len(issuerRawDn)))

		for i := range issuerRawDn {
			_arg1 = C.guint8(issuerRawDn)
		}
	}
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret *C.GList
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificates_issued_by(_arg0, _arg1, _arg2, _arg3, _arg4, _cerr)

	var _list *glib.List
	var _goerr error

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _list, _goerr
}

// LookupCertificatesIssuedByAsync: asynchronously look up certificates
// issued by this issuer in the database. See
// g_tls_database_lookup_certificates_issued_by() for more information.
//
// The database may choose to hold a reference to the issuer byte array for
// the duration of of this asynchronous operation. The byte array should not
// be modified during this time.
func (s tlsDatabase) LookupCertificatesIssuedByAsync() {
	var _arg0 *C.GTlsDatabase

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificates_issued_by_async(_arg0)
}

// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
// certificates. See g_tls_database_lookup_certificates_issued_by() for more
// information.
func (s tlsDatabase) LookupCertificatesIssuedByFinish(result AsyncResult) (*glib.List, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret *C.GList
	var _cerr *C.GError

	cret = C.g_tls_database_lookup_certificates_issued_by_finish(_arg0, _arg1, _cerr)

	var _list *glib.List
	var _goerr error

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _list, _goerr
}

// VerifyChain determines the validity of a certificate chain after looking
// up and adding any missing certificates to the chain.
//
// @chain is a chain of Certificate objects each pointing to the next
// certificate in the chain by its Certificate:issuer property. The chain
// may initially consist of one or more certificates. After the verification
// process is complete, @chain may be modified by adding missing
// certificates, or removing extra certificates. If a certificate anchor was
// found, then it is added to the @chain.
//
// @purpose describes the purpose (or usage) for which the certificate is
// being used. Typically @purpose will be set to
// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
// is being used to authenticate a server (and we are acting as the client).
//
// The @identity is used to ensure the server certificate is valid for the
// expected peer identity. If the identity does not match the certificate,
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
// @identity is nil, that bit will never be set in the return value. The
// peer identity may also be used to check for pinned certificates (trust
// exceptions) in the database. These may override the normal verification
// process on a host-by-host basis.
//
// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be
// used.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate the
// problems found. If the function is unable to determine whether @chain is
// valid or not (eg, because @cancellable is triggered before it completes)
// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
// will be set accordingly. @error is not set when @chain is successfully
// analyzed but found to be invalid.
//
// This function can block, use g_tls_database_verify_chain_async() to
// perform the verification operation asynchronously.
func (s tlsDatabase) VerifyChain(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GTlsCertificate
	var _arg2 *C.gchar
	var _arg3 *C.GSocketConnectable
	var _arg4 *C.GTlsInteraction
	var _arg5 C.GTlsDatabaseVerifyFlags
	var _arg6 *C.GCancellable

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	_arg2 = (*C.gchar)(C.CString(purpose))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	_arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg5 = (C.GTlsDatabaseVerifyFlags)(flags)
	_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret C.GTlsCertificateFlags
	var _cerr *C.GError

	cret = C.g_tls_database_verify_chain(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _cerr)

	var _tlsCertificateFlags TLSCertificateFlags
	var _goerr error

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}

// VerifyChainAsync: asynchronously determines the validity of a certificate
// chain after looking up and adding any missing certificates to the chain.
// See g_tls_database_verify_chain() for more information.
func (s tlsDatabase) VerifyChainAsync() {
	var _arg0 *C.GTlsDatabase

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_verify_chain_async(_arg0)
}

// VerifyChainFinish: finish an asynchronous verify chain operation. See
// g_tls_database_verify_chain() for more information.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate the
// problems found. If the function is unable to determine whether @chain is
// valid or not (eg, because @cancellable is triggered before it completes)
// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
// will be set accordingly. @error is not set when @chain is successfully
// analyzed but found to be invalid.
func (s tlsDatabase) VerifyChainFinish(result AsyncResult) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret C.GTlsCertificateFlags
	var _cerr *C.GError

	cret = C.g_tls_database_verify_chain_finish(_arg0, _arg1, _cerr)

	var _tlsCertificateFlags TLSCertificateFlags
	var _goerr error

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}
