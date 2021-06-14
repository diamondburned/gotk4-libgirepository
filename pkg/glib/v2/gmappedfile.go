// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mapped_file_get_type()), F: marshalMappedFile},
	})
}

// MappedFile: the File represents a file mapping created with
// g_mapped_file_new(). It has only private members and should not be accessed
// directly.
type MappedFile struct {
	native C.GMappedFile
}

// WrapMappedFile wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMappedFile(ptr unsafe.Pointer) *MappedFile {
	if ptr == nil {
		return nil
	}

	return (*MappedFile)(ptr)
}

func marshalMappedFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMappedFile(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MappedFile) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Free: this call existed before File had refcounting and is currently exactly
// the same as g_mapped_file_unref().
func (f *MappedFile) Free() {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	C.g_mapped_file_free(_arg0)
}

// Contents returns the contents of a File.
//
// Note that the contents may not be zero-terminated, even if the File is backed
// by a text file.
//
// If the file is empty then nil is returned.
func (f *MappedFile) Contents() string {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var _cret *C.gchar // in

	_cret = C.g_mapped_file_get_contents(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Length returns the length of the contents of a File.
func (f *MappedFile) Length() uint {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var _cret C.gsize // in

	_cret = C.g_mapped_file_get_length(_arg0)

	var _gsize uint // out

	_gsize = (uint)(_cret)

	return _gsize
}

// Unref decrements the reference count of @file by one. If the reference count
// drops to 0, unmaps the buffer of @file and frees it.
//
// It is safe to call this function from any thread.
//
// Since 2.22
func (f *MappedFile) Unref() {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	C.g_mapped_file_unref(_arg0)
}
