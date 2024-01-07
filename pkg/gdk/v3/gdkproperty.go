// Code generated by girgen. DO NOT EDIT.

package gdk

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
	GTypePropMode = coreglib.Type(girepository.MustFind("Gdk", "PropMode").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePropMode, F: marshalPropMode},
	})
}

// PropMode describes how existing data is combined with new data when using
// gdk_property_change().
type PropMode C.gint

const (
	// PropModeReplace: new data replaces the existing data.
	PropModeReplace PropMode = iota
	// PropModePrepend: new data is prepended to the existing data.
	PropModePrepend
	// PropModeAppend: new data is appended to the existing data.
	PropModeAppend
)

func marshalPropMode(p uintptr) (interface{}, error) {
	return PropMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PropMode.
func (p PropMode) String() string {
	switch p {
	case PropModeReplace:
		return "Replace"
	case PropModePrepend:
		return "Prepend"
	case PropModeAppend:
		return "Append"
	default:
		return fmt.Sprintf("PropMode(%d)", p)
	}
}
