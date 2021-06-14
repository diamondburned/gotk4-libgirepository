// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_file_attribute_info_list_get_type()), F: marshalFileAttributeInfoList},
	})
}

// FileAttributeInfo: information about a specific attribute.
type FileAttributeInfo struct {
	native C.GFileAttributeInfo
}

// WrapFileAttributeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeInfo(ptr unsafe.Pointer) *FileAttributeInfo {
	if ptr == nil {
		return nil
	}

	return (*FileAttributeInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (f *FileAttributeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// Name gets the field inside the struct.
func (f *FileAttributeInfo) Name() string {
	var v string // out
	v = C.GoString(f.native.name)
	return v
}

// FileAttributeInfoList acts as a lightweight registry for possible valid file
// attributes. The registry stores Key-Value pair formats as AttributeInfos.
type FileAttributeInfoList struct {
	native C.GFileAttributeInfoList
}

// WrapFileAttributeInfoList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeInfoList(ptr unsafe.Pointer) *FileAttributeInfoList {
	if ptr == nil {
		return nil
	}

	return (*FileAttributeInfoList)(ptr)
}

func marshalFileAttributeInfoList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFileAttributeInfoList(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FileAttributeInfoList) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// NInfos gets the field inside the struct.
func (f *FileAttributeInfoList) NInfos() int {
	var v int // out
	v = (int)(f.native.n_infos)
	return v
}

// Add adds a new attribute with @name to the @list, setting its @type and
// @flags.
func (l *FileAttributeInfoList) Add(name string, typ FileAttributeType, flags FileAttributeInfoFlags) {
	var _arg0 *C.GFileAttributeInfoList // out
	var _arg1 *C.char                   // out
	var _arg2 C.GFileAttributeType      // out
	var _arg3 C.GFileAttributeInfoFlags // out

	_arg0 = (*C.GFileAttributeInfoList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GFileAttributeType)(typ)
	_arg3 = (C.GFileAttributeInfoFlags)(flags)

	C.g_file_attribute_info_list_add(_arg0, _arg1, _arg2, _arg3)
}

// Unref removes a reference from the given @list. If the reference count falls
// to zero, the @list is deleted.
func (l *FileAttributeInfoList) Unref() {
	var _arg0 *C.GFileAttributeInfoList // out

	_arg0 = (*C.GFileAttributeInfoList)(unsafe.Pointer(l.Native()))

	C.g_file_attribute_info_list_unref(_arg0)
}
