// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeMultiSorter returns the GType for the type MultiSorter.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMultiSorter() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MultiSorter").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMultiSorter)
	return gtype
}

// MultiSorterOverrider contains methods that are overridable.
type MultiSorterOverrider interface {
}

// MultiSorter: GtkMultiSorter combines multiple sorters by trying them in turn.
//
// If the first sorter compares two items as equal, the second is tried next,
// and so on.
type MultiSorter struct {
	_ [0]func() // equal guard
	Sorter

	*coreglib.Object
	gio.ListModel
	Buildable
}

var (
	_ coreglib.Objector = (*MultiSorter)(nil)
)

func classInitMultiSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMultiSorter(obj *coreglib.Object) *MultiSorter {
	return &MultiSorter{
		Sorter: Sorter{
			Object: obj,
		},
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalMultiSorter(p uintptr) (interface{}, error) {
	return wrapMultiSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultiSorter creates a new multi sorter.
//
// This sorter compares items by trying each of the sorters in turn, until one
// returns non-zero. In particular, if no sorter has been added to it, it will
// always compare items as equal.
//
// The function returns the following values:
//
//    - multiSorter: new GtkMultiSorter.
//
func NewMultiSorter() *MultiSorter {
	_gret := girepository.MustFind("Gtk", "MultiSorter").InvokeMethod("new_MultiSorter", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _multiSorter *MultiSorter // out

	_multiSorter = wrapMultiSorter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multiSorter
}

// Append: add sorter to self to use for sorting at the end.
//
// self will consult all existing sorters before it will sort with the given
// sorter.
//
// The function takes the following parameters:
//
//    - sorter to add.
//
func (self *MultiSorter) Append(sorter *Sorter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sorter).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(sorter).Native()))

	girepository.MustFind("Gtk", "MultiSorter").InvokeMethod("append", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(sorter)
}

// Remove removes the sorter at the given position from the list of sorter used
// by self.
//
// If position is larger than the number of sorters, nothing happens.
//
// The function takes the following parameters:
//
//    - position of sorter to remove.
//
func (self *MultiSorter) Remove(position uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	girepository.MustFind("Gtk", "MultiSorter").InvokeMethod("remove", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
