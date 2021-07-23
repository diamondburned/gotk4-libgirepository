// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_center_layout_get_type()), F: marshalCenterLayouter},
	})
}

// CenterLayout: GtkCenterLayout is a layout manager that manages up to three
// children.
//
// The start widget is allocated at the start of the layout (left in
// left-to-right locales and right in right-to-left ones), and the end widget at
// the end.
//
// The center widget is centered regarding the full width of the layout's.
type CenterLayout struct {
	LayoutManager
}

func wrapCenterLayout(obj *externglib.Object) *CenterLayout {
	return &CenterLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalCenterLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCenterLayout(obj), nil
}

// NewCenterLayout creates a new GtkCenterLayout.
func NewCenterLayout() *CenterLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_center_layout_new()

	var _centerLayout *CenterLayout // out

	_centerLayout = wrapCenterLayout(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _centerLayout
}

// BaselinePosition returns the baseline position of the layout.
func (self *CenterLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterLayout    // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget returns the center widget of the layout.
func (self *CenterLayout) CenterWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_center_widget(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// EndWidget returns the end widget of the layout.
func (self *CenterLayout) EndWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_end_widget(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Orientation gets the current orienration of the layout manager.
func (self *CenterLayout) Orientation() Orientation {
	var _arg0 *C.GtkCenterLayout // out
	var _cret C.GtkOrientation   // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// StartWidget returns the start widget fo the layout.
func (self *CenterLayout) StartWidget() Widgetter {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_layout_get_start_widget(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// SetBaselinePosition sets the new baseline position of self
func (self *CenterLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	var _arg0 *C.GtkCenterLayout    // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkBaselinePosition(baselinePosition)

	C.gtk_center_layout_set_baseline_position(_arg0, _arg1)
}

// SetCenterWidget sets the new center widget of self.
//
// To remove the existing center widget, pass NULL.
func (self *CenterLayout) SetCenterWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_center_widget(_arg0, _arg1)
}

// SetEndWidget sets the new end widget of self.
//
// To remove the existing center widget, pass NULL.
func (self *CenterLayout) SetEndWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_end_widget(_arg0, _arg1)
}

// SetOrientation sets the orientation of self.
func (self *CenterLayout) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 C.GtkOrientation   // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_center_layout_set_orientation(_arg0, _arg1)
}

// SetStartWidget sets the new start widget of self.
//
// To remove the existing start widget, pass NULL.
func (self *CenterLayout) SetStartWidget(widget Widgetter) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_center_layout_set_start_widget(_arg0, _arg1)
}
