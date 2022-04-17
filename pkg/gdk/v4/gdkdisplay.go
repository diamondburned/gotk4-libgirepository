// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_Display_ConnectClosed(gpointer, gboolean, guintptr);
// extern void _gotk4_gdk4_Display_ConnectOpened(gpointer, guintptr);
// extern void _gotk4_gdk4_Display_ConnectSeatAdded(gpointer, GdkSeat*, guintptr);
// extern void _gotk4_gdk4_Display_ConnectSeatRemoved(gpointer, GdkSeat*, guintptr);
// extern void _gotk4_gdk4_Display_ConnectSettingChanged(gpointer, gchar*, guintptr);
import "C"

// glib.Type values for gdkdisplay.go.
var GTypeDisplay = externglib.Type(C.gdk_display_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDisplay, F: marshalDisplay},
	})
}

// Display: GdkDisplay objects are the GDK representation of a workstation.
//
// Their purpose are two-fold:
//
// - To manage and provide information about input devices (pointers, keyboards,
// etc)
//
// - To manage and provide information about output devices (monitors,
// projectors, etc)
//
// Most of the input device handling has been factored out into separate
// gdk.Seat objects. Every display has a one or more seats, which can be
// accessed with gdk.Display.GetDefaultSeat() and gdk.Display.ListSeats().
//
// Output devices are represented by gdk.Monitor objects, which can be accessed
// with gdk.Display.GetMonitorAtSurface() and similar APIs.
type Display struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Display)(nil)
)

func wrapDisplay(obj *externglib.Object) *Display {
	return &Display{
		Object: obj,
	}
}

func marshalDisplay(p uintptr) (interface{}, error) {
	return wrapDisplay(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk4_Display_ConnectClosed
func _gotk4_gdk4_Display_ConnectClosed(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(isError bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(isError bool))
	}

	var _isError bool // out

	if arg1 != 0 {
		_isError = true
	}

	f(_isError)
}

// ConnectClosed is emitted when the connection to the windowing system for
// display is closed.
func (display *Display) ConnectClosed(f func(isError bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(display, "closed", false, unsafe.Pointer(C._gotk4_gdk4_Display_ConnectClosed), f)
}

//export _gotk4_gdk4_Display_ConnectOpened
func _gotk4_gdk4_Display_ConnectOpened(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectOpened is emitted when the connection to the windowing system for
// display is opened.
func (display *Display) ConnectOpened(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(display, "opened", false, unsafe.Pointer(C._gotk4_gdk4_Display_ConnectOpened), f)
}

//export _gotk4_gdk4_Display_ConnectSeatAdded
func _gotk4_gdk4_Display_ConnectSeatAdded(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

// ConnectSeatAdded is emitted whenever a new seat is made known to the
// windowing system.
func (display *Display) ConnectSeatAdded(f func(seat Seater)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(display, "seat-added", false, unsafe.Pointer(C._gotk4_gdk4_Display_ConnectSeatAdded), f)
}

//export _gotk4_gdk4_Display_ConnectSeatRemoved
func _gotk4_gdk4_Display_ConnectSeatRemoved(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

// ConnectSeatRemoved is emitted whenever a seat is removed by the windowing
// system.
func (display *Display) ConnectSeatRemoved(f func(seat Seater)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(display, "seat-removed", false, unsafe.Pointer(C._gotk4_gdk4_Display_ConnectSeatRemoved), f)
}

//export _gotk4_gdk4_Display_ConnectSettingChanged
func _gotk4_gdk4_Display_ConnectSettingChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(setting string)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(setting string))
	}

	var _setting string // out

	_setting = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_setting)
}

// ConnectSettingChanged is emitted whenever a setting changes its value.
func (display *Display) ConnectSettingChanged(f func(setting string)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(display, "setting-changed", false, unsafe.Pointer(C._gotk4_gdk4_Display_ConnectSettingChanged), f)
}

// Beep emits a short beep on display.
func (display *Display) Beep() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_display_beep(_arg0)
	runtime.KeepAlive(display)
}

// Close closes the connection to the windowing system for the given display.
//
// This cleans up associated resources.
func (display *Display) Close() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_display_close(_arg0)
	runtime.KeepAlive(display)
}

