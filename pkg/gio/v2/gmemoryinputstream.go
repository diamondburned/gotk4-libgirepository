// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeMemoryInputStream returns the GType for the type MemoryInputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMemoryInputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "MemoryInputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMemoryInputStream)
	return gtype
}

// MemoryInputStreamOverrider contains methods that are overridable.
type MemoryInputStreamOverrider interface {
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream struct {
	_ [0]func() // equal guard
	InputStream

	*coreglib.Object
	PollableInputStream
	Seekable
}

var (
	_ InputStreamer     = (*MemoryInputStream)(nil)
	_ coreglib.Objector = (*MemoryInputStream)(nil)
)

func classInitMemoryInputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMemoryInputStream(obj *coreglib.Object) *MemoryInputStream {
	return &MemoryInputStream{
		InputStream: InputStream{
			Object: obj,
		},
		Object: obj,
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalMemoryInputStream(p uintptr) (interface{}, error) {
	return wrapMemoryInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMemoryInputStream creates a new empty InputStream.
//
// The function returns the following values:
//
//    - memoryInputStream: new Stream.
//
func NewMemoryInputStream() *MemoryInputStream {
	_gret := girepository.MustFind("Gio", "MemoryInputStream").InvokeMethod("new_MemoryInputStream", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// NewMemoryInputStreamFromBytes creates a new InputStream with data from the
// given bytes.
//
// The function takes the following parameters:
//
//    - bytes: #GBytes.
//
// The function returns the following values:
//
//    - memoryInputStream: new Stream read from bytes.
//
func NewMemoryInputStreamFromBytes(bytes *glib.Bytes) *MemoryInputStream {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(bytes)))

	_gret := girepository.MustFind("Gio", "MemoryInputStream").InvokeMethod("new_MemoryInputStream_from_bytes", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bytes)

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// AddBytes appends bytes to data that can be read from the input stream.
//
// The function takes the following parameters:
//
//    - bytes: input data.
//
func (stream *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(bytes)))

	girepository.MustFind("Gio", "MemoryInputStream").InvokeMethod("add_bytes", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(bytes)
}
