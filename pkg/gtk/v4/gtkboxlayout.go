// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_box_layout_get_type()), F: marshalBoxLayout},
	})
}

// BoxLayout: a GtkBoxLayout is a layout manager that arranges the children of
// any widget using it into a single row or column, depending on the value of
// its Orientable:orientation property. Within the other dimension all children
// all allocated the same size. The GtkBoxLayout will respect the Widget:halign
// and Widget:valign properties of each child widget.
//
// If you want all children to be assigned the same size, you can use the
// BoxLayout:homogeneous property.
//
// If you want to specify the amount of space placed between each child, you can
// use the BoxLayout:spacing property.
type BoxLayout interface {
	LayoutManager
	Orientable

	// BaselinePosition gets the value set by
	// gtk_box_layout_set_baseline_position().
	BaselinePosition() BaselinePosition
	// Homogeneous returns whether the layout is set to be homogeneous.
	Homogeneous() bool
	// Spacing returns the space that @box_layout puts between children.
	Spacing() uint
	// SetBaselinePosition sets the baseline position of a box layout.
	//
	// The baseline position affects only horizontal boxes with at least one
	// baseline aligned child. If there is more vertical space available than
	// requested, and the baseline is not allocated by the parent then the given
	// @position is used to allocate the baseline within the extra space
	// available.
	SetBaselinePosition(position BaselinePosition)
	// SetHomogeneous sets whether the box layout will allocate the same size to
	// all children.
	SetHomogeneous(homogeneous bool)
	// SetSpacing sets how much spacing to put between children.
	SetSpacing(spacing uint)
}

// boxLayout implements the BoxLayout interface.
type boxLayout struct {
	LayoutManager
	Orientable
}

var _ BoxLayout = (*boxLayout)(nil)

// WrapBoxLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapBoxLayout(obj *externglib.Object) BoxLayout {
	return BoxLayout{
		LayoutManager: WrapLayoutManager(obj),
		Orientable:    WrapOrientable(obj),
	}
}

func marshalBoxLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBoxLayout(obj), nil
}

// NewBoxLayout constructs a class BoxLayout.
func NewBoxLayout(orientation Orientation) BoxLayout {
	var _arg1 C.GtkOrientation

	_arg1 = (C.GtkOrientation)(orientation)

	var _cret C.GtkBoxLayout

	cret = C.gtk_box_layout_new(_arg1)

	var _boxLayout BoxLayout

	_boxLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(BoxLayout)

	return _boxLayout
}

// BaselinePosition gets the value set by
// gtk_box_layout_set_baseline_position().
func (b boxLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBoxLayout

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	var _cret C.GtkBaselinePosition

	cret = C.gtk_box_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// Homogeneous returns whether the layout is set to be homogeneous.
func (b boxLayout) Homogeneous() bool {
	var _arg0 *C.GtkBoxLayout

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_box_layout_get_homogeneous(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Spacing returns the space that @box_layout puts between children.
func (b boxLayout) Spacing() uint {
	var _arg0 *C.GtkBoxLayout

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	var _cret C.guint

	cret = C.gtk_box_layout_get_spacing(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// SetBaselinePosition sets the baseline position of a box layout.
//
// The baseline position affects only horizontal boxes with at least one
// baseline aligned child. If there is more vertical space available than
// requested, and the baseline is not allocated by the parent then the given
// @position is used to allocate the baseline within the extra space
// available.
func (b boxLayout) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBoxLayout
	var _arg1 C.GtkBaselinePosition

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = (C.GtkBaselinePosition)(position)

	C.gtk_box_layout_set_baseline_position(_arg0, _arg1)
}

// SetHomogeneous sets whether the box layout will allocate the same size to
// all children.
func (b boxLayout) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBoxLayout
	var _arg1 C.gboolean

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.gboolean(1)
	}

	C.gtk_box_layout_set_homogeneous(_arg0, _arg1)
}

// SetSpacing sets how much spacing to put between children.
func (b boxLayout) SetSpacing(spacing uint) {
	var _arg0 *C.GtkBoxLayout
	var _arg1 C.guint

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_box_layout_set_spacing(_arg0, _arg1)
}
