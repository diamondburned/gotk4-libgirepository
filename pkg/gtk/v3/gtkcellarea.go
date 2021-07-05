// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
//
// gboolean gotk4_CellAllocCallback(GtkCellRenderer*, GdkRectangle*, GdkRectangle*, gpointer);
// gboolean gotk4_CellCallback(GtkCellRenderer*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_area_get_type()), F: marshalCellArea},
	})
}

// CellAllocCallback: the type of the callback functions used for iterating over
// the cell renderers and their allocated areas inside a CellArea, see
// gtk_cell_area_foreach_alloc().
type CellAllocCallback func(renderer CellRenderer, cellArea *gdk.Rectangle, cellBackground *gdk.Rectangle) (ok bool)

//export gotk4_CellAllocCallback
func gotk4_CellAllocCallback(arg0 *C.GtkCellRenderer, arg1 *C.GdkRectangle, arg2 *C.GdkRectangle, arg3 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var renderer CellRenderer         // out
	var cellArea *gdk.Rectangle       // out
	var cellBackground *gdk.Rectangle // out

	renderer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(CellRenderer)
	cellArea = (*gdk.Rectangle)(unsafe.Pointer(arg1))
	cellBackground = (*gdk.Rectangle)(unsafe.Pointer(arg2))

	fn := v.(CellAllocCallback)
	ok := fn(renderer, cellArea, cellBackground)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CellCallback: the type of the callback functions used for iterating over the
// cell renderers of a CellArea, see gtk_cell_area_foreach().
type CellCallback func(renderer CellRenderer) (ok bool)

//export gotk4_CellCallback
func gotk4_CellCallback(arg0 *C.GtkCellRenderer, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var renderer CellRenderer // out

	renderer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(CellRenderer)

	fn := v.(CellCallback)
	ok := fn(renderer)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// CellArea: the CellArea is an abstract class for CellLayout widgets (also
// referred to as "layouting widgets") to interface with an arbitrary number of
// CellRenderers and interact with the user for a given TreeModel row.
//
// The cell area handles events, focus navigation, drawing and size requests and
// allocations for a given row of data.
//
// Usually users dont have to interact with the CellArea directly unless they
// are implementing a cell-layouting widget themselves.
//
//
// Requesting area sizes
//
// As outlined in [GtkWidget’s geometry management
// section][geometry-management], GTK+ uses a height-for-width geometry
// management system to compute the sizes of widgets and user interfaces.
// CellArea uses the same semantics to calculate the size of an area for an
// arbitrary number of TreeModel rows.
//
// When requesting the size of a cell area one needs to calculate the size for a
// handful of rows, and this will be done differently by different layouting
// widgets. For instance a TreeViewColumn always lines up the areas from top to
// bottom while a IconView on the other hand might enforce that all areas
// received the same width and wrap the areas around, requesting height for more
// cell areas when allocated less width.
//
// It’s also important for areas to maintain some cell alignments with areas
// rendered for adjacent rows (cells can appear “columnized” inside an area even
// when the size of cells are different in each row). For this reason the
// CellArea uses a CellAreaContext object to store the alignments and sizes
// along the way (as well as the overall largest minimum and natural size for
// all the rows which have been calculated with the said context).
//
// The CellAreaContext is an opaque object specific to the CellArea which
// created it (see gtk_cell_area_create_context()). The owning cell-layouting
// widget can create as many contexts as it wishes to calculate sizes of rows
// which should receive the same size in at least one orientation (horizontally
// or vertically), However, it’s important that the same CellAreaContext which
// was used to request the sizes for a given TreeModel row be used when
// rendering or processing events for that row.
//
// In order to request the width of all the rows at the root level of a
// TreeModel one would do the following:
//
//    static gboolean
//    foo_focus (GtkWidget       *widget,
//               GtkDirectionType direction)
//    {
//      Foo        *foo  = FOO (widget);
//      FooPrivate *priv = foo->priv;
//      gint        focus_row;
//      gboolean    have_focus = FALSE;
//
//      focus_row = priv->focus_row;
//
//      if (!gtk_widget_has_focus (widget))
//        gtk_widget_grab_focus (widget);
//
//      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, priv->focus_row);
//      while (valid)
//        {
//          gtk_cell_area_apply_attributes (priv->area, priv->model, &iter, FALSE, FALSE);
//
//          if (gtk_cell_area_focus (priv->area, direction))
//            {
//               priv->focus_row = focus_row;
//               have_focus = TRUE;
//               break;
//            }
//          else
//            {
//              if (direction == GTK_DIR_RIGHT ||
//                  direction == GTK_DIR_LEFT)
//                break;
//              else if (direction == GTK_DIR_UP ||
//                       direction == GTK_DIR_TAB_BACKWARD)
//               {
//                  if (focus_row == 0)
//                    break;
//                  else
//                   {
//                      focus_row--;
//                      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, focus_row);
//                   }
//                }
//              else
//                {
//                  if (focus_row == last_row)
//                    break;
//                  else
//                    {
//                      focus_row++;
//                      valid = gtk_tree_model_iter_next (priv->model, &iter);
//                    }
//                }
//            }
//        }
//        return have_focus;
//    }
//
// Note that the layouting widget is responsible for matching the
// GtkDirectionType values to the way it lays out its cells.
//
//
// Cell Properties
//
// The CellArea introduces cell properties for CellRenderers in very much the
// same way that Container introduces [child properties][child-properties] for
// Widgets. This provides some general interfaces for defining the relationship
// cell areas have with their cells. For instance in a CellAreaBox a cell might
// “expand” and receive extra space when the area is allocated more than its
// full natural request, or a cell might be configured to “align” with adjacent
// rows which were requested and rendered with the same CellAreaContext.
//
// Use gtk_cell_area_class_install_cell_property() to install cell properties
// for a cell area class and gtk_cell_area_class_find_cell_property() or
// gtk_cell_area_class_list_cell_properties() to get information about existing
// cell properties.
//
// To set the value of a cell property, use gtk_cell_area_cell_set_property(),
// gtk_cell_area_cell_set() or gtk_cell_area_cell_set_valist(). To obtain the
// value of a cell property, use gtk_cell_area_cell_get_property(),
// gtk_cell_area_cell_get() or gtk_cell_area_cell_get_valist().
type CellArea interface {
	gextras.Objector

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellLayout casts the class to the CellLayout interface.
	AsCellLayout() CellLayout

	// ActivateCellArea activates @area, usually by activating the currently
	// focused cell, however some subclasses which embed widgets in the area can
	// also activate a widget if it currently has the focus.
	ActivateCellArea(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool
	// AddCellArea adds @renderer to @area with the default child cell
	// properties.
	AddCellArea(renderer CellRenderer)
	// AddFocusSiblingCellArea adds @sibling to @renderer’s focusable area,
	// focus will be drawn around @renderer and all of its siblings if @renderer
	// can focus for a given row.
	//
	// Events handled by focus siblings can also activate the given focusable
	// @renderer.
	AddFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer)
	// ApplyAttributesCellArea applies any connected attributes to the renderers
	// in @area by pulling the values from @tree_model.
	ApplyAttributesCellArea(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)
	// AttributeConnectCellArea connects an @attribute to apply values from
	// @column for the TreeModel in use.
	AttributeConnectCellArea(renderer CellRenderer, attribute string, column int)
	// AttributeDisconnectCellArea disconnects @attribute for the @renderer in
	// @area so that attribute will no longer be updated with values from the
	// model.
	AttributeDisconnectCellArea(renderer CellRenderer, attribute string)
	// AttributeGetColumnCellArea returns the model column that an attribute has
	// been mapped to, or -1 if the attribute is not mapped.
	AttributeGetColumnCellArea(renderer CellRenderer, attribute string) int
	// CellGetPropertyCellArea gets the value of a cell property for @renderer
	// in @area.
	CellGetPropertyCellArea(renderer CellRenderer, propertyName string, value externglib.Value)
	// CellSetPropertyCellArea sets a cell property for @renderer in @area.
	CellSetPropertyCellArea(renderer CellRenderer, propertyName string, value externglib.Value)
	// CopyContextCellArea: this is sometimes needed for cases where rows need
	// to share alignments in one orientation but may be separately grouped in
	// the opposing orientation.
	//
	// For instance, IconView creates all icons (rows) to have the same width
	// and the cells theirin to have the same horizontal alignments. However
	// each row of icons may have a separate collective height. IconView uses
	// this to request the heights of each row based on a context which was
	// already used to request all the row widths that are to be displayed.
	CopyContextCellArea(context CellAreaContext) CellAreaContext
	// CreateContextCellArea creates a CellAreaContext to be used with @area for
	// all purposes. CellAreaContext stores geometry information for rows for
	// which it was operated on, it is important to use the same context for the
	// same row of data at all times (i.e. one should render and handle events
	// with the same CellAreaContext which was used to request the size of those
	// rows of data).
	CreateContextCellArea() CellAreaContext
	// FocusCellArea: this should be called by the @area’s owning layout widget
	// when focus is to be passed to @area, or moved within @area for a given
	// @direction and row data.
	//
	// Implementing CellArea classes should implement this method to receive and
	// navigate focus in its own way particular to how it lays out cells.
	FocusCellArea(direction DirectionType) bool
	// ForeachCellArea calls @callback for every CellRenderer in @area.
	ForeachCellArea(callback CellCallback)
	// ForeachAllocCellArea calls @callback for every CellRenderer in @area with
	// the allocated rectangle inside @cell_area.
	ForeachAllocCellArea(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, backgroundArea *gdk.Rectangle, callback CellAllocCallback)
	// CellAllocation derives the allocation of @renderer inside @area if @area
	// were to be renderered in @cell_area.
	CellAllocation(context CellAreaContext, widget Widget, renderer CellRenderer, cellArea *gdk.Rectangle) gdk.Rectangle
	// CellAtPosition gets the CellRenderer at @x and @y coordinates inside
	// @area and optionally returns the full cell allocation for it inside
	// @cell_area.
	CellAtPosition(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, x int, y int) (gdk.Rectangle, CellRenderer)
	// CurrentPathString gets the current TreePath string for the currently
	// applied TreeIter, this is implicitly updated when
	// gtk_cell_area_apply_attributes() is called and can be used to interact
	// with renderers from CellArea subclasses.
	CurrentPathString() string
	// EditWidget gets the CellEditable widget currently used to edit the
	// currently edited cell.
	EditWidget() CellEditable
	// EditedCell gets the CellRenderer in @area that is currently being edited.
	EditedCell() CellRenderer
	// FocusCell retrieves the currently focused cell for @area
	FocusCell() CellRenderer
	// FocusFromSibling gets the CellRenderer which is expected to be focusable
	// for which @renderer is, or may be a sibling.
	//
	// This is handy for CellArea subclasses when handling events, after
	// determining the renderer at the event location it can then chose to
	// activate the focus cell for which the event cell may have been a sibling.
	FocusFromSibling(renderer CellRenderer) CellRenderer
	// PreferredHeight retrieves a cell area’s initial minimum and natural
	// height.
	//
	// @area will store some geometrical information in @context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the @minimum_height and @natural_height of this call
	// but rather to consult gtk_cell_area_context_get_preferred_height() after
	// a series of requests.
	PreferredHeight(context CellAreaContext, widget Widget) (minimumHeight int, naturalHeight int)
	// PreferredHeightForWidth retrieves a cell area’s minimum and natural
	// height if it would be given the specified @width.
	//
	// @area stores some geometrical information in @context along the way while
	// calling gtk_cell_area_get_preferred_width(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_width() requests with @context
	// first and then call gtk_cell_area_get_preferred_height_for_width() on
	// each cell area individually to get the height for width of each fully
	// requested row.
	//
	// If at some point, the width of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_width() again and then the
	// full width of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_width().
	PreferredHeightForWidth(context CellAreaContext, widget Widget, width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth retrieves a cell area’s initial minimum and natural width.
	//
	// @area will store some geometrical information in @context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the @minimum_width and @natural_width of this call but
	// rather to consult gtk_cell_area_context_get_preferred_width() after a
	// series of requests.
	PreferredWidth(context CellAreaContext, widget Widget) (minimumWidth int, naturalWidth int)
	// PreferredWidthForHeight retrieves a cell area’s minimum and natural width
	// if it would be given the specified @height.
	//
	// @area stores some geometrical information in @context along the way while
	// calling gtk_cell_area_get_preferred_height(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_height() requests with @context
	// first and then call gtk_cell_area_get_preferred_width_for_height() on
	// each cell area individually to get the height for width of each fully
	// requested row.
	//
	// If at some point, the height of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_height() again and then the
	// full height of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_height().
	PreferredWidthForHeight(context CellAreaContext, widget Widget, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the area prefers a height-for-width layout or a
	// width-for-height layout.
	RequestMode() SizeRequestMode
	// HasRendererCellArea checks if @area contains @renderer.
	HasRendererCellArea(renderer CellRenderer) bool
	// InnerCellAreaCellArea: this is a convenience function for CellArea
	// implementations to get the inner area where a given CellRenderer will be
	// rendered. It removes any padding previously added by
	// gtk_cell_area_request_renderer().
	InnerCellAreaCellArea(widget Widget, cellArea *gdk.Rectangle) gdk.Rectangle
	// IsActivatableCellArea returns whether the area can do anything when
	// activated, after applying new attributes to @area.
	IsActivatableCellArea() bool
	// IsFocusSiblingCellArea returns whether @sibling is one of @renderer’s
	// focus siblings (see gtk_cell_area_add_focus_sibling()).
	IsFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer) bool
	// RemoveCellArea removes @renderer from @area.
	RemoveCellArea(renderer CellRenderer)
	// RemoveFocusSiblingCellArea removes @sibling from @renderer’s focus
	// sibling list (see gtk_cell_area_add_focus_sibling()).
	RemoveFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer)
	// RenderCellArea renders @area’s cells according to @area’s layout onto
	// @widget at the given coordinates.
	RenderCellArea(context CellAreaContext, widget Widget, cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool)
	// RequestRendererCellArea: this is a convenience function for CellArea
	// implementations to request size for cell renderers. It’s important to use
	// this function to request size and then use
	// gtk_cell_area_inner_cell_area() at render and event time since this
	// function will add padding around the cell for focus painting.
	RequestRendererCellArea(renderer CellRenderer, orientation Orientation, widget Widget, forSize int) (minimumSize int, naturalSize int)
	// SetFocusCellCellArea: explicitly sets the currently focused cell to
	// @renderer.
	//
	// This is generally called by implementations of CellAreaClass.focus() or
	// CellAreaClass.event(), however it can also be used to implement functions
	// such as gtk_tree_view_set_cursor_on_cell().
	SetFocusCellCellArea(renderer CellRenderer)
	// StopEditingCellArea: explicitly stops the editing of the currently edited
	// cell.
	//
	// If @canceled is true, the currently edited cell renderer will emit the
	// ::editing-canceled signal, otherwise the the ::editing-done signal will
	// be emitted on the current edit widget.
	//
	// See gtk_cell_area_get_edited_cell() and gtk_cell_area_get_edit_widget().
	StopEditingCellArea(canceled bool)
}

// cellArea implements the CellArea class.
type cellArea struct {
	gextras.Objector
}

// WrapCellArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellArea(obj *externglib.Object) CellArea {
	return cellArea{
		Objector: obj,
	}
}

func marshalCellArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellArea(obj), nil
}

