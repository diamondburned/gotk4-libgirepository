// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gtk4_TreeViewColumnDropFunc
func _gotk4_gtk4_TreeViewColumnDropFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 *C.void, arg5 C.gpointer) (cret C.gboolean) {
	var fn TreeViewColumnDropFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeViewColumnDropFunc)
	}

	var _treeView *TreeView         // out
	var _column *TreeViewColumn     // out
	var _prevColumn *TreeViewColumn // out
	var _nextColumn *TreeViewColumn // out

	_treeView = wrapTreeView(coreglib.Take(unsafe.Pointer(arg1)))
	_column = wrapTreeViewColumn(coreglib.Take(unsafe.Pointer(arg2)))
	_prevColumn = wrapTreeViewColumn(coreglib.Take(unsafe.Pointer(arg3)))
	_nextColumn = wrapTreeViewColumn(coreglib.Take(unsafe.Pointer(arg4)))

	ok := fn(_treeView, _column, _prevColumn, _nextColumn)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeViewMappingFunc
func _gotk4_gtk4_TreeViewMappingFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
	var fn TreeViewMappingFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeViewMappingFunc)
	}

	var _treeView *TreeView // out
	var _path *TreePath     // out

	_treeView = wrapTreeView(coreglib.Take(unsafe.Pointer(arg1)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn(_treeView, _path)
}

//export _gotk4_gtk4_TreeViewRowSeparatorFunc
func _gotk4_gtk4_TreeViewRowSeparatorFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gboolean) {
	var fn TreeViewRowSeparatorFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeViewRowSeparatorFunc)
	}

	var _model TreeModeller // out
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
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := fn(_model, _iter)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeViewSearchEqualFunc
func _gotk4_gtk4_TreeViewSearchEqualFunc(arg1 *C.void, arg2 C.int, arg3 *C.char, arg4 *C.void, arg5 C.gpointer) (cret C.gboolean) {
	var fn TreeViewSearchEqualFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeViewSearchEqualFunc)
	}

	var _model TreeModeller // out
	var _column int         // out
	var _key string         // out
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
	_column = int(arg2)
	_key = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	ok := fn(_model, _column, _key, _iter)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeView_ConnectColumnsChanged
func _gotk4_gtk4_TreeView_ConnectColumnsChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_TreeView_ConnectCursorChanged
func _gotk4_gtk4_TreeView_ConnectCursorChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_TreeView_ConnectRowActivated
func _gotk4_gtk4_TreeView_ConnectRowActivated(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(path *TreePath, column *TreeViewColumn)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path *TreePath, column *TreeViewColumn))
	}

	var _path *TreePath         // out
	var _column *TreeViewColumn // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_column = wrapTreeViewColumn(coreglib.Take(unsafe.Pointer(arg2)))

	f(_path, _column)
}

//export _gotk4_gtk4_TreeView_ConnectRowCollapsed
func _gotk4_gtk4_TreeView_ConnectRowCollapsed(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(iter *TreeIter, path *TreePath)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *TreeIter, path *TreePath))
	}

	var _iter *TreeIter // out
	var _path *TreePath // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_iter, _path)
}

//export _gotk4_gtk4_TreeView_ConnectRowExpanded
func _gotk4_gtk4_TreeView_ConnectRowExpanded(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) {
	var f func(iter *TreeIter, path *TreePath)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *TreeIter, path *TreePath))
	}

	var _iter *TreeIter // out
	var _path *TreePath // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_iter, _path)
}

//export _gotk4_gtk4_TreeView_ConnectTestCollapseRow
func _gotk4_gtk4_TreeView_ConnectTestCollapseRow(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(iter *TreeIter, path *TreePath) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *TreeIter, path *TreePath) (ok bool))
	}

	var _iter *TreeIter // out
	var _path *TreePath // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := f(_iter, _path)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeView_ConnectTestExpandRow
func _gotk4_gtk4_TreeView_ConnectTestExpandRow(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(iter *TreeIter, path *TreePath) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *TreeIter, path *TreePath) (ok bool))
	}

	var _iter *TreeIter // out
	var _path *TreePath // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := f(_iter, _path)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
