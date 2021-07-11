// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_cell_accessible_parent_get_type()), F: marshalCellAccessibleParenter},
	})
}

// CellAccessibleParentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellAccessibleParentOverrider interface {
	//
	Activate(cell CellAccessibler)
	//
	Edit(cell CellAccessibler)
	//
	ExpandCollapse(cell CellAccessibler)
	//
	CellArea(cell CellAccessibler) gdk.Rectangle
	//
	CellPosition(cell CellAccessibler) (row int, column int)
	//
	ChildIndex(cell CellAccessibler) int
	//
	RendererState(cell CellAccessibler) CellRendererState
	//
	GrabFocus(cell CellAccessibler) bool
	//
	UpdateRelationset(cell CellAccessibler, relationset atk.RelationSetter)
}

// CellAccessibleParenter describes CellAccessibleParent's methods.
type CellAccessibleParenter interface {
	//
	Activate(cell CellAccessibler)
	//
	Edit(cell CellAccessibler)
	//
	ExpandCollapse(cell CellAccessibler)
	//
	CellArea(cell CellAccessibler) gdk.Rectangle
	//
	CellPosition(cell CellAccessibler) (row int, column int)
	//
	ChildIndex(cell CellAccessibler) int
	//
	RendererState(cell CellAccessibler) CellRendererState
	//
	GrabFocus(cell CellAccessibler) bool
	//
	UpdateRelationset(cell CellAccessibler, relationset atk.RelationSetter)
}

//
type CellAccessibleParent struct {
	*externglib.Object
}

var (
	_ CellAccessibleParenter = (*CellAccessibleParent)(nil)
	_ gextras.Nativer        = (*CellAccessibleParent)(nil)
)

func wrapCellAccessibleParent(obj *externglib.Object) CellAccessibleParenter {
	return &CellAccessibleParent{
		Object: obj,
	}
}

func marshalCellAccessibleParenter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellAccessibleParent(obj), nil
}

//
func (parent *CellAccessibleParent) Activate(cell CellAccessibler) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_activate(_arg0, _arg1)
}

//
func (parent *CellAccessibleParent) Edit(cell CellAccessibler) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_edit(_arg0, _arg1)
}

//
func (parent *CellAccessibleParent) ExpandCollapse(cell CellAccessibler) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_expand_collapse(_arg0, _arg1)
}

//
func (parent *CellAccessibleParent) CellArea(cell CellAccessibler) gdk.Rectangle {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cellRect gdk.Rectangle

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_get_cell_area(_arg0, _arg1, (*C.GdkRectangle)(unsafe.Pointer(&_cellRect)))

	return _cellRect
}

//
func (parent *CellAccessibleParent) CellPosition(cell CellAccessibler) (row int, column int) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.gint                     // in
	var _arg3 C.gint                     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_get_cell_position(_arg0, _arg1, &_arg2, &_arg3)

	var _row int    // out
	var _column int // out

	_row = int(_arg2)
	_column = int(_arg3)

	return _row, _column
}

//
func (parent *CellAccessibleParent) ChildIndex(cell CellAccessibler) int {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.int                      // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	_cret = C.gtk_cell_accessible_parent_get_child_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

//
func (parent *CellAccessibleParent) RendererState(cell CellAccessibler) CellRendererState {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.GtkCellRendererState     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	_cret = C.gtk_cell_accessible_parent_get_renderer_state(_arg0, _arg1)

	var _cellRendererState CellRendererState // out

	_cellRendererState = CellRendererState(_cret)

	return _cellRendererState
}

//
func (parent *CellAccessibleParent) GrabFocus(cell CellAccessibler) bool {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	_cret = C.gtk_cell_accessible_parent_grab_focus(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

//
func (parent *CellAccessibleParent) UpdateRelationset(cell CellAccessibler, relationset atk.RelationSetter) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 *C.AtkRelationSet          // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer((cell).(gextras.Nativer).Native()))
	_arg2 = (*C.AtkRelationSet)(unsafe.Pointer((relationset).(gextras.Nativer).Native()))

	C.gtk_cell_accessible_parent_update_relationset(_arg0, _arg1, _arg2)
}
