// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GTypeEmblemedIcon returns the GType for the type EmblemedIcon.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEmblemedIcon() coreglib.Type {
	gtype := coreglib.Type(C.g_emblemed_icon_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalEmblemedIcon)
	return gtype
}

// EmblemedIconOverrider contains methods that are overridable.
type EmblemedIconOverrider interface {
}

// EmblemedIcon is an implementation of #GIcon that supports adding an emblem to
// an icon. Adding multiple emblems to an icon is ensured via
// g_emblemed_icon_add_emblem().
//
// Note that Icon allows no control over the position of the emblems. See also
// #GEmblem for more information.
type EmblemedIcon struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Icon
}

var (
	_ coreglib.Objector = (*EmblemedIcon)(nil)
)

func classInitEmblemedIconner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEmblemedIcon(obj *coreglib.Object) *EmblemedIcon {
	return &EmblemedIcon{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblemedIcon(p uintptr) (interface{}, error) {
	return wrapEmblemedIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEmblemedIcon creates a new emblemed icon for icon with the emblem emblem.
//
// The function takes the following parameters:
//
//    - icon: #GIcon.
//    - emblem (optional) or NULL.
//
// The function returns the following values:
//
//    - emblemedIcon: new #GIcon.
//
func NewEmblemedIcon(icon Iconner, emblem *Emblem) *EmblemedIcon {
	var _arg1 *C.GIcon   // out
	var _arg2 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	if emblem != nil {
		_arg2 = (*C.GEmblem)(unsafe.Pointer(coreglib.InternObject(emblem).Native()))
	}

	_cret = C.g_emblemed_icon_new(_arg1, _arg2)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(emblem)

	var _emblemedIcon *EmblemedIcon // out

	_emblemedIcon = wrapEmblemedIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblemedIcon
}

// AddEmblem adds emblem to the #GList of #GEmblems.
//
// The function takes the following parameters:
//
//    - emblem: #GEmblem.
//
func (emblemed *EmblemedIcon) AddEmblem(emblem *Emblem) {
	var _arg0 *C.GEmblemedIcon // out
	var _arg1 *C.GEmblem       // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(coreglib.InternObject(emblemed).Native()))
	_arg1 = (*C.GEmblem)(unsafe.Pointer(coreglib.InternObject(emblem).Native()))

	C.g_emblemed_icon_add_emblem(_arg0, _arg1)
	runtime.KeepAlive(emblemed)
	runtime.KeepAlive(emblem)
}

// ClearEmblems removes all the emblems from icon.
func (emblemed *EmblemedIcon) ClearEmblems() {
	var _arg0 *C.GEmblemedIcon // out

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(coreglib.InternObject(emblemed).Native()))

	C.g_emblemed_icon_clear_emblems(_arg0)
	runtime.KeepAlive(emblemed)
}

// Emblems gets the list of emblems for the icon.
//
// The function returns the following values:
//
//    - list of #GEmblems that is owned by emblemed.
//
func (emblemed *EmblemedIcon) Emblems() []*Emblem {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GList         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(coreglib.InternObject(emblemed).Native()))

	_cret = C.g_emblemed_icon_get_emblems(_arg0)
	runtime.KeepAlive(emblemed)

	var _list []*Emblem // out

	_list = make([]*Emblem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GEmblem)(v)
		var dst *Emblem // out
		dst = wrapEmblem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// GetIcon gets the main icon for emblemed.
//
// The function returns the following values:
//
//    - icon that is owned by emblemed.
//
func (emblemed *EmblemedIcon) GetIcon() *Icon {
	var _arg0 *C.GEmblemedIcon // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GEmblemedIcon)(unsafe.Pointer(coreglib.InternObject(emblemed).Native()))

	_cret = C.g_emblemed_icon_get_icon(_arg0)
	runtime.KeepAlive(emblemed)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.Take(unsafe.Pointer(_cret)))

	return _icon
}
