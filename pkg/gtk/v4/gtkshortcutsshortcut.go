// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
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
	GTypeShortcutType      = coreglib.Type(girepository.MustFind("Gtk", "ShortcutType").RegisteredGType())
	GTypeShortcutsShortcut = coreglib.Type(girepository.MustFind("Gtk", "ShortcutsShortcut").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutType, F: marshalShortcutType},
		coreglib.TypeMarshaler{T: GTypeShortcutsShortcut, F: marshalShortcutsShortcut},
	})
}

// ShortcutType specifies the kind of shortcut that is being described.
//
// More values may be added to this enumeration over time.
type ShortcutType C.gint

const (
	// ShortcutAccelerator: shortcut is a keyboard accelerator. The
	// GtkShortcutsShortcut:accelerator property will be used.
	ShortcutAccelerator ShortcutType = iota
	// ShortcutGesturePinch: shortcut is a pinch gesture. GTK provides an icon
	// and subtitle.
	ShortcutGesturePinch
	// ShortcutGestureStretch: shortcut is a stretch gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureStretch
	// ShortcutGestureRotateClockwise: shortcut is a clockwise rotation gesture.
	// GTK provides an icon and subtitle.
	ShortcutGestureRotateClockwise
	// ShortcutGestureRotateCounterclockwise: shortcut is a counterclockwise
	// rotation gesture. GTK provides an icon and subtitle.
	ShortcutGestureRotateCounterclockwise
	// ShortcutGestureTwoFingerSwipeLeft: shortcut is a two-finger swipe
	// gesture. GTK provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeLeft
	// ShortcutGestureTwoFingerSwipeRight: shortcut is a two-finger swipe
	// gesture. GTK provides an icon and subtitle.
	ShortcutGestureTwoFingerSwipeRight
	// ShortcutGesture: shortcut is a gesture. The GtkShortcutsShortcut:icon
	// property will be used.
	ShortcutGesture
	// ShortcutGestureSwipeLeft: shortcut is a swipe gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureSwipeLeft
	// ShortcutGestureSwipeRight: shortcut is a swipe gesture. GTK provides an
	// icon and subtitle.
	ShortcutGestureSwipeRight
)

func marshalShortcutType(p uintptr) (interface{}, error) {
	return ShortcutType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ShortcutType.
func (s ShortcutType) String() string {
	switch s {
	case ShortcutAccelerator:
		return "Accelerator"
	case ShortcutGesturePinch:
		return "GesturePinch"
	case ShortcutGestureStretch:
		return "GestureStretch"
	case ShortcutGestureRotateClockwise:
		return "GestureRotateClockwise"
	case ShortcutGestureRotateCounterclockwise:
		return "GestureRotateCounterclockwise"
	case ShortcutGestureTwoFingerSwipeLeft:
		return "GestureTwoFingerSwipeLeft"
	case ShortcutGestureTwoFingerSwipeRight:
		return "GestureTwoFingerSwipeRight"
	case ShortcutGesture:
		return "Gesture"
	case ShortcutGestureSwipeLeft:
		return "GestureSwipeLeft"
	case ShortcutGestureSwipeRight:
		return "GestureSwipeRight"
	default:
		return fmt.Sprintf("ShortcutType(%d)", s)
	}
}

// ShortcutsShortcut: GtkShortcutsShortcut represents a single keyboard shortcut
// or gesture with a short text.
//
// This widget is only meant to be used with GtkShortcutsWindow.
type ShortcutsShortcut struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ShortcutsShortcut)(nil)
)

func wrapShortcutsShortcut(obj *coreglib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsShortcut(p uintptr) (interface{}, error) {
	return wrapShortcutsShortcut(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
