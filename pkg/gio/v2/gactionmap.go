// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_action_map_get_type()), F: marshalActionMap},
	})
}

// ActionMapOverrider contains methods that are overridable. This
// interface is a subset of the interface ActionMap.
type ActionMapOverrider interface {
	// AddAction adds an action to the @action_map.
	//
	// If the action map already contains an action with the same name as
	// @action then the old action is dropped from the action map.
	//
	// The action map takes its own reference on @action.
	AddAction(action Action)
	// LookupAction looks up the action with the name @action_name in
	// @action_map.
	//
	// If no such action exists, returns nil.
	LookupAction(actionName string) Action
	// RemoveAction removes the named action from the action map.
	//
	// If no action of this name is in the map then nothing happens.
	RemoveAction(actionName string)
}

// ActionMap: the GActionMap interface is implemented by Group implementations
// that operate by containing a number of named #GAction instances, such as
// ActionGroup.
//
// One useful application of this interface is to map the names of actions from
// various action groups to unique, prefixed names (e.g. by prepending "app." or
// "win."). This is the motivation for the 'Map' part of the interface name.
type ActionMap interface {
	gextras.Objector
	ActionMapOverrider
}

// actionMap implements the ActionMap interface.
type actionMap struct {
	gextras.Objector
}

var _ ActionMap = (*actionMap)(nil)

// WrapActionMap wraps a GObject to a type that implements interface
// ActionMap. It is primarily used internally.
func WrapActionMap(obj *externglib.Object) ActionMap {
	return ActionMap{
		Objector: obj,
	}
}

func marshalActionMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActionMap(obj), nil
}

// AddAction adds an action to the @action_map.
//
// If the action map already contains an action with the same name as
// @action then the old action is dropped from the action map.
//
// The action map takes its own reference on @action.
func (a actionMap) AddAction(action Action) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.GAction    // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_action_map_add_action(_arg0, _arg1)
}

// LookupAction looks up the action with the name @action_name in
// @action_map.
//
// If no such action exists, returns nil.
func (a actionMap) LookupAction(actionName string) Action {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out
	var _cret *C.GAction    // in

	_arg0 = (*C.GActionMap)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_map_lookup_action(_arg0, _arg1)

	var _action Action // out

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Action)

	return _action
}

// RemoveAction removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
func (a actionMap) RemoveAction(actionName string) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_action_map_remove_action(_arg0, _arg1)
}
