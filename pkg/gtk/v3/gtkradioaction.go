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
// extern void _gotk4_gtk3_RadioAction_ConnectChanged(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeRadioAction = coreglib.Type(girepository.MustFind("Gtk", "RadioAction").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRadioAction, F: marshalRadioAction},
	})
}

// RadioActionOverrides contains methods that are overridable.
type RadioActionOverrides struct {
}

func defaultRadioActionOverrides(v *RadioAction) RadioActionOverrides {
	return RadioActionOverrides{}
}

// RadioAction is similar to RadioMenuItem. A number of radio actions can be
// linked together so that only one may be active at any one time.
type RadioAction struct {
	_ [0]func() // equal guard
	ToggleAction
}

var (
	_ coreglib.Objector = (*RadioAction)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RadioAction, *RadioActionClass, RadioActionOverrides](
		GTypeRadioAction,
		initRadioActionClass,
		wrapRadioAction,
		defaultRadioActionOverrides,
	)
}

func initRadioActionClass(gclass unsafe.Pointer, overrides RadioActionOverrides, classInitFunc func(*RadioActionClass)) {
	if classInitFunc != nil {
		class := (*RadioActionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRadioAction(obj *coreglib.Object) *RadioAction {
	return &RadioAction{
		ToggleAction: ToggleAction{
			Action: Action{
				Object: obj,
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioAction(p uintptr) (interface{}, error) {
	return wrapRadioAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged signal is emitted on every member of a radio group when the
// active member is changed. The signal gets emitted after the ::activate
// signals for the previous and current active members.
func (v *RadioAction) ConnectChanged(f func(current *RadioAction)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk3_RadioAction_ConnectChanged), f)
}

// RadioActionClass: instance of this type is always passed by reference.
type RadioActionClass struct {
	*radioActionClass
}

// radioActionClass is the struct that's finalized.
type radioActionClass struct {
	native unsafe.Pointer
}

var GIRInfoRadioActionClass = girepository.MustFind("Gtk", "RadioActionClass")
