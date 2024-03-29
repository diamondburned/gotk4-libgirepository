// Code generated by girgen. DO NOT EDIT.

package gtk

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
	GTypeNativeSurface = coreglib.Type(girepository.MustFind("Gtk", "NativeSurface").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNativeSurface, F: marshalNativeSurface},
	})
}

// NativeSurface: GtkNative is the interface implemented by all widgets that
// have their own GdkSurface.
//
// The obvious example of a GtkNative is GtkWindow.
//
// Every widget that is not itself a GtkNative is contained in one, and you can
// get it with gtk.Widget.GetNative().
//
// To get the surface of a GtkNative, use gtk.Native.GetSurface(). It is also
// possible to find the GtkNative to which a surface belongs, with
// gtk.Native().GetForSurface.
//
// In addition to a gdk.Surface, a GtkNative also provides a gsk.Renderer for
// rendering on that surface. To get the renderer, use gtk.Native.GetRenderer().
//
// This type has been renamed from Native.
//
// NativeSurface wraps an interface. This means the user can get the
// underlying type by calling Cast().
type NativeSurface struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*NativeSurface)(nil)
)

// NativeSurfacer describes NativeSurface's interface methods.
type NativeSurfacer interface {
	coreglib.Objector

	baseNativeSurface() *NativeSurface
}

var _ NativeSurfacer = (*NativeSurface)(nil)

func wrapNativeSurface(obj *coreglib.Object) *NativeSurface {
	return &NativeSurface{
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

func marshalNativeSurface(p uintptr) (interface{}, error) {
	return wrapNativeSurface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *NativeSurface) baseNativeSurface() *NativeSurface {
	return v
}

// BaseNativeSurface returns the underlying base object.
func BaseNativeSurface(obj NativeSurfacer) *NativeSurface {
	return obj.baseNativeSurface()
}
