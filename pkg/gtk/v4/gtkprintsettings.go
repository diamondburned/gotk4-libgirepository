// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_PrintSettingsFunc(char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_settings_get_type()), F: marshalPrintSettings},
	})
}

type PrintSettingsFunc func(key string, value string)

//export gotk4_PrintSettingsFunc
func gotk4_PrintSettingsFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key string   // out
	var value string // out

	key = C.GoString(arg0)
	value = C.GoString(arg1)

	fn := v.(PrintSettingsFunc)
	fn(key, value)
}

// PrintSettings: `GtkPrintSettings` object represents the settings of a print
// dialog in a system-independent way.
//
// The main use for this object is that once you’ve printed you can get a
// settings object that represents the settings the user chose, and the next
// time you print you can pass that object in so that the user doesn’t have to
// re-set all his settings.
//
// Its also possible to enumerate the settings so that you can easily save the
// settings for the next time your app runs, or even store them in a document.
// The predefined keys try to use shared values as much as possible so that
// moving such a document between systems still works.
type PrintSettings interface {
	gextras.Objector

	// CopyPrintSettings copies a `GtkPrintSettings` object.
	CopyPrintSettings() PrintSettings
	// ForeachPrintSettings calls @func for each key-value pair of @settings.
	ForeachPrintSettings(fn PrintSettingsFunc)
	// GetPrintSettings looks up the string value associated with @key.
	GetPrintSettings(key string) string
	// Bool returns the boolean represented by the value that is associated with
	// @key.
	//
	// The string “true” represents true, any other string false.
	Bool(key string) bool
	// Collate gets the value of GTK_PRINT_SETTINGS_COLLATE.
	Collate() bool
	// DefaultSource gets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
	DefaultSource() string
	// Dither gets the value of GTK_PRINT_SETTINGS_DITHER.
	Dither() string
	// Double returns the double value associated with @key, or 0.
	Double(key string) float64
	// DoubleWithDefault returns the floating point number represented by the
	// value that is associated with @key, or @default_val if the value does not
	// represent a floating point number.
	//
	// Floating point numbers are parsed with g_ascii_strtod().
	DoubleWithDefault(key string, def float64) float64
	// Duplex gets the value of GTK_PRINT_SETTINGS_DUPLEX.
	Duplex() PrintDuplex
	// Finishings gets the value of GTK_PRINT_SETTINGS_FINISHINGS.
	Finishings() string
	// Int returns the integer value of @key, or 0.
	Int(key string) int
	// IntWithDefault returns the value of @key, interpreted as an integer, or
	// the default value.
	IntWithDefault(key string, def int) int
	// Length returns the value associated with @key, interpreted as a length.
	//
	// The returned value is converted to @units.
	Length(key string, unit Unit) float64
	// MediaType gets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
	//
	// The set of media types is defined in PWG 5101.1-2002 PWG.
	MediaType() string
	// NCopies gets the value of GTK_PRINT_SETTINGS_N_COPIES.
	NCopies() int
	// NumberUp gets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
	NumberUp() int
	// NumberUpLayout gets the value of GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
	NumberUpLayout() NumberUpLayout
	// Orientation: get the value of GTK_PRINT_SETTINGS_ORIENTATION, converted
	// to a `GtkPageOrientation`.
	Orientation() PageOrientation
	// OutputBin gets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
	OutputBin() string
	// PageSet gets the value of GTK_PRINT_SETTINGS_PAGE_SET.
	PageSet() PageSet
	// PaperHeight gets the value of GTK_PRINT_SETTINGS_PAPER_HEIGHT, converted
	// to @unit.
	PaperHeight(unit Unit) float64
	// PaperSize gets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT, converted to
	// a `GtkPaperSize`.
	PaperSize() *PaperSize
	// PaperWidth gets the value of GTK_PRINT_SETTINGS_PAPER_WIDTH, converted to
	// @unit.
	PaperWidth(unit Unit) float64
	// PrintPages gets the value of GTK_PRINT_SETTINGS_PRINT_PAGES.
	PrintPages() PrintPages
	// Printer: convenience function to obtain the value of
	// GTK_PRINT_SETTINGS_PRINTER.
	Printer() string
	// PrinterLpi gets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
	PrinterLpi() float64
	// Quality gets the value of GTK_PRINT_SETTINGS_QUALITY.
	Quality() PrintQuality
	// Resolution gets the value of GTK_PRINT_SETTINGS_RESOLUTION.
	Resolution() int
	// ResolutionX gets the value of GTK_PRINT_SETTINGS_RESOLUTION_X.
	ResolutionX() int
	// ResolutionY gets the value of GTK_PRINT_SETTINGS_RESOLUTION_Y.
	ResolutionY() int
	// Reverse gets the value of GTK_PRINT_SETTINGS_REVERSE.
	Reverse() bool
	// Scale gets the value of GTK_PRINT_SETTINGS_SCALE.
	Scale() float64
	// UseColor gets the value of GTK_PRINT_SETTINGS_USE_COLOR.
	UseColor() bool
	// HasKeyPrintSettings returns true, if a value is associated with @key.
	HasKeyPrintSettings(key string) bool
	// LoadFilePrintSettings reads the print settings from @file_name.
	//
	// If the file could not be loaded then error is set to either a
	// `GFileError` or `GKeyFileError`.
	//
	// See [method@Gtk.PrintSettings.to_file].
	LoadFilePrintSettings(fileName string) error
	// LoadKeyFilePrintSettings reads the print settings from the group
	// @group_name in @key_file.
	//
	// If the file could not be loaded then error is set to either a
	// `GFileError` or `GKeyFileError`.
	LoadKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) error
	// SetPrintSettings associates @value with @key.
	SetPrintSettings(key string, value string)
	// SetBoolPrintSettings sets @key to a boolean value.
	SetBoolPrintSettings(key string, value bool)
	// SetCollatePrintSettings sets the value of GTK_PRINT_SETTINGS_COLLATE.
	SetCollatePrintSettings(collate bool)
	// SetDefaultSourcePrintSettings sets the value of
	// GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
	SetDefaultSourcePrintSettings(defaultSource string)
	// SetDitherPrintSettings sets the value of GTK_PRINT_SETTINGS_DITHER.
	SetDitherPrintSettings(dither string)
	// SetDoublePrintSettings sets @key to a double value.
	SetDoublePrintSettings(key string, value float64)
	// SetDuplexPrintSettings sets the value of GTK_PRINT_SETTINGS_DUPLEX.
	SetDuplexPrintSettings(duplex PrintDuplex)
	// SetFinishingsPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_FINISHINGS.
	SetFinishingsPrintSettings(finishings string)
	// SetIntPrintSettings sets @key to an integer value.
	SetIntPrintSettings(key string, value int)
	// SetLengthPrintSettings associates a length in units of @unit with @key.
	SetLengthPrintSettings(key string, value float64, unit Unit)
	// SetMediaTypePrintSettings sets the value of
	// GTK_PRINT_SETTINGS_MEDIA_TYPE.
	//
	// The set of media types is defined in PWG 5101.1-2002 PWG.
	SetMediaTypePrintSettings(mediaType string)
	// SetNCopiesPrintSettings sets the value of GTK_PRINT_SETTINGS_N_COPIES.
	SetNCopiesPrintSettings(numCopies int)
	// SetNumberUpPrintSettings sets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
	SetNumberUpPrintSettings(numberUp int)
	// SetNumberUpLayoutPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
	SetNumberUpLayoutPrintSettings(numberUpLayout NumberUpLayout)
	// SetOrientationPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_ORIENTATION.
	SetOrientationPrintSettings(orientation PageOrientation)
	// SetOutputBinPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_OUTPUT_BIN.
	SetOutputBinPrintSettings(outputBin string)
	// SetPageRangesPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PAGE_RANGES.
	SetPageRangesPrintSettings(pageRanges []PageRange)
	// SetPageSetPrintSettings sets the value of GTK_PRINT_SETTINGS_PAGE_SET.
	SetPageSetPrintSettings(pageSet PageSet)
	// SetPaperHeightPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PAPER_HEIGHT.
	SetPaperHeightPrintSettings(height float64, unit Unit)
	// SetPaperSizePrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PAPER_FORMAT, GTK_PRINT_SETTINGS_PAPER_WIDTH and
	// GTK_PRINT_SETTINGS_PAPER_HEIGHT.
	SetPaperSizePrintSettings(paperSize *PaperSize)
	// SetPaperWidthPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PAPER_WIDTH.
	SetPaperWidthPrintSettings(width float64, unit Unit)
	// SetPrintPagesPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PRINT_PAGES.
	SetPrintPagesPrintSettings(pages PrintPages)
	// SetPrinterPrintSettings: convenience function to set
	// GTK_PRINT_SETTINGS_PRINTER to @printer.
	SetPrinterPrintSettings(printer string)
	// SetPrinterLpiPrintSettings sets the value of
	// GTK_PRINT_SETTINGS_PRINTER_LPI.
	SetPrinterLpiPrintSettings(lpi float64)
	// SetQualityPrintSettings sets the value of GTK_PRINT_SETTINGS_QUALITY.
	SetQualityPrintSettings(quality PrintQuality)
	// SetResolutionPrintSettings sets the values of
	// GTK_PRINT_SETTINGS_RESOLUTION, GTK_PRINT_SETTINGS_RESOLUTION_X and
	// GTK_PRINT_SETTINGS_RESOLUTION_Y.
	SetResolutionPrintSettings(resolution int)
	// SetResolutionXYPrintSettings sets the values of
	// GTK_PRINT_SETTINGS_RESOLUTION, GTK_PRINT_SETTINGS_RESOLUTION_X and
	// GTK_PRINT_SETTINGS_RESOLUTION_Y.
	SetResolutionXYPrintSettings(resolutionX int, resolutionY int)
	// SetReversePrintSettings sets the value of GTK_PRINT_SETTINGS_REVERSE.
	SetReversePrintSettings(reverse bool)
	// SetScalePrintSettings sets the value of GTK_PRINT_SETTINGS_SCALE.
	SetScalePrintSettings(scale float64)
	// SetUseColorPrintSettings sets the value of GTK_PRINT_SETTINGS_USE_COLOR.
	SetUseColorPrintSettings(useColor bool)
	// ToFilePrintSettings: this function saves the print settings from
	// @settings to @file_name.
	//
	// If the file could not be written then error is set to either a
	// `GFileError` or `GKeyFileError`.
	ToFilePrintSettings(fileName string) error
	// ToGVariantPrintSettings: serialize print settings to an a{sv} variant.
	ToGVariantPrintSettings() *glib.Variant
	// ToKeyFilePrintSettings: this function adds the print settings from
	// @settings to @key_file.
	ToKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string)
	// UnsetPrintSettings removes any value associated with @key.
	//
	// This has the same effect as setting the value to nil.
	UnsetPrintSettings(key string)
}

