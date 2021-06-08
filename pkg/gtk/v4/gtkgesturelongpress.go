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
		{T: externglib.Type(C.gtk_gesture_long_press_get_type()), F: marshalGestureLongPress},
	})
}

// GestureLongPress: `GtkGestureLongPress` is a `GtkGesture` for long presses.
//
// This gesture is also known as “Press and Hold”.
//
// When the timeout is exceeded, the gesture is triggering the
// [signal@Gtk.GestureLongPress::pressed] signal.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the [signal@Gtk.GestureLongPress::cancelled]
// signal will be emitted.
//
// How long the timeout is before the ::pressed signal gets emitted is
// determined by the [property@Gtk.Settings:gtk-long-press-time] setting. It can
// be modified by the [property@Gtk.GestureLongPress:delay-factor] property.
type GestureLongPress interface {
	GestureSingle

	// DelayFactor returns the delay factor.
	DelayFactor() float64
	// SetDelayFactor applies the given delay factor.
	//
	// The default long press time will be multiplied by this value. Valid
	// values are in the range [0.5..2.0].
	SetDelayFactor(delayFactor float64)
}

// gestureLongPress implements the GestureLongPress interface.
type gestureLongPress struct {
	GestureSingle
}

var _ GestureLongPress = (*gestureLongPress)(nil)

// WrapGestureLongPress wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureLongPress(obj *externglib.Object) GestureLongPress {
	return GestureLongPress{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureLongPress(obj), nil
}

// NewGestureLongPress constructs a class GestureLongPress.
func NewGestureLongPress() GestureLongPress {
	cret := new(C.GtkGestureLongPress)
	var goret GestureLongPress

	cret = C.gtk_gesture_long_press_new()

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureLongPress)

	return goret
}

// DelayFactor returns the delay factor.
func (g gestureLongPress) DelayFactor() float64 {
	var arg0 *C.GtkGestureLongPress

	arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(g.Native()))

	var cret C.double
	var goret float64

	cret = C.gtk_gesture_long_press_get_delay_factor(arg0)

	goret = float64(cret)

	return goret
}

// SetDelayFactor applies the given delay factor.
//
// The default long press time will be multiplied by this value. Valid
// values are in the range [0.5..2.0].
func (g gestureLongPress) SetDelayFactor(delayFactor float64) {
	var arg0 *C.GtkGestureLongPress
	var arg1 C.double

	arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(g.Native()))
	arg1 = C.double(delayFactor)

	C.gtk_gesture_long_press_set_delay_factor(arg0, arg1)
}
