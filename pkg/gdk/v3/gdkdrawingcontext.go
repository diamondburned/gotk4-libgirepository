// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drawing_context_get_type()), F: marshalDrawingContext},
	})
}

// DrawingContext is an object that represents the current drawing state of a
// Window.
//
// It's possible to use a DrawingContext to draw on a Window via rendering API
// like Cairo or OpenGL.
//
// A DrawingContext can only be created by calling gdk_window_begin_draw_frame()
// and will be valid until a call to gdk_window_end_draw_frame().
//
// DrawingContext is available since GDK 3.22
type DrawingContext interface {
	gextras.Objector

	// CairoContext retrieves a Cairo context to be used to draw on the Window
	// that created the DrawingContext.
	//
	// The returned context is guaranteed to be valid as long as the
	// DrawingContext is valid, that is between a call to
	// gdk_window_begin_draw_frame() and gdk_window_end_draw_frame().
	CairoContext() *cairo.Context
	// Clip retrieves a copy of the clip region used when creating the @context.
	Clip() *cairo.Region
	// Window retrieves the window that created the drawing @context.
	Window() Window
	// IsValid checks whether the given DrawingContext is valid.
	IsValid() bool
}

// drawingContext implements the DrawingContext interface.
type drawingContext struct {
	*externglib.Object
}

var _ DrawingContext = (*drawingContext)(nil)

// WrapDrawingContext wraps a GObject to a type that implements
// interface DrawingContext. It is primarily used internally.
func WrapDrawingContext(obj *externglib.Object) DrawingContext {
	return drawingContext{obj}
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawingContext(obj), nil
}

func (c drawingContext) CairoContext() *cairo.Context {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_t           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drawing_context_get_cairo_context(_arg0)

	var _ret *cairo.Context // out

	_ret = (*cairo.Context)(unsafe.Pointer(_cret))

	return _ret
}

func (c drawingContext) Clip() *cairo.Region {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.cairo_region_t    // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drawing_context_get_clip(_arg0)

	var _region *cairo.Region // out

	_region = (*cairo.Region)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v))
	})

	return _region
}

func (c drawingContext) Window() Window {
	var _arg0 *C.GdkDrawingContext // out
	var _cret *C.GdkWindow         // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drawing_context_get_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (c drawingContext) IsValid() bool {
	var _arg0 *C.GdkDrawingContext // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_drawing_context_is_valid(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
