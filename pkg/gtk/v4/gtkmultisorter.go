// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_multi_sorter_get_type()), F: marshalMultiSorter},
	})
}

// MultiSorter: `GtkMultiSorter` combines multiple sorters by trying them in
// turn.
//
// If the first sorter compares two items as equal, the second is tried next,
// and so on.
type MultiSorter interface {
	gextras.Objector

	// AsSorter casts the class to the Sorter interface.
	AsSorter() Sorter
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// Changed emits the [signal@Gtk.Sorter::changed] signal to notify all users
	// of the sorter that it has changed.
	//
	// Users of the sorter should then update the sort order via
	// gtk_sorter_compare().
	//
	// Depending on the @change parameter, it may be possible to update the sort
	// order without a full resorting. Refer to the [enum@Gtk.SorterChange]
	// documentation for details.
	//
	// This function is intended for implementors of `GtkSorter` subclasses and
	// should not be called from other functions.
	//
	// This method is inherited from Sorter
	Changed(change SorterChange)
	// Compare compares two given items according to the sort order implemented
	// by the sorter.
	//
	// Sorters implement a partial order:
	//
	// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
	// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
	// then a ≤ c
	//
	// The sorter may signal it conforms to additional constraints via the
	// return value of [method@Gtk.Sorter.get_order].
	//
	// This method is inherited from Sorter
	Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering
	// GetOrder gets the order that @self conforms to.
	//
	// See [enum@Gtk.SorterOrder] for details of the possible return values.
	//
	// This function is intended to allow optimizations.
	//
	// This method is inherited from Sorter
	GetOrder() SorterOrder
	// GetBuildableID gets the ID of the @buildable object.
	//
	// `GtkBuilder` sets the name based on the ID attribute of the <object> tag
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetBuildableID() string

	// Append: add @sorter to @self to use for sorting at the end.
	//
	// @self will consult all existing sorters before it will sort with the
	// given @sorter.
	Append(sorter Sorter)
	// Remove removes the sorter at the given @position from the list of sorter
	// used by @self.
	//
	// If @position is larger than the number of sorters, nothing happens.
	Remove(position uint)
}

// multiSorter implements the MultiSorter interface.
type multiSorter struct {
	*externglib.Object
}

var _ MultiSorter = (*multiSorter)(nil)

// WrapMultiSorter wraps a GObject to a type that implements
// interface MultiSorter. It is primarily used internally.
func WrapMultiSorter(obj *externglib.Object) MultiSorter {
	return multiSorter{obj}
}

func marshalMultiSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiSorter(obj), nil
}

// NewMultiSorter creates a new multi sorter.
//
// This sorter compares items by trying each of the sorters in turn, until one
// returns non-zero. In particular, if no sorter has been added to it, it will
// always compare items as equal.
func NewMultiSorter() MultiSorter {
	var _cret *C.GtkMultiSorter // in

	_cret = C.gtk_multi_sorter_new()

	var _multiSorter MultiSorter // out

	_multiSorter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MultiSorter)

	return _multiSorter
}

func (m multiSorter) AsSorter() Sorter {
	return WrapSorter(gextras.InternObject(m))
}

func (m multiSorter) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(m))
}

func (s multiSorter) Changed(change SorterChange) {
	WrapSorter(gextras.InternObject(s)).Changed(change)
}

func (s multiSorter) Compare(item1 gextras.Objector, item2 gextras.Objector) Ordering {
	return WrapSorter(gextras.InternObject(s)).Compare(item1, item2)
}

func (s multiSorter) GetOrder() SorterOrder {
	return WrapSorter(gextras.InternObject(s)).GetOrder()
}

func (b multiSorter) GetBuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).GetBuildableID()
}

func (s multiSorter) Append(sorter Sorter) {
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 *C.GtkSorter      // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_multi_sorter_append(_arg0, _arg1)
}

func (s multiSorter) Remove(position uint) {
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_sorter_remove(_arg0, _arg1)
}
