// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_resolver_name_lookup_flags_get_type()), F: marshalResolverNameLookupFlags},
		{T: externglib.Type(C.g_resolver_get_type()), F: marshalResolver},
	})
}

// ResolverNameLookupFlags flags to modify lookup behavior.
type ResolverNameLookupFlags int

const (
	// ResolverNameLookupFlagsDefault behavior (same as
	// g_resolver_lookup_by_name())
	ResolverNameLookupFlagsDefault ResolverNameLookupFlags = 0b0
	// ResolverNameLookupFlagsIPv4Only: only resolve ipv4 addresses
	ResolverNameLookupFlagsIPv4Only ResolverNameLookupFlags = 0b1
	// ResolverNameLookupFlagsIPv6Only: only resolve ipv6 addresses
	ResolverNameLookupFlagsIPv6Only ResolverNameLookupFlags = 0b10
)

func marshalResolverNameLookupFlags(p uintptr) (interface{}, error) {
	return ResolverNameLookupFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Resolver provides cancellable synchronous and asynchronous DNS resolution,
// for hostnames (g_resolver_lookup_by_address(), g_resolver_lookup_by_name()
// and their async variants) and SRV (service) records
// (g_resolver_lookup_service()).
//
// Address and Service provide wrappers around #GResolver functionality that
// also implement Connectable, making it easy to connect to a remote
// host/service.
type Resolver interface {
	gextras.Objector

	// LookupByAddress: synchronously reverse-resolves @address to determine its
	// associated hostname.
	//
	// If the DNS resolution fails, @error (if non-nil) will be set to a value
	// from Error.
	//
	// If @cancellable is non-nil, it can be used to cancel the operation, in
	// which case @error (if non-nil) will be set to G_IO_ERROR_CANCELLED.
	LookupByAddress(address InetAddress, cancellable Cancellable) (string, error)
	// LookupByAddressAsync begins asynchronously reverse-resolving @address to
	// determine its associated hostname, and eventually calls @callback, which
	// must call g_resolver_lookup_by_address_finish() to get the final result.
	LookupByAddressAsync(address InetAddress, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupByAddressFinish retrieves the result of a previous call to
	// g_resolver_lookup_by_address_async().
	//
	// If the DNS resolution failed, @error (if non-nil) will be set to a value
	// from Error. If the operation was cancelled, @error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupByAddressFinish(result AsyncResult) (string, error)
	// LookupByNameAsync begins asynchronously resolving @hostname to determine
	// its associated IP address(es), and eventually calls @callback, which must
	// call g_resolver_lookup_by_name_finish() to get the result. See
	// g_resolver_lookup_by_name() for more details.
	LookupByNameAsync(hostname string, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupByNameWithFlagsAsync begins asynchronously resolving @hostname to
	// determine its associated IP address(es), and eventually calls @callback,
	// which must call g_resolver_lookup_by_name_with_flags_finish() to get the
	// result. See g_resolver_lookup_by_name() for more details.
	LookupByNameWithFlagsAsync(hostname string, flags ResolverNameLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupRecordsAsync begins asynchronously performing a DNS lookup for the
	// given @rrname, and eventually calls @callback, which must call
	// g_resolver_lookup_records_finish() to get the final result. See
	// g_resolver_lookup_records() for more details.
	LookupRecordsAsync(rrname string, recordType ResolverRecordType, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupServiceAsync begins asynchronously performing a DNS SRV lookup for
	// the given @service and @protocol in the given @domain, and eventually
	// calls @callback, which must call g_resolver_lookup_service_finish() to
	// get the final result. See g_resolver_lookup_service() for more details.
	LookupServiceAsync(service string, protocol string, domain string, cancellable Cancellable, callback AsyncReadyCallback)
	// SetDefault sets @resolver to be the application's default resolver
	// (reffing @resolver, and unreffing the previous default resolver, if any).
	// Future calls to g_resolver_get_default() will return this resolver.
	//
	// This can be used if an application wants to perform any sort of DNS
	// caching or "pinning"; it can implement its own #GResolver that calls the
	// original default resolver for DNS operations, and implements its own
	// cache policies on top of that, and then set itself as the default
	// resolver for all later code to use.
	SetDefault()
}

// resolver implements the Resolver interface.
type resolver struct {
	*externglib.Object
}

var _ Resolver = (*resolver)(nil)

// WrapResolver wraps a GObject to a type that implements
// interface Resolver. It is primarily used internally.
func WrapResolver(obj *externglib.Object) Resolver {
	return resolver{obj}
}

func marshalResolver(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapResolver(obj), nil
}

func (r resolver) LookupByAddress(address InetAddress, cancellable Cancellable) (string, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GInetAddress // out
	var _arg2 *C.GCancellable // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_resolver_lookup_by_address(_arg0, _arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

func (r resolver) LookupByAddressAsync(address InetAddress, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg1 *C.GInetAddress       // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_resolver_lookup_by_address_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (r resolver) LookupByAddressFinish(result AsyncResult) (string, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_by_address_finish(_arg0, _arg1, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

func (r resolver) LookupByNameAsync(hostname string, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_resolver_lookup_by_name_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (r resolver) LookupByNameWithFlagsAsync(hostname string, flags ResolverNameLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver               // out
	var _arg1 *C.gchar                   // out
	var _arg2 C.GResolverNameLookupFlags // out
	var _arg3 *C.GCancellable            // out
	var _arg4 C.GAsyncReadyCallback      // out
	var _arg5 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverNameLookupFlags(flags)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_resolver_lookup_by_name_with_flags_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (r resolver) LookupRecordsAsync(rrname string, recordType ResolverRecordType, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GResolverRecordType // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(rrname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverRecordType(recordType)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_resolver_lookup_records_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (r resolver) LookupServiceAsync(service string, protocol string, domain string, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg3 *C.gchar              // out
	var _arg4 *C.GCancellable       // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.g_resolver_lookup_service_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (r resolver) SetDefault() {
	var _arg0 *C.GResolver // out

	_arg0 = (*C.GResolver)(unsafe.Pointer(r.Native()))

	C.g_resolver_set_default(_arg0)
}
