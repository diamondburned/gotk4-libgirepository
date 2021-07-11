// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_css_provider_get_type()), F: marshalCSSProviderer},
	})
}

// CSSProviderer describes CSSProvider's methods.
type CSSProviderer interface {
	// LoadFromData loads @data into @css_provider.
	LoadFromData(data []byte)
	// LoadFromFile loads the data contained in @file into @css_provider.
	LoadFromFile(file gio.Filer)
	// LoadFromPath loads the data contained in @path into @css_provider.
	LoadFromPath(path string)
	// LoadFromResource loads the data contained in the resource at
	// @resource_path into the @css_provider.
	LoadFromResource(resourcePath string)
	// LoadNamed loads a theme from the usual theme paths.
	LoadNamed(name string, variant string)
	// String converts the @provider into a string representation in CSS format.
	String() string
}

// CSSProvider: `GtkCssProvider` is an object implementing the
// `GtkStyleProvider` interface for CSS.
//
// It is able to parse CSS-like input in order to style widgets.
//
// An application can make GTK parse a specific CSS style sheet by calling
// [method@Gtk.CssProvider.load_from_file] or
// [method@Gtk.CssProvider.load_from_resource] and adding the provider with
// [method@Gtk.StyleContext.add_provider] or
// [func@Gtk.StyleContext.add_provider_for_display].
//
// In addition, certain files will be read when GTK is initialized. First, the
// file `$XDG_CONFIG_HOME/gtk-4.0/gtk.css` is loaded if it exists. Then, GTK
// loads the first existing file among
// `XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk-VARIANT.css`,
// `$HOME/.themes/THEME/gtk-VERSION/gtk-VARIANT.css`,
// `$XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk-VARIANT.css` and
// `DATADIR/share/themes/THEME/gtk-VERSION/gtk-VARIANT.css`, where `THEME` is
// the name of the current theme (see the [property@Gtk.Settings:gtk-theme-name]
// setting), `VARIANT` is the variant to load (see the
// [property@Gtk.Settings:gtk-application-prefer-dark-theme] setting), `DATADIR`
// is the prefix configured when GTK was compiled (unless overridden by the
// `GTK_DATA_PREFIX` environment variable), and `VERSION` is the GTK version
// number. If no file is found for the current version, GTK tries older versions
// all the way back to 4.0.
//
// To track errors while loading CSS, connect to the
// [signal@Gtk.CssProvider::parsing-error] signal.
type CSSProvider struct {
	*externglib.Object

	StyleProvider
}

var (
	_ CSSProviderer   = (*CSSProvider)(nil)
	_ gextras.Nativer = (*CSSProvider)(nil)
)

func wrapCSSProvider(obj *externglib.Object) CSSProviderer {
	return &CSSProvider{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalCSSProviderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCSSProvider(obj), nil
}

// NewCSSProvider returns a newly created `GtkCssProvider`.
func NewCSSProvider() *CSSProvider {
	var _cret *C.GtkCssProvider // in

	_cret = C.gtk_css_provider_new()

	var _cssProvider *CSSProvider // out

	_cssProvider = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*CSSProvider)

	return _cssProvider
}

// LoadFromData loads @data into @css_provider.
//
// This clears any previously loaded information.
func (cssProvider *CSSProvider) LoadFromData(data []byte) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.char
	var _arg2 C.gssize

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg2 = C.gssize(len(data))
	_arg1 = (*C.char)(unsafe.Pointer(&data[0]))

	C.gtk_css_provider_load_from_data(_arg0, _arg1, _arg2)
}

// LoadFromFile loads the data contained in @file into @css_provider.
//
// This clears any previously loaded information.
func (cssProvider *CSSProvider) LoadFromFile(file gio.Filer) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.GFile          // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))

	C.gtk_css_provider_load_from_file(_arg0, _arg1)
}

// LoadFromPath loads the data contained in @path into @css_provider.
//
// This clears any previously loaded information.
func (cssProvider *CSSProvider) LoadFromPath(path string) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_path(_arg0, _arg1)
}

// LoadFromResource loads the data contained in the resource at @resource_path
// into the @css_provider.
//
// This clears any previously loaded information.
func (cssProvider *CSSProvider) LoadFromResource(resourcePath string) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_resource(_arg0, _arg1)
}

// LoadNamed loads a theme from the usual theme paths.
//
// The actual process of finding the theme might change between releases, but it
// is guaranteed that this function uses the same mechanism to load the theme
// that GTK uses for loading its own theme.
func (provider *CSSProvider) LoadNamed(name string, variant string) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(provider.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(variant))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_css_provider_load_named(_arg0, _arg1, _arg2)
}

// String converts the @provider into a string representation in CSS format.
//
// Using [method@Gtk.CssProvider.load_from_data] with the return value from this
// function on a new provider created with [ctor@Gtk.CssProvider.new] will
// basically create a duplicate of this @provider.
func (provider *CSSProvider) String() string {
	var _arg0 *C.GtkCssProvider // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gtk_css_provider_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
