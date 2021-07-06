// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_file_output_stream_get_type()), F: marshalFileOutputStream},
	})
}

// FileOutputStream provides output streams that write their content to a file.
//
// GFileOutputStream implements #GSeekable, which allows the output stream to
// jump to arbitrary positions in the file and to truncate the file, provided
// the filesystem of the file supports these operations.
//
// To find the position of a file output stream, use g_seekable_tell(). To find
// out if a file output stream supports seeking, use g_seekable_can_seek().To
// position a file output stream, use g_seekable_seek(). To find out if a file
// output stream supports truncating, use g_seekable_can_truncate(). To truncate
// a file output stream, use g_seekable_truncate().
type FileOutputStream interface {
	gextras.Objector

	// AsOutputStream casts the class to the OutputStream interface.
	AsOutputStream() OutputStream
	// AsSeekable casts the class to the Seekable interface.
	AsSeekable() Seekable

	// ClearPending clears the pending flag on @stream.
	//
	// This method is inherited from OutputStream
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Closing a stream will automatically flush any outstanding buffers in the
	// stream.
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
	// important to check and report the error to the user, otherwise there
	// might be a loss of data as all data might not be written.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but there some
	// streams can use a faster close that doesn't block to e.g. check errors.
	// On cancellation (as with any error) there is no guarantee that all
	// written data will reach the target.
	//
	// This method is inherited from OutputStream
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous close of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_output_stream_close_finish() to get the
	// result of the operation.
	//
	// For behaviour details see g_output_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// This method is inherited from OutputStream
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish closes an output stream.
	//
	// This method is inherited from OutputStream
	CloseFinish(result AsyncResult) error
	// Flush forces a write of all user-space buffered data for the given
	// @stream. Will block during the operation. Closing the stream will
	// implicitly cause a flush.
	//
	// This function is optional for inherited classes.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	//
	// This method is inherited from OutputStream
	Flush(cancellable Cancellable) error
	// FlushAsync forces an asynchronous write of all user-space buffered data
	// for the given @stream. For behaviour details see g_output_stream_flush().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_output_stream_flush_finish() to get the result of the operation.
	//
	// This method is inherited from OutputStream
	FlushAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// FlushFinish finishes flushing an output stream.
	//
	// This method is inherited from OutputStream
	FlushFinish(result AsyncResult) error
	// HasPending checks if an output stream has pending actions.
	//
	// This method is inherited from OutputStream
	HasPending() bool
	// IsClosed checks if an output stream has already been closed.
	//
	// This method is inherited from OutputStream
	IsClosed() bool
	// IsClosing checks if an output stream is being closed. This can be used
	// inside e.g. a flush implementation to see if the flush (or other i/o
	// operation) is called from within the closing operation.
	//
	// This method is inherited from OutputStream
	IsClosing() bool
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	//
	// This method is inherited from OutputStream
	SetPending() error
	// Splice splices an input stream into an output stream.
	//
	// This method is inherited from OutputStream
	Splice(source InputStream, flags OutputStreamSpliceFlags, cancellable Cancellable) (int, error)
	// SpliceAsync splices a stream asynchronously. When the operation is
	// finished @callback will be called. You can then call
	// g_output_stream_splice_finish() to get the result of the operation.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_splice().
	//
	// This method is inherited from OutputStream
	SpliceAsync(source InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// SpliceFinish finishes an asynchronous stream splice operation.
	//
	// This method is inherited from OutputStream
	SpliceFinish(result AsyncResult) (int, error)
	// Write tries to write @count bytes from @buffer into the stream. Will
	// block during the operation.
	//
	// If count is 0, returns 0 and does nothing. A value of @count larger than
	// G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes written to the stream is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. on a partial I/O error, or if there is not enough storage in
	// the stream. All writes block until at least one byte is written or an
	// error occurs; 0 is never returned (unless @count is 0).
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and @error is set accordingly.
	//
	// This method is inherited from OutputStream
	Write(buffer []byte, cancellable Cancellable) (int, error)
	// WriteAll tries to write @count bytes from @buffer into the stream. Will
	// block during the operation.
	//
	// This function is similar to g_output_stream_write(), except it tries to
	// write as many bytes as requested, only stopping on an error.
	//
	// On a successful write of @count bytes, true is returned, and
	// @bytes_written is set to @count.
	//
	// If there is an error during the operation false is returned and @error is
	// set to indicate the error status.
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_write().
	//
	// This method is inherited from OutputStream
	WriteAll(buffer []byte, cancellable Cancellable) (uint, error)
	// WriteAllAsync: request an asynchronous write of @count bytes from @buffer
	// into the stream. When the operation is finished @callback will be called.
	// You can then call g_output_stream_write_all_finish() to get the result of
	// the operation.
	//
	// This is the asynchronous version of g_output_stream_write_all().
	//
	// Call g_output_stream_write_all_finish() to collect the result.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// Note that no copy of @buffer will be made, so it must stay valid until
	// @callback is called.
	//
	// This method is inherited from OutputStream
	WriteAllAsync(buffer []byte, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// WriteAllFinish finishes an asynchronous stream write operation started
	// with g_output_stream_write_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_write_async().
	//
	// This method is inherited from OutputStream
	WriteAllFinish(result AsyncResult) (uint, error)
	// WriteAsync: request an asynchronous write of @count bytes from @buffer
	// into the stream. When the operation is finished @callback will be called.
	// You can then call g_output_stream_write_finish() to get the result of the
	// operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes written will be passed to the @callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. on a partial I/O error, but generally we try to write as
	// many bytes as requested.
	//
	// You are guaranteed that this method will never fail with
	// G_IO_ERROR_WOULD_BLOCK - if @stream can't accept more data, the method
	// will just wait until this changes.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_write().
	//
	// Note that no copy of @buffer will be made, so it must stay valid until
	// @callback is called. See g_output_stream_write_bytes_async() for a
	// #GBytes version that will automatically hold a reference to the contents
	// (without copying) for the duration of the call.
	//
	// This method is inherited from OutputStream
	WriteAsync(buffer []byte, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// WriteBytesFinish finishes a stream write-from-#GBytes operation.
	//
	// This method is inherited from OutputStream
	WriteBytesFinish(result AsyncResult) (int, error)
	// WriteFinish finishes a stream write operation.
	//
	// This method is inherited from OutputStream
	WriteFinish(result AsyncResult) (int, error)
	// Writev tries to write the bytes contained in the @n_vectors @vectors into
	// the stream. Will block during the operation.
	//
	// If @n_vectors is 0 or the sum of all bytes in @vectors is 0, returns 0
	// and does nothing.
	//
	// On success, the number of bytes written to the stream is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. on a partial I/O error, or if there is not enough storage in
	// the stream. All writes block until at least one byte is written or an
	// error occurs; 0 is never returned (unless @n_vectors is 0 or the sum of
	// all bytes in @vectors is 0).
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// Some implementations of g_output_stream_writev() may have limitations on
	// the aggregate buffer size, and will return G_IO_ERROR_INVALID_ARGUMENT if
	// these are exceeded. For example, when writing to a local file on UNIX
	// platforms, the aggregate buffer size must not exceed G_MAXSSIZE bytes.
	//
	// This method is inherited from OutputStream
	Writev(vectors []OutputVector, cancellable Cancellable) (uint, error)
	// WritevAll tries to write the bytes contained in the @n_vectors @vectors
	// into the stream. Will block during the operation.
	//
	// This function is similar to g_output_stream_writev(), except it tries to
	// write as many bytes as requested, only stopping on an error.
	//
	// On a successful write of all @n_vectors vectors, true is returned, and
	// @bytes_written is set to the sum of all the sizes of @vectors.
	//
	// If there is an error during the operation false is returned and @error is
	// set to indicate the error status.
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_write().
	//
	// The content of the individual elements of @vectors might be changed by
	// this function.
	//
	// This method is inherited from OutputStream
	WritevAll(vectors []OutputVector, cancellable Cancellable) (uint, error)
	// WritevAllAsync: request an asynchronous write of the bytes contained in
	// the @n_vectors @vectors into the stream. When the operation is finished
	// @callback will be called. You can then call
	// g_output_stream_writev_all_finish() to get the result of the operation.
	//
	// This is the asynchronous version of g_output_stream_writev_all().
	//
	// Call g_output_stream_writev_all_finish() to collect the result.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// Note that no copy of @vectors will be made, so it must stay valid until
	// @callback is called. The content of the individual elements of @vectors
	// might be changed by this function.
	//
	// This method is inherited from OutputStream
	WritevAllAsync(vectors []OutputVector, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// WritevAllFinish finishes an asynchronous stream write operation started
	// with g_output_stream_writev_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_writev_async().
	//
	// This method is inherited from OutputStream
	WritevAllFinish(result AsyncResult) (uint, error)
	// WritevAsync: request an asynchronous write of the bytes contained in
	// @n_vectors @vectors into the stream. When the operation is finished
	// @callback will be called. You can then call
	// g_output_stream_writev_finish() to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// On success, the number of bytes written will be passed to the @callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. on a partial I/O error, but generally we try to write as
	// many bytes as requested.
	//
	// You are guaranteed that this method will never fail with
	// G_IO_ERROR_WOULD_BLOCK — if @stream can't accept more data, the method
	// will just wait until this changes.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_writev().
	//
	// Note that no copy of @vectors will be made, so it must stay valid until
	// @callback is called.
	//
	// This method is inherited from OutputStream
	WritevAsync(vectors []OutputVector, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// WritevFinish finishes a stream writev operation.
	//
	// This method is inherited from OutputStream
	WritevFinish(result AsyncResult) (uint, error)
	// CanSeek tests if the stream supports the Iface.
	//
	// This method is inherited from Seekable
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	//
	// This method is inherited from Seekable
	CanTruncate() bool
	// Seek seeks in the stream by the given @offset, modified by @type.
	//
	// Attempting to seek past the end of the stream will have different results
	// depending on if the stream is fixed-sized or resizable. If the stream is
	// resizable then seeking past the end and then writing will result in zeros
	// filling the empty space. Seeking past the end of a resizable stream and
	// reading will result in EOF. Seeking past the end of a fixed-sized stream
	// will fail.
	//
	// Any operation that would result in a negative offset will fail.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	//
	// This method is inherited from Seekable
	Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error
	// Tell tells the current position within the stream.
	//
	// This method is inherited from Seekable
	Tell() int64
	// Truncate sets the length of the stream to @offset. If the stream was
	// previously larger than @offset, the extra data is discarded. If the
	// stream was previously shorter than @offset, it is extended with NUL
	// ('\0') bytes.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// This method is inherited from Seekable
	Truncate(offset int64, cancellable Cancellable) error

	// Etag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	Etag() string
	// QueryInfo queries a file output stream for the given @attributes. This
	// function blocks while querying the stream. For the asynchronous version
	// of this function, see g_file_output_stream_query_info_async(). While the
	// stream is blocked, the stream will set the pending flag internally, and
	// any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with @error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED).
	// In all cases of failure, nil will be returned.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
	// be returned.
	QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error)
	// QueryInfoAsync: asynchronously queries the @stream for a Info. When
	// completed, @callback will be called with a Result which can be used to
	// finish the operation with g_file_output_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_output_stream_query_info().
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_output_stream_query_info_async().
	QueryInfoFinish(result AsyncResult) (FileInfo, error)
}

// fileOutputStream implements the FileOutputStream interface.
type fileOutputStream struct {
	*externglib.Object
}

var _ FileOutputStream = (*fileOutputStream)(nil)

// WrapFileOutputStream wraps a GObject to a type that implements
// interface FileOutputStream. It is primarily used internally.
func WrapFileOutputStream(obj *externglib.Object) FileOutputStream {
	return fileOutputStream{obj}
}

func marshalFileOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileOutputStream(obj), nil
}

func (f fileOutputStream) AsOutputStream() OutputStream {
	return WrapOutputStream(gextras.InternObject(f))
}

func (f fileOutputStream) AsSeekable() Seekable {
	return WrapSeekable(gextras.InternObject(f))
}

func (s fileOutputStream) ClearPending() {
	WrapOutputStream(gextras.InternObject(s)).ClearPending()
}

func (s fileOutputStream) Close(cancellable Cancellable) error {
	return WrapOutputStream(gextras.InternObject(s)).Close(cancellable)
}

func (s fileOutputStream) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).CloseAsync(ioPriority, cancellable, callback)
}

func (s fileOutputStream) CloseFinish(result AsyncResult) error {
	return WrapOutputStream(gextras.InternObject(s)).CloseFinish(result)
}

func (s fileOutputStream) Flush(cancellable Cancellable) error {
	return WrapOutputStream(gextras.InternObject(s)).Flush(cancellable)
}

func (s fileOutputStream) FlushAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).FlushAsync(ioPriority, cancellable, callback)
}

