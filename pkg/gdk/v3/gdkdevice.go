// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk3_Device_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gdk3_Device_ConnectToolChanged(gpointer, void*, guintptr);
// struct TimeCoord {
//     guint32 time;
//     void    axes;
// };
import "C"

// glib.Type values for gdkdevice.go.
var (
	GTypeDeviceType  = coreglib.Type(C.gdk_device_type_get_type())
	GTypeInputMode   = coreglib.Type(C.gdk_input_mode_get_type())
	GTypeInputSource = coreglib.Type(C.gdk_input_source_get_type())
	GTypeDevice      = coreglib.Type(C.gdk_device_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDeviceType, F: marshalDeviceType},
		{T: GTypeInputMode, F: marshalInputMode},
		{T: GTypeInputSource, F: marshalInputSource},
		{T: GTypeDevice, F: marshalDevice},
	})
}

const MAX_TIMECOORD_AXES = 128

// DeviceType indicates the device type. See
// [above][GdkDeviceManager.description] for more information about the meaning
// of these device types.
type DeviceType C.gint

const (
	// DeviceTypeMaster: device is a master (or virtual) device. There will be
	// an associated focus indicator on the screen.
	DeviceTypeMaster DeviceType = iota
	// DeviceTypeSlave: device is a slave (or physical) device.
	DeviceTypeSlave
	// DeviceTypeFloating: device is a physical device, currently not attached
	// to any virtual device.
	DeviceTypeFloating
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	return DeviceType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DeviceType.
func (d DeviceType) String() string {
	switch d {
	case DeviceTypeMaster:
		return "Master"
	case DeviceTypeSlave:
		return "Slave"
	case DeviceTypeFloating:
		return "Floating"
	default:
		return fmt.Sprintf("DeviceType(%d)", d)
	}
}

// InputMode: enumeration that describes the mode of an input device.
type InputMode C.gint

const (
	// ModeDisabled: device is disabled and will not report any events.
	ModeDisabled InputMode = iota
	// ModeScreen: device is enabled. The device’s coordinate space maps to the
	// entire screen.
	ModeScreen
	// ModeWindow: device is enabled. The device’s coordinate space is mapped to
	// a single window. The manner in which this window is chosen is undefined,
	// but it will typically be the same way in which the focus window for key
	// events is determined.
	ModeWindow
)

func marshalInputMode(p uintptr) (interface{}, error) {
	return InputMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputMode.
func (i InputMode) String() string {
	switch i {
	case ModeDisabled:
		return "Disabled"
	case ModeScreen:
		return "Screen"
	case ModeWindow:
		return "Window"
	default:
		return fmt.Sprintf("InputMode(%d)", i)
	}
}

// InputSource: enumeration describing the type of an input device in general
// terms.
type InputSource C.gint

const (
	// SourceMouse: device is a mouse. (This will be reported for the core
	// pointer, even if it is something else, such as a trackball.).
	SourceMouse InputSource = iota
	// SourcePen: device is a stylus of a graphics tablet or similar device.
	SourcePen
	// SourceEraser: device is an eraser. Typically, this would be the other end
	// of a stylus on a graphics tablet.
	SourceEraser
	// SourceCursor: device is a graphics tablet “puck” or similar device.
	SourceCursor
	// SourceKeyboard: device is a keyboard.
	SourceKeyboard
	// SourceTouchscreen: device is a direct-input touch device, such as a
	// touchscreen or tablet. This device type has been added in 3.4.
	SourceTouchscreen
	// SourceTouchpad: device is an indirect touch device, such as a touchpad.
	// This device type has been added in 3.4.
	SourceTouchpad
	// SourceTrackpoint: device is a trackpoint. This device type has been added
	// in 3.22.
	SourceTrackpoint
	// SourceTabletPad: device is a "pad", a collection of buttons, rings and
	// strips found in drawing tablets. This device type has been added in 3.22.
	SourceTabletPad
)

func marshalInputSource(p uintptr) (interface{}, error) {
	return InputSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InputSource.
func (i InputSource) String() string {
	switch i {
	case SourceMouse:
		return "Mouse"
	case SourcePen:
		return "Pen"
	case SourceEraser:
		return "Eraser"
	case SourceCursor:
		return "Cursor"
	case SourceKeyboard:
		return "Keyboard"
	case SourceTouchscreen:
		return "Touchscreen"
	case SourceTouchpad:
		return "Touchpad"
	case SourceTrackpoint:
		return "Trackpoint"
	case SourceTabletPad:
		return "TabletPad"
	default:
		return fmt.Sprintf("InputSource(%d)", i)
	}
}

// Device object represents a single input device, such as a keyboard, a mouse,
// a touchpad, etc.
//
// See the DeviceManager documentation for more information about the various
// kinds of master and slave devices, and their relationships.
type Device struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Device)(nil)
)

// Devicer describes types inherited from class Device.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Devicer interface {
	coreglib.Objector
	baseDevice() *Device
}

