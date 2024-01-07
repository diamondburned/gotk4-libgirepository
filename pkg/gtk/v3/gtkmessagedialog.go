// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
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
// GtkWidget* _gotk4_gtk_message_dialog_new2(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons) {
// 	return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, NULL);
// }
import "C"

// GType values.
var (
	GTypeButtonsType   = coreglib.Type(girepository.MustFind("Gtk", "ButtonsType").RegisteredGType())
	GTypeMessageDialog = coreglib.Type(girepository.MustFind("Gtk", "MessageDialog").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeButtonsType, F: marshalButtonsType},
		coreglib.TypeMarshaler{T: GTypeMessageDialog, F: marshalMessageDialog},
	})
}

// ButtonsType: prebuilt sets of buttons for the dialog. If none of these
// choices are appropriate, simply use GTK_BUTTONS_NONE then call
// gtk_dialog_add_buttons().
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType C.gint

const (
	// ButtonsNone: no buttons at all.
	ButtonsNone ButtonsType = iota
	// ButtonsOK: OK button.
	ButtonsOK
	// ButtonsClose: close button.
	ButtonsClose
	// ButtonsCancel: cancel button.
	ButtonsCancel
	// ButtonsYesNo yes and No buttons.
	ButtonsYesNo
	// ButtonsOKCancel: OK and Cancel buttons.
	ButtonsOKCancel
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ButtonsType.
func (b ButtonsType) String() string {
	switch b {
	case ButtonsNone:
		return "None"
	case ButtonsOK:
		return "OK"
	case ButtonsClose:
		return "Close"
	case ButtonsCancel:
		return "Cancel"
	case ButtonsYesNo:
		return "YesNo"
	case ButtonsOKCancel:
		return "OKCancel"
	default:
		return fmt.Sprintf("ButtonsType(%d)", b)
	}
}

// MessageDialogOverrides contains methods that are overridable.
type MessageDialogOverrides struct {
}

func defaultMessageDialogOverrides(v *MessageDialog) MessageDialogOverrides {
	return MessageDialogOverrides{}
}

// MessageDialog presents a dialog with some message text. It’s simply a
// convenience widget; you could construct the equivalent of MessageDialog from
// Dialog without too much effort, but MessageDialog saves typing.
//
// One difference from Dialog is that MessageDialog sets the
// Window:skip-taskbar-hint property to TRUE, so that the dialog is hidden from
// the taskbar by default.
//
// The easiest way to do a modal message dialog is to use gtk_dialog_run(),
// though you can also pass in the GTK_DIALOG_MODAL flag, gtk_dialog_run()
// automatically makes the dialog modal and waits for the user to respond to it.
// gtk_dialog_run() returns when any dialog button is clicked.
//
// An example for using a modal dialog:
//
//     GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_message_dialog_new (parent_window,
//                                      flags,
//                                      GTK_MESSAGE_ERROR,
//                                      GTK_BUTTONS_CLOSE,
//                                      "Error reading “s”: s",
//                                      filename,
//                                      g_strerror (errno));
//
//     // Destroy the dialog when the user responds to it
//     // (e.g. clicks a button)
//
//     g_signal_connect_swapped (dialog, "response",
//                               G_CALLBACK (gtk_widget_destroy),
//                               dialog);
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog struct {
	_ [0]func() // equal guard
	Dialog
}

var (
	_ Binner = (*MessageDialog)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MessageDialog, *MessageDialogClass, MessageDialogOverrides](
		GTypeMessageDialog,
		initMessageDialogClass,
		wrapMessageDialog,
		defaultMessageDialogOverrides,
	)
}

func initMessageDialogClass(gclass unsafe.Pointer, overrides MessageDialogOverrides, classInitFunc func(*MessageDialogClass)) {
	if classInitFunc != nil {
		class := (*MessageDialogClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMessageDialog(obj *coreglib.Object) *MessageDialog {
	return &MessageDialog{
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
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	return wrapMessageDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MessageDialogClass: instance of this type is always passed by reference.
type MessageDialogClass struct {
	*messageDialogClass
}

// messageDialogClass is the struct that's finalized.
type messageDialogClass struct {
	native unsafe.Pointer
}

var GIRInfoMessageDialogClass = girepository.MustFind("Gtk", "MessageDialogClass")

// NewMessageDialog creates a new message dialog. This is a simple
// dialog with some text taht the user may want to see. When the user
// clicks a button, a "response" signal is emitted with response IDs
// from ResponseType.
func NewMessageDialog(parent *Window, flags DialogFlags, typ MessageType, buttons ButtonsType) *MessageDialog {
	w := C._gotk4_gtk_message_dialog_new2(
		(*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native())),
		(C.GtkDialogFlags)(flags),
		(C.GtkMessageType)(typ),
		(C.GtkButtonsType)(buttons),
	)
	runtime.KeepAlive(parent)

	return wrapMessageDialog(coreglib.Take(unsafe.Pointer(w)))
}
