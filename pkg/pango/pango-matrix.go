// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_matrix_get_type()), F: marshalMatrix},
	})
}

// Matrix: `PangoMatrix` specifies a transformation between user-space and
// device coordinates.
//
//
// The transformation is given by
//
// “` x_device = x_user * matrix->xx + y_user * matrix->xy + matrix->x0;
// y_device = x_user * matrix->yx + y_user * matrix->yy + matrix->y0; “`
type Matrix struct {
	native C.PangoMatrix
}

func marshalMatrix(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Matrix)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *Matrix) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Concat changes the transformation represented by @matrix to be the
// transformation given by first applying transformation given by @new_matrix
// then applying the original transformation.
func (matrix *Matrix) Concat(newMatrix *Matrix) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 *C.PangoMatrix // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = (*C.PangoMatrix)(unsafe.Pointer(newMatrix))

	C.pango_matrix_concat(_arg0, _arg1)
}

// Copy copies a `PangoMatrix`.
func (matrix *Matrix) Copy() *Matrix {
	var _arg0 *C.PangoMatrix // out
	var _cret *C.PangoMatrix // in

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	_cret = C.pango_matrix_copy(_arg0)

	var _ret *Matrix // out

	_ret = (*Matrix)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Matrix) {
		C.pango_matrix_free((*C.PangoMatrix)(unsafe.Pointer(v)))
	})

	return _ret
}

// Free: free a `PangoMatrix`.
func (matrix *Matrix) free() {
	var _arg0 *C.PangoMatrix // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	C.pango_matrix_free(_arg0)
}

// FontScaleFactor returns the scale factor of a matrix on the height of the
// font.
//
// That is, the scale factor in the direction perpendicular to the vector that
// the X coordinate is mapped to. If the scale in the X coordinate is needed as
// well, use [method@Pango.Matrix.get_font_scale_factors].
func (matrix *Matrix) FontScaleFactor() float64 {
	var _arg0 *C.PangoMatrix // out
	var _cret C.double       // in

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	_cret = C.pango_matrix_get_font_scale_factor(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// FontScaleFactors calculates the scale factor of a matrix on the width and
// height of the font.
//
// That is, @xscale is the scale factor in the direction of the X coordinate,
// and @yscale is the scale factor in the direction perpendicular to the vector
// that the X coordinate is mapped to.
//
// Note that output numbers will always be non-negative.
func (matrix *Matrix) FontScaleFactors() (xscale float64, yscale float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // in
	var _arg2 C.double       // in

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	C.pango_matrix_get_font_scale_factors(_arg0, &_arg1, &_arg2)

	var _xscale float64 // out
	var _yscale float64 // out

	_xscale = float64(_arg1)
	_yscale = float64(_arg2)

	return _xscale, _yscale
}

// Rotate changes the transformation represented by @matrix to be the
// transformation given by first rotating by @degrees degrees counter-clockwise
// then applying the original transformation.
func (matrix *Matrix) Rotate(degrees float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = C.double(degrees)

	C.pango_matrix_rotate(_arg0, _arg1)
}

// Scale changes the transformation represented by @matrix to be the
// transformation given by first scaling by @sx in the X direction and @sy in
// the Y direction then applying the original transformation.
func (matrix *Matrix) Scale(scaleX float64, scaleY float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out
	var _arg2 C.double       // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = C.double(scaleX)
	_arg2 = C.double(scaleY)

	C.pango_matrix_scale(_arg0, _arg1, _arg2)
}

// TransformDistance transforms the distance vector (@dx,@dy) by @matrix.
//
// This is similar to [method@Pango.Matrix.transform_point], except that the
// translation components of the transformation are ignored. The calculation of
// the returned vector is as follows:
//
// “` dx2 = dx1 * xx + dy1 * xy; dy2 = dx1 * yx + dy1 * yy; “`
//
// Affine transformations are position invariant, so the same vector always
// transforms to the same vector. If (@x1,@y1) transforms to (@x2,@y2) then
// (@x1+@dx1,@y1+@dy1) will transform to (@x1+@dx2,@y1+@dy2) for all values of
// @x1 and @x2.
func (matrix *Matrix) TransformDistance(dx *float64, dy *float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 *C.double      // out
	var _arg2 *C.double      // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = (*C.double)(unsafe.Pointer(dx))
	_arg2 = (*C.double)(unsafe.Pointer(dy))

	C.pango_matrix_transform_distance(_arg0, _arg1, _arg2)
}

// TransformPixelRectangle: first transforms the @rect using @matrix, then
// calculates the bounding box of the transformed rectangle.
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image should
// be and how much you should shift the layout when rendering.
//
// For better accuracy, you should use [method@Pango.Matrix.transform_rectangle]
// on original rectangle in Pango units and convert to pixels afterward using
// [func@extents_to_pixels]'s first argument.
func (matrix *Matrix) TransformPixelRectangle(rect *Rectangle) {
	var _arg0 *C.PangoMatrix    // out
	var _arg1 *C.PangoRectangle // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = (*C.PangoRectangle)(unsafe.Pointer(rect))

	C.pango_matrix_transform_pixel_rectangle(_arg0, _arg1)
}

// TransformPoint transforms the point (@x, @y) by @matrix.
func (matrix *Matrix) TransformPoint(x *float64, y *float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 *C.double      // out
	var _arg2 *C.double      // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = (*C.double)(unsafe.Pointer(x))
	_arg2 = (*C.double)(unsafe.Pointer(y))

	C.pango_matrix_transform_point(_arg0, _arg1, _arg2)
}

// TransformRectangle: first transforms @rect using @matrix, then calculates the
// bounding box of the transformed rectangle.
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image should
// be and how much you should shift the layout when rendering.
//
// If you have a rectangle in device units (pixels), use
// [method@Pango.Matrix.transform_pixel_rectangle].
//
// If you have the rectangle in Pango units and want to convert to transformed
// pixel bounding box, it is more accurate to transform it first (using this
// function) and pass the result to pango_extents_to_pixels(), first argument,
// for an inclusive rounded rectangle. However, there are valid reasons that you
// may want to convert to pixels first and then transform, for example when the
// transformed coordinates may overflow in Pango units (large matrix translation
// for example).
func (matrix *Matrix) TransformRectangle(rect *Rectangle) {
	var _arg0 *C.PangoMatrix    // out
	var _arg1 *C.PangoRectangle // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = (*C.PangoRectangle)(unsafe.Pointer(rect))

	C.pango_matrix_transform_rectangle(_arg0, _arg1)
}

// Translate changes the transformation represented by @matrix to be the
// transformation given by first translating by (@tx, @ty) then applying the
// original transformation.
func (matrix *Matrix) Translate(tx float64, ty float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out
	var _arg2 C.double       // out

	_arg0 = (*C.PangoMatrix)(unsafe.Pointer(matrix))
	_arg1 = C.double(tx)
	_arg2 = C.double(ty)

	C.pango_matrix_translate(_arg0, _arg1, _arg2)
}
