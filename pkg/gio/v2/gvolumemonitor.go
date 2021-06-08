// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_volume_monitor_get_type()), F: marshalVolumeMonitor},
	})
}

// VolumeMonitor is for listing the user interesting devices and volumes on the
// computer. In other words, what a file selector or file manager would show in
// a sidebar.
//
// Monitor is not [thread-default-context
// aware][g-main-context-push-thread-default], and so should not be used other
// than from the main thread, with no thread-default-context active.
//
// In order to receive updates about volumes and mounts monitored through GVFS,
// a main loop must be running.
type VolumeMonitor interface {
	gextras.Objector

	// ConnectedDrives gets a list of drives connected to the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	ConnectedDrives() *glib.List
	// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
	MountForUUID(uuiD string) Mount
	// Mounts gets a list of the mounts on the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	Mounts() *glib.List
	// VolumeForUUID finds a #GVolume object by its UUID (see
	// g_volume_get_uuid())
	VolumeForUUID(uuiD string) Volume
	// Volumes gets a list of the volumes on the system.
	//
	// The returned list should be freed with g_list_free(), after its elements
	// have been unreffed with g_object_unref().
	Volumes() *glib.List
}

// volumeMonitor implements the VolumeMonitor interface.
type volumeMonitor struct {
	gextras.Objector
}

var _ VolumeMonitor = (*volumeMonitor)(nil)

// WrapVolumeMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapVolumeMonitor(obj *externglib.Object) VolumeMonitor {
	return VolumeMonitor{
		Objector: obj,
	}
}

func marshalVolumeMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVolumeMonitor(obj), nil
}

// ConnectedDrives gets a list of drives connected to the system.
//
// The returned list should be freed with g_list_free(), after its elements
// have been unreffed with g_object_unref().
func (v volumeMonitor) ConnectedDrives() *glib.List {
	var arg0 *C.GVolumeMonitor

	arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))

	cret := new(C.GList)
	var goret *glib.List

	cret = C.g_volume_monitor_get_connected_drives(arg0)

	goret = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
func (v volumeMonitor) MountForUUID(uuiD string) Mount {
	var arg0 *C.GVolumeMonitor
	var arg1 *C.char

	arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(uuiD))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.GMount)
	var goret Mount

	cret = C.g_volume_monitor_get_mount_for_uuid(arg0, arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Mount)

	return goret
}

// Mounts gets a list of the mounts on the system.
//
// The returned list should be freed with g_list_free(), after its elements
// have been unreffed with g_object_unref().
func (v volumeMonitor) Mounts() *glib.List {
	var arg0 *C.GVolumeMonitor

	arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))

	cret := new(C.GList)
	var goret *glib.List

	cret = C.g_volume_monitor_get_mounts(arg0)

	goret = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// VolumeForUUID finds a #GVolume object by its UUID (see
// g_volume_get_uuid())
func (v volumeMonitor) VolumeForUUID(uuiD string) Volume {
	var arg0 *C.GVolumeMonitor
	var arg1 *C.char

	arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(uuiD))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.GVolume)
	var goret Volume

	cret = C.g_volume_monitor_get_volume_for_uuid(arg0, arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Volume)

	return goret
}

// Volumes gets a list of the volumes on the system.
//
// The returned list should be freed with g_list_free(), after its elements
// have been unreffed with g_object_unref().
func (v volumeMonitor) Volumes() *glib.List {
	var arg0 *C.GVolumeMonitor

	arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))

	cret := new(C.GList)
	var goret *glib.List

	cret = C.g_volume_monitor_get_volumes(arg0)

	goret = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}
