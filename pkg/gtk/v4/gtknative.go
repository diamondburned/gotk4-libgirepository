// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_get_type()), F: marshalNative},
	})
}

// Native: `GtkNative` is the interface implemented by all widgets that have
// their own `GdkSurface`.
//
// The obvious example of a `GtkNative` is `GtkWindow`.
//
// Every widget that is not itself a `GtkNative` is contained in one, and you
// can get it with [method@Gtk.Widget.get_native].
//
// To get the surface of a `GtkNative`, use [method@Gtk.Native.get_surface]. It
// is also possible to find the `GtkNative` to which a surface belongs, with
// [func@Gtk.Native.get_for_surface].
//
// In addition to a [class@Gdk.Surface], a `GtkNative` also provides a
// [class@Gsk.Renderer] for rendering on that surface. To get the renderer, use
// [method@Gtk.Native.get_renderer].
type Native interface {
	Widget

	// Renderer returns the renderer that is used for this `GtkNative`.
	Renderer() gsk.Renderer
	// Surface returns the surface of this `GtkNative`.
	Surface() gdk.Surface
	// SurfaceTransform retrieves the surface transform of @self.
	//
	// This is the translation from @self's surface coordinates into @self's
	// widget coordinates.
	SurfaceTransform() (x float64, y float64)
	// Realize realizes a `GtkNative`.
	//
	// This should only be used by subclasses.
	Realize()
	// Unrealize unrealizes a `GtkNative`.
	//
	// This should only be used by subclasses.
	Unrealize()
}

// native implements the Native interface.
type native struct {
	Widget
}

var _ Native = (*native)(nil)

// WrapNative wraps a GObject to a type that implements interface
// Native. It is primarily used internally.
func WrapNative(obj *externglib.Object) Native {
	return Native{
		Widget: WrapWidget(obj),
	}
}

func marshalNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNative(obj), nil
}

// Renderer returns the renderer that is used for this `GtkNative`.
func (s native) Renderer() gsk.Renderer {
	var _arg0 *C.GtkNative   // out
	var _cret *C.GskRenderer // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_native_get_renderer(_arg0)

	var _renderer gsk.Renderer // out

	_renderer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gsk.Renderer)

	return _renderer
}

// Surface returns the surface of this `GtkNative`.
func (s native) Surface() gdk.Surface {
	var _arg0 *C.GtkNative  // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_native_get_surface(_arg0)

	var _surface gdk.Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Surface)

	return _surface
}

// SurfaceTransform retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into @self's
// widget coordinates.
func (s native) SurfaceTransform() (x float64, y float64) {
	var _arg0 *C.GtkNative // out
	var _arg1 C.double     // in
	var _arg2 C.double     // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_get_surface_transform(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out

	_x = (float64)(_arg1)
	_y = (float64)(_arg2)

	return _x, _y
}

// Realize realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (s native) Realize() {
	var _arg0 *C.GtkNative // out

	_arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_realize(_arg0)
}

// Unrealize unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (s native) Unrealize() {
	var _arg0 *C.GtkNative // out

	_arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_unrealize(_arg0)
}
