// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_paintable_flags_get_type()), F: marshalPaintableFlags},
		{T: externglib.Type(C.gdk_paintable_get_type()), F: marshalPaintabler},
	})
}

// PaintableFlags flags about a paintable object.
//
// Implementations use these for optimizations such as caching.
type PaintableFlags C.guint

const (
	// PaintableStaticSize: size is immutable. The gdkpaintable::invalidate-size
	// signal will never be emitted.
	PaintableStaticSize PaintableFlags = 0b1
	// PaintableStaticContents: content is immutable. The
	// gdkpaintable::invalidate-contents signal will never be emitted.
	PaintableStaticContents PaintableFlags = 0b10
)

func marshalPaintableFlags(p uintptr) (interface{}, error) {
	return PaintableFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for PaintableFlags.
func (p PaintableFlags) String() string {
	if p == 0 {
		return "PaintableFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(43)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PaintableStaticSize:
			builder.WriteString("Size|")
		case PaintableStaticContents:
			builder.WriteString("Contents|")
		default:
			builder.WriteString(fmt.Sprintf("PaintableFlags(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PaintableFlags) Has(other PaintableFlags) bool {
	return (p & other) == other
}

// PaintableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PaintableOverrider interface {
	// CurrentImage gets an immutable paintable for the current contents
	// displayed by paintable.
	//
	// This is useful when you want to retain the current state of an animation,
	// for example to take a screenshot of a running animation.
	//
	// If the paintable is already immutable, it will return itself.
	CurrentImage() Paintabler
	// Flags: get flags for the paintable.
	//
	// This is oftentimes useful for optimizations.
	//
	// See gdk.PaintableFlags for the flags and what they mean.
	Flags() PaintableFlags
	// IntrinsicAspectRatio gets the preferred aspect ratio the paintable would
	// like to be displayed at.
	//
	// The aspect ratio is the width divided by the height, so a value of 0.5
	// means that the paintable prefers to be displayed twice as high as it is
	// wide. Consumers of this interface can use this to preserve aspect ratio
	// when displaying the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to gdk.Paintable.Snapshot().
	//
	// Usually when a paintable returns nonzero values from
	// gdk.Paintable.GetIntrinsicWidth() and gdk.Paintable.GetIntrinsicHeight()
	// the aspect ratio should conform to those values, though that is not
	// required.
	//
	// If the paintable does not have a preferred aspect ratio, it returns 0.
	// Negative values are never returned.
	IntrinsicAspectRatio() float64
	// IntrinsicHeight gets the preferred height the paintable would like to be
	// displayed at.
	//
	// Consumers of this interface can use this to reserve enough space to draw
	// the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to gdk.Paintable.Snapshot().
	//
	// If the paintable does not have a preferred height, it returns 0. Negative
	// values are never returned.
	IntrinsicHeight() int
	// IntrinsicWidth gets the preferred width the paintable would like to be
	// displayed at.
	//
	// Consumers of this interface can use this to reserve enough space to draw
	// the paintable.
	//
	// This is a purely informational value and does not in any way limit the
	// values that may be passed to gdk.Paintable.Snapshot().
	//
	// If the paintable does not have a preferred width, it returns 0. Negative
	// values are never returned.
	IntrinsicWidth() int
	// Snapshot snapshots the given paintable with the given width and height.
	//
	// The paintable is drawn at the current (0,0) offset of the snapshot. If
	// width and height are not larger than zero, this function will do nothing.
	Snapshot(snapshot Snapshotter, width, height float64)
}

// Paintable: GdkPaintable is a simple interface used by GTK to represent
// content that can be painted.
//
// The content of a GdkPaintable can be painted anywhere at any size without
// requiring any sort of layout. The interface is inspired by similar concepts
// elsewhere, such as ClutterContent
// (https://developer.gnome.org/clutter/stable/ClutterContent.html), HTML/CSS
// Paint Sources (https://www.w3.org/TR/css-images-4/#paint-source), or SVG
// Paint Servers (https://www.w3.org/TR/SVG2/pservers.html).
//
// A GdkPaintable can be snapshot at any time and size using
// gdk.Paintable.Snapshot(). How the paintable interprets that size and if it
// scales or centers itself into the given rectangle is implementation defined,
// though if you are implementing a GdkPaintable and don't know what to do, it
// is suggested that you scale your paintable ignoring any potential aspect
// ratio.
//
// The contents that a GdkPaintable produces may depend on the gdksnapshot
// passed to it. For example, paintables may decide to use more detailed images
// on higher resolution screens or when OpenGL is available. A GdkPaintable will
// however always produce the same output for the same snapshot.
//
// A GdkPaintable may change its contents, meaning that it will now produce a
// different output with the same snapshot. Once that happens, it will call
// gdk.Paintable.InvalidateContents() which will emit the
// gdkpaintable::invalidate-contents signal. If a paintable is known to never
// change its contents, it will set the GDK_PAINTABLE_STATIC_CONTENTS flag. If a
// consumer cannot deal with changing contents, it may call
// gdk.Paintable.GetCurrentImage() which will return a static paintable and use
// that.
//
// A paintable can report an intrinsic (or preferred) size or aspect ratio it
// wishes to be rendered at, though it doesn't have to. Consumers of the
// interface can use this information to layout thepaintable appropriately. Just
// like the contents, the size of a paintable can change. A paintable will
// indicate this by calling gdk.Paintable.InvalidateSize() which will emit the
// gdkpaintable::invalidate-size signal. And just like for contents, if a
// paintable is known to never change its size, it will set the
// GDK_PAINTABLE_STATIC_SIZE flag.
//
// Besides API for applications, there are some functions that are only useful
// for implementing subclasses and should not be used by applications:
// gdk.Paintable.InvalidateContents(), gdk.Paintable.InvalidateSize(),
// gdk.Paintable().NewEmpty.
type Paintable struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Paintable)(nil)
)

// Paintabler describes Paintable's interface methods.
type Paintabler interface {
	externglib.Objector

	// ComputeConcreteSize: compute a concrete size for the GdkPaintable.
	ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight float64) (concreteWidth float64, concreteHeight float64)
	// CurrentImage gets an immutable paintable for the current contents
	// displayed by paintable.
	CurrentImage() Paintabler
	// Flags: get flags for the paintable.
	Flags() PaintableFlags
	// IntrinsicAspectRatio gets the preferred aspect ratio the paintable would
	// like to be displayed at.
	IntrinsicAspectRatio() float64
	// IntrinsicHeight gets the preferred height the paintable would like to be
	// displayed at.
	IntrinsicHeight() int
	// IntrinsicWidth gets the preferred width the paintable would like to be
	// displayed at.
	IntrinsicWidth() int
	// InvalidateContents: called by implementations of GdkPaintable to
	// invalidate their contents.
	InvalidateContents()
	// InvalidateSize: called by implementations of GdkPaintable to invalidate
	// their size.
	InvalidateSize()
	// Snapshot snapshots the given paintable with the given width and height.
	Snapshot(snapshot Snapshotter, width, height float64)
}

