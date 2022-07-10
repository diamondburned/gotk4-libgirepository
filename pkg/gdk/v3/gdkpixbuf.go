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
// #include <glib-object.h>
import "C"

// PixbufGetFromSurface transfers image data from a #cairo_surface_t and
// converts it to an RGB(A) representation inside a Pixbuf. This allows you to
// efficiently read individual pixels from cairo surfaces. For Windows, use
// gdk_pixbuf_get_from_window() instead.
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

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(surface.Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(srcX)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(srcY)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(width)
	*(*C.gint)(unsafe.Pointer(&_args[4])) = C.gint(height)

	_info := girepository.MustFind("Gdk", "pixbuf_get_from_surface")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GdkPixbuf)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(surface)
	runtime.KeepAlive(srcX)
	runtime.KeepAlive(srcY)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if *(**C.GdkPixbuf)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(*(**C.GdkPixbuf)(unsafe.Pointer(&_cret))))
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

// PixbufGetFromWindow transfers image data from a Window and converts it to an
// RGB(A) representation inside a Pixbuf. In other words, copies image data from
// a server-side drawable to a client-side RGB(A) buffer. This allows you to
// efficiently read individual pixels on the client side.
//
// This function will create an RGB pixbuf with 8 bits per channel with the size
// specified by the width and height arguments scaled by the scale factor of
// window. The pixbuf will contain an alpha channel if the window contains one.
//
// If the window is off the screen, then there is no image data in the
// obscured/offscreen regions to be placed in the pixbuf. The contents of
// portions of the pixbuf corresponding to the offscreen region are undefined.
//
// If the window you’re obtaining data from is partially obscured by other
// windows, then the contents of the pixbuf areas corresponding to the obscured
// regions are undefined.
//
// If the window is not mapped (typically because it’s iconified/minimized or
// not on the current workspace), then NULL will be returned.
//
// If memory can’t be allocated for the return value, NULL will be returned
// instead.
//
// (In short, there are several ways this function can fail, and if it fails it
// returns NULL; so check the return value.).
//
// The function takes the following parameters:
//
//    - window: source window.
//    - srcX: source X coordinate within window.
//    - srcY: source Y coordinate within window.
//    - width: width in pixels of region to get.
//    - height: height in pixels of region to get.
//
// The function returns the following values:
//
//    - pixbuf (optional): newly-created pixbuf with a reference count of 1, or
//      NULL on error.
//
func PixbufGetFromWindow(window Windower, srcX, srcY, width, height int32) *gdkpixbuf.Pixbuf {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(srcX)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(srcY)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(width)
	*(*C.gint)(unsafe.Pointer(&_args[4])) = C.gint(height)

	_info := girepository.MustFind("Gdk", "pixbuf_get_from_window")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.GdkPixbuf)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(window)
	runtime.KeepAlive(srcX)
	runtime.KeepAlive(srcY)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if *(**C.GdkPixbuf)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(*(**C.GdkPixbuf)(unsafe.Pointer(&_cret))))
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
