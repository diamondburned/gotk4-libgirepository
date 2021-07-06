// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_simple_action_group_get_type()), F: marshalSimpleActionGroup},
	})
}

// SimpleActionGroup is a hash table filled with #GAction objects, implementing
// the Group and Map interfaces.
type SimpleActionGroup interface {
	gextras.Objector

	// AsActionGroup casts the class to the ActionGroup interface.
	AsActionGroup() ActionGroup
	// AsActionMap casts the class to the ActionMap interface.
	AsActionMap() ActionMap

	// ActionAdded emits the Group::action-added signal on @action_group.
	//
	// This function should only be called by Group implementations.
	//
	// This method is inherited from ActionGroup
	ActionAdded(actionName string)
	// ActionEnabledChanged emits the Group::action-enabled-changed signal on
	// @action_group.
	//
	// This function should only be called by Group implementations.
	//
	// This method is inherited from ActionGroup
	ActionEnabledChanged(actionName string, enabled bool)
	// ActionRemoved emits the Group::action-removed signal on @action_group.
	//
	// This function should only be called by Group implementations.
	//
	// This method is inherited from ActionGroup
	ActionRemoved(actionName string)
	// ActionStateChanged emits the Group::action-state-changed signal on
	// @action_group.
	//
	// This function should only be called by Group implementations.
	//
	// This method is inherited from ActionGroup
	ActionStateChanged(actionName string, state *glib.Variant)
	// ActivateAction: activate the named action within @action_group.
	//
	// If the action is expecting a parameter, then the correct type of
	// parameter must be given as @parameter. If the action is expecting no
	// parameters then @parameter must be nil. See
	// g_action_group_get_action_parameter_type().
	//
	// If the Group implementation supports asynchronous remote activation over
	// D-Bus, this call may return before the relevant D-Bus traffic has been
	// sent, or any replies have been received. In order to block on such
	// asynchronous activation calls, g_dbus_connection_flush() should be called
	// prior to the code, which depends on the result of the action activation.
	// Without flushing the D-Bus connection, there is no guarantee that the
	// action would have been activated.
	//
	// The following code which runs in a remote app instance, shows an example
	// of a "quit" action being activated on the primary app instance over
	// D-Bus. Here g_dbus_connection_flush() is called before `exit()`. Without
	// g_dbus_connection_flush(), the "quit" action may fail to be activated on
	// the primary instance.
	//
	//    // call "quit" action on primary instance
	//    g_action_group_activate_action (G_ACTION_GROUP (app), "quit", NULL);
	//
	//    // make sure the action is activated now
	//    g_dbus_connection_flush (...);
	//
	//    g_debug ("application has been terminated. exiting.");
	//
	//    exit (0);
	//
	// This method is inherited from ActionGroup
	ActivateAction(actionName string, parameter *glib.Variant)
	// ChangeActionState: request for the state of the named action within
	// @action_group to be changed to @value.
	//
	// The action must be stateful and @value must be of the correct type. See
	// g_action_group_get_action_state_type().
	//
	// This call merely requests a change. The action may refuse to change its
	// state or may change its state to something other than @value. See
	// g_action_group_get_action_state_hint().
	//
	// If the @value GVariant is floating, it is consumed.
	//
	// This method is inherited from ActionGroup
	ChangeActionState(actionName string, value *glib.Variant)
	// GetActionEnabled checks if the named action within @action_group is
	// currently enabled.
	//
	// An action must be enabled in order to be activated or in order to have
	// its state changed from outside callers.
	//
	// This method is inherited from ActionGroup
	GetActionEnabled(actionName string) bool
	// GetActionParameterType queries the type of the parameter that must be
	// given when activating the named action within @action_group.
	//
	// When activating the action using g_action_group_activate_action(), the
	// #GVariant given to that function must be of the type returned by this
	// function.
	//
	// In the case that this function returns nil, you must not give any
	// #GVariant, but nil instead.
	//
	// The parameter type of a particular action will never change but it is
	// possible for an action to be removed and for a new action to be added
	// with the same name but a different parameter type.
	//
	// This method is inherited from ActionGroup
	GetActionParameterType(actionName string) *glib.VariantType
	// GetActionState queries the current state of the named action within
	// @action_group.
	//
	// If the action is not stateful then nil will be returned. If the action is
	// stateful then the type of the return value is the type given by
	// g_action_group_get_action_state_type().
	//
	// The return value (if non-nil) should be freed with g_variant_unref() when
	// it is no longer required.
	//
	// This method is inherited from ActionGroup
	GetActionState(actionName string) *glib.Variant
	// GetActionStateHint requests a hint about the valid range of values for
	// the state of the named action within @action_group.
	//
	// If nil is returned it either means that the action is not stateful or
	// that there is no hint about the valid range of values for the state of
	// the action.
	//
	// If a #GVariant array is returned then each item in the array is a
	// possible value for the state. If a #GVariant pair (ie: two-tuple) is
	// returned then the tuple specifies the inclusive lower and upper bound of
	// valid values for the state.
	//
	// In any case, the information is merely a hint. It may be possible to have
	// a state value outside of the hinted range and setting a value within the
	// range may fail.
	//
	// The return value (if non-nil) should be freed with g_variant_unref() when
	// it is no longer required.
	//
	// This method is inherited from ActionGroup
	GetActionStateHint(actionName string) *glib.Variant
	// GetActionStateType queries the type of the state of the named action
	// within @action_group.
	//
	// If the action is stateful then this function returns the Type of the
	// state. All calls to g_action_group_change_action_state() must give a
	// #GVariant of this type and g_action_group_get_action_state() will return
	// a #GVariant of the same type.
	//
	// If the action is not stateful then this function will return nil. In that
	// case, g_action_group_get_action_state() will return nil and you must not
	// call g_action_group_change_action_state().
	//
	// The state type of a particular action will never change but it is
	// possible for an action to be removed and for a new action to be added
	// with the same name but a different state type.
	//
	// This method is inherited from ActionGroup
	GetActionStateType(actionName string) *glib.VariantType
	// HasAction checks if the named action exists within @action_group.
	//
	// This method is inherited from ActionGroup
	HasAction(actionName string) bool
	// ListActions lists the actions contained within @action_group.
	//
	// The caller is responsible for freeing the list with g_strfreev() when it
	// is no longer required.
	//
	// This method is inherited from ActionGroup
	ListActions() []string
	// QueryAction queries all aspects of the named action within an
	// @action_group.
	//
	// This function acquires the information available from
	// g_action_group_has_action(), g_action_group_get_action_enabled(),
	// g_action_group_get_action_parameter_type(),
	// g_action_group_get_action_state_type(),
	// g_action_group_get_action_state_hint() and
	// g_action_group_get_action_state() with a single function call.
	//
	// This provides two main benefits.
	//
	// The first is the improvement in efficiency that comes with not having to
	// perform repeated lookups of the action in order to discover different
	// things about it. The second is that implementing Group can now be done by
	// only overriding this one virtual function.
	//
	// The interface provides a default implementation of this function that
	// calls the individual functions, as required, to fetch the information.
	// The interface also provides default implementations of those functions
	// that call this function. All implementations, therefore, must override
	// either this function or all of the others.
	//
	// If the action exists, true is returned and any of the requested fields
	// (as indicated by having a non-nil reference passed in) are filled. If the
	// action doesn't exist, false is returned and the fields may or may not
	// have been modified.
	//
	// This method is inherited from ActionGroup
	QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool)
	// AddAction adds an action to the @action_map.
	//
	// If the action map already contains an action with the same name as
	// @action then the old action is dropped from the action map.
	//
	// The action map takes its own reference on @action.
	//
	// This method is inherited from ActionMap
	AddAction(action Action)
	// AddActionEntries: convenience function for creating multiple Action
	// instances and adding them to a Map.
	//
	// Each action is constructed as per one Entry.
	//
	//    static void
	//    activate_quit (GSimpleAction *simple,
	//                   GVariant      *parameter,
	//                   gpointer       user_data)
	//    {
	//      exit (0);
	//    }
	//
	//    static void
	//    activate_print_string (GSimpleAction *simple,
	//                           GVariant      *parameter,
	//                           gpointer       user_data)
	//    {
	//      g_print ("s\n", g_variant_get_string (parameter, NULL));
	//    }
	//
	//    static GActionGroup *
	//    create_action_group (void)
	//    {
	//      const GActionEntry entries[] = {
	//        { "quit",         activate_quit              },
	//        { "print-string", activate_print_string, "s" }
	//      };
	//      GSimpleActionGroup *group;
	//
	//      group = g_simple_action_group_new ();
	//      g_action_map_add_action_entries (G_ACTION_MAP (group), entries, G_N_ELEMENTS (entries), NULL);
	//
	//      return G_ACTION_GROUP (group);
	//    }
	//
	// This method is inherited from ActionMap
	AddActionEntries(entries []ActionEntry, userData interface{})
	// LookupAction looks up the action with the name @action_name in
	// @action_map.
	//
	// If no such action exists, returns nil.
	//
	// This method is inherited from ActionMap
	LookupAction(actionName string) Action
	// RemoveAction removes the named action from the action map.
	//
	// If no action of this name is in the map then nothing happens.
	//
	// This method is inherited from ActionMap
	RemoveAction(actionName string)

	// AddEntries: convenience function for creating multiple Action instances
	// and adding them to the action group.
	//
	// Deprecated: since version 2.38.
	AddEntries(entries []ActionEntry, userData interface{})
	// Insert adds an action to the action group.
	//
	// If the action group already contains an action with the same name as
	// @action then the old action is dropped from the group.
	//
	// The action group takes its own reference on @action.
	//
	// Deprecated: since version 2.38.
	Insert(action Action)
	// Lookup looks up the action with the name @action_name in the group.
	//
	// If no such action exists, returns nil.
	//
	// Deprecated: since version 2.38.
	Lookup(actionName string) Action
	// Remove removes the named action from the action group.
	//
	// If no action of this name is in the group then nothing happens.
	//
	// Deprecated: since version 2.38.
	Remove(actionName string)
}

