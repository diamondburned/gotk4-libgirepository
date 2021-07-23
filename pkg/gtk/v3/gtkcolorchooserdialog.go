// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialogger},
	})
}

// ColorChooserDialog widget is a dialog for choosing a color. It implements the
// ColorChooser interface.
type ColorChooserDialog struct {
	Dialog

	ColorChooser
	*externglib.Object
}

func wrapColorChooserDialog(obj *externglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
			},
		},
		ColorChooser: ColorChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalColorChooserDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorChooserDialog(obj), nil
}

// NewColorChooserDialog creates a new ColorChooserDialog.
func NewColorChooserDialog(title string, parent *Window) *ColorChooserDialog {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)

	var _colorChooserDialog *ColorChooserDialog // out

	_colorChooserDialog = wrapColorChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserDialog
}

func (*ColorChooserDialog) privateColorChooserDialog() {}
