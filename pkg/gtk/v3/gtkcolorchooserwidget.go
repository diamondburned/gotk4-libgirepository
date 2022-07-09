// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeColorChooserWidget returns the GType for the type ColorChooserWidget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorChooserWidget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ColorChooserWidget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorChooserWidget)
	return gtype
}

// ColorChooserWidgetOverrider contains methods that are overridable.
type ColorChooserWidgetOverrider interface {
}

// ColorChooserWidget widget lets the user select a color. By default, the
// chooser presents a predefined palette of colors, plus a small number of
// settable custom colors. It is also possible to select a different color with
// the single-color editor. To enter the single-color editing mode, use the
// context menu of any color of the palette, or use the '+' button to add a new
// custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To change the initially selected color, use gtk_color_chooser_set_rgba(). To
// get the selected color use gtk_color_chooser_get_rgba().
//
// The ColorChooserWidget is used in the ColorChooserDialog to provide a dialog
// for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserWidget)(nil)
	_ Containerer       = (*ColorChooserWidget)(nil)
)

func classInitColorChooserWidgetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapColorChooserWidget(obj *coreglib.Object) *ColorChooserWidget {
	return &ColorChooserWidget{
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
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserWidget(p uintptr) (interface{}, error) {
	return wrapColorChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserWidget creates a new ColorChooserWidget.
//
// The function returns the following values:
//
//    - colorChooserWidget: new ColorChooserWidget.
//
func NewColorChooserWidget() *ColorChooserWidget {
	_gret := girepository.MustFind("Gtk", "ColorChooserWidget").InvokeMethod("new_ColorChooserWidget", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _colorChooserWidget *ColorChooserWidget // out

	_colorChooserWidget = wrapColorChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserWidget
}
