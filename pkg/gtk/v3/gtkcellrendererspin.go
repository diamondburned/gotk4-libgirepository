// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_renderer_spin_get_type()), F: marshalCellRendererSpin},
	})
}

// CellRendererSpin renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererSpin offers a SpinButton widget. Of course, that means that
// the text has to be parseable as a floating point number.
//
// The range of the spinbutton is taken from the adjustment property of the cell
// renderer, which can be set explicitly or mapped to a column in the tree
// model, like all properties of cell renders. CellRendererSpin also has
// properties for the CellRendererSpin:climb-rate and the number of
// CellRendererSpin:digits to display. Other SpinButton properties can be set in
// a handler for the CellRenderer::editing-started signal.
//
// The CellRendererSpin cell renderer was added in GTK+ 2.10.
type CellRendererSpin interface {
	CellRendererText
}

// cellRendererSpin implements the CellRendererSpin class.
type cellRendererSpin struct {
	CellRendererText
}

var _ CellRendererSpin = (*cellRendererSpin)(nil)

// WrapCellRendererSpin wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererSpin(obj *externglib.Object) CellRendererSpin {
	return cellRendererSpin{
		CellRendererText: WrapCellRendererText(obj),
	}
}

func marshalCellRendererSpin(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererSpin(obj), nil
}

// NewCellRendererSpin constructs a class CellRendererSpin.
func NewCellRendererSpin() CellRendererSpin {
	var _cret C.GtkCellRendererSpin // in

	_cret = C.gtk_cell_renderer_spin_new()

	var _cellRendererSpin CellRendererSpin // out

	_cellRendererSpin = WrapCellRendererSpin(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererSpin
}