func (a cellArea) ActivateCellArea(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 C.GtkCellRendererState // out
	var _arg5 C.gboolean             // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))
	_arg4 = C.GtkCellRendererState(flags)
	if editOnly {
		_arg5 = C.TRUE
	}

	_cret = C.gtk_cell_area_activate(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a cellArea) AddCellArea(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_add(_arg0, _arg1)
}

func (a cellArea) AddFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	C.gtk_cell_area_add_focus_sibling(_arg0, _arg1, _arg2)
}

func (a cellArea) ApplyAttributesCellArea(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	var _arg0 *C.GtkCellArea  // out
	var _arg1 *C.GtkTreeModel // out
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.gboolean      // out
	var _arg4 C.gboolean      // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	if isExpander {
		_arg3 = C.TRUE
	}
	if isExpanded {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_apply_attributes(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (a cellArea) AttributeConnectCellArea(renderer CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out
	var _arg3 C.gint             // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(column)

	C.gtk_cell_area_attribute_connect(_arg0, _arg1, _arg2, _arg3)
}

func (a cellArea) AttributeDisconnectCellArea(renderer CellRenderer, attribute string) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_cell_area_attribute_disconnect(_arg0, _arg1, _arg2)
}

func (a cellArea) AttributeGetColumnCellArea(renderer CellRenderer, attribute string) int {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_cell_area_attribute_get_column(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a cellArea) CellGetPropertyCellArea(renderer CellRenderer, propertyName string, value externglib.Value) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.GValue          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_cell_area_cell_get_property(_arg0, _arg1, _arg2, _arg3)
}

func (a cellArea) CellSetPropertyCellArea(renderer CellRenderer, propertyName string, value externglib.Value) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.GValue          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_cell_area_cell_set_property(_arg0, _arg1, _arg2, _arg3)
}

