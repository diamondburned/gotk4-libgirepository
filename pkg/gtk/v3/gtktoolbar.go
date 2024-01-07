// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
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
// extern gboolean _gotk4_gtk3_Toolbar_ConnectPopupContextMenu(gpointer, gint, gint, gint, guintptr);
// extern gboolean _gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd(gpointer, gboolean, guintptr);
import "C"

// GType values.
var (
	GTypeToolbarSpaceStyle = coreglib.Type(girepository.MustFind("Gtk", "ToolbarSpaceStyle").RegisteredGType())
	GTypeToolbar           = coreglib.Type(girepository.MustFind("Gtk", "Toolbar").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToolbarSpaceStyle, F: marshalToolbarSpaceStyle},
		coreglib.TypeMarshaler{T: GTypeToolbar, F: marshalToolbar},
	})
}

// ToolbarSpaceStyle: whether spacers are vertical lines or just blank.
//
// Deprecated: since version 3.20.
type ToolbarSpaceStyle C.gint

const (
	// ToolbarSpaceEmpty: use blank spacers.
	ToolbarSpaceEmpty ToolbarSpaceStyle = iota
	// ToolbarSpaceLine: use vertical lines for spacers.
	ToolbarSpaceLine
)

func marshalToolbarSpaceStyle(p uintptr) (interface{}, error) {
	return ToolbarSpaceStyle(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ToolbarSpaceStyle.
func (t ToolbarSpaceStyle) String() string {
	switch t {
	case ToolbarSpaceEmpty:
		return "Empty"
	case ToolbarSpaceLine:
		return "Line"
	default:
		return fmt.Sprintf("ToolbarSpaceStyle(%d)", t)
	}
}

// ToolbarOverrides contains methods that are overridable.
type ToolbarOverrides struct {
}

func defaultToolbarOverrides(v *Toolbar) ToolbarOverrides {
	return ToolbarOverrides{}
}

// Toolbar: toolbar is created with a call to gtk_toolbar_new().
//
// A toolbar can contain instances of a subclass of ToolItem. To add a ToolItem
// to the a toolbar, use gtk_toolbar_insert(). To remove an item from the
// toolbar use gtk_container_remove(). To add a button to the toolbar, add an
// instance of ToolButton.
//
// Toolbar items can be visually grouped by adding instances of
// SeparatorToolItem to the toolbar. If the GtkToolbar child property “expand”
// is UE and the property SeparatorToolItem:draw is set to LSE, the effect is to
// force all following items to the end of the toolbar.
//
// By default, a toolbar can be shrunk, upon which it will add an arrow button
// to show an overflow menu offering access to any ToolItem child that has a
// proxy menu item. To disable this and request enough size for all children,
// call gtk_toolbar_set_show_arrow() to set Toolbar:show-arrow to FALSE.
//
// Creating a context menu for the toolbar can be done by connecting to the
// Toolbar::popup-context-menu signal.
//
//
// CSS nodes
//
// GtkToolbar has a single CSS node with name toolbar.
type Toolbar struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	atk.ImplementorIface
	coreglib.InitiallyUnowned
	Buildable
	Orientable
	ToolShell
	Widget
}

var (
	_ Containerer       = (*Toolbar)(nil)
	_ coreglib.Objector = (*Toolbar)(nil)
	_ Widgetter         = (*Toolbar)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Toolbar, *ToolbarClass, ToolbarOverrides](
		GTypeToolbar,
		initToolbarClass,
		wrapToolbar,
		defaultToolbarOverrides,
	)
}

func initToolbarClass(gclass unsafe.Pointer, overrides ToolbarOverrides, classInitFunc func(*ToolbarClass)) {
	if classInitFunc != nil {
		class := (*ToolbarClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapToolbar(obj *coreglib.Object) *Toolbar {
	return &Toolbar{
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
		Object: obj,
		ImplementorIface: atk.ImplementorIface{
			Object: obj,
		},
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		ToolShell: ToolShell{
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
	}
}

func marshalToolbar(p uintptr) (interface{}, error) {
	return wrapToolbar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectFocusHomeOrEnd: keybinding signal used internally by GTK+. This signal
// can't be used in application code.
func (v *Toolbar) ConnectFocusHomeOrEnd(f func(focusHome bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "focus-home-or-end", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd), f)
}

// ConnectPopupContextMenu is emitted when the user right-clicks the toolbar or
// uses the keybinding to display a popup menu.
//
// Application developers should handle this signal if they want to display a
// context menu on the toolbar. The context-menu should appear at the
// coordinates given by x and y. The mouse button number is given by the button
// parameter. If the menu was popped up using the keybaord, button is -1.
func (v *Toolbar) ConnectPopupContextMenu(f func(x, y, button int) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "popup-context-menu", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectPopupContextMenu), f)
}

// ToolbarClass: instance of this type is always passed by reference.
type ToolbarClass struct {
	*toolbarClass
}

// toolbarClass is the struct that's finalized.
type toolbarClass struct {
	native unsafe.Pointer
}

var GIRInfoToolbarClass = girepository.MustFind("Gtk", "ToolbarClass")
