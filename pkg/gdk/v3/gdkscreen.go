// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_Screen_ConnectCompositedChanged(gpointer, guintptr);
// extern void _gotk4_gdk3_Screen_ConnectMonitorsChanged(gpointer, guintptr);
// extern void _gotk4_gdk3_Screen_ConnectSizeChanged(gpointer, guintptr);
import "C"

// GTypeScreen returns the GType for the type Screen.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScreen() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "Screen").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalScreen)
	return gtype
}

// Screen objects are the GDK representation of the screen on which windows can
// be displayed and on which the pointer moves. X originally identified screens
// with physical screens, but nowadays it is more common to have a single Screen
// which combines several physical monitors (see gdk_screen_get_n_monitors()).
//
// GdkScreen is used throughout GDK and GTK+ to specify which screen the top
// level windows are to be displayed on. it is also used to query the screen
// specification and default settings such as the default visual
// (gdk_screen_get_system_visual()), the dimensions of the physical monitors
// (gdk_screen_get_monitor_geometry()), etc.
type Screen struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Screen)(nil)
)

func wrapScreen(obj *coreglib.Object) *Screen {
	return &Screen{
		Object: obj,
	}
}

func marshalScreen(p uintptr) (interface{}, error) {
	return wrapScreen(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk3_Screen_ConnectCompositedChanged
func _gotk4_gdk3_Screen_ConnectCompositedChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectCompositedChanged signal is emitted when the composited status of the
// screen changes.
func (screen *Screen) ConnectCompositedChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "composited-changed", false, unsafe.Pointer(C._gotk4_gdk3_Screen_ConnectCompositedChanged), f)
}

//export _gotk4_gdk3_Screen_ConnectMonitorsChanged
func _gotk4_gdk3_Screen_ConnectMonitorsChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectMonitorsChanged signal is emitted when the number, size or position of
// the monitors attached to the screen change.
//
// Only for X11 and OS X for now. A future implementation for Win32 may be a
// possibility.
func (screen *Screen) ConnectMonitorsChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "monitors-changed", false, unsafe.Pointer(C._gotk4_gdk3_Screen_ConnectMonitorsChanged), f)
}

//export _gotk4_gdk3_Screen_ConnectSizeChanged
func _gotk4_gdk3_Screen_ConnectSizeChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectSizeChanged signal is emitted when the pixel width or height of a
// screen changes.
func (screen *Screen) ConnectSizeChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(screen, "size-changed", false, unsafe.Pointer(C._gotk4_gdk3_Screen_ConnectSizeChanged), f)
}

// ActiveWindow returns the screen’s currently active window.
//
// On X11, this is done by inspecting the _NET_ACTIVE_WINDOW property on the
// root window, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec). If there is no currently
// currently active window, or the window manager does not support the
// _NET_ACTIVE_WINDOW hint, this function returns NULL.
//
// On other platforms, this function may return NULL, depending on whether it is
// implementable on that platform.
//
// The returned window should be unrefed using g_object_unref() when no longer
// needed.
//
// Deprecated: since version 3.22.
//
// The function returns the following values:
//
//    - window (optional): currently active window, or NULL.
//
func (screen *Screen) ActiveWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_active_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _window Windower // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Windower)
				return ok
			})
			rv, ok := casted.(Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// Display gets the display to which the screen belongs.
//
// The function returns the following values:
//
//    - display to which screen belongs.
//
func (screen *Screen) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_display", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _display
}

// FontOptions gets any options previously set with
// gdk_screen_set_font_options().
//
// The function returns the following values:
//
//    - fontOptions (optional): current font options, or NULL if no default font
//      options have been set.
//
func (screen *Screen) FontOptions() *cairo.FontOptions {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_font_options", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _fontOptions *cairo.FontOptions // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_fontOptions = (*cairo.FontOptions)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _fontOptions
}

