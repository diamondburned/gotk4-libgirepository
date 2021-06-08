// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

// ParamSpecBoolean creates a new SpecBoolean instance specifying a
// G_TYPE_BOOLEAN property. In many cases, it may be more appropriate to use an
// enum with g_param_spec_enum(), both to improve code clarity by using
// explicitly named values, and to allow for more values to be added in future
// without breaking API.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecBoolean(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gboolean
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	if defaultValue {
		arg4 = C.gboolean(1)
	}
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_boolean(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecBoxed creates a new SpecBoxed instance specifying a G_TYPE_BOXED
// derived property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecBoxed(name string, nick string, blurb string, boxedType externglib.Type, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(boxedType)
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_boxed(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecChar creates a new SpecChar instance specifying a G_TYPE_CHAR
// property.
func ParamSpecChar(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gint8
	var arg5 C.gint8
	var arg6 C.gint8
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gint8(minimum)
	arg5 = C.gint8(maximum)
	arg6 = C.gint8(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_char(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecDouble creates a new SpecDouble instance specifying a G_TYPE_DOUBLE
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecDouble(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gdouble
	var arg5 C.gdouble
	var arg6 C.gdouble
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gdouble(minimum)
	arg5 = C.gdouble(maximum)
	arg6 = C.gdouble(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_double(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecEnum creates a new SpecEnum instance specifying a G_TYPE_ENUM
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecEnum(name string, nick string, blurb string, enumType externglib.Type, defaultValue int, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.gint
	var arg6 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(enumType)
	arg5 = C.gint(defaultValue)
	arg6 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_enum(arg1, arg2, arg3, arg4, arg5, arg6)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecFlags creates a new SpecFlags instance specifying a G_TYPE_FLAGS
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecFlags(name string, nick string, blurb string, flagsType externglib.Type, defaultValue uint, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.guint
	var arg6 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(flagsType)
	arg5 = C.guint(defaultValue)
	arg6 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_flags(arg1, arg2, arg3, arg4, arg5, arg6)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecFloat creates a new SpecFloat instance specifying a G_TYPE_FLOAT
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecFloat(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gfloat
	var arg5 C.gfloat
	var arg6 C.gfloat
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gfloat(minimum)
	arg5 = C.gfloat(maximum)
	arg6 = C.gfloat(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_float(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecGType creates a new SpecGType instance specifying a G_TYPE_GTYPE
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecGType(name string, nick string, blurb string, isAType externglib.Type, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(isAType)
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_gtype(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecInt creates a new SpecInt instance specifying a G_TYPE_INT property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecInt(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gint
	var arg5 C.gint
	var arg6 C.gint
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gint(minimum)
	arg5 = C.gint(maximum)
	arg6 = C.gint(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_int(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecInt64 creates a new SpecInt64 instance specifying a G_TYPE_INT64
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecInt64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gint64
	var arg5 C.gint64
	var arg6 C.gint64
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gint64(minimum)
	arg5 = C.gint64(maximum)
	arg6 = C.gint64(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_int64(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecLong creates a new SpecLong instance specifying a G_TYPE_LONG
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecLong(name string, nick string, blurb string, minimum int32, maximum int32, defaultValue int32, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.glong
	var arg5 C.glong
	var arg6 C.glong
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.glong(minimum)
	arg5 = C.glong(maximum)
	arg6 = C.glong(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_long(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecObject creates a new SpecBoxed instance specifying a G_TYPE_OBJECT
// derived property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecObject(name string, nick string, blurb string, objectType externglib.Type, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(objectType)
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_object(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecOverride creates a new property of type SpecOverride. This is used
// to direct operations to another paramspec, and will not be directly useful
// unless you are implementing a new base type similar to GObject.
func ParamSpecOverride(name string, overridden ParamSpec) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.GParamSpec

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GParamSpec)(unsafe.Pointer(overridden.Native()))

	var cret *C.GParamSpec
	var goret ParamSpec

	cret = C.g_param_spec_override(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecParam creates a new SpecParam instance specifying a G_TYPE_PARAM
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecParam(name string, nick string, blurb string, paramType externglib.Type, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GType
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.GType(paramType)
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_param(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecPointer creates a new SpecPointer instance specifying a pointer
// property. Where possible, it is better to use g_param_spec_object() or
// g_param_spec_boxed() to expose memory management information.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecPointer(name string, nick string, blurb string, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_pointer(arg1, arg2, arg3, arg4)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecString creates a new SpecString instance.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecString(name string, nick string, blurb string, defaultValue string, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.gchar
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(defaultValue))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_string(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecUchar creates a new SpecUChar instance specifying a G_TYPE_UCHAR
// property.
func ParamSpecUchar(name string, nick string, blurb string, minimum byte, maximum byte, defaultValue byte, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.guint8
	var arg5 C.guint8
	var arg6 C.guint8
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.guint8(minimum)
	arg5 = C.guint8(maximum)
	arg6 = C.guint8(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_uchar(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecUint creates a new SpecUInt instance specifying a G_TYPE_UINT
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.guint
	var arg5 C.guint
	var arg6 C.guint
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.guint(minimum)
	arg5 = C.guint(maximum)
	arg6 = C.guint(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_uint(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecUint64 creates a new SpecUInt64 instance specifying a G_TYPE_UINT64
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.guint64
	var arg5 C.guint64
	var arg6 C.guint64
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.guint64(minimum)
	arg5 = C.guint64(maximum)
	arg6 = C.guint64(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_uint64(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecUlong creates a new SpecULong instance specifying a G_TYPE_ULONG
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint32, maximum uint32, defaultValue uint32, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gulong
	var arg5 C.gulong
	var arg6 C.gulong
	var arg7 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gulong(minimum)
	arg5 = C.gulong(maximum)
	arg6 = C.gulong(defaultValue)
	arg7 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_ulong(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecUnichar creates a new SpecUnichar instance specifying a G_TYPE_UINT
// property. #GValue structures for this property can be accessed with
// g_value_set_uint() and g_value_get_uint().
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUnichar(name string, nick string, blurb string, defaultValue uint32, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gunichar
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gunichar(defaultValue)
	arg5 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_unichar(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecValueArray creates a new SpecValueArray instance specifying a
// G_TYPE_VALUE_ARRAY property. G_TYPE_VALUE_ARRAY is a G_TYPE_BOXED type, as
// such, #GValue structures for this property can be accessed with
// g_value_set_boxed() and g_value_get_boxed().
//
// See g_param_spec_internal() for details on property names.
func ParamSpecValueArray(name string, nick string, blurb string, elementSpec ParamSpec, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.GParamSpec
	var arg5 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.GParamSpec)(unsafe.Pointer(elementSpec.Native()))
	arg5 = (C.GParamFlags)(flags)

	var cret *C.GParamSpec
	var goret ParamSpec

	cret = C.g_param_spec_value_array(arg1, arg2, arg3, arg4, arg5)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}

// ParamSpecVariant creates a new SpecVariant instance specifying a #GVariant
// property.
//
// If @default_value is floating, it is consumed.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecVariant(name string, nick string, blurb string, typ *glib.VariantType, defaultValue *glib.Variant, flags ParamFlags) ParamSpec {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.GVariantType
	var arg5 *C.GVariant
	var arg6 C.GParamFlags

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.GVariantType)(unsafe.Pointer(typ.Native()))
	arg5 = (*C.GVariant)(unsafe.Pointer(defaultValue.Native()))
	arg6 = (C.GParamFlags)(flags)

	cret := new(C.GParamSpec)
	var goret ParamSpec

	cret = C.g_param_spec_variant(arg1, arg2, arg3, arg4, arg5, arg6)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ParamSpec)

	return goret
}