// simpleActionGroup implements the SimpleActionGroup interface.
type simpleActionGroup struct {
	*externglib.Object
}

var _ SimpleActionGroup = (*simpleActionGroup)(nil)

// WrapSimpleActionGroup wraps a GObject to a type that implements
// interface SimpleActionGroup. It is primarily used internally.
func WrapSimpleActionGroup(obj *externglib.Object) SimpleActionGroup {
	return simpleActionGroup{obj}
}

func marshalSimpleActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSimpleActionGroup(obj), nil
}

// NewSimpleActionGroup creates a new, empty, ActionGroup.
func NewSimpleActionGroup() SimpleActionGroup {
	var _cret *C.GSimpleActionGroup // in

	_cret = C.g_simple_action_group_new()

	var _simpleActionGroup SimpleActionGroup // out

	_simpleActionGroup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(SimpleActionGroup)

	return _simpleActionGroup
}

func (s simpleActionGroup) AsActionGroup() ActionGroup {
	return WrapActionGroup(gextras.InternObject(s))
}

func (s simpleActionGroup) AsActionMap() ActionMap {
	return WrapActionMap(gextras.InternObject(s))
}

func (a simpleActionGroup) ActionAdded(actionName string) {
	WrapActionGroup(gextras.InternObject(a)).ActionAdded(actionName)
}

