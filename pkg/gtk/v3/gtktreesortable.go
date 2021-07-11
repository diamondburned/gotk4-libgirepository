// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_tree_sortable_get_type()), F: marshalTreeSortabler},
	})
}

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if @a sorts before @b, @a sorts with @b, or @a sorts after @b
// respectively. If two iters compare as equal, their order in the sorted model
// is undefined. In order to ensure that the TreeSortable behaves as expected,
// the GtkTreeIterCompareFunc must define a partial order on the model, i.e. it
// must be reflexive, antisymmetric and transitive.
//
// For example, if @model is a product catalogue, then a compare function for
// the “price” column could be one which returns `price_of(@a) - price_of(@b)`.
type TreeIterCompareFunc func(model *TreeModel, a *TreeIter, b *TreeIter, userData cgo.Handle) (gint int)

//export gotk4_TreeIterCompareFunc
func gotk4_TreeIterCompareFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var model *TreeModel    // out
	var a *TreeIter         // out
	var b *TreeIter         // out
	var userData cgo.Handle // out

	model = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*TreeModel)
	a = (*TreeIter)(unsafe.Pointer(arg1))
	b = (*TreeIter)(unsafe.Pointer(arg2))
	userData = (cgo.Handle)(arg3)

	fn := v.(TreeIterCompareFunc)
	gint := fn(model, a, b, userData)

	cret = C.gint(gint)

	return cret
}

// TreeSortableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeSortableOverrider interface {
	// SortColumnID fills in @sort_column_id and @order with the current sort
	// column and the order. It returns true unless the @sort_column_id is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
	// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns true if the model has a default sort function.
	// This is used primarily by GtkTreeViewColumns in order to determine if a
	// model can go back to the default state, or not.
	HasDefaultSortFunc() bool
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// @sortable.
	SortColumnChanged()
}

// TreeSortabler describes TreeSortable's methods.
type TreeSortabler interface {
	// SortColumnID fills in @sort_column_id and @order with the current sort
	// column and the order.
	SortColumnID() (int, SortType, bool)
	// HasDefaultSortFunc returns true if the model has a default sort function.
	HasDefaultSortFunc() bool
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// @sortable.
	SortColumnChanged()
}

// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
type TreeSortable struct {
	TreeModel
}

var (
	_ TreeSortabler   = (*TreeSortable)(nil)
	_ gextras.Nativer = (*TreeSortable)(nil)
)

func wrapTreeSortable(obj *externglib.Object) TreeSortabler {
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

// SortColumnID fills in @sort_column_id and @order with the current sort column
// and the order. It returns true unless the @sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (sortable *TreeSortable) SortColumnID() (int, SortType, bool) {
	var _arg0 *C.GtkTreeSortable // out
	var _arg1 C.gint             // in
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

// HasDefaultSortFunc returns true if the model has a default sort function.
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

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// @sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _arg0 *C.GtkTreeSortable // out

	_arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(sortable.Native()))

	C.gtk_tree_sortable_sort_column_changed(_arg0)
}
