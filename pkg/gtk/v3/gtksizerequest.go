// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// DistributeNaturalAllocation distributes @extra_space to child @sizes by
// bringing smaller children up to natural size first.
//
// The remaining space will be added to the @minimum_size member of the
// GtkRequestedSize struct. If all sizes reach their natural size then the
// remaining space is returned.
func DistributeNaturalAllocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) int {
	var arg1 C.gint
	var arg2 C.guint
	var arg3 *C.GtkRequestedSize

	arg1 = C.gint(extraSpace)
	arg2 = C.guint(nRequestedSizes)
	arg3 = (*C.GtkRequestedSize)(unsafe.Pointer(sizes.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_distribute_natural_allocation(arg1, arg2, arg3)

	goret = int(cret)

	return goret
}

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
type RequestedSize struct {
	native C.GtkRequestedSize
}

// WrapRequestedSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRequestedSize(ptr unsafe.Pointer) *RequestedSize {
	if ptr == nil {
		return nil
	}

	return (*RequestedSize)(ptr)
}

func marshalRequestedSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRequestedSize(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RequestedSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Data gets the field inside the struct.
func (r *RequestedSize) Data() interface{} {
	var v interface{}
	v = interface{}(r.native.data)
	return v
}

// MinimumSize gets the field inside the struct.
func (r *RequestedSize) MinimumSize() int {
	var v int
	v = int(r.native.minimum_size)
	return v
}

// NaturalSize gets the field inside the struct.
func (r *RequestedSize) NaturalSize() int {
	var v int
	v = int(r.native.natural_size)
	return v
}
