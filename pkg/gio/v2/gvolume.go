// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_volume_get_type()), F: marshalVolume},
	})
}

// Volume: the #GVolume interface represents user-visible objects that can be
// mounted. Note, when porting from GnomeVFS, #GVolume is the moral equivalent
// of VFSDrive.
//
// Mounting a #GVolume instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To mount a
// #GVolume, first call g_volume_mount() with (at least) the #GVolume instance,
// optionally a Operation object and a ReadyCallback.
//
// Typically, one will only want to pass nil for the Operation if automounting
// all volumes when a desktop session starts since it's not desirable to put up
// a lot of dialogs asking for credentials.
//
// The callback will be fired when the operation has resolved (either with
// success or failure), and a Result instance will be passed to the callback.
// That callback should then call g_volume_mount_finish() with the #GVolume
// instance and the Result data to see if the operation was completed
// successfully. If an @error is present when g_volume_mount_finish() is called,
// then it will be filled with any error information.
//
//
// Volume Identifiers
//
// It is sometimes necessary to directly access the underlying operating system
// object behind a volume (e.g. for passing a volume to an application via the
// commandline). For this purpose, GIO allows to obtain an 'identifier' for the
// volume. There can be different kinds of identifiers, such as Hal UDIs,
// filesystem labels, traditional Unix devices (e.g. `/dev/sda2`), UUIDs. GIO
// uses predefined strings as names for the different kinds of identifiers:
// VOLUME_IDENTIFIER_KIND_UUID, VOLUME_IDENTIFIER_KIND_LABEL, etc. Use
// g_volume_get_identifier() to obtain an identifier for a volume.
//
//    Note that VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available when the gvfs hal volume monitor is in use. Other volume monitors will generally be able to provide the VOLUME_IDENTIFIER_KIND_UNIX_DEVICE identifier, which can be used to obtain a hal device by means of libhal_manager_find_device_string_match().
type Volume interface {
	gextras.Objector

	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool
	// Eject ejects a volume. This is an asynchronous operation, and is finished
	// by calling g_volume_eject_finish() with the @volume and Result returned
	// in the @callback.
	//
	// Deprecated: since version 2.22.
	Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: since version 2.22.
	EjectFinish(result AsyncResult) error
	// EjectWithOperation ejects a volume. This is an asynchronous operation,
	// and is finished by calling g_volume_eject_with_operation_finish() with
	// the @volume and Result data returned in the @callback.
	EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a volume. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(result AsyncResult) error
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that @volume has. Use g_volume_get_identifier() to obtain the identifiers
	// themselves.
	EnumerateIdentifiers() []string
	// ActivationRoot gets the activation root for a #GVolume if it is known
	// ahead of mount time. Returns nil otherwise. If not nil and if @volume is
	// mounted, then the result of g_mount_get_root() on the #GMount object
	// obtained from g_volume_get_mount() will always either be equal or a
	// prefix of what this function returns. In other words, in code
	//
	//    (g_file_has_prefix (volume_activation_root, mount_root) ||
	//     g_file_equal (volume_activation_root, mount_root))
	//
	// will always be true.
	//
	// Activation roots are typically used in Monitor implementations to find
	// the underlying mount to shadow, see g_mount_is_shadowed() for more
	// details.
	ActivationRoot() File
	// Drive gets the drive for the @volume.
	Drive() Drive
	// Icon gets the icon for @volume.
	Icon() Icon
	// Identifier gets the identifier of the given kind for @volume. See the
	// [introduction][volume-identifier] for more information about volume
	// identifiers.
	Identifier(kind string) string
	// GetMount gets the mount for the @volume.
	GetMount() Mount
	// Name gets the name of @volume.
	Name() string
	// SortKey gets the sort key for @volume, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for @volume.
	SymbolicIcon() Icon
	// UUID gets the UUID for the @volume. The reference is typically based on
	// the file system UUID for the volume in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID() string
	// Mount mounts a volume. This is an asynchronous operation, and is finished
	// by calling g_volume_mount_finish() with the @volume and Result returned
	// in the @callback.
	Mount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// MountFinish finishes mounting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// If the mount operation succeeded, g_volume_get_mount() on @volume is
	// guaranteed to return the mount right after calling this function; there's
	// no need to listen for the 'mount-added' signal on Monitor.
	MountFinish(result AsyncResult) error
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool
}

// volume implements the Volume interface.
type volume struct {
	*externglib.Object
}

var _ Volume = (*volume)(nil)

// WrapVolume wraps a GObject to a type that implements
// interface Volume. It is primarily used internally.
func WrapVolume(obj *externglib.Object) Volume {
	return volume{obj}
}

func marshalVolume(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVolume(obj), nil
}

func (v volume) CanEject() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (v volume) CanMount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_can_mount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (v volume) Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_volume_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (v volume) EjectFinish(result AsyncResult) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_eject_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (v volume) EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_volume_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (v volume) EjectWithOperationFinish(result AsyncResult) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_eject_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (v volume) EnumerateIdentifiers() []string {
	var _arg0 *C.GVolume // out
	var _cret **C.char

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_enumerate_identifiers(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (v volume) ActivationRoot() File {
	var _arg0 *C.GVolume // out
	var _cret *C.GFile   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_activation_root(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

func (v volume) Drive() Drive {
	var _arg0 *C.GVolume // out
	var _cret *C.GDrive  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_drive(_arg0)

	var _drive Drive // out

	_drive = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Drive)

	return _drive
}

func (v volume) Icon() Icon {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (v volume) Identifier(kind string) string {
	var _arg0 *C.GVolume // out
	var _arg1 *C.char    // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_get_identifier(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (v volume) GetMount() Mount {
	var _arg0 *C.GVolume // out
	var _cret *C.GMount  // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_mount(_arg0)

	var _mount Mount // out

	_mount = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Mount)

	return _mount
}

func (v volume) Name() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (v volume) SortKey() string {
	var _arg0 *C.GVolume // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (v volume) SymbolicIcon() Icon {
	var _arg0 *C.GVolume // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (v volume) UUID() string {
	var _arg0 *C.GVolume // out
	var _cret *C.char    // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_get_uuid(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (v volume) Mount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GVolume            // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = C.GMountMountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_volume_mount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (v volume) MountFinish(result AsyncResult) error {
	var _arg0 *C.GVolume      // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_volume_mount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (v volume) ShouldAutomount() bool {
	var _arg0 *C.GVolume // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	_cret = C.g_volume_should_automount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
