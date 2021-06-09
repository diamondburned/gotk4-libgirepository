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
		{T: externglib.Type(C.graphene_plane_get_type()), F: marshalPlane},
	})
}

// Plane: a 2D plane that extends infinitely in a 3D volume.
//
// The contents of the `graphene_plane_t` are private, and should not be
// modified directly.
type Plane struct {
	native C.graphene_plane_t
}

// WrapPlane wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPlane(ptr unsafe.Pointer) *Plane {
	if ptr == nil {
		return nil
	}

	return (*Plane)(ptr)
}

func marshalPlane(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPlane(unsafe.Pointer(b)), nil
}

// NewPlaneAlloc constructs a struct Plane.
func NewPlaneAlloc() *Plane {
	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_alloc()

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_plane, func(v *Plane) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _plane
}

// Native returns the underlying C source pointer.
func (p *Plane) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Distance computes the distance of @point from a #graphene_plane_t.
func (p *Plane) Distance(point *Point3D) float32 {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var _cret C.float

	cret = C.graphene_plane_distance(_arg0, _arg1)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_plane_t are equal.
func (a *Plane) Equal(b *Plane) bool {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_plane_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool

	cret = C.graphene_plane_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_plane_alloc().
func (p *Plane) Free() {
	var _arg0 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	C.graphene_plane_free(_arg0)
}

// Constant retrieves the distance along the normal vector of the given
// #graphene_plane_t from the origin.
func (p *Plane) Constant() float32 {
	var _arg0 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var _cret C.float

	cret = C.graphene_plane_get_constant(_arg0)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Normal retrieves the normal vector pointing towards the origin of the given
// #graphene_plane_t.
func (p *Plane) Normal() Vec3 {
	var _arg0 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var _normal Vec3

	C.graphene_plane_get_normal(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_normal)))

	return _normal
}

// Init initializes the given #graphene_plane_t using the given @normal vector
// and @constant values.
func (p *Plane) Init(normal *Vec3, constant float32) *Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_vec3_t
	var _arg2 C.float

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(normal.Native()))
	_arg2 = C.float(constant)

	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_init(_arg0, _arg1, _arg2)

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))

	return _plane
}

// InitFromPlane initializes the given #graphene_plane_t using the normal vector
// and constant of another #graphene_plane_t.
func (p *Plane) InitFromPlane(src *Plane) *Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_plane_t)(unsafe.Pointer(src.Native()))

	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_init_from_plane(_arg0, _arg1)

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))

	return _plane
}

// InitFromPoint initializes the given #graphene_plane_t using the given normal
// vector and an arbitrary co-planar point.
func (p *Plane) InitFromPoint(normal *Vec3, point *Point3D) *Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_vec3_t
	var _arg2 *C.graphene_point3d_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(normal.Native()))
	_arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_init_from_point(_arg0, _arg1, _arg2)

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))

	return _plane
}

// InitFromPoints initializes the given #graphene_plane_t using the 3 provided
// co-planar points.
//
// The winding order is counter-clockwise, and determines which direction the
// normal vector will point.
func (p *Plane) InitFromPoints(a *Point3D, b *Point3D, c *Point3D) *Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_point3d_t
	var _arg2 *C.graphene_point3d_t
	var _arg3 *C.graphene_point3d_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	_arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))
	_arg3 = (*C.graphene_point3d_t)(unsafe.Pointer(c.Native()))

	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_init_from_points(_arg0, _arg1, _arg2, _arg3)

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))

	return _plane
}

// InitFromVec4 initializes the given #graphene_plane_t using the components of
// the given #graphene_vec4_t vector.
func (p *Plane) InitFromVec4(src *Vec4) *Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_vec4_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_vec4_t)(unsafe.Pointer(src.Native()))

	var _cret *C.graphene_plane_t

	cret = C.graphene_plane_init_from_vec4(_arg0, _arg1)

	var _plane *Plane

	_plane = WrapPlane(unsafe.Pointer(_cret))

	return _plane
}

// Negate negates the normal vector and constant of a #graphene_plane_t,
// effectively mirroring the plane across the origin.
func (p *Plane) Negate() Plane {
	var _arg0 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var _res Plane

	C.graphene_plane_negate(_arg0, (*C.graphene_plane_t)(unsafe.Pointer(&_res)))

	return _res
}

// Normalize normalizes the vector of the given #graphene_plane_t, and adjusts
// the constant accordingly.
func (p *Plane) Normalize() Plane {
	var _arg0 *C.graphene_plane_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var _res Plane

	C.graphene_plane_normalize(_arg0, (*C.graphene_plane_t)(unsafe.Pointer(&_res)))

	return _res
}

// Transform transforms a #graphene_plane_t @p using the given @matrix and
// @normal_matrix.
//
// If @normal_matrix is nil, a transformation matrix for the plane normal will
// be computed from @matrix. If you are transforming multiple planes using the
// same @matrix it's recommended to compute the normal matrix beforehand to
// avoid incurring in the cost of recomputing it every time.
func (p *Plane) Transform(matrix *Matrix, normalMatrix *Matrix) Plane {
	var _arg0 *C.graphene_plane_t
	var _arg1 *C.graphene_matrix_t
	var _arg2 *C.graphene_matrix_t

	_arg0 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(matrix.Native()))
	_arg2 = (*C.graphene_matrix_t)(unsafe.Pointer(normalMatrix.Native()))

	var _res Plane

	C.graphene_plane_transform(_arg0, _arg1, _arg2, (*C.graphene_plane_t)(unsafe.Pointer(&_res)))

	return _res
}