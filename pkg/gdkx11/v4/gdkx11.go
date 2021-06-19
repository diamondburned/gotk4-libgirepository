// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_app_launch_context_get_type()), F: marshalX11AppLaunchContext},
		{T: externglib.Type(C.gdk_x11_device_manager_xi2_get_type()), F: marshalX11DeviceManagerXI2},
		{T: externglib.Type(C.gdk_x11_device_xi2_get_type()), F: marshalX11DeviceXI2},
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
		{T: externglib.Type(C.gdk_x11_drag_get_type()), F: marshalX11Drag},
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitor},
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screen},
		{T: externglib.Type(C.gdk_x11_surface_get_type()), F: marshalX11Surface},
	})
}

type X11DeviceType int

const (
	X11DeviceTypeLogical  X11DeviceType = 0
	X11DeviceTypePhysical X11DeviceType = 1
	X11DeviceTypeFloating X11DeviceType = 2
)

// X11DeviceGetID returns the device ID as seen by XInput2.
func X11DeviceGetID(device X11DeviceXI2) int {
	var _arg1 *C.GdkDevice // out
	var _cret C.int        // in

	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_x11_device_get_id(_arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerXI2, deviceId int) X11DeviceXI2 {
	var _arg1 *C.GdkX11DeviceManagerXI2 // out
	var _arg2 C.int                     // out
	var _cret *C.GdkDevice              // in

	_arg1 = (*C.GdkX11DeviceManagerXI2)(unsafe.Pointer(deviceManager.Native()))
	_arg2 = (C.int)(deviceId)

	_cret = C.gdk_x11_device_manager_lookup(_arg1, _arg2)

	var _x11DeviceXI2 X11DeviceXI2 // out

	_x11DeviceXI2 = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(X11DeviceXI2)

	return _x11DeviceXI2
}

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
func X11FreeCompoundText(ctext *byte) {
	var _arg1 *C.guchar // out

	_arg1 = (*C.guchar)(unsafe.Pointer(ctext))

	C.gdk_x11_free_compound_text(_arg1)
}

// X11FreeTextList frees the array of strings created by
// gdk_x11_display_text_property_to_text_list().
func X11FreeTextList(list *string) {
	var _arg1 **C.char // out

	_arg1 = (**C.char)(C.CString(list))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_free_text_list(_arg1)
}

// X11GetServerTime: routine to get the current X server time stamp.
func X11GetServerTime(surface X11Surface) uint32 {
	var _arg1 *C.GdkSurface // out
	var _cret C.guint32     // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_x11_get_server_time(_arg1)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(smClientId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_set_sm_client_id(_arg1)
}

type X11AppLaunchContext interface {
}

// x11AppLaunchContext implements the X11AppLaunchContext class.
type x11AppLaunchContext struct {
	gdk.AppLaunchContext
}

// WrapX11AppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11AppLaunchContext(obj *externglib.Object) X11AppLaunchContext {
	return x11AppLaunchContext{
		AppLaunchContext: gdk.WrapAppLaunchContext(obj),
	}
}

func marshalX11AppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11AppLaunchContext(obj), nil
}

type X11DeviceManagerXI2 interface {
}

// x11DeviceManagerXI2 implements the X11DeviceManagerXI2 class.
type x11DeviceManagerXI2 struct {
	gextras.Objector
}

// WrapX11DeviceManagerXI2 wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceManagerXI2(obj *externglib.Object) X11DeviceManagerXI2 {
	return x11DeviceManagerXI2{
		Objector: obj,
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceManagerXI2(obj), nil
}

type X11DeviceXI2 interface {
}

// x11DeviceXI2 implements the X11DeviceXI2 class.
type x11DeviceXI2 struct {
	gdk.Device
}

// WrapX11DeviceXI2 wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceXI2(obj *externglib.Object) X11DeviceXI2 {
	return x11DeviceXI2{
		Device: gdk.WrapDevice(obj),
	}
}

func marshalX11DeviceXI2(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceXI2(obj), nil
}

