// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_TreeSelectionFunc(void*, void*, void*, gboolean, gpointer);
// extern void _gotk4_gtk3_TreeSelectionClass_changed(void*);
// extern void _gotk4_gtk3_TreeSelectionForEachFunc(void*, void*, void*, gpointer);
// extern void _gotk4_gtk3_TreeSelection_ConnectChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// GTypeTreeSelection returns the GType for the type TreeSelection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeSelection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeSelection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeSelection)
	return gtype
}

// TreeSelectionForEachFunc: function used by
// gtk_tree_selection_selected_foreach() to map all selected rows. It will be
// called on every selected row in the view.
type TreeSelectionForEachFunc func(model TreeModeller, path *TreePath, iter *TreeIter)

//export _gotk4_gtk3_TreeSelectionForEachFunc
func _gotk4_gtk3_TreeSelectionForEachFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) {
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
func _gotk4_gtk3_TreeSelectionFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gboolean, arg5 C.gpointer) (cret C.gboolean) {
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

func classInitTreeSelectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "TreeSelectionClass")

	if _, ok := goval.(interface{ Changed() }); ok {
		o := pclass.StructFieldOffset("changed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_TreeSelectionClass_changed)
	}
}

//export _gotk4_gtk3_TreeSelectionClass_changed
func _gotk4_gtk3_TreeSelectionClass_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
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
func (selection *TreeSelection) CountSelectedRows() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("count_selected_rows", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
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
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("get_selected", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)

	var _model *TreeModel // out
	var _iter *TreeIter   // out
	var _ok bool          // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_model = wrapTreeModel(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[1])))))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("get_selected_rows", _args[:], _outs[:])
	_cret := *(**C.GList)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)

	var _model *TreeModel // out
	var _list []*TreePath // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_model = wrapTreeModel(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	}
	_list = make([]*TreePath, 0, gextras.ListSize(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret)))))
	gextras.MoveList(unsafe.Pointer(*(**C.GList)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *TreePath // out
		dst = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					var args [1]girepository.Argument
					*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
					girepository.MustFind("Gtk", "TreePath").InvokeRecordMethod("free", args[:], nil)
				}
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("get_tree_view", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)

	var _treeView *TreeView // out

	_treeView = wrapTreeView(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("iter_is_selected", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_gret := _info.InvokeClassMethod("path_is_selected", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(selection)
	runtime.KeepAlive(path)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SelectAll selects all the nodes. selection must be set to
// K_SELECTION_MULTIPLE mode.
func (selection *TreeSelection) SelectAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("select_all", _args[:], nil)

	runtime.KeepAlive(selection)
}

// SelectIter selects the specified iterator.
//
// The function takes the following parameters:
//
//    - iter to be selected.
//
func (selection *TreeSelection) SelectIter(iter *TreeIter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("select_iter", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("select_path", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(startPath)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(endPath)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("select_range", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_TreeSelectionForEachFunc)
	_args[2] = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_args[2]))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("selected_foreach", _args[:], nil)

	runtime.KeepAlive(selection)
	runtime.KeepAlive(fn)
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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	if fn != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_TreeSelectionFunc)
		_args[2] = C.gpointer(gbox.Assign(fn))
		_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("set_select_function", _args[:], nil)

	runtime.KeepAlive(selection)
	runtime.KeepAlive(fn)
}

// UnselectAll unselects all the nodes.
func (selection *TreeSelection) UnselectAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("unselect_all", _args[:], nil)

	runtime.KeepAlive(selection)
}

// UnselectIter unselects the specified iterator.
//
// The function takes the following parameters:
//
//    - iter to be unselected.
//
func (selection *TreeSelection) UnselectIter(iter *TreeIter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("unselect_iter", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("unselect_path", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(selection).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(startPath)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(endPath)))

	_info := girepository.MustFind("Gtk", "TreeSelection")
	_info.InvokeClassMethod("unselect_range", _args[:], nil)

	runtime.KeepAlive(selection)
	runtime.KeepAlive(startPath)
	runtime.KeepAlive(endPath)
}
