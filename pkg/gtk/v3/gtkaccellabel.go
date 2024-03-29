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
	GTypeAccelLabel = coreglib.Type(girepository.MustFind("Gtk", "AccelLabel").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAccelLabel, F: marshalAccelLabel},
	})
}

// AccelLabelOverrides contains methods that are overridable.
type AccelLabelOverrides struct {
}

func defaultAccelLabelOverrides(v *AccelLabel) AccelLabelOverrides {
	return AccelLabelOverrides{}
}

// AccelLabel widget is a subclass of Label that also displays an accelerator
// key on the right of the label text, e.g. “Ctrl+S”. It is commonly used in
// menus to show the keyboard short-cuts for commands.
//
// The accelerator key to display is typically not set explicitly (although it
// can be, with gtk_accel_label_set_accel()). Instead, the AccelLabel displays
// the accelerators which have been added to a particular widget. This widget is
// set by calling gtk_accel_label_set_accel_widget().
//
// For example, a MenuItem widget may have an accelerator added to emit the
// “activate” signal when the “Ctrl+S” key combination is pressed. A AccelLabel
// is created and added to the MenuItem, and gtk_accel_label_set_accel_widget()
// is called with the MenuItem as the second argument. The AccelLabel will now
// display “Ctrl+S” after its label.
//
// Note that creating a MenuItem with gtk_menu_item_new_with_label() (or one of
// the similar functions for CheckMenuItem and RadioMenuItem) automatically adds
// a AccelLabel to the MenuItem and calls gtk_accel_label_set_accel_widget() to
// set it up for you.
//
// A AccelLabel will only display accelerators which have GTK_ACCEL_VISIBLE set
// (see AccelFlags). A AccelLabel can display multiple accelerators and even
// signal names, though it is almost always used to display just one accelerator
// key.
//
// Creating a simple menu item with an accelerator key.
//
//    label
//    ╰── accelerator
//
// Like Label, GtkAccelLabel has a main CSS node with the name label. It adds a
// subnode with name accelerator.
type AccelLabel struct {
	_ [0]func() // equal guard
	Label
}

var (
	_ Miscer = (*AccelLabel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AccelLabel, *AccelLabelClass, AccelLabelOverrides](
		GTypeAccelLabel,
		initAccelLabelClass,
		wrapAccelLabel,
		defaultAccelLabelOverrides,
	)
}

func initAccelLabelClass(gclass unsafe.Pointer, overrides AccelLabelOverrides, classInitFunc func(*AccelLabelClass)) {
	if classInitFunc != nil {
		class := (*AccelLabelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAccelLabel(obj *coreglib.Object) *AccelLabel {
	return &AccelLabel{
		Label: Label{
			Misc: Misc{
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
			},
		},
	}
}

func marshalAccelLabel(p uintptr) (interface{}, error) {
	return wrapAccelLabel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AccelLabelClass: instance of this type is always passed by reference.
type AccelLabelClass struct {
	*accelLabelClass
}

// accelLabelClass is the struct that's finalized.
type accelLabelClass struct {
	native unsafe.Pointer
}

var GIRInfoAccelLabelClass = girepository.MustFind("Gtk", "AccelLabelClass")
