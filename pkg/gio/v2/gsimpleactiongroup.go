// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gsimpleactiongroup.go.
var GTypeSimpleActionGroup = coreglib.Type(C.g_simple_action_group_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSimpleActionGroup, F: marshalSimpleActionGroup},
	})
}

// SimpleActionGroupOverrider contains methods that are overridable.
type SimpleActionGroupOverrider interface {
}

// SimpleActionGroup is a hash table filled with #GAction objects, implementing
// the Group and Map interfaces.
type SimpleActionGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object

	ActionGroup
	ActionMap
}

var (
	_ coreglib.Objector = (*SimpleActionGroup)(nil)
)

func classInitSimpleActionGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSimpleActionGroup(obj *coreglib.Object) *SimpleActionGroup {
	return &SimpleActionGroup{
		Object: obj,
		ActionGroup: ActionGroup{
			Object: obj,
		},
		ActionMap: ActionMap{
			Object: obj,
		},
	}
}

func marshalSimpleActionGroup(p uintptr) (interface{}, error) {
	return wrapSimpleActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSimpleActionGroup creates a new, empty, ActionGroup.
//
// The function returns the following values:
//
//    - simpleActionGroup: new ActionGroup.
//
func NewSimpleActionGroup() *SimpleActionGroup {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gio", "SimpleActionGroup").InvokeMethod("new_SimpleActionGroup", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _simpleActionGroup *SimpleActionGroup // out

	_simpleActionGroup = wrapSimpleActionGroup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleActionGroup
}

// Insert adds an action to the action group.
//
// If the action group already contains an action with the same name as action
// then the old action is dropped from the group.
//
// The action group takes its own reference on action.
//
// Deprecated: Use g_action_map_add_action().
//
// The function takes the following parameters:
//
//    - action: #GAction.
//
func (simple *SimpleActionGroup) Insert(action Actioner) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "SimpleActionGroup").InvokeMethod("insert", _args[:], nil)

	runtime.KeepAlive(simple)
	runtime.KeepAlive(action)
}

// Lookup looks up the action with the name action_name in the group.
//
// If no such action exists, returns NULL.
//
// Deprecated: Use g_action_map_lookup_action().
//
// The function takes the following parameters:
//
//    - actionName: name of an action.
//
// The function returns the following values:
//
//    - action or NULL.
//
func (simple *SimpleActionGroup) Lookup(actionName string) *Action {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "SimpleActionGroup").InvokeMethod("lookup", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(simple)
	runtime.KeepAlive(actionName)

	var _action *Action // out

	_action = wrapAction(coreglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// Remove removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
//
// Deprecated: Use g_action_map_remove_action().
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
func (simple *SimpleActionGroup) Remove(actionName string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gio", "SimpleActionGroup").InvokeMethod("remove", _args[:], nil)

	runtime.KeepAlive(simple)
	runtime.KeepAlive(actionName)
}
