// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expander_style_get_type()), F: marshalExpanderStyle},
		{T: externglib.Type(C.gtk_style_get_type()), F: marshalStyle},
	})
}

// ExpanderStyle: used to specify the style of the expanders drawn by a
// TreeView.
type ExpanderStyle int

const (
	// ExpanderStyleCollapsed: the style used for a collapsed subtree.
	ExpanderStyleCollapsed ExpanderStyle = 0
	// ExpanderStyleSemiCollapsed: intermediate style used during animation.
	ExpanderStyleSemiCollapsed ExpanderStyle = 1
	// ExpanderStyleSemiExpanded: intermediate style used during animation.
	ExpanderStyleSemiExpanded ExpanderStyle = 2
	// ExpanderStyleExpanded: the style used for an expanded subtree.
	ExpanderStyleExpanded ExpanderStyle = 3
)

func marshalExpanderStyle(p uintptr) (interface{}, error) {
	return ExpanderStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PaintArrow draws an arrow in the given rectangle on @cr using the given
// parameters. @arrow_type determines the direction of the arrow.
func PaintArrow(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, arrowType ArrowType, fill bool, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.GtkArrowType  // out
	var _arg8 C.gboolean      // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out
	var _arg11 C.gint         // out
	var _arg12 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.GtkArrowType)(arrowType)
	if fill {
		_arg8 = C.TRUE
	}
	_arg9 = (C.gint)(x)
	_arg10 = (C.gint)(y)
	_arg11 = (C.gint)(width)
	_arg12 = (C.gint)(height)

	C.gtk_paint_arrow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12)
}

