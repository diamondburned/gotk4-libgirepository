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
	GTypeWindowControls = coreglib.Type(girepository.MustFind("Gtk", "WindowControls").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowControls, F: marshalWindowControls},
	})
}

// WindowControlsOverrides contains methods that are overridable.
type WindowControlsOverrides struct {
}

func defaultWindowControlsOverrides(v *WindowControls) WindowControlsOverrides {
	return WindowControlsOverrides{}
}

// WindowControls: GtkWindowControls shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// GtkWindowControls only displays start or end side of the controls (see
// gtk.WindowControls:side), so it's intended to be always used in pair with
// another GtkWindowControls for the opposite side, for example:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">start</property>
//        </object>
//      </child>
//
//      ...
//
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">end</property>
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    windowcontrols
//    ├── [image.icon]
//    ├── [button.minimize]
//    ├── [button.maximize]
//    ╰── [button.close]
//
//
// A GtkWindowControls' CSS node is called windowcontrols. It contains subnodes
// corresponding to each title button. Which of the title buttons exist and
// where they are placed exactly depends on the desktop environment and
// gtk.WindowControls:decoration-layout value.
//
// When gtk.WindowControls:empty is TRUE, it gets the .empty style class.
//
//
// Accessibility
//
// GtkWindowControls uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*WindowControls)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WindowControls, *WindowControlsClass, WindowControlsOverrides](
		GTypeWindowControls,
		initWindowControlsClass,
		wrapWindowControls,
		defaultWindowControlsOverrides,
	)
}

func initWindowControlsClass(gclass unsafe.Pointer, overrides WindowControlsOverrides, classInitFunc func(*WindowControlsClass)) {
	if classInitFunc != nil {
		class := (*WindowControlsClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWindowControls(obj *coreglib.Object) *WindowControls {
	return &WindowControls{
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
	}
}

func marshalWindowControls(p uintptr) (interface{}, error) {
	return wrapWindowControls(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WindowControlsClass: instance of this type is always passed by reference.
type WindowControlsClass struct {
	*windowControlsClass
}

// windowControlsClass is the struct that's finalized.
type windowControlsClass struct {
	native unsafe.Pointer
}

var GIRInfoWindowControlsClass = girepository.MustFind("Gtk", "WindowControlsClass")
