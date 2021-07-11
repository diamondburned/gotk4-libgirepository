// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_expander_style_get_type()), F: marshalExpanderStyle},
		{T: externglib.Type(C.gtk_style_get_type()), F: marshalStyler},
	})
}

// ExpanderStyle: used to specify the style of the expanders drawn by a
// TreeView.
type ExpanderStyle int

const (
	// Collapsed: style used for a collapsed subtree.
	ExpanderStyleCollapsed ExpanderStyle = iota
	// SemiCollapsed: intermediate style used during animation.
	ExpanderStyleSemiCollapsed
	// SemiExpanded: intermediate style used during animation.
	ExpanderStyleSemiExpanded
	// Expanded: style used for an expanded subtree.
	ExpanderStyleExpanded
)

func marshalExpanderStyle(p uintptr) (interface{}, error) {
	return ExpanderStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StyleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StyleOverrider interface {
	//
	Copy(src Styler)
	//
	InitFromRC(rcStyle RCStyler)
	//
	Realize()
	//
	Unrealize()
}

// Styler describes Style's methods.
type Styler interface {
	// Copy creates a copy of the passed in Style object.
	Copy() *Style
	// Detach detaches a style from a window.
	Detach()
	// StyleProperty queries the value of a style property corresponding to a
	// widget class is in the given style.
	StyleProperty(widgetType externglib.Type, propertyName string) externglib.Value
	// HasContext returns whether @style has an associated StyleContext.
	HasContext() bool
	// LookupColor looks up @color_name in the style’s logical color mappings,
	// filling in @color and returning true if found, otherwise returning false.
	LookupColor(colorName string) (gdk.Color, bool)
	// LookupIconSet looks up @stock_id in the icon factories associated with
	// @style and the default icon factory, returning an icon set if found,
	// otherwise nil.
	LookupIconSet(stockId string) *IconSet
}

// Style object encapsulates the information that provides the look and feel for
// a widget.
//
// > In GTK+ 3.0, GtkStyle has been deprecated and replaced by > StyleContext.
//
// Each Widget has an associated Style object that is used when rendering that
// widget. Also, a Style holds information for the five possible widget states
// though not every widget supports all five states; see StateType.
//
// Usually the Style for a widget is the same as the default style that is set
// by GTK+ and modified the theme engine.
//
// Usually applications should not need to use or modify the Style of their
// widgets.
type Style struct {
	*externglib.Object
}

var (
	_ Styler          = (*Style)(nil)
	_ gextras.Nativer = (*Style)(nil)
)

func wrapStyle(obj *externglib.Object) Styler {
	return &Style{
		Object: obj,
	}
}

func marshalStyler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStyle(obj), nil
}

// NewStyle creates a new Style.
//
// Deprecated: Use StyleContext.
func NewStyle() *Style {
	var _cret *C.GtkStyle // in

	_cret = C.gtk_style_new()

	var _style *Style // out

	_style = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Style)

	return _style
}

// Copy creates a copy of the passed in Style object.
//
// Deprecated: Use StyleContext instead.
func (style *Style) Copy() *Style {
	var _arg0 *C.GtkStyle // out
	var _cret *C.GtkStyle // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))

	_cret = C.gtk_style_copy(_arg0)

	var _ret *Style // out

	_ret = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Style)

	return _ret
}

// Detach detaches a style from a window. If the style is not attached to any
// windows anymore, it is unrealized. See gtk_style_attach().
//
// Deprecated: Use StyleContext instead.
func (style *Style) Detach() {
	var _arg0 *C.GtkStyle // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))

	C.gtk_style_detach(_arg0)
}

// StyleProperty queries the value of a style property corresponding to a widget
// class is in the given style.
func (style *Style) StyleProperty(widgetType externglib.Type, propertyName string) externglib.Value {
	var _arg0 *C.GtkStyle // out
	var _arg1 C.GType     // out
	var _arg2 *C.gchar    // out
	var _arg3 C.GValue    // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = C.GType(widgetType)
	_arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_style_get_style_property(_arg0, _arg1, _arg2, &_arg3)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))

	return _value
}

// HasContext returns whether @style has an associated StyleContext.
func (style *Style) HasContext() bool {
	var _arg0 *C.GtkStyle // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))

	_cret = C.gtk_style_has_context(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupColor looks up @color_name in the style’s logical color mappings,
// filling in @color and returning true if found, otherwise returning false. Do
// not cache the found mapping, because it depends on the Style and might change
// when a theme switch occurs.
//
// Deprecated: Use gtk_style_context_lookup_color() instead.
func (style *Style) LookupColor(colorName string) (gdk.Color, bool) {
	var _arg0 *C.GtkStyle // out
	var _arg1 *C.gchar    // out
	var _color gdk.Color
	var _cret C.gboolean // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_color(_arg0, _arg1, (*C.GdkColor)(unsafe.Pointer(&_color)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// LookupIconSet looks up @stock_id in the icon factories associated with @style
// and the default icon factory, returning an icon set if found, otherwise nil.
//
// Deprecated: Use gtk_style_context_lookup_icon_set() instead.
func (style *Style) LookupIconSet(stockId string) *IconSet {
	var _arg0 *C.GtkStyle   // out
	var _arg1 *C.gchar      // out
	var _cret *C.GtkIconSet // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_icon_set(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(unsafe.Pointer(_cret))
	C.gtk_icon_set_ref(_cret)
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(unsafe.Pointer(v)))
	})

	return _iconSet
}
