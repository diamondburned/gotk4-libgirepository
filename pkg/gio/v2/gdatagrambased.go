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
		{T: externglib.Type(C.g_datagram_based_get_type()), F: marshalDatagramBased},
	})
}

// DatagramBasedOverrider contains methods that are overridable. This
// interface is a subset of the interface DatagramBased.
type DatagramBasedOverrider interface {
	// ConditionCheck checks on the readiness of @datagram_based to perform
	// operations. The operations specified in @condition are checked for and
	// masked against the currently-satisfied conditions on @datagram_based. The
	// result is returned.
	//
	// G_IO_IN will be set in the return value if data is available to read with
	// g_datagram_based_receive_messages(), or if the connection is closed
	// remotely (EOS); and if the datagram_based has not been closed locally
	// using some implementation-specific method (such as g_socket_close() or
	// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket).
	//
	// If the connection is shut down or closed (by calling g_socket_close() or
	// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket, for
	// example), all calls to this function will return G_IO_ERROR_CLOSED.
	//
	// G_IO_OUT will be set if it is expected that at least one byte can be sent
	// using g_datagram_based_send_messages() without blocking. It will not be
	// set if the datagram_based has been closed locally.
	//
	// G_IO_HUP will be set if the connection has been closed locally.
	//
	// G_IO_ERR will be set if there was an asynchronous error in transmitting
	// data previously enqueued using g_datagram_based_send_messages().
	//
	// Note that on Windows, it is possible for an operation to return
	// G_IO_ERROR_WOULD_BLOCK even immediately after
	// g_datagram_based_condition_check() has claimed that the Based is ready
	// for writing. Rather than calling g_datagram_based_condition_check() and
	// then writing to the Based if it succeeds, it is generally better to
	// simply try writing right away, and try again later if the initial attempt
	// returns G_IO_ERROR_WOULD_BLOCK.
	//
	// It is meaningless to specify G_IO_ERR or G_IO_HUP in @condition; these
	// conditions will always be set in the output if they are true. Apart from
	// these flags, the output is guaranteed to be masked by @condition.
	//
	// This call never blocks.
	ConditionCheck(condition glib.IOCondition) glib.IOCondition
	// ConditionWait waits for up to @timeout microseconds for condition to
	// become true on @datagram_based. If the condition is met, true is
	// returned.
	//
	// If @cancellable is cancelled before the condition is met, or if @timeout
	// is reached before the condition is met, then false is returned and @error
	// is set appropriately (G_IO_ERROR_CANCELLED or G_IO_ERROR_TIMED_OUT).
	ConditionWait(condition glib.IOCondition, timeout int64, cancellable Cancellable) error
	// CreateSource creates a #GSource that can be attached to a Context to
	// monitor for the availability of the specified @condition on the Based.
	// The #GSource keeps a reference to the @datagram_based.
	//
	// The callback on the source is of the BasedSourceFunc type.
	//
	// It is meaningless to specify G_IO_ERR or G_IO_HUP in @condition; these
	// conditions will always be reported in the callback if they are true.
	//
	// If non-nil, @cancellable can be used to cancel the source, which will
	// cause the source to trigger, reporting the current condition (which is
	// likely 0 unless cancellation happened at the same time as a condition
	// change). You can check for this in the callback using
	// g_cancellable_is_cancelled().
	CreateSource(condition glib.IOCondition, cancellable Cancellable) *glib.Source
}

// DatagramBased: a Based is a networking interface for representing
// datagram-based communications. It is a more or less direct mapping of the
// core parts of the BSD socket API in a portable GObject interface. It is
// implemented by #GSocket, which wraps the UNIX socket API on UNIX and winsock2
// on Windows.
//
// Based is entirely platform independent, and is intended to be used alongside
// higher-level networking APIs such as OStream.
//
// It uses vectored scatter/gather I/O by default, allowing for many messages to
// be sent or received in a single call. Where possible, implementations of the
// interface should take advantage of vectored I/O to minimise processing or
// system calls. For example, #GSocket uses recvmmsg() and sendmmsg() where
// possible. Callers should take advantage of scatter/gather I/O (the use of
// multiple buffers per message) to avoid unnecessary copying of data to
// assemble or disassemble a message.
//
// Each Based operation has a timeout parameter which may be negative for
// blocking behaviour, zero for non-blocking behaviour, or positive for timeout
// behaviour. A blocking operation blocks until finished or there is an error. A
// non-blocking operation will return immediately with a G_IO_ERROR_WOULD_BLOCK
// error if it cannot make progress. A timeout operation will block until the
// operation is complete or the timeout expires; if the timeout expires it will
// return what progress it made, or G_IO_ERROR_TIMED_OUT if no progress was
// made. To know when a call would successfully run you can call
// g_datagram_based_condition_check() or g_datagram_based_condition_wait(). You
// can also use g_datagram_based_create_source() and attach it to a Context to
// get callbacks when I/O is possible.
//
// When running a non-blocking operation applications should always be able to
// handle getting a G_IO_ERROR_WOULD_BLOCK error even when some other function
// said that I/O was possible. This can easily happen in case of a race
// condition in the application, but it can also happen for other reasons. For
// instance, on Windows a socket is always seen as writable until a write
// returns G_IO_ERROR_WOULD_BLOCK.
//
// As with #GSocket, Baseds can be either connection oriented (for example,
// SCTP) or connectionless (for example, UDP). Baseds must be datagram-based,
// not stream-based. The interface does not cover connection establishment — use
// methods on the underlying type to establish a connection before sending and
// receiving data through the Based API. For connectionless socket types the
// target/source address is specified or received in each I/O operation.
//
// Like most other APIs in GLib, Based is not inherently thread safe. To use a
// Based concurrently from multiple threads, you must implement your own
// locking.
type DatagramBased interface {
	gextras.Objector
	DatagramBasedOverrider
}

