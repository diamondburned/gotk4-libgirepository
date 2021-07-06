// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_pixbuf_get_type()), F: marshalCellRendererPixbuf},
	})
}

// CellRendererPixbuf can be used to render an image in a cell. It allows to
// render either a given Pixbuf (set via the CellRendererPixbuf:pixbuf property)
// or a named icon (set via the CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is true. If
// the CellRenderer:is-expanded property is true and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is false and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf interface {
	gextras.Objector

	// AsCellRenderer casts the class to the CellRenderer interface.
	AsCellRenderer() CellRenderer

	// GetAlignedArea gets the aligned area used by @cell inside @cell_area.
	// Used for finding the appropriate edit and focus rectangle.
	//
	// This method is inherited from CellRenderer
	GetAlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// GetAlignment fills in @xalign and @yalign with the appropriate values of
	// @cell.
	//
	// This method is inherited from CellRenderer
	GetAlignment() (xalign float32, yalign float32)
	// GetFixedSize fills in @width and @height with the appropriate size of
	// @cell.
	//
	// This method is inherited from CellRenderer
	GetFixedSize() (width int, height int)
	// GetPadding fills in @xpad and @ypad with the appropriate values of @cell.
	//
	// This method is inherited from CellRenderer
	GetPadding() (xpad int, ypad int)
	// GetPreferredHeight retreives a renderer’s natural size when rendered to
	// @widget.
	//
	// This method is inherited from CellRenderer
	GetPreferredHeight(widget Widget) (minimumSize int, naturalSize int)
	// GetPreferredHeightForWidth retreives a cell renderers’s minimum and
	// natural height if it were rendered to @widget with the specified @width.
	//
	// This method is inherited from CellRenderer
	GetPreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int)
	// GetPreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	//
	// This method is inherited from CellRenderer
	GetPreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition)
	// GetPreferredWidth retreives a renderer’s natural size when rendered to
	// @widget.
	//
	// This method is inherited from CellRenderer
	GetPreferredWidth(widget Widget) (minimumSize int, naturalSize int)
	// GetPreferredWidthForHeight retreives a cell renderers’s minimum and
	// natural width if it were rendered to @widget with the specified @height.
	//
	// This method is inherited from CellRenderer
	GetPreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int)
	// GetRequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	//
	// This method is inherited from CellRenderer
	GetRequestMode() SizeRequestMode
	// GetSensitive returns the cell renderer’s sensitivity.
	//
	// This method is inherited from CellRenderer
	GetSensitive() bool
	// GetSize obtains the width and height needed to render the cell. Used by
	// view widgets to determine the appropriate size for the cell_area passed
	// to gtk_cell_renderer_render(). If @cell_area is not nil, fills in the x
	// and y offsets (if set) of the cell relative to this location.
	//
	// Please note that the values set in @width and @height, as well as those
	// in @x_offset and @y_offset are inclusive of the xpad and ypad properties.
	//
	// Deprecated: since version 3.0.
	//
	// This method is inherited from CellRenderer
	GetSize(widget Widget, cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int)
	// GetState translates the cell renderer state to StateFlags, based on the
	// cell renderer and widget sensitivity, and the given CellRendererState.
	//
	// This method is inherited from CellRenderer
	GetState(widget Widget, cellState CellRendererState) StateFlags
	// GetVisible returns the cell renderer’s visibility.
	//
	// This method is inherited from CellRenderer
	GetVisible() bool
	// IsActivatable checks whether the cell renderer can do something when
	// activated.
	//
	// This method is inherited from CellRenderer
	IsActivatable() bool
	// Render invokes the virtual render function of the CellRenderer. The three
	// passed-in rectangles are areas in @cr. Most renderers will draw within
	// @cell_area; the xalign, yalign, xpad, and ypad fields of the CellRenderer
	// should be honored with respect to @cell_area. @background_area includes
	// the blank space around the cell, and also the area containing the tree
	// expander; so the @background_area rectangles for all cells tile to cover
	// the entire @window.
	//
	// This method is inherited from CellRenderer
	Render(cr *cairo.Context, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// SetAlignment sets the renderer’s alignment within its available space.
	//
	// This method is inherited from CellRenderer
	SetAlignment(xalign float32, yalign float32)
	// SetFixedSize sets the renderer size to be explicit, independent of the
	// properties set.
	//
	// This method is inherited from CellRenderer
	SetFixedSize(width int, height int)
	// SetPadding sets the renderer’s padding.
	//
	// This method is inherited from CellRenderer
	SetPadding(xpad int, ypad int)
	// SetSensitive sets the cell renderer’s sensitivity.
	//
	// This method is inherited from CellRenderer
	SetSensitive(sensitive bool)
	// SetVisible sets the cell renderer’s visibility.
	//
	// This method is inherited from CellRenderer
	SetVisible(visible bool)
	// StopEditing informs the cell renderer that the editing is stopped. If
	// @canceled is true, the cell renderer will emit the
	// CellRenderer::editing-canceled signal.
	//
	// This function should be called by cell renderer implementations in
	// response to the CellEditable::editing-done signal of CellEditable.
	//
	// This method is inherited from CellRenderer
	StopEditing(canceled bool)
}

