// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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

// Quad: 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
//
// An instance of this type is always passed by reference.
type Quad struct {
	*quad
}

// quad is the struct that's finalized.
type quad struct {
	native *C.graphene_quad_t
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Quad{&quad{(*C.graphene_quad_t)(b)}}, nil
}

// NewQuadAlloc constructs a struct Quad.
func NewQuadAlloc() *Quad {
	var _cret *C.graphene_quad_t // in

	_cret = C.graphene_quad_alloc()

	var _quad *Quad // out

	_quad = (*Quad)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_quad)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_quad_free((*C.graphene_quad_t)(intern.C))
		},
	)

	return _quad
}

// Bounds computes the bounding rectangle of q and places it into r.
func (q *Quad) Bounds() Rect {
	var _arg0 *C.graphene_quad_t // out
	var _arg1 C.graphene_rect_t  // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quad_bounds(_arg0, &_arg1)
	runtime.KeepAlive(q)

	var _r Rect // out

	_r = *(*Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _r
}

// Contains checks if the given #graphene_quad_t contains the given
// #graphene_point_t.
func (q *Quad) Contains(p *Point) bool {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 *C.graphene_point_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_quad_contains(_arg0, _arg1)
	runtime.KeepAlive(q)
	runtime.KeepAlive(p)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Point retrieves the point of a #graphene_quad_t at the given index.
func (q *Quad) Point(index_ uint) *Point {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 C.uint              // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.uint(index_)

	_cret = C.graphene_quad_get_point(_arg0, _arg1)
	runtime.KeepAlive(q)
	runtime.KeepAlive(index_)

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point
}

// Init initializes a #graphene_quad_t with the given points.
func (q *Quad) Init(p1 *Point, p2 *Point, p3 *Point, p4 *Point) *Quad {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.graphene_point_t // out
	var _cret *C.graphene_quad_t  // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p1)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p2)))
	_arg3 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p3)))
	_arg4 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p4)))

	_cret = C.graphene_quad_init(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(q)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
	runtime.KeepAlive(p3)
	runtime.KeepAlive(p4)

	var _quad *Quad // out

	_quad = (*Quad)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quad
}

// InitFromPoints initializes a #graphene_quad_t using an array of points.
func (q *Quad) InitFromPoints(points [4]Point) *Quad {
	var _arg0 *C.graphene_quad_t  // out
	var _arg1 *C.graphene_point_t // out
	var _cret *C.graphene_quad_t  // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(&points))

	_cret = C.graphene_quad_init_from_points(_arg0, _arg1)
	runtime.KeepAlive(q)
	runtime.KeepAlive(points)

	var _quad *Quad // out

	_quad = (*Quad)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quad
}

// InitFromRect initializes a #graphene_quad_t using the four corners of the
// given #graphene_rect_t.
func (q *Quad) InitFromRect(r *Rect) *Quad {
	var _arg0 *C.graphene_quad_t // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.graphene_quad_t // in

	_arg0 = (*C.graphene_quad_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(r)))

	_cret = C.graphene_quad_init_from_rect(_arg0, _arg1)
	runtime.KeepAlive(q)
	runtime.KeepAlive(r)

	var _quad *Quad // out

	_quad = (*Quad)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quad
}
