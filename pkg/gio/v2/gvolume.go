// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
	// that @volume has. Use g_volume_get_identifier() to obtain the identifiers
	// themselves.
	EnumerateIdentifiers() []string
	// Identifier gets the identifier of the given kind for @volume. See the
	// [introduction][volume-identifier] for more information about volume
	// identifiers.
	Identifier(kind string) string
	// Name gets the name of @volume.
	Name() string
	// SortKey gets the sort key for @volume, if any.
	SortKey() string
	// UUID gets the UUID for the @volume. The reference is typically based on
	// the file system UUID for the volume in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID() string

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
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret C.gboolean // in

	_cret = C.g_volume_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanMount checks if a volume can be mounted.
func (v volume) CanMount() bool {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret C.gboolean // in

	_cret = C.g_volume_can_mount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EnumerateIdentifiers gets the kinds of [identifiers][volume-identifier]
// that @volume has. Use g_volume_get_identifier() to obtain the identifiers
// themselves.
func (v volume) EnumerateIdentifiers() []string {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret **C.char

	_cret = C.g_volume_enumerate_identifiers(_arg0)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != nil; p = (**C.char)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		src := unsafe.Slice(_cret, length)
		_utf8s = make([]string, length)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Identifier gets the identifier of the given kind for @volume. See the
// [introduction][volume-identifier] for more information about volume
// identifiers.
func (v volume) Identifier(kind string) string {
	var _arg0 *C.GVolume // out
	var _arg1 *C.char    // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_volume_get_identifier(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Name gets the name of @volume.
func (v volume) Name() string {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret *C.char // in

	_cret = C.g_volume_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SortKey gets the sort key for @volume, if any.
func (v volume) SortKey() string {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret *C.gchar // in

	_cret = C.g_volume_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UUID gets the UUID for the @volume. The reference is typically based on
// the file system UUID for the volume in question and should be considered
// an opaque string. Returns nil if there is no UUID available.
func (v volume) UUID() string {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret *C.char // in

	_cret = C.g_volume_get_uuid(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShouldAutomount returns whether the volume should be automatically
// mounted.
func (v volume) ShouldAutomount() bool {
	var _arg0 *C.GVolume // out

	_arg0 = (*C.GVolume)(unsafe.Pointer(v.Native()))

	var _cret C.gboolean // in

	_cret = C.g_volume_should_automount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
