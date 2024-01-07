// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

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
	GTypeRecentFilterFlags = coreglib.Type(girepository.MustFind("Gtk", "RecentFilterFlags").RegisteredGType())
	GTypeRecentFilter      = coreglib.Type(girepository.MustFind("Gtk", "RecentFilter").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentFilterFlags, F: marshalRecentFilterFlags},
		coreglib.TypeMarshaler{T: GTypeRecentFilter, F: marshalRecentFilter},
	})
}

// RecentFilterFlags: these flags indicate what parts of a RecentFilterInfo
// struct are filled or need to be filled.
type RecentFilterFlags C.guint

const (
	// RecentFilterURI: URI of the file being tested.
	RecentFilterURI RecentFilterFlags = 0b1
	// RecentFilterDisplayName: string that will be used to display the file in
	// the recent chooser.
	RecentFilterDisplayName RecentFilterFlags = 0b10
	// RecentFilterMIMEType: mime type of the file.
	RecentFilterMIMEType RecentFilterFlags = 0b100
	// RecentFilterApplication: list of applications that have registered the
	// file.
	RecentFilterApplication RecentFilterFlags = 0b1000
	// RecentFilterGroup groups to which the file belongs to.
	RecentFilterGroup RecentFilterFlags = 0b10000
	// RecentFilterAge: number of days elapsed since the file has been
	// registered.
	RecentFilterAge RecentFilterFlags = 0b100000
)

func marshalRecentFilterFlags(p uintptr) (interface{}, error) {
	return RecentFilterFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for RecentFilterFlags.
func (r RecentFilterFlags) String() string {
	if r == 0 {
		return "RecentFilterFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(118)

	for r != 0 {
		next := r & (r - 1)
		bit := r - next

		switch bit {
		case RecentFilterURI:
			builder.WriteString("URI|")
		case RecentFilterDisplayName:
			builder.WriteString("DisplayName|")
		case RecentFilterMIMEType:
			builder.WriteString("MIMEType|")
		case RecentFilterApplication:
			builder.WriteString("Application|")
		case RecentFilterGroup:
			builder.WriteString("Group|")
		case RecentFilterAge:
			builder.WriteString("Age|")
		default:
			builder.WriteString(fmt.Sprintf("RecentFilterFlags(0b%b)|", bit))
		}

		r = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if r contains other.
func (r RecentFilterFlags) Has(other RecentFilterFlags) bool {
	return (r & other) == other
}

// RecentFilterFunc: type of function that is used with custom filters, see
// gtk_recent_filter_add_custom().
type RecentFilterFunc func(filterInfo *RecentFilterInfo) (ok bool)

// RecentFilter can be used to restrict the files being shown in a
// RecentChooser. Files can be filtered based on their name (with
// gtk_recent_filter_add_pattern()), on their mime type (with
// gtk_file_filter_add_mime_type()), on the application that has registered them
// (with gtk_recent_filter_add_application()), or by a custom filter function
// (with gtk_recent_filter_add_custom()).
//
// Filtering by mime type handles aliasing and subclassing of mime types; e.g. a
// filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that RecentFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, filters are used by adding them to a RecentChooser, see
// gtk_recent_chooser_add_filter(), but it is also possible to manually use a
// filter on a file with gtk_recent_filter_filter().
//
// Recently used files are supported since GTK+ 2.10.
//
//
// GtkRecentFilter as GtkBuildable
//
// The GtkRecentFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types>, <patterns> and <applications> elements
// and listing the rules within. Specifying a <mime-type>, <pattern> or
// <application> has the same effect as calling
// gtk_recent_filter_add_mime_type(), gtk_recent_filter_add_pattern() or
// gtk_recent_filter_add_application().
//
// An example of a UI definition fragment specifying GtkRecentFilter rules:
//
//    <object class="GtkRecentFilter">
//      <mime-types>
//        <mime-type>text/plain</mime-type>
//        <mime-type>image/png</mime-type>
//      </mime-types>
//      <patterns>
//        <pattern>*.txt</pattern>
//        <pattern>*.png</pattern>
//      </patterns>
//      <applications>
//        <application>gimp</application>
//        <application>gedit</application>
//        <application>glade</application>
//      </applications>
//    </object>.
type RecentFilter struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned

	*coreglib.Object
	Buildable
}

var (
	_ coreglib.Objector = (*RecentFilter)(nil)
)

func wrapRecentFilter(obj *coreglib.Object) *RecentFilter {
	return &RecentFilter{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalRecentFilter(p uintptr) (interface{}, error) {
	return wrapRecentFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RecentFilterInfo struct is used to pass information about the tested file to
// gtk_recent_filter_filter().
//
// An instance of this type is always passed by reference.
type RecentFilterInfo struct {
	*recentFilterInfo
}

// recentFilterInfo is the struct that's finalized.
type recentFilterInfo struct {
	native unsafe.Pointer
}

var GIRInfoRecentFilterInfo = girepository.MustFind("Gtk", "RecentFilterInfo")
