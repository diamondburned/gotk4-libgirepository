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
	GTypeFontChooserWidget = coreglib.Type(girepository.MustFind("Gtk", "FontChooserWidget").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFontChooserWidget, F: marshalFontChooserWidget},
	})
}

// FontChooserWidgetOverrides contains methods that are overridable.
type FontChooserWidgetOverrides struct {
}

func defaultFontChooserWidgetOverrides(v *FontChooserWidget) FontChooserWidgetOverrides {
	return FontChooserWidgetOverrides{}
}

// FontChooserWidget widget lists the available fonts, styles and sizes,
// allowing the user to select a font. It is used in the FontChooserDialog
// widget to provide a dialog box for selecting fonts.
//
// To set the font which is initially selected, use gtk_font_chooser_set_font()
// or gtk_font_chooser_set_font_desc().
//
// To get the selected font use gtk_font_chooser_get_font() or
// gtk_font_chooser_get_font_desc().
//
// To change the text which is shown in the preview area, use
// gtk_font_chooser_set_preview_text().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*coreglib.Object
	FontChooser
}

var (
	_ coreglib.Objector = (*FontChooserWidget)(nil)
	_ Containerer       = (*FontChooserWidget)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FontChooserWidget, *FontChooserWidgetClass, FontChooserWidgetOverrides](
		GTypeFontChooserWidget,
		initFontChooserWidgetClass,
		wrapFontChooserWidget,
		defaultFontChooserWidgetOverrides,
	)
}

func initFontChooserWidgetClass(gclass unsafe.Pointer, overrides FontChooserWidgetOverrides, classInitFunc func(*FontChooserWidgetClass)) {
	if classInitFunc != nil {
		class := (*FontChooserWidgetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFontChooserWidget(obj *coreglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidget(p uintptr) (interface{}, error) {
	return wrapFontChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
