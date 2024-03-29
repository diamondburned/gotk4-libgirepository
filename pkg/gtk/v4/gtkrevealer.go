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
	GTypeRevealerTransitionType = coreglib.Type(girepository.MustFind("Gtk", "RevealerTransitionType").RegisteredGType())
	GTypeRevealer               = coreglib.Type(girepository.MustFind("Gtk", "Revealer").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRevealerTransitionType, F: marshalRevealerTransitionType},
		coreglib.TypeMarshaler{T: GTypeRevealer, F: marshalRevealer},
	})
}

// RevealerTransitionType: these enumeration values describe the possible
// transitions when the child of a GtkRevealer widget is shown or hidden.
type RevealerTransitionType C.gint

const (
	// RevealerTransitionTypeNone: no transition.
	RevealerTransitionTypeNone RevealerTransitionType = iota
	// RevealerTransitionTypeCrossfade: fade in.
	RevealerTransitionTypeCrossfade
	// RevealerTransitionTypeSlideRight: slide in from the left.
	RevealerTransitionTypeSlideRight
	// RevealerTransitionTypeSlideLeft: slide in from the right.
	RevealerTransitionTypeSlideLeft
	// RevealerTransitionTypeSlideUp: slide in from the bottom.
	RevealerTransitionTypeSlideUp
	// RevealerTransitionTypeSlideDown: slide in from the top.
	RevealerTransitionTypeSlideDown
	// RevealerTransitionTypeSwingRight: floop in from the left.
	RevealerTransitionTypeSwingRight
	// RevealerTransitionTypeSwingLeft: floop in from the right.
	RevealerTransitionTypeSwingLeft
	// RevealerTransitionTypeSwingUp: floop in from the bottom.
	RevealerTransitionTypeSwingUp
	// RevealerTransitionTypeSwingDown: floop in from the top.
	RevealerTransitionTypeSwingDown
)

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	return RevealerTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RevealerTransitionType.
func (r RevealerTransitionType) String() string {
	switch r {
	case RevealerTransitionTypeNone:
		return "None"
	case RevealerTransitionTypeCrossfade:
		return "Crossfade"
	case RevealerTransitionTypeSlideRight:
		return "SlideRight"
	case RevealerTransitionTypeSlideLeft:
		return "SlideLeft"
	case RevealerTransitionTypeSlideUp:
		return "SlideUp"
	case RevealerTransitionTypeSlideDown:
		return "SlideDown"
	case RevealerTransitionTypeSwingRight:
		return "SwingRight"
	case RevealerTransitionTypeSwingLeft:
		return "SwingLeft"
	case RevealerTransitionTypeSwingUp:
		return "SwingUp"
	case RevealerTransitionTypeSwingDown:
		return "SwingDown"
	default:
		return fmt.Sprintf("RevealerTransitionType(%d)", r)
	}
}

// Revealer: GtkRevealer animates the transition of its child from invisible to
// visible.
//
// The style of transition can be controlled with
// gtk.Revealer.SetTransitionType().
//
// These animations respect the gtk.Settings:gtk-enable-animations setting.
//
//
// CSS nodes
//
// GtkRevealer has a single CSS node with name revealer. When styling
// GtkRevealer using CSS, remember that it only hides its contents, not itself.
// That means applied margin, padding and borders will be visible even when the
// gtk.Revealer:reveal-child property is set to FALSE.
//
//
// Accessibility
//
// GtkRevealer uses the GTK_ACCESSIBLE_ROLE_GROUP role.
//
// The child of GtkRevealer, if set, is always available in the accessibility
// tree, regardless of the state of the revealer widget.
type Revealer struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Revealer)(nil)
)

func wrapRevealer(obj *coreglib.Object) *Revealer {
	return &Revealer{
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

func marshalRevealer(p uintptr) (interface{}, error) {
	return wrapRevealer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