func (a cellArea) CopyContextCellArea(context CellAreaContext) CellAreaContext {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _cret *C.GtkCellAreaContext // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_cell_area_copy_context(_arg0, _arg1)

	var _cellAreaContext CellAreaContext // out

	_cellAreaContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(CellAreaContext)

	return _cellAreaContext
}

func (a cellArea) CreateContextCellArea() CellAreaContext {
	var _arg0 *C.GtkCellArea        // out
	var _cret *C.GtkCellAreaContext // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_create_context(_arg0)

	var _cellAreaContext CellAreaContext // out

	_cellAreaContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(CellAreaContext)

	return _cellAreaContext
}

func (a cellArea) FocusCellArea(direction DirectionType) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 C.GtkDirectionType // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkDirectionType(direction)

	_cret = C.gtk_cell_area_focus(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a cellArea) ForeachCellArea(callback CellCallback) {
	var _arg0 *C.GtkCellArea    // out
	var _arg1 C.GtkCellCallback // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*[0]byte)(C.gotk4_CellCallback)
	_arg2 = C.gpointer(box.Assign(callback))

	C.gtk_cell_area_foreach(_arg0, _arg1, _arg2)
}

func (a cellArea) ForeachAllocCellArea(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, backgroundArea *gdk.Rectangle, callback CellAllocCallback) {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellAllocCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea))
	_arg5 = (*[0]byte)(C.gotk4_CellAllocCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.gtk_cell_area_foreach_alloc(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (a cellArea) CellAllocation(context CellAreaContext, widget Widget, renderer CellRenderer, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 *C.GtkCellRenderer    // out
	var _arg4 *C.GdkRectangle       // out
	var _arg5 C.GdkRectangle        // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	C.gtk_cell_area_get_cell_allocation(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5)

	var _allocation gdk.Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg5
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_allocation = *refTmpOut
	}

	return _allocation
}

func (a cellArea) CellAtPosition(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, x int, y int) (gdk.Rectangle, CellRenderer) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 *C.GdkRectangle       // out
	var _arg4 C.gint                // out
	var _arg5 C.gint                // out
	var _arg6 C.GdkRectangle        // in
	var _cret *C.GtkCellRenderer    // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))
	_arg4 = C.gint(x)
	_arg5 = C.gint(y)

	_cret = C.gtk_cell_area_get_cell_at_position(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_arg6)

	var _allocArea gdk.Rectangle   // out
	var _cellRenderer CellRenderer // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg6
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_allocArea = *refTmpOut
	}
	_cellRenderer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRenderer)

	return _allocArea, _cellRenderer
}