var _ Devicer = (*Device)(nil)

func wrapDevice(obj *coreglib.Object) *Device {
	return &Device{
		Object: obj,
	}
}

func marshalDevice(p uintptr) (interface{}, error) {
	return wrapDevice(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (device *Device) baseDevice() *Device {
	return device
}

// BaseDevice returns the underlying base object.
func BaseDevice(obj Devicer) *Device {
	return obj.baseDevice()
}

//export _gotk4_gdk3_Device_ConnectChanged
func _gotk4_gdk3_Device_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectChanged signal is emitted either when the Device has changed the
// number of either axes or keys. For example In X this will normally happen
// when the slave device routing events through the master device changes (for
// example, user switches from the USB mouse to a tablet), in that case the
// master device will change to reflect the new slave device axes and keys.
func (device *Device) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(device, "changed", false, unsafe.Pointer(C._gotk4_gdk3_Device_ConnectChanged), f)
}

//export _gotk4_gdk3_Device_ConnectToolChanged
func _gotk4_gdk3_Device_ConnectToolChanged(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(tool *DeviceTool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(tool *DeviceTool))
	}

	var _tool *DeviceTool // out

	_tool = wrapDeviceTool(coreglib.Take(unsafe.Pointer(arg1)))

	f(_tool)
}

// ConnectToolChanged signal is emitted on pen/eraser Devices whenever tools
// enter or leave proximity.
func (device *Device) ConnectToolChanged(f func(tool *DeviceTool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(device, "tool-changed", false, unsafe.Pointer(C._gotk4_gdk3_Device_ConnectToolChanged), f)
}

// AssociatedDevice returns the associated device to device, if device is of
// type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or keyboard.
//
// If device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master device
// to which device is attached to.
//
// If device is of type GDK_DEVICE_TYPE_FLOATING, NULL will be returned, as
// there is no associated device.
//
// The function returns the following values:
//
//    - ret (optional): associated device, or NULL.
//
func (device *Device) AssociatedDevice() Devicer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_associated_device", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ret Devicer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Devicer)
				return ok
			})
			rv, ok := casted.(Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_ret = rv
		}
	}

	return _ret
}

// Display returns the Display to which device pertains.
//
// The function returns the following values:
//
//    - display This memory is owned by GTK+, and must not be freed or unreffed.
//
func (device *Device) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// HasCursor determines whether the pointer follows device motion. This is not
// meaningful for keyboard devices, which don't have a pointer.
//
// The function returns the following values:
//
//    - ok: TRUE if the pointer follows device motion.
//
func (device *Device) HasCursor() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_has_cursor", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Key: if index_ has a valid keyval, this function will return TRUE and fill in
// keyval and modifiers with the keyval settings.
//
// The function takes the following parameters:
//
//    - index_: index of the macro button to get.
//
// The function returns the following values:
//
//    - keyval: return value for the keyval.
//    - modifiers: return value for modifiers.
//    - ok: TRUE if keyval is set for index.
//
func (device *Device) Key(index_ uint32) (uint32, ModifierType, bool) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(index_)

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_key", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)

	var _keyval uint32          // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = *(*uint32)(unsafe.Pointer(_outs[0]))
	_modifiers = *(*ModifierType)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

// LastEventWindow gets information about which window the given pointer device
// is in, based on events that have been received so far from the display
// server. If another application has a pointer grab, or this application has a
// grab with owner_events = FALSE, NULL may be returned even if the pointer is
// physically over one of this application's windows.
//
// The function returns the following values:
//
//    - window (optional): last window the device.
//
func (device *Device) LastEventWindow() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_last_event_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _window Windower // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _window
}

// NAxes returns the number of axes the device currently has.
//
// The function returns the following values:
//
//    - gint: number of axes.
//
func (device *Device) NAxes() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_n_axes", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// NKeys returns the number of keys the device currently has.
//
// The function returns the following values:
//
//    - gint: number of keys.
//
func (device *Device) NKeys() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_n_keys", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Name determines the name of the device.
//
// The function returns the following values:
//
//    - utf8: name.
//
func (device *Device) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Position gets the current location of device. As a slave device coordinates
// are those of its master pointer, This function may not be called on devices
// of type GDK_DEVICE_TYPE_SLAVE, unless there is an ongoing grab on them, see
// gdk_device_grab().
//
// The function returns the following values:
//
//    - screen (optional): location to store the Screen the device is on, or
//      NULL.
//    - x (optional): location to store root window X coordinate of device, or
//      NULL.
//    - y (optional): location to store root window Y coordinate of device, or
//      NULL.
//
func (device *Device) Position() (screen *Screen, x, y int32) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	girepository.MustFind("Gdk", "Device").InvokeMethod("get_position", _args[:], _outs[:])

	runtime.KeepAlive(device)

	var _screen *Screen // out
	var _x int32        // out
	var _y int32        // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_screen = wrapScreen(coreglib.Take(unsafe.Pointer(_outs[0])))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_x = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_y = *(*int32)(unsafe.Pointer(_outs[2]))
	}

	return _screen, _x, _y
}

