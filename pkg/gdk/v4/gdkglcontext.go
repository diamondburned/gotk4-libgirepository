// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeGLContext returns the GType for the type GLContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGLContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "GLContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGLContext)
	return gtype
}

// GLContext: GdkGLContext is an object representing a platform-specific OpenGL
// draw context.
//
// GdkGLContexts are created for a surface using gdk.Surface.CreateGLContext(),
// and the context will match the the characteristics of the surface.
//
// A GdkGLContext is not tied to any particular normal framebuffer. For
// instance, it cannot draw to the surface back buffer. The GDK repaint system
// is in full control of the painting to that. Instead, you can create render
// buffers or textures and use cairo_draw_from_gl in the draw function of your
// widget to draw them. Then GDK will handle the integration of your rendering
// with that of other widgets.
//
// Support for GdkGLContext is platform-specific and context creation can fail,
// returning NULL context.
//
// A GdkGLContext has to be made "current" in order to start using it, otherwise
// any OpenGL call will be ignored.
//
//
// Creating a new OpenGL context
//
// In order to create a new GdkGLContext instance you need a GdkSurface, which
// you typically get during the realize call of a widget.
//
// A GdkGLContext is not realized until either gdk.GLContext.MakeCurrent() or
// gdk.GLContext.Realize() is called. It is possible to specify details of the
// GL context like the OpenGL version to be used, or whether the GL context
// should have extra state validation enabled after calling
// gdk.Surface.CreateGLContext() by calling gdk.GLContext.Realize(). If the
// realization fails you have the option to change the settings of the
// GdkGLContext and try again.
//
//
// Using a GdkGLContext
//
// You will need to make the GdkGLContext the current context before issuing
// OpenGL calls; the system sends OpenGL commands to whichever context is
// current. It is possible to have multiple contexts, so you always need to
// ensure that the one which you want to draw with is the current one before
// issuing commands:
//
//    gdk_gl_context_make_current (context);
//
//
// You can now perform your drawing using OpenGL commands.
//
// You can check which GdkGLContext is the current one by using
// gdk.GLContext().GetCurrent; you can also unset any GdkGLContext that is
// currently set by calling gdk.GLContext().ClearCurrent.
type GLContext struct {
	_ [0]func() // equal guard
	DrawContext
}

var (
	_ DrawContexter = (*GLContext)(nil)
)

// GLContexter describes types inherited from class GLContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type GLContexter interface {
	coreglib.Objector
	baseGLContext() *GLContext
}

var _ GLContexter = (*GLContext)(nil)

func wrapGLContext(obj *coreglib.Object) *GLContext {
	return &GLContext{
		DrawContext: DrawContext{
			Object: obj,
		},
	}
}

func marshalGLContext(p uintptr) (interface{}, error) {
	return wrapGLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (context *GLContext) baseGLContext() *GLContext {
	return context
}

// BaseGLContext returns the underlying base object.
func BaseGLContext(obj GLContexter) *GLContext {
	return obj.baseGLContext()
}

// DebugEnabled retrieves whether the context is doing extra validations and
// runtime checking.
//
// See gdk.GLContext.SetDebugEnabled().
//
// The function returns the following values:
//
//    - ok: TRUE if debugging is enabled.
//
func (context *GLContext) DebugEnabled() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_debug_enabled", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Display retrieves the display the context is created for.
//
// The function returns the following values:
//
//    - display (optional): GdkDisplay or NULL.
//
func (context *GLContext) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_display", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _display *Display // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _display
}

// ForwardCompatible retrieves whether the context is forward-compatible.
//
// See gdk.GLContext.SetForwardCompatible().
//
// The function returns the following values:
//
//    - ok: TRUE if the context should be forward-compatible.
//
func (context *GLContext) ForwardCompatible() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_forward_compatible", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves required OpenGL version.
//
// See gdk.GLContext.SetRequiredVersion().
//
// The function returns the following values:
//
//    - major (optional): return location for the major version to request.
//    - minor (optional): return location for the minor version to request.
//
func (context *GLContext) RequiredVersion() (major, minor int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("get_required_version", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _major int32 // out
	var _minor int32 // out

	_major = int32(*(*C.int)(unsafe.Pointer(&_outs[0])))
	_minor = int32(*(*C.int)(unsafe.Pointer(&_outs[1])))

	return _major, _minor
}

// SharedContext retrieves the GdkGLContext that this context share data with.
//
// The function returns the following values:
//
//    - glContext (optional): GdkGLContext or NULL.
//
func (context *GLContext) SharedContext() GLContexter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_shared_context", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _glContext GLContexter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(GLContexter)
				return ok
			})
			rv, ok := casted.(GLContexter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
			}
			_glContext = rv
		}
	}

	return _glContext
}

