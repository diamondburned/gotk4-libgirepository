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
		{T: externglib.Type(C.gtk_tree_model_sort_get_type()), F: marshalTreeModelSort},
	})
}

// TreeModelSort: the TreeModelSort is a model which implements the TreeSortable
// interface. It does not hold any data itself, but rather is created with a
// child model and proxies its data. It has identical column types to this child
// model, and the changes in the child are propagated. The primary purpose of
// this model is to provide a way to sort a different model without modifying
// it. Note that the sort function used by TreeModelSort is not guaranteed to be
// stable.
//
// The use of this is best demonstrated through an example. In the following
// sample code we create two TreeView widgets each with a view of the same data.
// As the model is wrapped here by a TreeModelSort, the two TreeViews can each
// sort their view of the data without affecting the other. By contrast, if we
// simply put the same model in each widget, then sorting the first would sort
// the second.
//
// Using a TreeModelSort
//
//    void
//    selection_changed (GtkTreeSelection *selection, gpointer data)
//    {
//      GtkTreeModel *sort_model = NULL;
//      GtkTreeModel *child_model;
//      GtkTreeIter sort_iter;
//      GtkTreeIter child_iter;
//      char *some_data = NULL;
//      char *modified_data;
//
//      // Get the current selected row and the model.
//      if (! gtk_tree_selection_get_selected (selection,
//                                             &sort_model,
//                                             &sort_iter))
//        return;
//
//      // Look up the current value on the selected row and get
//      // a new value to change it to.
//      gtk_tree_model_get (GTK_TREE_MODEL (sort_model), &sort_iter,
//                          COLUMN_1, &some_data,
//                          -1);
//
//      modified_data = change_the_data (some_data);
//      g_free (some_data);
//
//      // Get an iterator on the child model, instead of the sort model.
//      gtk_tree_model_sort_convert_iter_to_child_iter (GTK_TREE_MODEL_SORT (sort_model),
//                                                      &child_iter,
//                                                      &sort_iter);
//
//      // Get the child model and change the value of the row. In this
//      // example, the child model is a GtkListStore. It could be any other
//      // type of model, though.
//      child_model = gtk_tree_model_sort_get_model (GTK_TREE_MODEL_SORT (sort_model));
//      gtk_list_store_set (GTK_LIST_STORE (child_model), &child_iter,
//                          COLUMN_1, &modified_data,
//                          -1);
//      g_free (modified_data);
//    }
type TreeModelSort interface {
	gextras.Objector
	TreeDragSource
	TreeModel
	TreeSortable

	// ClearCache: this function should almost never be called. It clears the
	// @tree_model_sort of any cached iterators that haven’t been reffed with
	// gtk_tree_model_ref_node(). This might be useful if the child model being
	// sorted is static (and doesn’t change often) and there has been a lot of
	// unreffed access to nodes. As a side effect of this function, all unreffed
	// iters will be invalid.
	ClearCache()
	// ConvertChildIterToIter sets @sort_iter to point to the row in
	// @tree_model_sort that corresponds to the row pointed at by @child_iter.
	// If @sort_iter was not set, false is returned. Note: a boolean is only
	// returned since 2.14.
	ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool)
	// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
	// @sorted_iter.
	ConvertIterToChildIter(sortedIter *TreeIter) TreeIter
	// IterIsValid: > This function is slow. Only use it for debugging and/or
	// testing > purposes.
	//
	// Checks if the given iter is a valid iter for this TreeModelSort.
	IterIsValid(iter *TreeIter) bool
	// ResetDefaultSortFunc: this resets the default sort function to be in the
	// “unsorted” state. That is, it is in the same order as the child model. It
	// will re-sort the model to be in the same order as the child model only if
	// the TreeModelSort is in “unsorted” state.
	ResetDefaultSortFunc()
}

// treeModelSort implements the TreeModelSort class.
type treeModelSort struct {
	gextras.Objector
	TreeDragSource
	TreeModel
	TreeSortable
}

var _ TreeModelSort = (*treeModelSort)(nil)

// WrapTreeModelSort wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeModelSort(obj *externglib.Object) TreeModelSort {
	return treeModelSort{
		Objector:       obj,
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
		TreeSortable:   WrapTreeSortable(obj),
	}
}

func marshalTreeModelSort(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeModelSort(obj), nil
}

// ClearCache: this function should almost never be called. It clears the
// @tree_model_sort of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// sorted is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (t treeModelSort) ClearCache() {
	var _arg0 *C.GtkTreeModelSort // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))

	C.gtk_tree_model_sort_clear_cache(_arg0)
}

// ConvertChildIterToIter sets @sort_iter to point to the row in
// @tree_model_sort that corresponds to the row pointed at by @child_iter.
// If @sort_iter was not set, false is returned. Note: a boolean is only
// returned since 2.14.
func (t treeModelSort) ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool) {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg2 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(childIter.Native()))

	var _sortIter TreeIter
	var _cret C.gboolean // in

	_cret = C.gtk_tree_model_sort_convert_child_iter_to_iter(_arg0, _arg2, (*C.GtkTreeIter)(unsafe.Pointer(&_sortIter)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _sortIter, _ok
}

// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
// @sorted_iter.
func (t treeModelSort) ConvertIterToChildIter(sortedIter *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg2 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sortedIter.Native()))

	var _childIter TreeIter

	C.gtk_tree_model_sort_convert_iter_to_child_iter(_arg0, _arg2, (*C.GtkTreeIter)(unsafe.Pointer(&_childIter)))

	return _childIter
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this TreeModelSort.
func (t treeModelSort) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg1 *C.GtkTreeIter      // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_model_sort_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResetDefaultSortFunc: this resets the default sort function to be in the
// “unsorted” state. That is, it is in the same order as the child model. It
// will re-sort the model to be in the same order as the child model only if
// the TreeModelSort is in “unsorted” state.
func (t treeModelSort) ResetDefaultSortFunc() {
	var _arg0 *C.GtkTreeModelSort // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))

	C.gtk_tree_model_sort_reset_default_sort_func(_arg0)
}
