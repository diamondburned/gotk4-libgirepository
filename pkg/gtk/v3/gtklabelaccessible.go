// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeLabelAccessible returns the GType for the type LabelAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeLabelAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "LabelAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalLabelAccessible)
	return gtype
}

// LabelAccessibleOverrider contains methods that are overridable.
type LabelAccessibleOverrider interface {
}

type LabelAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	*coreglib.Object
	atk.Hypertext
	atk.Text
}

var (
	_ coreglib.Objector = (*LabelAccessible)(nil)
)

func classInitLabelAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLabelAccessible(obj *coreglib.Object) *LabelAccessible {
	return &LabelAccessible{
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
		Object: obj,
		Hypertext: atk.Hypertext{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
	}
}

func marshalLabelAccessible(p uintptr) (interface{}, error) {
	return wrapLabelAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
