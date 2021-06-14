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
		{T: externglib.Type(C.graphene_matrix_get_type()), F: marshalMatrix},
	})
}

// Matrix: a structure capable of holding a 4x4 matrix.
//
// The contents of the #graphene_matrix_t structure are private and should never
// be accessed directly.
type Matrix struct {
	native C.graphene_matrix_t
}

// WrapMatrix wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMatrix(ptr unsafe.Pointer) *Matrix {
	if ptr == nil {
		return nil
	}

	return (*Matrix)(ptr)
}

func marshalMatrix(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMatrix(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *Matrix) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Decompose decomposes a transformation matrix into its component
// transformations.
//
// The algorithm for decomposing a matrix is taken from the CSS3 Transforms
// specification (http://dev.w3.org/csswg/css-transforms/); specifically, the
// decomposition code is based on the equivalent code published in "Graphics
// Gems II", edited by Jim Arvo, and available online
// (http://tog.acm.org/resources/GraphicsGems/gemsii/unmatrix.c).
func (m *Matrix) Decompose() (translate Vec3, scale Vec3, rotate Quaternion, shear Vec3, perspective Vec4, ok bool) {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _translate Vec3
	var _scale Vec3
	var _rotate Quaternion
	var _shear Vec3
	var _perspective Vec4
	var _cret C._Bool // in

	_cret = C.graphene_matrix_decompose(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_translate)), (*C.graphene_vec3_t)(unsafe.Pointer(&_scale)), (*C.graphene_quaternion_t)(unsafe.Pointer(&_rotate)), (*C.graphene_vec3_t)(unsafe.Pointer(&_shear)), (*C.graphene_vec4_t)(unsafe.Pointer(&_perspective)))

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _translate, _scale, _rotate, _shear, _perspective, _ok
}

