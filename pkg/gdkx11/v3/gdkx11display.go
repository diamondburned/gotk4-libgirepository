// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11display.go.
var GTypeX11Display = externglib.Type(C.gdk_x11_display_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11Display, F: marshalX11Display},
	})
}

// X11RegisterStandardEventType registers interest in receiving extension events
// with type codes between event_base and event_base + n_events - 1. The
// registered events must have the window field in the same place as core X
// events (this is not the case for e.g. XKB extension events).
//
// If an event type is registered, events of this type will go through global
// and window-specific filters (see gdk_window_add_filter()). Unregistered
// events will only go through global filters. GDK may register the events of
// some X extensions on its own.
//
// This function should only be needed in unusual circumstances, e.g. when
// filtering XInput extension events on the root window.
//
// The function takes the following parameters:
//
//    - display: Display.
//    - eventBase: first event type code to register.
//    - nEvents: number of event type codes to register.
//
func X11RegisterStandardEventType(display *X11Display, eventBase, nEvents int) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg2 = C.gint(eventBase)
	_arg3 = C.gint(nEvents)

	C.gdk_x11_register_standard_event_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(display)
	runtime.KeepAlive(eventBase)
	runtime.KeepAlive(nEvents)
}

// X11SetSmClientID sets the SM_CLIENT_ID property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual.
//
// The function takes the following parameters:
//
//    - smClientId (optional): client id assigned by the session manager when the
//      connection was opened, or NULL to remove the property.
//
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.gchar // out

	if smClientId != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(smClientId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gdk_x11_set_sm_client_id(_arg1)
	runtime.KeepAlive(smClientId)
}

// X11DisplayOverrider contains methods that are overridable.
type X11DisplayOverrider interface {
	externglib.Objector
}

type X11Display struct {
	_ [0]func() // equal guard
	gdk.Display
}

var (
	_ externglib.Objector = (*X11Display)(nil)
)

func classInitX11Displayer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Display(obj *externglib.Object) *X11Display {
	return &X11Display{
		Display: gdk.Display{
			Object: obj,
		},
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	return wrapX11Display(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ErrorTrapPop pops the error trap pushed by gdk_x11_display_error_trap_push().
// Will XSync() if necessary and will always block until the error is known to
// have occurred or not occurred, so the error code can be returned.
//
// If you don’t need to use the return value,
// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
//
// See gdk_error_trap_pop() for the all-displays-at-once equivalent.
//
// The function returns the following values:
//
//    - gint: x error code or 0 on success.
//
func (display *X11Display) ErrorTrapPop() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)
	runtime.KeepAlive(display)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ErrorTrapPopIgnored pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Does not block to see if an error
// occurred; merely records the range of requests to ignore errors for, and
// ignores those errors if they arrive asynchronously.
//
// See gdk_error_trap_pop_ignored() for the all-displays-at-once equivalent.
func (display *X11Display) ErrorTrapPopIgnored() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
	runtime.KeepAlive(display)
}

// ErrorTrapPush begins a range of X requests on display for which X error
// events will be ignored. Unignored errors (when no trap is pushed) will abort
// the application. Use gdk_x11_display_error_trap_pop() or
// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
// function.
//
// See also gdk_error_trap_push() to push a trap on all displays.
func (display *X11Display) ErrorTrapPush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
	runtime.KeepAlive(display)
}

// StartupNotificationID gets the startup notification ID for a display.
//
// The function returns the following values:
//
//    - utf8: startup notification ID for display.
//
func (display *X11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UserTime returns the timestamp of the last user interaction on display. The
// timestamp is taken from events caused by user interaction such as key presses
// or pointer movements. See gdk_x11_window_set_user_time().
//
// The function returns the following values:
//
//    - guint32: timestamp of the last user interaction.
//
func (display *X11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)
	runtime.KeepAlive(display)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Grab: call XGrabServer() on display. To ungrab the display again, use
// gdk_x11_display_ungrab().
//
// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
func (display *X11Display) Grab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_x11_display_grab(_arg0)
	runtime.KeepAlive(display)
}

// SetCursorTheme sets the cursor theme from which the images for cursor should
// be taken.
//
// If the windowing system supports it, existing cursors created with
// gdk_cursor_new(), gdk_cursor_new_for_display() and gdk_cursor_new_from_name()
// are updated to reflect the theme change. Custom cursors constructed with
// gdk_cursor_new_from_pixbuf() will have to be handled by the application (GTK+
// applications can learn about cursor theme changes by listening for change
// notification for the corresponding Setting).
//
// The function takes the following parameters:
//
//    - theme (optional): name of the cursor theme to use, or NULL to unset a
//      previously set value.
//    - size: cursor size to use, or 0 to keep the previous size.
//
func (display *X11Display) SetCursorTheme(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	if theme != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(theme)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.gint(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(theme)
	runtime.KeepAlive(size)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID environment
// variable, but in some cases (such as the application not being launched using
// exec()) it can come from other sources.
//
// If the ID contains the string "_TIME" then the portion following that string
// is taken to be the X11 timestamp of the event that triggered the application
// to be launched and the GDK current event time is set accordingly.
//
// The startup ID is also what is used to signal that the startup is complete
// (for example, when opening a window or when calling
// gdk_notify_startup_complete()).
//
// The function takes the following parameters:
//
//    - startupId: startup notification ID (must be valid utf8).
//
func (display *X11Display) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}

// SetWindowScale forces a specific window scale for all windows on this
// display, instead of using the default or user configured scale. This is can
// be used to disable scaling support by setting scale to 1, or to
// programmatically set the window scale.
//
// Once the scale is set by this call it will not change in response to later
// user configuration changes.
//
// The function takes the following parameters:
//
//    - scale: new scale value.
//
func (display *X11Display) SetWindowScale(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = C.gint(scale)

	C.gdk_x11_display_set_window_scale(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(scale)
}

// Ungrab display after it has been grabbed with gdk_x11_display_grab().
func (display *X11Display) Ungrab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_x11_display_ungrab(_arg0)
	runtime.KeepAlive(display)
}
