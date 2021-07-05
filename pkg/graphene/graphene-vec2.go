// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_vec2_get_type()), F: marshalVec2},
	})
}

// Vec2: structure capable of holding a vector with two dimensions, x and y.
//
// The contents of the #graphene_vec2_t structure are private and should never
// be accessed directly.
type Vec2 struct {
	native C.graphene_vec2_t
}

// WrapVec2 wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapVec2(ptr unsafe.Pointer) *Vec2 {
	return (*Vec2)(ptr)
}

func marshalVec2(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Vec2)(unsafe.Pointer(b)), nil
}

// NewVec2Alloc constructs a struct Vec2.
func NewVec2Alloc() *Vec2 {
	var _cret *C.graphene_vec2_t // in

	_cret = C.graphene_vec2_alloc()

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_vec2, func(v *Vec2) {
		C.free(unsafe.Pointer(v))
	})

	return _vec2
}

// Native returns the underlying C source pointer.
func (v *Vec2) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// Add adds each component of the two passed vectors and places each result into
// the components of @res.
func (a *Vec2) Add(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_add(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Divide divides each component of the first operand @a by the corresponding
// component of the second operand @b, and places the results into the vector
// @res.
func (a *Vec2) Divide(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_divide(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Dot computes the dot product of the two given vectors.
func (a *Vec2) Dot(b *Vec2) float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	_cret = C.graphene_vec2_dot(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_vec2_t are equal.
func (v *Vec2) Equal(v2 *Vec2) bool {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(v2))

	_cret = C.graphene_vec2_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by @v
func (v *Vec2) Free() {
	var _arg0 *C.graphene_vec2_t // out

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	C.graphene_vec2_free(_arg0)
}

// X retrieves the X component of the #graphene_vec2_t.
func (v *Vec2) X() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec2_get_x(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Y retrieves the Y component of the #graphene_vec2_t.
func (v *Vec2) Y() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec2_get_y(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_vec2_t using the given values.
//
// This function can be called multiple times.
func (v *Vec2) Init(x float32, y float32) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = C.float(x)
	_arg2 = C.float(y)

	_cret = C.graphene_vec2_init(_arg0, _arg1, _arg2)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(unsafe.Pointer(_cret))

	return _vec2
}

// InitFromFloat initializes @v with the contents of the given array.
func (v *Vec2) InitFromFloat(src [2]float32) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.float
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = (*C.float)(unsafe.Pointer(&src))

	_cret = C.graphene_vec2_init_from_float(_arg0, _arg1)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(unsafe.Pointer(_cret))

	return _vec2
}

// InitFromVec2 copies the contents of @src into @v.
func (v *Vec2) InitFromVec2(src *Vec2) *Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _cret *C.graphene_vec2_t // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(src))

	_cret = C.graphene_vec2_init_from_vec2(_arg0, _arg1)

	var _vec2 *Vec2 // out

	_vec2 = (*Vec2)(unsafe.Pointer(_cret))

	return _vec2
}

// Interpolate: linearly interpolates @v1 and @v2 using the given @factor.
func (v *Vec2) Interpolate(v2 *Vec2, factor float64) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(v2))
	_arg2 = C.double(factor)

	C.graphene_vec2_interpolate(_arg0, _arg1, _arg2, &_arg3)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Length computes the length of the given vector.
func (v *Vec2) Length() float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec2_length(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Max compares the two given vectors and places the maximum values of each
// component into @res.
func (a *Vec2) Max(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_max(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Min compares the two given vectors and places the minimum values of each
// component into @res.
func (a *Vec2) Min(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_min(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Multiply multiplies each component of the two passed vectors and places each
// result into the components of @res.
func (a *Vec2) Multiply(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_multiply(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Near compares the two given #graphene_vec2_t vectors and checks whether their
// values are within the given @epsilon.
func (v *Vec2) Near(v2 *Vec2, epsilon float32) bool {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.float            // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(v2))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_vec2_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Negate negates the given #graphene_vec2_t.
func (v *Vec2) Negate() Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	C.graphene_vec2_negate(_arg0, &_arg1)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Normalize computes the normalized vector for the given vector @v.
func (v *Vec2) Normalize() Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	C.graphene_vec2_normalize(_arg0, &_arg1)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Scale multiplies all components of the given vector with the given scalar
// @factor.
func (v *Vec2) Scale(factor float32) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 C.float            // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))
	_arg1 = C.float(factor)

	C.graphene_vec2_scale(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Subtract subtracts from each component of the first operand @a the
// corresponding component of the second operand @b and places each result into
// the components of @res.
func (a *Vec2) Subtract(b *Vec2) Vec2 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(b))

	C.graphene_vec2_subtract(_arg0, _arg1, &_arg2)

	var _res Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// ToFloat stores the components of @v into an array.
func (v *Vec2) ToFloat() [2]float32 {
	var _arg0 *C.graphene_vec2_t // out
	var _arg1 [2]C.float

	_arg0 = (*C.graphene_vec2_t)(unsafe.Pointer(v))

	C.graphene_vec2_to_float(_arg0, &_arg1[0])

	var _dest [2]float32

	_dest = *(*[2]float32)(unsafe.Pointer(&_arg1))

	return _dest
}
