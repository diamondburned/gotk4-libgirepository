// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_page_setup_get_type()), F: marshalPageSetup},
	})
}

// PageSetup object stores the page size, orientation and margins. The idea is
// that you can get one of these from the page setup dialog and then pass it to
// the PrintOperation when printing. The benefit of splitting this out of the
// PrintSettings is that these affect the actual layout of the page, and thus
// need to be set long before user prints.
//
//
// Margins
//
// The margins specified in this object are the “print margins”, i.e. the parts
// of the page that the printer cannot print on. These are different from the
// layout margins that a word processor uses; they are typically used to
// determine the minimal size for the layout margins.
//
// To obtain a PageSetup use gtk_page_setup_new() to get the defaults, or use
// gtk_print_run_page_setup_dialog() to show the page setup dialog and receive
// the resulting page setup.
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
//    }
//
// Printing support was added in GTK+ 2.10.
type PageSetup interface {
	gextras.Objector

	// CopyPageSetup copies a PageSetup.
	CopyPageSetup() PageSetup
	// BottomMargin gets the bottom margin in units of @unit.
	BottomMargin(unit Unit) float64
	// LeftMargin gets the left margin in units of @unit.
	LeftMargin(unit Unit) float64
	// Orientation gets the page orientation of the PageSetup.
	Orientation() PageOrientation
	// PageHeight returns the page height in units of @unit.
	//
	// Note that this function takes orientation and margins into consideration.
	// See gtk_page_setup_get_paper_height().
	PageHeight(unit Unit) float64
	// PageWidth returns the page width in units of @unit.
	//
	// Note that this function takes orientation and margins into consideration.
	// See gtk_page_setup_get_paper_width().
	PageWidth(unit Unit) float64
	// PaperHeight returns the paper height in units of @unit.
	//
	// Note that this function takes orientation, but not margins into
	// consideration. See gtk_page_setup_get_page_height().
	PaperHeight(unit Unit) float64
	// PaperSize gets the paper size of the PageSetup.
	PaperSize() *PaperSize
	// PaperWidth returns the paper width in units of @unit.
	//
	// Note that this function takes orientation, but not margins into
	// consideration. See gtk_page_setup_get_page_width().
	PaperWidth(unit Unit) float64
	// RightMargin gets the right margin in units of @unit.
	RightMargin(unit Unit) float64
	// TopMargin gets the top margin in units of @unit.
	TopMargin(unit Unit) float64
	// LoadFilePageSetup reads the page setup from the file @file_name. See
	// gtk_page_setup_to_file().
	LoadFilePageSetup(fileName string) error
	// LoadKeyFilePageSetup reads the page setup from the group @group_name in
	// the key file @key_file.
	LoadKeyFilePageSetup(keyFile *glib.KeyFile, groupName string) error
	// SetBottomMarginPageSetup sets the bottom margin of the PageSetup.
	SetBottomMarginPageSetup(margin float64, unit Unit)
	// SetLeftMarginPageSetup sets the left margin of the PageSetup.
	SetLeftMarginPageSetup(margin float64, unit Unit)
	// SetOrientationPageSetup sets the page orientation of the PageSetup.
	SetOrientationPageSetup(orientation PageOrientation)
	// SetPaperSizePageSetup sets the paper size of the PageSetup without
	// changing the margins. See
	// gtk_page_setup_set_paper_size_and_default_margins().
	SetPaperSizePageSetup(size *PaperSize)
	// SetPaperSizeAndDefaultMarginsPageSetup sets the paper size of the
	// PageSetup and modifies the margins according to the new paper size.
	SetPaperSizeAndDefaultMarginsPageSetup(size *PaperSize)
	// SetRightMarginPageSetup sets the right margin of the PageSetup.
	SetRightMarginPageSetup(margin float64, unit Unit)
	// SetTopMarginPageSetup sets the top margin of the PageSetup.
	SetTopMarginPageSetup(margin float64, unit Unit)
	// ToFilePageSetup: this function saves the information from @setup to
	// @file_name.
	ToFilePageSetup(fileName string) error
	// ToGVariantPageSetup: serialize page setup to an a{sv} variant.
	ToGVariantPageSetup() *glib.Variant
	// ToKeyFilePageSetup: this function adds the page setup from @setup to
	// @key_file.
	ToKeyFilePageSetup(keyFile *glib.KeyFile, groupName string)
}

