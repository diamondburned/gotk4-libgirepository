// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_snapshot_get_type()), F: marshalSnapshot},
	})
}

// Snapshot: base type for snapshot operations.
//
// The subclass of `GdkSnapshot` used by GTK is [class@Gtk.Snapshot].
type Snapshot interface {
	gextras.Objector
}

// snapshot implements the Snapshot class.
type snapshot struct {
	gextras.Objector
}

var _ Snapshot = (*snapshot)(nil)

// WrapSnapshot wraps a GObject to the right type. It is
// primarily used internally.
func WrapSnapshot(obj *externglib.Object) Snapshot {
	return snapshot{
		Objector: obj,
	}
}

func marshalSnapshot(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSnapshot(obj), nil
}
