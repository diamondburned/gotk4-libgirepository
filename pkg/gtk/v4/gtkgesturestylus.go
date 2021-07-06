// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylus},
	})
}

// GestureStylus: `GtkGestureStylus` is a `GtkGesture` specific to stylus input.
//
// The provided signals just relay the basic information of the stylus events.
type GestureStylus interface {
	gextras.Objector

	// AsGestureSingle casts the class to the GestureSingle interface.
	AsGestureSingle() GestureSingle

	// GetButton returns the button number @gesture listens for.
	//
	// If this is 0, the gesture reacts to any button press.
	//
	// This method is inherited from GestureSingle
	GetButton() uint
	// GetCurrentButton returns the button number currently interacting with
	// @gesture, or 0 if there is none.
	//
	// This method is inherited from GestureSingle
	GetCurrentButton() uint
	// GetCurrentSequence returns the event sequence currently interacting with
	// @gesture.
	//
	// This is only meaningful if [method@Gtk.Gesture.is_active] returns true.
	//
	// This method is inherited from GestureSingle
	GetCurrentSequence() *gdk.EventSequence
	// GetExclusive gets whether a gesture is exclusive.
	//
	// For more information, see [method@Gtk.GestureSingle.set_exclusive].
	//
	// This method is inherited from GestureSingle
	GetExclusive() bool
	// GetTouchOnly returns true if the gesture is only triggered by touch
	// events.
	//
	// This method is inherited from GestureSingle
	GetTouchOnly() bool
	// SetButton sets the button number @gesture listens to.
	//
	// If non-0, every button press from a different button number will be
	// ignored. Touch events implicitly match with button 1.
	//
	// This method is inherited from GestureSingle
	SetButton(button uint)
	// SetExclusive sets whether @gesture is exclusive.
	//
	// An exclusive gesture will only handle pointer and "pointer emulated"
	// touch events, so at any given time, there is only one sequence able to
	// interact with those.
	//
	// This method is inherited from GestureSingle
	SetExclusive(exclusive bool)
	// SetTouchOnly sets whether to handle only touch events.
	//
	// If @touch_only is true, @gesture will only handle events of type
	// GDK_TOUCH_BEGIN, GDK_TOUCH_UPDATE or GDK_TOUCH_END. If false, mouse
	// events will be handled too.
	//
	// This method is inherited from GestureSingle
	SetTouchOnly(touchOnly bool)
	// GetBoundingBox: if there are touch sequences being currently handled by
	// @gesture, returns true and fills in @rect with the bounding box
	// containing all active touches.
	//
	// Otherwise, false will be returned.
	//
	// Note: This function will yield unexpected results on touchpad gestures.
	// Since there is no correlation between physical and pixel distances, these
	// will look as if constrained in an infinitely small area, @rect width and
	// height will thus be 0 regardless of the number of touchpoints.
	//
	// This method is inherited from Gesture
	GetBoundingBox() (gdk.Rectangle, bool)
	// GetBoundingBoxCenter: if there are touch sequences being currently
	// handled by @gesture, returns true and fills in @x and @y with the center
	// of the bounding box containing all active touches.
	//
	// Otherwise, false will be returned.
	//
	// This method is inherited from Gesture
	GetBoundingBoxCenter() (x float64, y float64, ok bool)
	// GetDevice returns the logical `GdkDevice` that is currently operating on
	// @gesture.
	//
	// This returns nil if the gesture is not being interacted.
	//
	// This method is inherited from Gesture
	GetDevice() gdk.Device
	// GetLastEvent returns the last event that was processed for @sequence.
	//
	// Note that the returned pointer is only valid as long as the @sequence is
	// still interpreted by the @gesture. If in doubt, you should make a copy of
	// the event.
	//
	// This method is inherited from Gesture
	GetLastEvent(sequence *gdk.EventSequence) gdk.Event
	// GetLastUpdatedSequence returns the `GdkEventSequence` that was last
	// updated on @gesture.
	//
	// This method is inherited from Gesture
	GetLastUpdatedSequence() *gdk.EventSequence
	// GetPoint: if @sequence is currently being interpreted by @gesture,
	// returns true and fills in @x and @y with the last coordinates stored for
	// that event sequence.
	//
	// The coordinates are always relative to the widget allocation.
	//
	// This method is inherited from Gesture
	GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool)
	// GetSequenceState returns the @sequence state, as seen by @gesture.
	//
	// This method is inherited from Gesture
	GetSequenceState(sequence *gdk.EventSequence) EventSequenceState
	// Group adds @gesture to the same group than @group_gesture.
	//
	// Gestures are by default isolated in their own groups.
	//
	// Both gestures must have been added to the same widget before they can be
	// grouped.
	//
	// When gestures are grouped, the state of `GdkEventSequences` is kept in
	// sync for all of those, so calling
	// [method@Gtk.Gesture.set_sequence_state], on one will transfer the same
	// value to the others.
	//
	// Groups also perform an "implicit grabbing" of sequences, if a
	// `GdkEventSequence` state is set to GTK_EVENT_SEQUENCE_CLAIMED on one
	// group, every other gesture group attached to the same `GtkWidget` will
	// switch the state for that sequence to GTK_EVENT_SEQUENCE_DENIED.
	//
	// This method is inherited from Gesture
	Group(gesture Gesture)
	// HandlesSequence returns true if @gesture is currently handling events
	// corresponding to @sequence.
	//
	// This method is inherited from Gesture
	HandlesSequence(sequence *gdk.EventSequence) bool
	// IsActive returns true if the gesture is currently active.
	//
	// A gesture is active while there are touch sequences interacting with it.
	//
	// This method is inherited from Gesture
	IsActive() bool
	// IsGroupedWith returns true if both gestures pertain to the same group.
	//
	// This method is inherited from Gesture
	IsGroupedWith(other Gesture) bool
	// IsRecognized returns true if the gesture is currently recognized.
	//
	// A gesture is recognized if there are as many interacting touch sequences
	// as required by @gesture.
	//
	// This method is inherited from Gesture
	IsRecognized() bool
	// SetSequenceState sets the state of @sequence in @gesture.
	//
	// Sequences start in state GTK_EVENT_SEQUENCE_NONE, and whenever they
	// change state, they can never go back to that state. Likewise, sequences
	// in state GTK_EVENT_SEQUENCE_DENIED cannot turn back to a not denied
	// state. With these rules, the lifetime of an event sequence is constrained
	// to the next four:
	//
	// * None * None → Denied * None → Claimed * None → Claimed → Denied
	//
	// Note: Due to event handling ordering, it may be unsafe to set the state
	// on another gesture within a [signal@Gtk.Gesture::begin] signal handler,
	// as the callback might be executed before the other gesture knows about
	// the sequence. A safe way to perform this could be:
	//
	// “`c static void first_gesture_begin_cb (GtkGesture *first_gesture,
	// GdkEventSequence *sequence, gpointer user_data) {
	// gtk_gesture_set_sequence_state (first_gesture, sequence,
	// GTK_EVENT_SEQUENCE_CLAIMED); gtk_gesture_set_sequence_state
	// (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED); }
	//
	// static void second_gesture_begin_cb (GtkGesture *second_gesture,
	// GdkEventSequence *sequence, gpointer user_data) { if
	// (gtk_gesture_get_sequence_state (first_gesture, sequence) ==
	// GTK_EVENT_SEQUENCE_CLAIMED) gtk_gesture_set_sequence_state
	// (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED); } “`
	//
	// If both gestures are in the same group, just set the state on the gesture
	// emitting the event, the sequence will be already be initialized to the
	// group's global state when the second gesture processes the event.
	//
	// This method is inherited from Gesture
	SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool
	// SetState sets the state of all sequences that @gesture is currently
	// interacting with.
	//
	// See [method@Gtk.Gesture.set_sequence_state] for more details on sequence
	// states.
	//
	// This method is inherited from Gesture
	SetState(state EventSequenceState) bool
	// Ungroup separates @gesture into an isolated group.
	//
	// This method is inherited from Gesture
	Ungroup()
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

	// Axis returns the current value for the requested @axis.
	//
	// This function must be called from the handler of one of the
	// [signal@Gtk.GestureStylus::down], [signal@Gtk.GestureStylus::motion],
	// [signal@Gtk.GestureStylus::up] or [signal@Gtk.GestureStylus::proximity]
	// signals.
	Axis(axis gdk.AxisUse) (float64, bool)
	// Backlog returns the accumulated backlog of tracking information.
	//
	// By default, GTK will limit rate of input events. On stylus input where
	// accuracy of strokes is paramount, this function returns the accumulated
	// coordinate/timing state before the emission of the current
	// [Gtk.GestureStylus::motion] signal.
	//
	// This function may only be called within a
	// [signal@Gtk.GestureStylus::motion] signal handler, the state given in
	// this signal and obtainable through [method@Gtk.GestureStylus.get_axis]
	// express the latest (most up-to-date) state in motion history.
	//
	// The @backlog is provided in chronological order.
	Backlog() ([]gdk.TimeCoord, bool)
	// DeviceTool returns the `GdkDeviceTool` currently driving input through
	// this gesture.
	//
	// This function must be called from the handler of one of the
	// [signal@Gtk.GestureStylus::down], [signal@Gtk.GestureStylus::motion],
	// [signal@Gtk.GestureStylus::up] or [signal@Gtk.GestureStylus::proximity]
	// signals.
	DeviceTool() gdk.DeviceTool
}

