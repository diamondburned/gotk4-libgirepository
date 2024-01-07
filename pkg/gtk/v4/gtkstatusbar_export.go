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

//export _gotk4_gtk4_Statusbar_ConnectTextPopped
func _gotk4_gtk4_Statusbar_ConnectTextPopped(arg0 C.gpointer, arg1 C.guint, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(contextId uint, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(contextId uint, text string))
	}

	var _contextId uint // out
	var _text string    // out

	_contextId = uint(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_contextId, _text)
}

//export _gotk4_gtk4_Statusbar_ConnectTextPushed
func _gotk4_gtk4_Statusbar_ConnectTextPushed(arg0 C.gpointer, arg1 C.guint, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(contextId uint, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(contextId uint, text string))
	}

	var _contextId uint // out
	var _text string    // out

	_contextId = uint(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_contextId, _text)
}
