// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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

// NewPollableSource: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource that expects a callback of type
// SourceFunc. The new source does not actually do anything on its own; use
// g_source_add_child_source() to add other sources to it to cause it to
// trigger.
//
// The function takes the following parameters:
//
//    - pollableStream: stream associated with the new source.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func NewPollableSource(pollableStream *coreglib.Object) *glib.Source {
	var _args [1]girepository.Argument

	*(**C.GObject)(unsafe.Pointer(&_args[0])) = (*C.GObject)(unsafe.Pointer(pollableStream.Native()))

	_info := girepository.MustFind("Gio", "pollable_source_new")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GSource)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pollableStream)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(*(**C.GSource)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("GLib", "Source").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _source
}

// PollableSourceNewFull: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource, as with g_pollable_source_new(), but
// also attaching child_source (with a dummy callback), and cancellable, if they
// are non-NULL.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable to attach.
//    - pollableStream: stream associated with the new source.
//    - childSource (optional): optional child source to attach.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func PollableSourceNewFull(ctx context.Context, pollableStream *coreglib.Object, childSource *glib.Source) *glib.Source {
	var _args [3]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(unsafe.Pointer(pollableStream.Native()))
	if childSource != nil {
		*(**C.GSource)(unsafe.Pointer(&_args[1])) = (*C.GSource)(gextras.StructNative(unsafe.Pointer(childSource)))
	}

	_info := girepository.MustFind("Gio", "pollable_source_new_full")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GSource)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(pollableStream)
	runtime.KeepAlive(childSource)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(*(**C.GSource)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("GLib", "Source").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _source
}

// PollableStreamRead tries to read from stream, as with g_input_stream_read()
// (if blocking is TRUE) or g_pollable_input_stream_read_nonblocking() (if
// blocking is FALSE). This can be used to more easily share code between
// blocking and non-blocking implementations of a method.
//
// If blocking is FALSE, then stream must be a InputStream for which
// g_pollable_input_stream_can_poll() returns TRUE, or else the behavior is
// undefined. If blocking is TRUE, then stream does not need to be a
// InputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer to read data into.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - gssize: number of bytes read, or -1 on error.
//
func PollableStreamRead(ctx context.Context, stream InputStreamer, buffer []byte, blocking bool) (int, error) {
	var _args [5]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[4] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[2])) = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "pollable_stream_read")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GError)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _gssize, _goerr
}

// PollableStreamWrite tries to write to stream, as with g_output_stream_write()
// (if blocking is TRUE) or g_pollable_output_stream_write_nonblocking() (if
// blocking is FALSE). This can be used to more easily share code between
// blocking and non-blocking implementations of a method.
//
// If blocking is FALSE, then stream must be a OutputStream for which
// g_pollable_output_stream_can_poll() returns TRUE or else the behavior is
// undefined. If blocking is TRUE, then stream does not need to be a
// OutputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer: buffer containing the data to write.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - gssize: number of bytes written, or -1 on error.
//
func PollableStreamWrite(ctx context.Context, stream OutputStreamer, buffer []byte, blocking bool) (int, error) {
	var _args [5]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[4] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[2])) = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "pollable_stream_write")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GError)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _gssize, _goerr
}

// PollableStreamWriteAll tries to write count bytes to stream, as with
// g_output_stream_write_all(), but using g_pollable_stream_write() rather than
// g_output_stream_write().
//
// On a successful write of count bytes, TRUE is returned, and bytes_written is
// set to count.
//
// If there is an error during the operation (including G_IO_ERROR_WOULD_BLOCK
// in the non-blocking case), FALSE is returned and error is set to indicate the
// error status, bytes_written is updated to contain the number of bytes written
// into the stream before the error occurred.
//
// As with g_pollable_stream_write(), if blocking is FALSE, then stream must be
// a OutputStream for which g_pollable_output_stream_can_poll() returns TRUE or
// else the behavior is undefined. If blocking is TRUE, then stream does not
// need to be a OutputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer: buffer containing the data to write.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - bytesWritten: location to store the number of bytes that was written to
//      the stream.
//
func PollableStreamWriteAll(ctx context.Context, stream OutputStreamer, buffer []byte, blocking bool) (uint, error) {
	var _args [5]girepository.Argument
	var _outs [1]girepository.Argument

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[4] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[2])) = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "pollable_stream_write_all")
	_info.InvokeFunction(_args[:], _outs[:])

	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _bytesWritten uint // out
	var _goerr error       // out

	_bytesWritten = uint(*(*C.gsize)(unsafe.Pointer(&_outs[0])))
	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _bytesWritten, _goerr
}
