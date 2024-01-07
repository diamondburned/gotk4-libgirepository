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
	GTypeGesturePan = coreglib.Type(girepository.MustFind("Gtk", "GesturePan").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGesturePan, F: marshalGesturePan},
	})
}

// GesturePan: GtkGesturePan is a GtkGesture for pan gestures.
//
// These are drags that are locked to happen along one axis. The axis that a
// GtkGesturePan handles is defined at construct time, and can be changed
// through gtk.GesturePan.SetOrientation().
//
// When the gesture starts to be recognized, GtkGesturePan will attempt to
// determine as early as possible whether the sequence is moving in the expected
// direction, and denying the sequence if this does not happen.
//
// Once a panning gesture along the expected axis is recognized, the
// gtk.GesturePan::pan signal will be emitted as input events are received,
// containing the offset in the given axis.
type GesturePan struct {
	_ [0]func() // equal guard
	GestureDrag
}

var (
	_ Gesturer = (*GesturePan)(nil)
)

func wrapGesturePan(obj *coreglib.Object) *GesturePan {
	return &GesturePan{
		GestureDrag: GestureDrag{
			GestureSingle: GestureSingle{
				Gesture: Gesture{
					EventController: EventController{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalGesturePan(p uintptr) (interface{}, error) {
	return wrapGesturePan(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
