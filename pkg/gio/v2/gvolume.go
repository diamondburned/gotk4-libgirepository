// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_volume_get_type()), F: marshalVolumer},
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
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type VolumeOverrider interface {
	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	Changed()
	// Eject ejects a volume. This is an asynchronous operation, and is finished
	// by calling g_volume_eject_finish() with the volume and Result returned in
	// the callback.
	//
	// Deprecated: Use g_volume_eject_with_operation() instead.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a volume. If any errors occurred during the
	// operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// Deprecated: Use g_volume_eject_with_operation_finish() instead.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a volume. This is an asynchronous operation,
	// and is finished by calling g_volume_eject_with_operation_finish() with
	// the volume and Result data returned in the callback.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a volume. If any errors
	// occurred during the operation, error will be set to contain the errors
	// and FALSE will be returned.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that volume has. Use g_volume_get_identifier() to obtain the identifiers
	// themselves.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time. Returns NULL otherwise. If not NULL and if volume is
	// mounted, then the result of g_mount_get_root() on the #GMount object
	// obtained from g_volume_get_mount() will always either be equal or a
	// prefix of what this function returns. In other words, in code
	//
	//    (g_file_has_prefix (volume_activation_root, mount_root) ||
	//     g_file_equal (volume_activation_root, mount_root))
	//
	// will always be TRUE.
	//
	// Activation roots are typically used in Monitor implementations to find
	// the underlying mount to shadow, see g_mount_is_shadowed() for more
	// details.
	ActivationRoot() Filer
	// Drive gets the drive for the volume.
	Drive() Driver
	// Icon gets the icon for volume.
	Icon() Iconner
	// Identifier gets the identifier of the given kind for volume. See the
	// [introduction][volume-identifier] for more information about volume
	// identifiers.
	Identifier(kind string) string
	// Mount gets the mount for the volume.
	Mount() Mounter
	// Name gets the name of volume.
	Name() string
	// SortKey gets the sort key for volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for volume.
	SymbolicIcon() Iconner
	// UUID gets the UUID for the volume. The reference is typically based on
	// the file system UUID for the volume in question and should be considered
	// an opaque string. Returns NULL if there is no UUID available.
	UUID() string
	// MountFinish finishes mounting a volume. If any errors occurred during the
	// operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// If the mount operation succeeded, g_volume_get_mount() on volume is
	// guaranteed to return the mount right after calling this function; there's
	// no need to listen for the 'mount-added' signal on Monitor.
	MountFinish(result AsyncResulter) error
	// MountFn mounts a volume. This is an asynchronous operation, and is
	// finished by calling g_volume_mount_finish() with the volume and Result
	// returned in the callback.
	MountFn(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	Removed()
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool
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
type Volume struct {
	*externglib.Object
}

// Volumer describes Volume's abstract methods.
type Volumer interface {
	externglib.Objector

	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	// Eject ejects a volume.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a volume.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a volume.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a volume.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that volume has.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time.
	ActivationRoot() Filer
	// Drive gets the drive for the volume.
	Drive() Driver
	// Icon gets the icon for volume.
	Icon() Iconner
	// Identifier gets the identifier of the given kind for volume.
	Identifier(kind string) string
	// GetMount gets the mount for the volume.
	GetMount() Mounter
	// Name gets the name of volume.
	Name() string
	// SortKey gets the sort key for volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for volume.
	SymbolicIcon() Iconner
	// UUID gets the UUID for the volume.
	UUID() string
	// Mount mounts a volume.
	Mount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// MountFinish finishes mounting a volume.
	MountFinish(result AsyncResulter) error
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool
}

var _ Volumer = (*Volume)(nil)

func wrapVolume(obj *externglib.Object) *Volume {
	return &Volume{
		Object: obj,
	}
}

func marshalVolumer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVolume(obj), nil
}

// CanEject checks if a volume can be ejected.
func (volume *Volume) CanEject() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_can_eject(_arg0)
	runtime.KeepAlive(volume)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanMount checks if a volume can be mounted.
func (volume *Volume) CanMount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_can_mount(_arg0)
	runtime.KeepAlive(volume)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Eject ejects a volume. This is an asynchronous operation, and is finished by
// calling g_volume_eject_finish() with the volume and Result returned in the
// callback.
//
// Deprecated: Use g_volume_eject_with_operation() instead.
func (volume *Volume) Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_volume_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// EjectFinish finishes ejecting a volume. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_volume_eject_with_operation_finish() instead.
func (volume *Volume) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_eject_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EjectWithOperation ejects a volume. This is an asynchronous operation, and is
// finished by calling g_volume_eject_with_operation_finish() with the volume
// and Result data returned in the callback.
func (volume *Volume) EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_volume_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// EjectWithOperationFinish finishes ejecting a volume. If any errors occurred
// during the operation, error will be set to contain the errors and FALSE will
// be returned.
func (volume *Volume) EjectWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_eject_with_operation_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier] that
// volume has. Use g_volume_get_identifier() to obtain the identifiers
// themselves.
func (volume *Volume) EnumerateIdentifiers() []string {
	var _arg0 *C.GVolume // out
	var _cret **C.char   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_enumerate_identifiers(_arg0)
	runtime.KeepAlive(volume)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ActivationRoot gets the activation root for a #GVolume if it is known ahead
// of mount time. Returns NULL otherwise. If not NULL and if volume is mounted,
// then the result of g_mount_get_root() on the #GMount object obtained from
// g_volume_get_mount() will always either be equal or a prefix of what this
// function returns. In other words, in code
//
//    (g_file_has_prefix (volume_activation_root, mount_root) ||
//     g_file_equal (volume_activation_root, mount_root))
//
// will always be TRUE.
//
// Activation roots are typically used in Monitor implementations to find the
// underlying mount to shadow, see g_mount_is_shadowed() for more details.
func (volume *Volume) ActivationRoot() Filer {
	var _arg0 *C.GVolume // out
	var _cret *C.GFile   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_activation_root(_arg0)
	runtime.KeepAlive(volume)

	var _file Filer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Filer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.Filer")
			}
			_file = rv
		}
	}

	return _file
}

