// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_theming_engine_get_type()), F: marshalThemingEngine},
	})
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine interface {
	gextras.Objector

	// BackgroundColor gets the background color for a given state.
	//
	// Deprecated: since version 3.14.
	BackgroundColor(state StateFlags) gdk.RGBA
	// Border gets the border for a given state as a Border.
	//
	// Deprecated: since version 3.14.
	Border(state StateFlags) Border
	// BorderColor gets the border color for a given state.
	//
	// Deprecated: since version 3.14.
	BorderColor(state StateFlags) gdk.RGBA
	// Color gets the foreground color for a given state.
	//
	// Deprecated: since version 3.14.
	Color(state StateFlags) gdk.RGBA
	// Direction returns the widget direction used for rendering.
	//
	// Deprecated: since version 3.8.
	Direction() TextDirection
	// Font returns the font description for a given state.
	//
	// Deprecated: since version 3.8.
	Font(state StateFlags) *pango.FontDescription
	// JunctionSides returns the widget direction used for rendering.
	//
	// Deprecated: since version 3.14.
	JunctionSides() JunctionSides
	// Margin gets the margin for a given state as a Border.
	//
	// Deprecated: since version 3.14.
	Margin(state StateFlags) Border
	// Padding gets the padding for a given state as a Border.
	//
	// Deprecated: since version 3.14.
	Padding(state StateFlags) Border
	// Path returns the widget path used for style matching.
	//
	// Deprecated: since version 3.14.
	Path() *WidgetPath
	// Property gets a property value as retrieved from the style settings that
	// apply to the currently rendered element.
	//
	// Deprecated: since version 3.14.
	Property(property string, state StateFlags) externglib.Value
	// Screen returns the Screen to which @engine currently rendering to.
	//
	// Deprecated: since version 3.14.
	Screen() gdk.Screen
	// State returns the state used when rendering.
	//
	// Deprecated: since version 3.14.
	State() StateFlags
	// StyleProperty gets the value for a widget style property.
	//
	// Deprecated: since version 3.14.
	StyleProperty(propertyName string) externglib.Value
	// HasClassThemingEngine returns true if the currently rendered contents
	// have defined the given class name.
	//
	// Deprecated: since version 3.14.
	HasClassThemingEngine(styleClass string) bool
	// HasRegionThemingEngine returns true if the currently rendered contents
	// have the region defined. If @flags_return is not nil, it is set to the
	// flags affecting the region.
	//
	// Deprecated: since version 3.14.
	HasRegionThemingEngine(styleRegion string) (RegionFlags, bool)
	// LookupColorThemingEngine looks up and resolves a color name in the
	// current style’s color map.
	//
	// Deprecated: since version 3.14.
	LookupColorThemingEngine(colorName string) (gdk.RGBA, bool)
	// StateIsRunningThemingEngine returns true if there is a transition
	// animation running for the current region (see
	// gtk_style_context_push_animatable_region()).
	//
	// If @progress is not nil, the animation progress will be returned there,
	// 0.0 means the state is closest to being false, while 1.0 means it’s
	// closest to being true. This means transition animations will run from 0
	// to 1 when @state is being set to true and from 1 to 0 when it’s being set
	// to false.
	//
	// Deprecated: since version 3.6.
	StateIsRunningThemingEngine(state StateType) (float64, bool)
}

// themingEngine implements the ThemingEngine class.
type themingEngine struct {
	gextras.Objector
}

// WrapThemingEngine wraps a GObject to the right type. It is
// primarily used internally.
func WrapThemingEngine(obj *externglib.Object) ThemingEngine {
	return themingEngine{
		Objector: obj,
	}
}

func marshalThemingEngine(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapThemingEngine(obj), nil
}

func (e themingEngine) BackgroundColor(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_background_color(_arg0, _arg1, &_arg2)

	var _color gdk.RGBA // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (e themingEngine) Border(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border(_arg0, _arg1, &_arg2)

	var _border Border // out

	{
		var refTmpIn *C.GtkBorder
		var refTmpOut *Border

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Border)(unsafe.Pointer(refTmpIn))

		_border = *refTmpOut
	}

	return _border
}

