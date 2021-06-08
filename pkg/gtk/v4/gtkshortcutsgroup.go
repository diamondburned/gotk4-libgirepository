// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_group_get_type()), F: marshalShortcutsGroup},
	})
}

// ShortcutsGroup: a `GtkShortcutsGroup` represents a group of related keyboard
// shortcuts or gestures.
//
// The group has a title. It may optionally be associated with a view of the
// application, which can be used to show only relevant shortcuts depending on
// the application context.
//
// This widget is only meant to be used with [class@Gtk.ShortcutsWindow].
type ShortcutsGroup interface {
	Box
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

// shortcutsGroup implements the ShortcutsGroup interface.
type shortcutsGroup struct {
	Box
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ ShortcutsGroup = (*shortcutsGroup)(nil)

// WrapShortcutsGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsGroup(obj *externglib.Object) ShortcutsGroup {
	return ShortcutsGroup{
		Box:              WrapBox(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsGroup(obj), nil
}
