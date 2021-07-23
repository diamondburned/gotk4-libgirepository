// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalPasswordEntrier},
	})
}

// PasswordEntry: GtkPasswordEntry is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, GtkPasswordEntry will also place the text in a
// non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// GtkPasswordEntry provides only minimal API and should be used with the
// gtk.Editable API.
//
// CSS Nodes
//
//    entry.password
//    ╰── text
//        ├── image.caps-lock-indicator
//        ┊
//
//
// GtkPasswordEntry has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// GtkPasswordEntry uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	Widget

	Editable
	*externglib.Object
}

func wrapPasswordEntry(obj *externglib.Object) *PasswordEntry {
	return &PasswordEntry{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalPasswordEntrier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPasswordEntry(obj), nil
}

// NewPasswordEntry creates a GtkPasswordEntry.
func NewPasswordEntry() *PasswordEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_password_entry_new()

	var _passwordEntry *PasswordEntry // out

	_passwordEntry = wrapPasswordEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _passwordEntry
}

// Native solves the ambiguous selector of this class or interface.
func (entry *PasswordEntry) Native() uintptr {
	return entry.Object.Native()
}

// ExtraMenu gets the menu model set with gtk_password_entry_set_extra_menu().
func (entry *PasswordEntry) ExtraMenu() gio.MenuModeller {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret *C.GMenuModel       // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_extra_menu(_arg0)

	var _menuModel gio.MenuModeller // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.MenuModeller)

	return _menuModel
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
func (entry *PasswordEntry) ShowPeekIcon() bool {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_show_peek_icon(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// entry.
func (entry *PasswordEntry) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 *C.GMenuModel       // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_password_entry_set_extra_menu(_arg0, _arg1)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to reveal
// the contents.
//
// Setting this to FALSE also hides the text again.
func (entry *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	if showPeekIcon {
		_arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(_arg0, _arg1)
}
