// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tree_get_type()), F: marshalTree},
	})
}

// Tree: the GTree struct is an opaque data structure representing a [balanced
// binary tree][glib-Balanced-Binary-Trees]. It should be accessed only by using
// the following functions.
type Tree struct {
	native C.GTree
}

// WrapTree wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTree(ptr unsafe.Pointer) *Tree {
	if ptr == nil {
		return nil
	}

	return (*Tree)(ptr)
}

func marshalTree(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTree(unsafe.Pointer(b)), nil
}

// NewTreeFull constructs a struct Tree.
func NewTreeFull() *Tree {
	cret := new(C.GTree)
	var goret *Tree

	cret = C.g_tree_new_full(arg1, arg2, arg3, arg4)

	goret = WrapTree(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Tree) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// NewTreeWithData constructs a struct Tree.
func NewTreeWithData() *Tree {
	cret := new(C.GTree)
	var goret *Tree

	cret = C.g_tree_new_with_data(arg1, arg2)

	goret = WrapTree(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Tree) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (t *Tree) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Destroy removes all keys and values from the #GTree and decreases its
// reference count by one. If keys and/or values are dynamically allocated, you
// should either free them first or create the #GTree using g_tree_new_full().
// In the latter case the destroy functions you supplied will be called on all
// keys and values before destroying the #GTree.
func (t *Tree) Destroy() {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_destroy(arg0)
}

// Foreach calls the given function for each of the key/value pairs in the
// #GTree. The function is passed the key and value of each pair, and the given
// @data parameter. The tree is traversed in sorted order.
//
// The tree may not be modified while iterating over it (you can't add/remove
// items). To remove all items matching a predicate, you need to add each item
// to a list in your Func as you walk over the tree, then walk the list and
// remove each item.
func (t *Tree) Foreach() {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_foreach(arg0, arg1, arg2)
}

// ForeachNode calls the given function for each of the nodes in the #GTree. The
// function is passed the pointer to the particular node, and the given @data
// parameter. The tree traversal happens in-order.
//
// The tree may not be modified while iterating over it (you can't add/remove
// items). To remove all items matching a predicate, you need to add each item
// to a list in your Func as you walk over the tree, then walk the list and
// remove each item.
func (t *Tree) ForeachNode() {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_foreach_node(arg0, arg1, arg2)
}

// Height gets the height of a #GTree.
//
// If the #GTree contains no nodes, the height is 0. If the #GTree contains only
// one root node the height is 1. If the root node has children the height is 2,
// etc.
func (t *Tree) Height() int {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var goret int

	cret = C.g_tree_height(arg0)

	goret = int(cret)

	return goret
}

// Insert inserts a key/value pair into a #GTree.
//
// Inserts a new key and value into a #GTree as g_tree_insert_node() does, only
// this function does not return the inserted or set node.
func (t *Tree) Insert(key interface{}, value interface{}) {
	var arg0 *C.GTree
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)
	arg2 = C.gpointer(value)

	C.g_tree_insert(arg0, arg1, arg2)
}

// InsertNode inserts a key/value pair into a #GTree.
//
// If the given key already exists in the #GTree its corresponding value is set
// to the new value. If you supplied a @value_destroy_func when creating the
// #GTree, the old value is freed using that function. If you supplied a
// @key_destroy_func when creating the #GTree, the passed key is freed using
// that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added, so
// that the distance from the root to every leaf is as small as possible. The
// cost of maintaining a balanced tree while inserting new key/value result in a
// O(n log(n)) operation where most of the other operations are O(log(n)).
func (t *Tree) InsertNode(key interface{}, value interface{}) *TreeNode {
	var arg0 *C.GTree
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)
	arg2 = C.gpointer(value)

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_insert_node(arg0, arg1, arg2)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Lookup gets the value corresponding to the given key. Since a #GTree is
// automatically balanced as key/value pairs are added, key lookup is O(log n)
// (where n is the number of key/value pairs in the tree).
func (t *Tree) Lookup(key interface{}) interface{} {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret C.gpointer
	var goret interface{}

	cret = C.g_tree_lookup(arg0, arg1)

	goret = interface{}(cret)

	return goret
}

// LookupExtended looks up a key in the #GTree, returning the original key and
// the associated value. This is useful if you need to free the memory allocated
// for the original key, for example before calling g_tree_remove().
func (t *Tree) LookupExtended(lookupKey interface{}) (origKey interface{}, value interface{}, ok bool) {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(lookupKey)

	arg2 := new(C.gpointer)
	var ret2 interface{}
	arg3 := new(C.gpointer)
	var ret3 interface{}
	var cret C.gboolean
	var goret bool

	cret = C.g_tree_lookup_extended(arg0, arg1, arg2, arg3)

	ret2 = interface{}(*arg2)
	ret3 = interface{}(*arg3)
	if cret {
		goret = true
	}

	return ret2, ret3, goret
}

// LookupNode gets the tree node corresponding to the given key. Since a #GTree
// is automatically balanced as key/value pairs are added, key lookup is O(log
// n) (where n is the number of key/value pairs in the tree).
func (t *Tree) LookupNode(key interface{}) *TreeNode {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_lookup_node(arg0, arg1)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// LowerBound gets the lower bound node corresponding to the given key, or nil
// if the tree is empty or all the nodes in the tree have keys that are strictly
// lower than the searched key.
//
// The lower bound is the first node that has its key greater than or equal to
// the searched key.
func (t *Tree) LowerBound(key interface{}) *TreeNode {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_lower_bound(arg0, arg1)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Nnodes gets the number of nodes in a #GTree.
func (t *Tree) Nnodes() int {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret C.gint
	var goret int

	cret = C.g_tree_nnodes(arg0)

	goret = int(cret)

	return goret
}

// NodeFirst returns the first in-order node of the tree, or nil for an empty
// tree.
func (t *Tree) NodeFirst() *TreeNode {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_node_first(arg0)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// NodeLast returns the last in-order node of the tree, or nil for an empty
// tree.
func (t *Tree) NodeLast() *TreeNode {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_node_last(arg0)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Ref increments the reference count of @tree by one.
//
// It is safe to call this function from any thread.
func (t *Tree) Ref() *Tree {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	cret := new(C.GTree)
	var goret *Tree

	cret = C.g_tree_ref(arg0)

	goret = WrapTree(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *Tree) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Remove removes a key/value pair from a #GTree.
//
// If the #GTree was created using g_tree_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself. If the key does not
// exist in the #GTree, the function does nothing.
//
// The cost of maintaining a balanced tree while removing a key/value result in
// a O(n log(n)) operation where most of the other operations are O(log(n)).
func (t *Tree) Remove(key interface{}) bool {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret C.gboolean
	var goret bool

	cret = C.g_tree_remove(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Replace inserts a new key and value into a #GTree as g_tree_replace_node()
// does, only this function does not return the inserted or set node.
func (t *Tree) Replace(key interface{}, value interface{}) {
	var arg0 *C.GTree
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)
	arg2 = C.gpointer(value)

	C.g_tree_replace(arg0, arg1, arg2)
}

// ReplaceNode inserts a new key and value into a #GTree similar to
// g_tree_insert_node(). The difference is that if the key already exists in the
// #GTree, it gets replaced by the new key. If you supplied a
// @value_destroy_func when creating the #GTree, the old value is freed using
// that function. If you supplied a @key_destroy_func when creating the #GTree,
// the old key is freed using that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added, so
// that the distance from the root to every leaf is as small as possible.
func (t *Tree) ReplaceNode(key interface{}, value interface{}) *TreeNode {
	var arg0 *C.GTree
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)
	arg2 = C.gpointer(value)

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_replace_node(arg0, arg1, arg2)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Search searches a #GTree using @search_func.
//
// The @search_func is called with a pointer to the key of a key/value pair in
// the tree, and the passed in @user_data. If @search_func returns 0 for a
// key/value pair, then the corresponding value is returned as the result of
// g_tree_search(). If @search_func returns -1, searching will proceed among the
// key/value pairs that have a smaller key; if @search_func returns 1, searching
// will proceed among the key/value pairs that have a larger key.
func (t *Tree) Search() interface{} {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret C.gpointer
	var goret interface{}

	cret = C.g_tree_search(arg0, arg1, arg2)

	goret = interface{}(cret)

	return goret
}

// SearchNode searches a #GTree using @search_func.
//
// The @search_func is called with a pointer to the key of a key/value pair in
// the tree, and the passed in @user_data. If @search_func returns 0 for a
// key/value pair, then the corresponding node is returned as the result of
// g_tree_search(). If @search_func returns -1, searching will proceed among the
// key/value pairs that have a smaller key; if @search_func returns 1, searching
// will proceed among the key/value pairs that have a larger key.
func (t *Tree) SearchNode() *TreeNode {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_search_node(arg0, arg1, arg2)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Steal removes a key and its associated value from a #GTree without calling
// the key and value destroy functions.
//
// If the key does not exist in the #GTree, the function does nothing.
func (t *Tree) Steal(key interface{}) bool {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret C.gboolean
	var goret bool

	cret = C.g_tree_steal(arg0, arg1)

	if cret {
		goret = true
	}

	return goret
}

// Traverse calls the given function for each node in the #GTree.
func (t *Tree) Traverse() {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_traverse(arg0, arg1, arg2, arg3)
}

// Unref decrements the reference count of @tree by one. If the reference count
// drops to 0, all keys and values will be destroyed (if destroy functions were
// specified) and all memory allocated by @tree will be released.
//
// It is safe to call this function from any thread.
func (t *Tree) Unref() {
	var arg0 *C.GTree

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_unref(arg0)
}

// UpperBound gets the upper bound node corresponding to the given key, or nil
// if the tree is empty or all the nodes in the tree have keys that are lower
// than or equal to the searched key.
//
// The upper bound is the first node that has its key strictly greater than the
// searched key.
func (t *Tree) UpperBound(key interface{}) *TreeNode {
	var arg0 *C.GTree
	var arg1 C.gpointer

	arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))
	arg1 = C.gpointer(key)

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_upper_bound(arg0, arg1)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// TreeNode: an opaque type which identifies a specific node in a #GTree.
type TreeNode struct {
	native C.GTreeNode
}

// WrapTreeNode wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTreeNode(ptr unsafe.Pointer) *TreeNode {
	if ptr == nil {
		return nil
	}

	return (*TreeNode)(ptr)
}

func marshalTreeNode(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTreeNode(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TreeNode) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Key gets the key stored at a particular tree node.
func (n *TreeNode) Key() interface{} {
	var arg0 *C.GTreeNode

	arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var cret C.gpointer
	var goret interface{}

	cret = C.g_tree_node_key(arg0)

	goret = interface{}(cret)

	return goret
}

// Next returns the next in-order node of the tree, or nil if the passed node
// was already the last one.
func (n *TreeNode) Next() *TreeNode {
	var arg0 *C.GTreeNode

	arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_node_next(arg0)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Previous returns the previous in-order node of the tree, or nil if the passed
// node was already the first one.
func (n *TreeNode) Previous() *TreeNode {
	var arg0 *C.GTreeNode

	arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var cret *C.GTreeNode
	var goret *TreeNode

	cret = C.g_tree_node_previous(arg0)

	goret = WrapTreeNode(unsafe.Pointer(cret))

	return goret
}

// Value gets the value stored at a particular tree node.
func (n *TreeNode) Value() interface{} {
	var arg0 *C.GTreeNode

	arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var cret C.gpointer
	var goret interface{}

	cret = C.g_tree_node_value(arg0)

	goret = interface{}(cret)

	return goret
}
