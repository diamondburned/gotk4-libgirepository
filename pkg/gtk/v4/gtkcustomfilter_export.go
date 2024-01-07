// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_CustomFilterFunc
func _gotk4_gtk4_CustomFilterFunc(arg1 C.gpointer, arg2 C.gpointer) (cret C.gboolean) {
	var fn CustomFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CustomFilterFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	ok := fn(_item)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
