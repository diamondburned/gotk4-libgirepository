// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_tree_selection_get_type()), F: marshalTreeSelection},
	})
}

// TreeSelection: the TreeSelection object is a helper object to manage the
// selection for a TreeView widget. The TreeSelection object is automatically
// created when a new TreeView widget is created, and cannot exist independently
// of this widget. The primary reason the TreeSelection objects exists is for
// cleanliness of code and API. That is, there is no conceptual reason all these
// functions could not be methods on the TreeView widget instead of a separate
// function.
//
// The TreeSelection object is gotten from a TreeView by calling
// gtk_tree_view_get_selection(). It can be manipulated to check the selection
// status of the tree, as well as select and deselect individual rows. Selection
// is done completely view side. As a result, multiple views of the same model
// can have completely different selections. Additionally, you cannot change the
// selection of a row on the model that is not currently displayed by the view
// without expanding its parents first.
//
// One of the important things to remember when monitoring the selection of a
// view is that the TreeSelection::changed signal is mostly a hint. That is, it
// may only emit one signal when a range of rows is selected. Additionally, it
// may on occasion emit a TreeSelection::changed signal when nothing has
// happened (mostly as a result of programmers calling select_row on an already
// selected row).
type TreeSelection interface {
	gextras.Objector

	// CountSelectedRows returns the number of rows that have been selected in
	// @tree.
	CountSelectedRows() int
	// IterIsSelected returns true if the row at @iter is currently selected.
	IterIsSelected(iter *TreeIter) bool
	// PathIsSelected returns true if the row pointed to by @path is currently
	// selected. If @path does not point to a valid location, false is returned
	PathIsSelected(path *TreePath) bool
	// SelectAll selects all the nodes. @selection must be set to
	// K_SELECTION_MULTIPLE mode.
	SelectAll()
	// SelectIter selects the specified iterator.
	SelectIter(iter *TreeIter)
	// SelectPath: select the row at @path.
	SelectPath(path *TreePath)
	// SelectRange selects a range of nodes, determined by @start_path and
	// @end_path inclusive. @selection must be set to K_SELECTION_MULTIPLE mode.
	SelectRange(startPath *TreePath, endPath *TreePath)
	// SetMode sets the selection mode of the @selection. If the previous type
	// was K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
	// previously selected.
	SetMode(typ SelectionMode)
	// UnselectAll unselects all the nodes.
	UnselectAll()
	// UnselectIter unselects the specified iterator.
	UnselectIter(iter *TreeIter)
	// UnselectPath unselects the row at @path.
	UnselectPath(path *TreePath)
	// UnselectRange unselects a range of nodes, determined by @start_path and
	// @end_path inclusive.
	UnselectRange(startPath *TreePath, endPath *TreePath)
}

// treeSelection implements the TreeSelection class.
type treeSelection struct {
	gextras.Objector
}

var _ TreeSelection = (*treeSelection)(nil)

// WrapTreeSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeSelection(obj *externglib.Object) TreeSelection {
	return treeSelection{
		Objector: obj,
	}
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeSelection(obj), nil
}

// CountSelectedRows returns the number of rows that have been selected in
// @tree.
func (s treeSelection) CountSelectedRows() int {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_selection_count_selected_rows(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// IterIsSelected returns true if the row at @iter is currently selected.
func (s treeSelection) IterIsSelected(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_selection_iter_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PathIsSelected returns true if the row pointed to by @path is currently
// selected. If @path does not point to a valid location, false is returned
func (s treeSelection) PathIsSelected(path *TreePath) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_selection_path_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAll selects all the nodes. @selection must be set to
// K_SELECTION_MULTIPLE mode.
func (s treeSelection) SelectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_select_all(_arg0)
}

// SelectIter selects the specified iterator.
func (s treeSelection) SelectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_tree_selection_select_iter(_arg0, _arg1)
}

// SelectPath: select the row at @path.
func (s treeSelection) SelectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_tree_selection_select_path(_arg0, _arg1)
}

// SelectRange selects a range of nodes, determined by @start_path and
// @end_path inclusive. @selection must be set to K_SELECTION_MULTIPLE mode.
func (s treeSelection) SelectRange(startPath *TreePath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath.Native()))

	C.gtk_tree_selection_select_range(_arg0, _arg1, _arg2)
}

// SetMode sets the selection mode of the @selection. If the previous type
// was K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
// previously selected.
func (s treeSelection) SetMode(typ SelectionMode) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 C.GtkSelectionMode  // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkSelectionMode)(typ)

	C.gtk_tree_selection_set_mode(_arg0, _arg1)
}

// UnselectAll unselects all the nodes.
func (s treeSelection) UnselectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_unselect_all(_arg0)
}

// UnselectIter unselects the specified iterator.
func (s treeSelection) UnselectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_tree_selection_unselect_iter(_arg0, _arg1)
}

// UnselectPath unselects the row at @path.
func (s treeSelection) UnselectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_tree_selection_unselect_path(_arg0, _arg1)
}

// UnselectRange unselects a range of nodes, determined by @start_path and
// @end_path inclusive.
func (s treeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath.Native()))

	C.gtk_tree_selection_unselect_range(_arg0, _arg1, _arg2)
}
