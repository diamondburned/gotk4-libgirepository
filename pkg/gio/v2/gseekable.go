// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSeekable = coreglib.Type(girepository.MustFind("Gio", "Seekable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeekable, F: marshalSeekable},
	})
}

// SeekableOverrider contains methods that are overridable.
type SeekableOverrider interface {
}

// Seekable is implemented by streams (implementations of Stream or Stream) that
// support seeking.
//
// Seekable streams largely fall into two categories: resizable and fixed-size.
//
// #GSeekable on fixed-sized streams is approximately the same as POSIX lseek()
// on a block device (for example: attempting to seek past the end of the device
// is an error). Fixed streams typically cannot be truncated.
//
// #GSeekable on resizable streams is approximately the same as POSIX lseek() on
// a normal file. Seeking past the end and writing data will usually cause the
// stream to resize by introducing zero bytes.
//
// Seekable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Seekable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Seekable)(nil)
)

// Seekabler describes Seekable's interface methods.
type Seekabler interface {
	coreglib.Objector

	baseSeekable() *Seekable
}

var _ Seekabler = (*Seekable)(nil)

func ifaceInitSeekabler(gifacePtr, data C.gpointer) {
}

func wrapSeekable(obj *coreglib.Object) *Seekable {
	return &Seekable{
		Object: obj,
	}
}

func marshalSeekable(p uintptr) (interface{}, error) {
	return wrapSeekable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Seekable) baseSeekable() *Seekable {
	return v
}

// BaseSeekable returns the underlying base object.
func BaseSeekable(obj Seekabler) *Seekable {
	return obj.baseSeekable()
}

// SeekableIface provides an interface for implementing seekable functionality
// on I/O Streams.
//
// An instance of this type is always passed by reference.
type SeekableIface struct {
	*seekableIface
}

// seekableIface is the struct that's finalized.
type seekableIface struct {
	native unsafe.Pointer
}

var GIRInfoSeekableIface = girepository.MustFind("Gio", "SeekableIface")
