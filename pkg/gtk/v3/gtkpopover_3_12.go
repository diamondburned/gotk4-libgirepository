// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_Popover_ConnectClosed(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypePopover = coreglib.Type(girepository.MustFind("Gtk", "Popover").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePopover, F: marshalPopover},
	})
}

// PopoverOverrides contains methods that are overridable.
type PopoverOverrides struct {
}

func defaultPopoverOverrides(v *Popover) PopoverOverrides {
	return PopoverOverrides{}
}

// Popover is a bubble-like context window, primarily meant to provide
// context-dependent information or options. Popovers are attached to a widget,
// passed at construction time on gtk_popover_new(), or updated afterwards
// through gtk_popover_set_relative_to(), by default they will point to the
// whole widget area, although this behavior can be changed through
// gtk_popover_set_pointing_to().
//
// The position of a popover relative to the widget it is attached to can also
// be changed through gtk_popover_set_position().
//
// By default, Popover performs a GTK+ grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Esc key being
// pressed). If no such modal behavior is desired on a popover,
// gtk_popover_set_modal() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. To facilitate this, it supports
// being populated from a Model, using gtk_popover_new_from_model(). In addition
// to all the regular menu model features, this function supports rendering
// sections in the model in a more compact form, as a row of icon buttons
// instead of menu items.
//
// To use this rendering, set the ”display-hint” attribute of the section to
// ”horizontal-buttons” and set the icons of your items with the ”verb-icon”
// attribute.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
// GtkPopover has a single css node called popover. It always gets the
// .background style class and it gets the .menu style class if it is menu-like
// (e.g. PopoverMenu or created using gtk_popover_new_from_model().
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in Entry or TextView get style classes like .touch-selection or .magnifier to
// differentiate from plain popovers.
type Popover struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Popover)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Popover, *PopoverClass, PopoverOverrides](
		GTypePopover,
		initPopoverClass,
		wrapPopover,
		defaultPopoverOverrides,
	)
}

func initPopoverClass(gclass unsafe.Pointer, overrides PopoverOverrides, classInitFunc func(*PopoverClass)) {
	if classInitFunc != nil {
		class := (*PopoverClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPopover(obj *coreglib.Object) *Popover {
	return &Popover{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	return wrapPopover(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClosed: this signal is emitted when the popover is dismissed either
// through API or user interaction.
func (v *Popover) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "closed", false, unsafe.Pointer(C._gotk4_gtk3_Popover_ConnectClosed), f)
}
