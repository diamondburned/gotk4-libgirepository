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
	GTypeScrollable = coreglib.Type(girepository.MustFind("Gtk", "Scrollable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrollable, F: marshalScrollable},
	})
}

// ScrollableOverrider contains methods that are overridable.
type ScrollableOverrider interface {
}

// Scrollable: GtkScrollable is an interface for widgets with native scrolling
// ability.
//
// To implement this interface you should override the
// gtk.Scrollable:hadjustment and gtk.Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ gtk.Adjustment:lower,
// gtk.Adjustment:upper, gtk.Adjustment:step-increment,
// gtk.Adjustment:page-increment and gtk.Adjustment:page-size properties and
// connect to the gtk.Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its GtkWidgetClass.size_allocate()
// function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the gtk.Adjustment::value-changed signal,
// the scrollable widget should scroll its contents.
//
// Scrollable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Scrollable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Scrollable)(nil)
)

// Scrollabler describes Scrollable's interface methods.
type Scrollabler interface {
	coreglib.Objector

	baseScrollable() *Scrollable
}

var _ Scrollabler = (*Scrollable)(nil)

func ifaceInitScrollabler(gifacePtr, data C.gpointer) {
}

func wrapScrollable(obj *coreglib.Object) *Scrollable {
	return &Scrollable{
		Object: obj,
	}
}

func marshalScrollable(p uintptr) (interface{}, error) {
	return wrapScrollable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Scrollable) baseScrollable() *Scrollable {
	return v
}

// BaseScrollable returns the underlying base object.
func BaseScrollable(obj Scrollabler) *Scrollable {
	return obj.baseScrollable()
}

// ScrollableInterface: instance of this type is always passed by reference.
type ScrollableInterface struct {
	*scrollableInterface
}

// scrollableInterface is the struct that's finalized.
type scrollableInterface struct {
	native unsafe.Pointer
}

var GIRInfoScrollableInterface = girepository.MustFind("Gtk", "ScrollableInterface")