var _ Paintabler = (*Paintable)(nil)

func wrapPaintable(obj *externglib.Object) *Paintable {
	return &Paintable{
		Object: obj,
	}
}

func marshalPaintabler(p uintptr) (interface{}, error) {
	return wrapPaintable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ComputeConcreteSize: compute a concrete size for the GdkPaintable.
//
// Applies the sizing algorithm outlined in the CSS Image spec
// (https://drafts.csswg.org/css-images-3/#default-sizing) to the given
// paintable. See that link for more details.
//
// It is not necessary to call this function when both specified_width and
// specified_height are known, but it is useful to call this function in
// GtkWidget:measure implementations to compute the other dimension when only
// one dimension is given.
//
// The function takes the following parameters:
//
//    - specifiedWidth: width paintable could be drawn into or 0.0 if unknown.
//    - specifiedHeight: height paintable could be drawn into or 0.0 if
//    unknown.
//    - defaultWidth: width paintable would be drawn into if no other
//    constraints were given.
//    - defaultHeight: height paintable would be drawn into if no other
//    constraints were given.
//
func (paintable *Paintable) ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight float64) (concreteWidth float64, concreteHeight float64) {
	var _arg0 *C.GdkPaintable // out
	var _arg1 C.double        // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out
	var _arg5 C.double        // in
	var _arg6 C.double        // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))
	_arg1 = C.double(specifiedWidth)
	_arg2 = C.double(specifiedHeight)
	_arg3 = C.double(defaultWidth)
	_arg4 = C.double(defaultHeight)

	C.gdk_paintable_compute_concrete_size(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)
	runtime.KeepAlive(paintable)
	runtime.KeepAlive(specifiedWidth)
	runtime.KeepAlive(specifiedHeight)
	runtime.KeepAlive(defaultWidth)
	runtime.KeepAlive(defaultHeight)

	var _concreteWidth float64  // out
	var _concreteHeight float64 // out

	_concreteWidth = float64(_arg5)
	_concreteHeight = float64(_arg6)

	return _concreteWidth, _concreteHeight
}

