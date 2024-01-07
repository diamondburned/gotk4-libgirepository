// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_FontFilterFunc
func _gotk4_gtk4_FontFilterFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gboolean) {
	var fn FontFilterFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FontFilterFunc)
	}

	var _family pango.FontFamilier // out
	var _face pango.FontFacer      // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFamilier)
			return ok
		})
		rv, ok := casted.(pango.FontFamilier)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
		}
		_family = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type pango.FontFacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFacer)
			return ok
		})
		rv, ok := casted.(pango.FontFacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFacer")
		}
		_face = rv
	}

	ok := fn(_family, _face)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_FontChooser_ConnectFontActivated
func _gotk4_gtk4_FontChooser_ConnectFontActivated(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(fontname string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(fontname string))
	}

	var _fontname string // out

	_fontname = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_fontname)
}
