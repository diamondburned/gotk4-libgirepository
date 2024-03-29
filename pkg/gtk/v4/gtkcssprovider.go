// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_CssProvider_ConnectParsingError(gpointer, void*, GError*, guintptr);
import "C"

// GType values.
var (
	GTypeCSSProvider = coreglib.Type(girepository.MustFind("Gtk", "CssProvider").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCSSProvider, F: marshalCSSProvider},
	})
}

// CSSProvider: GtkCssProvider is an object implementing the GtkStyleProvider
// interface for CSS.
//
// It is able to parse CSS-like input in order to style widgets.
//
// An application can make GTK parse a specific CSS style sheet by calling
// gtk.CSSProvider.LoadFromFile() or gtk.CSSProvider.LoadFromResource() and
// adding the provider with gtk.StyleContext.AddProvider() or
// gtk.StyleContext().AddProviderForDisplay.
//
// In addition, certain files will be read when GTK is initialized. First, the
// file $XDG_CONFIG_HOME/gtk-4.0/gtk.css is loaded if it exists. Then, GTK loads
// the first existing file among
// XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk-VARIANT.css,
// $HOME/.themes/THEME/gtk-VERSION/gtk-VARIANT.css,
// $XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk-VARIANT.css and
// DATADIR/share/themes/THEME/gtk-VERSION/gtk-VARIANT.css, where THEME is the
// name of the current theme (see the gtk.Settings:gtk-theme-name setting),
// VARIANT is the variant to load (see the
// gtk.Settings:gtk-application-prefer-dark-theme setting), DATADIR is the
// prefix configured when GTK was compiled (unless overridden by the
// GTK_DATA_PREFIX environment variable), and VERSION is the GTK version number.
// If no file is found for the current version, GTK tries older versions all the
// way back to 4.0.
//
// To track errors while loading CSS, connect to the
// gtk.CSSProvider::parsing-error signal.
type CSSProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object

	StyleProvider
}

var (
	_ coreglib.Objector = (*CSSProvider)(nil)
)

func wrapCSSProvider(obj *coreglib.Object) *CSSProvider {
	return &CSSProvider{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalCSSProvider(p uintptr) (interface{}, error) {
	return wrapCSSProvider(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectParsingError signals that a parsing error occurred.
//
// The path, line and position describe the actual location of the error as
// accurately as possible.
//
// Parsing errors are never fatal, so the parsing will resume after the error.
// Errors may however cause parts of the given data or even all of it to not be
// parsed at all. So it is a useful idea to check that the parsing succeeds by
// connecting to this signal.
//
// Note that this signal may be emitted at any time as the css provider may opt
// to defer parsing parts or all of the input to a later time than when a
// loading function was called.
func (v *CSSProvider) ConnectParsingError(f func(section *CSSSection, err error)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "parsing-error", false, unsafe.Pointer(C._gotk4_gtk4_CssProvider_ConnectParsingError), f)
}
