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
		{T: externglib.Type(C.gtk_center_box_get_type()), F: marshalCenterBox},
	})
}

// CenterBox: `GtkCenterBox` arranges three children in a row, keeping the
// middle child centered as well as possible.
//
// !An example GtkCenterBox (centerbox.png)
//
// To add children to `GtkCenterBox`, use
// [method@Gtk.CenterBox.set_start_widget],
// [method@Gtk.CenterBox.set_center_widget] and
// [method@Gtk.CenterBox.set_end_widget].
//
// The sizing and positioning of children can be influenced with the align and
// expand properties of the children.
//
//
// GtkCenterBox as GtkBuildable
//
// The `GtkCenterBox` implementation of the `GtkBuildable` interface supports
// placing children in the 3 positions by specifying “start”, “center” or “end”
// as the “type” attribute of a <child> element.
//
//
// CSS nodes
//
// `GtkCenterBox` uses a single CSS node with the name “box”,
//
// The first child of the `GtkCenterBox` will be allocated depending on the text
// direction, i.e. in left-to-right layouts it will be allocated on the left and
// in right-to-left layouts on the right.
//
// In vertical orientation, the nodes of the children are arranged from top to
// bottom.
//
//
// Accessibility
//
// `GtkCenterBox` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type CenterBox interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable

	// BaselinePosition gets the value set by
	// gtk_center_box_set_baseline_position().
	BaselinePosition() BaselinePosition
	// CenterWidget gets the center widget, or nil if there is none.
	CenterWidget() Widget
	// EndWidget gets the end widget, or nil if there is none.
	EndWidget() Widget
	// StartWidget gets the start widget, or nil if there is none.
	StartWidget() Widget
	// SetBaselinePosition sets the baseline position of a center box.
	//
	// This affects only horizontal boxes with at least one baseline aligned
	// child. If there is more vertical space available than requested, and the
	// baseline is not allocated by the parent then @position is used to
	// allocate the baseline wrt. the extra space available.
	SetBaselinePosition(position BaselinePosition)
	// SetCenterWidget sets the center widget.
	//
	// To remove the existing center widget, pas nil.
	SetCenterWidget(child Widget)
	// SetEndWidget sets the end widget.
	//
	// To remove the existing end widget, pass nil.
	SetEndWidget(child Widget)
	// SetStartWidget sets the start widget.
	//
	// To remove the existing start widget, pass nil.
	SetStartWidget(child Widget)
}

// centerBox implements the CenterBox interface.
type centerBox struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ CenterBox = (*centerBox)(nil)

// WrapCenterBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapCenterBox(obj *externglib.Object) CenterBox {
	return CenterBox{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalCenterBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterBox(obj), nil
}

// NewCenterBox constructs a class CenterBox.
func NewCenterBox() CenterBox {
	var _cret C.GtkCenterBox

	cret = C.gtk_center_box_new()

	var _centerBox CenterBox

	_centerBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(CenterBox)

	return _centerBox
}

// BaselinePosition gets the value set by
// gtk_center_box_set_baseline_position().
func (s centerBox) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterBox

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	var _cret C.GtkBaselinePosition

	cret = C.gtk_center_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget gets the center widget, or nil if there is none.
func (s centerBox) CenterWidget() Widget {
	var _arg0 *C.GtkCenterBox

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_box_get_center_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// EndWidget gets the end widget, or nil if there is none.
func (s centerBox) EndWidget() Widget {
	var _arg0 *C.GtkCenterBox

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_box_get_end_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// StartWidget gets the start widget, or nil if there is none.
func (s centerBox) StartWidget() Widget {
	var _arg0 *C.GtkCenterBox

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_box_get_start_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetBaselinePosition sets the baseline position of a center box.
//
// This affects only horizontal boxes with at least one baseline aligned
// child. If there is more vertical space available than requested, and the
// baseline is not allocated by the parent then @position is used to
// allocate the baseline wrt. the extra space available.
func (s centerBox) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkCenterBox
	var _arg1 C.GtkBaselinePosition

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkBaselinePosition)(position)

	C.gtk_center_box_set_baseline_position(_arg0, _arg1)
}

// SetCenterWidget sets the center widget.
//
// To remove the existing center widget, pas nil.
func (s centerBox) SetCenterWidget(child Widget) {
	var _arg0 *C.GtkCenterBox
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_center_widget(_arg0, _arg1)
}

// SetEndWidget sets the end widget.
//
// To remove the existing end widget, pass nil.
func (s centerBox) SetEndWidget(child Widget) {
	var _arg0 *C.GtkCenterBox
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_end_widget(_arg0, _arg1)
}

// SetStartWidget sets the start widget.
//
// To remove the existing start widget, pass nil.
func (s centerBox) SetStartWidget(child Widget) {
	var _arg0 *C.GtkCenterBox
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_start_widget(_arg0, _arg1)
}