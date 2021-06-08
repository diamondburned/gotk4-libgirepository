// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

// RoundedRect: a rectangular region with rounded corners.
//
// Application code should normalize rectangles using
// [method@Gsk.RoundedRect.normalize]; this function will ensure that the bounds
// of the rectangle are normalized and ensure that the corner values are
// positive and the corners do not overlap.
//
// All functions taking a `GskRoundedRect` as an argument will internally
// operate on a normalized copy; all functions returning a `GskRoundedRect` will
// always return a normalized one.
//
// The algorithm used for normalizing corner sizes is described in the CSS
// specification (https://drafts.csswg.org/css-backgrounds-3/#border-radius).
type RoundedRect struct {
	native C.GskRoundedRect
}

// WrapRoundedRect wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRoundedRect(ptr unsafe.Pointer) *RoundedRect {
	if ptr == nil {
		return nil
	}

	return (*RoundedRect)(ptr)
}

func marshalRoundedRect(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRoundedRect(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RoundedRect) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Bounds gets the field inside the struct.
func (r *RoundedRect) Bounds() graphene.Rect {
	var v graphene.Rect
	v = *graphene.WrapRect(unsafe.Pointer(&r.native.bounds))
	return v
}

// Corner gets the field inside the struct.
func (r *RoundedRect) Corner() [4]graphene.Size {
	var v [4]graphene.Size
	v = *(*[4]graphene.Size)(unsafe.Pointer(r.native.corner))
	return v
}

// ContainsPoint checks if the given @point is inside the rounded rectangle.
func (s *RoundedRect) ContainsPoint(point *graphene.Point) bool {
	var arg0 *C.GskRoundedRect
	var arg1 *C.graphene_point_t

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(point.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gsk_rounded_rect_contains_point(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// ContainsRect checks if the given @rect is contained inside the rounded
// rectangle.
func (s *RoundedRect) ContainsRect(rect *graphene.Rect) bool {
	var arg0 *C.GskRoundedRect
	var arg1 *C.graphene_rect_t

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(rect.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gsk_rounded_rect_contains_rect(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Init initializes the given `GskRoundedRect` with the given values.
//
// This function will implicitly normalize the `GskRoundedRect` before
// returning.
func (s *RoundedRect) Init(bounds *graphene.Rect, topLeft *graphene.Size, topRight *graphene.Size, bottomRight *graphene.Size, bottomLeft *graphene.Size) *RoundedRect {
	var arg0 *C.GskRoundedRect
	var arg1 *C.graphene_rect_t
	var arg2 *C.graphene_size_t
	var arg3 *C.graphene_size_t
	var arg4 *C.graphene_size_t
	var arg5 *C.graphene_size_t

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	arg2 = (*C.graphene_size_t)(unsafe.Pointer(topLeft.Native()))
	arg3 = (*C.graphene_size_t)(unsafe.Pointer(topRight.Native()))
	arg4 = (*C.graphene_size_t)(unsafe.Pointer(bottomRight.Native()))
	arg5 = (*C.graphene_size_t)(unsafe.Pointer(bottomLeft.Native()))

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_init(arg0, arg1, arg2, arg3, arg4, arg5)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}

// InitCopy initializes @self using the given @src rectangle.
//
// This function will not normalize the `GskRoundedRect`, so make sure the
// source is normalized.
func (s *RoundedRect) InitCopy(src *RoundedRect) *RoundedRect {
	var arg0 *C.GskRoundedRect
	var arg1 *C.GskRoundedRect

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GskRoundedRect)(unsafe.Pointer(src.Native()))

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_init_copy(arg0, arg1)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}

// InitFromRect initializes @self to the given @bounds and sets the radius of
// all four corners to @radius.
func (s *RoundedRect) InitFromRect(bounds *graphene.Rect, radius float32) *RoundedRect {
	var arg0 *C.GskRoundedRect
	var arg1 *C.graphene_rect_t
	var arg2 C.float

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	arg2 = C.float(radius)

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_init_from_rect(arg0, arg1, arg2)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}

// IntersectsRect checks if part of the given @rect is contained inside the
// rounded rectangle.
func (s *RoundedRect) IntersectsRect(rect *graphene.Rect) bool {
	var arg0 *C.GskRoundedRect
	var arg1 *C.graphene_rect_t

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(rect.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gsk_rounded_rect_intersects_rect(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// IsRectilinear checks if all corners of @self are right angles and the
// rectangle covers all of its bounds.
//
// This information can be used to decide if [ctor@Gsk.ClipNode.new] or
// [ctor@Gsk.RoundedClipNode.new] should be called.
func (s *RoundedRect) IsRectilinear() bool {
	var arg0 *C.GskRoundedRect

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gsk_rounded_rect_is_rectilinear(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Normalize normalizes the passed rectangle.
//
// This function will ensure that the bounds of the rectangle are normalized and
// ensure that the corner values are positive and the corners do not overlap.
func (s *RoundedRect) Normalize() *RoundedRect {
	var arg0 *C.GskRoundedRect

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_normalize(arg0)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}

// Offset offsets the bound's origin by @dx and @dy.
//
// The size and corners of the rectangle are unchanged.
func (s *RoundedRect) Offset(dx float32, dy float32) *RoundedRect {
	var arg0 *C.GskRoundedRect
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = C.float(dx)
	arg2 = C.float(dy)

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_offset(arg0, arg1, arg2)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}

// Shrink shrinks (or grows) the given rectangle by moving the 4 sides according
// to the offsets given.
//
// The corner radii will be changed in a way that tries to keep the center of
// the corner circle intact. This emulates CSS behavior.
//
// This function also works for growing rectangles if you pass negative values
// for the @top, @right, @bottom or @left.
func (s *RoundedRect) Shrink(top float32, right float32, bottom float32, left float32) *RoundedRect {
	var arg0 *C.GskRoundedRect
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float
	var arg4 C.float

	arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	arg1 = C.float(top)
	arg2 = C.float(right)
	arg3 = C.float(bottom)
	arg4 = C.float(left)

	var cret *C.GskRoundedRect
	var goret *RoundedRect

	cret = C.gsk_rounded_rect_shrink(arg0, arg1, arg2, arg3, arg4)

	goret = WrapRoundedRect(unsafe.Pointer(cret))

	return goret
}
