// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// CairoCreate creates a Cairo context for drawing to window.
//
// Note that calling cairo_reset_clip() on the resulting #cairo_t will produce
// undefined results, so avoid it at all costs.
//
// Typically, this function is used to draw on a Window out of the paint cycle
// of the toolkit; this should be avoided, as it breaks various assumptions and
// optimizations.
//
// If you are drawing on a native Window in response to a GDK_EXPOSE event you
// should use gdk_window_begin_draw_frame() and
// gdk_drawing_context_get_cairo_context() instead. GTK will automatically do
// this for you when drawing a widget.
//
// Deprecated: Use gdk_window_begin_draw_frame() and
// gdk_drawing_context_get_cairo_context() instead.
//
// The function takes the following parameters:
//
//    - window: Window.
//
// The function returns the following values:
//
//    - context: newly created Cairo context. Free with cairo_destroy() when you
//      are done drawing.
//
func CairoCreate(window Windower) *cairo.Context {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_gret := girepository.MustFind("Gdk", "cairo_create").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(window)

	var _context *cairo.Context // out

	_context = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_context, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	return _context
}

// CairoDrawFromGL: this is the main way to draw GL content in GTK+. It takes a
// render buffer ID (source_type == RENDERBUFFER) or a texture id (source_type
// == TEXTURE) and draws it onto cr with an OVER operation, respecting the
// current clip. The top left corner of the rectangle specified by x, y, width
// and height will be drawn at the current (0,0) position of the cairo_t.
//
// This will work for *all* cairo_t, as long as window is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a window with no
// special effects applied to cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - window we're rendering for (not necessarily into).
//    - source: GL ID of the source buffer.
//    - sourceType: type of the source.
//    - bufferScale: scale-factor that the source buffer is allocated for.
//    - x: source x position in source to start copying from in GL coordinates.
//    - y: source y position in source to start copying from in GL coordinates.
//    - width of the region to draw.
//    - height of the region to draw.
//
func CairoDrawFromGL(cr *cairo.Context, window Windower, source, sourceType, bufferScale, x, y, width, height int32) {
	var _args [9]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(source)
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(sourceType)
	*(*C.int)(unsafe.Pointer(&_args[4])) = C.int(bufferScale)
	*(*C.int)(unsafe.Pointer(&_args[5])) = C.int(x)
	*(*C.int)(unsafe.Pointer(&_args[6])) = C.int(y)
	*(*C.int)(unsafe.Pointer(&_args[7])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[8])) = C.int(height)

	girepository.MustFind("Gdk", "cairo_draw_from_gl").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(window)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceType)
	runtime.KeepAlive(bufferScale)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// CairoGetClipRectangle: this is a convenience function around
// cairo_clip_extents(). It rounds the clip extents to integer coordinates and
// returns a boolean indicating if a clip area exists.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - rect (optional): return location for the clip, or NULL.
//    - ok: TRUE if a clip rectangle exists, FALSE if all of cr is clipped and
//      all drawing can be skipped.
//
func CairoGetClipRectangle(cr *cairo.Context) (*Rectangle, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))

	_gret := girepository.MustFind("Gdk", "cairo_get_clip_rectangle").Invoke(_args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cr)

	var _rect *Rectangle // out
	var _ok bool         // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_rect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _rect, _ok
}

// CairoGetDrawingContext retrieves the DrawingContext that created the Cairo
// context cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - drawingContext (optional) if any is set.
//
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))

	_gret := girepository.MustFind("Gdk", "cairo_get_drawing_context").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cr)

	var _drawingContext *DrawingContext // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_drawingContext = wrapDrawingContext(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _drawingContext
}

// CairoRectangle adds the given rectangle to the current path of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - rectangle: Rectangle.
//
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rectangle)))

	girepository.MustFind("Gdk", "cairo_rectangle").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(rectangle)
}

// CairoRegion adds the given region to the current path of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - region: #cairo_region_t.
//
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(region.Native()))

	girepository.MustFind("Gdk", "cairo_region").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(region)
}

// CairoRegionCreateFromSurface creates region that describes covers the area
// where the given surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
//
// The function takes the following parameters:
//
//    - surface: cairo surface.
//
// The function returns the following values:
//
//    - region must be freed with cairo_region_destroy().
//
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(surface.Native()))

	_gret := girepository.MustFind("Gdk", "cairo_region_create_from_surface").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(surface)

	var _region *cairo.Region // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_region = (*cairo.Region)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	return _region
}

// CairoSetSourceColor sets the specified Color as the source color of cr.
//
// Deprecated: Use gdk_cairo_set_source_rgba() instead.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - color: Color.
//
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))

	girepository.MustFind("Gdk", "cairo_set_source_color").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(color)
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of pixbuf is pixbuf_x, pixbuf_y.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - pixbuf: Pixbuf.
//    - pixbufX: x coordinate of location to place upper left corner of pixbuf.
//    - pixbufY: y coordinate of location to place upper left corner of pixbuf.
//
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX, pixbufY float64) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(pixbufX)
	*(*C.gdouble)(unsafe.Pointer(&_args[3])) = C.gdouble(pixbufY)

	girepository.MustFind("Gdk", "cairo_set_source_pixbuf").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(pixbufX)
	runtime.KeepAlive(pixbufY)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - rgba: RGBA.
//
func CairoSetSourceRGBA(cr *cairo.Context, rgba *RGBA) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	girepository.MustFind("Gdk", "cairo_set_source_rgba").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(rgba)
}

// CairoSetSourceWindow sets the given window as the source pattern for cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of window is x, y. The window contains all its subwindows when
// rendering.
//
// Note that the contents of window are undefined outside of the visible part of
// window, so use this function with care.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - window: Window.
//    - x: x coordinate of location to place upper left corner of window.
//    - y: y coordinate of location to place upper left corner of window.
//
func CairoSetSourceWindow(cr *cairo.Context, window Windower, x, y float64) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(cr.Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(x)
	*(*C.gdouble)(unsafe.Pointer(&_args[3])) = C.gdouble(y)

	girepository.MustFind("Gdk", "cairo_set_source_window").Invoke(_args[:], nil)

	runtime.KeepAlive(cr)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// CairoSurfaceCreateFromPixbuf creates an image surface with the same contents
// as the pixbuf.
//
// The function takes the following parameters:
//
//    - pixbuf: Pixbuf.
//    - scale of the new surface, or 0 to use same as window.
//    - forWindow (optional): window this will be drawn to, or NULL.
//
// The function returns the following values:
//
//    - surface: new cairo surface, must be freed with cairo_surface_destroy().
//
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int32, forWindow Windower) *cairo.Surface {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(scale)
	if forWindow != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(forWindow).Native()))
	}

	_gret := girepository.MustFind("Gdk", "cairo_surface_create_from_pixbuf").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(forWindow)

	var _surface *cairo.Surface // out

	_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	return _surface
}
