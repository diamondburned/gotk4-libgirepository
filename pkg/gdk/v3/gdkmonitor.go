// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_Monitor_ConnectInvalidate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeMonitor = coreglib.Type(girepository.MustFind("Gdk", "Monitor").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMonitor, F: marshalMonitor},
	})
}

// Monitor objects represent the individual outputs that are associated with a
// Display. GdkDisplay has APIs to enumerate monitors with
// gdk_display_get_n_monitors() and gdk_display_get_monitor(), and to find
// particular monitors with gdk_display_get_primary_monitor() or
// gdk_display_get_monitor_at_window().
//
// GdkMonitor was introduced in GTK+ 3.22 and supersedes earlier APIs in
// GdkScreen to obtain monitor-related information.
type Monitor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Monitor)(nil)
)

func wrapMonitor(obj *coreglib.Object) *Monitor {
	return &Monitor{
		Object: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	return wrapMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Monitor) ConnectInvalidate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "invalidate", false, unsafe.Pointer(C._gotk4_gdk3_Monitor_ConnectInvalidate), f)
}
