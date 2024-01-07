// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ColorSelection_ConnectColorChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeColorSelection = coreglib.Type(girepository.MustFind("Gtk", "ColorSelection").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorSelection, F: marshalColorSelection},
	})
}

// ColorSelectionOverrides contains methods that are overridable.
type ColorSelectionOverrides struct {
}

func defaultColorSelectionOverrides(v *ColorSelection) ColorSelectionOverrides {
	return ColorSelectionOverrides{}
}

type ColorSelection struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ColorSelection)(nil)
	_ coreglib.Objector = (*ColorSelection)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ColorSelection, *ColorSelectionClass, ColorSelectionOverrides](
		GTypeColorSelection,
		initColorSelectionClass,
		wrapColorSelection,
		defaultColorSelectionOverrides,
	)
}

func initColorSelectionClass(gclass unsafe.Pointer, overrides ColorSelectionOverrides, classInitFunc func(*ColorSelectionClass)) {
	if classInitFunc != nil {
		class := (*ColorSelectionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapColorSelection(obj *coreglib.Object) *ColorSelection {
	return &ColorSelection{
		Box: Box{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalColorSelection(p uintptr) (interface{}, error) {
	return wrapColorSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectColorChanged: this signal is emitted when the color changes in the
// ColorSelection according to its update policy.
func (v *ColorSelection) ConnectColorChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "color-changed", false, unsafe.Pointer(C._gotk4_gtk3_ColorSelection_ConnectColorChanged), f)
}

// ColorSelectionClass: instance of this type is always passed by reference.
type ColorSelectionClass struct {
	*colorSelectionClass
}

// colorSelectionClass is the struct that's finalized.
type colorSelectionClass struct {
	native unsafe.Pointer
}

var GIRInfoColorSelectionClass = girepository.MustFind("Gtk", "ColorSelectionClass")
