// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
	})
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

type X11Display interface {
	gextras.Objector

	// AsDisplay casts the class to the gdk.Display interface.
	AsDisplay() gdk.Display

	// Beep emits a short beep on @display
	//
	// This method is inherited from gdk.Display
	Beep()
	// Close closes the connection to the windowing system for the given
	// display.
	//
	// This cleans up associated resources.
	//
	// This method is inherited from gdk.Display
	Close()
	// DeviceIsGrabbed returns true if there is an ongoing grab on @device for
	// @display.
	//
	// This method is inherited from gdk.Display
	DeviceIsGrabbed(device gdk.Device) bool
	// Flush flushes any requests queued for the windowing system.
	//
	// This happens automatically when the main loop blocks waiting for new
	// events, but if your application is drawing without returning control to
	// the main loop, you may need to call this function explicitly. A common
	// case where this function needs to be called is when an application is
	// executing drawing commands from a thread other than the thread where the
	// main loop is running.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	//
	// This method is inherited from gdk.Display
	Flush()
	// GetAppLaunchContext returns a `GdkAppLaunchContext` suitable for
	// launching applications on the given display.
	//
	// This method is inherited from gdk.Display
	GetAppLaunchContext() gdk.AppLaunchContext
	// GetClipboard gets the clipboard used for copy/paste operations.
	//
	// This method is inherited from gdk.Display
	GetClipboard() gdk.Clipboard
	// GetDefaultSeat returns the default `GdkSeat` for this display.
	//
	// Note that a display may not have a seat. In this case, this function will
	// return nil.
	//
	// This method is inherited from gdk.Display
	GetDefaultSeat() gdk.Seat
	// GetMonitorAtSurface gets the monitor in which the largest area of
	// @surface resides.
	//
	// Returns a monitor close to @surface if it is outside of all monitors.
	//
	// This method is inherited from gdk.Display
	GetMonitorAtSurface(surface gdk.Surface) gdk.Monitor
	// GetName gets the name of the display.
	//
	// This method is inherited from gdk.Display
	GetName() string
	// GetPrimaryClipboard gets the clipboard used for the primary selection.
	//
	// On backends where the primary clipboard is not supported natively, GDK
	// emulates this clipboard locally.
	//
	// This method is inherited from gdk.Display
	GetPrimaryClipboard() gdk.Clipboard
	// GetSetting retrieves a desktop-wide setting such as double-click time for
	// the @display.
	//
	// This method is inherited from gdk.Display
	GetSetting(name string, value externglib.Value) bool
	// GetStartupNotificationID gets the startup notification ID for a Wayland
	// display, or nil if no ID has been defined.
	//
	// This method is inherited from gdk.Display
	GetStartupNotificationID() string
	// IsClosed finds out if the display has been closed.
	//
	// This method is inherited from gdk.Display
	IsClosed() bool
	// IsComposited returns whether surfaces can reasonably be expected to have
	// their alpha channel drawn correctly on the screen.
	//
	// Check [method@Gdk.Display.is_rgba] for whether the display supports an
	// alpha channel.
	//
	// On X11 this function returns whether a compositing manager is compositing
	// on @display.
	//
	// On modern displays, this value is always true.
	//
	// This method is inherited from gdk.Display
	IsComposited() bool
	// IsRGBA returns whether surfaces on this @display are created with an
	// alpha channel.
	//
	// Even if a true is returned, it is possible that the surface’s alpha
	// channel won’t be honored when displaying the surface on the screen: in
	// particular, for X an appropriate windowing manager and compositing
	// manager must be running to provide appropriate display. Use
	// [method@Gdk.Display.is_composited] to check if that is the case.
	//
	// On modern displays, this value is always true.
	//
	// This method is inherited from gdk.Display
	IsRGBA() bool
	// MapKeycode returns the keyvals bound to @keycode.
	//
	// The Nth `GdkKeymapKey` in @keys is bound to the Nth keyval in @keyvals.
	//
	// When a keycode is pressed by the user, the keyval from this list of
	// entries is selected by considering the effective keyboard group and
	// level.
	//
	// Free the returned arrays with g_free().
	//
	// This method is inherited from gdk.Display
	MapKeycode(keycode uint) ([]gdk.KeymapKey, []uint, bool)
	// MapKeyval obtains a list of keycode/group/level combinations that will
	// generate @keyval.
	//
	// Groups and levels are two kinds of keyboard mode; in general, the level
	// determines whether the top or bottom symbol on a key is used, and the
	// group determines whether the left or right symbol is used.
	//
	// On US keyboards, the shift key changes the keyboard level, and there are
	// no groups. A group switch key might convert a keyboard between Hebrew to
	// English modes, for example.
	//
	// `GdkEventKey` contains a group field that indicates the active keyboard
	// group. The level is computed from the modifier mask.
	//
	// The returned array should be freed with g_free().
	//
	// This method is inherited from gdk.Display
	MapKeyval(keyval uint) ([]gdk.KeymapKey, bool)
	// NotifyStartupComplete indicates to the GUI environment that the
	// application has finished loading, using a given identifier.
	//
	// GTK will call this function automatically for [class@Gtk.Window] with
	// custom startup-notification identifier unless
	// [method@Gtk.Window.set_auto_startup_notification] is called to disable
	// that feature.
	//
	// This method is inherited from gdk.Display
	NotifyStartupComplete(startupId string)
	// PutEvent appends the given event onto the front of the event queue for
	// @display.
	//
	// This function is only useful in very special situations and should not be
	// used by applications.
	//
	// This method is inherited from gdk.Display
	PutEvent(event gdk.Event)
	// SupportsInputShapes returns true if the display supports input shapes.
	//
	// This means that [method@Gdk.Surface.set_input_region] can be used to
	// modify the input shape of surfaces on @display.
	//
	// On modern displays, this value is always true.
	//
	// This method is inherited from gdk.Display
	SupportsInputShapes() bool
	// Sync flushes any requests queued for the windowing system and waits until
	// all requests have been handled.
	//
	// This is often used for making sure that the display is synchronized with
	// the current state of the program. Calling [method@Gdk.Display.sync]
	// before [method@GdkX11.Display.error_trap_pop] makes sure that any errors
	// generated from earlier requests are handled before the error trap is
	// removed.
	//
	// This is most useful for X11. On windowing systems where requests are
	// handled synchronously, this function will do nothing.
	//
	// This method is inherited from gdk.Display
	Sync()
	// TranslateKey translates the contents of a `GdkEventKey` into a keyval,
	// effective group, and level.
	//
	// Modifiers that affected the translation and are thus unavailable for
	// application use are returned in @consumed_modifiers.
	//
	// The @effective_group is the group that was actually used for the
	// translation; some keys such as Enter are not affected by the active
	// keyboard group. The @level is derived from @state.
	//
	// @consumed_modifiers gives modifiers that should be masked out from @state
	// when comparing this key press to a keyboard shortcut. For instance, on a
	// US keyboard, the `plus` symbol is shifted, so when comparing a key press
	// to a `<Control>plus` accelerator `<Shift>` should be masked out.
	//
	// This function should rarely be needed, since `GdkEventKey` already
	// contains the translated keyval. It is exported for the benefit of
	// virtualized test environments.
	//
	// This method is inherited from gdk.Display
	TranslateKey(keycode uint, state gdk.ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed gdk.ModifierType, ok bool)

	// ErrorTrapPop pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	ErrorTrapPop() int
	// ErrorTrapPopIgnored pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	ErrorTrapPopIgnored()
	// ErrorTrapPush begins a range of X requests on @display for which X error
	// events will be ignored. Unignored errors (when no trap is pushed) will
	// abort the application. Use gdk_x11_display_error_trap_pop() or
	// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
	// function.
	ErrorTrapPush()
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
	// Grab: call XGrabServer() on @display. To ungrab the display again, use
	// gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	Grab()
	// SetCursorTheme sets the cursor theme from which the images for cursor
	// should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_texture() will have
	// to be handled by the application (GTK applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorTheme(theme string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
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
	SetStartupNotificationID(startupId string)
	// SetSurfaceScale forces a specific window scale for all windows on this
	// display, instead of using the default or user configured scale. This is
	// can be used to disable scaling support by setting @scale to 1, or to
	// programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetSurfaceScale(scale int)
	// StringToCompoundText: convert a string from the encoding of the current
	// locale into a form suitable for storing in a window property.
	StringToCompoundText(str string) (encoding string, format int, ctext []byte, gint int)
	// TextPropertyToTextList: convert a text string from the encoding as it is
	// stored in a property into an array of strings in the encoding of the
	// current locale. (The elements of the array represent the nul-separated
	// elements of the original text string.)
	TextPropertyToTextList(encoding string, format int, text *byte, length int, list **string) int
	// Ungrab @display after it has been grabbed with gdk_x11_display_grab().
	Ungrab()
	// UTF8ToCompoundText converts from UTF-8 to compound text.
	UTF8ToCompoundText(str string) (string, int, []byte, bool)
}

