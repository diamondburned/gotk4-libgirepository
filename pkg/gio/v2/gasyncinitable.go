// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_AsyncInitableIface_init_finish(GAsyncInitable*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// GType values.
var (
	GTypeAsyncInitable = coreglib.Type(C.g_async_initable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAsyncInitable, F: marshalAsyncInitable},
	})
}

// AsyncInitable: this is the asynchronous version of #GInitable; it behaves the
// same in all ways except that initialization is asynchronous. For more details
// see the descriptions on #GInitable.
//
// A class may implement both the #GInitable and Initable interfaces.
//
// Users of objects implementing this are not intended to use the interface
// method directly; instead it will be used automatically in various ways. For C
// applications you generally just call g_async_initable_new_async() directly,
// or indirectly via a foo_thing_new_async() wrapper. This will call
// g_async_initable_init_async() under the cover, calling back with NULL and a
// set GError on failure.
//
// A typical implementation might look something like this:
//
//    enum {
//       NOT_INITIALIZED,
//       INITIALIZING,
//       INITIALIZED
//    };
//
//    static void
//    _foo_ready_cb (Foo *self)
//    {
//      GList *l;
//
//      self->priv->state = INITIALIZED;
//
//      for (l = self->priv->init_results; l != NULL; l = l->next)
//        {
//          GTask *task = l->data;
//
//          if (self->priv->success)
//            g_task_return_boolean (task, TRUE);
//          else
//            g_task_return_new_error (task, ...);
//          g_object_unref (task);
//        }
//
//      g_list_free (self->priv->init_results);
//      self->priv->init_results = NULL;
//    }
//
//    static void
//    foo_init_async (GAsyncInitable       *initable,
//                    int                   io_priority,
//                    GCancellable         *cancellable,
//                    GAsyncReadyCallback   callback,
//                    gpointer              user_data)
//    {
//      Foo *self = FOO (initable);
//      GTask *task;
//
//      task = g_task_new (initable, cancellable, callback, user_data);
//      g_task_set_name (task, G_STRFUNC);
//
//      switch (self->priv->state)
//        {
//          case NOT_INITIALIZED:
//            _foo_get_ready (self);
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            self->priv->state = INITIALIZING;
//            break;
//          case INITIALIZING:
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            break;
//          case INITIALIZED:
//            if (!self->priv->success)
//              g_task_return_new_error (task, ...);
//            else
//              g_task_return_boolean (task, TRUE);
//            g_object_unref (task);
//            break;
//        }
//    }
//
//    static gboolean
//    foo_init_finish (GAsyncInitable       *initable,
//                     GAsyncResult         *result,
//                     GError              **error)
//    {
//      g_return_val_if_fail (g_task_is_valid (result, initable), FALSE);
//
//      return g_task_propagate_boolean (G_TASK (result), error);
//    }
//
//    static void
//    foo_async_initable_iface_init (gpointer g_iface,
//                                   gpointer data)
//    {
//      GAsyncInitableIface *iface = g_iface;
//
//      iface->init_async = foo_init_async;
//      iface->init_finish = foo_init_finish;
//    }.
//
// AsyncInitable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type AsyncInitable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AsyncInitable)(nil)
)

// AsyncInitabler describes AsyncInitable's interface methods.
type AsyncInitabler interface {
	coreglib.Objector

	// InitAsync starts asynchronous initialization of the object implementing
	// the interface.
	InitAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback)
	// InitFinish finishes asynchronous initialization and returns the result.
	InitFinish(res AsyncResulter) error
	// NewFinish finishes the async construction for the various
	// g_async_initable_new calls, returning the created object or NULL on
	// error.
	NewFinish(res AsyncResulter) (*coreglib.Object, error)
}

var _ AsyncInitabler = (*AsyncInitable)(nil)

func wrapAsyncInitable(obj *coreglib.Object) *AsyncInitable {
	return &AsyncInitable{
		Object: obj,
	}
}

func marshalAsyncInitable(p uintptr) (interface{}, error) {
	return wrapAsyncInitable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// InitAsync starts asynchronous initialization of the object implementing the
// interface. This must be done before any real use of the object after initial
// construction. If the object also implements #GInitable you can optionally
// call g_initable_init() instead.
//
// This method is intended for language bindings. If writing in C,
// g_async_initable_new_async() should typically be used instead.
//
// When the initialization is finished, callback will be called. You can then
// call g_async_initable_init_finish() to get the result of the initialization.
//
// Implementations may also support cancellation. If cancellable is not NULL,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// G_IO_ERROR_CANCELLED will be returned. If cancellable is not NULL, and the
// object doesn't support cancellable initialization, the error
// G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// As with #GInitable, if the object is not initialized, or initialization
// returns with an error, then all operations on the object except
// g_object_ref() and g_object_unref() are considered to be invalid, and have
// undefined behaviour. They will often fail with g_critical() or g_warning(),
// but this must not be relied on.
//
// Callers should not assume that a class which implements Initable can be
// initialized multiple times; for more information, see g_initable_init(). If a
// class explicitly supports being initialized multiple times, implementation
// requires yielding all subsequent calls to init_async() on the results of the
// first call.
//
// For classes that also support the #GInitable interface, the default
// implementation of this method will run the g_initable_init() function in a
// thread, so if you want to support asynchronous initialization via threads,
// just implement the Initable interface without overriding any interface
// methods.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: [I/O priority][io-priority] of the operation.
//    - callback (optional) to call when the request is satisfied.
//
func (initable *AsyncInitable) InitAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GAsyncInitable     // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(ioPriority)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_async_initable_init_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(initable)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// InitFinish finishes asynchronous initialization and returns the result. See
// g_async_initable_init_async().
//
// The function takes the following parameters:
//
//    - res: Result.
//
func (initable *AsyncInitable) InitFinish(res AsyncResulter) error {
	var _arg0 *C.GAsyncInitable // out
	var _arg1 *C.GAsyncResult   // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	C.g_async_initable_init_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(initable)
	runtime.KeepAlive(res)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// NewFinish finishes the async construction for the various
// g_async_initable_new calls, returning the created object or NULL on error.
//
// The function takes the following parameters:
//
//    - res from the callback.
//
// The function returns the following values:
//
//    - object: newly created #GObject, or NULL on error. Free with
//      g_object_unref().
//
func (initable *AsyncInitable) NewFinish(res AsyncResulter) (*coreglib.Object, error) {
	var _arg0 *C.GAsyncInitable // out
	var _arg1 *C.GAsyncResult   // out
	var _cret *C.GObject        // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_async_initable_new_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(initable)
	runtime.KeepAlive(res)

	var _object *coreglib.Object // out
	var _goerr error             // out

	_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _object, _goerr
}

// AsyncInitableIface provides an interface for asynchronous initializing object
// such that initialization may fail.
//
// An instance of this type is always passed by reference.
type AsyncInitableIface struct {
	*asyncInitableIface
}

// asyncInitableIface is the struct that's finalized.
type asyncInitableIface struct {
	native *C.GAsyncInitableIface
}
