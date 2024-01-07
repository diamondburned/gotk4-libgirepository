// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// RoundedRect: rectangular region with rounded corners.
//
// Application code should normalize rectangles using
// gsk.RoundedRect.Normalize(); this function will ensure that the bounds of the
// rectangle are normalized and ensure that the corner values are positive and
// the corners do not overlap.
//
// All functions taking a GskRoundedRect as an argument will internally operate
// on a normalized copy; all functions returning a GskRoundedRect will always
// return a normalized one.
//
// The algorithm used for normalizing corner sizes is described in the CSS
// specification (https://drafts.csswg.org/css-backgrounds-3/#border-radius).
//
// An instance of this type is always passed by reference.
type RoundedRect struct {
	*roundedRect
}

// roundedRect is the struct that's finalized.
type roundedRect struct {
	native unsafe.Pointer
}

var GIRInfoRoundedRect = girepository.MustFind("Gsk", "RoundedRect")

// Bounds bounds of the rectangle.
func (r *RoundedRect) Bounds() *graphene.Rect {
	offset := GIRInfoRoundedRect.StructFieldOffset("bounds")
	valptr := (**graphene.Rect)(unsafe.Add(r.native, offset))
	var _v *graphene.Rect // out
	_v = (*graphene.Rect)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Corner: size of the 4 rounded corners.
func (r *RoundedRect) Corner() [4]graphene.Size {
	offset := GIRInfoRoundedRect.StructFieldOffset("corner")
	valptr := (*[4]graphene.Size)(unsafe.Add(r.native, offset))
	var _v [4]graphene.Size // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = *(*graphene.Size)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}
	return _v
}
