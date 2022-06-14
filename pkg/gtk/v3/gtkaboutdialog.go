// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_AboutDialogClass_activate_link(void*, void*);
// extern gboolean _gotk4_gtk3_AboutDialog_ConnectActivateLink(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtkaboutdialog.go.
var (
	GTypeLicense     = coreglib.Type(C.gtk_license_get_type())
	GTypeAboutDialog = coreglib.Type(C.gtk_about_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeLicense, F: marshalLicense},
		{T: GTypeAboutDialog, F: marshalAboutDialog},
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
	// LicenseGPL20_Only: GNU General Public License, version 2.0 only. Since
	// 3.12.
	LicenseGPL20_Only
	// LicenseGPL30_Only: GNU General Public License, version 3.0 only. Since
	// 3.12.
	LicenseGPL30_Only
	// LicenseLGPL21_Only: GNU Lesser General Public License, version 2.1 only.
	// Since 3.12.
	LicenseLGPL21_Only
	// LicenseLGPL30_Only: GNU Lesser General Public License, version 3.0 only.
	// Since 3.12.
	LicenseLGPL30_Only
	// LicenseAGPL30: GNU Affero General Public License, version 3.0 or later.
	// Since: 3.22.
	LicenseAGPL30
	// LicenseAGPL30_Only: GNU Affero General Public License, version 3.0 only.
	// Since: 3.22.27.
	LicenseAGPL30_Only
	// LicenseBSD3: 3-clause BSD licence. Since: 3.24.20.
	LicenseBSD3
	// LicenseApache20: apache License, version 2.0. Since: 3.24.20.
	LicenseApache20
	// LicenseMPL20: mozilla Public License, version 2.0. Since: 3.24.20.
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

// AboutDialogOverrider contains methods that are overridable.
type AboutDialogOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	ActivateLink(uri string) bool
}

// AboutDialog offers a simple way to display information about a program like
// its logo, name, copyright, website and license. It is also possible to give
// credits to the authors, documenters, translators and artists who have worked
// on the program. An about dialog is typically opened when the user selects the
// About option from the Help menu. All parts of the dialog are optional.
//
// About dialogs often contain links and email addresses. GtkAboutDialog
// displays these as clickable links. By default, it calls
// gtk_show_uri_on_window() when a user clicks one. The behaviour can be
// overridden with the AboutDialog::activate-link signal.
//
// To specify a person with an email address, use a string like "Edgar Allan Poe
// <edgar\poe.com>". To specify a website with a title, use a string like "GTK+
// team http://www.gtk.org".
//
// To make constructing a GtkAboutDialog as convenient as possible, you can use
// the function gtk_show_about_dialog() which constructs and shows a dialog and
// keeps it around so that it can be shown again.
//
// Note that GTK+ sets a default title of _("About s") on the dialog window
// (where \s is replaced by the name of the application, but in order to ensure
// proper translation of the title, applications should set the title property
// explicitly when constructing a GtkAboutDialog, as shown in the following
// example:
//
//    GdkPixbuf *example_logo = gdk_pixbuf_new_from_file ("./logo.png", NULL);
//    gtk_show_about_dialog (NULL,
//                           "program-name", "ExampleCode",
//                           "logo", example_logo,
//                           "title", _("About ExampleCode"),
//                           NULL);
//
// It is also possible to show a AboutDialog like any other Dialog, e.g. using
// gtk_dialog_run(). In this case, you might need to know that the “Close”
// button returns the K_RESPONSE_CANCEL response id.
type AboutDialog struct {
	_ [0]func() // equal guard
	Dialog
}

var (
	_ Binner = (*AboutDialog)(nil)
)

func classInitAboutDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "AboutDialogClass")

	if _, ok := goval.(interface{ ActivateLink(uri string) bool }); ok {
		o := pclass.StructFieldOffset("activate_link")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_AboutDialogClass_activate_link)
	}
}

//export _gotk4_gtk3_AboutDialogClass_activate_link
func _gotk4_gtk3_AboutDialogClass_activate_link(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateLink(uri string) bool })

	var _uri string // out

	_uri = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := iface.ActivateLink(_uri)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapAboutDialog(obj *coreglib.Object) *AboutDialog {
	return &AboutDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: coreglib.InitiallyUnowned{
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
			},
		},
	}
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	return wrapAboutDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_AboutDialog_ConnectActivateLink
func _gotk4_gtk3_AboutDialog_ConnectActivateLink(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) (cret C.gboolean) {
	var f func(uri string) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(uri string) (ok bool))
	}

	var _uri string // out

	_uri = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := f(_uri)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectActivateLink: signal which gets emitted to activate a URI.
