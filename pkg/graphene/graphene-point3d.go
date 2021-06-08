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
		{T: externglib.Type(C.graphene_point3d_get_type()), F: marshalPoint3D},
	})
}

// Point3D: a point with three components: X, Y, and Z.
type Point3D struct {
	native C.graphene_point3d_t
}

// WrapPoint3D wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint3D(ptr unsafe.Pointer) *Point3D {
	if ptr == nil {
		return nil
	}

	return (*Point3D)(ptr)
}

func marshalPoint3D(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPoint3D(unsafe.Pointer(b)), nil
}

// NewPoint3DAlloc constructs a struct Point3D.
func NewPoint3DAlloc() *Point3D {
	cret := new(C.graphene_point3d_t)
	var goret *Point3D

	cret = C.graphene_point3d_alloc()

	goret = WrapPoint3D(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Point3D) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (p *Point3D) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// X gets the field inside the struct.
func (p *Point3D) X() float32 {
	var v float32
	v = float32(p.native.x)
	return v
}

// Y gets the field inside the struct.
func (p *Point3D) Y() float32 {
	var v float32
	v = float32(p.native.y)
	return v
}

// Z gets the field inside the struct.
func (p *Point3D) Z() float32 {
	var v float32
	v = float32(p.native.z)
	return v
}

// Cross computes the cross product of the two given #graphene_point3d_t.
func (a *Point3D) Cross(b *Point3D) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	arg2 := new(C.graphene_point3d_t)
	var ret2 *Point3D

	C.graphene_point3d_cross(arg0, arg1, arg2)

	ret2 = WrapPoint3D(unsafe.Pointer(arg2))

	return ret2
}

// Distance computes the distance between the two given #graphene_point3d_t.
func (a *Point3D) Distance(b *Point3D) (delta *Vec3, gfloat float32) {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	arg2 := new(C.graphene_vec3_t)
	var ret2 *Vec3
	var cret C.float
	var goret float32

	cret = C.graphene_point3d_distance(arg0, arg1, arg2)

	ret2 = WrapVec3(unsafe.Pointer(arg2))
	goret = float32(cret)

	return ret2, goret
}

// Dot computes the dot product of the two given #graphene_point3d_t.
func (a *Point3D) Dot(b *Point3D) float32 {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var cret C.float
	var goret float32

	cret = C.graphene_point3d_dot(arg0, arg1)

	goret = float32(cret)

	return goret
}

// Equal checks whether two given points are equal.
func (a *Point3D) Equal(b *Point3D) bool {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var goret bool

	cret = C.graphene_point3d_equal(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Free frees the resources allocated via graphene_point3d_alloc().
func (p *Point3D) Free() {
	var arg0 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	C.graphene_point3d_free(arg0)
}

// Init initializes a #graphene_point3d_t with the given coordinates.
func (p *Point3D) Init(x float32, y float32, z float32) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(z)

	var cret *C.graphene_point3d_t
	var goret *Point3D

	cret = C.graphene_point3d_init(arg0, arg1, arg2, arg3)

	goret = WrapPoint3D(unsafe.Pointer(cret))

	return goret
}

// InitFromPoint initializes a #graphene_point3d_t using the coordinates of
// another #graphene_point3d_t.
func (p *Point3D) InitFromPoint(src *Point3D) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_point3d_t
	var goret *Point3D

	cret = C.graphene_point3d_init_from_point(arg0, arg1)

	goret = WrapPoint3D(unsafe.Pointer(cret))

	return goret
}

// InitFromVec3 initializes a #graphene_point3d_t using the components of a
// #graphene_vec3_t.
func (p *Point3D) InitFromVec3(v *Vec3) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_vec3_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v.Native()))

	var cret *C.graphene_point3d_t
	var goret *Point3D

	cret = C.graphene_point3d_init_from_vec3(arg0, arg1)

	goret = WrapPoint3D(unsafe.Pointer(cret))

	return goret
}

// Interpolate: linearly interpolates each component of @a and @b using the
// provided @factor, and places the result in @res.
func (a *Point3D) Interpolate(b *Point3D, factor float64) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t
	var arg2 C.double

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))
	arg2 = C.double(factor)

	arg3 := new(C.graphene_point3d_t)
	var ret3 *Point3D

	C.graphene_point3d_interpolate(arg0, arg1, arg2, arg3)

	ret3 = WrapPoint3D(unsafe.Pointer(arg3))

	return ret3
}

// Length computes the length of the vector represented by the coordinates of
// the given #graphene_point3d_t.
func (p *Point3D) Length() float32 {
	var arg0 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var cret C.float
	var goret float32

	cret = C.graphene_point3d_length(arg0)

	goret = float32(cret)

	return goret
}

// Near checks whether the two points are near each other, within an @epsilon
// factor.
func (a *Point3D) Near(b *Point3D, epsilon float32) bool {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_point3d_t
	var arg2 C.float

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(b.Native()))
	arg2 = C.float(epsilon)

	var cret C._Bool
	var goret bool

	cret = C.graphene_point3d_near(arg0, arg1, arg2)

	if cret {
		goret = true
	}

	return goret
}

// Normalize computes the normalization of the vector represented by the
// coordinates of the given #graphene_point3d_t.
func (p *Point3D) Normalize() *Point3D {
	var arg0 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	arg1 := new(C.graphene_point3d_t)
	var ret1 *Point3D

	C.graphene_point3d_normalize(arg0, arg1)

	ret1 = WrapPoint3D(unsafe.Pointer(arg1))

	return ret1
}

// NormalizeViewport normalizes the coordinates of a #graphene_point3d_t using
// the given viewport and clipping planes.
//
// The coordinates of the resulting #graphene_point3d_t will be in the [ -1, 1 ]
// range.
func (p *Point3D) NormalizeViewport(viewport *Rect, zNear float32, zFar float32) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 *C.graphene_rect_t
	var arg2 C.float
	var arg3 C.float

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(viewport.Native()))
	arg2 = C.float(zNear)
	arg3 = C.float(zFar)

	arg4 := new(C.graphene_point3d_t)
	var ret4 *Point3D

	C.graphene_point3d_normalize_viewport(arg0, arg1, arg2, arg3, arg4)

	ret4 = WrapPoint3D(unsafe.Pointer(arg4))

	return ret4
}

// Scale scales the coordinates of the given #graphene_point3d_t by the given
// @factor.
func (p *Point3D) Scale(factor float32) *Point3D {
	var arg0 *C.graphene_point3d_t
	var arg1 C.float

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))
	arg1 = C.float(factor)

	arg2 := new(C.graphene_point3d_t)
	var ret2 *Point3D

	C.graphene_point3d_scale(arg0, arg1, arg2)

	ret2 = WrapPoint3D(unsafe.Pointer(arg2))

	return ret2
}

// ToVec3 stores the coordinates of a #graphene_point3d_t into a
// #graphene_vec3_t.
func (p *Point3D) ToVec3() *Vec3 {
	var arg0 *C.graphene_point3d_t

	arg0 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	arg1 := new(C.graphene_vec3_t)
	var ret1 *Vec3

	C.graphene_point3d_to_vec3(arg0, arg1)

	ret1 = WrapVec3(unsafe.Pointer(arg1))

	return ret1
}
