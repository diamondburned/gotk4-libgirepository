// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_iter_get_type()), F: marshalPixbufSimpleAnimIterer},
	})
}

func PixbufErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_pixbuf_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type PixbufSimpleAnimIter struct {
	PixbufAnimationIter
}

func wrapPixbufSimpleAnimIter(obj *externglib.Object) *PixbufSimpleAnimIter {
	return &PixbufSimpleAnimIter{
		PixbufAnimationIter: PixbufAnimationIter{
			Object: obj,
		},
	}
}

func marshalPixbufSimpleAnimIterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufSimpleAnimIter(obj), nil
}

func (*PixbufSimpleAnimIter) privatePixbufSimpleAnimIter() {}