// printSettings implements the PrintSettings class.
type printSettings struct {
	gextras.Objector
}

// WrapPrintSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrintSettings(obj *externglib.Object) PrintSettings {
	return printSettings{
		Objector: obj,
	}
}

func marshalPrintSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintSettings(obj), nil
}

// NewPrintSettings creates a new `GtkPrintSettings` object.
func NewPrintSettings() PrintSettings {
	var _cret *C.GtkPrintSettings // in

	_cret = C.gtk_print_settings_new()

	var _printSettings PrintSettings // out

	_printSettings = WrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printSettings
}

// NewPrintSettingsFromFile reads the print settings from @file_name.
//
// Returns a new `GtkPrintSettings` object with the restored settings, or nil if
// an error occurred. If the file could not be loaded then error is set to
// either a `GFileError` or `GKeyFileError`.
//
// See [method@Gtk.PrintSettings.to_file].
func NewPrintSettingsFromFile(fileName string) (PrintSettings, error) {
	var _arg1 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_new_from_file(_arg1, &_cerr)

	var _printSettings PrintSettings // out
	var _goerr error                 // out

	_printSettings = WrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

// NewPrintSettingsFromGVariant: deserialize print settings from an a{sv}
// variant.
//
// The variant must be in the format produced by
// [method@Gtk.PrintSettings.to_gvariant].
func NewPrintSettingsFromGVariant(variant *glib.Variant) PrintSettings {
	var _arg1 *C.GVariant         // out
	var _cret *C.GtkPrintSettings // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant))

	_cret = C.gtk_print_settings_new_from_gvariant(_arg1)

	var _printSettings PrintSettings // out

	_printSettings = WrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printSettings
}

// NewPrintSettingsFromKeyFile reads the print settings from the group
// @group_name in @key_file.
//
// Returns a new `GtkPrintSettings` object with the restored settings, or nil if
// an error occurred. If the file could not be loaded then error is set to
// either `GFileError` or `GKeyFileError`.
func NewPrintSettingsFromKeyFile(keyFile *glib.KeyFile, groupName string) (PrintSettings, error) {
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_print_settings_new_from_key_file(_arg1, _arg2, &_cerr)

	var _printSettings PrintSettings // out
	var _goerr error                 // out

	_printSettings = WrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

func (o printSettings) CopyPrintSettings() PrintSettings {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPrintSettings // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_settings_copy(_arg0)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func (s printSettings) ForeachPrintSettings(fn PrintSettingsFunc) {
	var _arg0 *C.GtkPrintSettings    // out
	var _arg1 C.GtkPrintSettingsFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*[0]byte)(C.gotk4_PrintSettingsFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_print_settings_foreach(_arg0, _arg1, _arg2)
}

func (s printSettings) GetPrintSettings(key string) string {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Bool(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_bool(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) Collate() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_collate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) DefaultSource() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_default_source(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Dither() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_dither(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Double(key string) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_double(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) DoubleWithDefault(key string, def float64) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(def)

	_cret = C.gtk_print_settings_get_double_with_default(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) Duplex() PrintDuplex {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintDuplex    // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_duplex(_arg0)

	var _printDuplex PrintDuplex // out

	_printDuplex = PrintDuplex(_cret)

	return _printDuplex
}

func (s printSettings) Finishings() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_finishings(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Int(key string) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_int(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) IntWithDefault(key string, def int) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(def)

	_cret = C.gtk_print_settings_get_int_with_default(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) Length(key string, unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_length(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) MediaType() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_media_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) NCopies() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_n_copies(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) NumberUp() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_number_up(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) NumberUpLayout() NumberUpLayout {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkNumberUpLayout // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_number_up_layout(_arg0)

	var _numberUpLayout NumberUpLayout // out

	_numberUpLayout = NumberUpLayout(_cret)

	return _numberUpLayout
}

func (s printSettings) Orientation() PageOrientation {
	var _arg0 *C.GtkPrintSettings  // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

func (s printSettings) OutputBin() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_output_bin(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) PageSet() PageSet {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPageSet        // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_page_set(_arg0)

	var _pageSet PageSet // out

	_pageSet = PageSet(_cret)

	return _pageSet
}

func (s printSettings) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) PaperSize() *PaperSize {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPaperSize     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(unsafe.Pointer(v)))
	})

	return _paperSize
}

func (s printSettings) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) PrintPages() PrintPages {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintPages     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_print_pages(_arg0)

	var _printPages PrintPages // out

	_printPages = PrintPages(_cret)

	return _printPages
}

func (s printSettings) Printer() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_printer(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) PrinterLpi() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_printer_lpi(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) Quality() PrintQuality {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintQuality   // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_quality(_arg0)

	var _printQuality PrintQuality // out

	_printQuality = PrintQuality(_cret)

	return _printQuality
}

func (s printSettings) Resolution() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) ResolutionX() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution_x(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) ResolutionY() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution_y(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) Reverse() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_reverse(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) Scale() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_scale(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) UseColor() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_use_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) HasKeyPrintSettings(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) LoadFilePrintSettings(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) LoadKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) SetPrintSettings(key string, value string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_set(_arg0, _arg1, _arg2)
}

