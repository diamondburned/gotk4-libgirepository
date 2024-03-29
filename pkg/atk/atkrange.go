// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeRange = coreglib.Type(girepository.MustFind("Atk", "Range").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRange, F: marshalRange},
	})
}

// Range are used on Value, in order to represent the full range of a given
// component (for example an slider or a range control), or to define each
// individual subrange this full range is splitted if available. See Value
// documentation for further details.
//
// An instance of this type is always passed by reference.
type Range struct {
	*_range
}

// _range is the struct that's finalized.
type _range struct {
	native unsafe.Pointer
}

var GIRInfoRange = girepository.MustFind("Atk", "Range")

func marshalRange(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Range{&_range{(unsafe.Pointer)(b)}}, nil
}
