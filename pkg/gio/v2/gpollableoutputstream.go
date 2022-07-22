// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GPollableReturn _gotk4_gio2_PollableOutputStreamInterface_writev_nonblocking(GPollableOutputStream*, GOutputVector*, gsize, gsize*, GError**);
// extern GSource* _gotk4_gio2_PollableOutputStreamInterface_create_source(GPollableOutputStream*, GCancellable*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_can_poll(GPollableOutputStream*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_is_writable(GPollableOutputStream*);
// extern gssize _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking(GPollableOutputStream*, void*, gsize, GError**);
import "C"

// GType values.
var (
	GTypePollableOutputStream = coreglib.Type(C.g_pollable_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePollableOutputStream, F: marshalPollableOutputStream},
	})
}

// PollableOutputStreamOverrider contains methods that are overridable.
type PollableOutputStreamOverrider interface {
	// CanPoll checks if stream is actually pollable. Some classes may implement
	// OutputStream but have only certain instances of that class be pollable.
	// If this method returns FALSE, then the behavior of other OutputStream
	// methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is pollable, FALSE if not.
	//
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be written,
	// or cancellable is triggered or an error occurs. The callback on the
	// source is of the SourceFunc type.
	//
	// As with g_pollable_output_stream_is_writable(), it is possible that the
	// stream may not actually be writable even after the source triggers, so
	// you should use g_pollable_output_stream_write_nonblocking() rather than
	// g_output_stream_write() from the callback.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//
	// The function returns the following values:
	//
	//    - source: new #GSource.
	//
	CreateSource(ctx context.Context) *glib.Source
	// IsWritable checks if stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns TRUE would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is writable, FALSE if not. If an error has
	//      occurred on stream, this will result in
	//      g_pollable_output_stream_is_writable() returning TRUE, and the next
	//      attempt to write will return the error.
	//
	IsWritable() bool
	// WriteNonblocking attempts to write up to count bytes from buffer to
	// stream, as with g_output_stream_write(). If stream is not currently
	// writable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_output_stream_create_source() to create a #GSource
	// that will be triggered when stream is writable.
	//
	// Note that since this method never blocks, you cannot actually use
	// cancellable to cancel it. However, it will return an error if cancellable
	// has already been cancelled when you call, which may happen if you call
	// this method after a source triggers due to having been cancelled.
	//
	// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
	// transports like D/TLS require that you re-send the same buffer and count
	// in the next write call.
	//
	// The function takes the following parameters:
	//
	//    - buffer (optional) to write data from.
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes written, or -1 on error (including
	//      G_IO_ERROR_WOULD_BLOCK).
	//
	WriteNonblocking(buffer []byte) (int, error)
	// WritevNonblocking attempts to write the bytes contained in the n_vectors
	// vectors to stream, as with g_output_stream_writev(). If stream is not
	// currently writable, this will immediately return
	// G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
	// g_pollable_output_stream_create_source() to create a #GSource that will
	// be triggered when stream is writable. error will *not* be set in that
	// case.
	//
	// Note that since this method never blocks, you cannot actually use
	// cancellable to cancel it. However, it will return an error if cancellable
	// has already been cancelled when you call, which may happen if you call
	// this method after a source triggers due to having been cancelled.
	//
	// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some
	// underlying transports like D/TLS require that you re-send the same
	// vectors and n_vectors in the next write call.
	//
	// The function takes the following parameters:
	//
	//    - vectors: buffer containing the Vectors to write.
	//
	// The function returns the following values:
	//
	//    - bytesWritten (optional): location to store the number of bytes that
	//      were written to the stream.
	//    - pollableReturn: G_POLLABLE_RETURN_OK on success,
	//      G_POLLABLE_RETURN_WOULD_BLOCK if the stream is not currently writable
	//      (and error is *not* set), or G_POLLABLE_RETURN_FAILED if there was an
	//      error in which case error will be set.
	//
	WritevNonblocking(vectors []OutputVector) (uint, PollableReturn, error)
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
//
// PollableOutputStream wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PollableOutputStream struct {
	_ [0]func() // equal guard
	OutputStream
}

var (
	_ OutputStreamer = (*PollableOutputStream)(nil)
)

// PollableOutputStreamer describes PollableOutputStream's interface methods.
type PollableOutputStreamer interface {
	coreglib.Objector

	// CanPoll checks if stream is actually pollable.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be written,
	// or cancellable is triggered or an error occurs.
	CreateSource(ctx context.Context) *glib.Source
	// IsWritable checks if stream can be written.
	IsWritable() bool
	// WriteNonblocking attempts to write up to count bytes from buffer to
	// stream, as with g_output_stream_write().
	WriteNonblocking(ctx context.Context, buffer []byte) (int, error)
	// WritevNonblocking attempts to write the bytes contained in the n_vectors
	// vectors to stream, as with g_output_stream_writev().
	WritevNonblocking(ctx context.Context, vectors []OutputVector) (uint, PollableReturn, error)
}

var _ PollableOutputStreamer = (*PollableOutputStream)(nil)

func ifaceInitPollableOutputStreamer(gifacePtr, data C.gpointer) {
	iface := (*C.GPollableOutputStreamInterface)(unsafe.Pointer(gifacePtr))
	iface.can_poll = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_can_poll)
	iface.create_source = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_create_source)
	iface.is_writable = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_is_writable)
	iface.write_nonblocking = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_write_nonblocking)
	iface.writev_nonblocking = (*[0]byte)(C._gotk4_gio2_PollableOutputStreamInterface_writev_nonblocking)
}

