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
import "C"

// GType values.
var (
	GTypeColorChooserDialog = coreglib.Type(girepository.MustFind("Gtk", "ColorChooserDialog").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorChooserDialog, F: marshalColorChooserDialog},
	})
}

// ColorChooserDialogOverrides contains methods that are overridable.
type ColorChooserDialogOverrides struct {
}

func defaultColorChooserDialogOverrides(v *ColorChooserDialog) ColorChooserDialogOverrides {
	return ColorChooserDialogOverrides{}
}

// ColorChooserDialog widget is a dialog for choosing a color. It implements the
// ColorChooser interface.
type ColorChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserDialog)(nil)
	_ Binner            = (*ColorChooserDialog)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ColorChooserDialog, *ColorChooserDialogClass, ColorChooserDialogOverrides](
		GTypeColorChooserDialog,
		initColorChooserDialogClass,
		wrapColorChooserDialog,
		defaultColorChooserDialogOverrides,
	)
}

func initColorChooserDialogClass(gclass unsafe.Pointer, overrides ColorChooserDialogOverrides, classInitFunc func(*ColorChooserDialogClass)) {
	if classInitFunc != nil {
		class := (*ColorChooserDialogClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapColorChooserDialog(obj *coreglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
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
				},
			},
		},
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	return wrapColorChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}