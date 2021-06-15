// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_get_type()), F: marshalGesture},
	})
}

// Gesture is the base object for gesture recognition, although this object is
// quite generalized to serve as a base for multi-touch gestures, it is suitable
// to implement single-touch and pointer-based gestures (using the special nil
// EventSequence value for these).
//
// The number of touches that a Gesture need to be recognized is controlled by
// the Gesture:n-points property, if a gesture is keeping track of less or more
// than that number of sequences, it won't check wether the gesture is
// recognized.
//
// As soon as the gesture has the expected number of touches, the gesture will
// run the Gesture::check signal regularly on input events until the gesture is
// recognized, the criteria to consider a gesture as "recognized" is left to
// Gesture subclasses.
//
// A recognized gesture will then emit the following signals: - Gesture::begin
// when the gesture is recognized. - A number of Gesture::update, whenever an
// input event is processed. - Gesture::end when the gesture is no longer
// recognized.
//
//
// Event propagation
//
// In order to receive events, a gesture needs to either set a propagation phase
// through gtk_event_controller_set_propagation_phase(), or feed those manually
// through gtk_event_controller_handle_event().
//
// In the capture phase, events are propagated from the toplevel down to the
// target widget, and gestures that are attached to containers above the widget
// get a chance to interact with the event before it reaches the target.
//
// After the capture phase, GTK+ emits the traditional
// Widget::button-press-event, Widget::button-release-event,
// Widget::touch-event, etc signals. Gestures with the GTK_PHASE_TARGET phase
// are fed events from the default Widget::event handlers.
//
// In the bubble phase, events are propagated up from the target widget to the
// toplevel, and gestures that are attached to containers above the widget get a
// chance to interact with events that have not been handled yet.
//
//
// States of a sequence
//
// Whenever input interaction happens, a single event may trigger a cascade of
// Gestures, both across the parents of the widget receiving the event and in
// parallel within an individual widget. It is a responsibility of the widgets
// using those gestures to set the state of touch sequences accordingly in order
// to enable cooperation of gestures around the EventSequences triggering those.
//
// Within a widget, gestures can be grouped through gtk_gesture_group(), grouped
// gestures synchronize the state of sequences, so calling
// gtk_gesture_set_sequence_state() on one will effectively propagate the state
// throughout the group.
//
// By default, all sequences start out in the K_EVENT_SEQUENCE_NONE state,
// sequences in this state trigger the gesture event handler, but event
// propagation will continue unstopped by gestures.
//
// If a sequence enters into the K_EVENT_SEQUENCE_DENIED state, the gesture
// group will effectively ignore the sequence, letting events go unstopped
// through the gesture, but the "slot" will still remain occupied while the
// touch is active.
//
// If a sequence enters in the K_EVENT_SEQUENCE_CLAIMED state, the gesture group
// will grab all interaction on the sequence, by: - Setting the same sequence to
// K_EVENT_SEQUENCE_DENIED on every other gesture group within the widget, and
// every gesture on parent widgets in the propagation chain. - calling
// Gesture::cancel on every gesture in widgets underneath in the propagation
// chain. - Stopping event propagation after the gesture group handles the
// event.
//
// Note: if a sequence is set early to K_EVENT_SEQUENCE_CLAIMED on
// K_TOUCH_BEGIN/K_BUTTON_PRESS (so those events are captured before reaching
// the event widget, this implies K_PHASE_CAPTURE), one similar event will
// emulated if the sequence changes to K_EVENT_SEQUENCE_DENIED. This way event
// coherence is preserved before event propagation is unstopped again.
//
// Sequence states can't be changed freely, see gtk_gesture_set_sequence_state()
// to know about the possible lifetimes of a EventSequence.
//
//
// Touchpad gestures
//
// On the platforms that support it, Gesture will handle transparently touchpad
// gesture events. The only precautions users of Gesture should do to enable
// this support are: - Enabling GDK_TOUCHPAD_GESTURE_MASK on their Windows - If
// the gesture has GTK_PHASE_NONE, ensuring events of type GDK_TOUCHPAD_SWIPE
// and GDK_TOUCHPAD_PINCH are handled by the Gesture
type Gesture interface {
	EventController

	// BoundingBox: if there are touch sequences being currently handled by
	// @gesture, this function returns true and fills in @rect with the bounding
	// box containing all active touches. Otherwise, false will be returned.
	//
	// Note: This function will yield unexpected results on touchpad gestures.
	// Since there is no correlation between physical and pixel distances, these
	// will look as if constrained in an infinitely small area, @rect width and
	// height will thus be 0 regardless of the number of touchpoints.
	BoundingBox() (gdk.Rectangle, bool)
	// BoundingBoxCenter: if there are touch sequences being currently handled
	// by @gesture, this function returns true and fills in @x and @y with the
	// center of the bounding box containing all active touches. Otherwise,
	// false will be returned.
	BoundingBoxCenter() (x float64, y float64, ok bool)
	// Device returns the master Device that is currently operating on @gesture,
	// or nil if the gesture is not being interacted.
	Device() gdk.Device
	// LastUpdatedSequence returns the EventSequence that was last updated on
	// @gesture.
	LastUpdatedSequence() *gdk.EventSequence
	// Point: if @sequence is currently being interpreted by @gesture, this
	// function returns true and fills in @x and @y with the last coordinates
	// stored for that event sequence. The coordinates are always relative to
	// the widget allocation.
	Point(sequence *gdk.EventSequence) (x float64, y float64, ok bool)
	// SequenceState returns the @sequence state, as seen by @gesture.
	SequenceState(sequence *gdk.EventSequence) EventSequenceState
	// Window returns the user-defined window that receives the events handled
	// by @gesture. See gtk_gesture_set_window() for more information.
	Window() gdk.Window
	// Group adds @gesture to the same group than @group_gesture. Gestures are
	// by default isolated in their own groups.
	//
	// When gestures are grouped, the state of EventSequences is kept in sync
	// for all of those, so calling gtk_gesture_set_sequence_state(), on one
	// will transfer the same value to the others.
	//
	// Groups also perform an "implicit grabbing" of sequences, if a
	// EventSequence state is set to K_EVENT_SEQUENCE_CLAIMED on one group,
	// every other gesture group attached to the same Widget will switch the
	// state for that sequence to K_EVENT_SEQUENCE_DENIED.
	Group(gesture Gesture)
	// HandlesSequence returns true if @gesture is currently handling events
	// corresponding to @sequence.
	HandlesSequence(sequence *gdk.EventSequence) bool
	// IsActive returns true if the gesture is currently active. A gesture is
	// active meanwhile there are touch sequences interacting with it.
	IsActive() bool
	// IsGroupedWith returns true if both gestures pertain to the same group.
	IsGroupedWith(other Gesture) bool
	// IsRecognized returns true if the gesture is currently recognized. A
	// gesture is recognized if there are as many interacting touch sequences as
	// required by @gesture, and Gesture::check returned true for the sequences
	// being currently interpreted.
	IsRecognized() bool
	// SetSequenceState sets the state of @sequence in @gesture. Sequences start
	// in state K_EVENT_SEQUENCE_NONE, and whenever they change state, they can
	// never go back to that state. Likewise, sequences in state
	// K_EVENT_SEQUENCE_DENIED cannot turn back to a not denied state. With
	// these rules, the lifetime of an event sequence is constrained to the next
	// four:
	//
	// * None * None → Denied * None → Claimed * None → Claimed → Denied
	//
	// Note: Due to event handling ordering, it may be unsafe to set the state
	// on another gesture within a Gesture::begin signal handler, as the
	// callback might be executed before the other gesture knows about the
	// sequence. A safe way to perform this could be:
	//
	//    static void
	//    first_gesture_begin_cb (GtkGesture       *first_gesture,
	//                            GdkEventSequence *sequence,
	//                            gpointer          user_data)
	//    {
	//      gtk_gesture_set_sequence_state (first_gesture, sequence, GTK_EVENT_SEQUENCE_CLAIMED);
	//      gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
	//    }
	//
	//    static void
	//    second_gesture_begin_cb (GtkGesture       *second_gesture,
	//                             GdkEventSequence *sequence,
	//                             gpointer          user_data)
	//    {
	//      if (gtk_gesture_get_sequence_state (first_gesture, sequence) == GTK_EVENT_SEQUENCE_CLAIMED)
	//        gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
	//    }
	//
	// If both gestures are in the same group, just set the state on the gesture
	// emitting the event, the sequence will be already be initialized to the
	// group's global state when the second gesture processes the event.
	SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool
	// SetState sets the state of all sequences that @gesture is currently
	// interacting with. See gtk_gesture_set_sequence_state() for more details
	// on sequence states.
	SetState(state EventSequenceState) bool
	// SetWindow sets a specific window to receive events about, so @gesture
	// will effectively handle only events targeting @window, or a child of it.
	// @window must pertain to gtk_event_controller_get_widget().
	SetWindow(window gdk.Window)
	// Ungroup separates @gesture into an isolated group.
	Ungroup()
}

