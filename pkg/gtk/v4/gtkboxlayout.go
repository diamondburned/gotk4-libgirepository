// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_box_layout_get_type()), F: marshalBoxLayout},
	})
}

// BoxLayout: `GtkBoxLayout` is a layout manager that arranges children in a
// single row or column.
//
// Whether it is a row or column depends on the value of its
// [property@Gtk.Orientable:orientation] property. Within the other dimension
// all children all allocated the same size. The `GtkBoxLayout` will respect the
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties of
// each child widget.
//
// If you want all children to be assigned the same size, you can use the
// [property@Gtk.BoxLayout:homogeneous] property.
//
// If you want to specify the amount of space placed between each child, you can
// use the [property@Gtk.BoxLayout:spacing] property.
type BoxLayout interface {
	gextras.Objector

	// AsLayoutManager casts the class to the LayoutManager interface.
	AsLayoutManager() LayoutManager
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// Allocate assigns the given @width, @height, and @baseline to a @widget,
	// and computes the position and sizes of the children of the @widget using
	// the layout management policy of @manager.
	//
	// This method is inherited from LayoutManager
	Allocate(widget Widget, width int, height int, baseline int)
	// GetLayoutChild retrieves a `GtkLayoutChild` instance for the
	// `GtkLayoutManager`, creating one if necessary.
	//
	// The @child widget must be a child of the widget using @manager.
	//
	// The `GtkLayoutChild` instance is owned by the `GtkLayoutManager`, and is
	// guaranteed to exist as long as @child is a child of the `GtkWidget` using
	// the given `GtkLayoutManager`.
	//
	// This method is inherited from LayoutManager
	GetLayoutChild(child Widget) LayoutChild
	// GetRequestMode retrieves the request mode of @manager.
	//
	// This method is inherited from LayoutManager
	GetRequestMode() SizeRequestMode
	// GetWidget retrieves the `GtkWidget` using the given `GtkLayoutManager`.
	//
	// This method is inherited from LayoutManager
	GetWidget() Widget
	// LayoutChanged queues a resize on the `GtkWidget` using @manager, if any.
	//
	// This function should be called by subclasses of `GtkLayoutManager` in
	// response to changes to their layout management policies.
	//
	// This method is inherited from LayoutManager
	LayoutChanged()
	// Measure measures the size of the @widget using @manager, for the given
	// @orientation and size.
	//
	// See the [class@Gtk.Widget] documentation on layout management for more
	// details.
	//
	// This method is inherited from LayoutManager
	Measure(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int)
	// GetOrientation retrieves the orientation of the @orientable.
	//
	// This method is inherited from Orientable
	GetOrientation() Orientation
	// SetOrientation sets the orientation of the @orientable.
	//
	// This method is inherited from Orientable
	SetOrientation(orientation Orientation)

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
	*externglib.Object
}

var _ BoxLayout = (*boxLayout)(nil)

// WrapBoxLayout wraps a GObject to a type that implements
// interface BoxLayout. It is primarily used internally.
func WrapBoxLayout(obj *externglib.Object) BoxLayout {
	return boxLayout{obj}
}

func marshalBoxLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBoxLayout(obj), nil
}

// NewBoxLayout creates a new `GtkBoxLayout`.
func NewBoxLayout(orientation Orientation) BoxLayout {
	var _arg1 C.GtkOrientation    // out
	var _cret *C.GtkLayoutManager // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_box_layout_new(_arg1)

	var _boxLayout BoxLayout // out

	_boxLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(BoxLayout)

	return _boxLayout
}

func (b boxLayout) AsLayoutManager() LayoutManager {
	return WrapLayoutManager(gextras.InternObject(b))
}

func (b boxLayout) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(b))
}

func (m boxLayout) Allocate(widget Widget, width int, height int, baseline int) {
	WrapLayoutManager(gextras.InternObject(m)).Allocate(widget, width, height, baseline)
}

func (m boxLayout) GetLayoutChild(child Widget) LayoutChild {
	return WrapLayoutManager(gextras.InternObject(m)).GetLayoutChild(child)
}

func (m boxLayout) GetRequestMode() SizeRequestMode {
	return WrapLayoutManager(gextras.InternObject(m)).GetRequestMode()
}

func (m boxLayout) GetWidget() Widget {
	return WrapLayoutManager(gextras.InternObject(m)).GetWidget()
}

func (m boxLayout) LayoutChanged() {
	WrapLayoutManager(gextras.InternObject(m)).LayoutChanged()
}

func (m boxLayout) Measure(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int) {
	return WrapLayoutManager(gextras.InternObject(m)).Measure(widget, orientation, forSize)
}

func (o boxLayout) GetOrientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).GetOrientation()
}

func (o boxLayout) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}

func (b boxLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBoxLayout       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (b boxLayout) Homogeneous() bool {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b boxLayout) Spacing() uint {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_layout_get_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b boxLayout) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBoxLayout       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_layout_set_baseline_position(_arg0, _arg1)
}

func (b boxLayout) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_layout_set_homogeneous(_arg0, _arg1)
}

func (b boxLayout) SetSpacing(spacing uint) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_box_layout_set_spacing(_arg0, _arg1)
}
