// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_value_type_get_type()), F: marshalValueType},
		{T: externglib.Type(C.atk_value_get_type()), F: marshalValue},
	})
}

// ValueType: default types for a given value. Those are defined in order to
// easily get localized strings to describe a given value or a given subrange,
// using atk_value_type_get_localized_name().
type ValueType int

const (
	ValueVeryWeak ValueType = iota
	ValueWeak
	ValueAcceptable
	ValueStrong
	ValueVeryStrong
	ValueVeryLow
	ValueLow
	ValueMedium
	ValueHigh
	ValueVeryHigh
	ValueVeryBad
	ValueBad
	ValueGood
	ValueVeryGood
	ValueBest
	ValueLastDefined
)

func marshalValueType(p uintptr) (interface{}, error) {
	return ValueType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ValueOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ValueOverrider interface {
	// CurrentValue gets the value of this object.
	//
	// Deprecated.
	CurrentValue() externglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed. If zero, the minimum increment is undefined, which may
	// mean that it is limited only by the floating point precision of the
	// platform.
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	//
	// Deprecated.
	MaximumValue() externglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed. If zero, the minimum increment is undefined, which
	// may mean that it is limited only by the floating point precision of the
	// platform.
	//
	// Deprecated.
	MinimumIncrement() externglib.Value
	// MinimumValue gets the minimum value of this object.
	//
	// Deprecated.
	MinimumValue() externglib.Value
	// Range gets the range of this object.
	Range() *Range
	// ValueAndText gets the current value and the human readable text
	// alternative of @obj. @text is a newly created string, that must be freed
	// by the caller. Can be NULL if no descriptor is available.
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	//
	// Deprecated.
	SetCurrentValue(value externglib.Value) bool
	// SetValue sets the value of this object.
	//
	// This method is intended to provide a way to change the value of the
	// object. In any case, it is possible that the value can't be modified (ie:
	// a read-only component). If the value changes due this call, it is
	// possible that the text could change, and will trigger an
	// Value::value-changed signal emission.
	//
	// Note for implementors: the deprecated atk_value_set_current_value()
	// method returned TRUE or FALSE depending if the value was assigned or not.
	// In the practice several implementors were not able to decide it, and
	// returned TRUE in any case. For that reason it is not required anymore to
	// return if the value was properly assigned or not.
	SetValue(newValue float64)
}

// Value should be implemented for components which either display a value from
// a bounded range, or which allow the user to specify a value from a bounded
// range, or both. For instance, most sliders and range controls, as well as
// dials, should have Object representations which implement Value on the
// component's behalf. KValues may be read-only, in which case attempts to alter
// the value return would fail.
//
// <refsect1 id="current-value-text"> <title>On the subject of current value
// text</title> <para> In addition to providing the current value, implementors
// can optionally provide an end-user-consumable textual description associated
// with this value. This description should be included when the numeric value
// fails to convey the full, on-screen representation seen by users. </para>
//
// <example> <title>Password strength</title> A password strength meter whose
// value changes as the user types their new password. Red is used for values
// less than 4.0, yellow for values between 4.0 and 7.0, and green for values
// greater than 7.0. In this instance, value text should be provided by the
// implementor. Appropriate value text would be "weak", "acceptable," and
// "strong" respectively. </example>
//
// A level bar whose value changes to reflect the battery charge. The color
// remains the same regardless of the charge and there is no on-screen text
// reflecting the fullness of the battery. In this case, because the position
// within the bar is the only indication the user has of the current charge,
// value text should not be provided by the implementor.
//
// <refsect2 id="implementor-notes"> <title>Implementor Notes</title> <para>
// Implementors should bear in mind that assistive technologies will likely
// prefer the value text provided over the numeric value when presenting a
// widget's value. As a result, strings not intended for end users should not be
// exposed in the value text, and strings which are exposed should be localized.
// In the case of widgets which display value text on screen, for instance
// through a separate label in close proximity to the value-displaying widget,
// it is still expected that implementors will expose the value text using the
// above API. </para>
//
// <para> Value should NOT be implemented for widgets whose displayed value is
// not reflective of a meaningful amount. For instance, a progress pulse
// indicator whose value alternates between 0.0 and 1.0 to indicate that some
// process is still taking place should not implement Value because the current
// value does not reflect progress towards completion. </para> </refsect2>
// </refsect1>
//
// <refsect1 id="ranges"> <title>On the subject of ranges</title> <para> In
// addition to providing the minimum and maximum values, implementors can
// optionally provide details about subranges associated with the widget. These
// details should be provided by the implementor when both of the following are
// communicated visually to the end user: </para> <itemizedlist> <listitem>The
// existence of distinct ranges such as "weak", "acceptable", and "strong"
// indicated by color, bar tick marks, and/or on-screen text.</listitem>
// <listitem>Where the current value stands within a given subrange, for
// instance illustrating progression from very "weak" towards nearly
// "acceptable" through changes in shade and/or position on the bar within the
// "weak" subrange.</listitem> </itemizedlist> <para> If both of the above do
// not apply to the widget, it should be sufficient to expose the numeric value,
// along with the value text if appropriate, to make the widget accessible.
// </para>
//
// <refsect2 id="ranges-implementor-notes"> <title>Implementor Notes</title>
// <para> If providing subrange details is deemed necessary, all possible values
// of the widget are expected to fall within one of the subranges defined by the
// implementor. </para> </refsect2> </refsect1>
//
// <refsect1 id="localization"> <title>On the subject of localization of
// end-user-consumable text values</title> <para> Because value text and
// subrange descriptors are human-consumable, implementors are expected to
// provide localized strings which can be directly presented to end users via
// their assistive technology. In order to simplify this for implementors,
// implementors can use atk_value_type_get_localized_name() with the following
// already-localized constants for commonly-needed values can be used: </para>
//
// <itemizedlist> <listitem>ATK_VALUE_VERY_WEAK</listitem>
// <listitem>ATK_VALUE_WEAK</listitem> <listitem>ATK_VALUE_ACCEPTABLE</listitem>
// <listitem>ATK_VALUE_STRONG</listitem>
// <listitem>ATK_VALUE_VERY_STRONG</listitem>
// <listitem>ATK_VALUE_VERY_LOW</listitem> <listitem>ATK_VALUE_LOW</listitem>
// <listitem>ATK_VALUE_MEDIUM</listitem> <listitem>ATK_VALUE_HIGH</listitem>
// <listitem>ATK_VALUE_VERY_HIGH</listitem>
// <listitem>ATK_VALUE_VERY_BAD</listitem> <listitem>ATK_VALUE_BAD</listitem>
// <listitem>ATK_VALUE_GOOD</listitem> <listitem>ATK_VALUE_VERY_GOOD</listitem>
// <listitem>ATK_VALUE_BEST</listitem>
// <listitem>ATK_VALUE_SUBSUBOPTIMAL</listitem>
// <listitem>ATK_VALUE_SUBOPTIMAL</listitem>
// <listitem>ATK_VALUE_OPTIMAL</listitem> </itemizedlist> <para> Proposals for
// additional constants, along with their use cases, should be submitted to the
// GNOME Accessibility Team. </para> </refsect1>
//
// <refsect1 id="changes"> <title>On the subject of changes</title> <para> Note
// that if there is a textual description associated with the new numeric value,
// that description should be included regardless of whether or not it has also
// changed. </para> </refsect1>
type Value interface {
	gextras.Objector

	// CurrentValue gets the value of this object.
	//
	// Deprecated.
	CurrentValue() externglib.Value
	// Increment gets the minimum increment by which the value of this object
	// may be changed. If zero, the minimum increment is undefined, which may
	// mean that it is limited only by the floating point precision of the
	// platform.
	Increment() float64
	// MaximumValue gets the maximum value of this object.
	//
	// Deprecated.
	MaximumValue() externglib.Value
	// MinimumIncrement gets the minimum increment by which the value of this
	// object may be changed. If zero, the minimum increment is undefined, which
	// may mean that it is limited only by the floating point precision of the
	// platform.
	//
	// Deprecated.
	MinimumIncrement() externglib.Value
	// MinimumValue gets the minimum value of this object.
	//
	// Deprecated.
	MinimumValue() externglib.Value
	// Range gets the range of this object.
	Range() *Range
	// ValueAndText gets the current value and the human readable text
	// alternative of @obj. @text is a newly created string, that must be freed
	// by the caller. Can be NULL if no descriptor is available.
	ValueAndText() (float64, string)
	// SetCurrentValue sets the value of this object.
	//
	// Deprecated.
	SetCurrentValue(value externglib.Value) bool
	// SetValue sets the value of this object.
	//
	// This method is intended to provide a way to change the value of the
	// object. In any case, it is possible that the value can't be modified (ie:
	// a read-only component). If the value changes due this call, it is
	// possible that the text could change, and will trigger an
	// Value::value-changed signal emission.
	//
	// Note for implementors: the deprecated atk_value_set_current_value()
	// method returned TRUE or FALSE depending if the value was assigned or not.
	// In the practice several implementors were not able to decide it, and
	// returned TRUE in any case. For that reason it is not required anymore to
	// return if the value was properly assigned or not.
	SetValue(newValue float64)
}

// value implements the Value interface.
type value struct {
	*externglib.Object
}

var _ Value = (*value)(nil)

// WrapValue wraps a GObject to a type that implements
// interface Value. It is primarily used internally.
func WrapValue(obj *externglib.Object) Value {
	return value{obj}
}

func marshalValue(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapValue(obj), nil
}

func (o value) CurrentValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	C.atk_value_get_current_value(_arg0, &_arg1)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}

	return _value
}

