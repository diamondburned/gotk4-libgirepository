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
// extern void _gotk4_gtk4_GestureLongPress_ConnectCancelled(gpointer, guintptr);
// extern void _gotk4_gtk4_GestureLongPress_ConnectPressed(gpointer, gdouble, gdouble, guintptr);
import "C"

// glib.Type values for gtkgesturelongpress.go.
var GTypeGestureLongPress = externglib.Type(C.gtk_gesture_long_press_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGestureLongPress, F: marshalGestureLongPress},
	})
}

// GestureLongPressOverrider contains methods that are overridable.
type GestureLongPressOverrider interface {
	externglib.Objector
}

// GestureLongPress: GtkGestureLongPress is a GtkGesture for long presses.
//
// This gesture is also known as “Press and Hold”.
//
// When the timeout is exceeded, the gesture is triggering the
// gtk.GestureLongPress::pressed signal.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the gtk.GestureLongPress::cancelled signal
// will be emitted.
//
// How long the timeout is before the ::pressed signal gets emitted is
// determined by the gtk.Settings:gtk-long-press-time setting. It can be
// modified by the gtk.GestureLongPress:delay-factor property.
type GestureLongPress struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureLongPress)(nil)
)

func classInitGestureLongPresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureLongPress(obj *externglib.Object) *GestureLongPress {
	return &GestureLongPress{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	return wrapGestureLongPress(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GestureLongPress_ConnectCancelled
func _gotk4_gtk4_GestureLongPress_ConnectCancelled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectCancelled is emitted whenever a press moved too far, or was released
// before gtk.GestureLongPress::pressed happened.
func (gesture *GestureLongPress) ConnectCancelled(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "cancelled", false, unsafe.Pointer(C._gotk4_gtk4_GestureLongPress_ConnectCancelled), f)
}

//export _gotk4_gtk4_GestureLongPress_ConnectPressed
func _gotk4_gtk4_GestureLongPress_ConnectPressed(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectPressed is emitted whenever a press goes unmoved/unreleased longer
// than what the GTK defaults tell.
func (gesture *GestureLongPress) ConnectPressed(f func(x, y float64)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(gesture, "pressed", false, unsafe.Pointer(C._gotk4_gtk4_GestureLongPress_ConnectPressed), f)
}

// NewGestureLongPress returns a newly created GtkGesture that recognizes long
// presses.
//
// The function returns the following values:
//
//    - gestureLongPress: newly created GtkGestureLongPress.
//
func NewGestureLongPress() *GestureLongPress {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_long_press_new()

	var _gestureLongPress *GestureLongPress // out

	_gestureLongPress = wrapGestureLongPress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureLongPress
}

// DelayFactor returns the delay factor.
//
// The function returns the following values:
//
//    - gdouble: delay factor.
//
func (gesture *GestureLongPress) DelayFactor() float64 {
	var _arg0 *C.GtkGestureLongPress // out
	var _cret C.double               // in

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(externglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_long_press_get_delay_factor(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDelayFactor applies the given delay factor.
//
// The default long press time will be multiplied by this value. Valid values
// are in the range [0.5..2.0].
//
// The function takes the following parameters:
//
//    - delayFactor: delay factor to apply.
//
func (gesture *GestureLongPress) SetDelayFactor(delayFactor float64) {
	var _arg0 *C.GtkGestureLongPress // out
	var _arg1 C.double               // out

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(externglib.InternObject(gesture).Native()))
	_arg1 = C.double(delayFactor)

	C.gtk_gesture_long_press_set_delay_factor(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(delayFactor)
}
