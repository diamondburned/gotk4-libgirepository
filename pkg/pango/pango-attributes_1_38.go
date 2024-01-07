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

// AttrFontFeatures: PangoAttrFontFeatures structure is used to represent
// OpenType font features as an attribute.
//
// An instance of this type is always passed by reference.
type AttrFontFeatures struct {
	*attrFontFeatures
}

// attrFontFeatures is the struct that's finalized.
type attrFontFeatures struct {
	native unsafe.Pointer
}

var GIRInfoAttrFontFeatures = girepository.MustFind("Pango", "AttrFontFeatures")