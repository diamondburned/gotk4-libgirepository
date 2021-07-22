// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_loader_get_type()), F: marshalPixbufLoaderer},
	})
}

// PixbufLoaderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PixbufLoaderOverrider interface {
	AreaPrepared()
	AreaUpdated(x int, y int, width int, height int)
	Closed()
	SizePrepared(width int, height int)
}

// PixbufLoader: incremental image loader.
//
// GdkPixbufLoader provides a way for applications to drive the process of
// loading an image, by letting them send the image data directly to the loader
// instead of having the loader read the data from a file. Applications can use
// this functionality instead of gdk_pixbuf_new_from_file() or
// gdk_pixbuf_animation_new_from_file() when they need to parse image data in
// small chunks. For example, it should be used when reading an image from a
// (potentially) slow network connection, or when loading an extremely large
// file.
//
// To use GdkPixbufLoader to load an image, create a new instance, and call
// gdkpixbuf.PixbufLoader.Write() to send the data to it. When done,
// gdkpixbuf.PixbufLoader.Close() should be called to end the stream and
// finalize everything.
//
// The loader will emit three important signals throughout the process:
//
//    - gdkpixbuf.PixbufLoader::size-prepared will be emitted as
//      soon as the image has enough information to determine the size of
//      the image to be used. If you want to scale the image while loading
//      it, you can call gdkpixbuf.PixbufLoader.SetSize() in
//      response to this signal.
//    - gdkpixbuf.PixbufLoader::area-prepared will be emitted as
//      soon as the pixbuf of the desired has been allocated. You can obtain
//      the GdkPixbuf instance by calling gdkpixbuf.PixbufLoader.GetPixbuf().
//      If you want to use it, simply acquire a reference to it. You can
//      also call gdk_pixbuf_loader_get_pixbuf() later to get the same
//      pixbuf.
//    - gdkpixbuf.PixbufLoader::area-updated will be emitted every
//      time a region is updated. This way you can update a partially
//      completed image. Note that you do not know anything about the
//      completeness of an image from the updated area. For example, in an
//      interlaced image you will need to make several passes before the
//      image is done loading.
//
//
// Loading an animation
//
// Loading an animation is almost as easy as loading an image. Once the first
// gdkpixbuf.PixbufLoader::area-prepared signal has been emitted, you can call
// gdkpixbuf.PixbufLoader.GetAnimation() to get the gdkpixbuf.PixbufAnimation
// instance, and then call and gdkpixbuf.PixbufAnimation.GetIter() to get a
// gdkpixbuf.PixbufAnimationIter to retrieve the pixbuf for the desired time
// stamp.
type PixbufLoader struct {
	*externglib.Object
}

func wrapPixbufLoader(obj *externglib.Object) *PixbufLoader {
	return &PixbufLoader{
		Object: obj,
	}
}

func marshalPixbufLoaderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufLoader(obj), nil
}

