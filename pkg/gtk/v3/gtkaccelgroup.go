// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
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
	GTypeAccelFlags = coreglib.Type(girepository.MustFind("Gtk", "AccelFlags").RegisteredGType())
	GTypeAccelGroup = coreglib.Type(girepository.MustFind("Gtk", "AccelGroup").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAccelFlags, F: marshalAccelFlags},
		coreglib.TypeMarshaler{T: GTypeAccelGroup, F: marshalAccelGroup},
	})
}

// AccelFlags: accelerator flags used with gtk_accel_group_connect().
type AccelFlags C.guint

const (
	// AccelVisible: accelerator is visible.
	AccelVisible AccelFlags = 0b1
	// AccelLocked: accelerator not removable.
	AccelLocked AccelFlags = 0b10
	// AccelMask: mask.
	AccelMask AccelFlags = 0b111
)

func marshalAccelFlags(p uintptr) (interface{}, error) {
	return AccelFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AccelFlags.
func (a AccelFlags) String() string {
	if a == 0 {
		return "AccelFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(34)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AccelVisible:
			builder.WriteString("Visible|")
		case AccelLocked:
			builder.WriteString("Locked|")
		case AccelMask:
			builder.WriteString("Mask|")
		default:
			builder.WriteString(fmt.Sprintf("AccelFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AccelFlags) Has(other AccelFlags) bool {
	return (a & other) == other
}

// AccelGroupOverrides contains methods that are overridable.
type AccelGroupOverrides struct {
}

func defaultAccelGroupOverrides(v *AccelGroup) AccelGroupOverrides {
	return AccelGroupOverrides{}
}

// AccelGroup represents a group of keyboard accelerators, typically attached to
// a toplevel Window (with gtk_window_add_accel_group()). Usually you won’t need
// to create a AccelGroup directly; instead, when using UIManager, GTK+
// automatically sets up the accelerators for your menus in the ui manager’s
// AccelGroup.
//
// Note that “accelerators” are different from “mnemonics”. Accelerators are
// shortcuts for activating a menu item; they appear alongside the menu item
// they’re a shortcut for. For example “Ctrl+Q” might appear alongside the
// “Quit” menu item. Mnemonics are shortcuts for GUI elements such as text
// entries or buttons; they appear as underlined characters. See
// gtk_label_new_with_mnemonic(). Menu items can have both accelerators and
// mnemonics, of course.
type AccelGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AccelGroup)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AccelGroup, *AccelGroupClass, AccelGroupOverrides](
		GTypeAccelGroup,
		initAccelGroupClass,
		wrapAccelGroup,
		defaultAccelGroupOverrides,
	)
}

func initAccelGroupClass(gclass unsafe.Pointer, overrides AccelGroupOverrides, classInitFunc func(*AccelGroupClass)) {
	if classInitFunc != nil {
		class := (*AccelGroupClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAccelGroup(obj *coreglib.Object) *AccelGroup {
	return &AccelGroup{
		Object: obj,
	}
}

func marshalAccelGroup(p uintptr) (interface{}, error) {
	return wrapAccelGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AccelGroupClass: instance of this type is always passed by reference.
type AccelGroupClass struct {
	*accelGroupClass
}

// accelGroupClass is the struct that's finalized.
type accelGroupClass struct {
	native unsafe.Pointer
}

var GIRInfoAccelGroupClass = girepository.MustFind("Gtk", "AccelGroupClass")

// AccelGroupEntry: instance of this type is always passed by reference.
type AccelGroupEntry struct {
	*accelGroupEntry
}

// accelGroupEntry is the struct that's finalized.
type accelGroupEntry struct {
	native unsafe.Pointer
}

var GIRInfoAccelGroupEntry = girepository.MustFind("Gtk", "AccelGroupEntry")

// AccelKey: instance of this type is always passed by reference.
type AccelKey struct {
	*accelKey
}

// accelKey is the struct that's finalized.
type accelKey struct {
	native unsafe.Pointer
}

var GIRInfoAccelKey = girepository.MustFind("Gtk", "AccelKey")
