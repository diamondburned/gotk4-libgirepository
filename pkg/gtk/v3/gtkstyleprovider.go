// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern GtkIconFactory* _gotk4_gtk3_StyleProviderIface_get_icon_factory(GtkStyleProvider*, GtkWidgetPath*);
// extern GtkStyleProperties* _gotk4_gtk3_StyleProviderIface_get_style(GtkStyleProvider*, GtkWidgetPath*);
import "C"

// GType values.
var (
	GTypeStyleProvider = coreglib.Type(C.gtk_style_provider_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleProvider, F: marshalStyleProvider},
	})
}

// STYLE_PROVIDER_PRIORITY_APPLICATION: priority that can be used when adding a
// StyleProvider for application-specific style information.
const STYLE_PROVIDER_PRIORITY_APPLICATION = 600

// STYLE_PROVIDER_PRIORITY_FALLBACK: priority used for default style information
// that is used in the absence of themes.
//
// Note that this is not very useful for providing default styling for custom
// style classes - themes are likely to override styling provided at this
// priority with catch-all * {...} rules.
const STYLE_PROVIDER_PRIORITY_FALLBACK = 1

// STYLE_PROVIDER_PRIORITY_SETTINGS: priority used for style information
// provided via Settings.
//
// This priority is higher than K_STYLE_PROVIDER_PRIORITY_THEME to let settings
// override themes.
const STYLE_PROVIDER_PRIORITY_SETTINGS = 400

// STYLE_PROVIDER_PRIORITY_THEME: priority used for style information provided
// by themes.
const STYLE_PROVIDER_PRIORITY_THEME = 200

// STYLE_PROVIDER_PRIORITY_USER: priority used for the style information from
// XDG_CONFIG_HOME/gtk-3.0/gtk.css.
//
// You should not use priorities higher than this, to give the user the last
// word.
const STYLE_PROVIDER_PRIORITY_USER = 800

// StyleProviderOverrider contains methods that are overridable.
type StyleProviderOverrider interface {
	// IconFactory returns the IconFactory defined to be in use for path, or
	// NULL if none is defined.
	//
	// Deprecated: Will always return NULL for all GTK-provided style providers.
	//
	// The function takes the following parameters:
	//
	//    - path to query.
	//
	// The function returns the following values:
	//
	//    - iconFactory (optional): icon factory to use for path, or NULL.
	//
	IconFactory(path *WidgetPath) *IconFactory
	// Style returns the style settings affecting a widget defined by path, or
	// NULL if provider doesn’t contemplate styling path.
	//
	// Deprecated: Will always return NULL for all GTK-provided style providers
	// as the interface cannot correctly work the way CSS is specified.
	//
	// The function takes the following parameters:
	//
	//    - path to query.
	//
	// The function returns the following values:
	//
	//    - styleProperties (optional) containing the style settings affecting
	//      path.
	//
	Style(path *WidgetPath) *StyleProperties
}

// StyleProvider is an interface used to provide style information to a
// StyleContext. See gtk_style_context_add_provider() and
// gtk_style_context_add_provider_for_screen().
//
// StyleProvider wraps an interface. This means the user can get the
// underlying type by calling Cast().
type StyleProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleProvider)(nil)
)

// StyleProviderer describes StyleProvider's interface methods.
type StyleProviderer interface {
	coreglib.Objector

	// IconFactory returns the IconFactory defined to be in use for path, or
	// NULL if none is defined.
	IconFactory(path *WidgetPath) *IconFactory
	// Style returns the style settings affecting a widget defined by path, or
	// NULL if provider doesn’t contemplate styling path.
	Style(path *WidgetPath) *StyleProperties
}

var _ StyleProviderer = (*StyleProvider)(nil)

func ifaceInitStyleProviderer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkStyleProviderIface)(unsafe.Pointer(gifacePtr))
	iface.get_icon_factory = (*[0]byte)(C._gotk4_gtk3_StyleProviderIface_get_icon_factory)
	iface.get_style = (*[0]byte)(C._gotk4_gtk3_StyleProviderIface_get_style)
}

//export _gotk4_gtk3_StyleProviderIface_get_icon_factory
func _gotk4_gtk3_StyleProviderIface_get_icon_factory(arg0 *C.GtkStyleProvider, arg1 *C.GtkWidgetPath) (cret *C.GtkIconFactory) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(StyleProviderOverrider)

	var _path *WidgetPath // out

	_path = (*WidgetPath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_widget_path_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_path)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_widget_path_unref((*C.GtkWidgetPath)(intern.C))
		},
	)

	iconFactory := iface.IconFactory(_path)

	if iconFactory != nil {
		cret = (*C.GtkIconFactory)(unsafe.Pointer(coreglib.InternObject(iconFactory).Native()))
	}

	return cret
}

//export _gotk4_gtk3_StyleProviderIface_get_style
func _gotk4_gtk3_StyleProviderIface_get_style(arg0 *C.GtkStyleProvider, arg1 *C.GtkWidgetPath) (cret *C.GtkStyleProperties) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(StyleProviderOverrider)

	var _path *WidgetPath // out

	_path = (*WidgetPath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.gtk_widget_path_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_path)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_widget_path_unref((*C.GtkWidgetPath)(intern.C))
		},
	)

	styleProperties := iface.Style(_path)

	if styleProperties != nil {
		cret = (*C.GtkStyleProperties)(unsafe.Pointer(coreglib.InternObject(styleProperties).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(styleProperties).Native()))
	}

	return cret
}

func wrapStyleProvider(obj *coreglib.Object) *StyleProvider {
	return &StyleProvider{
		Object: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	return wrapStyleProvider(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// IconFactory returns the IconFactory defined to be in use for path, or NULL if
// none is defined.
//
// Deprecated: Will always return NULL for all GTK-provided style providers.
//
// The function takes the following parameters:
//
//    - path to query.
//
// The function returns the following values:
//
//    - iconFactory (optional): icon factory to use for path, or NULL.
//
func (provider *StyleProvider) IconFactory(path *WidgetPath) *IconFactory {
	var _arg0 *C.GtkStyleProvider // out
	var _arg1 *C.GtkWidgetPath    // out
	var _cret *C.GtkIconFactory   // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkWidgetPath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_style_provider_get_icon_factory(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(path)

	var _iconFactory *IconFactory // out

	if _cret != nil {
		_iconFactory = wrapIconFactory(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _iconFactory
}

// Style returns the style settings affecting a widget defined by path, or NULL
// if provider doesn’t contemplate styling path.
//
// Deprecated: Will always return NULL for all GTK-provided style providers as
// the interface cannot correctly work the way CSS is specified.
//
// The function takes the following parameters:
//
//    - path to query.
//
// The function returns the following values:
//
//    - styleProperties (optional) containing the style settings affecting path.
//
func (provider *StyleProvider) Style(path *WidgetPath) *StyleProperties {
	var _arg0 *C.GtkStyleProvider   // out
	var _arg1 *C.GtkWidgetPath      // out
	var _cret *C.GtkStyleProperties // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkWidgetPath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_style_provider_get_style(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(path)

	var _styleProperties *StyleProperties // out

	if _cret != nil {
		_styleProperties = wrapStyleProperties(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _styleProperties
}