// Height gets the height of screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information instead.
//
// The function returns the following values:
//
//    - gint: height of screen in pixels.
//
func (screen *Screen) Height() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_height", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// HeightMm returns the height of screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen has multiple
// monitors of different resolution. It is recommended to use the monitor
// dimensions instead.
//
// Deprecated: Use per-monitor information instead.
//
// The function returns the following values:
//
//    - gint: heigth of screen in millimeters.
//
func (screen *Screen) HeightMm() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_height_mm", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorAtPoint returns the monitor number in which the point (x,y) is
// located.
//
// Deprecated: Use gdk_display_get_monitor_at_point() instead.
//
// The function takes the following parameters:
//
//    - x coordinate in the virtual screen.
//    - y coordinate in the virtual screen.
//
// The function returns the following values:
//
//    - gint: monitor number in which the point (x,y) lies, or a monitor close to
//      (x,y) if the point is not in any monitor.
//
func (screen *Screen) MonitorAtPoint(x, y int32) int32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(x)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(y)

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_at_point", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorAtWindow returns the number of the monitor in which the largest area
// of the bounding rectangle of window resides.
//
// Deprecated: Use gdk_display_get_monitor_at_window() instead.
//
// The function takes the following parameters:
//
//    - window: Window.
//
// The function returns the following values:
//
//    - gint: monitor number in which most of window is located, or if window
//      does not intersect any monitors, a monitor, close to window.
//
func (screen *Screen) MonitorAtWindow(window Windower) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_at_window", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(window)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorGeometry retrieves the Rectangle representing the size and position of
// the individual monitor within the entire screen area. The returned geometry
// is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Monitor numbers start at 0. To obtain the number of monitors of screen, use
// gdk_screen_get_n_monitors().
//
// Note that the size of the entire screen area can be retrieved via
// gdk_screen_get_width() and gdk_screen_get_height().
//
// Deprecated: Use gdk_monitor_get_geometry() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: monitor number.
//
// The function returns the following values:
//
//    - dest (optional) to be filled with the monitor geometry.
//
func (screen *Screen) MonitorGeometry(monitorNum int32) *Rectangle {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_info.InvokeClassMethod("get_monitor_geometry", _args[:], _outs[:])

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _dest *Rectangle // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	}

	return _dest
}

// MonitorHeightMm gets the height in millimeters of the specified monitor.
//
// Deprecated: Use gdk_monitor_get_height_mm() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: number of the monitor, between 0 and
//      gdk_screen_get_n_monitors (screen).
//
// The function returns the following values:
//
//    - gint: height of the monitor, or -1 if not available.
//
func (screen *Screen) MonitorHeightMm(monitorNum int32) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_height_mm", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorPlugName returns the output name of the specified monitor. Usually
// something like VGA, DVI, or TV, not the actual product name of the display
// device.
//
// Deprecated: Use gdk_monitor_get_model() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: number of the monitor, between 0 and
//      gdk_screen_get_n_monitors (screen).
//
// The function returns the following values:
//
//    - utf8 (optional): newly-allocated string containing the name of the
//      monitor, or NULL if the name cannot be determined.
//
func (screen *Screen) MonitorPlugName(monitorNum int32) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_plug_name", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _utf8 string // out

	if *(**C.gchar)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
		defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret))))
	}

	return _utf8
}

// MonitorScaleFactor returns the internal scale factor that maps from monitor
// coordinates to the actual device pixels. On traditional systems this is 1,
// but on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a window where it is better
// to use gdk_window_get_scale_factor() instead.
//
// Deprecated: Use gdk_monitor_get_scale_factor() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: number of the monitor, between 0 and
//      gdk_screen_get_n_monitors (screen).
//
// The function returns the following values:
//
//    - gint: scale factor.
//
func (screen *Screen) MonitorScaleFactor(monitorNum int32) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_scale_factor", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorWidthMm gets the width in millimeters of the specified monitor, if
// available.
//
// Deprecated: Use gdk_monitor_get_width_mm() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: number of the monitor, between 0 and
//      gdk_screen_get_n_monitors (screen).
//
// The function returns the following values:
//
//    - gint: width of the monitor, or -1 if not available.
//
func (screen *Screen) MonitorWidthMm(monitorNum int32) int32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_monitor_width_mm", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MonitorWorkarea retrieves the Rectangle representing the size and position of
// the “work area” on a monitor within the entire screen area. The returned
// geometry is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// The work area should be considered when positioning menus and similar popups,
// to avoid placing them below panels, docks or other desktop components.
//
// Note that not all backends may have a concept of workarea. This function will
// return the monitor geometry if a workarea is not available, or does not
// apply.
//
// Monitor numbers start at 0. To obtain the number of monitors of screen, use
// gdk_screen_get_n_monitors().
//
// Deprecated: Use gdk_monitor_get_workarea() instead.
//
// The function takes the following parameters:
//
//    - monitorNum: monitor number.
//
// The function returns the following values:
//
//    - dest (optional) to be filled with the monitor workarea.
//
func (screen *Screen) MonitorWorkarea(monitorNum int32) *Rectangle {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(monitorNum)

	_info := girepository.MustFind("Gdk", "Screen")
	_info.InvokeClassMethod("get_monitor_workarea", _args[:], _outs[:])

	runtime.KeepAlive(screen)
	runtime.KeepAlive(monitorNum)

	var _dest *Rectangle // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	}

	return _dest
}

