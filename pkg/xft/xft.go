// Code generated by girgen. DO NOT EDIT.

package xft

import (
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{

		// Objects/Classes
	})
}

type Color struct {
	native *C.XftColor
}

type Draw struct {
	native *C.XftDraw
}

type Font struct {
	native *C.XftFont
}

type GlyphSpec struct {
	native *C.XftGlyphSpec
}