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
		{T: externglib.Type(C.gdk_x11_display_manager_get_type()), F: marshalX11DisplayManagerer},
	})
}

type X11DisplayManager struct {
	gdk.DisplayManager
}

func wrapX11DisplayManager(obj *externglib.Object) *X11DisplayManager {
	return &X11DisplayManager{
		DisplayManager: gdk.DisplayManager{
			Object: obj,
		},
	}
}

func marshalX11DisplayManagerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DisplayManager(obj), nil
}

func (*X11DisplayManager) privateX11DisplayManager() {}