func (a cellArea) CurrentPathString() string {
	var _arg0 *C.GtkCellArea // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_get_current_path_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a cellArea) EditWidget() CellEditable {
	var _arg0 *C.GtkCellArea     // out
	var _cret *C.GtkCellEditable // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_get_edit_widget(_arg0)

	var _cellEditable CellEditable // out

	_cellEditable = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellEditable)

	return _cellEditable
}

func (a cellArea) EditedCell() CellRenderer {
	var _arg0 *C.GtkCellArea     // out
	var _cret *C.GtkCellRenderer // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_get_edited_cell(_arg0)

	var _cellRenderer CellRenderer // out

	_cellRenderer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRenderer)

	return _cellRenderer
}

func (a cellArea) FocusCell() CellRenderer {
	var _arg0 *C.GtkCellArea     // out
	var _cret *C.GtkCellRenderer // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_get_focus_cell(_arg0)

	var _cellRenderer CellRenderer // out

	_cellRenderer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRenderer)

	return _cellRenderer
}

func (a cellArea) FocusFromSibling(renderer CellRenderer) CellRenderer {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _cret *C.GtkCellRenderer // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gtk_cell_area_get_focus_from_sibling(_arg0, _arg1)

	var _cellRenderer CellRenderer // out

	_cellRenderer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRenderer)

	return _cellRenderer
}

