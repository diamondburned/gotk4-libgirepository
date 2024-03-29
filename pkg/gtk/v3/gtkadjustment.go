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
// extern void _gotk4_gtk3_Adjustment_ConnectValueChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_Adjustment_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeAdjustment = coreglib.Type(girepository.MustFind("Gtk", "Adjustment").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAdjustment, F: marshalAdjustment},
	})
}

// AdjustmentOverrides contains methods that are overridable.
type AdjustmentOverrides struct {
}

func defaultAdjustmentOverrides(v *Adjustment) AdjustmentOverrides {
	return AdjustmentOverrides{}
}

// Adjustment object represents a value which has an associated lower and upper
// bound, together with step and page increments, and a page size. It is used
// within several GTK+ widgets, including SpinButton, Viewport, and Range (which
// is a base class for Scrollbar and Scale).
//
// The Adjustment object does not update the value itself. Instead it is left up
// to the owner of the Adjustment to control the value.
type Adjustment struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

func init() {
	coreglib.RegisterClassInfo[*Adjustment, *AdjustmentClass, AdjustmentOverrides](
		GTypeAdjustment,
		initAdjustmentClass,
		wrapAdjustment,
		defaultAdjustmentOverrides,
	)
}

func initAdjustmentClass(gclass unsafe.Pointer, overrides AdjustmentOverrides, classInitFunc func(*AdjustmentClass)) {
	if classInitFunc != nil {
		class := (*AdjustmentClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAdjustment(obj *coreglib.Object) *Adjustment {
	return &Adjustment{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	return wrapAdjustment(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when one or more of the Adjustment properties have
// been changed, other than the Adjustment:value property.
func (v *Adjustment) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk3_Adjustment_ConnectChanged), f)
}

// ConnectValueChanged is emitted when the Adjustment:value property has been
// changed.
func (v *Adjustment) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_Adjustment_ConnectValueChanged), f)
}

// AdjustmentClass: instance of this type is always passed by reference.
type AdjustmentClass struct {
	*adjustmentClass
}

// adjustmentClass is the struct that's finalized.
type adjustmentClass struct {
	native unsafe.Pointer
}

var GIRInfoAdjustmentClass = girepository.MustFind("Gtk", "AdjustmentClass")