//export _gotk4_gio2_PollableOutputStreamInterface_can_poll
func _gotk4_gio2_PollableOutputStreamInterface_can_poll(arg0 *C.GPollableOutputStream) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.CanPoll()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_create_source
func _gotk4_gio2_PollableOutputStreamInterface_create_source(arg0 *C.GPollableOutputStream, arg1 *C.GCancellable) (cret *C.GSource) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	source := iface.CreateSource(_cancellable)

	cret = (*C.GSource)(gextras.StructNative(unsafe.Pointer(source)))

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_is_writable
func _gotk4_gio2_PollableOutputStreamInterface_is_writable(arg0 *C.GPollableOutputStream) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.IsWritable()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking
func _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking(arg0 *C.GPollableOutputStream, arg1 *C.void, arg2 C.gsize, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	var _buffer []byte // out

	if arg1 != nil {
		_buffer = make([]byte, arg2)
		copy(_buffer, unsafe.Slice((*byte)(unsafe.Pointer(arg1)), arg2))
	}

	gssize, _goerr := iface.WriteNonblocking(_buffer)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_writev_nonblocking
func _gotk4_gio2_PollableOutputStreamInterface_writev_nonblocking(arg0 *C.GPollableOutputStream, arg1 *C.GOutputVector, arg2 C.gsize, arg3 *C.gsize, _cerr **C.GError) (cret C.GPollableReturn) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	var _vectors []OutputVector // out

	{
		src := unsafe.Slice((*C.GOutputVector)(arg1), arg2)
		_vectors = make([]OutputVector, arg2)
		for i := 0; i < int(arg2); i++ {
			_vectors[i] = *(*OutputVector)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}

	bytesWritten, pollableReturn, _goerr := iface.WritevNonblocking(_vectors)

	*arg3 = C.gsize(bytesWritten)
	cret = C.GPollableReturn(pollableReturn)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapPollableOutputStream(obj *coreglib.Object) *PollableOutputStream {
	return &PollableOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	return wrapPollableOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanPoll checks if stream is actually pollable. Some classes may implement
// OutputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other OutputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableOutputStream) CanPoll() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_pollable_output_stream_can_poll(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when stream can be written, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so you
// should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableOutputStream) CreateSource(ctx context.Context) *glib.Source {
	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GCancellable          // out
	var _cret *C.GSource               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_pollable_output_stream_create_source(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// IsWritable checks if stream can be written.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_output_stream_write() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_output_stream_write_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is writable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_output_stream_is_writable()
//      returning TRUE, and the next attempt to write will return the error.
//
func (stream *PollableOutputStream) IsWritable() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_pollable_output_stream_is_writable(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WriteNonblocking attempts to write up to count bytes from buffer to stream,
// as with g_output_stream_write(). If stream is not currently writable, this
// will immediately return G_IO_ERROR_WOULD_BLOCK, and you can use
// g_pollable_output_stream_create_source() to create a #GSource that will be
// triggered when stream is writable.
//
// Note that since this method never blocks, you cannot actually use cancellable
// to cancel it. However, it will return an error if cancellable has already
// been cancelled when you call, which may happen if you call this method after
// a source triggers due to having been cancelled.
//
// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same buffer and count in
// the next write call.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - buffer to write data from.
//
// The function returns the following values:
//
//    - gssize: number of bytes written, or -1 on error (including
//      G_IO_ERROR_WOULD_BLOCK).
//
func (stream *PollableOutputStream) WriteNonblocking(ctx context.Context, buffer []byte) (int, error) {
	var _arg0 *C.GPollableOutputStream // out
	var _arg3 *C.GCancellable          // out
	var _arg1 *C.void                  // out
	var _arg2 C.gsize
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.g_pollable_output_stream_write_nonblocking(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// WritevNonblocking attempts to write the bytes contained in the n_vectors
// vectors to stream, as with g_output_stream_writev(). If stream is not
// currently writable, this will immediately return
// G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
// g_pollable_output_stream_create_source() to create a #GSource that will be
// triggered when stream is writable. error will *not* be set in that case.
//
// Note that since this method never blocks, you cannot actually use cancellable
// to cancel it. However, it will return an error if cancellable has already
// been cancelled when you call, which may happen if you call this method after
// a source triggers due to having been cancelled.
//
// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same vectors and n_vectors
// in the next write call.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - vectors: buffer containing the Vectors to write.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that were
//      written to the stream.
//    - pollableReturn: G_POLLABLE_RETURN_OK on success,
//      G_POLLABLE_RETURN_WOULD_BLOCK if the stream is not currently writable
//      (and error is *not* set), or G_POLLABLE_RETURN_FAILED if there was an
//      error in which case error will be set.
//
func (stream *PollableOutputStream) WritevNonblocking(ctx context.Context, vectors []OutputVector) (uint, PollableReturn, error) {
	var _arg0 *C.GPollableOutputStream // out
	var _arg4 *C.GCancellable          // out
	var _arg1 *C.GOutputVector         // out
	var _arg2 C.gsize
	var _arg3 C.gsize           // in
	var _cret C.GPollableReturn // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(vectors))
	_arg1 = (*C.GOutputVector)(C.calloc(C.size_t(len(vectors)), C.size_t(C.sizeof_GOutputVector)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GOutputVector)(_arg1), len(vectors))
		for i := range vectors {
			out[i] = *(*C.GOutputVector)(gextras.StructNative(unsafe.Pointer((&vectors[i]))))
		}
	}

	_cret = C.g_pollable_output_stream_writev_nonblocking(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(vectors)

	var _bytesWritten uint             // out
	var _pollableReturn PollableReturn // out
	var _goerr error                   // out

	_bytesWritten = uint(_arg3)
	_pollableReturn = PollableReturn(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _pollableReturn, _goerr
}
