// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_action_get_type()), F: marshalAction},
	})
}

// ActionOverrider contains methods that are overridable. This
// interface is a subset of the interface Action.
type ActionOverrider interface {
	// Activate activates the action.
	//
	// @parameter must be the correct type of parameter for the action (ie: the
	// parameter type given at construction time). If the parameter type was nil
	// then @parameter must also be nil.
	//
	// If the @parameter GVariant is floating, it is consumed.
	Activate(parameter *glib.Variant)
	// ChangeState: request for the state of @action to be changed to @value.
	//
	// The action must be stateful and @value must be of the correct type. See
	// g_action_get_state_type().
	//
	// This call merely requests a change. The action may refuse to change its
	// state or may change its state to something other than @value. See
	// g_action_get_state_hint().
	//
	// If the @value GVariant is floating, it is consumed.
	ChangeState(value *glib.Variant)
	// Enabled checks if @action is currently enabled.
	//
	// An action must be enabled in order to be activated or in order to have
	// its state changed from outside callers.
	Enabled() bool
	// Name queries the name of @action.
	Name() string
	// ParameterType queries the type of the parameter that must be given when
	// activating @action.
	//
	// When activating the action using g_action_activate(), the #GVariant given
	// to that function must be of the type returned by this function.
	//
	// In the case that this function returns nil, you must not give any
	// #GVariant, but nil instead.
	ParameterType() *glib.VariantType
	// State queries the current state of @action.
	//
	// If the action is not stateful then nil will be returned. If the action is
	// stateful then the type of the return value is the type given by
	// g_action_get_state_type().
	//
	// The return value (if non-nil) should be freed with g_variant_unref() when
	// it is no longer required.
	State() *glib.Variant
	// StateHint requests a hint about the valid range of values for the state
	// of @action.
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
	StateHint() *glib.Variant
	// StateType queries the type of the state of @action.
	//
	// If the action is stateful (e.g. created with
	// g_simple_action_new_stateful()) then this function returns the Type of
	// the state. This is the type of the initial value given as the state. All
	// calls to g_action_change_state() must give a #GVariant of this type and
	// g_action_get_state() will return a #GVariant of the same type.
	//
	// If the action is not stateful (e.g. created with g_simple_action_new())
	// then this function will return nil. In that case, g_action_get_state()
	// will return nil and you must not call g_action_change_state().
	StateType() *glib.VariantType
}

// Action represents a single named action.
//
// The main interface to an action is that it can be activated with
// g_action_activate(). This results in the 'activate' signal being emitted. An
// activation has a #GVariant parameter (which may be nil). The correct type for
// the parameter is determined by a static parameter type (which is given at
// construction time).
//
// An action may optionally have a state, in which case the state may be set
// with g_action_change_state(). This call takes a #GVariant. The correct type
// for the state is determined by a static state type (which is given at
// construction time).
//
// The state may have a hint associated with it, specifying its valid range.
//
// #GAction is merely the interface to the concept of an action, as described
// above. Various implementations of actions exist, including Action.
//
// In all cases, the implementing class is responsible for storing the name of
// the action, the parameter type, the enabled state, the optional state type
// and the state and emitting the appropriate signals when these change. The
// implementor is responsible for filtering calls to g_action_activate() and
// g_action_change_state() for type safety and for the state being enabled.
//
// Probably the only useful thing to do with a #GAction is to put it inside of a
// ActionGroup.
type Action interface {
	gextras.Objector
	ActionOverrider
}

// action implements the Action interface.
type action struct {
	gextras.Objector
}

var _ Action = (*action)(nil)

// WrapAction wraps a GObject to a type that implements interface
// Action. It is primarily used internally.
func WrapAction(obj *externglib.Object) Action {
	return Action{
		Objector: obj,
	}
}

func marshalAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAction(obj), nil
}

// Activate activates the action.
//
// @parameter must be the correct type of parameter for the action (ie: the
// parameter type given at construction time). If the parameter type was nil
// then @parameter must also be nil.
//
// If the @parameter GVariant is floating, it is consumed.
func (a action) Activate(parameter *glib.Variant) {
	var _arg0 *C.GAction
	var _arg1 *C.GVariant

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GVariant)(unsafe.Pointer(parameter.Native()))

	C.g_action_activate(_arg0, _arg1)
}

// ChangeState: request for the state of @action to be changed to @value.
//
// The action must be stateful and @value must be of the correct type. See
// g_action_get_state_type().
//
// This call merely requests a change. The action may refuse to change its
// state or may change its state to something other than @value. See
// g_action_get_state_hint().
//
// If the @value GVariant is floating, it is consumed.
func (a action) ChangeState(value *glib.Variant) {
	var _arg0 *C.GAction
	var _arg1 *C.GVariant

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	C.g_action_change_state(_arg0, _arg1)
}

// Enabled checks if @action is currently enabled.
//
// An action must be enabled in order to be activated or in order to have
// its state changed from outside callers.
func (a action) Enabled() bool {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean

	cret = C.g_action_get_enabled(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Name queries the name of @action.
func (a action) Name() string {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret *C.gchar

	cret = C.g_action_get_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ParameterType queries the type of the parameter that must be given when
// activating @action.
//
// When activating the action using g_action_activate(), the #GVariant given
// to that function must be of the type returned by this function.
//
// In the case that this function returns nil, you must not give any
// #GVariant, but nil instead.
func (a action) ParameterType() *glib.VariantType {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret *C.GVariantType

	cret = C.g_action_get_parameter_type(_arg0)

	var _variantType *glib.VariantType

	_variantType = glib.WrapVariantType(unsafe.Pointer(_cret))

	return _variantType
}

// State queries the current state of @action.
//
// If the action is not stateful then nil will be returned. If the action is
// stateful then the type of the return value is the type given by
// g_action_get_state_type().
//
// The return value (if non-nil) should be freed with g_variant_unref() when
// it is no longer required.
func (a action) State() *glib.Variant {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret *C.GVariant

	cret = C.g_action_get_state(_arg0)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _variant
}

// StateHint requests a hint about the valid range of values for the state
// of @action.
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
func (a action) StateHint() *glib.Variant {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret *C.GVariant

	cret = C.g_action_get_state_hint(_arg0)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _variant
}

// StateType queries the type of the state of @action.
//
// If the action is stateful (e.g. created with
// g_simple_action_new_stateful()) then this function returns the Type of
// the state. This is the type of the initial value given as the state. All
// calls to g_action_change_state() must give a #GVariant of this type and
// g_action_get_state() will return a #GVariant of the same type.
//
// If the action is not stateful (e.g. created with g_simple_action_new())
// then this function will return nil. In that case, g_action_get_state()
// will return nil and you must not call g_action_change_state().
func (a action) StateType() *glib.VariantType {
	var _arg0 *C.GAction

	_arg0 = (*C.GAction)(unsafe.Pointer(a.Native()))

	var _cret *C.GVariantType

	cret = C.g_action_get_state_type(_arg0)

	var _variantType *glib.VariantType

	_variantType = glib.WrapVariantType(unsafe.Pointer(_cret))

	return _variantType
}
