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
		{T: externglib.Type(C.g_async_result_get_type()), F: marshalAsyncResult},
	})
}

// AsyncResultOverrider contains methods that are overridable. This
// interface is a subset of the interface AsyncResult.
type AsyncResultOverrider interface {
	// SourceObject gets the source object from a Result.
	SourceObject() gextras.Objector
}

// AsyncResult provides a base class for implementing asynchronous function
// results.
//
// Asynchronous operations are broken up into two separate operations which are
// chained together by a ReadyCallback. To begin an asynchronous operation,
// provide a ReadyCallback to the asynchronous function. This callback will be
// triggered when the operation has completed, and must be run in a later
// iteration of the [thread-default main
// context][g-main-context-push-thread-default] from where the operation was
// initiated. It will be passed a Result instance filled with the details of the
// operation's success or failure, the object the asynchronous function was
// started for and any error codes returned. The asynchronous callback function
// is then expected to call the corresponding "_finish()" function, passing the
// object the function was called for, the Result instance, and (optionally) an
// @error to grab any error conditions that may have occurred.
//
// The "_finish()" function for an operation takes the generic result (of type
// Result) and returns the specific result that the operation in question yields
// (e.g. a Enumerator for a "enumerate children" operation). If the result or
// error status of the operation is not needed, there is no need to call the
// "_finish()" function; GIO will take care of cleaning up the result and error
// information after the ReadyCallback returns. You can pass nil for the
// ReadyCallback if you don't need to take any action at all after the operation
// completes. Applications may also take a reference to the Result and call
// "_finish()" later; however, the "_finish()" function may be called at most
// once.
//
// Example of a typical asynchronous operation flow:
//
//    void _theoretical_frobnitz_async (Theoretical         *t,
//                                      GCancellable        *c,
//                                      GAsyncReadyCallback  cb,
//                                      gpointer             u);
//
//    gboolean _theoretical_frobnitz_finish (Theoretical   *t,
//                                           GAsyncResult  *res,
//                                           GError       **e);
//
//    static void
//    frobnitz_result_func (GObject      *source_object,
//    		 GAsyncResult *res,
//    		 gpointer      user_data)
//    {
//      gboolean success = FALSE;
//
//      success = _theoretical_frobnitz_finish (source_object, res, NULL);
//
//      if (success)
//        g_printf ("Hurray!\n");
//      else
//        g_printf ("Uh oh!\n");
//
//      ...
//
//    }
//
//    int main (int argc, void *argv[])
//    {
//       ...
//
//       _theoretical_frobnitz_async (theoretical_data,
//                                    NULL,
//                                    frobnitz_result_func,
//                                    NULL);
//
//       ...
//    }
//
// The callback for an asynchronous operation is called only once, and is always
// called, even in the case of a cancelled operation. On cancellation the result
// is a G_IO_ERROR_CANCELLED error.
//
// I/O Priority
//
// Many I/O-related asynchronous operations have a priority parameter, which is
// used in certain cases to determine the order in which operations are
// executed. They are not used to determine system-wide I/O scheduling.
// Priorities are integers, with lower numbers indicating higher priority. It is
// recommended to choose priorities between G_PRIORITY_LOW and G_PRIORITY_HIGH,
// with G_PRIORITY_DEFAULT as a default.
type AsyncResult interface {
	gextras.Objector
	AsyncResultOverrider

	// LegacyPropagateError: if @res is a AsyncResult, this is equivalent to
	// g_simple_async_result_propagate_error(). Otherwise it returns false.
	//
	// This can be used for legacy error handling in async *_finish() wrapper
	// functions that traditionally handled AsyncResult error returns themselves
	// rather than calling into the virtual method. This should not be used in
	// new code; Result errors that are set by virtual methods should also be
	// extracted by virtual methods, to enable subclasses to chain up correctly.
	LegacyPropagateError() error
}

// asyncResult implements the AsyncResult interface.
type asyncResult struct {
	gextras.Objector
}

var _ AsyncResult = (*asyncResult)(nil)

// WrapAsyncResult wraps a GObject to a type that implements interface
// AsyncResult. It is primarily used internally.
func WrapAsyncResult(obj *externglib.Object) AsyncResult {
	return AsyncResult{
		Objector: obj,
	}
}

func marshalAsyncResult(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAsyncResult(obj), nil
}

// SourceObject gets the source object from a Result.
func (r asyncResult) SourceObject() gextras.Objector {
	var _arg0 *C.GAsyncResult // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(r.Native()))

	_cret = C.g_async_result_get_source_object(_arg0)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

// LegacyPropagateError: if @res is a AsyncResult, this is equivalent to
// g_simple_async_result_propagate_error(). Otherwise it returns false.
//
// This can be used for legacy error handling in async *_finish() wrapper
// functions that traditionally handled AsyncResult error returns themselves
// rather than calling into the virtual method. This should not be used in
// new code; Result errors that are set by virtual methods should also be
// extracted by virtual methods, to enable subclasses to chain up correctly.
func (r asyncResult) LegacyPropagateError() error {
	var _arg0 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(r.Native()))

	C.g_async_result_legacy_propagate_error(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
