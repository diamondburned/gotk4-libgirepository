// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_EventControllerScroll_ConnectDecelerate(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_EventControllerScroll_ConnectScroll(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_EventControllerScroll_ConnectScrollBegin(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerScroll_ConnectScrollEnd(gpointer, guintptr);
import "C"

// glib.Type values for gtkeventcontrollerscroll.go.
var (
	GTypeEventControllerScrollFlags = coreglib.Type(C.gtk_event_controller_scroll_flags_get_type())
	GTypeEventControllerScroll      = coreglib.Type(C.gtk_event_controller_scroll_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEventControllerScrollFlags, F: marshalEventControllerScrollFlags},
		{T: GTypeEventControllerScroll, F: marshalEventControllerScroll},
	})
}

// EventControllerScrollFlags describes the behavior of a EventControllerScroll.
type EventControllerScrollFlags C.guint

const (
	// EventControllerScrollNone: don't emit scroll.
	EventControllerScrollNone EventControllerScrollFlags = 0b0
	// EventControllerScrollVertical: emit scroll with vertical deltas.
	EventControllerScrollVertical EventControllerScrollFlags = 0b1
	// EventControllerScrollHorizontal: emit scroll with horizontal deltas.
	EventControllerScrollHorizontal EventControllerScrollFlags = 0b10
	// EventControllerScrollDiscrete: only emit deltas that are multiples of 1.
	EventControllerScrollDiscrete EventControllerScrollFlags = 0b100
	// EventControllerScrollKinetic: emit EventControllerScroll::decelerate
	// after continuous scroll finishes.
	EventControllerScrollKinetic EventControllerScrollFlags = 0b1000
	// EventControllerScrollBothAxes: emit scroll on both axes.
	EventControllerScrollBothAxes EventControllerScrollFlags = 0b11
)

func marshalEventControllerScrollFlags(p uintptr) (interface{}, error) {
	return EventControllerScrollFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for EventControllerScrollFlags.
func (e EventControllerScrollFlags) String() string {
	if e == 0 {
		return "EventControllerScrollFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(176)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case EventControllerScrollNone:
			builder.WriteString("None|")
		case EventControllerScrollVertical:
			builder.WriteString("Vertical|")
		case EventControllerScrollHorizontal:
			builder.WriteString("Horizontal|")
		case EventControllerScrollDiscrete:
			builder.WriteString("Discrete|")
		case EventControllerScrollKinetic:
			builder.WriteString("Kinetic|")
		case EventControllerScrollBothAxes:
			builder.WriteString("BothAxes|")
		default:
			builder.WriteString(fmt.Sprintf("EventControllerScrollFlags(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if e contains other.
func (e EventControllerScrollFlags) Has(other EventControllerScrollFlags) bool {
	return (e & other) == other
}

// EventControllerScroll is an event controller meant to handle scroll events
// from mice and touchpads. It is capable of handling both discrete and
// continuous scroll events, abstracting them both on the
// EventControllerScroll::scroll signal (deltas in the discrete case are
// multiples of 1).
//
// In the case of continuous scroll events, EventControllerScroll encloses all
// EventControllerScroll::scroll events between two
// EventControllerScroll::scroll-begin and EventControllerScroll::scroll-end
// signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// gtk_event_controller_scroll_set_flags() (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through K_EVENT_CONTROLLER_SCROLL_VERTICAL,
// K_EVENT_CONTROLLER_SCROLL_HORIZONTAL and K_EVENT_CONTROLLER_SCROLL_BOTH. If
// any axis is disabled, the respective EventControllerScroll::scroll delta will
// be 0. Vertical scroll events will be translated to horizontal motion for the
// devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through K_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The K_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// EventControllerScroll::decelerate signal, emitted at the end of scrolling
// with two X/Y velocity arguments that are consistent with the motion that was
// received.
//
// This object was added in 3.24.
type EventControllerScroll struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerScroll)(nil)
)

func wrapEventControllerScroll(obj *coreglib.Object) *EventControllerScroll {
	return &EventControllerScroll{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerScroll(p uintptr) (interface{}, error) {
	return wrapEventControllerScroll(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_EventControllerScroll_ConnectDecelerate
func _gotk4_gtk3_EventControllerScroll_ConnectDecelerate(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velX, velY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velX, velY float64))
	}

	var _velX float64 // out
	var _velY float64 // out

	_velX = float64(arg1)
	_velY = float64(arg2)

	f(_velX, _velY)
}

// ConnectDecelerate is emitted after scroll is finished if the
// K_EVENT_CONTROLLER_SCROLL_KINETIC flag is set. vel_x and vel_y express the
// initial velocity that was imprinted by the scroll events. vel_x and vel_y are
// expressed in pixels/ms.
func (v *EventControllerScroll) ConnectDecelerate(f func(velX, velY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "decelerate", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerScroll_ConnectDecelerate), f)
}

//export _gotk4_gtk3_EventControllerScroll_ConnectScroll
func _gotk4_gtk3_EventControllerScroll_ConnectScroll(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(dx, dy float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(dx, dy float64))
	}

	var _dx float64 // out
	var _dy float64 // out

	_dx = float64(arg1)
	_dy = float64(arg2)

	f(_dx, _dy)
}

// ConnectScroll signals that the widget should scroll by the amount specified
// by dx and dy.
func (v *EventControllerScroll) ConnectScroll(f func(dx, dy float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "scroll", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerScroll_ConnectScroll), f)
}

//export _gotk4_gtk3_EventControllerScroll_ConnectScrollBegin
func _gotk4_gtk3_EventControllerScroll_ConnectScrollBegin(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectScrollBegin signals that a new scrolling operation has begun. It will
// only be emitted on devices capable of it.
func (v *EventControllerScroll) ConnectScrollBegin(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "scroll-begin", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerScroll_ConnectScrollBegin), f)
}

//export _gotk4_gtk3_EventControllerScroll_ConnectScrollEnd
func _gotk4_gtk3_EventControllerScroll_ConnectScrollEnd(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectScrollEnd signals that a new scrolling operation has finished. It will
// only be emitted on devices capable of it.
func (v *EventControllerScroll) ConnectScrollEnd(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "scroll-end", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerScroll_ConnectScrollEnd), f)
}
