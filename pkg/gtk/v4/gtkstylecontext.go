// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeStyleContextPrintFlags = coreglib.Type(girepository.MustFind("Gtk", "StyleContextPrintFlags").RegisteredGType())
	GTypeStyleContext           = coreglib.Type(girepository.MustFind("Gtk", "StyleContext").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleContextPrintFlags, F: marshalStyleContextPrintFlags},
		coreglib.TypeMarshaler{T: GTypeStyleContext, F: marshalStyleContext},
	})
}

// StyleContextPrintFlags flags that modify the behavior of
// gtk_style_context_to_string().
//
// New values may be added to this enumeration.
type StyleContextPrintFlags C.guint

const (
	// StyleContextPrintNone: default value.
	StyleContextPrintNone StyleContextPrintFlags = 0b0
	// StyleContextPrintRecurse: print the entire tree of CSS nodes starting at
	// the style context's node.
	StyleContextPrintRecurse StyleContextPrintFlags = 0b1
	// StyleContextPrintShowStyle: show the values of the CSS properties for
	// each node.
	StyleContextPrintShowStyle StyleContextPrintFlags = 0b10
	// StyleContextPrintShowChange: show information about what changes affect
	// the styles.
	StyleContextPrintShowChange StyleContextPrintFlags = 0b100
)

func marshalStyleContextPrintFlags(p uintptr) (interface{}, error) {
	return StyleContextPrintFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for StyleContextPrintFlags.
func (s StyleContextPrintFlags) String() string {
	if s == 0 {
		return "StyleContextPrintFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(101)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case StyleContextPrintNone:
			builder.WriteString("None|")
		case StyleContextPrintRecurse:
			builder.WriteString("Recurse|")
		case StyleContextPrintShowStyle:
			builder.WriteString("ShowStyle|")
		case StyleContextPrintShowChange:
			builder.WriteString("ShowChange|")
		default:
			builder.WriteString(fmt.Sprintf("StyleContextPrintFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s StyleContextPrintFlags) Has(other StyleContextPrintFlags) bool {
	return (s & other) == other
}

// StyleContextOverrides contains methods that are overridable.
type StyleContextOverrides struct {
}

func defaultStyleContextOverrides(v *StyleContext) StyleContextOverrides {
	return StyleContextOverrides{}
}

// StyleContext: GtkStyleContext stores styling information affecting a widget.
//
// In order to construct the final style information, GtkStyleContext queries
// information from all attached GtkStyleProviders. Style providers can be
// either attached explicitly to the context through
// gtk.StyleContext.AddProvider(), or to the display through
// gtk.StyleContext().AddProviderForDisplay. The resulting style is a
// combination of all providers’ information in priority order.
//
// For GTK widgets, any GtkStyleContext returned by gtk.Widget.GetStyleContext()
// will already have a GdkDisplay and RTL/LTR information set. The style context
// will also be updated automatically if any of these settings change on the
// widget.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
//
// Custom styling in UI libraries and applications
//
// If you are developing a library with custom widgets that render differently
// than standard components, you may need to add a GtkStyleProvider yourself
// with the GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority, either a
// GtkCssProvider or a custom object implementing the GtkStyleProvider
// interface. This way themes may still attempt to style your UI elements in a
// different way if needed so.
//
// If you are using custom styling on an applications, you probably want then to
// make your style information prevail to the theme’s, so you must use a
// GtkStyleProvider with the GTK_STYLE_PROVIDER_PRIORITY_APPLICATION priority,
// keep in mind that the user settings in XDG_CONFIG_HOME/gtk-4.0/gtk.css will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleContext)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StyleContext, *StyleContextClass, StyleContextOverrides](
		GTypeStyleContext,
		initStyleContextClass,
		wrapStyleContext,
		defaultStyleContextOverrides,
	)
}

func initStyleContextClass(gclass unsafe.Pointer, overrides StyleContextOverrides, classInitFunc func(*StyleContextClass)) {
	if classInitFunc != nil {
		class := (*StyleContextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyleContext(obj *coreglib.Object) *StyleContext {
	return &StyleContext{
		Object: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	return wrapStyleContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StyleContextClass: instance of this type is always passed by reference.
type StyleContextClass struct {
	*styleContextClass
}

// styleContextClass is the struct that's finalized.
type styleContextClass struct {
	native unsafe.Pointer
}

var GIRInfoStyleContextClass = girepository.MustFind("Gtk", "StyleContextClass")
