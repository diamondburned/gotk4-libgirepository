// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gl_area_get_type()), F: marshalGLArea},
	})
}

// GLArea is a widget that allows drawing with OpenGL.
//
// GLArea sets up its own GLContext for the window it creates, and creates a
// custom GL framebuffer that the widget will do GL rendering onto. It also
// ensures that this framebuffer is the default GL rendering target when
// rendering.
//
// In order to draw, you have to connect to the GLArea::render signal, or
// subclass GLArea and override the @GtkGLAreaClass.render() virtual function.
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
type GLArea interface {
	Widget
	Buildable

	// AttachBuffers ensures that the @area framebuffer object is made the
	// current draw and read target, and that all the required buffers for the
	// @area are created and bound to the frambuffer.
	//
	// This function is automatically called before emitting the GLArea::render
	// signal, and doesn't normally need to be called by application code.
	AttachBuffers()
	// AutoRender returns whether the area is in auto render mode or not.
	AutoRender() bool
	// Context retrieves the GLContext used by @area.
	Context() gdk.GLContext
	// Error gets the current error set on the @area.
	Error() *error
	// HasAlpha returns whether the area has an alpha component.
	HasAlpha() bool
	// HasDepthBuffer returns whether the area has a depth buffer.
	HasDepthBuffer() bool
	// HasStencilBuffer returns whether the area has a stencil buffer.
	HasStencilBuffer() bool
	// RequiredVersion retrieves the required version of OpenGL set using
	// gtk_gl_area_set_required_version().
	RequiredVersion() (major int, minor int)
	// UseES retrieves the value set by gtk_gl_area_set_use_es().
	UseES() bool
	// MakeCurrent ensures that the GLContext used by @area is associated with
	// the GLArea.
	//
	// This function is automatically called before emitting the GLArea::render
	// signal, and doesn't normally need to be called by application code.
	MakeCurrent()
	// QueueRender marks the currently rendered data (if any) as invalid, and
	// queues a redraw of the widget, ensuring that the GLArea::render signal is
	// emitted during the draw.
	//
	// This is only needed when the gtk_gl_area_set_auto_render() has been
	// called with a false value. The default behaviour is to emit
	// GLArea::render on each draw.
	QueueRender()
	// SetAutoRender: if @auto_render is true the GLArea::render signal will be
	// emitted every time the widget draws. This is the default and is useful if
	// drawing the widget is faster.
	//
	// If @auto_render is false the data from previous rendering is kept around
	// and will be used for drawing the widget the next time, unless the window
	// is resized. In order to force a rendering gtk_gl_area_queue_render() must
	// be called. This mode is useful when the scene changes seldomly, but takes
	// a long time to redraw.
	SetAutoRender(autoRender bool)
	// SetError sets an error on the area which will be shown instead of the GL
	// rendering. This is useful in the GLArea::create-context signal if GL
	// context creation fails.
	SetError(err *error)
	// SetHasAlpha: if @has_alpha is true the buffer allocated by the widget
	// will have an alpha channel component, and when rendering to the window
	// the result will be composited over whatever is below the widget.
	//
	// If @has_alpha is false there will be no alpha channel, and the buffer
	// will fully replace anything below the widget.
	SetHasAlpha(hasAlpha bool)
	// SetHasDepthBuffer: if @has_depth_buffer is true the widget will allocate
	// and enable a depth buffer for the target framebuffer. Otherwise there
	// will be none.
	SetHasDepthBuffer(hasDepthBuffer bool)
	// SetHasStencilBuffer: if @has_stencil_buffer is true the widget will
	// allocate and enable a stencil buffer for the target framebuffer.
	// Otherwise there will be none.
	SetHasStencilBuffer(hasStencilBuffer bool)
	// SetRequiredVersion sets the required version of OpenGL to be used when
	// creating the context for the widget.
	//
	// This function must be called before the area has been realized.
	SetRequiredVersion(major int, minor int)
	// SetUseES sets whether the @area should create an OpenGL or an OpenGL ES
	// context.
	//
	// You should check the capabilities of the GLContext before drawing with
	// either API.
	SetUseES(useEs bool)
}

// glArea implements the GLArea interface.
type glArea struct {
	Widget
	Buildable
}

var _ GLArea = (*glArea)(nil)

// WrapGLArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLArea(obj *externglib.Object) GLArea {
	return GLArea{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalGLArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLArea(obj), nil
}

// NewGLArea constructs a class GLArea.
func NewGLArea() GLArea {
	var _cret C.GtkGLArea

	cret = C.gtk_gl_area_new()

	var _glArea GLArea

	_glArea = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(GLArea)

	return _glArea
}

// AttachBuffers ensures that the @area framebuffer object is made the
// current draw and read target, and that all the required buffers for the
// @area are created and bound to the frambuffer.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (a glArea) AttachBuffers() {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_attach_buffers(_arg0)
}

