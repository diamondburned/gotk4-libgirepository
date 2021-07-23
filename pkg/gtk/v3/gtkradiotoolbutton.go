// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_tool_button_get_type()), F: marshalRadioToolButtonner},
	})
}

// RadioToolButton is a ToolItem that contains a radio button, that is, a button
// that is part of a group of toggle buttons where only one button can be active
// at a time.
//
// Use gtk_radio_tool_button_new() to create a new GtkRadioToolButton. Use
// gtk_radio_tool_button_new_from_widget() to create a new GtkRadioToolButton
// that is part of the same group as an existing GtkRadioToolButton.
//
//
// CSS nodes
//
// GtkRadioToolButton has a single CSS node with name toolbutton.
type RadioToolButton struct {
	ToggleToolButton
}

func wrapRadioToolButton(obj *externglib.Object) *RadioToolButton {
	return &RadioToolButton{
		ToggleToolButton: ToggleToolButton{
			ToolButton: ToolButton{
				ToolItem: ToolItem{
					Bin: Bin{
						Container: Container{
							Widget: Widget{
								InitiallyUnowned: externglib.InitiallyUnowned{
									Object: obj,
								},
								ImplementorIface: atk.ImplementorIface{
									Object: obj,
								},
								Buildable: Buildable{
									Object: obj,
								},
								Object: obj,
							},
						},
					},
					Activatable: Activatable{
						Object: obj,
					},
					Object: obj,
				},
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
				Object: obj,
			},
		},
	}
}

func marshalRadioToolButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioToolButton(obj), nil
}

// NewRadioToolButtonFromWidget creates a new RadioToolButton adding it to the
// same group as gruup
func NewRadioToolButtonFromWidget(group *RadioToolButton) *RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _cret *C.GtkToolItem        // in

	if group != nil {
		_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))
	}

	_cret = C.gtk_radio_tool_button_new_from_widget(_arg1)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonWithStockFromWidget creates a new RadioToolButton adding it
// to the same group as group. The new RadioToolButton will contain an icon and
// label from the stock item indicated by stock_id.
//
// Deprecated: gtk_radio_tool_button_new_from_widget.
func NewRadioToolButtonWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _arg2 *C.gchar              // out
	var _cret *C.GtkToolItem        // in

	if group != nil {
		_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_tool_button_new_with_stock_from_widget(_arg1, _arg2)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

func (*RadioToolButton) privateRadioToolButton() {}
