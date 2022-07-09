// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_CellAccessibleClass_update_cache(void*, gboolean);
import "C"

// GTypeCellAccessible returns the GType for the type CellAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CellAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellAccessible)
	return gtype
}

// CellAccessibleOverrider contains methods that are overridable.
type CellAccessibleOverrider interface {
	// The function takes the following parameters:
	//
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	_ [0]func() // equal guard
	Accessible

	*coreglib.Object
	atk.Action
	atk.Component
	atk.ObjectClass
	atk.TableCell
}

var (
	_ coreglib.Objector = (*CellAccessible)(nil)
)

func classInitCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "CellAccessibleClass")

	if _, ok := goval.(interface{ UpdateCache(emitSignal bool) }); ok {
		o := pclass.StructFieldOffset("update_cache")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_CellAccessibleClass_update_cache)
	}
}

//export _gotk4_gtk3_CellAccessibleClass_update_cache
func _gotk4_gtk3_CellAccessibleClass_update_cache(arg0 *C.void, arg1 C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ UpdateCache(emitSignal bool) })

	var _emitSignal bool // out

	if arg1 != 0 {
		_emitSignal = true
	}

	iface.UpdateCache(_emitSignal)
}

func wrapCellAccessible(obj *coreglib.Object) *CellAccessible {
	return &CellAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		TableCell: atk.TableCell{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalCellAccessible(p uintptr) (interface{}, error) {
	return wrapCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