// x11Display implements the X11Display interface.
type x11Display struct {
	*externglib.Object
}

var _ X11Display = (*x11Display)(nil)

// WrapX11Display wraps a GObject to a type that implements
// interface X11Display. It is primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return x11Display{obj}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

func (x x11Display) AsDisplay() gdk.Display {
	return gdk.WrapDisplay(gextras.InternObject(x))
}

func (d x11Display) Beep() {
	gdk.WrapDisplay(gextras.InternObject(d)).Beep()
}

func (d x11Display) Close() {
	gdk.WrapDisplay(gextras.InternObject(d)).Close()
}

func (d x11Display) DeviceIsGrabbed(device gdk.Device) bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).DeviceIsGrabbed(device)
}

func (d x11Display) Flush() {
	gdk.WrapDisplay(gextras.InternObject(d)).Flush()
}

func (d x11Display) GetAppLaunchContext() gdk.AppLaunchContext {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetAppLaunchContext()
}

func (d x11Display) GetClipboard() gdk.Clipboard {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetClipboard()
}

func (d x11Display) GetDefaultSeat() gdk.Seat {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetDefaultSeat()
}

func (d x11Display) GetMonitorAtSurface(surface gdk.Surface) gdk.Monitor {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetMonitorAtSurface(surface)
}

