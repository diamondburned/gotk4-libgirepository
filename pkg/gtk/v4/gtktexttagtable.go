// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// TextTagTableForeach: a function used with gtk_text_tag_table_foreach(), to
// iterate over every `GtkTextTag` inside a `GtkTextTagTable`.
type TextTagTableForeach func()

//export gotk4_TextTagTableForeach
func gotk4_TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TextTagTableForeach)
	fn()
}