type X11Display interface {

	// ErrorTrapPopX11Display pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	ErrorTrapPopX11Display() int
	// ErrorTrapPopIgnoredX11Display pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	ErrorTrapPopIgnoredX11Display()
	// ErrorTrapPushX11Display begins a range of X requests on @display for
	// which X error events will be ignored. Unignored errors (when no trap is
	// pushed) will abort the application. Use gdk_x11_display_error_trap_pop()
	// or gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with
	// this function.
	ErrorTrapPushX11Display()
	// DefaultGroup returns the default group leader surface for all toplevel
	// surfaces on @display. This surface is implicitly created by GDK. See
	// gdk_x11_surface_set_group().
	DefaultGroup() gdk.Surface
	// GlxVersion retrieves the version of the GLX implementation.
	GlxVersion() (major int, minor int, ok bool)
	// PrimaryMonitor gets the primary monitor for the display.
	//
	// The primary monitor is considered the monitor where the “main desktop”
	// lives. While normal application surfaces typically allow the window
	// manager to place the surfaces, specialized desktop applications such as
	// panels should place themselves on the primary monitor.
	//
	// If no monitor is the designated primary monitor, any monitor (usually the
	// first) may be returned.
	PrimaryMonitor() gdk.Monitor
	// Screen retrieves the X11Screen of the @display.
	Screen() X11Screen
	// StartupNotificationID gets the startup notification ID for a display.
	StartupNotificationID() string
	// UserTime returns the timestamp of the last user interaction on @display.
	// The timestamp is taken from events caused by user interaction such as key
	// presses or pointer movements. See gdk_x11_surface_set_user_time().
	UserTime() uint32
	// GrabX11Display: call XGrabServer() on @display. To ungrab the display
	// again, use gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	GrabX11Display()
	// SetCursorThemeX11Display sets the cursor theme from which the images for
	// cursor should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_texture() will have
	// to be handled by the application (GTK applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorThemeX11Display(theme string, size int)
	// SetStartupNotificationIDX11Display sets the startup notification ID for a
	// display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// If the ID contains the string "_TIME" then the portion following that
	// string is taken to be the X11 timestamp of the event that triggered the
	// application to be launched and the GDK current event time is set
	// accordingly.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_display_notify_startup_complete()).
	SetStartupNotificationIDX11Display(startupId string)
	// SetSurfaceScaleX11Display forces a specific window scale for all windows
	// on this display, instead of using the default or user configured scale.
	// This is can be used to disable scaling support by setting @scale to 1, or
	// to programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetSurfaceScaleX11Display(scale int)
	// StringToCompoundTextX11Display: convert a string from the encoding of the
	// current locale into a form suitable for storing in a window property.
	StringToCompoundTextX11Display(str string) (encoding string, format int, ctext []byte, gint int)
	// TextPropertyToTextListX11Display: convert a text string from the encoding
	// as it is stored in a property into an array of strings in the encoding of
	// the current locale. (The elements of the array represent the
	// nul-separated elements of the original text string.)
	TextPropertyToTextListX11Display(encoding string, format int, text *byte, length int, list **string) int
	// UngrabX11Display: ungrab @display after it has been grabbed with
	// gdk_x11_display_grab().
	UngrabX11Display()
	// UTF8ToCompoundTextX11Display converts from UTF-8 to compound text.
	UTF8ToCompoundTextX11Display(str string) (string, int, []byte, bool)
}

// x11Display implements the X11Display class.
type x11Display struct {
	gdk.Display
}

// WrapX11Display wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return x11Display{
		Display: gdk.WrapDisplay(obj),
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

func (d x11Display) ErrorTrapPopX11Display() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (d x11Display) ErrorTrapPopIgnoredX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
}

func (d x11Display) ErrorTrapPushX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_push(_arg0)
}

func (d x11Display) DefaultGroup() gdk.Surface {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_default_group(_arg0)

	var _surface gdk.Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Surface)

	return _surface
}

func (d x11Display) GlxVersion() (major int, minor int, ok bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // in
	var _arg2 C.int         // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = (int)(_arg1)
	_minor = (int)(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

func (d x11Display) PrimaryMonitor() gdk.Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_primary_monitor(_arg0)

	var _monitor gdk.Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Monitor)

	return _monitor
}

func (d x11Display) Screen() X11Screen {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkX11Screen // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_screen(_arg0)

	var _x11Screen X11Screen // out

	_x11Screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(X11Screen)

	return _x11Screen
}

func (d x11Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d x11Display) UserTime() uint32 {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_get_user_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (d x11Display) GrabX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(_arg0)
}

func (d x11Display) SetCursorThemeX11Display(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(theme))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.int)(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d x11Display) SetStartupNotificationIDX11Display(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
}

func (d x11Display) SetSurfaceScaleX11Display(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (C.int)(scale)

	C.gdk_x11_display_set_surface_scale(_arg0, _arg1)
}

