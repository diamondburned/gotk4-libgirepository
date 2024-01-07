// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_TreeSelection_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeTreeSelection = coreglib.Type(girepository.MustFind("Gtk", "TreeSelection").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeSelection, F: marshalTreeSelection},
	})
}

// TreeSelectionForEachFunc: function used by
// gtk_tree_selection_selected_foreach() to map all selected rows. It will be
// called on every selected row in the view.
type TreeSelectionForEachFunc func(model TreeModeller, path *TreePath, iter *TreeIter)

// TreeSelectionFunc: function used by gtk_tree_selection_set_select_function()
// to filter whether or not a row may be selected. It is called whenever a row's
// state might change.
//
// A return value of TRUE indicates to selection that it is okay to change the
// selection.
type TreeSelectionFunc func(selection *TreeSelection, model TreeModeller, path *TreePath, pathCurrentlySelected bool) (ok bool)

// TreeSelection: selection object for GtkTreeView
//
// The TreeSelection object is a helper object to manage the selection for a
// TreeView widget. The TreeSelection object is automatically created when a new
// TreeView widget is created, and cannot exist independently of this widget.
// The primary reason the TreeSelection objects exists is for cleanliness of
// code and API. That is, there is no conceptual reason all these functions
// could not be methods on the TreeView widget instead of a separate function.
//
// The TreeSelection object is gotten from a TreeView by calling
// gtk_tree_view_get_selection(). It can be manipulated to check the selection
// status of the tree, as well as select and deselect individual rows. Selection
// is done completely view side. As a result, multiple views of the same model
// can have completely different selections. Additionally, you cannot change the
// selection of a row on the model that is not currently displayed by the view
// without expanding its parents first.
//
// One of the important things to remember when monitoring the selection of a
// view is that the TreeSelection::changed signal is mostly a hint. That is, it
// may only emit one signal when a range of rows is selected. Additionally, it
// may on occasion emit a TreeSelection::changed signal when nothing has
// happened (mostly as a result of programmers calling select_row on an already
// selected row).
type TreeSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeSelection)(nil)
)

func wrapTreeSelection(obj *coreglib.Object) *TreeSelection {
	return &TreeSelection{
		Object: obj,
	}
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	return wrapTreeSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted whenever the selection has (possibly) changed.
// Please note that this signal is mostly a hint. It may only be emitted once
// when a range of rows are selected, and it may occasionally be emitted when
// nothing has happened.
func (v *TreeSelection) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk4_TreeSelection_ConnectChanged), f)
}
