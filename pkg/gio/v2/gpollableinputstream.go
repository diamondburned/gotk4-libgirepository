// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GSource* _gotk4_gio2_PollableInputStreamInterface_create_source(GPollableInputStream*, GCancellable*);
// extern gboolean _gotk4_gio2_PollableInputStreamInterface_can_poll(GPollableInputStream*);
// extern gboolean _gotk4_gio2_PollableInputStreamInterface_is_readable(GPollableInputStream*);
// extern gssize _gotk4_gio2_PollableInputStreamInterface_read_nonblocking(GPollableInputStream*, void*, gsize, GError**);
import "C"

// glib.Type values for gpollableinputstream.go.
var GTypePollableInputStream = externglib.Type(C.g_pollable_input_stream_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePollableInputStream, F: marshalPollableInputStream},
	})
}

// PollableInputStreamOverrider contains methods that are overridable.
type PollableInputStreamOverrider interface {
	// CanPoll checks if stream is actually pollable. Some classes may implement
	// InputStream but have only certain instances of that class be pollable. If
	// this method returns FALSE, then the behavior of other InputStream methods
	// is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is pollable, FALSE if not.
	//
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be read, or
	// cancellable is triggered or an error occurs. The callback on the source
	// is of the SourceFunc type.
	//
	// As with g_pollable_input_stream_is_readable(), it is possible that the
	// stream may not actually be readable even after the source triggers, so
	// you should use g_pollable_input_stream_read_nonblocking() rather than
	// g_input_stream_read() from the callback.
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
	// IsReadable checks if stream can be read.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_input_stream_read() after
	// this returns TRUE would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_input_stream_read_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if stream is readable, FALSE if not. If an error has
	//      occurred on stream, this will result in
	//      g_pollable_input_stream_is_readable() returning TRUE, and the next
	//      attempt to read will return the error.
	//
	IsReadable() bool
	// ReadNonblocking attempts to read up to count bytes from stream into
	// buffer, as with g_input_stream_read(). If stream is not currently
	// readable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_input_stream_create_source() to create a #GSource that
	// will be triggered when stream is readable.
	//
	// Note that since this method never blocks, you cannot actually use
	// cancellable to cancel it. However, it will return an error if cancellable
	// has already been cancelled when you call, which may happen if you call
	// this method after a source triggers due to having been cancelled.
	//
	// The function takes the following parameters:
	//
	//    - buffer (optional) to read data into (which should be at least count
	//      bytes long).
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes read, or -1 on error (including
	//      G_IO_ERROR_WOULD_BLOCK).
	//
	ReadNonblocking(buffer []byte) (int, error)
}

// PollableInputStream is implemented by Streams that can be polled for
// readiness to read. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
//
// PollableInputStream wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PollableInputStream struct {
	_ [0]func() // equal guard
	InputStream
}

var (
	_ InputStreamer = (*PollableInputStream)(nil)
)

// PollableInputStreamer describes PollableInputStream's interface methods.
type PollableInputStreamer interface {
	externglib.Objector

	// CanPoll checks if stream is actually pollable.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be read, or
	// cancellable is triggered or an error occurs.
	CreateSource(ctx context.Context) *glib.Source
	// IsReadable checks if stream can be read.
	IsReadable() bool
	// ReadNonblocking attempts to read up to count bytes from stream into
	// buffer, as with g_input_stream_read().
	ReadNonblocking(ctx context.Context, buffer []byte) (int, error)
}

var _ PollableInputStreamer = (*PollableInputStream)(nil)

func ifaceInitPollableInputStreamer(gifacePtr, data C.gpointer) {
	iface := (*C.GPollableInputStreamInterface)(unsafe.Pointer(gifacePtr))
	iface.can_poll = (*[0]byte)(C._gotk4_gio2_PollableInputStreamInterface_can_poll)
	iface.create_source = (*[0]byte)(C._gotk4_gio2_PollableInputStreamInterface_create_source)
	iface.is_readable = (*[0]byte)(C._gotk4_gio2_PollableInputStreamInterface_is_readable)
	iface.read_nonblocking = (*[0]byte)(C._gotk4_gio2_PollableInputStreamInterface_read_nonblocking)
}

