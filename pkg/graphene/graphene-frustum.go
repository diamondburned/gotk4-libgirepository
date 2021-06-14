// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 graphene-1.0 graphene-gobject-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_frustum_get_type()), F: marshalFrustum},
	})
}

// Frustum: a 3D volume delimited by 2D clip planes.
//
// The contents of the `graphene_frustum_t` are private, and should not be
// modified directly.
type Frustum struct {
	native C.graphene_frustum_t
}

// WrapFrustum wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFrustum(ptr unsafe.Pointer) *Frustum {
	if ptr == nil {
		return nil
	}

	return (*Frustum)(ptr)
}

func marshalFrustum(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFrustum(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *Frustum) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// ContainsPoint checks whether a point is inside the volume defined by the
// given #graphene_frustum_t.
func (f *Frustum) ContainsPoint(point *Point3D) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_frustum_contains_point(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Equal checks whether the two given #graphene_frustum_t are equal.
func (a *Frustum) Equal(b *Frustum) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_frustum_t // out

	_arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_frustum_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_frustum_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_frustum_alloc().
func (f *Frustum) Free() {
	var _arg0 *C.graphene_frustum_t // out

	_arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))

	C.graphene_frustum_free(_arg0)
}

// IntersectsBox checks whether the given @box intersects a plane of a
// #graphene_frustum_t.
func (f *Frustum) IntersectsBox(box *Box) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_box_t     // out

	_arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(box.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_frustum_intersects_box(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsSphere checks whether the given @sphere intersects a plane of a
// #graphene_frustum_t.
func (f *Frustum) IntersectsSphere(sphere *Sphere) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_sphere_t  // out

	_arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(sphere.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_frustum_intersects_sphere(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
