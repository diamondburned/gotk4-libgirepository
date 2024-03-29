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
// extern void _gotk4_gio2_Mount_ConnectUnmounted(gpointer, guintptr);
// extern void _gotk4_gio2_Mount_ConnectPreUnmount(gpointer, guintptr);
// extern void _gotk4_gio2_Mount_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeMount = coreglib.Type(girepository.MustFind("Gio", "Mount").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMount, F: marshalMount},
	})
}

// MountOverrider contains methods that are overridable.
type MountOverrider interface {
}

// Mount interface represents user-visible mounts. Note, when porting from
// GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
//
// Mount wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Mount struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Mount)(nil)
)

// Mounter describes Mount's interface methods.
type Mounter interface {
	coreglib.Objector

	baseMount() *Mount
}

var _ Mounter = (*Mount)(nil)

func ifaceInitMounter(gifacePtr, data C.gpointer) {
}

func wrapMount(obj *coreglib.Object) *Mount {
	return &Mount{
		Object: obj,
	}
}

func marshalMount(p uintptr) (interface{}, error) {
	return wrapMount(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Mount) baseMount() *Mount {
	return v
}

// BaseMount returns the underlying base object.
func BaseMount(obj Mounter) *Mount {
	return obj.baseMount()
}

// ConnectChanged is emitted when the mount has been changed.
func (v *Mount) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectChanged), f)
}

// ConnectPreUnmount: this signal may be emitted when the #GMount is about to be
// unmounted.
//
// This signal depends on the backend and is only emitted if GIO was used to
// unmount.
func (v *Mount) ConnectPreUnmount(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "pre-unmount", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectPreUnmount), f)
}

// ConnectUnmounted: this signal is emitted when the #GMount have been
// unmounted. If the recipient is holding references to the object they should
// release them so the object can be finalized.
func (v *Mount) ConnectUnmounted(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "unmounted", false, unsafe.Pointer(C._gotk4_gio2_Mount_ConnectUnmounted), f)
}

// MountIface: interface for implementing operations for mounts.
//
// An instance of this type is always passed by reference.
type MountIface struct {
	*mountIface
}

// mountIface is the struct that's finalized.
type mountIface struct {
	native unsafe.Pointer
}

var GIRInfoMountIface = girepository.MustFind("Gio", "MountIface")
