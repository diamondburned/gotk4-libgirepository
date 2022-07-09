// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeRecentAction returns the GType for the type RecentAction.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRecentAction() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "RecentAction").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRecentAction)
	return gtype
}

// RecentActionOverrider contains methods that are overridable.
type RecentActionOverrider interface {
}

// RecentAction represents a list of recently used files, which can be shown by
// widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction struct {
	_ [0]func() // equal guard
	Action

	*coreglib.Object
	RecentChooser
}

var (
	_ coreglib.Objector = (*RecentAction)(nil)
)

func classInitRecentActioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRecentAction(obj *coreglib.Object) *RecentAction {
	return &RecentAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentAction(p uintptr) (interface{}, error) {
	return wrapRecentAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRecentAction creates a new RecentAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//
// The function returns the following values:
//
//    - recentAction: newly created RecentAction.
//
func NewRecentAction(name, label, tooltip, stockId string) *RecentAction {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[0]))
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}
	if tooltip != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_args[2]))
	}
	if stockId != "" {
		*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_args[3]))
	}

	_gret := girepository.MustFind("Gtk", "RecentAction").InvokeMethod("new_RecentAction", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)

	var _recentAction *RecentAction // out

	_recentAction = wrapRecentAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _recentAction
}

// NewRecentActionForManager creates a new RecentAction object. To add the
// action to a ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name: unique name for the action.
//    - label (optional) displayed in menu items and on buttons, or NULL.
//    - tooltip (optional) for the action, or NULL.
//    - stockId (optional): stock icon to display in widgets representing the
//      action, or NULL.
//    - manager (optional) or NULL for using the default RecentManager.
//
// The function returns the following values:
//
//    - recentAction: newly created RecentAction.
//
func NewRecentActionForManager(name, label, tooltip, stockId string, manager *RecentManager) *RecentAction {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[0]))
	if label != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_args[1]))
	}
	if tooltip != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_args[2]))
	}
	if stockId != "" {
		*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_args[3]))
	}
	if manager != nil {
		*(**C.void)(unsafe.Pointer(&_args[4])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	}

	_gret := girepository.MustFind("Gtk", "RecentAction").InvokeMethod("new_RecentAction_for_manager", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(manager)

	var _recentAction *RecentAction // out

	_recentAction = wrapRecentAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _recentAction
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if numbers should be shown.
//
func (action *RecentAction) ShowNumbers() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_gret := girepository.MustFind("Gtk", "RecentAction").InvokeMethod("get_show_numbers", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetShowNumbers sets whether a number should be added to the items shown by
// the widgets representing action. The numbers are shown to provide a unique
// character for a mnemonic to be used inside the menu item's label. Only the
// first ten items get a number to avoid clashes.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - showNumbers: TRUE if the shown items should be numbered.
//
func (action *RecentAction) SetShowNumbers(showNumbers bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if showNumbers {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "RecentAction").InvokeMethod("set_show_numbers", _args[:], nil)

	runtime.KeepAlive(action)
	runtime.KeepAlive(showNumbers)
}
