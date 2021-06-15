// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButton},
	})
}

// AppChooserButton: the `GtkAppChooserButton` lets the user select an
// application.
//
// !An example GtkAppChooserButton (appchooserbutton.png)
//
// Initially, a `GtkAppChooserButton` selects the first application in its list,
// which will either be the most-recently used application or, if
// [property@Gtk.AppChooserButton:show-default-item] is true, the default
// application.
//
// The list of applications shown in a `GtkAppChooserButton` includes the
// recommended applications for the given content type. When
// [property@Gtk.AppChooserButton:show-default-item] is set, the default
// application is also included. To let the user chooser other applications, you
// can set the [property@Gtk.AppChooserButton:show-dialog-item] property, which
// allows to open a full [class@Gtk.AppChooserDialog].
//
// It is possible to add custom items to the list, using
// [method@Gtk.AppChooserButton.append_custom_item]. These items cause the
// [signal@Gtk.AppChooserButton::custom-item-activated] signal to be emitted
// when they are selected.
//
// To track changes in the selected application, use the
// [signal@Gtk.AppChooserButton::changed] signal.
//
//
// CSS nodes
//
// `GtkAppChooserButton` has a single CSS node with the name “appchooserbutton”.
type AppChooserButton interface {
	Widget
	Accessible
	AppChooser
	Buildable
	ConstraintTarget

	// AppendCustomItem appends a custom item to the list of applications that
	// is shown in the popup.
	//
	// The item name must be unique per-widget. Clients can use the provided
	// name as a detail for the
	// [signal@Gtk.AppChooserButton::custom-item-activated] signal, to add a
	// callback for the activation of a particular custom item in the list.
	//
	// See also [method@Gtk.AppChooserButton.append_separator].
	AppendCustomItem(name string, label string, icon gio.Icon)
	// AppendSeparator appends a separator to the list of applications that is
	// shown in the popup.
	AppendSeparator()
	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// Modal gets whether the dialog is modal.
	Modal() bool
	// ShowDefaultItem returns whether the dropdown menu should show the default
	// application at the top.
	ShowDefaultItem() bool
	// ShowDialogItem returns whether the dropdown menu shows an item for a
	// `GtkAppChooserDialog`.
	ShowDialogItem() bool
	// SetActiveCustomItem selects a custom item.
	//
	// See [method@Gtk.AppChooserButton.append_custom_item].
	//
	// Use [method@Gtk.AppChooser.refresh] to bring the selection to its initial
	// state.
	SetActiveCustomItem(name string)
	// SetHeading sets the text to display at the top of the dialog.
	//
	// If the heading is not set, the dialog displays a default text.
	SetHeading(heading string)
	// SetModal sets whether the dialog should be modal.
	SetModal(modal bool)
	// SetShowDefaultItem sets whether the dropdown menu of this button should
	// show the default application for the given content type at top.
	SetShowDefaultItem(setting bool)
	// SetShowDialogItem sets whether the dropdown menu of this button should
	// show an entry to trigger a `GtkAppChooserDialog`.
	SetShowDialogItem(setting bool)
}

// appChooserButton implements the AppChooserButton class.
type appChooserButton struct {
	Widget
	Accessible
	AppChooser
	Buildable
	ConstraintTarget
}

var _ AppChooserButton = (*appChooserButton)(nil)

// WrapAppChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserButton(obj *externglib.Object) AppChooserButton {
	return appChooserButton{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		AppChooser:       WrapAppChooser(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserButton(obj), nil
}

// NewAppChooserButton constructs a class AppChooserButton.
func NewAppChooserButton(contentType string) AppChooserButton {
	var _arg1 *C.char               // out
	var _cret C.GtkAppChooserButton // in

	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_button_new(_arg1)

	var _appChooserButton AppChooserButton // out

	_appChooserButton = WrapAppChooserButton(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserButton
}

// AppendCustomItem appends a custom item to the list of applications that
// is shown in the popup.
//
// The item name must be unique per-widget. Clients can use the provided
// name as a detail for the
// [signal@Gtk.AppChooserButton::custom-item-activated] signal, to add a
// callback for the activation of a particular custom item in the list.
//
// See also [method@Gtk.AppChooserButton.append_separator].
func (s appChooserButton) AppendCustomItem(name string, label string, icon gio.Icon) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out
	var _arg2 *C.char                // out
	var _arg3 *C.GIcon               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_app_chooser_button_append_custom_item(_arg0, _arg1, _arg2, _arg3)
}

// AppendSeparator appends a separator to the list of applications that is
// shown in the popup.
func (s appChooserButton) AppendSeparator() {
	var _arg0 *C.GtkAppChooserButton // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
}

// Heading returns the text to display at the top of the dialog.
func (s appChooserButton) Heading() string {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Modal gets whether the dialog is modal.
func (s appChooserButton) Modal() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefaultItem returns whether the dropdown menu should show the default
// application at the top.
func (s appChooserButton) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns whether the dropdown menu shows an item for a
// `GtkAppChooserDialog`.
func (s appChooserButton) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item.
//
// See [method@Gtk.AppChooserButton.append_custom_item].
//
// Use [method@Gtk.AppChooser.refresh] to bring the selection to its initial
// state.
func (s appChooserButton) SetActiveCustomItem(name string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
}

// SetHeading sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
func (s appChooserButton) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
}

// SetModal sets whether the dialog should be modal.
func (s appChooserButton) SetModal(modal bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_modal(_arg0, _arg1)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should
// show the default application for the given content type at top.
func (s appChooserButton) SetShowDefaultItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
}

// SetShowDialogItem sets whether the dropdown menu of this button should
// show an entry to trigger a `GtkAppChooserDialog`.
func (s appChooserButton) SetShowDialogItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
}
