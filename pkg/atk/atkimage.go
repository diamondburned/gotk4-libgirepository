// Code generated by girgen. DO NOT EDIT.

package atk

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
	GTypeImage = coreglib.Type(girepository.MustFind("Atk", "Image").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeImage, F: marshalImage},
	})
}

// ImageOverrider contains methods that are overridable.
type ImageOverrider interface {
}

// Image should be implemented by Object subtypes on behalf of components which
// display image/pixmap information onscreen, and which provide information
// (other than just widget borders, etc.) via that image content. For instance,
// icons, buttons with icons, toolbar elements, and image viewing panes
// typically should implement Image.
//
// Image primarily provides two types of information: coordinate information
// (useful for screen review mode of screenreaders, and for use by onscreen
// magnifiers), and descriptive information. The descriptive information is
// provided for alternative, text-only presentation of the most significant
// information present in the image.
//
// Image wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Image struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Image)(nil)
)

// Imager describes Image's interface methods.
type Imager interface {
	coreglib.Objector

	baseImage() *Image
}

var _ Imager = (*Image)(nil)

func ifaceInitImager(gifacePtr, data C.gpointer) {
}

func wrapImage(obj *coreglib.Object) *Image {
	return &Image{
		Object: obj,
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	return wrapImage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Image) baseImage() *Image {
	return v
}

// BaseImage returns the underlying base object.
func BaseImage(obj Imager) *Image {
	return obj.baseImage()
}

// ImageIface: instance of this type is always passed by reference.
type ImageIface struct {
	*imageIface
}

// imageIface is the struct that's finalized.
type imageIface struct {
	native unsafe.Pointer
}

var GIRInfoImageIface = girepository.MustFind("Atk", "ImageIface")