// Drive gets the drive for the volume.
func (volume *Volume) Drive() Driver {
	var _arg0 *C.GVolume // out
	var _cret *C.GDrive  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_drive(_arg0)
	runtime.KeepAlive(volume)

	var _drive Driver // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Driver)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.Driver")
			}
			_drive = rv
		}
	}

	return _drive
}

// Icon gets the icon for volume.
func (volume *Volume) Icon() Iconner {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_icon(_arg0)
	runtime.KeepAlive(volume)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// Identifier gets the identifier of the given kind for volume. See the
// [introduction][volume-identifier] for more information about volume
// identifiers.
func (volume *Volume) Identifier(kind string) string {
	var _arg0 *C.GVolume // out
	var _arg1 *C.char    // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(kind)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_get_identifier(_arg0, _arg1)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(kind)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// GetMount gets the mount for the volume.
func (volume *Volume) GetMount() Mounter {
	var _arg0 *C.GVolume // out
	var _cret *C.GMount  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_mount(_arg0)
	runtime.KeepAlive(volume)

	var _mount Mounter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Mounter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.Mounter")
			}
			_mount = rv
		}
	}

	return _mount
}

// Name gets the name of volume.
func (volume *Volume) Name() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_name(_arg0)
	runtime.KeepAlive(volume)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for volume, if any.
func (volume *Volume) SortKey() string {
	var _arg0 *C.GVolume // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_sort_key(_arg0)
	runtime.KeepAlive(volume)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SymbolicIcon gets the symbolic icon for volume.
func (volume *Volume) SymbolicIcon() Iconner {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_symbolic_icon(_arg0)
	runtime.KeepAlive(volume)

	var _icon Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Iconner)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// UUID gets the UUID for the volume. The reference is typically based on the
// file system UUID for the volume in question and should be considered an
// opaque string. Returns NULL if there is no UUID available.
func (volume *Volume) UUID() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_get_uuid(_arg0)
	runtime.KeepAlive(volume)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Mount mounts a volume. This is an asynchronous operation, and is finished by
// calling g_volume_mount_finish() with the volume and Result returned in the
// callback.
func (volume *Volume) Mount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountMountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_volume_mount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// MountFinish finishes mounting a volume. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// If the mount operation succeeded, g_volume_get_mount() on volume is
// guaranteed to return the mount right after calling this function; there's no
// need to listen for the 'mount-added' signal on Monitor.
func (volume *Volume) MountFinish(result AsyncResulter) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_mount_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ShouldAutomount returns whether the volume should be automatically mounted.
func (volume *Volume) ShouldAutomount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(volume.Native()))

	_cret = C.g_volume_should_automount(_arg0)
	runtime.KeepAlive(volume)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
