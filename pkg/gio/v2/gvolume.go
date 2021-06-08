// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.g_volume_get_type()), F: marshalVolume},
	})
}

// VolumeOverrider contains methods that are overridable. This
// interface is a subset of the interface Volume.
type VolumeOverrider interface {
	// CanEject checks if a volume can be ejected.
	CanEject() bool
	// CanMount checks if a volume can be mounted.
	CanMount() bool

	Changed()
	// Eject ejects a volume. This is an asynchronous operation, and is finished
	// by calling g_volume_eject_finish() with the @volume and Result returned
	// in the @callback.
	Eject()
	// EjectFinish finishes ejecting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	EjectFinish(result AsyncResult) error
	// EjectWithOperation ejects a volume. This is an asynchronous operation,
	// and is finished by calling g_volume_eject_with_operation_finish() with
	// the @volume and Result data returned in the @callback.
	EjectWithOperation()
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
	// Mount gets the mount for the @volume.
	Mount() Mount
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
	// MountFinish finishes mounting a volume. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// If the mount operation succeeded, g_volume_get_mount() on @volume is
	// guaranteed to return the mount right after calling this function; there's
	// no need to listen for the 'mount-added' signal on Monitor.
	MountFinish(result AsyncResult) error
	// MountFn mounts a volume. This is an asynchronous operation, and is
	// finished by calling g_volume_mount_finish() with the @volume and Result
	// returned in the @callback.
	MountFn()

	Removed()
	// ShouldAutomount returns whether the volume should be automatically
	// mounted.
	ShouldAutomount() bool
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
	VolumeOverrider

	// GetMount gets the mount for the @volume.
	GetMount() Mount
}

// volume implements the Volume interface.
type volume struct {
	gextras.Objector
}

var _ Volume = (*volume)(nil)

// WrapVolume wraps a GObject to a type that implements interface
// Volume. It is primarily used internally.
func WrapVolume(obj *externglib.Object) Volume {
	return Volume{
		Objector: obj,
	}
}

func marshalVolume(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVolume(obj), nil
}

// CanEject checks if a volume can be ejected.
func (v volume) CanEject() bool {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_volume_can_eject(arg0)

	if cret {
		goret = true
	}

	return goret
}

// CanMount checks if a volume can be mounted.
func (v volume) CanMount() bool {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_volume_can_mount(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Eject ejects a volume. This is an asynchronous operation, and is finished
// by calling g_volume_eject_finish() with the @volume and Result returned
// in the @callback.
func (v volume) Eject() {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	C.g_volume_eject(arg0, arg1, arg2, arg3, arg4)
}

// EjectFinish finishes ejecting a volume. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
func (v volume) EjectFinish(result AsyncResult) error {
	var arg0 *C.GVolume
	var arg1 *C.GAsyncResult

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.g_volume_eject_finish(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// EjectWithOperation ejects a volume. This is an asynchronous operation,
// and is finished by calling g_volume_eject_with_operation_finish() with
// the @volume and Result data returned in the @callback.
func (v volume) EjectWithOperation() {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	C.g_volume_eject_with_operation(arg0, arg1, arg2, arg3, arg4, arg5)
}

// EjectWithOperationFinish finishes ejecting a volume. If any errors
// occurred during the operation, @error will be set to contain the errors
// and false will be returned.
func (v volume) EjectWithOperationFinish(result AsyncResult) error {
	var arg0 *C.GVolume
	var arg1 *C.GAsyncResult

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.g_volume_eject_with_operation_finish(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
// that @volume has. Use g_volume_get_identifier() to obtain the identifiers
// themselves.
func (v volume) EnumerateIdentifiers() []string {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var cret **C.char
	var goret []string

	cret = C.g_volume_enumerate_identifiers(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return goret
}

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
func (v volume) ActivationRoot() File {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.GFile)
	var goret File

	cret = C.g_volume_get_activation_root(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(File)

	return goret
}

// Drive gets the drive for the @volume.
func (v volume) Drive() Drive {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.GDrive)
	var goret Drive

	cret = C.g_volume_get_drive(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Drive)

	return goret
}

// Icon gets the icon for @volume.
func (v volume) Icon() Icon {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.GIcon)
	var goret Icon

	cret = C.g_volume_get_icon(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Icon)

	return goret
}

// Identifier gets the identifier of the given kind for @volume. See the
// [introduction][volume-identifier] for more information about volume
// identifiers.
func (v volume) Identifier(kind string) string {
	var arg0 *C.GVolume
	var arg1 *C.char

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.char)
	var goret string

	cret = C.g_volume_get_identifier(arg0, arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// GetMount gets the mount for the @volume.
func (v volume) GetMount() Mount {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.GMount)
	var goret Mount

	cret = C.g_volume_get_mount(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Mount)

	return goret
}

// Name gets the name of @volume.
func (v volume) Name() string {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.char)
	var goret string

	cret = C.g_volume_get_name(arg0)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// SortKey gets the sort key for @volume, if any.
func (v volume) SortKey() string {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var cret *C.gchar
	var goret string

	cret = C.g_volume_get_sort_key(arg0)

	goret = C.GoString(cret)

	return goret
}

// SymbolicIcon gets the symbolic icon for @volume.
func (v volume) SymbolicIcon() Icon {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.GIcon)
	var goret Icon

	cret = C.g_volume_get_symbolic_icon(arg0)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Icon)

	return goret
}

// UUID gets the UUID for the @volume. The reference is typically based on
// the file system UUID for the volume in question and should be considered
// an opaque string. Returns nil if there is no UUID available.
func (v volume) UUID() string {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	cret := new(C.char)
	var goret string

	cret = C.g_volume_get_uuid(arg0)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// Mount mounts a volume. This is an asynchronous operation, and is finished
// by calling g_volume_mount_finish() with the @volume and Result returned
// in the @callback.
func (v volume) Mount() {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	C.g_volume_mount(arg0, arg1, arg2, arg3, arg4, arg5)
}

// MountFinish finishes mounting a volume. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
//
// If the mount operation succeeded, g_volume_get_mount() on @volume is
// guaranteed to return the mount right after calling this function; there's
// no need to listen for the 'mount-added' signal on Monitor.
func (v volume) MountFinish(result AsyncResult) error {
	var arg0 *C.GVolume
	var arg1 *C.GAsyncResult

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.g_volume_mount_finish(arg0, arg1, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// ShouldAutomount returns whether the volume should be automatically
// mounted.
func (v volume) ShouldAutomount() bool {
	var arg0 *C.GVolume

	arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.g_volume_should_automount(arg0)

	if cret {
		goret = true
	}

	return goret
}
