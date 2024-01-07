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
	GTypeLayout = coreglib.Type(girepository.MustFind("Gtk", "Layout").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLayout, F: marshalLayout},
	})
}

// LayoutOverrides contains methods that are overridable.
type LayoutOverrides struct {
}

func defaultLayoutOverrides(v *Layout) LayoutOverrides {
	return LayoutOverrides{}
}

// Layout is similar to DrawingArea in that it’s a “blank slate” and doesn’t do
// anything except paint a blank background by default. It’s different in that
// it supports scrolling natively due to implementing Scrollable, and can
// contain child widgets since it’s a Container.
//
// If you just want to draw, a DrawingArea is a better choice since it has lower
// overhead. If you just need to position child widgets at specific points, then
// Fixed provides that functionality on its own.
//
// When handling expose events on a Layout, you must draw to the Window returned
// by gtk_layout_get_bin_window(), rather than to the one returned by
// gtk_widget_get_window() as you would for a DrawingArea.
type Layout struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Scrollable
}

var (
	_ Containerer       = (*Layout)(nil)
	_ coreglib.Objector = (*Layout)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Layout, *LayoutClass, LayoutOverrides](
		GTypeLayout,
		initLayoutClass,
		wrapLayout,
		defaultLayoutOverrides,
	)
}

func initLayoutClass(gclass unsafe.Pointer, overrides LayoutOverrides, classInitFunc func(*LayoutClass)) {
	if classInitFunc != nil {
		class := (*LayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLayout(obj *coreglib.Object) *Layout {
	return &Layout{
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
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalLayout(p uintptr) (interface{}, error) {
	return wrapLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LayoutClass: instance of this type is always passed by reference.
type LayoutClass struct {
	*layoutClass
}

// layoutClass is the struct that's finalized.
type layoutClass struct {
	native unsafe.Pointer
}

var GIRInfoLayoutClass = girepository.MustFind("Gtk", "LayoutClass")
