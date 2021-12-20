// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_emblem_get_type()), F: marshalEmblemmer},
	})
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem struct {
	*externglib.Object

	Icon

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*Emblem)(nil)
)

func wrapEmblem(obj *externglib.Object) *Emblem {
	return &Emblem{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblemmer(p uintptr) (interface{}, error) {
	return wrapEmblem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEmblem creates a new emblem for icon.
//
// The function takes the following parameters:
//
//    - icon: GIcon containing the icon.
//
// The function returns the following values:
//
//    - emblem: new #GEmblem.
//
func NewEmblem(icon Iconner) *Emblem {
	var _arg1 *C.GIcon   // out
	var _cret *C.GEmblem // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_emblem_new(_arg1)
	runtime.KeepAlive(icon)

	var _emblem *Emblem // out

	_emblem = wrapEmblem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblem
}

// NewEmblemWithOrigin creates a new emblem for icon.
//
// The function takes the following parameters:
//
//    - icon: GIcon containing the icon.
//    - origin: GEmblemOrigin enum defining the emblem's origin.
//
// The function returns the following values:
//
//    - emblem: new #GEmblem.
//
func NewEmblemWithOrigin(icon Iconner, origin EmblemOrigin) *Emblem {
	var _arg1 *C.GIcon        // out
	var _arg2 C.GEmblemOrigin // out
	var _cret *C.GEmblem      // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = C.GEmblemOrigin(origin)

	_cret = C.g_emblem_new_with_origin(_arg1, _arg2)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(origin)

	var _emblem *Emblem // out

	_emblem = wrapEmblem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblem
}

// GetIcon gives back the icon from emblem.
//
// The function returns the following values:
//
//    - icon The returned object belongs to the emblem and should not be modified
//      or freed.
//
func (emblem *Emblem) GetIcon() Iconner {
	var _arg0 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	_cret = C.g_emblem_get_icon(_arg0)
	runtime.KeepAlive(emblem)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Iconner)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// Origin gets the origin of the emblem.
//
// The function returns the following values:
//
//    - emblemOrigin: origin of the emblem.
//
func (emblem *Emblem) Origin() EmblemOrigin {
	var _arg0 *C.GEmblem      // out
	var _cret C.GEmblemOrigin // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	_cret = C.g_emblem_get_origin(_arg0)
	runtime.KeepAlive(emblem)

	var _emblemOrigin EmblemOrigin // out

	_emblemOrigin = EmblemOrigin(_cret)

	return _emblemOrigin
}
