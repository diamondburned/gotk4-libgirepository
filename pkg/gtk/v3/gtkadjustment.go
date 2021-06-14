// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_adjustment_get_type()), F: marshalAdjustment},
	})
}

// Adjustment: the Adjustment object represents a value which has an associated
// lower and upper bound, together with step and page increments, and a page
// size. It is used within several GTK+ widgets, including SpinButton, Viewport,
// and Range (which is a base class for Scrollbar and Scale).
//
// The Adjustment object does not update the value itself. Instead it is left up
// to the owner of the Adjustment to control the value.
type Adjustment interface {
	gextras.Objector

	// Changed emits a Adjustment::changed signal from the Adjustment. This is
	// typically called by the owner of the Adjustment after it has changed any
	// of the Adjustment properties other than the value.
	Changed()
	// ClampPage updates the Adjustment:value property to ensure that the range
	// between @lower and @upper is in the current page (i.e. between
	// Adjustment:value and Adjustment:value + Adjustment:page-size). If the
	// range is larger than the page size, then only the start of it will be in
	// the current page.
	//
	// A Adjustment::value-changed signal will be emitted if the value is
	// changed.
	ClampPage(lower float64, upper float64)
	// Configure sets all properties of the adjustment at once.
	//
	// Use this function to avoid multiple emissions of the Adjustment::changed
	// signal. See gtk_adjustment_set_lower() for an alternative way of
	// compressing multiple emissions of Adjustment::changed into one.
	Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64)
	// Lower retrieves the minimum value of the adjustment.
	Lower() float64
	// MinimumIncrement gets the smaller of step increment and page increment.
	MinimumIncrement() float64
	// PageIncrement retrieves the page increment of the adjustment.
	PageIncrement() float64
	// PageSize retrieves the page size of the adjustment.
	PageSize() float64
	// StepIncrement retrieves the step increment of the adjustment.
	StepIncrement() float64
	// Upper retrieves the maximum value of the adjustment.
	Upper() float64
	// Value gets the current value of the adjustment. See
	// gtk_adjustment_set_value().
	Value() float64
	// SetLower sets the minimum value of the adjustment.
	//
	// When setting multiple adjustment properties via their individual setters,
	// multiple Adjustment::changed signals will be emitted. However, since the
	// emission of the Adjustment::changed signal is tied to the emission of the
	// #GObject::notify signals of the changed properties, it’s possible to
	// compress the Adjustment::changed signals into one by calling
	// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
	// the individual setters.
	//
	// Alternatively, using a single g_object_set() for all the properties to
	// change, or using gtk_adjustment_configure() has the same effect of
	// compressing Adjustment::changed emissions.
	SetLower(lower float64)
	// SetPageIncrement sets the page increment of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetPageIncrement(pageIncrement float64)
	// SetPageSize sets the page size of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the GtkAdjustment::changed signal when setting multiple adjustment
	// properties.
	SetPageSize(pageSize float64)
	// SetStepIncrement sets the step increment of the adjustment.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetStepIncrement(stepIncrement float64)
	// SetUpper sets the maximum value of the adjustment.
	//
	// Note that values will be restricted by `upper - page-size` if the
	// page-size property is nonzero.
	//
	// See gtk_adjustment_set_lower() about how to compress multiple emissions
	// of the Adjustment::changed signal when setting multiple adjustment
	// properties.
	SetUpper(upper float64)
	// SetValue sets the Adjustment value. The value is clamped to lie between
	// Adjustment:lower and Adjustment:upper.
	//
	// Note that for adjustments which are used in a Scrollbar, the effective
	// range of allowed values goes from Adjustment:lower to Adjustment:upper -
	// Adjustment:page-size.
	SetValue(value float64)
	// ValueChanged emits a Adjustment::value-changed signal from the
	// Adjustment. This is typically called by the owner of the Adjustment after
	// it has changed the Adjustment:value property.
	ValueChanged()
}

// adjustment implements the Adjustment class.
type adjustment struct {
	gextras.Objector
}

var _ Adjustment = (*adjustment)(nil)

// WrapAdjustment wraps a GObject to the right type. It is
// primarily used internally.
func WrapAdjustment(obj *externglib.Object) Adjustment {
	return adjustment{
		Objector: obj,
	}
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAdjustment(obj), nil
}

// Changed emits a Adjustment::changed signal from the Adjustment. This is
// typically called by the owner of the Adjustment after it has changed any
// of the Adjustment properties other than the value.
func (a adjustment) Changed() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	C.gtk_adjustment_changed(_arg0)
}

// ClampPage updates the Adjustment:value property to ensure that the range
// between @lower and @upper is in the current page (i.e. between
// Adjustment:value and Adjustment:value + Adjustment:page-size). If the
// range is larger than the page size, then only the start of it will be in
// the current page.
//
// A Adjustment::value-changed signal will be emitted if the value is
// changed.
func (a adjustment) ClampPage(lower float64, upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(lower)
	_arg2 = C.gdouble(upper)

	C.gtk_adjustment_clamp_page(_arg0, _arg1, _arg2)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the Adjustment::changed
// signal. See gtk_adjustment_set_lower() for an alternative way of
// compressing multiple emissions of Adjustment::changed into one.
func (a adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

	C.gtk_adjustment_configure(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// Lower retrieves the minimum value of the adjustment.
func (a adjustment) Lower() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_lower(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// MinimumIncrement gets the smaller of step increment and page increment.
func (a adjustment) MinimumIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_minimum_increment(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
func (a adjustment) PageIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_page_increment(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
func (a adjustment) PageSize() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_page_size(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
func (a adjustment) StepIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_step_increment(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
func (a adjustment) Upper() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_upper(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Value gets the current value of the adjustment. See
// gtk_adjustment_set_value().
func (a adjustment) Value() float64 {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_adjustment_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple Adjustment::changed signals will be emitted. However, since the
// emission of the Adjustment::changed signal is tied to the emission of the
// #GObject::notify signals of the changed properties, it’s possible to
// compress the Adjustment::changed signals into one by calling
// g_object_freeze_notify() and g_object_thaw_notify() around the calls to
// the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk_adjustment_configure() has the same effect of
// compressing Adjustment::changed emissions.
func (a adjustment) SetLower(lower float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(lower)

	C.gtk_adjustment_set_lower(_arg0, _arg1)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetPageIncrement(pageIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(pageIncrement)

	C.gtk_adjustment_set_page_increment(_arg0, _arg1)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the GtkAdjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetPageSize(pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(pageSize)

	C.gtk_adjustment_set_page_size(_arg0, _arg1)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetStepIncrement(stepIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(stepIncrement)

	C.gtk_adjustment_set_step_increment(_arg0, _arg1)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by `upper - page-size` if the
// page-size property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions
// of the Adjustment::changed signal when setting multiple adjustment
// properties.
func (a adjustment) SetUpper(upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(upper)

	C.gtk_adjustment_set_upper(_arg0, _arg1)
}

// SetValue sets the Adjustment value. The value is clamped to lie between
// Adjustment:lower and Adjustment:upper.
//
// Note that for adjustments which are used in a Scrollbar, the effective
// range of allowed values goes from Adjustment:lower to Adjustment:upper -
// Adjustment:page-size.
func (a adjustment) SetValue(value float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_adjustment_set_value(_arg0, _arg1)
}

// ValueChanged emits a Adjustment::value-changed signal from the
// Adjustment. This is typically called by the owner of the Adjustment after
// it has changed the Adjustment:value property.
func (a adjustment) ValueChanged() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(a.Native()))

	C.gtk_adjustment_value_changed(_arg0)
}
