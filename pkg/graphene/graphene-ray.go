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
	cret := new(C.graphene_ray_t)
	var goret *Ray

	cret = C.graphene_ray_alloc()

	goret = WrapRay(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Ray) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (r *Ray) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Equal checks whether the two given #graphene_ray_t are equal.
func (a *Ray) Equal(b *Ray) bool {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_ray_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_ray_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_ray_equal(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Free frees the resources allocated by graphene_ray_alloc().
func (r *Ray) Free() {
	var arg0 *C.graphene_ray_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	C.graphene_ray_free(arg0)
}

// ClosestPointToPoint computes the point on the given #graphene_ray_t that is
// closest to the given point @p.
func (r *Ray) ClosestPointToPoint(p *Point3D) *Point3D {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	arg2 := new(C.graphene_point3d_t)
	var ret2 *Point3D

	C.graphene_ray_get_closest_point_to_point(arg0, arg1, arg2)

	ret2 = WrapPoint3D(unsafe.Pointer(arg2))

	return ret2
}

// Direction retrieves the direction of the given #graphene_ray_t.
func (r *Ray) Direction() *Vec3 {
	var arg0 *C.graphene_ray_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	arg1 := new(C.graphene_vec3_t)
	var ret1 *Vec3

	C.graphene_ray_get_direction(arg0, arg1)

	ret1 = WrapVec3(unsafe.Pointer(arg1))

	return ret1
}

// DistanceToPlane computes the distance of the origin of the given
// #graphene_ray_t from the given plane.
//
// If the ray does not intersect the plane, this function returns `INFINITY`.
func (r *Ray) DistanceToPlane(p *Plane) float32 {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_plane_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	var cret C.float
	var goret float32

	cret = C.graphene_ray_get_distance_to_plane(arg0, arg1)

	goret = float32(cret)

	return goret
}

// DistanceToPoint computes the distance of the closest approach between the
// given #graphene_ray_t @r and the point @p.
//
// The closest approach to a ray from a point is the distance between the point
// and the projection of the point on the ray itself.
func (r *Ray) DistanceToPoint(p *Point3D) float32 {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var cret C.float
	var goret float32

	cret = C.graphene_ray_get_distance_to_point(arg0, arg1)

	goret = float32(cret)

	return goret
}

// Origin retrieves the origin of the given #graphene_ray_t.
func (r *Ray) Origin() *Point3D {
	var arg0 *C.graphene_ray_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	arg1 := new(C.graphene_point3d_t)
	var ret1 *Point3D

	C.graphene_ray_get_origin(arg0, arg1)

	ret1 = WrapPoint3D(unsafe.Pointer(arg1))

	return ret1
}

// PositionAt retrieves the coordinates of a point at the distance @t along the
// given #graphene_ray_t.
func (r *Ray) PositionAt(t float32) *Point3D {
	var arg0 *C.graphene_ray_t
	var arg1 C.float

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(t)

	arg2 := new(C.graphene_point3d_t)
	var ret2 *Point3D

	C.graphene_ray_get_position_at(arg0, arg1, arg2)

	ret2 = WrapPoint3D(unsafe.Pointer(arg2))

	return ret2
}

// Init initializes the given #graphene_ray_t using the given @origin and
// @direction values.
func (r *Ray) Init(origin *Point3D, direction *Vec3) *Ray {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_point3d_t
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(origin.Native()))
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	var cret *C.graphene_ray_t
	var goret *Ray

	cret = C.graphene_ray_init(arg0, arg1, arg2)

	goret = WrapRay(unsafe.Pointer(cret))

	return goret
}

// InitFromRay initializes the given #graphene_ray_t using the origin and
// direction values of another #graphene_ray_t.
func (r *Ray) InitFromRay(src *Ray) *Ray {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_ray_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_ray_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_ray_t
	var goret *Ray

	cret = C.graphene_ray_init_from_ray(arg0, arg1)

	goret = WrapRay(unsafe.Pointer(cret))

	return goret
}

// InitFromVec3 initializes the given #graphene_ray_t using the given vectors.
func (r *Ray) InitFromVec3(origin *Vec3, direction *Vec3) *Ray {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_vec3_t
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(origin.Native()))
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	var cret *C.graphene_ray_t
	var goret *Ray

	cret = C.graphene_ray_init_from_vec3(arg0, arg1, arg2)

	goret = WrapRay(unsafe.Pointer(cret))

	return goret
}

// IntersectBox intersects the given #graphene_ray_t @r with the given
// #graphene_box_t @b.
func (r *Ray) IntersectBox(b *Box) (tOut float32, rayIntersectionKind RayIntersectionKind) {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	arg2 := new(C.float)
	var ret2 float32
	var cret C.graphene_ray_intersection_kind_t
	var goret RayIntersectionKind

	cret = C.graphene_ray_intersect_box(arg0, arg1, arg2)

	ret2 = float32(*arg2)
	goret = RayIntersectionKind(cret)

	return ret2, goret
}

// IntersectSphere intersects the given #graphene_ray_t @r with the given
// #graphene_sphere_t @s.
func (r *Ray) IntersectSphere(s *Sphere) (tOut float32, rayIntersectionKind RayIntersectionKind) {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_sphere_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	arg2 := new(C.float)
	var ret2 float32
	var cret C.graphene_ray_intersection_kind_t
	var goret RayIntersectionKind

	cret = C.graphene_ray_intersect_sphere(arg0, arg1, arg2)

	ret2 = float32(*arg2)
	goret = RayIntersectionKind(cret)

	return ret2, goret
}

// IntersectTriangle intersects the given #graphene_ray_t @r with the given
// #graphene_triangle_t @t.
func (r *Ray) IntersectTriangle(t *Triangle) (tOut float32, rayIntersectionKind RayIntersectionKind) {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_triangle_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	arg2 := new(C.float)
	var ret2 float32
	var cret C.graphene_ray_intersection_kind_t
	var goret RayIntersectionKind

	cret = C.graphene_ray_intersect_triangle(arg0, arg1, arg2)

	ret2 = float32(*arg2)
	goret = RayIntersectionKind(cret)

	return ret2, goret
}

// IntersectsBox checks whether the given #graphene_ray_t @r intersects the
// given #graphene_box_t @b.
//
// See also: graphene_ray_intersect_box()
func (r *Ray) IntersectsBox(b *Box) bool {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_ray_intersects_box(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// IntersectsSphere checks if the given #graphene_ray_t @r intersects the given
// #graphene_sphere_t @s.
//
// See also: graphene_ray_intersect_sphere()
func (r *Ray) IntersectsSphere(s *Sphere) bool {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_sphere_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_ray_intersects_sphere(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// IntersectsTriangle checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectsTriangle(t *Triangle) bool {
	var arg0 *C.graphene_ray_t
	var arg1 *C.graphene_triangle_t

	arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_ray_intersects_triangle(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}
