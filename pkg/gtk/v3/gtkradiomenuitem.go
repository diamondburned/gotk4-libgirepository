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
// extern void _gotk4_gtk3_RadioMenuItem_ConnectGroupChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeRadioMenuItem = coreglib.Type(girepository.MustFind("Gtk", "RadioMenuItem").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRadioMenuItem, F: marshalRadioMenuItem},
	})
}

// RadioMenuItemOverrides contains methods that are overridable.
type RadioMenuItemOverrides struct {
}

func defaultRadioMenuItemOverrides(v *RadioMenuItem) RadioMenuItemOverrides {
	return RadioMenuItemOverrides{}
}

// RadioMenuItem: radio menu item is a check menu item that belongs to a group.
// At each instant exactly one of the radio menu items from a group is selected.
//
// The group list does not need to be freed, as each RadioMenuItem will remove
// itself and its list item when it is destroyed.
//
// The correct way to create a group of radio menu items is approximatively
// this:
//
// How to create a group of radio menu items.
//
//    menuitem
//    ├── radio.left
//    ╰── <child>
//
// GtkRadioMenuItem has a main CSS node with name menuitem, and a subnode with
// name radio, which gets the .left or .right style class.
type RadioMenuItem struct {
	_ [0]func() // equal guard
	CheckMenuItem
}

var (
	_ Binner            = (*RadioMenuItem)(nil)
	_ coreglib.Objector = (*RadioMenuItem)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RadioMenuItem, *RadioMenuItemClass, RadioMenuItemOverrides](
		GTypeRadioMenuItem,
		initRadioMenuItemClass,
		wrapRadioMenuItem,
		defaultRadioMenuItemOverrides,
	)
}

func initRadioMenuItemClass(gclass unsafe.Pointer, overrides RadioMenuItemOverrides, classInitFunc func(*RadioMenuItemClass)) {
	if classInitFunc != nil {
		class := (*RadioMenuItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRadioMenuItem(obj *coreglib.Object) *RadioMenuItem {
	return &RadioMenuItem{
		CheckMenuItem: CheckMenuItem{
			MenuItem: MenuItem{
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
	}
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	return wrapRadioMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *RadioMenuItem) ConnectGroupChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "group-changed", false, unsafe.Pointer(C._gotk4_gtk3_RadioMenuItem_ConnectGroupChanged), f)
}

// RadioMenuItemClass: instance of this type is always passed by reference.
type RadioMenuItemClass struct {
	*radioMenuItemClass
}

// radioMenuItemClass is the struct that's finalized.
type radioMenuItemClass struct {
	native unsafe.Pointer
}

var GIRInfoRadioMenuItemClass = girepository.MustFind("Gtk", "RadioMenuItemClass")
