// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filter_input_stream_get_type()), F: marshalFilterInputStream},
	})
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream interface {
	gextras.Objector

	// AsInputStream casts the class to the InputStream interface.
	AsInputStream() InputStream

	// ClearPending clears the pending flag on @stream.
	//
	// This method is inherited from InputStream
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Streams will be automatically closed when the last reference is dropped,
	// but you might want to call this function to make sure resources are
	// released as early as possible.
	//
	// Some streams might keep the backing store of the stream (e.g. a file
	// descriptor) open after the stream is closed. See the documentation for
	// the individual stream for details.
	//
	// On failure the first error that happened will be reported, but the close
	// operation will finish as much as possible. A stream that failed to close
	// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
	// important to check and report the error to the user.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but some streams
	// can use a faster close that doesn't block to e.g. check errors.
	//
	// This method is inherited from InputStream
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous closes of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_input_stream_close_finish() to get the result
	// of the operation.
	//
	// For behaviour details see g_input_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// This method is inherited from InputStream
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish finishes closing a stream asynchronously, started from
	// g_input_stream_close_async().
	//
	// This method is inherited from InputStream
	CloseFinish(result AsyncResult) error
	// HasPending checks if an input stream has pending actions.
	//
	// This method is inherited from InputStream
	HasPending() bool
	// IsClosed checks if an input stream is closed.
	//
	// This method is inherited from InputStream
	IsClosed() bool
	// ReadAllFinish finishes an asynchronous stream read operation started with
	// g_input_stream_read_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_read will be set to the number of bytes that were successfully
	// read before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_input_stream_read_async().
	//
	// This method is inherited from InputStream
	ReadAllFinish(result AsyncResult) (uint, error)
	// ReadBytesAsync: request an asynchronous read of @count bytes from the
	// stream into a new #GBytes. When the operation is finished @callback will
	// be called. You can then call g_input_stream_read_bytes_finish() to get
	// the result of the operation.
	//
	// During an async request no other sync and async calls are allowed on
	// @stream, and will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the new #GBytes will be passed to the callback. It is not an
	// error if this is smaller than the requested size, as it can happen e.g.
	// near the end of a file, but generally we try to read as many bytes as
	// requested. Zero is returned on end of file (or if @count is zero), but
	// never otherwise.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// This method is inherited from InputStream
	ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadFinish finishes an asynchronous stream read operation.
	//
	// This method is inherited from InputStream
	ReadFinish(result AsyncResult) (int, error)
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	//
	// This method is inherited from InputStream
	SetPending() error
	// Skip tries to skip @count bytes from the stream. Will block during the
	// operation.
	//
	// This is identical to g_input_stream_read(), from a behaviour standpoint,
	// but the bytes that are skipped are not returned to the user. Some streams
	// have an implementation that is more efficient than reading the data.
	//
	// This function is optional for inherited classes, as the default
	// implementation emulates it using read.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// This method is inherited from InputStream
	Skip(count uint, cancellable Cancellable) (int, error)
	// SkipAsync: request an asynchronous skip of @count bytes from the stream.
	// When the operation is finished @callback will be called. You can then
	// call g_input_stream_skip_finish() to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes skipped will be passed to the callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. near the end of a file, but generally we try to skip as
	// many bytes as requested. Zero is returned on end of file (or if @count is
	// zero), but never otherwise.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one, you must override all.
	//
	// This method is inherited from InputStream
	SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// SkipFinish finishes a stream skip operation.
	//
	// This method is inherited from InputStream
	SkipFinish(result AsyncResult) (int, error)

	// BaseStream gets the base stream for the filter stream.
	BaseStream() InputStream
	// CloseBaseStream returns whether the base stream will be closed when
	// @stream is closed.
	CloseBaseStream() bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// @stream is closed.
	SetCloseBaseStream(closeBase bool)
}

// filterInputStream implements the FilterInputStream interface.
type filterInputStream struct {
	*externglib.Object
}

var _ FilterInputStream = (*filterInputStream)(nil)

// WrapFilterInputStream wraps a GObject to a type that implements
// interface FilterInputStream. It is primarily used internally.
func WrapFilterInputStream(obj *externglib.Object) FilterInputStream {
	return filterInputStream{obj}
}

func marshalFilterInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilterInputStream(obj), nil
}

func (f filterInputStream) AsInputStream() InputStream {
	return WrapInputStream(gextras.InternObject(f))
}

func (s filterInputStream) ClearPending() {
	WrapInputStream(gextras.InternObject(s)).ClearPending()
}

func (s filterInputStream) Close(cancellable Cancellable) error {
	return WrapInputStream(gextras.InternObject(s)).Close(cancellable)
}

func (s filterInputStream) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).CloseAsync(ioPriority, cancellable, callback)
}

func (s filterInputStream) CloseFinish(result AsyncResult) error {
	return WrapInputStream(gextras.InternObject(s)).CloseFinish(result)
}

func (s filterInputStream) HasPending() bool {
	return WrapInputStream(gextras.InternObject(s)).HasPending()
}

func (s filterInputStream) IsClosed() bool {
	return WrapInputStream(gextras.InternObject(s)).IsClosed()
}

func (s filterInputStream) ReadAllFinish(result AsyncResult) (uint, error) {
	return WrapInputStream(gextras.InternObject(s)).ReadAllFinish(result)
}

func (s filterInputStream) ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).ReadBytesAsync(count, ioPriority, cancellable, callback)
}

func (s filterInputStream) ReadFinish(result AsyncResult) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).ReadFinish(result)
}

func (s filterInputStream) SetPending() error {
	return WrapInputStream(gextras.InternObject(s)).SetPending()
}

func (s filterInputStream) Skip(count uint, cancellable Cancellable) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).Skip(count, cancellable)
}

func (s filterInputStream) SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).SkipAsync(count, ioPriority, cancellable, callback)
}

func (s filterInputStream) SkipFinish(result AsyncResult) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).SkipFinish(result)
}

func (s filterInputStream) BaseStream() InputStream {
	var _arg0 *C.GFilterInputStream // out
	var _cret *C.GInputStream       // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_filter_input_stream_get_base_stream(_arg0)

	var _inputStream InputStream // out

	_inputStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(InputStream)

	return _inputStream
}

func (s filterInputStream) CloseBaseStream() bool {
	var _arg0 *C.GFilterInputStream // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_filter_input_stream_get_close_base_stream(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s filterInputStream) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterInputStream // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_input_stream_set_close_base_stream(_arg0, _arg1)
}
