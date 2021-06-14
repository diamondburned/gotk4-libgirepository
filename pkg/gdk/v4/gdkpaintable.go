// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_paintable_flags_get_type()), F: marshalPaintableFlags},
		{T: externglib.Type(C.gdk_paintable_get_type()), F: marshalPaintable},
	})
}

// PaintableFlags flags about a paintable object.
//
// Implementations use these for optimizations such as caching.
type PaintableFlags int

const (
	// PaintableFlagsSize: the size is immutable. The
	// [signal@GdkPaintable::invalidate-size] signal will never be emitted.
	PaintableFlagsSize PaintableFlags = 1
	// PaintableFlagsContents: the content is immutable. The
	// [signal@GdkPaintable::invalidate-contents] signal will never be emitted.
	PaintableFlagsContents PaintableFlags = 2
)

func marshalPaintableFlags(p uintptr) (interface{}, error) {
	return PaintableFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PaintableOverrider contains methods that are overridable. This
// interface is a subset of the interface Paintable.
type PaintableOverrider interface {
	// IntrinsicAspectRatio gets the preferred aspect ratio the @paintable would
	// like to be displayed at.
	//
	// The aspect ratio is the width divided by the height, so a value of 0.5
	// means that the @paintable prefers to be displayed twice as high as it is
	// wide. Consumers of this interface can use this to preserve aspect ratio
	// when displaying the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to [method@Gdk.Paintable.snapshot].
	//
	// Usually when a @paintable returns nonzero values from
	// [method@Gdk.Paintable.get_intrinsic_width] and
	// [method@Gdk.Paintable.get_intrinsic_height] the aspect ratio should
	// conform to those values, though that is not required.
	//
	// If the @paintable does not have a preferred aspect ratio, it returns 0.
	// Negative values are never returned.
	IntrinsicAspectRatio() float64
	// IntrinsicHeight gets the preferred height the @paintable would like to be
	// displayed at.
	//
	// Consumers of this interface can use this to reserve enough space to draw
	// the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to [method@Gdk.Paintable.snapshot].
	//
	// If the @paintable does not have a preferred height, it returns 0.
	// Negative values are never returned.
	IntrinsicHeight() int
	// IntrinsicWidth gets the preferred width the @paintable would like to be
	// displayed at.
	//
	// Consumers of this interface can use this to reserve enough space to draw
	// the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to [method@Gdk.Paintable.snapshot].
	//
	// If the @paintable does not have a preferred width, it returns 0. Negative
	// values are never returned.
	IntrinsicWidth() int
	// Snapshot snapshots the given paintable with the given @width and @height.
	//
	// The paintable is drawn at the current (0,0) offset of the @snapshot. If
	// @width and @height are not larger than zero, this function will do
	// nothing.
	Snapshot(snapshot Snapshot, width float64, height float64)
}

// Paintable: `GdkPaintable` is a simple interface used by GTK to represent
// content that can be painted.
//
// The content of a `GdkPaintable` can be painted anywhere at any size without
// requiring any sort of layout. The interface is inspired by similar concepts
// elsewhere, such as ClutterContent
// (https://developer.gnome.org/clutter/stable/ClutterContent.html), HTML/CSS
// Paint Sources (https://www.w3.org/TR/css-images-4/#paint-source), or SVG
// Paint Servers (https://www.w3.org/TR/SVG2/pservers.html).
//
// A `GdkPaintable` can be snapshot at any time and size using
// [method@Gdk.Paintable.snapshot]. How the paintable interprets that size and
// if it scales or centers itself into the given rectangle is implementation
// defined, though if you are implementing a `GdkPaintable` and don't know what
// to do, it is suggested that you scale your paintable ignoring any potential
// aspect ratio.
//
// The contents that a `GdkPaintable` produces may depend on the
// [class@GdkSnapshot] passed to it. For example, paintables may decide to use
// more detailed images on higher resolution screens or when OpenGL is
// available. A `GdkPaintable` will however always produce the same output for
// the same snapshot.
//
// A `GdkPaintable` may change its contents, meaning that it will now produce a
// different output with the same snapshot. Once that happens, it will call
// [method@Gdk.Paintable.invalidate_contents] which will emit the
// [signal@GdkPaintable::invalidate-contents] signal. If a paintable is known to
// never change its contents, it will set the GDK_PAINTABLE_STATIC_CONTENTS
// flag. If a consumer cannot deal with changing contents, it may call
// [method@Gdk.Paintable.get_current_image] which will return a static paintable
// and use that.
//
// A paintable can report an intrinsic (or preferred) size or aspect ratio it
// wishes to be rendered at, though it doesn't have to. Consumers of the
// interface can use this information to layout thepaintable appropriately. Just
// like the contents, the size of a paintable can change. A paintable will
// indicate this by calling [method@Gdk.Paintable.invalidate_size] which will
// emit the [signal@GdkPaintable::invalidate-size] signal. And just like for
// contents, if a paintable is known to never change its size, it will set the
// GDK_PAINTABLE_STATIC_SIZE flag.
//
// Besides API for applications, there are some functions that are only useful
// for implementing subclasses and should not be used by applications:
// [method@Gdk.Paintable.invalidate_contents],
// [method@Gdk.Paintable.invalidate_size], [func@Gdk.Paintable.new_empty].
type Paintable interface {
	gextras.Objector
	PaintableOverrider

	// ComputeConcreteSize: compute a concrete size for the `GdkPaintable`.
	//
	// Applies the sizing algorithm outlined in the CSS Image spec
	// (https://drafts.csswg.org/css-images-3/#default-sizing) to the given
	// @paintable. See that link for more details.
	//
	// It is not necessary to call this function when both @specified_width and
	// @specified_height are known, but it is useful to call this function in
	// GtkWidget:measure implementations to compute the other dimension when
	// only one dimension is given.
	ComputeConcreteSize(specifiedWidth float64, specifiedHeight float64, defaultWidth float64, defaultHeight float64) (concreteWidth float64, concreteHeight float64)
	// InvalidateContents: called by implementations of `GdkPaintable` to
	// invalidate their contents.
	//
	// Unless the contents are invalidated, implementations must guarantee that
	// multiple calls of [method@Gdk.Paintable.snapshot] produce the same
	// output.
	//
	// This function will emit the [signal@Gdk.Paintable::invalidate-contents]
	// signal.
	//
	// If a @paintable reports the GDK_PAINTABLE_STATIC_CONTENTS flag, it must
	// not call this function.
	InvalidateContents()
	// InvalidateSize: called by implementations of `GdkPaintable` to invalidate
	// their size.
	//
	// As long as the size is not invalidated, @paintable must return the same
	// values for its intrinsic width, height and aspect ratio.
	//
	// This function will emit the [signal@Gdk.Paintable::invalidate-size]
	// signal.
	//
	// If a @paintable reports the GDK_PAINTABLE_STATIC_SIZE flag, it must not
	// call this function.
	InvalidateSize()
}

// paintable implements the Paintable interface.
type paintable struct {
	gextras.Objector
}

var _ Paintable = (*paintable)(nil)

// WrapPaintable wraps a GObject to a type that implements interface
// Paintable. It is primarily used internally.
func WrapPaintable(obj *externglib.Object) Paintable {
	return Paintable{
		Objector: obj,
	}
}

func marshalPaintable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPaintable(obj), nil
}

