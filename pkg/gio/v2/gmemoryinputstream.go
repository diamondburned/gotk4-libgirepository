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
		{T: externglib.Type(C.g_memory_input_stream_get_type()), F: marshalMemoryInputStream},
	})
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream interface {
	InputStream
	PollableInputStream
	Seekable
}

// memoryInputStream implements the MemoryInputStream class.
type memoryInputStream struct {
	InputStream
	PollableInputStream
	Seekable
}

var _ MemoryInputStream = (*memoryInputStream)(nil)

// WrapMemoryInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapMemoryInputStream(obj *externglib.Object) MemoryInputStream {
	return memoryInputStream{
		InputStream:         WrapInputStream(obj),
		PollableInputStream: WrapPollableInputStream(obj),
		Seekable:            WrapSeekable(obj),
	}
}

func marshalMemoryInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMemoryInputStream(obj), nil
}

// NewMemoryInputStream constructs a class MemoryInputStream.
func NewMemoryInputStream() MemoryInputStream {
	var _cret C.GMemoryInputStream // in

	_cret = C.g_memory_input_stream_new()

	var _memoryInputStream MemoryInputStream // out

	_memoryInputStream = WrapMemoryInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}
