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
		{T: externglib.Type(C.graphene_ray_get_type()), F: marshalRay},
	})
}

// RayIntersectionKind: the type of intersection.
type RayIntersectionKind int

const (
	// RayIntersectionKindNone: no intersection
	RayIntersectionKindNone RayIntersectionKind = 0
	// RayIntersectionKindEnter: the ray is entering the intersected object
	RayIntersectionKindEnter RayIntersectionKind = 1
	// RayIntersectionKindLeave: the ray is leaving the intersected object
	RayIntersectionKindLeave RayIntersectionKind = 2
)

// Ray: a ray emitted from an origin in a given direction.
//
// The contents of the `graphene_ray_t` structure are private, and should not be
// modified directly.
type Ray struct {
	native C.graphene_ray_t
}

// WrapRay wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRay(ptr unsafe.Pointer) *Ray {
	if ptr == nil {
		return nil
	}

	return (*Ray)(ptr)
}

func marshalRay(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRay(unsafe.Pointer(b)), nil
}

// NewRayAlloc constructs a struct Ray.
func NewRayAlloc() *Ray {
	var _cret *C.graphene_ray_t

	cret = C.graphene_ray_alloc()

	var _ray *Ray

	_ray = WrapRay(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ray, func(v *Ray) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _ray
}

// Native returns the underlying C source pointer.
func (r *Ray) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Equal checks whether the two given #graphene_ray_t are equal.
func (a *Ray) Equal(b *Ray) bool {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_ray_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_ray_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool

	cret = C.graphene_ray_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_ray_alloc().
func (r *Ray) Free() {
	var _arg0 *C.graphene_ray_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	C.graphene_ray_free(_arg0)
}

// ClosestPointToPoint computes the point on the given #graphene_ray_t that is
// closest to the given point @p.
func (r *Ray) ClosestPointToPoint(p *Point3D) Point3D {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _res Point3D

	C.graphene_ray_get_closest_point_to_point(_arg0, _arg1, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// Direction retrieves the direction of the given #graphene_ray_t.
func (r *Ray) Direction() Vec3 {
	var _arg0 *C.graphene_ray_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	var _direction Vec3

	C.graphene_ray_get_direction(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_direction)))

	return _direction
}

// DistanceToPlane computes the distance of the origin of the given
// #graphene_ray_t from the given plane.
//
// If the ray does not intersect the plane, this function returns `INFINITY`.
func (r *Ray) DistanceToPlane(p *Plane) float32 {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_plane_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var _cret C.float

	cret = C.graphene_ray_get_distance_to_plane(_arg0, _arg1)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// DistanceToPoint computes the distance of the closest approach between the
// given #graphene_ray_t @r and the point @p.
//
// The closest approach to a ray from a point is the distance between the point
// and the projection of the point on the ray itself.
func (r *Ray) DistanceToPoint(p *Point3D) float32 {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_point3d_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _cret C.float

	cret = C.graphene_ray_get_distance_to_point(_arg0, _arg1)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Origin retrieves the origin of the given #graphene_ray_t.
func (r *Ray) Origin() Point3D {
	var _arg0 *C.graphene_ray_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	var _origin Point3D

	C.graphene_ray_get_origin(_arg0, (*C.graphene_point3d_t)(unsafe.Pointer(&_origin)))

	return _origin
}

// PositionAt retrieves the coordinates of a point at the distance @t along the
// given #graphene_ray_t.
func (r *Ray) PositionAt(t float32) Point3D {
	var _arg0 *C.graphene_ray_t
	var _arg1 C.float

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = C.float(t)

	var _position Point3D

	C.graphene_ray_get_position_at(_arg0, _arg1, (*C.graphene_point3d_t)(unsafe.Pointer(&_position)))

	return _position
}

// Init initializes the given #graphene_ray_t using the given @origin and
// @direction values.
func (r *Ray) Init(origin *Point3D, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_point3d_t
	var _arg2 *C.graphene_vec3_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(origin.Native()))
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	var _cret *C.graphene_ray_t

	cret = C.graphene_ray_init(_arg0, _arg1, _arg2)

	var _ray *Ray

	_ray = WrapRay(unsafe.Pointer(_cret))

	return _ray
}

// InitFromRay initializes the given #graphene_ray_t using the origin and
// direction values of another #graphene_ray_t.
func (r *Ray) InitFromRay(src *Ray) *Ray {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_ray_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_ray_t)(unsafe.Pointer(src.Native()))

	var _cret *C.graphene_ray_t

	cret = C.graphene_ray_init_from_ray(_arg0, _arg1)

	var _ray *Ray

	_ray = WrapRay(unsafe.Pointer(_cret))

	return _ray
}

// InitFromVec3 initializes the given #graphene_ray_t using the given vectors.
func (r *Ray) InitFromVec3(origin *Vec3, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_vec3_t
	var _arg2 *C.graphene_vec3_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(origin.Native()))
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	var _cret *C.graphene_ray_t

	cret = C.graphene_ray_init_from_vec3(_arg0, _arg1, _arg2)

	var _ray *Ray

	_ray = WrapRay(unsafe.Pointer(_cret))

	return _ray
}

// IntersectBox intersects the given #graphene_ray_t @r with the given
// #graphene_box_t @b.
func (r *Ray) IntersectBox(b *Box) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_box_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var _arg2 C.float
	var _cret C.graphene_ray_intersection_kind_t

	cret = C.graphene_ray_intersect_box(_arg0, _arg1, &_arg2)

	var _tOut float32
	var _rayIntersectionKind RayIntersectionKind

	_tOut = (float32)(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectSphere intersects the given #graphene_ray_t @r with the given
// #graphene_sphere_t @s.
func (r *Ray) IntersectSphere(s *Sphere) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_sphere_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	var _arg2 C.float
	var _cret C.graphene_ray_intersection_kind_t

	cret = C.graphene_ray_intersect_sphere(_arg0, _arg1, &_arg2)

	var _tOut float32
	var _rayIntersectionKind RayIntersectionKind

	_tOut = (float32)(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectTriangle intersects the given #graphene_ray_t @r with the given
// #graphene_triangle_t @t.
func (r *Ray) IntersectTriangle(t *Triangle) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_triangle_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	var _arg2 C.float
	var _cret C.graphene_ray_intersection_kind_t

	cret = C.graphene_ray_intersect_triangle(_arg0, _arg1, &_arg2)

	var _tOut float32
	var _rayIntersectionKind RayIntersectionKind

	_tOut = (float32)(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectsBox checks whether the given #graphene_ray_t @r intersects the
// given #graphene_box_t @b.
//
// See also: graphene_ray_intersect_box()
func (r *Ray) IntersectsBox(b *Box) bool {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_box_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool

	cret = C.graphene_ray_intersects_box(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsSphere checks if the given #graphene_ray_t @r intersects the given
// #graphene_sphere_t @s.
//
// See also: graphene_ray_intersect_sphere()
func (r *Ray) IntersectsSphere(s *Sphere) bool {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_sphere_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	var _cret C._Bool

	cret = C.graphene_ray_intersects_sphere(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsTriangle checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectsTriangle(t *Triangle) bool {
	var _arg0 *C.graphene_ray_t
	var _arg1 *C.graphene_triangle_t

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	var _cret C._Bool

	cret = C.graphene_ray_intersects_triangle(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}