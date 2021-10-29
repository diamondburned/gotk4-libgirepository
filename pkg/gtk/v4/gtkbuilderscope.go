// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_builder_closure_flags_get_type()), F: marshalBuilderClosureFlags},
		{T: externglib.Type(C.gtk_builder_scope_get_type()), F: marshalBuilderScoper},
		{T: externglib.Type(C.gtk_builder_cscope_get_type()), F: marshalBuilderCScoper},
	})
}

// BuilderClosureFlags: list of flags that can be passed to
// gtk_builder_create_closure().
//
// New values may be added in the future for new features, so external
// implementations of gtk.BuilderScope should test the flags for unknown values
// and raise a GTK_BUILDER_ERROR_INVALID_ATTRIBUTE error when they encounter
// one.
type BuilderClosureFlags C.guint

const (
	// BuilderClosureSwapped: closure should be created swapped. See
	// g_cclosure_new_swap() for details.
	BuilderClosureSwapped BuilderClosureFlags = 0b1
)

func marshalBuilderClosureFlags(p uintptr) (interface{}, error) {
	return BuilderClosureFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for BuilderClosureFlags.
func (b BuilderClosureFlags) String() string {
	if b == 0 {
		return "BuilderClosureFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(21)

	for b != 0 {
		next := b & (b - 1)
		bit := b - next

		switch bit {
		case BuilderClosureSwapped:
			builder.WriteString("Swapped|")
		default:
			builder.WriteString(fmt.Sprintf("BuilderClosureFlags(0b%b)|", bit))
		}

		b = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if b contains other.
func (b BuilderClosureFlags) Has(other BuilderClosureFlags) bool {
	return (b & other) == other
}

// BuilderScopeOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BuilderScopeOverrider interface {
	TypeFromFunction(builder *Builder, functionName string) externglib.Type
	TypeFromName(builder *Builder, typeName string) externglib.Type
}

// BuilderScope: GtkBuilderScope is an interface to provide language binding
// support to GtkBuilder.
//
// The goal of GtkBuilderScope is to look up programming-language-specific
// values for strings that are given in a GtkBuilder UI file.
//
// The primary intended audience is bindings that want to provide deeper
// integration of GtkBuilder into the language.
//
// A GtkBuilderScope instance may be used with multiple GtkBuilder objects, even
// at once.
//
// By default, GTK will use its own implementation of GtkBuilderScope for the C
// language which can be created via gtk.BuilderCScope.New.
type BuilderScope struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*BuilderScope)(nil)
)

// BuilderScoper describes BuilderScope's interface methods.
type BuilderScoper interface {
	externglib.Objector

	baseBuilderScope() *BuilderScope
}

var _ BuilderScoper = (*BuilderScope)(nil)

func wrapBuilderScope(obj *externglib.Object) *BuilderScope {
	return &BuilderScope{
		Object: obj,
	}
}

func marshalBuilderScoper(p uintptr) (interface{}, error) {
	return wrapBuilderScope(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *BuilderScope) baseBuilderScope() *BuilderScope {
	return self
}

// BaseBuilderScope returns the underlying base object.
func BaseBuilderScope(obj BuilderScoper) *BuilderScope {
	return obj.baseBuilderScope()
}

// BuilderCScope: GtkBuilderScope implementation for the C language.
//
// GtkBuilderCScope instances use symbols explicitly added to builder with prior
// calls to gtk.BuilderCScope.AddCallbackSymbol(). If developers want to do
// that, they are encouraged to create their own scopes for that purpose.
//
// In the case that symbols are not explicitly added; GTK will uses GModule’s
// introspective features (by opening the module NULL) to look at the
// application’s symbol table. From here it tries to match the signal function
// names given in the interface description with symbols in the application.
//
// Note that unless gtk.BuilderCScope.AddCallbackSymbol() is called for all
// signal callbacks which are referenced by the loaded XML, this functionality
// will require that GModule be supported on the platform.
type BuilderCScope struct {
	*externglib.Object

	BuilderScope
}

var (
	_ externglib.Objector = (*BuilderCScope)(nil)
)

func wrapBuilderCScope(obj *externglib.Object) *BuilderCScope {
	return &BuilderCScope{
		Object: obj,
		BuilderScope: BuilderScope{
			Object: obj,
		},
	}
}

func marshalBuilderCScoper(p uintptr) (interface{}, error) {
	return wrapBuilderCScope(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBuilderCScope creates a new GtkBuilderCScope object to use with future
// GtkBuilder instances.
//
// Calling this function is only necessary if you want to add custom callbacks
// via gtk.BuilderCScope.AddCallbackSymbol().
func NewBuilderCScope() *BuilderCScope {
	var _cret *C.GtkBuilderScope // in

	_cret = C.gtk_builder_cscope_new()

	var _builderCScope *BuilderCScope // out

	_builderCScope = wrapBuilderCScope(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderCScope
}
