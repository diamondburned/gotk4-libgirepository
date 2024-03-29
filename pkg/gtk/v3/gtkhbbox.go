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
	GTypeHButtonBox = coreglib.Type(girepository.MustFind("Gtk", "HButtonBox").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHButtonBox, F: marshalHButtonBox},
	})
}

// HButtonBoxOverrides contains methods that are overridable.
type HButtonBoxOverrides struct {
}

func defaultHButtonBoxOverrides(v *HButtonBox) HButtonBoxOverrides {
	return HButtonBoxOverrides{}
}

type HButtonBox struct {
	_ [0]func() // equal guard
	ButtonBox
}

var (
	_ Containerer       = (*HButtonBox)(nil)
	_ coreglib.Objector = (*HButtonBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HButtonBox, *HButtonBoxClass, HButtonBoxOverrides](
		GTypeHButtonBox,
		initHButtonBoxClass,
		wrapHButtonBox,
		defaultHButtonBoxOverrides,
	)
}

func initHButtonBoxClass(gclass unsafe.Pointer, overrides HButtonBoxOverrides, classInitFunc func(*HButtonBoxClass)) {
	if classInitFunc != nil {
		class := (*HButtonBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHButtonBox(obj *coreglib.Object) *HButtonBox {
	return &HButtonBox{
		ButtonBox: ButtonBox{
			Box: Box{
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
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHButtonBox(p uintptr) (interface{}, error) {
	return wrapHButtonBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// HButtonBoxClass: instance of this type is always passed by reference.
type HButtonBoxClass struct {
	*hButtonBoxClass
}

// hButtonBoxClass is the struct that's finalized.
type hButtonBoxClass struct {
	native unsafe.Pointer
}

var GIRInfoHButtonBoxClass = girepository.MustFind("Gtk", "HButtonBoxClass")
