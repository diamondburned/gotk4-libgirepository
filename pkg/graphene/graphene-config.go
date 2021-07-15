// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <graphene-gobject.h>
import "C"

type SIMD4F struct {
	nocopy gextras.NoCopy
	native *C.graphene_simd4f_t
}

type SIMD4X4F struct {
	nocopy gextras.NoCopy
	native *C.graphene_simd4x4f_t
}
