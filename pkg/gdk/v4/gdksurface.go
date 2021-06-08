// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_surface_get_type()), F: marshalSurface},
	})
}

// Surface: a `GdkSurface` is a rectangular region on the screen.
//
// It’s a low-level object, used to implement high-level objects such as
// [class@Gtk.Window] or [class@Gtk.Dialog] in GTK.
//
// The surfaces you see in practice are either [class@Gdk.Toplevel] or
// [class@Gdk.Popup], and those interfaces provide much of the required API to
// interact with these surfaces. Other, more specialized surface types exist,
// but you will rarely interact with them directly.
type Surface interface {
	gextras.Objector

	// Beep emits a short beep associated to @surface.
	//
	// If the display of @surface does not support per-surface beeps, emits a
	// short beep on the display just as [method@Gdk.Display.beep].
	Beep()
	// CreateCairoContext creates a new `GdkCairoContext` for rendering on
	// @surface.
	CreateCairoContext() CairoContext
	// CreateGLContext creates a new `GdkGLContext` for the `GdkSurface`.
	//
	// The context is disconnected from any particular surface or surface. If
	// the creation of the `GdkGLContext` failed, @error will be set. Before
	// using the returned `GdkGLContext`, you will need to call
	// [method@Gdk.GLContext.make_current] or [method@Gdk.GLContext.realize].
	CreateGLContext() (glContext GLContext, err error)
	// CreateSimilarSurface: create a new Cairo surface that is as compatible as
	// possible with the given @surface.
	//
	// For example the new surface will have the same fallback resolution and
	// font options as @surface. Generally, the new surface will also use the
	// same backend as @surface, unless that is not possible for some reason.
	// The type of the returned surface may be examined with
	// cairo_surface_get_type().
	//
	// Initially the surface contents are all 0 (transparent if contents have
	// transparency, black otherwise.)
	//
	// This function always returns a valid pointer, but it will return a
	// pointer to a “nil” surface if @other is already in an error state or any
	// other error occurs.
	CreateSimilarSurface(content cairo.Content, width int, height int) *cairo.Surface
	// CreateVulkanContext creates a new `GdkVulkanContext` for rendering on
	// @surface.
	//
	// If the creation of the `GdkVulkanContext` failed, @error will be set.
	CreateVulkanContext() (vulkanContext VulkanContext, err error)
	// Destroy destroys the window system resources associated with @surface and
	// decrements @surface's reference count.
	//
	// The window system resources for all children of @surface are also
	// destroyed, but the children’s reference counts are not decremented.
	//
	// Note that a surface will not be destroyed automatically when its
	// reference count reaches zero. You must call this function yourself before
	// that happens.
	Destroy()
	// Cursor retrieves a `GdkCursor` pointer for the cursor currently set on
	// the `GdkSurface`.
	//
	// If the return value is nil then there is no custom cursor set on the
	// surface, and it is using the cursor for its parent surface.
	Cursor() Cursor
	// DeviceCursor retrieves a `GdkCursor` pointer for the @device currently
	// set on the specified `GdkSurface`.
	//
	// If the return value is nil then there is no custom cursor set on the
	// specified surface, and it is using the cursor for its parent surface.
	DeviceCursor(device Device) Cursor
	// DevicePosition obtains the current device position and modifier state.
	//
	// The position is given in coordinates relative to the upper left corner of
	// @surface.
	DevicePosition(device Device) (x float64, y float64, mask *ModifierType, ok bool)
	// Display gets the `GdkDisplay` associated with a `GdkSurface`.
	Display() Display
	// FrameClock gets the frame clock for the surface.
	//
	// The frame clock for a surface never changes unless the surface is
	// reparented to a new toplevel surface.
	FrameClock() FrameClock
	// Height returns the height of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Height() int
	// Mapped checks whether the surface has been mapped.
	//
	// A surface is mapped with [method@Gdk.Toplevel.present] or
	// [method@Gdk.Popup.present].
	Mapped() bool
	// ScaleFactor returns the internal scale factor that maps from surface
	// coordinates to the actual device pixels.
	//
	// On traditional systems this is 1, but on very high density outputs this
	// can be a higher value (often 2). A higher value means that drawing is
	// automatically scaled up to a higher resolution, so any code doing drawing
	// will automatically look nicer. However, if you are supplying pixel-based
	// data the scale value can be used to determine whether to use a pixel
	// resource with higher resolution data.
	//
	// The scale of a surface may change during runtime.
	ScaleFactor() int
	// Width returns the width of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Width() int
	// Hide: hide the surface.
	//
	// For toplevel surfaces, withdraws them, so they will no longer be known to
	// the window manager; for all surfaces, unmaps them, so they won’t be
	// displayed. Normally done automatically as part of
	// [method@Gtk.Widget.hide].
	Hide()
	// IsDestroyed: check to see if a surface is destroyed.
	IsDestroyed() bool
	// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
	// @surface to be scheduled.
	//
	// This function is useful for implementations that track invalid regions on
	// their own.
	QueueRender()
	// RequestLayout: request a layout phase from the surface's frame clock.
	//
	// See [method@Gdk.FrameClock.request_phase].
	RequestLayout()
	// SetCursor sets the default mouse pointer for a `GdkSurface`.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default. Note
	// that @cursor must be for the same display as @surface.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetCursor(cursor Cursor)
	// SetDeviceCursor sets a specific `GdkCursor` for a given device when it
	// gets inside @surface.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetDeviceCursor(device Device, cursor Cursor)
	// SetInputRegion: apply the region to the surface for the purpose of event
	// handling.
	//
	// Mouse events which happen while the pointer position corresponds to an
	// unset bit in the mask will be passed on the surface below @surface.
	//
	// An input region is typically used with RGBA surfaces. The alpha channel
	// of the surface defines which pixels are invisible and allows for nicely
	// antialiased borders, and the input region controls where the surface is
	// “clickable”.
	//
	// Use [method@Gdk.Display.supports_input_shapes] to find out if a
	// particular backend supports input regions.
	SetInputRegion(region *cairo.Region)
	// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
	//
	// For optimisation purposes, compositing window managers may like to not
	// draw obscured regions of surfaces, or turn off blending during for these
	// regions. With RGB windows with no transparency, this is just the shape of
	// the window, but with ARGB32 windows, the compositor does not know what
	// regions of the window are transparent or not.
	//
	// This function only works for toplevel surfaces.
	//
	// GTK will update this property automatically if the @surface background is
	// opaque, as we know where the opaque regions are. If your surface
	// background is not opaque, please update this property in your
	// WidgetClass.css_changed() handler.
	SetOpaqueRegion(region *cairo.Region)
	// TranslateCoordinates translates coordinates between two surfaces.
	//
	// Note that this only works if @to and @from are popups or transient-for to
	// the same toplevel (directly or indirectly).
	TranslateCoordinates(to Surface, x float64, y float64) bool
}

