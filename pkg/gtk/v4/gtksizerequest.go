// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// DistributeNaturalAllocation distributes extra_space to child sizes by
// bringing smaller children up to natural size first.
//
// The remaining space will be added to the minimum_size member of the
// GtkRequestedSize struct. If all sizes reach their natural size then the
// remaining space is returned.
func DistributeNaturalAllocation(extraSpace int, sizes []RequestedSize) int {
	var _arg1 C.int // out
	var _arg3 *C.GtkRequestedSize
	var _arg2 C.guint
	var _cret C.int // in

	_arg1 = C.int(extraSpace)
	_arg2 = (C.guint)(len(sizes))
	if len(sizes) > 0 {
		_arg3 = (*C.GtkRequestedSize)(unsafe.Pointer(&sizes[0]))
	}

	_cret = C.gtk_distribute_natural_allocation(_arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
type RequestedSize struct {
	native C.GtkRequestedSize
}

// Native returns the underlying C source pointer.
func (r *RequestedSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Data: client pointer
func (r *RequestedSize) Data() cgo.Handle {
	var v cgo.Handle // out
	v = (cgo.Handle)(unsafe.Pointer(r.native.data))
	return v
}

// MinimumSize: minimum size needed for allocation in a given orientation
func (r *RequestedSize) MinimumSize() int {
	var v int // out
	v = int(r.native.minimum_size)
	return v
}

// NaturalSize: natural size for allocation in a given orientation
func (r *RequestedSize) NaturalSize() int {
	var v int // out
	v = int(r.native.natural_size)
	return v
}