// datagramBased implements the DatagramBased interface.
type datagramBased struct {
	gextras.Objector
}

var _ DatagramBased = (*datagramBased)(nil)

// WrapDatagramBased wraps a GObject to a type that implements interface
// DatagramBased. It is primarily used internally.
func WrapDatagramBased(obj *externglib.Object) DatagramBased {
	return DatagramBased{
		Objector: obj,
	}
}

func marshalDatagramBased(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDatagramBased(obj), nil
}

// ConditionCheck checks on the readiness of @datagram_based to perform
// operations. The operations specified in @condition are checked for and
// masked against the currently-satisfied conditions on @datagram_based. The
// result is returned.
//
// G_IO_IN will be set in the return value if data is available to read with
// g_datagram_based_receive_messages(), or if the connection is closed
// remotely (EOS); and if the datagram_based has not been closed locally
// using some implementation-specific method (such as g_socket_close() or
// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket, for
// example), all calls to this function will return G_IO_ERROR_CLOSED.
//
// G_IO_OUT will be set if it is expected that at least one byte can be sent
// using g_datagram_based_send_messages() without blocking. It will not be
// set if the datagram_based has been closed locally.
//
// G_IO_HUP will be set if the connection has been closed locally.
//
// G_IO_ERR will be set if there was an asynchronous error in transmitting
// data previously enqueued using g_datagram_based_send_messages().
//
// Note that on Windows, it is possible for an operation to return
// G_IO_ERROR_WOULD_BLOCK even immediately after
// g_datagram_based_condition_check() has claimed that the Based is ready
// for writing. Rather than calling g_datagram_based_condition_check() and
// then writing to the Based if it succeeds, it is generally better to
// simply try writing right away, and try again later if the initial attempt
// returns G_IO_ERROR_WOULD_BLOCK.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in @condition; these
// conditions will always be set in the output if they are true. Apart from
// these flags, the output is guaranteed to be masked by @condition.
//
// This call never blocks.
func (d datagramBased) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	var arg0 *C.GDatagramBased
	var arg1 C.GIOCondition

	arg0 = (*C.GDatagramBased)(unsafe.Pointer(d.Native()))
	arg1 = (C.GIOCondition)(condition)

	var cret C.GIOCondition
	var goret glib.IOCondition

	cret = C.g_datagram_based_condition_check(arg0, arg1)

	goret = glib.IOCondition(cret)

	return goret
}

// ConditionWait waits for up to @timeout microseconds for condition to
// become true on @datagram_based. If the condition is met, true is
// returned.
//
// If @cancellable is cancelled before the condition is met, or if @timeout
// is reached before the condition is met, then false is returned and @error
// is set appropriately (G_IO_ERROR_CANCELLED or G_IO_ERROR_TIMED_OUT).
func (d datagramBased) ConditionWait(condition glib.IOCondition, timeout int64, cancellable Cancellable) error {
	var arg0 *C.GDatagramBased
	var arg1 C.GIOCondition
	var arg2 C.gint64
	var arg3 *C.GCancellable

	arg0 = (*C.GDatagramBased)(unsafe.Pointer(d.Native()))
	arg1 = (C.GIOCondition)(condition)
	arg2 = C.gint64(timeout)
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError
	var goerr error

	C.g_datagram_based_condition_wait(arg0, arg1, arg2, arg3, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// CreateSource creates a #GSource that can be attached to a Context to
// monitor for the availability of the specified @condition on the Based.
// The #GSource keeps a reference to the @datagram_based.
//
// The callback on the source is of the BasedSourceFunc type.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in @condition; these
// conditions will always be reported in the callback if they are true.
//
// If non-nil, @cancellable can be used to cancel the source, which will
// cause the source to trigger, reporting the current condition (which is
// likely 0 unless cancellation happened at the same time as a condition
// change). You can check for this in the callback using
// g_cancellable_is_cancelled().
func (d datagramBased) CreateSource(condition glib.IOCondition, cancellable Cancellable) *glib.Source {
	var arg0 *C.GDatagramBased
	var arg1 C.GIOCondition
	var arg2 *C.GCancellable

	arg0 = (*C.GDatagramBased)(unsafe.Pointer(d.Native()))
	arg1 = (C.GIOCondition)(condition)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	cret := new(C.GSource)
	var goret *glib.Source

	cret = C.g_datagram_based_create_source(arg0, arg1, arg2)

	goret = glib.WrapSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.Source) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}
