// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkx11monitor.go.
var GTypeX11Monitor = coreglib.Type(C.gdk_x11_monitor_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeX11Monitor, F: marshalX11Monitor},
	})
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

// Workarea retrieves the size and position of the “work area” on a monitor
// within the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
//
// The function returns the following values:
//
//    - workarea to be filled with the monitor workarea.
//
func (monitor *X11Monitor) Workarea() *gdk.Rectangle {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	girepository.MustFind("GdkX11", "X11Monitor").InvokeMethod("get_workarea", _args[:], _outs[:])

	runtime.KeepAlive(monitor)

	var _workarea *gdk.Rectangle // out

	_workarea = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))

	return _workarea
}
