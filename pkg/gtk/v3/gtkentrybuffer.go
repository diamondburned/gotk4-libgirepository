// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// EntryBufferClass: instance of this type is always passed by reference.
type EntryBufferClass struct {
	*entryBufferClass
}

// entryBufferClass is the struct that's finalized.
type entryBufferClass struct {
	native unsafe.Pointer
}

var GIRInfoEntryBufferClass = girepository.MustFind("Gtk", "EntryBufferClass")
