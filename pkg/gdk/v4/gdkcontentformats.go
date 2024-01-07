// Code generated by girgen. DO NOT EDIT.

package gdk

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
	GTypeContentFormatsBuilder = coreglib.Type(girepository.MustFind("Gdk", "ContentFormatsBuilder").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContentFormatsBuilder, F: marshalContentFormatsBuilder},
	})
}

// ContentFormatsBuilder: GdkContentFormatsBuilder is an auxiliary struct used
// to create new GdkContentFormats, and should not be kept around.
//
// An instance of this type is always passed by reference.
type ContentFormatsBuilder struct {
	*contentFormatsBuilder
}

// contentFormatsBuilder is the struct that's finalized.
type contentFormatsBuilder struct {
	native unsafe.Pointer
}

var GIRInfoContentFormatsBuilder = girepository.MustFind("Gdk", "ContentFormatsBuilder")

func marshalContentFormatsBuilder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ContentFormatsBuilder{&contentFormatsBuilder{(unsafe.Pointer)(b)}}, nil
}