// CurrentImage gets an immutable paintable for the current contents displayed
// by paintable.
//
// This is useful when you want to retain the current state of an animation, for
// example to take a screenshot of a running animation.
//
// If the paintable is already immutable, it will return itself.
func (paintable *Paintable) CurrentImage() Paintabler {
	var _arg0 *C.GdkPaintable // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	_cret = C.gdk_paintable_get_current_image(_arg0)
	runtime.KeepAlive(paintable)

	var _ret Paintabler // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Paintabler is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Paintabler)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Paintabler")
		}
		_ret = rv
	}

	return _ret
}

// Flags: get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See gdk.PaintableFlags for the flags and what they mean.
func (paintable *Paintable) Flags() PaintableFlags {
	var _arg0 *C.GdkPaintable     // out
	var _cret C.GdkPaintableFlags // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	_cret = C.gdk_paintable_get_flags(_arg0)
	runtime.KeepAlive(paintable)

	var _paintableFlags PaintableFlags // out

	_paintableFlags = PaintableFlags(_cret)

	return _paintableFlags
}

// IntrinsicAspectRatio gets the preferred aspect ratio the paintable would like
// to be displayed at.
//
// The aspect ratio is the width divided by the height, so a value of 0.5 means
// that the paintable prefers to be displayed twice as high as it is wide.
// Consumers of this interface can use this to preserve aspect ratio when
// displaying the paintable.
//
// This is a purely informational value and does not in any way limit the values
// that may be passed to gdk.Paintable.Snapshot().
//
// Usually when a paintable returns nonzero values from
// gdk.Paintable.GetIntrinsicWidth() and gdk.Paintable.GetIntrinsicHeight() the
// aspect ratio should conform to those values, though that is not required.
//
// If the paintable does not have a preferred aspect ratio, it returns 0.
// Negative values are never returned.
func (paintable *Paintable) IntrinsicAspectRatio() float64 {
	var _arg0 *C.GdkPaintable // out
	var _cret C.double        // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	_cret = C.gdk_paintable_get_intrinsic_aspect_ratio(_arg0)
	runtime.KeepAlive(paintable)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// IntrinsicHeight gets the preferred height the paintable would like to be
// displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw the
// paintable.
//
// This is a purely informational value and does not in any way limit the values
// that may be passed to gdk.Paintable.Snapshot().
//
// If the paintable does not have a preferred height, it returns 0. Negative
// values are never returned.
func (paintable *Paintable) IntrinsicHeight() int {
	var _arg0 *C.GdkPaintable // out
	var _cret C.int           // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	_cret = C.gdk_paintable_get_intrinsic_height(_arg0)
	runtime.KeepAlive(paintable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IntrinsicWidth gets the preferred width the paintable would like to be
// displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw the
// paintable.
//
// This is a purely informational value and does not in any way limit the values
// that may be passed to gdk.Paintable.Snapshot().
//
// If the paintable does not have a preferred width, it returns 0. Negative
// values are never returned.
func (paintable *Paintable) IntrinsicWidth() int {
	var _arg0 *C.GdkPaintable // out
	var _cret C.int           // in

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	_cret = C.gdk_paintable_get_intrinsic_width(_arg0)
	runtime.KeepAlive(paintable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InvalidateContents: called by implementations of GdkPaintable to invalidate
// their contents.
//
// Unless the contents are invalidated, implementations must guarantee that
// multiple calls of gdk.Paintable.Snapshot() produce the same output.
//
// This function will emit the gdk.Paintable::invalidate-contents signal.
//
// If a paintable reports the GDK_PAINTABLE_STATIC_CONTENTS flag, it must not
// call this function.
func (paintable *Paintable) InvalidateContents() {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	C.gdk_paintable_invalidate_contents(_arg0)
	runtime.KeepAlive(paintable)
}

// InvalidateSize: called by implementations of GdkPaintable to invalidate their
// size.
//
// As long as the size is not invalidated, paintable must return the same values
// for its intrinsic width, height and aspect ratio.
//
// This function will emit the gdk.Paintable::invalidate-size signal.
//
// If a paintable reports the GDK_PAINTABLE_STATIC_SIZE flag, it must not call
// this function.
func (paintable *Paintable) InvalidateSize() {
	var _arg0 *C.GdkPaintable // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	C.gdk_paintable_invalidate_size(_arg0)
	runtime.KeepAlive(paintable)
}

// Snapshot snapshots the given paintable with the given width and height.
//
// The paintable is drawn at the current (0,0) offset of the snapshot. If width
// and height are not larger than zero, this function will do nothing.
//
// The function takes the following parameters:
//
//    - snapshot: GdkSnapshot to snapshot to.
//    - width to snapshot in.
//    - height to snapshot in.
//
func (paintable *Paintable) Snapshot(snapshot Snapshotter, width, height float64) {
	var _arg0 *C.GdkPaintable // out
	var _arg1 *C.GdkSnapshot  // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out

	_arg0 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))
	_arg1 = (*C.GdkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg2 = C.double(width)
	_arg3 = C.double(height)

	C.gdk_paintable_snapshot(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(paintable)
	runtime.KeepAlive(snapshot)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// ConnectInvalidateContents: emitted when the contents of the paintable change.
//
// Examples for such an event would be videos changing to the next frame or the
// icon theme for an icon changing.
func (paintable *Paintable) ConnectInvalidateContents(f func()) externglib.SignalHandle {
	return paintable.Connect("invalidate-contents", f)
}

// ConnectInvalidateSize: emitted when the intrinsic size of the paintable
// changes.
//
// This means the values reported by at least one of
// gdk.Paintable.GetIntrinsicWidth(), gdk.Paintable.GetIntrinsicHeight() or
// gdk.Paintable.GetIntrinsicAspectRatio() has changed.
//
// Examples for such an event would be a paintable displaying the contents of a
// toplevel surface being resized.
func (paintable *Paintable) ConnectInvalidateSize(f func()) externglib.SignalHandle {
	return paintable.Connect("invalidate-size", f)
}

// NewPaintableEmpty returns a paintable that has the given intrinsic size and
// draws nothing.
//
// This is often useful for implementing the
// PaintableInterface.get_current_image() virtual function when the paintable is
// in an incomplete state (like a gtk.MediaStream before receiving the first
// frame).
//
// The function takes the following parameters:
//
//    - intrinsicWidth: intrinsic width to report. Can be 0 for no width.
//    - intrinsicHeight: intrinsic height to report. Can be 0 for no height.
//
func NewPaintableEmpty(intrinsicWidth, intrinsicHeight int) Paintabler {
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret *C.GdkPaintable // in

	_arg1 = C.int(intrinsicWidth)
	_arg2 = C.int(intrinsicHeight)

	_cret = C.gdk_paintable_new_empty(_arg1, _arg2)
	runtime.KeepAlive(intrinsicWidth)
	runtime.KeepAlive(intrinsicHeight)

	var _paintable Paintabler // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Paintabler is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Paintabler)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Paintabler")
		}
		_paintable = rv
	}

	return _paintable
}
