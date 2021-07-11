// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
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
		{T: externglib.Type(C.g_drive_get_type()), F: marshalDriver},
	})
}

// DriveOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DriveOverrider interface {
	// CanEject checks if a drive can be ejected.
	CanEject() bool
	// CanPollForMedia checks if a drive can be polled for media changes.
	CanPollForMedia() bool
	// CanStart checks if a drive can be started.
	CanStart() bool
	// CanStartDegraded checks if a drive can be started degraded.
	CanStartDegraded() bool
	// CanStop checks if a drive can be stopped.
	CanStop() bool
	//
	Changed()
	//
	Disconnected()
	//
	EjectButton()
	// EjectFinish finishes ejecting a drive.
	//
	// Deprecated: Use g_drive_eject_with_operation_finish() instead.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a drive. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of identifiers that @drive has. Use
	// g_drive_get_identifier() to obtain the identifiers themselves.
	EnumerateIdentifiers() []string
	// Icon gets the icon for @drive.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for @drive. The only
	// identifier currently available is DRIVE_IDENTIFIER_KIND_UNIX_DEVICE.
	Identifier(kind string) string
	// Name gets the name of @drive.
	Name() string
	// SortKey gets the sort key for @drive, if any.
	SortKey() string
	// StartStopType gets a hint about how a drive can be started/stopped.
	StartStopType() DriveStartStopType
	// SymbolicIcon gets the icon for @drive.
	SymbolicIcon() *Icon
	// HasMedia checks if the @drive has media. Note that the OS may not be
	// polling the drive for media changes; see
	// g_drive_is_media_check_automatic() for more details.
	HasMedia() bool
	// HasVolumes: check if @drive has any mountable volumes.
	HasVolumes() bool
	// IsMediaCheckAutomatic checks if @drive is capable of automatically
	// detecting media changes.
	IsMediaCheckAutomatic() bool
	// IsMediaRemovable checks if the @drive supports removable media.
	IsMediaRemovable() bool
	// IsRemovable checks if the #GDrive and/or its media is considered
	// removable by the user. See g_drive_is_media_removable().
	IsRemovable() bool
	// PollForMedia: asynchronously polls @drive to see if media has been
	// inserted or removed.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_drive_poll_for_media_finish() to obtain the result of the
	// operation.
	PollForMedia(cancellable Cancellabler, callback AsyncReadyCallback)
	// PollForMediaFinish finishes an operation started with
	// g_drive_poll_for_media() on a drive.
	PollForMediaFinish(result AsyncResulter) error
	// StartFinish finishes starting a drive.
	StartFinish(result AsyncResulter) error
	//
	StopButton()
	// StopFinish finishes stopping a drive.
	StopFinish(result AsyncResulter) error
}

// Driver describes Drive's methods.
type Driver interface {
	// CanEject checks if a drive can be ejected.
	CanEject() bool
	// CanPollForMedia checks if a drive can be polled for media changes.
	CanPollForMedia() bool
	// CanStart checks if a drive can be started.
	CanStart() bool
	// CanStartDegraded checks if a drive can be started degraded.
	CanStartDegraded() bool
	// CanStop checks if a drive can be stopped.
	CanStop() bool
	// EjectFinish finishes ejecting a drive.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperationFinish finishes ejecting a drive.
	EjectWithOperationFinish(result AsyncResulter) error
	// EnumerateIdentifiers gets the kinds of identifiers that @drive has.
	EnumerateIdentifiers() []string
	// Icon gets the icon for @drive.
	Icon() *Icon
	// Identifier gets the identifier of the given kind for @drive.
	Identifier(kind string) string
	// Name gets the name of @drive.
	Name() string
	// SortKey gets the sort key for @drive, if any.
	SortKey() string
	// StartStopType gets a hint about how a drive can be started/stopped.
	StartStopType() DriveStartStopType
	// SymbolicIcon gets the icon for @drive.
	SymbolicIcon() *Icon
	// HasMedia checks if the @drive has media.
	HasMedia() bool
	// HasVolumes: check if @drive has any mountable volumes.
	HasVolumes() bool
	// IsMediaCheckAutomatic checks if @drive is capable of automatically
	// detecting media changes.
	IsMediaCheckAutomatic() bool
	// IsMediaRemovable checks if the @drive supports removable media.
	IsMediaRemovable() bool
	// IsRemovable checks if the #GDrive and/or its media is considered
	// removable by the user.
	IsRemovable() bool
	// PollForMedia: asynchronously polls @drive to see if media has been
	// inserted or removed.
	PollForMedia(cancellable Cancellabler, callback AsyncReadyCallback)
	// PollForMediaFinish finishes an operation started with
	// g_drive_poll_for_media() on a drive.
	PollForMediaFinish(result AsyncResulter) error
	// StartFinish finishes starting a drive.
	StartFinish(result AsyncResulter) error
	// StopFinish finishes stopping a drive.
	StopFinish(result AsyncResulter) error
}

