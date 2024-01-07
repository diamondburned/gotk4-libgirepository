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

//export _gotk4_gtk4_GestureRotate_ConnectAngleChanged
func _gotk4_gtk4_GestureRotate_ConnectAngleChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(angle, angleDelta float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(angle, angleDelta float64))
	}

	var _angle float64      // out
	var _angleDelta float64 // out

	_angle = float64(arg1)
	_angleDelta = float64(arg2)

	f(_angle, _angleDelta)
}
