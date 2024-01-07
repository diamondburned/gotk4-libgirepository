// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk4_AboutDialog_ConnectActivateLink(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeLicense     = coreglib.Type(girepository.MustFind("Gtk", "License").RegisteredGType())
	GTypeAboutDialog = coreglib.Type(girepository.MustFind("Gtk", "AboutDialog").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLicense, F: marshalLicense},
		coreglib.TypeMarshaler{T: GTypeAboutDialog, F: marshalAboutDialog},
	})
}

// License: type of license for an application.
//
// This enumeration can be expanded at later date.
type License C.gint

const (
	// LicenseUnknown: no license specified.
	LicenseUnknown License = iota
	// LicenseCustom: license text is going to be specified by the developer.
	LicenseCustom
	// LicenseGPL20: GNU General Public License, version 2.0 or later.
	LicenseGPL20
	// LicenseGPL30: GNU General Public License, version 3.0 or later.
	LicenseGPL30
	// LicenseLGPL21: GNU Lesser General Public License, version 2.1 or later.
	LicenseLGPL21
	// LicenseLGPL30: GNU Lesser General Public License, version 3.0 or later.
	LicenseLGPL30
	// LicenseBSD: BSD standard license.
	LicenseBSD
	// LicenseMITX11: MIT/X11 standard license.
	LicenseMITX11
	// LicenseArtistic: artistic License, version 2.0.
	LicenseArtistic
	// LicenseGPL20_Only: GNU General Public License, version 2.0 only.
	LicenseGPL20_Only
	// LicenseGPL30_Only: GNU General Public License, version 3.0 only.
	LicenseGPL30_Only
	// LicenseLGPL21_Only: GNU Lesser General Public License, version 2.1 only.
	LicenseLGPL21_Only
	// LicenseLGPL30_Only: GNU Lesser General Public License, version 3.0 only.
	LicenseLGPL30_Only
	// LicenseAGPL30: GNU Affero General Public License, version 3.0 or later.
	LicenseAGPL30
	// LicenseAGPL30_Only: GNU Affero General Public License, version 3.0 only.
	LicenseAGPL30_Only
	// LicenseBSD3: 3-clause BSD licence.
	LicenseBSD3
	// LicenseApache20: apache License, version 2.0.
	LicenseApache20
	// LicenseMPL20: mozilla Public License, version 2.0.
	LicenseMPL20
)

func marshalLicense(p uintptr) (interface{}, error) {
	return License(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for License.
func (l License) String() string {
	switch l {
	case LicenseUnknown:
		return "Unknown"
	case LicenseCustom:
		return "Custom"
	case LicenseGPL20:
		return "GPL20"
	case LicenseGPL30:
		return "GPL30"
	case LicenseLGPL21:
		return "LGPL21"
	case LicenseLGPL30:
		return "LGPL30"
	case LicenseBSD:
		return "BSD"
	case LicenseMITX11:
		return "MITX11"
	case LicenseArtistic:
		return "Artistic"
	case LicenseGPL20_Only:
		return "GPL20_Only"
	case LicenseGPL30_Only:
		return "GPL30_Only"
	case LicenseLGPL21_Only:
		return "LGPL21_Only"
	case LicenseLGPL30_Only:
		return "LGPL30_Only"
	case LicenseAGPL30:
		return "AGPL30"
	case LicenseAGPL30_Only:
		return "AGPL30_Only"
	case LicenseBSD3:
		return "BSD3"
	case LicenseApache20:
		return "Apache20"
	case LicenseMPL20:
		return "MPL20"
	default:
		return fmt.Sprintf("License(%d)", l)
	}
}

// AboutDialog: GtkAboutDialog offers a simple way to display information about
// a program.
//
// The shown information includes the programs' logo, name, copyright, website
// and license. It is also possible to give credits to the authors, documenters,
// translators and artists who have worked on the program.
//
// An about dialog is typically opened when the user selects the About option
// from the Help menu. All parts of the dialog are optional.
//
// !An example GtkAboutDialog (aboutdialog.png)
//
// About dialogs often contain links and email addresses. GtkAboutDialog
// displays these as clickable links. By default, it calls gtk.ShowURI() when a
// user clicks one. The behaviour can be overridden with the
// gtk.AboutDialog::activate-link signal.
//
// To specify a person with an email address, use a string like Edgar Allan Poe
// <edgarpoe.com>. To specify a website with a title, use a string like GTK team
// https://www.gtk.org.
//
// To make constructing a GtkAboutDialog as convenient as possible, you can use
// the function gtk.ShowAboutDialog() which constructs and shows a dialog and
// keeps it around so that it can be shown again.
//
// Note that GTK sets a default title of _("About s") on the dialog window
// (where s is replaced by the name of the application, but in order to ensure
// proper translation of the title, applications should set the title property
// explicitly when constructing a GtkAboutDialog, as shown in the following
// example:
//
//    GFile *logo_file = g_file_new_for_path ("./logo.png");
//    GdkTexture *example_logo = gdk_texture_new_from_file (logo_file, NULL);
//    g_object_unref (logo_file);
//
//    gtk_show_about_dialog (NULL,
//                           "program-name", "ExampleCode",
//                           "logo", example_logo,
//                           "title", _("About ExampleCode"),
//                           NULL);
//
//
//
// CSS nodes
//
// GtkAboutDialog has a single CSS node with the name window and style class
// .aboutdialog.
type AboutDialog struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Widgetter         = (*AboutDialog)(nil)
	_ coreglib.Objector = (*AboutDialog)(nil)
)

func wrapAboutDialog(obj *coreglib.Object) *AboutDialog {
	return &AboutDialog{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						Accessible: Accessible{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						ConstraintTarget: ConstraintTarget{
							Object: obj,
						},
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	return wrapAboutDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivateLink is emitted every time a URL is activated.
//
// Applications may connect to it to override the default behaviour, which is to
// call gtk.ShowURI().
func (v *AboutDialog) ConnectActivateLink(f func(uri string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "activate-link", false, unsafe.Pointer(C._gotk4_gtk4_AboutDialog_ConnectActivateLink), f)
}