// AutoRender returns whether the area is in auto render mode or not.
func (a glArea) AutoRender() bool {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_gl_area_get_auto_render(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Context retrieves the GLContext used by @area.
func (a glArea) Context() gdk.GLContext {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret *C.GdkGLContext

	cret = C.gtk_gl_area_get_context(_arg0)

	var _glContext gdk.GLContext

	_glContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gdk.GLContext)

	return _glContext
}

// Error gets the current error set on the @area.
func (a glArea) Error() *error {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret *C.GError

	cret = C.gtk_gl_area_get_error(_arg0)

	var _err *error

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

// HasAlpha returns whether the area has an alpha component.
func (a glArea) HasAlpha() bool {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_gl_area_get_has_alpha(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// HasDepthBuffer returns whether the area has a depth buffer.
func (a glArea) HasDepthBuffer() bool {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_gl_area_get_has_depth_buffer(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// HasStencilBuffer returns whether the area has a stencil buffer.
func (a glArea) HasStencilBuffer() bool {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_gl_area_get_has_stencil_buffer(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves the required version of OpenGL set using
// gtk_gl_area_set_required_version().
func (a glArea) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _arg1 C.gint
	var _arg2 C.gint

	C.gtk_gl_area_get_required_version(_arg0, &_arg1, &_arg2)

	var _major int
	var _minor int

	_major = (int)(_arg1)
	_minor = (int)(_arg2)

	return _major, _minor
}

// UseES retrieves the value set by gtk_gl_area_set_use_es().
func (a glArea) UseES() bool {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.gtk_gl_area_get_use_es(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// MakeCurrent ensures that the GLContext used by @area is associated with
// the GLArea.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (a glArea) MakeCurrent() {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_make_current(_arg0)
}

// QueueRender marks the currently rendered data (if any) as invalid, and
// queues a redraw of the widget, ensuring that the GLArea::render signal is
// emitted during the draw.
//
// This is only needed when the gtk_gl_area_set_auto_render() has been
// called with a false value. The default behaviour is to emit
// GLArea::render on each draw.
func (a glArea) QueueRender() {
	var _arg0 *C.GtkGLArea

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))

	C.gtk_gl_area_queue_render(_arg0)
}

// SetAutoRender: if @auto_render is true the GLArea::render signal will be
// emitted every time the widget draws. This is the default and is useful if
// drawing the widget is faster.
//
// If @auto_render is false the data from previous rendering is kept around
// and will be used for drawing the widget the next time, unless the window
// is resized. In order to force a rendering gtk_gl_area_queue_render() must
// be called. This mode is useful when the scene changes seldomly, but takes
// a long time to redraw.
func (a glArea) SetAutoRender(autoRender bool) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if autoRender {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gl_area_set_auto_render(_arg0, _arg1)
}

// SetError sets an error on the area which will be shown instead of the GL
// rendering. This is useful in the GLArea::create-context signal if GL
// context creation fails.
func (a glArea) SetError(err *error) {
	var _arg0 *C.GtkGLArea
	var _arg1 *C.GError

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GError)(gerror.New(unsafe.Pointer(err)))
	defer C.g_error_free(_arg1)

	C.gtk_gl_area_set_error(_arg0, _arg1)
}

// SetHasAlpha: if @has_alpha is true the buffer allocated by the widget
// will have an alpha channel component, and when rendering to the window
// the result will be composited over whatever is below the widget.
//
// If @has_alpha is false there will be no alpha channel, and the buffer
// will fully replace anything below the widget.
func (a glArea) SetHasAlpha(hasAlpha bool) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if hasAlpha {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gl_area_set_has_alpha(_arg0, _arg1)
}

// SetHasDepthBuffer: if @has_depth_buffer is true the widget will allocate
// and enable a depth buffer for the target framebuffer. Otherwise there
// will be none.
func (a glArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if hasDepthBuffer {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gl_area_set_has_depth_buffer(_arg0, _arg1)
}

// SetHasStencilBuffer: if @has_stencil_buffer is true the widget will
// allocate and enable a stencil buffer for the target framebuffer.
// Otherwise there will be none.
func (a glArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if hasStencilBuffer {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gl_area_set_has_stencil_buffer(_arg0, _arg1)
}

// SetRequiredVersion sets the required version of OpenGL to be used when
// creating the context for the widget.
//
// This function must be called before the area has been realized.
func (a glArea) SetRequiredVersion(major int, minor int) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gint
	var _arg2 C.gint

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(major)
	_arg2 = C.gint(minor)

	C.gtk_gl_area_set_required_version(_arg0, _arg1, _arg2)
}

// SetUseES sets whether the @area should create an OpenGL or an OpenGL ES
// context.
//
// You should check the capabilities of the GLContext before drawing with
// either API.
func (a glArea) SetUseES(useEs bool) {
	var _arg0 *C.GtkGLArea
	var _arg1 C.gboolean

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(a.Native()))
	if useEs {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
}