// PaintBox draws a box on @cr with the given parameters.
func PaintBox(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_box(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintBoxGap draws a box in @cr using the given style and state and shadow
// type, leaving a gap in one side.
func PaintBoxGap(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
	var _arg1 *C.GtkStyle        // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.GtkStateType     // out
	var _arg4 C.GtkShadowType    // out
	var _arg5 *C.GtkWidget       // out
	var _arg6 *C.gchar           // out
	var _arg7 C.gint             // out
	var _arg8 C.gint             // out
	var _arg9 C.gint             // out
	var _arg10 C.gint            // out
	var _arg11 C.GtkPositionType // out
	var _arg12 C.gint            // out
	var _arg13 C.gint            // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)
	_arg11 = (C.GtkPositionType)(gapSide)
	_arg12 = (C.gint)(gapX)
	_arg13 = (C.gint)(gapWidth)

	C.gtk_paint_box_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, _arg13)
}

// PaintCheck draws a check button indicator in the given rectangle on @cr with
// the given parameters.
func PaintCheck(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_check(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintDiamond draws a diamond in the given rectangle on @window using the
// given parameters.
func PaintDiamond(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_diamond(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintExpander draws an expander as used in TreeView. @x and @y specify the
// center the expander. The size of the expander is determined by the
// “expander-size” style property of @widget. (If widget is not specified or
// doesn’t have an “expander-size” property, an unspecified default size will be
// used, since the caller doesn't have sufficient information to position the
// expander, this is likely not useful.) The expander is expander_size pixels
// tall in the collapsed position and expander_size pixels wide in the expanded
// position.
func PaintExpander(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, x int, y int, expanderStyle ExpanderStyle) {
	var _arg1 *C.GtkStyle        // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.GtkStateType     // out
	var _arg4 *C.GtkWidget       // out
	var _arg5 *C.gchar           // out
	var _arg6 C.gint             // out
	var _arg7 C.gint             // out
	var _arg8 C.GtkExpanderStyle // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.gint)(x)
	_arg7 = (C.gint)(y)
	_arg8 = (C.GtkExpanderStyle)(expanderStyle)

	C.gtk_paint_expander(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// PaintExtension draws an extension, i.e. a notebook tab.
func PaintExtension(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int, gapSide PositionType) {
	var _arg1 *C.GtkStyle        // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.GtkStateType     // out
	var _arg4 C.GtkShadowType    // out
	var _arg5 *C.GtkWidget       // out
	var _arg6 *C.gchar           // out
	var _arg7 C.gint             // out
	var _arg8 C.gint             // out
	var _arg9 C.gint             // out
	var _arg10 C.gint            // out
	var _arg11 C.GtkPositionType // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)
	_arg11 = (C.GtkPositionType)(gapSide)

	C.gtk_paint_extension(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintFlatBox draws a flat box on @cr with the given parameters.
func PaintFlatBox(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_flat_box(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintFocus draws a focus indicator around the given rectangle on @cr using
// the given style.
func PaintFocus(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle    // out
	var _arg2 *C.cairo_t     // out
	var _arg3 C.GtkStateType // out
	var _arg4 *C.GtkWidget   // out
	var _arg5 *C.gchar       // out
	var _arg6 C.gint         // out
	var _arg7 C.gint         // out
	var _arg8 C.gint         // out
	var _arg9 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.gint)(x)
	_arg7 = (C.gint)(y)
	_arg8 = (C.gint)(width)
	_arg9 = (C.gint)(height)

	C.gtk_paint_focus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// PaintHandle draws a handle as used in HandleBox and Paned.
func PaintHandle(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
	var _arg1 *C.GtkStyle       // out
	var _arg2 *C.cairo_t        // out
	var _arg3 C.GtkStateType    // out
	var _arg4 C.GtkShadowType   // out
	var _arg5 *C.GtkWidget      // out
	var _arg6 *C.gchar          // out
	var _arg7 C.gint            // out
	var _arg8 C.gint            // out
	var _arg9 C.gint            // out
	var _arg10 C.gint           // out
	var _arg11 C.GtkOrientation // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)
	_arg11 = (C.GtkOrientation)(orientation)

	C.gtk_paint_handle(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintHline draws a horizontal line from (@x1, @y) to (@x2, @y) in @cr using
// the given style and state.
func PaintHline(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, x1 int, x2 int, y int) {
	var _arg1 *C.GtkStyle    // out
	var _arg2 *C.cairo_t     // out
	var _arg3 C.GtkStateType // out
	var _arg4 *C.GtkWidget   // out
	var _arg5 *C.gchar       // out
	var _arg6 C.gint         // out
	var _arg7 C.gint         // out
	var _arg8 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.gint)(x1)
	_arg7 = (C.gint)(x2)
	_arg8 = (C.gint)(y)

	C.gtk_paint_hline(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// PaintLayout draws a layout on @cr using the given parameters.
func PaintLayout(style Style, cr *cairo.Context, stateType StateType, useText bool, widget Widget, detail string, x int, y int, layout pango.Layout) {
	var _arg1 *C.GtkStyle    // out
	var _arg2 *C.cairo_t     // out
	var _arg3 C.GtkStateType // out
	var _arg4 C.gboolean     // out
	var _arg5 *C.GtkWidget   // out
	var _arg6 *C.gchar       // out
	var _arg7 C.gint         // out
	var _arg8 C.gint         // out
	var _arg9 *C.PangoLayout // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	if useText {
		_arg4 = C.TRUE
	}
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_paint_layout(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// PaintOption draws a radio button indicator in the given rectangle on @cr with
// the given parameters.
func PaintOption(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_option(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintResizeGrip draws a resize grip in the given rectangle on @cr using the
// given parameters.
func PaintResizeGrip(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, edge gdk.WindowEdge, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 *C.GtkWidget    // out
	var _arg5 *C.gchar        // out
	var _arg6 C.GdkWindowEdge // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.GdkWindowEdge)(edge)
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_resize_grip(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintShadow draws a shadow around the given rectangle in @cr using the given
// style and state and shadow type.
func PaintShadow(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_shadow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintShadowGap draws a shadow around the given rectangle in @cr using the
// given style and state and shadow type, leaving a gap in one side.
func PaintShadowGap(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
	var _arg1 *C.GtkStyle        // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.GtkStateType     // out
	var _arg4 C.GtkShadowType    // out
	var _arg5 *C.GtkWidget       // out
	var _arg6 *C.gchar           // out
	var _arg7 C.gint             // out
	var _arg8 C.gint             // out
	var _arg9 C.gint             // out
	var _arg10 C.gint            // out
	var _arg11 C.GtkPositionType // out
	var _arg12 C.gint            // out
	var _arg13 C.gint            // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)
	_arg11 = (C.GtkPositionType)(gapSide)
	_arg12 = (C.gint)(gapX)
	_arg13 = (C.gint)(gapWidth)

	C.gtk_paint_shadow_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, _arg13)
}

// PaintSlider draws a slider in the given rectangle on @cr using the given
// style and orientation.
func PaintSlider(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
	var _arg1 *C.GtkStyle       // out
	var _arg2 *C.cairo_t        // out
	var _arg3 C.GtkStateType    // out
	var _arg4 C.GtkShadowType   // out
	var _arg5 *C.GtkWidget      // out
	var _arg6 *C.gchar          // out
	var _arg7 C.gint            // out
	var _arg8 C.gint            // out
	var _arg9 C.gint            // out
	var _arg10 C.gint           // out
	var _arg11 C.GtkOrientation // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)
	_arg11 = (C.GtkOrientation)(orientation)

	C.gtk_paint_slider(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintSpinner draws a spinner on @window using the given parameters.
func PaintSpinner(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, step uint, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle    // out
	var _arg2 *C.cairo_t     // out
	var _arg3 C.GtkStateType // out
	var _arg4 *C.GtkWidget   // out
	var _arg5 *C.gchar       // out
	var _arg6 C.guint        // out
	var _arg7 C.gint         // out
	var _arg8 C.gint         // out
	var _arg9 C.gint         // out
	var _arg10 C.gint        // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.guint)(step)
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_spinner(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintTab draws an option menu tab (i.e. the up and down pointing arrows) in
// the given rectangle on @cr using the given parameters.
func PaintTab(style Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widget, detail string, x int, y int, width int, height int) {
	var _arg1 *C.GtkStyle     // out
	var _arg2 *C.cairo_t      // out
	var _arg3 C.GtkStateType  // out
	var _arg4 C.GtkShadowType // out
	var _arg5 *C.GtkWidget    // out
	var _arg6 *C.gchar        // out
	var _arg7 C.gint          // out
	var _arg8 C.gint          // out
	var _arg9 C.gint          // out
	var _arg10 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.GtkShadowType)(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (C.gint)(x)
	_arg8 = (C.gint)(y)
	_arg9 = (C.gint)(width)
	_arg10 = (C.gint)(height)

	C.gtk_paint_tab(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintVline draws a vertical line from (@x, @y1_) to (@x, @y2_) in @cr using
// the given style and state.
func PaintVline(style Style, cr *cairo.Context, stateType StateType, widget Widget, detail string, y1 int, y2 int, x int) {
	var _arg1 *C.GtkStyle    // out
	var _arg2 *C.cairo_t     // out
	var _arg3 C.GtkStateType // out
	var _arg4 *C.GtkWidget   // out
	var _arg5 *C.gchar       // out
	var _arg6 C.gint         // out
	var _arg7 C.gint         // out
	var _arg8 C.gint         // out

	_arg1 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (C.gint)(y1)
	_arg7 = (C.gint)(y2)
	_arg8 = (C.gint)(x)

	C.gtk_paint_vline(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// Style: a Style object encapsulates the information that provides the look and
// feel for a widget.
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
type Style interface {
	gextras.Objector

	ApplyDefaultBackground(cr *cairo.Context, window gdk.Window, stateType StateType, x int, y int, width int, height int)
	// Copy creates a copy of the passed in Style object.
	Copy() Style
	// Detach detaches a style from a window. If the style is not attached to
	// any windows anymore, it is unrealized. See gtk_style_attach().
	Detach()
	// StyleProperty queries the value of a style property corresponding to a
	// widget class is in the given style.
	StyleProperty(widgetType externglib.Type, propertyName string) *externglib.Value
	// HasContext returns whether @style has an associated StyleContext.
	HasContext() bool
	// LookupColor looks up @color_name in the style’s logical color mappings,
	// filling in @color and returning true if found, otherwise returning false.
	// Do not cache the found mapping, because it depends on the Style and might
	// change when a theme switch occurs.
	LookupColor(colorName string) (gdk.Color, bool)
	// LookupIconSet looks up @stock_id in the icon factories associated with
	// @style and the default icon factory, returning an icon set if found,
	// otherwise nil.
	LookupIconSet(stockId string) *IconSet
	// RenderIcon renders the icon specified by @source at the given @size
	// according to the given parameters and returns the result in a pixbuf.
	RenderIcon(source *IconSource, direction TextDirection, state StateType, size int, widget Widget, detail string) gdkpixbuf.Pixbuf
	// SetBackground sets the background of @window to the background color or
	// pixmap specified by @style for the given state.
	SetBackground(window gdk.Window, stateType StateType)
}

// style implements the Style class.
type style struct {
	gextras.Objector
}

var _ Style = (*style)(nil)

// WrapStyle wraps a GObject to the right type. It is
// primarily used internally.
func WrapStyle(obj *externglib.Object) Style {
	return style{
		Objector: obj,
	}
}

func marshalStyle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyle(obj), nil
}

// NewStyle constructs a class Style.
func NewStyle() Style {
	var _cret C.GtkStyle // in

	_cret = C.gtk_style_new()

	var _style Style // out

	_style = WrapStyle(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _style
}

func (s style) ApplyDefaultBackground(cr *cairo.Context, window gdk.Window, stateType StateType, x int, y int, width int, height int) {
	var _arg0 *C.GtkStyle    // out
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.GdkWindow   // out
	var _arg3 C.GtkStateType // out
	var _arg4 C.gint         // out
	var _arg5 C.gint         // out
	var _arg6 C.gint         // out
	var _arg7 C.gint         // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg3 = (C.GtkStateType)(stateType)
	_arg4 = (C.gint)(x)
	_arg5 = (C.gint)(y)
	_arg6 = (C.gint)(width)
	_arg7 = (C.gint)(height)

	C.gtk_style_apply_default_background(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// Copy creates a copy of the passed in Style object.
func (s style) Copy() Style {
	var _arg0 *C.GtkStyle // out
	var _cret *C.GtkStyle // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_style_copy(_arg0)

	var _ret Style // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Style)

	return _ret
}

// Detach detaches a style from a window. If the style is not attached to
// any windows anymore, it is unrealized. See gtk_style_attach().
func (s style) Detach() {
	var _arg0 *C.GtkStyle // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))

	C.gtk_style_detach(_arg0)
}

// StyleProperty queries the value of a style property corresponding to a
// widget class is in the given style.
func (s style) StyleProperty(widgetType externglib.Type, propertyName string) *externglib.Value {
	var _arg0 *C.GtkStyle // out
	var _arg1 C.GType     // out
	var _arg2 *C.gchar    // out
	var _arg3 C.GValue    // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = C.GType(widgetType)
	_arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_style_get_style_property(_arg0, _arg1, _arg2, &_arg3)

	var _value *externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_arg3))

	return _value
}

// HasContext returns whether @style has an associated StyleContext.
func (s style) HasContext() bool {
	var _arg0 *C.GtkStyle // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_style_has_context(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupColor looks up @color_name in the style’s logical color mappings,
// filling in @color and returning true if found, otherwise returning false.
// Do not cache the found mapping, because it depends on the Style and might
// change when a theme switch occurs.
func (s style) LookupColor(colorName string) (gdk.Color, bool) {
	var _arg0 *C.GtkStyle // out
	var _arg1 *C.gchar    // out
	var _color gdk.Color
	var _cret C.gboolean // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_color(_arg0, _arg1, (*C.GdkColor)(unsafe.Pointer(&_color)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// LookupIconSet looks up @stock_id in the icon factories associated with
// @style and the default icon factory, returning an icon set if found,
// otherwise nil.
func (s style) LookupIconSet(stockId string) *IconSet {
	var _arg0 *C.GtkStyle   // out
	var _arg1 *C.gchar      // out
	var _cret *C.GtkIconSet // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_icon_set(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = WrapIconSet(unsafe.Pointer(_cret))

	return _iconSet
}

// RenderIcon renders the icon specified by @source at the given @size
// according to the given parameters and returns the result in a pixbuf.
func (s style) RenderIcon(source *IconSource, direction TextDirection, state StateType, size int, widget Widget, detail string) gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkStyle        // out
	var _arg1 *C.GtkIconSource   // out
	var _arg2 C.GtkTextDirection // out
	var _arg3 C.GtkStateType     // out
	var _arg4 C.GtkIconSize      // out
	var _arg5 *C.GtkWidget       // out
	var _arg6 *C.gchar           // out
	var _cret *C.GdkPixbuf       // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkIconSource)(unsafe.Pointer(source.Native()))
	_arg2 = (C.GtkTextDirection)(direction)
	_arg3 = (C.GtkStateType)(state)
	_arg4 = (C.GtkIconSize)(size)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(C.CString(detail))
	defer C.free(unsafe.Pointer(_arg6))

	_cret = C.gtk_style_render_icon(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

// SetBackground sets the background of @window to the background color or
// pixmap specified by @style for the given state.
func (s style) SetBackground(window gdk.Window, stateType StateType) {
	var _arg0 *C.GtkStyle    // out
	var _arg1 *C.GdkWindow   // out
	var _arg2 C.GtkStateType // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (C.GtkStateType)(stateType)

	C.gtk_style_set_background(_arg0, _arg1, _arg2)
}

type ThemeEngine struct {
	native C.GtkThemeEngine
}

// WrapThemeEngine wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapThemeEngine(ptr unsafe.Pointer) *ThemeEngine {
	if ptr == nil {
		return nil
	}

	return (*ThemeEngine)(ptr)
}

// Native returns the underlying C source pointer.
func (t *ThemeEngine) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
