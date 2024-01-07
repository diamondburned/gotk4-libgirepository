// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk3_ClipboardImageReceivedFunc
func _gotk4_gtk3_ClipboardImageReceivedFunc(arg1 *C.void, arg2 *C.GdkPixbuf, arg3 C.gpointer) {
	var fn ClipboardImageReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardImageReceivedFunc)
	}

	var _clipboard *Clipboard     // out
	var _pixbuf *gdkpixbuf.Pixbuf // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	{
		obj := coreglib.Take(unsafe.Pointer(arg2))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	fn(_clipboard, _pixbuf)
}