// NMonitors returns the number of monitors which screen consists of.
//
// Deprecated: Use gdk_display_get_n_monitors() instead.
//
// The function returns the following values:
//
//    - gint: number of monitors which screen consists of.
//
func (screen *Screen) NMonitors() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_n_monitors", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Number gets the index of screen among the screens in the display to which it
// belongs. (See gdk_screen_get_display())
//
// Deprecated: since version 3.22.
//
// The function returns the following values:
//
//    - gint: index.
//
func (screen *Screen) Number() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_number", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// PrimaryMonitor gets the primary monitor for screen. The primary monitor is
// considered the monitor where the “main desktop” lives. While normal
// application windows typically allow the window manager to place the windows,
// specialized desktop applications such as panels should place themselves on
// the primary monitor.
//
// If no primary monitor is configured by the user, the return value will be 0,
// defaulting to the first monitor.
//
// Deprecated: Use gdk_display_get_primary_monitor() instead.
//
// The function returns the following values:
//
//    - gint: integer index for the primary monitor, or 0 if none is configured.
//
func (screen *Screen) PrimaryMonitor() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_primary_monitor", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Resolution gets the resolution for font handling on the screen; see
// gdk_screen_set_resolution() for full details.
//
// The function returns the following values:
//
//    - gdouble: current resolution, or -1 if no resolution has been set.
//
func (screen *Screen) Resolution() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_resolution", _args[:], nil)
	_cret := *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// RGBAVisual gets a visual to use for creating windows with an alpha channel.
// The windowing system on which GTK+ is running may not support this
// capability, in which case NULL will be returned. Even if a non-NULL value is
// returned, its possible that the window’s alpha channel won’t be honored when
// displaying the window on the screen: in particular, for X an appropriate
// windowing manager and compositing manager must be running to provide
// appropriate display.
//
// This functionality is not implemented in the Windows backend.
//
// For setting an overall opacity for a top-level window, see
// gdk_window_set_opacity().
//
// The function returns the following values:
//
//    - visual (optional) to use for windows with an alpha channel or NULL if the
//      capability is not available.
//
func (screen *Screen) RGBAVisual() *Visual {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_rgba_visual", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _visual *Visual // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_visual = wrapVisual(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _visual
}

// RootWindow gets the root window of screen.
//
// The function returns the following values:
//
//    - window: root window.
//
func (screen *Screen) RootWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_root_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// Setting retrieves a desktop-wide setting such as double-click time for the
// Screen screen.
//
// FIXME needs a list of valid settings here, or a link to more information.
//
// The function takes the following parameters:
//
//    - name of the setting.
//    - value: location to store the value of the setting.
//
// The function returns the following values:
//
//    - ok: TRUE if the setting existed and a value was stored in value, FALSE
//      otherwise.
//
func (screen *Screen) Setting(name string, value *coreglib.Value) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))
	*(**C.GValue)(unsafe.Pointer(&_args[2])) = (*C.GValue)(unsafe.Pointer(value.Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_setting", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SystemVisual: get the system’s default visual for screen. This is the visual
// for the root window of the display. The return value should not be freed.
//
// The function returns the following values:
//
//    - visual: system visual.
//
func (screen *Screen) SystemVisual() *Visual {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_system_visual", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _visual *Visual // out

	_visual = wrapVisual(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _visual
}

// ToplevelWindows obtains a list of all toplevel windows known to GDK on the
// screen screen. A toplevel window is a child of the root window (see
// gdk_get_default_root_window()).
//
// The returned list should be freed with g_list_free(), but its elements need
// not be freed.
//
// The function returns the following values:
//
//    - list: list of toplevel windows, free with g_list_free().
//
func (screen *Screen) ToplevelWindows() []Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_toplevel_windows", _args[:], nil)
	_cret := *(**C.GList)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _list []Windower // out

	_list = make([]Windower, 0, gextras.ListSize(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret)))))
	gextras.MoveList(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst Windower // out
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))
			if objptr == nil {
				panic("object of type gdk.Windower is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Windower)
				return ok
			})
			rv, ok := casted.(Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// Width gets the width of screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information instead.
//
// The function returns the following values:
//
//    - gint: width of screen in pixels.
//
func (screen *Screen) Width() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_width", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// WidthMm gets the width of screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen has multiple
// monitors of different resolution. It is recommended to use the monitor
// dimensions instead.
//
// Deprecated: Use per-monitor information instead.
//
// The function returns the following values:
//
//    - gint: width of screen in millimeters.
//
func (screen *Screen) WidthMm() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_width_mm", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// WindowStack returns a #GList of Windows representing the current window
// stack.
//
// On X11, this is done by inspecting the _NET_CLIENT_LIST_STACKING property on
// the root window, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec). If the window manager does
// not support the _NET_CLIENT_LIST_STACKING hint, this function returns NULL.
//
// On other platforms, this function may return NULL, depending on whether it is
// implementable on that platform.
//
// The returned list is newly allocated and owns references to the windows it
// contains, so it should be freed using g_list_free() and its windows unrefed
// using g_object_unref() when no longer needed.
//
// The function returns the following values:
//
//    - list (optional): a list of Windows for the current window stack, or NULL.
//
func (screen *Screen) WindowStack() []Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("get_window_stack", _args[:], nil)
	_cret := *(**C.GList)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _list []Windower // out

	if *(**C.GList)(unsafe.Pointer(&_cret)) != nil {
		_list = make([]Windower, 0, gextras.ListSize(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret)))))
		gextras.MoveList(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
			src := (*C.void)(v)
			var dst Windower // out
			{
				objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))
				if objptr == nil {
					panic("object of type gdk.Windower is nil")
				}

				object := coreglib.AssumeOwnership(objptr)
				casted := object.WalkCast(func(obj coreglib.Objector) bool {
					_, ok := obj.(Windower)
					return ok
				})
				rv, ok := casted.(Windower)
				if !ok {
					panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
				}
				dst = rv
			}
			_list = append(_list, dst)
		})
	}

	return _list
}

