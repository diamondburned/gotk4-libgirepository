// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_TextCharPredicate
func _gotk4_gtk3_TextCharPredicate(arg1 C.gunichar, arg2 C.gpointer) (cret C.gboolean) {
	var fn TextCharPredicate
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TextCharPredicate)
	}

	var _ch uint32 // out

	_ch = uint32(arg1)

	ok := fn(_ch)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}