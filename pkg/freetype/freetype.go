// Code generated by girgen. DO NOT EDIT.

package freetype

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

type Int32 int32

type Bitmap struct {
	native *C.FT_Bitmap
}

type Face struct {
	native *C.FT_Face
}

type Library struct {
	native *C.FT_Library
}