// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkcontainercellaccessible.go.
var GTypeContainerCellAccessible = coreglib.Type(C.gtk_container_cell_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeContainerCellAccessible, F: marshalContainerCellAccessible},
	})
}

// ContainerCellAccessibleOverrider contains methods that are overridable.
type ContainerCellAccessibleOverrider interface {
}

type ContainerCellAccessible struct {
	_ [0]func() // equal guard
	CellAccessible
}

var (
	_ coreglib.Objector = (*ContainerCellAccessible)(nil)
)

func classInitContainerCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapContainerCellAccessible(obj *coreglib.Object) *ContainerCellAccessible {
	return &ContainerCellAccessible{
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

func marshalContainerCellAccessible(p uintptr) (interface{}, error) {
	return wrapContainerCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func NewContainerCellAccessible() *ContainerCellAccessible {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ContainerCellAccessible").InvokeMethod("new_ContainerCellAccessible", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _containerCellAccessible *ContainerCellAccessible // out

	_containerCellAccessible = wrapContainerCellAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _containerCellAccessible
}

// The function takes the following parameters:
//
func (container *ContainerCellAccessible) AddChild(child *CellAccessible) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(container).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ContainerCellAccessible").InvokeMethod("add_child", _args[:], nil)

	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}

// Children: get a list of children.
//
// The function returns the following values:
//
func (container *ContainerCellAccessible) Children() []*CellAccessible {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(container).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ContainerCellAccessible").InvokeMethod("get_children", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(container)

	var _list []*CellAccessible // out

	_list = make([]*CellAccessible, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *CellAccessible // out
		dst = wrapCellAccessible(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// The function takes the following parameters:
//
func (container *ContainerCellAccessible) RemoveChild(child *CellAccessible) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(container).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "ContainerCellAccessible").InvokeMethod("remove_child", _args[:], nil)

	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}
