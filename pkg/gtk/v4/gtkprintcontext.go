// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
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
		{T: externglib.Type(C.gtk_print_context_get_type()), F: marshalPrintContext},
	})
}

// PrintContext: `GtkPrintContext` encapsulates context information that is
// required when drawing pages for printing.
//
// This includes the cairo context and important parameters like page size and
// resolution. It also lets you easily create [class@Pango.Layout] and
// [class@Pango.Context] objects that match the font metrics of the cairo
// surface.
//
// `GtkPrintContext` objects get passed to the
// [signal@Gtk.PrintOperation::begin-print],
// [signal@Gtk.PrintOperation::end-print],
// [signal@Gtk.PrintOperation::request-page-setup] and
// [signal@Gtk.PrintOperation::draw-page] signals on the
// [class@Gtk.PrintOperation] object.
//
// Using GtkPrintContext in a ::draw-page callback
//
// “`c static void draw_page (GtkPrintOperation *operation, GtkPrintContext
// *context, int page_nr) { cairo_t *cr; PangoLayout *layout;
// PangoFontDescription *desc;
//
//    cr = gtk_print_context_get_cairo_context (context);
//
//    // Draw a red rectangle, as wide as the paper (inside the margins)
//    cairo_set_source_rgb (cr, 1.0, 0, 0);
//    cairo_rectangle (cr, 0, 0, gtk_print_context_get_width (context), 50);
//
//    cairo_fill (cr);
//
//    // Draw some lines
//    cairo_move_to (cr, 20, 10);
//    cairo_line_to (cr, 40, 20);
//    cairo_arc (cr, 60, 60, 20, 0, M_PI);
//    cairo_line_to (cr, 80, 20);
//
//    cairo_set_source_rgb (cr, 0, 0, 0);
//    cairo_set_line_width (cr, 5);
//    cairo_set_line_cap (cr, CAIRO_LINE_CAP_ROUND);
//    cairo_set_line_join (cr, CAIRO_LINE_JOIN_ROUND);
//
//    cairo_stroke (cr);
//
//    // Draw some text
//    layout = gtk_print_context_create_pango_layout (context);
//    pango_layout_set_text (layout, "Hello World! Printing is easy", -1);
//    desc = pango_font_description_from_string ("sans 28");
//    pango_layout_set_font_description (layout, desc);
//    pango_font_description_free (desc);
//
//    cairo_move_to (cr, 30, 20);
//    pango_cairo_layout_path (cr, layout);
//
//    // Font Outline
//    cairo_set_source_rgb (cr, 0.93, 1.0, 0.47);
//    cairo_set_line_width (cr, 0.5);
//    cairo_stroke_preserve (cr);
//
//    // Font Fill
//    cairo_set_source_rgb (cr, 0, 0.0, 1.0);
//    cairo_fill (cr);
//
//    g_object_unref (layout);
//
// } “`
type PrintContext interface {
	gextras.Objector

	// CreatePangoContextPrintContext creates a new `PangoContext` that can be
	// used with the `GtkPrintContext`.
	CreatePangoContextPrintContext() pango.Context
	// CreatePangoLayoutPrintContext creates a new `PangoLayout` that is
	// suitable for use with the `GtkPrintContext`.
	CreatePangoLayoutPrintContext() pango.Layout
	// CairoContext obtains the cairo context that is associated with the
	// `GtkPrintContext`.
	CairoContext() *cairo.Context
	// DPIX obtains the horizontal resolution of the `GtkPrintContext`, in dots
	// per inch.
	DPIX() float64
	// DPIY obtains the vertical resolution of the `GtkPrintContext`, in dots
	// per inch.
	DPIY() float64
	// HardMargins obtains the hardware printer margins of the
	// `GtkPrintContext`, in units.
	HardMargins() (top float64, bottom float64, left float64, right float64, ok bool)
	// Height obtains the height of the `GtkPrintContext`, in pixels.
	Height() float64
	// PageSetup obtains the `GtkPageSetup` that determines the page dimensions
	// of the `GtkPrintContext`.
	PageSetup() PageSetup
	// PangoFontmap returns a `PangoFontMap` that is suitable for use with the
	// `GtkPrintContext`.
	PangoFontmap() pango.FontMap
	// Width obtains the width of the `GtkPrintContext`, in pixels.
	Width() float64
	// SetCairoContextPrintContext sets a new cairo context on a print context.
	//
	// This function is intended to be used when implementing an internal print
	// preview, it is not needed for printing, since GTK itself creates a
	// suitable cairo context in that case.
	SetCairoContextPrintContext(cr *cairo.Context, dpiX float64, dpiY float64)
}

// printContext implements the PrintContext class.
type printContext struct {
	gextras.Objector
}

// WrapPrintContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrintContext(obj *externglib.Object) PrintContext {
	return printContext{
		Objector: obj,
	}
}

func marshalPrintContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintContext(obj), nil
}

func (c printContext) CreatePangoContextPrintContext() pango.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoContext    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_create_pango_context(_arg0)

	var _ret pango.Context // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(pango.Context)

	return _ret
}

func (c printContext) CreatePangoLayoutPrintContext() pango.Layout {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoLayout     // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_create_pango_layout(_arg0)

	var _layout pango.Layout // out

	_layout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(pango.Layout)

	return _layout
}

func (c printContext) CairoContext() *cairo.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_cairo_context(_arg0)

	var _ret *cairo.Context // out

	_ret = (*cairo.Context)(unsafe.Pointer(_cret))

	return _ret
}

func (c printContext) DPIX() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_dpi_x(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (c printContext) DPIY() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_dpi_y(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (c printContext) HardMargins() (top float64, bottom float64, left float64, right float64, ok bool) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 C.double           // in
	var _arg2 C.double           // in
	var _arg3 C.double           // in
	var _arg4 C.double           // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_hard_margins(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _top float64    // out
	var _bottom float64 // out
	var _left float64   // out
	var _right float64  // out
	var _ok bool        // out

	_top = float64(_arg1)
	_bottom = float64(_arg2)
	_left = float64(_arg3)
	_right = float64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _top, _bottom, _left, _right, _ok
}

func (c printContext) Height() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_height(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (c printContext) PageSetup() PageSetup {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.GtkPageSetup    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_page_setup(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

func (c printContext) PangoFontmap() pango.FontMap {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoFontMap    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_pango_fontmap(_arg0)

	var _fontMap pango.FontMap // out

	_fontMap = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.FontMap)

	return _fontMap
}

func (c printContext) Width() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_print_context_get_width(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (c printContext) SetCairoContextPrintContext(cr *cairo.Context, dpiX float64, dpiY float64) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 *C.cairo_t         // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg2 = C.double(dpiX)
	_arg3 = C.double(dpiY)

	C.gtk_print_context_set_cairo_context(_arg0, _arg1, _arg2, _arg3)
}
