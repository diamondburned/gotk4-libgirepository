// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GVariant* _gotk4_gio2_ActionInterface_get_state(void*);
// extern GVariant* _gotk4_gio2_ActionInterface_get_state_hint(void*);
// extern GVariantType* _gotk4_gio2_ActionInterface_get_parameter_type(void*);
// extern GVariantType* _gotk4_gio2_ActionInterface_get_state_type(void*);
// extern gboolean _gotk4_gio2_ActionInterface_get_enabled(void*);
// extern gchar* _gotk4_gio2_ActionInterface_get_name(void*);
// extern void _gotk4_gio2_ActionInterface_activate(void*, void*);
// extern void _gotk4_gio2_ActionInterface_change_state(void*, void*);
import "C"

// glib.Type values for gaction.go.
var GTypeAction = coreglib.Type(C.g_action_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAction, F: marshalAction},
	})
}

// ActionOverrider contains methods that are overridable.
type ActionOverrider interface {
	// Activate activates the action.
	//
	// parameter must be the correct type of parameter for the action (ie: the
	// parameter type given at construction time). If the parameter type was
	// NULL then parameter must also be NULL.
	//
	// If the parameter GVariant is floating, it is consumed.
	//
	// The function takes the following parameters:
	//
	//    - parameter (optional) to the activation.
	//
	Activate(parameter *glib.Variant)
	// ChangeState: request for the state of action to be changed to value.
	//
	// The action must be stateful and value must be of the correct type. See
	// g_action_get_state_type().
	//
	// This call merely requests a change. The action may refuse to change its
	// state or may change its state to something other than value. See
	// g_action_get_state_hint().
	//
	// If the value GVariant is floating, it is consumed.
	//
	// The function takes the following parameters:
	//
	//    - value: new state.
	//
	ChangeState(value *glib.Variant)
	// Enabled checks if action is currently enabled.
	//
	// An action must be enabled in order to be activated or in order to have
	// its state changed from outside callers.
	//
	// The function returns the following values:
	//
	//    - ok: whether the action is enabled.
	//
	Enabled() bool
	// Name queries the name of action.
	//
	// The function returns the following values:
	//
	//    - utf8: name of the action.
	//
	Name() string
	// ParameterType queries the type of the parameter that must be given when
	// activating action.
	//
	// When activating the action using g_action_activate(), the #GVariant given
	// to that function must be of the type returned by this function.
	//
	// In the case that this function returns NULL, you must not give any
	// #GVariant, but NULL instead.
	//
	// The function returns the following values:
	//
	//    - variantType (optional): parameter type.
	//
	ParameterType() *glib.VariantType
	// State queries the current state of action.
	//
	// If the action is not stateful then NULL will be returned. If the action
	// is stateful then the type of the return value is the type given by
	// g_action_get_state_type().
	//
	// The return value (if non-NULL) should be freed with g_variant_unref()
	// when it is no longer required.
	//
	// The function returns the following values:
	//
	//    - variant (optional): current state of the action.
	//
	State() *glib.Variant
	// StateHint requests a hint about the valid range of values for the state
	// of action.
	//
	// If NULL is returned it either means that the action is not stateful or
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
	// The return value (if non-NULL) should be freed with g_variant_unref()
	// when it is no longer required.
	//
	// The function returns the following values:
	//
	//    - variant (optional): state range hint.
	//
	StateHint() *glib.Variant
	// StateType queries the type of the state of action.
	//
	// If the action is stateful (e.g. created with
	// g_simple_action_new_stateful()) then this function returns the Type of
	// the state. This is the type of the initial value given as the state. All
	// calls to g_action_change_state() must give a #GVariant of this type and
	// g_action_get_state() will return a #GVariant of the same type.
	//
	// If the action is not stateful (e.g. created with g_simple_action_new())
	// then this function will return NULL. In that case, g_action_get_state()
	// will return NULL and you must not call g_action_change_state().
	//
	// The function returns the following values:
	//
	//    - variantType (optional): state type, if the action is stateful.
	//
	StateType() *glib.VariantType
}

// Action represents a single named action.
//
// The main interface to an action is that it can be activated with
// g_action_activate(). This results in the 'activate' signal being emitted. An
// activation has a #GVariant parameter (which may be NULL). The correct type
// for the parameter is determined by a static parameter type (which is given at
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
//
// Action wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Action struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Action)(nil)
)

// Actioner describes Action's interface methods.
type Actioner interface {
	coreglib.Objector

	// Activate activates the action.
	Activate(parameter *glib.Variant)
	// ChangeState: request for the state of action to be changed to value.
	ChangeState(value *glib.Variant)
	// Enabled checks if action is currently enabled.
	Enabled() bool
	// Name queries the name of action.
	Name() string
	// ParameterType queries the type of the parameter that must be given when
	// activating action.
	ParameterType() *glib.VariantType
	// State queries the current state of action.
	State() *glib.Variant
	// StateHint requests a hint about the valid range of values for the state
	// of action.
	StateHint() *glib.Variant
	// StateType queries the type of the state of action.
	StateType() *glib.VariantType
}

