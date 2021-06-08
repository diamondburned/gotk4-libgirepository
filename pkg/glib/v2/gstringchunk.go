// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// StringChunk: an opaque data structure representing String Chunks. It should
// only be accessed by using the following functions.
type StringChunk struct {
	native C.GStringChunk
}

// WrapStringChunk wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStringChunk(ptr unsafe.Pointer) *StringChunk {
	if ptr == nil {
		return nil
	}

	return (*StringChunk)(ptr)
}

func marshalStringChunk(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStringChunk(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StringChunk) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Clear frees all strings contained within the Chunk. After calling
// g_string_chunk_clear() it is not safe to access any of the strings which were
// contained within it.
func (c *StringChunk) Clear() {
	var arg0 *C.GStringChunk

	arg0 = (*C.GStringChunk)(unsafe.Pointer(c.Native()))

	C.g_string_chunk_clear(arg0)
}

// Free frees all memory allocated by the Chunk. After calling
// g_string_chunk_free() it is not safe to access any of the strings which were
// contained within it.
func (c *StringChunk) Free() {
	var arg0 *C.GStringChunk

	arg0 = (*C.GStringChunk)(unsafe.Pointer(c.Native()))

	C.g_string_chunk_free(arg0)
}

// Insert adds a copy of @string to the Chunk. It returns a pointer to the new
// copy of the string in the Chunk. The characters in the string can be changed,
// if necessary, though you should not change anything after the end of the
// string.
//
// Unlike g_string_chunk_insert_const(), this function does not check for
// duplicates. Also strings added with g_string_chunk_insert() will not be
// searched by g_string_chunk_insert_const() when looking for duplicates.
func (c *StringChunk) Insert(string string) string {
	var arg0 *C.GStringChunk
	var arg1 *C.gchar

	arg0 = (*C.GStringChunk)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.gchar)
	var goret string

	cret = C.g_string_chunk_insert(arg0, arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// InsertConst adds a copy of @string to the Chunk, unless the same string has
// already been added to the Chunk with g_string_chunk_insert_const().
//
// This function is useful if you need to copy a large number of strings but do
// not want to waste space storing duplicates. But you must remember that there
// may be several pointers to the same string, and so any changes made to the
// strings should be done very carefully.
//
// Note that g_string_chunk_insert_const() will not return a pointer to a string
// added with g_string_chunk_insert(), even if they do match.
func (c *StringChunk) InsertConst(string string) string {
	var arg0 *C.GStringChunk
	var arg1 *C.gchar

	arg0 = (*C.GStringChunk)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.gchar)
	var goret string

	cret = C.g_string_chunk_insert_const(arg0, arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// InsertLen adds a copy of the first @len bytes of @string to the Chunk. The
// copy is nul-terminated.
//
// Since this function does not stop at nul bytes, it is the caller's
// responsibility to ensure that @string has at least @len addressable bytes.
//
// The characters in the returned string can be changed, if necessary, though
// you should not change anything after the end of the string.
func (c *StringChunk) InsertLen(string string, len int) string {
	var arg0 *C.GStringChunk
	var arg1 *C.gchar
	var arg2 C.gssize

	arg0 = (*C.GStringChunk)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	cret := new(C.gchar)
	var goret string

	cret = C.g_string_chunk_insert_len(arg0, arg1, arg2)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}
