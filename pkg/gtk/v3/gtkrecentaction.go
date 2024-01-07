// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
	GTypeRecentAction = coreglib.Type(girepository.MustFind("Gtk", "RecentAction").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentAction, F: marshalRecentAction},
	})
}

// RecentActionOverrides contains methods that are overridable.
type RecentActionOverrides struct {
}

func defaultRecentActionOverrides(v *RecentAction) RecentActionOverrides {
	return RecentActionOverrides{}
}

// RecentAction represents a list of recently used files, which can be shown by
// widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction struct {
	_ [0]func() // equal guard
	Action

	*coreglib.Object
	RecentChooser
}

var (
	_ coreglib.Objector = (*RecentAction)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RecentAction, *RecentActionClass, RecentActionOverrides](
		GTypeRecentAction,
		initRecentActionClass,
		wrapRecentAction,
		defaultRecentActionOverrides,
	)
}

func initRecentActionClass(gclass unsafe.Pointer, overrides RecentActionOverrides, classInitFunc func(*RecentActionClass)) {
	if classInitFunc != nil {
		class := (*RecentActionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRecentAction(obj *coreglib.Object) *RecentAction {
	return &RecentAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentAction(p uintptr) (interface{}, error) {
	return wrapRecentAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RecentActionClass: instance of this type is always passed by reference.
type RecentActionClass struct {
	*recentActionClass
}

// recentActionClass is the struct that's finalized.
type recentActionClass struct {
	native unsafe.Pointer
}

var GIRInfoRecentActionClass = girepository.MustFind("Gtk", "RecentActionClass")
