// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_paper_size_get_type()), F: marshalPaperSize},
	})
}

// PaperSize: gtkPaperSize handles paper sizes. It uses the standard called PWG
// 5101.1-2002 PWG: Standard for Media Standardized Names
// (http://www.pwg.org/standards.html) to name the paper sizes (and to get the
// data for the page sizes). In addition to standard paper sizes, GtkPaperSize
// allows to construct custom paper sizes with arbitrary dimensions.
//
// The PaperSize object stores not only the dimensions (width and height) of a
// paper size and its name, it also provides default [print
// margins][print-margins].
//
// Printing support has been added in GTK+ 2.10.
type PaperSize struct {
	native C.GtkPaperSize
}

// WrapPaperSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPaperSize(ptr unsafe.Pointer) *PaperSize {
	if ptr == nil {
		return nil
	}

	return (*PaperSize)(ptr)
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPaperSize(unsafe.Pointer(b)), nil
}

// NewPaperSize constructs a struct PaperSize.
func NewPaperSize(name string) *PaperSize {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_new(_arg1)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// NewPaperSizeCustom constructs a struct PaperSize.
func NewPaperSizeCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _arg3 C.gdouble
	var _arg4 C.gdouble
	var _arg5 C.GtkUnit

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(displayName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gdouble(width)
	_arg4 = C.gdouble(height)
	_arg5 = (C.GtkUnit)(unit)

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_new_custom(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// NewPaperSizeFromGVariant constructs a struct PaperSize.
func NewPaperSizeFromGVariant(variant *glib.Variant) *PaperSize {
	var _arg1 *C.GVariant

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_new_from_gvariant(_arg1)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// NewPaperSizeFromIpp constructs a struct PaperSize.
func NewPaperSizeFromIpp(ippName string, width float64, height float64) *PaperSize {
	var _arg1 *C.gchar
	var _arg2 C.gdouble
	var _arg3 C.gdouble

	_arg1 = (*C.gchar)(C.CString(ippName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(width)
	_arg3 = C.gdouble(height)

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_new_from_ipp(_arg1, _arg2, _arg3)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// NewPaperSizeFromKeyFile constructs a struct PaperSize.
func NewPaperSizeFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PaperSize, error) {
	var _arg1 *C.GKeyFile
	var _arg2 *C.gchar

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret *C.GtkPaperSize
	var _cerr *C.GError

	cret = C.gtk_paper_size_new_from_key_file(_arg1, _arg2, _cerr)

	var _paperSize *PaperSize
	var _goerr error

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _paperSize, _goerr
}

// NewPaperSizeFromPpd constructs a struct PaperSize.
func NewPaperSizeFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	var _arg1 *C.gchar
	var _arg2 *C.gchar
	var _arg3 C.gdouble
	var _arg4 C.gdouble

	_arg1 = (*C.gchar)(C.CString(ppdName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(ppdDisplayName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gdouble(width)
	_arg4 = C.gdouble(height)

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_new_from_ppd(_arg1, _arg2, _arg3, _arg4)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// Native returns the underlying C source pointer.
func (p *PaperSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Copy copies an existing PaperSize.
func (o *PaperSize) Copy() *PaperSize {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(o.Native()))

	var _cret *C.GtkPaperSize

	cret = C.gtk_paper_size_copy(_arg0)

	var _paperSize *PaperSize

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _paperSize
}

// Free: free the given PaperSize object.
func (s *PaperSize) Free() {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	C.gtk_paper_size_free(_arg0)
}

// DefaultBottomMargin gets the default bottom margin for the PaperSize.
func (s *PaperSize) DefaultBottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_default_bottom_margin(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// DefaultLeftMargin gets the default left margin for the PaperSize.
func (s *PaperSize) DefaultLeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_default_left_margin(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// DefaultRightMargin gets the default right margin for the PaperSize.
func (s *PaperSize) DefaultRightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_default_right_margin(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// DefaultTopMargin gets the default top margin for the PaperSize.
func (s *PaperSize) DefaultTopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_default_top_margin(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// DisplayName gets the human-readable name of the PaperSize.
func (s *PaperSize) DisplayName() string {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar

	cret = C.gtk_paper_size_get_display_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Height gets the paper height of the PaperSize, in units of @unit.
func (s *PaperSize) Height(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_height(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Name gets the name of the PaperSize.
func (s *PaperSize) Name() string {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar

	cret = C.gtk_paper_size_get_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PpdName gets the PPD name of the PaperSize, which may be nil.
func (s *PaperSize) PpdName() string {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar

	cret = C.gtk_paper_size_get_ppd_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Width gets the paper width of the PaperSize, in units of @unit.
func (s *PaperSize) Width(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble

	cret = C.gtk_paper_size_get_width(_arg0, _arg1)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// IsCustom returns true if @size is not a standard paper size.
func (s *PaperSize) IsCustom() bool {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_paper_size_is_custom(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsEqual compares two PaperSize objects.
func (s *PaperSize) IsEqual(size2 *PaperSize) bool {
	var _arg0 *C.GtkPaperSize
	var _arg1 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size2.Native()))

	var _cret C.gboolean

	cret = C.gtk_paper_size_is_equal(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsIpp returns true if @size is an IPP standard paper size.
func (s *PaperSize) IsIpp() bool {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gtk_paper_size_is_ipp(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetSize changes the dimensions of a @size to @width x @height.
func (s *PaperSize) SetSize(width float64, height float64, unit Unit) {
	var _arg0 *C.GtkPaperSize
	var _arg1 C.gdouble
	var _arg2 C.gdouble
	var _arg3 C.GtkUnit

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(width)
	_arg2 = C.gdouble(height)
	_arg3 = (C.GtkUnit)(unit)

	C.gtk_paper_size_set_size(_arg0, _arg1, _arg2, _arg3)
}

// ToGVariant: serialize a paper size to an a{sv} variant.
func (p *PaperSize) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPaperSize

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(p.Native()))

	var _cret *C.GVariant

	cret = C.gtk_paper_size_to_gvariant(_arg0)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))

	return _variant
}

// ToKeyFile: this function adds the paper size from @size to @key_file.
func (s *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPaperSize
	var _arg1 *C.GKeyFile
	var _arg2 *C.gchar

	_arg0 = (*C.GtkPaperSize)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_paper_size_to_key_file(_arg0, _arg1, _arg2)
}
