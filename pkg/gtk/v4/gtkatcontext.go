// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_at_context_get_type()), F: marshalATContexter},
	})
}

// ATContexter describes ATContext's methods.
type ATContexter interface {
	// Accessible retrieves the `GtkAccessible` using this context.
	Accessible() *Accessible
	// AccessibleRole retrieves the accessible role of this context.
	AccessibleRole() AccessibleRole
}

// ATContext: `GtkATContext` is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a `GtkATContext` subclass, and is
// responsible for updating the accessible state in response to state changes in
// `GtkAccessible`.
type ATContext struct {
	*externglib.Object
}

var (
	_ ATContexter     = (*ATContext)(nil)
	_ gextras.Nativer = (*ATContext)(nil)
)

func wrapATContext(obj *externglib.Object) ATContexter {
	return &ATContext{
		Object: obj,
	}
}

func marshalATContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapATContext(obj), nil
}

// Accessible retrieves the `GtkAccessible` using this context.
func (self *ATContext) Accessible() *Accessible {
	var _arg0 *C.GtkATContext  // out
	var _cret *C.GtkAccessible // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_at_context_get_accessible(_arg0)

	var _accessible *Accessible // out

	_accessible = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Accessible)

	return _accessible
}

// AccessibleRole retrieves the accessible role of this context.
func (self *ATContext) AccessibleRole() AccessibleRole {
	var _arg0 *C.GtkATContext     // out
	var _cret C.GtkAccessibleRole // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_at_context_get_accessible_role(_arg0)

	var _accessibleRole AccessibleRole // out

	_accessibleRole = AccessibleRole(_cret)

	return _accessibleRole
}
