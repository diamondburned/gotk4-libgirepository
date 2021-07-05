// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_draw_context_get_type()), F: marshalDrawContext},
	})
}

// DrawContext: base class for objects implementing different rendering methods.
//
// `GdkDrawContext` is the base object used by contexts implementing different
// rendering methods, such as [class@Gdk.CairoContext] or [class@Gdk.GLContext].
// It provides shared functionality between those contexts.
//
// You will always interact with one of those subclasses.
//
// A `GdkDrawContext` is always associated with a single toplevel surface.
type DrawContext interface {
	gextras.Objector

	// BeginFrameDrawContext indicates that you are beginning the process of
	// redrawing @region on the @context's surface.
	//
	// Calling this function begins a drawing operation using @context on the
	// surface that @context was created from. The actual requirements and
	// guarantees for the drawing operation vary for different implementations
	// of drawing, so a [class@Gdk.CairoContext] and a [class@Gdk.GLContext]
	// need to be treated differently.
	//
	// A call to this function is a requirement for drawing and must be followed
	// by a call to [method@Gdk.DrawContext.end_frame], which will complete the
	// drawing operation and ensure the contents become visible on screen.
	//
	// Note that the @region passed to this function is the minimum region that
	// needs to be drawn and depending on implementation, windowing system and
	// hardware in use, it might be necessary to draw a larger region. Drawing
	// implementation must use [method@Gdk.DrawContext.get_frame_region() to
	// query the region that must be drawn.
	//
	// When using GTK, the widget system automatically places calls to
	// gdk_draw_context_begin_frame() and gdk_draw_context_end_frame() via the
	// use of [class@Gsk.Renderer]s, so application code does not need to call
	// these functions explicitly.
	BeginFrameDrawContext(region *cairo.Region)
	// EndFrameDrawContext ends a drawing operation started with
	// gdk_draw_context_begin_frame().
	//
	// This makes the drawing available on screen. See
	// [method@Gdk.DrawContext.begin_frame] for more details about drawing.
	//
	// When using a [class@Gdk.GLContext], this function may call `glFlush()`
	// implicitly before returning; it is not recommended to call `glFlush()`
	// explicitly before calling this function.
	EndFrameDrawContext()
	// Display retrieves the `GdkDisplay` the @context is created for
	Display() Display
	// FrameRegion retrieves the region that is currently being repainted.
	//
	// After a call to [method@Gdk.DrawContext.begin_frame] this function will
	// return a union of the region passed to that function and the area of the
	// surface that the @context determined needs to be repainted.
	//
	// If @context is not in between calls to
	// [method@Gdk.DrawContext.begin_frame] and
	// [method@Gdk.DrawContext.end_frame], nil will be returned.
	FrameRegion() *cairo.Region
	// Surface retrieves the surface that @context is bound to.
	Surface() Surface
	// IsInFrameDrawContext returns true if @context is in the process of
	// drawing to its surface.
	//
	// This is the case between calls to [method@Gdk.DrawContext.begin_frame]
	// and [method@Gdk.DrawContext.end_frame]. In this situation, drawing
	// commands may be effecting the contents of the @context's surface.
	IsInFrameDrawContext() bool
}

// drawContext implements the DrawContext class.
type drawContext struct {
	gextras.Objector
}

// WrapDrawContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawContext(obj *externglib.Object) DrawContext {
	return drawContext{
		Objector: obj,
	}
}

func marshalDrawContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawContext(obj), nil
}

func (c drawContext) BeginFrameDrawContext(region *cairo.Region) {
	var _arg0 *C.GdkDrawContext // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_draw_context_begin_frame(_arg0, _arg1)
}

func (c drawContext) EndFrameDrawContext() {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	C.gdk_draw_context_end_frame(_arg0)
}

func (c drawContext) Display() Display {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.GdkDisplay     // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_draw_context_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (c drawContext) FrameRegion() *cairo.Region {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.cairo_region_t // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_draw_context_get_frame_region(_arg0)

	var _region *cairo.Region // out

	_region = (*cairo.Region)(unsafe.Pointer(_cret))

	return _region
}

func (c drawContext) Surface() Surface {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.GdkSurface     // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_draw_context_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (c drawContext) IsInFrameDrawContext() bool {
	var _arg0 *C.GdkDrawContext // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_draw_context_is_in_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
