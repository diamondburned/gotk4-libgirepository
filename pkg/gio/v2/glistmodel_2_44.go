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

// ListModelInterface: virtual function table for Model.
//
// An instance of this type is always passed by reference.
type ListModelInterface struct {
	*listModelInterface
}

// listModelInterface is the struct that's finalized.
type listModelInterface struct {
	native unsafe.Pointer
}

var GIRInfoListModelInterface = girepository.MustFind("Gio", "ListModelInterface")
