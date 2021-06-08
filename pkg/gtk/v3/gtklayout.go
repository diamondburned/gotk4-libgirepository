// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_layout_get_type()), F: marshalLayout},
	})
}

// Layout is similar to DrawingArea in that it’s a “blank slate” and doesn’t do
// anything except paint a blank background by default. It’s different in that
// it supports scrolling natively due to implementing Scrollable, and can
// contain child widgets since it’s a Container.
//
// If you just want to draw, a DrawingArea is a better choice since it has lower
// overhead. If you just need to position child widgets at specific points, then
// Fixed provides that functionality on its own.
//
// When handling expose events on a Layout, you must draw to the Window returned
// by gtk_layout_get_bin_window(), rather than to the one returned by
// gtk_widget_get_window() as you would for a DrawingArea.
type Layout interface {
	Container
	Buildable
	Scrollable

	// BinWindow: retrieve the bin window of the layout used for drawing
	// operations.
	BinWindow() gdk.Window
	// HAdjustment: this function should only be called after the layout has
	// been placed in a ScrolledWindow or otherwise configured for scrolling. It
	// returns the Adjustment used for communication between the horizontal
	// scrollbar and @layout.
	//
	// See ScrolledWindow, Scrollbar, Adjustment for details.
	HAdjustment() Adjustment
	// Size gets the size that has been set on the layout, and that determines
	// the total extents of the layout’s scrollbar area. See gtk_layout_set_size
	// ().
	Size() (width uint, height uint)
	// VAdjustment: this function should only be called after the layout has
	// been placed in a ScrolledWindow or otherwise configured for scrolling. It
	// returns the Adjustment used for communication between the vertical
	// scrollbar and @layout.
	//
	// See ScrolledWindow, Scrollbar, Adjustment for details.
	VAdjustment() Adjustment
	// Move moves a current child of @layout to a new position.
	Move(childWidget Widget, x int, y int)
	// Put adds @child_widget to @layout, at position (@x,@y). @layout becomes
	// the new parent container of @child_widget.
	Put(childWidget Widget, x int, y int)
	// SetHAdjustment sets the horizontal scroll adjustment for the layout.
	//
	// See ScrolledWindow, Scrollbar, Adjustment for details.
	SetHAdjustment(adjustment Adjustment)
	// SetSize sets the size of the scrollable area of the layout.
	SetSize(width uint, height uint)
	// SetVAdjustment sets the vertical scroll adjustment for the layout.
	//
	// See ScrolledWindow, Scrollbar, Adjustment for details.
	SetVAdjustment(adjustment Adjustment)
}

// layout implements the Layout interface.
type layout struct {
	Container
	Buildable
	Scrollable
}

var _ Layout = (*layout)(nil)

// WrapLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapLayout(obj *externglib.Object) Layout {
	return Layout{
		Container:  WrapContainer(obj),
		Buildable:  WrapBuildable(obj),
		Scrollable: WrapScrollable(obj),
	}
}

func marshalLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLayout(obj), nil
}

// NewLayout constructs a class Layout.
func NewLayout(hAdjustment Adjustment, vAdjustment Adjustment) Layout {
	var arg1 *C.GtkAdjustment
	var arg2 *C.GtkAdjustment

	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hAdjustment.Native()))
	arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vAdjustment.Native()))

	var cret C.GtkLayout
	var goret Layout

	cret = C.gtk_layout_new(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Layout)

	return goret
}

// BinWindow: retrieve the bin window of the layout used for drawing
// operations.
func (l layout) BinWindow() gdk.Window {
	var arg0 *C.GtkLayout

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	var cret *C.GdkWindow
	var goret gdk.Window

	cret = C.gtk_layout_get_bin_window(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Window)

	return goret
}

// HAdjustment: this function should only be called after the layout has
// been placed in a ScrolledWindow or otherwise configured for scrolling. It
// returns the Adjustment used for communication between the horizontal
// scrollbar and @layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
func (l layout) HAdjustment() Adjustment {
	var arg0 *C.GtkLayout

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	var cret *C.GtkAdjustment
	var goret Adjustment

	cret = C.gtk_layout_get_hadjustment(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return goret
}

// Size gets the size that has been set on the layout, and that determines
// the total extents of the layout’s scrollbar area. See gtk_layout_set_size
// ().
func (l layout) Size() (width uint, height uint) {
	var arg0 *C.GtkLayout

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	arg1 := new(C.guint)
	var ret1 uint
	arg2 := new(C.guint)
	var ret2 uint

	C.gtk_layout_get_size(arg0, arg1, arg2)

	ret1 = uint(*arg1)
	ret2 = uint(*arg2)

	return ret1, ret2
}

// VAdjustment: this function should only be called after the layout has
// been placed in a ScrolledWindow or otherwise configured for scrolling. It
// returns the Adjustment used for communication between the vertical
// scrollbar and @layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
func (l layout) VAdjustment() Adjustment {
	var arg0 *C.GtkLayout

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	var cret *C.GtkAdjustment
	var goret Adjustment

	cret = C.gtk_layout_get_vadjustment(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return goret
}

// Move moves a current child of @layout to a new position.
func (l layout) Move(childWidget Widget, x int, y int) {
	var arg0 *C.GtkLayout
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 C.gint

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)

	C.gtk_layout_move(arg0, arg1, arg2, arg3)
}

// Put adds @child_widget to @layout, at position (@x,@y). @layout becomes
// the new parent container of @child_widget.
func (l layout) Put(childWidget Widget, x int, y int) {
	var arg0 *C.GtkLayout
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 C.gint

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)

	C.gtk_layout_put(arg0, arg1, arg2, arg3)
}

// SetHAdjustment sets the horizontal scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
func (l layout) SetHAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkLayout
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_hadjustment(arg0, arg1)
}

// SetSize sets the size of the scrollable area of the layout.
func (l layout) SetSize(width uint, height uint) {
	var arg0 *C.GtkLayout
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	arg1 = C.guint(width)
	arg2 = C.guint(height)

	C.gtk_layout_set_size(arg0, arg1, arg2)
}

// SetVAdjustment sets the vertical scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
func (l layout) SetVAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkLayout
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_vadjustment(arg0, arg1)
}
