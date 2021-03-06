// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GTypeNativeVolumeMonitor returns the GType for the type NativeVolumeMonitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNativeVolumeMonitor() coreglib.Type {
	gtype := coreglib.Type(C.g_native_volume_monitor_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalNativeVolumeMonitor)
	return gtype
}

const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-native-volume-monitor"

// NativeVolumeMonitorOverrider contains methods that are overridable.
type NativeVolumeMonitorOverrider interface {
}

type NativeVolumeMonitor struct {
	_ [0]func() // equal guard
	VolumeMonitor
}

var (
	_ coreglib.Objector = (*NativeVolumeMonitor)(nil)
)

// NativeVolumeMonitorrer describes types inherited from class NativeVolumeMonitor.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type NativeVolumeMonitorrer interface {
	coreglib.Objector
	baseNativeVolumeMonitor() *NativeVolumeMonitor
}

var _ NativeVolumeMonitorrer = (*NativeVolumeMonitor)(nil)

func classInitNativeVolumeMonitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNativeVolumeMonitor(obj *coreglib.Object) *NativeVolumeMonitor {
	return &NativeVolumeMonitor{
		VolumeMonitor: VolumeMonitor{
			Object: obj,
		},
	}
}

func marshalNativeVolumeMonitor(p uintptr) (interface{}, error) {
	return wrapNativeVolumeMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *NativeVolumeMonitor) baseNativeVolumeMonitor() *NativeVolumeMonitor {
	return v
}

// BaseNativeVolumeMonitor returns the underlying base object.
func BaseNativeVolumeMonitor(obj NativeVolumeMonitorrer) *NativeVolumeMonitor {
	return obj.baseNativeVolumeMonitor()
}
