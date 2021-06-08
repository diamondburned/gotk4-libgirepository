// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
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
	BackgroundColor(state StateFlags) *gdk.RGBA
	// Border gets the border for a given state as a Border.
	Border(state StateFlags) *Border
	// BorderColor gets the border color for a given state.
	BorderColor(state StateFlags) *gdk.RGBA
	// Color gets the foreground color for a given state.
	Color(state StateFlags) *gdk.RGBA
	// Direction returns the widget direction used for rendering.
	Direction() TextDirection
	// Font returns the font description for a given state.
	Font(state StateFlags) *pango.FontDescription
	// JunctionSides returns the widget direction used for rendering.
	JunctionSides() JunctionSides
	// Margin gets the margin for a given state as a Border.
	Margin(state StateFlags) *Border
	// Padding gets the padding for a given state as a Border.
	Padding(state StateFlags) *Border
	// Path returns the widget path used for style matching.
	Path() *WidgetPath
	// Property gets a property value as retrieved from the style settings that
	// apply to the currently rendered element.
	Property(property string, state StateFlags) *externglib.Value
	// Screen returns the Screen to which @engine currently rendering to.
	Screen() gdk.Screen
	// State returns the state used when rendering.
	State() StateFlags
	// StyleProperty gets the value for a widget style property.
	StyleProperty(propertyName string) *externglib.Value
	// HasClass returns true if the currently rendered contents have defined the
	// given class name.
	HasClass(styleClass string) bool
	// HasRegion returns true if the currently rendered contents have the region
	// defined. If @flags_return is not nil, it is set to the flags affecting
	// the region.
	HasRegion(styleRegion string) (flags *RegionFlags, ok bool)
	// LookupColor looks up and resolves a color name in the current style’s
	// color map.
	LookupColor(colorName string) (color *gdk.RGBA, ok bool)
	// StateIsRunning returns true if there is a transition animation running
	// for the current region (see gtk_style_context_push_animatable_region()).
	//
	// If @progress is not nil, the animation progress will be returned there,
	// 0.0 means the state is closest to being false, while 1.0 means it’s
	// closest to being true. This means transition animations will run from 0
	// to 1 when @state is being set to true and from 1 to 0 when it’s being set
	// to false.
	StateIsRunning(state StateType) (progress float64, ok bool)
}

// themingEngine implements the ThemingEngine interface.
type themingEngine struct {
	gextras.Objector
}

var _ ThemingEngine = (*themingEngine)(nil)

// WrapThemingEngine wraps a GObject to the right type. It is
// primarily used internally.
func WrapThemingEngine(obj *externglib.Object) ThemingEngine {
	return ThemingEngine{
		Objector: obj,
	}
}

func marshalThemingEngine(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapThemingEngine(obj), nil
}

// BackgroundColor gets the background color for a given state.
func (e themingEngine) BackgroundColor(state StateFlags) *gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GdkRGBA)
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_background_color(arg0, arg1, arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Border gets the border for a given state as a Border.
func (e themingEngine) Border(state StateFlags) *Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GtkBorder)
	var ret2 *Border

	C.gtk_theming_engine_get_border(arg0, arg1, arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// BorderColor gets the border color for a given state.
func (e themingEngine) BorderColor(state StateFlags) *gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GdkRGBA)
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_border_color(arg0, arg1, arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Color gets the foreground color for a given state.
func (e themingEngine) Color(state StateFlags) *gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GdkRGBA)
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_color(arg0, arg1, arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Direction returns the widget direction used for rendering.
func (e themingEngine) Direction() TextDirection {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkTextDirection
	var goret TextDirection

	cret = C.gtk_theming_engine_get_direction(arg0)

	goret = TextDirection(cret)

	return goret
}

// Font returns the font description for a given state.
func (e themingEngine) Font(state StateFlags) *pango.FontDescription {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var cret *C.PangoFontDescription
	var goret *pango.FontDescription

	cret = C.gtk_theming_engine_get_font(arg0, arg1)

	goret = pango.WrapFontDescription(unsafe.Pointer(cret))

	return goret
}

// JunctionSides returns the widget direction used for rendering.
func (e themingEngine) JunctionSides() JunctionSides {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkJunctionSides
	var goret JunctionSides

	cret = C.gtk_theming_engine_get_junction_sides(arg0)

	goret = JunctionSides(cret)

	return goret
}

// Margin gets the margin for a given state as a Border.
func (e themingEngine) Margin(state StateFlags) *Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GtkBorder)
	var ret2 *Border

	C.gtk_theming_engine_get_margin(arg0, arg1, arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// Padding gets the padding for a given state as a Border.
func (e themingEngine) Padding(state StateFlags) *Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	arg2 := new(C.GtkBorder)
	var ret2 *Border

	C.gtk_theming_engine_get_padding(arg0, arg1, arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// Path returns the widget path used for style matching.
func (e themingEngine) Path() *WidgetPath {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret *C.GtkWidgetPath
	var goret *WidgetPath

	cret = C.gtk_theming_engine_get_path(arg0)

	goret = WrapWidgetPath(unsafe.Pointer(cret))

	return goret
}

// Property gets a property value as retrieved from the style settings that
// apply to the currently rendered element.
func (e themingEngine) Property(property string, state StateFlags) *externglib.Value {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar
	var arg2 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GtkStateFlags)(state)

	arg3 := new(C.GValue)
	var ret3 *externglib.Value

	C.gtk_theming_engine_get_property(arg0, arg1, arg2, arg3)

	ret3 = externglib.ValueFromNative(unsafe.Pointer(*arg3))
	runtime.SetFinalizer(ret3, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(v.GValue))
	})

	return ret3
}

