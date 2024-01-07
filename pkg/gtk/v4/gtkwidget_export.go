// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_TickCallback
func _gotk4_gtk4_TickCallback(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gboolean) {
	var fn TickCallback
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TickCallback)
	}

	var _widget Widgetter            // out
	var _frameClock gdk.FrameClocker // out

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
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gdk.FrameClocker is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.FrameClocker)
			return ok
		})
		rv, ok := casted.(gdk.FrameClocker)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.FrameClocker")
		}
		_frameClock = rv
	}

	ok := fn(_widget, _frameClock)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_Widget_ConnectDestroy
func _gotk4_gtk4_Widget_ConnectDestroy(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectHide
func _gotk4_gtk4_Widget_ConnectHide(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectMap
func _gotk4_gtk4_Widget_ConnectMap(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectMnemonicActivate
func _gotk4_gtk4_Widget_ConnectMnemonicActivate(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) (cret C.gboolean) {
	var f func(groupCycling bool) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(groupCycling bool) (ok bool))
	}

	var _groupCycling bool // out

	if arg1 != 0 {
		_groupCycling = true
	}

	ok := f(_groupCycling)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_Widget_ConnectQueryTooltip
func _gotk4_gtk4_Widget_ConnectQueryTooltip(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.gboolean, arg4 *C.void, arg5 C.guintptr) (cret C.gboolean) {
	var f func(x, y int, keyboardMode bool, tooltip *Tooltip) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg5))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y int, keyboardMode bool, tooltip *Tooltip) (ok bool))
	}

	var _x int             // out
	var _y int             // out
	var _keyboardMode bool // out
	var _tooltip *Tooltip  // out

	_x = int(arg1)
	_y = int(arg2)
	if arg3 != 0 {
		_keyboardMode = true
	}
	_tooltip = wrapTooltip(coreglib.Take(unsafe.Pointer(arg4)))

	ok := f(_x, _y, _keyboardMode, _tooltip)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_Widget_ConnectRealize
func _gotk4_gtk4_Widget_ConnectRealize(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectShow
func _gotk4_gtk4_Widget_ConnectShow(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectUnmap
func _gotk4_gtk4_Widget_ConnectUnmap(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Widget_ConnectUnrealize
func _gotk4_gtk4_Widget_ConnectUnrealize(arg0 C.gpointer, arg1 C.guintptr) {
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