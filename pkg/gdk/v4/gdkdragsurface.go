// Code generated by girgen. DO NOT EDIT.

package gdk

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

// glib.Type values for gdkdragsurface.go.
var GTypeDragSurface = coreglib.Type(C.gdk_drag_surface_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDragSurface, F: marshalDragSurface},
	})
}

// DragSurface is an interface for surfaces used during DND.
//
// DragSurface wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DragSurface struct {
	_ [0]func() // equal guard
	Surface
}

var (
	_ Surfacer = (*DragSurface)(nil)
)

// DragSurfacer describes DragSurface's interface methods.
type DragSurfacer interface {
	coreglib.Objector

	// Present drag_surface.
	Present(width, height int32) bool
}

var _ DragSurfacer = (*DragSurface)(nil)

func wrapDragSurface(obj *coreglib.Object) *DragSurface {
	return &DragSurface{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalDragSurface(p uintptr) (interface{}, error) {
	return wrapDragSurface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Present drag_surface.
//
// The function takes the following parameters:
//
//    - width: unconstrained drag_surface width to layout.
//    - height: unconstrained drag_surface height to layout.
//
// The function returns the following values:
//
//    - ok: FALSE if it failed to be presented, otherwise TRUE.
//
func (dragSurface *DragSurface) Present(width, height int32) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragSurface).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(height)

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragSurface)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
