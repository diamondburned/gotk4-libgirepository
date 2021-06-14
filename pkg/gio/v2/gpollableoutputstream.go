// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_pollable_output_stream_get_type()), F: marshalPollableOutputStream},
	})
}

// PollableOutputStreamOverrider contains methods that are overridable. This
// interface is a subset of the interface PollableOutputStream.
type PollableOutputStreamOverrider interface {
	// CanPoll checks if @stream is actually pollable. Some classes may
	// implement OutputStream but have only certain instances of that class be
	// pollable. If this method returns false, then the behavior of other
	// OutputStream methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	CanPoll() bool
	// IsWritable checks if @stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns true would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	IsWritable() bool
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
type PollableOutputStream interface {
	OutputStream
	PollableOutputStreamOverrider
}

// pollableOutputStream implements the PollableOutputStream interface.
type pollableOutputStream struct {
	OutputStream
}

var _ PollableOutputStream = (*pollableOutputStream)(nil)

// WrapPollableOutputStream wraps a GObject to a type that implements interface
// PollableOutputStream. It is primarily used internally.
func WrapPollableOutputStream(obj *externglib.Object) PollableOutputStream {
	return PollableOutputStream{
		OutputStream: WrapOutputStream(obj),
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPollableOutputStream(obj), nil
}

// CanPoll checks if @stream is actually pollable. Some classes may
// implement OutputStream but have only certain instances of that class be
// pollable. If this method returns false, then the behavior of other
// OutputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant; a
// stream cannot switch from pollable to non-pollable or vice versa.
func (s pollableOutputStream) CanPoll() bool {
	var _arg0 *C.GPollableOutputStream // out

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.g_pollable_output_stream_can_poll(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsWritable checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_output_stream_write() after
// this returns true would still block. To guarantee non-blocking behavior,
// you should always use g_pollable_output_stream_write_nonblocking(), which
// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (s pollableOutputStream) IsWritable() bool {
	var _arg0 *C.GPollableOutputStream // out

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.g_pollable_output_stream_is_writable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
