// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
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
import "C"

// ResourcesEnumerateChildren returns all the names of children at the specified
// path in the set of globally registered resources. The return result is a NULL
// terminated list of strings which should be released with g_strfreev().
//
// lookup_flags controls the behaviour of the lookup.
func ResourcesEnumerateChildren(path string, lookupFlags ResourceLookupFlags) ([]string, error) {
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret **C.char               // in
	var _cerr *C.GError              // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resources_enumerate_children(_arg1, _arg2, &_cerr)

	var _utf8s []string // out
	var _goerr error    // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

// ResourcesGetInfo looks for a file at the specified path in the set of
// globally registered resources and if found returns information about it.
//
// lookup_flags controls the behaviour of the lookup.
func ResourcesGetInfo(path string, lookupFlags ResourceLookupFlags) (uint, uint32, error) {
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _arg3 C.gsize                // in
	var _arg4 C.guint32              // in
	var _cerr *C.GError              // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	C.g_resources_get_info(_arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _size uint    // out
	var _flags uint32 // out
	var _goerr error  // out

	_size = uint(_arg3)
	_flags = uint32(_arg4)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _size, _flags, _goerr
}

// ResourcesOpenStream looks for a file at the specified path in the set of
// globally registered resources and returns a Stream that lets you read the
// data.
//
// lookup_flags controls the behaviour of the lookup.
func ResourcesOpenStream(path string, lookupFlags ResourceLookupFlags) (InputStreamer, error) {
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret *C.GInputStream        // in
	var _cerr *C.GError              // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resources_open_stream(_arg1, _arg2, &_cerr)

	var _inputStream InputStreamer // out
	var _goerr error               // out

	_inputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(InputStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inputStream, _goerr
}

// ResourcesRegister registers the resource with the process-global set of
// resources. Once a resource is registered the files in it can be accessed with
// the global resource lookup functions like g_resources_lookup_data().
func ResourcesRegister(resource *Resource) {
	var _arg1 *C.GResource // out

	_arg1 = (*C.GResource)(gextras.StructNative(unsafe.Pointer(resource)))

	C.g_resources_register(_arg1)
}

// ResourcesUnregister unregisters the resource from the process-global set of
// resources.
func ResourcesUnregister(resource *Resource) {
	var _arg1 *C.GResource // out

	_arg1 = (*C.GResource)(gextras.StructNative(unsafe.Pointer(resource)))

	C.g_resources_unregister(_arg1)
}

// ResourceLoad loads a binary resource bundle and creates a #GResource
// representation of it, allowing you to query it for data.
//
// If you want to use this resource in the global resource namespace you need to
// register it with g_resources_register().
//
// If filename is empty or the data in it is corrupt, G_RESOURCE_ERROR_INTERNAL
// will be returned. If filename doesn’t exist, or there is an error in reading
// it, an error from g_mapped_file_new() will be returned.
func ResourceLoad(filename string) (*Resource, error) {
	var _arg1 *C.gchar     // out
	var _cret *C.GResource // in
	var _cerr *C.GError    // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))

	_cret = C.g_resource_load(_arg1, &_cerr)

	var _resource *Resource // out
	var _goerr error        // out

	_resource = (*Resource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_resource_ref(_cret)
	runtime.SetFinalizer(_resource, func(v *Resource) {
		C.g_resource_unref((*C.GResource)(gextras.StructNative(unsafe.Pointer(v))))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _resource, _goerr
}

// StaticResource is an opaque data structure and can only be accessed using the
// following functions.
type StaticResource struct {
	nocopy gextras.NoCopy
	native *C.GStaticResource
}

// Fini: finalized a GResource initialized by g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (staticResource *StaticResource) Fini() {
	var _arg0 *C.GStaticResource // out

	_arg0 = (*C.GStaticResource)(gextras.StructNative(unsafe.Pointer(staticResource)))

	C.g_static_resource_fini(_arg0)
}

// Resource gets the GResource that was registered by a call to
// g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (staticResource *StaticResource) Resource() *Resource {
	var _arg0 *C.GStaticResource // out
	var _cret *C.GResource       // in

	_arg0 = (*C.GStaticResource)(gextras.StructNative(unsafe.Pointer(staticResource)))

	_cret = C.g_static_resource_get_resource(_arg0)

	var _resource *Resource // out

	_resource = (*Resource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_resource, func(v *Resource) {
		C.g_resource_unref((*C.GResource)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _resource
}

// Init initializes a GResource from static data using a GStaticResource.
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (staticResource *StaticResource) Init() {
	var _arg0 *C.GStaticResource // out

	_arg0 = (*C.GStaticResource)(gextras.StructNative(unsafe.Pointer(staticResource)))

	C.g_static_resource_init(_arg0)
}
