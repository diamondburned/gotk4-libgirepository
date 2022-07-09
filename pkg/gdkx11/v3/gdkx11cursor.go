// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeX11Cursor returns the GType for the type X11Cursor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11Cursor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11Cursor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11Cursor)
	return gtype
}

type X11Cursor struct {
	_ [0]func() // equal guard
	gdk.Cursor
}

var (
	_ gdk.Cursorrer = (*X11Cursor)(nil)
)

func wrapX11Cursor(obj *coreglib.Object) *X11Cursor {
	return &X11Cursor{
		Cursor: gdk.Cursor{
			Object: obj,
		},
	}
}

func marshalX11Cursor(p uintptr) (interface{}, error) {
	return wrapX11Cursor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
