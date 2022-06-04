// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtktreeexpander.go.
var GTypeTreeExpander = coreglib.Type(C.gtk_tree_expander_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeExpander, F: marshalTreeExpander},
	})
}

// TreeExpanderOverrider contains methods that are overridable.
type TreeExpanderOverrider interface {
}

// TreeExpander: GtkTreeExpander is a widget that provides an expander for a
// list.
//
// It is typically placed as a bottommost child into a GtkListView to allow
// users to expand and collapse children in a list with a gtk.TreeListModel.
// GtkTreeExpander provides the common UI elements, gestures and keybindings for
// this purpose.
//
// On top of this, the "listitem.expand", "listitem.collapse" and
// "listitem.toggle-expand" actions are provided to allow adding custom UI for
// managing expanded state.
//
// The GtkTreeListModel must be set to not be passthrough. Then it will provide
// gtk.TreeListRow items which can be set via gtk.TreeExpander.SetListRow() on
// the expander. The expander will then watch that row item automatically.
// gtk.TreeExpander.SetChild() sets the widget that displays the actual row
// contents.
//
// CSS nodes
//
//    treeexpander
//    ├── [indent]*
//    ├── [expander]
//    ╰── <child>
//
//
// GtkTreeExpander has zero or one CSS nodes with the name "expander" that
// should display the expander icon. The node will be :checked when it is
// expanded. If the node is not expandable, an "indent" node will be displayed
// instead.
//
// For every level of depth, another "indent" node is prepended.
//
//
// Accessibility
//
// GtkTreeExpander uses the GTK_ACCESSIBLE_ROLE_GROUP role. The expander icon is
// represented as a GTK_ACCESSIBLE_ROLE_BUTTON, labelled by the expander's
// child, and toggling it will change the GTK_ACCESSIBLE_STATE_EXPANDED state.
type TreeExpander struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*TreeExpander)(nil)
)

func classInitTreeExpanderer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeExpander(obj *coreglib.Object) *TreeExpander {
	return &TreeExpander{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalTreeExpander(p uintptr) (interface{}, error) {
	return wrapTreeExpander(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeExpander creates a new GtkTreeExpander.
//
// The function returns the following values:
//
//    - treeExpander: new GtkTreeExpander.
//
func NewTreeExpander() *TreeExpander {
	_gret := girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("new_TreeExpander", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _treeExpander *TreeExpander // out

	_treeExpander = wrapTreeExpander(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeExpander
}

// Child gets the child widget displayed by self.
//
// The function returns the following values:
//
//    - widget (optional): child displayed by self.
//
func (self *TreeExpander) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Item forwards the item set on the GtkTreeListRow that self is managing.
//
// This call is essentially equivalent to calling:
//
//    gtk_tree_list_row_get_item (gtk_tree_expander_get_list_row (self));.
//
// The function returns the following values:
//
//    - object (optional): item of the row.
//
func (self *TreeExpander) Item() *coreglib.Object {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("get_item", _args[:], nil)
	_cret = *(*C.gpointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _object *coreglib.Object // out

	_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// ListRow gets the list row managed by self.
//
// The function returns the following values:
//
//    - treeListRow (optional): list row displayed by self.
//
func (self *TreeExpander) ListRow() *TreeListRow {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("get_list_row", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _treeListRow *TreeListRow // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeListRow = wrapTreeListRow(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// SetChild sets the content widget to display.
//
// The function takes the following parameters:
//
//    - child (optional): GtkWidget, or NULL.
//
func (self *TreeExpander) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetListRow sets the tree list row that this expander should manage.
//
// The function takes the following parameters:
//
//    - listRow (optional): GtkTreeListRow, or NULL.
//
func (self *TreeExpander) SetListRow(listRow *TreeListRow) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if listRow != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(listRow).Native()))
	}

	girepository.MustFind("Gtk", "TreeExpander").InvokeMethod("set_list_row", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(listRow)
}