// PositionDouble gets the current location of device in double precision. As a
// slave device's coordinates are those of its master pointer, this function may
// not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is an
// ongoing grab on them. See gdk_device_grab().
//
// The function returns the following values:
//
//    - screen (optional): location to store the Screen the device is on, or
//      NULL.
//    - x (optional): location to store root window X coordinate of device, or
//      NULL.
//    - y (optional): location to store root window Y coordinate of device, or
//      NULL.
//
func (device *Device) PositionDouble() (screen *Screen, x, y float64) {
	var _args [1]girepository.Argument
	var _outs [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	girepository.MustFind("Gdk", "Device").InvokeMethod("get_position_double", _args[:], _outs[:])

	runtime.KeepAlive(device)

	var _screen *Screen // out
	var _x float64      // out
	var _y float64      // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_screen = wrapScreen(coreglib.Take(unsafe.Pointer(_outs[0])))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_x = *(*float64)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[2])) != nil {
		_y = *(*float64)(unsafe.Pointer(_outs[2]))
	}

	return _screen, _x, _y
}

// ProductID returns the product ID of this device, or NULL if this information
// couldn't be obtained. This ID is retrieved from the device, and is thus
// constant for it. See gdk_device_get_vendor_id() for more information.
//
// The function returns the following values:
//
//    - utf8 (optional): product ID, or NULL.
//
func (device *Device) ProductID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_product_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Seat returns the Seat the device belongs to.
//
// The function returns the following values:
//
//    - seat This memory is owned by GTK+ and must not be freed.
//
func (device *Device) Seat() Seater {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_seat", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	return _seat
}

// VendorID returns the vendor ID of this device, or NULL if this information
// couldn't be obtained. This ID is retrieved from the device, and is thus
// constant for it.
//
// This function, together with gdk_device_get_product_id(), can be used to eg.
// compose #GSettings paths to store settings for this device.
//
//     static GSettings *
//     get_device_settings (GdkDevice *device)
//     {
//       const gchar *vendor, *product;
//       GSettings *settings;
//       GdkDevice *device;
//       gchar *path;
//
//       vendor = gdk_device_get_vendor_id (device);
//       product = gdk_device_get_product_id (device);
//
//       path = g_strdup_printf ("/org/example/app/devices/s:s/", vendor, product);
//       settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
//       g_free (path);
//
//       return settings;
//     }.
//
// The function returns the following values:
//
//    - utf8 (optional): vendor ID, or NULL.
//
func (device *Device) VendorID() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_vendor_id", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// WindowAtPosition obtains the window underneath device, returning the location
// of the device in win_x and win_y. Returns NULL if the window tree under
// device is not known to GDK (for example, belongs to another application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
//
// The function returns the following values:
//
//    - winX (optional): return location for the X coordinate of the device
//      location, relative to the window origin, or NULL.
//    - winY (optional): return location for the Y coordinate of the device
//      location, relative to the window origin, or NULL.
//    - window (optional) under the device position, or NULL.
//
func (device *Device) WindowAtPosition() (winX, winY int32, window Windower) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_window_at_position", _args[:], _outs[:])
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _winX int32      // out
	var _winY int32      // out
	var _window Windower // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_winX = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_winY = *(*int32)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _winX, _winY, _window
}

// WindowAtPositionDouble obtains the window underneath device, returning the
// location of the device in win_x and win_y in double precision. Returns NULL
// if the window tree under device is not known to GDK (for example, belongs to
// another application).
//
// As a slave device coordinates are those of its master pointer, This function
// may not be called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is
// an ongoing grab on them, see gdk_device_grab().
//
// The function returns the following values:
//
//    - winX (optional): return location for the X coordinate of the device
//      location, relative to the window origin, or NULL.
//    - winY (optional): return location for the Y coordinate of the device
//      location, relative to the window origin, or NULL.
//    - window (optional) under the device position, or NULL.
//
func (device *Device) WindowAtPositionDouble() (winX, winY float64, window Windower) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("get_window_at_position_double", _args[:], _outs[:])
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _winX float64    // out
	var _winY float64    // out
	var _window Windower // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_winX = *(*float64)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_winY = *(*float64)(unsafe.Pointer(_outs[1]))
	}
	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _winX, _winY, _window
}

