// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_AppChooserButton_ConnectChanged(gpointer, guintptr);
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

// AppChooserButton: GtkAppChooserButton lets the user select an application.
//
// !An example GtkAppChooserButton (appchooserbutton.png)
//
// Initially, a GtkAppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// gtk.AppChooserButton:show-default-item is TRUE, the default application.
//
// The list of applications shown in a GtkAppChooserButton includes the
// recommended applications for the given content type. When
// gtk.AppChooserButton:show-default-item is set, the default application is
// also included. To let the user chooser other applications, you can set the
// gtk.AppChooserButton:show-dialog-item property, which allows to open a full
// gtk.AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk.AppChooserButton.AppendCustomItem(). These items cause the
// gtk.AppChooserButton::custom-item-activated signal to be emitted when they
// are selected.
//
// To track changes in the selected application, use the
// gtk.AppChooserButton::changed signal.
//
//
// CSS nodes
//
// GtkAppChooserButton has a single CSS node with the name “appchooserbutton”.
type AppChooserButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	AppChooser
}

var (
	_ Widgetter         = (*AppChooserButton)(nil)
	_ coreglib.Objector = (*AppChooserButton)(nil)
)

func wrapAppChooserButton(obj *coreglib.Object) *AppChooserButton {
	return &AppChooserButton{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
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
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	return wrapAppChooserButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when the active application changes.
func (v *AppChooserButton) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserButton_ConnectChanged), f)
}

// ConnectCustomItemActivated is emitted when a custom item is activated.
//
// Use gtk.AppChooserButton.AppendCustomItem(), to add custom items.
func (v *AppChooserButton) ConnectCustomItemActivated(f func(itemName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "custom-item-activated", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated), f)
}
