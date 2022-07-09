// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk4_TreeDragDestIface_drag_data_received(void*, void*, void*);
// extern gboolean _gotk4_gtk4_TreeDragDestIface_row_drop_possible(void*, void*, void*);
// extern gboolean _gotk4_gtk4_TreeDragSourceIface_drag_data_delete(void*, void*);
// extern gboolean _gotk4_gtk4_TreeDragSourceIface_row_draggable(void*, void*);
// extern void* _gotk4_gtk4_TreeDragSourceIface_drag_data_get(void*, void*);
import "C"

// GTypeTreeDragDest returns the GType for the type TreeDragDest.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeDragDest() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeDragDest").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeDragDest)
	return gtype
}

// GTypeTreeDragSource returns the GType for the type TreeDragSource.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeDragSource() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeDragSource").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeDragSource)
	return gtype
}

// TreeCreateRowDragContent creates a content provider for dragging path from
// tree_model.
//
// The function takes the following parameters:
//
//    - treeModel: TreeModel.
//    - path: row in tree_model.
//
// The function returns the following values:
//
//    - contentProvider: new ContentProvider.
//
func TreeCreateRowDragContent(treeModel TreeModeller, path *TreePath) *gdk.ContentProvider {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(treeModel).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "tree_create_row_drag_content")
	_gret := _info.Invoke(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(treeModel)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_contentProvider = &gdk.ContentProvider{
			Object: obj,
		}
	}

	return _contentProvider
}

// TreeGetRowDragData obtains a tree_model and path from value of target type
// GTK_TYPE_TREE_ROW_DATA.
//
// The returned path must be freed with gtk_tree_path_free().
//
// The function takes the following parameters:
//
//    - value: #GValue.
//
// The function returns the following values:
//
//    - treeModel (optional): TreeModel.
//    - path (optional): row in tree_model.
//    - ok: TRUE if selection_data had target type GTK_TYPE_TREE_ROW_DATA is
//      otherwise valid.
//
func TreeGetRowDragData(value *coreglib.Value) (*TreeModel, *TreePath, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))

	_info := girepository.MustFind("Gtk", "tree_get_row_drag_data")
	_gret := _info.Invoke(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)

	var _treeModel *TreeModel // out
	var _path *TreePath       // out
	var _ok bool              // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_outs[0])))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_outs[1])))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_path)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					args := [1]girepository.Argument{(*C.void)(intern.C)}
					girepository.MustFind("Gtk", "TreePath").InvokeRecordMethod("free", args[:], nil)
				}
			},
		)
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeDragDestOverrider contains methods that are overridable.
type TreeDragDestOverrider interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from value. If dest is outside the
	// tree so that inserting before it is impossible, FALSE will be returned.
	// Also, FALSE may be returned if the new row is not created for some
	// model-specific reason. Should robustly handle a dest no longer found in
	// the model!.
	//
	// The function takes the following parameters:
	//
	//    - dest: row to drop in front of.
	//    - value: data to drop.
	//
	// The function returns the following values:
	//
	//    - ok: whether a new row was created before position dest.
	//
	DragDataReceived(dest *TreePath, value *coreglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path. i.e., can we drop the data in
	// value at that location. dest_path does not have to exist; the return
	// value will almost certainly be FALSE if the parent of dest_path doesn’t
	// exist, though.
	//
	// The function takes the following parameters:
	//
	//    - destPath: destination row.
	//    - value: data being dropped.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if a drop is possible before dest_path.
	//
	RowDropPossible(destPath *TreePath, value *coreglib.Value) bool
}

// TreeDragDest: interface for Drag-and-Drop destinations in GtkTreeView.
//
// TreeDragDest wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeDragDest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeDragDest)(nil)
)