// ListSlaveDevices: if the device if of type GDK_DEVICE_TYPE_MASTER, it will
// return the list of slave devices attached to it, otherwise it will return
// NULL.
//
// The function returns the following values:
//
//    - list (optional): the list of slave devices, or NULL. The list must be
//      freed with g_list_free(), the contents of the list are owned by GTK+ and
//      should not be freed.
//
func (device *Device) ListSlaveDevices() []Devicer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "Device").InvokeMethod("list_slave_devices", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _list []Devicer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_list = make([]Devicer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
		gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
			src := (*C.void)(v)
			var dst Devicer // out
			{
				objptr := unsafe.Pointer(src)
				if objptr == nil {
					panic("object of type gdk.Devicer is nil")
				}

				object := coreglib.Take(objptr)
				casted := object.WalkCast(func(obj coreglib.Objector) bool {
					_, ok := obj.(Devicer)
					return ok
				})
				rv, ok := casted.(Devicer)
				if !ok {
					panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
				}
				dst = rv
			}
			_list = append(_list, dst)
		})
	}

	return _list
}

// Ungrab: release any grab on device.
//
// Deprecated: Use gdk_seat_ungrab() instead.
//
// The function takes the following parameters:
//
//    - time_: timestap (e.g. GDK_CURRENT_TIME).
//
func (device *Device) Ungrab(time_ uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(time_)

	girepository.MustFind("Gdk", "Device").InvokeMethod("ungrab", _args[:], nil)

	runtime.KeepAlive(device)
	runtime.KeepAlive(time_)
}

// Warp warps device in display to the point x,y on the screen screen, unless
// the device is confined to a window by a grab, in which case it will be moved
// as far as allowed by the grab. Warping the pointer creates events as if the
// user had moved the mouse instantaneously to the destination.
//
// Note that the pointer should normally be under the control of the user. This
// function was added to cover some rare use cases like keyboard navigation
// support for the color picker in the ColorSelectionDialog.
//
// The function takes the following parameters:
//
//    - screen to warp device to.
//    - x: x coordinate of the destination.
//    - y: y coordinate of the destination.
//
func (device *Device) Warp(screen *Screen, x, y int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(x)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(y)

	girepository.MustFind("Gdk", "Device").InvokeMethod("warp", _args[:], nil)

	runtime.KeepAlive(device)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DeviceGrabInfoLibgtkOnly determines information about the current keyboard
// grab. This is not public API and must not be used by applications.
//
// Deprecated: The symbol was never meant to be used outside of GTK+.
//
// The function takes the following parameters:
//
//    - display for which to get the grab information.
//    - device to get the grab information from.
//
// The function returns the following values:
//
//    - grabWindow: location to store current grab window.
//    - ownerEvents: location to store boolean indicating whether the
//      owner_events flag to gdk_keyboard_grab() or gdk_pointer_grab() was TRUE.
//    - ok: TRUE if this application currently has the keyboard grabbed.
//
func DeviceGrabInfoLibgtkOnly(display *Display, device Devicer) (grabWindow Windower, ownerEvents, ok bool) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_gret := girepository.MustFind("Gdk", "grab_info_libgtk_only").Invoke(_args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)
	runtime.KeepAlive(device)

	var _grabWindow Windower // out
	var _ownerEvents bool    // out
	var _ok bool             // out

	{
		objptr := unsafe.Pointer(_outs[0])
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
		_grabWindow = rv
	}
	if **(**C.void)(unsafe.Pointer(&_outs[1])) != 0 {
		_ownerEvents = true
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _grabWindow, _ownerEvents, _ok
}

// TimeCoord stores a single event in a motion history.
//
// An instance of this type is always passed by reference.
type TimeCoord struct {
	*timeCoord
}

// timeCoord is the struct that's finalized.
type timeCoord struct {
	native unsafe.Pointer
}

// Time: timestamp for this event.
func (t *TimeCoord) Time() uint32 {
	offset := girepository.MustFind("Gdk", "TimeCoord").StructFieldOffset("time")
	valptr := unsafe.Add(unsafe.Pointer(t), offset)
	var v uint32 // out
	v = uint32(*(*C.guint32)(unsafe.Pointer(&valptr)))
	return v
}

// Axes values of the device’s axes.
func (t *TimeCoord) Axes() [128]float64 {
	offset := girepository.MustFind("Gdk", "TimeCoord").StructFieldOffset("axes")
	valptr := unsafe.Add(unsafe.Pointer(t), offset)
	var v [128]float64 // out
	v = *(*[128]float64)(unsafe.Pointer(&valptr))
	return v
}
