// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_cancellable_get_type()), F: marshalCancellabler},
	})
}

// CancellableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CancellableOverrider interface {
	Cancelled()
}

// Cancellable is a thread-safe operation cancellation stack used throughout GIO
// to allow for cancellation of synchronous and asynchronous operations.
type Cancellable struct {
	*externglib.Object
}

func wrapCancellable(obj *externglib.Object) *Cancellable {
	return &Cancellable{
		Object: obj,
	}
}

func marshalCancellabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCancellable(obj), nil
}

// NewCancellable creates a new #GCancellable object.
//
// Applications that want to start one or more operations that should be
// cancellable should create a #GCancellable and pass it to the operations.
//
// One #GCancellable can be used in multiple consecutive operations or in
// multiple concurrent operations.
func NewCancellable() *Cancellable {
	var _cret *C.GCancellable // in

	_cret = C.g_cancellable_new()

	var _cancellable *Cancellable // out

	_cancellable = wrapCancellable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cancellable
}

// Cancel will set cancellable to cancelled, and will emit the
// #GCancellable::cancelled signal. (However, see the warning about race
// conditions in the documentation for that signal if you are planning to
// connect to it.)
//
// This function is thread-safe. In other words, you can safely call it from a
// thread other than the one running the operation that was passed the
// cancellable.
//
// If cancellable is NULL, this function returns immediately for convenience.
//
// The convention within GIO is that cancelling an asynchronous operation causes
// it to complete asynchronously. That is, if you cancel the operation from the
// same thread in which it is running, then the operation's ReadyCallback will
// not be invoked until the application returns to the main loop.
func (cancellable *Cancellable) Cancel() {
	var _arg0 *C.GCancellable // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_cancel(_arg0)
}

// Disconnect disconnects a handler from a cancellable instance similar to
// g_signal_handler_disconnect(). Additionally, in the event that a signal
// handler is currently running, this call will block until the handler has
// finished. Calling this function from a #GCancellable::cancelled signal
// handler will therefore result in a deadlock.
//
// This avoids a race condition where a thread cancels at the same time as the
// cancellable operation is finished and the signal handler is removed. See
// #GCancellable::cancelled for details on how to use this.
//
// If cancellable is NULL or handler_id is 0 this function does nothing.
func (cancellable *Cancellable) Disconnect(handlerId uint32) {
	var _arg0 *C.GCancellable // out
	var _arg1 C.gulong        // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg1 = C.gulong(handlerId)

	C.g_cancellable_disconnect(_arg0, _arg1)
}

// Fd gets the file descriptor for a cancellable job. This can be used to
// implement cancellable operations on Unix systems. The returned fd will turn
// readable when cancellable is cancelled.
//
// You are not supposed to read from the fd yourself, just check for readable
// status. Reading to unset the readable status is done with
// g_cancellable_reset().
//
// After a successful return from this function, you should use
// g_cancellable_release_fd() to free up resources allocated for the returned
// file descriptor.
//
// See also g_cancellable_make_pollfd().
func (cancellable *Cancellable) Fd() int {
	var _arg0 *C.GCancellable // out
	var _cret C.int           // in

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_cancellable_get_fd(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsCancelled checks if a cancellable job has been cancelled.
func (cancellable *Cancellable) IsCancelled() bool {
	var _arg0 *C.GCancellable // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_cancellable_is_cancelled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakePollfd creates a FD corresponding to cancellable; this can be passed to
// g_poll() and used to poll for cancellation. This is useful both for unix
// systems without a native poll and for portability to windows.
//
// When this function returns TRUE, you should use g_cancellable_release_fd() to
// free up resources allocated for the pollfd. After a FALSE return, do not call
// g_cancellable_release_fd().
//
// If this function returns FALSE, either no cancellable was given or resource
// limits prevent this function from allocating the necessary structures for
// polling. (On Linux, you will likely have reached the maximum number of file
// descriptors.) The suggested way to handle these cases is to ignore the
// cancellable.
//
// You are not supposed to read from the fd yourself, just check for readable
// status. Reading to unset the readable status is done with
// g_cancellable_reset().
func (cancellable *Cancellable) MakePollfd(pollfd *glib.PollFD) bool {
	var _arg0 *C.GCancellable // out
	var _arg1 *C.GPollFD      // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg1 = (*C.GPollFD)(gextras.StructNative(unsafe.Pointer(pollfd)))

	_cret = C.g_cancellable_make_pollfd(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopCurrent pops cancellable off the cancellable stack (verifying that
// cancellable is on the top of the stack).
func (cancellable *Cancellable) PopCurrent() {
	var _arg0 *C.GCancellable // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_pop_current(_arg0)
}

// PushCurrent pushes cancellable onto the cancellable stack. The current
// cancellable can then be received using g_cancellable_get_current().
//
// This is useful when implementing cancellable operations in code that does not
// allow you to pass down the cancellable object.
//
// This is typically called automatically by e.g. #GFile operations, so you
// rarely have to call this yourself.
func (cancellable *Cancellable) PushCurrent() {
	var _arg0 *C.GCancellable // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_push_current(_arg0)
}

// ReleaseFd releases a resources previously allocated by g_cancellable_get_fd()
// or g_cancellable_make_pollfd().
//
// For compatibility reasons with older releases, calling this function is not
// strictly required, the resources will be automatically freed when the
// cancellable is finalized. However, the cancellable will block scarce file
// descriptors until it is finalized if this function is not called. This can
// cause the application to run out of file descriptors when many #GCancellables
// are used at the same time.
func (cancellable *Cancellable) ReleaseFd() {
	var _arg0 *C.GCancellable // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_release_fd(_arg0)
}

// Reset resets cancellable to its uncancelled state.
//
// If cancellable is currently in use by any cancellable operation then the
// behavior of this function is undefined.
//
// Note that it is generally not a good idea to reuse an existing cancellable
// for more operations after it has been cancelled once, as this function might
// tempt you to do. The recommended practice is to drop the reference to a
// cancellable after cancelling it, and let it die with the outstanding async
// operations. You should create a fresh cancellable for further async
// operations.
func (cancellable *Cancellable) Reset() {
	var _arg0 *C.GCancellable // out

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_reset(_arg0)
}

// SetErrorIfCancelled: if the cancellable is cancelled, sets the error to
// notify that the operation was cancelled.
func (cancellable *Cancellable) SetErrorIfCancelled() error {
	var _arg0 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_cancellable_set_error_if_cancelled(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// NewSource creates a source that triggers if cancellable is cancelled and
// calls its callback of type SourceFunc. This is primarily useful for attaching
// to another (non-cancellable) source with g_source_add_child_source() to add
// cancellability to it.
//
// For convenience, you can call this with a NULL #GCancellable, in which case
// the source will never trigger.
//
// The new #GSource will hold a reference to the #GCancellable.
func (cancellable *Cancellable) NewSource() *glib.Source {
	var _arg0 *C.GCancellable // out
	var _cret *C.GSource      // in

	_arg0 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_cancellable_source_new(_arg0)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_source_ref(_cret)
	runtime.SetFinalizer(_source, func(v *glib.Source) {
		C.g_source_unref((*C.GSource)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _source
}

// CancellableGetCurrent gets the top cancellable from the stack.
func CancellableGetCurrent() *Cancellable {
	var _cret *C.GCancellable // in

	_cret = C.g_cancellable_get_current()

	var _cancellable *Cancellable // out

	_cancellable = wrapCancellable(externglib.Take(unsafe.Pointer(_cret)))

	return _cancellable
}
