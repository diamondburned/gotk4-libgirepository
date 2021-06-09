// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerXI2, deviceId int) X11DeviceXI2 {
	var _arg1 *C.GdkX11DeviceManagerXI2
	var _arg2 C.int

	_arg1 = (*C.GdkX11DeviceManagerXI2)(unsafe.Pointer(deviceManager.Native()))
	_arg2 = C.int(deviceId)

	var _cret *C.GdkDevice

	cret = C.gdk_x11_device_manager_lookup(_arg1, _arg2)

	var _x11DeviceXI2 X11DeviceXI2

	_x11DeviceXI2 = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(X11DeviceXI2)

	return _x11DeviceXI2
}
