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
	GTypeLabelAccessible = coreglib.Type(girepository.MustFind("Gtk", "LabelAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLabelAccessible, F: marshalLabelAccessible},
	})
}

// LabelAccessibleOverrides contains methods that are overridable.
type LabelAccessibleOverrides struct {
}

func defaultLabelAccessibleOverrides(v *LabelAccessible) LabelAccessibleOverrides {
	return LabelAccessibleOverrides{}
}

type LabelAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	*coreglib.Object
	atk.Hypertext
	atk.Text
}

var (
	_ coreglib.Objector = (*LabelAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*LabelAccessible, *LabelAccessibleClass, LabelAccessibleOverrides](
		GTypeLabelAccessible,
		initLabelAccessibleClass,
		wrapLabelAccessible,
		defaultLabelAccessibleOverrides,
	)
}

func initLabelAccessibleClass(gclass unsafe.Pointer, overrides LabelAccessibleOverrides, classInitFunc func(*LabelAccessibleClass)) {
	if classInitFunc != nil {
		class := (*LabelAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLabelAccessible(obj *coreglib.Object) *LabelAccessible {
	return &LabelAccessible{
		WidgetAccessible: WidgetAccessible{
			Accessible: Accessible{
				AtkObject: atk.AtkObject{
					Object: obj,
				},
			},
			Component: atk.Component{
				Object: obj,
			},
		},
		Object: obj,
		Hypertext: atk.Hypertext{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
	}
}

func marshalLabelAccessible(p uintptr) (interface{}, error) {
	return wrapLabelAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LabelAccessibleClass: instance of this type is always passed by reference.
type LabelAccessibleClass struct {
	*labelAccessibleClass
}

// labelAccessibleClass is the struct that's finalized.
type labelAccessibleClass struct {
	native unsafe.Pointer
}

var GIRInfoLabelAccessibleClass = girepository.MustFind("Gtk", "LabelAccessibleClass")
