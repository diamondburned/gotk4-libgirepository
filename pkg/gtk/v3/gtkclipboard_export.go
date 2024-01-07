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

//export _gotk4_gtk3_ClipboardReceivedFunc
func _gotk4_gtk3_ClipboardReceivedFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
	var fn ClipboardReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardReceivedFunc)
	}

	var _clipboard *Clipboard         // out
	var _selectionData *SelectionData // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn(_clipboard, _selectionData)
}

//export _gotk4_gtk3_ClipboardTextReceivedFunc
func _gotk4_gtk3_ClipboardTextReceivedFunc(arg1 *C.void, arg2 *C.gchar, arg3 C.gpointer) {
	var fn ClipboardTextReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardTextReceivedFunc)
	}

	var _clipboard *Clipboard // out
	var _text string          // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	}

	fn(_clipboard, _text)
}

//export _gotk4_gtk3_Clipboard_ConnectOwnerChange
func _gotk4_gtk3_Clipboard_ConnectOwnerChange(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(event *gdk.EventOwnerChange)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(event *gdk.EventOwnerChange))
	}

	var _event *gdk.EventOwnerChange // out

	_event = (*gdk.EventOwnerChange)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_event)
}
