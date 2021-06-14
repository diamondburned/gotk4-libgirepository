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
		{T: externglib.Type(C.gtk_cell_editable_get_type()), F: marshalCellEditable},
	})
}

// CellEditableOverrider contains methods that are overridable. This
// interface is a subset of the interface CellEditable.
type CellEditableOverrider interface {
	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
}

// CellEditable: the CellEditable interface must be implemented for widgets to
// be usable to edit the contents of a TreeView cell. It provides a way to
// specify how temporary widgets should be configured for editing, get the new
// value, etc.
type CellEditable interface {
	Widget
	CellEditableOverrider
}

// cellEditable implements the CellEditable interface.
type cellEditable struct {
	Widget
}

var _ CellEditable = (*cellEditable)(nil)

// WrapCellEditable wraps a GObject to a type that implements interface
// CellEditable. It is primarily used internally.
func WrapCellEditable(obj *externglib.Object) CellEditable {
	return CellEditable{
		Widget: WrapWidget(obj),
	}
}

func marshalCellEditable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellEditable(obj), nil
}

// EditingDone emits the CellEditable::editing-done signal.
func (c cellEditable) EditingDone() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(c.Native()))

	C.gtk_cell_editable_editing_done(_arg0)
}

// RemoveWidget emits the CellEditable::remove-widget signal.
func (c cellEditable) RemoveWidget() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(c.Native()))

	C.gtk_cell_editable_remove_widget(_arg0)
}