// DeviceIsGrabbed returns TRUE if there is an ongoing grab on device for
// display.
//
// The function takes the following parameters:
//
//    - device: GdkDevice.
//
// The function returns the following values:
//
//    - ok: TRUE if there is a grab in effect for device.
//
func (display *Display) DeviceIsGrabbed(device Devicer) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(externglib.InternObject(device).Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Flush flushes any requests queued for the windowing system.
//
// This happens automatically when the main loop blocks waiting for new events,
// but if your application is drawing without returning control to the main
// loop, you may need to call this function explicitly. A common case where this
// function needs to be called is when an application is executing drawing
// commands from a thread other than the thread where the main loop is running.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *Display) Flush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_display_flush(_arg0)
	runtime.KeepAlive(display)
}

// AppLaunchContext returns a GdkAppLaunchContext suitable for launching
// applications on the given display.
//
// The function returns the following values:
//
//    - appLaunchContext: new GdkAppLaunchContext for display. Free with
//      g_object_unref() when done.
//
func (display *Display) AppLaunchContext() *AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)
	runtime.KeepAlive(display)

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = wrapAppLaunchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// Clipboard gets the clipboard used for copy/paste operations.
//
// The function returns the following values:
//
//    - clipboard display's clipboard.
//
func (display *Display) Clipboard() *Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_clipboard(_arg0)
	runtime.KeepAlive(display)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}

// DefaultSeat returns the default GdkSeat for this display.
//
// Note that a display may not have a seat. In this case, this function will
// return NULL.
//
// The function returns the following values:
//
//    - seat (optional): default seat.
//
func (display *Display) DefaultSeat() Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)
	runtime.KeepAlive(display)

	var _seat Seater // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Seater)
				return ok
			})
			rv, ok := casted.(Seater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
			}
			_seat = rv
		}
	}

	return _seat
}

// MonitorAtSurface gets the monitor in which the largest area of surface
// resides.
//
// Returns a monitor close to surface if it is outside of all monitors.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface.
//
// The function returns the following values:
//
//    - monitor with the largest overlap with surface.
//
func (display *Display) MonitorAtSurface(surface Surfacer) *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkSurface // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(externglib.InternObject(surface).Native()))

	_cret = C.gdk_display_get_monitor_at_surface(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(surface)

	var _monitor *Monitor // out

	_monitor = wrapMonitor(externglib.Take(unsafe.Pointer(_cret)))

	return _monitor
}

