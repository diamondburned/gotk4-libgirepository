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

// BitsetIter: opaque, stack-allocated struct for iterating over the elements of
// a GtkBitset.
//
// Before a GtkBitsetIter can be used, it needs to be initialized with
// gtk.BitsetIter().InitFirst, gtk.BitsetIter().InitLast or
// gtk.BitsetIter().InitAt.
//
// An instance of this type is always passed by reference.
type BitsetIter struct {
	*bitsetIter
}

// bitsetIter is the struct that's finalized.
type bitsetIter struct {
	native unsafe.Pointer
}

var GIRInfoBitsetIter = girepository.MustFind("Gtk", "BitsetIter")