// cellRendererPixbuf implements the CellRendererPixbuf interface.
type cellRendererPixbuf struct {
	*externglib.Object
}

var _ CellRendererPixbuf = (*cellRendererPixbuf)(nil)

// WrapCellRendererPixbuf wraps a GObject to a type that implements
// interface CellRendererPixbuf. It is primarily used internally.
func WrapCellRendererPixbuf(obj *externglib.Object) CellRendererPixbuf {
	return cellRendererPixbuf{obj}
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererPixbuf(obj), nil
}

// NewCellRendererPixbuf creates a new CellRendererPixbuf. Adjust rendering
// parameters using object properties. Object properties can be set globally
// (with g_object_set()). Also, with TreeViewColumn, you can bind a property to
// a value in a TreeModel. For example, you can bind the “pixbuf” property on
// the cell renderer to a pixbuf value in the model, thus rendering a different
// image in each row of the TreeView.
func NewCellRendererPixbuf() CellRendererPixbuf {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_pixbuf_new()

	var _cellRendererPixbuf CellRendererPixbuf // out

	_cellRendererPixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererPixbuf)

	return _cellRendererPixbuf
}

func (c cellRendererPixbuf) AsCellRenderer() CellRenderer {
	return WrapCellRenderer(gextras.InternObject(c))
}

func (c cellRendererPixbuf) GetAlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignedArea(widget, flags, cellArea)
}

func (c cellRendererPixbuf) GetAlignment() (xalign float32, yalign float32) {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignment()
}

func (c cellRendererPixbuf) GetFixedSize() (width int, height int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetFixedSize()
}

func (c cellRendererPixbuf) GetPadding() (xpad int, ypad int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPadding()
}

func (c cellRendererPixbuf) GetPreferredHeight(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeight(widget)
}

func (c cellRendererPixbuf) GetPreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeightForWidth(widget, width)
}

func (c cellRendererPixbuf) GetPreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredSize(widget)
}

func (c cellRendererPixbuf) GetPreferredWidth(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidth(widget)
}

func (c cellRendererPixbuf) GetPreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidthForHeight(widget, height)
}

func (c cellRendererPixbuf) GetRequestMode() SizeRequestMode {
	return WrapCellRenderer(gextras.InternObject(c)).GetRequestMode()
}

func (c cellRendererPixbuf) GetSensitive() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetSensitive()
}

func (c cellRendererPixbuf) GetSize(widget Widget, cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetSize(widget, cellArea)
}

func (c cellRendererPixbuf) GetState(widget Widget, cellState CellRendererState) StateFlags {
	return WrapCellRenderer(gextras.InternObject(c)).GetState(widget, cellState)
}

func (c cellRendererPixbuf) GetVisible() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetVisible()
}

func (c cellRendererPixbuf) IsActivatable() bool {
	return WrapCellRenderer(gextras.InternObject(c)).IsActivatable()
}

func (c cellRendererPixbuf) Render(cr *cairo.Context, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	WrapCellRenderer(gextras.InternObject(c)).Render(cr, widget, backgroundArea, cellArea, flags)
}

func (c cellRendererPixbuf) SetAlignment(xalign float32, yalign float32) {
	WrapCellRenderer(gextras.InternObject(c)).SetAlignment(xalign, yalign)
}

func (c cellRendererPixbuf) SetFixedSize(width int, height int) {
	WrapCellRenderer(gextras.InternObject(c)).SetFixedSize(width, height)
}

func (c cellRendererPixbuf) SetPadding(xpad int, ypad int) {
	WrapCellRenderer(gextras.InternObject(c)).SetPadding(xpad, ypad)
}

func (c cellRendererPixbuf) SetSensitive(sensitive bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetSensitive(sensitive)
}

func (c cellRendererPixbuf) SetVisible(visible bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetVisible(visible)
}

func (c cellRendererPixbuf) StopEditing(canceled bool) {
	WrapCellRenderer(gextras.InternObject(c)).StopEditing(canceled)
}
