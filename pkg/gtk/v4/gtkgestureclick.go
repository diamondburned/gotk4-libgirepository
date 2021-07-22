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
		{T: externglib.Type(C.gtk_gesture_click_get_type()), F: marshalGestureClicker},
	})
}

// GestureClick: GtkGestureClick is a GtkGesture implementation for clicks.
//
// It is able to recognize multiple clicks on a nearby zone, which can be
// listened for through the gtk.GestureClick::pressed signal. Whenever time or
// distance between clicks exceed the GTK defaults, gtk.GestureClick::stopped is
// emitted, and the click counter is reset.
type GestureClick struct {
	GestureSingle
}

func wrapGestureClick(obj *externglib.Object) *GestureClick {
	return &GestureClick{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureClicker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureClick(obj), nil
}

// NewGestureClick returns a newly created GtkGesture that recognizes single and
// multiple presses.
func NewGestureClick() *GestureClick {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_click_new()

	var _gestureClick *GestureClick // out

	_gestureClick = wrapGestureClick(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureClick
}

func (*GestureClick) privateGestureClick() {}
