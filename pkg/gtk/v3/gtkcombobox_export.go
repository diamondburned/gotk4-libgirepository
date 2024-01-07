// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_ComboBox_ConnectChanged
func _gotk4_gtk3_ComboBox_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_ComboBox_ConnectFormatEntryText
func _gotk4_gtk3_ComboBox_ConnectFormatEntryText(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret *C.gchar) {
	var f func(path string) (utf8 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string) (utf8 string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := f(_path)

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_gtk3_ComboBox_ConnectPopdown
func _gotk4_gtk3_ComboBox_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_ComboBox_ConnectPopup
func _gotk4_gtk3_ComboBox_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
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