// gesture implements the Gesture class.
type gesture struct {
	EventController
}

var _ Gesture = (*gesture)(nil)

// WrapGesture wraps a GObject to the right type. It is
// primarily used internally.
func WrapGesture(obj *externglib.Object) Gesture {
	return gesture{
		EventController: WrapEventController(obj),
	}
}

func marshalGesture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGesture(obj), nil
}

// BoundingBox: if there are touch sequences being currently handled by
// @gesture, this function returns true and fills in @rect with the bounding
// box containing all active touches. Otherwise, false will be returned.
//
// Note: This function will yield unexpected results on touchpad gestures.
// Since there is no correlation between physical and pixel distances, these
// will look as if constrained in an infinitely small area, @rect width and
// height will thus be 0 regardless of the number of touchpoints.
func (g gesture) BoundingBox() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkGesture // out
	var _rect gdk.Rectangle
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_get_bounding_box(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_rect)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// BoundingBoxCenter: if there are touch sequences being currently handled
// by @gesture, this function returns true and fills in @x and @y with the
// center of the bounding box containing all active touches. Otherwise,
// false will be returned.
func (g gesture) BoundingBoxCenter() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture // out
	var _arg1 C.gdouble     // in
	var _arg2 C.gdouble     // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_get_bounding_box_center(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = (float64)(_arg1)
	_y = (float64)(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// Device returns the master Device that is currently operating on @gesture,
// or nil if the gesture is not being interacted.
func (g gesture) Device() gdk.Device {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GdkDevice  // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_get_device(_arg0)

	var _device gdk.Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Device)

	return _device
}

