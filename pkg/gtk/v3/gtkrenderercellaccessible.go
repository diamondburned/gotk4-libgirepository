// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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

// GTypeRendererCellAccessible returns the GType for the type RendererCellAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRendererCellAccessible() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "RendererCellAccessible").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRendererCellAccessible)
	return gtype
}

// RendererCellAccessibleOverrider contains methods that are overridable.
type RendererCellAccessibleOverrider interface {
}

type RendererCellAccessible struct {
	_ [0]func() // equal guard
	CellAccessible
}

var (
	_ coreglib.Objector = (*RendererCellAccessible)(nil)
)

func classInitRendererCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRendererCellAccessible(obj *coreglib.Object) *RendererCellAccessible {
	return &RendererCellAccessible{
		CellAccessible: CellAccessible{
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
		},
	}
}

func marshalRendererCellAccessible(p uintptr) (interface{}, error) {
	return wrapRendererCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewRendererCellAccessible(renderer CellRendererer) *RendererCellAccessible {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_gret := girepository.MustFind("Gtk", "RendererCellAccessible").InvokeMethod("new_RendererCellAccessible", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(renderer)

	var _rendererCellAccessible *RendererCellAccessible // out

	_rendererCellAccessible = wrapRendererCellAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _rendererCellAccessible
}
