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

//export _gotk4_gtk3_GestureDrag_ConnectDragBegin
func _gotk4_gtk3_GestureDrag_ConnectDragBegin(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(startX, startY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(startX, startY float64))
	}

	var _startX float64 // out
	var _startY float64 // out

	_startX = float64(arg1)
	_startY = float64(arg2)

	f(_startX, _startY)
}

//export _gotk4_gtk3_GestureDrag_ConnectDragEnd
func _gotk4_gtk3_GestureDrag_ConnectDragEnd(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}

//export _gotk4_gtk3_GestureDrag_ConnectDragUpdate
func _gotk4_gtk3_GestureDrag_ConnectDragUpdate(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}
