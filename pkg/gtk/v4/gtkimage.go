// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeImageType returns the GType for the type ImageType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeImageType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ImageType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalImageType)
	return gtype
}

// GTypeImage returns the GType for the type Image.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeImage() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Image").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalImage)
	return gtype
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

// NewImage creates a new empty GtkImage widget.
//
// The function returns the following values:
//
//    - image: newly created GtkImage widget.
//
func NewImage() *Image {
	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromFile creates a new GtkImage displaying the file filename.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use gdk.Texture.NewFromFile
// to load the file yourself, then create the GtkImage from the texture.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
//
// The function takes the following parameters:
//
//    - filename: filename.
//
// The function returns the following values:
//
//    - image: new GtkImage.
//
func NewImageFromFile(filename string) *Image {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filename)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromGIcon creates a GtkImage displaying an icon from the current icon
// theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - icon: icon.
//
// The function returns the following values:
//
//    - image: new GtkImage displaying the themed icon.
//
func NewImageFromGIcon(icon gio.Iconner) *Image {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_gicon", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(icon)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconName creates a GtkImage displaying an icon from the current
// icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name or NULL.
//
// The function returns the following values:
//
//    - image: new GtkImage displaying the themed icon.
//
func NewImageFromIconName(iconName string) *Image {
	var _args [1]girepository.Argument

	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iconName)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPaintable creates a new GtkImage displaying paintable.
//
// The GtkImage does not assume a reference to the paintable; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// The GtkImage will track changes to the paintable and update its size and
// contents in response to it.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable, or NULL.
//
// The function returns the following values:
//
//    - image: new GtkImage.
//
func NewImageFromPaintable(paintable gdk.Paintabler) *Image {
	var _args [1]girepository.Argument

	if paintable != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_paintable", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paintable)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPixbuf creates a new GtkImage displaying pixbuf.
//
// The GtkImage does not assume a reference to the pixbuf; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// This is a helper for gtk.Image.NewFromPaintable, and you can't get back the
// exact pixbuf once this is called, only a texture.
//
// Note that this function just creates an GtkImage from the pixbuf. The
// GtkImage created will not react to state changes. Should you want that, you
// should use gtk.Image.NewFromIconName.
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf, or NULL.
//
// The function returns the following values:
//
//    - image: new GtkImage.
//
func NewImageFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	var _args [1]girepository.Argument

	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_pixbuf", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixbuf)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromResource creates a new GtkImage displaying the resource file
// resource_path.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use
// gdkpixbuf.Pixbuf.NewFromFile to load the file yourself, then create the
// GtkImage from the pixbuf.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
//
// The function takes the following parameters:
//
//    - resourcePath: resource path.
//
// The function returns the following values:
//
//    - image: new GtkImage.
//
func NewImageFromResource(resourcePath string) *Image {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("new_Image_from_resource", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resourcePath)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// Clear resets the image to be empty.
func (image *Image) Clear() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	girepository.MustFind("Gtk", "Image").InvokeMethod("clear", _args[:], nil)

	runtime.KeepAlive(image)
}

// GIcon gets the GIcon being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned GIcon.
//
// The function returns the following values:
//
//    - icon (optional): GIcon, or NULL.
//
func (image *Image) GIcon() *gio.Icon {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("get_gicon", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)

	var _icon *gio.Icon // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_icon = &gio.Icon{
				Object: obj,
			}
		}
	}

	return _icon
}

// IconName gets the icon name and size being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
// (see gtk.Image.GetStorageType()). The returned string is owned by the
// GtkImage and should not be freed.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name, or NULL.
//
func (image *Image) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("get_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Paintable gets the image GdkPaintable being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PAINTABLE
// (see gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned paintable.
//
// The function returns the following values:
//
//    - paintable (optional): displayed paintable, or NULL if the image is empty.
//
func (image *Image) Paintable() *gdk.Paintable {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("get_paintable", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)

	var _paintable *gdk.Paintable // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_paintable = &gdk.Paintable{
				Object: obj,
			}
		}
	}

	return _paintable
}

// PixelSize gets the pixel size used for named icons.
//
// The function returns the following values:
//
//    - gint: pixel size used for named icons.
//
func (image *Image) PixelSize() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_gret := girepository.MustFind("Gtk", "Image").InvokeMethod("get_pixel_size", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(image)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// SetFromFile sets a GtkImage to show a file.
//
// See gtk.Image.NewFromFile for details.
//
// The function takes the following parameters:
//
//    - filename (optional) or NULL.
//
func (image *Image) SetFromFile(filename string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if filename != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_file", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(filename)
}

// SetFromGIcon sets a GtkImage to show a GIcon.
//
// See gtk.Image.NewFromGIcon for details.
//
// The function takes the following parameters:
//
//    - icon: icon.
//
func (image *Image) SetFromGIcon(icon gio.Iconner) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_gicon", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(icon)
}

// SetFromIconName sets a GtkImage to show a named icon.
//
// See gtk.Image.NewFromIconName for details.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name or NULL.
//
func (image *Image) SetFromIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_icon_name", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(iconName)
}

// SetFromPaintable sets a GtkImage to show a GdkPaintable.
//
// See gtk.Image.NewFromPaintable for details.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable or NULL.
//
func (image *Image) SetFromPaintable(paintable gdk.Paintabler) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if paintable != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_paintable", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(paintable)
}

// SetFromPixbuf sets a GtkImage to show a GdkPixbuf.
//
// See gtk.Image.NewFromPixbuf for details.
//
// Note: This is a helper for gtk.Image.SetFromPaintable(), and you can't get
// back the exact pixbuf once this is called, only a paintable.
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf or NULL.
//
func (image *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_pixbuf", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(pixbuf)
}

// SetFromResource sets a GtkImage to show a resource.
//
// See gtk.Image.NewFromResource for details.
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource path or NULL.
//
func (image *Image) SetFromResource(resourcePath string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if resourcePath != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_from_resource", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(resourcePath)
}

// SetPixelSize sets the pixel size to use for named icons.
//
// If the pixel size is set to a value != -1, it is used instead of the icon
// size set by gtk.Image.SetFromIconName().
//
// The function takes the following parameters:
//
//    - pixelSize: new pixel size.
//
func (image *Image) SetPixelSize(pixelSize int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(pixelSize)

	girepository.MustFind("Gtk", "Image").InvokeMethod("set_pixel_size", _args[:], nil)

	runtime.KeepAlive(image)
	runtime.KeepAlive(pixelSize)
}
