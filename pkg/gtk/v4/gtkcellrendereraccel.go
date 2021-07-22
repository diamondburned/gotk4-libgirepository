// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_accel_mode_get_type()), F: marshalCellRendererAccelMode},
		{T: externglib.Type(C.gtk_cell_renderer_accel_get_type()), F: marshalCellRendererAcceller},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode int

const (
	// CellRendererAccelModeGTK: GTK accelerators mode
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// CellRendererAccelModeOther: other accelerator mode
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CellRendererAccelMode.
func (c CellRendererAccelMode) String() string {
	switch c {
	case CellRendererAccelModeGTK:
		return "GTK"
	case CellRendererAccelModeOther:
		return "Other"
	default:
		return fmt.Sprintf("CellRendererAccelMode(%d)", c)
	}
}

// CellRendererAccel renders a keyboard accelerator in a cell
//
// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like Control + a). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
type CellRendererAccel struct {
	CellRendererText
}

func wrapCellRendererAccel(obj *externglib.Object) *CellRendererAccel {
	return &CellRendererAccel{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererAcceller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererAccel(obj), nil
}

// NewCellRendererAccel creates a new CellRendererAccel.
func NewCellRendererAccel() *CellRendererAccel {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_accel_new()

	var _cellRendererAccel *CellRendererAccel // out

	_cellRendererAccel = wrapCellRendererAccel(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererAccel
}

func (*CellRendererAccel) privateCellRendererAccel() {}