// Monitors gets the list of monitors associated with this display.
//
// Subsequent calls to this function will always return the same list for the
// same display.
//
// You can listen to the GListModel::items-changed signal on this list to
// monitor changes to the monitor of this display.
//
// The function returns the following values:
//
//    - listModel of GdkMonitor.
//
func (self *Display) Monitors() gio.ListModelOverrider {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GListModel // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gdk_display_get_monitors(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModelOverrider // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.ListModelOverrider)
			return ok
		})
		rv, ok := casted.(gio.ListModelOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// Name gets the name of the display.
//
// The function returns the following values:
//
//    - utf8: string representing the display name. This string is owned by GDK
//      and should not be modified or freed.
//
func (display *Display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_name(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PrimaryClipboard gets the clipboard used for the primary selection.
//
// On backends where the primary clipboard is not supported natively, GDK
// emulates this clipboard locally.
//
// The function returns the following values:
//
//    - clipboard: primary clipboard.
//
func (display *Display) PrimaryClipboard() *Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_primary_clipboard(_arg0)
	runtime.KeepAlive(display)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}

// Setting retrieves a desktop-wide setting such as double-click time for the
// display.
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
func (display *Display) Setting(name string, value *externglib.Value) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.GValue     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gdk_display_get_setting(_arg0, _arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartupNotificationID gets the startup notification ID for a Wayland display,
// or NULL if no ID has been defined.
//
// The function returns the following values:
//
//    - utf8 (optional): startup notification ID for display, or NULL.
//
func (display *Display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_get_startup_notification_id(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IsClosed finds out if the display has been closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the display is closed.
//
func (display *Display) IsClosed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_is_closed(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsComposited returns whether surfaces can reasonably be expected to have
// their alpha channel drawn correctly on the screen.
//
// Check gdk.Display.IsRGBA() for whether the display supports an alpha channel.
//
// On X11 this function returns whether a compositing manager is compositing on
// display.
//
// On modern displays, this value is always TRUE.
//
// The function returns the following values:
//
//    - ok: whether surfaces with RGBA visuals can reasonably be expected to have
//      their alpha channels drawn correctly on the screen.
//
func (display *Display) IsComposited() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_is_composited(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRGBA returns whether surfaces on this display are created with an alpha
// channel.
//
// Even if a TRUE is returned, it is possible that the surface’s alpha channel
// won’t be honored when displaying the surface on the screen: in particular,
// for X an appropriate windowing manager and compositing manager must be
// running to provide appropriate display. Use gdk.Display.IsComposited() to
// check if that is the case.
//
// On modern displays, this value is always TRUE.
//
// The function returns the following values:
//
//    - ok: TRUE if surfaces are created with an alpha channel or FALSE if the
//      display does not support this functionality.
//
func (display *Display) IsRGBA() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_is_rgba(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListSeats returns the list of seats known to display.
//
// The function returns the following values:
//
//    - list: the list of seats known to the GdkDisplay.
//
func (display *Display) ListSeats() []Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GList      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_list_seats(_arg0)
	runtime.KeepAlive(display)

	var _list []Seater // out

	_list = make([]Seater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkSeat)(v)
		var dst Seater // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gdk.Seater is nil")
			}

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Seater)
				return ok
			})
			rv, ok := casted.(Seater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// MapKeycode returns the keyvals bound to keycode.
//
// The Nth GdkKeymapKey in keys is bound to the Nth keyval in keyvals.
//
// When a keycode is pressed by the user, the keyval from this list of entries
// is selected by considering the effective keyboard group and level.
//
// Free the returned arrays with g_free().
//
// The function takes the following parameters:
//
//    - keycode: keycode.
//
// The function returns the following values:
//
//    - keys (optional): return location for array of GdkKeymapKey, or NULL.
//    - keyvals (optional): return location for array of keyvals, or NULL.
//    - ok: TRUE if there were any entries.
//
func (display *Display) MapKeycode(keycode uint) ([]KeymapKey, []uint, bool) {
	var _arg0 *C.GdkDisplay   // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg4 C.int           // in
	var _arg3 *C.guint        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_display_map_keycode(_arg0, _arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(display)
	runtime.KeepAlive(keycode)

	var _keys []KeymapKey // out
	var _keyvals []uint   // out
	var _ok bool          // out

	if _arg2 != nil {
		defer C.free(unsafe.Pointer(_arg2))
		{
			src := unsafe.Slice((*C.GdkKeymapKey)(_arg2), _arg4)
			_keys = make([]KeymapKey, _arg4)
			for i := 0; i < int(_arg4); i++ {
				_keys[i] = *(*KeymapKey)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
				runtime.SetFinalizer(
					gextras.StructIntern(unsafe.Pointer(&_keys[i])),
					func(intern *struct{ C unsafe.Pointer }) {
						C.free(intern.C)
					},
				)
			}
		}
	}
	if _arg3 != nil {
		defer C.free(unsafe.Pointer(_arg3))
		{
			src := unsafe.Slice((*C.guint)(_arg3), _arg4)
			_keyvals = make([]uint, _arg4)
			for i := 0; i < int(_arg4); i++ {
				_keyvals[i] = uint(src[i])
			}
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _keys, _keyvals, _ok
}

// MapKeyval obtains a list of keycode/group/level combinations that will
// generate keyval.
//
// Groups and levels are two kinds of keyboard mode; in general, the level
// determines whether the top or bottom symbol on a key is used, and the group
// determines whether the left or right symbol is used.
//
// On US keyboards, the shift key changes the keyboard level, and there are no
// groups. A group switch key might convert a keyboard between Hebrew to English
// modes, for example.
//
// GdkEventKey contains a group field that indicates the active keyboard group.
// The level is computed from the modifier mask.
//
// The returned array should be freed with g_free().
//
// The function takes the following parameters:
//
//    - keyval: keyval, such as GDK_KEY_a, GDK_KEY_Up, GDK_KEY_Return, etc.
//
// The function returns the following values:
//
//    - keys: return location for an array of GdkKeymapKey.
//    - ok: TRUE if keys were found and returned.
//
func (display *Display) MapKeyval(keyval uint) ([]KeymapKey, bool) {
	var _arg0 *C.GdkDisplay   // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg3 C.int           // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = C.guint(keyval)

	_cret = C.gdk_display_map_keyval(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(display)
	runtime.KeepAlive(keyval)

	var _keys []KeymapKey // out
	var _ok bool          // out

	defer C.free(unsafe.Pointer(_arg2))
	{
		src := unsafe.Slice((*C.GdkKeymapKey)(_arg2), _arg3)
		_keys = make([]KeymapKey, _arg3)
		for i := 0; i < int(_arg3); i++ {
			_keys[i] = *(*KeymapKey)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_keys[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _keys, _ok
}

// NotifyStartupComplete indicates to the GUI environment that the application
// has finished loading, using a given identifier.
//
// GTK will call this function automatically for gtk.Window with custom
// startup-notification identifier unless
// gtk.Window.SetAutoStartupNotification() is called to disable that feature.
//
// The function takes the following parameters:
//
//    - startupId: startup-notification identifier, for which notification
//      process should be completed.
//
func (display *Display) NotifyStartupComplete(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}

// PutEvent appends the given event onto the front of the event queue for
// display.
//
// This function is only useful in very special situations and should not be
// used by applications.
//
// The function takes the following parameters:
//
//    - event: GdkEvent.
//
func (display *Display) PutEvent(event Eventer) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkEvent   // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(externglib.InternObject(event).Native()))

	C.gdk_display_put_event(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(event)
}

// SupportsInputShapes returns TRUE if the display supports input shapes.
//
// This means that gdk.Surface.SetInputRegion() can be used to modify the input
// shape of surfaces on display.
//
// On modern displays, this value is always TRUE.
//
// The function returns the following values:
//
//    - ok: TRUE if surfaces with modified input shape are supported.
//
func (display *Display) SupportsInputShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sync flushes any requests queued for the windowing system and waits until all
// requests have been handled.
//
// This is often used for making sure that the display is synchronized with the
// current state of the program. Calling gdk.Display.Sync() before
// gdkx11.Display.ErrorTrapPop() makes sure that any errors generated from
// earlier requests are handled before the error trap is removed.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *Display) Sync() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	C.gdk_display_sync(_arg0)
	runtime.KeepAlive(display)
}

// TranslateKey translates the contents of a GdkEventKey into a keyval,
// effective group, and level.
//
// Modifiers that affected the translation and are thus unavailable for
// application use are returned in consumed_modifiers.
//
// The effective_group is the group that was actually used for the translation;
// some keys such as Enter are not affected by the active keyboard group. The
// level is derived from state.
//
// consumed_modifiers gives modifiers that should be masked out from state when
// comparing this key press to a keyboard shortcut. For instance, on a US
// keyboard, the plus symbol is shifted, so when comparing a key press to a
// <Control>plus accelerator <Shift> should be masked out.
//
// This function should rarely be needed, since GdkEventKey already contains the
// translated keyval. It is exported for the benefit of virtualized test
// environments.
//
// The function takes the following parameters:
//
//    - keycode: keycode.
//    - state: modifier state.
//    - group: active keyboard group.
//
// The function returns the following values:
//
//    - keyval (optional): return location for keyval, or NULL.
//    - effectiveGroup (optional): return location for effective group, or NULL.
//    - level (optional): return location for level, or NULL.
//    - consumed (optional): return location for modifiers that were used to
//      determine the group or level, or NULL.
//    - ok: TRUE if there was a keyval bound to keycode/state/group.
//
func (display *Display) TranslateKey(keycode uint, state ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed ModifierType, ok bool) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 C.int             // out
	var _arg4 C.guint           // in
	var _arg5 C.int             // in
	var _arg6 C.int             // in
	var _arg7 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))
	_arg1 = C.guint(keycode)
	_arg2 = C.GdkModifierType(state)
	_arg3 = C.int(group)

	_cret = C.gdk_display_translate_key(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)
	runtime.KeepAlive(display)
	runtime.KeepAlive(keycode)
	runtime.KeepAlive(state)
	runtime.KeepAlive(group)

	var _keyval uint           // out
	var _effectiveGroup int    // out
	var _level int             // out
	var _consumed ModifierType // out
	var _ok bool               // out

	_keyval = uint(_arg4)
	_effectiveGroup = int(_arg5)
	_level = int(_arg6)
	_consumed = ModifierType(_arg7)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _effectiveGroup, _level, _consumed, _ok
}

// DisplayGetDefault gets the default GdkDisplay.
//
// This is a convenience function for: gdk_display_manager_get_default_display
// (gdk_display_manager_get ()).
//
// The function returns the following values:
//
//    - display (optional): GdkDisplay, or NULL if there is no default display.
//
func DisplayGetDefault() *Display {
	var _cret *C.GdkDisplay // in

	_cret = C.gdk_display_get_default()

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// DisplayOpen opens a display.
//
// The function takes the following parameters:
//
//    - displayName: name of the display to open.
//
// The function returns the following values:
//
//    - display (optional): GdkDisplay, or NULL if the display could not be
//      opened.
//
func DisplayOpen(displayName string) *Display {
	var _arg1 *C.char       // out
	var _cret *C.GdkDisplay // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(displayName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_open(_arg1)
	runtime.KeepAlive(displayName)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}
