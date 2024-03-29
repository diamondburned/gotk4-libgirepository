// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeX11DeviceXI2 = coreglib.Type(girepository.MustFind("GdkX11", "X11DeviceXI2").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeX11DeviceXI2, F: marshalX11DeviceXI2},
	})
}

type X11DeviceXI2 struct {
	_ [0]func() // equal guard
	gdk.Device
}

var (
	_ gdk.Devicer = (*X11DeviceXI2)(nil)
)

func wrapX11DeviceXI2(obj *coreglib.Object) *X11DeviceXI2 {
	return &X11DeviceXI2{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalX11DeviceXI2(p uintptr) (interface{}, error) {
	return wrapX11DeviceXI2(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
