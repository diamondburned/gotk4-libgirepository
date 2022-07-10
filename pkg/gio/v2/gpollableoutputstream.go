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
// extern GSource* _gotk4_gio2_PollableOutputStreamInterface_create_source(void*, void*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_can_poll(void*);
// extern gboolean _gotk4_gio2_PollableOutputStreamInterface_is_writable(void*);
// extern gssize _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking(void*, void*, gsize, GError**);
import "C"

// GTypePollableOutputStream returns the GType for the type PollableOutputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePollableOutputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "PollableOutputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPollableOutputStream)
	return gtype
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
}

var _ PollableOutputStreamer = (*PollableOutputStream)(nil)

func ifaceInitPollableOutputStreamer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "PollableOutputStreamInterface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("can_poll"))) = unsafe.Pointer(C._gotk4_gio2_PollableOutputStreamInterface_can_poll)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("create_source"))) = unsafe.Pointer(C._gotk4_gio2_PollableOutputStreamInterface_create_source)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("is_writable"))) = unsafe.Pointer(C._gotk4_gio2_PollableOutputStreamInterface_is_writable)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("write_nonblocking"))) = unsafe.Pointer(C._gotk4_gio2_PollableOutputStreamInterface_write_nonblocking)
}

//export _gotk4_gio2_PollableOutputStreamInterface_can_poll
func _gotk4_gio2_PollableOutputStreamInterface_can_poll(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.CanPoll()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_create_source
func _gotk4_gio2_PollableOutputStreamInterface_create_source(arg0 *C.void, arg1 *C.void) (cret *C.GSource) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
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
func _gotk4_gio2_PollableOutputStreamInterface_is_writable(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableOutputStreamOverrider)

	ok := iface.IsWritable()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking
func _gotk4_gio2_PollableOutputStreamInterface_write_nonblocking(arg0 *C.void, arg1 *C.void, arg2 C.gsize, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "PollableOutputStream")
	_gret := _info.InvokeIfaceMethod("can_poll", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "PollableOutputStream")
	_gret := _info.InvokeIfaceMethod("create_source", _args[:], nil)
	_cret := *(**C.GSource)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "PollableOutputStream")
	_gret := _info.InvokeIfaceMethod("is_writable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gsize)(unsafe.Pointer(&_args[2])) = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(&buffer[0]))
	}

	_info := girepository.MustFind("Gio", "PollableOutputStream")
	_gret := _info.InvokeIfaceMethod("write_nonblocking", _args[:], nil)
	_cret := *(**C.GError)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _gssize, _goerr
}
