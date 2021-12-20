// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_loadable_icon_get_type()), F: marshalLoadableIconner},
	})
}

// LoadableIconOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LoadableIconOverrider interface {
	// Load loads a loadable icon. For the asynchronous version of this
	// function, see g_loadable_icon_load_async().
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - size: integer.
	//
	// The function returns the following values:
	//
	//    - typ (optional): location to store the type of the loaded icon, NULL
	//      to ignore.
	//    - inputStream to read the icon from.
	//
	Load(ctx context.Context, size int) (string, InputStreamer, error)
	// LoadAsync loads an icon asynchronously. To finish this function, see
	// g_loadable_icon_load_finish(). For the synchronous, blocking version of
	// this function, see g_loadable_icon_load().
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - size: integer.
	//    - callback (optional) to call when the request is satisfied.
	//
	LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback)
	// LoadFinish finishes an asynchronous icon load started in
	// g_loadable_icon_load_async().
	//
	// The function takes the following parameters:
	//
	//    - res: Result.
	//
	// The function returns the following values:
	//
	//    - typ (optional): location to store the type of the loaded icon, NULL
	//      to ignore.
	//    - inputStream to read the icon from.
	//
	LoadFinish(res AsyncResulter) (string, InputStreamer, error)
}

// LoadableIcon extends the #GIcon interface and adds the ability to load icons
// from streams.
type LoadableIcon struct {
	Icon

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var ()

// LoadableIconner describes LoadableIcon's interface methods.
type LoadableIconner interface {
	externglib.Objector

	// Load loads a loadable icon.
	Load(ctx context.Context, size int) (string, InputStreamer, error)
	// LoadAsync loads an icon asynchronously.
	LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback)
	// LoadFinish finishes an asynchronous icon load started in
	// g_loadable_icon_load_async().
	LoadFinish(res AsyncResulter) (string, InputStreamer, error)
}

var _ LoadableIconner = (*LoadableIcon)(nil)

func wrapLoadableIcon(obj *externglib.Object) *LoadableIcon {
	return &LoadableIcon{
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalLoadableIconner(p uintptr) (interface{}, error) {
	return wrapLoadableIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Load loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - size: integer.
//
// The function returns the following values:
//
//    - typ (optional): location to store the type of the loaded icon, NULL to
//      ignore.
//    - inputStream to read the icon from.
//
func (icon *LoadableIcon) Load(ctx context.Context, size int) (string, InputStreamer, error) {
	var _arg0 *C.GLoadableIcon // out
	var _arg3 *C.GCancellable  // out
	var _arg1 C.int            // out
	var _arg2 *C.char          // in
	var _cret *C.GInputStream  // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(size)

	_cret = C.g_loadable_icon_load(_arg0, _arg1, &_arg2, _arg3, &_cerr)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(size)

	var _typ string                // out
	var _inputStream InputStreamer // out
	var _goerr error               // out

	if _arg2 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(InputStreamer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.InputStreamer")
		}
		_inputStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _typ, _inputStream, _goerr
}

// LoadAsync loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking version of this
// function, see g_loadable_icon_load().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - size: integer.
//    - callback (optional) to call when the request is satisfied.
//
func (icon *LoadableIcon) LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback) {
	var _arg0 *C.GLoadableIcon      // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(size)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_loadable_icon_load_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(size)
	runtime.KeepAlive(callback)
}

// LoadFinish finishes an asynchronous icon load started in
// g_loadable_icon_load_async().
//
// The function takes the following parameters:
//
//    - res: Result.
//
// The function returns the following values:
//
//    - typ (optional): location to store the type of the loaded icon, NULL to
//      ignore.
//    - inputStream to read the icon from.
//
func (icon *LoadableIcon) LoadFinish(res AsyncResulter) (string, InputStreamer, error) {
	var _arg0 *C.GLoadableIcon // out
	var _arg1 *C.GAsyncResult  // out
	var _arg2 *C.char          // in
	var _cret *C.GInputStream  // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_loadable_icon_load_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(res)

	var _typ string                // out
	var _inputStream InputStreamer // out
	var _goerr error               // out

	if _arg2 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.Cast()
		rv, ok := casted.(InputStreamer)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.InputStreamer")
		}
		_inputStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _typ, _inputStream, _goerr
}
