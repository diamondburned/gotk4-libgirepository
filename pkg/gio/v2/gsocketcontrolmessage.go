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

// SocketControlMessageClass class structure for ControlMessage.
//
// An instance of this type is always passed by reference.
type SocketControlMessageClass struct {
	*socketControlMessageClass
}

// socketControlMessageClass is the struct that's finalized.
type socketControlMessageClass struct {
	native unsafe.Pointer
}

var GIRInfoSocketControlMessageClass = girepository.MustFind("Gio", "SocketControlMessageClass")
