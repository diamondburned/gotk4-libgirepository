// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSection defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the containing
// region.
type CSSSection struct {
	native C.GtkCssSection
}

// WrapCSSSection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCSSSection(ptr unsafe.Pointer) *CSSSection {
	if ptr == nil {
		return nil
	}

	return (*CSSSection)(ptr)
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCSSSection(unsafe.Pointer(b)), nil
}

// NewCSSSection constructs a struct CSSSection.
func NewCSSSection(file gio.File, start *CSSLocation, end *CSSLocation) *CSSSection {
	var _arg1 *C.GFile
	var _arg2 *C.GtkCssLocation
	var _arg3 *C.GtkCssLocation

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))
	_arg2 = (*C.GtkCssLocation)(unsafe.Pointer(start.Native()))
	_arg3 = (*C.GtkCssLocation)(unsafe.Pointer(end.Native()))

	var _cret *C.GtkCssSection

	cret = C.gtk_css_section_new(_arg1, _arg2, _arg3)

	var _cssSection *CSSSection

	_cssSection = WrapCSSSection(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _cssSection
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// EndLocation returns the location in the CSS document where this section ends.
func (s *CSSSection) EndLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkCssLocation

	cret = C.gtk_css_section_get_end_location(_arg0)

	var _cssLocation *CSSLocation

	_cssLocation = WrapCSSLocation(unsafe.Pointer(_cret))

	return _cssLocation
}

// File gets the file that @section was parsed from. If no such file exists, for
// example because the CSS was loaded via @gtk_css_provider_load_from_data(),
// then nil is returned.
func (s *CSSSection) File() gio.File {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.GFile

	cret = C.gtk_css_section_get_file(_arg0)

	var _file gio.File

	_file = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.File)

	return _file
}

// Parent gets the parent section for the given @section. The parent section is
// the section that contains this @section. A special case are sections of type
// K_CSS_SECTION_DOCUMENT. Their parent will either be nil if they are the
// original CSS document that was loaded by gtk_css_provider_load_from_file() or
// a section of type K_CSS_SECTION_IMPORT if it was loaded with an import rule
// from a different file.
func (s *CSSSection) Parent() *CSSSection {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkCssSection

	cret = C.gtk_css_section_get_parent(_arg0)

	var _cssSection *CSSSection

	_cssSection = WrapCSSSection(unsafe.Pointer(_cret))

	return _cssSection
}

// StartLocation returns the location in the CSS document where this section
// starts.
func (s *CSSSection) StartLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkCssLocation

	cret = C.gtk_css_section_get_start_location(_arg0)

	var _cssLocation *CSSLocation

	_cssLocation = WrapCSSLocation(unsafe.Pointer(_cret))

	return _cssLocation
}

// Print prints the @section into @string in a human-readable form. This is a
// form like `gtk.css:32:1-23` to denote line 32, characters 1 to 23 in the file
// gtk.css.
func (s *CSSSection) Print(string *glib.String) {
	var _arg0 *C.GtkCssSection
	var _arg1 *C.GString

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gtk_css_section_print(_arg0, _arg1)
}

// Ref increments the reference count on @section.
func (s *CSSSection) Ref() *CSSSection {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkCssSection

	cret = C.gtk_css_section_ref(_arg0)

	var _cssSection *CSSSection

	_cssSection = WrapCSSSection(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _cssSection
}

// String prints the section into a human-readable text form using
// gtk_css_section_print().
func (s *CSSSection) String() string {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	cret = C.gtk_css_section_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Unref() {
	var _arg0 *C.GtkCssSection

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	C.gtk_css_section_unref(_arg0)
}
