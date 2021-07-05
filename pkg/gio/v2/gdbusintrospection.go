// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_dbus_annotation_info_get_type()), F: marshalDBusAnnotationInfo},
		{T: externglib.Type(C.g_dbus_arg_info_get_type()), F: marshalDBusArgInfo},
		{T: externglib.Type(C.g_dbus_interface_info_get_type()), F: marshalDBusInterfaceInfo},
		{T: externglib.Type(C.g_dbus_method_info_get_type()), F: marshalDBusMethodInfo},
		{T: externglib.Type(C.g_dbus_node_info_get_type()), F: marshalDBusNodeInfo},
		{T: externglib.Type(C.g_dbus_property_info_get_type()), F: marshalDBusPropertyInfo},
		{T: externglib.Type(C.g_dbus_signal_info_get_type()), F: marshalDBusSignalInfo},
	})
}

// DBusAnnotationInfo: information about an annotation.
type DBusAnnotationInfo struct {
	native C.GDBusAnnotationInfo
}

// WrapDBusAnnotationInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusAnnotationInfo(ptr unsafe.Pointer) *DBusAnnotationInfo {
	return (*DBusAnnotationInfo)(ptr)
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusAnnotationInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusAnnotationInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	var _arg0 *C.GDBusAnnotationInfo // out
	var _cret *C.GDBusAnnotationInfo // in

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_annotation_info_ref(_arg0)

	var _dBusAnnotationInfo *DBusAnnotationInfo // out

	_dBusAnnotationInfo = (*DBusAnnotationInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusAnnotationInfo, func(v *DBusAnnotationInfo) {
		C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
	})

	return _dBusAnnotationInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusAnnotationInfo) Unref() {
	var _arg0 *C.GDBusAnnotationInfo // out

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i))

	C.g_dbus_annotation_info_unref(_arg0)
}

// DBusArgInfo: information about an argument for a method or a signal.
type DBusArgInfo struct {
	native C.GDBusArgInfo
}

// WrapDBusArgInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusArgInfo(ptr unsafe.Pointer) *DBusArgInfo {
	return (*DBusArgInfo)(ptr)
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusArgInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusArgInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusArgInfo) Ref() *DBusArgInfo {
	var _arg0 *C.GDBusArgInfo // out
	var _cret *C.GDBusArgInfo // in

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_arg_info_ref(_arg0)

	var _dBusArgInfo *DBusArgInfo // out

	_dBusArgInfo = (*DBusArgInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusArgInfo, func(v *DBusArgInfo) {
		C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
	})

	return _dBusArgInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusArgInfo) Unref() {
	var _arg0 *C.GDBusArgInfo // out

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i))

	C.g_dbus_arg_info_unref(_arg0)
}

// DBusInterfaceInfo: information about a D-Bus interface.
type DBusInterfaceInfo struct {
	native C.GDBusInterfaceInfo
}

// WrapDBusInterfaceInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusInterfaceInfo(ptr unsafe.Pointer) *DBusInterfaceInfo {
	return (*DBusInterfaceInfo)(ptr)
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusInterfaceInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with @info, the existing cache is used and
// its use count is increased.
//
// Note that @info cannot be modified until
// g_dbus_interface_info_cache_release() is called.
func (i *DBusInterfaceInfo) CacheBuild() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_cache_build(_arg0)
}

// CacheRelease decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (i *DBusInterfaceInfo) CacheRelease() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_cache_release(_arg0)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusMethodInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_method(_arg0, _arg1)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	C.g_dbus_method_info_ref(_cret)

	return _dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusPropertyInfo  // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_property(_arg0, _arg1)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	C.g_dbus_property_info_ref(_cret)

	return _dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusSignalInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_signal(_arg0, _arg1)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	C.g_dbus_signal_info_ref(_cret)

	return _dBusSignalInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_interface_info_ref(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusInterfaceInfo) Unref() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_unref(_arg0)
}

// DBusMethodInfo: information about a method on an D-Bus interface.
type DBusMethodInfo struct {
	native C.GDBusMethodInfo
}

// WrapDBusMethodInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusMethodInfo(ptr unsafe.Pointer) *DBusMethodInfo {
	return (*DBusMethodInfo)(ptr)
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusMethodInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusMethodInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusMethodInfo) Ref() *DBusMethodInfo {
	var _arg0 *C.GDBusMethodInfo // out
	var _cret *C.GDBusMethodInfo // in

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_method_info_ref(_arg0)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusMethodInfo, func(v *DBusMethodInfo) {
		C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
	})

	return _dBusMethodInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusMethodInfo) Unref() {
	var _arg0 *C.GDBusMethodInfo // out

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i))

	C.g_dbus_method_info_unref(_arg0)
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
type DBusNodeInfo struct {
	native C.GDBusNodeInfo
}

// WrapDBusNodeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusNodeInfo(ptr unsafe.Pointer) *DBusNodeInfo {
	return (*DBusNodeInfo)(ptr)
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusNodeInfo)(unsafe.Pointer(b)), nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (*DBusNodeInfo, error) {
	var _arg1 *C.gchar         // out
	var _cret *C.GDBusNodeInfo // in
	var _cerr *C.GError        // in

	_arg1 = (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_new_for_xml(_arg1, &_cerr)

	var _dBusNodeInfo *DBusNodeInfo // out
	var _goerr error                // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusNodeInfo, _goerr
}

// Native returns the underlying C source pointer.
func (d *DBusNodeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
func (i *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var _arg0 *C.GDBusNodeInfo      // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_lookup_interface(_arg0, _arg1)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)

	return _dBusInterfaceInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusNodeInfo) Ref() *DBusNodeInfo {
	var _arg0 *C.GDBusNodeInfo // out
	var _cret *C.GDBusNodeInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_node_info_ref(_arg0)

	var _dBusNodeInfo *DBusNodeInfo // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})

	return _dBusNodeInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusNodeInfo) Unref() {
	var _arg0 *C.GDBusNodeInfo // out

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))

	C.g_dbus_node_info_unref(_arg0)
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
type DBusPropertyInfo struct {
	native C.GDBusPropertyInfo
}

// WrapDBusPropertyInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusPropertyInfo(ptr unsafe.Pointer) *DBusPropertyInfo {
	return (*DBusPropertyInfo)(ptr)
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusPropertyInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusPropertyInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	var _arg0 *C.GDBusPropertyInfo // out
	var _cret *C.GDBusPropertyInfo // in

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_property_info_ref(_arg0)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
	})

	return _dBusPropertyInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusPropertyInfo) Unref() {
	var _arg0 *C.GDBusPropertyInfo // out

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i))

	C.g_dbus_property_info_unref(_arg0)
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
type DBusSignalInfo struct {
	native C.GDBusSignalInfo
}

// WrapDBusSignalInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusSignalInfo(ptr unsafe.Pointer) *DBusSignalInfo {
	return (*DBusSignalInfo)(ptr)
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusSignalInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusSignalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusSignalInfo) Ref() *DBusSignalInfo {
	var _arg0 *C.GDBusSignalInfo // out
	var _cret *C.GDBusSignalInfo // in

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_signal_info_ref(_arg0)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_dBusSignalInfo, func(v *DBusSignalInfo) {
		C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
	})

	return _dBusSignalInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusSignalInfo) Unref() {
	var _arg0 *C.GDBusSignalInfo // out

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i))

	C.g_dbus_signal_info_unref(_arg0)
}
