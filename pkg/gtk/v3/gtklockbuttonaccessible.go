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
	GTypeLockButtonAccessible = coreglib.Type(girepository.MustFind("Gtk", "LockButtonAccessible").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLockButtonAccessible, F: marshalLockButtonAccessible},
	})
}

// LockButtonAccessibleOverrides contains methods that are overridable.
type LockButtonAccessibleOverrides struct {
}

func defaultLockButtonAccessibleOverrides(v *LockButtonAccessible) LockButtonAccessibleOverrides {
	return LockButtonAccessibleOverrides{}
}

type LockButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible
}

var (
	_ coreglib.Objector = (*LockButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*LockButtonAccessible, *LockButtonAccessibleClass, LockButtonAccessibleOverrides](
		GTypeLockButtonAccessible,
		initLockButtonAccessibleClass,
		wrapLockButtonAccessible,
		defaultLockButtonAccessibleOverrides,
	)
}

func initLockButtonAccessibleClass(gclass unsafe.Pointer, overrides LockButtonAccessibleOverrides, classInitFunc func(*LockButtonAccessibleClass)) {
	if classInitFunc != nil {
		class := (*LockButtonAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLockButtonAccessible(obj *coreglib.Object) *LockButtonAccessible {
	return &LockButtonAccessible{
		ButtonAccessible: ButtonAccessible{
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
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
	}
}

func marshalLockButtonAccessible(p uintptr) (interface{}, error) {
	return wrapLockButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LockButtonAccessibleClass: instance of this type is always passed by
// reference.
type LockButtonAccessibleClass struct {
	*lockButtonAccessibleClass
}

// lockButtonAccessibleClass is the struct that's finalized.
type lockButtonAccessibleClass struct {
	native unsafe.Pointer
}

var GIRInfoLockButtonAccessibleClass = girepository.MustFind("Gtk", "LockButtonAccessibleClass")
