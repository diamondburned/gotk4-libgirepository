// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern GVariant* _gotk4_gtk3_ActionableInterface_get_action_target_value(GtkActionable*);
// extern gchar* _gotk4_gtk3_ActionableInterface_get_action_name(GtkActionable*);
// extern void _gotk4_gtk3_ActionableInterface_set_action_name(GtkActionable*, gchar*);
// extern void _gotk4_gtk3_ActionableInterface_set_action_target_value(GtkActionable*, GVariant*);
import "C"

// GType values.
var (
	GTypeActionable = coreglib.Type(C.gtk_actionable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionable, F: marshalActionable},
	})
}

// ActionableOverrider contains methods that are overridable.
type ActionableOverrider interface {
	// ActionName gets the action name for actionable.
	//
	// See gtk_actionable_set_action_name() for more information.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): action name, or NULL if none is set.
	//
	ActionName() string
	// ActionTargetValue gets the current target value of actionable.
	//
	// See gtk_actionable_set_action_target_value() for more information.
	//
	// The function returns the following values:
	//
	//    - variant: current target value.
	//
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated. If action_name is NULL then the widget will be
	// unassociated from any previous action.
	//
	// Usually this function is used when the widget is located (or will be
	// located) within the hierarchy of a ApplicationWindow.
	//
	// Names are of the form “win.save” or “app.quit” for actions on the
	// containing ApplicationWindow or its associated Application, respectively.
	// This is the same form used for actions in the #GMenu associated with the
	// window.
	//
	// The function takes the following parameters:
	//
	//    - actionName (optional): action name, or NULL.
	//
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	//
	// If target_value is NULL then the target value is unset.
	//
	// The target value has two purposes. First, it is used as the parameter to
	// activation of the action associated with the Actionable widget. Second,
	// it is used to determine if the widget should be rendered as “active” —
	// the widget is active if the state is equal to the given target.
	//
	// Consider the example of associating a set of buttons with a #GAction with
	// string state in a typical “radio button” situation. Each button will be
	// associated with the same action, but with a different target value for
	// that action. Clicking on a particular button will activate the action
	// with the target of that button, which will typically cause the action’s
	// state to change to that value. Since the action’s state is now equal to
	// the target value of the button, the button will now be rendered as active
	// (and the other buttons, with different targets, rendered inactive).
	//
	// The function takes the following parameters:
	//
	//    - targetValue (optional) to set as the target value, or NULL.
	//
	SetActionTargetValue(targetValue *glib.Variant)
}

// Actionable: this interface provides a convenient way of associating widgets
// with actions on a ApplicationWindow or Application.
//
// It primarily consists of two properties: Actionable:action-name and
// Actionable:action-target. There are also some convenience APIs for setting
// these properties.
//
// The action will be looked up in action groups that are found among the
// widgets ancestors. Most commonly, these will be the actions with the “win.”
// or “app.” prefix that are associated with the ApplicationWindow or
// Application, but other action groups that are added with
// gtk_widget_insert_action_group() will be consulted as well.
//
// Actionable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Actionable struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Actionable)(nil)
)

// Actionabler describes Actionable's interface methods.
type Actionabler interface {
	coreglib.Objector

	// ActionName gets the action name for actionable.
	ActionName() string
	// ActionTargetValue gets the current target value of actionable.
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated.
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	SetActionTargetValue(targetValue *glib.Variant)
	// SetDetailedActionName sets the action-name and associated string target
	// value of an actionable widget.
	SetDetailedActionName(detailedActionName string)
}

var _ Actionabler = (*Actionable)(nil)

func ifaceInitActionabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkActionableInterface)(unsafe.Pointer(gifacePtr))
	iface.get_action_name = (*[0]byte)(C._gotk4_gtk3_ActionableInterface_get_action_name)
	iface.get_action_target_value = (*[0]byte)(C._gotk4_gtk3_ActionableInterface_get_action_target_value)
	iface.set_action_name = (*[0]byte)(C._gotk4_gtk3_ActionableInterface_set_action_name)
	iface.set_action_target_value = (*[0]byte)(C._gotk4_gtk3_ActionableInterface_set_action_target_value)
}

//export _gotk4_gtk3_ActionableInterface_get_action_name
func _gotk4_gtk3_ActionableInterface_get_action_name(arg0 *C.GtkActionable) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionableOverrider)

	utf8 := iface.ActionName()

	if utf8 != "" {
		cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_gtk3_ActionableInterface_get_action_target_value
func _gotk4_gtk3_ActionableInterface_get_action_target_value(arg0 *C.GtkActionable) (cret *C.GVariant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionableOverrider)

	variant := iface.ActionTargetValue()

	cret = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	return cret
}

