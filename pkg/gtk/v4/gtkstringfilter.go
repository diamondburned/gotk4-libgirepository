// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeStringFilterMatchMode = coreglib.Type(girepository.MustFind("Gtk", "StringFilterMatchMode").RegisteredGType())
	GTypeStringFilter          = coreglib.Type(girepository.MustFind("Gtk", "StringFilter").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStringFilterMatchMode, F: marshalStringFilterMatchMode},
		coreglib.TypeMarshaler{T: GTypeStringFilter, F: marshalStringFilter},
	})
}

// StringFilterMatchMode specifies how search strings are matched inside text.
type StringFilterMatchMode C.gint

const (
	// StringFilterMatchModeExact: search string and text must match exactly.
	StringFilterMatchModeExact StringFilterMatchMode = iota
	// StringFilterMatchModeSubstring: search string must be contained as a
	// substring inside the text.
	StringFilterMatchModeSubstring
	// StringFilterMatchModePrefix: text must begin with the search string.
	StringFilterMatchModePrefix
)

func marshalStringFilterMatchMode(p uintptr) (interface{}, error) {
	return StringFilterMatchMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for StringFilterMatchMode.
func (s StringFilterMatchMode) String() string {
	switch s {
	case StringFilterMatchModeExact:
		return "Exact"
	case StringFilterMatchModeSubstring:
		return "Substring"
	case StringFilterMatchModePrefix:
		return "Prefix"
	default:
		return fmt.Sprintf("StringFilterMatchMode(%d)", s)
	}
}

// StringFilterOverrides contains methods that are overridable.
type StringFilterOverrides struct {
}

func defaultStringFilterOverrides(v *StringFilter) StringFilterOverrides {
	return StringFilterOverrides{}
}

// StringFilter: GtkStringFilter determines whether to include items by
// comparing strings to a fixed search term.
//
// The strings are obtained from the items by evaluating a GtkExpression set
// with gtk.StringFilter.SetExpression(), and they are compared against a search
// term set with gtk.StringFilter.SetSearch().
//
// GtkStringFilter has several different modes of comparison - it can match the
// whole string, just a prefix, or any substring. Use
// gtk.StringFilter.SetMatchMode() choose a mode.
//
// It is also possible to make case-insensitive comparisons, with
// gtk.StringFilter.SetIgnoreCase().
type StringFilter struct {
	_ [0]func() // equal guard
	Filter
}

var (
	_ coreglib.Objector = (*StringFilter)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StringFilter, *StringFilterClass, StringFilterOverrides](
		GTypeStringFilter,
		initStringFilterClass,
		wrapStringFilter,
		defaultStringFilterOverrides,
	)
}

func initStringFilterClass(gclass unsafe.Pointer, overrides StringFilterOverrides, classInitFunc func(*StringFilterClass)) {
	if classInitFunc != nil {
		class := (*StringFilterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStringFilter(obj *coreglib.Object) *StringFilter {
	return &StringFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalStringFilter(p uintptr) (interface{}, error) {
	return wrapStringFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StringFilterClass: instance of this type is always passed by reference.
type StringFilterClass struct {
	*stringFilterClass
}

// stringFilterClass is the struct that's finalized.
type stringFilterClass struct {
	native unsafe.Pointer
}

var GIRInfoStringFilterClass = girepository.MustFind("Gtk", "StringFilterClass")
