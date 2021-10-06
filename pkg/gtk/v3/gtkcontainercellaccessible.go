// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_container_cell_accessible_get_type()), F: marshalContainerCellAccessibler},
	})
}

type ContainerCellAccessible struct {
	CellAccessible
}

func wrapContainerCellAccessible(obj *externglib.Object) *ContainerCellAccessible {
	return &ContainerCellAccessible{
		CellAccessible: CellAccessible{
			Accessible: Accessible{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Action: atk.Action{
				Object: obj,
			},
			Component: atk.Component{
				Object: obj,
			},
			TableCell: atk.TableCell{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Object: obj,
		},
	}
}

func marshalContainerCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapContainerCellAccessible(obj), nil
}

func NewContainerCellAccessible() *ContainerCellAccessible {
	var _cret *C.GtkContainerCellAccessible // in

	_cret = C.gtk_container_cell_accessible_new()

	var _containerCellAccessible *ContainerCellAccessible // out

	_containerCellAccessible = wrapContainerCellAccessible(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _containerCellAccessible
}

//
// The function takes the following parameters:
//

//
func (container *ContainerCellAccessible) AddChild(child *CellAccessible) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(container.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(child.Native()))

	C.gtk_container_cell_accessible_add_child(_arg0, _arg1)
	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}

// Children: get a list of children.
func (container *ContainerCellAccessible) Children() []CellAccessible {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _cret *C.GList                      // in

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(container.Native()))

	_cret = C.gtk_container_cell_accessible_get_children(_arg0)
	runtime.KeepAlive(container)

	var _list []CellAccessible // out

	_list = make([]CellAccessible, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkCellAccessible)(v)
		var dst CellAccessible // out
		dst = *wrapCellAccessible(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

//
// The function takes the following parameters:
//

//
func (container *ContainerCellAccessible) RemoveChild(child *CellAccessible) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(container.Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(child.Native()))

	C.gtk_container_cell_accessible_remove_child(_arg0, _arg1)
	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}