func (s fileOutputStream) FlushFinish(result AsyncResult) error {
	return WrapOutputStream(gextras.InternObject(s)).FlushFinish(result)
}

func (s fileOutputStream) HasPending() bool {
	return WrapOutputStream(gextras.InternObject(s)).HasPending()
}

func (s fileOutputStream) IsClosed() bool {
	return WrapOutputStream(gextras.InternObject(s)).IsClosed()
}

func (s fileOutputStream) IsClosing() bool {
	return WrapOutputStream(gextras.InternObject(s)).IsClosing()
}

func (s fileOutputStream) SetPending() error {
	return WrapOutputStream(gextras.InternObject(s)).SetPending()
}

func (s fileOutputStream) Splice(source InputStream, flags OutputStreamSpliceFlags, cancellable Cancellable) (int, error) {
	return WrapOutputStream(gextras.InternObject(s)).Splice(source, flags, cancellable)
}

func (s fileOutputStream) SpliceAsync(source InputStream, flags OutputStreamSpliceFlags, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).SpliceAsync(source, flags, ioPriority, cancellable, callback)
}

func (s fileOutputStream) SpliceFinish(result AsyncResult) (int, error) {
	return WrapOutputStream(gextras.InternObject(s)).SpliceFinish(result)
}

func (s fileOutputStream) Write(buffer []byte, cancellable Cancellable) (int, error) {
	return WrapOutputStream(gextras.InternObject(s)).Write(buffer, cancellable)
}

