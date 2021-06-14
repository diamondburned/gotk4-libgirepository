// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_role_get_type()), F: marshalButtonRole},
		{T: externglib.Type(C.gtk_model_button_get_type()), F: marshalModelButton},
	})
}

// ButtonRole: the role specifies the desired appearance of a ModelButton.
type ButtonRole int

const (
	// ButtonRoleNormal: a plain button
	ButtonRoleNormal ButtonRole = 0
	// ButtonRoleCheck: a check button
	ButtonRoleCheck ButtonRole = 1
	// ButtonRoleRadio: a radio button
	ButtonRoleRadio ButtonRole = 2
)

func marshalButtonRole(p uintptr) (interface{}, error) {
	return ButtonRole(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ModelButton: gtkModelButton is a button class that can use a #GAction as its
// model. In contrast to ToggleButton or RadioButton, which can also be backed
// by a #GAction via the Actionable:action-name property, GtkModelButton will
// adapt its appearance according to the kind of action it is backed by, and
// appear either as a plain, check or radio button.
//
// Model buttons are used when popovers from a menu model with
// gtk_popover_new_from_model(); they can also be used manually in a
// PopoverMenu.
//
// When the action is specified via the Actionable:action-name and
// Actionable:action-target properties, the role of the button (i.e. whether it
// is a plain, check or radio button) is determined by the type of the action
// and doesn't have to be explicitly specified with the ModelButton:role
// property.
//
// The content of the button is specified by the ModelButton:text and
// ModelButton:icon properties.
//
// The appearance of model buttons can be influenced with the
// ModelButton:centered and ModelButton:iconic properties.
//
// Model buttons have built-in support for submenus in PopoverMenu. To make a
// GtkModelButton that opens a submenu when activated, set the
// ModelButton:menu-name property. To make a button that goes back to the parent
// menu, you should set the ModelButton:inverted property to place the submenu
// indicator at the opposite side.
//
// Example
//
//    <object class="GtkPopoverMenu">
//      <child>
//        <object class="GtkBox">
//          <property name="visible">True</property>
//          <property name="margin">10</property>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.cut</property>
//              <property name="text" translatable="yes">Cut</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.copy</property>
//              <property name="text" translatable="yes">Copy</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.paste</property>
//              <property name="text" translatable="yes">Paste</property>
//            </object>
//          </child>
//        </object>
//      </child>
//    </object>
//
// CSS nodes
//
//    button.model
//    ├── <child>
//    ╰── check
//
// Iconic model buttons (see ModelButton:iconic) change the name of their main
// node to button and add a .model style class to it. The indicator subnode is
// invisible in this case.
type ModelButton interface {
	Button
	Actionable
	Activatable
	Buildable
}

// modelButton implements the ModelButton class.
type modelButton struct {
	Button
	Actionable
	Activatable
	Buildable
}

var _ ModelButton = (*modelButton)(nil)

// WrapModelButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapModelButton(obj *externglib.Object) ModelButton {
	return modelButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalModelButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapModelButton(obj), nil
}