// surface implements the Surface interface.
type surface struct {
	gextras.Objector
}

var _ Surface = (*surface)(nil)

// WrapSurface wraps a GObject to the right type. It is
// primarily used internally.
func WrapSurface(obj *externglib.Object) Surface {
	return Surface{
		Objector: obj,
	}
}

func marshalSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSurface(obj), nil
}

// NewSurfacePopup constructs a class Surface.
func NewSurfacePopup(parent Surface, autohide bool) Surface {
	var arg1 *C.GdkSurface
	var arg2 C.gboolean

	arg1 = (*C.GdkSurface)(unsafe.Pointer(parent.Native()))
	if autohide {
		arg2 = C.gboolean(1)
	}

	cret := new(C.GdkSurface)
	var goret Surface

	cret = C.gdk_surface_new_popup(arg1, arg2)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Surface)

	return goret
}

// NewSurfaceToplevel constructs a class Surface.
func NewSurfaceToplevel(display Display) Surface {
	var arg1 *C.GdkDisplay

	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	cret := new(C.GdkSurface)
	var goret Surface

	cret = C.gdk_surface_new_toplevel(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Surface)

	return goret
}

// Beep emits a short beep associated to @surface.
//
// If the display of @surface does not support per-surface beeps, emits a
// short beep on the display just as [method@Gdk.Display.beep].
func (s surface) Beep() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_beep(arg0)
}