func (a cellArea) PreferredHeight(context CellAreaContext, widget Widget) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.gint                // in
	var _arg4 C.gint                // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_area_get_preferred_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg3)
	_naturalHeight = int(_arg4)

	return _minimumHeight, _naturalHeight
}

func (a cellArea) PreferredHeightForWidth(context CellAreaContext, widget Widget, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.gint                // out
	var _arg4 C.gint                // in
	var _arg5 C.gint                // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.gint(width)

	C.gtk_cell_area_get_preferred_height_for_width(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg4)
	_naturalHeight = int(_arg5)

	return _minimumHeight, _naturalHeight
}

func (a cellArea) PreferredWidth(context CellAreaContext, widget Widget) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.gint                // in
	var _arg4 C.gint                // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_area_get_preferred_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg3)
	_naturalWidth = int(_arg4)

	return _minimumWidth, _naturalWidth
}

func (a cellArea) PreferredWidthForHeight(context CellAreaContext, widget Widget, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.gint                // out
	var _arg4 C.gint                // in
	var _arg5 C.gint                // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.gint(height)

	C.gtk_cell_area_get_preferred_width_for_height(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg4)
	_naturalWidth = int(_arg5)

	return _minimumWidth, _naturalWidth
}

func (a cellArea) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkCellArea       // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_get_request_mode(_arg0)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