// ComputeConcreteSize: compute a concrete size for the `GdkPaintable`.
//
// Applies the sizing algorithm outlined in the CSS Image spec
// (https://drafts.csswg.org/css-images-3/#default-sizing) to the given
// @paintable. See that link for more details.
//
// It is not necessary to call this function when both @specified_width and
// @specified_height are known, but it is useful to call this function in
// GtkWidget:measure implementations to compute the other dimension when
// only one dimension is given.
func (p paintable) ComputeConcreteSize(specifiedWidth float64, specifiedHeight float64, defaultWidth float64, defaultHeight float64) (concreteWidth float64, concreteHeight float64) {
	var _arg0 *C.GdkPaintable // out
	var _arg1 C.double        // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))
	_arg1 = C.double(specifiedWidth)
	_arg2 = C.double(specifiedHeight)
	_arg3 = C.double(defaultWidth)
	_arg4 = C.double(defaultHeight)

	var _arg5 C.double // in
	var _arg6 C.double // in

	C.gdk_paintable_compute_concrete_size(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)

	var _concreteWidth float64  // out
	var _concreteHeight float64 // out

	_concreteWidth = (float64)(_arg5)
	_concreteHeight = (float64)(_arg6)

	return _concreteWidth, _concreteHeight
}

// IntrinsicAspectRatio gets the preferred aspect ratio the @paintable would
// like to be displayed at.
//
// The aspect ratio is the width divided by the height, so a value of 0.5
// means that the @paintable prefers to be displayed twice as high as it is
// wide. Consumers of this interface can use this to preserve aspect ratio
// when displaying the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// Usually when a @paintable returns nonzero values from
// [method@Gdk.Paintable.get_intrinsic_width] and
// [method@Gdk.Paintable.get_intrinsic_height] the aspect ratio should
// conform to those values, though that is not required.
//
// If the @paintable does not have a preferred aspect ratio, it returns 0.
// Negative values are never returned.
func (p paintable) IntrinsicAspectRatio() float64 {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))

	var _cret C.double // in

	_cret = C.gdk_paintable_get_intrinsic_aspect_ratio(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// IntrinsicHeight gets the preferred height the @paintable would like to be
// displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred height, it returns 0.
// Negative values are never returned.
func (p paintable) IntrinsicHeight() int {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))

	var _cret C.int // in

	_cret = C.gdk_paintable_get_intrinsic_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// IntrinsicWidth gets the preferred width the @paintable would like to be
// displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred width, it returns 0. Negative
// values are never returned.
func (p paintable) IntrinsicWidth() int {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))

	var _cret C.int // in

	_cret = C.gdk_paintable_get_intrinsic_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// InvalidateContents: called by implementations of `GdkPaintable` to
// invalidate their contents.
//
// Unless the contents are invalidated, implementations must guarantee that
// multiple calls of [method@Gdk.Paintable.snapshot] produce the same
// output.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-contents]
// signal.
//
// If a @paintable reports the GDK_PAINTABLE_STATIC_CONTENTS flag, it must
// not call this function.
func (p paintable) InvalidateContents() {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))

	C.gdk_paintable_invalidate_contents(_arg0)
}

// InvalidateSize: called by implementations of `GdkPaintable` to invalidate
// their size.
//
// As long as the size is not invalidated, @paintable must return the same
// values for its intrinsic width, height and aspect ratio.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-size]
// signal.
//
// If a @paintable reports the GDK_PAINTABLE_STATIC_SIZE flag, it must not
// call this function.
func (p paintable) InvalidateSize() {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))

	C.gdk_paintable_invalidate_size(_arg0)
}

// Snapshot snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot. If
// @width and @height are not larger than zero, this function will do
// nothing.
func (p paintable) Snapshot(snapshot Snapshot, width float64, height float64) {
	var _arg0 *C.GdkPaintable // out
	var _arg1 *C.GdkSnapshot  // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GdkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg2 = C.double(width)
	_arg3 = C.double(height)

	C.gdk_paintable_snapshot(_arg0, _arg1, _arg2, _arg3)
}