// Applications may connect to it to override the default behaviour, which is to
// call gtk_show_uri_on_window().
func (about *AboutDialog) ConnectActivateLink(f func(uri string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(about, "activate-link", false, unsafe.Pointer(C._gotk4_gtk3_AboutDialog_ConnectActivateLink), f)
}

// NewAboutDialog creates a new AboutDialog.
//
// The function returns the following values:
//
//    - aboutDialog: newly created AboutDialog.
//
func NewAboutDialog() *AboutDialog {
	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("new_AboutDialog", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _aboutDialog *AboutDialog // out

	_aboutDialog = wrapAboutDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _aboutDialog
}

// AddCreditSection creates a new section in the Credits page.
//
// The function takes the following parameters:
//
//    - sectionName: name of the section.
//    - people who belong to that section.
//
func (about *AboutDialog) AddCreditSection(sectionName string, people []string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(sectionName)))
	defer C.free(unsafe.Pointer(_args[1]))
	{
		*(***C.void)(unsafe.Pointer(&_args[2])) = (**C.void)(C.calloc(C.size_t((len(people) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[2]))
		{
			out := unsafe.Slice(_args[2], len(people)+1)
			var zero *C.void
			out[len(people)] = zero
			for i := range people {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(people[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("add_credit_section", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(sectionName)
	runtime.KeepAlive(people)
}

// Artists returns the string which are displayed in the artists tab of the
// secondary credits dialog.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated string array containing the artists. The array
//      is owned by the about dialog and must not be modified.
//
func (about *AboutDialog) Artists() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_artists", _args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// Authors returns the string which are displayed in the authors tab of the
// secondary credits dialog.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated string array containing the authors. The array
//      is owned by the about dialog and must not be modified.
//
func (about *AboutDialog) Authors() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_authors", _args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// Comments returns the comments string.
//
// The function returns the following values:
//
//    - utf8: comments. The string is owned by the about dialog and must not be
//      modified.
//
func (about *AboutDialog) Comments() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_comments", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Copyright returns the copyright string.
//
// The function returns the following values:
//
//    - utf8: copyright string. The string is owned by the about dialog and must
//      not be modified.
//
func (about *AboutDialog) Copyright() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_copyright", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Documenters returns the string which are displayed in the documenters tab of
// the secondary credits dialog.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated string array containing the documenters. The
//      array is owned by the about dialog and must not be modified.
//
func (about *AboutDialog) Documenters() []string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_documenters", _args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8s []string // out

	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// License returns the license information.
//
// The function returns the following values:
//
//    - utf8: license information. The string is owned by the about dialog and
//      must not be modified.
//
func (about *AboutDialog) License() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_license", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Logo returns the pixbuf displayed as logo in the about dialog.
//
// The function returns the following values:
//
//    - pixbuf displayed as logo. The pixbuf is owned by the about dialog. If you
//      want to keep a reference to it, you have to call g_object_ref() on it.
//
func (about *AboutDialog) Logo() *gdkpixbuf.Pixbuf {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_logo", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// LogoIconName returns the icon name displayed as logo in the about dialog.
//
// The function returns the following values:
//
//    - utf8: icon name displayed as logo. The string is owned by the dialog. If
//      you want to keep a reference to it, you have to call g_strdup() on it.
//
func (about *AboutDialog) LogoIconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_logo_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ProgramName returns the program name displayed in the about dialog.
//
// The function returns the following values:
//
//    - utf8: program name. The string is owned by the about dialog and must not
//      be modified.
//
func (about *AboutDialog) ProgramName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_program_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TranslatorCredits returns the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
//
// The function returns the following values:
//
//    - utf8: translator credits string. The string is owned by the about dialog
//      and must not be modified.
//
func (about *AboutDialog) TranslatorCredits() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_translator_credits", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Version returns the version string.
//
// The function returns the following values:
//
//    - utf8: version string. The string is owned by the about dialog and must
//      not be modified.
//
func (about *AboutDialog) Version() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_version", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Website returns the website URL.
//
// The function returns the following values:
//
//    - utf8: website URL. The string is owned by the about dialog and must not
//      be modified.
//
func (about *AboutDialog) Website() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_website", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebsiteLabel returns the label used for the website link.
//
// The function returns the following values:
//
//    - utf8: label used for the website link. The string is owned by the about
//      dialog and must not be modified.
//
func (about *AboutDialog) WebsiteLabel() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_website_label", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WrapLicense returns whether the license text in about is automatically
// wrapped.
//
// The function returns the following values:
//
//    - ok: TRUE if the license text is wrapped.
//
func (about *AboutDialog) WrapLicense() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_gret := girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("get_wrap_license", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(about)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetArtists sets the strings which are displayed in the artists tab of the
// secondary credits dialog.
//
// The function takes the following parameters:
//
//    - artists: NULL-terminated array of strings.
//
func (about *AboutDialog) SetArtists(artists []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(artists) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(artists)+1)
			var zero *C.void
			out[len(artists)] = zero
			for i := range artists {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(artists[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_artists", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(artists)
}

// SetAuthors sets the strings which are displayed in the authors tab of the
// secondary credits dialog.
//
// The function takes the following parameters:
//
//    - authors: NULL-terminated array of strings.
//
func (about *AboutDialog) SetAuthors(authors []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(authors) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(authors)+1)
			var zero *C.void
			out[len(authors)] = zero
			for i := range authors {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(authors[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_authors", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(authors)
}

// SetComments sets the comments string to display in the about dialog. This
// should be a short string of one or two lines.
//
// The function takes the following parameters:
//
//    - comments (optional) string.
//
func (about *AboutDialog) SetComments(comments string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if comments != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(comments)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_comments", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(comments)
}

// SetCopyright sets the copyright string to display in the about dialog. This
// should be a short string of one or two lines.
//
// The function takes the following parameters:
//
//    - copyright (optional) string.
//
func (about *AboutDialog) SetCopyright(copyright string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if copyright != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(copyright)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_copyright", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(copyright)
}

// SetDocumenters sets the strings which are displayed in the documenters tab of
// the secondary credits dialog.
//
// The function takes the following parameters:
//
//    - documenters: NULL-terminated array of strings.
//
func (about *AboutDialog) SetDocumenters(documenters []string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	{
		*(***C.void)(unsafe.Pointer(&_args[1])) = (**C.void)(C.calloc(C.size_t((len(documenters) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_args[1]))
		{
			out := unsafe.Slice(_args[1], len(documenters)+1)
			var zero *C.void
			out[len(documenters)] = zero
			for i := range documenters {
				*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(unsafe.Pointer(C.CString(documenters[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_documenters", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(documenters)
}

// SetLicense sets the license information to be displayed in the secondary
// license dialog. If license is NULL, the license button is hidden.
//
// The function takes the following parameters:
//
//    - license (optional) information or NULL.
//
func (about *AboutDialog) SetLicense(license string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if license != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(license)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_license", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(license)
}

// SetLogo sets the pixbuf to be displayed as logo in the about dialog. If it is
// NULL, the default window icon set with gtk_window_set_default_icon() will be
// used.
//
// The function takes the following parameters:
//
//    - logo (optional) or NULL.
//
func (about *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if logo != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(logo).Native()))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_logo", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(logo)
}

// SetLogoIconName sets the pixbuf to be displayed as logo in the about dialog.
// If it is NULL, the default window icon set with gtk_window_set_default_icon()
// will be used.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (about *AboutDialog) SetLogoIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if iconName != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_logo_icon_name", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(iconName)
}

// SetProgramName sets the name to display in the about dialog. If this is not
// set, it defaults to g_get_application_name().
//
// The function takes the following parameters:
//
//    - name: program name.
//
func (about *AboutDialog) SetProgramName(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_program_name", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(name)
}

// SetTranslatorCredits sets the translator credits string which is displayed in
// the translators tab of the secondary credits dialog.
//
// The intended use for this string is to display the translator of the language
// which is currently used in the user interface. Using gettext(), a simple way
// to achieve that is to mark the string for translation:
//
//    GtkWidget *about = gtk_about_dialog_new ();
//    gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
//                                             _("translator-credits"));
//
// It is a good idea to use the customary msgid “translator-credits” for this
// purpose, since translators will already know the purpose of that msgid, and
// since AboutDialog will detect if “translator-credits” is untranslated and
// hide the tab.
//
// The function takes the following parameters:
//
//    - translatorCredits (optional): translator credits.
//
func (about *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if translatorCredits != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(translatorCredits)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_translator_credits", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(translatorCredits)
}

// SetVersion sets the version string to display in the about dialog.
//
// The function takes the following parameters:
//
//    - version (optional) string.
//
func (about *AboutDialog) SetVersion(version string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if version != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(version)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_version", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(version)
}

// SetWebsite sets the URL to use for the website link.
//
// The function takes the following parameters:
//
//    - website (optional): URL string starting with "http://".
//
func (about *AboutDialog) SetWebsite(website string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if website != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(website)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_website", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(website)
}

// SetWebsiteLabel sets the label to be used for the website link.
//
// The function takes the following parameters:
//
//    - websiteLabel: label used for the website link.
//
func (about *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(websiteLabel)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_website_label", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(websiteLabel)
}

// SetWrapLicense sets whether the license text in about is automatically
// wrapped.
//
// The function takes the following parameters:
//
//    - wrapLicense: whether to wrap the license.
//
func (about *AboutDialog) SetWrapLicense(wrapLicense bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if wrapLicense {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AboutDialog").InvokeMethod("set_wrap_license", _args[:], nil)

	runtime.KeepAlive(about)
	runtime.KeepAlive(wrapLicense)
}
