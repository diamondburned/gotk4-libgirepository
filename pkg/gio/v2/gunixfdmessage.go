// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_unix_fd_message_get_type()), F: marshalUnixFDMessage},
	})
}

// UnixFDMessage: this ControlMessage contains a FDList. It may be sent using
// g_socket_send_message() and received using g_socket_receive_message() over
// UNIX sockets (ie: sockets in the G_SOCKET_FAMILY_UNIX family). The file
// descriptors are copied between processes by the kernel.
//
// For an easier way to send and receive file descriptors over stream-oriented
// UNIX sockets, see g_unix_connection_send_fd() and
// g_unix_connection_receive_fd().
//
// Note that `<gio/gunixfdmessage.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixFDMessage interface {
	SocketControlMessage

	// AppendFd adds a file descriptor to @message.
	//
	// The file descriptor is duplicated using dup(). You keep your copy of the
	// descriptor and the copy contained in @message will be closed when
	// @message is finalized.
	//
	// A possible cause of failure is exceeding the per-process or system-wide
	// file descriptor limit.
	AppendFd(fd int) error
	// FdList gets the FDList contained in @message. This function does not
	// return a reference to the caller, but the returned list is valid for the
	// lifetime of @message.
	FdList() UnixFDList
}

// unixFDMessage implements the UnixFDMessage class.
type unixFDMessage struct {
	SocketControlMessage
}

var _ UnixFDMessage = (*unixFDMessage)(nil)

// WrapUnixFDMessage wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixFDMessage(obj *externglib.Object) UnixFDMessage {
	return unixFDMessage{
		SocketControlMessage: WrapSocketControlMessage(obj),
	}
}

func marshalUnixFDMessage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixFDMessage(obj), nil
}

// NewUnixFDMessage constructs a class UnixFDMessage.
func NewUnixFDMessage() UnixFDMessage {
	var _cret C.GUnixFDMessage // in

	_cret = C.g_unix_fd_message_new()

	var _unixFDMessage UnixFDMessage // out

	_unixFDMessage = WrapUnixFDMessage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixFDMessage
}

// NewUnixFDMessageWithFdList constructs a class UnixFDMessage.
func NewUnixFDMessageWithFdList(fdList UnixFDList) UnixFDMessage {
	var _arg1 *C.GUnixFDList   // out
	var _cret C.GUnixFDMessage // in

	_arg1 = (*C.GUnixFDList)(unsafe.Pointer(fdList.Native()))

	_cret = C.g_unix_fd_message_new_with_fd_list(_arg1)

	var _unixFDMessage UnixFDMessage // out

	_unixFDMessage = WrapUnixFDMessage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixFDMessage
}

// AppendFd adds a file descriptor to @message.
//
// The file descriptor is duplicated using dup(). You keep your copy of the
// descriptor and the copy contained in @message will be closed when
// @message is finalized.
//
// A possible cause of failure is exceeding the per-process or system-wide
// file descriptor limit.
func (m unixFDMessage) AppendFd(fd int) error {
	var _arg0 *C.GUnixFDMessage // out
	var _arg1 C.gint            // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GUnixFDMessage)(unsafe.Pointer(m.Native()))
	_arg1 = (C.gint)(fd)

	C.g_unix_fd_message_append_fd(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// FdList gets the FDList contained in @message. This function does not
// return a reference to the caller, but the returned list is valid for the
// lifetime of @message.
func (m unixFDMessage) FdList() UnixFDList {
	var _arg0 *C.GUnixFDMessage // out
	var _cret *C.GUnixFDList    // in

	_arg0 = (*C.GUnixFDMessage)(unsafe.Pointer(m.Native()))

	_cret = C.g_unix_fd_message_get_fd_list(_arg0)

	var _unixFDList UnixFDList // out

	_unixFDList = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(UnixFDList)

	return _unixFDList
}
