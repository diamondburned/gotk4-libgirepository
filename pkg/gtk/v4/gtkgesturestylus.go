// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylus},
	})
}

// GestureStylus is a Gesture implementation specific to stylus input. The
// provided signals just relay the basic information of the stylus events.
type GestureStylus interface {
	GestureSingle

	// Axis returns the current value for the requested @axis.
	//
	// This function must be called from the handler of one of the
	// GestureStylus::down, GestureStylus::motion, GestureStylus::up or
	// GestureStylus::proximity signals.
	Axis(axis gdk.AxisUse) (float64, bool)
	// Backlog: by default, GTK will limit rate of input events. On stylus input
	// where accuracy of strokes is paramount, this function returns the
	// accumulated coordinate/timing state before the emission of the current
	// GestureStylus::motion signal.
	//
	// This function may only be called within a GestureStylus::motion signal
	// handler, the state given in this signal and obtainable through
	// gtk_gesture_stylus_get_axis() call express the latest (most up-to-date)
	// state in motion history.
	//
	// The @backlog is provided in chronological order.
	Backlog() bool
	// DeviceTool returns the DeviceTool currently driving input through this
	// gesture. This function must be called from either the
	// GestureStylus::down, GestureStylus::motion, GestureStylus::up or
	// GestureStylus::proximity signal handlers.
	DeviceTool() gdk.DeviceTool
}

// gestureStylus implements the GestureStylus interface.
type gestureStylus struct {
	GestureSingle
}

var _ GestureStylus = (*gestureStylus)(nil)

// WrapGestureStylus wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureStylus(obj *externglib.Object) GestureStylus {
	return GestureStylus{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureStylus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureStylus(obj), nil
}

// NewGestureStylus constructs a class GestureStylus.
func NewGestureStylus() GestureStylus {
	var _cret C.GtkGestureStylus

	cret = C.gtk_gesture_stylus_new()

	var _gestureStylus GestureStylus

	_gestureStylus = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(GestureStylus)

	return _gestureStylus
}

// Axis returns the current value for the requested @axis.
//
// This function must be called from the handler of one of the
// GestureStylus::down, GestureStylus::motion, GestureStylus::up or
// GestureStylus::proximity signals.
func (g gestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus
	var _arg1 C.GdkAxisUse

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))
	_arg1 = (C.GdkAxisUse)(axis)

	var _arg2 C.double
	var _cret C.gboolean

	cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)

	var _value float64
	var _ok bool

	_value = (float64)(_arg2)
	if _cret {
		_ok = true
	}

	return _value, _ok
}

// Backlog: by default, GTK will limit rate of input events. On stylus input
// where accuracy of strokes is paramount, this function returns the
// accumulated coordinate/timing state before the emission of the current
// GestureStylus::motion signal.
//
// This function may only be called within a GestureStylus::motion signal
// handler, the state given in this signal and obtainable through
// gtk_gesture_stylus_get_axis() call express the latest (most up-to-date)
// state in motion history.
//
// The @backlog is provided in chronological order.
func (g gestureStylus) Backlog() bool {
	var _arg0 *C.GtkGestureStylus

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean

	cret = C.gtk_gesture_stylus_get_backlog(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DeviceTool returns the DeviceTool currently driving input through this
// gesture. This function must be called from either the
// GestureStylus::down, GestureStylus::motion, GestureStylus::up or
// GestureStylus::proximity signal handlers.
func (g gestureStylus) DeviceTool() gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	var _cret *C.GdkDeviceTool

	cret = C.gtk_gesture_stylus_get_device_tool(_arg0)

	var _deviceTool gdk.DeviceTool

	_deviceTool = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gdk.DeviceTool)

	return _deviceTool
}
