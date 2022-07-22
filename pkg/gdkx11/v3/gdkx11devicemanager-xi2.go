// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeX11DeviceManagerXI2 = coreglib.Type(C.gdk_x11_device_manager_xi2_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeX11DeviceManagerXI2, F: marshalX11DeviceManagerXI2},
	})
}

// X11DeviceManagerXI2Overrider contains methods that are overridable.
type X11DeviceManagerXI2Overrider interface {
}

type X11DeviceManagerXI2 struct {
	_ [0]func() // equal guard
	X11DeviceManagerCore
}

var (
	_ gdk.DeviceManagerer = (*X11DeviceManagerXI2)(nil)
)

func initClassX11DeviceManagerXI2(gclass unsafe.Pointer, goval any) {
}

func wrapX11DeviceManagerXI2(obj *coreglib.Object) *X11DeviceManagerXI2 {
	return &X11DeviceManagerXI2{
		X11DeviceManagerCore: X11DeviceManagerCore{
			DeviceManager: gdk.DeviceManager{
				Object: obj,
			},
		},
	}
}

func marshalX11DeviceManagerXI2(p uintptr) (interface{}, error) {
	return wrapX11DeviceManagerXI2(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
