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
// extern void _gotk4_gtk4_ShortcutsWindow_ConnectSearch(gpointer, guintptr);
// extern void _gotk4_gtk4_ShortcutsWindow_ConnectClose(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeShortcutsWindow = coreglib.Type(girepository.MustFind("Gtk", "ShortcutsWindow").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutsWindow, F: marshalShortcutsWindow},
	})
}

// ShortcutsWindow: GtkShortcutsWindow shows information about the keyboard
// shortcuts and gestures of an application.
//
// The shortcuts can be grouped, and you can have multiple sections in this
// window, corresponding to the major modes of your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with gtk.Builder, by
// populating a GtkShortcutsWindow with one or more GtkShortcutsSection objects,
// which contain GtkShortcutsGroups that in turn contain objects of class
// GtkShortcutsShortcut.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a GtkShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a GtkShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Widgetter         = (*ShortcutsWindow)(nil)
	_ coreglib.Objector = (*ShortcutsWindow)(nil)
)

func wrapShortcutsWindow(obj *coreglib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						Accessible: Accessible{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						ConstraintTarget: ConstraintTarget{
							Object: obj,
						},
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	return wrapShortcutsWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose is emitted when the user uses a keybinding to close the window.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is the Escape key.
func (v *ShortcutsWindow) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "close", false, unsafe.Pointer(C._gotk4_gtk4_ShortcutsWindow_ConnectClose), f)
}

// ConnectSearch is emitted when the user uses a keybinding to start a search.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Control-F.
func (v *ShortcutsWindow) ConnectSearch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "search", false, unsafe.Pointer(C._gotk4_gtk4_ShortcutsWindow_ConnectSearch), f)
}
