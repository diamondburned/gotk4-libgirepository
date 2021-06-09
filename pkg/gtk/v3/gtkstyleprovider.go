// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
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
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProvider},
	})
}

// StyleProviderOverrider contains methods that are overridable. This
// interface is a subset of the interface StyleProvider.
type StyleProviderOverrider interface {
	// IconFactory returns the IconFactory defined to be in use for @path, or
	// nil if none is defined.
	IconFactory(path *WidgetPath) IconFactory
	// Style returns the style settings affecting a widget defined by @path, or
	// nil if @provider doesn’t contemplate styling @path.
	Style(path *WidgetPath) StyleProperties
	// StyleProperty looks up a widget style property as defined by @provider
	// for the widget represented by @path.
	StyleProperty(path *WidgetPath, state StateFlags, pspec gobject.ParamSpec) (*externglib.Value, bool)
}

// StyleProvider: gtkStyleProvider is an interface used to provide style
// information to a StyleContext. See gtk_style_context_add_provider() and
// gtk_style_context_add_provider_for_screen().
type StyleProvider interface {
	gextras.Objector
	StyleProviderOverrider
}

// styleProvider implements the StyleProvider interface.
type styleProvider struct {
	gextras.Objector
}

var _ StyleProvider = (*styleProvider)(nil)

// WrapStyleProvider wraps a GObject to a type that implements interface
// StyleProvider. It is primarily used internally.
func WrapStyleProvider(obj *externglib.Object) StyleProvider {
	return StyleProvider{
		Objector: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProvider(obj), nil
}

// IconFactory returns the IconFactory defined to be in use for @path, or
// nil if none is defined.
func (p styleProvider) IconFactory(path *WidgetPath) IconFactory {
	var _arg0 *C.GtkStyleProvider
	var _arg1 *C.GtkWidgetPath

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))

	var _cret *C.GtkIconFactory

	cret = C.gtk_style_provider_get_icon_factory(_arg0, _arg1)

	var _iconFactory IconFactory

	_iconFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(IconFactory)

	return _iconFactory
}

// Style returns the style settings affecting a widget defined by @path, or
// nil if @provider doesn’t contemplate styling @path.
func (p styleProvider) Style(path *WidgetPath) StyleProperties {
	var _arg0 *C.GtkStyleProvider
	var _arg1 *C.GtkWidgetPath

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))

	var _cret *C.GtkStyleProperties

	cret = C.gtk_style_provider_get_style(_arg0, _arg1)

	var _styleProperties StyleProperties

	_styleProperties = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(StyleProperties)

	return _styleProperties
}

// StyleProperty looks up a widget style property as defined by @provider
// for the widget represented by @path.
func (p styleProvider) StyleProperty(path *WidgetPath, state StateFlags, pspec gobject.ParamSpec) (*externglib.Value, bool) {
	var _arg0 *C.GtkStyleProvider
	var _arg1 *C.GtkWidgetPath
	var _arg2 C.GtkStateFlags
	var _arg3 *C.GParamSpec

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))
	_arg2 = (C.GtkStateFlags)(state)
	_arg3 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))

	var _arg4 C.GValue
	var _cret C.gboolean

	cret = C.gtk_style_provider_get_style_property(_arg0, _arg1, _arg2, _arg3, &_arg4)

	var _value *externglib.Value
	var _ok bool

	_value = externglib.ValueFromNative(unsafe.Pointer(_arg4))
	if _cret {
		_ok = true
	}

	return _value, _ok
}