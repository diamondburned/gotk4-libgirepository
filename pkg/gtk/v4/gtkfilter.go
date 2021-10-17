// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_filter_change_get_type()), F: marshalFilterChange},
		{T: externglib.Type(C.gtk_filter_match_get_type()), F: marshalFilterMatch},
		{T: externglib.Type(C.gtk_filter_get_type()), F: marshalFilterer},
	})
}

// FilterChange describes changes in a filter in more detail and allows objects
// using the filter to optimize refiltering items.
//
// If you are writing an implementation and are not sure which value to pass,
// GTK_FILTER_CHANGE_DIFFERENT is always a correct choice.
type FilterChange int

const (
	// FilterChangeDifferent: filter change cannot be described with any of the
	// other enumeration values.
	FilterChangeDifferent FilterChange = iota
	// FilterChangeLessStrict: filter is less strict than it was before: All
	// items that it used to return TRUE for still return TRUE, others now may,
	// too.
	FilterChangeLessStrict
	// FilterChangeMoreStrict: filter is more strict than it was before: All
	// items that it used to return FALSE for still return FALSE, others now
	// may, too.
	FilterChangeMoreStrict
)

func marshalFilterChange(p uintptr) (interface{}, error) {
	return FilterChange(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FilterChange.
func (f FilterChange) String() string {
	switch f {
	case FilterChangeDifferent:
		return "Different"
	case FilterChangeLessStrict:
		return "LessStrict"
	case FilterChangeMoreStrict:
		return "MoreStrict"
	default:
		return fmt.Sprintf("FilterChange(%d)", f)
	}
}

// FilterMatch describes the known strictness of a filter.
//
// Note that for filters where the strictness is not known,
// GTK_FILTER_MATCH_SOME is always an acceptable value, even if a filter does
// match all or no items.
type FilterMatch int

const (
	// FilterMatchSome: filter matches some items, gtk_filter_match() may return
	// TRUE or FALSE.
	FilterMatchSome FilterMatch = iota
	// FilterMatchNone: filter does not match any item, gtk_filter_match() will
	// always return FALSE.
	FilterMatchNone
	// FilterMatchAll: filter matches all items, gtk_filter_match() will alays
	// return TRUE.
	FilterMatchAll
)

func marshalFilterMatch(p uintptr) (interface{}, error) {
	return FilterMatch(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FilterMatch.
func (f FilterMatch) String() string {
	switch f {
	case FilterMatchSome:
		return "Some"
	case FilterMatchNone:
		return "None"
	case FilterMatchAll:
		return "All"
	default:
		return fmt.Sprintf("FilterMatch(%d)", f)
	}
}

// FilterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FilterOverrider interface {
	// Strictness gets the known strictness of filters. If the strictness is not
	// known, GTK_FILTER_MATCH_SOME is returned.
	//
	// This value may change after emission of the Filter::changed signal.
	//
	// This function is meant purely for optimization purposes, filters can
	// choose to omit implementing it, but FilterListModel uses it.
	Strictness() FilterMatch
	// Match checks if the given item is matched by the filter or not.
	Match(item *externglib.Object) bool
}

// Filter: GtkFilter object describes the filtering to be performed by a
// GtkFilterListModel.
//
// The model will use the filter to determine if it should include items or not
// by calling gtk.Filter.Match() for each item and only keeping the ones that
// the function returns TRUE for.
//
// Filters may change what items they match through their lifetime. In that
// case, they will emit the gtk.Filter::changed signal to notify that previous
// filter results are no longer valid and that items should be checked again via
// gtk.Filter.Match().
//
// GTK provides various pre-made filter implementations for common filtering
// operations. These filters often include properties that can be linked to
// various widgets to easily allow searches.
//
// However, in particular for large lists or complex search methods, it is also
// possible to subclass Filter and provide one's own filter.
type Filter struct {
	*externglib.Object
}

func wrapFilter(obj *externglib.Object) *Filter {
	return &Filter{
		Object: obj,
	}
}

func marshalFilterer(p uintptr) (interface{}, error) {
	return wrapFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Changed emits the Filter::changed signal to notify all users of the filter
// that the filter changed. Users of the filter should then check items again
// via gtk_filter_match().
//
// Depending on the change parameter, not all items need to be changed, but only
// some. Refer to the FilterChange documentation for details.
//
// This function is intended for implementors of Filter subclasses and should
// not be called from other functions.
//
// The function takes the following parameters:
//
//    - change: how the filter changed.
//
func (self *Filter) Changed(change FilterChange) {
	var _arg0 *C.GtkFilter      // out
	var _arg1 C.GtkFilterChange // out

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkFilterChange(change)

	C.gtk_filter_changed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(change)
}

// Strictness gets the known strictness of filters. If the strictness is not
// known, GTK_FILTER_MATCH_SOME is returned.
//
// This value may change after emission of the Filter::changed signal.
//
// This function is meant purely for optimization purposes, filters can choose
// to omit implementing it, but FilterListModel uses it.
func (self *Filter) Strictness() FilterMatch {
	var _arg0 *C.GtkFilter     // out
	var _cret C.GtkFilterMatch // in

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_filter_get_strictness(_arg0)
	runtime.KeepAlive(self)

	var _filterMatch FilterMatch // out

	_filterMatch = FilterMatch(_cret)

	return _filterMatch
}

// Match checks if the given item is matched by the filter or not.
//
// The function takes the following parameters:
//
//    - item to check.
//
func (self *Filter) Match(item *externglib.Object) bool {
	var _arg0 *C.GtkFilter // out
	var _arg1 C.gpointer   // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item.Native()))

	_cret = C.gtk_filter_match(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConnectChanged: emitted whenever the filter changed.
//
// Users of the filter should then check items again via gtk.Filter.Match().
//
// GtkFilterListModel handles this signal automatically.
//
// Depending on the change parameter, not all items need to be checked, but only
// some. Refer to the gtk.FilterChange documentation for details.
func (self *Filter) ConnectChanged(f func(change FilterChange)) externglib.SignalHandle {
	return self.Connect("changed", f)
}
