// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_file_chooser_dialog_get_type()), F: marshalFileChooserDialogger},
	})
}

// FileChooserDialog is a dialog box suitable for use with “File/Open” or
// “File/Save as” commands. This widget works by putting a FileChooserWidget
// inside a Dialog. It exposes the FileChooser interface, so you can use all of
// the FileChooser functions on the file chooser dialog as well as those for
// Dialog.
//
// Note that FileChooserDialog does not have any methods of its own. Instead,
// you should use the functions that work on a FileChooser.
//
// If you want to integrate well with the platform you should use the
// FileChooserNative API, which will use a platform-specific dialog if available
// and fall back to GtkFileChooserDialog otherwise.
//
//
// Typical usage
//
// In the simplest of cases, you can the following code to use FileChooserDialog
// to select a file for opening:
//
//    GtkWidget *dialog;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
//    gint res;
//
//    dialog = gtk_file_chooser_dialog_new ("Open File",
//                                          parent_window,
//                                          action,
//                                          _("_Cancel"),
//                                          GTK_RESPONSE_CANCEL,
//                                          _("_Open"),
//                                          GTK_RESPONSE_ACCEPT,
//                                          NULL);
//
//    res = gtk_dialog_run (GTK_DIALOG (dialog));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        char *filename;
//        GtkFileChooser *chooser = GTK_FILE_CHOOSER (dialog);
//        filename = gtk_file_chooser_get_filename (chooser);
//        open_file (filename);
//        g_free (filename);
//      }
//
//    gtk_widget_destroy (dialog);
//
// To use a dialog for saving, you can use this:
//
//    GtkWidget *dialog;
//    GtkFileChooser *chooser;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
//    gint res;
//
//    dialog = gtk_file_chooser_dialog_new ("Save File",
//                                          parent_window,
//                                          action,
//                                          _("_Cancel"),
//                                          GTK_RESPONSE_CANCEL,
//                                          _("_Save"),
//                                          GTK_RESPONSE_ACCEPT,
//                                          NULL);
//    chooser = GTK_FILE_CHOOSER (dialog);
//
//    gtk_file_chooser_set_do_overwrite_confirmation (chooser, TRUE);
//
//    if (user_edited_a_new_document)
//      gtk_file_chooser_set_current_name (chooser,
//                                         _("Untitled document"));
//    else
//      gtk_file_chooser_set_filename (chooser,
//                                     existing_filename);
//
//    res = gtk_dialog_run (GTK_DIALOG (dialog));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        char *filename;
//
//        filename = gtk_file_chooser_get_filename (chooser);
//        save_to_file (filename);
//        g_free (filename);
//      }
//
//    gtk_widget_destroy (dialog);
//
//
// Setting up a file chooser dialog
//
// There are various cases in which you may need to use a FileChooserDialog:
//
// - To select a file for opening. Use K_FILE_CHOOSER_ACTION_OPEN.
//
// - To save a file for the first time. Use K_FILE_CHOOSER_ACTION_SAVE, and
// suggest a name such as “Untitled” with gtk_file_chooser_set_current_name().
//
// - To save a file under a different name. Use K_FILE_CHOOSER_ACTION_SAVE, and
// set the existing filename with gtk_file_chooser_set_filename().
//
// - To choose a folder instead of a file. Use
// K_FILE_CHOOSER_ACTION_SELECT_FOLDER.
//
// Note that old versions of the file chooser’s documentation suggested using
// gtk_file_chooser_set_current_folder() in various situations, with the
// intention of letting the application suggest a reasonable default folder.
// This is no longer considered to be a good policy, as now the file chooser is
// able to make good suggestions on its own. In general, you should only cause
// the file chooser to show a specific folder when it is appropriate to use
// gtk_file_chooser_set_filename(), i.e. when you are doing a Save As command
// and you already have a file saved somewhere.
//
//
// Response Codes
//
// FileChooserDialog inherits from Dialog, so buttons that go in its action area
// have response codes such as K_RESPONSE_ACCEPT and K_RESPONSE_CANCEL. For
// example, you could call gtk_file_chooser_dialog_new() as follows:
//
//    GtkWidget *dialog;
//    GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
//
//    dialog = gtk_file_chooser_dialog_new ("Open File",
//                                          parent_window,
//                                          action,
//                                          _("_Cancel"),
//                                          GTK_RESPONSE_CANCEL,
//                                          _("_Open"),
//                                          GTK_RESPONSE_ACCEPT,
//                                          NULL);
//
// This will create buttons for “Cancel” and “Open” that use stock response
// identifiers from ResponseType. For most dialog boxes you can use your own
// custom response codes rather than the ones in ResponseType, but
// FileChooserDialog assumes that its “accept”-type action, e.g. an “Open” or
// “Save” button, will have one of the following response codes:
//
// - K_RESPONSE_ACCEPT
//
// - K_RESPONSE_OK
//
// - K_RESPONSE_YES
//
// - K_RESPONSE_APPLY
//
// This is because FileChooserDialog must intercept responses and switch to
// folders if appropriate, rather than letting the dialog terminate — the
// implementation uses these known response codes to know which responses can be
// blocked if appropriate.
//
// To summarize, make sure you use a [stock response
// code][gtkfilechooserdialog-responses] when you use FileChooserDialog to
// ensure proper operation.
type FileChooserDialog struct {
	Dialog

	FileChooser
	*externglib.Object
}

func wrapFileChooserDialog(obj *externglib.Object) *FileChooserDialog {
	return &FileChooserDialog{
		Dialog: Dialog{
			Window: Window{
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
							Object: obj,
						},
					},
				},
			},
		},
		FileChooser: FileChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFileChooserDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooserDialog(obj), nil
}

func (*FileChooserDialog) privateFileChooserDialog() {}
