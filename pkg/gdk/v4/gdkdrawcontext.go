// Code generated by girgen. DO NOT EDIT.

package gdk

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
	GTypeDrawContext = coreglib.Type(girepository.MustFind("Gdk", "DrawContext").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDrawContext, F: marshalDrawContext},
	})
}

// DrawContext: base class for objects implementing different rendering methods.
//
// GdkDrawContext is the base object used by contexts implementing different
// rendering methods, such as gdk.CairoContext or gdk.GLContext. It provides
// shared functionality between those contexts.
//
// You will always interact with one of those subclasses.
//
// A GdkDrawContext is always associated with a single toplevel surface.
type DrawContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DrawContext)(nil)
)

// DrawContexter describes types inherited from class DrawContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DrawContexter interface {
	coreglib.Objector
	baseDrawContext() *DrawContext
}

var _ DrawContexter = (*DrawContext)(nil)

func wrapDrawContext(obj *coreglib.Object) *DrawContext {
	return &DrawContext{
		Object: obj,
	}
}

func marshalDrawContext(p uintptr) (interface{}, error) {
	return wrapDrawContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *DrawContext) baseDrawContext() *DrawContext {
	return v
}

// BaseDrawContext returns the underlying base object.
func BaseDrawContext(obj DrawContexter) *DrawContext {
	return obj.baseDrawContext()
}
