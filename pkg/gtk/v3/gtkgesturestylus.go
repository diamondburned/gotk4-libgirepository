// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylus},
	})
}

// GestureStylus is a Gesture implementation specific to stylus input. The
// provided signals just provide the basic information
type GestureStylus interface {
	GestureSingle

	// Axis returns the current value for the requested @axis. This function
	// must be called from either the GestureStylus:down, GestureStylus:motion,
	// GestureStylus:up or GestureStylus:proximity signals.
	Axis(axis gdk.AxisUse) (value float64, ok bool)
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
func NewGestureStylus(widget Widget) GestureStylus {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	cret := new(C.GtkGestureStylus)
	var goret GestureStylus

	cret = C.gtk_gesture_stylus_new(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureStylus)

	return goret
}

// Axis returns the current value for the requested @axis. This function
// must be called from either the GestureStylus:down, GestureStylus:motion,
// GestureStylus:up or GestureStylus:proximity signals.
func (g gestureStylus) Axis(axis gdk.AxisUse) (value float64, ok bool) {
	var arg0 *C.GtkGestureStylus
	var arg1 C.GdkAxisUse

	arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))
	arg1 = (C.GdkAxisUse)(axis)

	arg2 := new(C.gdouble)
	var ret2 float64
	var cret C.gboolean
	var goret bool

	cret = C.gtk_gesture_stylus_get_axis(arg0, arg1, arg2)

	ret2 = float64(*arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// DeviceTool returns the DeviceTool currently driving input through this
// gesture. This function must be called from either the
// GestureStylus::down, GestureStylus::motion, GestureStylus::up or
// GestureStylus::proximity signal handlers.
func (g gestureStylus) DeviceTool() gdk.DeviceTool {
	var arg0 *C.GtkGestureStylus

	arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	var cret *C.GdkDeviceTool
	var goret gdk.DeviceTool

	cret = C.gtk_gesture_stylus_get_device_tool(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.DeviceTool)

	return goret
}
