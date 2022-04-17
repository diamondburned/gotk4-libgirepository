// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GestureSwipe_ConnectSwipe(gpointer, gdouble, gdouble, guintptr);
import "C"

// glib.Type values for gtkgestureswipe.go.
var GTypeGestureSwipe = externglib.Type(C.gtk_gesture_swipe_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGestureSwipe, F: marshalGestureSwipe},
	})
}

// GestureSwipeOverrider contains methods that are overridable.
type GestureSwipeOverrider interface {
	externglib.Objector
}

// GestureSwipe: GtkGestureSwipe is a GtkGesture for swipe gestures.
//
// After a press/move/.../move/release sequence happens, the
// gtk.GestureSwipe::swipe signal will be emitted, providing the velocity and
// directionality of the sequence at the time it was lifted.
//
// If the velocity is desired in intermediate points,
// gtk.GestureSwipe.GetVelocity() can be called in a gtk.Gesture::update
// handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureSwipe)(nil)
)

func classInitGestureSwiper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

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

func marshalGestureSwipe(p uintptr) (interface{}, error) {
	return wrapGestureSwipe(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GestureSwipe_ConnectSwipe
func _gotk4_gtk4_GestureSwipe_ConnectSwipe(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velocityX, velocityY float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velocityX, velocityY float64))
	}

	var _velocityX float64 // out
	var _velocityY float64 // out

	_velocityX = float64(arg1)
	_velocityY = float64(arg2)

	f(_velocityX, _velocityY)
}

// ConnectSwipe is emitted when the recognized gesture is finished.
//
// Velocity and direction are a product of previously recorded events.
func (gesture *GestureSwipe) ConnectSwipe(f func(velocityX, velocityY float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "swipe", false, unsafe.Pointer(C._gotk4_gtk4_GestureSwipe_ConnectSwipe), f)
}

// NewGestureSwipe returns a newly created GtkGesture that recognizes swipes.
//
// The function returns the following values:
//
//    - gestureSwipe: newly created GtkGestureSwipe.
//
func NewGestureSwipe() *GestureSwipe {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_swipe_new()

	var _gestureSwipe *GestureSwipe // out

	_gestureSwipe = wrapGestureSwipe(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureSwipe
}

// Velocity gets the current velocity.
//
// If the gesture is recognized, this function returns TRUE and fills in
// velocity_x and velocity_y with the recorded velocity, as per the last events
// processed.
//
// The function returns the following values:
//
//    - velocityX: return value for the velocity in the X axis, in pixels/sec.
//    - velocityY: return value for the velocity in the Y axis, in pixels/sec.
//    - ok: whether velocity could be calculated.
//
func (gesture *GestureSwipe) Velocity() (velocityX float64, velocityY float64, ok bool) {
	var _arg0 *C.GtkGestureSwipe // out
	var _arg1 C.double           // in
	var _arg2 C.double           // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkGestureSwipe)(unsafe.Pointer(externglib.InternObject(gesture).Native()))

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
