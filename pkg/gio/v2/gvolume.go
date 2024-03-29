// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_Volume_ConnectRemoved(gpointer, guintptr);
// extern void _gotk4_gio2_Volume_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeVolume = coreglib.Type(girepository.MustFind("Gio", "Volume").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVolume, F: marshalVolume},
	})
}

// VOLUME_IDENTIFIER_KIND_CLASS: string used to obtain the volume class with
// g_volume_get_identifier().
//
// Known volume classes include device, network, and loop. Other classes may be
// added in the future.
//
// This is intended to be used by applications to classify #GVolume instances
// into different sections - for example a file manager or file chooser can use
// this information to show network volumes under a "Network" heading and device
// volumes under a "Devices" heading.
const VOLUME_IDENTIFIER_KIND_CLASS = "class"

// VOLUME_IDENTIFIER_KIND_HAL_UDI: string used to obtain a Hal UDI with
// g_volume_get_identifier().
//
// Deprecated: Do not use, HAL is deprecated.
const VOLUME_IDENTIFIER_KIND_HAL_UDI = "hal-udi"

// VOLUME_IDENTIFIER_KIND_LABEL: string used to obtain a filesystem label with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_LABEL = "label"

// VOLUME_IDENTIFIER_KIND_NFS_MOUNT: string used to obtain a NFS mount with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT = "nfs-mount"

// VOLUME_IDENTIFIER_KIND_UNIX_DEVICE: string used to obtain a Unix device path
// with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE = "unix-device"

// VOLUME_IDENTIFIER_KIND_UUID: string used to obtain a UUID with
// g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UUID = "uuid"

// VolumeOverrider contains methods that are overridable.
type VolumeOverrider interface {
}

// Volume interface represents user-visible objects that can be mounted. Note,
// when porting from GnomeVFS, #GVolume is the moral equivalent of VFSDrive.
//
// Mounting a #GVolume instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To mount a
// #GVolume, first call g_volume_mount() with (at least) the #GVolume instance,
// optionally a Operation object and a ReadyCallback.
//
// Typically, one will only want to pass NULL for the Operation if automounting
// all volumes when a desktop session starts since it's not desirable to put up
// a lot of dialogs asking for credentials.
//
// The callback will be fired when the operation has resolved (either with
// success or failure), and a Result instance will be passed to the callback.
// That callback should then call g_volume_mount_finish() with the #GVolume
// instance and the Result data to see if the operation was completed
// successfully. If an error is present when g_volume_mount_finish() is called,
// then it will be filled with any error information.
//
//
// Volume Identifiers
//
// It is sometimes necessary to directly access the underlying operating system
// object behind a volume (e.g. for passing a volume to an application via the
// commandline). For this purpose, GIO allows to obtain an 'identifier' for the
// volume. There can be different kinds of identifiers, such as Hal UDIs,
// filesystem labels, traditional Unix devices (e.g. /dev/sda2), UUIDs. GIO uses
// predefined strings as names for the different kinds of identifiers:
// VOLUME_IDENTIFIER_KIND_UUID, VOLUME_IDENTIFIER_KIND_LABEL, etc. Use
// g_volume_get_identifier() to obtain an identifier for a volume.
//
//    Note that VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available when the gvfs hal volume monitor is in use. Other volume monitors will generally be able to provide the VOLUME_IDENTIFIER_KIND_UNIX_DEVICE identifier, which can be used to obtain a hal device by means of libhal_manager_find_device_string_match().
//
// Volume wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Volume struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Volume)(nil)
)

// Volumer describes Volume's interface methods.
type Volumer interface {
	coreglib.Objector

	baseVolume() *Volume
}

var _ Volumer = (*Volume)(nil)

func ifaceInitVolumer(gifacePtr, data C.gpointer) {
}

func wrapVolume(obj *coreglib.Object) *Volume {
	return &Volume{
		Object: obj,
	}
}

func marshalVolume(p uintptr) (interface{}, error) {
	return wrapVolume(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Volume) baseVolume() *Volume {
	return v
}

// BaseVolume returns the underlying base object.
func BaseVolume(obj Volumer) *Volume {
	return obj.baseVolume()
}

// ConnectChanged is emitted when the volume has been changed.
func (v *Volume) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gio2_Volume_ConnectChanged), f)
}

// ConnectRemoved: this signal is emitted when the #GVolume have been removed.
// If the recipient is holding references to the object they should release them
// so the object can be finalized.
func (v *Volume) ConnectRemoved(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "removed", false, unsafe.Pointer(C._gotk4_gio2_Volume_ConnectRemoved), f)
}

// VolumeIface: interface for implementing operations for mountable volumes.
//
// An instance of this type is always passed by reference.
type VolumeIface struct {
	*volumeIface
}

// volumeIface is the struct that's finalized.
type volumeIface struct {
	native unsafe.Pointer
}

var GIRInfoVolumeIface = girepository.MustFind("Gio", "VolumeIface")
