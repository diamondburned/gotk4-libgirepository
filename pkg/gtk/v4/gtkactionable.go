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
	GTypeActionable = coreglib.Type(girepository.MustFind("Gtk", "Actionable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionable, F: marshalActionable},
	})
}

// ActionableOverrider contains methods that are overridable.
type ActionableOverrider interface {
}

// Actionable: GtkActionable interface provides a convenient way of asscociating
// widgets with actions.
//
// It primarily consists of two properties: gtk.Actionable:action-name and
// gtk.Actionable:action-target. There are also some convenience APIs for
// setting these properties.
//
// The action will be looked up in action groups that are found among the
// widgets ancestors. Most commonly, these will be the actions with the “win.”
// or “app.” prefix that are associated with the GtkApplicationWindow or
// GtkApplication, but other action groups that are added with
// gtk.Widget.InsertActionGroup() will be consulted as well.
//
// Actionable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Actionable struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Actionable)(nil)
)

// Actionabler describes Actionable's interface methods.
type Actionabler interface {
	coreglib.Objector

	baseActionable() *Actionable
}

var _ Actionabler = (*Actionable)(nil)

func ifaceInitActionabler(gifacePtr, data C.gpointer) {
}

func wrapActionable(obj *coreglib.Object) *Actionable {
	return &Actionable{
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

func marshalActionable(p uintptr) (interface{}, error) {
	return wrapActionable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Actionable) baseActionable() *Actionable {
	return v
}

// BaseActionable returns the underlying base object.
func BaseActionable(obj Actionabler) *Actionable {
	return obj.baseActionable()
}

// ActionableInterface: interface vtable for GtkActionable.
//
// An instance of this type is always passed by reference.
type ActionableInterface struct {
	*actionableInterface
}

// actionableInterface is the struct that's finalized.
type actionableInterface struct {
	native unsafe.Pointer
}

var GIRInfoActionableInterface = girepository.MustFind("Gtk", "ActionableInterface")
