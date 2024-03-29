// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
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
	GTypeOverline = coreglib.Type(girepository.MustFind("Pango", "Overline").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOverline, F: marshalOverline},
	})
}

// Overline: PangoOverline enumeration is used to specify whether text should be
// overlined, and if so, the type of line.
type Overline C.gint

const (
	// OverlineNone: no overline should be drawn.
	OverlineNone Overline = iota
	// OverlineSingle: draw a single line above the ink extents of the text
	// being underlined.
	OverlineSingle
)

func marshalOverline(p uintptr) (interface{}, error) {
	return Overline(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Overline.
func (o Overline) String() string {
	switch o {
	case OverlineNone:
		return "None"
	case OverlineSingle:
		return "Single"
	default:
		return fmt.Sprintf("Overline(%d)", o)
	}
}
