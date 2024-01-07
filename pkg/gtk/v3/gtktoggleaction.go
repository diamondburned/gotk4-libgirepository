// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ToggleAction_ConnectToggled(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeToggleAction = coreglib.Type(girepository.MustFind("Gtk", "ToggleAction").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToggleAction, F: marshalToggleAction},
	})
}

// ToggleActionOverrides contains methods that are overridable.
type ToggleActionOverrides struct {
}

func defaultToggleActionOverrides(v *ToggleAction) ToggleActionOverrides {
	return ToggleActionOverrides{}
}

// ToggleAction corresponds roughly to a CheckMenuItem. It has an “active” state
// specifying whether the action has been checked or not.
type ToggleAction struct {
	_ [0]func() // equal guard
	Action
}

var (
	_ coreglib.Objector = (*ToggleAction)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ToggleAction, *ToggleActionClass, ToggleActionOverrides](
		GTypeToggleAction,
		initToggleActionClass,
		wrapToggleAction,
		defaultToggleActionOverrides,
	)
}

func initToggleActionClass(gclass unsafe.Pointer, overrides ToggleActionOverrides, classInitFunc func(*ToggleActionClass)) {
	if classInitFunc != nil {
		class := (*ToggleActionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapToggleAction(obj *coreglib.Object) *ToggleAction {
	return &ToggleAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToggleAction(p uintptr) (interface{}, error) {
	return wrapToggleAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectToggled: should be connected if you wish to perform an action whenever
// the ToggleAction state is changed.
func (v *ToggleAction) ConnectToggled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_ToggleAction_ConnectToggled), f)
}

// ToggleActionClass: instance of this type is always passed by reference.
type ToggleActionClass struct {
	*toggleActionClass
}

// toggleActionClass is the struct that's finalized.
type toggleActionClass struct {
	native unsafe.Pointer
}

var GIRInfoToggleActionClass = girepository.MustFind("Gtk", "ToggleActionClass")