// pageSetup implements the PageSetup class.
type pageSetup struct {
	gextras.Objector
}

// WrapPageSetup wraps a GObject to the right type. It is
// primarily used internally.
func WrapPageSetup(obj *externglib.Object) PageSetup {
	return pageSetup{
		Objector: obj,
	}
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPageSetup(obj), nil
}

// NewPageSetup creates a new PageSetup.
func NewPageSetup() PageSetup {
	var _cret *C.GtkPageSetup // in

	_cret = C.gtk_page_setup_new()

	var _pageSetup PageSetup // out

	_pageSetup = WrapPageSetup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromFile reads the page setup from the file @file_name. Returns a
// new PageSetup object with the restored page setup, or nil if an error
// occurred. See gtk_page_setup_to_file().
func NewPageSetupFromFile(fileName string) (PageSetup, error) {
	var _arg1 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_page_setup_new_from_file(_arg1, &_cerr)

	var _pageSetup PageSetup // out
	var _goerr error         // out

	_pageSetup = WrapPageSetup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pageSetup, _goerr
}

// NewPageSetupFromGVariant: desrialize a page setup from an a{sv} variant in
// the format produced by gtk_page_setup_to_gvariant().
func NewPageSetupFromGVariant(variant *glib.Variant) PageSetup {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPageSetup // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant))

	_cret = C.gtk_page_setup_new_from_gvariant(_arg1)

	var _pageSetup PageSetup // out

	_pageSetup = WrapPageSetup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromKeyFile reads the page setup from the group @group_name in
// the key file @key_file. Returns a new PageSetup object with the restored page
// setup, or nil if an error occurred.
func NewPageSetupFromKeyFile(keyFile *glib.KeyFile, groupName string) (PageSetup, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_page_setup_new_from_key_file(_arg1, _arg2, &_cerr)

	var _pageSetup PageSetup // out
	var _goerr error         // out

	_pageSetup = WrapPageSetup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pageSetup, _goerr
}

func (o pageSetup) CopyPageSetup() PageSetup {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPageSetup // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_page_setup_copy(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

func (s pageSetup) BottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_bottom_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) LeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_left_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) Orientation() PageOrientation {
	var _arg0 *C.GtkPageSetup      // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_page_setup_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

func (s pageSetup) PageHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_page_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) PageWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_page_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_paper_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) PaperSize() *PaperSize {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_page_setup_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(unsafe.Pointer(_cret))

	return _paperSize
}

func (s pageSetup) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_paper_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) RightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_right_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) TopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_top_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s pageSetup) LoadFilePageSetup(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s pageSetup) LoadKeyFilePageSetup(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_page_setup_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s pageSetup) SetBottomMarginPageSetup(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_bottom_margin(_arg0, _arg1, _arg2)
}

func (s pageSetup) SetLeftMarginPageSetup(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_left_margin(_arg0, _arg1, _arg2)
}

func (s pageSetup) SetOrientationPageSetup(orientation PageOrientation) {
	var _arg0 *C.GtkPageSetup      // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPageOrientation(orientation)

	C.gtk_page_setup_set_orientation(_arg0, _arg1)
}

func (s pageSetup) SetPaperSizePageSetup(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size))

	C.gtk_page_setup_set_paper_size(_arg0, _arg1)
}

func (s pageSetup) SetPaperSizeAndDefaultMarginsPageSetup(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size))

	C.gtk_page_setup_set_paper_size_and_default_margins(_arg0, _arg1)
}

func (s pageSetup) SetRightMarginPageSetup(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_right_margin(_arg0, _arg1, _arg2)
}

func (s pageSetup) SetTopMarginPageSetup(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_top_margin(_arg0, _arg1, _arg2)
}

func (s pageSetup) ToFilePageSetup(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s pageSetup) ToGVariantPageSetup() *glib.Variant {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_page_setup_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)

	return _variant
}

func (s pageSetup) ToKeyFilePageSetup(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_page_setup_to_key_file(_arg0, _arg1, _arg2)
}
