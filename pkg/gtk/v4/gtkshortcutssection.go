// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_section_get_type()), F: marshalShortcutsSection},
	})
}

// ShortcutsSection: a `GtkShortcutsSection` collects all the keyboard shortcuts
// and gestures for a major application mode.
//
// If your application needs multiple sections, you should give each section a
// unique [property@Gtk.ShortcutsSection:section-name] and a
// [property@Gtk.ShortcutsSection:title] that can be shown in the section
// selector of the [class@Gtk.ShortcutsWindow].
//
// The [property@Gtk.ShortcutsSection:max-height] property can be used to
// influence how the groups in the section are distributed over pages and
// columns.
//
// This widget is only meant to be used with [class@Gtk.ShortcutsWindow].
type ShortcutsSection interface {
	Box
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

// shortcutsSection implements the ShortcutsSection class.
type shortcutsSection struct {
	Box
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ ShortcutsSection = (*shortcutsSection)(nil)

// WrapShortcutsSection wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsSection(obj *externglib.Object) ShortcutsSection {
	return shortcutsSection{
		Box:              WrapBox(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalShortcutsSection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsSection(obj), nil
}
