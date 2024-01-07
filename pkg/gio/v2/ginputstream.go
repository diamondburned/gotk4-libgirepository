// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeInputStream = coreglib.Type(girepository.MustFind("Gio", "InputStream").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInputStream, F: marshalInputStream},
	})
}

// InputStreamOverrides contains methods that are overridable.
type InputStreamOverrides struct {
}

func defaultInputStreamOverrides(v *InputStream) InputStreamOverrides {
	return InputStreamOverrides{}
}

// InputStream has functions to read from a stream (g_input_stream_read()), to
// close a stream (g_input_stream_close()) and to skip some content
// (g_input_stream_skip()).
//
// To copy the content of an input stream to an output stream without manually
// handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for OStream for details of thread safety of streaming
// APIs.
//
// All of these functions have async variants too.
type InputStream struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*InputStream)(nil)
)

// InputStreamer describes types inherited from class InputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type InputStreamer interface {
	coreglib.Objector
	baseInputStream() *InputStream
}

var _ InputStreamer = (*InputStream)(nil)

func init() {
	coreglib.RegisterClassInfo[*InputStream, *InputStreamClass, InputStreamOverrides](
		GTypeInputStream,
		initInputStreamClass,
		wrapInputStream,
		defaultInputStreamOverrides,
	)
}

func initInputStreamClass(gclass unsafe.Pointer, overrides InputStreamOverrides, classInitFunc func(*InputStreamClass)) {
	if classInitFunc != nil {
		class := (*InputStreamClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapInputStream(obj *coreglib.Object) *InputStream {
	return &InputStream{
		Object: obj,
	}
}

func marshalInputStream(p uintptr) (interface{}, error) {
	return wrapInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *InputStream) baseInputStream() *InputStream {
	return v
}

// BaseInputStream returns the underlying base object.
func BaseInputStream(obj InputStreamer) *InputStream {
	return obj.baseInputStream()
}

// InputStreamClass: instance of this type is always passed by reference.
type InputStreamClass struct {
	*inputStreamClass
}

// inputStreamClass is the struct that's finalized.
type inputStreamClass struct {
	native unsafe.Pointer
}

var GIRInfoInputStreamClass = girepository.MustFind("Gio", "InputStreamClass")