func (s printSettings) SetBoolPrintSettings(key string, value bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	C.gtk_print_settings_set_bool(_arg0, _arg1, _arg2)
}

func (s printSettings) SetCollatePrintSettings(collate bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if collate {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_collate(_arg0, _arg1)
}

func (s printSettings) SetDefaultSourcePrintSettings(defaultSource string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(defaultSource))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_default_source(_arg0, _arg1)
}

func (s printSettings) SetDitherPrintSettings(dither string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(dither))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_dither(_arg0, _arg1)
}

func (s printSettings) SetDoublePrintSettings(key string, value float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)

	C.gtk_print_settings_set_double(_arg0, _arg1, _arg2)
}

func (s printSettings) SetDuplexPrintSettings(duplex PrintDuplex) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintDuplex    // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintDuplex(duplex)

	C.gtk_print_settings_set_duplex(_arg0, _arg1)
}

func (s printSettings) SetFinishingsPrintSettings(finishings string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(finishings))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_finishings(_arg0, _arg1)
}

func (s printSettings) SetIntPrintSettings(key string, value int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(value)

	C.gtk_print_settings_set_int(_arg0, _arg1, _arg2)
}

func (s printSettings) SetLengthPrintSettings(key string, value float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _arg3 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)
	_arg3 = C.GtkUnit(unit)

	C.gtk_print_settings_set_length(_arg0, _arg1, _arg2, _arg3)
}

