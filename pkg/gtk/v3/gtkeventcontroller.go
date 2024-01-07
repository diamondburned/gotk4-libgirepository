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
	GTypeEventController = coreglib.Type(girepository.MustFind("Gtk", "EventController").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEventController, F: marshalEventController},
	})
}

// EventController is a base, low-level implementation for event controllers.
// Those react to a series of Events, and possibly trigger actions as a
// consequence of those.
type EventController struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EventController)(nil)
)

// EventControllerer describes types inherited from class EventController.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type EventControllerer interface {
	coreglib.Objector
	baseEventController() *EventController
}

var _ EventControllerer = (*EventController)(nil)

func wrapEventController(obj *coreglib.Object) *EventController {
	return &EventController{
		Object: obj,
	}
}

func marshalEventController(p uintptr) (interface{}, error) {
	return wrapEventController(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *EventController) baseEventController() *EventController {
	return v
}

// BaseEventController returns the underlying base object.
func BaseEventController(obj EventControllerer) *EventController {
	return obj.baseEventController()
}