func (d x11Display) GetName() string {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetName()
}

func (d x11Display) GetPrimaryClipboard() gdk.Clipboard {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetPrimaryClipboard()
}

func (d x11Display) GetSetting(name string, value externglib.Value) bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetSetting(name, value)
}

func (d x11Display) GetStartupNotificationID() string {
	return gdk.WrapDisplay(gextras.InternObject(d)).GetStartupNotificationID()
}

func (d x11Display) IsClosed() bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).IsClosed()
}

func (d x11Display) IsComposited() bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).IsComposited()
}

func (d x11Display) IsRGBA() bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).IsRGBA()
}

func (d x11Display) MapKeycode(keycode uint) ([]gdk.KeymapKey, []uint, bool) {
	return gdk.WrapDisplay(gextras.InternObject(d)).MapKeycode(keycode)
}

func (d x11Display) MapKeyval(keyval uint) ([]gdk.KeymapKey, bool) {
	return gdk.WrapDisplay(gextras.InternObject(d)).MapKeyval(keyval)
}

func (d x11Display) NotifyStartupComplete(startupId string) {
	gdk.WrapDisplay(gextras.InternObject(d)).NotifyStartupComplete(startupId)
}

func (d x11Display) PutEvent(event gdk.Event) {
	gdk.WrapDisplay(gextras.InternObject(d)).PutEvent(event)
}

func (d x11Display) SupportsInputShapes() bool {
	return gdk.WrapDisplay(gextras.InternObject(d)).SupportsInputShapes()
}

func (d x11Display) Sync() {
	gdk.WrapDisplay(gextras.InternObject(d)).Sync()
}

func (d x11Display) TranslateKey(keycode uint, state gdk.ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed gdk.ModifierType, ok bool) {
	return gdk.WrapDisplay(gextras.InternObject(d)).TranslateKey(keycode, state, group)
}

func (d x11Display) ErrorTrapPop() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_x11_display_error_trap_pop(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d x11Display) ErrorTrapPopIgnored() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(_arg0)
}

func (d x11Display) ErrorTrapPush() {
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

	_major = int(_arg1)
	_minor = int(_arg2)
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

	_guint32 = uint32(_cret)

	return _guint32
}

func (d x11Display) Grab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(_arg0)
}

func (d x11Display) SetCursorTheme(theme string, size int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(theme))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(size)

	C.gdk_x11_display_set_cursor_theme(_arg0, _arg1, _arg2)
}

func (d x11Display) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_display_set_startup_notification_id(_arg0, _arg1)
}

func (d x11Display) SetSurfaceScale(scale int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(scale)

	C.gdk_x11_display_set_surface_scale(_arg0, _arg1)
}

func (d x11Display) StringToCompoundText(str string) (encoding string, format int, ctext []byte, gint int) {
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
	_format = int(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_gint = int(_cret)

	return _encoding, _format, _ctext, _gint
}

func (d x11Display) TextPropertyToTextList(encoding string, format int, text *byte, length int, list **string) int {
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
	_arg2 = C.int(format)
	_arg3 = (*C.guchar)(unsafe.Pointer(text))
	_arg4 = C.int(length)
	{
		var refTmpIn string
		var refTmpOut *C.char

		refTmpIn = **list

		refTmpOut = (*C.char)(C.CString(refTmpIn))
		defer C.free(unsafe.Pointer(refTmpOut))

		if refTmpOut != nil {
			out0 := &refTmpOut
			if out0 != nil {
				out1 := &out0
				_arg5 = out1
			}
		}
	}

	_cret = C.gdk_x11_display_text_property_to_text_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d x11Display) Ungrab() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(_arg0)
}

func (d x11Display) UTF8ToCompoundText(str string) (string, int, []byte, bool) {
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
	_format = int(_arg3)
	_ctext = unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5)
	runtime.SetFinalizer(&_ctext, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _encoding, _format, _ctext, _ok
}
