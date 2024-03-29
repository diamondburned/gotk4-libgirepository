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
// extern void _gotk4_gtk3_RecentChooser_ConnectSelectionChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_RecentChooser_ConnectItemActivated(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeRecentChooser = coreglib.Type(girepository.MustFind("Gtk", "RecentChooser").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentChooser, F: marshalRecentChooser},
	})
}

type RecentSortFunc func(a, b *RecentInfo) (gint int)

// RecentChooserOverrider contains methods that are overridable.
type RecentChooserOverrider interface {
}

// RecentChooser is an interface that can be implemented by widgets displaying
// the list of recently used files. In GTK+, the main objects that implement
// this interface are RecentChooserWidget, RecentChooserDialog and
// RecentChooserMenu.
//
// Recently used files are supported since GTK+ 2.10.
//
// RecentChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type RecentChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*RecentChooser)(nil)
)

// RecentChooserer describes RecentChooser's interface methods.
type RecentChooserer interface {
	coreglib.Objector

	baseRecentChooser() *RecentChooser
}

var _ RecentChooserer = (*RecentChooser)(nil)

func ifaceInitRecentChooserer(gifacePtr, data C.gpointer) {
}

func wrapRecentChooser(obj *coreglib.Object) *RecentChooser {
	return &RecentChooser{
		Object: obj,
	}
}

func marshalRecentChooser(p uintptr) (interface{}, error) {
	return wrapRecentChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *RecentChooser) baseRecentChooser() *RecentChooser {
	return v
}

// BaseRecentChooser returns the underlying base object.
func BaseRecentChooser(obj RecentChooserer) *RecentChooser {
	return obj.baseRecentChooser()
}

// ConnectItemActivated: this signal is emitted when the user "activates" a
// recent item in the recent chooser. This can happen by double-clicking on an
// item in the recently used resources list, or by pressing Enter.
func (v *RecentChooser) ConnectItemActivated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "item-activated", false, unsafe.Pointer(C._gotk4_gtk3_RecentChooser_ConnectItemActivated), f)
}

// ConnectSelectionChanged: this signal is emitted when there is a change in the
// set of selected recently used resources. This can happen when a user modifies
// the selection with the mouse or the keyboard, or when explicitly calling
// functions to change the selection.
func (v *RecentChooser) ConnectSelectionChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "selection-changed", false, unsafe.Pointer(C._gotk4_gtk3_RecentChooser_ConnectSelectionChanged), f)
}

// RecentChooserIface: instance of this type is always passed by reference.
type RecentChooserIface struct {
	*recentChooserIface
}

// recentChooserIface is the struct that's finalized.
type recentChooserIface struct {
	native unsafe.Pointer
}

var GIRInfoRecentChooserIface = girepository.MustFind("Gtk", "RecentChooserIface")
