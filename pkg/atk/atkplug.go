// Code generated by girgen. DO NOT EDIT.

package atk

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
	GTypePlug = coreglib.Type(girepository.MustFind("Atk", "Plug").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePlug, F: marshalPlug},
	})
}

// PlugOverrides contains methods that are overridable.
type PlugOverrides struct {
}

func defaultPlugOverrides(v *Plug) PlugOverrides {
	return PlugOverrides{}
}

// Plug: see Socket.
type Plug struct {
	_ [0]func() // equal guard
	AtkObject

	*coreglib.Object
	Component
}

var (
	_ coreglib.Objector = (*Plug)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Plug, *PlugClass, PlugOverrides](
		GTypePlug,
		initPlugClass,
		wrapPlug,
		defaultPlugOverrides,
	)
}

func initPlugClass(gclass unsafe.Pointer, overrides PlugOverrides, classInitFunc func(*PlugClass)) {
	if classInitFunc != nil {
		class := (*PlugClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPlug(obj *coreglib.Object) *Plug {
	return &Plug{
		AtkObject: AtkObject{
			Object: obj,
		},
		Object: obj,
		Component: Component{
			Object: obj,
		},
	}
}

func marshalPlug(p uintptr) (interface{}, error) {
	return wrapPlug(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PlugClass: instance of this type is always passed by reference.
type PlugClass struct {
	*plugClass
}

// plugClass is the struct that's finalized.
type plugClass struct {
	native unsafe.Pointer
}

var GIRInfoPlugClass = girepository.MustFind("Atk", "PlugClass")
