// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_menu_item_get_type()), F: marshalCheckMenuItemer},
	})
}

// CheckMenuItemOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CheckMenuItemOverrider interface {
	DrawIndicator(cr *cairo.Context)
	// Toggled emits the CheckMenuItem::toggled signal.
	Toggled()
}

// CheckMenuItemer describes CheckMenuItem's methods.
type CheckMenuItemer interface {
	// Active returns whether the check menu item is active.
	Active() bool
	// DrawAsRadio returns whether check_menu_item looks like a RadioMenuItem
	DrawAsRadio() bool
	// Inconsistent retrieves the value set by
	// gtk_check_menu_item_set_inconsistent().
	Inconsistent() bool
	// SetActive sets the active state of the menu item’s check box.
	SetActive(isActive bool)
	// SetDrawAsRadio sets whether check_menu_item is drawn like a RadioMenuItem
	SetDrawAsRadio(drawAsRadio bool)
	// SetInconsistent: if the user has selected a range of elements (such as
	// some text or spreadsheet cells) that are affected by a boolean setting,
	// and the current values in that range are inconsistent, you may want to
	// display the check in an “in between” state.
	SetInconsistent(setting bool)
	// Toggled emits the CheckMenuItem::toggled signal.
	Toggled()
}

// CheckMenuItem is a menu item that maintains the state of a boolean value in
// addition to a MenuItem usual role in activating application code.
//
// A check box indicating the state of the boolean value is displayed at the
// left side of the MenuItem. Activating the MenuItem toggles the value.
//
// CSS nodes
//
//    menuitem
//    ├── check.left
//    ╰── <child>
//
// GtkCheckMenuItem has a main CSS node with name menuitem, and a subnode with
// name check, which gets the .left or .right style class.
type CheckMenuItem struct {
	MenuItem
}

var (
	_ CheckMenuItemer = (*CheckMenuItem)(nil)
	_ gextras.Nativer = (*CheckMenuItem)(nil)
)

func wrapCheckMenuItem(obj *externglib.Object) *CheckMenuItem {
	return &CheckMenuItem{
		MenuItem: MenuItem{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
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

func marshalCheckMenuItemer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCheckMenuItem(obj), nil
}

// NewCheckMenuItem creates a new CheckMenuItem.
func NewCheckMenuItem() *CheckMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_check_menu_item_new()

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _checkMenuItem
}

// NewCheckMenuItemWithLabel creates a new CheckMenuItem with a label.
func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_check_menu_item_new_with_label(_arg1)

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _checkMenuItem
}

// NewCheckMenuItemWithMnemonic creates a new CheckMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_check_menu_item_new_with_mnemonic(_arg1)

	var _checkMenuItem *CheckMenuItem // out

	_checkMenuItem = wrapCheckMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _checkMenuItem
}

// Active returns whether the check menu item is active. See
// gtk_check_menu_item_set_active ().
func (checkMenuItem *CheckMenuItem) Active() bool {
	var _arg0 *C.GtkCheckMenuItem // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))

	_cret = C.gtk_check_menu_item_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether check_menu_item looks like a RadioMenuItem
func (checkMenuItem *CheckMenuItem) DrawAsRadio() bool {
	var _arg0 *C.GtkCheckMenuItem // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))

	_cret = C.gtk_check_menu_item_get_draw_as_radio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inconsistent retrieves the value set by
// gtk_check_menu_item_set_inconsistent().
func (checkMenuItem *CheckMenuItem) Inconsistent() bool {
	var _arg0 *C.GtkCheckMenuItem // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))

	_cret = C.gtk_check_menu_item_get_inconsistent(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the active state of the menu item’s check box.
func (checkMenuItem *CheckMenuItem) SetActive(isActive bool) {
	var _arg0 *C.GtkCheckMenuItem // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_check_menu_item_set_active(_arg0, _arg1)
}

// SetDrawAsRadio sets whether check_menu_item is drawn like a RadioMenuItem
func (checkMenuItem *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	var _arg0 *C.GtkCheckMenuItem // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))
	if drawAsRadio {
		_arg1 = C.TRUE
	}

	C.gtk_check_menu_item_set_draw_as_radio(_arg0, _arg1)
}

// SetInconsistent: if the user has selected a range of elements (such as some
// text or spreadsheet cells) that are affected by a boolean setting, and the
// current values in that range are inconsistent, you may want to display the
// check in an “in between” state. This function turns on “in between” display.
// Normally you would turn off the inconsistent state again if the user
// explicitly selects a setting. This has to be done manually,
// gtk_check_menu_item_set_inconsistent() only affects visual appearance, it
// doesn’t affect the semantics of the widget.
func (checkMenuItem *CheckMenuItem) SetInconsistent(setting bool) {
	var _arg0 *C.GtkCheckMenuItem // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_check_menu_item_set_inconsistent(_arg0, _arg1)
}

// Toggled emits the CheckMenuItem::toggled signal.
func (checkMenuItem *CheckMenuItem) Toggled() {
	var _arg0 *C.GtkCheckMenuItem // out

	_arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem.Native()))

	C.gtk_check_menu_item_toggled(_arg0)
}
