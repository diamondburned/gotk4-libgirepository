// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_border_get_type()), F: marshalBorder},
	})
}

// Border: struct that specifies a border around a rectangular area.
//
// Each side can have different width.
//
// An instance of this type is always passed by reference.
type Border struct {
	*border
}

// border is the struct that's finalized.
type border struct {
	native *C.GtkBorder
}

func marshalBorder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Border{&border{(*C.GtkBorder)(unsafe.Pointer(b))}}, nil
}

// NewBorder constructs a struct Border.
func NewBorder() *Border {
	var _cret *C.GtkBorder // in

	_cret = C.gtk_border_new()

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_border)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_border_free((*C.GtkBorder)(intern.C))
		},
	)

	return _border
}

// Left: width of the left border
func (b *Border) Left() int16 {
	var v int16 // out
	v = int16(b.native.left)
	return v
}

// Right: width of the right border
func (b *Border) Right() int16 {
	var v int16 // out
	v = int16(b.native.right)
	return v
}

// Top: width of the top border
func (b *Border) Top() int16 {
	var v int16 // out
	v = int16(b.native.top)
	return v
}

// Bottom: width of the bottom border
func (b *Border) Bottom() int16 {
	var v int16 // out
	v = int16(b.native.bottom)
	return v
}

// Copy copies a Border-struct.
func (border_ *Border) Copy() *Border {
	var _arg0 *C.GtkBorder // out
	var _cret *C.GtkBorder // in

	_arg0 = (*C.GtkBorder)(gextras.StructNative(unsafe.Pointer(border_)))

	_cret = C.gtk_border_copy(_arg0)
	runtime.KeepAlive(border_)

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_border)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_border_free((*C.GtkBorder)(intern.C))
		},
	)

	return _border
}
