// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_icon_get_type()), F: marshalIconner},
	})
}

// IconOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type IconOverrider interface {
	// Equal checks if two icons are equal.
	Equal(icon2 Iconner) bool
	// Hash gets a hash for an icon.
	Hash() uint
	// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can
	// be retrieved back by calling g_icon_deserialize() on the returned value.
	// As serialization will avoid using raw icon data when possible, it only
	// makes sense to transfer the #GVariant between processes on the same
	// machine, (as opposed to over the network), and within the same file
	// system namespace.
	Serialize() *glib.Variant
}

// Icon is a very minimal interface for icons. It provides functions for
// checking the equality of two icons, hashing of icons and serializing an icon
// to and from strings.
//
// #GIcon does not provide the actual pixmap for the icon as this is out of
// GIO's scope, however implementations of #GIcon may contain the name of an
// icon (see Icon), or the path to an icon (see Icon).
//
// To obtain a hash of a #GIcon, see g_icon_hash().
//
// To check if two #GIcons are equal, see g_icon_equal().
//
// For serializing a #GIcon, use g_icon_serialize() and g_icon_deserialize().
//
// If you want to consume #GIcon (for example, in a toolkit) you must be
// prepared to handle at least the three following cases: Icon, Icon and Icon.
// It may also make sense to have fast-paths for other cases (like handling
// Pixbuf directly, for example) but all compliant #GIcon implementations
// outside of GIO must implement Icon.
//
// If your application or library provides one or more #GIcon implementations
// you need to ensure that your new implementation also implements Icon.
// Additionally, you must provide an implementation of g_icon_serialize() that
// gives a result that is understood by g_icon_deserialize(), yielding one of
// the built-in icon types.
type Icon struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*Icon)(nil)
)

// Iconner describes Icon's interface methods.
type Iconner interface {
	externglib.Objector

	// Equal checks if two icons are equal.
	Equal(icon2 Iconner) bool
	// Serialize serializes a #GIcon into a #GVariant.
	Serialize() *glib.Variant
	// String generates a textual representation of icon that can be used for
	// serialization such as when passing icon to a different process or saving
	// it to persistent storage.
	String() string
}

var _ Iconner = (*Icon)(nil)

func wrapIcon(obj *externglib.Object) *Icon {
	return &Icon{
		Object: obj,
	}
}

func marshalIconner(p uintptr) (interface{}, error) {
	return wrapIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Equal checks if two icons are equal.
//
// The function takes the following parameters:
//
//    - icon2: pointer to the second #GIcon.
//
func (icon1 *Icon) Equal(icon2 Iconner) bool {
	var _arg0 *C.GIcon   // out
	var _arg1 *C.GIcon   // out
	var _cret C.gboolean // in

	if icon1 != nil {
		_arg0 = (*C.GIcon)(unsafe.Pointer(icon1.Native()))
	}
	if icon2 != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(icon2.Native()))
	}

	_cret = C.g_icon_equal(_arg0, _arg1)
	runtime.KeepAlive(icon1)
	runtime.KeepAlive(icon2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Serialize serializes a #GIcon into a #GVariant. An equivalent #GIcon can be
// retrieved back by calling g_icon_deserialize() on the returned value. As
// serialization will avoid using raw icon data when possible, it only makes
// sense to transfer the #GVariant between processes on the same machine, (as
// opposed to over the network), and within the same file system namespace.
func (icon *Icon) Serialize() *glib.Variant {
	var _arg0 *C.GIcon    // out
	var _cret *C.GVariant // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_icon_serialize(_arg0)
	runtime.KeepAlive(icon)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// String generates a textual representation of icon that can be used for
// serialization such as when passing icon to a different process or saving it
// to persistent storage. Use g_icon_new_for_string() to get icon back from the
// returned string.
//
// The encoding of the returned string is proprietary to #GIcon except in the
// following two cases
//
// - If icon is a Icon, the returned string is a native path (such as
// /path/to/my icon.png) without escaping if the #GFile for icon is a native
// file. If the file is not native, the returned string is the result of
// g_file_get_uri() (such as sftp://path/to/my20icon.png).
//
// - If icon is a Icon with exactly one name and no fallbacks, the encoding is
// simply the name (such as network-server).
func (icon *Icon) String() string {
	var _arg0 *C.GIcon // out
	var _cret *C.gchar // in

	_arg0 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_icon_to_string(_arg0)
	runtime.KeepAlive(icon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// IconDeserialize deserializes a #GIcon previously serialized using
// g_icon_serialize().
//
// The function takes the following parameters:
//
//    - value created with g_icon_serialize().
//
func IconDeserialize(value *glib.Variant) Iconner {
	var _arg1 *C.GVariant // out
	var _cret *C.GIcon    // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.g_icon_deserialize(_arg1)
	runtime.KeepAlive(value)

	var _icon Iconner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Iconner)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.Iconner")
			}
			_icon = rv
		}
	}

	return _icon
}

// IconHash gets a hash for an icon.
//
// The function takes the following parameters:
//
//    - icon to an icon object.
//
func IconHash(icon cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(icon))

	_cret = C.g_icon_hash(_arg1)
	runtime.KeepAlive(icon)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NewIconForString: generate a #GIcon instance from str. This function can fail
// if str is not valid - see g_icon_to_string() for discussion.
//
// If your application or library provides one or more #GIcon implementations
// you need to ensure that each #GType is registered with the type system prior
// to calling g_icon_new_for_string().
//
// The function takes the following parameters:
//
//    - str: string obtained via g_icon_to_string().
//
func NewIconForString(str string) (Iconner, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.GIcon  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_icon_new_for_string(_arg1, &_cerr)
	runtime.KeepAlive(str)

	var _icon Iconner // out
	var _goerr error  // out

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
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _icon, _goerr
}
