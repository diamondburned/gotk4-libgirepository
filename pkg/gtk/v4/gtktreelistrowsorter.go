// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeTreeListRowSorter returns the GType for the type TreeListRowSorter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeListRowSorter() coreglib.Type {
	gtype := coreglib.Type(C.gtk_tree_list_row_sorter_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeListRowSorter)
	return gtype
}

// TreeListRowSorterOverrider contains methods that are overridable.
type TreeListRowSorterOverrider interface {
}

// TreeListRowSorter: GtkTreeListRowSorter is a special-purpose sorter that will
// apply a given sorter to the levels in a tree.
//
// Here is an example for setting up a column view with a tree model and a
// GtkTreeListSorter:
//
//    column_sorter = gtk_column_view_get_sorter (view);
//    sorter = gtk_tree_list_row_sorter_new (g_object_ref (column_sorter));
//    sort_model = gtk_sort_list_model_new (tree_model, sorter);
//    selection = gtk_single_selection_new (sort_model);
//    gtk_column_view_set_model (view, G_LIST_MODEL (selection));.
type TreeListRowSorter struct {
	_ [0]func() // equal guard
	Sorter
}

var (
	_ coreglib.Objector = (*TreeListRowSorter)(nil)
)

func classInitTreeListRowSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeListRowSorter(obj *coreglib.Object) *TreeListRowSorter {
	return &TreeListRowSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalTreeListRowSorter(p uintptr) (interface{}, error) {
	return wrapTreeListRowSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeListRowSorter: create a special-purpose sorter that applies the
// sorting of sorter to the levels of a GtkTreeListModel.
//
// Note that this sorter relies on gtk.TreeListModel:passthrough being FALSE as
// it can only sort gtk.TreeListRows.
//
// The function takes the following parameters:
//
//    - sorter (optional): GtkSorter, or NULL.
//
// The function returns the following values:
//
//    - treeListRowSorter: new GtkTreeListRowSorter.
//
func NewTreeListRowSorter(sorter *Sorter) *TreeListRowSorter {
	var _arg1 *C.GtkSorter            // out
	var _cret *C.GtkTreeListRowSorter // in

	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(sorter).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(sorter).Native()))
	}

	_cret = C.gtk_tree_list_row_sorter_new(_arg1)
	runtime.KeepAlive(sorter)

	var _treeListRowSorter *TreeListRowSorter // out

	_treeListRowSorter = wrapTreeListRowSorter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeListRowSorter
}

// GetSorter returns the sorter used by self.
//
// The function returns the following values:
//
//    - sorter (optional) used.
//
func (self *TreeListRowSorter) GetSorter() *Sorter {
	var _arg0 *C.GtkTreeListRowSorter // out
	var _cret *C.GtkSorter            // in

	_arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_tree_list_row_sorter_get_sorter(_arg0)
	runtime.KeepAlive(self)

	var _sorter *Sorter // out

	if _cret != nil {
		_sorter = wrapSorter(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _sorter
}

// SetSorter sets the sorter to use for items with the same parent.
//
// This sorter will be passed the gtk.TreeListRow:item of the tree list rows
// passed to self.
//
// The function takes the following parameters:
//
//    - sorter (optional) to use, or NULL.
//
func (self *TreeListRowSorter) SetSorter(sorter *Sorter) {
	var _arg0 *C.GtkTreeListRowSorter // out
	var _arg1 *C.GtkSorter            // out

	_arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(coreglib.InternObject(sorter).Native()))
	}

	C.gtk_tree_list_row_sorter_set_sorter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sorter)
}
