// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocus},
	})
}

// EventControllerFocus: `GtkEventControllerFocus` is an event controller to
// keep track of keyboard focus.
//
// The event controller offers [signal@Gtk.EventControllerFocus::enter] and
// [signal@Gtk.EventControllerFocus::leave] signals, as well as
// [property@Gtk.EventControllerFocus:is-focus] and
// [property@Gtk.EventControllerFocus:contains-focus] properties which are
// updated to reflect focus changes inside the widget hierarchy that is rooted
// at the controllers widget.
type EventControllerFocus interface {
	gextras.Objector

	// AsEventController casts the class to the EventController interface.
	AsEventController() EventController

	// GetCurrentEvent returns the event that is currently being handled by the
	// controller, and nil at other times.
	//
	// This method is inherited from EventController
	GetCurrentEvent() gdk.Event
	// GetCurrentEventDevice returns the device of the event that is currently
	// being handled by the controller, and nil otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventDevice() gdk.Device
	// GetCurrentEventState returns the modifier state of the event that is
	// currently being handled by the controller, and 0 otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventState() gdk.ModifierType
	// GetCurrentEventTime returns the timestamp of the event that is currently
	// being handled by the controller, and 0 otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventTime() uint32
	// GetName gets the name of @controller.
	//
	// This method is inherited from EventController
	GetName() string
	// GetPropagationLimit gets the propagation limit of the event controller.
	//
	// This method is inherited from EventController
	GetPropagationLimit() PropagationLimit
	// GetPropagationPhase gets the propagation phase at which @controller
	// handles events.
	//
	// This method is inherited from EventController
	GetPropagationPhase() PropagationPhase
	// GetWidget returns the Widget this controller relates to.
	//
	// This method is inherited from EventController
	GetWidget() Widget
	// Reset resets the @controller to a clean state.
	//
	// This method is inherited from EventController
	Reset()
	// SetName sets a name on the controller that can be used for debugging.
	//
	// This method is inherited from EventController
	SetName(name string)
	// SetPropagationLimit sets the event propagation limit on the event
	// controller.
	//
	// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
	// events that are targeted at widgets on a different surface, such as
	// popovers.
	//
	// This method is inherited from EventController
	SetPropagationLimit(limit PropagationLimit)
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	//
	// If @phase is GTK_PHASE_NONE, no automatic event handling will be
	// performed, but other additional gesture maintenance will.
	//
	// This method is inherited from EventController
	SetPropagationPhase(phase PropagationPhase)

	// ContainsFocus returns true if focus is within @self or one of its
	// children.
	ContainsFocus() bool
	// IsFocus returns true if focus is within @self, but not one of its
	// children.
	IsFocus() bool
}

// eventControllerFocus implements the EventControllerFocus interface.
type eventControllerFocus struct {
	*externglib.Object
}

var _ EventControllerFocus = (*eventControllerFocus)(nil)

// WrapEventControllerFocus wraps a GObject to a type that implements
// interface EventControllerFocus. It is primarily used internally.
func WrapEventControllerFocus(obj *externglib.Object) EventControllerFocus {
	return eventControllerFocus{obj}
}

func marshalEventControllerFocus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerFocus(obj), nil
}

// NewEventControllerFocus creates a new event controller that will handle focus
// events.
func NewEventControllerFocus() EventControllerFocus {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_focus_new()

	var _eventControllerFocus EventControllerFocus // out

	_eventControllerFocus = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EventControllerFocus)

	return _eventControllerFocus
}

func (e eventControllerFocus) AsEventController() EventController {
	return WrapEventController(gextras.InternObject(e))
}

func (c eventControllerFocus) GetCurrentEvent() gdk.Event {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEvent()
}

func (c eventControllerFocus) GetCurrentEventDevice() gdk.Device {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventDevice()
}

func (c eventControllerFocus) GetCurrentEventState() gdk.ModifierType {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventState()
}

func (c eventControllerFocus) GetCurrentEventTime() uint32 {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventTime()
}

func (c eventControllerFocus) GetName() string {
	return WrapEventController(gextras.InternObject(c)).GetName()
}

func (c eventControllerFocus) GetPropagationLimit() PropagationLimit {
	return WrapEventController(gextras.InternObject(c)).GetPropagationLimit()
}

func (c eventControllerFocus) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c eventControllerFocus) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c eventControllerFocus) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c eventControllerFocus) SetName(name string) {
	WrapEventController(gextras.InternObject(c)).SetName(name)
}

func (c eventControllerFocus) SetPropagationLimit(limit PropagationLimit) {
	WrapEventController(gextras.InternObject(c)).SetPropagationLimit(limit)
}

func (c eventControllerFocus) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (s eventControllerFocus) ContainsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_event_controller_focus_contains_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s eventControllerFocus) IsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_event_controller_focus_is_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
