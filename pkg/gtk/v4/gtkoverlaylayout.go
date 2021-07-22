// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_overlay_layout_get_type()), F: marshalOverlayLayouter},
		{T: externglib.Type(C.gtk_overlay_layout_child_get_type()), F: marshalOverlayLayoutChilder},
	})
}

// OverlayLayout: GtkOverlayLayout is the layout manager used by GtkOverlay.
//
// It places widgets as overlays on top of the main child.
//
// This is not a reusable layout manager, since it expects its widget to be a
// GtkOverlay. It only listed here so that its layout properties get documented.
type OverlayLayout struct {
	LayoutManager
}

func wrapOverlayLayout(obj *externglib.Object) *OverlayLayout {
	return &OverlayLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalOverlayLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOverlayLayout(obj), nil
}

// NewOverlayLayout creates a new GtkOverlayLayout instance.
func NewOverlayLayout() *OverlayLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_overlay_layout_new()

	var _overlayLayout *OverlayLayout // out

	_overlayLayout = wrapOverlayLayout(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _overlayLayout
}

func (*OverlayLayout) privateOverlayLayout() {}

// OverlayLayoutChild: GtkLayoutChild subclass for children in a
// GtkOverlayLayout.
type OverlayLayoutChild struct {
	LayoutChild
}

func wrapOverlayLayoutChild(obj *externglib.Object) *OverlayLayoutChild {
	return &OverlayLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalOverlayLayoutChilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOverlayLayoutChild(obj), nil
}

// ClipOverlay retrieves whether the child is clipped.
func (child *OverlayLayoutChild) ClipOverlay() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_overlay_layout_child_get_clip_overlay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Measure retrieves whether the child is measured.
func (child *OverlayLayoutChild) Measure() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_overlay_layout_child_get_measure(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetClipOverlay sets whether to clip this child.
func (child *OverlayLayoutChild) SetClipOverlay(clipOverlay bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(child.Native()))
	if clipOverlay {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_clip_overlay(_arg0, _arg1)
}

// SetMeasure sets whether to measure this child.
func (child *OverlayLayoutChild) SetMeasure(measure bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(child.Native()))
	if measure {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_measure(_arg0, _arg1)
}
