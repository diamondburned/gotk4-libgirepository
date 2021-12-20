// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_monitor_get_type()), F: marshalWaylandMonitorrer},
	})
}

// WaylandMonitor: wayland implementation of GdkMonitor.
//
// Beyond the gdk.Monitor API, the Wayland implementation offers access to the
// Wayland wl_output object with gdkwayland.WaylandMonitor.GetWlOutput().
type WaylandMonitor struct {
	gdk.Monitor

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*WaylandMonitor)(nil)
)

func wrapWaylandMonitor(obj *externglib.Object) *WaylandMonitor {
	return &WaylandMonitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalWaylandMonitorrer(p uintptr) (interface{}, error) {
	return wrapWaylandMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
