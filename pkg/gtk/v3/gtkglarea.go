// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gl_area_get_type()), F: marshalGLAreaer},
	})
}

// GLAreaOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type GLAreaOverrider interface {
	Render(context gdk.GLContexter) bool
	Resize(width int, height int)
}

// GLArea is a widget that allows drawing with OpenGL.
//
// GLArea sets up its own GLContext for the window it creates, and creates a
// custom GL framebuffer that the widget will do GL rendering onto. It also
// ensures that this framebuffer is the default GL rendering target when
// rendering.
//
// In order to draw, you have to connect to the GLArea::render signal, or
// subclass GLArea and override the GtkGLAreaClass.render() virtual function.
//
// The GLArea widget ensures that the GLContext is associated with the widget's
// drawing area, and it is kept updated when the size and position of the
// drawing area changes.
//
// Drawing with GtkGLArea ##
//
// The simplest way to draw using OpenGL commands in a GLArea is to create a
// widget instance and connect to the GLArea::render signal:
//
//      static void
//      on_realize (GtkGLarea *area)
//      {
//        // We need to make the context current if we want to
//        // call GL API
//        gtk_gl_area_make_current (area);
//
//        // If there were errors during the initialization or
//        // when trying to make the context current, this
//        // function will return a #GError for you to catch
//        if (gtk_gl_area_get_error (area) != NULL)
//          return;
//
//        // You can also use gtk_gl_area_set_error() in order
//        // to show eventual initialization errors on the
//        // GtkGLArea widget itself
//        GError *internal_error = NULL;
//        init_buffer_objects (&error);
//        if (error != NULL)
//          {
//            gtk_gl_area_set_error (area, error);
//            g_error_free (error);
//            return;
//          }
//
//        init_shaders (&error);
//        if (error != NULL)
//          {
//            gtk_gl_area_set_error (area, error);
//            g_error_free (error);
//            return;
//          }
//      }
//
// If you need to change the options for creating the GLContext you should use
// the GLArea::create-context signal.
type GLArea struct {
	Widget
}

func wrapGLArea(obj *externglib.Object) *GLArea {
	return &GLArea{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalGLAreaer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGLArea(obj), nil
}

// NewGLArea creates a new GLArea widget.
func NewGLArea() *GLArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_gl_area_new()

	var _glArea *GLArea // out

	_glArea = wrapGLArea(externglib.Take(unsafe.Pointer(_cret)))

	return _glArea
}

// AttachBuffers ensures that the area framebuffer object is made the current
// draw and read target, and that all the required buffers for the area are
// created and bound to the frambuffer.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) AttachBuffers() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_attach_buffers(_arg0)
	runtime.KeepAlive(area)
}

// AutoRender returns whether the area is in auto render mode or not.
func (area *GLArea) AutoRender() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_auto_render(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Context retrieves the GLContext used by area.
func (area *GLArea) Context() gdk.GLContexter {
	var _arg0 *C.GtkGLArea    // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_context(_arg0)
	runtime.KeepAlive(area)

	var _glContext gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gdk.GLContexter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.GLContexter")
		}
		_glContext = rv
	}

	return _glContext
}

// Error gets the current error set on the area.
func (area *GLArea) Error() error {
	var _arg0 *C.GtkGLArea // out
	var _cret *C.GError    // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_error(_arg0)
	runtime.KeepAlive(area)

	var _err error // out

	if _cret != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// HasAlpha returns whether the area has an alpha component.
func (area *GLArea) HasAlpha() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_has_alpha(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasDepthBuffer returns whether the area has a depth buffer.
func (area *GLArea) HasDepthBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_has_depth_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasStencilBuffer returns whether the area has a stencil buffer.
func (area *GLArea) HasStencilBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_has_stencil_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves the required version of OpenGL set using
// gtk_gl_area_set_required_version().
func (area *GLArea) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gint       // in
	var _arg2 C.gint       // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_get_required_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(area)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// UseES retrieves the value set by gtk_gl_area_set_use_es().
func (area *GLArea) UseES() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_gl_area_get_use_es(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent ensures that the GLContext used by area is associated with the
// GLArea.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) MakeCurrent() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_make_current(_arg0)
	runtime.KeepAlive(area)
}

// QueueRender marks the currently rendered data (if any) as invalid, and queues
// a redraw of the widget, ensuring that the GLArea::render signal is emitted
// during the draw.
//
// This is only needed when the gtk_gl_area_set_auto_render() has been called
// with a FALSE value. The default behaviour is to emit GLArea::render on each
// draw.
func (area *GLArea) QueueRender() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))

	C.gtk_gl_area_queue_render(_arg0)
	runtime.KeepAlive(area)
}

// SetAutoRender: if auto_render is TRUE the GLArea::render signal will be
// emitted every time the widget draws. This is the default and is useful if
// drawing the widget is faster.
//
// If auto_render is FALSE the data from previous rendering is kept around and
// will be used for drawing the widget the next time, unless the window is
// resized. In order to force a rendering gtk_gl_area_queue_render() must be
// called. This mode is useful when the scene changes seldomly, but takes a long
// time to redraw.
func (area *GLArea) SetAutoRender(autoRender bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if autoRender {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_auto_render(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(autoRender)
}

// SetError sets an error on the area which will be shown instead of the GL
// rendering. This is useful in the GLArea::create-context signal if GL context
// creation fails.
func (area *GLArea) SetError(err error) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 *C.GError    // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.gtk_gl_area_set_error(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(err)
}

// SetHasAlpha: if has_alpha is TRUE the buffer allocated by the widget will
// have an alpha channel component, and when rendering to the window the result
// will be composited over whatever is below the widget.
//
// If has_alpha is FALSE there will be no alpha channel, and the buffer will
// fully replace anything below the widget.
func (area *GLArea) SetHasAlpha(hasAlpha bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if hasAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_alpha(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasAlpha)
}

// SetHasDepthBuffer: if has_depth_buffer is TRUE the widget will allocate and
// enable a depth buffer for the target framebuffer. Otherwise there will be
// none.
func (area *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if hasDepthBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_depth_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasDepthBuffer)
}

// SetHasStencilBuffer: if has_stencil_buffer is TRUE the widget will allocate
// and enable a stencil buffer for the target framebuffer. Otherwise there will
// be none.
func (area *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if hasStencilBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_stencil_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasStencilBuffer)
}

// SetRequiredVersion sets the required version of OpenGL to be used when
// creating the context for the widget.
//
// This function must be called before the area has been realized.
func (area *GLArea) SetRequiredVersion(major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	_arg1 = C.gint(major)
	_arg2 = C.gint(minor)

	C.gtk_gl_area_set_required_version(_arg0, _arg1, _arg2)
	runtime.KeepAlive(area)
	runtime.KeepAlive(major)
	runtime.KeepAlive(minor)
}

// SetUseES sets whether the area should create an OpenGL or an OpenGL ES
// context.
//
// You should check the capabilities of the GLContext before drawing with either
// API.
func (area *GLArea) SetUseES(useEs bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(area.Native()))
	if useEs {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(useEs)
}
