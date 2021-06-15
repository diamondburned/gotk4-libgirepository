// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_type_get_type()), F: marshalImageType},
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImage},
	})
}

// ImageType describes the image data representation used by a Image. If you
// want to get the image from the widget, you can only get the currently-stored
// representation. e.g. if the gtk_image_get_storage_type() returns
// K_IMAGE_PIXBUF, then you can call gtk_image_get_pixbuf() but not
// gtk_image_get_stock(). For empty images, you can request any storage type
// (call any of the "get" functions), but they will all return nil values.
type ImageType int

const (
	// ImageTypeEmpty: there is no image displayed by the widget
	ImageTypeEmpty ImageType = 0
	// ImageTypePixbuf: the widget contains a Pixbuf
	ImageTypePixbuf ImageType = 1
	// ImageTypeStock: the widget contains a [stock item name][gtkstock]
	ImageTypeStock ImageType = 2
	// ImageTypeIconSet: the widget contains a IconSet
	ImageTypeIconSet ImageType = 3
	// ImageTypeAnimation: the widget contains a PixbufAnimation
	ImageTypeAnimation ImageType = 4
	// ImageTypeIconName: the widget contains a named icon. This image type was
	// added in GTK+ 2.6
	ImageTypeIconName ImageType = 5
	// ImageTypeGIcon: the widget contains a #GIcon. This image type was added
	// in GTK+ 2.14
	ImageTypeGIcon ImageType = 6
	// ImageTypeSurface: the widget contains a #cairo_surface_t. This image type
	// was added in GTK+ 3.10
	ImageTypeSurface ImageType = 7
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Image: the Image widget displays an image. Various kinds of object can be
// displayed as an image; most typically, you would load a Pixbuf ("pixel
// buffer") from a file, and then display that. There’s a convenience function
// to do this, gtk_image_new_from_file(), used as follows:
//
//      static gboolean
//      button_press_callback (GtkWidget      *event_box,
//                             GdkEventButton *event,
//                             gpointer        data)
//      {
//        g_print ("Event box clicked at coordinates f,f\n",
//                 event->x, event->y);
//
//        // Returning TRUE means we handled the event, so the signal
//        // emission should be stopped (don’t call any further callbacks
//        // that may be connected). Return FALSE to continue invoking callbacks.
//        return TRUE;
//      }
//
//      static GtkWidget*
//      create_image (void)
//      {
//        GtkWidget *image;
//        GtkWidget *event_box;
//
//        image = gtk_image_new_from_file ("myfile.png");
//
//        event_box = gtk_event_box_new ();
//
//        gtk_container_add (GTK_CONTAINER (event_box), image);
//
//        g_signal_connect (G_OBJECT (event_box),
//                          "button_press_event",
//                          G_CALLBACK (button_press_callback),
//                          image);
//
//        return image;
//      }
//
// When handling events on the event box, keep in mind that coordinates in the
// image may be different from event box coordinates due to the alignment and
// padding settings on the image (see Misc). The simplest way to solve this is
// to set the alignment to 0.0 (left/top), and set the padding to zero. Then the
// origin of the image will be the same as the origin of the event box.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. GTK+ comes with a program to avoid this, called
// “gdk-pixbuf-csource”. This library allows you to convert an image into a C
// variable declaration, which can then be loaded into a Pixbuf using
// gdk_pixbuf_new_from_inline().
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes may
// appear on image CSS nodes: .icon-dropshadow, .lowres-icon.
type Image interface {
	Misc
	Buildable

	// Clear resets the image to be empty.
	Clear()
	// Animation gets the PixbufAnimation being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION
	// (see gtk_image_get_storage_type()). The caller of this function does not
	// own a reference to the returned animation.
	Animation() gdkpixbuf.PixbufAnimation
	// GIcon gets the #GIcon and size being displayed by the Image. The storage
	// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
	// gtk_image_get_storage_type()). The caller of this function does not own a
	// reference to the returned #GIcon.
	GIcon() (gio.Icon, int)
	// IconName gets the icon name and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
	// (see gtk_image_get_storage_type()). The returned string is owned by the
	// Image and should not be freed.
	IconName() (string, int)
	// IconSet gets the icon set and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET
	// (see gtk_image_get_storage_type()).
	IconSet() (*IconSet, int)
	// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of
	// the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
	// gtk_image_get_storage_type()). The caller of this function does not own a
	// reference to the returned pixbuf.
	Pixbuf() gdkpixbuf.Pixbuf
	// PixelSize gets the pixel size used for named icons.
	PixelSize() int
	// Stock gets the stock icon name and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK (see
	// gtk_image_get_storage_type()). The returned string is owned by the Image
	// and should not be freed.
	Stock() (string, int)
	// StorageType gets the type of representation being used by the Image to
	// store image data. If the Image has no image data, the return value will
	// be GTK_IMAGE_EMPTY.
	StorageType() ImageType
	// SetFromAnimation causes the Image to display the given animation (or
	// display nothing, if you set the animation to nil).
	SetFromAnimation(animation gdkpixbuf.PixbufAnimation)
	// SetFromFile: see gtk_image_new_from_file() for details.
	SetFromFile(filename string)
	// SetFromGIcon: see gtk_image_new_from_gicon() for details.
	SetFromGIcon(icon gio.Icon, size int)
	// SetFromIconName: see gtk_image_new_from_icon_name() for details.
	SetFromIconName(iconName string, size int)
	// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
	SetFromIconSet(iconSet *IconSet, size int)
	// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
	SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf)
	// SetFromResource: see gtk_image_new_from_resource() for details.
	SetFromResource(resourcePath string)
	// SetFromStock: see gtk_image_new_from_stock() for details.
	SetFromStock(stockId string, size int)
	// SetFromSurface: see gtk_image_new_from_surface() for details.
	SetFromSurface(surface *cairo.Surface)
	// SetPixelSize sets the pixel size to use for named icons. If the pixel
	// size is set to a value != -1, it is used instead of the icon size set by
	// gtk_image_set_from_icon_name().
	SetPixelSize(pixelSize int)
}

