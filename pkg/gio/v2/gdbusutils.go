// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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

// DBusGenerateGuid: generate a D-Bus GUID that can be used with e.g.
// g_dbus_connection_new().
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusGenerateGuid() string {
	var _cret *C.gchar

	cret = C.g_dbus_generate_guid()

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusGValueToGVariant converts a #GValue to a #GVariant of the type indicated
// by the @type parameter.
//
// The conversion is using the following rules:
//
// - TYPE_STRING: 's', 'o', 'g' or 'ay' - TYPE_STRV: 'as', 'ao' or 'aay' -
// TYPE_BOOLEAN: 'b' - TYPE_UCHAR: 'y' - TYPE_INT: 'i', 'n' - TYPE_UINT: 'u',
// 'q' - TYPE_INT64 'x' - TYPE_UINT64: 't' - TYPE_DOUBLE: 'd' - TYPE_VARIANT:
// Any Type
//
// This can fail if e.g. @gvalue is of type TYPE_STRING and @type is
// ['i'][G-VARIANT-TYPE-INT32:CAPS]. It will also fail for any #GType (including
// e.g. TYPE_OBJECT and TYPE_BOXED derived-types) not in the table above.
//
// Note that if @gvalue is of type TYPE_VARIANT and its value is nil, the empty
// #GVariant instance (never nil) for @type is returned (e.g. 0 for scalar
// types, the empty string for string types, '/' for object path types, the
// empty array for any array type and so on).
//
// See the g_dbus_gvariant_to_gvalue() function for how to convert a #GVariant
// to a #GValue.
func DBusGValueToGVariant(gvalue **externglib.Value, typ *glib.VariantType) *glib.Variant {
	var _arg1 *C.GValue
	var _arg2 *C.GVariantType

	_arg1 = (*C.GValue)(gvalue.GValue)
	_arg2 = (*C.GVariantType)(unsafe.Pointer(typ.Native()))

	var _cret *C.GVariant

	cret = C.g_dbus_gvalue_to_gvariant(_arg1, _arg2)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _variant
}

// DBusGVariantToGValue converts a #GVariant to a #GValue. If @value is
// floating, it is consumed.
//
// The rules specified in the g_dbus_gvalue_to_gvariant() function are used -
// this function is essentially its reverse form. So, a #GVariant containing any
// basic or string array type will be converted to a #GValue containing a basic
// value or string array. Any other #GVariant (handle, variant, tuple, dict
// entry) will be converted to a #GValue containing that #GVariant.
//
// The conversion never fails - a valid #GValue is always returned in
// @out_gvalue.
func DBusGVariantToGValue(value *glib.Variant) *externglib.Value {
	var _arg1 *C.GVariant

	_arg1 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	var _arg2 C.GValue

	C.g_dbus_gvariant_to_gvalue(_arg1, &_arg2)

	var _outGvalue *externglib.Value

	_outGvalue = externglib.ValueFromNative(unsafe.Pointer(_arg2))

	return _outGvalue
}

// DBusIsGuid checks if @string is a D-Bus GUID.
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusIsGuid(string string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_dbus_is_guid(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DBusIsInterfaceName checks if @string is a valid D-Bus interface name.
func DBusIsInterfaceName(string string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_dbus_is_interface_name(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DBusIsMemberName checks if @string is a valid D-Bus member (e.g. signal or
// method) name.
func DBusIsMemberName(string string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_dbus_is_member_name(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DBusIsName checks if @string is a valid D-Bus bus name (either unique or
// well-known).
func DBusIsName(string string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_dbus_is_name(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DBusIsUniqueName checks if @string is a valid D-Bus unique bus name.
func DBusIsUniqueName(string string) bool {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	cret = C.g_dbus_is_unique_name(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}
