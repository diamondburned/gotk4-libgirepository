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
	GTypeSearchBar = coreglib.Type(girepository.MustFind("Gtk", "SearchBar").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSearchBar, F: marshalSearchBar},
	})
}

// SearchBar: GtkSearchBar is a container made to have a search entry.
//
// !An example GtkSearchBar (search-bar.png)
//
// It can also contain additional widgets, such as drop-down menus, or buttons.
// The search bar would appear when a search is started through typing on the
// keyboard, or the application’s search mode is toggled on.
//
// For keyboard presses to start a search, the search bar must be told of a
// widget to capture key events from through
// gtk.SearchBar.SetKeyCaptureWidget(). This widget will typically be the
// top-level window, or a parent container of the search bar. Common shortcuts
// such as Ctrl+F should be handled as an application action, or through the
// menu items.
//
// You will also need to tell the search bar about which entry you are using as
// your search entry using gtk.SearchBar.ConnectEntry().
//
//
// Creating a search bar
//
// The following example shows you how to create a more complex search entry.
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/search-bar.c)
//
// CSS nodes
//
//    searchbar
//    ╰── revealer
//        ╰── box
//             ├── [child]
//             ╰── [button.close]
//
//
// GtkSearchBar has a main CSS node with name searchbar. It has a child node
// with name revealer that contains a node with name box. The box node contains
// both the CSS node of the child widget as well as an optional button node
// which gets the .close style class applied.
//
//
// Accessibility
//
// GtkSearchBar uses the GTK_ACCESSIBLE_ROLE_SEARCH role.
type SearchBar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*SearchBar)(nil)
)

func wrapSearchBar(obj *coreglib.Object) *SearchBar {
	return &SearchBar{
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
	}
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	return wrapSearchBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
