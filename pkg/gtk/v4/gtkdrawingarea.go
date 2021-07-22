// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// void _gotk4_gtk4_DrawingAreaDrawFunc(GtkDrawingArea*, cairo_t*, int, int, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drawing_area_get_type()), F: marshalDrawingAreaer},
	})
}

// DrawingAreaDrawFunc: whenever drawing_area needs to redraw, this function
// will be called.
//
// This function should exclusively redraw the contents of the drawing area and
// must not call any widget functions that cause changes.
type DrawingAreaDrawFunc func(drawingArea *DrawingArea, cr *cairo.Context, width int, height int)

//export _gotk4_gtk4_DrawingAreaDrawFunc
func _gotk4_gtk4_DrawingAreaDrawFunc(arg0 *C.GtkDrawingArea, arg1 *C.cairo_t, arg2 C.int, arg3 C.int, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var drawingArea *DrawingArea // out
	var cr *cairo.Context        // out
	var width int                // out
	var height int               // out

	drawingArea = wrapDrawingArea(externglib.Take(unsafe.Pointer(arg0)))
	cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})
	width = int(arg2)
	height = int(arg3)

	fn := v.(DrawingAreaDrawFunc)
	fn(drawingArea, cr, width, height)
}

// DrawingAreaOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DrawingAreaOverrider interface {
	Resize(width int, height int)
}

// DrawingArea: GtkDrawingArea is a widget that allows drawing with cairo.
//
// !An example GtkDrawingArea (drawingarea.png)
//
// It’s essentially a blank widget; you can draw on it. After creating a drawing
// area, the application may want to connect to:
//
// - The gtk.Widget::realize signal to take any necessary actions when the
// widget is instantiated on a particular display. (Create GDK resources in
// response to this signal.)
//
// - The gtk.DrawingArea::resize signal to take any necessary actions when the
// widget changes size.
//
// - Call gtk.DrawingArea.SetDrawFunc() to handle redrawing the contents of the
// widget.
//
// The following code portion demonstrates using a drawing area to display a
// circle in the normal widget foreground color.
//
// Simple GtkDrawingArea usage
//
//    static void
//    draw_function (GtkDrawingArea *area,
//                   cairo_t        *cr,
//                   int             width,
//                   int             height,
//                   gpointer        data)
//    {
//      GdkRGBA color;
//      GtkStyleContext *context;
//
//      context = gtk_widget_get_style_context (GTK_WIDGET (area));
//
//      cairo_arc (cr,
//                 width / 2.0, height / 2.0,
//                 MIN (width, height) / 2.0,
//                 0, 2 * G_PI);
//
//      gtk_style_context_get_color (context,
//                                   &color);
//      gdk_cairo_set_source_rgba (cr, &color);
//
//      cairo_fill (cr);
//    }
//
//    int
//    main (int argc, char **argv)
//    {
//      gtk_init ();
//
//      GtkWidget *area = gtk_drawing_area_new ();
//      gtk_drawing_area_set_content_width (GTK_DRAWING_AREA (area), 100);
//      gtk_drawing_area_set_content_height (GTK_DRAWING_AREA (area), 100);
//      gtk_drawing_area_set_draw_func (GTK_DRAWING_AREA (area),
//                                      draw_function,
//                                      NULL, NULL);
//      return 0;
//    }
//
//
// The draw function is normally called when a drawing area first comes
// onscreen, or when it’s covered by another window and then uncovered. You can
// also force a redraw by adding to the “damage region” of the drawing area’s
// window using gtk.Widget.QueueDraw(). This will cause the drawing area to call
// the draw function again.
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
// creating your own GtkWidget subclass.
type DrawingArea struct {
	Widget
}

func wrapDrawingArea(obj *externglib.Object) *DrawingArea {
	return &DrawingArea{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalDrawingAreaer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrawingArea(obj), nil
}

// NewDrawingArea creates a new drawing area.
func NewDrawingArea() *DrawingArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_drawing_area_new()

	var _drawingArea *DrawingArea // out

	_drawingArea = wrapDrawingArea(externglib.Take(unsafe.Pointer(_cret)))

	return _drawingArea
}

// ContentHeight retrieves the content height of the GtkDrawingArea.
func (self *DrawingArea) ContentHeight() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drawing_area_get_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ContentWidth retrieves the content width of the GtkDrawingArea.
func (self *DrawingArea) ContentWidth() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drawing_area_get_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetContentHeight sets the desired height of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual height passed to your draw function is larger
// than the height set here. You can use gtk.Widget.SetVAlign() to avoid that.
//
// If the height is set to 0 (the default), the drawing area may disappear.
func (self *DrawingArea) SetContentHeight(height int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(height)

	C.gtk_drawing_area_set_content_height(_arg0, _arg1)
}

// SetContentWidth sets the desired width of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual width passed to your draw function is larger
// than the width set here. You can use gtk.Widget.SetHAlign() to avoid that.
//
// If the width is set to 0 (the default), the drawing area may disappear.
func (self *DrawingArea) SetContentWidth(width int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(width)

	C.gtk_drawing_area_set_content_width(_arg0, _arg1)
}

// SetDrawFunc: setting a draw function is the main thing you want to do when
// using a drawing area.
//
// The draw function is called whenever GTK needs to draw the contents of the
// drawing area to the screen.
//
// The draw function will be called during the drawing stage of GTK. In the
// drawing stage it is not allowed to change properties of any GTK widgets or
// call any functions that would cause any properties to be changed. You should
// restrict yourself exclusively to drawing your contents in the draw function.
//
// If what you are drawing does change, call gtk.Widget.QueueDraw() on the
// drawing area. This will cause a redraw and will call draw_func again.
func (self *DrawingArea) SetDrawFunc(drawFunc DrawingAreaDrawFunc) {
	var _arg0 *C.GtkDrawingArea        // out
	var _arg1 C.GtkDrawingAreaDrawFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_DrawingAreaDrawFunc)
	_arg2 = C.gpointer(gbox.Assign(drawFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_drawing_area_set_draw_func(_arg0, _arg1, _arg2, _arg3)
}