func (e themingEngine) BorderColor(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border_color(_arg0, _arg1, &_arg2)

	var _color gdk.RGBA // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (e themingEngine) Color(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_color(_arg0, _arg1, &_arg2)

	var _color gdk.RGBA // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (e themingEngine) Direction() TextDirection {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkTextDirection  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_theming_engine_get_direction(_arg0)

	var _textDirection TextDirection // out

	_textDirection = TextDirection(_cret)

	return _textDirection
}

func (e themingEngine) Font(state StateFlags) *pango.FontDescription {
	var _arg0 *C.GtkThemingEngine     // out
	var _arg1 C.GtkStateFlags         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	_cret = C.gtk_theming_engine_get_font(_arg0, _arg1)

	var _fontDescription *pango.FontDescription // out

	_fontDescription = (*pango.FontDescription)(unsafe.Pointer(_cret))

	return _fontDescription
}

func (e themingEngine) JunctionSides() JunctionSides {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkJunctionSides  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_theming_engine_get_junction_sides(_arg0)

	var _junctionSides JunctionSides // out

	_junctionSides = JunctionSides(_cret)

	return _junctionSides
}

func (e themingEngine) Margin(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_margin(_arg0, _arg1, &_arg2)

	var _margin Border // out

	{
		var refTmpIn *C.GtkBorder
		var refTmpOut *Border

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Border)(unsafe.Pointer(refTmpIn))

		_margin = *refTmpOut
	}

	return _margin
}

func (e themingEngine) Padding(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_padding(_arg0, _arg1, &_arg2)

	var _padding Border // out

	{
		var refTmpIn *C.GtkBorder
		var refTmpOut *Border

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Border)(unsafe.Pointer(refTmpIn))

		_padding = *refTmpOut
	}

	return _padding
}

func (e themingEngine) Path() *WidgetPath {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GtkWidgetPath    // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_theming_engine_get_path(_arg0)

	var _widgetPath *WidgetPath // out

	_widgetPath = (*WidgetPath)(unsafe.Pointer(_cret))
	C.gtk_widget_path_ref(_cret)
	runtime.SetFinalizer(_widgetPath, func(v *WidgetPath) {
		C.gtk_widget_path_unref((*C.GtkWidgetPath)(unsafe.Pointer(v)))
	})

	return _widgetPath
}

func (e themingEngine) Property(property string, state StateFlags) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkStateFlags     // out
	var _arg3 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_property(_arg0, _arg1, _arg2, &_arg3)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))
		runtime.SetFinalizer(refTmpOut, func(v *externglib.Value) {
			C.g_value_unset((*C.GValue)(v.GValue))
		})

		_value = *refTmpOut
	}

	return _value
}

func (e themingEngine) Screen() gdk.Screen {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GdkScreen        // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_theming_engine_get_screen(_arg0)

	var _screen gdk.Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Screen)

	return _screen
}

func (e themingEngine) State() StateFlags {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkStateFlags     // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_theming_engine_get_state(_arg0)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

func (e themingEngine) StyleProperty(propertyName string) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_theming_engine_get_style_property(_arg0, _arg1, &_arg2)

	var _value externglib.Value // out

	{
		var refTmpIn *C.GValue
		var refTmpOut *externglib.Value

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = externglib.ValueFromNative(unsafe.Pointer(refTmpIn))

		_value = *refTmpOut
	}

	return _value
}

func (e themingEngine) HasClassThemingEngine(styleClass string) bool {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(styleClass))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_has_class(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e themingEngine) HasRegionThemingEngine(styleRegion string) (RegionFlags, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkRegionFlags    // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(styleRegion))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_has_region(_arg0, _arg1, &_arg2)

	var _flags RegionFlags // out
	var _ok bool           // out

	_flags = RegionFlags(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _flags, _ok
}

func (e themingEngine) LookupColorThemingEngine(colorName string) (gdk.RGBA, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GdkRGBA           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_lookup_color(_arg0, _arg1, &_arg2)

	var _color gdk.RGBA // out
	var _ok bool        // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

func (e themingEngine) StateIsRunningThemingEngine(state StateType) (float64, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateType      // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkStateType(state)

	_cret = C.gtk_theming_engine_state_is_running(_arg0, _arg1, &_arg2)

	var _progress float64 // out
	var _ok bool          // out

	_progress = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _progress, _ok
}
