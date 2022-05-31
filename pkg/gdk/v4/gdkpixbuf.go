// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// PixbufGetFromSurface transfers image data from a cairo_surface_t and converts
// it to a GdkPixbuf.
//
// This allows you to efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the surface contains one.
//
// The function takes the following parameters:
//
//    - surface to copy from.
//    - srcX: source X coordinate within surface.
//    - srcY: source Y coordinate within surface.
//    - width: width in pixels of region to get.
//    - height: height in pixels of region to get.
//
// The function returns the following values:
//
//    - pixbuf (optional): newly-created pixbuf with a reference count of 1, or
//      NULL on error.
//
func PixbufGetFromSurface(surface *cairo.Surface, srcX, srcY, width, height int32) *gdkpixbuf.Pixbuf {
	var _args [5]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out
	var _arg2 C.int   // out
	var _arg3 C.int   // out
	var _arg4 C.int   // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(surface.Native()))
	_arg1 = C.int(srcX)
	_arg2 = C.int(srcY)
	_arg3 = C.int(width)
	_arg4 = C.int(height)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.int)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.int)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.int)(unsafe.Pointer(&_args[4])) = _arg4

	_gret := girepository.MustFind("Gdk", "pixbuf_get_from_surface").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(surface)
	runtime.KeepAlive(srcX)
	runtime.KeepAlive(srcY)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// PixbufGetFromTexture creates a new pixbuf from texture.
//
// This should generally not be used in newly written code as later stages will
// almost certainly convert the pixbuf back into a texture to draw it on screen.
//
// The function takes the following parameters:
//
//    - texture: GdkTexture.
//
// The function returns the following values:
//
//    - pixbuf (optional): new Pixbuf or NULL in case of an error.
//
func PixbufGetFromTexture(texture Texturer) *gdkpixbuf.Pixbuf {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(texture).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "pixbuf_get_from_texture").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(texture)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}