var _ Actioner = (*Action)(nil)

func ifaceInitActioner(gifacePtr, data C.gpointer) {
	iface := (*C.GActionInterface)(unsafe.Pointer(gifacePtr))
	iface.activate = (*[0]byte)(C._gotk4_gio2_ActionInterface_activate)
	iface.change_state = (*[0]byte)(C._gotk4_gio2_ActionInterface_change_state)
	iface.get_enabled = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_enabled)
	iface.get_name = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_name)
	iface.get_parameter_type = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_parameter_type)
	iface.get_state = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_state)
	iface.get_state_hint = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_state_hint)
	iface.get_state_type = (*[0]byte)(C._gotk4_gio2_ActionInterface_get_state_type)
}

//export _gotk4_gio2_ActionInterface_activate
func _gotk4_gio2_ActionInterface_activate(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _parameter *glib.Variant // out

	if arg1 != nil {
		_parameter = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		C.g_variant_ref(arg1)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_parameter)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	iface.Activate(_parameter)
}

//export _gotk4_gio2_ActionInterface_change_state
func _gotk4_gio2_ActionInterface_change_state(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	var _value *glib.Variant // out

	_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_variant_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	iface.ChangeState(_value)
}

//export _gotk4_gio2_ActionInterface_get_enabled
func _gotk4_gio2_ActionInterface_get_enabled(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	ok := iface.Enabled()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_ActionInterface_get_name
func _gotk4_gio2_ActionInterface_get_name(arg0 *C.void) (cret *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	utf8 := iface.Name()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gio2_ActionInterface_get_parameter_type
func _gotk4_gio2_ActionInterface_get_parameter_type(arg0 *C.void) (cret *C.GVariantType) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	variantType := iface.ParameterType()

	if variantType != nil {
		cret = (*C.void)(gextras.StructNative(unsafe.Pointer(variantType)))
	}

	return cret
}

//export _gotk4_gio2_ActionInterface_get_state
func _gotk4_gio2_ActionInterface_get_state(arg0 *C.void) (cret *C.GVariant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	variant := iface.State()

	if variant != nil {
		cret = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))
	}

	return cret
}

//export _gotk4_gio2_ActionInterface_get_state_hint
func _gotk4_gio2_ActionInterface_get_state_hint(arg0 *C.void) (cret *C.GVariant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	variant := iface.StateHint()

	if variant != nil {
		cret = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))
	}

	return cret
}

//export _gotk4_gio2_ActionInterface_get_state_type
func _gotk4_gio2_ActionInterface_get_state_type(arg0 *C.void) (cret *C.GVariantType) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ActionOverrider)

	variantType := iface.StateType()

	if variantType != nil {
		cret = (*C.void)(gextras.StructNative(unsafe.Pointer(variantType)))
	}

	return cret
}

func wrapAction(obj *coreglib.Object) *Action {
	return &Action{
		Object: obj,
	}
}

func marshalAction(p uintptr) (interface{}, error) {
	return wrapAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Activate activates the action.
//
// parameter must be the correct type of parameter for the action (ie: the
// parameter type given at construction time). If the parameter type was NULL
// then parameter must also be NULL.
//
// If the parameter GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - parameter (optional) to the activation.
//
func (action *Action) Activate(parameter *glib.Variant) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if parameter != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(parameter)))
	}

	runtime.KeepAlive(action)
	runtime.KeepAlive(parameter)
}

// ChangeState: request for the state of action to be changed to value.
//
// The action must be stateful and value must be of the correct type. See
// g_action_get_state_type().
//
// This call merely requests a change. The action may refuse to change its state
// or may change its state to something other than value. See
// g_action_get_state_hint().
//
// If the value GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - value: new state.
//
func (action *Action) ChangeState(value *glib.Variant) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(value)))

	runtime.KeepAlive(action)
	runtime.KeepAlive(value)
}

