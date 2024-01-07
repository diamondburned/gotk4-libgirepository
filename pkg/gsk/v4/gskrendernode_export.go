// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gsk4_ParseErrorFunc
func _gotk4_gsk4_ParseErrorFunc(arg1 *C.void, arg2 *C.void, arg3 *C.GError, arg4 C.gpointer) {
	var fn ParseErrorFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ParseErrorFunc)
	}

	var _start *ParseLocation // out
	var _end *ParseLocation   // out
	var _err error            // out

	_start = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_end = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_err = gerror.Take(unsafe.Pointer(arg3))

	fn(_start, _end, _err)
}