// IsComposited returns whether windows with an RGBA visual can reasonably be
// expected to have their alpha channel drawn correctly on the screen.
//
// On X11 this function returns whether a compositing manager is compositing
// screen.
//
// The function returns the following values:
//
//    - ok: whether windows with RGBA visuals can reasonably be expected to have
//      their alpha channels drawn correctly on the screen.
//
func (screen *Screen) IsComposited() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("is_composited", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ListVisuals lists the available visuals for the specified screen. A visual
// describes a hardware image data format. For example, a visual might support
// 24-bit color, or 8-bit color, and might expect pixels to be in a certain
// format.
//
// Call g_list_free() on the return value when you’re finished with it.
//
// The function returns the following values:
//
//    - list: a list of visuals; the list must be freed, but not its contents.
//
func (screen *Screen) ListVisuals() []*Visual {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("list_visuals", _args[:], nil)
	_cret := *(**C.GList)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _list []*Visual // out

	_list = make([]*Visual, 0, gextras.ListSize(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret)))))
	gextras.MoveList(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *Visual // out
		dst = wrapVisual(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))))
		_list = append(_list, dst)
	})

	return _list
}

// MakeDisplayName determines the name to pass to gdk_display_open() to get a
// Display with this screen as the default screen.
//
// Deprecated: since version 3.22.
//
// The function returns the following values:
//
//    - utf8: newly allocated string, free with g_free().
//
func (screen *Screen) MakeDisplayName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gdk", "Screen")
	_gret := _info.InvokeClassMethod("make_display_name", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret))))

	return _utf8
}

// SetFontOptions sets the default font options for the screen. These options
// will be set on any Context’s newly created with
// gdk_pango_context_get_for_screen(). Changing the default set of font options
// does not affect contexts that have already been created.
//
// The function takes the following parameters:
//
//    - options (optional) or NULL to unset any previously set default font
//      options.
//
func (screen *Screen) SetFontOptions(options *cairo.FontOptions) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	if options != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(options)))
	}

	_info := girepository.MustFind("Gdk", "Screen")
	_info.InvokeClassMethod("set_font_options", _args[:], nil)

	runtime.KeepAlive(screen)
	runtime.KeepAlive(options)
}

// SetResolution sets the resolution for font handling on the screen. This is a
// scale factor between points specified in a FontDescription and cairo units.
// The default value is 96, meaning that a 10 point font will be 13 units high.
// (10 * 96. / 72. = 13.3).
//
// The function takes the following parameters:
//
//    - dpi: resolution in “dots per inch”. (Physical inches aren’t actually
//      involved; the terminology is conventional.).
//
func (screen *Screen) SetResolution(dpi float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(dpi)

	_info := girepository.MustFind("Gdk", "Screen")
	_info.InvokeClassMethod("set_resolution", _args[:], nil)

	runtime.KeepAlive(screen)
	runtime.KeepAlive(dpi)
}

// ScreenGetDefault gets the default screen for the default display. (See
// gdk_display_get_default ()).
//
// The function returns the following values:
//
//    - screen (optional) or NULL if there is no default display.
//
func ScreenGetDefault() *Screen {
	_info := girepository.MustFind("Gdk", "get_default")
	_gret := _info.InvokeFunction(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _screen *Screen // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_screen = wrapScreen(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _screen
}
