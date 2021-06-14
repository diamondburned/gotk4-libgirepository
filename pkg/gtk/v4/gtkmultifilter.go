// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_any_filter_get_type()), F: marshalAnyFilter},
		{T: externglib.Type(C.gtk_every_filter_get_type()), F: marshalEveryFilter},
		{T: externglib.Type(C.gtk_multi_filter_get_type()), F: marshalMultiFilter},
	})
}

// AnyFilter: `GtkAnyFilter` matches an item when at least one of its filters
// matches.
//
// To add filters to a `GtkAnyFilter`, use [method@Gtk.MultiFilter.append].
type AnyFilter interface {
	MultiFilter
	gio.ListModel
	Buildable
}

// anyFilter implements the AnyFilter class.
type anyFilter struct {
	MultiFilter
	gio.ListModel
	Buildable
}

var _ AnyFilter = (*anyFilter)(nil)

// WrapAnyFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapAnyFilter(obj *externglib.Object) AnyFilter {
	return anyFilter{
		MultiFilter:   WrapMultiFilter(obj),
		gio.ListModel: gio.WrapListModel(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalAnyFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAnyFilter(obj), nil
}

// EveryFilter: `GtkEveryFilter` matches an item when each of its filters
// matches.
//
// To add filters to a `GtkEveryFilter`, use [method@Gtk.MultiFilter.append].
type EveryFilter interface {
	MultiFilter
	gio.ListModel
	Buildable
}

// everyFilter implements the EveryFilter class.
type everyFilter struct {
	MultiFilter
	gio.ListModel
	Buildable
}

var _ EveryFilter = (*everyFilter)(nil)

// WrapEveryFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapEveryFilter(obj *externglib.Object) EveryFilter {
	return everyFilter{
		MultiFilter:   WrapMultiFilter(obj),
		gio.ListModel: gio.WrapListModel(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalEveryFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEveryFilter(obj), nil
}

// MultiFilter: `GtkMultiFilter` is the base class for filters that combine
// multiple filters.
type MultiFilter interface {
	Filter
	gio.ListModel
	Buildable

	// Append adds a @filter to @self to use for matching.
	Append(filter Filter)
	// Remove removes the filter at the given @position from the list of filters
	// used by @self.
	//
	// If @position is larger than the number of filters, nothing happens and
	// the function returns.
	Remove(position uint)
}

// multiFilter implements the MultiFilter class.
type multiFilter struct {
	Filter
	gio.ListModel
	Buildable
}

var _ MultiFilter = (*multiFilter)(nil)

// WrapMultiFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapMultiFilter(obj *externglib.Object) MultiFilter {
	return multiFilter{
		Filter:        WrapFilter(obj),
		gio.ListModel: gio.WrapListModel(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalMultiFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiFilter(obj), nil
}

// Append adds a @filter to @self to use for matching.
func (s multiFilter) Append(filter Filter) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 *C.GtkFilter      // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_multi_filter_append(_arg0, _arg1)
}

// Remove removes the filter at the given @position from the list of filters
// used by @self.
//
// If @position is larger than the number of filters, nothing happens and
// the function returns.
func (s multiFilter) Remove(position uint) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_filter_remove(_arg0, _arg1)
}