// NewPixbufLoader creates a new pixbuf loader object.
func NewPixbufLoader() *PixbufLoader {
	var _cret *C.GdkPixbufLoader // in

	_cret = C.gdk_pixbuf_loader_new()

	var _pixbufLoader *PixbufLoader // out

	_pixbufLoader = wrapPixbufLoader(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pixbufLoader
}

// NewPixbufLoaderWithMIMEType creates a new pixbuf loader object that always
// attempts to parse image data as if it were an image of MIME type mime_type,
// instead of identifying the type automatically.
//
// This function is useful if you want an error if the image isn't the expected
// MIME type; for loading image formats that can't be reliably identified by
// looking at the data; or if the user manually forces a specific MIME type.
//
// The list of supported mime types depends on what image loaders are installed,
// but typically "image/png", "image/jpeg", "image/gif", "image/tiff" and
// "image/x-xpixmap" are among the supported mime types. To obtain the full list
// of supported mime types, call gdk_pixbuf_format_get_mime_types() on each of
// the PixbufFormat structs returned by gdk_pixbuf_get_formats().
func NewPixbufLoaderWithMIMEType(mimeType string) (*PixbufLoader, error) {
	var _arg1 *C.char            // out
	var _cret *C.GdkPixbufLoader // in
	var _cerr *C.GError          // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_loader_new_with_mime_type(_arg1, &_cerr)

	var _pixbufLoader *PixbufLoader // out
	var _goerr error                // out

	_pixbufLoader = wrapPixbufLoader(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufLoader, _goerr
}

// NewPixbufLoaderWithType creates a new pixbuf loader object that always
// attempts to parse image data as if it were an image of type image_type,
// instead of identifying the type automatically.
//
// This function is useful if you want an error if the image isn't the expected
// type; for loading image formats that can't be reliably identified by looking
// at the data; or if the user manually forces a specific type.
//
// The list of supported image formats depends on what image loaders are
// installed, but typically "png", "jpeg", "gif", "tiff" and "xpm" are among the
// supported formats. To obtain the full list of supported image formats, call
// gdk_pixbuf_format_get_name() on each of the PixbufFormat structs returned by
// gdk_pixbuf_get_formats().
func NewPixbufLoaderWithType(imageType string) (*PixbufLoader, error) {
	var _arg1 *C.char            // out
	var _cret *C.GdkPixbufLoader // in
	var _cerr *C.GError          // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(imageType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_loader_new_with_type(_arg1, &_cerr)

	var _pixbufLoader *PixbufLoader // out
	var _goerr error                // out

	_pixbufLoader = wrapPixbufLoader(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufLoader, _goerr
}

// Close informs a pixbuf loader that no further writes with
// gdk_pixbuf_loader_write() will occur, so that it can free its internal
// loading structures.
//
// This function also tries to parse any data that hasn't yet been parsed; if
// the remaining data is partial or corrupt, an error will be returned.
//
// If FALSE is returned, error will be set to an error from the GDK_PIXBUF_ERROR
// or G_FILE_ERROR domains.
//
// If you're just cancelling a load rather than expecting it to be finished,
// passing NULL for error to ignore it is reasonable.
//
// Remember that this function does not release a reference on the loader, so
// you will need to explicitly release any reference you hold.
func (loader *PixbufLoader) Close() error {
	var _arg0 *C.GdkPixbufLoader // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))

	C.gdk_pixbuf_loader_close(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Animation queries the PixbufAnimation that a pixbuf loader is currently
// creating.
//
// In general it only makes sense to call this function after the
// gdkpixbuf.PixbufLoader::area-prepared signal has been emitted by the loader.
//
// If the loader doesn't have enough bytes yet, and hasn't emitted the
// area-prepared signal, this function will return NULL.
func (loader *PixbufLoader) Animation() *PixbufAnimation {
	var _arg0 *C.GdkPixbufLoader    // out
	var _cret *C.GdkPixbufAnimation // in

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))

	_cret = C.gdk_pixbuf_loader_get_animation(_arg0)

	var _pixbufAnimation *PixbufAnimation // out

	_pixbufAnimation = wrapPixbufAnimation(externglib.Take(unsafe.Pointer(_cret)))

	return _pixbufAnimation
}

// Format obtains the available information about the format of the currently
// loading image file.
func (loader *PixbufLoader) Format() *PixbufFormat {
	var _arg0 *C.GdkPixbufLoader // out
	var _cret *C.GdkPixbufFormat // in

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))

	_cret = C.gdk_pixbuf_loader_get_format(_arg0)

	var _pixbufFormat *PixbufFormat // out

	_pixbufFormat = (*PixbufFormat)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _pixbufFormat
}

// Pixbuf queries the Pixbuf that a pixbuf loader is currently creating.
//
// In general it only makes sense to call this function after the
// gdkpixbuf.PixbufLoader::area-prepared signal has been emitted by the loader;
// this means that enough data has been read to know the size of the image that
// will be allocated.
//
// If the loader has not received enough data via gdk_pixbuf_loader_write(),
// then this function returns NULL.
//
// The returned pixbuf will be the same in all future calls to the loader, so if
// you want to keep using it, you should acquire a reference to it.
//
// Additionally, if the loader is an animation, it will return the "static
// image" of the animation (see gdk_pixbuf_animation_get_static_image()).
func (loader *PixbufLoader) Pixbuf() *Pixbuf {
	var _arg0 *C.GdkPixbufLoader // out
	var _cret *C.GdkPixbuf       // in

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))

	_cret = C.gdk_pixbuf_loader_get_pixbuf(_arg0)

	var _pixbuf *Pixbuf // out

	_pixbuf = wrapPixbuf(externglib.Take(unsafe.Pointer(_cret)))

	return _pixbuf
}

// SetSize causes the image to be scaled while it is loaded.
//
// The desired image size can be determined relative to the original size of the
// image by calling gdk_pixbuf_loader_set_size() from a signal handler for the
// ::size-prepared signal.
//
// Attempts to set the desired image size are ignored after the emission of the
// ::size-prepared signal.
func (loader *PixbufLoader) SetSize(width int, height int) {
	var _arg0 *C.GdkPixbufLoader // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gdk_pixbuf_loader_set_size(_arg0, _arg1, _arg2)
}

// Write parses the next count bytes in the given image buffer.
func (loader *PixbufLoader) Write(buf []byte) error {
	var _arg0 *C.GdkPixbufLoader // out
	var _arg1 *C.guchar          // out
	var _arg2 C.gsize
	var _cerr *C.GError // in

	_arg0 = (*C.GdkPixbufLoader)(unsafe.Pointer(loader.Native()))
	_arg2 = (C.gsize)(len(buf))
	if len(buf) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&buf[0]))
	}

	C.gdk_pixbuf_loader_write(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
