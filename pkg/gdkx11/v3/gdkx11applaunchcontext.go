// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_app_launch_context_get_type()), F: marshalX11AppLaunchContexter},
	})
}

// X11AppLaunchContexter describes X11AppLaunchContext's methods.
type X11AppLaunchContexter interface {
	privateX11AppLaunchContext()
}

//
type X11AppLaunchContext struct {
	gdk.AppLaunchContext
}

var (
	_ X11AppLaunchContexter = (*X11AppLaunchContext)(nil)
	_ gextras.Nativer       = (*X11AppLaunchContext)(nil)
)

func wrapX11AppLaunchContext(obj *externglib.Object) X11AppLaunchContexter {
	return &X11AppLaunchContext{
		AppLaunchContext: gdk.AppLaunchContext{
			AppLaunchContext: gio.AppLaunchContext{
				Object: obj,
			},
		},
	}
}

func marshalX11AppLaunchContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11AppLaunchContext(obj), nil
}

func (*X11AppLaunchContext) privateX11AppLaunchContext() {}
