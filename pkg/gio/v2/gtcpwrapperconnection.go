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

// TCPWrapperConnectionClass: instance of this type is always passed by
// reference.
type TCPWrapperConnectionClass struct {
	*tcpWrapperConnectionClass
}

// tcpWrapperConnectionClass is the struct that's finalized.
type tcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

var GIRInfoTCPWrapperConnectionClass = girepository.MustFind("Gio", "TcpWrapperConnectionClass")
