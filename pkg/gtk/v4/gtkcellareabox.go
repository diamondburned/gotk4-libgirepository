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
		{T: externglib.Type(C.gtk_cell_area_box_get_type()), F: marshalCellAreaBox},
	})
}

// CellAreaBox: a cell area that renders GtkCellRenderers into a row or a column
//
// The CellAreaBox renders cell renderers into a row or a column depending on
// its Orientation.
//
// GtkCellAreaBox uses a notion of packing. Packing refers to adding cell
// renderers with reference to a particular position in a CellAreaBox. There are
// two reference positions: the start and the end of the box. When the
// CellAreaBox is oriented in the GTK_ORIENTATION_VERTICAL orientation, the
// start is defined as the top of the box and the end is defined as the bottom.
// In the GTK_ORIENTATION_HORIZONTAL orientation start is defined as the left
// side and the end is defined as the right side.
//
// Alignments of CellRenderers rendered in adjacent rows can be configured by
// configuring the CellAreaBox align child cell property with
// gtk_cell_area_cell_set_property() or by specifying the "align" argument to
// gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox interface {
	CellArea
	Buildable
	CellLayout
	Orientable

	// Spacing gets the spacing added between cell renderers.
	Spacing() int
	// PackEnd adds @renderer to @box, packed with reference to the end of @box.
	//
	// The @renderer is packed after (away from end of) any other CellRenderer
	// packed with reference to the end of @box.
	PackEnd(renderer CellRenderer, expand bool, align bool, fixed bool)
	// PackStart adds @renderer to @box, packed with reference to the start of
	// @box.
	//
	// The @renderer is packed after any other CellRenderer packed with
	// reference to the start of @box.
	PackStart(renderer CellRenderer, expand bool, align bool, fixed bool)
	// SetSpacing sets the spacing to add between cell renderers in @box.
	SetSpacing(spacing int)
}

// cellAreaBox implements the CellAreaBox class.
type cellAreaBox struct {
	CellArea
	Buildable
	CellLayout
	Orientable
}

var _ CellAreaBox = (*cellAreaBox)(nil)

// WrapCellAreaBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAreaBox(obj *externglib.Object) CellAreaBox {
	return cellAreaBox{
		CellArea:   WrapCellArea(obj),
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAreaBox(obj), nil
}

// Spacing gets the spacing added between cell renderers.
func (b cellAreaBox) Spacing() int {
	var _arg0 *C.GtkCellAreaBox // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))

	var _cret C.int // in

	_cret = C.gtk_cell_area_box_get_spacing(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// PackEnd adds @renderer to @box, packed with reference to the end of @box.
//
// The @renderer is packed after (away from end of) any other CellRenderer
// packed with reference to the end of @box.
func (b cellAreaBox) PackEnd(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// PackStart adds @renderer to @box, packed with reference to the start of
// @box.
//
// The @renderer is packed after any other CellRenderer packed with
// reference to the start of @box.
func (b cellAreaBox) PackStart(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetSpacing sets the spacing to add between cell renderers in @box.
func (b cellAreaBox) SetSpacing(spacing int) {
	var _arg0 *C.GtkCellAreaBox // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(spacing)

	C.gtk_cell_area_box_set_spacing(_arg0, _arg1)
}
