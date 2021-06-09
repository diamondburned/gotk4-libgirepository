// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
	ActionGroup
	ActionMap

	// Insert adds an action to the action group.
	//
	// If the action group already contains an action with the same name as
	// @action then the old action is dropped from the group.
	//
	// The action group takes its own reference on @action.
	Insert(action Action)
	// Lookup looks up the action with the name @action_name in the group.
	//
	// If no such action exists, returns nil.
	Lookup(actionName string) Action
	// Remove removes the named action from the action group.
	//
	// If no action of this name is in the group then nothing happens.
	Remove(actionName string)
}

// simpleActionGroup implements the SimpleActionGroup interface.
type simpleActionGroup struct {
	gextras.Objector
	ActionGroup
	ActionMap
}

var _ SimpleActionGroup = (*simpleActionGroup)(nil)

// WrapSimpleActionGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapSimpleActionGroup(obj *externglib.Object) SimpleActionGroup {
	return SimpleActionGroup{
		Objector:    obj,
		ActionGroup: WrapActionGroup(obj),
		ActionMap:   WrapActionMap(obj),
	}
}

func marshalSimpleActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSimpleActionGroup(obj), nil
}

// NewSimpleActionGroup constructs a class SimpleActionGroup.
func NewSimpleActionGroup() SimpleActionGroup {
	var _cret C.GSimpleActionGroup

	cret = C.g_simple_action_group_new()

	var _simpleActionGroup SimpleActionGroup

	_simpleActionGroup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(SimpleActionGroup)

	return _simpleActionGroup
}

// Insert adds an action to the action group.
//
// If the action group already contains an action with the same name as
// @action then the old action is dropped from the group.
//
// The action group takes its own reference on @action.
func (s simpleActionGroup) Insert(action Action) {
	var _arg0 *C.GSimpleActionGroup
	var _arg1 *C.GAction

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_simple_action_group_insert(_arg0, _arg1)
}

// Lookup looks up the action with the name @action_name in the group.
//
// If no such action exists, returns nil.
func (s simpleActionGroup) Lookup(actionName string) Action {
	var _arg0 *C.GSimpleActionGroup
	var _arg1 *C.gchar

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GAction

	cret = C.g_simple_action_group_lookup(_arg0, _arg1)

	var _action Action

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Action)

	return _action
}

// Remove removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
func (s simpleActionGroup) Remove(actionName string) {
	var _arg0 *C.GSimpleActionGroup
	var _arg1 *C.gchar

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_action_group_remove(_arg0, _arg1)
}
