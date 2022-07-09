// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeCSSSection returns the GType for the type CSSSection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCSSSection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CssSection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCSSSection)
	return gtype
}

// CSSSection defines a part of a CSS document.
//
// Because sections are nested into one another, you can use
// gtk_css_section_get_parent() to get the containing region.
//
// An instance of this type is always passed by reference.
type CSSSection struct {
	*cssSection
}

// cssSection is the struct that's finalized.
type cssSection struct {
	native unsafe.Pointer
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &CSSSection{&cssSection{(unsafe.Pointer)(b)}}, nil
}

// NewCSSSection constructs a struct CSSSection.
func NewCSSSection(file gio.Filer, start *CSSLocation, end *CSSLocation) *CSSSection {
	var _args [3]girepository.Argument

	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(start)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(end)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(file)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_cssSection)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _cssSection
}

// EndLocation returns the location in the CSS document where this section ends.
//
// The function returns the following values:
//
//    - cssLocation: end location of this section.
//
func (section *CSSSection) EndLocation() *CSSLocation {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _cssLocation
}

// File gets the file that section was parsed from.
//
// If no such file exists, for example because the CSS was loaded via
// gtk.CSSProvider.LoadFromData(), then NULL is returned.
//
// The function returns the following values:
//
//    - file: GFile from which the section was parsed.
//
func (section *CSSSection) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _file *gio.File // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// Parent gets the parent section for the given section.
//
// The parent section is the section that contains this section. A special case
// are sections of type GTK_CSS_SECTION_DOCUMENT. Their parent will either be
// NULL if they are the original CSS document that was loaded by
// gtk.CSSProvider.LoadFromFile() or a section of type GTK_CSS_SECTION_IMPORT if
// it was loaded with an import rule from a different file.
//
// The function returns the following values:
//
//    - cssSection (optional): parent section.
//
func (section *CSSSection) Parent() *CSSSection {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _cssSection *CSSSection // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_cssSection = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_css_section_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_cssSection)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _cssSection
}

// StartLocation returns the location in the CSS document where this section
// starts.
//
// The function returns the following values:
//
//    - cssLocation: start location of this section.
//
func (section *CSSSection) StartLocation() *CSSLocation {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _cssLocation
}

// String prints the section into a human-readable text form using
// gtk.CSSSection.Print().
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (section *CSSSection) String() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
