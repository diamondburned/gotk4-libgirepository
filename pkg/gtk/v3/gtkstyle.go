// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	// ExpanderCollapsed: style used for a collapsed subtree.
	ExpanderCollapsed ExpanderStyle = iota
	// ExpanderSemiCollapsed: intermediate style used during animation.
	ExpanderSemiCollapsed
	// ExpanderSemiExpanded: intermediate style used during animation.
	ExpanderSemiExpanded
	// ExpanderExpanded: style used for an expanded subtree.
	ExpanderExpanded
)

func marshalExpanderStyle(p uintptr) (interface{}, error) {
	return ExpanderStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ExpanderStyle.
func (e ExpanderStyle) String() string {
	switch e {
	case ExpanderCollapsed:
		return "Collapsed"
	case ExpanderSemiCollapsed:
		return "SemiCollapsed"
	case ExpanderSemiExpanded:
		return "SemiExpanded"
	case ExpanderExpanded:
		return "Expanded"
	default:
		return fmt.Sprintf("ExpanderStyle(%d)", e)
	}
}

// PaintArrow draws an arrow in the given rectangle on cr using the given
// parameters. arrow_type determines the direction of the arrow.
//
// Deprecated: Use gtk_render_arrow() instead.
func PaintArrow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, arrowType ArrowType, fill bool, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.GtkArrowType(arrowType)
	if fill {
		_arg8 = C.TRUE
	}
	_arg9 = C.gint(x)
	_arg10 = C.gint(y)
	_arg11 = C.gint(width)
	_arg12 = C.gint(height)

	C.gtk_paint_arrow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12)
}

// PaintBox draws a box on cr with the given parameters.
//
// Deprecated: Use gtk_render_frame() and gtk_render_background() instead.
func PaintBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_box(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintBoxGap draws a box in cr using the given style and state and shadow
// type, leaving a gap in one side.
//
// Deprecated: Use gtk_render_frame_gap() instead.
func PaintBoxGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)
	_arg11 = C.GtkPositionType(gapSide)
	_arg12 = C.gint(gapX)
	_arg13 = C.gint(gapWidth)

	C.gtk_paint_box_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, _arg13)
}

