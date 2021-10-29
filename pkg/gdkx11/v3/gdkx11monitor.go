// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitorrer},
	})
}

type X11Monitor struct {
	gdk.Monitor
}

var (
	_ externglib.Objector = (*X11Monitor)(nil)
)

func wrapX11Monitor(obj *externglib.Object) *X11Monitor {
	return &X11Monitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalX11Monitorrer(p uintptr) (interface{}, error) {
	return wrapX11Monitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
