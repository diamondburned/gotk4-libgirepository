// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// SocketSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_socket_create_source().
type SocketSourceFunc func(socket *Socket, condition glib.IOCondition) (ok bool)

// InputVector: structure used for scatter/gather data input. You generally pass
// in an array of Vectors and the operation will store the read data starting in
// the first buffer, switching to the next as needed.
//
// An instance of this type is always passed by reference.
type InputVector struct {
	*inputVector
}

// inputVector is the struct that's finalized.
type inputVector struct {
	native unsafe.Pointer
}

var GIRInfoInputVector = girepository.MustFind("Gio", "InputVector")

// Buffer: pointer to a buffer where data will be written.
func (i *InputVector) Buffer() unsafe.Pointer {
	offset := GIRInfoInputVector.StructFieldOffset("buffer")
	valptr := (*unsafe.Pointer)(unsafe.Add(i.native, offset))
	var _v unsafe.Pointer // out
	_v = (unsafe.Pointer)(unsafe.Pointer(*valptr))
	return _v
}

// Size: available size in buffer.
func (i *InputVector) Size() uint {
	offset := GIRInfoInputVector.StructFieldOffset("size")
	valptr := (*uint)(unsafe.Add(i.native, offset))
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Size: available size in buffer.
func (i *InputVector) SetSize(size uint) {
	offset := GIRInfoInputVector.StructFieldOffset("size")
	valptr := (*C.gsize)(unsafe.Add(i.native, offset))
	*valptr = C.gsize(size)
}

// OutputVector: structure used for scatter/gather data output. You generally
// pass in an array of Vectors and the operation will use all the buffers as if
// they were one buffer.
//
// An instance of this type is always passed by reference.
type OutputVector struct {
	*outputVector
}

// outputVector is the struct that's finalized.
type outputVector struct {
	native unsafe.Pointer
}

var GIRInfoOutputVector = girepository.MustFind("Gio", "OutputVector")

// Buffer: pointer to a buffer of data to read.
func (o *OutputVector) Buffer() unsafe.Pointer {
	offset := GIRInfoOutputVector.StructFieldOffset("buffer")
	valptr := (*unsafe.Pointer)(unsafe.Add(o.native, offset))
	var _v unsafe.Pointer // out
	_v = (unsafe.Pointer)(unsafe.Pointer(*valptr))
	return _v
}

// Size: size of buffer.
func (o *OutputVector) Size() uint {
	offset := GIRInfoOutputVector.StructFieldOffset("size")
	valptr := (*uint)(unsafe.Add(o.native, offset))
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Size: size of buffer.
func (o *OutputVector) SetSize(size uint) {
	offset := GIRInfoOutputVector.StructFieldOffset("size")
	valptr := (*C.gsize)(unsafe.Add(o.native, offset))
	*valptr = C.gsize(size)
}
