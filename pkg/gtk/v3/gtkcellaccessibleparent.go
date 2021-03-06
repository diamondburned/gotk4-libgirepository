// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern GtkCellRendererState _gotk4_gtk3_CellAccessibleParentIface_get_renderer_state(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern gboolean _gotk4_gtk3_CellAccessibleParentIface_grab_focus(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern int _gotk4_gtk3_CellAccessibleParentIface_get_child_index(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_activate(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_edit(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_expand_collapse(GtkCellAccessibleParent*, GtkCellAccessible*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_get_cell_area(GtkCellAccessibleParent*, GtkCellAccessible*, GdkRectangle*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_get_cell_extents(GtkCellAccessibleParent*, GtkCellAccessible*, gint*, gint*, gint*, gint*, AtkCoordType);
// extern void _gotk4_gtk3_CellAccessibleParentIface_get_cell_position(GtkCellAccessibleParent*, GtkCellAccessible*, gint*, gint*);
// extern void _gotk4_gtk3_CellAccessibleParentIface_update_relationset(GtkCellAccessibleParent*, GtkCellAccessible*, AtkRelationSet*);
import "C"

// GTypeCellAccessibleParent returns the GType for the type CellAccessibleParent.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellAccessibleParent() coreglib.Type {
	gtype := coreglib.Type(C.gtk_cell_accessible_parent_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellAccessibleParent)
	return gtype
}

// CellAccessibleParentOverrider contains methods that are overridable.
type CellAccessibleParentOverrider interface {
	// The function takes the following parameters:
	//
	Activate(cell *CellAccessible)
	// The function takes the following parameters:
	//
	Edit(cell *CellAccessible)
	// The function takes the following parameters:
	//
	ExpandCollapse(cell *CellAccessible)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	CellArea(cell *CellAccessible) *gdk.Rectangle
	// The function takes the following parameters:
	//
	//    - cell
	//    - coordType
	//
	// The function returns the following values:
	//
	//    - x
	//    - y
	//    - width
	//    - height
	//
	CellExtents(cell *CellAccessible, coordType atk.CoordType) (x, y, width, height int32)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	//    - row
	//    - column
	//
	CellPosition(cell *CellAccessible) (row, column int32)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	ChildIndex(cell *CellAccessible) int32
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	RendererState(cell *CellAccessible) CellRendererState
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	GrabFocus(cell *CellAccessible) bool
	// The function takes the following parameters:
	//
	//    - cell
	//    - relationset
	//
	UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet)
}

//
// CellAccessibleParent wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CellAccessibleParent struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellAccessibleParent)(nil)
)

// CellAccessibleParenter describes CellAccessibleParent's interface methods.
type CellAccessibleParenter interface {
	coreglib.Objector

	Activate(cell *CellAccessible)
	Edit(cell *CellAccessible)
	ExpandCollapse(cell *CellAccessible)
	CellArea(cell *CellAccessible) *gdk.Rectangle
	CellExtents(cell *CellAccessible, coordType atk.CoordType) (x, y, width, height int32)
	CellPosition(cell *CellAccessible) (row, column int32)
	ChildIndex(cell *CellAccessible) int32
	RendererState(cell *CellAccessible) CellRendererState
	GrabFocus(cell *CellAccessible) bool
	UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet)
}

var _ CellAccessibleParenter = (*CellAccessibleParent)(nil)

func ifaceInitCellAccessibleParenter(gifacePtr, data C.gpointer) {
	iface := (*C.GtkCellAccessibleParentIface)(unsafe.Pointer(gifacePtr))
	iface.activate = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_activate)
	iface.edit = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_edit)
	iface.expand_collapse = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_expand_collapse)
	iface.get_cell_area = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_get_cell_area)
	iface.get_cell_extents = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_get_cell_extents)
	iface.get_cell_position = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_get_cell_position)
	iface.get_child_index = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_get_child_index)
	iface.get_renderer_state = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_get_renderer_state)
	iface.grab_focus = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_grab_focus)
	iface.update_relationset = (*[0]byte)(C._gotk4_gtk3_CellAccessibleParentIface_update_relationset)
}

//export _gotk4_gtk3_CellAccessibleParentIface_activate
func _gotk4_gtk3_CellAccessibleParentIface_activate(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	iface.Activate(_cell)
}

//export _gotk4_gtk3_CellAccessibleParentIface_edit
func _gotk4_gtk3_CellAccessibleParentIface_edit(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	iface.Edit(_cell)
}

//export _gotk4_gtk3_CellAccessibleParentIface_expand_collapse
func _gotk4_gtk3_CellAccessibleParentIface_expand_collapse(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	iface.ExpandCollapse(_cell)
}

//export _gotk4_gtk3_CellAccessibleParentIface_get_cell_area
func _gotk4_gtk3_CellAccessibleParentIface_get_cell_area(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible, arg2 *C.GdkRectangle) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	cellRect := iface.CellArea(_cell)

	*arg2 = *(*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellRect)))
}

//export _gotk4_gtk3_CellAccessibleParentIface_get_cell_extents
func _gotk4_gtk3_CellAccessibleParentIface_get_cell_extents(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible, arg2 *C.gint, arg3 *C.gint, arg4 *C.gint, arg5 *C.gint, arg6 C.AtkCoordType) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible    // out
	var _coordType atk.CoordType // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))
	_coordType = atk.CoordType(arg6)

	x, y, width, height := iface.CellExtents(_cell, _coordType)

	*arg2 = C.gint(x)
	*arg3 = C.gint(y)
	*arg4 = C.gint(width)
	*arg5 = C.gint(height)
}