func (s printSettings) SetMediaTypePrintSettings(mediaType string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(mediaType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_media_type(_arg0, _arg1)
}

func (s printSettings) SetNCopiesPrintSettings(numCopies int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(numCopies)

	C.gtk_print_settings_set_n_copies(_arg0, _arg1)
}

func (s printSettings) SetNumberUpPrintSettings(numberUp int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(numberUp)

	C.gtk_print_settings_set_number_up(_arg0, _arg1)
}

func (s printSettings) SetNumberUpLayoutPrintSettings(numberUpLayout NumberUpLayout) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkNumberUpLayout // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkNumberUpLayout(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout(_arg0, _arg1)
}

func (s printSettings) SetOrientationPrintSettings(orientation PageOrientation) {
	var _arg0 *C.GtkPrintSettings  // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPageOrientation(orientation)

	C.gtk_print_settings_set_orientation(_arg0, _arg1)
}

func (s printSettings) SetOutputBinPrintSettings(outputBin string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(outputBin))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_output_bin(_arg0, _arg1)
}

func (s printSettings) SetPageRangesPrintSettings(pageRanges []PageRange) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPageRange
	var _arg2 C.int

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg2 = C.int(len(pageRanges))
	_arg1 = (*C.GtkPageRange)(unsafe.Pointer(&pageRanges[0]))

	C.gtk_print_settings_set_page_ranges(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPageSetPrintSettings(pageSet PageSet) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPageSet        // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPageSet(pageSet)

	C.gtk_print_settings_set_page_set(_arg0, _arg1)
}

