// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_multi_sorter_get_type()), F: marshalMultiSorterer},
	})
}

// MultiSorter: GtkMultiSorter combines multiple sorters by trying them in turn.
//
// If the first sorter compares two items as equal, the second is tried next,
// and so on.
type MultiSorter struct {
	Sorter

	gio.ListModel
	Buildable
	*externglib.Object
}

func wrapMultiSorter(obj *externglib.Object) *MultiSorter {
	return &MultiSorter{
		Sorter: Sorter{
			Object: obj,
		},
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalMultiSorterer(p uintptr) (interface{}, error) {
	return wrapMultiSorter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultiSorter creates a new multi sorter.
//
// This sorter compares items by trying each of the sorters in turn, until one
// returns non-zero. In particular, if no sorter has been added to it, it will
// always compare items as equal.
func NewMultiSorter() *MultiSorter {
	var _cret *C.GtkMultiSorter // in

	_cret = C.gtk_multi_sorter_new()

	var _multiSorter *MultiSorter // out

	_multiSorter = wrapMultiSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 *C.GtkSorter      // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))
	C.g_object_ref(C.gpointer(sorter.Native()))

	C.gtk_multi_sorter_append(_arg0, _arg1)
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
func (self *MultiSorter) Remove(position uint) {
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_sorter_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