// image implements the Image class.
type image struct {
	Misc
	Buildable
}

var _ Image = (*image)(nil)

// WrapImage wraps a GObject to the right type. It is
// primarily used internally.
func WrapImage(obj *externglib.Object) Image {
	return image{
		Misc:      WrapMisc(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImage(obj), nil
}

// NewImage constructs a class Image.
func NewImage() Image {
	var _cret C.GtkImage // in

	_cret = C.gtk_image_new()

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromAnimation constructs a class Image.
func NewImageFromAnimation(animation gdkpixbuf.PixbufAnimation) Image {
	var _arg1 *C.GdkPixbufAnimation // out
	var _cret C.GtkImage            // in

	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gtk_image_new_from_animation(_arg1)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromFile constructs a class Image.
func NewImageFromFile(filename string) Image {
	var _arg1 *C.gchar   // out
	var _cret C.GtkImage // in

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_file(_arg1)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromGIcon constructs a class Image.
func NewImageFromGIcon(icon gio.Icon, size int) Image {
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out
	var _cret C.GtkImage    // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = (C.GtkIconSize)(size)

	_cret = C.gtk_image_new_from_gicon(_arg1, _arg2)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconName constructs a class Image.
func NewImageFromIconName(iconName string, size int) Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret C.GtkImage    // in

	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkIconSize)(size)

	_cret = C.gtk_image_new_from_icon_name(_arg1, _arg2)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconSet constructs a class Image.
func NewImageFromIconSet(iconSet *IconSet, size int) Image {
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out
	var _cret C.GtkImage    // in

	_arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet.Native()))
	_arg2 = (C.GtkIconSize)(size)

	_cret = C.gtk_image_new_from_icon_set(_arg1, _arg2)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPixbuf constructs a class Image.
func NewImageFromPixbuf(pixbuf gdkpixbuf.Pixbuf) Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret C.GtkImage   // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_image_new_from_pixbuf(_arg1)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromResource constructs a class Image.
func NewImageFromResource(resourcePath string) Image {
	var _arg1 *C.gchar   // out
	var _cret C.GtkImage // in

	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_resource(_arg1)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromStock constructs a class Image.
func NewImageFromStock(stockId string, size int) Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret C.GtkImage    // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkIconSize)(size)

	_cret = C.gtk_image_new_from_stock(_arg1, _arg2)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromSurface constructs a class Image.
func NewImageFromSurface(surface *cairo.Surface) Image {
	var _arg1 *C.cairo_surface_t // out
	var _cret C.GtkImage         // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	_cret = C.gtk_image_new_from_surface(_arg1)

	var _image Image // out

	_image = WrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// Clear resets the image to be empty.
func (i image) Clear() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_clear(_arg0)
}

// Animation gets the PixbufAnimation being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION
// (see gtk_image_get_storage_type()). The caller of this function does not
// own a reference to the returned animation.
func (i image) Animation() gdkpixbuf.PixbufAnimation {
	var _arg0 *C.GtkImage           // out
	var _cret *C.GdkPixbufAnimation // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_animation(_arg0)

	var _pixbufAnimation gdkpixbuf.PixbufAnimation // out

	_pixbufAnimation = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.PixbufAnimation)

	return _pixbufAnimation
}

// GIcon gets the #GIcon and size being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned #GIcon.
func (i image) GIcon() (gio.Icon, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_get_gicon(_arg0, &_arg1, &_arg2)

	var _gicon gio.Icon // out
	var _size int       // out

	_gicon = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1))).(gio.Icon)
	_size = (int)(_arg2)

	return _gicon, _size
}

// IconName gets the icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
// (see gtk_image_get_storage_type()). The returned string is owned by the
// Image and should not be freed.
func (i image) IconName() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_get_icon_name(_arg0, &_arg1, &_arg2)

	var _iconName string // out
	var _size int        // out

	_iconName = C.GoString(_arg1)
	_size = (int)(_arg2)

	return _iconName, _size
}