func (s printSettings) SetPaperHeightPrintSettings(height float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(height)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_height(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPaperSizePrintSettings(paperSize *PaperSize) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPaperSize     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(paperSize))

	C.gtk_print_settings_set_paper_size(_arg0, _arg1)
}

func (s printSettings) SetPaperWidthPrintSettings(width float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(width)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_width(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPrintPagesPrintSettings(pages PrintPages) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintPages     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintPages(pages)

	C.gtk_print_settings_set_print_pages(_arg0, _arg1)
}

func (s printSettings) SetPrinterPrintSettings(printer string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(printer))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_printer(_arg0, _arg1)
}

func (s printSettings) SetPrinterLpiPrintSettings(lpi float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(lpi)

	C.gtk_print_settings_set_printer_lpi(_arg0, _arg1)
}

func (s printSettings) SetQualityPrintSettings(quality PrintQuality) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintQuality   // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintQuality(quality)

	C.gtk_print_settings_set_quality(_arg0, _arg1)
}

func (s printSettings) SetResolutionPrintSettings(resolution int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(resolution)

	C.gtk_print_settings_set_resolution(_arg0, _arg1)
}

func (s printSettings) SetResolutionXYPrintSettings(resolutionX int, resolutionY int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(resolutionX)
	_arg2 = C.int(resolutionY)

	C.gtk_print_settings_set_resolution_xy(_arg0, _arg1, _arg2)
}

func (s printSettings) SetReversePrintSettings(reverse bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_reverse(_arg0, _arg1)
}

func (s printSettings) SetScalePrintSettings(scale float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(scale)

	C.gtk_print_settings_set_scale(_arg0, _arg1)
}

func (s printSettings) SetUseColorPrintSettings(useColor bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if useColor {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_use_color(_arg0, _arg1)
}

func (s printSettings) ToFilePrintSettings(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) ToGVariantPrintSettings() *glib.Variant {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GVariant         // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)

	return _variant
}

func (s printSettings) ToKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_to_key_file(_arg0, _arg1, _arg2)
}

func (s printSettings) UnsetPrintSettings(key string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_unset(_arg0, _arg1)
}

// PageRange: range of pages to print.
//
// See also [method@Gtk.PrintSettings.set_page_ranges].
type PageRange struct {
	native C.GtkPageRange
}

// WrapPageRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPageRange(ptr unsafe.Pointer) *PageRange {
	return (*PageRange)(ptr)
}

// Native returns the underlying C source pointer.
func (p *PageRange) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
