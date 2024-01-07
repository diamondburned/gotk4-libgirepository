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
// extern void _gotk4_gtk3_RadioButton_ConnectGroupChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeRadioButton = coreglib.Type(girepository.MustFind("Gtk", "RadioButton").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRadioButton, F: marshalRadioButton},
	})
}

// RadioButtonOverrides contains methods that are overridable.
type RadioButtonOverrides struct {
}

func defaultRadioButtonOverrides(v *RadioButton) RadioButtonOverrides {
	return RadioButtonOverrides{}
}

// RadioButton: single radio button performs the same basic function as a
// CheckButton, as its position in the object hierarchy reflects. It is only
// when multiple radio buttons are grouped together that they become a different
// user interface component in their own right.
//
// Every radio button is a member of some group of radio buttons. When one is
// selected, all other radio buttons in the same group are deselected. A
// RadioButton is one way of giving the user a choice from many options.
//
// Radio button widgets are created with gtk_radio_button_new(), passing NULL as
// the argument if this is the first radio button in a group. In subsequent
// calls, the group you wish to add this button to should be passed as an
// argument. Optionally, gtk_radio_button_new_with_label() can be used if you
// want a text label on the radio button.
//
// Alternatively, when adding widgets to an existing group of radio buttons, use
// gtk_radio_button_new_from_widget() with a RadioButton that already has a
// group assigned to it. The convenience function
// gtk_radio_button_new_with_label_from_widget() is also provided.
//
// To retrieve the group a RadioButton is assigned to, use
// gtk_radio_button_get_group().
//
// To remove a RadioButton from one group and make it part of a new one, use
// gtk_radio_button_set_group().
//
// The group list does not need to be freed, as each RadioButton will remove
// itself and its list item when it is destroyed.
//
// CSS nodes
//
//    void create_radio_buttons (void) {
//
//       GtkWidget *window, *radio1, *radio2, *box, *entry;
//       window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//       box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 2);
//       gtk_box_set_homogeneous (GTK_BOX (box), TRUE);
//
//       // Create a radio button with a GtkEntry widget
//       radio1 = gtk_radio_button_new (NULL);
//       entry = gtk_entry_new ();
//       gtk_container_add (GTK_CONTAINER (radio1), entry);
//
//
//       // Create a radio button with a label
//       radio2 = gtk_radio_button_new_with_label_from_widget (GTK_RADIO_BUTTON (radio1),
//                                                             "I’m the second radio button.");
//
//       // Pack them into a box, then show all the widgets
//       gtk_box_pack_start (GTK_BOX (box), radio1);
//       gtk_box_pack_start (GTK_BOX (box), radio2);
//       gtk_container_add (GTK_CONTAINER (window), box);
//       gtk_widget_show_all (window);
//       return;
//    }
//
// When an unselected button in the group is clicked the clicked button receives
// the ToggleButton::toggled signal, as does the previously selected button.
// Inside the ToggleButton::toggled handler, gtk_toggle_button_get_active() can
// be used to determine if the button has been selected or deselected.
type RadioButton struct {
	_ [0]func() // equal guard
	CheckButton
}

var (
	_ Binner            = (*RadioButton)(nil)
	_ coreglib.Objector = (*RadioButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RadioButton, *RadioButtonClass, RadioButtonOverrides](
		GTypeRadioButton,
		initRadioButtonClass,
		wrapRadioButton,
		defaultRadioButtonOverrides,
	)
}

func initRadioButtonClass(gclass unsafe.Pointer, overrides RadioButtonOverrides, classInitFunc func(*RadioButtonClass)) {
	if classInitFunc != nil {
		class := (*RadioButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRadioButton(obj *coreglib.Object) *RadioButton {
	return &RadioButton{
		CheckButton: CheckButton{
			ToggleButton: ToggleButton{
				Button: Button{
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
					Actionable: Actionable{
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
					Activatable: Activatable{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalRadioButton(p uintptr) (interface{}, error) {
	return wrapRadioButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectGroupChanged is emitted when the group of radio buttons that a radio
// button belongs to changes. This is emitted when a radio button switches from
// being alone to being part of a group of 2 or more buttons, or vice-versa, and
// when a button is moved from one group of 2 or more buttons to a different
// one, but not when the composition of the group that a button belongs to
// changes.
func (v *RadioButton) ConnectGroupChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "group-changed", false, unsafe.Pointer(C._gotk4_gtk3_RadioButton_ConnectGroupChanged), f)
}

// RadioButtonClass: instance of this type is always passed by reference.
type RadioButtonClass struct {
	*radioButtonClass
}

// radioButtonClass is the struct that's finalized.
type radioButtonClass struct {
	native unsafe.Pointer
}

var GIRInfoRadioButtonClass = girepository.MustFind("Gtk", "RadioButtonClass")
