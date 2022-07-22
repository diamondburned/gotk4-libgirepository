// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeTearoffMenuItem = coreglib.Type(C.gtk_tearoff_menu_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTearoffMenuItem, F: marshalTearoffMenuItem},
	})
}

// TearoffMenuItemOverrider contains methods that are overridable.
type TearoffMenuItemOverrider interface {
}

// TearoffMenuItem is a special MenuItem which is used to tear off and reattach
// its menu.
//
// When its menu is shown normally, the TearoffMenuItem is drawn as a dotted
// line indicating that the menu can be torn off. Activating it causes its menu
// to be torn off and displayed in its own window as a tearoff menu.
//
// When its menu is shown as a tearoff menu, the TearoffMenuItem is drawn as a
// dotted line which has a left pointing arrow graphic indicating that the
// tearoff menu can be reattached. Activating it will erase the tearoff menu
// window.
//
// > TearoffMenuItem is deprecated and should not be used in newly > written
// code. Menus are not meant to be torn around.
type TearoffMenuItem struct {
	_ [0]func() // equal guard
	MenuItem
}

var (
	_ Binner            = (*TearoffMenuItem)(nil)
	_ coreglib.Objector = (*TearoffMenuItem)(nil)
)

func initClassTearoffMenuItem(gclass unsafe.Pointer, goval any) {
}

func wrapTearoffMenuItem(obj *coreglib.Object) *TearoffMenuItem {
	return &TearoffMenuItem{
		MenuItem: MenuItem{
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
			Object: obj,
			Actionable: Actionable{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalTearoffMenuItem(p uintptr) (interface{}, error) {
	return wrapTearoffMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTearoffMenuItem creates a new TearoffMenuItem.
//
// Deprecated: TearoffMenuItem is deprecated and should not be used in newly
// written code.
//
// The function returns the following values:
//
//    - tearoffMenuItem: new TearoffMenuItem.
//
func NewTearoffMenuItem() *TearoffMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tearoff_menu_item_new()

	var _tearoffMenuItem *TearoffMenuItem // out

	_tearoffMenuItem = wrapTearoffMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _tearoffMenuItem
}
