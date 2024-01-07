// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_CssProvider_ConnectParsingError
func _gotk4_gtk4_CssProvider_ConnectParsingError(arg0 C.gpointer, arg1 *C.void, arg2 *C.GError, arg3 C.guintptr) {
	var f func(section *CSSSection, err error)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(section *CSSSection, err error))
	}

	var _section *CSSSection // out
	var _err error           // out

	_section = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_section)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				Gtk.GIRInfoCSSSection.InvokeRecordMethod("unref", args[:], nil)
			}
		},
	)
	_err = gerror.Take(unsafe.Pointer(arg2))

	f(_section, _err)
}
