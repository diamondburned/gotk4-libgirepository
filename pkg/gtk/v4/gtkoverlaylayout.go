// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkoverlaylayout.go.
var (
	GTypeOverlayLayout      = coreglib.Type(C.gtk_overlay_layout_get_type())
	GTypeOverlayLayoutChild = coreglib.Type(C.gtk_overlay_layout_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeOverlayLayout, F: marshalOverlayLayout},
		{T: GTypeOverlayLayoutChild, F: marshalOverlayLayoutChild},
	})
}

// OverlayLayoutOverrider contains methods that are overridable.
type OverlayLayoutOverrider interface {
}

// OverlayLayout: GtkOverlayLayout is the layout manager used by GtkOverlay.
//
// It places widgets as overlays on top of the main child.
//
// This is not a reusable layout manager, since it expects its widget to be a
// GtkOverlay. It only listed here so that its layout properties get documented.
type OverlayLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*OverlayLayout)(nil)
)

func classInitOverlayLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapOverlayLayout(obj *coreglib.Object) *OverlayLayout {
	return &OverlayLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalOverlayLayout(p uintptr) (interface{}, error) {
	return wrapOverlayLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewOverlayLayout creates a new GtkOverlayLayout instance.
//
// The function returns the following values:
//
//    - overlayLayout: newly created instance.
//
func NewOverlayLayout() *OverlayLayout {
	_gret := girepository.MustFind("Gtk", "OverlayLayout").InvokeMethod("new_OverlayLayout", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _overlayLayout *OverlayLayout // out

	_overlayLayout = wrapOverlayLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _overlayLayout
}

// OverlayLayoutChildOverrider contains methods that are overridable.
type OverlayLayoutChildOverrider interface {
}

// OverlayLayoutChild: GtkLayoutChild subclass for children in a
// GtkOverlayLayout.
type OverlayLayoutChild struct {
	_ [0]func() // equal guard
	LayoutChild
}

var (
	_ LayoutChilder = (*OverlayLayoutChild)(nil)
)

func classInitOverlayLayoutChilder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapOverlayLayoutChild(obj *coreglib.Object) *OverlayLayoutChild {
	return &OverlayLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalOverlayLayoutChild(p uintptr) (interface{}, error) {
	return wrapOverlayLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClipOverlay retrieves whether the child is clipped.
//
// The function returns the following values:
//
//    - ok: whether the child is clipped.
//
func (child *OverlayLayoutChild) ClipOverlay() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_gret := girepository.MustFind("Gtk", "OverlayLayoutChild").InvokeMethod("get_clip_overlay", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Measure retrieves whether the child is measured.
//
// The function returns the following values:
//
//    - ok: whether the child is measured.
//
func (child *OverlayLayoutChild) Measure() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_gret := girepository.MustFind("Gtk", "OverlayLayoutChild").InvokeMethod("get_measure", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetClipOverlay sets whether to clip this child.
//
// The function takes the following parameters:
//
//    - clipOverlay: whether to clip this child.
//
func (child *OverlayLayoutChild) SetClipOverlay(clipOverlay bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if clipOverlay {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "OverlayLayoutChild").InvokeMethod("set_clip_overlay", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(clipOverlay)
}

// SetMeasure sets whether to measure this child.
//
// The function takes the following parameters:
//
//    - measure: whether to measure this child.
//
func (child *OverlayLayoutChild) SetMeasure(measure bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if measure {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "OverlayLayoutChild").InvokeMethod("set_measure", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(measure)
}
