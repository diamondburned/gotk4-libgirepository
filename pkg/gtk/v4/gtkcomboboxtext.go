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
import "C"

// GType values.
var (
	GTypeComboBoxText = coreglib.Type(girepository.MustFind("Gtk", "ComboBoxText").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeComboBoxText, F: marshalComboBoxText},
	})
}

// ComboBoxText: GtkComboBoxText is a simple variant of GtkComboBox for
// text-only use cases.
//
// !An example GtkComboBoxText (combo-box-text.png)
//
// GtkComboBoxText hides the model-view complexity of GtkComboBox.
//
// To create a GtkComboBoxText, use gtk.ComboBoxText.New or
// gtk.ComboBoxText.NewWithEntry.
//
// You can add items to a GtkComboBoxText with gtk.ComboBoxText.AppendText(),
// gtk.ComboBoxText.InsertText() or gtk.ComboBoxText.PrependText() and remove
// options with gtk.ComboBoxText.Remove().
//
// If the GtkComboBoxText contains an entry (via the gtk.ComboBox:has-entry
// property), its contents can be retrieved using
// gtk.ComboBoxText.GetActiveText().
//
// You should not call gtk.ComboBox.SetModel() or attempt to pack more cells
// into this combo box via its gtk.CellLayout interface.
//
//
// GtkComboBoxText as GtkBuildable
//
// The GtkComboBoxText implementation of the GtkBuildable interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element can specify the “id”
// corresponding to the appended text and also supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying GtkComboBoxText items:
//
//    <object class="GtkComboBoxText">
//      <items>
//        <item translatable="yes" id="factory">Factory</item>
//        <item translatable="yes" id="home">Home</item>
//        <item translatable="yes" id="subway">Subway</item>
//      </items>
//    </object>
//
//
// CSS nodes
//
//    combobox
//    ╰── box.linked
//        ├── entry.combo
//        ├── button.combo
//        ╰── window.popup
//
//
// GtkComboBoxText has a single CSS node with name combobox. It adds the style
// class .combo to the main CSS nodes of its entry and button children, and the
// .linked class to the node of its internal box.
type ComboBoxText struct {
	_ [0]func() // equal guard
	ComboBox
}

var (
	_ Widgetter         = (*ComboBoxText)(nil)
	_ coreglib.Objector = (*ComboBoxText)(nil)
)

func wrapComboBoxText(obj *coreglib.Object) *ComboBoxText {
	return &ComboBoxText{
		ComboBox: ComboBox{
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
			CellEditable: CellEditable{
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
			CellLayout: CellLayout{
				Object: obj,
			},
		},
	}
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	return wrapComboBoxText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
