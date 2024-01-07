// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

//export _gotk4_glib2_HRFunc
func _gotk4_glib2_HRFunc(arg1 C.gpointer, arg2 C.gpointer, arg3 C.gpointer) (cret C.gboolean) {
	var fn HRFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(HRFunc)
	}

	var _key unsafe.Pointer   // out
	var _value unsafe.Pointer // out

	_key = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_value = (unsafe.Pointer)(unsafe.Pointer(arg2))

	ok := fn(_key, _value)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}