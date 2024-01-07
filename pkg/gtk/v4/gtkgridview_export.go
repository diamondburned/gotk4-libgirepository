// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_GridView_ConnectActivate
func _gotk4_gtk4_GridView_ConnectActivate(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(position uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position uint))
	}

	var _position uint // out

	_position = uint(arg1)

	f(_position)
}
