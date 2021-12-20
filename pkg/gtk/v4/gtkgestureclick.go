// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_click_get_type()), F: marshalGestureClicker},
	})
}

// GestureClick: GtkGestureClick is a GtkGesture implementation for clicks.
//
// It is able to recognize multiple clicks on a nearby zone, which can be
// listened for through the gtk.GestureClick::pressed signal. Whenever time or
// distance between clicks exceed the GTK defaults, gtk.GestureClick::stopped is
// emitted, and the click counter is reset.
type GestureClick struct {
	GestureSingle

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Gesturer = (*GestureClick)(nil)
)

func wrapGestureClick(obj *externglib.Object) *GestureClick {
	return &GestureClick{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureClicker(p uintptr) (interface{}, error) {
	return wrapGestureClick(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectPressed: emitted whenever a button or touch press happens.
func (v *GestureClick) ConnectPressed(f func(nPress int, x, y float64)) externglib.SignalHandle {
	return v.Connect("pressed", f)
}

// ConnectReleased: emitted when a button or touch is released.
//
// n_press will report the number of press that is paired to this event, note
// that gtk.GestureClick::stopped may have been emitted between the press and
// its release, n_press will only start over at the next press.
func (v *GestureClick) ConnectReleased(f func(nPress int, x, y float64)) externglib.SignalHandle {
	return v.Connect("released", f)
}

// ConnectStopped: emitted whenever any time/distance threshold has been
// exceeded.
func (v *GestureClick) ConnectStopped(f func()) externglib.SignalHandle {
	return v.Connect("stopped", f)
}

// ConnectUnpairedRelease: emitted whenever the gesture receives a release event
// that had no previous corresponding press.
//
// Due to implicit grabs, this can only happen on situations where input is
// grabbed elsewhere mid-press or the pressed widget voluntarily relinquishes
// its implicit grab.
func (v *GestureClick) ConnectUnpairedRelease(f func(x, y float64, button uint, sequence *gdk.EventSequence)) externglib.SignalHandle {
	return v.Connect("unpaired-release", f)
}

// NewGestureClick returns a newly created GtkGesture that recognizes single and
// multiple presses.
//
// The function returns the following values:
//
//    - gestureClick: newly created GtkGestureClick.
//
func NewGestureClick() *GestureClick {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_click_new()

	var _gestureClick *GestureClick // out

	_gestureClick = wrapGestureClick(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureClick
}
