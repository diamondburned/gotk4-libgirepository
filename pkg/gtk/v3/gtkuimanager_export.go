// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_UIManager_ConnectActionsChanged
func _gotk4_gtk3_UIManager_ConnectActionsChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_UIManager_ConnectAddWidget
func _gotk4_gtk3_UIManager_ConnectAddWidget(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

//export _gotk4_gtk3_UIManager_ConnectConnectProxy
func _gotk4_gtk3_UIManager_ConnectConnectProxy(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(action *Action, proxy Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(action *Action, proxy Widgetter))
	}

	var _action *Action  // out
	var _proxy Widgetter // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_proxy = rv
	}

	f(_action, _proxy)
}

//export _gotk4_gtk3_UIManager_ConnectDisconnectProxy
func _gotk4_gtk3_UIManager_ConnectDisconnectProxy(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(action *Action, proxy Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(action *Action, proxy Widgetter))
	}

	var _action *Action  // out
	var _proxy Widgetter // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_proxy = rv
	}

	f(_action, _proxy)
}

//export _gotk4_gtk3_UIManager_ConnectPostActivate
func _gotk4_gtk3_UIManager_ConnectPostActivate(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(action *Action)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(action *Action))
	}

	var _action *Action // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))

	f(_action)
}

//export _gotk4_gtk3_UIManager_ConnectPreActivate
func _gotk4_gtk3_UIManager_ConnectPreActivate(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(action *Action)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(action *Action))
	}

	var _action *Action // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(arg1)))

	f(_action)
}
