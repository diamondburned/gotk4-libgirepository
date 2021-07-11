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
		{T: externglib.Type(C.gtk_gesture_pan_get_type()), F: marshalGesturePanner},
	})
}

// GesturePanner describes GesturePan's methods.
type GesturePanner interface {
	// Orientation returns the orientation of the pan gestures that this
	// @gesture expects.
	Orientation() Orientation
}

// GesturePan: `GtkGesturePan` is a `GtkGesture` for pan gestures.
//
// These are drags that are locked to happen along one axis. The axis that a
// `GtkGesturePan` handles is defined at construct time, and can be changed
// through [method@Gtk.GesturePan.set_orientation].
//
// When the gesture starts to be recognized, `GtkGesturePan` will attempt to
// determine as early as possible whether the sequence is moving in the expected
// direction, and denying the sequence if this does not happen.
//
// Once a panning gesture along the expected axis is recognized, the
// [signal@Gtk.GesturePan::pan] signal will be emitted as input events are
// received, containing the offset in the given axis.
type GesturePan struct {
	GestureDrag
}

var (
	_ GesturePanner   = (*GesturePan)(nil)
	_ gextras.Nativer = (*GesturePan)(nil)
)

func wrapGesturePan(obj *externglib.Object) GesturePanner {
	return &GesturePan{
		GestureDrag: GestureDrag{
			GestureSingle: GestureSingle{
				Gesture: Gesture{
					EventController: EventController{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalGesturePanner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGesturePan(obj), nil
}

// Orientation returns the orientation of the pan gestures that this @gesture
// expects.
func (gesture *GesturePan) Orientation() Orientation {
	var _arg0 *C.GtkGesturePan // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_pan_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}
