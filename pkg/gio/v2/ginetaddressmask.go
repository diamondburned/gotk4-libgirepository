// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// InetAddressMaskClass: instance of this type is always passed by reference.
type InetAddressMaskClass struct {
	*inetAddressMaskClass
}

// inetAddressMaskClass is the struct that's finalized.
type inetAddressMaskClass struct {
	native unsafe.Pointer
}

var GIRInfoInetAddressMaskClass = girepository.MustFind("Gio", "InetAddressMaskClass")
