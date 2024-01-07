// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_TranslateFunc
func _gotk4_gtk3_TranslateFunc(arg1 *C.gchar, arg2 C.gpointer) (cret *C.gchar) {
	var fn TranslateFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TranslateFunc)
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := fn(_path)

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}