func (a cellArea) HasRendererCellArea(renderer CellRenderer) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gtk_cell_area_has_renderer(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a cellArea) InnerCellAreaCellArea(widget Widget, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellArea  // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkRectangle // out
	var _arg3 C.GdkRectangle  // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	C.gtk_cell_area_inner_cell_area(_arg0, _arg1, _arg2, &_arg3)

	var _innerArea gdk.Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_innerArea = *refTmpOut
	}

	return _innerArea
}

func (a cellArea) IsActivatableCellArea() bool {
	var _arg0 *C.GtkCellArea // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_cell_area_is_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a cellArea) IsFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	_cret = C.gtk_cell_area_is_focus_sibling(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a cellArea) RemoveCellArea(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_remove(_arg0, _arg1)
}

func (a cellArea) RemoveFocusSiblingCellArea(renderer CellRenderer, sibling CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	C.gtk_cell_area_remove_focus_sibling(_arg0, _arg1, _arg2)
}

func (a cellArea) RenderCellArea(context CellAreaContext, widget Widget, cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool) {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.cairo_t             // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _arg7 C.gboolean             // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea))
	_arg5 = (*C.GdkRectangle)(unsafe.Pointer(cellArea))
	_arg6 = C.GtkCellRendererState(flags)
	if paintFocus {
		_arg7 = C.TRUE
	}

	C.gtk_cell_area_render(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

func (a cellArea) RequestRendererCellArea(renderer CellRenderer, orientation Orientation, widget Widget, forSize int) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.GtkOrientation   // out
	var _arg3 *C.GtkWidget       // out
	var _arg4 C.gint             // out
	var _arg5 C.gint             // in
	var _arg6 C.gint             // in

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = C.GtkOrientation(orientation)
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg4 = C.gint(forSize)

	C.gtk_cell_area_request_renderer(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg5)
	_naturalSize = int(_arg6)

	return _minimumSize, _naturalSize
}

func (a cellArea) SetFocusCellCellArea(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_set_focus_cell(_arg0, _arg1)
}

func (a cellArea) StopEditingCellArea(canceled bool) {
	var _arg0 *C.GtkCellArea // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	if canceled {
		_arg1 = C.TRUE
	}

	C.gtk_cell_area_stop_editing(_arg0, _arg1)
}

func (c cellArea) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(c))
}

func (c cellArea) AsCellLayout() CellLayout {
	return WrapCellLayout(gextras.InternObject(c))
}
