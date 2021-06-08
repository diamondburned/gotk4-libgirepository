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
		{T: externglib.Type(C.gtk_gesture_click_get_type()), F: marshalGestureClick},
	})
}

// GestureClick: `GtkGestureClick` is a `GtkGesture` implementation for clicks.
//
// It is able to recognize multiple clicks on a nearby zone, which can be
// listened for through the [signal@Gtk.GestureClick::pressed] signal. Whenever
// time or distance between clicks exceed the GTK defaults,
// [signal@Gtk.GestureClick::stopped] is emitted, and the click counter is
// reset.
type GestureClick interface {
	GestureSingle
}

// gestureClick implements the GestureClick interface.
type gestureClick struct {
	GestureSingle
}

var _ GestureClick = (*gestureClick)(nil)

// WrapGestureClick wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureClick(obj *externglib.Object) GestureClick {
	return GestureClick{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureClick(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureClick(obj), nil
}

// NewGestureClick constructs a class GestureClick.
func NewGestureClick() GestureClick {
	cret := new(C.GtkGestureClick)
	var goret GestureClick

	cret = C.gtk_gesture_click_new()

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureClick)

	return goret
}
