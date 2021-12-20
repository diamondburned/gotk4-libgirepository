// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_swipe_get_type()), F: marshalGestureSwiper},
	})
}

// GestureSwipe is a Gesture implementation able to recognize swipes, after a
// press/move/.../move/release sequence happens, the GestureSwipe::swipe signal
// will be emitted, providing the velocity and directionality of the sequence at
// the time it was lifted.
//
// If the velocity is desired in intermediate points,
// gtk_gesture_swipe_get_velocity() can be called on eg. a Gesture::update
// handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe struct {
	GestureSingle

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Gesturer = (*GestureSwipe)(nil)
)

func wrapGestureSwipe(obj *externglib.Object) *GestureSwipe {
	return &GestureSwipe{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureSwiper(p uintptr) (interface{}, error) {
	return wrapGestureSwipe(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectSwipe: this signal is emitted when the recognized gesture is finished,
// velocity and direction are a product of previously recorded events.
func (gesture *GestureSwipe) ConnectSwipe(f func(velocityX, velocityY float64)) externglib.SignalHandle {
	return gesture.Connect("swipe", f)
}

// NewGestureSwipe returns a newly created Gesture that recognizes swipes.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureSwipe: newly created GestureSwipe.
//
func NewGestureSwipe(widget Widgetter) *GestureSwipe {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_swipe_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureSwipe *GestureSwipe // out

	_gestureSwipe = wrapGestureSwipe(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureSwipe
}

// Velocity: if the gesture is recognized, this function returns TRUE and fill
// in velocity_x and velocity_y with the recorded velocity, as per the last
// event(s) processed.
//
// The function returns the following values:
//
//    - velocityX: return value for the velocity in the X axis, in pixels/sec.
//    - velocityY: return value for the velocity in the Y axis, in pixels/sec.
//    - ok: whether velocity could be calculated.
//
func (gesture *GestureSwipe) Velocity() (velocityX float64, velocityY float64, ok bool) {
	var _arg0 *C.GtkGestureSwipe // out
	var _arg1 C.gdouble          // in
	var _arg2 C.gdouble          // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkGestureSwipe)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_swipe_get_velocity(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _velocityX float64 // out
	var _velocityY float64 // out
	var _ok bool           // out

	_velocityX = float64(_arg1)
	_velocityY = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _velocityX, _velocityY, _ok
}
