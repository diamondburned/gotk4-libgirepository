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
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey: `GtkEventControllerKey` is an event controller that
// provides access to key events.
type EventControllerKey interface {
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

	// Forward forwards the current event of this @controller to a @widget.
	//
	// This function can only be used in handlers for the
	// [signal@Gtk.EventControllerKey::key-pressed],
	// [signal@Gtk.EventControllerKey::key-released] or
	// [signal@Gtk.EventControllerKey::modifiers] signals.
	Forward(widget Widget) bool
	// Group gets the key group of the current event of this @controller.
	//
	// See [method@Gdk.KeyEvent.get_layout].
	Group() uint
	// ImContext gets the input method context of the key @controller.
	ImContext() IMContext
	// SetImContext sets the input method context of the key @controller.
	SetImContext(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey interface.
type eventControllerKey struct {
	*externglib.Object
}

var _ EventControllerKey = (*eventControllerKey)(nil)

// WrapEventControllerKey wraps a GObject to a type that implements
// interface EventControllerKey. It is primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return eventControllerKey{obj}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

// NewEventControllerKey creates a new event controller that will handle key
// events.
func NewEventControllerKey() EventControllerKey {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_key_new()

	var _eventControllerKey EventControllerKey // out

	_eventControllerKey = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EventControllerKey)

	return _eventControllerKey
}

func (e eventControllerKey) AsEventController() EventController {
	return WrapEventController(gextras.InternObject(e))
}

func (c eventControllerKey) GetCurrentEvent() gdk.Event {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEvent()
}

func (c eventControllerKey) GetCurrentEventDevice() gdk.Device {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventDevice()
}

func (c eventControllerKey) GetCurrentEventState() gdk.ModifierType {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventState()
}

func (c eventControllerKey) GetCurrentEventTime() uint32 {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventTime()
}

func (c eventControllerKey) GetName() string {
	return WrapEventController(gextras.InternObject(c)).GetName()
}

func (c eventControllerKey) GetPropagationLimit() PropagationLimit {
	return WrapEventController(gextras.InternObject(c)).GetPropagationLimit()
}

func (c eventControllerKey) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c eventControllerKey) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c eventControllerKey) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c eventControllerKey) SetName(name string) {
	WrapEventController(gextras.InternObject(c)).SetName(name)
}

func (c eventControllerKey) SetPropagationLimit(limit PropagationLimit) {
	WrapEventController(gextras.InternObject(c)).SetPropagationLimit(limit)
}

func (c eventControllerKey) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (c eventControllerKey) Forward(widget Widget) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c eventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (c eventControllerKey) ImContext() IMContext {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)

	var _imContext IMContext // out

	_imContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IMContext)

	return _imContext
}

func (c eventControllerKey) SetImContext(imContext IMContext) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