// Screen returns the Screen to which @engine currently rendering to.
func (e themingEngine) Screen() gdk.Screen {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret *C.GdkScreen
	var goret gdk.Screen

	cret = C.gtk_theming_engine_get_screen(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Screen)

	return goret
}

// State returns the state used when rendering.
func (e themingEngine) State() StateFlags {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkStateFlags
	var goret StateFlags

	cret = C.gtk_theming_engine_get_state(arg0)

	goret = StateFlags(cret)

	return goret
}

// StyleProperty gets the value for a widget style property.
func (e themingEngine) StyleProperty(propertyName string) *externglib.Value {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg1))

	arg2 := new(C.GValue)
	var ret2 *externglib.Value

	C.gtk_theming_engine_get_style_property(arg0, arg1, arg2)

	ret2 = externglib.ValueFromNative(unsafe.Pointer(*arg2))

	return ret2
}

// HasClass returns true if the currently rendered contents have defined the
// given class name.
func (e themingEngine) HasClass(styleClass string) bool {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(styleClass))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_theming_engine_has_class(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// HasRegion returns true if the currently rendered contents have the region
// defined. If @flags_return is not nil, it is set to the flags affecting
// the region.
func (e themingEngine) HasRegion(styleRegion string) (flags *RegionFlags, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(styleRegion))
	defer C.free(unsafe.Pointer(arg1))

	arg2 := new(C.GtkRegionFlags)
	var ret2 *RegionFlags
	var cret C.gboolean
	var goret bool

	cret = C.gtk_theming_engine_has_region(arg0, arg1, arg2)

	ret2 = *RegionFlags(arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}

// LookupColor looks up and resolves a color name in the current style’s
// color map.
func (e themingEngine) LookupColor(colorName string) (color *gdk.RGBA, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(arg1))

	arg2 := new(C.GdkRGBA)
	var ret2 *gdk.RGBA
	var cret C.gboolean
	var goret bool

	cret = C.gtk_theming_engine_lookup_color(arg0, arg1, arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))
	if cret {
		goret = true
	}

	return ret2, goret
}

// StateIsRunning returns true if there is a transition animation running
// for the current region (see gtk_style_context_push_animatable_region()).
//
// If @progress is not nil, the animation progress will be returned there,
// 0.0 means the state is closest to being false, while 1.0 means it’s
// closest to being true. This means transition animations will run from 0
// to 1 when @state is being set to true and from 1 to 0 when it’s being set
// to false.
func (e themingEngine) StateIsRunning(state StateType) (progress float64, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateType

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateType)(state)

	arg2 := new(C.gdouble)
	var ret2 float64
	var cret C.gboolean
	var goret bool

	cret = C.gtk_theming_engine_state_is_running(arg0, arg1, arg2)

	ret2 = float64(*arg2)
	if cret {
		goret = true
	}

	return ret2, goret
}
