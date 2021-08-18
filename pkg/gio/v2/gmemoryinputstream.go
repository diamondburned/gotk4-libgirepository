// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_memory_input_stream_get_type()), F: marshalMemoryInputStreamer},
	})
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream struct {
	InputStream

	PollableInputStream
	Seekable
	*externglib.Object
}

func wrapMemoryInputStream(obj *externglib.Object) *MemoryInputStream {
	return &MemoryInputStream{
		InputStream: InputStream{
			Object: obj,
		},
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalMemoryInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryInputStream(obj), nil
}

// NewMemoryInputStream creates a new empty InputStream.
func NewMemoryInputStream() *MemoryInputStream {
	var _cret *C.GInputStream // in

	_cret = C.g_memory_input_stream_new()

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// NewMemoryInputStreamFromBytes creates a new InputStream with data from the
// given bytes.
func NewMemoryInputStreamFromBytes(bytes []byte) *MemoryInputStream {
	var _arg1 *C.GBytes       // out
	var _cret *C.GInputStream // in

	_arg1 = C.g_bytes_new_with_free_func(
		C.gconstpointer(unsafe.Pointer(&bytes[0])),
		C.gsize(len(bytes)),
		C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
		C.gpointer(gbox.Assign(bytes)),
	)
	defer C.g_bytes_unref(_arg1)

	_cret = C.g_memory_input_stream_new_from_bytes(_arg1)
	runtime.KeepAlive(bytes)

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// AddBytes appends bytes to data that can be read from the input stream.
func (stream *MemoryInputStream) AddBytes(bytes []byte) {
	var _arg0 *C.GMemoryInputStream // out
	var _arg1 *C.GBytes             // out

	_arg0 = (*C.GMemoryInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.g_bytes_new_with_free_func(
		C.gconstpointer(unsafe.Pointer(&bytes[0])),
		C.gsize(len(bytes)),
		C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
		C.gpointer(gbox.Assign(bytes)),
	)
	defer C.g_bytes_unref(_arg1)

	C.g_memory_input_stream_add_bytes(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(bytes)
}
