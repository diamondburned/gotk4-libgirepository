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
	GTypeNoOpObjectFactory = coreglib.Type(girepository.MustFind("Atk", "NoOpObjectFactory").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNoOpObjectFactory, F: marshalNoOpObjectFactory},
	})
}

// NoOpObjectFactoryOverrides contains methods that are overridable.
type NoOpObjectFactoryOverrides struct {
}

func defaultNoOpObjectFactoryOverrides(v *NoOpObjectFactory) NoOpObjectFactoryOverrides {
	return NoOpObjectFactoryOverrides{}
}

// NoOpObjectFactory: atkObjectFactory which creates an AtkNoOpObject. An
// instance of this is created by an AtkRegistry if no factory type has not been
// specified to create an accessible object of a particular type.
type NoOpObjectFactory struct {
	_ [0]func() // equal guard
	ObjectFactory
}

var (
	_ coreglib.Objector = (*NoOpObjectFactory)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*NoOpObjectFactory, *NoOpObjectFactoryClass, NoOpObjectFactoryOverrides](
		GTypeNoOpObjectFactory,
		initNoOpObjectFactoryClass,
		wrapNoOpObjectFactory,
		defaultNoOpObjectFactoryOverrides,
	)
}

func initNoOpObjectFactoryClass(gclass unsafe.Pointer, overrides NoOpObjectFactoryOverrides, classInitFunc func(*NoOpObjectFactoryClass)) {
	if classInitFunc != nil {
		class := (*NoOpObjectFactoryClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNoOpObjectFactory(obj *coreglib.Object) *NoOpObjectFactory {
	return &NoOpObjectFactory{
		ObjectFactory: ObjectFactory{
			Object: obj,
		},
	}
}

func marshalNoOpObjectFactory(p uintptr) (interface{}, error) {
	return wrapNoOpObjectFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NoOpObjectFactoryClass: instance of this type is always passed by reference.
type NoOpObjectFactoryClass struct {
	*noOpObjectFactoryClass
}

// noOpObjectFactoryClass is the struct that's finalized.
type noOpObjectFactoryClass struct {
	native unsafe.Pointer
}

var GIRInfoNoOpObjectFactoryClass = girepository.MustFind("Atk", "NoOpObjectFactoryClass")
