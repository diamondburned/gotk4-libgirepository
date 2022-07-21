// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gint _gotk4_gio2_CompareDataFunc(gconstpointer, gconstpointer, gpointer);
// extern gint _gotk4_glib2_CompareDataFunc(gconstpointer, gconstpointer, gpointer);
import "C"

// GTypeListStore returns the GType for the type ListStore.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListStore() coreglib.Type {
	gtype := coreglib.Type(C.g_list_store_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalListStore)
	return gtype
}

// ListStoreOverrider contains methods that are overridable.
type ListStoreOverrider interface {
}

// ListStore is a simple implementation of Model that stores all items in
// memory.
//
// It provides insertions, deletions, and lookups in logarithmic time with a
// fast path for the common case of iterating the list linearly.
type ListStore struct {
	_ [0]func() // equal guard
	*coreglib.Object

	ListModel
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
		ListModel: ListModel{
			Object: obj,
		},
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	return wrapListStore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewListStore creates a new Store with items of type item_type. item_type must
// be a subclass of #GObject.
//
// The function takes the following parameters:
//
//    - itemType of items in the list.
//
// The function returns the following values:
//
//    - listStore: new Store.
//
func NewListStore(itemType coreglib.Type) *ListStore {
	var _arg1 C.GType       // out
	var _cret *C.GListStore // in

	_arg1 = C.GType(itemType)

	_cret = C.g_list_store_new(_arg1)
	runtime.KeepAlive(itemType)

	var _listStore *ListStore // out

	_listStore = wrapListStore(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _listStore
}

// Append appends item to store. item must be of type Store:item-type.
//
// This function takes a ref on item.
//
// Use g_list_store_splice() to append multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - item: new item.
//
func (store *ListStore) Append(item *coreglib.Object) {
	var _arg0 *C.GListStore // out
	var _arg1 C.gpointer    // out

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item.Native()))

	C.g_list_store_append(_arg0, _arg1)
	runtime.KeepAlive(store)
	runtime.KeepAlive(item)
}

// Find looks up the given item in the list store by looping over the items
// until the first occurrence of item. If item was not found, then position will
// not be set, and this method will return FALSE.
//
// If you need to compare the two items with a custom comparison function, use
// g_list_store_find_with_equal_func() with a custom Func instead.
//
// The function takes the following parameters:
//
//    - item: item.
//
// The function returns the following values:
//
//    - position (optional): first position of item, if it was found.
//    - ok: whether store contains item. If it was found, position will be set to
//      the position where item occurred for the first time.
//
func (store *ListStore) Find(item *coreglib.Object) (uint, bool) {
	var _arg0 *C.GListStore // out
	var _arg1 C.gpointer    // out
	var _arg2 C.guint       // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item.Native()))

	_cret = C.g_list_store_find(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(item)

	var _position uint // out
	var _ok bool       // out

	_position = uint(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _position, _ok
}

// Insert inserts item into store at position. item must be of type
// Store:item-type or derived from it. position must be smaller than the length
// of the list, or equal to it to append.
//
// This function takes a ref on item.
//
// Use g_list_store_splice() to insert multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - position at which to insert the new item.
//    - item: new item.
//
func (store *ListStore) Insert(position uint, item *coreglib.Object) {
	var _arg0 *C.GListStore // out
	var _arg1 C.guint       // out
	var _arg2 C.gpointer    // out

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.gpointer(unsafe.Pointer(item.Native()))

	C.g_list_store_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
	runtime.KeepAlive(item)
}

// InsertSorted inserts item into store at a position to be determined by the
// compare_func.
//
// The list must already be sorted before calling this function or the result is
// undefined. Usually you would approach this by only ever inserting items by
// way of this function.
//
// This function takes a ref on item.
//
// The function takes the following parameters:
//
//    - item: new item.
//    - compareFunc: pairwise comparison function for sorting.
//
// The function returns the following values:
//
//    - guint: position at which item was inserted.
//
func (store *ListStore) InsertSorted(item *coreglib.Object, compareFunc glib.CompareDataFunc) uint {
	var _arg0 *C.GListStore      // out
	var _arg1 C.gpointer         // out
	var _arg2 C.GCompareDataFunc // out
	var _arg3 C.gpointer
	var _cret C.guint // in

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item.Native()))
	_arg2 = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
	_arg3 = C.gpointer(gbox.Assign(compareFunc))
	defer gbox.Delete(uintptr(_arg3))

	_cret = C.g_list_store_insert_sorted(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(store)
	runtime.KeepAlive(item)
	runtime.KeepAlive(compareFunc)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Remove removes the item from store that is at position. position must be
// smaller than the current length of the list.
//
// Use g_list_store_splice() to remove multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - position of the item that is to be removed.
//
func (store *ListStore) Remove(position uint) {
	var _arg0 *C.GListStore // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.guint(position)

	C.g_list_store_remove(_arg0, _arg1)
	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
}

// RemoveAll removes all items from store.
func (store *ListStore) RemoveAll() {
	var _arg0 *C.GListStore // out

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))

	C.g_list_store_remove_all(_arg0)
	runtime.KeepAlive(store)
}

// Sort the items in store according to compare_func.
//
// The function takes the following parameters:
//
//    - compareFunc: pairwise comparison function for sorting.
//
func (store *ListStore) Sort(compareFunc glib.CompareDataFunc) {
	var _arg0 *C.GListStore      // out
	var _arg1 C.GCompareDataFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
	_arg2 = C.gpointer(gbox.Assign(compareFunc))
	defer gbox.Delete(uintptr(_arg2))

	C.g_list_store_sort(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(compareFunc)
}

// Splice changes store by removing n_removals items and adding n_additions
// items to it. additions must contain n_additions items of type
// Store:item-type. NULL is not permitted.
//
// This function is more efficient than g_list_store_insert() and
// g_list_store_remove(), because it only emits Model::items-changed once for
// the change.
//
// This function takes a ref on each item in additions.
//
// The parameters position and n_removals must be correct (ie: position +
// n_removals must be less than or equal to the length of the list at the time
// this function is called).
//
// The function takes the following parameters:
//
//    - position at which to make the change.
//    - nRemovals: number of items to remove.
//    - additions items to add.
//
func (store *ListStore) Splice(position, nRemovals uint, additions []*coreglib.Object) {
	var _arg0 *C.GListStore // out
	var _arg1 C.guint       // out
	var _arg2 C.guint       // out
	var _arg3 *C.gpointer   // out
	var _arg4 C.guint

	_arg0 = (*C.GListStore)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nRemovals)
	_arg4 = (C.guint)(len(additions))
	_arg3 = (*C.gpointer)(C.calloc(C.size_t(len(additions)), C.size_t(C.sizeof_gpointer)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.gpointer)(_arg3), len(additions))
		for i := range additions {
			out[i] = C.gpointer(unsafe.Pointer(additions[i].Native()))
		}
	}

	C.g_list_store_splice(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nRemovals)
	runtime.KeepAlive(additions)
}
