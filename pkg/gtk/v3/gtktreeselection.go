// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_TreeSelectionFunc(GtkTreeSelection*, GtkTreeModel*, GtkTreePath*, gboolean, gpointer);
// extern void _gotk4_gtk3_TreeSelectionClass_changed(GtkTreeSelection*);
// extern void _gotk4_gtk3_TreeSelectionForEachFunc(GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);
// extern void _gotk4_gtk3_TreeSelection_ConnectChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// GType values.
var (
	GTypeTreeSelection = coreglib.Type(C.gtk_tree_selection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeSelection, F: marshalTreeSelection},
	})
}

// TreeSelectionForEachFunc: function used by
// gtk_tree_selection_selected_foreach() to map all selected rows. It will be
// called on every selected row in the view.
type TreeSelectionForEachFunc func(model TreeModeller, path *TreePath, iter *TreeIter)

//export _gotk4_gtk3_TreeSelectionForEachFunc
func _gotk4_gtk3_TreeSelectionForEachFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreePath, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	var fn TreeSelectionForEachFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeSelectionForEachFunc)
	}

	var _model TreeModeller // out
	var _path *TreePath     // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	fn(_model, _path, _iter)
}

// TreeSelectionFunc: function used by gtk_tree_selection_set_select_function()
// to filter whether or not a row may be selected. It is called whenever a row's
// state might change. A return value of TRUE indicates to selection that it is
// okay to change the selection.
type TreeSelectionFunc func(selection *TreeSelection, model TreeModeller, path *TreePath, pathCurrentlySelected bool) (ok bool)

//export _gotk4_gtk3_TreeSelectionFunc
func _gotk4_gtk3_TreeSelectionFunc(arg1 *C.GtkTreeSelection, arg2 *C.GtkTreeModel, arg3 *C.GtkTreePath, arg4 C.gboolean, arg5 C.gpointer) (cret C.gboolean) {
	var fn TreeSelectionFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeSelectionFunc)
	}

	var _selection *TreeSelection   // out
	var _model TreeModeller         // out
	var _path *TreePath             // out
	var _pathCurrentlySelected bool // out

	_selection = wrapTreeSelection(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	if arg4 != 0 {
		_pathCurrentlySelected = true
	}

	ok := fn(_selection, _model, _path, _pathCurrentlySelected)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TreeSelectionOverrider contains methods that are overridable.
type TreeSelectionOverrider interface {
	Changed()
}

// TreeSelection object is a helper object to manage the selection for a
// TreeView widget. The TreeSelection object is automatically created when a new
// TreeView widget is created, and cannot exist independently of this widget.
// The primary reason the TreeSelection objects exists is for cleanliness of
// code and API. That is, there is no conceptual reason all these functions
// could not be methods on the TreeView widget instead of a separate function.
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
type TreeSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeSelection)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeTreeSelection,
		GoType:    reflect.TypeOf((*TreeSelection)(nil)),
		InitClass: initClassTreeSelection,
	})
}

func initClassTreeSelection(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkTreeSelectionClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gtk3_TreeSelectionClass_changed)
	}
	if goval, ok := goval.(interface{ InitTreeSelection(*TreeSelectionClass) }); ok {
		klass := (*TreeSelectionClass)(gextras.NewStructNative(gclass))
		goval.InitTreeSelection(klass)
	}
}

//export _gotk4_gtk3_TreeSelectionClass_changed
func _gotk4_gtk3_TreeSelectionClass_changed(arg0 *C.GtkTreeSelection) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

func wrapTreeSelection(obj *coreglib.Object) *TreeSelection {
	return &TreeSelection{
		Object: obj,
	}
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	return wrapTreeSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_TreeSelection_ConnectChanged
func _gotk4_gtk3_TreeSelection_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted whenever the selection has (possibly) changed.
// Please note that this signal is mostly a hint. It may only be emitted once
// when a range of rows are selected, and it may occasionally be emitted when
// nothing has happened.
func (selection *TreeSelection) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(selection, "changed", false, unsafe.Pointer(C._gotk4_gtk3_TreeSelection_ConnectChanged), f)
}

