// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// LogAttr: PangoLogAttr structure stores information about the attributes of a
// single character.
//
// An instance of this type is always passed by reference.
type LogAttr struct {
	*logAttr
}

// logAttr is the struct that's finalized.
type logAttr struct {
	native unsafe.Pointer
}

var GIRInfoLogAttr = girepository.MustFind("Pango", "LogAttr")
