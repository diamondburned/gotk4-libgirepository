// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drawing_area_get_type()), F: marshalDrawingArea},
	})
}

// DrawingArea: `GtkDrawingArea` is a widget that allows drawing with cairo.
//
// !An example GtkDrawingArea (drawingarea.png)
//
// It’s essentially a blank widget; you can draw on it. After creating a drawing
// area, the application may want to connect to:
//
// - The [signal@Gtk.Widget::realize] signal to take any necessary actions when
// the widget is instantiated on a particular display. (Create GDK resources in
// response to this signal.)
//
// - The [signal@Gtk.DrawingArea::resize] signal to take any necessary actions
// when the widget changes size.
//
// - Call [method@Gtk.DrawingArea.set_draw_func] to handle redrawing the
// contents of the widget.
//
// The following code portion demonstrates using a drawing area to display a
// circle in the normal widget foreground color.
//
//
// Simple GtkDrawingArea usage
//
// “`c static void draw_function (GtkDrawingArea *area, cairo_t *cr, int width,
// int height, gpointer data) { GdkRGBA color; GtkStyleContext *context;
//
//    context = gtk_widget_get_style_context (GTK_WIDGET (area));
//
//    cairo_arc (cr,
//               width / 2.0, height / 2.0,
//               MIN (width, height) / 2.0,
//               0, 2 * G_PI);
//
//    gtk_style_context_get_color (context,
//                                 &color);
//    gdk_cairo_set_source_rgba (cr, &color);
//
//    cairo_fill (cr);
//
// }
//
// int main (int argc, char **argv) { gtk_init ();
//
//    GtkWidget *area = gtk_drawing_area_new ();
//    gtk_drawing_area_set_content_width (GTK_DRAWING_AREA (area), 100);
//    gtk_drawing_area_set_content_height (GTK_DRAWING_AREA (area), 100);
//    gtk_drawing_area_set_draw_func (GTK_DRAWING_AREA (area),
//                                    draw_function,
//                                    NULL, NULL);
//    return 0;
//
// } “`
//
// The draw function is normally called when a drawing area first comes
// onscreen, or when it’s covered by another window and then uncovered. You can
// also force a redraw by adding to the “damage region” of the drawing area’s
// window using [method@Gtk.Widget.queue_draw]. This will cause the drawing area
// to call the draw function again.
//
// The available routines for drawing are documented on the [GDK Drawing
// Primitives][gdk4-Cairo-Interaction] page and the cairo documentation.
//
// To receive mouse events on a drawing area, you will need to use event
// controllers. To receive keyboard events, you will need to set the “can-focus”
// property on the drawing area, and you should probably draw some user-visible
// indication that the drawing area is focused.
//
// If you need more complex control over your widget, you should consider
// creating your own `GtkWidget` subclass.
type DrawingArea interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// ContentHeight retrieves the content height of the `GtkDrawingArea`.
	ContentHeight() int
	// ContentWidth retrieves the content width of the `GtkDrawingArea`.
	ContentWidth() int
	// SetContentHeight sets the desired height of the contents of the drawing
	// area.
	//
	// Note that because widgets may be allocated larger sizes than they
	// requested, it is possible that the actual height passed to your draw
	// function is larger than the height set here. You can use
	// [method@Gtk.Widget.set_valign] to avoid that.
	//
	// If the height is set to 0 (the default), the drawing area may disappear.
	SetContentHeight(height int)
	// SetContentWidth sets the desired width of the contents of the drawing
	// area.
	//
	// Note that because widgets may be allocated larger sizes than they
	// requested, it is possible that the actual width passed to your draw
	// function is larger than the width set here. You can use
	// [method@Gtk.Widget.set_halign] to avoid that.
	//
	// If the width is set to 0 (the default), the drawing area may disappear.
	SetContentWidth(width int)
}

// drawingArea implements the DrawingArea class.
type drawingArea struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ DrawingArea = (*drawingArea)(nil)

// WrapDrawingArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawingArea(obj *externglib.Object) DrawingArea {
	return drawingArea{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawingArea(obj), nil
}

// ContentHeight retrieves the content height of the `GtkDrawingArea`.
func (s drawingArea) ContentHeight() int {
	var _arg0 *C.GtkDrawingArea // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gtk_drawing_area_get_content_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// ContentWidth retrieves the content width of the `GtkDrawingArea`.
func (s drawingArea) ContentWidth() int {
	var _arg0 *C.GtkDrawingArea // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gtk_drawing_area_get_content_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// SetContentHeight sets the desired height of the contents of the drawing
// area.
//
// Note that because widgets may be allocated larger sizes than they
// requested, it is possible that the actual height passed to your draw
// function is larger than the height set here. You can use
// [method@Gtk.Widget.set_valign] to avoid that.
//
// If the height is set to 0 (the default), the drawing area may disappear.
func (s drawingArea) SetContentHeight(height int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(height)

	C.gtk_drawing_area_set_content_height(_arg0, _arg1)
}

// SetContentWidth sets the desired width of the contents of the drawing
// area.
//
// Note that because widgets may be allocated larger sizes than they
// requested, it is possible that the actual width passed to your draw
// function is larger than the width set here. You can use
// [method@Gtk.Widget.set_halign] to avoid that.
//
// If the width is set to 0 (the default), the drawing area may disappear.
func (s drawingArea) SetContentWidth(width int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(width)

	C.gtk_drawing_area_set_content_width(_arg0, _arg1)
}
