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
	GTypeActionBar = coreglib.Type(girepository.MustFind("Gtk", "ActionBar").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionBar, F: marshalActionBar},
	})
}

// ActionBarOverrides contains methods that are overridable.
type ActionBarOverrides struct {
}

func defaultActionBarOverrides(v *ActionBar) ActionBarOverrides {
	return ActionBarOverrides{}
}

// ActionBar is designed to present contextual actions. It is expected to be
// displayed below the content and expand horizontally to fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*ActionBar)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ActionBar, *ActionBarClass, ActionBarOverrides](
		GTypeActionBar,
		initActionBarClass,
		wrapActionBar,
		defaultActionBarOverrides,
	)
}

func initActionBarClass(gclass unsafe.Pointer, overrides ActionBarOverrides, classInitFunc func(*ActionBarClass)) {
	if classInitFunc != nil {
		class := (*ActionBarClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapActionBar(obj *coreglib.Object) *ActionBar {
	return &ActionBar{
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
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	return wrapActionBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActionBarClass: instance of this type is always passed by reference.
type ActionBarClass struct {
	*actionBarClass
}

// actionBarClass is the struct that's finalized.
type actionBarClass struct {
	native unsafe.Pointer
}

var GIRInfoActionBarClass = girepository.MustFind("Gtk", "ActionBarClass")
