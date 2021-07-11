// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_plug_get_type()), F: marshalPlugger},
	})
}

// PlugOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PlugOverrider interface {
	//
	ObjectID() string
}

// Plugger describes Plug's methods.
type Plugger interface {
	// ID gets the unique ID of an Plug object, which can be used to embed
	// inside of an Socket using atk_socket_embed().
	ID() string
	// SetChild sets @child as accessible child of @plug and @plug as accessible
	// parent of @child.
	SetChild(child ObjectClasser)
}

// Plug: see Socket
type Plug struct {
	ObjectClass

	Component
}

var (
	_ Plugger         = (*Plug)(nil)
	_ gextras.Nativer = (*Plug)(nil)
)

func wrapPlug(obj *externglib.Object) Plugger {
	return &Plug{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
	}
}

func marshalPlugger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPlug(obj), nil
}

// NewPlug creates a new Plug instance.
func NewPlug() *Plug {
	var _cret *C.AtkObject // in

	_cret = C.atk_plug_new()

	var _plug *Plug // out

	_plug = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Plug)

	return _plug
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *Plug) Native() uintptr {
	return v.ObjectClass.Object.Native()
}

// ID gets the unique ID of an Plug object, which can be used to embed inside of
// an Socket using atk_socket_embed().
//
// Internally, this calls a class function that should be registered by the IPC
// layer (usually at-spi2-atk). The implementor of an Plug object should call
// this function (after atk-bridge is loaded) and pass the value to the process
// implementing the Socket, so it could embed the plug.
func (plug *Plug) ID() string {
	var _arg0 *C.AtkPlug // out
	var _cret *C.gchar   // in

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(plug.Native()))

	_cret = C.atk_plug_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetChild sets @child as accessible child of @plug and @plug as accessible
// parent of @child. @child can be NULL.
//
// In some cases, one can not use the AtkPlug type directly as accessible object
// for the toplevel widget of the application. For instance in the gtk case,
// GtkPlugAccessible can not inherit both from GtkWindowAccessible and from
// AtkPlug. In such a case, one can create, in addition to the standard
// accessible object for the toplevel widget, an AtkPlug object, and make the
// former the child of the latter by calling atk_plug_set_child().
func (plug *Plug) SetChild(child ObjectClasser) {
	var _arg0 *C.AtkPlug   // out
	var _arg1 *C.AtkObject // out

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(plug.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.atk_plug_set_child(_arg0, _arg1)
}
