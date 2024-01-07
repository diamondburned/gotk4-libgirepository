// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeShortcutActionFlags = coreglib.Type(girepository.MustFind("Gtk", "ShortcutActionFlags").RegisteredGType())
	GTypeActivateAction      = coreglib.Type(girepository.MustFind("Gtk", "ActivateAction").RegisteredGType())
	GTypeCallbackAction      = coreglib.Type(girepository.MustFind("Gtk", "CallbackAction").RegisteredGType())
	GTypeMnemonicAction      = coreglib.Type(girepository.MustFind("Gtk", "MnemonicAction").RegisteredGType())
	GTypeNamedAction         = coreglib.Type(girepository.MustFind("Gtk", "NamedAction").RegisteredGType())
	GTypeNothingAction       = coreglib.Type(girepository.MustFind("Gtk", "NothingAction").RegisteredGType())
	GTypeShortcutAction      = coreglib.Type(girepository.MustFind("Gtk", "ShortcutAction").RegisteredGType())
	GTypeSignalAction        = coreglib.Type(girepository.MustFind("Gtk", "SignalAction").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutActionFlags, F: marshalShortcutActionFlags},
		coreglib.TypeMarshaler{T: GTypeActivateAction, F: marshalActivateAction},
		coreglib.TypeMarshaler{T: GTypeCallbackAction, F: marshalCallbackAction},
		coreglib.TypeMarshaler{T: GTypeMnemonicAction, F: marshalMnemonicAction},
		coreglib.TypeMarshaler{T: GTypeNamedAction, F: marshalNamedAction},
		coreglib.TypeMarshaler{T: GTypeNothingAction, F: marshalNothingAction},
		coreglib.TypeMarshaler{T: GTypeShortcutAction, F: marshalShortcutAction},
		coreglib.TypeMarshaler{T: GTypeSignalAction, F: marshalSignalAction},
	})
}

// ShortcutActionFlags: list of flags that can be passed to action activation.
//
// More flags may be added in the future.
type ShortcutActionFlags C.guint

const (
	// ShortcutActionExclusive: action is the only action that can be activated.
	// If this flag is not set, a future activation may select a different
	// action.
	ShortcutActionExclusive ShortcutActionFlags = 0b1
)

func marshalShortcutActionFlags(p uintptr) (interface{}, error) {
	return ShortcutActionFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ShortcutActionFlags.
func (s ShortcutActionFlags) String() string {
	if s == 0 {
		return "ShortcutActionFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(23)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case ShortcutActionExclusive:
			builder.WriteString("Exclusive|")
		default:
			builder.WriteString(fmt.Sprintf("ShortcutActionFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s ShortcutActionFlags) Has(other ShortcutActionFlags) bool {
	return (s & other) == other
}

// ShortcutFunc: prototype for shortcuts based on user callbacks.
type ShortcutFunc func(widget Widgetter, args *glib.Variant) (ok bool)

// ActivateAction: GtkShortcutAction that calls gtk_widget_activate().
type ActivateAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*ActivateAction)(nil)
)

func wrapActivateAction(obj *coreglib.Object) *ActivateAction {
	return &ActivateAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalActivateAction(p uintptr) (interface{}, error) {
	return wrapActivateAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CallbackAction: GtkShortcutAction that invokes a callback.
type CallbackAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*CallbackAction)(nil)
)

func wrapCallbackAction(obj *coreglib.Object) *CallbackAction {
	return &CallbackAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalCallbackAction(p uintptr) (interface{}, error) {
	return wrapCallbackAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MnemonicAction: GtkShortcutAction that calls gtk_widget_mnemonic_activate().
type MnemonicAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*MnemonicAction)(nil)
)

func wrapMnemonicAction(obj *coreglib.Object) *MnemonicAction {
	return &MnemonicAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalMnemonicAction(p uintptr) (interface{}, error) {
	return wrapMnemonicAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NamedAction: GtkShortcutAction that activates an action by name.
type NamedAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*NamedAction)(nil)
)

func wrapNamedAction(obj *coreglib.Object) *NamedAction {
	return &NamedAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNamedAction(p uintptr) (interface{}, error) {
	return wrapNamedAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NothingAction: GtkShortcutAction that does nothing.
type NothingAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*NothingAction)(nil)
)

func wrapNothingAction(obj *coreglib.Object) *NothingAction {
	return &NothingAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNothingAction(p uintptr) (interface{}, error) {
	return wrapNothingAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ShortcutAction: GtkShortcutAction encodes an action that can be triggered by
// a keyboard shortcut.
//
// GtkShortcutActions contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutActions are immutable, you can only specify their properties
// during construction. If you want to change a action, you have to replace it
// with a new one. If you need to pass arguments to an action, these are
// specified by the higher-level GtkShortcut object.
//
// To activate a GtkShortcutAction manually, gtk.ShortcutAction.Activate() can
// be called.
//
// GTK provides various actions:
//
//    - gtk.MnemonicAction: a shortcut action that calls
//      gtk_widget_mnemonic_activate()
//    - gtk.CallbackAction: a shortcut action that invokes
//      a given callback
//    - gtk.SignalAction: a shortcut action that emits a
//      given signal
//    - gtk.ActivateAction: a shortcut action that calls
//      gtk_widget_activate()
//    - gtk.NamedAction: a shortcut action that calls
//      gtk_widget_activate_action()
//    - gtk.NothingAction: a shortcut action that does nothing.
type ShortcutAction struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutAction)(nil)
)

// ShortcutActioner describes types inherited from class ShortcutAction.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ShortcutActioner interface {
	coreglib.Objector
	baseShortcutAction() *ShortcutAction
}

var _ ShortcutActioner = (*ShortcutAction)(nil)

func wrapShortcutAction(obj *coreglib.Object) *ShortcutAction {
	return &ShortcutAction{
		Object: obj,
	}
}

func marshalShortcutAction(p uintptr) (interface{}, error) {
	return wrapShortcutAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ShortcutAction) baseShortcutAction() *ShortcutAction {
	return v
}

// BaseShortcutAction returns the underlying base object.
func BaseShortcutAction(obj ShortcutActioner) *ShortcutAction {
	return obj.baseShortcutAction()
}

// SignalAction: GtkShortcutAction that emits a signal.
//
// Signals that are used in this way are referred to as keybinding signals, and
// they are expected to be defined with the G_SIGNAL_ACTION flag.
type SignalAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*SignalAction)(nil)
)

func wrapSignalAction(obj *coreglib.Object) *SignalAction {
	return &SignalAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalSignalAction(p uintptr) (interface{}, error) {
	return wrapSignalAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