// Drive - this represent a piece of hardware connected to the machine. It's
// generally only created for removable hardware or hardware with removable
// media.
//
// #GDrive is a container class for #GVolume objects that stem from the same
// piece of media. As such, #GDrive abstracts a drive with (or without)
// removable media and provides operations for querying whether media is
// available, determining whether media change is automatically detected and
// ejecting the media.
//
// If the #GDrive reports that media isn't automatically detected, one can poll
// for media; typically one should not do this periodically as a poll for media
// operation is potentially expensive and may spin up the drive creating noise.
//
// #GDrive supports starting and stopping drives with authentication support for
// the former. This can be used to support a diverse set of use cases including
// connecting/disconnecting iSCSI devices, powering down external disk
// enclosures and starting/stopping multi-disk devices such as RAID devices.
// Note that the actual semantics and side-effects of starting/stopping a
// #GDrive may vary according to implementation. To choose the correct verbs in
// e.g. a file manager, use g_drive_get_start_stop_type().
//
// For porting from GnomeVFS note that there is no equivalent of #GDrive in that
// API.
type Drive struct {
	*externglib.Object
}

var (
	_ Driver          = (*Drive)(nil)
	_ gextras.Nativer = (*Drive)(nil)
)

func wrapDrive(obj *externglib.Object) Driver {
	return &Drive{
		Object: obj,
	}
}

func marshalDriver(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrive(obj), nil
}

// CanEject checks if a drive can be ejected.
func (drive *Drive) CanEject() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanPollForMedia checks if a drive can be polled for media changes.
func (drive *Drive) CanPollForMedia() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_can_poll_for_media(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStart checks if a drive can be started.
func (drive *Drive) CanStart() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_can_start(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStartDegraded checks if a drive can be started degraded.
func (drive *Drive) CanStartDegraded() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_can_start_degraded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanStop checks if a drive can be stopped.
func (drive *Drive) CanStop() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_can_stop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EjectFinish finishes ejecting a drive.
//
// Deprecated: Use g_drive_eject_with_operation_finish() instead.
func (drive *Drive) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_drive_eject_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// EjectWithOperationFinish finishes ejecting a drive. If any errors occurred
// during the operation, @error will be set to contain the errors and false will
// be returned.
func (drive *Drive) EjectWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_drive_eject_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// EnumerateIdentifiers gets the kinds of identifiers that @drive has. Use
// g_drive_get_identifier() to obtain the identifiers themselves.
func (drive *Drive) EnumerateIdentifiers() []string {
	var _arg0 *C.GDrive // out
	var _cret **C.char

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_enumerate_identifiers(_arg0)

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
			_utf8s[i] = C.GoString((*C.gchar)(src[i]))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Icon gets the icon for @drive.
func (drive *Drive) Icon() *Icon {
	var _arg0 *C.GDrive // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_get_icon(_arg0)

	var _icon *Icon // out

	_icon = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Icon)

	return _icon
}

// Identifier gets the identifier of the given kind for @drive. The only
// identifier currently available is DRIVE_IDENTIFIER_KIND_UNIX_DEVICE.
func (drive *Drive) Identifier(kind string) string {
	var _arg0 *C.GDrive // out
	var _arg1 *C.char   // out
	var _cret *C.char   // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_drive_get_identifier(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Name gets the name of @drive.
func (drive *Drive) Name() string {
	var _arg0 *C.GDrive // out
	var _cret *C.char   // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for @drive, if any.
func (drive *Drive) SortKey() string {
	var _arg0 *C.GDrive // out
	var _cret *C.gchar  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(_cret))

	return _utf8
}

// StartStopType gets a hint about how a drive can be started/stopped.
func (drive *Drive) StartStopType() DriveStartStopType {
	var _arg0 *C.GDrive             // out
	var _cret C.GDriveStartStopType // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_get_start_stop_type(_arg0)

	var _driveStartStopType DriveStartStopType // out

	_driveStartStopType = DriveStartStopType(_cret)

	return _driveStartStopType
}

// SymbolicIcon gets the icon for @drive.
func (drive *Drive) SymbolicIcon() *Icon {
	var _arg0 *C.GDrive // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_get_symbolic_icon(_arg0)

	var _icon *Icon // out

	_icon = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Icon)

	return _icon
}

// HasMedia checks if the @drive has media. Note that the OS may not be polling
// the drive for media changes; see g_drive_is_media_check_automatic() for more
// details.
func (drive *Drive) HasMedia() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_has_media(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasVolumes: check if @drive has any mountable volumes.
func (drive *Drive) HasVolumes() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_has_volumes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMediaCheckAutomatic checks if @drive is capable of automatically detecting
// media changes.
func (drive *Drive) IsMediaCheckAutomatic() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_is_media_check_automatic(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMediaRemovable checks if the @drive supports removable media.
func (drive *Drive) IsMediaRemovable() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_is_media_removable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRemovable checks if the #GDrive and/or its media is considered removable by
// the user. See g_drive_is_media_removable().
func (drive *Drive) IsRemovable() bool {
	var _arg0 *C.GDrive  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))

	_cret = C.g_drive_is_removable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PollForMedia: asynchronously polls @drive to see if media has been inserted
// or removed.
//
// When the operation is finished, @callback will be called. You can then call
// g_drive_poll_for_media_finish() to obtain the result of the operation.
func (drive *Drive) PollForMedia(cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GDrive             // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(gbox.Assign(callback))

	C.g_drive_poll_for_media(_arg0, _arg1, _arg2, _arg3)
}

// PollForMediaFinish finishes an operation started with
// g_drive_poll_for_media() on a drive.
func (drive *Drive) PollForMediaFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_drive_poll_for_media_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// StartFinish finishes starting a drive.
func (drive *Drive) StartFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_drive_start_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// StopFinish finishes stopping a drive.
func (drive *Drive) StopFinish(result AsyncResulter) error {
	var _arg0 *C.GDrive       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GDrive)(unsafe.Pointer(drive.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_drive_stop_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
