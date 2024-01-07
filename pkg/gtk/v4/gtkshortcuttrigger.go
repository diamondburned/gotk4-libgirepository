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
	GTypeAlternativeTrigger = coreglib.Type(girepository.MustFind("Gtk", "AlternativeTrigger").RegisteredGType())
	GTypeKeyvalTrigger      = coreglib.Type(girepository.MustFind("Gtk", "KeyvalTrigger").RegisteredGType())
	GTypeMnemonicTrigger    = coreglib.Type(girepository.MustFind("Gtk", "MnemonicTrigger").RegisteredGType())
	GTypeNeverTrigger       = coreglib.Type(girepository.MustFind("Gtk", "NeverTrigger").RegisteredGType())
	GTypeShortcutTrigger    = coreglib.Type(girepository.MustFind("Gtk", "ShortcutTrigger").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAlternativeTrigger, F: marshalAlternativeTrigger},
		coreglib.TypeMarshaler{T: GTypeKeyvalTrigger, F: marshalKeyvalTrigger},
		coreglib.TypeMarshaler{T: GTypeMnemonicTrigger, F: marshalMnemonicTrigger},
		coreglib.TypeMarshaler{T: GTypeNeverTrigger, F: marshalNeverTrigger},
		coreglib.TypeMarshaler{T: GTypeShortcutTrigger, F: marshalShortcutTrigger},
	})
}

// AlternativeTrigger: GtkShortcutTrigger that combines two triggers.
//
// The GtkAlternativeTrigger triggers when either of two trigger.
//
// This can be cascaded to combine more than two triggers.
type AlternativeTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*AlternativeTrigger)(nil)
)

func wrapAlternativeTrigger(obj *coreglib.Object) *AlternativeTrigger {
	return &AlternativeTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalAlternativeTrigger(p uintptr) (interface{}, error) {
	return wrapAlternativeTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// KeyvalTrigger: GtkShortcutTrigger that triggers when a specific keyval and
// modifiers are pressed.
type KeyvalTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*KeyvalTrigger)(nil)
)

func wrapKeyvalTrigger(obj *coreglib.Object) *KeyvalTrigger {
	return &KeyvalTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalKeyvalTrigger(p uintptr) (interface{}, error) {
	return wrapKeyvalTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MnemonicTrigger: GtkShortcutTrigger that triggers when a specific mnemonic is
// pressed.
//
// Mnemonics require a *mnemonic modifier* (typically <kbd>Alt</kbd>) to be
// pressed together with the mnemonic key.
type MnemonicTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*MnemonicTrigger)(nil)
)

func wrapMnemonicTrigger(obj *coreglib.Object) *MnemonicTrigger {
	return &MnemonicTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalMnemonicTrigger(p uintptr) (interface{}, error) {
	return wrapMnemonicTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NeverTrigger: GtkShortcutTrigger that never triggers.
type NeverTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*NeverTrigger)(nil)
)

func wrapNeverTrigger(obj *coreglib.Object) *NeverTrigger {
	return &NeverTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalNeverTrigger(p uintptr) (interface{}, error) {
	return wrapNeverTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ShortcutTrigger: GtkShortcutTrigger tracks how a GtkShortcut should be
// activated.
//
// To find out if a GtkShortcutTrigger triggers, you can call
// gtk.ShortcutTrigger.Trigger() on a GdkEvent.
//
// GtkShortcutTriggers contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutTriggers are immutable, you can only specify their properties
// during construction. If you want to change a trigger, you have to replace it
// with a new one.
type ShortcutTrigger struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutTrigger)(nil)
)

// ShortcutTriggerer describes types inherited from class ShortcutTrigger.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ShortcutTriggerer interface {
	coreglib.Objector
	baseShortcutTrigger() *ShortcutTrigger
}

var _ ShortcutTriggerer = (*ShortcutTrigger)(nil)

func wrapShortcutTrigger(obj *coreglib.Object) *ShortcutTrigger {
	return &ShortcutTrigger{
		Object: obj,
	}
}

func marshalShortcutTrigger(p uintptr) (interface{}, error) {
	return wrapShortcutTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ShortcutTrigger) baseShortcutTrigger() *ShortcutTrigger {
	return v
}

// BaseShortcutTrigger returns the underlying base object.
func BaseShortcutTrigger(obj ShortcutTriggerer) *ShortcutTrigger {
	return obj.baseShortcutTrigger()
}
