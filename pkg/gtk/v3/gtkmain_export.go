// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_KeySnoopFunc
func _gotk4_gtk3_KeySnoopFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gint) {
	var fn KeySnoopFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(KeySnoopFunc)
	}

	var _grabWidget Widgetter // out
	var _event *gdk.EventKey  // out

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
		_grabWidget = rv
	}
	_event = (*gdk.EventKey)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	gint := fn(_grabWidget, _event)

	var _ int

	cret = C.gint(gint)

	return cret
}
