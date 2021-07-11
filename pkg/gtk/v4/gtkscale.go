// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scale_get_type()), F: marshalScaler},
	})
}

//
type ScaleFormatValueFunc func(scale *Scale, value float64, userData cgo.Handle) (utf8 string)

//export gotk4_ScaleFormatValueFunc
func gotk4_ScaleFormatValueFunc(arg0 *C.GtkScale, arg1 C.double, arg2 C.gpointer) (cret *C.char) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var scale *Scale        // out
	var value float64       // out
	var userData cgo.Handle // out

	scale = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*Scale)
	value = float64(arg1)
	userData = (cgo.Handle)(arg2)

	fn := v.(ScaleFormatValueFunc)
	utf8 := fn(scale, value, userData)

	cret = (*C.char)(C.CString(utf8))

	return cret
}

// ScaleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleOverrider interface {
	// LayoutOffsets obtains the coordinates where the scale will draw the
	// `PangoLayout` representing the text in the scale.
	//
	// Remember when using the `PangoLayout` function you need to convert to and
	// from pixels using `PANGO_PIXELS()` or `PANGO_SCALE`.
	//
	// If the [property@GtkScale:draw-value] property is false, the return
	// values are undefined.
	LayoutOffsets() (x int, y int)
}

// Scaler describes Scale's methods.
type Scaler interface {
	// ClearMarks removes any marks that have been added.
	ClearMarks()
	// Digits gets the number of decimal places that are displayed in the value.
	Digits() int
	// DrawValue returns whether the current value is displayed as a string next
	// to the slider.
	DrawValue() bool
	// HasOrigin returns whether the scale has an origin.
	HasOrigin() bool
	// Layout gets the `PangoLayout` used to display the scale.
	Layout() *pango.Layout
	// LayoutOffsets obtains the coordinates where the scale will draw the
	// `PangoLayout` representing the text in the scale.
	LayoutOffsets() (x int, y int)
	// ValuePos gets the position in which the current value is displayed.
	ValuePos() PositionType
	// SetDigits sets the number of decimal places that are displayed in the
	// value.
	SetDigits(digits int)
	// SetDrawValue specifies whether the current value is displayed as a string
	// next to the slider.
	SetDrawValue(drawValue bool)
	// SetHasOrigin sets whether the scale has an origin.
	SetHasOrigin(hasOrigin bool)
}

// Scale: `GtkScale` is a slider control used to select a numeric value.
//
// !An example GtkScale (scales.png)
//
// To use it, you’ll probably want to investigate the methods on its base class,
// [class@GtkRange], in addition to the methods for `GtkScale` itself. To set
// the value of a scale, you would normally use [method@Gtk.Range.set_value]. To
// detect changes to the value, you would normally use the
// [signal@Gtk.Range::value-changed] signal.
//
// Note that using the same upper and lower bounds for the `GtkScale` (through
// the `GtkRange` methods) will hide the slider itself. This is useful for
// applications that want to show an undeterminate value on the scale, without
// changing the layout of the application (such as movie or music players).
//
//
// GtkScale as GtkBuildable
//
// `GtkScale` supports a custom <marks> element, which can contain multiple
// <mark\> elements. The “value” and “position” attributes have the same meaning
// as [method@Gtk.Scale.add_mark] parameters of the same name. If the element is
// not empty, its content is taken as the markup to show at the mark. It can be
// translated with the usual ”translatable” and “context” attributes.
//
//
// CSS nodes
//
// “` scale[.fine-tune][.marks-before][.marks-after] ├──
// [value][.top][.right][.bottom][.left] ├── marks.top │ ├── mark │ ┊ ├──
// [label] │ ┊ ╰── indicator ┊ ┊ │ ╰── mark ├── marks.bottom │ ├── mark │ ┊ ├──
// indicator │ ┊ ╰── [label] ┊ ┊ │ ╰── mark ╰── trough ├── [fill] ├──
// [highlight] ╰── slider “`
//
// `GtkScale` has a main CSS node with name scale and a subnode for its
// contents, with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see [method@Gtk.Scale.set_has_origin]), there is
// a subnode with name highlight below the trough node that is used for
// rendering the highlighted part of the trough.
//
// If the scale is showing a fill level (see
// [method@Gtk.Range.set_show_fill_level]), there is a subnode with name fill
// below the trough node that is used for rendering the filled in part of the
// trough.
//
// If marks are present, there is a marks subnode before or after the trough
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see [property@Gtk.Scale:draw-value]),
// there is subnode with name value. This node will get the .top or .bottom
// style classes similar to the marks node.
//
//
// Accessibility
//
// `GtkScale` uses the GTK_ACCESSIBLE_ROLE_SLIDER role.
type Scale struct {
	Range
}

