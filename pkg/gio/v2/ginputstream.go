// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"

	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.g_input_stream_get_type()), F: marshalInputStream},
	})
}

// InputStream has functions to read from a stream (g_input_stream_read()), to
// close a stream (g_input_stream_close()) and to skip some content
// (g_input_stream_skip()).
//
// To copy the content of an input stream to an output stream without manually
// handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for OStream for details of thread safety of streaming
// APIs.
//
// All of these functions have async variants too.
type InputStream interface {
	gextras.Objector

	// ClearPending clears the pending flag on @stream.
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
	CloseAsync()
	// CloseFinish finishes closing a stream asynchronously, started from
	// g_input_stream_close_async().
	CloseFinish(result AsyncResult) error
	// HasPending checks if an input stream has pending actions.
	HasPending() bool
	// IsClosed checks if an input stream is closed.
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
	ReadAllFinish(result AsyncResult) (bytesRead uint, err error)
	// ReadBytes: like g_input_stream_read(), this tries to read @count bytes
	// from the stream in a blocking fashion. However, rather than reading into
	// a user-supplied buffer, this will create a new #GBytes containing the
	// data that was read. This may be easier to use from language bindings.
	//
	// If count is zero, returns a zero-length #GBytes and does nothing. A value
	// of @count larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT
	// error.
	//
	// On success, a new #GBytes is returned. It is not an error if the size of
	// this object is not the same as the requested size, as it can happen e.g.
	// near the end of a file. A zero-length #GBytes is returned on end of file
	// (or if @count is zero), but never otherwise.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error nil is returned and @error is set accordingly.
	ReadBytes(count uint, cancellable Cancellable) (bytes *glib.Bytes, err error)
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
	ReadBytesAsync()
	// ReadBytesFinish finishes an asynchronous stream read-into-#GBytes
	// operation.
	ReadBytesFinish(result AsyncResult) (bytes *glib.Bytes, err error)
	// ReadFinish finishes an asynchronous stream read operation.
	ReadFinish(result AsyncResult) (gssize int, err error)
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
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
	Skip(count uint, cancellable Cancellable) (gssize int, err error)
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
	SkipAsync()
	// SkipFinish finishes a stream skip operation.
	SkipFinish(result AsyncResult) (gssize int, err error)
}

// inputStream implements the InputStream interface.
type inputStream struct {
	gextras.Objector
}

var _ InputStream = (*inputStream)(nil)

// WrapInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapInputStream(obj *externglib.Object) InputStream {
	return InputStream{
		Objector: obj,
	}
}

func marshalInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInputStream(obj), nil
}

// ClearPending clears the pending flag on @stream.
func (s inputStream) ClearPending() {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	C.g_input_stream_clear_pending(arg0)
}

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
func (s inputStream) Close(cancellable Cancellable) error {
	var arg0 *C.GInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError
	var goerr error

	C.g_input_stream_close(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

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
func (s inputStream) CloseAsync() {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	C.g_input_stream_close_async(arg0, arg1, arg2, arg3, arg4)
}

// CloseFinish finishes closing a stream asynchronously, started from
// g_input_stream_close_async().
func (s inputStream) CloseFinish(result AsyncResult) error {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.g_input_stream_close_finish(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// HasPending checks if an input stream has pending actions.
func (s inputStream) HasPending() bool {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_input_stream_has_pending(arg0)

	if cret {
		goret = true
	}

	return goret
}

// IsClosed checks if an input stream is closed.
func (s inputStream) IsClosed() bool {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_input_stream_is_closed(arg0)

	if cret {
		goret = true
	}

	return goret
}

// ReadAllFinish finishes an asynchronous stream read operation started with
// g_input_stream_read_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns false (and sets @error) then
// @bytes_read will be set to the number of bytes that were successfully
// read before the error was encountered. This functionality is only
// available from C. If you need it from another language then you must
// write your own loop around g_input_stream_read_async().
func (s inputStream) ReadAllFinish(result AsyncResult) (bytesRead uint, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	arg2 := new(C.gsize)
	var ret2 uint
	var cerr *C.GError
	var goerr error

	C.g_input_stream_read_all_finish(arg0, arg1, arg2, &cerr)

	ret2 = uint(*arg2)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return ret2, goerr
}

// ReadBytes: like g_input_stream_read(), this tries to read @count bytes
// from the stream in a blocking fashion. However, rather than reading into
// a user-supplied buffer, this will create a new #GBytes containing the
// data that was read. This may be easier to use from language bindings.
//
// If count is zero, returns a zero-length #GBytes and does nothing. A value
// of @count larger than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT
// error.
//
// On success, a new #GBytes is returned. It is not an error if the size of
// this object is not the same as the requested size, as it can happen e.g.
// near the end of a file. A zero-length #GBytes is returned on end of file
// (or if @count is zero), but never otherwise.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error nil is returned and @error is set accordingly.
func (s inputStream) ReadBytes(count uint, cancellable Cancellable) (bytes *glib.Bytes, err error) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 *C.GCancellable

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gsize(count)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	cret := new(C.GBytes)
	var goret *glib.Bytes
	var cerr *C.GError
	var goerr error

	cret = C.g_input_stream_read_bytes(arg0, arg1, arg2, &cerr)

	goret = glib.WrapBytes(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.Bytes) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

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
func (s inputStream) ReadBytesAsync() {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	C.g_input_stream_read_bytes_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReadBytesFinish finishes an asynchronous stream read-into-#GBytes
// operation.
func (s inputStream) ReadBytesFinish(result AsyncResult) (bytes *glib.Bytes, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	cret := new(C.GBytes)
	var goret *glib.Bytes
	var cerr *C.GError
	var goerr error

	cret = C.g_input_stream_read_bytes_finish(arg0, arg1, &cerr)

	goret = glib.WrapBytes(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.Bytes) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

// ReadFinish finishes an asynchronous stream read operation.
func (s inputStream) ReadFinish(result AsyncResult) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.gssize
	var goret int
	var cerr *C.GError
	var goerr error

	cret = C.g_input_stream_read_finish(arg0, arg1, &cerr)

	goret = int(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

// SetPending sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return false and set @error.
func (s inputStream) SetPending() error {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	var cerr *C.GError
	var goerr error

	C.g_input_stream_set_pending(arg0, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

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
func (s inputStream) Skip(count uint, cancellable Cancellable) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 C.gsize
	var arg2 *C.GCancellable

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gsize(count)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.gssize
	var goret int
	var cerr *C.GError
	var goerr error

	cret = C.g_input_stream_skip(arg0, arg1, arg2, &cerr)

	goret = int(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

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
func (s inputStream) SkipAsync() {
	var arg0 *C.GInputStream

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))

	C.g_input_stream_skip_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// SkipFinish finishes a stream skip operation.
func (s inputStream) SkipFinish(result AsyncResult) (gssize int, err error) {
	var arg0 *C.GInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.gssize
	var goret int
	var cerr *C.GError
	var goerr error

	cret = C.g_input_stream_skip_finish(arg0, arg1, &cerr)

	goret = int(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}
