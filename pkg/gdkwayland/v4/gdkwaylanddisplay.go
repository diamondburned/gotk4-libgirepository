// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkwaylanddisplay.go.
var GTypeWaylandDisplay = coreglib.Type(C.gdk_wayland_display_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeWaylandDisplay, F: marshalWaylandDisplay},
	})
}

// WaylandDisplay: wayland implementation of GdkDisplay.
//
// Beyond the regular gdk.Display API, the Wayland implementation provides
// access to Wayland objects such as the wl_display with
// gdkwayland.WaylandDisplay.GetWlDisplay(), the wl_compositor with
// gdkwayland.WaylandDisplay.GetWlCompositor().
//
// You can find out what Wayland globals are supported by a display with
// gdkwayland.WaylandDisplay.QueryRegistry().
type WaylandDisplay struct {
	_ [0]func() // equal guard
	gdk.Display
}

var (
	_ coreglib.Objector = (*WaylandDisplay)(nil)
)

func wrapWaylandDisplay(obj *coreglib.Object) *WaylandDisplay {
	return &WaylandDisplay{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	return wrapWaylandDisplay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or NULL if no ID has been defined.
//
// The function returns the following values:
//
//    - utf8 (optional): startup notification ID for display, or NULL.
//
func (display *WaylandDisplay) StartupNotificationID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("GdkWayland", "WaylandDisplay").InvokeMethod("get_startup_notification_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// QueryRegistry returns TRUE if the the interface was found in the display
// wl_registry.global handler.
//
// The function takes the following parameters:
//
//    - global interface to query in the registry.
//
// The function returns the following values:
//
//    - ok: TRUE if the global is offered by the compositor.
//
func (display *WaylandDisplay) QueryRegistry(global string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(global)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("GdkWayland", "WaylandDisplay").InvokeMethod("query_registry", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)
	runtime.KeepAlive(global)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetCursorTheme sets the cursor theme for the given display.
//
// The function takes the following parameters:
//
//    - name: new cursor theme.
//    - size to use for cursors.
//
func (display *WaylandDisplay) SetCursorTheme(name string, size int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(size)

	girepository.MustFind("GdkWayland", "WaylandDisplay").InvokeMethod("set_cursor_theme", _args[:], nil)

	runtime.KeepAlive(display)
	runtime.KeepAlive(name)
	runtime.KeepAlive(size)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk.Display.NotifyStartupComplete()).
//
// The function takes the following parameters:
//
//    - startupId: startup notification ID (must be valid utf8).
//
func (display *WaylandDisplay) SetStartupNotificationID(startupId string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("GdkWayland", "WaylandDisplay").InvokeMethod("set_startup_notification_id", _args[:], nil)

	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}
