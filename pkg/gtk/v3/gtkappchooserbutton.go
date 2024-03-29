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
// extern void _gotk4_gtk3_AppChooserButton_ConnectCustomItemActivated(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeAppChooserButton = coreglib.Type(girepository.MustFind("Gtk", "AppChooserButton").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAppChooserButton, F: marshalAppChooserButton},
	})
}

// AppChooserButtonOverrides contains methods that are overridable.
type AppChooserButtonOverrides struct {
}

func defaultAppChooserButtonOverrides(v *AppChooserButton) AppChooserButtonOverrides {
	return AppChooserButtonOverrides{}
}

// AppChooserButton is a widget that lets the user select an application. It
// implements the AppChooser interface.
//
// Initially, a AppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// AppChooserButton:show-default-item is TRUE, the default application.
//
// The list of applications shown in a AppChooserButton includes the recommended
// applications for the given content type. When
// AppChooserButton:show-default-item is set, the default application is also
// included. To let the user chooser other applications, you can set the
// AppChooserButton:show-dialog-item property, which allows to open a full
// AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk_app_chooser_button_append_custom_item(). These items cause the
// AppChooserButton::custom-item-activated signal to be emitted when they are
// selected.
//
// To track changes in the selected application, use the ComboBox::changed
// signal.
type AppChooserButton struct {
	_ [0]func() // equal guard
	ComboBox

	*coreglib.Object
	AppChooser
}

var (
	_ coreglib.Objector = (*AppChooserButton)(nil)
	_ Binner            = (*AppChooserButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AppChooserButton, *AppChooserButtonClass, AppChooserButtonOverrides](
		GTypeAppChooserButton,
		initAppChooserButtonClass,
		wrapAppChooserButton,
		defaultAppChooserButtonOverrides,
	)
}

func initAppChooserButtonClass(gclass unsafe.Pointer, overrides AppChooserButtonOverrides, classInitFunc func(*AppChooserButtonClass)) {
	if classInitFunc != nil {
		class := (*AppChooserButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAppChooserButton(obj *coreglib.Object) *AppChooserButton {
	return &AppChooserButton{
		ComboBox: ComboBox{
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
			Object: obj,
			CellEditable: CellEditable{
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
			CellLayout: CellLayout{
				Object: obj,
			},
		},
		Object: obj,
		AppChooser: AppChooser{
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
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	return wrapAppChooserButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCustomItemActivated is emitted when a custom item, previously added
// with gtk_app_chooser_button_append_custom_item(), is activated from the
// dropdown menu.
func (v *AppChooserButton) ConnectCustomItemActivated(f func(itemName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "custom-item-activated", false, unsafe.Pointer(C._gotk4_gtk3_AppChooserButton_ConnectCustomItemActivated), f)
}

// AppChooserButtonClass: instance of this type is always passed by reference.
type AppChooserButtonClass struct {
	*appChooserButtonClass
}

// appChooserButtonClass is the struct that's finalized.
type appChooserButtonClass struct {
	native unsafe.Pointer
}

var GIRInfoAppChooserButtonClass = girepository.MustFind("Gtk", "AppChooserButtonClass")
