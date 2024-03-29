// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
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
	GTypeImageType = coreglib.Type(girepository.MustFind("Gtk", "ImageType").RegisteredGType())
	GTypeImage     = coreglib.Type(girepository.MustFind("Gtk", "Image").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeImageType, F: marshalImageType},
		coreglib.TypeMarshaler{T: GTypeImage, F: marshalImage},
	})
}

// ImageType describes the image data representation used by a gtk.Image.
//
// If you want to get the image from the widget, you can only get the
// currently-stored representation; for instance, if the
// gtk_image_get_storage_type() returns GTK_IMAGE_PAINTABLE, then you can call
// gtk_image_get_paintable().
//
// For empty images, you can request any storage type (call any of the "get"
// functions), but they will all return NULL values.
type ImageType C.gint

const (
	// ImageEmpty: there is no image displayed by the widget.
	ImageEmpty ImageType = iota
	// ImageIconName: widget contains a named icon.
	ImageIconName
	// ImageGIcon: widget contains a #GIcon.
	ImageGIcon
	// ImagePaintable: widget contains a Paintable.
	ImagePaintable
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ImageType.
func (i ImageType) String() string {
	switch i {
	case ImageEmpty:
		return "Empty"
	case ImageIconName:
		return "IconName"
	case ImageGIcon:
		return "GIcon"
	case ImagePaintable:
		return "Paintable"
	default:
		return fmt.Sprintf("ImageType(%d)", i)
	}
}

// Image: GtkImage widget displays an image.
//
// !An example GtkImage (image.png)
//
// Various kinds of object can be displayed as an image; most typically, you
// would load a GdkTexture from a file, using the convenience function
// gtk.Image.NewFromFile, for instance:
//
//    GtkWidget *image = gtk_image_new_from_file ("myfile.png");
//
//
// If the file isn’t loaded successfully, the image will contain a “broken
// image” icon similar to that used in many web browsers.
//
// If you want to handle errors in loading the file yourself, for example by
// displaying an error message, then load the image with
// gdk.Texture.NewFromFile, then create the GtkImage with
// gtk.Image.NewFromPaintable.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of GResource inside GIO, for
// details. In this case, gtk.Image:resource, gtk.Image.NewFromResource, and
// gtk.Image.SetFromResource() should be used.
//
// GtkImage displays its image as an icon, with a size that is determined by the
// application. See gtk.Picture if you want to show an image at is actual size.
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes
// .normal-icons or .large-icons may appear, depending on the
// gtk.Image:icon-size property.
//
//
// Accessibility
//
// GtkImage uses the GTK_ACCESSIBLE_ROLE_IMG role.
type Image struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Image)(nil)
)

func wrapImage(obj *coreglib.Object) *Image {
	return &Image{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	return wrapImage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
