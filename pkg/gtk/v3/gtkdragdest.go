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
	GTypeDestDefaults = coreglib.Type(girepository.MustFind("Gtk", "DestDefaults").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDestDefaults, F: marshalDestDefaults},
	})
}

// DestDefaults enumeration specifies the various types of action that will be
// taken on behalf of the user for a drag destination site.
type DestDefaults C.guint

const (
	// DestDefaultMotion: if set for a widget, GTK+, during a drag over this
	// widget will check if the drag matches this widget’s list of possible
	// targets and actions. GTK+ will then call gdk_drag_status() as
	// appropriate.
	DestDefaultMotion DestDefaults = 0b1
	// DestDefaultHighlight: if set for a widget, GTK+ will draw a highlight on
	// this widget as long as a drag is over this widget and the widget drag
	// format and action are acceptable.
	DestDefaultHighlight DestDefaults = 0b10
	// DestDefaultDrop: if set for a widget, when a drop occurs, GTK+ will will
	// check if the drag matches this widget’s list of possible targets and
	// actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the
	// widget. Whether or not the drop is successful, GTK+ will call
	// gtk_drag_finish(). If the action was a move, then if the drag was
	// successful, then TRUE will be passed for the delete parameter to
	// gtk_drag_finish().
	DestDefaultDrop DestDefaults = 0b100
	// DestDefaultAll: if set, specifies that all default actions should be
	// taken.
	DestDefaultAll DestDefaults = 0b111
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	return DestDefaults(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DestDefaults.
func (d DestDefaults) String() string {
	if d == 0 {
		return "DestDefaults(0)"
	}

	var builder strings.Builder
	builder.Grow(69)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DestDefaultMotion:
			builder.WriteString("Motion|")
		case DestDefaultHighlight:
			builder.WriteString("Highlight|")
		case DestDefaultDrop:
			builder.WriteString("Drop|")
		case DestDefaultAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("DestDefaults(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DestDefaults) Has(other DestDefaults) bool {
	return (d & other) == other
}
