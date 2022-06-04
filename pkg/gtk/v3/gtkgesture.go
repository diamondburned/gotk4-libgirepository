// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_Gesture_ConnectBegin(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_Gesture_ConnectCancel(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_Gesture_ConnectEnd(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_Gesture_ConnectUpdate(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtkgesture.go.
var GTypeGesture = coreglib.Type(C.gtk_gesture_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeGesture, F: marshalGesture},
	})
}

// GestureOverrider contains methods that are overridable.
type GestureOverrider interface {
}

// Gesture is the base object for gesture recognition, although this object is
// quite generalized to serve as a base for multi-touch gestures, it is suitable
// to implement single-touch and pointer-based gestures (using the special NULL
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
// A recognized gesture will then emit the following signals:
//
// - Gesture::begin when the gesture is recognized.
//
// - A number of Gesture::update, whenever an input event is processed.
//
// - Gesture::end when the gesture is no longer recognized.
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
// will grab all interaction on the sequence, by:
//
// - Setting the same sequence to K_EVENT_SEQUENCE_DENIED on every other gesture
// group within the widget, and every gesture on parent widgets in the
// propagation chain.
//
// - calling Gesture::cancel on every gesture in widgets underneath in the
// propagation chain.
//
// - Stopping event propagation after the gesture group handles the event.
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
// this support are:
//
// - Enabling GDK_TOUCHPAD_GESTURE_MASK on their Windows
//
// - If the gesture has GTK_PHASE_NONE, ensuring events of type
// GDK_TOUCHPAD_SWIPE and GDK_TOUCHPAD_PINCH are handled by the Gesture.
type Gesture struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*Gesture)(nil)
)

// Gesturer describes types inherited from class Gesture.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Gesturer interface {
	coreglib.Objector
	baseGesture() *Gesture
}

var _ Gesturer = (*Gesture)(nil)

func classInitGesturer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGesture(obj *coreglib.Object) *Gesture {
	return &Gesture{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalGesture(p uintptr) (interface{}, error) {
	return wrapGesture(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (gesture *Gesture) baseGesture() *Gesture {
	return gesture
}

// BaseGesture returns the underlying base object.
func BaseGesture(obj Gesturer) *Gesture {
	return obj.baseGesture()
}

//export _gotk4_gtk3_Gesture_ConnectBegin
func _gotk4_gtk3_Gesture_ConnectBegin(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(sequence *gdk.EventSequence)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(sequence *gdk.EventSequence))
	}

	var _sequence *gdk.EventSequence // out

	if arg1 != nil {
		_sequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	}

	f(_sequence)
}

// ConnectBegin: this signal is emitted when the gesture is recognized. This
// means the number of touch sequences matches Gesture:n-points, and the
// Gesture::check handler(s) returned UE.
//
// Note: These conditions may also happen when an extra touch (eg. a third touch
// on a 2-touches gesture) is lifted, in that situation sequence won't pertain
// to the current set of active touches, so don't rely on this being true.
func (gesture *Gesture) ConnectBegin(f func(sequence *gdk.EventSequence)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "begin", false, unsafe.Pointer(C._gotk4_gtk3_Gesture_ConnectBegin), f)
}

//export _gotk4_gtk3_Gesture_ConnectCancel
func _gotk4_gtk3_Gesture_ConnectCancel(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(sequence *gdk.EventSequence)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(sequence *gdk.EventSequence))
	}

	var _sequence *gdk.EventSequence // out

	if arg1 != nil {
		_sequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	}

	f(_sequence)
}

// ConnectCancel: this signal is emitted whenever a sequence is cancelled. This
// usually happens on active touches when gtk_event_controller_reset() is called
// on gesture (manually, due to grabs...), or the individual sequence was
// claimed by parent widgets' controllers (see
// gtk_gesture_set_sequence_state()).
//
// gesture must forget everything about sequence as a reaction to this signal.
func (gesture *Gesture) ConnectCancel(f func(sequence *gdk.EventSequence)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "cancel", false, unsafe.Pointer(C._gotk4_gtk3_Gesture_ConnectCancel), f)
}

//export _gotk4_gtk3_Gesture_ConnectEnd
func _gotk4_gtk3_Gesture_ConnectEnd(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(sequence *gdk.EventSequence)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(sequence *gdk.EventSequence))
	}

	var _sequence *gdk.EventSequence // out

	if arg1 != nil {
		_sequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	}

	f(_sequence)
}

