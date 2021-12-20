// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_bar_get_type()), F: marshalMenuBarrer},
	})
}

// MenuBar is a subclass of MenuShell which contains one or more MenuItems. The
// result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar struct {
	MenuShell

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ MenuSheller = (*MenuBar)(nil)
)

func wrapMenuBar(obj *externglib.Object) *MenuBar {
	return &MenuBar{
		MenuShell: MenuShell{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalMenuBarrer(p uintptr) (interface{}, error) {
	return wrapMenuBar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenuBar creates a new MenuBar.
//
// The function returns the following values:
//
//    - menuBar: new menu bar, as a Widget.
//
func NewMenuBar() *MenuBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_bar_new()

	var _menuBar *MenuBar // out

	_menuBar = wrapMenuBar(externglib.Take(unsafe.Pointer(_cret)))

	return _menuBar
}

// NewMenuBarFromModel creates a new MenuBar and populates it with menu items
// and submenus according to model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu bar belongs - typically by means of being
// contained within the ApplicationWindows widget hierarchy.
//
// The function takes the following parameters:
//
//    - model: Model.
//
// The function returns the following values:
//
//    - menuBar: new MenuBar.
//
func NewMenuBarFromModel(model gio.MenuModeller) *MenuBar {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_menu_bar_new_from_model(_arg1)
	runtime.KeepAlive(model)

	var _menuBar *MenuBar // out

	_menuBar = wrapMenuBar(externglib.Take(unsafe.Pointer(_cret)))

	return _menuBar
}

// ChildPackDirection retrieves the current child pack direction of the menubar.
// See gtk_menu_bar_set_child_pack_direction().
//
// The function returns the following values:
//
//    - packDirection: child pack direction.
//
func (menubar *MenuBar) ChildPackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))

	_cret = C.gtk_menu_bar_get_child_pack_direction(_arg0)
	runtime.KeepAlive(menubar)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

// PackDirection retrieves the current pack direction of the menubar. See
// gtk_menu_bar_set_pack_direction().
//
// The function returns the following values:
//
//    - packDirection: pack direction.
//
func (menubar *MenuBar) PackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))

	_cret = C.gtk_menu_bar_get_pack_direction(_arg0)
	runtime.KeepAlive(menubar)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

// SetChildPackDirection sets how widgets should be packed inside the children
// of a menubar.
//
// The function takes the following parameters:
//
//    - childPackDir: new PackDirection.
//
func (menubar *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))
	_arg1 = C.GtkPackDirection(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction(_arg0, _arg1)
	runtime.KeepAlive(menubar)
	runtime.KeepAlive(childPackDir)
}

// SetPackDirection sets how items should be packed inside a menubar.
//
// The function takes the following parameters:
//
//    - packDir: new PackDirection.
//
func (menubar *MenuBar) SetPackDirection(packDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))
	_arg1 = C.GtkPackDirection(packDir)

	C.gtk_menu_bar_set_pack_direction(_arg0, _arg1)
	runtime.KeepAlive(menubar)
	runtime.KeepAlive(packDir)
}