// gestureStylus implements the GestureStylus interface.
type gestureStylus struct {
	*externglib.Object
}

var _ GestureStylus = (*gestureStylus)(nil)

// WrapGestureStylus wraps a GObject to a type that implements
// interface GestureStylus. It is primarily used internally.
func WrapGestureStylus(obj *externglib.Object) GestureStylus {
	return gestureStylus{obj}
}

func marshalGestureStylus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureStylus(obj), nil
}

// NewGestureStylus creates a new `GtkGestureStylus`.
func NewGestureStylus() GestureStylus {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_stylus_new()

	var _gestureStylus GestureStylus // out

	_gestureStylus = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(GestureStylus)

	return _gestureStylus
}

func (g gestureStylus) AsGestureSingle() GestureSingle {
	return WrapGestureSingle(gextras.InternObject(g))
}

func (g gestureStylus) GetButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetButton()
}

func (g gestureStylus) GetCurrentButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentButton()
}

func (g gestureStylus) GetCurrentSequence() *gdk.EventSequence {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentSequence()
}

func (g gestureStylus) GetExclusive() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetExclusive()
}

func (g gestureStylus) GetTouchOnly() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetTouchOnly()
}

func (g gestureStylus) SetButton(button uint) {
	WrapGestureSingle(gextras.InternObject(g)).SetButton(button)
}