func (o value) Increment() float64 {
	var _arg0 *C.AtkValue // out
	var _cret C.gdouble   // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	_cret = C.atk_value_get_increment(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (o value) MaximumValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	C.atk_value_get_maximum_value(_arg0, &_arg1)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}

	return _value
}

func (o value) MinimumIncrement() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	C.atk_value_get_minimum_increment(_arg0, &_arg1)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}

	return _value
}

func (o value) MinimumValue() externglib.Value {
	var _arg0 *C.AtkValue // out
	var _arg1 C.GValue    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	C.atk_value_get_minimum_value(_arg0, &_arg1)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}

	return _value
}

func (o value) Range() *Range {
	var _arg0 *C.AtkValue // out
	var _cret *C.AtkRange // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	_cret = C.atk_value_get_range(_arg0)

	var __range *Range // out

	__range = (*Range)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(__range, func(v *Range) {
		C.free(unsafe.Pointer(v))
	})

	return __range
}

func (o value) ValueAndText() (float64, string) {
	var _arg0 *C.AtkValue // out
	var _arg1 C.gdouble   // in
	var _arg2 *C.gchar    // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))

	C.atk_value_get_value_and_text(_arg0, &_arg1, &_arg2)

	var _value float64 // out
	var _text string   // out

	_value = float64(_arg1)
	_text = C.GoString(_arg2)
	defer C.free(unsafe.Pointer(_arg2))

	return _value, _text
}

func (o value) SetCurrentValue(value externglib.Value) bool {
	var _arg0 *C.AtkValue // out
	var _arg1 *C.GValue   // out
	var _cret C.gboolean  // in

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.atk_value_set_current_value(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o value) SetValue(newValue float64) {
	var _arg0 *C.AtkValue // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.AtkValue)(unsafe.Pointer(o.Native()))
	_arg1 = C.gdouble(newValue)

	C.atk_value_set_value(_arg0, _arg1)
}
