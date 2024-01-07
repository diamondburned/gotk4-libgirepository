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
	GTypePageSetup = coreglib.Type(girepository.MustFind("Gtk", "PageSetup").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePageSetup, F: marshalPageSetup},
	})
}

// PageSetup: GtkPageSetup object stores the page size, orientation and margins.
//
// The idea is that you can get one of these from the page setup dialog and then
// pass it to the GtkPrintOperation when printing. The benefit of splitting this
// out of the GtkPrintSettings is that these affect the actual layout of the
// page, and thus need to be set long before user prints.
//
//
// Margins
//
// The margins specified in this object are the “print margins”, i.e. the parts
// of the page that the printer cannot print on. These are different from the
// layout margins that a word processor uses; they are typically used to
// determine the minimal size for the layout margins.
//
// To obtain a GtkPageSetup use gtk.PageSetup.New to get the defaults, or use
// gtk.PrintRunPageSetupDialog() to show the page setup dialog and receive the
// resulting page setup.
//
// A page setup dialog
//
//    static GtkPrintSettings *settings = NULL;
//    static GtkPageSetup *page_setup = NULL;
//
//    static void
//    do_page_setup (void)
//    {
//      GtkPageSetup *new_page_setup;
//
//      if (settings == NULL)
//        settings = gtk_print_settings_new ();
//
//      new_page_setup = gtk_print_run_page_setup_dialog (GTK_WINDOW (main_window),
//                                                        page_setup, settings);
//
//      if (page_setup)
//        g_object_unref (page_setup);
//
//      page_setup = new_page_setup;
//    }.
type PageSetup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PageSetup)(nil)
)

func wrapPageSetup(obj *coreglib.Object) *PageSetup {
	return &PageSetup{
		Object: obj,
	}
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	return wrapPageSetup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
