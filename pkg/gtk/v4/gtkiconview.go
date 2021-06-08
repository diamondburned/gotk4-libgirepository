// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// IconViewForeachFunc: a function used by gtk_icon_view_selected_foreach() to
// map all selected rows.
//
// It will be called on every selected row in the view.
type IconViewForeachFunc func()

//export gotk4_IconViewForeachFunc
func gotk4_IconViewForeachFunc(arg0 *C.GtkIconView, arg1 *C.GtkTreePath, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(IconViewForeachFunc)
	fn()
}
