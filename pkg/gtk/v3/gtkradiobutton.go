// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_button_get_type()), F: marshalRadioButton},
	})
}

// RadioButton: a single radio button performs the same basic function as a
// CheckButton, as its position in the object hierarchy reflects. It is only
// when multiple radio buttons are grouped together that they become a different
// user interface component in their own right.
//
// Every radio button is a member of some group of radio buttons. When one is
// selected, all other radio buttons in the same group are deselected. A
// RadioButton is one way of giving the user a choice from many options.
//
// Radio button widgets are created with gtk_radio_button_new(), passing nil as
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
type RadioButton interface {
	CheckButton
	Actionable
	Activatable
	Buildable

	// Group retrieves the group assigned to a radio button.
	Group() *glib.SList
	// JoinGroup joins a RadioButton object to the group of another RadioButton
	// object
	//
	// Use this in language bindings instead of the gtk_radio_button_get_group()
	// and gtk_radio_button_set_group() methods
	//
	// A common way to set up a group of radio buttons is the following:
	//
	//      GtkRadioButton *radio_button;
	//      GtkRadioButton *last_button;
	//
	//      while (some_condition)
	//        {
	//           radio_button = gtk_radio_button_new (NULL);
	//
	//           gtk_radio_button_join_group (radio_button, last_button);
	//           last_button = radio_button;
	//        }
	JoinGroup(groupSource RadioButton)
	// SetGroup sets a RadioButton’s group. It should be noted that this does
	// not change the layout of your interface in any way, so if you are
	// changing the group, it is likely you will need to re-arrange the user
	// interface to reflect these changes.
	SetGroup(group *glib.SList)
}

// radioButton implements the RadioButton interface.
type radioButton struct {
	CheckButton
	Actionable
	Activatable
	Buildable
}

var _ RadioButton = (*radioButton)(nil)

// WrapRadioButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioButton(obj *externglib.Object) RadioButton {
	return RadioButton{
		CheckButton: WrapCheckButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalRadioButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioButton(obj), nil
}

// NewRadioButton constructs a class RadioButton.
func NewRadioButton(group *glib.SList) RadioButton {
	var arg1 *C.GSList

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// NewRadioButtonFromWidget constructs a class RadioButton.
func NewRadioButtonFromWidget(radioGroupMember RadioButton) RadioButton {
	var arg1 *C.GtkRadioButton

	arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new_from_widget(arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// NewRadioButtonWithLabel constructs a class RadioButton.
func NewRadioButtonWithLabel(group *glib.SList, label string) RadioButton {
	var arg1 *C.GSList
	var arg2 *C.gchar

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new_with_label(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// NewRadioButtonWithLabelFromWidget constructs a class RadioButton.
func NewRadioButtonWithLabelFromWidget(radioGroupMember RadioButton, label string) RadioButton {
	var arg1 *C.GtkRadioButton
	var arg2 *C.gchar

	arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new_with_label_from_widget(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// NewRadioButtonWithMnemonic constructs a class RadioButton.
func NewRadioButtonWithMnemonic(group *glib.SList, label string) RadioButton {
	var arg1 *C.GSList
	var arg2 *C.gchar

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new_with_mnemonic(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// NewRadioButtonWithMnemonicFromWidget constructs a class RadioButton.
func NewRadioButtonWithMnemonicFromWidget(radioGroupMember RadioButton, label string) RadioButton {
	var arg1 *C.GtkRadioButton
	var arg2 *C.gchar

	arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioButton
	var goret RadioButton

	cret = C.gtk_radio_button_new_with_mnemonic_from_widget(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioButton)

	return goret
}

// Group retrieves the group assigned to a radio button.
func (r radioButton) Group() *glib.SList {
	var arg0 *C.GtkRadioButton

	arg0 = (*C.GtkRadioButton)(unsafe.Pointer(r.Native()))

	var cret *C.GSList
	var goret *glib.SList

	cret = C.gtk_radio_button_get_group(arg0)

	goret = glib.WrapSList(unsafe.Pointer(cret))

	return goret
}

// JoinGroup joins a RadioButton object to the group of another RadioButton
// object
//
// Use this in language bindings instead of the gtk_radio_button_get_group()
// and gtk_radio_button_set_group() methods
//
// A common way to set up a group of radio buttons is the following:
//
//      GtkRadioButton *radio_button;
//      GtkRadioButton *last_button;
//
//      while (some_condition)
//        {
//           radio_button = gtk_radio_button_new (NULL);
//
//           gtk_radio_button_join_group (radio_button, last_button);
//           last_button = radio_button;
//        }
func (r radioButton) JoinGroup(groupSource RadioButton) {
	var arg0 *C.GtkRadioButton
	var arg1 *C.GtkRadioButton

	arg0 = (*C.GtkRadioButton)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GtkRadioButton)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_button_join_group(arg0, arg1)
}

// SetGroup sets a RadioButton’s group. It should be noted that this does
// not change the layout of your interface in any way, so if you are
// changing the group, it is likely you will need to re-arrange the user
// interface to reflect these changes.
func (r radioButton) SetGroup(group *glib.SList) {
	var arg0 *C.GtkRadioButton
	var arg1 *C.GSList

	arg0 = (*C.GtkRadioButton)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	C.gtk_radio_button_set_group(arg0, arg1)
}
