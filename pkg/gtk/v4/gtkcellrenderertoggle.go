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
		{T: externglib.Type(C.gtk_cell_renderer_toggle_get_type()), F: marshalCellRendererToggle},
	})
}

// CellRendererToggle renders a toggle button in a cell
//
// CellRendererToggle renders a toggle button in a cell. The button is drawn as
// a radio or a checkbutton, depending on the CellRendererToggle:radio property.
// When activated, it emits the CellRendererToggle::toggled signal.
type CellRendererToggle interface {
	CellRenderer

	// Activatable returns whether the cell renderer is activatable. See
	// gtk_cell_renderer_toggle_set_activatable().
	Activatable() bool
	// Active returns whether the cell renderer is active. See
	// gtk_cell_renderer_toggle_set_active().
	Active() bool
	// Radio returns whether we’re rendering radio toggles rather than
	// checkboxes.
	Radio() bool
	// SetActivatable makes the cell renderer activatable.
	SetActivatable(setting bool)
	// SetActive activates or deactivates a cell renderer.
	SetActive(setting bool)
	// SetRadio: if @radio is true, the cell renderer renders a radio toggle
	// (i.e. a toggle in a group of mutually-exclusive toggles). If false, it
	// renders a check toggle (a standalone boolean option). This can be set
	// globally for the cell renderer, or changed just before rendering each
	// cell in the model (for TreeView, you set up a per-row setting using
	// TreeViewColumn to associate model columns with cell renderer properties).
	SetRadio(radio bool)
}

// cellRendererToggle implements the CellRendererToggle class.
type cellRendererToggle struct {
	CellRenderer
}

var _ CellRendererToggle = (*cellRendererToggle)(nil)

// WrapCellRendererToggle wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererToggle(obj *externglib.Object) CellRendererToggle {
	return cellRendererToggle{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererToggle(obj), nil
}

// NewCellRendererToggle constructs a class CellRendererToggle.
func NewCellRendererToggle() CellRendererToggle {
	var _cret C.GtkCellRendererToggle // in

	_cret = C.gtk_cell_renderer_toggle_new()

	var _cellRendererToggle CellRendererToggle // out

	_cellRendererToggle = WrapCellRendererToggle(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererToggle
}

// Activatable returns whether the cell renderer is activatable. See
// gtk_cell_renderer_toggle_set_activatable().
func (t cellRendererToggle) Activatable() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_cell_renderer_toggle_get_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Active returns whether the cell renderer is active. See
// gtk_cell_renderer_toggle_set_active().
func (t cellRendererToggle) Active() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_cell_renderer_toggle_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Radio returns whether we’re rendering radio toggles rather than
// checkboxes.
func (t cellRendererToggle) Radio() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_cell_renderer_toggle_get_radio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable makes the cell renderer activatable.
func (t cellRendererToggle) SetActivatable(setting bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_activatable(_arg0, _arg1)
}

// SetActive activates or deactivates a cell renderer.
func (t cellRendererToggle) SetActive(setting bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_active(_arg0, _arg1)
}

// SetRadio: if @radio is true, the cell renderer renders a radio toggle
// (i.e. a toggle in a group of mutually-exclusive toggles). If false, it
// renders a check toggle (a standalone boolean option). This can be set
// globally for the cell renderer, or changed just before rendering each
// cell in the model (for TreeView, you set up a per-row setting using
// TreeViewColumn to associate model columns with cell renderer properties).
func (t cellRendererToggle) SetRadio(radio bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(t.Native()))
	if radio {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_radio(_arg0, _arg1)
}
