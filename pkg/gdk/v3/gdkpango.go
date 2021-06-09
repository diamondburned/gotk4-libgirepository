// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PangoContextGet creates a Context for the default GDK screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the default screen; if these options change it
// will not be updated. Using gtk_widget_get_pango_context() is more convenient
// if you want to keep a context around and track changes to the screen’s font
// rendering settings.
func PangoContextGet() pango.Context {
	var _cret *C.PangoContext

	cret = C.gdk_pango_context_get()

	var _context pango.Context

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(pango.Context)

	return _context
}

// PangoContextGetForDisplay creates a Context for @display.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the display; if these options change it will not
// be updated. Using gtk_widget_get_pango_context() is more convenient if you
// want to keep a context around and track changes to the font rendering
// settings.
func PangoContextGetForDisplay(display Display) pango.Context {
	var _arg1 *C.GdkDisplay

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	var _cret *C.PangoContext

	cret = C.gdk_pango_context_get_for_display(_arg1)

	var _context pango.Context

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(pango.Context)

	return _context
}

// PangoContextGetForScreen creates a Context for @screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the screen; if these options change it will not be
// updated. Using gtk_widget_get_pango_context() is more convenient if you want
// to keep a context around and track changes to the screen’s font rendering
// settings.
func PangoContextGetForScreen(screen Screen) pango.Context {
	var _arg1 *C.GdkScreen

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	var _cret *C.PangoContext

	cret = C.gdk_pango_context_get_for_screen(_arg1)

	var _context pango.Context

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(pango.Context)

	return _context
}

// PangoLayoutGetClipRegion obtains a clip region which contains the areas where
// the given ranges of text would be drawn. @x_origin and @y_origin are the top
// left point to center the layout. @index_ranges should contain ranges of bytes
// in the layout’s text.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn layout may in fact touch areas out of
// the clip region. The clip region is mainly useful for highlightling parts of
// text, such as when text is selected.
func PangoLayoutGetClipRegion(layout pango.Layout, xOrigin int, yOrigin int, indexRanges *int, nRanges int) *cairo.Region {
	var _arg1 *C.PangoLayout
	var _arg2 C.gint
	var _arg3 C.gint
	var _arg4 *C.gint
	var _arg5 C.gint

	_arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = C.gint(xOrigin)
	_arg3 = C.gint(yOrigin)
	_arg4 = *C.gint(indexRanges)
	_arg5 = C.gint(nRanges)

	var _cret *C.cairo_region_t

	cret = C.gdk_pango_layout_get_clip_region(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _region *cairo.Region

	_region = cairo.WrapRegion(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _region
}
