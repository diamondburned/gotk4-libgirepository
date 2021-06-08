// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// EndLine returns the line in the CSS document where this section end. The line
// number is 0-indexed, so the first line of the document will return 0. This
// value may change in future invocations of this function if @section is not
// yet parsed completely. This will for example happen in the
// GtkCssProvider::parsing-error signal. The end position and line may be
// identical to the start position and line for sections which failed to parse
// anything successfully.
func (s *CSSSection) EndLine() uint {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_css_section_get_end_line(arg0)

	goret = uint(cret)

	return goret
}

// EndPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_end_line(). This value may change in future
// invocations of this function if @section is not yet parsed completely. This
// will for example happen in the GtkCssProvider::parsing-error signal. The end
// position and line may be identical to the start position and line for
// sections which failed to parse anything successfully.
func (s *CSSSection) EndPosition() uint {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_css_section_get_end_position(arg0)

	goret = uint(cret)

	return goret
}

// File gets the file that @section was parsed from. If no such file exists, for
// example because the CSS was loaded via @gtk_css_provider_load_from_data(),
// then nil is returned.
func (s *CSSSection) File() gio.File {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret *C.GFile
	var goret gio.File

	cret = C.gtk_css_section_get_file(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gio.File)

	return goret
}

// Parent gets the parent section for the given @section. The parent section is
// the section that contains this @section. A special case are sections of type
// K_CSS_SECTION_DOCUMENT. Their parent will either be nil if they are the
// original CSS document that was loaded by gtk_css_provider_load_from_file() or
// a section of type K_CSS_SECTION_IMPORT if it was loaded with an import rule
// from a different file.
func (s *CSSSection) Parent() *CSSSection {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret *C.GtkCssSection
	var goret *CSSSection

	cret = C.gtk_css_section_get_parent(arg0)

	goret = WrapCSSSection(unsafe.Pointer(cret))

	return goret
}

// SectionType gets the type of information that @section describes.
func (s *CSSSection) SectionType() CSSSectionType {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret C.GtkCssSectionType
	var goret CSSSectionType

	cret = C.gtk_css_section_get_section_type(arg0)

	goret = CSSSectionType(cret)

	return goret
}

// StartLine returns the line in the CSS document where this section starts. The
// line number is 0-indexed, so the first line of the document will return 0.
func (s *CSSSection) StartLine() uint {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_css_section_get_start_line(arg0)

	goret = uint(cret)

	return goret
}

// StartPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_start_line().
func (s *CSSSection) StartPosition() uint {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_css_section_get_start_position(arg0)

	goret = uint(cret)

	return goret
}

// Ref increments the reference count on @section.
func (s *CSSSection) Ref() *CSSSection {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	cret := new(C.GtkCssSection)
	var goret *CSSSection

	cret = C.gtk_css_section_ref(arg0)

	goret = WrapCSSSection(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *CSSSection) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Unref decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Unref() {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	C.gtk_css_section_unref(arg0)
}
