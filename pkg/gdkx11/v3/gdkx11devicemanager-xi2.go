// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_manager_xi2_get_type()), F: marshalX11DeviceManagerXI2er},
	})
}

type X11DeviceManagerXI2 struct {
	X11DeviceManagerCore
}

func wrapX11DeviceManagerXI2(obj *externglib.Object) *X11DeviceManagerXI2 {
	return &X11DeviceManagerXI2{
		X11DeviceManagerCore: X11DeviceManagerCore{
			DeviceManager: gdk.DeviceManager{
				Object: obj,
			},
		},
	}
}

func marshalX11DeviceManagerXI2er(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DeviceManagerXI2(obj), nil
}

func (*X11DeviceManagerXI2) privateX11DeviceManagerXI2() {}
