// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_cursor_get_type()), F: marshalX11Cursor},
	})
}

type X11Cursor interface {
	gextras.Objector

	// AsCursor casts the class to the gdk.Cursor interface.
	AsCursor() gdk.Cursor

	// GetCursorType returns the cursor type for this cursor.
	//
	// This method is inherited from gdk.Cursor
	GetCursorType() gdk.CursorType
	// GetDisplay returns the display on which the Cursor is defined.
	//
	// This method is inherited from gdk.Cursor
	GetDisplay() gdk.Display
	// GetImage returns a Pixbuf with the image used to display the cursor.
	//
	// Note that depending on the capabilities of the windowing system and on
	// the cursor, GDK may not be able to obtain the image data. In this case,
	// nil is returned.
	//
	// This method is inherited from gdk.Cursor
	GetImage() gdkpixbuf.Pixbuf
	// GetSurface returns a cairo image surface with the image used to display
	// the cursor.
	//
	// Note that depending on the capabilities of the windowing system and on
	// the cursor, GDK may not be able to obtain the image data. In this case,
	// nil is returned.
	//
	// This method is inherited from gdk.Cursor
	GetSurface() (xHot float64, yHot float64, surface *cairo.Surface)
	// Ref adds a reference to @cursor.
	//
	// Deprecated: since version 3.0.
	//
	// This method is inherited from gdk.Cursor
	Ref() gdk.Cursor
	// Unref removes a reference from @cursor, deallocating the cursor if no
	// references remain.
	//
	// Deprecated: since version 3.0.
	//
	// This method is inherited from gdk.Cursor
	Unref()
}

// x11Cursor implements the X11Cursor interface.
type x11Cursor struct {
	*externglib.Object
}

var _ X11Cursor = (*x11Cursor)(nil)

// WrapX11Cursor wraps a GObject to a type that implements
// interface X11Cursor. It is primarily used internally.
func WrapX11Cursor(obj *externglib.Object) X11Cursor {
	return x11Cursor{obj}
}

func marshalX11Cursor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Cursor(obj), nil
}

func (x x11Cursor) AsCursor() gdk.Cursor {
	return gdk.WrapCursor(gextras.InternObject(x))
}

func (c x11Cursor) GetCursorType() gdk.CursorType {
	return gdk.WrapCursor(gextras.InternObject(c)).GetCursorType()
}

func (c x11Cursor) GetDisplay() gdk.Display {
	return gdk.WrapCursor(gextras.InternObject(c)).GetDisplay()
}

func (c x11Cursor) GetImage() gdkpixbuf.Pixbuf {
	return gdk.WrapCursor(gextras.InternObject(c)).GetImage()
}

func (c x11Cursor) GetSurface() (xHot float64, yHot float64, surface *cairo.Surface) {
	return gdk.WrapCursor(gextras.InternObject(c)).GetSurface()
}

func (c x11Cursor) Ref() gdk.Cursor {
	return gdk.WrapCursor(gextras.InternObject(c)).Ref()
}

func (c x11Cursor) Unref() {
	gdk.WrapCursor(gextras.InternObject(c)).Unref()
}