// Enabled checks if action is currently enabled.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
//
// The function returns the following values:
//
//    - ok: whether the action is enabled.
//
func (action *Action) Enabled() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Name queries the name of action.
//
// The function returns the following values:
//
//    - utf8: name of the action.
//
func (action *Action) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ParameterType queries the type of the parameter that must be given when
// activating action.
//
// When activating the action using g_action_activate(), the #GVariant given to
// that function must be of the type returned by this function.
//
// In the case that this function returns NULL, you must not give any #GVariant,
// but NULL instead.
//
// The function returns the following values:
//
//    - variantType (optional): parameter type.
//
func (action *Action) ParameterType() *glib.VariantType {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _variantType *glib.VariantType // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// State queries the current state of action.
//
// If the action is not stateful then NULL will be returned. If the action is
// stateful then the type of the return value is the type given by
// g_action_get_state_type().
//
// The return value (if non-NULL) should be freed with g_variant_unref() when it
// is no longer required.
//
// The function returns the following values:
//
//    - variant (optional): current state of the action.
//
func (action *Action) State() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _variant *glib.Variant // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// StateHint requests a hint about the valid range of values for the state of
// action.
//
// If NULL is returned it either means that the action is not stateful or that
// there is no hint about the valid range of values for the state of the action.
//
// If a #GVariant array is returned then each item in the array is a possible
// value for the state. If a #GVariant pair (ie: two-tuple) is returned then the
// tuple specifies the inclusive lower and upper bound of valid values for the
// state.
//
// In any case, the information is merely a hint. It may be possible to have a
// state value outside of the hinted range and setting a value within the range
// may fail.
//
// The return value (if non-NULL) should be freed with g_variant_unref() when it
// is no longer required.
//
// The function returns the following values:
//
//    - variant (optional): state range hint.
//
func (action *Action) StateHint() *glib.Variant {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _variant *glib.Variant // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// StateType queries the type of the state of action.
//
// If the action is stateful (e.g. created with g_simple_action_new_stateful())
// then this function returns the Type of the state. This is the type of the
// initial value given as the state. All calls to g_action_change_state() must
// give a #GVariant of this type and g_action_get_state() will return a
// #GVariant of the same type.
//
// If the action is not stateful (e.g. created with g_simple_action_new()) then
// this function will return NULL. In that case, g_action_get_state() will
// return NULL and you must not call g_action_change_state().
//
// The function returns the following values:
//
//    - variantType (optional): state type, if the action is stateful.
//
func (action *Action) StateType() *glib.VariantType {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(action)

	var _variantType *glib.VariantType // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_variantType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _variantType
}

// ActionNameIsValid checks if action_name is valid.
//
// action_name is valid if it consists only of alphanumeric characters, plus '-'
// and '.'. The empty string is not a valid action name.
//
// It is an error to call this function with a non-utf8 action_name. action_name
// must not be NULL.
//
// The function takes the following parameters:
//
//    - actionName: potential action name.
//
// The function returns the following values:
//
//    - ok: TRUE if action_name is valid.
//
func ActionNameIsValid(actionName string) bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gio", "name_is_valid").Invoke(_args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionName)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ActionParseDetailedName parses a detailed action name into its separate name
// and target components.
//
// Detailed action names can have three formats.
//
// The first format is used to represent an action name with no target value and
// consists of just an action name containing no whitespace nor the characters
// ':', '(' or ')'. For example: "app.action".
//
// The second format is used to represent an action with a target value that is
// a non-empty string consisting only of alphanumerics, plus '-' and '.'. In
// that case, the action name and target value are separated by a double colon
// ("::"). For example: "app.action::target".
//
// The third format is used to represent an action with any type of target
// value, including strings. The target value follows the action name,
// surrounded in parens. For example: "app.action(42)". The target value is
// parsed using g_variant_parse(). If a tuple-typed value is desired, it must be
// specified in the same way, resulting in two sets of parens, for example:
// "app.action((1,2,3))". A string target can be specified this way as well:
// "app.action('target')". For strings, this third format must be used if *
// target value is empty or contains characters other than alphanumerics, '-'
// and '.'.
//
// The function takes the following parameters:
//
//    - detailedName: detailed action name.
//
// The function returns the following values:
//
//    - actionName: action name.
//    - targetValue: target value, or NULL for no target.
//
func ActionParseDetailedName(detailedName string) (string, *glib.Variant, error) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(detailedName)))
	defer C.free(unsafe.Pointer(_args[0]))

	girepository.MustFind("Gio", "parse_detailed_name").Invoke(_args[:], _outs[:])

	runtime.KeepAlive(detailedName)

	var _actionName string         // out
	var _targetValue *glib.Variant // out
	var _goerr error               // out

	_actionName = C.GoString((*C.gchar)(unsafe.Pointer(_outs[0])))
	defer C.free(unsafe.Pointer(_outs[0]))
	_targetValue = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_outs[1])))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetValue)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _actionName, _targetValue, _goerr
}

// ActionPrintDetailedName formats a detailed action name from action_name and
// target_value.
//
// It is an error to call this function with an invalid action name.
//
// This function is the opposite of g_action_parse_detailed_name(). It will
// produce a string that can be parsed back to the action_name and target_value
// by that function.
//
// See that function for the types of strings that will be printed by this
// function.
//
// The function takes the following parameters:
//
//    - actionName: valid action name.
//    - targetValue (optional) target value, or NULL.
//
// The function returns the following values:
//
//    - utf8: detailed format string.
//
func ActionPrintDetailedName(actionName string, targetValue *glib.Variant) string {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_args[0]))
	if targetValue != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(targetValue)))
	}

	_gret := girepository.MustFind("Gio", "print_detailed_name").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(actionName)
	runtime.KeepAlive(targetValue)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
