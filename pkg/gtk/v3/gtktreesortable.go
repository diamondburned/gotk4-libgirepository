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
// extern gboolean _gotk4_gtk3_TreeSortableIface_get_sort_column_id(void*, void*, void*);
// extern gboolean _gotk4_gtk3_TreeSortableIface_has_default_sort_func(void*);
// extern gint _gotk4_gtk3_TreeIterCompareFunc(void*, void*, void*, gpointer);
// extern void _gotk4_gtk3_TreeSortableIface_sort_column_changed(void*);
// extern void _gotk4_gtk3_TreeSortable_ConnectSortColumnChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtktreesortable.go.
var GTypeTreeSortable = coreglib.Type(C.gtk_tree_sortable_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeSortable, F: marshalTreeSortable},
	})
}

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID can be used to make a TreeSortable use
// the default sort function.
//
// See also gtk_tree_sortable_set_sort_column_id().
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID:
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID can be used to make a TreeSortable
// use no sorting.
//
// See also gtk_tree_sortable_set_sort_column_id().
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if a sorts before b, a sorts with b, or a sorts after b respectively.
// If two iters compare as equal, their order in the sorted model is undefined.
// In order to ensure that the TreeSortable behaves as expected, the
// GtkTreeIterCompareFunc must define a partial order on the model, i.e. it must
// be reflexive, antisymmetric and transitive.
//
// For example, if model is a product catalogue, then a compare function for the
// “price” column could be one which returns price_of(a) - price_of(b).
type TreeIterCompareFunc func(model TreeModeller, a, b *TreeIter) (gint int32)

//export _gotk4_gtk3_TreeIterCompareFunc
func _gotk4_gtk3_TreeIterCompareFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) (cret C.gint) {
	var fn TreeIterCompareFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeIterCompareFunc)
	}

	var _model TreeModeller // out
	var _a *TreeIter        // out
	var _b *TreeIter        // out

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
	_a = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_b = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	gint := fn(_model, _a, _b)

	cret = C.gint(gint)

	return cret
}

// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
//
// TreeSortable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeSortable struct {
	_ [0]func() // equal guard
	TreeModel
}

var ()

// TreeSortabler describes TreeSortable's interface methods.
type TreeSortabler interface {
	coreglib.Objector

	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order.
	SortColumnID() (int32, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	HasDefaultSortFunc() bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be sort_func.
	SetDefaultSortFunc(sortFunc TreeIterCompareFunc)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortColumnId int32, sortFunc TreeIterCompareFunc)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// sortable.
	SortColumnChanged()

	// Sort-column-changed signal is emitted when the sort column or sort order
	// of sortable is changed.
	ConnectSortColumnChanged(func()) coreglib.SignalHandle
}

var _ TreeSortabler = (*TreeSortable)(nil)

func wrapTreeSortable(obj *coreglib.Object) *TreeSortable {
	return &TreeSortable{
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	return wrapTreeSortable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_TreeSortable_ConnectSortColumnChanged
func _gotk4_gtk3_TreeSortable_ConnectSortColumnChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectSortColumnChanged signal is emitted when the sort column or sort order
// of sortable is changed. The signal is emitted before the contents of sortable
// are resorted.
func (sortable *TreeSortable) ConnectSortColumnChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sortable, "sort-column-changed", false, unsafe.Pointer(C._gotk4_gtk3_TreeSortable_ConnectSortColumnChanged), f)
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
func (sortable *TreeSortable) SortColumnID() (int32, SortType, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument
	var _arg0 *C.void    // out
	var _out0 *C.void    // in
	var _out1 *C.void    // in
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sortable)

	var _sortColumnId int32 // out
	var _order SortType     // out
	var _ok bool            // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))
	_out1 = *(**C.void)(unsafe.Pointer(&_outs[1]))

	_sortColumnId = *(*int32)(unsafe.Pointer(_out0))
	_order = *(*SortType)(unsafe.Pointer(_out1))
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
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var _args [4]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gpointer // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeIterCompareFunc)
	_arg2 = C.gpointer(gbox.Assign(sortFunc))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = _arg1

	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortFunc)
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
func (sortable *TreeSortable) SetSortFunc(sortColumnId int32, sortFunc TreeIterCompareFunc) {
	var _args [5]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gint     // out
	var _arg2 C.gpointer // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	_arg1 = C.gint(sortColumnId)
	_arg2 = (*[0]byte)(C._gotk4_gtk3_TreeIterCompareFunc)
	_arg3 = C.gpointer(gbox.Assign(sortFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gpointer)(unsafe.Pointer(&_args[2])) = _arg2

	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(sortFunc)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	runtime.KeepAlive(sortable)
}
