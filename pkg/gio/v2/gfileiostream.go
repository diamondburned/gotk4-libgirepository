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
		{T: externglib.Type(C.g_file_io_stream_get_type()), F: marshalFileIOStream},
	})
}

// FileIOStream provides io streams that both read and write to the same file
// handle.
//
// GFileIOStream implements #GSeekable, which allows the io stream to jump to
// arbitrary positions in the file and to truncate the file, provided the
// filesystem of the file supports these operations.
//
// To find the position of a file io stream, use g_seekable_tell().
//
// To find out if a file io stream supports seeking, use g_seekable_can_seek().
// To position a file io stream, use g_seekable_seek(). To find out if a file io
// stream supports truncating, use g_seekable_can_truncate(). To truncate a file
// io stream, use g_seekable_truncate().
//
// The default implementation of all the IOStream operations and the
// implementation of #GSeekable just call into the same operations on the output
// stream.
type FileIOStream interface {
	IOStream

	// AsSeekable casts the class to the Seekable interface.
	AsSeekable() Seekable

	// Etag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	Etag() string
	// QueryInfoFileIOStream queries a file io stream for the given @attributes.
	// This function blocks while querying the stream. For the asynchronous
	// version of this function, see g_file_io_stream_query_info_async(). While
	// the stream is blocked, the stream will set the pending flag internally,
	// and any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with @error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED). I
	// all cases of failure, nil will be returned.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
	// be returned.
	QueryInfoFileIOStream(attributes string, cancellable Cancellable) (FileInfo, error)
	// QueryInfoAsyncFileIOStream: asynchronously queries the @stream for a
	// Info. When completed, @callback will be called with a Result which can be
	// used to finish the operation with g_file_io_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_io_stream_query_info().
	QueryInfoAsyncFileIOStream(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinishFileIOStream finalizes the asynchronous query started by
	// g_file_io_stream_query_info_async().
	QueryInfoFinishFileIOStream(result AsyncResult) (FileInfo, error)
}

// fileIOStream implements the FileIOStream class.
type fileIOStream struct {
	IOStream
}

// WrapFileIOStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileIOStream(obj *externglib.Object) FileIOStream {
	return fileIOStream{
		IOStream: WrapIOStream(obj),
	}
}

func marshalFileIOStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileIOStream(obj), nil
}

func (s fileIOStream) Etag() string {
	var _arg0 *C.GFileIOStream // out
	var _cret *C.char          // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_file_io_stream_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (s fileIOStream) QueryInfoFileIOStream(attributes string, cancellable Cancellable) (FileInfo, error) {
	var _arg0 *C.GFileIOStream // out
	var _arg1 *C.char          // out
	var _arg2 *C.GCancellable  // out
	var _cret *C.GFileInfo     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_file_io_stream_query_info(_arg0, _arg1, _arg2, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

func (s fileIOStream) QueryInfoAsyncFileIOStream(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileIOStream      // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_file_io_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s fileIOStream) QueryInfoFinishFileIOStream(result AsyncResult) (FileInfo, error) {
	var _arg0 *C.GFileIOStream // out
	var _arg1 *C.GAsyncResult  // out
	var _cret *C.GFileInfo     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GFileIOStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_file_io_stream_query_info_finish(_arg0, _arg1, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

func (f fileIOStream) AsSeekable() Seekable {
	return WrapSeekable(gextras.InternObject(f))
}
