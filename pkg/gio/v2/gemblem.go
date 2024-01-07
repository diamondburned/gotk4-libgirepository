// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
	GTypeEmblem = coreglib.Type(girepository.MustFind("Gio", "Emblem").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEmblem, F: marshalEmblem},
	})
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Icon
}

var (
	_ coreglib.Objector = (*Emblem)(nil)
)

func wrapEmblem(obj *coreglib.Object) *Emblem {
	return &Emblem{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblem(p uintptr) (interface{}, error) {
	return wrapEmblem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
