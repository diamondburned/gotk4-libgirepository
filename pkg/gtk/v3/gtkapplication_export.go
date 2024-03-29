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

//export _gotk4_gtk3_Application_ConnectQueryEnd
func _gotk4_gtk3_Application_ConnectQueryEnd(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_Application_ConnectWindowAdded
func _gotk4_gtk3_Application_ConnectWindowAdded(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(window *Window)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(window *Window))
	}

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	f(_window)
}

//export _gotk4_gtk3_Application_ConnectWindowRemoved
func _gotk4_gtk3_Application_ConnectWindowRemoved(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(window *Window)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(window *Window))
	}

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	f(_window)
}
