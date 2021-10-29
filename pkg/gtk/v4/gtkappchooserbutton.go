// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButtonner},
	})
}

// AppChooserButton: GtkAppChooserButton lets the user select an application.
//
// !An example GtkAppChooserButton (appchooserbutton.png)
//
// Initially, a GtkAppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// gtk.AppChooserButton:show-default-item is TRUE, the default application.
//
// The list of applications shown in a GtkAppChooserButton includes the
// recommended applications for the given content type. When
// gtk.AppChooserButton:show-default-item is set, the default application is
// also included. To let the user chooser other applications, you can set the
// gtk.AppChooserButton:show-dialog-item property, which allows to open a full
// gtk.AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk.AppChooserButton.AppendCustomItem(). These items cause the
// gtk.AppChooserButton::custom-item-activated signal to be emitted when they
// are selected.
//
// To track changes in the selected application, use the
// gtk.AppChooserButton::changed signal.
//
//
// CSS nodes
//
// GtkAppChooserButton has a single CSS node with the name “appchooserbutton”.
type AppChooserButton struct {
	Widget

	AppChooser
	*externglib.Object
}

var (
	_ Widgetter           = (*AppChooserButton)(nil)
	_ externglib.Objector = (*AppChooserButton)(nil)
)

func wrapAppChooserButton(obj *externglib.Object) *AppChooserButton {
	return &AppChooserButton{
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
		AppChooser: AppChooser{
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

func marshalAppChooserButtonner(p uintptr) (interface{}, error) {
	return wrapAppChooserButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAppChooserButton creates a new GtkAppChooserButton for applications that
// can handle content of the given type.
//
// The function takes the following parameters:
//
//    - contentType: content type to show applications for.
//
func NewAppChooserButton(contentType string) *AppChooserButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_button_new(_arg1)
	runtime.KeepAlive(contentType)

	var _appChooserButton *AppChooserButton // out

	_appChooserButton = wrapAppChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserButton
}

// AppendCustomItem appends a custom item to the list of applications that is
// shown in the popup.
//
// The item name must be unique per-widget. Clients can use the provided name as
// a detail for the gtk.AppChooserButton::custom-item-activated signal, to add a
// callback for the activation of a particular custom item in the list.
//
// See also gtk.AppChooserButton.AppendSeparator().
//
// The function takes the following parameters:
//
//    - name of the custom item.
//    - label for the custom item.
//    - icon for the custom item.
//
func (self *AppChooserButton) AppendCustomItem(name, label string, icon gio.Iconner) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out
	var _arg2 *C.char                // out
	var _arg3 *C.GIcon               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_app_chooser_button_append_custom_item(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(icon)
}

// AppendSeparator appends a separator to the list of applications that is shown
// in the popup.
func (self *AppChooserButton) AppendSeparator() {
	var _arg0 *C.GtkAppChooserButton // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
	runtime.KeepAlive(self)
}

// Heading returns the text to display at the top of the dialog.
func (self *AppChooserButton) Heading() string {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_heading(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Modal gets whether the dialog is modal.
func (self *AppChooserButton) Modal() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_modal(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefaultItem returns whether the dropdown menu should show the default
// application at the top.
func (self *AppChooserButton) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns whether the dropdown menu shows an item for a
// GtkAppChooserDialog.
func (self *AppChooserButton) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item.
//
// See gtk.AppChooserButton.AppendCustomItem().
//
// Use gtk.AppChooser.Refresh() to bring the selection to its initial state.
//
// The function takes the following parameters:
//
//    - name of the custom item.
//
func (self *AppChooserButton) SetActiveCustomItem(name string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetHeading sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
//
// The function takes the following parameters:
//
//    - heading: string containing Pango markup.
//
func (self *AppChooserButton) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}

// SetModal sets whether the dialog should be modal.
//
// The function takes the following parameters:
//
//    - modal: TRUE to make the dialog modal.
//
func (self *AppChooserButton) SetModal(modal bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_modal(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(modal)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should show
// the default application for the given content type at top.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserButton:show-default-item.
//
func (self *AppChooserButton) SetShowDefaultItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowDialogItem sets whether the dropdown menu of this button should show
// an entry to trigger a GtkAppChooserDialog.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserButton:show-dialog-item.
//
func (self *AppChooserButton) SetShowDialogItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// ConnectChanged: emitted when the active application changes.
func (self *AppChooserButton) ConnectChanged(f func()) externglib.SignalHandle {
	return self.Connect("changed", f)
}

// ConnectCustomItemActivated: emitted when a custom item is activated.
//
// Use gtk.AppChooserButton.AppendCustomItem(), to add custom items.
func (self *AppChooserButton) ConnectCustomItemActivated(f func(itemName string)) externglib.SignalHandle {
	return self.Connect("custom-item-activated", f)
}
