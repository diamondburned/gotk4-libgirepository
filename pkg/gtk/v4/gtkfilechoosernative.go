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
		{T: externglib.Type(C.gtk_file_chooser_native_get_type()), F: marshalFileChooserNative},
	})
}

// FileChooserNative: `GtkFileChooserNative` is an abstraction of a dialog
// suitable for use with “File Open” or “File Save as” commands.
//
// By default, this just uses a `GtkFileChooserDialog` to implement the actual
// dialog. However, on some platforms, such as Windows and macOS, the native
// platform file chooser is used instead. When the application is running in a
// sandboxed environment without direct filesystem access (such as Flatpak),
// `GtkFileChooserNative` may call the proper APIs (portals) to let the user
// choose a file and make it available to the application.
//
// While the API of `GtkFileChooserNative` closely mirrors
// `GtkFileChooserDialog`, the main difference is that there is no access to any
// `GtkWindow` or `GtkWidget` for the dialog. This is required, as there may not
// be one in the case of a platform native dialog.
//
// Showing, hiding and running the dialog is handled by the
// [class@Gtk.NativeDialog] functions.
//
// Note that unlike `GtkFileChooserDialog`, `GtkFileChooserNative` objects are
// not toplevel widgets, and GTK does not keep them alive. It is your
// responsibility to keep a reference until you are done with the object.
//
//
// Typical usage
//
// In the simplest of cases, you can the following code to use
// `GtkFileChooserNative` to select a file for opening:
//
// “`c static void on_response (GtkNativeDialog *native, int response) { if
// (response == GTK_RESPONSE_ACCEPT) { GtkFileChooser *chooser =
// GTK_FILE_CHOOSER (native); GFile *file = gtk_file_chooser_get_file (chooser);
//
//        open_file (file);
//
//        g_object_unref (file);
//      }
//
//    g_object_unref (native);
//
// }
//
//    // ...
//    GtkFileChooserNative *native;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
//
//    native = gtk_file_chooser_native_new ("Open File",
//                                          parent_window,
//                                          action,
//                                          "_Open",
//                                          "_Cancel");
//
//    g_signal_connect (native, "response", G_CALLBACK (on_response), NULL);
//    gtk_native_dialog_show (GTK_NATIVE_DIALOG (native));
//
// “`
//
// To use a `GtkFileChooserNative` for saving, you can use this:
//
// “`c static void on_response (GtkNativeDialog *native, int response) { if
// (response == GTK_RESPONSE_ACCEPT) { GtkFileChooser *chooser =
// GTK_FILE_CHOOSER (native); GFile *file = gtk_file_chooser_get_file (chooser);
//
//        save_to_file (file);
//
//        g_object_unref (file);
//      }
//
//    g_object_unref (native);
//
// }
//
//    // ...
//    GtkFileChooserNative *native;
//    GtkFileChooser *chooser;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
//
//    native = gtk_file_chooser_native_new ("Save File",
//                                          parent_window,
//                                          action,
//                                          "_Save",
//                                          "_Cancel");
//    chooser = GTK_FILE_CHOOSER (native);
//
//    if (user_edited_a_new_document)
//      gtk_file_chooser_set_current_name (chooser, _("Untitled document"));
//    else
//      gtk_file_chooser_set_file (chooser, existing_file, NULL);
//
//    g_signal_connect (native, "response", G_CALLBACK (on_response), NULL);
//    gtk_native_dialog_show (GTK_NATIVE_DIALOG (native));
//
// “`
//
// For more information on how to best set up a file dialog, see the
// [class@Gtk.FileChooserDialog] documentation.
//
//
// Response Codes
//
// `GtkFileChooserNative` inherits from [class@Gtk.NativeDialog], which means it
// will return GTK_RESPONSE_ACCEPT if the user accepted, and GTK_RESPONSE_CANCEL
// if he pressed cancel. It can also return GTK_RESPONSE_DELETE_EVENT if the
// window was unexpectedly closed.
//
//
// Differences from FileChooserDialog
//
// There are a few things in the [iface@Gtk.FileChooser] interface that are not
// possible to use with `GtkFileChooserNative`, as such use would prohibit the
// use of a native dialog.
//
// No operations that change the dialog work while the dialog is visible. Set
// all the properties that are required before showing the dialog.
//
//
// Win32 details
//
// On windows the `IFileDialog` implementation (added in Windows Vista) is used.
// It supports many of the features that `GtkFileChooser` has, but there are
// some things it does not handle:
//
// * Any [class@Gtk.FileFilter] added using a mimetype
//
// If any of these features are used the regular `GtkFileChooserDialog` will be
// used in place of the native one.
//
//
// Portal details
//
// When the `org.freedesktop.portal.FileChooser` portal is available on the
// session bus, it is used to bring up an out-of-process file chooser. Depending
// on the kind of session the application is running in, this may or may not be
// a GTK file chooser.
//
// macOS details
//
// On macOS the `NSSavePanel` and `NSOpenPanel` classes are used to provide
// native file chooser dialogs. Some features provided by `GtkFileChooser` are
// not supported:
//
// * Shortcut folders.
type FileChooserNative interface {
	NativeDialog
	FileChooser

	// AcceptLabel retrieves the custom label text for the accept button.
	AcceptLabel() string
	// CancelLabel retrieves the custom label text for the cancel button.
	CancelLabel() string
	// SetAcceptLabel sets the custom label text for the accept button.
	//
	// If characters in @label are preceded by an underscore, they are
	// underlined. If you need a literal underscore character in a label, use
	// “__” (two underscores). The first underlined character represents a
	// keyboard accelerator called a mnemonic.
	//
	// Pressing Alt and that key should activate the button.
	SetAcceptLabel(acceptLabel string)
	// SetCancelLabel sets the custom label text for the cancel button.
	//
	// If characters in @label are preceded by an underscore, they are
	// underlined. If you need a literal underscore character in a label, use
	// “__” (two underscores). The first underlined character represents a
	// keyboard accelerator called a mnemonic.
	//
	// Pressing Alt and that key should activate the button.
	SetCancelLabel(cancelLabel string)
}