// CreateCairoContext creates a new `GdkCairoContext` for rendering on
// @surface.
func (s surface) CreateCairoContext() CairoContext {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	cret := new(C.GdkCairoContext)
	var goret CairoContext

	cret = C.gdk_surface_create_cairo_context(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(CairoContext)

	return goret
}

// CreateGLContext creates a new `GdkGLContext` for the `GdkSurface`.
//
// The context is disconnected from any particular surface or surface. If
// the creation of the `GdkGLContext` failed, @error will be set. Before
// using the returned `GdkGLContext`, you will need to call
// [method@Gdk.GLContext.make_current] or [method@Gdk.GLContext.realize].
func (s surface) CreateGLContext() (glContext GLContext, err error) {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	cret := new(C.GdkGLContext)
	var goret GLContext
	var cerr *C.GError
	var goerr error

	cret = C.gdk_surface_create_gl_context(arg0, &cerr)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GLContext)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

// CreateSimilarSurface: create a new Cairo surface that is as compatible as
// possible with the given @surface.
//
// For example the new surface will have the same fallback resolution and
// font options as @surface. Generally, the new surface will also use the
// same backend as @surface, unless that is not possible for some reason.
// The type of the returned surface may be examined with
// cairo_surface_get_type().
//
// Initially the surface contents are all 0 (transparent if contents have
// transparency, black otherwise.)
//
// This function always returns a valid pointer, but it will return a
// pointer to a “nil” surface if @other is already in an error state or any
// other error occurs.
func (s surface) CreateSimilarSurface(content cairo.Content, width int, height int) *cairo.Surface {
	var arg0 *C.GdkSurface
	var arg1 C.cairo_content_t
	var arg2 C.int
	var arg3 C.int

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (C.cairo_content_t)(content)
	arg2 = C.int(width)
	arg3 = C.int(height)

	cret := new(C.cairo_surface_t)
	var goret *cairo.Surface

	cret = C.gdk_surface_create_similar_surface(arg0, arg1, arg2, arg3)

	goret = cairo.WrapSurface(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *cairo.Surface) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// CreateVulkanContext creates a new `GdkVulkanContext` for rendering on
// @surface.
//
// If the creation of the `GdkVulkanContext` failed, @error will be set.
func (s surface) CreateVulkanContext() (vulkanContext VulkanContext, err error) {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	cret := new(C.GdkVulkanContext)
	var goret VulkanContext
	var cerr *C.GError
	var goerr error

	cret = C.gdk_surface_create_vulkan_context(arg0, &cerr)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(VulkanContext)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goret, goerr
}

// Destroy destroys the window system resources associated with @surface and
// decrements @surface's reference count.
//
// The window system resources for all children of @surface are also
// destroyed, but the children’s reference counts are not decremented.
//
// Note that a surface will not be destroyed automatically when its
// reference count reaches zero. You must call this function yourself before
// that happens.
func (s surface) Destroy() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_destroy(arg0)
}

// Cursor retrieves a `GdkCursor` pointer for the cursor currently set on
// the `GdkSurface`.
//
// If the return value is nil then there is no custom cursor set on the
// surface, and it is using the cursor for its parent surface.
func (s surface) Cursor() Cursor {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret *C.GdkCursor
	var goret Cursor

	cret = C.gdk_surface_get_cursor(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Cursor)

	return goret
}

// DeviceCursor retrieves a `GdkCursor` pointer for the @device currently
// set on the specified `GdkSurface`.
//
// If the return value is nil then there is no custom cursor set on the
// specified surface, and it is using the cursor for its parent surface.
func (s surface) DeviceCursor(device Device) Cursor {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkDevice

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	var cret *C.GdkCursor
	var goret Cursor

	cret = C.gdk_surface_get_device_cursor(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Cursor)

	return goret
}

// DevicePosition obtains the current device position and modifier state.
//
// The position is given in coordinates relative to the upper left corner of
// @surface.
func (s surface) DevicePosition(device Device) (x float64, y float64, mask *ModifierType, ok bool) {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkDevice

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	arg2 := new(C.double)
	var ret2 float64
	arg3 := new(C.double)
	var ret3 float64
	arg4 := new(C.GdkModifierType)
	var ret4 *ModifierType
	var cret C.gboolean
	var goret bool

	cret = C.gdk_surface_get_device_position(arg0, arg1, arg2, arg3, arg4)

	ret2 = float64(*arg2)
	ret3 = float64(*arg3)
	ret4 = *ModifierType(arg4)
	if cret {
		goret = true
	}

	return ret2, ret3, ret4, goret
}

// Display gets the `GdkDisplay` associated with a `GdkSurface`.
func (s surface) Display() Display {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret *C.GdkDisplay
	var goret Display

	cret = C.gdk_surface_get_display(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Display)

	return goret
}

// FrameClock gets the frame clock for the surface.
//
// The frame clock for a surface never changes unless the surface is
// reparented to a new toplevel surface.
func (s surface) FrameClock() FrameClock {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret *C.GdkFrameClock
	var goret FrameClock

	cret = C.gdk_surface_get_frame_clock(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FrameClock)

	return goret
}

// Height returns the height of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels”
// (see [method@Gdk.Surface.get_scale_factor]).
func (s surface) Height() int {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.int
	var goret int

	cret = C.gdk_surface_get_height(arg0)

	goret = int(cret)

	return goret
}

// Mapped checks whether the surface has been mapped.
//
// A surface is mapped with [method@Gdk.Toplevel.present] or
// [method@Gdk.Popup.present].
func (s surface) Mapped() bool {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gdk_surface_get_mapped(arg0)

	if cret {
		goret = true
	}

	return goret
}

// ScaleFactor returns the internal scale factor that maps from surface
// coordinates to the actual device pixels.
//
// On traditional systems this is 1, but on very high density outputs this
// can be a higher value (often 2). A higher value means that drawing is
// automatically scaled up to a higher resolution, so any code doing drawing
// will automatically look nicer. However, if you are supplying pixel-based
// data the scale value can be used to determine whether to use a pixel
// resource with higher resolution data.
//
// The scale of a surface may change during runtime.
func (s surface) ScaleFactor() int {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.int
	var goret int

	cret = C.gdk_surface_get_scale_factor(arg0)

	goret = int(cret)

	return goret
}

// Width returns the width of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels”
// (see [method@Gdk.Surface.get_scale_factor]).
func (s surface) Width() int {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.int
	var goret int

	cret = C.gdk_surface_get_width(arg0)

	goret = int(cret)

	return goret
}

// Hide: hide the surface.
//
// For toplevel surfaces, withdraws them, so they will no longer be known to
// the window manager; for all surfaces, unmaps them, so they won’t be
// displayed. Normally done automatically as part of
// [method@Gtk.Widget.hide].
func (s surface) Hide() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_hide(arg0)
}