// Surface retrieves the surface used by the context.
//
// The function returns the following values:
//
//    - surface (optional): GdkSurface or NULL.
//
func (context *GLContext) Surface() Surfacer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_surface", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _surface Surfacer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Surfacer)
				return ok
			})
			rv, ok := casted.(Surfacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// UseES checks whether the context is using an OpenGL or OpenGL ES profile.
//
// The function returns the following values:
//
//    - ok: TRUE if the GdkGLContext is using an OpenGL ES profile.
//
func (context *GLContext) UseES() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("get_use_es", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Version retrieves the OpenGL version of the context.
//
// The context must be realized prior to calling this function.
//
// The function returns the following values:
//
//    - major: return location for the major version.
//    - minor: return location for the minor version.
//
func (context *GLContext) Version() (major, minor int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("get_version", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _major int32 // out
	var _minor int32 // out

	_major = int32(*(*C.int)(unsafe.Pointer(&_outs[0])))
	_minor = int32(*(*C.int)(unsafe.Pointer(&_outs[1])))

	return _major, _minor
}

// IsLegacy: whether the GdkGLContext is in legacy mode or not.
//
// The GdkGLContext must be realized before calling this function.
//
// When realizing a GL context, GDK will try to use the OpenGL 3.2 core profile;
// this profile removes all the OpenGL API that was deprecated prior to the 3.2
// version of the specification. If the realization is successful, this function
// will return FALSE.
//
// If the underlying OpenGL implementation does not support core profiles, GDK
// will fall back to a pre-3.2 compatibility profile, and this function will
// return TRUE.
//
// You can use the value returned by this function to decide which kind of
// OpenGL API to use, or whether to do extension discovery, or what kind of
// shader programs to load.
//
// The function returns the following values:
//
//    - ok: TRUE if the GL context is in legacy mode.
//
func (context *GLContext) IsLegacy() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_gret := _info.InvokeClassMethod("is_legacy", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent makes the context the current one.
func (context *GLContext) MakeCurrent() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("make_current", _args[:], nil)

	runtime.KeepAlive(context)
}

// Realize realizes the given GdkGLContext.
//
// It is safe to call this function on a realized GdkGLContext.
func (context *GLContext) Realize() error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("realize", _args[:], nil)

	runtime.KeepAlive(context)

	var _goerr error // out

	if *(**C.GError)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.GError)(unsafe.Pointer(&_cerr))))
	}

	return _goerr
}

// SetDebugEnabled sets whether the GdkGLContext should perform extra
// validations and runtime checking.
//
// This is useful during development, but has additional overhead.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - enabled: whether to enable debugging in the context.
//
func (context *GLContext) SetDebugEnabled(enabled bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if enabled {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("set_debug_enabled", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(enabled)
}

// SetForwardCompatible sets whether the GdkGLContext should be
// forward-compatible.
//
// Forward-compatible contexts must not support OpenGL functionality that has
// been marked as deprecated in the requested version; non-forward compatible
// contexts, on the other hand, must support both deprecated and non deprecated
// functionality.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - compatible: whether the context should be forward-compatible.
//
func (context *GLContext) SetForwardCompatible(compatible bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if compatible {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("set_forward_compatible", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(compatible)
}

// SetRequiredVersion sets the major and minor version of OpenGL to request.
//
// Setting major and minor to zero will use the default values.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - major version to request.
//    - minor version to request.
//
func (context *GLContext) SetRequiredVersion(major, minor int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(major)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(minor)

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("set_required_version", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(major)
	runtime.KeepAlive(minor)
}

// SetUseES requests that GDK create an OpenGL ES context instead of an OpenGL
// one.
//
// Not all platforms support OpenGL ES.
//
// The context must not have been realized.
//
// By default, GDK will attempt to automatically detect whether the underlying
// GL implementation is OpenGL or OpenGL ES once the context is realized.
//
// You should check the return value of gdk.GLContext.GetUseES() after calling
// gdk.GLContext.Realize() to decide whether to use the OpenGL or OpenGL ES API,
// extensions, or shaders.
//
// The function takes the following parameters:
//
//    - useEs: whether the context should use OpenGL ES instead of OpenGL, or -1
//      to allow auto-detection.
//
func (context *GLContext) SetUseES(useEs int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(useEs)

	_info := girepository.MustFind("Gdk", "GLContext")
	_info.InvokeClassMethod("set_use_es", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(useEs)
}

// GLContextClearCurrent clears the current GdkGLContext.
//
// Any OpenGL call after this function returns will be ignored until
// gdk.GLContext.MakeCurrent() is called.
func GLContextClearCurrent() {
	_info := girepository.MustFind("Gdk", "clear_current")
	_info.InvokeFunction(nil, nil)
}

// GLContextGetCurrent retrieves the current GdkGLContext.
//
// The function returns the following values:
//
//    - glContext (optional): current GdkGLContext, or NULL.
//
func GLContextGetCurrent() GLContexter {
	_info := girepository.MustFind("Gdk", "get_current")
	_gret := _info.InvokeFunction(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _glContext GLContexter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(GLContexter)
				return ok
			})
			rv, ok := casted.(GLContexter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
			}
			_glContext = rv
		}
	}

	return _glContext
}