// Determinant computes the determinant of the given matrix.
func (m *Matrix) Determinant() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_determinant(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_matrix_t matrices are equal.
func (a *Matrix) Equal(b *Matrix) bool {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// EqualFast checks whether the two given #graphene_matrix_t matrices are
// byte-by-byte equal.
//
// While this function is faster than graphene_matrix_equal(), it can also
// return false negatives, so it should be used in conjuction with either
// graphene_matrix_equal() or graphene_matrix_near(). For instance:
//
//    if (graphene_matrix_equal_fast (a, b))
//      {
//        // matrices are definitely the same
//      }
//    else
//      {
//        if (graphene_matrix_equal (a, b))
//          // matrices contain the same values within an epsilon of FLT_EPSILON
//        else if (graphene_matrix_near (a, b, 0.0001))
//          // matrices contain the same values within an epsilon of 0.0001
//        else
//          // matrices are not equal
//      }
func (a *Matrix) EqualFast(b *Matrix) bool {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(b.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_equal_fast(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_matrix_alloc().
func (m *Matrix) Free() {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	C.graphene_matrix_free(_arg0)
}

// Row retrieves the given row vector at @index_ inside a matrix.
func (m *Matrix) Row(index_ uint) Vec4 {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.uint               // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.uint(index_)

	var _res Vec4

	C.graphene_matrix_get_row(_arg0, _arg1, (*C.graphene_vec4_t)(unsafe.Pointer(&_res)))

	return _res
}

// Value retrieves the value at the given @row and @col index.
func (m *Matrix) Value(row uint, col uint) float32 {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.uint               // out
	var _arg2 C.uint               // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.uint(row)
	_arg2 = C.uint(col)

	var _cret C.float // in

	_cret = C.graphene_matrix_get_value(_arg0, _arg1, _arg2)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// XScale retrieves the scaling factor on the X axis in @m.
func (m *Matrix) XScale() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_x_scale(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// XTranslation retrieves the translation component on the X axis from @m.
func (m *Matrix) XTranslation() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_x_translation(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// YScale retrieves the scaling factor on the Y axis in @m.
func (m *Matrix) YScale() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_y_scale(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// YTranslation retrieves the translation component on the Y axis from @m.
func (m *Matrix) YTranslation() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_y_translation(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// ZScale retrieves the scaling factor on the Z axis in @m.
func (m *Matrix) ZScale() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_z_scale(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// ZTranslation retrieves the translation component on the Z axis from @m.
func (m *Matrix) ZTranslation() float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C.float // in

	_cret = C.graphene_matrix_get_z_translation(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Interpolate: linearly interpolates the two given #graphene_matrix_t by
// interpolating the decomposed transformations separately.
//
// If either matrix cannot be reduced to their transformations then the
// interpolation cannot be performed, and this function will return an identity
// matrix.
func (a *Matrix) Interpolate(b *Matrix, factor float64) Matrix {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 C.double             // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.double(factor)

	var _res Matrix

	C.graphene_matrix_interpolate(_arg0, _arg1, _arg2, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	return _res
}

// Inverse inverts the given matrix.
func (m *Matrix) Inverse() (Matrix, bool) {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _res Matrix
	var _cret C._Bool // in

	_cret = C.graphene_matrix_inverse(_arg0, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _res, _ok
}

// Is2D checks whether the given #graphene_matrix_t is compatible with an a 2D
// affine transformation matrix.
func (m *Matrix) Is2D() bool {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_is_2d(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IsBackfaceVisible checks whether a #graphene_matrix_t has a visible back
// face.
func (m *Matrix) IsBackfaceVisible() bool {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_is_backface_visible(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IsIdentity checks whether the given #graphene_matrix_t is the identity
// matrix.
func (m *Matrix) IsIdentity() bool {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_is_identity(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IsSingular checks whether a matrix is singular.
func (m *Matrix) IsSingular() bool {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _cret C._Bool // in

	_cret = C.graphene_matrix_is_singular(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Multiply multiplies two #graphene_matrix_t.
//
// Matrix multiplication is not commutative in general; the order of the factors
// matters. The product of this multiplication is (@a × @b)
func (a *Matrix) Multiply(b *Matrix) Matrix {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(b.Native()))

	var _res Matrix

	C.graphene_matrix_multiply(_arg0, _arg1, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	return _res
}

// Near compares the two given #graphene_matrix_t matrices and checks whether
// their values are within the given @epsilon of each other.
func (a *Matrix) Near(b *Matrix, epsilon float32) bool {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.float(epsilon)

	var _cret C._Bool // in

	_cret = C.graphene_matrix_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Normalize normalizes the given #graphene_matrix_t.
func (m *Matrix) Normalize() Matrix {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _res Matrix

	C.graphene_matrix_normalize(_arg0, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	return _res
}

// Perspective applies a perspective of @depth to the matrix.
func (m *Matrix) Perspective(depth float32) Matrix {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(depth)

	var _res Matrix

	C.graphene_matrix_perspective(_arg0, _arg1, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	return _res
}

// Print prints the contents of a matrix to the standard error stream.
//
// This function is only useful for debugging; there are no guarantees made on
// the format of the output.
func (m *Matrix) Print() {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	C.graphene_matrix_print(_arg0)
}

// ProjectPoint projects a #graphene_point_t using the matrix @m.
func (m *Matrix) ProjectPoint(p *Point) Point {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_point_t  // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var _res Point

	C.graphene_matrix_project_point(_arg0, _arg1, (*C.graphene_point_t)(unsafe.Pointer(&_res)))

	return _res
}

// ProjectRect projects all corners of a #graphene_rect_t using the given
// matrix.
//
// See also: graphene_matrix_project_point()
func (m *Matrix) ProjectRect(r *Rect) Quad {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var _res Quad

	C.graphene_matrix_project_rect(_arg0, _arg1, (*C.graphene_quad_t)(unsafe.Pointer(&_res)))

	return _res
}

// ProjectRectBounds projects a #graphene_rect_t using the given matrix.
//
// The resulting rectangle is the axis aligned bounding rectangle capable of
// fully containing the projected rectangle.
func (m *Matrix) ProjectRectBounds(r *Rect) Rect {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var _res Rect

	C.graphene_matrix_project_rect_bounds(_arg0, _arg1, (*C.graphene_rect_t)(unsafe.Pointer(&_res)))

	return _res
}

// Rotate adds a rotation transformation to @m, using the given @angle and @axis
// vector.
//
// This is the equivalent of calling graphene_matrix_init_rotate() and then
// multiplying the matrix @m with the rotation matrix.
func (m *Matrix) Rotate(angle float32, axis *Vec3) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out
	var _arg2 *C.graphene_vec3_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(angle)
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(axis.Native()))

	C.graphene_matrix_rotate(_arg0, _arg1, _arg2)
}

// RotateEuler adds a rotation transformation to @m, using the given
// #graphene_euler_t.
func (m *Matrix) RotateEuler(e *Euler) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_euler_t  // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_euler_t)(unsafe.Pointer(e.Native()))

	C.graphene_matrix_rotate_euler(_arg0, _arg1)
}

// RotateQuaternion adds a rotation transformation to @m, using the given
// #graphene_quaternion_t.
//
// This is the equivalent of calling graphene_quaternion_to_matrix() and then
// multiplying @m with the rotation matrix.
func (m *Matrix) RotateQuaternion(q *Quaternion) {
	var _arg0 *C.graphene_matrix_t     // out
	var _arg1 *C.graphene_quaternion_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_quaternion_t)(unsafe.Pointer(q.Native()))

	C.graphene_matrix_rotate_quaternion(_arg0, _arg1)
}

// RotateX adds a rotation transformation around the X axis to @m, using the
// given @angle.
//
// See also: graphene_matrix_rotate()
func (m *Matrix) RotateX(angle float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(angle)

	C.graphene_matrix_rotate_x(_arg0, _arg1)
}

// RotateY adds a rotation transformation around the Y axis to @m, using the
// given @angle.
//
// See also: graphene_matrix_rotate()
func (m *Matrix) RotateY(angle float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(angle)

	C.graphene_matrix_rotate_y(_arg0, _arg1)
}

// RotateZ adds a rotation transformation around the Z axis to @m, using the
// given @angle.
//
// See also: graphene_matrix_rotate()
func (m *Matrix) RotateZ(angle float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(angle)

	C.graphene_matrix_rotate_z(_arg0, _arg1)
}

// Scale adds a scaling transformation to @m, using the three given factors.
//
// This is the equivalent of calling graphene_matrix_init_scale() and then
// multiplying the matrix @m with the scale matrix.
func (m *Matrix) Scale(factorX float32, factorY float32, factorZ float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out
	var _arg2 C.float              // out
	var _arg3 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)
	_arg3 = C.float(factorZ)

	C.graphene_matrix_scale(_arg0, _arg1, _arg2, _arg3)
}

// SkewXY adds a skew of @factor on the X and Y axis to the given matrix.
func (m *Matrix) SkewXY(factor float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(factor)

	C.graphene_matrix_skew_xy(_arg0, _arg1)
}

// SkewXZ adds a skew of @factor on the X and Z axis to the given matrix.
func (m *Matrix) SkewXZ(factor float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(factor)

	C.graphene_matrix_skew_xz(_arg0, _arg1)
}

// SkewYZ adds a skew of @factor on the Y and Z axis to the given matrix.
func (m *Matrix) SkewYZ(factor float32) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 C.float              // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = C.float(factor)

	C.graphene_matrix_skew_yz(_arg0, _arg1)
}

// To2D converts a #graphene_matrix_t to an affine transformation matrix, if the
// given matrix is compatible.
//
// The returned values have the following layout:
//
//    ⎛ xx  yx ⎞   ⎛  a   b  0 ⎞
//    ⎜ xy  yy ⎟ = ⎜  c   d  0 ⎟
//    ⎝ x0  y0 ⎠   ⎝ tx  ty  1 ⎠
//
// This function can be used to convert between a #graphene_matrix_t and an
// affine matrix type from other libraries.
func (m *Matrix) To2D() (xx float64, yx float64, xy float64, yy float64, x0 float64, y0 float64, ok bool) {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _arg1 C.double // in
	var _arg2 C.double // in
	var _arg3 C.double // in
	var _arg4 C.double // in
	var _arg5 C.double // in
	var _arg6 C.double // in
	var _cret C._Bool  // in

	_cret = C.graphene_matrix_to_2d(_arg0, &_arg1, &_arg2, &_arg3, &_arg4, &_arg5, &_arg6)

	var _xx float64 // out
	var _yx float64 // out
	var _xy float64 // out
	var _yy float64 // out
	var _x0 float64 // out
	var _y0 float64 // out
	var _ok bool    // out

	_xx = (float64)(_arg1)
	_yx = (float64)(_arg2)
	_xy = (float64)(_arg3)
	_yy = (float64)(_arg4)
	_x0 = (float64)(_arg5)
	_y0 = (float64)(_arg6)
	if _cret {
		_ok = true
	}

	return _xx, _yx, _xy, _yy, _x0, _y0, _ok
}

// ToFloat converts a #graphene_matrix_t to an array of floating point values.
func (m *Matrix) ToFloat() [16]float32 {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _arg1 [16]C.float

	C.graphene_matrix_to_float(_arg0, &_arg1[0])

	var _v [16]float32

	_v = *(*[16]float32)(unsafe.Pointer(&_arg1))

	return _v
}

// TransformBounds transforms each corner of a #graphene_rect_t using the given
// matrix @m.
//
// The result is the axis aligned bounding rectangle containing the coplanar
// quadrilateral.
//
// See also: graphene_matrix_transform_point()
func (m *Matrix) TransformBounds(r *Rect) Rect {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var _res Rect

	C.graphene_matrix_transform_bounds(_arg0, _arg1, (*C.graphene_rect_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformBox transforms the vertices of a #graphene_box_t using the given
// matrix @m.
//
// The result is the axis aligned bounding box containing the transformed
// vertices.
func (m *Matrix) TransformBox(b *Box) Box {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_box_t    // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	var _res Box

	C.graphene_matrix_transform_box(_arg0, _arg1, (*C.graphene_box_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformPoint transforms the given #graphene_point_t using the matrix @m.
//
// Unlike graphene_matrix_transform_vec3(), this function will take into account
// the fourth row vector of the #graphene_matrix_t when computing the dot
// product of each row vector of the matrix.
//
// See also: graphene_simd4x4f_point3_mul()
func (m *Matrix) TransformPoint(p *Point) Point {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_point_t  // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var _res Point

	C.graphene_matrix_transform_point(_arg0, _arg1, (*C.graphene_point_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformPoint3D transforms the given #graphene_point3d_t using the matrix
// @m.
//
// Unlike graphene_matrix_transform_vec3(), this function will take into account
// the fourth row vector of the #graphene_matrix_t when computing the dot
// product of each row vector of the matrix.
//
// See also: graphene_simd4x4f_point3_mul()
func (m *Matrix) TransformPoint3D(p *Point3D) Point3D {
	var _arg0 *C.graphene_matrix_t  // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	var _res Point3D

	C.graphene_matrix_transform_point3d(_arg0, _arg1, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformRay: transform a #graphene_ray_t using the given matrix @m.
func (m *Matrix) TransformRay(r *Ray) Ray {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_ray_t    // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	var _res Ray

	C.graphene_matrix_transform_ray(_arg0, _arg1, (*C.graphene_ray_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformRect transforms each corner of a #graphene_rect_t using the given
// matrix @m.
//
// The result is a coplanar quadrilateral.
//
// See also: graphene_matrix_transform_point()
func (m *Matrix) TransformRect(r *Rect) Quad {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var _res Quad

	C.graphene_matrix_transform_rect(_arg0, _arg1, (*C.graphene_quad_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformSphere transforms a #graphene_sphere_t using the given matrix @m.
// The result is the bounding sphere containing the transformed sphere.
func (m *Matrix) TransformSphere(s *Sphere) Sphere {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_sphere_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	var _res Sphere

	C.graphene_matrix_transform_sphere(_arg0, _arg1, (*C.graphene_sphere_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformVec3 transforms the given #graphene_vec3_t using the matrix @m.
//
// This function will multiply the X, Y, and Z row vectors of the matrix @m with
// the corresponding components of the vector @v. The W row vector will be
// ignored.
//
// See also: graphene_simd4x4f_vec3_mul()
func (m *Matrix) TransformVec3(v *Vec3) Vec3 {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_vec3_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v.Native()))

	var _res Vec3

	C.graphene_matrix_transform_vec3(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// TransformVec4 transforms the given #graphene_vec4_t using the matrix @m.
//
// See also: graphene_simd4x4f_vec4_mul()
func (m *Matrix) TransformVec4(v *Vec4) Vec4 {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_vec4_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_vec4_t)(unsafe.Pointer(v.Native()))

	var _res Vec4

	C.graphene_matrix_transform_vec4(_arg0, _arg1, (*C.graphene_vec4_t)(unsafe.Pointer(&_res)))

	return _res
}

// Translate adds a translation transformation to @m using the coordinates of
// the given #graphene_point3d_t.
//
// This is the equivalent of calling graphene_matrix_init_translate() and then
// multiplying @m with the translation matrix.
func (m *Matrix) Translate(pos *Point3D) {
	var _arg0 *C.graphene_matrix_t  // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(pos.Native()))

	C.graphene_matrix_translate(_arg0, _arg1)
}

// Transpose transposes the given matrix.
func (m *Matrix) Transpose() Matrix {
	var _arg0 *C.graphene_matrix_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))

	var _res Matrix

	C.graphene_matrix_transpose(_arg0, (*C.graphene_matrix_t)(unsafe.Pointer(&_res)))

	return _res
}

// UnprojectPoint3D unprojects the given @point using the @projection matrix and
// a @modelview matrix.
func (p *Matrix) UnprojectPoint3D(modelview *Matrix, point *Point3D) Point3D {
	var _arg0 *C.graphene_matrix_t  // out
	var _arg1 *C.graphene_matrix_t  // out
	var _arg2 *C.graphene_point3d_t // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(modelview.Native()))
	_arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var _res Point3D

	C.graphene_matrix_unproject_point3d(_arg0, _arg1, _arg2, (*C.graphene_point3d_t)(unsafe.Pointer(&_res)))

	return _res
}

// UntransformBounds undoes the transformation on the corners of a
// #graphene_rect_t using the given matrix, within the given axis aligned
// rectangular @bounds.
func (m *Matrix) UntransformBounds(r *Rect, bounds *Rect) Rect {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_rect_t   // out
	var _arg2 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	var _res Rect

	C.graphene_matrix_untransform_bounds(_arg0, _arg1, _arg2, (*C.graphene_rect_t)(unsafe.Pointer(&_res)))

	return _res
}

// UntransformPoint undoes the transformation of a #graphene_point_t using the
// given matrix, within the given axis aligned rectangular @bounds.
func (m *Matrix) UntransformPoint(p *Point, bounds *Rect) (Point, bool) {
	var _arg0 *C.graphene_matrix_t // out
	var _arg1 *C.graphene_point_t  // out
	var _arg2 *C.graphene_rect_t   // out

	_arg0 = (*C.graphene_matrix_t)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	var _res Point
	var _cret C._Bool // in

	_cret = C.graphene_matrix_untransform_point(_arg0, _arg1, _arg2, (*C.graphene_point_t)(unsafe.Pointer(&_res)))

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _res, _ok
}
