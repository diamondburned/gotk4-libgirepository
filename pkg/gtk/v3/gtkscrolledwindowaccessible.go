// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeScrolledWindowAccessible = coreglib.Type(C.gtk_scrolled_window_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrolledWindowAccessible, F: marshalScrolledWindowAccessible},
	})
}

// ScrolledWindowAccessibleOverrider contains methods that are overridable.
type ScrolledWindowAccessibleOverrider interface {
}

type ScrolledWindowAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*ScrolledWindowAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeScrolledWindowAccessible,
		GoType:    reflect.TypeOf((*ScrolledWindowAccessible)(nil)),
		InitClass: initClassScrolledWindowAccessible,
	})
}

func initClassScrolledWindowAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitScrolledWindowAccessible(*ScrolledWindowAccessibleClass)
	}); ok {
		klass := (*ScrolledWindowAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitScrolledWindowAccessible(klass)
	}
}

func wrapScrolledWindowAccessible(obj *coreglib.Object) *ScrolledWindowAccessible {
	return &ScrolledWindowAccessible{
		ContainerAccessible: ContainerAccessible{
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
		},
	}
}

func marshalScrolledWindowAccessible(p uintptr) (interface{}, error) {
	return wrapScrolledWindowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ScrolledWindowAccessibleClass: instance of this type is always passed by
// reference.
type ScrolledWindowAccessibleClass struct {
	*scrolledWindowAccessibleClass
}

// scrolledWindowAccessibleClass is the struct that's finalized.
type scrolledWindowAccessibleClass struct {
	native *C.GtkScrolledWindowAccessibleClass
}

func (s *ScrolledWindowAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &s.native.parent_class
	var v *ContainerAccessibleClass // out
	v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