// LastUpdatedSequence returns the EventSequence that was last updated on
// @gesture.
func (g gesture) LastUpdatedSequence() *gdk.EventSequence {
	var _arg0 *C.GtkGesture       // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_get_last_updated_sequence(_arg0)

	var _eventSequence *gdk.EventSequence // out

	_eventSequence = gdk.WrapEventSequence(unsafe.Pointer(_cret))

	return _eventSequence
}

// Point: if @sequence is currently being interpreted by @gesture, this
// function returns true and fills in @x and @y with the last coordinates
// stored for that event sequence. The coordinates are always relative to
// the widget allocation.
func (g gesture) Point(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _arg2 C.gdouble           // in
	var _arg3 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkEventSequence)(unsafe.Pointer(sequence.Native()))

	_cret = C.gtk_gesture_get_point(_arg0, _arg1, &_arg2, &_arg3)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = (float64)(_arg2)
	_y = (float64)(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// SequenceState returns the @sequence state, as seen by @gesture.
func (g gesture) SequenceState(sequence *gdk.EventSequence) EventSequenceState {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _cret C.GtkEventSequenceState // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkEventSequence)(unsafe.Pointer(sequence.Native()))

	_cret = C.gtk_gesture_get_sequence_state(_arg0, _arg1)

	var _eventSequenceState EventSequenceState // out

	_eventSequenceState = EventSequenceState(_cret)

	return _eventSequenceState
}

// Window returns the user-defined window that receives the events handled
// by @gesture. See gtk_gesture_set_window() for more information.
func (g gesture) Window() gdk.Window {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_get_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

// Group adds @gesture to the same group than @group_gesture. Gestures are
// by default isolated in their own groups.
//
// When gestures are grouped, the state of EventSequences is kept in sync
// for all of those, so calling gtk_gesture_set_sequence_state(), on one
// will transfer the same value to the others.
//
// Groups also perform an "implicit grabbing" of sequences, if a
// EventSequence state is set to K_EVENT_SEQUENCE_CLAIMED on one group,
// every other gesture group attached to the same Widget will switch the
// state for that sequence to K_EVENT_SEQUENCE_DENIED.
func (g gesture) Group(gesture Gesture) {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	C.gtk_gesture_group(_arg0, _arg1)
}

// HandlesSequence returns true if @gesture is currently handling events
// corresponding to @sequence.
func (g gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkEventSequence)(unsafe.Pointer(sequence.Native()))

	_cret = C.gtk_gesture_handles_sequence(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActive returns true if the gesture is currently active. A gesture is
// active meanwhile there are touch sequences interacting with it.
func (g gesture) IsActive() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_is_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupedWith returns true if both gestures pertain to the same group.
func (g gesture) IsGroupedWith(other Gesture) bool {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(other.Native()))

	_cret = C.gtk_gesture_is_grouped_with(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRecognized returns true if the gesture is currently recognized. A
// gesture is recognized if there are as many interacting touch sequences as
// required by @gesture, and Gesture::check returned true for the sequences
// being currently interpreted.
func (g gesture) IsRecognized() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_is_recognized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSequenceState sets the state of @sequence in @gesture. Sequences start
// in state K_EVENT_SEQUENCE_NONE, and whenever they change state, they can
// never go back to that state. Likewise, sequences in state
// K_EVENT_SEQUENCE_DENIED cannot turn back to a not denied state. With
// these rules, the lifetime of an event sequence is constrained to the next
// four:
//
// * None * None → Denied * None → Claimed * None → Claimed → Denied
//
// Note: Due to event handling ordering, it may be unsafe to set the state
// on another gesture within a Gesture::begin signal handler, as the
// callback might be executed before the other gesture knows about the
// sequence. A safe way to perform this could be:
//
//    static void
//    first_gesture_begin_cb (GtkGesture       *first_gesture,
//                            GdkEventSequence *sequence,
//                            gpointer          user_data)
//    {
//      gtk_gesture_set_sequence_state (first_gesture, sequence, GTK_EVENT_SEQUENCE_CLAIMED);
//      gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
//    static void
//    second_gesture_begin_cb (GtkGesture       *second_gesture,
//                             GdkEventSequence *sequence,
//                             gpointer          user_data)
//    {
//      if (gtk_gesture_get_sequence_state (first_gesture, sequence) == GTK_EVENT_SEQUENCE_CLAIMED)
//        gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
// If both gestures are in the same group, just set the state on the gesture
// emitting the event, the sequence will be already be initialized to the
// group's global state when the second gesture processes the event.
func (g gesture) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _arg2 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkEventSequence)(unsafe.Pointer(sequence.Native()))
	_arg2 = (C.GtkEventSequenceState)(state)

	_cret = C.gtk_gesture_set_sequence_state(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetState sets the state of all sequences that @gesture is currently
// interacting with. See gtk_gesture_set_sequence_state() for more details
// on sequence states.
func (g gesture) SetState(state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (C.GtkEventSequenceState)(state)

	_cret = C.gtk_gesture_set_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetWindow sets a specific window to receive events about, so @gesture
// will effectively handle only events targeting @window, or a child of it.
// @window must pertain to gtk_event_controller_get_widget().
func (g gesture) SetWindow(window gdk.Window) {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GdkWindow  // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_gesture_set_window(_arg0, _arg1)
}

// Ungroup separates @gesture into an isolated group.
func (g gesture) Ungroup() {
	var _arg0 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(g.Native()))

	C.gtk_gesture_ungroup(_arg0)
}
