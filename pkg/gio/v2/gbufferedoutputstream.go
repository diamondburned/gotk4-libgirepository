// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gbufferedoutputstream.go.
var GTypeBufferedOutputStream = coreglib.Type(C.g_buffered_output_stream_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBufferedOutputStream, F: marshalBufferedOutputStream},
	})
}

// BufferedOutputStreamOverrider contains methods that are overridable.
type BufferedOutputStreamOverrider interface {
}

// BufferedOutputStream: buffered output stream implements OutputStream and
// provides for buffered writes.
//
// By default, OutputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered output stream, use g_buffered_output_stream_new(), or
// g_buffered_output_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_output_stream_get_buffer_size(). To change the size of a buffered
// output stream's buffer, use g_buffered_output_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedOutputStream struct {
	_ [0]func() // equal guard
	FilterOutputStream

	Seekable
}

var (
	_ FilterOutputStreamer = (*BufferedOutputStream)(nil)
)

func classInitBufferedOutputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBufferedOutputStream(obj *coreglib.Object) *BufferedOutputStream {
	return &BufferedOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalBufferedOutputStream(p uintptr) (interface{}, error) {
	return wrapBufferedOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBufferedOutputStream creates a new buffered output stream for a base
// stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//
// The function returns the following values:
//
//    - bufferedOutputStream for the given base_stream.
//
func NewBufferedOutputStream(baseStream OutputStreamer) *BufferedOutputStream {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("new_BufferedOutputStream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)

	var _bufferedOutputStream *BufferedOutputStream // out

	_bufferedOutputStream = wrapBufferedOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedOutputStream
}

// NewBufferedOutputStreamSized creates a new buffered output stream with a
// given buffer size.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - size: #gsize.
//
// The function returns the following values:
//
//    - bufferedOutputStream with an internal buffer set to size.
//
func NewBufferedOutputStreamSized(baseStream OutputStreamer, size uint) *BufferedOutputStream {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[1])) = C.gsize(size)

	_gret := girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("new_BufferedOutputStream_sized", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(size)

	var _bufferedOutputStream *BufferedOutputStream // out

	_bufferedOutputStream = wrapBufferedOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedOutputStream
}

// AutoGrow checks if the buffer automatically grows as data is added.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream's buffer automatically grows, FALSE otherwise.
//
func (stream *BufferedOutputStream) AutoGrow() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("get_auto_grow", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// BufferSize gets the size of the buffer in the stream.
//
// The function returns the following values:
//
//    - gsize: current size of the buffer.
//
func (stream *BufferedOutputStream) BufferSize() uint {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("get_buffer_size", _args[:], nil)
	_cret = *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(*(*C.gsize)(unsafe.Pointer(&_cret)))

	return _gsize
}

// SetAutoGrow sets whether or not the stream's buffer should automatically
// grow. If auto_grow is true, then each write will just make the buffer larger,
// and you must manually flush the buffer to actually write out the data to the
// underlying stream.
//
// The function takes the following parameters:
//
//    - autoGrow: #gboolean.
//
func (stream *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	if autoGrow {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("set_auto_grow", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(autoGrow)
}

// SetBufferSize sets the size of the internal buffer to size.
//
// The function takes the following parameters:
//
//    - size: #gsize.
//
func (stream *BufferedOutputStream) SetBufferSize(size uint) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[1])) = C.gsize(size)

	girepository.MustFind("Gio", "BufferedOutputStream").InvokeMethod("set_buffer_size", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(size)
}