// ConnectEnd: this signal is emitted when gesture either stopped recognizing
// the event sequences as something to be handled (the Gesture::check handler
// returned FALSE), or the number of touch sequences became higher or lower than
// Gesture:n-points.
//
// Note: sequence might not pertain to the group of sequences that were
// previously triggering recognition on gesture (ie. a just pressed touch
// sequence that exceeds Gesture:n-points). This situation may be detected by
// checking through gtk_gesture_handles_sequence().
func (gesture *Gesture) ConnectEnd(f func(sequence *gdk.EventSequence)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "end", false, unsafe.Pointer(C._gotk4_gtk3_Gesture_ConnectEnd), f)
}

//export _gotk4_gtk3_Gesture_ConnectUpdate
func _gotk4_gtk3_Gesture_ConnectUpdate(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(sequence *gdk.EventSequence)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(sequence *gdk.EventSequence))
	}

	var _sequence *gdk.EventSequence // out

	if arg1 != nil {
		_sequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	}

	f(_sequence)
}

// ConnectUpdate: this signal is emitted whenever an event is handled while the
// gesture is recognized. sequence is guaranteed to pertain to the set of active
// touches.
func (gesture *Gesture) ConnectUpdate(f func(sequence *gdk.EventSequence)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "update", false, unsafe.Pointer(C._gotk4_gtk3_Gesture_ConnectUpdate), f)
}

// BoundingBox: if there are touch sequences being currently handled by gesture,
// this function returns TRUE and fills in rect with the bounding box containing
// all active touches. Otherwise, FALSE will be returned.
//
// Note: This function will yield unexpected results on touchpad gestures. Since
// there is no correlation between physical and pixel distances, these will look
// as if constrained in an infinitely small area, rect width and height will
// thus be 0 regardless of the number of touchpoints.
//
// The function returns the following values:
//
//    - rect: bounding box containing all active touches.
//    - ok: TRUE if there are active touches, FALSE otherwise.
//
func (gesture *Gesture) BoundingBox() (*gdk.Rectangle, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_bounding_box", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _rect, _ok
}

// BoundingBoxCenter: if there are touch sequences being currently handled by
// gesture, this function returns TRUE and fills in x and y with the center of
// the bounding box containing all active touches. Otherwise, FALSE will be
// returned.
//
// The function returns the following values:
//
//    - x: x coordinate for the bounding box center.
//    - y: y coordinate for the bounding box center.
//    - ok: FALSE if no active touches are present, TRUE otherwise.
//
func (gesture *Gesture) BoundingBoxCenter() (x, y float64, ok bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_bounding_box_center", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = *(*float64)(unsafe.Pointer(_outs[0]))
	_y = *(*float64)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// Device returns the master Device that is currently operating on gesture, or
// NULL if the gesture is not being interacted.
//
// The function returns the following values:
//
//    - device (optional) or NULL.
//
func (gesture *Gesture) Device() gdk.Devicer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_device", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _device gdk.Devicer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Devicer)
				return ok
			})
			rv, ok := casted.(gdk.Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// GetGroup returns all gestures in the group of gesture.
//
// The function returns the following values:
//
//    - list: list of Gestures, free with g_list_free().
//
func (gesture *Gesture) GetGroup() []Gesturer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_group", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _list []Gesturer // out

	_list = make([]Gesturer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst Gesturer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Gesturer is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Gesturer)
				return ok
			})
			rv, ok := casted.(Gesturer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Gesturer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// LastEvent returns the last event that was processed for sequence.
//
// Note that the returned pointer is only valid as long as the sequence is still
// interpreted by the gesture. If in doubt, you should make a copy of the event.
//
// The function takes the following parameters:
//
//    - sequence (optional): EventSequence.
//
// The function returns the following values:
//
//    - event (optional): last event from sequence.
//
func (gesture *Gesture) LastEvent(sequence *gdk.EventSequence) *gdk.Event {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if sequence != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_last_event", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _event *gdk.Event // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer(_cret)))
			v = gdk.CopyEventer(v)
			_event = v
		}
	}

	return _event
}

// LastUpdatedSequence returns the EventSequence that was last updated on
// gesture.
//
// The function returns the following values:
//
//    - eventSequence (optional): last updated sequence.
//
func (gesture *Gesture) LastUpdatedSequence() *gdk.EventSequence {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_last_updated_sequence", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _eventSequence *gdk.EventSequence // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_eventSequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _eventSequence
}