// CountSelectedRows returns the number of rows that have been selected in tree.
//
// The function returns the following values:
//
//    - gint: number of rows selected.
//
func (selection *TreeSelection) CountSelectedRows() int {
	var _arg0 *C.GtkTreeSelection // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_cret = C.gtk_tree_selection_count_selected_rows(_arg0)
	runtime.KeepAlive(selection)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Mode gets the selection mode for selection. See
// gtk_tree_selection_set_mode().
//
// The function returns the following values:
//
//    - selectionMode: current selection mode.
//
func (selection *TreeSelection) Mode() SelectionMode {
	var _arg0 *C.GtkTreeSelection // out
	var _cret C.GtkSelectionMode  // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_cret = C.gtk_tree_selection_get_mode(_arg0)
	runtime.KeepAlive(selection)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Selected sets iter to the currently selected node if selection is set to
// K_SELECTION_SINGLE or K_SELECTION_BROWSE. iter may be NULL if you just want
// to test if selection has any selected nodes. model is filled with the current
// model as a convenience. This function will not work if you use selection is
// K_SELECTION_MULTIPLE.
//
// The function returns the following values:
//
//    - model (optional): pointer to set to the TreeModel, or NULL.
//    - iter (optional) or NULL.
//    - ok: TRUE, if there is a selected node.
//
func (selection *TreeSelection) Selected() (*TreeModel, *TreeIter, bool) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeModel     // in
	var _arg2 C.GtkTreeIter       // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_cret = C.gtk_tree_selection_get_selected(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(selection)

	var _model *TreeModel // out
	var _iter *TreeIter   // out
	var _ok bool          // out

	if _arg1 != nil {
		_model = wrapTreeModel(coreglib.Take(unsafe.Pointer(_arg1)))
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _model, _iter, _ok
}

// SelectedRows creates a list of path of all selected rows. Additionally, if
// you are planning on modifying the model after calling this function, you may
// want to convert the returned list into a list of TreeRowReferences. To do
// this, you can use gtk_tree_row_reference_new().
//
// To free the return value, use:
//
//    g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);.
//
// The function returns the following values:
//
//    - model (optional): pointer to set to the TreeModel, or NULL.
//    - list containing a TreePath for each selected row.
//
func (selection *TreeSelection) SelectedRows() (*TreeModel, []*TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeModel     // in
	var _cret *C.GList            // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_cret = C.gtk_tree_selection_get_selected_rows(_arg0, &_arg1)
	runtime.KeepAlive(selection)

	var _model *TreeModel // out
	var _list []*TreePath // out

	if _arg1 != nil {
		_model = wrapTreeModel(coreglib.Take(unsafe.Pointer(_arg1)))
	}
	_list = make([]*TreePath, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkTreePath)(v)
		var dst *TreePath // out
		dst = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _model, _list
}

// TreeView returns the tree view associated with selection.
//
// The function returns the following values:
//
//    - treeView: TreeView.
//
func (selection *TreeSelection) TreeView() *TreeView {
	var _arg0 *C.GtkTreeSelection // out
	var _cret *C.GtkTreeView      // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_cret = C.gtk_tree_selection_get_tree_view(_arg0)
	runtime.KeepAlive(selection)

	var _treeView *TreeView // out

	_treeView = wrapTreeView(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeView
}

// IterIsSelected returns TRUE if the row at iter is currently selected.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE, if iter is selected.
//
func (selection *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_tree_selection_iter_is_selected(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PathIsSelected returns TRUE if the row pointed to by path is currently
// selected. If path does not point to a valid location, FALSE is returned.
//
// The function takes the following parameters:
//
//    - path to check selection on.
//
// The function returns the following values:
//
//    - ok: TRUE if path is selected.
//
func (selection *TreeSelection) PathIsSelected(path *TreePath) bool {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_selection_path_is_selected(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAll selects all the nodes. selection must be set to
// K_SELECTION_MULTIPLE mode.
func (selection *TreeSelection) SelectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	C.gtk_tree_selection_select_all(_arg0)
	runtime.KeepAlive(selection)
}

// SelectIter selects the specified iterator.
//
// The function takes the following parameters:
//
//    - iter to be selected.
//
func (selection *TreeSelection) SelectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.gtk_tree_selection_select_iter(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(iter)
}

// SelectPath: select the row at path.
//
// The function takes the following parameters:
//
//    - path to be selected.
//
func (selection *TreeSelection) SelectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	C.gtk_tree_selection_select_path(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(path)
}

// SelectRange selects a range of nodes, determined by start_path and end_path
// inclusive. selection must be set to K_SELECTION_MULTIPLE mode.
//
// The function takes the following parameters:
//
//    - startPath: initial node of the range.
//    - endPath: final node of the range.
//
func (selection *TreeSelection) SelectRange(startPath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(startPath)))
	_arg2 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(endPath)))

	C.gtk_tree_selection_select_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(startPath)
	runtime.KeepAlive(endPath)
}

// SelectedForEach calls a function for each selected node. Note that you cannot
// modify the tree or selection from within this function. As a result,
// gtk_tree_selection_get_selected_rows() might be more useful.
//
// The function takes the following parameters:
//
//    - fn: function to call for each selected node.
//
func (selection *TreeSelection) SelectedForEach(fn TreeSelectionForEachFunc) {
	var _arg0 *C.GtkTreeSelection           // out
	var _arg1 C.GtkTreeSelectionForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeSelectionForEachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_tree_selection_selected_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(fn)
}

// SetMode sets the selection mode of the selection. If the previous type was
// K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was previously
// selected.
//
// The function takes the following parameters:
//
//    - typ: selection mode.
//
func (selection *TreeSelection) SetMode(typ SelectionMode) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 C.GtkSelectionMode  // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = C.GtkSelectionMode(typ)

	C.gtk_tree_selection_set_mode(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(typ)
}

// SetSelectFunction sets the selection function.
//
// If set, this function is called before any node is selected or unselected,
// giving some control over which nodes are selected. The select function should
// return TRUE if the state of the node may be toggled, and FALSE if the state
// of the node should be left unchanged.
//
// The function takes the following parameters:
//
//    - fn (optional): selection function. May be NULL.
//
func (selection *TreeSelection) SetSelectFunction(fn TreeSelectionFunc) {
	var _arg0 *C.GtkTreeSelection    // out
	var _arg1 C.GtkTreeSelectionFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeSelectionFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_tree_selection_set_select_function(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(fn)
}

// UnselectAll unselects all the nodes.
func (selection *TreeSelection) UnselectAll() {
	var _arg0 *C.GtkTreeSelection // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	C.gtk_tree_selection_unselect_all(_arg0)
	runtime.KeepAlive(selection)
}

// UnselectIter unselects the specified iterator.
//
// The function takes the following parameters:
//
//    - iter to be unselected.
//
func (selection *TreeSelection) UnselectIter(iter *TreeIter) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.gtk_tree_selection_unselect_iter(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(iter)
}

// UnselectPath unselects the row at path.
//
// The function takes the following parameters:
//
//    - path to be unselected.
//
func (selection *TreeSelection) UnselectPath(path *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	C.gtk_tree_selection_unselect_path(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(path)
}

// UnselectRange unselects a range of nodes, determined by start_path and
// end_path inclusive.
//
// The function takes the following parameters:
//
//    - startPath: initial node of the range.
//    - endPath: initial node of the range.
//
func (selection *TreeSelection) UnselectRange(startPath, endPath *TreePath) {
	var _arg0 *C.GtkTreeSelection // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkTreePath      // out

	_arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(startPath)))
	_arg2 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(endPath)))

	C.gtk_tree_selection_unselect_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(startPath)
	runtime.KeepAlive(endPath)
}

// TreeSelectionClass: instance of this type is always passed by reference.
type TreeSelectionClass struct {
	*treeSelectionClass
}

// treeSelectionClass is the struct that's finalized.
type treeSelectionClass struct {
	native *C.GtkTreeSelectionClass
}