//export _gotk4_gtk3_CellAccessibleParentIface_get_cell_position
func _gotk4_gtk3_CellAccessibleParentIface_get_cell_position(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible, arg2 *C.gint, arg3 *C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	row, column := iface.CellPosition(_cell)

	*arg2 = C.gint(row)
	*arg3 = C.gint(column)
}

//export _gotk4_gtk3_CellAccessibleParentIface_get_child_index
func _gotk4_gtk3_CellAccessibleParentIface_get_child_index(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) (cret C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	gint := iface.ChildIndex(_cell)

	cret = C.int(gint)

	return cret
}

//export _gotk4_gtk3_CellAccessibleParentIface_get_renderer_state
func _gotk4_gtk3_CellAccessibleParentIface_get_renderer_state(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) (cret C.GtkCellRendererState) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	cellRendererState := iface.RendererState(_cell)

	cret = C.GtkCellRendererState(cellRendererState)

	return cret
}

//export _gotk4_gtk3_CellAccessibleParentIface_grab_focus
func _gotk4_gtk3_CellAccessibleParentIface_grab_focus(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))

	ok := iface.GrabFocus(_cell)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_CellAccessibleParentIface_update_relationset
func _gotk4_gtk3_CellAccessibleParentIface_update_relationset(arg0 *C.GtkCellAccessibleParent, arg1 *C.GtkCellAccessible, arg2 *C.AtkRelationSet) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellAccessibleParentOverrider)

	var _cell *CellAccessible         // out
	var _relationset *atk.RelationSet // out

	_cell = wrapCellAccessible(coreglib.Take(unsafe.Pointer(arg1)))
	{
		obj := coreglib.Take(unsafe.Pointer(arg2))
		_relationset = &atk.RelationSet{
			Object: obj,
		}
	}

	iface.UpdateRelationset(_cell, _relationset)
}

func wrapCellAccessibleParent(obj *coreglib.Object) *CellAccessibleParent {
	return &CellAccessibleParent{
		Object: obj,
	}
}

func marshalCellAccessibleParent(p uintptr) (interface{}, error) {
	return wrapCellAccessibleParent(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) Activate(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	C.gtk_cell_accessible_parent_activate(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) Edit(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	C.gtk_cell_accessible_parent_edit(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
func (parent *CellAccessibleParent) ExpandCollapse(cell *CellAccessible) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	C.gtk_cell_accessible_parent_expand_collapse(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) CellArea(cell *CellAccessible) *gdk.Rectangle {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.GdkRectangle             // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	C.gtk_cell_accessible_parent_get_cell_area(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _cellRect *gdk.Rectangle // out

	_cellRect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _cellRect
}

// The function takes the following parameters:
//
//    - cell
//    - coordType
//
// The function returns the following values:
//
//    - x
//    - y
//    - width
//    - height
//
func (parent *CellAccessibleParent) CellExtents(cell *CellAccessible, coordType atk.CoordType) (x, y, width, height int32) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.gint                     // in
	var _arg3 C.gint                     // in
	var _arg4 C.gint                     // in
	var _arg5 C.gint                     // in
	var _arg6 C.AtkCoordType             // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	_arg6 = C.AtkCoordType(coordType)

	C.gtk_cell_accessible_parent_get_cell_extents(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5, _arg6)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(coordType)

	var _x int32      // out
	var _y int32      // out
	var _width int32  // out
	var _height int32 // out

	_x = int32(_arg2)
	_y = int32(_arg3)
	_width = int32(_arg4)
	_height = int32(_arg5)

	return _x, _y, _width, _height
}

// The function takes the following parameters:
//
// The function returns the following values:
//
//    - row
//    - column
//
func (parent *CellAccessibleParent) CellPosition(cell *CellAccessible) (row, column int32) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 C.gint                     // in
	var _arg3 C.gint                     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	C.gtk_cell_accessible_parent_get_cell_position(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _row int32    // out
	var _column int32 // out

	_row = int32(_arg2)
	_column = int32(_arg3)

	return _row, _column
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) ChildIndex(cell *CellAccessible) int32 {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.int                      // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.gtk_cell_accessible_parent_get_child_index(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) RendererState(cell *CellAccessible) CellRendererState {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.GtkCellRendererState     // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.gtk_cell_accessible_parent_get_renderer_state(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _cellRendererState CellRendererState // out

	_cellRendererState = CellRendererState(_cret)

	return _cellRendererState
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (parent *CellAccessibleParent) GrabFocus(cell *CellAccessible) bool {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))

	_cret = C.gtk_cell_accessible_parent_grab_focus(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - cell
//    - relationset
//
func (parent *CellAccessibleParent) UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet) {
	var _arg0 *C.GtkCellAccessibleParent // out
	var _arg1 *C.GtkCellAccessible       // out
	var _arg2 *C.AtkRelationSet          // out

	_arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	_arg2 = (*C.AtkRelationSet)(unsafe.Pointer(coreglib.InternObject(relationset).Native()))

	C.gtk_cell_accessible_parent_update_relationset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(relationset)
}
