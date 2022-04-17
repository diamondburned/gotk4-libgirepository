// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_TreeSortableIface_get_sort_column_id(GtkTreeSortable*, int*, GtkSortType*);
// extern gboolean _gotk4_gtk4_TreeSortableIface_has_default_sort_func(GtkTreeSortable*);
// extern int _gotk4_gtk4_TreeIterCompareFunc(GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer);
// extern void _gotk4_gtk4_TreeSortableIface_set_sort_column_id(GtkTreeSortable*, int, GtkSortType);
// extern void _gotk4_gtk4_TreeSortableIface_sort_column_changed(GtkTreeSortable*);
// extern void _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtktreesortable.go.
var GTypeTreeSortable = externglib.Type(C.gtk_tree_sortable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTreeSortable, F: marshalTreeSortable},
	})
}

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID uses the default sort function in a
// gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID disables sorting in a gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
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
type TreeIterCompareFunc func(model TreeModelOverrider, a, b *TreeIter) (gint int)

//export _gotk4_gtk4_TreeIterCompareFunc
func _gotk4_gtk4_TreeIterCompareFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 *C.GtkTreeIter, arg4 C.gpointer) (cret C.int) {
	var fn TreeIterCompareFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeIterCompareFunc)
	}

	var _model TreeModelOverrider // out
	var _a *TreeIter              // out
	var _b *TreeIter              // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TreeModelOverrider)
			return ok
		})
		rv, ok := casted.(TreeModelOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_a = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_b = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	gint := fn(_model, _a, _b)

	cret = C.int(gint)

	return cret
}

// TreeSortableOverrider contains methods that are overridable.
type TreeSortableOverrider interface {
	externglib.Objector
	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order. It returns TRUE unless the sort_column_id is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
	// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
	//
	// The function returns the following values:
	//
	//    - sortColumnId: sort column id to be filled in.
	//    - order to be filled in.
	//    - ok: TRUE if the sort column is not one of the special sort column
	//      ids.
	//
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	// This is used primarily by GtkTreeViewColumns in order to determine if a
	// model can go back to the default state, or not.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE, if the model has a default sort function.
	//
	HasDefaultSortFunc() bool
	// SetSortColumnID sets the current sort column to be sort_column_id. The
	// sortable will resort itself to reflect this change, after emitting a
	// TreeSortable::sort-column-changed signal. sort_column_id may either be a
	// regular column id, or one of the following special values:
	//
	// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
	// will be used, if it is set
	//
	// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur.
	//
	// The function takes the following parameters:
	//
	//    - sortColumnId: sort column id to set.
	//    - order: sort order of the column.
	//
	SetSortColumnID(sortColumnId int, order SortType)
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
	_ [0]func() // equal guard
	TreeModel
}

var ()

// TreeSortabler describes TreeSortable's interface methods.
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

	// Sort-column-changed signal is emitted when the sort column or sort order
	// of sortable is changed.
	ConnectSortColumnChanged(func()) externglib.SignalHandle
}

var _ TreeSortabler = (*TreeSortable)(nil)

func ifaceInitTreeSortabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkTreeSortableIface)(unsafe.Pointer(gifacePtr))
	iface.get_sort_column_id = (*[0]byte)(C._gotk4_gtk4_TreeSortableIface_get_sort_column_id)
	iface.has_default_sort_func = (*[0]byte)(C._gotk4_gtk4_TreeSortableIface_has_default_sort_func)
	iface.set_sort_column_id = (*[0]byte)(C._gotk4_gtk4_TreeSortableIface_set_sort_column_id)
	iface.sort_column_changed = (*[0]byte)(C._gotk4_gtk4_TreeSortableIface_sort_column_changed)
}

//export _gotk4_gtk4_TreeSortableIface_get_sort_column_id
func _gotk4_gtk4_TreeSortableIface_get_sort_column_id(arg0 *C.GtkTreeSortable, arg1 *C.int, arg2 *C.GtkSortType) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeSortableOverrider)

	sortColumnId, order, ok := iface.SortColumnID()

	*arg1 = C.int(sortColumnId)
	*arg2 = C.GtkSortType(order)
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeSortableIface_has_default_sort_func
func _gotk4_gtk4_TreeSortableIface_has_default_sort_func(arg0 *C.GtkTreeSortable) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeSortableOverrider)

	ok := iface.HasDefaultSortFunc()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeSortableIface_set_sort_column_id
func _gotk4_gtk4_TreeSortableIface_set_sort_column_id(arg0 *C.GtkTreeSortable, arg1 C.int, arg2 C.GtkSortType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeSortableOverrider)

	var _sortColumnId int // out
	var _order SortType   // out

	_sortColumnId = int(arg1)
	_order = SortType(arg2)

	iface.SetSortColumnID(_sortColumnId, _order)
}

//export _gotk4_gtk4_TreeSortableIface_sort_column_changed
func _gotk4_gtk4_TreeSortableIface_sort_column_changed(arg0 *C.GtkTreeSortable) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeSortableOverrider)

	iface.SortColumnChanged()
}

func wrapTreeSortable(obj *externglib.Object) *TreeSortable {
	return &TreeSortable{
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	return wrapTreeSortable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged
func _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSortColumnChanged signal is emitted when the sort column or sort order
// of sortable is changed. The signal is emitted before the contents of sortable
// are resorted.
func (sortable *TreeSortable) ConnectSortColumnChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(sortable, "sort-column-changed", false, unsafe.Pointer(C._gotk4_gtk4_TreeSortable_ConnectSortColumnChanged), f)
}

// SortColumnID fills in sort_column_id and order with the current sort column
// and the order. It returns TRUE unless the sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
//
// The function returns the following values:
//
//    - sortColumnId: sort column id to be filled in.
//    - order to be filled in.
//    - ok: TRUE if the sort column is not one of the special sort column ids.
//
func (sortable *TreeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // in
	var _arg2 C.GtkSortType      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))

	_cret = C.gtk_tree_sortable_get_sort_column_id(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(sortable)

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
//
// The function returns the following values:
//
//    - ok: TRUE, if the model has a default sort function.
//
func (sortable *TreeSortable) HasDefaultSortFunc() bool {
	var _arg0 *C.GtkTreeSortable // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))

	_cret = C.gtk_tree_sortable_has_default_sort_func(_arg0)
	runtime.KeepAlive(sortable)

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
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.GtkTreeIterCompareFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_default_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortFunc)
}

// SetSortColumnID sets the current sort column to be sort_column_id. The
// sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function will be
// used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set.
//    - order: sort order of the column.
//
func (sortable *TreeSortable) SetSortColumnID(sortColumnId int, order SortType) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.int              // out
	var _arg2 C.GtkSortType      // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = C.GtkSortType(order)

	C.gtk_tree_sortable_set_sort_column_id(_arg0, _arg1, _arg2)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(order)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the current sort column id of sortable is the same as sort_column_id, then
// the model will sort using this function.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set the function for.
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetSortFunc(sortColumnId int, sortFunc TreeIterCompareFunc) {
	var _arg0 *C.GtkTreeSortable       // out
	var _arg1 C.int                    // out
	var _arg2 C.GtkTreeIterCompareFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))
	_arg1 = C.int(sortColumnId)
	_arg2 = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_arg3 = C.gpointer(gbox.Assign(sortFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_sortable_set_sort_func(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(sortFunc)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(externglib.InternObject(sortable).Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
	runtime.KeepAlive(sortable)
}
