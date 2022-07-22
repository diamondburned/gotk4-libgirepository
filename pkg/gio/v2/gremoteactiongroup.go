// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_RemoteActionGroupInterface_activate_action_full(GRemoteActionGroup*, gchar*, GVariant*, GVariant*);
// extern void _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full(GRemoteActionGroup*, gchar*, GVariant*, GVariant*);
import "C"

// GType values.
var (
	GTypeRemoteActionGroup = coreglib.Type(C.g_remote_action_group_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRemoteActionGroup, F: marshalRemoteActionGroup},
	})
}

// RemoteActionGroupOverrider contains methods that are overridable.
type RemoteActionGroupOverrider interface {
	// ActivateActionFull activates the remote action.
	//
	// This is the same as g_action_group_activate_action() except that it
	// allows for provision of "platform data" to be sent along with the
	// activation request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// platform_data must be non-NULL and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action to activate.
	//    - parameter (optional): optional parameter to the activation.
	//    - platformData: platform data to send.
	//
	ActivateActionFull(actionName string, parameter, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	//
	// This is the same as g_action_group_change_action_state() except that it
	// allows for provision of "platform data" to be sent along with the state
	// change request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// platform_data must be non-NULL and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action to change the state of.
	//    - value: new requested value for the state.
	//    - platformData: platform data to send.
	//
	ChangeActionStateFull(actionName string, value, platformData *glib.Variant)
}

// RemoteActionGroup interface is implemented by Group instances that either
// transmit action invocations to other processes or receive action invocations
// in the local process from other processes.
//
// The interface has _full variants of the two methods on Group used to activate
// actions: g_action_group_activate_action() and
// g_action_group_change_action_state(). These variants allow a "platform data"
// #GVariant to be specified: a dictionary providing context for the action
// invocation (for example: timestamps, startup notification IDs, etc).
//
// BusActionGroup implements ActionGroup. This provides a mechanism to send
// platform data for action invocations over D-Bus.
//
// Additionally, g_dbus_connection_export_action_group() will check if the
// exported Group implements ActionGroup and use the _full variants of the calls
// if available. This provides a mechanism by which to receive platform data for
// action invocations that arrive by way of D-Bus.
//
// RemoteActionGroup wraps an interface. This means the user can get the
// underlying type by calling Cast().
type RemoteActionGroup struct {
	_ [0]func() // equal guard
	ActionGroup
}

var ()

// RemoteActionGrouper describes RemoteActionGroup's interface methods.
type RemoteActionGrouper interface {
	coreglib.Objector

	// ActivateActionFull activates the remote action.
	ActivateActionFull(actionName string, parameter, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	ChangeActionStateFull(actionName string, value, platformData *glib.Variant)
}

var _ RemoteActionGrouper = (*RemoteActionGroup)(nil)

func ifaceInitRemoteActionGrouper(gifacePtr, data C.gpointer) {
	iface := (*C.GRemoteActionGroupInterface)(unsafe.Pointer(gifacePtr))
	iface.activate_action_full = (*[0]byte)(C._gotk4_gio2_RemoteActionGroupInterface_activate_action_full)
	iface.change_action_state_full = (*[0]byte)(C._gotk4_gio2_RemoteActionGroupInterface_change_action_state_full)
}

//export _gotk4_gio2_RemoteActionGroupInterface_activate_action_full
func _gotk4_gio2_RemoteActionGroupInterface_activate_action_full(arg0 *C.GRemoteActionGroup, arg1 *C.gchar, arg2 *C.GVariant, arg3 *C.GVariant) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(RemoteActionGroupOverrider)

	var _actionName string          // out
	var _parameter *glib.Variant    // out
	var _platformData *glib.Variant // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_parameter = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
		C.g_variant_ref(arg2)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_parameter)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}
	_platformData = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_platformData)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	iface.ActivateActionFull(_actionName, _parameter, _platformData)
}

//export _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full
func _gotk4_gio2_RemoteActionGroupInterface_change_action_state_full(arg0 *C.GRemoteActionGroup, arg1 *C.gchar, arg2 *C.GVariant, arg3 *C.GVariant) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(RemoteActionGroupOverrider)

	var _actionName string          // out
	var _value *glib.Variant        // out
	var _platformData *glib.Variant // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_variant_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)
	_platformData = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	C.g_variant_ref(arg3)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_platformData)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	iface.ChangeActionStateFull(_actionName, _value, _platformData)
}

func wrapRemoteActionGroup(obj *coreglib.Object) *RemoteActionGroup {
	return &RemoteActionGroup{
		ActionGroup: ActionGroup{
			Object: obj,
		},
	}
}

func marshalRemoteActionGroup(p uintptr) (interface{}, error) {
	return wrapRemoteActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateActionFull activates the remote action.
//
// This is the same as g_action_group_activate_action() except that it allows
// for provision of "platform data" to be sent along with the activation
// request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter (optional): optional parameter to the activation.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ActivateActionFull(actionName string, parameter, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C.g_remote_action_group_activate_action_full(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
	runtime.KeepAlive(platformData)
}

// ChangeActionStateFull changes the state of a remote action.
//
// This is the same as g_action_group_change_action_state() except that it
// allows for provision of "platform data" to be sent along with the state
// change request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to change the state of.
//    - value: new requested value for the state.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ChangeActionStateFull(actionName string, value, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C.g_remote_action_group_change_action_state_full(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
	runtime.KeepAlive(platformData)
}

// RemoteActionGroupInterface: virtual function table for ActionGroup.
//
// An instance of this type is always passed by reference.
type RemoteActionGroupInterface struct {
	*remoteActionGroupInterface
}

// remoteActionGroupInterface is the struct that's finalized.
type remoteActionGroupInterface struct {
	native *C.GRemoteActionGroupInterface
}
