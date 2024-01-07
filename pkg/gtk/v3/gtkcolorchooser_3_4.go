// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ColorChooser_ConnectColorActivated(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeColorChooser = coreglib.Type(girepository.MustFind("Gtk", "ColorChooser").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorChooser, F: marshalColorChooser},
	})
}

// ColorChooserOverrider contains methods that are overridable.
type ColorChooserOverrider interface {
}

// ColorChooser is an interface that is implemented by widgets for choosing
// colors. Depending on the situation, colors may be allowed to have alpha
// (translucency).
//
// In GTK+, the main widgets that implement this interface are
// ColorChooserWidget, ColorChooserDialog and ColorButton.
//
// ColorChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ColorChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ColorChooser)(nil)
)

// ColorChooserer describes ColorChooser's interface methods.
type ColorChooserer interface {
	coreglib.Objector

	baseColorChooser() *ColorChooser
}

var _ ColorChooserer = (*ColorChooser)(nil)

func ifaceInitColorChooserer(gifacePtr, data C.gpointer) {
}

func wrapColorChooser(obj *coreglib.Object) *ColorChooser {
	return &ColorChooser{
		Object: obj,
	}
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	return wrapColorChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ColorChooser) baseColorChooser() *ColorChooser {
	return v
}

// BaseColorChooser returns the underlying base object.
func BaseColorChooser(obj ColorChooserer) *ColorChooser {
	return obj.baseColorChooser()
}

// ConnectColorActivated is emitted when a color is activated from the color
// chooser. This usually happens when the user clicks a color swatch, or a color
// is selected and the user presses one of the keys Space, Shift+Space, Return
// or Enter.
func (v *ColorChooser) ConnectColorActivated(f func(color *gdk.RGBA)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "color-activated", false, unsafe.Pointer(C._gotk4_gtk3_ColorChooser_ConnectColorActivated), f)
}