// fileChooserNative implements the FileChooserNative class.
type fileChooserNative struct {
	NativeDialog
	FileChooser
}

var _ FileChooserNative = (*fileChooserNative)(nil)

// WrapFileChooserNative wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserNative(obj *externglib.Object) FileChooserNative {
	return fileChooserNative{
		NativeDialog: WrapNativeDialog(obj),
		FileChooser:  WrapFileChooser(obj),
	}
}

func marshalFileChooserNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserNative(obj), nil
}

// AcceptLabel retrieves the custom label text for the accept button.
func (s fileChooserNative) AcceptLabel() string {
	var _arg0 *C.GtkFileChooserNative // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gtk_file_chooser_native_get_accept_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// CancelLabel retrieves the custom label text for the cancel button.
func (s fileChooserNative) CancelLabel() string {
	var _arg0 *C.GtkFileChooserNative // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gtk_file_chooser_native_get_cancel_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetAcceptLabel sets the custom label text for the accept button.
//
// If characters in @label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use
// “__” (two underscores). The first underlined character represents a
// keyboard accelerator called a mnemonic.
//
// Pressing Alt and that key should activate the button.
func (s fileChooserNative) SetAcceptLabel(acceptLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(acceptLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_native_set_accept_label(_arg0, _arg1)
}

// SetCancelLabel sets the custom label text for the cancel button.
//
// If characters in @label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use
// “__” (two underscores). The first underlined character represents a
// keyboard accelerator called a mnemonic.
//
// Pressing Alt and that key should activate the button.
func (s fileChooserNative) SetCancelLabel(cancelLabel string) {
	var _arg0 *C.GtkFileChooserNative // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(cancelLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_native_set_cancel_label(_arg0, _arg1)
}
