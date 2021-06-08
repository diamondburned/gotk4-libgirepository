// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRange},
	})
}

// Range is the common base class for widgets which visualize an adjustment, e.g
// Scale or Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, Range
// provides properties and methods for influencing the sensitivity of the
// “steppers”. It also provides properties and methods for setting a “fill
// level” on range widgets. See gtk_range_set_fill_level().
type Range interface {
	Widget
	Buildable
	Orientable

	// Adjustment: get the Adjustment which is the “model” object for Range. See
	// gtk_range_set_adjustment() for details. The return value does not have a
	// reference added, so should not be unreferenced.
	Adjustment() Adjustment
	// FillLevel gets the current position of the fill level indicator.
	FillLevel() float64
	// Flippable gets the value set by gtk_range_set_flippable().
	Flippable() bool
	// Inverted gets the value set by gtk_range_set_inverted().
	Inverted() bool
	// LowerStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'lower' end of the GtkRange’s adjustment.
	LowerStepperSensitivity() SensitivityType
	// MinSliderSize: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_min_slider_size().
	MinSliderSize() int
	// RangeRect: this function returns the area that contains the range’s
	// trough and its steppers, in widget->window coordinates.
	//
	// This function is useful mainly for Range subclasses.
	RangeRect() *gdk.Rectangle
	// RestrictToFillLevel gets whether the range is restricted to the fill
	// level.
	RestrictToFillLevel() bool
	// RoundDigits gets the number of digits to round the value to when it
	// changes. See Range::change-value.
	RoundDigits() int
	// ShowFillLevel gets whether the range displays the fill level graphically.
	ShowFillLevel() bool
	// SliderRange: this function returns sliders range along the long
	// dimension, in widget->window coordinates.
	//
	// This function is useful mainly for Range subclasses.
	SliderRange() (sliderStart int, sliderEnd int)
	// SliderSizeFixed: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_slider_size_fixed().
	SliderSizeFixed() bool
	// UpperStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'upper' end of the GtkRange’s adjustment.
	UpperStepperSensitivity() SensitivityType
	// Value gets the current value of the range.
	Value() float64
	// SetAdjustment sets the adjustment to be used as the “model” object for
	// this range widget. The adjustment indicates the current range value, the
	// minimum and maximum range values, the step/page increments used for
	// keybindings and scrolling, and the page size. The page size is normally 0
	// for Scale and nonzero for Scrollbar, and indicates the size of the
	// visible area of the widget being scrolled. The page size affects the size
	// of the scrollbar slider.
	SetAdjustment(adjustment Adjustment)
	// SetFillLevel: set the new position of the fill level indicator.
	//
	// The “fill level” is probably best described by its most prominent use
	// case, which is an indicator for the amount of pre-buffering in a
	// streaming media player. In that use case, the value of the range would
	// indicate the current play position, and the fill level would be the
	// position up to which the file/stream has been downloaded.
	//
	// This amount of prebuffering can be displayed on the range’s trough and is
	// themeable separately from the trough. To enable fill level display, use
	// gtk_range_set_show_fill_level(). The range defaults to not showing the
	// fill level.
	//
	// Additionally, it’s possible to restrict the range’s slider position to
	// values which are smaller than the fill level. This is controller by
	// gtk_range_set_restrict_to_fill_level() and is by default enabled.
	SetFillLevel(fillLevel float64)
	// SetFlippable: if a range is flippable, it will switch its direction if it
	// is horizontal and its direction is GTK_TEXT_DIR_RTL.
	//
	// See gtk_widget_get_direction().
	SetFlippable(flippable bool)
	// SetIncrements sets the step and page sizes for the range. The step size
	// is used when the user clicks the Scrollbar arrows or moves Scale via
	// arrow keys. The page size is used for example when moving via Page Up or
	// Page Down keys.
	SetIncrements(step float64, page float64)
	// SetInverted ranges normally move from lower to higher values as the
	// slider moves from top to bottom or left to right. Inverted ranges have
	// higher values at the top or on the right rather than on the bottom or
	// left.
	SetInverted(setting bool)
	// SetLowerStepperSensitivity sets the sensitivity policy for the stepper
	// that points to the 'lower' end of the GtkRange’s adjustment.
	SetLowerStepperSensitivity(sensitivity SensitivityType)
	// SetMinSliderSize sets the minimum size of the range’s slider.
	//
	// This function is useful mainly for Range subclasses.
	SetMinSliderSize(minSize int)
	// SetRange sets the allowable values in the Range, and clamps the range
	// value to be between @min and @max. (If the range has a non-zero page
	// size, it is clamped between @min and @max - page-size.)
	SetRange(min float64, max float64)
	// SetRestrictToFillLevel sets whether the slider is restricted to the fill
	// level. See gtk_range_set_fill_level() for a general description of the
	// fill level concept.
	SetRestrictToFillLevel(restrictToFillLevel bool)
	// SetRoundDigits sets the number of digits to round the value to when it
	// changes. See Range::change-value.
	SetRoundDigits(roundDigits int)
	// SetShowFillLevel sets whether a graphical fill level is show on the
	// trough. See gtk_range_set_fill_level() for a general description of the
	// fill level concept.
	SetShowFillLevel(showFillLevel bool)
	// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
	// size that depends on its adjustment’s page size.
	//
	// This function is useful mainly for Range subclasses.
	SetSliderSizeFixed(sizeFixed bool)
	// SetUpperStepperSensitivity sets the sensitivity policy for the stepper
	// that points to the 'upper' end of the GtkRange’s adjustment.
	SetUpperStepperSensitivity(sensitivity SensitivityType)
	// SetValue sets the current value of the range; if the value is outside the
	// minimum or maximum range values, it will be clamped to fit inside them.
	// The range emits the Range::value-changed signal if the value changes.
	SetValue(value float64)
}

