// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// ColorChooserInterface: instance of this type is always passed by reference.
type ColorChooserInterface struct {
	*colorChooserInterface
}

// colorChooserInterface is the struct that's finalized.
type colorChooserInterface struct {
	native unsafe.Pointer
}

var GIRInfoColorChooserInterface = girepository.MustFind("Gtk", "ColorChooserInterface")

func (c *ColorChooserInterface) Padding() [12]unsafe.Pointer {
	offset := GIRInfoColorChooserInterface.StructFieldOffset("padding")
	valptr := (*[12]unsafe.Pointer)(unsafe.Add(c.native, offset))
	var _v [12]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 12; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
