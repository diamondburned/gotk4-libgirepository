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
	GTypeTabAlign = coreglib.Type(girepository.MustFind("Pango", "TabAlign").RegisteredGType())
	GTypeTabArray = coreglib.Type(girepository.MustFind("Pango", "TabArray").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTabAlign, F: marshalTabAlign},
		coreglib.TypeMarshaler{T: GTypeTabArray, F: marshalTabArray},
	})
}

// TabAlign: PangoTabAlign specifies where a tab stop appears relative to the
// text.
type TabAlign C.gint

const (
	// TabLeft: tab stop appears to the left of the text.
	TabLeft TabAlign = iota
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	return TabAlign(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for TabAlign.
func (t TabAlign) String() string {
	switch t {
	case TabLeft:
		return "Left"
	default:
		return fmt.Sprintf("TabAlign(%d)", t)
	}
}

// TabArray: PangoTabArray contains an array of tab stops.
//
// PangoTabArray can be used to set tab stops in a PangoLayout. Each tab stop
// has an alignment and a position.
//
// An instance of this type is always passed by reference.
type TabArray struct {
	*tabArray
}

// tabArray is the struct that's finalized.
type tabArray struct {
	native unsafe.Pointer
}

var GIRInfoTabArray = girepository.MustFind("Pango", "TabArray")

func marshalTabArray(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TabArray{&tabArray{(unsafe.Pointer)(b)}}, nil
}
