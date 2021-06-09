// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.g_buffered_input_stream_get_type()), F: marshalBufferedInputStream},
	})
}

// BufferedInputStream: buffered input stream implements InputStream and
// provides for buffered reads.
//
// By default, InputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered input stream, use g_buffered_input_stream_new(), or
// g_buffered_input_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_input_stream_get_buffer_size(). To change the size of a buffered
// input stream's buffer, use g_buffered_input_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedInputStream interface {
	FilterInputStream
	Seekable

	// Fill tries to read @count bytes from the stream into the buffer. Will
	// block during this read.
	//
	// If @count is zero, returns zero and does nothing. A value of @count
	// larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. near the end of a file. Zero is returned on end of file (or
	// if @count is zero), but never otherwise.
	//
	// If @count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and @error is set accordingly.
	//
	// For the asynchronous, non-blocking, version of this function, see
	// g_buffered_input_stream_fill_async().
	Fill(count int, cancellable Cancellable) (int, error)
	// FillAsync reads data into @stream's buffer asynchronously, up to @count
	// size. @io_priority can be used to prioritize reads. For the synchronous
	// version of this function, see g_buffered_input_stream_fill().
	//
	// If @count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	FillAsync()
	// FillFinish finishes an asynchronous read.
	FillFinish(result AsyncResult) (int, error)
	// Available gets the size of the available data within the stream.
	Available() uint
	// BufferSize gets the size of the input buffer.
	BufferSize() uint
	// Peek peeks in the buffer, copying data of size @count into @buffer,
	// offset @offset bytes.
	Peek() uint
	// PeekBuffer returns the buffer with the currently available bytes. The
	// returned buffer must not be modified and will become invalid when reading
	// from the stream or filling the buffer.
	PeekBuffer() []byte
	// ReadByte tries to read a single byte from the stream or the buffer. Will
	// block during this read.
	//
	// On success, the byte read from the stream is returned. On end of stream
	// -1 is returned but it's not an exceptional error and @error is not set.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and @error is set accordingly.
	ReadByte(cancellable Cancellable) (int, error)
	// SetBufferSize sets the size of the internal buffer of @stream to @size,
	// or to the size of the contents of the buffer. The buffer can never be
	// resized smaller than its current contents.
	SetBufferSize(size uint)
}

// bufferedInputStream implements the BufferedInputStream interface.
type bufferedInputStream struct {
	FilterInputStream
	Seekable
}

var _ BufferedInputStream = (*bufferedInputStream)(nil)

// WrapBufferedInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapBufferedInputStream(obj *externglib.Object) BufferedInputStream {
	return BufferedInputStream{
		FilterInputStream: WrapFilterInputStream(obj),
		Seekable:          WrapSeekable(obj),
	}
}

func marshalBufferedInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBufferedInputStream(obj), nil
}

// NewBufferedInputStream constructs a class BufferedInputStream.
func NewBufferedInputStream(baseStream InputStream) BufferedInputStream {
	var _arg1 *C.GInputStream

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))

	var _cret C.GBufferedInputStream

	cret = C.g_buffered_input_stream_new(_arg1)

	var _bufferedInputStream BufferedInputStream

	_bufferedInputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(BufferedInputStream)

	return _bufferedInputStream
}

// NewBufferedInputStreamSized constructs a class BufferedInputStream.
func NewBufferedInputStreamSized(baseStream InputStream, size uint) BufferedInputStream {
	var _arg1 *C.GInputStream
	var _arg2 C.gsize

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = C.gsize(size)

	var _cret C.GBufferedInputStream

	cret = C.g_buffered_input_stream_new_sized(_arg1, _arg2)

	var _bufferedInputStream BufferedInputStream

	_bufferedInputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(BufferedInputStream)

	return _bufferedInputStream
}

