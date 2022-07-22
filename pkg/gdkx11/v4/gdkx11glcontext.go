// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeX11GLContext = coreglib.Type(C.gdk_x11_gl_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeX11GLContext, F: marshalX11GLContext},
	})
}

// GLXVersion retrieves the version of the GLX implementation.
//
// The function returns the following values:
//
//    - major: return location for the GLX major version.
//    - minor: return location for the GLX minor version.
//    - ok: TRUE if GLX is available.
//
func (display *X11Display) GLXVersion() (major, minor int, ok bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // in
	var _arg2 C.int         // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(display)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = int(_arg1)
	_minor = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

// X11GLContextOverrider contains methods that are overridable.
type X11GLContextOverrider interface {
}

type X11GLContext struct {
	_ [0]func() // equal guard
	gdk.GLContext
}

var (
	_ gdk.GLContexter = (*X11GLContext)(nil)
)

func initClassX11GLContext(gclass unsafe.Pointer, goval any) {
}

func wrapX11GLContext(obj *coreglib.Object) *X11GLContext {
	return &X11GLContext{
		GLContext: gdk.GLContext{
			DrawContext: gdk.DrawContext{
				Object: obj,
			},
		},
	}
}

func marshalX11GLContext(p uintptr) (interface{}, error) {
	return wrapX11GLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
