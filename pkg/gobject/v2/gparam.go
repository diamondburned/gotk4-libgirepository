// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

// ParamFlags: through the Flags flag values, certain aspects of parameters can
// be configured. See also PARAM_STATIC_STRINGS.
type ParamFlags int

const (
	// ParamFlagsReadable: the parameter is readable
	ParamFlagsReadable ParamFlags = 1
	// ParamFlagsWritable: the parameter is writable
	ParamFlagsWritable ParamFlags = 2
	// ParamFlagsReadwrite alias for G_PARAM_READABLE | G_PARAM_WRITABLE
	ParamFlagsReadwrite ParamFlags = 3
	// ParamFlagsConstruct: the parameter will be set upon object construction
	ParamFlagsConstruct ParamFlags = 4
	// ParamFlagsConstructOnly: the parameter can only be set upon object
	// construction
	ParamFlagsConstructOnly ParamFlags = 8
	// ParamFlagsLaxValidation: upon parameter conversion (see
	// g_param_value_convert()) strict validation is not required
	ParamFlagsLaxValidation ParamFlags = 16
	// ParamFlagsStaticName: the string used as name when constructing the
	// parameter is guaranteed to remain valid and unmodified for the lifetime
	// of the parameter. Since 2.8
	ParamFlagsStaticName ParamFlags = 32
	// ParamFlagsPrivate: internal
	ParamFlagsPrivate ParamFlags = 32
	// ParamFlagsStaticNick: the string used as nick when constructing the
	// parameter is guaranteed to remain valid and unmmodified for the lifetime
	// of the parameter. Since 2.8
	ParamFlagsStaticNick ParamFlags = 64
	// ParamFlagsStaticBlurb: the string used as blurb when constructing the
	// parameter is guaranteed to remain valid and unmodified for the lifetime
	// of the parameter. Since 2.8
	ParamFlagsStaticBlurb ParamFlags = 128
	// ParamFlagsExplicitNotify calls to g_object_set_property() for this
	// property will not automatically result in a "notify" signal being
	// emitted: the implementation must call g_object_notify() themselves in
	// case the property actually changes. Since: 2.42.
	ParamFlagsExplicitNotify ParamFlags = 1073741824
	// ParamFlagsDeprecated: the parameter is deprecated and will be removed in
	// a future version. A warning will be generated if it is used while running
	// with G_ENABLE_DIAGNOSTIC=1. Since 2.26
	ParamFlagsDeprecated ParamFlags = 2147483648
)

