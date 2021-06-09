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

// GestureLongPress is a Gesture implementation able to recognize long presses,
// triggering the GestureLongPress::pressed after the timeout is exceeded.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the GestureLongPress::cancelled signal will
// be emitted.
type GestureLongPress interface {
	GestureSingle

	// DelayFactor returns the delay factor as set by
	// gtk_gesture_long_press_set_delay_factor().
	DelayFactor() float64
	// SetDelayFactor applies the given delay factor. The default long press
	// time will be multiplied by this value. Valid values are in the range
	// [0.5..2.0].
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
	var _cret C.GtkGestureLongPress

	cret = C.gtk_gesture_long_press_new()

	var _gestureLongPress GestureLongPress

	_gestureLongPress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(GestureLongPress)

	return _gestureLongPress
}

// DelayFactor returns the delay factor as set by
// gtk_gesture_long_press_set_delay_factor().
func (g gestureLongPress) DelayFactor() float64 {
	var _arg0 *C.GtkGestureLongPress

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(g.Native()))

	var _cret C.double

	cret = C.gtk_gesture_long_press_get_delay_factor(_arg0)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// SetDelayFactor applies the given delay factor. The default long press
// time will be multiplied by this value. Valid values are in the range
// [0.5..2.0].
func (g gestureLongPress) SetDelayFactor(delayFactor float64) {
	var _arg0 *C.GtkGestureLongPress
	var _arg1 C.double

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(g.Native()))
	_arg1 = C.double(delayFactor)

	C.gtk_gesture_long_press_set_delay_factor(_arg0, _arg1)
}
