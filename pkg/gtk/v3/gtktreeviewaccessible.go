// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_view_accessible_get_type()), F: marshalTreeViewAccessible},
	})
}

type TreeViewAccessible interface {
	ContainerAccessible
	CellAccessibleParent
}

// treeViewAccessible implements the TreeViewAccessible class.
type treeViewAccessible struct {
	ContainerAccessible
	CellAccessibleParent
}

var _ TreeViewAccessible = (*treeViewAccessible)(nil)

// WrapTreeViewAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewAccessible(obj *externglib.Object) TreeViewAccessible {
	return treeViewAccessible{
		ContainerAccessible:  WrapContainerAccessible(obj),
		CellAccessibleParent: WrapCellAccessibleParent(obj),
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewAccessible(obj), nil
}
