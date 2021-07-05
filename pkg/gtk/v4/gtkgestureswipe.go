// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_gesture_swipe_get_type()), F: marshalGestureSwipe},
	})
}

// GestureSwipe: `GtkGestureSwipe` is a `GtkGesture` for swipe gestures.
//
// After a press/move/.../move/release sequence happens, the
// [signal@Gtk.GestureSwipe::swipe] signal will be emitted, providing the
// velocity and directionality of the sequence at the time it was lifted.
//
// If the velocity is desired in intermediate points,
// [method@Gtk.GestureSwipe.get_velocity] can be called in a
// [signal@Gtk.Gesture::update] handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe interface {
	GestureSingle

	// Velocity gets the current velocity.
	//
	// If the gesture is recognized, this function returns true and fills in
	// @velocity_x and @velocity_y with the recorded velocity, as per the last
	// events processed.
	Velocity() (velocityX float64, velocityY float64, ok bool)
}

// gestureSwipe implements the GestureSwipe class.
type gestureSwipe struct {
	GestureSingle
}

// WrapGestureSwipe wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureSwipe(obj *externglib.Object) GestureSwipe {
	return gestureSwipe{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureSwipe(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureSwipe(obj), nil
}

// NewGestureSwipe returns a newly created `GtkGesture` that recognizes swipes.
func NewGestureSwipe() GestureSwipe {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_swipe_new()

	var _gestureSwipe GestureSwipe // out

	_gestureSwipe = WrapGestureSwipe(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureSwipe
}

func (g gestureSwipe) Velocity() (velocityX float64, velocityY float64, ok bool) {
	var _arg0 *C.GtkGestureSwipe // out
	var _arg1 C.double           // in
	var _arg2 C.double           // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkGestureSwipe)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_swipe_get_velocity(_arg0, &_arg1, &_arg2)

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
