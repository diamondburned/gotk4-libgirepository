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
		{T: externglib.Type(C.gtk_menu_tool_button_get_type()), F: marshalMenuToolButton},
	})
}

// MenuToolButton: a MenuToolButton is a ToolItem that contains a button and a
// small additional button with an arrow. When clicked, the arrow button pops up
// a dropdown menu.
//
// Use gtk_menu_tool_button_new() to create a new MenuToolButton.
//
//
// GtkMenuToolButton as GtkBuildable
//
// The GtkMenuToolButton implementation of the GtkBuildable interface supports
// adding a menu by specifying “menu” as the “type” attribute of a <child>
// element.
//
// An example for a UI definition fragment with menus:
//
//    <object class="GtkMenuToolButton">
//      <child type="menu">
//        <object class="GtkMenu"/>
//      </child>
//    </object>
type MenuToolButton interface {
	ToolButton
	Actionable
	Activatable
	Buildable

	// SetArrowTooltipMarkup sets the tooltip markup text to be used as tooltip
	// for the arrow button which pops up the menu. See
	// gtk_tool_item_set_tooltip_text() for setting a tooltip on the whole
	// MenuToolButton.
	SetArrowTooltipMarkup(markup string)
	// SetArrowTooltipText sets the tooltip text to be used as tooltip for the
	// arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text()
	// for setting a tooltip on the whole MenuToolButton.
	SetArrowTooltipText(text string)
	// SetMenu sets the Menu that is popped up when the user clicks on the
	// arrow. If @menu is NULL, the arrow button becomes insensitive.
	SetMenu(menu Widget)
}

// menuToolButton implements the MenuToolButton class.
type menuToolButton struct {
	ToolButton
	Actionable
	Activatable
	Buildable
}

var _ MenuToolButton = (*menuToolButton)(nil)

// WrapMenuToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuToolButton(obj *externglib.Object) MenuToolButton {
	return menuToolButton{
		ToolButton:  WrapToolButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalMenuToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuToolButton(obj), nil
}

// SetArrowTooltipMarkup sets the tooltip markup text to be used as tooltip
// for the arrow button which pops up the menu. See
// gtk_tool_item_set_tooltip_text() for setting a tooltip on the whole
// MenuToolButton.
func (b menuToolButton) SetArrowTooltipMarkup(markup string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup(_arg0, _arg1)
}

// SetArrowTooltipText sets the tooltip text to be used as tooltip for the
// arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text()
// for setting a tooltip on the whole MenuToolButton.
func (b menuToolButton) SetArrowTooltipText(text string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_text(_arg0, _arg1)
}

// SetMenu sets the Menu that is popped up when the user clicks on the
// arrow. If @menu is NULL, the arrow button becomes insensitive.
func (b menuToolButton) SetMenu(menu Widget) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_tool_button_set_menu(_arg0, _arg1)
}
