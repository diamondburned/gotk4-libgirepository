// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_ATContext_ConnectStateChange(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeATContext = coreglib.Type(girepository.MustFind("Gtk", "ATContext").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeATContext, F: marshalATContext},
	})
}

// ATContext: GtkATContext is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a GtkATContext subclass, and is
// responsible for updating the accessible state in response to state changes in
// GtkAccessible.
type ATContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ATContext)(nil)
)

// ATContexter describes types inherited from class ATContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ATContexter interface {
	coreglib.Objector
	baseATContext() *ATContext
}

var _ ATContexter = (*ATContext)(nil)

func wrapATContext(obj *coreglib.Object) *ATContext {
	return &ATContext{
		Object: obj,
	}
}

func marshalATContext(p uintptr) (interface{}, error) {
	return wrapATContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ATContext) baseATContext() *ATContext {
	return v
}

// BaseATContext returns the underlying base object.
func BaseATContext(obj ATContexter) *ATContext {
	return obj.baseATContext()
}

// ConnectStateChange is emitted when the attributes of the accessible for the
// GtkATContext instance change.
func (v *ATContext) ConnectStateChange(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "state-change", false, unsafe.Pointer(C._gotk4_gtk4_ATContext_ConnectStateChange), f)
}
