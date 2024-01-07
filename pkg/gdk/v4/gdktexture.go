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
	GTypeTexture = coreglib.Type(girepository.MustFind("Gdk", "Texture").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTexture, F: marshalTexture},
	})
}

// Texture: GdkTexture is the basic element used to refer to pixel data.
//
// It is primarily meant for pixel data that will not change over multiple
// frames, and will be used for a long time.
//
// There are various ways to create GdkTexture objects from a GdkPixbuf, or a
// Cairo surface, or other pixel data.
//
// The ownership of the pixel data is transferred to the GdkTexture instance;
// you can only make a copy of it, via gdk.Texture.Download().
//
// GdkTexture is an immutable object: That means you cannot change anything
// about it other than increasing the reference count via g_object_ref().
type Texture struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Paintable
}

var (
	_ coreglib.Objector = (*Texture)(nil)
)

// Texturer describes types inherited from class Texture.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Texturer interface {
	coreglib.Objector
	baseTexture() *Texture
}

var _ Texturer = (*Texture)(nil)

func wrapTexture(obj *coreglib.Object) *Texture {
	return &Texture{
		Object: obj,
		Paintable: Paintable{
			Object: obj,
		},
	}
}

func marshalTexture(p uintptr) (interface{}, error) {
	return wrapTexture(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Texture) baseTexture() *Texture {
	return v
}

// BaseTexture returns the underlying base object.
func BaseTexture(obj Texturer) *Texture {
	return obj.baseTexture()
}
