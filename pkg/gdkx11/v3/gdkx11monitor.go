// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// GTypeX11Monitor returns the GType for the type X11Monitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11Monitor() coreglib.Type {
	gtype := coreglib.Type(C.gdk_x11_monitor_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11Monitor)
	return gtype
}

// X11MonitorOverrider contains methods that are overridable.
type X11MonitorOverrider interface {
}

type X11Monitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ coreglib.Objector = (*X11Monitor)(nil)
)

func classInitX11Monitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Monitor(obj *coreglib.Object) *X11Monitor {
	return &X11Monitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	return wrapX11Monitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
