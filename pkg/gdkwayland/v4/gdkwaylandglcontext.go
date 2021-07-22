// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_gl_context_get_type()), F: marshalWaylandGLContexter},
	})
}

// WaylandGLContext: wayland implementation of GdkGLContext.
type WaylandGLContext struct {
	gdk.GLContext
}

func wrapWaylandGLContext(obj *externglib.Object) *WaylandGLContext {
	return &WaylandGLContext{
		GLContext: gdk.GLContext{
			DrawContext: gdk.DrawContext{
				Object: obj,
			},
		},
	}
}

func marshalWaylandGLContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWaylandGLContext(obj), nil
}

func (*WaylandGLContext) privateWaylandGLContext() {}