//export _gotk4_gio2_PollableInputStreamInterface_can_poll
func _gotk4_gio2_PollableInputStreamInterface_can_poll(arg0 *C.GPollableInputStream) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	ok := iface.CanPoll()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableInputStreamInterface_create_source
func _gotk4_gio2_PollableInputStreamInterface_create_source(arg0 *C.GPollableInputStream, arg1 *C.GCancellable) (cret *C.GSource) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	source := iface.CreateSource(_cancellable)

	cret = (*C.GSource)(gextras.StructNative(unsafe.Pointer(source)))

	return cret
}

//export _gotk4_gio2_PollableInputStreamInterface_is_readable
func _gotk4_gio2_PollableInputStreamInterface_is_readable(arg0 *C.GPollableInputStream) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	ok := iface.IsReadable()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_PollableInputStreamInterface_read_nonblocking
func _gotk4_gio2_PollableInputStreamInterface_read_nonblocking(arg0 *C.GPollableInputStream, arg1 *C.void, arg2 C.gsize, _cerr **C.GError) (cret C.gssize) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PollableInputStreamOverrider)

	var _buffer []byte // out

	if arg1 != nil {
		_buffer = make([]byte, arg2)
		copy(_buffer, unsafe.Slice((*byte)(unsafe.Pointer(arg1)), arg2))
	}

	gssize, _goerr := iface.ReadNonblocking(_buffer)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapPollableInputStream(obj *externglib.Object) *PollableInputStream {
	return &PollableInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalPollableInputStream(p uintptr) (interface{}, error) {
	return wrapPollableInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanPoll checks if stream is actually pollable. Some classes may implement
// InputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other InputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableInputStream) CanPoll() bool {
	var _arg0 *C.GPollableInputStream // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))

	_cret = C.g_pollable_input_stream_can_poll(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when stream can be read, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_input_stream_is_readable(), it is possible that the stream
// may not actually be readable even after the source triggers, so you should
// use g_pollable_input_stream_read_nonblocking() rather than
// g_input_stream_read() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableInputStream) CreateSource(ctx context.Context) *glib.Source {
	var _arg0 *C.GPollableInputStream // out
	var _arg1 *C.GCancellable         // out
	var _cret *C.GSource              // in

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_pollable_input_stream_create_source(_arg0, _arg1)
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

// IsReadable checks if stream can be read.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_input_stream_read() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_input_stream_read_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is readable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_input_stream_is_readable()
//      returning TRUE, and the next attempt to read will return the error.
//
func (stream *PollableInputStream) IsReadable() bool {
	var _arg0 *C.GPollableInputStream // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))

	_cret = C.g_pollable_input_stream_is_readable(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReadNonblocking attempts to read up to count bytes from stream into buffer,
// as with g_input_stream_read(). If stream is not currently readable, this will
// immediately return G_IO_ERROR_WOULD_BLOCK, and you can use
// g_pollable_input_stream_create_source() to create a #GSource that will be
// triggered when stream is readable.
//
// Note that since this method never blocks, you cannot actually use cancellable
// to cancel it. However, it will return an error if cancellable has already
// been cancelled when you call, which may happen if you call this method after
// a source triggers due to having been cancelled.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - buffer to read data into (which should be at least count bytes long).
//
// The function returns the following values:
//
//    - gssize: number of bytes read, or -1 on error (including
//      G_IO_ERROR_WOULD_BLOCK).
//
func (stream *PollableInputStream) ReadNonblocking(ctx context.Context, buffer []byte) (int, error) {
	var _arg0 *C.GPollableInputStream // out
	var _arg3 *C.GCancellable         // out
	var _arg1 *C.void                 // out
	var _arg2 C.gsize
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.g_pollable_input_stream_read_nonblocking(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, &_cerr)
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
