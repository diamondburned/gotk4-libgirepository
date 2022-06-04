// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkfontchooserdialog.go.
var GTypeFontChooserDialog = coreglib.Type(C.gtk_font_chooser_dialog_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFontChooserDialog, F: marshalFontChooserDialog},
	})
}

// FontChooserDialogOverrider contains methods that are overridable.
type FontChooserDialogOverrider interface {
}

// FontChooserDialog widget is a dialog for selecting a font. It implements the
// FontChooser interface.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the Buildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	FontChooser
}

var (
	_ coreglib.Objector = (*FontChooserDialog)(nil)
	_ Binner            = (*FontChooserDialog)(nil)
)

func classInitFontChooserDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFontChooserDialog(obj *coreglib.Object) *FontChooserDialog {
	return &FontChooserDialog{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserDialog(p uintptr) (interface{}, error) {
	return wrapFontChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserDialog creates a new FontChooserDialog.
//
// The function takes the following parameters:
//
//    - title (optional): title of the dialog, or NULL.
//    - parent (optional): transient parent of the dialog, or NULL.
//
// The function returns the following values:
//
//    - fontChooserDialog: new FontChooserDialog.
//
func NewFontChooserDialog(title string, parent *Window) *FontChooserDialog {
	var _args [2]girepository.Argument

	if title != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_args[0]))
	}
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_gret := girepository.MustFind("Gtk", "FontChooserDialog").InvokeMethod("new_FontChooserDialog", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _fontChooserDialog *FontChooserDialog // out

	_fontChooserDialog = wrapFontChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserDialog
}
