// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// int _gotk4_gtk4_TreeIterCompareFunc(GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_sortable_get_type()), F: marshalTreeSortabler},
	})
}

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID uses the default sort function in a
// gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID()
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID disables sorting in a gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID()
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if a sorts before b, a sorts with b, or a sorts after b respectively.
//
// If two iters compare as equal, their order in the sorted model is undefined.
// In order to ensure that the TreeSortable behaves as expected, the
// GtkTreeIterCompareFunc must define a partial order on the model, i.e. it must
// be reflexive, antisymmetric and transitive.
//
// For example, if model is a product catalogue, then a compare function for the
// “price” column could be one which returns price_of(a) - price_of(b).
type TreeIterCompareFunc func(model TreeModeller, a *TreeIter, b *TreeIter) (gint int)

//export _gotk4_gtk4_TreeIterCompareFunc
func _gotk4_gtk4_TreeIterCompareFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.int) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var model TreeModeller // out
	var a *TreeIter        // out
	var b *TreeIter        // out

	model = (externglib.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(TreeModeller)
	a = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(a, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(v))))
	})
	b = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	runtime.SetFinalizer(b, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(v))))
	})

	fn := v.(TreeIterCompareFunc)
	gint := fn(model, a, b)

	cret = C.int(gint)

	return cret
}

// TreeSortableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeSortableOverrider interface {
	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order. It returns TRUE unless the sort_column_id is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
	// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	// This is used primarily by GtkTreeViewColumns in order to determine if a
	// model can go back to the default state, or not.
	HasDefaultSortFunc() bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be sort_func. If the current sort column id of sortable is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using
	// this function.
	//
	// If sort_func is NULL, then there will be no default comparison function.
	// This means that once the model has been sorted, it can’t go back to the
	// default state. In this case, when the current sort column id of sortable
	// is GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
	SetDefaultSortFunc(sortFunc TreeIterCompareFunc)
	// SetSortColumnID sets the current sort column to be sort_column_id. The
	// sortable will resort itself to reflect this change, after emitting a
	// TreeSortable::sort-column-changed signal. sort_column_id may either be a
	// regular column id, or one of the following special values:
	//
	// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
	// will be used, if it is set
	//
	// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
	SetSortColumnID(sortColumnId int, order SortType)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func. If the current sort column id of sortable is the same as
	// sort_column_id, then the model will sort using this function.
	SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// sortable.
	SortColumnChanged()
}

// TreeSortable: interface for sortable models used by GtkTreeView
//
// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
type TreeSortable struct {
	TreeModel
}

// TreeSortabler describes TreeSortable's abstract methods.
type TreeSortabler interface {
	externglib.Objector

	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	HasDefaultSortFunc() bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be sort_func.
	SetDefaultSortFunc(sortFunc TreeIterCompareFunc)
	// SetSortColumnID sets the current sort column to be sort_column_id.
	SetSortColumnID(sortColumnId int, order SortType)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// sortable.
	SortColumnChanged()
}

var _ TreeSortabler = (*TreeSortable)(nil)

func wrapTreeSortable(obj *externglib.Object) *TreeSortable {
	return &TreeSortable{
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeSortabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeSortable(obj), nil
}

// SortColumnID fills in sort_column_id and order with the current sort column
// and the order. It returns TRUE unless the sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (sortable *TreeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // in
	var _arg2 C.GtkSortType      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))

	_cret = C.gtk_tree_sortable_get_sort_column_id(_arg0, &_arg1, &_arg2)

	var _sortColumnId int // out
	var _order SortType   // out
	var _ok bool          // out

	_sortColumnId = int(_arg1)
	_order = SortType(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

// HasDefaultSortFunc returns TRUE if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a model
// can go back to the default state, or not.
func (sortable *TreeSortable) HasDefaultSortFunc() bool {
	var _arg0 *C.GtkTreeSortable // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))

	_cret = C.gtk_tree_sortable_has_default_sort_func(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultSortFunc sets the default comparison function used when sorting to
// be sort_func. If the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using this
// function.
//
// If sort_func is NULL, then there will be no default comparison function. This
// means that once the model has been sorted, it can’t go back to the default
// state. In this case, when the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
func (sortable *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.GtkTreeIterCompareFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_default_sort_func(_arg0, _arg1, _arg2, _arg3)
}

// SetSortColumnID sets the current sort column to be sort_column_id. The
// sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function will be
// used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
func (sortable *TreeSortable) SetSortColumnID(sortColumnId int, order SortType) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = C.GtkSortType(order)

	C.gtk_tree_sortable_set_sort_column_id(_arg0, _arg1, _arg2)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the current sort column id of sortable is the same as sort_column_id, then
// the model will sort using this function.
func (sortable *TreeSortable) SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.int                    // out
	var _arg2 C.GtkTreeIterCompareFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg3 = C.gpointer(gbox.Assign(sortFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_sort_func(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
}
