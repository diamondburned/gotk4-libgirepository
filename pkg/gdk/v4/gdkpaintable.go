// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GdkPaintable* _gotk4_gdk4_PaintableInterface_get_current_image(void*);
// extern double _gotk4_gdk4_PaintableInterface_get_intrinsic_aspect_ratio(void*);
// extern int _gotk4_gdk4_PaintableInterface_get_intrinsic_height(void*);
// extern int _gotk4_gdk4_PaintableInterface_get_intrinsic_width(void*);
// extern void _gotk4_gdk4_PaintableInterface_snapshot(void*, void*, double, double);
// extern void _gotk4_gdk4_Paintable_ConnectInvalidateContents(gpointer, guintptr);
// extern void _gotk4_gdk4_Paintable_ConnectInvalidateSize(gpointer, guintptr);
import "C"

// glib.Type values for gdkpaintable.go.
var (
	GTypePaintableFlags = coreglib.Type(C.gdk_paintable_flags_get_type())
	GTypePaintable      = coreglib.Type(C.gdk_paintable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePaintableFlags, F: marshalPaintableFlags},
		{T: GTypePaintable, F: marshalPaintable},
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
	return PaintableFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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
type PaintableOverrider interface {
	// CurrentImage gets an immutable paintable for the current contents
	// displayed by paintable.
	//
	// This is useful when you want to retain the current state of an animation,
	// for example to take a screenshot of a running animation.
	//
	// If the paintable is already immutable, it will return itself.
	//
	// The function returns the following values:
	//
	//    - ret: immutable paintable for the current contents of paintable.
	//
	CurrentImage() *Paintable
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
	//
	// The function returns the following values:
	//
	//    - gdouble: intrinsic aspect ratio of paintable or 0 if none.
	//
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
	//
	// The function returns the following values:
	//
	//    - gint: intrinsic height of paintable or 0 if none.
	//
	IntrinsicHeight() int32
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
	//
	// The function returns the following values:
	//
	//    - gint: intrinsic width of paintable or 0 if none.
	//
	IntrinsicWidth() int32
	// Snapshot snapshots the given paintable with the given width and height.
	//
	// The paintable is drawn at the current (0,0) offset of the snapshot. If
	// width and height are not larger than zero, this function will do nothing.
	//
	// The function takes the following parameters:
	//
	//    - snapshot: GdkSnapshot to snapshot to.
	//    - width to snapshot in.
	//    - height to snapshot in.
	//
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
//
// Paintable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Paintable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Paintable)(nil)
)

// Paintabler describes Paintable's interface methods.
type Paintabler interface {
	coreglib.Objector

	// ComputeConcreteSize: compute a concrete size for the GdkPaintable.
	ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight float64) (concreteWidth, concreteHeight float64)
	// CurrentImage gets an immutable paintable for the current contents
	// displayed by paintable.
	CurrentImage() *Paintable
	// IntrinsicAspectRatio gets the preferred aspect ratio the paintable would
	// like to be displayed at.
	IntrinsicAspectRatio() float64
	// IntrinsicHeight gets the preferred height the paintable would like to be
	// displayed at.
	IntrinsicHeight() int32
	// IntrinsicWidth gets the preferred width the paintable would like to be
	// displayed at.
	IntrinsicWidth() int32
	// InvalidateContents: called by implementations of GdkPaintable to
	// invalidate their contents.
	InvalidateContents()
	// InvalidateSize: called by implementations of GdkPaintable to invalidate
	// their size.
	InvalidateSize()
	// Snapshot snapshots the given paintable with the given width and height.
	Snapshot(snapshot Snapshotter, width, height float64)

	// Invalidate-contents is emitted when the contents of the paintable change.
	ConnectInvalidateContents(func()) coreglib.SignalHandle
	// Invalidate-size is emitted when the intrinsic size of the paintable
	// changes.
	ConnectInvalidateSize(func()) coreglib.SignalHandle
}

var _ Paintabler = (*Paintable)(nil)

func ifaceInitPaintabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gdk", "PaintableInterface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_current_image"))) = unsafe.Pointer(C._gotk4_gdk4_PaintableInterface_get_current_image)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_intrinsic_aspect_ratio"))) = unsafe.Pointer(C._gotk4_gdk4_PaintableInterface_get_intrinsic_aspect_ratio)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_intrinsic_height"))) = unsafe.Pointer(C._gotk4_gdk4_PaintableInterface_get_intrinsic_height)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_intrinsic_width"))) = unsafe.Pointer(C._gotk4_gdk4_PaintableInterface_get_intrinsic_width)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("snapshot"))) = unsafe.Pointer(C._gotk4_gdk4_PaintableInterface_snapshot)
}

//export _gotk4_gdk4_PaintableInterface_get_current_image
func _gotk4_gdk4_PaintableInterface_get_current_image(arg0 *C.void) (cret *C.GdkPaintable) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PaintableOverrider)

	ret := iface.CurrentImage()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(ret).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(ret).Native()))

	return cret
}

//export _gotk4_gdk4_PaintableInterface_get_intrinsic_aspect_ratio
func _gotk4_gdk4_PaintableInterface_get_intrinsic_aspect_ratio(arg0 *C.void) (cret C.double) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PaintableOverrider)

	gdouble := iface.IntrinsicAspectRatio()

	cret = C.double(gdouble)

	return cret
}

//export _gotk4_gdk4_PaintableInterface_get_intrinsic_height
func _gotk4_gdk4_PaintableInterface_get_intrinsic_height(arg0 *C.void) (cret C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PaintableOverrider)

	gint := iface.IntrinsicHeight()

	cret = C.int(gint)

	return cret
}

