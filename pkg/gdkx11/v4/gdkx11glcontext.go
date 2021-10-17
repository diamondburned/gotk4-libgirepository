// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_gl_context_get_type()), F: marshalX11GLContexter},
	})
}

type X11GLContext struct {
	gdk.GLContext
}

func wrapX11GLContext(obj *externglib.Object) *X11GLContext {
	return &X11GLContext{
		GLContext: gdk.GLContext{
			DrawContext: gdk.DrawContext{
				Object: obj,
			},
		},
	}
}

func marshalX11GLContexter(p uintptr) (interface{}, error) {
	return wrapX11GLContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*X11GLContext) privateX11GLContext() {}
