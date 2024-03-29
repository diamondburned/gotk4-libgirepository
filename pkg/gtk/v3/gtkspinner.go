// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeSpinner = coreglib.Type(girepository.MustFind("Gtk", "Spinner").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinner, F: marshalSpinner},
	})
}

// SpinnerOverrides contains methods that are overridable.
type SpinnerOverrides struct {
}

func defaultSpinnerOverrides(v *Spinner) SpinnerOverrides {
	return SpinnerOverrides{}
}

// Spinner widget displays an icon-size spinning animation. It is often used as
// an alternative to a ProgressBar for displaying indefinite activity, instead
// of actual progress.
//
// To start the animation, use gtk_spinner_start(), to stop it use
// gtk_spinner_stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Spinner)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Spinner, *SpinnerClass, SpinnerOverrides](
		GTypeSpinner,
		initSpinnerClass,
		wrapSpinner,
		defaultSpinnerOverrides,
	)
}

func initSpinnerClass(gclass unsafe.Pointer, overrides SpinnerOverrides, classInitFunc func(*SpinnerClass)) {
	if classInitFunc != nil {
		class := (*SpinnerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSpinner(obj *coreglib.Object) *Spinner {
	return &Spinner{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalSpinner(p uintptr) (interface{}, error) {
	return wrapSpinner(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SpinnerClass: instance of this type is always passed by reference.
type SpinnerClass struct {
	*spinnerClass
}

// spinnerClass is the struct that's finalized.
type spinnerClass struct {
	native unsafe.Pointer
}

var GIRInfoSpinnerClass = girepository.MustFind("Gtk", "SpinnerClass")