// IsDestroyed: check to see if a surface is destroyed.
func (s surface) IsDestroyed() bool {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gdk_surface_is_destroyed(arg0)

	if cret {
		goret = true
	}

	return goret
}

// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
// @surface to be scheduled.
//
// This function is useful for implementations that track invalid regions on
// their own.
func (s surface) QueueRender() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_queue_render(arg0)
}

// RequestLayout: request a layout phase from the surface's frame clock.
//
// See [method@Gdk.FrameClock.request_phase].
func (s surface) RequestLayout() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_request_layout(arg0)
}

// SetCursor sets the default mouse pointer for a `GdkSurface`.
//
// Passing nil for the @cursor argument means that @surface will use the
// cursor of its parent surface. Most surfaces should use this default. Note
// that @cursor must be for the same display as @surface.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s surface) SetCursor(cursor Cursor) {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkCursor

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_cursor(arg0, arg1)
}

// SetDeviceCursor sets a specific `GdkCursor` for a given device when it
// gets inside @surface.
//
// Passing nil for the @cursor argument means that @surface will use the
// cursor of its parent surface. Most surfaces should use this default.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s surface) SetDeviceCursor(device Device, cursor Cursor) {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkDevice
	var arg2 *C.GdkCursor

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	arg2 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_device_cursor(arg0, arg1, arg2)
}

// SetInputRegion: apply the region to the surface for the purpose of event
// handling.
//
// Mouse events which happen while the pointer position corresponds to an
// unset bit in the mask will be passed on the surface below @surface.
//
// An input region is typically used with RGBA surfaces. The alpha channel
// of the surface defines which pixels are invisible and allows for nicely
// antialiased borders, and the input region controls where the surface is
// “clickable”.
//
// Use [method@Gdk.Display.supports_input_shapes] to find out if a
// particular backend supports input regions.
func (s surface) SetInputRegion(region *cairo.Region) {
	var arg0 *C.GdkSurface
	var arg1 *C.cairo_region_t

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_input_region(arg0, arg1)
}

// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
//
// For optimisation purposes, compositing window managers may like to not
// draw obscured regions of surfaces, or turn off blending during for these
// regions. With RGB windows with no transparency, this is just the shape of
// the window, but with ARGB32 windows, the compositor does not know what
// regions of the window are transparent or not.
//
// This function only works for toplevel surfaces.
//
// GTK will update this property automatically if the @surface background is
// opaque, as we know where the opaque regions are. If your surface
// background is not opaque, please update this property in your
// WidgetClass.css_changed() handler.
func (s surface) SetOpaqueRegion(region *cairo.Region) {
	var arg0 *C.GdkSurface
	var arg1 *C.cairo_region_t

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_opaque_region(arg0, arg1)
}

// TranslateCoordinates translates coordinates between two surfaces.
//
// Note that this only works if @to and @from are popups or transient-for to
// the same toplevel (directly or indirectly).
func (f surface) TranslateCoordinates(to Surface, x float64, y float64) bool {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkSurface
	var arg2 *C.double
	var arg3 *C.double

	arg0 = (*C.GdkSurface)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkSurface)(unsafe.Pointer(to.Native()))
	arg2 = *C.double(x)
	arg3 = *C.double(y)

	var cret C.gboolean
	var goret bool

	cret = C.gdk_surface_translate_coordinates(arg0, arg1, arg2, arg3)

	if cret {
		goret = true
	}

	return goret
}