func (g gestureStylus) SetExclusive(exclusive bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetExclusive(exclusive)
}

func (g gestureStylus) SetTouchOnly(touchOnly bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetTouchOnly(touchOnly)
}

func (g gestureStylus) GetBoundingBox() (gdk.Rectangle, bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBox()
}

func (g gestureStylus) GetBoundingBoxCenter() (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBoxCenter()
}

func (g gestureStylus) GetDevice() gdk.Device {
	return WrapGesture(gextras.InternObject(g)).GetDevice()
}

func (g gestureStylus) GetLastEvent(sequence *gdk.EventSequence) gdk.Event {
	return WrapGesture(gextras.InternObject(g)).GetLastEvent(sequence)
}

func (g gestureStylus) GetLastUpdatedSequence() *gdk.EventSequence {
	return WrapGesture(gextras.InternObject(g)).GetLastUpdatedSequence()
}

func (g gestureStylus) GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetPoint(sequence)
}

func (g gestureStylus) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	return WrapGesture(gextras.InternObject(g)).GetSequenceState(sequence)
}

func (g gestureStylus) Group(gesture Gesture) {
	WrapGesture(gextras.InternObject(g)).Group(gesture)
}

func (g gestureStylus) HandlesSequence(sequence *gdk.EventSequence) bool {
	return WrapGesture(gextras.InternObject(g)).HandlesSequence(sequence)
}

func (g gestureStylus) IsActive() bool {
	return WrapGesture(gextras.InternObject(g)).IsActive()
}

func (g gestureStylus) IsGroupedWith(other Gesture) bool {
	return WrapGesture(gextras.InternObject(g)).IsGroupedWith(other)
}

func (g gestureStylus) IsRecognized() bool {
	return WrapGesture(gextras.InternObject(g)).IsRecognized()
}

func (g gestureStylus) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetSequenceState(sequence, state)
}

func (g gestureStylus) SetState(state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetState(state)
}

func (g gestureStylus) Ungroup() {
	WrapGesture(gextras.InternObject(g)).Ungroup()
}

func (c gestureStylus) GetCurrentEvent() gdk.Event {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEvent()
}

func (c gestureStylus) GetCurrentEventDevice() gdk.Device {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventDevice()
}

func (c gestureStylus) GetCurrentEventState() gdk.ModifierType {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventState()
}

func (c gestureStylus) GetCurrentEventTime() uint32 {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventTime()
}

func (c gestureStylus) GetName() string {
	return WrapEventController(gextras.InternObject(c)).GetName()
}

func (c gestureStylus) GetPropagationLimit() PropagationLimit {
	return WrapEventController(gextras.InternObject(c)).GetPropagationLimit()
}

func (c gestureStylus) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c gestureStylus) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c gestureStylus) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c gestureStylus) SetName(name string) {
	WrapEventController(gextras.InternObject(c)).SetName(name)
}

func (c gestureStylus) SetPropagationLimit(limit PropagationLimit) {
	WrapEventController(gextras.InternObject(c)).SetPropagationLimit(limit)
}

func (c gestureStylus) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (g gestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.double            // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))
	_arg1 = C.GdkAxisUse(axis)

	_cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

func (g gestureStylus) Backlog() ([]gdk.TimeCoord, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 *C.GdkTimeCoord
	var _arg2 C.guint    // in
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_stylus_get_backlog(_arg0, &_arg1, &_arg2)

	var _backlog []gdk.TimeCoord
	var _ok bool // out

	_backlog = unsafe.Slice((*gdk.TimeCoord)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_backlog, func(v *[]gdk.TimeCoord) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _backlog, _ok
}

func (g gestureStylus) DeviceTool() gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)

	var _deviceTool gdk.DeviceTool // out

	_deviceTool = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.DeviceTool)

	return _deviceTool
}
