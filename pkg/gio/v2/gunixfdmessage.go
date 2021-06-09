// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
	// StealFds returns the array of file descriptors that is contained in this
	// object.
	//
	// After this call, the descriptors are no longer contained in @message.
	// Further calls will return an empty list (unless more descriptors have
	// been added).
	//
	// The return result of this function must be freed with g_free(). The
	// caller is also responsible for closing all of the file descriptors.
	//
	// If @length is non-nil then it is set to the number of file descriptors in
	// the returned array. The returned array is also terminated with -1.
	//
	// This function never returns nil. In case there are no file descriptors
	// contained in @message, an empty array is returned.
	StealFds() []int
}

// unixFDMessage implements the UnixFDMessage interface.
type unixFDMessage struct {
	SocketControlMessage
}

var _ UnixFDMessage = (*unixFDMessage)(nil)

// WrapUnixFDMessage wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixFDMessage(obj *externglib.Object) UnixFDMessage {
	return UnixFDMessage{
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
	var _cret C.GUnixFDMessage

	cret = C.g_unix_fd_message_new()

	var _unixFDMessage UnixFDMessage

	_unixFDMessage = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(UnixFDMessage)

	return _unixFDMessage
}

// NewUnixFDMessageWithFdList constructs a class UnixFDMessage.
func NewUnixFDMessageWithFdList(fdList UnixFDList) UnixFDMessage {
	var _arg1 *C.GUnixFDList

	_arg1 = (*C.GUnixFDList)(unsafe.Pointer(fdList.Native()))

	var _cret C.GUnixFDMessage

	cret = C.g_unix_fd_message_new_with_fd_list(_arg1)

	var _unixFDMessage UnixFDMessage

	_unixFDMessage = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(UnixFDMessage)

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
	var _arg0 *C.GUnixFDMessage
	var _arg1 C.gint

	_arg0 = (*C.GUnixFDMessage)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(fd)

	var _cerr *C.GError

	C.g_unix_fd_message_append_fd(_arg0, _arg1, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// FdList gets the FDList contained in @message. This function does not
// return a reference to the caller, but the returned list is valid for the
// lifetime of @message.
func (m unixFDMessage) FdList() UnixFDList {
	var _arg0 *C.GUnixFDMessage

	_arg0 = (*C.GUnixFDMessage)(unsafe.Pointer(m.Native()))

	var _cret *C.GUnixFDList

	cret = C.g_unix_fd_message_get_fd_list(_arg0)

	var _unixFDList UnixFDList

	_unixFDList = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(UnixFDList)

	return _unixFDList
}

// StealFds returns the array of file descriptors that is contained in this
// object.
//
// After this call, the descriptors are no longer contained in @message.
// Further calls will return an empty list (unless more descriptors have
// been added).
//
// The return result of this function must be freed with g_free(). The
// caller is also responsible for closing all of the file descriptors.
//
// If @length is non-nil then it is set to the number of file descriptors in
// the returned array. The returned array is also terminated with -1.
//
// This function never returns nil. In case there are no file descriptors
// contained in @message, an empty array is returned.
func (m unixFDMessage) StealFds() []int {
	var _arg0 *C.GUnixFDMessage

	_arg0 = (*C.GUnixFDMessage)(unsafe.Pointer(m.Native()))

	var _cret *C.gint
	var _arg1 *C.gint

	cret = C.g_unix_fd_message_steal_fds(_arg0)

	var _gints []int

	ptr.SetSlice(unsafe.Pointer(&_gints), unsafe.Pointer(_cret), int(_arg1))
	runtime.SetFinalizer(&_gints, func(v *[]int) {
		C.free(ptr.Slice(unsafe.Pointer(v)))
	})

	return _gints
}