//export _gotk4_gdk4_PaintableInterface_get_intrinsic_width
func _gotk4_gdk4_PaintableInterface_get_intrinsic_width(arg0 *C.void) (cret C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PaintableOverrider)

	gint := iface.IntrinsicWidth()

	cret = C.int(gint)

	return cret
}

//export _gotk4_gdk4_PaintableInterface_snapshot
func _gotk4_gdk4_PaintableInterface_snapshot(arg0 *C.void, arg1 *C.void, arg2 C.double, arg3 C.double) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PaintableOverrider)

	var _snapshot Snapshotter // out
	var _width float64        // out
	var _height float64       // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Snapshotter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Snapshotter)
			return ok
		})
		rv, ok := casted.(Snapshotter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Snapshotter")
		}
		_snapshot = rv
	}
	_width = float64(arg2)
	_height = float64(arg3)

	iface.Snapshot(_snapshot, _width, _height)
}

func wrapPaintable(obj *coreglib.Object) *Paintable {
	return &Paintable{
		Object: obj,
	}
}

func marshalPaintable(p uintptr) (interface{}, error) {
	return wrapPaintable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gdk4_Paintable_ConnectInvalidateContents
func _gotk4_gdk4_Paintable_ConnectInvalidateContents(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectInvalidateContents is emitted when the contents of the paintable
// change.
//
// Examples for such an event would be videos changing to the next frame or the
// icon theme for an icon changing.
func (paintable *Paintable) ConnectInvalidateContents(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paintable, "invalidate-contents", false, unsafe.Pointer(C._gotk4_gdk4_Paintable_ConnectInvalidateContents), f)
}

//export _gotk4_gdk4_Paintable_ConnectInvalidateSize
func _gotk4_gdk4_Paintable_ConnectInvalidateSize(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectInvalidateSize is emitted when the intrinsic size of the paintable
// changes.
//
// This means the values reported by at least one of
// gdk.Paintable.GetIntrinsicWidth(), gdk.Paintable.GetIntrinsicHeight() or
// gdk.Paintable.GetIntrinsicAspectRatio() has changed.
//
// Examples for such an event would be a paintable displaying the contents of a
// toplevel surface being resized.
func (paintable *Paintable) ConnectInvalidateSize(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paintable, "invalidate-size", false, unsafe.Pointer(C._gotk4_gdk4_Paintable_ConnectInvalidateSize), f)
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
//    - specifiedHeight: height paintable could be drawn into or 0.0 if unknown.
//    - defaultWidth: width paintable would be drawn into if no other constraints
//      were given.
//    - defaultHeight: height paintable would be drawn into if no other
//      constraints were given.
//
// The function returns the following values:
//
//    - concreteWidth will be set to the concrete width computed.
//    - concreteHeight will be set to the concrete height computed.
//
func (paintable *Paintable) ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight float64) (concreteWidth, concreteHeight float64) {
	var _args [5]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(specifiedWidth)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(specifiedHeight)
	*(*C.double)(unsafe.Pointer(&_args[3])) = C.double(defaultWidth)
	*(*C.double)(unsafe.Pointer(&_args[4])) = C.double(defaultHeight)

	runtime.KeepAlive(paintable)
	runtime.KeepAlive(specifiedWidth)
	runtime.KeepAlive(specifiedHeight)
	runtime.KeepAlive(defaultWidth)
	runtime.KeepAlive(defaultHeight)

	var _concreteWidth float64  // out
	var _concreteHeight float64 // out

	_concreteWidth = *(*float64)(unsafe.Pointer(_outs[0]))
	_concreteHeight = *(*float64)(unsafe.Pointer(_outs[1]))

	return _concreteWidth, _concreteHeight
}

// CurrentImage gets an immutable paintable for the current contents displayed
// by paintable.
//
// This is useful when you want to retain the current state of an animation, for
// example to take a screenshot of a running animation.
//
// If the paintable is already immutable, it will return itself.
//
// The function returns the following values:
//
//    - ret: immutable paintable for the current contents of paintable.
//
func (paintable *Paintable) CurrentImage() *Paintable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _ret *Paintable // out

	_ret = wrapPaintable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
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
//
// The function returns the following values:
//
//    - gdouble: intrinsic aspect ratio of paintable or 0 if none.
//
func (paintable *Paintable) IntrinsicAspectRatio() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

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
//
// The function returns the following values:
//
//    - gint: intrinsic height of paintable or 0 if none.
//
func (paintable *Paintable) IntrinsicHeight() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
//
// The function returns the following values:
//
//    - gint: intrinsic width of paintable or 0 if none.
//
func (paintable *Paintable) IntrinsicWidth() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(snapshot).Native()))
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(width)
	*(*C.double)(unsafe.Pointer(&_args[3])) = C.double(height)

	runtime.KeepAlive(paintable)
	runtime.KeepAlive(snapshot)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
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
// The function returns the following values:
//
//    - paintable: GdkPaintable.
//
func NewPaintableEmpty(intrinsicWidth, intrinsicHeight int32) *Paintable {
	var _args [2]girepository.Argument

	*(*C.int)(unsafe.Pointer(&_args[0])) = C.int(intrinsicWidth)
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(intrinsicHeight)

	_gret := girepository.MustFind("Gdk", "new_empty").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(intrinsicWidth)
	runtime.KeepAlive(intrinsicHeight)

	var _paintable *Paintable // out

	_paintable = wrapPaintable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _paintable
}