func (d x11Display) StringToCompoundTextX11Display(str string) (encoding string, format int, ctext []byte, gint int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int // in
	var _cret C.int // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_string_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _gint int // out

	_encoding = C.GoString(_arg2)
	_format = (int)(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_gint = (int)(_cret)

	return _encoding, _format, _ctext, _gint
}

func (d x11Display) TextPropertyToTextListX11Display(encoding string, format int, text *byte, length int, list **string) int {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out
	var _arg3 *C.guchar     // out
	var _arg4 C.int         // out
	var _arg5 ***C.char     // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(encoding))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.int)(format)
	_arg3 = (*C.guchar)(unsafe.Pointer(text))
	_arg4 = (C.int)(length)
	_arg5 = (***C.char)(C.CString(list))
	defer C.free(unsafe.Pointer(_arg5))

	_cret = C.gdk_x11_display_text_property_to_text_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (d x11Display) UngrabX11Display() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(_arg0)
}

func (d x11Display) UTF8ToCompoundTextX11Display(str string) (string, int, []byte, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // in
	var _arg3 C.int         // in
	var _arg4 *C.guchar
	var _arg5 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_display_utf8_to_compound_text(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _encoding string // out
	var _format int      // out
	var _ctext []byte
	var _ok bool // out

	_encoding = C.GoString(_arg2)
	_format = (int)(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}

type X11Drag interface {
}

// x11Drag implements the X11Drag class.
type x11Drag struct {
	gdk.Drag
}

// WrapX11Drag wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Drag(obj *externglib.Object) X11Drag {
	return x11Drag{
		Drag: gdk.WrapDrag(obj),
	}
}

func marshalX11Drag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Drag(obj), nil
}

type X11Monitor interface {

	// Workarea retrieves the size and position of the “work area” on a monitor
	// within the display coordinate space. The returned geometry is in
	// ”application pixels”, not in ”device pixels” (see
	// gdk_monitor_get_scale_factor()).
	Workarea() gdk.Rectangle
}

// x11Monitor implements the X11Monitor class.
type x11Monitor struct {
	gdk.Monitor
}

// WrapX11Monitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Monitor(obj *externglib.Object) X11Monitor {
	return x11Monitor{
		Monitor: gdk.WrapMonitor(obj),
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Monitor(obj), nil
}

func (m x11Monitor) Workarea() gdk.Rectangle {
	var _arg0 *C.GdkMonitor // out
	var _workarea gdk.Rectangle

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	C.gdk_x11_monitor_get_workarea(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_workarea)))

	return _workarea
}

type X11Screen interface {

	// CurrentDesktop returns the current workspace for @screen when running
	// under a window manager that supports multiple workspaces, as described in
	// the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	CurrentDesktop() uint32
	// NumberOfDesktops returns the number of workspaces for @screen when
	// running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	NumberOfDesktops() uint32
	// ScreenNumber returns the index of a X11Screen.
	ScreenNumber() int
	// WindowManagerName returns the name of the window manager for @screen.
	WindowManagerName() string
	// SupportsNetWmHintX11Screen: this function is specific to the X11 backend
	// of GDK, and indicates whether the window manager supports a certain hint
	// from the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	//
	// When using this function, keep in mind that the window manager can change
	// over time; so you shouldn’t use this function in a way that impacts
	// persistent application state. A common bug is that your application can
	// start up before the window manager does when the user logs in, and before
	// the window manager starts gdk_x11_screen_supports_net_wm_hint() will
	// return false for every property. You can monitor the
	// window_manager_changed signal on X11Screen to detect a window manager
	// change.
	SupportsNetWmHintX11Screen(propertyName string) bool
}

// x11Screen implements the X11Screen class.
type x11Screen struct {
	gextras.Objector
}

// WrapX11Screen wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Screen(obj *externglib.Object) X11Screen {
	return x11Screen{
		Objector: obj,
	}
}

func marshalX11Screen(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Screen(obj), nil
}

