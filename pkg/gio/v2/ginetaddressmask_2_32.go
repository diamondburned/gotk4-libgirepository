// Code generated by girgen. DO NOT EDIT.

package gio

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
	GTypeInetAddressMask = coreglib.Type(girepository.MustFind("Gio", "InetAddressMask").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInetAddressMask, F: marshalInetAddressMask},
	})
}

// InetAddressMaskOverrides contains methods that are overridable.
type InetAddressMaskOverrides struct {
}

func defaultInetAddressMaskOverrides(v *InetAddressMask) InetAddressMaskOverrides {
	return InetAddressMaskOverrides{}
}

// InetAddressMask represents a range of IPv4 or IPv6 addresses described by a
// base address and a length indicating how many bits of the base address are
// relevant for matching purposes. These are often given in string form. Eg,
// "10.0.0.0/8", or "fe80::/10".
type InetAddressMask struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Initable
}

var (
	_ coreglib.Objector = (*InetAddressMask)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*InetAddressMask, *InetAddressMaskClass, InetAddressMaskOverrides](
		GTypeInetAddressMask,
		initInetAddressMaskClass,
		wrapInetAddressMask,
		defaultInetAddressMaskOverrides,
	)
}

func initInetAddressMaskClass(gclass unsafe.Pointer, overrides InetAddressMaskOverrides, classInitFunc func(*InetAddressMaskClass)) {
	if classInitFunc != nil {
		class := (*InetAddressMaskClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapInetAddressMask(obj *coreglib.Object) *InetAddressMask {
	return &InetAddressMask{
		Object: obj,
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalInetAddressMask(p uintptr) (interface{}, error) {
	return wrapInetAddressMask(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}