// _range implements the Range interface.
type _range struct {
	Widget
	Buildable
	Orientable
}

var _ Range = (*_range)(nil)

// WrapRange wraps a GObject to the right type. It is
// primarily used internally.
func WrapRange(obj *externglib.Object) Range {
	return Range{
		Widget:     WrapWidget(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalRange(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRange(obj), nil
}

// Adjustment: get the Adjustment which is the “model” object for Range. See
// gtk_range_set_adjustment() for details. The return value does not have a
// reference added, so should not be unreferenced.
func (r _range) Adjustment() Adjustment {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret *C.GtkAdjustment
	var goret Adjustment

	cret = C.gtk_range_get_adjustment(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return goret
}

// FillLevel gets the current position of the fill level indicator.
func (r _range) FillLevel() float64 {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gdouble
	var goret float64

	cret = C.gtk_range_get_fill_level(arg0)

	goret = float64(cret)

	return goret
}

// Flippable gets the value set by gtk_range_set_flippable().
func (r _range) Flippable() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_range_get_flippable(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Inverted gets the value set by gtk_range_set_inverted().
func (r _range) Inverted() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_range_get_inverted(arg0)

	if cret {
		goret = true
	}

	return goret
}

// LowerStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'lower' end of the GtkRange’s adjustment.
func (r _range) LowerStepperSensitivity() SensitivityType {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.GtkSensitivityType
	var goret SensitivityType

	cret = C.gtk_range_get_lower_stepper_sensitivity(arg0)

	goret = SensitivityType(cret)

	return goret
}

// MinSliderSize: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_min_slider_size().
func (r _range) MinSliderSize() int {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_range_get_min_slider_size(arg0)

	goret = int(cret)

	return goret
}

// RangeRect: this function returns the area that contains the range’s
// trough and its steppers, in widget->window coordinates.
//
// This function is useful mainly for Range subclasses.
func (r _range) RangeRect() *gdk.Rectangle {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	arg1 := new(C.GdkRectangle)
	var ret1 *gdk.Rectangle

	C.gtk_range_get_range_rect(arg0, arg1)

	ret1 = gdk.WrapRectangle(unsafe.Pointer(arg1))

	return ret1
}

// RestrictToFillLevel gets whether the range is restricted to the fill
// level.
func (r _range) RestrictToFillLevel() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_range_get_restrict_to_fill_level(arg0)

	if cret {
		goret = true
	}

	return goret
}

// RoundDigits gets the number of digits to round the value to when it
// changes. See Range::change-value.
func (r _range) RoundDigits() int {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gint
	var goret int

	cret = C.gtk_range_get_round_digits(arg0)

	goret = int(cret)

	return goret
}

// ShowFillLevel gets whether the range displays the fill level graphically.
func (r _range) ShowFillLevel() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_range_get_show_fill_level(arg0)

	if cret {
		goret = true
	}

	return goret
}

// SliderRange: this function returns sliders range along the long
// dimension, in widget->window coordinates.
//
// This function is useful mainly for Range subclasses.
func (r _range) SliderRange() (sliderStart int, sliderEnd int) {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	arg1 := new(C.gint)
	var ret1 int
	arg2 := new(C.gint)
	var ret2 int

	C.gtk_range_get_slider_range(arg0, arg1, arg2)

	ret1 = int(*arg1)
	ret2 = int(*arg2)

	return ret1, ret2
}

// SliderSizeFixed: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_slider_size_fixed().
func (r _range) SliderSizeFixed() bool {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_range_get_slider_size_fixed(arg0)

	if cret {
		goret = true
	}

	return goret
}

// UpperStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'upper' end of the GtkRange’s adjustment.
func (r _range) UpperStepperSensitivity() SensitivityType {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.GtkSensitivityType
	var goret SensitivityType

	cret = C.gtk_range_get_upper_stepper_sensitivity(arg0)

	goret = SensitivityType(cret)

	return goret
}

// Value gets the current value of the range.
func (r _range) Value() float64 {
	var arg0 *C.GtkRange

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	var cret C.gdouble
	var goret float64

	cret = C.gtk_range_get_value(arg0)

	goret = float64(cret)

	return goret
}

// SetAdjustment sets the adjustment to be used as the “model” object for
// this range widget. The adjustment indicates the current range value, the
// minimum and maximum range values, the step/page increments used for
// keybindings and scrolling, and the page size. The page size is normally 0
// for Scale and nonzero for Scrollbar, and indicates the size of the
// visible area of the widget being scrolled. The page size affects the size
// of the scrollbar slider.
func (r _range) SetAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkRange
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_range_set_adjustment(arg0, arg1)
}

// SetFillLevel: set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent use
// case, which is an indicator for the amount of pre-buffering in a
// streaming media player. In that use case, the value of the range would
// indicate the current play position, and the fill level would be the
// position up to which the file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough and is
// themeable separately from the trough. To enable fill level display, use
// gtk_range_set_show_fill_level(). The range defaults to not showing the
// fill level.
//
// Additionally, it’s possible to restrict the range’s slider position to
// values which are smaller than the fill level. This is controller by
// gtk_range_set_restrict_to_fill_level() and is by default enabled.
func (r _range) SetFillLevel(fillLevel float64) {
	var arg0 *C.GtkRange
	var arg1 C.gdouble

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gdouble(fillLevel)

	C.gtk_range_set_fill_level(arg0, arg1)
}

// SetFlippable: if a range is flippable, it will switch its direction if it
// is horizontal and its direction is GTK_TEXT_DIR_RTL.
//
// See gtk_widget_get_direction().
func (r _range) SetFlippable(flippable bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if flippable {
		arg1 = C.gboolean(1)
	}

	C.gtk_range_set_flippable(arg0, arg1)
}

// SetIncrements sets the step and page sizes for the range. The step size
// is used when the user clicks the Scrollbar arrows or moves Scale via
// arrow keys. The page size is used for example when moving via Page Up or
// Page Down keys.
func (r _range) SetIncrements(step float64, page float64) {
	var arg0 *C.GtkRange
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gdouble(step)
	arg2 = C.gdouble(page)

	C.gtk_range_set_increments(arg0, arg1, arg2)
}

// SetInverted ranges normally move from lower to higher values as the
// slider moves from top to bottom or left to right. Inverted ranges have
// higher values at the top or on the right rather than on the bottom or
// left.
func (r _range) SetInverted(setting bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_range_set_inverted(arg0, arg1)
}

// SetLowerStepperSensitivity sets the sensitivity policy for the stepper
// that points to the 'lower' end of the GtkRange’s adjustment.
func (r _range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	var arg0 *C.GtkRange
	var arg1 C.GtkSensitivityType

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = (C.GtkSensitivityType)(sensitivity)

	C.gtk_range_set_lower_stepper_sensitivity(arg0, arg1)
}

// SetMinSliderSize sets the minimum size of the range’s slider.
//
// This function is useful mainly for Range subclasses.
func (r _range) SetMinSliderSize(minSize int) {
	var arg0 *C.GtkRange
	var arg1 C.gint

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gint(minSize)

	C.gtk_range_set_min_slider_size(arg0, arg1)
}

// SetRange sets the allowable values in the Range, and clamps the range
// value to be between @min and @max. (If the range has a non-zero page
// size, it is clamped between @min and @max - page-size.)
func (r _range) SetRange(min float64, max float64) {
	var arg0 *C.GtkRange
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gdouble(min)
	arg2 = C.gdouble(max)

	C.gtk_range_set_range(arg0, arg1, arg2)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level. See gtk_range_set_fill_level() for a general description of the
// fill level concept.
func (r _range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if restrictToFillLevel {
		arg1 = C.gboolean(1)
	}

	C.gtk_range_set_restrict_to_fill_level(arg0, arg1)
}

// SetRoundDigits sets the number of digits to round the value to when it
// changes. See Range::change-value.
func (r _range) SetRoundDigits(roundDigits int) {
	var arg0 *C.GtkRange
	var arg1 C.gint

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gint(roundDigits)

	C.gtk_range_set_round_digits(arg0, arg1)
}

// SetShowFillLevel sets whether a graphical fill level is show on the
// trough. See gtk_range_set_fill_level() for a general description of the
// fill level concept.
func (r _range) SetShowFillLevel(showFillLevel bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if showFillLevel {
		arg1 = C.gboolean(1)
	}

	C.gtk_range_set_show_fill_level(arg0, arg1)
}

// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
// size that depends on its adjustment’s page size.
//
// This function is useful mainly for Range subclasses.
func (r _range) SetSliderSizeFixed(sizeFixed bool) {
	var arg0 *C.GtkRange
	var arg1 C.gboolean

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if sizeFixed {
		arg1 = C.gboolean(1)
	}

	C.gtk_range_set_slider_size_fixed(arg0, arg1)
}

// SetUpperStepperSensitivity sets the sensitivity policy for the stepper
// that points to the 'upper' end of the GtkRange’s adjustment.
func (r _range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	var arg0 *C.GtkRange
	var arg1 C.GtkSensitivityType

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = (C.GtkSensitivityType)(sensitivity)

	C.gtk_range_set_upper_stepper_sensitivity(arg0, arg1)
}

// SetValue sets the current value of the range; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them.
// The range emits the Range::value-changed signal if the value changes.
func (r _range) SetValue(value float64) {
	var arg0 *C.GtkRange
	var arg1 C.gdouble

	arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	arg1 = C.gdouble(value)

	C.gtk_range_set_value(arg0, arg1)
}
