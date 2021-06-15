// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scrollbar_get_type()), F: marshalScrollbar},
	})
}

// Scrollbar: the `GtkScrollbar` widget is a horizontal or vertical scrollbar.
//
// !An example GtkScrollbar (scrollbar.png)
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by [ctor@Gtk.Scrollbar.new]. See [class.Gtk.Adjustment] for more
// details. The [property@Gtk.Adjustment:value] field sets the position of the
// thumb and must be between [property@Gtk.Adjustment:lower] and
// [property@Gtk.Adjustment:upper] - [property@Gtk.Adjustment:page-size]. The
// [property@Gtk.Adjustment:page-size] represents the size of the visible
// scrollable area.
//
// The fields [property@Gtk.Adjustment:step-increment] and
// [property@Gtk.Adjustment:page-increment] fields are added to or subtracted
// from the [property@Gtk.Adjustment:value] when the user asks to move by a step
// (using e.g. the cursor arrow keys) or by a page (using e.g. the Page Down/Up
// keys).
//
//
// CSS nodes
//
// “` scrollbar ╰── range[.fine-tune] ╰── trough ╰── slider “`
//
// `GtkScrollbar` has a main CSS node with name scrollbar and a subnode for its
// contents. The main node gets the .horizontal or .vertical style classes
// applied, depending on the scrollbar's orientation.
//
// The range node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// Other style classes that may be added to scrollbars inside
// [class@Gtk.ScrolledWindow] include the positional classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering).
//
//
// Accessibility
//
// `GtkScrollbar` uses the GTK_ACCESSIBLE_ROLE_SCROLLBAR role.
type Scrollbar interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable

	// Adjustment returns the scrollbar's adjustment.
	Adjustment() Adjustment
	// SetAdjustment makes the scrollbar use the given adjustment.
	SetAdjustment(adjustment Adjustment)
}

// scrollbar implements the Scrollbar class.
type scrollbar struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Scrollbar = (*scrollbar)(nil)

// WrapScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrollbar(obj *externglib.Object) Scrollbar {
	return scrollbar{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollbar(obj), nil
}

// NewScrollbar constructs a class Scrollbar.
func NewScrollbar(orientation Orientation, adjustment Adjustment) Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret C.GtkScrollbar   // in

	_arg1 = (C.GtkOrientation)(orientation)
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)

	var _scrollbar Scrollbar // out

	_scrollbar = WrapScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _scrollbar
}

// Adjustment returns the scrollbar's adjustment.
func (s scrollbar) Adjustment() Adjustment {
	var _arg0 *C.GtkScrollbar  // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollbar_get_adjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

// SetAdjustment makes the scrollbar use the given adjustment.
func (s scrollbar) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkScrollbar  // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_scrollbar_set_adjustment(_arg0, _arg1)
}