// Fill tries to read @count bytes from the stream into the buffer. Will
// block during this read.
//
// If @count is zero, returns zero and does nothing. A value of @count
// larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is
// not an error if this is not the same as the requested size, as it can
// happen e.g. near the end of a file. Zero is returned on end of file (or
// if @count is zero), but never otherwise.
//
// If @count is -1 then the attempted read size is equal to the number of
// bytes that are required to fill the buffer.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
func (s bufferedInputStream) Fill(count int, cancellable Cancellable) (int, error) {
	var _arg0 *C.GBufferedInputStream
	var _arg1 C.gssize
	var _arg2 *C.GCancellable

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gssize(count)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret C.gssize
	var _cerr *C.GError

	cret = C.g_buffered_input_stream_fill(_arg0, _arg1, _arg2, _cerr)

	var _gssize int
	var _goerr error

	_gssize = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// FillAsync reads data into @stream's buffer asynchronously, up to @count
// size. @io_priority can be used to prioritize reads. For the synchronous
// version of this function, see g_buffered_input_stream_fill().
//
// If @count is -1 then the attempted read size is equal to the number of
// bytes that are required to fill the buffer.
func (s bufferedInputStream) FillAsync() {
	var _arg0 *C.GBufferedInputStream

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))

	C.g_buffered_input_stream_fill_async(_arg0)
}

// FillFinish finishes an asynchronous read.
func (s bufferedInputStream) FillFinish(result AsyncResult) (int, error) {
	var _arg0 *C.GBufferedInputStream
	var _arg1 *C.GAsyncResult

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cret C.gssize
	var _cerr *C.GError

	cret = C.g_buffered_input_stream_fill_finish(_arg0, _arg1, _cerr)

	var _gssize int
	var _goerr error

	_gssize = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// Available gets the size of the available data within the stream.
func (s bufferedInputStream) Available() uint {
	var _arg0 *C.GBufferedInputStream

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gsize

	cret = C.g_buffered_input_stream_get_available(_arg0)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// BufferSize gets the size of the input buffer.
func (s bufferedInputStream) BufferSize() uint {
	var _arg0 *C.GBufferedInputStream

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gsize

	cret = C.g_buffered_input_stream_get_buffer_size(_arg0)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// Peek peeks in the buffer, copying data of size @count into @buffer,
// offset @offset bytes.
func (s bufferedInputStream) Peek() uint {
	var _arg0 *C.GBufferedInputStream

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gsize

	cret = C.g_buffered_input_stream_peek(_arg0)

	var _gsize uint

	_gsize = (uint)(_cret)

	return _gsize
}

// PeekBuffer returns the buffer with the currently available bytes. The
// returned buffer must not be modified and will become invalid when reading
// from the stream or filling the buffer.
func (s bufferedInputStream) PeekBuffer() []byte {
	var _arg0 *C.GBufferedInputStream

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))

	var _cret *C.void
	var _arg1 *C.gsize

	cret = C.g_buffered_input_stream_peek_buffer(_arg0)

	var _guint8s []byte

	{
		var src []C.guint8
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg1))

		_guint8s = make([]byte, _arg1)
		for i := 0; i < uintptr(_arg1); i++ {
			_guint8s = (byte)(_cret)
		}
	}

	return _guint8s
}

// ReadByte tries to read a single byte from the stream or the buffer. Will
// block during this read.
//
// On success, the byte read from the stream is returned. On end of stream
// -1 is returned but it's not an exceptional error and @error is not set.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
func (s bufferedInputStream) ReadByte(cancellable Cancellable) (int, error) {
	var _arg0 *C.GBufferedInputStream
	var _arg1 *C.GCancellable

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret C.int
	var _cerr *C.GError

	cret = C.g_buffered_input_stream_read_byte(_arg0, _arg1, _cerr)

	var _gint int
	var _goerr error

	_gint = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// SetBufferSize sets the size of the internal buffer of @stream to @size,
// or to the size of the contents of the buffer. The buffer can never be
// resized smaller than its current contents.
func (s bufferedInputStream) SetBufferSize(size uint) {
	var _arg0 *C.GBufferedInputStream
	var _arg1 C.gsize

	_arg0 = (*C.GBufferedInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gsize(size)

	C.g_buffered_input_stream_set_buffer_size(_arg0, _arg1)
}