func (s x11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (s x11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (s x11Screen) ScreenNumber() int {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.int           // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

func (s x11Screen) WindowManagerName() string {
	var _arg0 *C.GdkX11Screen // out
	var _cret *C.char         // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s x11Screen) SupportsNetWmHintX11Screen(propertyName string) bool {
	var _arg0 *C.GdkX11Screen // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type X11Surface interface {

	// Desktop gets the number of the workspace @surface is on.
	Desktop() uint32
	// Group returns the group this surface belongs to.
	Group() gdk.Surface
	// MoveToCurrentDesktopX11Surface moves the surface to the correct workspace
	// when running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification. Will not do
	// anything if the surface is already on all workspaces.
	MoveToCurrentDesktopX11Surface()
	// MoveToDesktopX11Surface moves the surface to the given workspace when
	// running unde a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	MoveToDesktopX11Surface(desktop uint32)
	// SetFrameSyncEnabledX11Surface: this function can be used to disable frame
	// synchronization for a surface. Normally frame synchronziation will be
	// enabled or disabled based on whether the system has a compositor that
	// supports frame synchronization, but if the surface is not directly
	// managed by the window manager, then frame synchronziation may need to be
	// disabled. This is the case for a surface embedded via the XEMBED
	// protocol.
	SetFrameSyncEnabledX11Surface(frameSyncEnabled bool)
	// SetGroupX11Surface sets the group leader of @surface to be @leader. See
	// the ICCCM for details.
	SetGroupX11Surface(leader gdk.Surface)
	// SetSkipPagerHintX11Surface sets a hint on @surface that pagers should not
	// display it. See the EWMH for details.
	SetSkipPagerHintX11Surface(skipsPager bool)
	// SetSkipTaskbarHintX11Surface sets a hint on @surface that taskbars should
	// not display it. See the EWMH for details.
	SetSkipTaskbarHintX11Surface(skipsTaskbar bool)
	// SetThemeVariantX11Surface: GTK applications can request a dark theme
	// variant. In order to make other applications - namely window managers
	// using GTK for themeing - aware of this choice, GTK uses this function to
	// export the requested theme variant as _GTK_THEME_VARIANT property on
	// toplevel surfaces.
	//
	// Note that this property is automatically updated by GTK, so this function
	// should only be used by applications which do not use GTK to create
	// toplevel surfaces.
	SetThemeVariantX11Surface(variant string)
	// SetUrgencyHintX11Surface sets a hint on @surface that it needs user
	// attention. See the ICCCM for details.
	SetUrgencyHintX11Surface(urgent bool)
	// SetUserTimeX11Surface: the application can use this call to update the
	// _NET_WM_USER_TIME property on a toplevel surface. This property stores an
	// Xserver time which represents the time of the last user input event
	// received for this surface. This property may be used by the window
	// manager to alter the focus, stacking, and/or placement behavior of
	// surfaces when they are mapped depending on whether the new surface was
	// created by a user action or is a "pop-up" surface activated by a timer or
	// some other event.
	//
	// Note that this property is automatically updated by GDK, so this function
	// should only be used by applications which handle input events bypassing
	// GDK.
	SetUserTimeX11Surface(timestamp uint32)
	// SetUTF8PropertyX11Surface: this function modifies or removes an arbitrary
	// X11 window property of type UTF8_STRING. If the given @surface is not a
	// toplevel surface, it is ignored.
	SetUTF8PropertyX11Surface(name string, value string)
}

// x11Surface implements the X11Surface class.
type x11Surface struct {
	gdk.Surface
}

// WrapX11Surface wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Surface(obj *externglib.Object) X11Surface {
	return x11Surface{
		Surface: gdk.WrapSurface(obj),
	}
}

func marshalX11Surface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Surface(obj), nil
}

func (s x11Surface) Desktop() uint32 {
	var _arg0 *C.GdkSurface // out
	var _cret C.guint32     // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_surface_get_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

func (s x11Surface) Group() gdk.Surface {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_surface_get_group(_arg0)

	var _ret gdk.Surface // out

	_ret = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Surface)

	return _ret
}

func (s x11Surface) MoveToCurrentDesktopX11Surface() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_x11_surface_move_to_current_desktop(_arg0)
}

func (s x11Surface) MoveToDesktopX11Surface(desktop uint32) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (C.guint32)(desktop)

	C.gdk_x11_surface_move_to_desktop(_arg0, _arg1)
}

func (s x11Surface) SetFrameSyncEnabledX11Surface(frameSyncEnabled bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if frameSyncEnabled {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_frame_sync_enabled(_arg0, _arg1)
}

func (s x11Surface) SetGroupX11Surface(leader gdk.Surface) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(leader.Native()))

	C.gdk_x11_surface_set_group(_arg0, _arg1)
}

func (s x11Surface) SetSkipPagerHintX11Surface(skipsPager bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if skipsPager {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_skip_pager_hint(_arg0, _arg1)
}

func (s x11Surface) SetSkipTaskbarHintX11Surface(skipsTaskbar bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if skipsTaskbar {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_skip_taskbar_hint(_arg0, _arg1)
}

func (s x11Surface) SetThemeVariantX11Surface(variant string) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(variant))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_surface_set_theme_variant(_arg0, _arg1)
}

func (s x11Surface) SetUrgencyHintX11Surface(urgent bool) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if urgent {
		_arg1 = C.TRUE
	}

	C.gdk_x11_surface_set_urgency_hint(_arg0, _arg1)
}

func (s x11Surface) SetUserTimeX11Surface(timestamp uint32) {
	var _arg0 *C.GdkSurface // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (C.guint32)(timestamp)

	C.gdk_x11_surface_set_user_time(_arg0, _arg1)
}

func (s x11Surface) SetUTF8PropertyX11Surface(name string, value string) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.char       // out
	var _arg2 *C.char       // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.gdk_x11_surface_set_utf8_property(_arg0, _arg1, _arg2)
}
