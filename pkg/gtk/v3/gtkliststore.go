// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkliststore.go.
var GTypeListStore = coreglib.Type(C.gtk_list_store_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeListStore, F: marshalListStore},
	})
}

// ListStoreOverrider contains methods that are overridable.
type ListStoreOverrider interface {
}

// ListStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequentialy, can use all of the
// methods available there. It also implements the TreeSortable interface so it
// can be sorted by the view. Finally, it also implements the tree [drag and
// drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
//
// The ListStore can accept most GObject types as a column type, though it can’t
// accept all custom types. Internally, it will keep a copy of data passed in
// (such as a string or a boxed pointer). Columns that accept #GObjects are
// handled a little differently. The ListStore will keep a reference to the
// object instead of copying the value. As a result, if the object is modified,
// it is up to the application writer to call gtk_tree_model_row_changed() to
// emit the TreeModel::row_changed signal. This most commonly affects lists with
// Pixbufs stored.
//
// An example for creating a simple list store:
//
//    <object class="GtkListStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//      <data>
//        <row>
//          <col id="0">John</col>
//          <col id="1">Doe</col>
//          <col id="2">25</col>
//        </row>
//        <row>
//          <col id="0">Johan</col>
//          <col id="1">Dahlin</col>
//          <col id="2">50</col>
//        </row>
//      </data>
//    </object>.
type ListStore struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

var (
	_ coreglib.Objector = (*ListStore)(nil)
)

func classInitListStorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapListStore(obj *coreglib.Object) *ListStore {
	return &ListStore{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		TreeDragDest: TreeDragDest{
			Object: obj,
		},
		TreeDragSource: TreeDragSource{
			Object: obj,
		},
		TreeSortable: TreeSortable{
			TreeModel: TreeModel{
				Object: obj,
			},
		},
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	return wrapListStore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Append appends a new row to list_store. iter will be changed to point to this
// new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the appended row.
//
func (listStore *ListStore) Append() *TreeIter {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("append", _args[:], _outs[:])

	runtime.KeepAlive(listStore)

	var _iter *TreeIter // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_out0)))

	return _iter
}

// Clear removes all rows from the list store.
func (listStore *ListStore) Clear() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("clear", _args[:], nil)

	runtime.KeepAlive(listStore)
}

// Insert creates a new row at position. iter will be changed to point to this
// new row. If position is -1 or is larger than the number of rows on the list,
// then the new row will be appended to the list. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - position to insert the new row, or -1 for last.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) Insert(position int32) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = C.gint(position)

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gint)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("insert", _args[:], _outs[:])

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(position)

	var _iter *TreeIter // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_out0)))

	return _iter
}

// InsertAfter inserts a new row after sibling. If sibling is NULL, then the row
// will be prepended to the beginning of the list. iter will be changed to point
// to this new row. The row will be empty after this function is called. To fill
// in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	if sibling != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("insert_after", _args[:], _outs[:])

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_out0)))

	return _iter
}

// InsertBefore inserts a new row before sibling. If sibling is NULL, then the
// row will be appended to the end of the list. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function takes the following parameters:
//
//    - sibling (optional): valid TreeIter, or NULL.
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the new row.
//
func (listStore *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	if sibling != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("insert_before", _args[:], _outs[:])

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter *TreeIter // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_out0)))

	return _iter
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this ListStore.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE if the iter is valid, FALSE if the iter is invalid.
//
func (listStore *ListStore) IterIsValid(iter *TreeIter) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ListStore").InvokeMethod("iter_is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves iter in store to the position after position. Note that this
// function only works with unsorted stores. If position is NULL, iter will be
// moved to the start of the list.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position (optional) or NULL.
//
func (store *ListStore) MoveAfter(iter, position *TreeIter) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("move_after", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// MoveBefore moves iter in store to the position before position. Note that
// this function only works with unsorted stores. If position is NULL, iter will
// be moved to the end of the list.
//
// The function takes the following parameters:
//
//    - iter: TreeIter.
//    - position (optional) or NULL.
//
func (store *ListStore) MoveBefore(iter, position *TreeIter) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(position)))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("move_before", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// Prepend prepends a new row to list_store. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
//
// The function returns the following values:
//
//    - iter: unset TreeIter to set to the prepend row.
//
func (listStore *ListStore) Prepend() *TreeIter {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void // out
	var _out0 *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("prepend", _args[:], _outs[:])

	runtime.KeepAlive(listStore)

	var _iter *TreeIter // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(_out0)))

	return _iter
}

// Remove removes the given row from the list store. After being removed, iter
// is set to be the next valid row, or invalidated if it pointed to the last row
// in list_store.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter.
//
// The function returns the following values:
//
//    - ok: TRUE if iter is valid, FALSE if not.
//
func (listStore *ListStore) Remove(iter *TreeIter) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ListStore").InvokeMethod("remove", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reorder reorders store to follow the order indicated by new_order. Note that
// this function only works with unsorted stores.
//
// The function takes the following parameters:
//
//    - newOrder: array of integers mapping the new position of each child to its
//      old position before the re-ordering, i.e. new_order[newpos] = oldpos. It
//      must have exactly as many items as the list store’s length.
//
func (store *ListStore) Reorder(newOrder []int32) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	{
		_arg1 = (*C.void)(C.calloc(C.size_t((len(newOrder) + 1)), C.size_t(C.sizeof_gint)))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(newOrder)+1)
			var zero C.gint
			out[len(newOrder)] = zero
			for i := range newOrder {
				out[i] = C.gint(newOrder[i])
			}
		}
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("reorder", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(newOrder)
}

// SetValue sets the data in the cell specified by iter and column. The type of
// value must be convertible to the type of the column.
//
// The function takes the following parameters:
//
//    - iter: valid TreeIter for the row being modified.
//    - column number to modify.
//    - value: new value for the cell.
//
func (listStore *ListStore) SetValue(iter *TreeIter, column int32, value *coreglib.Value) {
	var _args [4]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gint  // out
	var _arg3 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(listStore).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.gint(column)
	_arg3 = (*C.void)(unsafe.Pointer(value.Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(*C.gint)(unsafe.Pointer(&_args[2])) = _arg2
	*(**C.void)(unsafe.Pointer(&_args[3])) = _arg3

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("set_value", _args[:], nil)

	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(column)
	runtime.KeepAlive(value)
}

// Swap swaps a and b in store. Note that this function only works with unsorted
// stores.
//
// The function takes the following parameters:
//
//    - a: TreeIter.
//    - b: another TreeIter.
//
func (store *ListStore) Swap(a, b *TreeIter) {
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(b)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "ListStore").InvokeMethod("swap", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
