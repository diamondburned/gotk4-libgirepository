// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_cell_renderer_accel_mode_get_type()), F: marshalCellRendererAccelMode},
		{T: externglib.Type(C.gtk_cell_renderer_accel_get_type()), F: marshalCellRendererAcceller},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK+
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK+ are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode C.gint

const (
	// CellRendererAccelModeGTK: GTK+ accelerators mode.
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// CellRendererAccelModeOther: other accelerator mode.
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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

// CellRendererAccelOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererAccelOverrider interface {
	AccelCleared(pathString string)
	AccelEdited(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
}

// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like Control + a). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
//
// The CellRendererAccel cell renderer was added in GTK+ 2.10.
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
	return wrapCellRendererAccel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererAccel creates a new CellRendererAccel.
func NewCellRendererAccel() *CellRendererAccel {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_accel_new()

	var _cellRendererAccel *CellRendererAccel // out

	_cellRendererAccel = wrapCellRendererAccel(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererAccel
}

// ConnectAccelCleared gets emitted when the user has removed the accelerator.
func (accel *CellRendererAccel) ConnectAccelCleared(f func(pathString string)) externglib.SignalHandle {
	return accel.Connect("accel-cleared", f)
}

// ConnectAccelEdited gets emitted when the user has selected a new accelerator.
func (accel *CellRendererAccel) ConnectAccelEdited(f func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)) externglib.SignalHandle {
	return accel.Connect("accel-edited", f)
}