//export _gotk4_gtk3_ActionableInterface_set_action_name
func _gotk4_gtk3_ActionableInterface_set_action_name(arg0 *C.GtkActionable, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionableOverrider)

	var _actionName string // out

	if arg1 != nil {
		_actionName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	}

	iface.SetActionName(_actionName)
}

//export _gotk4_gtk3_ActionableInterface_set_action_target_value
func _gotk4_gtk3_ActionableInterface_set_action_target_value(arg0 *C.GtkActionable, arg1 *C.GVariant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionableOverrider)

	var _targetValue *glib.Variant // out

	if arg1 != nil {
		_targetValue = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		C.g_variant_ref(arg1)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_targetValue)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	iface.SetActionTargetValue(_targetValue)
}

func wrapActionable(obj *coreglib.Object) *Actionable {
	return &Actionable{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalActionable(p uintptr) (interface{}, error) {
	return wrapActionable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActionName gets the action name for actionable.
//
// See gtk_actionable_set_action_name() for more information.
//
// The function returns the following values:
//
//    - utf8 (optional): action name, or NULL if none is set.
//
func (actionable *Actionable) ActionName() string {
	var _arg0 *C.GtkActionable // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(coreglib.InternObject(actionable).Native()))

	_cret = C.gtk_actionable_get_action_name(_arg0)
	runtime.KeepAlive(actionable)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActionTargetValue gets the current target value of actionable.
//
// See gtk_actionable_set_action_target_value() for more information.
//
// The function returns the following values:
//
//    - variant: current target value.
//
func (actionable *Actionable) ActionTargetValue() *glib.Variant {
	var _arg0 *C.GtkActionable // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(coreglib.InternObject(actionable).Native()))

	_cret = C.gtk_actionable_get_action_target_value(_arg0)
	runtime.KeepAlive(actionable)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// SetActionName specifies the name of the action with which this widget should
// be associated. If action_name is NULL then the widget will be unassociated
// from any previous action.
//
// Usually this function is used when the widget is located (or will be located)
// within the hierarchy of a ApplicationWindow.
//
// Names are of the form “win.save” or “app.quit” for actions on the containing
// ApplicationWindow or its associated Application, respectively. This is the
// same form used for actions in the #GMenu associated with the window.
//
// The function takes the following parameters:
//
//    - actionName (optional): action name, or NULL.
//
func (actionable *Actionable) SetActionName(actionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(coreglib.InternObject(actionable).Native()))
	if actionName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_actionable_set_action_name(_arg0, _arg1)
	runtime.KeepAlive(actionable)
	runtime.KeepAlive(actionName)
}

// SetActionTargetValue sets the target value of an actionable widget.
//
// If target_value is NULL then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter to
// activation of the action associated with the Actionable widget. Second, it is
// used to determine if the widget should be rendered as “active” — the widget
// is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a #GAction with
// string state in a typical “radio button” situation. Each button will be
// associated with the same action, but with a different target value for that
// action. Clicking on a particular button will activate the action with the
// target of that button, which will typically cause the action’s state to
// change to that value. Since the action’s state is now equal to the target
// value of the button, the button will now be rendered as active (and the other
// buttons, with different targets, rendered inactive).
//
// The function takes the following parameters:
//
//    - targetValue (optional) to set as the target value, or NULL.
//
func (actionable *Actionable) SetActionTargetValue(targetValue *glib.Variant) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(coreglib.InternObject(actionable).Native()))
	if targetValue != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(targetValue)))
	}

	C.gtk_actionable_set_action_target_value(_arg0, _arg1)
	runtime.KeepAlive(actionable)
	runtime.KeepAlive(targetValue)
}

// SetDetailedActionName sets the action-name and associated string target value
// of an actionable widget.
//
// detailed_action_name is a string in the format accepted by
// g_action_parse_detailed_name().
//
// (Note that prior to version 3.22.25, this function is only usable for actions
// with a simple "s" target, and detailed_action_name must be of the form
// "action::target" where action is the action name and target is the string to
// use as the target.).
//
// The function takes the following parameters:
//
//    - detailedActionName: detailed action name.
//
func (actionable *Actionable) SetDetailedActionName(detailedActionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer(coreglib.InternObject(actionable).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_actionable_set_detailed_action_name(_arg0, _arg1)
	runtime.KeepAlive(actionable)
	runtime.KeepAlive(detailedActionName)
}
