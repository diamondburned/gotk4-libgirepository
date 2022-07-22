// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_Monitor_ConnectInvalidate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSubpixelLayout = coreglib.Type(C.gdk_subpixel_layout_get_type())
	GTypeMonitor        = coreglib.Type(C.gdk_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSubpixelLayout, F: marshalSubpixelLayout},
		coreglib.TypeMarshaler{T: GTypeMonitor, F: marshalMonitor},
	})
}

// SubpixelLayout: this enumeration describes how the red, green and blue
// components of physical pixels on an output device are laid out.
type SubpixelLayout C.gint

const (
	// SubpixelLayoutUnknown: layout is not known.
	SubpixelLayoutUnknown SubpixelLayout = iota
	// SubpixelLayoutNone: not organized in this way.
	SubpixelLayoutNone
	// SubpixelLayoutHorizontalRGB: layout is horizontal, the order is RGB.
	SubpixelLayoutHorizontalRGB
	// SubpixelLayoutHorizontalBGR: layout is horizontal, the order is BGR.
	SubpixelLayoutHorizontalBGR
	// SubpixelLayoutVerticalRGB: layout is vertical, the order is RGB.
	SubpixelLayoutVerticalRGB
	// SubpixelLayoutVerticalBGR: layout is vertical, the order is BGR.
	SubpixelLayoutVerticalBGR
)

func marshalSubpixelLayout(p uintptr) (interface{}, error) {
	return SubpixelLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SubpixelLayout.
func (s SubpixelLayout) String() string {
	switch s {
	case SubpixelLayoutUnknown:
		return "Unknown"
	case SubpixelLayoutNone:
		return "None"
	case SubpixelLayoutHorizontalRGB:
		return "HorizontalRGB"
	case SubpixelLayoutHorizontalBGR:
		return "HorizontalBGR"
	case SubpixelLayoutVerticalRGB:
		return "VerticalRGB"
	case SubpixelLayoutVerticalBGR:
		return "VerticalBGR"
	default:
		return fmt.Sprintf("SubpixelLayout(%d)", s)
	}
}

// Monitor: GdkMonitor objects represent the individual outputs that are
// associated with a GdkDisplay.
//
// GdkDisplay keeps a GListModel to enumerate and monitor monitors with
// gdk.Display.GetMonitors(). You can use gdk.Display.GetMonitorAtSurface() to
// find a particular monitor.
type Monitor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Monitor)(nil)
)

func wrapMonitor(obj *coreglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	return wrapMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk4_Monitor_ConnectInvalidate
func _gotk4_gdk4_Monitor_ConnectInvalidate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectInvalidate is emitted when the output represented by monitor gets
// disconnected.
func (monitor *Monitor) ConnectInvalidate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(monitor, "invalidate", false, unsafe.Pointer(C._gotk4_gdk4_Monitor_ConnectInvalidate), f)
}

// Connector gets the name of the monitor's connector, if available.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the connector.
//
func (monitor *Monitor) Connector() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_connector(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Display gets the display that this monitor belongs to.
//
// The function returns the following values:
//
//    - display: display.
//
func (monitor *Monitor) Display() *Display {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_display(_arg0)
	runtime.KeepAlive(monitor)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Geometry retrieves the size and position of the monitor within the display
// coordinate space.
//
// The returned geometry is in ”application pixels”, not in ”device pixels” (see
// gdk.Monitor.GetScaleFactor()).
//
// The function returns the following values:
//
//    - geometry: GdkRectangle to be filled with the monitor geometry.
//
func (monitor *Monitor) Geometry() *Rectangle {
	var _arg0 *C.GdkMonitor  // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	C.gdk_monitor_get_geometry(_arg0, &_arg1)
	runtime.KeepAlive(monitor)

	var _geometry *Rectangle // out

	_geometry = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _geometry
}

// HeightMm gets the height in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical height of the monitor.
//
func (monitor *Monitor) HeightMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_height_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer.
//
// Note that this value might also vary depending on actual display backend.
//
// The PNP ID registry is located at https://uefi.org/pnp_id_list
// (https://uefi.org/pnp_id_list).
//
// The function returns the following values:
//
//    - utf8 (optional): name of the manufacturer, or NULL.
//
func (monitor *Monitor) Manufacturer() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_manufacturer(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Model gets the string identifying the monitor model, if available.
//
// The function returns the following values:
//
//    - utf8 (optional): monitor model, or NULL.
//
func (monitor *Monitor) Model() string {
	var _arg0 *C.GdkMonitor // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_model(_arg0)
	runtime.KeepAlive(monitor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as 60000.
//
// The function returns the following values:
//
//    - gint: refresh rate in milli-Hertz, or 0.
//
func (monitor *Monitor) RefreshRate() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_refresh_rate(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScaleFactor gets the internal scale factor that maps from monitor coordinates
// to device pixels.
//
// On traditional systems this is 1, but on very high density outputs it can be
// a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a surface where it is better
// to use gdk.Surface.GetScaleFactor() instead.
//
// The function returns the following values:
//
//    - gint: scale factor.
//
func (monitor *Monitor) ScaleFactor() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_scale_factor(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for pixels.
//
// The function returns the following values:
//
//    - subpixelLayout: subpixel layout.
//
func (monitor *Monitor) SubpixelLayout() SubpixelLayout {
	var _arg0 *C.GdkMonitor       // out
	var _cret C.GdkSubpixelLayout // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_subpixel_layout(_arg0)
	runtime.KeepAlive(monitor)

	var _subpixelLayout SubpixelLayout // out

	_subpixelLayout = SubpixelLayout(_cret)

	return _subpixelLayout
}

// WidthMm gets the width in millimeters of the monitor.
//
// The function returns the following values:
//
//    - gint: physical width of the monitor.
//
func (monitor *Monitor) WidthMm() int {
	var _arg0 *C.GdkMonitor // out
	var _cret C.int         // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_get_width_mm(_arg0)
	runtime.KeepAlive(monitor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsValid returns TRUE if the monitor object corresponds to a physical monitor.
//
// The monitor becomes invalid when the physical monitor is unplugged or
// removed.
//
// The function returns the following values:
//
//    - ok: TRUE if the object corresponds to a physical monitor.
//
func (monitor *Monitor) IsValid() bool {
	var _arg0 *C.GdkMonitor // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.gdk_monitor_is_valid(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
