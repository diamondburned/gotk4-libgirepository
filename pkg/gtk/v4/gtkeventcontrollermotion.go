// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_motion_get_type()), F: marshalEventControllerMotioner},
	})
}

// EventControllerMotion: GtkEventControllerMotion is an event controller
// tracking the pointer position.
//
// The event controller offers gtk.EventControllerMotion::enter and
// gtk.EventControllerMotion::leave signals, as well as
// gtk.EventControllerMotion:is-pointer and
// gtk.EventControllerMotion:contains-pointer properties which are updated to
// reflect changes in the pointer position as it moves over the widget.
type EventControllerMotion struct {
	EventController
}

func wrapEventControllerMotion(obj *externglib.Object) *EventControllerMotion {
	return &EventControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerMotioner(p uintptr) (interface{}, error) {
	return wrapEventControllerMotion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEventControllerMotion creates a new event controller that will handle
// motion events.
func NewEventControllerMotion() *EventControllerMotion {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_motion_new()

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = wrapEventControllerMotion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerMotion
}

// ContainsPointer returns if a pointer is within self or one of its children.
func (self *EventControllerMotion) ContainsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_contains_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPointer returns if a pointer is within self, but not one of its children.
func (self *EventControllerMotion) IsPointer() bool {
	var _arg0 *C.GtkEventControllerMotion // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkEventControllerMotion)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_event_controller_motion_is_pointer(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectEnter signals that the pointer has entered the widget.
func (self *EventControllerMotion) ConnectEnter(f func(x, y float64)) externglib.SignalHandle {
	return self.Connect("enter", f)
}

// ConnectLeave signals that the pointer has left the widget.
func (self *EventControllerMotion) ConnectLeave(f func()) externglib.SignalHandle {
	return self.Connect("leave", f)
}

// ConnectMotion: emitted when the pointer moves inside the widget.
func (self *EventControllerMotion) ConnectMotion(f func(x, y float64)) externglib.SignalHandle {
	return self.Connect("motion", f)
}
