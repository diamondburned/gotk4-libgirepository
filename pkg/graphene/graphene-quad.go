// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_quad_get_type()), F: marshalQuad},
	})
}

// Quad: a 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
type Quad struct {
	native C.graphene_quad_t
}

// WrapQuad wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapQuad(ptr unsafe.Pointer) *Quad {
	if ptr == nil {
		return nil
	}

	return (*Quad)(ptr)
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapQuad(unsafe.Pointer(b)), nil
}

// NewQuadAlloc constructs a struct Quad.
func NewQuadAlloc() *Quad {
	cret := new(C.graphene_quad_t)
	var goret *Quad

	cret = C.graphene_quad_alloc()

	goret = WrapQuad(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Quad) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (q *Quad) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// Bounds computes the bounding rectangle of @q and places it into @r.
func (q *Quad) Bounds() *Rect {
	var arg0 *C.graphene_quad_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))

	arg1 := new(C.graphene_rect_t)
	var ret1 *Rect

	C.graphene_quad_bounds(arg0, arg1)

	ret1 = WrapRect(unsafe.Pointer(arg1))

	return ret1
}

// Contains checks if the given #graphene_quad_t contains the given
// #graphene_point_t.
func (q *Quad) Contains(p *Point) bool {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_quad_contains(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Free frees the resources allocated by graphene_quad_alloc()
func (q *Quad) Free() {
	var arg0 *C.graphene_quad_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))

	C.graphene_quad_free(arg0)
}

// Point retrieves the point of a #graphene_quad_t at the given index.
func (q *Quad) Point(index_ uint) *Point {
	var arg0 *C.graphene_quad_t
	var arg1 C.uint

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = C.uint(index_)

	var cret *C.graphene_point_t
	var goret *Point

	cret = C.graphene_quad_get_point(arg0, arg1)

	goret = WrapPoint(unsafe.Pointer(cret))

	return goret
}

// Init initializes a #graphene_quad_t with the given points.
func (q *Quad) Init(p1 *Point, p2 *Point, p3 *Point, p4 *Point) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t
	var arg2 *C.graphene_point_t
	var arg3 *C.graphene_point_t
	var arg4 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p1.Native()))
	arg2 = (*C.graphene_point_t)(unsafe.Pointer(p2.Native()))
	arg3 = (*C.graphene_point_t)(unsafe.Pointer(p3.Native()))
	arg4 = (*C.graphene_point_t)(unsafe.Pointer(p4.Native()))

	var cret *C.graphene_quad_t
	var goret *Quad

	cret = C.graphene_quad_init(arg0, arg1, arg2, arg3, arg4)

	goret = WrapQuad(unsafe.Pointer(cret))

	return goret
}

// InitFromPoints initializes a #graphene_quad_t using an array of points.
func (q *Quad) InitFromPoints(points [4]Point) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(&points))
	defer runtime.KeepAlive(&arg1)

	var cret *C.graphene_quad_t
	var goret *Quad

	cret = C.graphene_quad_init_from_points(arg0, arg1)

	goret = WrapQuad(unsafe.Pointer(cret))

	return goret
}

// InitFromRect initializes a #graphene_quad_t using the four corners of the
// given #graphene_rect_t.
func (q *Quad) InitFromRect(r *Rect) *Quad {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret *C.graphene_quad_t
	var goret *Quad

	cret = C.graphene_quad_init_from_rect(arg0, arg1)

	goret = WrapQuad(unsafe.Pointer(cret))

	return goret
}
