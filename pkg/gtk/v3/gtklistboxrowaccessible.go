// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeListBoxRowAccessible returns the GType for the type ListBoxRowAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListBoxRowAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_list_box_row_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalListBoxRowAccessible)
	return gtype
}

// ListBoxRowAccessibleOverrider contains methods that are overridable.
type ListBoxRowAccessibleOverrider interface {
}

type ListBoxRowAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*ListBoxRowAccessible)(nil)
)

func classInitListBoxRowAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapListBoxRowAccessible(obj *coreglib.Object) *ListBoxRowAccessible {
	return &ListBoxRowAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
	}
}

func marshalListBoxRowAccessible(p uintptr) (interface{}, error) {
	return wrapListBoxRowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
