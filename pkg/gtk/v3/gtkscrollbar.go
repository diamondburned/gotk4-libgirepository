// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scrollbar_get_type()), F: marshalScrollbar},
	})
}

// Scrollbar: the Scrollbar widget is a horizontal or vertical scrollbar,
// depending on the value of the Orientable:orientation property.
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by gtk_scrollbar_new(). See Adjustment for more details. The
// Adjustment:value field sets the position of the thumb and must be between
// Adjustment:lower and Adjustment:upper - Adjustment:page-size. The
// Adjustment:page-size represents the size of the visible scrollable area. The
// fields Adjustment:step-increment and Adjustment:page-increment fields are
// added to or subtracted from the Adjustment:value when the user asks to move
// by a step (using e.g. the cursor arrow keys or, if present, the stepper
// buttons) or by a page (using e.g. the Page Down/Up keys).
//
// CSS nodes
//
//    scrollbar[.fine-tune]
//    ╰── contents
//        ├── [button.up]
//        ├── [button.down]
//        ├── trough
//        │   ╰── slider
//        ├── [button.up]
//        ╰── [button.down]
//
// GtkScrollbar has a main CSS node with name scrollbar and a subnode for its
// contents, with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// If steppers are enabled, they are represented by up to four additional
// subnodes with name button. These get the style classes .up and .down to
// indicate in which direction they are moving.
//
// Other style classes that may be added to scrollbars inside ScrolledWindow
// include the positional classes (.left, .right, .top, .bottom) and style
// classes related to overlay scrolling (.overlay-indicator, .dragging,
// .hovering).
type Scrollbar interface {
	Range
	Buildable
	Orientable
}

// scrollbar implements the Scrollbar class.
type scrollbar struct {
	Range
	Buildable
	Orientable
}

var _ Scrollbar = (*scrollbar)(nil)

// WrapScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrollbar(obj *externglib.Object) Scrollbar {
	return scrollbar{
		Range:      WrapRange(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollbar(obj), nil
}