// PaintCheck draws a check button indicator in the given rectangle on cr with
// the given parameters.
//
// Deprecated: Use gtk_render_check() instead.
func PaintCheck(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_check(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintDiamond draws a diamond in the given rectangle on window using the given
// parameters.
//
// Deprecated: Use cairo instead.
func PaintDiamond(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_diamond(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintExpander draws an expander as used in TreeView. x and y specify the
// center the expander. The size of the expander is determined by the
// “expander-size” style property of widget. (If widget is not specified or
// doesn’t have an “expander-size” property, an unspecified default size will be
// used, since the caller doesn't have sufficient information to position the
// expander, this is likely not useful.) The expander is expander_size pixels
// tall in the collapsed position and expander_size pixels wide in the expanded
// position.
//
// Deprecated: Use gtk_render_expander() instead.
func PaintExpander(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x int, y int, expanderStyle ExpanderStyle) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.gint(x)
	_arg7 = C.gint(y)
	_arg8 = C.GtkExpanderStyle(expanderStyle)

	C.gtk_paint_expander(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// PaintExtension draws an extension, i.e. a notebook tab.
//
// Deprecated: Use gtk_render_extension() instead.
func PaintExtension(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)
	_arg11 = C.GtkPositionType(gapSide)

	C.gtk_paint_extension(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintFlatBox draws a flat box on cr with the given parameters.
//
// Deprecated: Use gtk_render_frame() and gtk_render_background() instead.
func PaintFlatBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_flat_box(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintFocus draws a focus indicator around the given rectangle on cr using the
// given style.
//
// Deprecated: Use gtk_render_focus() instead.
func PaintFocus(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.gint(x)
	_arg7 = C.gint(y)
	_arg8 = C.gint(width)
	_arg9 = C.gint(height)

	C.gtk_paint_focus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// PaintHandle draws a handle as used in HandleBox and Paned.
//
// Deprecated: Use gtk_render_handle() instead.
func PaintHandle(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, orientation Orientation) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)
	_arg11 = C.GtkOrientation(orientation)

	C.gtk_paint_handle(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintHline draws a horizontal line from (x1, y) to (x2, y) in cr using the
// given style and state.
//
// Deprecated: Use gtk_render_line() instead.
func PaintHline(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x1 int, x2 int, y int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.gint(x1)
	_arg7 = C.gint(x2)
	_arg8 = C.gint(y)

	C.gtk_paint_hline(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// PaintLayout draws a layout on cr using the given parameters.
//
// Deprecated: Use gtk_render_layout() instead.
func PaintLayout(style *Style, cr *cairo.Context, stateType StateType, useText bool, widget Widgetter, detail string, x int, y int, layout *pango.Layout) {
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
	_arg3 = C.GtkStateType(stateType)
	if useText {
		_arg4 = C.TRUE
	}
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_paint_layout(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// PaintOption draws a radio button indicator in the given rectangle on cr with
// the given parameters.
//
// Deprecated: Use gtk_render_option() instead.
func PaintOption(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_option(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintResizeGrip draws a resize grip in the given rectangle on cr using the
// given parameters.
//
// Deprecated: Use gtk_render_handle() instead.
func PaintResizeGrip(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, edge gdk.WindowEdge, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.GdkWindowEdge(edge)
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_resize_grip(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintShadow draws a shadow around the given rectangle in cr using the given
// style and state and shadow type.
//
// Deprecated: Use gtk_render_frame() instead.
func PaintShadow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_shadow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintShadowGap draws a shadow around the given rectangle in cr using the
// given style and state and shadow type, leaving a gap in one side.
//
// Deprecated: Use gtk_render_frame_gap() instead.
func PaintShadowGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)
	_arg11 = C.GtkPositionType(gapSide)
	_arg12 = C.gint(gapX)
	_arg13 = C.gint(gapWidth)

	C.gtk_paint_shadow_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, _arg13)
}

// PaintSlider draws a slider in the given rectangle on cr using the given style
// and orientation.
//
// Deprecated: Use gtk_render_slider() instead.
func PaintSlider(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, orientation Orientation) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)
	_arg11 = C.GtkOrientation(orientation)

	C.gtk_paint_slider(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

// PaintSpinner draws a spinner on window using the given parameters.
//
// Deprecated: Use gtk_render_icon() and the StyleContext you are drawing
// instead.
func PaintSpinner(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, step uint, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.guint(step)
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_spinner(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintTab draws an option menu tab (i.e. the up and down pointing arrows) in
// the given rectangle on cr using the given parameters.
//
// Deprecated: Use cairo instead.
func PaintTab(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.GtkShadowType(shadowType)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = C.gint(x)
	_arg8 = C.gint(y)
	_arg9 = C.gint(width)
	_arg10 = C.gint(height)

	C.gtk_paint_tab(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

// PaintVline draws a vertical line from (x, y1_) to (x, y2_) in cr using the
// given style and state.
//
// Deprecated: Use gtk_render_line() instead.
func PaintVline(style *Style, cr *cairo.Context, stateType StateType, widget Widgetter, detail string, y1 int, y2 int, x int) {
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
	_arg3 = C.GtkStateType(stateType)
	_arg4 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = C.gint(y1)
	_arg7 = C.gint(y2)
	_arg8 = C.gint(x)

	C.gtk_paint_vline(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// StyleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StyleOverrider interface {
	Copy(src *Style)
	DrawArrow(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, arrowType ArrowType, fill bool, x int, y int, width int, height int)
	DrawBox(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawBoxGap(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int)
	DrawCheck(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawDiamond(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawExpander(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x int, y int, expanderStyle ExpanderStyle)
	DrawExtension(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType)
	DrawFlatBox(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawFocus(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawHandle(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, orientation Orientation)
	DrawHline(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, x1 int, x2 int, y int)
	DrawLayout(cr *cairo.Context, stateType StateType, useText bool, widget Widgetter, detail string, x int, y int, layout *pango.Layout)
	DrawOption(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawResizeGrip(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, edge gdk.WindowEdge, x int, y int, width int, height int)
	DrawShadow(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawShadowGap(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int)
	DrawSlider(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int, orientation Orientation)
	DrawSpinner(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, step uint, x int, y int, width int, height int)
	DrawTab(cr *cairo.Context, stateType StateType, shadowType ShadowType, widget Widgetter, detail string, x int, y int, width int, height int)
	DrawVline(cr *cairo.Context, stateType StateType, widget Widgetter, detail string, y1 int, y2 int, x int)
	InitFromRC(rcStyle *RCStyle)
	Realize()
	// RenderIcon renders the icon specified by source at the given size
	// according to the given parameters and returns the result in a pixbuf.
	//
	// Deprecated: Use gtk_render_icon_pixbuf() instead.
	RenderIcon(source *IconSource, direction TextDirection, state StateType, size int, widget Widgetter, detail string) *gdkpixbuf.Pixbuf
	// SetBackground sets the background of window to the background color or
	// pixmap specified by style for the given state.
	//
	// Deprecated: Use gtk_style_context_set_background() instead.
	SetBackground(window gdk.Windower, stateType StateType)
	Unrealize()
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

func wrapStyle(obj *externglib.Object) *Style {
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

	_style = wrapStyle(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _style
}

// ApplyDefaultBackground: deprecated: Use StyleContext instead.
func (style *Style) ApplyDefaultBackground(cr *cairo.Context, window gdk.Windower, stateType StateType, x int, y int, width int, height int) {
	var _arg0 *C.GtkStyle    // out
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.GdkWindow   // out
	var _arg3 C.GtkStateType // out
	var _arg4 C.gint         // out
	var _arg5 C.gint         // out
	var _arg6 C.gint         // out
	var _arg7 C.gint         // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg3 = C.GtkStateType(stateType)
	_arg4 = C.gint(x)
	_arg5 = C.gint(y)
	_arg6 = C.gint(width)
	_arg7 = C.gint(height)

	C.gtk_style_apply_default_background(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
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

	_ret = wrapStyle(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_style_get_style_property(_arg0, _arg1, _arg2, &_arg3)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))

	return _value
}

// HasContext returns whether style has an associated StyleContext.
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

// LookupColor looks up color_name in the style’s logical color mappings,
// filling in color and returning TRUE if found, otherwise returning FALSE. Do
// not cache the found mapping, because it depends on the Style and might change
// when a theme switch occurs.
//
// Deprecated: Use gtk_style_context_lookup_color() instead.
func (style *Style) LookupColor(colorName string) (gdk.Color, bool) {
	var _arg0 *C.GtkStyle // out
	var _arg1 *C.gchar    // out
	var _arg2 C.GdkColor  // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(colorName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_color(_arg0, _arg1, &_arg2)

	var _color gdk.Color // out
	var _ok bool         // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// LookupIconSet looks up stock_id in the icon factories associated with style
// and the default icon factory, returning an icon set if found, otherwise NULL.
//
// Deprecated: Use gtk_style_context_lookup_icon_set() instead.
func (style *Style) LookupIconSet(stockId string) *IconSet {
	var _arg0 *C.GtkStyle   // out
	var _arg1 *C.gchar      // out
	var _cret *C.GtkIconSet // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_lookup_icon_set(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _iconSet
}

// RenderIcon renders the icon specified by source at the given size according
// to the given parameters and returns the result in a pixbuf.
//
// Deprecated: Use gtk_render_icon_pixbuf() instead.
func (style *Style) RenderIcon(source *IconSource, direction TextDirection, state StateType, size int, widget Widgetter, detail string) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkStyle        // out
	var _arg1 *C.GtkIconSource   // out
	var _arg2 C.GtkTextDirection // out
	var _arg3 C.GtkStateType     // out
	var _arg4 C.GtkIconSize      // out
	var _arg5 *C.GtkWidget       // out
	var _arg6 *C.gchar           // out
	var _cret *C.GdkPixbuf       // in

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.GtkIconSource)(gextras.StructNative(unsafe.Pointer(source)))
	_arg2 = C.GtkTextDirection(direction)
	_arg3 = C.GtkStateType(state)
	_arg4 = C.GtkIconSize(size)
	_arg5 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(detail)))
	defer C.free(unsafe.Pointer(_arg6))

	_cret = C.gtk_style_render_icon(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// SetBackground sets the background of window to the background color or pixmap
// specified by style for the given state.
//
// Deprecated: Use gtk_style_context_set_background() instead.
func (style *Style) SetBackground(window gdk.Windower, stateType StateType) {
	var _arg0 *C.GtkStyle    // out
	var _arg1 *C.GdkWindow   // out
	var _arg2 C.GtkStateType // out

	_arg0 = (*C.GtkStyle)(unsafe.Pointer(style.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkStateType(stateType)

	C.gtk_style_set_background(_arg0, _arg1, _arg2)
}

// WidgetGetDefaultStyle returns the default style used by all widgets
// initially.
//
// Deprecated: Use StyleContext instead, and gtk_css_provider_get_default() to
// obtain a StyleProvider with the default widget style information.
func WidgetGetDefaultStyle() *Style {
	var _cret *C.GtkStyle // in

	_cret = C.gtk_widget_get_default_style()

	var _style *Style // out

	_style = wrapStyle(externglib.Take(unsafe.Pointer(_cret)))

	return _style
}
