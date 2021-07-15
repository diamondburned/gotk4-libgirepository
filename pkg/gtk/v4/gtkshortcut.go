// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_get_type()), F: marshalShortcuter},
	})
}

// Shortcuter describes Shortcut's methods.
type Shortcuter interface {
	// Action gets the action that is activated by this shortcut.
	Action() *ShortcutAction
	// Arguments gets the arguments that are passed when activating the
	// shortcut.
	Arguments() *glib.Variant
	// Trigger gets the trigger used to trigger self.
	Trigger() *ShortcutTrigger
	// SetAction sets the new action for self to be action.
	SetAction(action ShortcutActioner)
	// SetArguments sets the arguments to pass when activating the shortcut.
	SetArguments(args *glib.Variant)
	// SetTrigger sets the new trigger for self to be trigger.
	SetTrigger(trigger ShortcutTriggerer)
}

// Shortcut: GtkShortcut describes a keyboard shortcut.
//
// It contains a description of how to trigger the shortcut via a
// gtk.ShortcutTrigger and a way to activate the shortcut on a widget via a
// gtk.ShortcutAction.
//
// The actual work is usually done via gtk.ShortcutController, which decides if
// and when to activate a shortcut. Using that controller directly however is
// rarely necessary as various higher level convenience APIs exist on Widgets
// that make it easier to use shortcuts in GTK.
//
// GtkShortcut does provide functionality to make it easy for users to work with
// shortcuts, either by providing informational strings for display purposes or
// by allowing shortcuts to be configured.
type Shortcut struct {
	*externglib.Object
}

var (
	_ Shortcuter      = (*Shortcut)(nil)
	_ gextras.Nativer = (*Shortcut)(nil)
)

func wrapShortcut(obj *externglib.Object) *Shortcut {
	return &Shortcut{
		Object: obj,
	}
}

func marshalShortcuter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcut(obj), nil
}

// NewShortcut creates a new GtkShortcut that is triggered by trigger and then
// activates action.
func NewShortcut(trigger ShortcutTriggerer, action ShortcutActioner) *Shortcut {
	var _arg1 *C.GtkShortcutTrigger // out
	var _arg2 *C.GtkShortcutAction  // out
	var _cret *C.GtkShortcut        // in

	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer((trigger).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkShortcutAction)(unsafe.Pointer((action).(gextras.Nativer).Native()))

	_cret = C.gtk_shortcut_new(_arg1, _arg2)

	var _shortcut *Shortcut // out

	_shortcut = wrapShortcut(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcut
}

// Action gets the action that is activated by this shortcut.
func (self *Shortcut) Action() *ShortcutAction {
	var _arg0 *C.GtkShortcut       // out
	var _cret *C.GtkShortcutAction // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_action(_arg0)

	var _shortcutAction *ShortcutAction // out

	_shortcutAction = wrapShortcutAction(externglib.Take(unsafe.Pointer(_cret)))

	return _shortcutAction
}

// Arguments gets the arguments that are passed when activating the shortcut.
func (self *Shortcut) Arguments() *glib.Variant {
	var _arg0 *C.GtkShortcut // out
	var _cret *C.GVariant    // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_arguments(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variant
}

// Trigger gets the trigger used to trigger self.
func (self *Shortcut) Trigger() *ShortcutTrigger {
	var _arg0 *C.GtkShortcut        // out
	var _cret *C.GtkShortcutTrigger // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_trigger(_arg0)

	var _shortcutTrigger *ShortcutTrigger // out

	_shortcutTrigger = wrapShortcutTrigger(externglib.Take(unsafe.Pointer(_cret)))

	return _shortcutTrigger
}

// SetAction sets the new action for self to be action.
func (self *Shortcut) SetAction(action ShortcutActioner) {
	var _arg0 *C.GtkShortcut       // out
	var _arg1 *C.GtkShortcutAction // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkShortcutAction)(unsafe.Pointer((action).(gextras.Nativer).Native()))

	C.gtk_shortcut_set_action(_arg0, _arg1)
}

// SetArguments sets the arguments to pass when activating the shortcut.
func (self *Shortcut) SetArguments(args *glib.Variant) {
	var _arg0 *C.GtkShortcut // out
	var _arg1 *C.GVariant    // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(args)))

	C.gtk_shortcut_set_arguments(_arg0, _arg1)
}

// SetTrigger sets the new trigger for self to be trigger.
func (self *Shortcut) SetTrigger(trigger ShortcutTriggerer) {
	var _arg0 *C.GtkShortcut        // out
	var _arg1 *C.GtkShortcutTrigger // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer((trigger).(gextras.Nativer).Native()))

	C.gtk_shortcut_set_trigger(_arg0, _arg1)
}