func (s fileOutputStream) WriteAll(buffer []byte, cancellable Cancellable) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).WriteAll(buffer, cancellable)
}

func (s fileOutputStream) WriteAllAsync(buffer []byte, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).WriteAllAsync(buffer, ioPriority, cancellable, callback)
}

func (s fileOutputStream) WriteAllFinish(result AsyncResult) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).WriteAllFinish(result)
}

func (s fileOutputStream) WriteAsync(buffer []byte, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).WriteAsync(buffer, ioPriority, cancellable, callback)
}

func (s fileOutputStream) WriteBytesFinish(result AsyncResult) (int, error) {
	return WrapOutputStream(gextras.InternObject(s)).WriteBytesFinish(result)
}

func (s fileOutputStream) WriteFinish(result AsyncResult) (int, error) {
	return WrapOutputStream(gextras.InternObject(s)).WriteFinish(result)
}

func (s fileOutputStream) Writev(vectors []OutputVector, cancellable Cancellable) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).Writev(vectors, cancellable)
}

func (s fileOutputStream) WritevAll(vectors []OutputVector, cancellable Cancellable) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).WritevAll(vectors, cancellable)
}

func (s fileOutputStream) WritevAllAsync(vectors []OutputVector, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).WritevAllAsync(vectors, ioPriority, cancellable, callback)
}

