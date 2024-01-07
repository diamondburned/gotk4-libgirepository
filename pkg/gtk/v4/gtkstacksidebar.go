// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeStackSidebar = coreglib.Type(girepository.MustFind("Gtk", "StackSidebar").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStackSidebar, F: marshalStackSidebar},
	})
}

// StackSidebar: GtkStackSidebar uses a sidebar to switch between GtkStack
// pages.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk.StackSidebar.SetStack() to connect the GtkStackSidebar to the GtkStack.
//
//
// CSS nodes
//
// GtkStackSidebar has a single CSS node with name stacksidebar and style class
// .sidebar.
//
// When circumstances require it, GtkStackSidebar adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*StackSidebar)(nil)
)

func wrapStackSidebar(obj *coreglib.Object) *StackSidebar {
	return &StackSidebar{
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

func marshalStackSidebar(p uintptr) (interface{}, error) {
	return wrapStackSidebar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
