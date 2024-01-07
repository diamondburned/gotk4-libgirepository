// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_ActionGroup_ConnectActionStateChanged(gpointer, gchar*, GVariant*, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionRemoved(gpointer, gchar*, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionEnabledChanged(gpointer, gchar*, gboolean, guintptr);
// extern void _gotk4_gio2_ActionGroup_ConnectActionAdded(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeActionGroup = coreglib.Type(girepository.MustFind("Gio", "ActionGroup").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionGroup, F: marshalActionGroup},
	})
}

// ActionGroupOverrider contains methods that are overridable.
type ActionGroupOverrider interface {
}

// ActionGroup represents a group of actions. Actions can be used to expose
// functionality in a structured way, either from one part of a program to
// another, or to the outside world. Action groups are often used together with
// a Model that provides additional representation data for displaying the
// actions to the user, e.g. in a menu.
//
// The main way to interact with the actions in a GActionGroup is to activate
// them with g_action_group_activate_action(). Activating an action may require
// a #GVariant parameter. The required type of the parameter can be inquired
// with g_action_group_get_action_parameter_type(). Actions may be disabled, see
// g_action_group_get_action_enabled(). Activating a disabled action has no
// effect.
//
// Actions may optionally have a state in the form of a #GVariant. The current
// state of an action can be inquired with g_action_group_get_action_state().
// Activating a stateful action may change its state, but it is also possible to
// set the state by calling g_action_group_change_action_state().
//
// As typical example, consider a text editing application which has an option
// to change the current font to 'bold'. A good way to represent this would be a
// stateful action, with a boolean state. Activating the action would toggle the
// state.
//
// Each action in the group has a unique name (which is a string). All method
// calls, except g_action_group_list_actions() take the name of an action as an
// argument.
//
// The Group API is meant to be the 'public' API to the action group. The calls
// here are exactly the interaction that 'external forces' (eg: UI, incoming
// D-Bus messages, etc.) are supposed to have with actions. 'Internal' APIs (ie:
// ones meant only to be accessed by the action group implementation) are found
// on subclasses. This is why you will find - for example -
// g_action_group_get_action_enabled() but not an equivalent set() call.
//
// Signals are emitted on the action group in response to state changes on
// individual actions.
//
// Implementations of Group should provide implementations for the virtual
// functions g_action_group_list_actions() and g_action_group_query_action().
// The other virtual functions should not be implemented - their "wrappers" are
// actually implemented with calls to g_action_group_query_action().
//
// ActionGroup wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ActionGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ActionGroup)(nil)
)

// ActionGrouper describes ActionGroup's interface methods.
type ActionGrouper interface {
	coreglib.Objector

	baseActionGroup() *ActionGroup
}

var _ ActionGrouper = (*ActionGroup)(nil)

func ifaceInitActionGrouper(gifacePtr, data C.gpointer) {
}

func wrapActionGroup(obj *coreglib.Object) *ActionGroup {
	return &ActionGroup{
		Object: obj,
	}
}

func marshalActionGroup(p uintptr) (interface{}, error) {
	return wrapActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ActionGroup) baseActionGroup() *ActionGroup {
	return v
}

// BaseActionGroup returns the underlying base object.
func BaseActionGroup(obj ActionGrouper) *ActionGroup {
	return obj.baseActionGroup()
}

// ConnectActionAdded signals that a new action was just added to the group.
// This signal is emitted after the action has been added and is now visible.
func (v *ActionGroup) ConnectActionAdded(f func(actionName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "action-added", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionAdded), f)
}

// ConnectActionEnabledChanged signals that the enabled status of the named
// action has changed.
func (v *ActionGroup) ConnectActionEnabledChanged(f func(actionName string, enabled bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "action-enabled-changed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionEnabledChanged), f)
}

// ConnectActionRemoved signals that an action is just about to be removed from
// the group. This signal is emitted before the action is removed, so the action
// is still visible and can be queried from the signal handler.
func (v *ActionGroup) ConnectActionRemoved(f func(actionName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "action-removed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionRemoved), f)
}

// ConnectActionStateChanged signals that the state of the named action has
// changed.
func (v *ActionGroup) ConnectActionStateChanged(f func(actionName string, value *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "action-state-changed", false, unsafe.Pointer(C._gotk4_gio2_ActionGroup_ConnectActionStateChanged), f)
}
