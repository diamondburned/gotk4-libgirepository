// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_accel_get_type()), F: marshalCellRendererAccel},
	})
}

// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like `Control + a`). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
//
// The CellRendererAccel cell renderer was added in GTK+ 2.10.
type CellRendererAccel interface {
	CellRendererText
}

// cellRendererAccel implements the CellRendererAccel interface.
type cellRendererAccel struct {
	CellRendererText
}

var _ CellRendererAccel = (*cellRendererAccel)(nil)

// WrapCellRendererAccel wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererAccel(obj *externglib.Object) CellRendererAccel {
	return CellRendererAccel{
		CellRendererText: WrapCellRendererText(obj),
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererAccel(obj), nil
}

// NewCellRendererAccel constructs a class CellRendererAccel.
func NewCellRendererAccel() CellRendererAccel {
	var cret C.GtkCellRendererAccel
	var goret CellRendererAccel

	cret = C.gtk_cell_renderer_accel_new()

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CellRendererAccel)

	return goret
}