// Point: if sequence is currently being interpreted by gesture, this function
// returns TRUE and fills in x and y with the last coordinates stored for that
// event sequence. The coordinates are always relative to the widget allocation.
//
// The function takes the following parameters:
//
//    - sequence (optional) or NULL for pointer events.
//
// The function returns the following values:
//
//    - x (optional): return location for X axis of the sequence coordinates.
//    - y (optional): return location for Y axis of the sequence coordinates.
//    - ok: TRUE if sequence is currently interpreted.
//
func (gesture *Gesture) Point(sequence *gdk.EventSequence) (x, y float64, ok bool) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if sequence != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_point", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_x = *(*float64)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_y = *(*float64)(unsafe.Pointer(_outs[1]))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// Sequences returns the list of EventSequences currently being interpreted by
// gesture.
//
// The function returns the following values:
//
//    - list: list of EventSequences, the list elements are owned by GTK+ and
//      must not be freed or modified, the list itself must be deleted through
//      g_list_free().
//
func (gesture *Gesture) Sequences() []*gdk.EventSequence {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_sequences", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _list []*gdk.EventSequence // out

	_list = make([]*gdk.EventSequence, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *gdk.EventSequence // out
		dst = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Window returns the user-defined window that receives the events handled by
// gesture. See gtk_gesture_set_window() for more information.
//
// The function returns the following values:
//
//    - window (optional): user defined window, or NULL if none.
//
func (gesture *Gesture) Window() gdk.Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("get_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _window gdk.Windower // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Windower)
				return ok
			})
			rv, ok := casted.(gdk.Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// Group adds gesture to the same group than group_gesture. Gestures are by
// default isolated in their own groups.
//
// When gestures are grouped, the state of EventSequences is kept in sync for
// all of those, so calling gtk_gesture_set_sequence_state(), on one will
// transfer the same value to the others.
//
// Groups also perform an "implicit grabbing" of sequences, if a EventSequence
// state is set to K_EVENT_SEQUENCE_CLAIMED on one group, every other gesture
// group attached to the same Widget will switch the state for that sequence to
// K_EVENT_SEQUENCE_DENIED.
//
// The function takes the following parameters:
//
//    - gesture: Gesture.
//
func (groupGesture *Gesture) Group(gesture Gesturer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(groupGesture).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	girepository.MustFind("Gtk", "Gesture").InvokeMethod("group", _args[:], nil)

	runtime.KeepAlive(groupGesture)
	runtime.KeepAlive(gesture)
}

// HandlesSequence returns TRUE if gesture is currently handling events
// corresponding to sequence.
//
// The function takes the following parameters:
//
//    - sequence (optional) or NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if gesture is handling sequence, FALSE otherwise.
//
func (gesture *Gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if sequence != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("handles_sequence", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsActive returns TRUE if the gesture is currently active. A gesture is active
// meanwhile there are touch sequences interacting with it.
//
// The function returns the following values:
//
//    - ok: TRUE if gesture is active.
//
func (gesture *Gesture) IsActive() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("is_active", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupedWith returns TRUE if both gestures pertain to the same group.
//
// The function takes the following parameters:
//
//    - other Gesture.
//
// The function returns the following values:
//
//    - ok: whether the gestures are grouped.
//
func (gesture *Gesture) IsGroupedWith(other Gesturer) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("is_grouped_with", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(other)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsRecognized returns TRUE if the gesture is currently recognized. A gesture
// is recognized if there are as many interacting touch sequences as required by
// gesture, and Gesture::check returned TRUE for the sequences being currently
// interpreted.
//
// The function returns the following values:
//
//    - ok: TRUE if gesture is recognized.
//
func (gesture *Gesture) IsRecognized() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "Gesture").InvokeMethod("is_recognized", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetWindow sets a specific window to receive events about, so gesture will
// effectively handle only events targeting window, or a child of it. window
// must pertain to gtk_event_controller_get_widget().
//
// The function takes the following parameters:
//
//    - window (optional) or NULL.
//
func (gesture *Gesture) SetWindow(window gdk.Windower) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if window != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	}

	girepository.MustFind("Gtk", "Gesture").InvokeMethod("set_window", _args[:], nil)

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(window)
}

// Ungroup separates gesture into an isolated group.
func (gesture *Gesture) Ungroup() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	girepository.MustFind("Gtk", "Gesture").InvokeMethod("ungroup", _args[:], nil)

	runtime.KeepAlive(gesture)
}
