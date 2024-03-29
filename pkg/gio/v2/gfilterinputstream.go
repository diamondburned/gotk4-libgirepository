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
	GTypeFilterInputStream = coreglib.Type(girepository.MustFind("Gio", "FilterInputStream").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFilterInputStream, F: marshalFilterInputStream},
	})
}

// FilterInputStreamOverrides contains methods that are overridable.
type FilterInputStreamOverrides struct {
}

func defaultFilterInputStreamOverrides(v *FilterInputStream) FilterInputStreamOverrides {
	return FilterInputStreamOverrides{}
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream struct {
	_ [0]func() // equal guard
	InputStream
}

var (
	_ InputStreamer = (*FilterInputStream)(nil)
)

// FilterInputStreamer describes types inherited from class FilterInputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FilterInputStreamer interface {
	coreglib.Objector
	baseFilterInputStream() *FilterInputStream
}

var _ FilterInputStreamer = (*FilterInputStream)(nil)

func init() {
	coreglib.RegisterClassInfo[*FilterInputStream, *FilterInputStreamClass, FilterInputStreamOverrides](
		GTypeFilterInputStream,
		initFilterInputStreamClass,
		wrapFilterInputStream,
		defaultFilterInputStreamOverrides,
	)
}

func initFilterInputStreamClass(gclass unsafe.Pointer, overrides FilterInputStreamOverrides, classInitFunc func(*FilterInputStreamClass)) {
	if classInitFunc != nil {
		class := (*FilterInputStreamClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFilterInputStream(obj *coreglib.Object) *FilterInputStream {
	return &FilterInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalFilterInputStream(p uintptr) (interface{}, error) {
	return wrapFilterInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *FilterInputStream) baseFilterInputStream() *FilterInputStream {
	return v
}

// BaseFilterInputStream returns the underlying base object.
func BaseFilterInputStream(obj FilterInputStreamer) *FilterInputStream {
	return obj.baseFilterInputStream()
}

// FilterInputStreamClass: instance of this type is always passed by reference.
type FilterInputStreamClass struct {
	*filterInputStreamClass
}

// filterInputStreamClass is the struct that's finalized.
type filterInputStreamClass struct {
	native unsafe.Pointer
}

var GIRInfoFilterInputStreamClass = girepository.MustFind("Gio", "FilterInputStreamClass")
