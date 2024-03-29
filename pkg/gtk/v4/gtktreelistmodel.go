// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeTreeListModel = coreglib.Type(girepository.MustFind("Gtk", "TreeListModel").RegisteredGType())
	GTypeTreeListRow   = coreglib.Type(girepository.MustFind("Gtk", "TreeListRow").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeListModel, F: marshalTreeListModel},
		coreglib.TypeMarshaler{T: GTypeTreeListRow, F: marshalTreeListRow},
	})
}

// TreeListModelCreateModelFunc: prototype of the function called to create new
// child models when gtk_tree_list_row_set_expanded() is called.
//
// This function can return NULL to indicate that item is guaranteed to be a
// leaf node and will never have children. If it does not have children but may
// get children later, it should return an empty model that is filled once
// children arrive.
type TreeListModelCreateModelFunc func(item *coreglib.Object) (listModel *gio.ListModel)

// TreeListModelOverrides contains methods that are overridable.
type TreeListModelOverrides struct {
}

func defaultTreeListModelOverrides(v *TreeListModel) TreeListModelOverrides {
	return TreeListModelOverrides{}
}

// TreeListModel: GtkTreeListModel is a list model that can create child models
// on demand.
type TreeListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*TreeListModel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TreeListModel, *TreeListModelClass, TreeListModelOverrides](
		GTypeTreeListModel,
		initTreeListModelClass,
		wrapTreeListModel,
		defaultTreeListModelOverrides,
	)
}

func initTreeListModelClass(gclass unsafe.Pointer, overrides TreeListModelOverrides, classInitFunc func(*TreeListModelClass)) {
	if classInitFunc != nil {
		class := (*TreeListModelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTreeListModel(obj *coreglib.Object) *TreeListModel {
	return &TreeListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalTreeListModel(p uintptr) (interface{}, error) {
	return wrapTreeListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TreeListRowOverrides contains methods that are overridable.
type TreeListRowOverrides struct {
}

func defaultTreeListRowOverrides(v *TreeListRow) TreeListRowOverrides {
	return TreeListRowOverrides{}
}

// TreeListRow: GtkTreeListRow is used by GtkTreeListModel to represent items.
//
// It allows navigating the model as a tree and modify the state of rows.
//
// GtkTreeListRow instances are created by a GtkTreeListModel only when the
// gtk.TreeListModel:passthrough property is not set.
//
// There are various support objects that can make use of GtkTreeListRow
// objects, such as the gtk.TreeExpander widget that allows displaying an icon
// to expand or collapse a row or gtk.TreeListRowSorter that makes it possible
// to sort trees properly.
type TreeListRow struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeListRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TreeListRow, *TreeListRowClass, TreeListRowOverrides](
		GTypeTreeListRow,
		initTreeListRowClass,
		wrapTreeListRow,
		defaultTreeListRowOverrides,
	)
}

func initTreeListRowClass(gclass unsafe.Pointer, overrides TreeListRowOverrides, classInitFunc func(*TreeListRowClass)) {
	if classInitFunc != nil {
		class := (*TreeListRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTreeListRow(obj *coreglib.Object) *TreeListRow {
	return &TreeListRow{
		Object: obj,
	}
}

func marshalTreeListRow(p uintptr) (interface{}, error) {
	return wrapTreeListRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TreeListModelClass: instance of this type is always passed by reference.
type TreeListModelClass struct {
	*treeListModelClass
}

// treeListModelClass is the struct that's finalized.
type treeListModelClass struct {
	native unsafe.Pointer
}

var GIRInfoTreeListModelClass = girepository.MustFind("Gtk", "TreeListModelClass")

// TreeListRowClass: instance of this type is always passed by reference.
type TreeListRowClass struct {
	*treeListRowClass
}

// treeListRowClass is the struct that's finalized.
type treeListRowClass struct {
	native unsafe.Pointer
}

var GIRInfoTreeListRowClass = girepository.MustFind("Gtk", "TreeListRowClass")