func (a simpleActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	WrapActionGroup(gextras.InternObject(a)).ActionEnabledChanged(actionName, enabled)
}

func (a simpleActionGroup) ActionRemoved(actionName string) {
	WrapActionGroup(gextras.InternObject(a)).ActionRemoved(actionName)
}

func (a simpleActionGroup) ActionStateChanged(actionName string, state *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ActionStateChanged(actionName, state)
}

func (a simpleActionGroup) ActivateAction(actionName string, parameter *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ActivateAction(actionName, parameter)
}

func (a simpleActionGroup) ChangeActionState(actionName string, value *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ChangeActionState(actionName, value)
}

func (a simpleActionGroup) GetActionEnabled(actionName string) bool {
	return WrapActionGroup(gextras.InternObject(a)).GetActionEnabled(actionName)
}

func (a simpleActionGroup) GetActionParameterType(actionName string) *glib.VariantType {
	return WrapActionGroup(gextras.InternObject(a)).GetActionParameterType(actionName)
}

func (a simpleActionGroup) GetActionState(actionName string) *glib.Variant {
	return WrapActionGroup(gextras.InternObject(a)).GetActionState(actionName)
}

func (a simpleActionGroup) GetActionStateHint(actionName string) *glib.Variant {
	return WrapActionGroup(gextras.InternObject(a)).GetActionStateHint(actionName)
}

func (a simpleActionGroup) GetActionStateType(actionName string) *glib.VariantType {
	return WrapActionGroup(gextras.InternObject(a)).GetActionStateType(actionName)
}

func (a simpleActionGroup) HasAction(actionName string) bool {
	return WrapActionGroup(gextras.InternObject(a)).HasAction(actionName)
}

func (a simpleActionGroup) ListActions() []string {
	return WrapActionGroup(gextras.InternObject(a)).ListActions()
}

func (a simpleActionGroup) QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool) {
	return WrapActionGroup(gextras.InternObject(a)).QueryAction(actionName)
}

func (a simpleActionGroup) AddAction(action Action) {
	WrapActionMap(gextras.InternObject(a)).AddAction(action)
}

func (a simpleActionGroup) AddActionEntries(entries []ActionEntry, userData interface{}) {
	WrapActionMap(gextras.InternObject(a)).AddActionEntries(entries, userData)
}

func (a simpleActionGroup) LookupAction(actionName string) Action {
	return WrapActionMap(gextras.InternObject(a)).LookupAction(actionName)
}

func (a simpleActionGroup) RemoveAction(actionName string) {
	WrapActionMap(gextras.InternObject(a)).RemoveAction(actionName)
}

func (s simpleActionGroup) AddEntries(entries []ActionEntry, userData interface{}) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GActionEntry
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg2 = C.gint(len(entries))
	_arg1 = (*C.GActionEntry)(unsafe.Pointer(&entries[0]))
	_arg3 = (C.gpointer)(box.Assign(userData))

	C.g_simple_action_group_add_entries(_arg0, _arg1, _arg2, _arg3)
}

func (s simpleActionGroup) Insert(action Action) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GAction            // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_simple_action_group_insert(_arg0, _arg1)
}

func (s simpleActionGroup) Lookup(actionName string) Action {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out
	var _cret *C.GAction            // in

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_simple_action_group_lookup(_arg0, _arg1)

	var _action Action // out

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (s simpleActionGroup) Remove(actionName string) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_action_group_remove(_arg0, _arg1)
}