// IconSet gets the icon set and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET
// (see gtk_image_get_storage_type()).
func (i image) IconSet() (*IconSet, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_get_icon_set(_arg0, &_arg1, &_arg2)

	var _iconSet *IconSet // out
	var _size int         // out

	_iconSet = WrapIconSet(unsafe.Pointer(_arg1))
	_size = (int)(_arg2)

	return _iconSet, _size
}

// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of
// the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned pixbuf.
func (i image) Pixbuf() gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkImage  // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_pixbuf(_arg0)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PixelSize gets the pixel size used for named icons.
func (i image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Stock gets the stock icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK (see
// gtk_image_get_storage_type()). The returned string is owned by the Image
// and should not be freed.
func (i image) Stock() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_get_stock(_arg0, &_arg1, &_arg2)

	var _stockId string // out
	var _size int       // out

	_stockId = C.GoString(_arg1)
	_size = (int)(_arg2)

	return _stockId, _size
}

// StorageType gets the type of representation being used by the Image to
// store image data. If the Image has no image data, the return value will
// be GTK_IMAGE_EMPTY.
func (i image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

// SetFromAnimation causes the Image to display the given animation (or
// display nothing, if you set the animation to nil).
func (i image) SetFromAnimation(animation gdkpixbuf.PixbufAnimation) {
	var _arg0 *C.GtkImage           // out
	var _arg1 *C.GdkPixbufAnimation // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	C.gtk_image_set_from_animation(_arg0, _arg1)
}

// SetFromFile: see gtk_image_new_from_file() for details.
func (i image) SetFromFile(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_file(_arg0, _arg1)
}

// SetFromGIcon: see gtk_image_new_from_gicon() for details.
func (i image) SetFromGIcon(icon gio.Icon, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = (C.GtkIconSize)(size)

	C.gtk_image_set_from_gicon(_arg0, _arg1, _arg2)
}

// SetFromIconName: see gtk_image_new_from_icon_name() for details.
func (i image) SetFromIconName(iconName string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkIconSize)(size)

	C.gtk_image_set_from_icon_name(_arg0, _arg1, _arg2)
}

// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
func (i image) SetFromIconSet(iconSet *IconSet, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet.Native()))
	_arg2 = (C.GtkIconSize)(size)

	C.gtk_image_set_from_icon_set(_arg0, _arg1, _arg2)
}

// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
func (i image) SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
}

// SetFromResource: see gtk_image_new_from_resource() for details.
func (i image) SetFromResource(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_resource(_arg0, _arg1)
}

// SetFromStock: see gtk_image_new_from_stock() for details.
func (i image) SetFromStock(stockId string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GtkIconSize)(size)

	C.gtk_image_set_from_stock(_arg0, _arg1, _arg2)
}

// SetFromSurface: see gtk_image_new_from_surface() for details.
func (i image) SetFromSurface(surface *cairo.Surface) {
	var _arg0 *C.GtkImage        // out
	var _arg1 *C.cairo_surface_t // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	C.gtk_image_set_from_surface(_arg0, _arg1)
}

// SetPixelSize sets the pixel size to use for named icons. If the pixel
// size is set to a value != -1, it is used instead of the icon size set by
// gtk_image_set_from_icon_name().
func (i image) SetPixelSize(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	_arg1 = (C.gint)(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
}