// TreeDragDester describes TreeDragDest's interface methods.
type TreeDragDester interface {
	coreglib.Objector

	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from value.
	DragDataReceived(dest *TreePath, value *coreglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path.
	RowDropPossible(destPath *TreePath, value *coreglib.Value) bool
}

var _ TreeDragDester = (*TreeDragDest)(nil)

func ifaceInitTreeDragDester(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "TreeDragDestIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("drag_data_received"))) = unsafe.Pointer(C._gotk4_gtk4_TreeDragDestIface_drag_data_received)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("row_drop_possible"))) = unsafe.Pointer(C._gotk4_gtk4_TreeDragDestIface_row_drop_possible)
}

//export _gotk4_gtk4_TreeDragDestIface_drag_data_received
func _gotk4_gtk4_TreeDragDestIface_drag_data_received(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _dest *TreePath        // out
	var _value *coreglib.Value // out

	_dest = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_value = coreglib.ValueFromNative(unsafe.Pointer(arg2))

	ok := iface.DragDataReceived(_dest, _value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeDragDestIface_row_drop_possible
func _gotk4_gtk4_TreeDragDestIface_row_drop_possible(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _destPath *TreePath    // out
	var _value *coreglib.Value // out

	_destPath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_value = coreglib.ValueFromNative(unsafe.Pointer(arg2))

	ok := iface.RowDropPossible(_destPath, _value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragDest(obj *coreglib.Object) *TreeDragDest {
	return &TreeDragDest{
		Object: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	return wrapTreeDragDest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from value. If dest is outside the tree so
// that inserting before it is impossible, FALSE will be returned. Also, FALSE
// may be returned if the new row is not created for some model-specific reason.
// Should robustly handle a dest no longer found in the model!.
//
// The function takes the following parameters:
//
//    - dest: row to drop in front of.
//    - value: data to drop.
//
// The function returns the following values:
//
//    - ok: whether a new row was created before position dest.
//
func (dragDest *TreeDragDest) DragDataReceived(dest *TreePath, value *coreglib.Value) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(dest)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(value.Native()))

	_info := girepository.MustFind("Gtk", "TreeDragDest")
	_gret := _info.InvokeIfaceMethod("drag_data_received", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// value at that location. dest_path does not have to exist; the return value
// will almost certainly be FALSE if the parent of dest_path doesn’t exist,
// though.
//
// The function takes the following parameters:
//
//    - destPath: destination row.
//    - value: data being dropped.
//
// The function returns the following values:
//
//    - ok: TRUE if a drop is possible before dest_path.
//
func (dragDest *TreeDragDest) RowDropPossible(destPath *TreePath, value *coreglib.Value) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(destPath)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(value.Native()))

	_info := girepository.MustFind("Gtk", "TreeDragDest")
	_gret := _info.InvokeIfaceMethod("row_drop_possible", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(destPath)
	runtime.KeepAlive(value)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSourceOverrider contains methods that are overridable.
type TreeDragSourceOverrider interface {
	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop. Returns FALSE if the
	// deletion fails because path no longer exists, or for some model-specific
	// reason. Should robustly handle a path no longer found in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was being dragged.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row was successfully deleted.
	//
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at path. Should robustly handle a path no longer
	// found in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was dragged.
	//
	// The function returns the following values:
	//
	//    - contentProvider (optional) for the given path or NULL if none exists.
	//
	DragDataGet(path *TreePath) *gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	//
	// The function takes the following parameters:
	//
	//    - path: row on which user is initiating a drag.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row can be dragged.
	//
	RowDraggable(path *TreePath) bool
}

// TreeDragSource: interface for Drag-and-Drop destinations in GtkTreeView.
//
// TreeDragSource wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeDragSource struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeDragSource)(nil)
)

// TreeDragSourcer describes TreeDragSource's interface methods.
type TreeDragSourcer interface {
	coreglib.Objector

	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at path.
	DragDataGet(path *TreePath) *gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation.
	RowDraggable(path *TreePath) bool
}

var _ TreeDragSourcer = (*TreeDragSource)(nil)

func ifaceInitTreeDragSourcer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "TreeDragSourceIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("drag_data_delete"))) = unsafe.Pointer(C._gotk4_gtk4_TreeDragSourceIface_drag_data_delete)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("drag_data_get"))) = unsafe.Pointer(C._gotk4_gtk4_TreeDragSourceIface_drag_data_get)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("row_draggable"))) = unsafe.Pointer(C._gotk4_gtk4_TreeDragSourceIface_row_draggable)
}

//export _gotk4_gtk4_TreeDragSourceIface_drag_data_delete
func _gotk4_gtk4_TreeDragSourceIface_drag_data_delete(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.DragDataDelete(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeDragSourceIface_drag_data_get
func _gotk4_gtk4_TreeDragSourceIface_drag_data_get(arg0 *C.void, arg1 *C.void) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	contentProvider := iface.DragDataGet(_path)

	if contentProvider != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(contentProvider).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(contentProvider).Native()))
	}

	return cret
}

//export _gotk4_gtk4_TreeDragSourceIface_row_draggable
func _gotk4_gtk4_TreeDragSourceIface_row_draggable(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.RowDraggable(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragSource(obj *coreglib.Object) *TreeDragSource {
	return &TreeDragSource{
		Object: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	return wrapTreeDragSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if the row was successfully deleted.
//
func (dragSource *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeDragSource")
	_gret := _info.InvokeIfaceMethod("drag_data_delete", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to return a ContentProvider representing
// the row at path. Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was dragged.
//
// The function returns the following values:
//
//    - contentProvider (optional) for the given path or NULL if none exists.
//
func (dragSource *TreeDragSource) DragDataGet(path *TreePath) *gdk.ContentProvider {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeDragSource")
	_gret := _info.InvokeIfaceMethod("drag_data_get", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_contentProvider = &gdk.ContentProvider{
				Object: obj,
			}
		}
	}

	return _contentProvider
}

// RowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
//
// The function takes the following parameters:
//
//    - path: row on which user is initiating a drag.
//
// The function returns the following values:
//
//    - ok: TRUE if the row can be dragged.
//
func (dragSource *TreeDragSource) RowDraggable(path *TreePath) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(path)))

	_info := girepository.MustFind("Gtk", "TreeDragSource")
	_gret := _info.InvokeIfaceMethod("row_draggable", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
