// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeShortcutManager = coreglib.Type(girepository.MustFind("Gtk", "ShortcutManager").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutManager, F: marshalShortcutManager},
	})
}

// ShortcutManagerOverrider contains methods that are overridable.
type ShortcutManagerOverrider interface {
}

// ShortcutManager: GtkShortcutManager interface is used to implement shortcut
// scopes.
//
// This is important for gtk.Native widgets that have their own surface, since
// the event controllers that are used to implement managed and global scopes
// are limited to the same native.
//
// Examples for widgets implementing GtkShortcutManager are gtk.Window and
// gtk.Popover.
//
// Every widget that implements GtkShortcutManager will be used as a
// GTK_SHORTCUT_SCOPE_MANAGED.
//
// ShortcutManager wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ShortcutManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutManager)(nil)
)

// ShortcutManagerer describes ShortcutManager's interface methods.
type ShortcutManagerer interface {
	coreglib.Objector

	baseShortcutManager() *ShortcutManager
}

var _ ShortcutManagerer = (*ShortcutManager)(nil)

func ifaceInitShortcutManagerer(gifacePtr, data C.gpointer) {
}

func wrapShortcutManager(obj *coreglib.Object) *ShortcutManager {
	return &ShortcutManager{
		Object: obj,
	}
}

func marshalShortcutManager(p uintptr) (interface{}, error) {
	return wrapShortcutManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ShortcutManager) baseShortcutManager() *ShortcutManager {
	return v
}

// BaseShortcutManager returns the underlying base object.
func BaseShortcutManager(obj ShortcutManagerer) *ShortcutManager {
	return obj.baseShortcutManager()
}

// ShortcutManagerInterface: list of functions that can be implemented for the
// ShortcutManager interface.
//
// Note that no function is mandatory to implement, the default implementation
// will work fine.
//
// An instance of this type is always passed by reference.
type ShortcutManagerInterface struct {
	*shortcutManagerInterface
}

// shortcutManagerInterface is the struct that's finalized.
type shortcutManagerInterface struct {
	native unsafe.Pointer
}

var GIRInfoShortcutManagerInterface = girepository.MustFind("Gtk", "ShortcutManagerInterface")
