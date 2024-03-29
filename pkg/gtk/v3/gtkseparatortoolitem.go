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
import "C"

// GType values.
var (
	GTypeSeparatorToolItem = coreglib.Type(girepository.MustFind("Gtk", "SeparatorToolItem").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeparatorToolItem, F: marshalSeparatorToolItem},
	})
}

// SeparatorToolItemOverrides contains methods that are overridable.
type SeparatorToolItemOverrides struct {
}

func defaultSeparatorToolItemOverrides(v *SeparatorToolItem) SeparatorToolItemOverrides {
	return SeparatorToolItemOverrides{}
}

// SeparatorToolItem is a ToolItem that separates groups of other ToolItems.
// Depending on the theme, a SeparatorToolItem will often look like a vertical
// line on horizontally docked toolbars.
//
// If the Toolbar child property “expand” is TRUE and the property
// SeparatorToolItem:draw is FALSE, a SeparatorToolItem will act as a “spring”
// that forces other items to the ends of the toolbar.
//
// Use gtk_separator_tool_item_new() to create a new SeparatorToolItem.
//
//
// CSS nodes
//
// GtkSeparatorToolItem has a single CSS node with name separator.
type SeparatorToolItem struct {
	_ [0]func() // equal guard
	ToolItem
}

var (
	_ Binner            = (*SeparatorToolItem)(nil)
	_ coreglib.Objector = (*SeparatorToolItem)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SeparatorToolItem, *SeparatorToolItemClass, SeparatorToolItemOverrides](
		GTypeSeparatorToolItem,
		initSeparatorToolItemClass,
		wrapSeparatorToolItem,
		defaultSeparatorToolItemOverrides,
	)
}

func initSeparatorToolItemClass(gclass unsafe.Pointer, overrides SeparatorToolItemOverrides, classInitFunc func(*SeparatorToolItemClass)) {
	if classInitFunc != nil {
		class := (*SeparatorToolItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSeparatorToolItem(obj *coreglib.Object) *SeparatorToolItem {
	return &SeparatorToolItem{
		ToolItem: ToolItem{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalSeparatorToolItem(p uintptr) (interface{}, error) {
	return wrapSeparatorToolItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SeparatorToolItemClass: instance of this type is always passed by reference.
type SeparatorToolItemClass struct {
	*separatorToolItemClass
}

// separatorToolItemClass is the struct that's finalized.
type separatorToolItemClass struct {
	native unsafe.Pointer
}

var GIRInfoSeparatorToolItemClass = girepository.MustFind("Gtk", "SeparatorToolItemClass")