var (
	_ Scaler          = (*Scale)(nil)
	_ gextras.Nativer = (*Scale)(nil)
)

func wrapScale(obj *externglib.Object) Scaler {
	return &Scale{
		Range: Range{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalScaler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScale(obj), nil
}

// ClearMarks removes any marks that have been added.
func (scale *Scale) ClearMarks() {
	var _arg0 *C.GtkScale // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	C.gtk_scale_clear_marks(_arg0)
}

// Digits gets the number of decimal places that are displayed in the value.
func (scale *Scale) Digits() int {
	var _arg0 *C.GtkScale // out
	var _cret C.int       // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_digits(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DrawValue returns whether the current value is displayed as a string next to
// the slider.
func (scale *Scale) DrawValue() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_draw_value(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasOrigin returns whether the scale has an origin.
func (scale *Scale) HasOrigin() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_has_origin(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout gets the `PangoLayout` used to display the scale.
//
// The returned object is owned by the scale so does not need to be freed by the
// caller.
func (scale *Scale) Layout() *pango.Layout {
	var _arg0 *C.GtkScale    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_layout(_arg0)

	var _layout *pango.Layout // out

	_layout = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*pango.Layout)

	return _layout
}

// LayoutOffsets obtains the coordinates where the scale will draw the
// `PangoLayout` representing the text in the scale.
//
// Remember when using the `PangoLayout` function you need to convert to and
// from pixels using `PANGO_PIXELS()` or `PANGO_SCALE`.
//
// If the [property@GtkScale:draw-value] property is false, the return values
// are undefined.
func (scale *Scale) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	C.gtk_scale_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// ValuePos gets the position in which the current value is displayed.
func (scale *Scale) ValuePos() PositionType {
	var _arg0 *C.GtkScale       // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_value_pos(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// SetDigits sets the number of decimal places that are displayed in the value.
//
// Also causes the value of the adjustment to be rounded to this number of
// digits, so the retrieved value matches the displayed one, if
// [property@GtkScale:draw-value] is true when the value changes. If you want to
// enforce rounding the value when [property@GtkScale:draw-value] is false, you
// can set [property@GtkRange:round-digits] instead.
//
// Note that rounding to a small number of digits can interfere with the smooth
// autoscrolling that is built into `GtkScale`. As an alternative, you can use
// [method@Gtk.Scale.set_format_value_func] to format the displayed value
// yourself.
func (scale *Scale) SetDigits(digits int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	_arg1 = C.int(digits)

	C.gtk_scale_set_digits(_arg0, _arg1)
}

// SetDrawValue specifies whether the current value is displayed as a string
// next to the slider.
func (scale *Scale) SetDrawValue(drawValue bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	if drawValue {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_draw_value(_arg0, _arg1)
}

// SetHasOrigin sets whether the scale has an origin.
//
// If [property@GtkScale:has-origin] is set to true (the default), the scale
// will highlight the part of the trough between the origin (bottom or left
// side) and the current value.
func (scale *Scale) SetHasOrigin(hasOrigin bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	if hasOrigin {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_has_origin(_arg0, _arg1)
}