// ParamValueConvert transforms @src_value into @dest_value if possible, and
// then validates @dest_value, in order for it to conform to @pspec. If
// @strict_validation is true this function will only succeed if the transformed
// @dest_value complied to @pspec without modifications.
//
// See also g_value_type_transformable(), g_value_transform() and
// g_param_value_validate().
func ParamValueConvert(pspec ParamSpec, srcValue **externglib.Value, destValue **externglib.Value, strictValidation bool) bool {
	var _arg1 *C.GParamSpec
	var _arg2 *C.GValue
	var _arg3 *C.GValue
	var _arg4 C.gboolean

	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = (*C.GValue)(srcValue.GValue)
	_arg3 = (*C.GValue)(destValue.GValue)
	if strictValidation {
		_arg4 = C.gboolean(1)
	}

	var _cret C.gboolean

	cret = C.g_param_value_convert(_arg1, _arg2, _arg3, _arg4)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ParamValueDefaults checks whether @value contains the default value as
// specified in @pspec.
func ParamValueDefaults(pspec ParamSpec, value **externglib.Value) bool {
	var _arg1 *C.GParamSpec
	var _arg2 *C.GValue

	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	var _cret C.gboolean

	cret = C.g_param_value_defaults(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ParamValueSetDefault sets @value to its default value as specified in @pspec.
func ParamValueSetDefault(pspec ParamSpec, value **externglib.Value) {
	var _arg1 *C.GParamSpec
	var _arg2 *C.GValue

	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	C.g_param_value_set_default(_arg1, _arg2)
}

// ParamValueValidate ensures that the contents of @value comply with the
// specifications set out by @pspec. For example, a SpecInt might require that
// integers stored in @value may not be smaller than -42 and not be greater than
// +42. If @value contains an integer outside of this range, it is modified
// accordingly, so the resulting value will fit into the range -42 .. +42.
func ParamValueValidate(pspec ParamSpec, value **externglib.Value) bool {
	var _arg1 *C.GParamSpec
	var _arg2 *C.GValue

	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = (*C.GValue)(value.GValue)

	var _cret C.gboolean

	cret = C.g_param_value_validate(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ParamValuesCmp compares @value1 with @value2 according to @pspec, and return
// -1, 0 or +1, if @value1 is found to be less than, equal to or greater than
// @value2, respectively.
func ParamValuesCmp(pspec ParamSpec, value1 **externglib.Value, value2 **externglib.Value) int {
	var _arg1 *C.GParamSpec
	var _arg2 *C.GValue
	var _arg3 *C.GValue

	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = (*C.GValue)(value1.GValue)
	_arg3 = (*C.GValue)(value2.GValue)

	var _cret C.gint

	cret = C.g_param_values_cmp(_arg1, _arg2, _arg3)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// ParamSpecPool: a SpecPool maintains a collection of Specs which can be
// quickly accessed by owner and name. The implementation of the #GObject
// property system uses such a pool to store the Specs of the properties all
// object types.
type ParamSpecPool struct {
	native C.GParamSpecPool
}

// WrapParamSpecPool wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapParamSpecPool(ptr unsafe.Pointer) *ParamSpecPool {
	if ptr == nil {
		return nil
	}

	return (*ParamSpecPool)(ptr)
}

func marshalParamSpecPool(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapParamSpecPool(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *ParamSpecPool) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Insert inserts a Spec in the pool.
func (p *ParamSpecPool) Insert(pspec ParamSpec, ownerType externglib.Type) {
	var _arg0 *C.GParamSpecPool
	var _arg1 *C.GParamSpec
	var _arg2 C.GType

	_arg0 = (*C.GParamSpecPool)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	_arg2 = C.GType(ownerType)

	C.g_param_spec_pool_insert(_arg0, _arg1, _arg2)
}

// List gets an array of all Specs owned by @owner_type in the pool.
func (p *ParamSpecPool) List(ownerType externglib.Type) []ParamSpec {
	var _arg0 *C.GParamSpecPool
	var _arg1 C.GType

	_arg0 = (*C.GParamSpecPool)(unsafe.Pointer(p.Native()))
	_arg1 = C.GType(ownerType)

	var _cret **C.GParamSpec
	var _arg2 *C.guint

	cret = C.g_param_spec_pool_list(_arg0, _arg1)

	var _paramSpecs []ParamSpec

	{
		var src []*C.GParamSpec
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg2))

		_paramSpecs = make([]ParamSpec, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_paramSpecs = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ParamSpec)
		}
	}

	return _paramSpecs
}

// ListOwned gets an #GList of all Specs owned by @owner_type in the pool.
func (p *ParamSpecPool) ListOwned(ownerType externglib.Type) *glib.List {
	var _arg0 *C.GParamSpecPool
	var _arg1 C.GType

	_arg0 = (*C.GParamSpecPool)(unsafe.Pointer(p.Native()))
	_arg1 = C.GType(ownerType)

	var _cret *C.GList

	cret = C.g_param_spec_pool_list_owned(_arg0, _arg1)

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _list
}

// Lookup looks up a Spec in the pool.
func (p *ParamSpecPool) Lookup(paramName string, ownerType externglib.Type, walkAncestors bool) ParamSpec {
	var _arg0 *C.GParamSpecPool
	var _arg1 *C.gchar
	var _arg2 C.GType
	var _arg3 C.gboolean

	_arg0 = (*C.GParamSpecPool)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(paramName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GType(ownerType)
	if walkAncestors {
		_arg3 = C.gboolean(1)
	}

	var _cret *C.GParamSpec

	cret = C.g_param_spec_pool_lookup(_arg0, _arg1, _arg2, _arg3)

	var _paramSpec ParamSpec

	_paramSpec = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// Remove removes a Spec from the pool.
func (p *ParamSpecPool) Remove(pspec ParamSpec) {
	var _arg0 *C.GParamSpecPool
	var _arg1 *C.GParamSpec

	_arg0 = (*C.GParamSpecPool)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))

	C.g_param_spec_pool_remove(_arg0, _arg1)
}

// Parameter: the GParameter struct is an auxiliary structure used to hand
// parameter name/value pairs to g_object_newv().
type Parameter struct {
	native C.GParameter
}

// WrapParameter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapParameter(ptr unsafe.Pointer) *Parameter {
	if ptr == nil {
		return nil
	}

	return (*Parameter)(ptr)
}

func marshalParameter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapParameter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *Parameter) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Name gets the field inside the struct.
func (p *Parameter) Name() string {
	var v string
	v = C.GoString(p.native.name)
	return v
}

// Value gets the field inside the struct.
func (p *Parameter) Value() **externglib.Value {
	var v **externglib.Value
	v = externglib.ValueFromNative(unsafe.Pointer(p.native.value))
	return v
}