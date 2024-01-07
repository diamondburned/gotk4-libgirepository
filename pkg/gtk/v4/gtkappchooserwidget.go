// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected(gpointer, void*, guintptr);
// extern void _gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeAppChooserWidget = coreglib.Type(girepository.MustFind("Gtk", "AppChooserWidget").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAppChooserWidget, F: marshalAppChooserWidget},
	})
}

// AppChooserWidget: GtkAppChooserWidget is a widget for selecting applications.
//
// It is the main building block for gtk.AppChooserDialog. Most applications
// only need to use the latter; but you can use this widget as part of a larger
// widget if you have special needs.
//
// GtkAppChooserWidget offers detailed control over what applications are shown,
// using the gtk.AppChooserWidget:show-default,
// gtk.AppChooserWidget:show-recommended, gtk.AppChooserWidget:show-fallback,
// gtk.AppChooserWidget:show-other and gtk.AppChooserWidget:show-all properties.
// See the gtk.AppChooser documentation for more information about these groups
// of applications.
//
// To keep track of the selected application, use the
// gtk.AppChooserWidget::application-selected and
// gtk.AppChooserWidget::application-activated signals.
//
//
// CSS nodes
//
// GtkAppChooserWidget has a single CSS node with name appchooser.
type AppChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	AppChooser
}

var (
	_ Widgetter         = (*AppChooserWidget)(nil)
	_ coreglib.Objector = (*AppChooserWidget)(nil)
)

func wrapAppChooserWidget(obj *coreglib.Object) *AppChooserWidget {
	return &AppChooserWidget{
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

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	return wrapAppChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectApplicationActivated is emitted when an application item is activated
// from the widget's list.
//
// This usually happens when the user double clicks an item, or an item is
// selected and the user presses one of the keys Space, Shift+Space, Return or
// Enter.
func (v *AppChooserWidget) ConnectApplicationActivated(f func(application gio.AppInfor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "application-activated", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated), f)
}

// ConnectApplicationSelected is emitted when an application item is selected
// from the widget's list.
func (v *AppChooserWidget) ConnectApplicationSelected(f func(application gio.AppInfor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "application-selected", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected), f)
}
