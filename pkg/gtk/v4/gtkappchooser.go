// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooser},
	})
}

// AppChooser: `GtkAppChooser` is an interface for widgets which allow the user
// to choose an application.
//
// The main objects that implement this interface are
// [class@Gtk.AppChooserWidget], [class@Gtk.AppChooserDialog] and
// [class@Gtk.AppChooserButton].
//
// Applications are represented by GIO `GAppInfo` objects here. GIO has a
// concept of recommended and fallback applications for a given content type.
// Recommended applications are those that claim to handle the content type
// itself, while fallback also includes applications that handle a more generic
// content type. GIO also knows the default and last-used application for a
// given content type. The `GtkAppChooserWidget` provides detailed control over
// whether the shown list of applications should include default, recommended or
// fallback applications.
//
// To obtain the application that has been selected in a `GtkAppChooser`, use
// [method@Gtk.AppChooser.get_app_info].
type AppChooser interface {
	Widget

	// ContentType returns the content type for which the `GtkAppChooser` shows
	// applications.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

// appChooser implements the AppChooser interface.
type appChooser struct {
	Widget
}

var _ AppChooser = (*appChooser)(nil)

// WrapAppChooser wraps a GObject to a type that implements interface
// AppChooser. It is primarily used internally.
func WrapAppChooser(obj *externglib.Object) AppChooser {
	return AppChooser{
		Widget: WrapWidget(obj),
	}
}

func marshalAppChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooser(obj), nil
}

// ContentType returns the content type for which the `GtkAppChooser` shows
// applications.
func (s appChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gtk_app_chooser_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Refresh reloads the list of applications.
func (s appChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_refresh(_arg0)
}