func (s fileOutputStream) WritevAllFinish(result AsyncResult) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).WritevAllFinish(result)
}

func (s fileOutputStream) WritevAsync(vectors []OutputVector, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapOutputStream(gextras.InternObject(s)).WritevAsync(vectors, ioPriority, cancellable, callback)
}

func (s fileOutputStream) WritevFinish(result AsyncResult) (uint, error) {
	return WrapOutputStream(gextras.InternObject(s)).WritevFinish(result)
}

func (s fileOutputStream) CanSeek() bool {
	return WrapSeekable(gextras.InternObject(s)).CanSeek()
}

func (s fileOutputStream) CanTruncate() bool {
	return WrapSeekable(gextras.InternObject(s)).CanTruncate()
}

func (s fileOutputStream) Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error {
	return WrapSeekable(gextras.InternObject(s)).Seek(offset, typ, cancellable)
}

func (s fileOutputStream) Tell() int64 {
	return WrapSeekable(gextras.InternObject(s)).Tell()
}

func (s fileOutputStream) Truncate(offset int64, cancellable Cancellable) error {
	return WrapSeekable(gextras.InternObject(s)).Truncate(offset, cancellable)
}

func (s fileOutputStream) Etag() string {
	var _arg0 *C.GFileOutputStream // out
	var _cret *C.char              // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_file_output_stream_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (s fileOutputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	var _arg0 *C.GFileOutputStream // out
	var _arg1 *C.char              // out
	var _arg2 *C.GCancellable      // out
	var _cret *C.GFileInfo         // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_file_output_stream_query_info(_arg0, _arg1, _arg2, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

func (s fileOutputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileOutputStream  // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_file_output_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s fileOutputStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var _arg0 *C.GFileOutputStream // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GFileInfo         // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_file_output_stream_query_info_finish(_arg0, _arg1, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}
