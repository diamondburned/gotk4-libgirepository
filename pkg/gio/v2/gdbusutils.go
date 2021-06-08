// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
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

// DBusEscapeObjectPath: this is a language binding friendly version of
// g_dbus_escape_object_path_bytestring().
func DBusEscapeObjectPath(s string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.gchar)
	var goret string

	cret = C.g_dbus_escape_object_path(arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// DBusEscapeObjectPathBytestring escapes @bytes for use in a D-Bus object path
// component. @bytes is an array of zero or more nonzero bytes in an unspecified
// encoding, followed by a single zero byte.
//
// The escaping method consists of replacing all non-alphanumeric characters
// (see g_ascii_isalnum()) with their hexadecimal value preceded by an
// underscore (`_`). For example: `foo.bar.baz` will become `foo_2ebar_2ebaz`.
//
// This method is appropriate to use when the input is nearly a valid object
// path component but is not when your input is far from being a valid object
// path component. Other escaping algorithms are also valid to use with D-Bus
// object paths.
//
// This can be reversed with g_dbus_unescape_object_path().
func DBusEscapeObjectPathBytestring(bytes []byte) string {
	var arg1 *C.guint8

	arg1 = (*C.guint8)(C.malloc((len(bytes) + 1) * C.sizeof_guint8))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []C.guint8
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(bytes)))

		for i := range bytes {
			out[i] = C.guint8(bytes[i])
		}
	}

	cret := new(C.gchar)
	var goret string

	cret = C.g_dbus_escape_object_path_bytestring(arg1)

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
}

// DBusGenerateGuid: generate a D-Bus GUID that can be used with e.g.
// g_dbus_connection_new().
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusGenerateGuid() string {
	cret := new(C.gchar)
	var goret string

	cret = C.g_dbus_generate_guid()

	goret = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret
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
func DBusGValueToGVariant(gValue *externglib.Value, typ *glib.VariantType) *glib.Variant {
	var arg1 *C.GValue
	var arg2 *C.GVariantType

	arg1 = (*C.GValue)(gValue.GValue)
	arg2 = (*C.GVariantType)(unsafe.Pointer(typ.Native()))

	cret := new(C.GVariant)
	var goret *glib.Variant

	cret = C.g_dbus_gvalue_to_gvariant(arg1, arg2)

	goret = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
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
	var arg1 *C.GVariant

	arg1 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	arg2 := new(C.GValue)
	var ret2 *externglib.Value

	C.g_dbus_gvariant_to_gvalue(arg1, arg2)

	ret2 = externglib.ValueFromNative(unsafe.Pointer(*arg2))

	return ret2
}

// DBusIsGuid checks if @string is a D-Bus GUID.
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusIsGuid(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_dbus_is_guid(arg1)

	if cret {
		goret = true
	}

	return goret
}

// DBusIsInterfaceName checks if @string is a valid D-Bus interface name.
func DBusIsInterfaceName(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_dbus_is_interface_name(arg1)

	if cret {
		goret = true
	}

	return goret
}

// DBusIsMemberName checks if @string is a valid D-Bus member (e.g. signal or
// method) name.
func DBusIsMemberName(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_dbus_is_member_name(arg1)

	if cret {
		goret = true
	}

	return goret
}

// DBusIsName checks if @string is a valid D-Bus bus name (either unique or
// well-known).
func DBusIsName(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_dbus_is_name(arg1)

	if cret {
		goret = true
	}

	return goret
}

// DBusIsUniqueName checks if @string is a valid D-Bus unique bus name.
func DBusIsUniqueName(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret bool

	cret = C.g_dbus_is_unique_name(arg1)

	if cret {
		goret = true
	}

	return goret
}

// DBusUnescapeObjectPath unescapes an string that was previously escaped with
// g_dbus_escape_object_path(). If the string is in a format that could not have
// been returned by g_dbus_escape_object_path(), this function returns nil.
//
// Encoding alphanumeric characters which do not need to be encoded is not
// allowed (e.g `_63` is not valid, the string should contain `c` instead).
func DBusUnescapeObjectPath(s string) []byte {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.guint8
	var goret []byte

	cret = C.g_dbus_unescape_object_path(arg1)

	{
		var length int
		for p := cret; *p != 0; p = (*C.guint8)(ptr.Add(unsafe.Pointer(p), C.sizeof_guint8)) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret = make([]byte, length)
		for i := uintptr(0); i < uintptr(length); i += C.sizeof_guint8 {
			src := (C.guint8)(ptr.Add(unsafe.Pointer(cret), i))
			goret[i] = byte(src)
		}
	}

	return goret
}
