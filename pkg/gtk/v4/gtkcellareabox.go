// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkcellareabox.go.
var GTypeCellAreaBox = coreglib.Type(C.gtk_cell_area_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellAreaBox, F: marshalCellAreaBox},
	})
}

// CellAreaBox: cell area that renders GtkCellRenderers into a row or a column
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
type CellAreaBox struct {
	_ [0]func() // equal guard
	CellArea

	*coreglib.Object
	Orientable
}

var (
	_ CellAreaer        = (*CellAreaBox)(nil)
	_ coreglib.Objector = (*CellAreaBox)(nil)
)

func wrapCellAreaBox(obj *coreglib.Object) *CellAreaBox {
	return &CellAreaBox{
		CellArea: CellArea{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
			CellLayout: CellLayout{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	return wrapCellAreaBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellAreaBox creates a new CellAreaBox.
//
// The function returns the following values:
//
//    - cellAreaBox: newly created CellAreaBox.
//
func NewCellAreaBox() *CellAreaBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "CellAreaBox").InvokeMethod("new_CellAreaBox", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _cellAreaBox *CellAreaBox // out

	_cellAreaBox = wrapCellAreaBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellAreaBox
}

// Spacing gets the spacing added between cell renderers.
//
// The function returns the following values:
//
//    - gint: space added between cell renderers in box.
//
func (box *CellAreaBox) Spacing() int32 {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.int   // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellAreaBox").InvokeMethod("get_spacing", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(box)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// PackEnd adds renderer to box, packed with reference to the end of box.
//
// The renderer is packed after (away from end of) any other CellRenderer packed
// with reference to the end of box.
//
// The function takes the following parameters:
//
//    - renderer to add.
//    - expand: whether renderer should receive extra space when the area
//      receives more than its natural size.
//    - align: whether renderer should be aligned in adjacent rows.
//    - fixed: whether renderer should have the same size in all rows.
//
func (box *CellAreaBox) PackEnd(renderer CellRendererer, expand, align, fixed bool) {
	var _args [5]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out
	var _arg3 C.gboolean // out
	var _arg4 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gboolean)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.gboolean)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.gboolean)(unsafe.Pointer(&_args[4])) = _arg4

	girepository.MustFind("Gtk", "CellAreaBox").InvokeMethod("pack_end", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(align)
	runtime.KeepAlive(fixed)
}

// PackStart adds renderer to box, packed with reference to the start of box.
//
// The renderer is packed after any other CellRenderer packed with reference to
// the start of box.
//
// The function takes the following parameters:
//
//    - renderer to add.
//    - expand: whether renderer should receive extra space when the area
//      receives more than its natural size.
//    - align: whether renderer should be aligned in adjacent rows.
//    - fixed: whether renderer should have the same size in all rows.
//
func (box *CellAreaBox) PackStart(renderer CellRendererer, expand, align, fixed bool) {
	var _args [5]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out
	var _arg3 C.gboolean // out
	var _arg4 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gboolean)(unsafe.Pointer(&_args[2])) = _arg2
	*(*C.gboolean)(unsafe.Pointer(&_args[3])) = _arg3
	*(*C.gboolean)(unsafe.Pointer(&_args[4])) = _arg4

	girepository.MustFind("Gtk", "CellAreaBox").InvokeMethod("pack_start", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(align)
	runtime.KeepAlive(fixed)
}

// SetSpacing sets the spacing to add between cell renderers in box.
//
// The function takes the following parameters:
//
//    - spacing: space to add between CellRenderers.
//
func (box *CellAreaBox) SetSpacing(spacing int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.int   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = C.int(spacing)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.int)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "CellAreaBox").InvokeMethod("set_spacing", _args[:], nil)

	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
