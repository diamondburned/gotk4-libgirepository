// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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

	// AddBytes appends @bytes to data that can be read from the input stream.
	AddBytes(bytes *glib.Bytes)
}

// memoryInputStream implements the MemoryInputStream interface.
type memoryInputStream struct {
	InputStream
	PollableInputStream
	Seekable
}

var _ MemoryInputStream = (*memoryInputStream)(nil)

// WrapMemoryInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapMemoryInputStream(obj *externglib.Object) MemoryInputStream {
	return MemoryInputStream{
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
	cret := new(C.GMemoryInputStream)
	var goret MemoryInputStream

	cret = C.g_memory_input_stream_new()

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MemoryInputStream)

	return goret
}

// NewMemoryInputStreamFromBytes constructs a class MemoryInputStream.
func NewMemoryInputStreamFromBytes(bytes *glib.Bytes) MemoryInputStream {
	var arg1 *C.GBytes

	arg1 = (*C.GBytes)(unsafe.Pointer(bytes.Native()))

	cret := new(C.GMemoryInputStream)
	var goret MemoryInputStream

	cret = C.g_memory_input_stream_new_from_bytes(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MemoryInputStream)

	return goret
}

// AddBytes appends @bytes to data that can be read from the input stream.
func (s memoryInputStream) AddBytes(bytes *glib.Bytes) {
	var arg0 *C.GMemoryInputStream
	var arg1 *C.GBytes

	arg0 = (*C.GMemoryInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(bytes.Native()))

	C.g_memory_input_stream_add_bytes(arg0, arg1)
}
