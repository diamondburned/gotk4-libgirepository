// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

// Type: a numerical value which represents the unique identifier of a
// registered type.
type Type uint

// TypeDebugFlags: these flags used to be passed to
// g_type_init_with_debug_flags() which is now deprecated.
//
// If you need to enable debugging features, use the GOBJECT_DEBUG environment
// variable.
type TypeDebugFlags int

const (
	// TypeDebugFlagsNone: print no messages
	TypeDebugFlagsNone TypeDebugFlags = 0
	// TypeDebugFlagsObjects: print messages about object bookkeeping
	TypeDebugFlagsObjects TypeDebugFlags = 1
	// TypeDebugFlagsSignals: print messages about signal emissions
	TypeDebugFlagsSignals TypeDebugFlags = 2
	// TypeDebugFlagsInstanceCount: keep a count of instances of each type
	TypeDebugFlagsInstanceCount TypeDebugFlags = 4
	// TypeDebugFlagsMask: mask covering all debug flags
	TypeDebugFlagsMask TypeDebugFlags = 7
)

// TypeFlags: bit masks used to check or determine characteristics of a type.
type TypeFlags int

const (
	// TypeFlagsAbstract indicates an abstract type. No instances can be created
	// for an abstract type
	TypeFlagsAbstract TypeFlags = 16
	// TypeFlagsValueAbstract indicates an abstract value type, i.e. a type that
	// introduces a value table, but can't be used for g_value_init()
	TypeFlagsValueAbstract TypeFlags = 32
)

// TypeFundamentalFlags: bit masks used to check or determine specific
// characteristics of a fundamental type.
type TypeFundamentalFlags int

const (
	// TypeFundamentalFlagsClassed indicates a classed type
	TypeFundamentalFlagsClassed TypeFundamentalFlags = 1
	// TypeFundamentalFlagsInstantiatable indicates an instantiable type
	// (implies classed)
	TypeFundamentalFlagsInstantiatable TypeFundamentalFlags = 2
	// TypeFundamentalFlagsDerivable indicates a flat derivable type
	TypeFundamentalFlagsDerivable TypeFundamentalFlags = 4
	// TypeFundamentalFlagsDeepDerivable indicates a deep derivable type
	// (implies derivable)
	TypeFundamentalFlagsDeepDerivable TypeFundamentalFlags = 8
)

// TypeAddClassPrivate registers a private class structure for a classed type;
// when the class is allocated, the private structures for the class and all of
// its parent types are allocated sequentially in the same memory block as the
// public structures, and are zero-filled.
//
// This function should be called in the type's get_type() function after the
// type is registered. The private structure can be retrieved using the
// G_TYPE_CLASS_GET_PRIVATE() macro.
func TypeAddClassPrivate(classType externglib.Type, privateSize uint) {
	var _arg1 C.GType
	var _arg2 C.gsize

	_arg1 = C.GType(classType)
	_arg2 = C.gsize(privateSize)

	C.g_type_add_class_private(_arg1, _arg2)
}

func TypeAddInstancePrivate(classType externglib.Type, privateSize uint) int {
	var _arg1 C.GType
	var _arg2 C.gsize

	_arg1 = C.GType(classType)
	_arg2 = C.gsize(privateSize)

	var _cret C.gint

	cret = C.g_type_add_instance_private(_arg1, _arg2)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// TypeAddInterfaceDynamic adds @interface_type to the dynamic
// @instantiable_type. The information contained in the Plugin structure pointed
// to by @plugin is used to manage the relationship.
func TypeAddInterfaceDynamic(instanceType externglib.Type, interfaceType externglib.Type, plugin TypePlugin) {
	var _arg1 C.GType
	var _arg2 C.GType
	var _arg3 *C.GTypePlugin

	_arg1 = C.GType(instanceType)
	_arg2 = C.GType(interfaceType)
	_arg3 = (*C.GTypePlugin)(unsafe.Pointer(plugin.Native()))

	C.g_type_add_interface_dynamic(_arg1, _arg2, _arg3)
}

// TypeAddInterfaceStatic adds @interface_type to the static @instantiable_type.
// The information contained in the Info structure pointed to by @info is used
// to manage the relationship.
func TypeAddInterfaceStatic(instanceType externglib.Type, interfaceType externglib.Type, info *InterfaceInfo) {
	var _arg1 C.GType
	var _arg2 C.GType
	var _arg3 *C.GInterfaceInfo

	_arg1 = C.GType(instanceType)
	_arg2 = C.GType(interfaceType)
	_arg3 = (*C.GInterfaceInfo)(unsafe.Pointer(info.Native()))

	C.g_type_add_interface_static(_arg1, _arg2, _arg3)
}

// TypeCheckInstance: private helper function to aid implementation of the
// G_TYPE_CHECK_INSTANCE() macro.
func TypeCheckInstance(instance *TypeInstance) bool {
	var _arg1 *C.GTypeInstance

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	var _cret C.gboolean

	cret = C.g_type_check_instance(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

func TypeCheckInstanceCast(instance *TypeInstance, ifaceType externglib.Type) *TypeInstance {
	var _arg1 *C.GTypeInstance
	var _arg2 C.GType

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	_arg2 = C.GType(ifaceType)

	var _cret *C.GTypeInstance

	cret = C.g_type_check_instance_cast(_arg1, _arg2)

	var _typeInstance *TypeInstance

	_typeInstance = WrapTypeInstance(unsafe.Pointer(_cret))

	return _typeInstance
}

func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType externglib.Type) bool {
	var _arg1 *C.GTypeInstance
	var _arg2 C.GType

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	_arg2 = C.GType(ifaceType)

	var _cret C.gboolean

	cret = C.g_type_check_instance_is_a(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType externglib.Type) bool {
	var _arg1 *C.GTypeInstance
	var _arg2 C.GType

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	_arg2 = C.GType(fundamentalType)

	var _cret C.gboolean

	cret = C.g_type_check_instance_is_fundamentally_a(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

func TypeCheckIsValueType(typ externglib.Type) bool {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret C.gboolean

	cret = C.g_type_check_is_value_type(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

func TypeCheckValue(value **externglib.Value) bool {
	var _arg1 *C.GValue

	_arg1 = (*C.GValue)(value.GValue)

	var _cret C.gboolean

	cret = C.g_type_check_value(_arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

func TypeCheckValueHolds(value **externglib.Value, typ externglib.Type) bool {
	var _arg1 *C.GValue
	var _arg2 C.GType

	_arg1 = (*C.GValue)(value.GValue)
	_arg2 = C.GType(typ)

	var _cret C.gboolean

	cret = C.g_type_check_value_holds(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// TypeChildren: return a newly allocated and 0-terminated array of type IDs,
// listing the child types of @type.
func TypeChildren(typ externglib.Type) []externglib.Type {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret *C.GType
	var _arg2 *C.guint

	cret = C.g_type_children(_arg1)

	var _gTypes []externglib.Type

	{
		var src []C.GType
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg2))

		_gTypes = make([]externglib.Type, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_gTypes = externglib.Type(_cret)
		}
	}

	return _gTypes
}

// TypeCreateInstance creates and initializes an instance of @type if @type is
// valid and can be instantiated. The type system only performs basic allocation
// and structure setups for instances: actual instance creation should happen
// through functions supplied by the type's fundamental type implementation. So
// use of g_type_create_instance() is reserved for implementers of fundamental
// types only. E.g. instances of the #GObject hierarchy should be created via
// g_object_new() and never directly through g_type_create_instance() which
// doesn't handle things like singleton objects or object construction.
//
// The extended members of the returned instance are guaranteed to be filled
// with zeros.
//
// Note: Do not use this function, unless you're implementing a fundamental
// type. Also language bindings should not use this function, but g_object_new()
// instead.
func TypeCreateInstance(typ externglib.Type) *TypeInstance {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret *C.GTypeInstance

	cret = C.g_type_create_instance(_arg1)

	var _typeInstance *TypeInstance

	_typeInstance = WrapTypeInstance(unsafe.Pointer(_cret))

	return _typeInstance
}

// TypeDepth returns the length of the ancestry of the passed in type. This
// includes the type itself, so that e.g. a fundamental type has depth 1.
func TypeDepth(typ externglib.Type) uint {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret C.guint

	cret = C.g_type_depth(_arg1)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// TypeEnsure ensures that the indicated @type has been registered with the type
// system, and its _class_init() method has been run.
//
// In theory, simply calling the type's _get_type() method (or using the
// corresponding macro) is supposed take care of this. However, _get_type()
// methods are often marked G_GNUC_CONST for performance reasons, even though
// this is technically incorrect (since G_GNUC_CONST requires that the function
// not have side effects, which _get_type() methods do on the first call). As a
// result, if you write a bare call to a _get_type() macro, it may get optimized
// out by the compiler. Using g_type_ensure() guarantees that the type's
// _get_type() method is called.
func TypeEnsure(typ externglib.Type) {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	C.g_type_ensure(_arg1)
}

// TypeFreeInstance frees an instance of a type, returning it to the instance
// pool for the type, if there is one.
//
// Like g_type_create_instance(), this function is reserved for implementors of
// fundamental types.
func TypeFreeInstance(instance *TypeInstance) {
	var _arg1 *C.GTypeInstance

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	C.g_type_free_instance(_arg1)
}

// TypeFromName: look up the type ID from a given type name, returning 0 if no
// type has been registered under this name (this is the preferred method to
// find out by name whether a specific type has been registered yet).
func TypeFromName(name string) externglib.Type {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GType

	cret = C.g_type_from_name(_arg1)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeFundamental: internal function, used to extract the fundamental type ID
// portion. Use G_TYPE_FUNDAMENTAL() instead.
func TypeFundamental(typeId externglib.Type) externglib.Type {
	var _arg1 C.GType

	_arg1 = C.GType(typeId)

	var _cret C.GType

	cret = C.g_type_fundamental(_arg1)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeFundamentalNext returns the next free fundamental type id which can be
// used to register a new fundamental type with g_type_register_fundamental().
// The returned type ID represents the highest currently registered fundamental
// type identifier.
func TypeFundamentalNext() externglib.Type {
	var _cret C.GType

	cret = C.g_type_fundamental_next()

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeGetInstanceCount returns the number of instances allocated of the
// particular type; this is only available if GLib is built with debugging
// support and the instance_count debug flag is set (by setting the
// GOBJECT_DEBUG variable to include instance-count).
func TypeGetInstanceCount(typ externglib.Type) int {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret C.int

	cret = C.g_type_get_instance_count(_arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// TypeGetPlugin returns the Plugin structure for @type.
func TypeGetPlugin(typ externglib.Type) TypePlugin {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret *C.GTypePlugin

	cret = C.g_type_get_plugin(_arg1)

	var _typePlugin TypePlugin

	_typePlugin = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TypePlugin)

	return _typePlugin
}

// TypeGetTypeRegistrationSerial returns an opaque serial number that represents
// the state of the set of registered types. Any time a type is registered this
// serial changes, which means you can cache information based on type lookups
// (such as g_type_from_name()) and know if the cache is still valid at a later
// time by comparing the current serial with the one at the type lookup.
func TypeGetTypeRegistrationSerial() uint {
	var _cret C.guint

	cret = C.g_type_get_type_registration_serial()

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// TypeInit: this function used to initialise the type system. Since GLib 2.36,
// the type system is initialised automatically and this function does nothing.
func TypeInit() {
	C.g_type_init()
}

// TypeInitWithDebugFlags: this function used to initialise the type system with
// debugging flags. Since GLib 2.36, the type system is initialised
// automatically and this function does nothing.
//
// If you need to enable debugging features, use the GOBJECT_DEBUG environment
// variable.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	var _arg1 C.GTypeDebugFlags

	_arg1 = (C.GTypeDebugFlags)(debugFlags)

	C.g_type_init_with_debug_flags(_arg1)
}

// TypeInterfaces: return a newly allocated and 0-terminated array of type IDs,
// listing the interface types that @type conforms to.
func TypeInterfaces(typ externglib.Type) []externglib.Type {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret *C.GType
	var _arg2 *C.guint

	cret = C.g_type_interfaces(_arg1)

	var _gTypes []externglib.Type

	{
		var src []C.GType
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg2))

		_gTypes = make([]externglib.Type, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_gTypes = externglib.Type(_cret)
		}
	}

	return _gTypes
}

// TypeIsA: if @is_a_type is a derivable type, check whether @type is a
// descendant of @is_a_type. If @is_a_type is an interface, check whether @type
// conforms to it.
func TypeIsA(typ externglib.Type, isAType externglib.Type) bool {
	var _arg1 C.GType
	var _arg2 C.GType

	_arg1 = C.GType(typ)
	_arg2 = C.GType(isAType)

	var _cret C.gboolean

	cret = C.g_type_is_a(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// TypeName: get the unique name that is assigned to a type ID. Note that this
// function (like all other GType API) cannot cope with invalid type IDs.
// G_TYPE_INVALID may be passed to this function, as may be any other validly
// registered type ID, but randomized type IDs should not be passed in and will
// most likely lead to a crash.
func TypeName(typ externglib.Type) string {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret *C.gchar

	cret = C.g_type_name(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

func TypeNameFromInstance(instance *TypeInstance) string {
	var _arg1 *C.GTypeInstance

	_arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	var _cret *C.gchar

	cret = C.g_type_name_from_instance(_arg1)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// TypeNextBase: given a @leaf_type and a @root_type which is contained in its
// anchestry, return the type that @root_type is the immediate parent of. In
// other words, this function determines the type that is derived directly from
// @root_type which is also a base class of @leaf_type. Given a root type and a
// leaf type, this function can be used to determine the types and order in
// which the leaf type is descended from the root type.
func TypeNextBase(leafType externglib.Type, rootType externglib.Type) externglib.Type {
	var _arg1 C.GType
	var _arg2 C.GType

	_arg1 = C.GType(leafType)
	_arg2 = C.GType(rootType)

	var _cret C.GType

	cret = C.g_type_next_base(_arg1, _arg2)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeParent: return the direct parent type of the passed in type. If the
// passed in type has no parent, i.e. is a fundamental type, 0 is returned.
func TypeParent(typ externglib.Type) externglib.Type {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _cret C.GType

	cret = C.g_type_parent(_arg1)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeQuery queries the type system for information about a specific type. This
// function will fill in a user-provided structure to hold type-specific
// information. If an invalid #GType is passed in, the @type member of the Query
// is 0. All members filled into the Query structure should be considered
// constant and have to be left untouched.
func TypeQuery(typ externglib.Type) TypeQuery {
	var _arg1 C.GType

	_arg1 = C.GType(typ)

	var _query TypeQuery

	C.g_type_query(_arg1, (*C.GTypeQuery)(unsafe.Pointer(&_query)))

	return _query
}

// TypeRegisterDynamic registers @type_name as the name of a new dynamic type
// derived from @parent_type. The type system uses the information contained in
// the Plugin structure pointed to by @plugin to manage the type and its
// instances (if not abstract). The value of @flags determines the nature (e.g.
// abstract or not) of the type.
func TypeRegisterDynamic(parentType externglib.Type, typeName string, plugin TypePlugin, flags TypeFlags) externglib.Type {
	var _arg1 C.GType
	var _arg2 *C.gchar
	var _arg3 *C.GTypePlugin
	var _arg4 C.GTypeFlags

	_arg1 = C.GType(parentType)
	_arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GTypePlugin)(unsafe.Pointer(plugin.Native()))
	_arg4 = (C.GTypeFlags)(flags)

	var _cret C.GType

	cret = C.g_type_register_dynamic(_arg1, _arg2, _arg3, _arg4)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeRegisterFundamental registers @type_id as the predefined identifier and
// @type_name as the name of a fundamental type. If @type_id is already
// registered, or a type named @type_name is already registered, the behaviour
// is undefined. The type system uses the information contained in the Info
// structure pointed to by @info and the FundamentalInfo structure pointed to by
// @finfo to manage the type and its instances. The value of @flags determines
// additional characteristics of the fundamental type.
func TypeRegisterFundamental(typeId externglib.Type, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) externglib.Type {
	var _arg1 C.GType
	var _arg2 *C.gchar
	var _arg3 *C.GTypeInfo
	var _arg4 *C.GTypeFundamentalInfo
	var _arg5 C.GTypeFlags

	_arg1 = C.GType(typeId)
	_arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GTypeInfo)(unsafe.Pointer(info.Native()))
	_arg4 = (*C.GTypeFundamentalInfo)(unsafe.Pointer(finfo.Native()))
	_arg5 = (C.GTypeFlags)(flags)

	var _cret C.GType

	cret = C.g_type_register_fundamental(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// TypeRegisterStatic registers @type_name as the name of a new static type
// derived from @parent_type. The type system uses the information contained in
// the Info structure pointed to by @info to manage the type and its instances
// (if not abstract). The value of @flags determines the nature (e.g. abstract
// or not) of the type.
func TypeRegisterStatic(parentType externglib.Type, typeName string, info *TypeInfo, flags TypeFlags) externglib.Type {
	var _arg1 C.GType
	var _arg2 *C.gchar
	var _arg3 *C.GTypeInfo
	var _arg4 C.GTypeFlags

	_arg1 = C.GType(parentType)
	_arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GTypeInfo)(unsafe.Pointer(info.Native()))
	_arg4 = (C.GTypeFlags)(flags)

	var _cret C.GType

	cret = C.g_type_register_static(_arg1, _arg2, _arg3, _arg4)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

func TypeTestFlags(typ externglib.Type, flags uint) bool {
	var _arg1 C.GType
	var _arg2 C.guint

	_arg1 = C.GType(typ)
	_arg2 = C.guint(flags)

	var _cret C.gboolean

	cret = C.g_type_test_flags(_arg1, _arg2)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// InterfaceInfo: a structure that provides information to the type system which
// is used specifically for managing interface types.
type InterfaceInfo struct {
	native C.GInterfaceInfo
}

// WrapInterfaceInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInterfaceInfo(ptr unsafe.Pointer) *InterfaceInfo {
	if ptr == nil {
		return nil
	}

	return (*InterfaceInfo)(ptr)
}

func marshalInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapInterfaceInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *InterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// InterfaceData gets the field inside the struct.
func (i *InterfaceInfo) InterfaceData() interface{} {
	var v interface{}
	v = (interface{})(i.native.interface_data)
	return v
}

// TypeFundamentalInfo: a structure that provides information to the type system
// which is used specifically for managing fundamental types.
type TypeFundamentalInfo struct {
	native C.GTypeFundamentalInfo
}

// WrapTypeFundamentalInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeFundamentalInfo(ptr unsafe.Pointer) *TypeFundamentalInfo {
	if ptr == nil {
		return nil
	}

	return (*TypeFundamentalInfo)(ptr)
}

func marshalTypeFundamentalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeFundamentalInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeFundamentalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TypeFlags gets the field inside the struct.
func (t *TypeFundamentalInfo) TypeFlags() TypeFundamentalFlags {
	var v TypeFundamentalFlags
	v = TypeFundamentalFlags(t.native.type_flags)
	return v
}

// TypeInfo: this structure is used to provide the type system with the
// information required to initialize and destruct (finalize) a type's class and
// its instances.
//
// The initialized structure is passed to the g_type_register_static() function
// (or is copied into the provided Info structure in the
// g_type_plugin_complete_type_info()). The type system will perform a deep copy
// of this structure, so its memory does not need to be persistent across
// invocation of g_type_register_static().
type TypeInfo struct {
	native C.GTypeInfo
}

// WrapTypeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeInfo(ptr unsafe.Pointer) *TypeInfo {
	if ptr == nil {
		return nil
	}

	return (*TypeInfo)(ptr)
}

func marshalTypeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// ClassSize gets the field inside the struct.
func (t *TypeInfo) ClassSize() uint16 {
	var v uint16
	v = (uint16)(t.native.class_size)
	return v
}

// ClassData gets the field inside the struct.
func (t *TypeInfo) ClassData() interface{} {
	var v interface{}
	v = (interface{})(t.native.class_data)
	return v
}

// InstanceSize gets the field inside the struct.
func (t *TypeInfo) InstanceSize() uint16 {
	var v uint16
	v = (uint16)(t.native.instance_size)
	return v
}

// NPreallocs gets the field inside the struct.
func (t *TypeInfo) NPreallocs() uint16 {
	var v uint16
	v = (uint16)(t.native.n_preallocs)
	return v
}

// TypeInstance: an opaque structure used as the base of all type instances.
type TypeInstance struct {
	native C.GTypeInstance
}

// WrapTypeInstance wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeInstance(ptr unsafe.Pointer) *TypeInstance {
	if ptr == nil {
		return nil
	}

	return (*TypeInstance)(ptr)
}

func marshalTypeInstance(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeInstance(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeInstance) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

func (i *TypeInstance) Private(privateType externglib.Type) interface{} {
	var _arg0 *C.GTypeInstance
	var _arg1 C.GType

	_arg0 = (*C.GTypeInstance)(unsafe.Pointer(i.Native()))
	_arg1 = C.GType(privateType)

	var _cret C.gpointer

	cret = C.g_type_instance_get_private(_arg0, _arg1)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// TypeQuery: a structure holding information for a specific type. It is filled
// in by the g_type_query() function.
type TypeQuery struct {
	native C.GTypeQuery
}

// WrapTypeQuery wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeQuery(ptr unsafe.Pointer) *TypeQuery {
	if ptr == nil {
		return nil
	}

	return (*TypeQuery)(ptr)
}

func marshalTypeQuery(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeQuery(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeQuery) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Type gets the field inside the struct.
func (t *TypeQuery) Type() externglib.Type {
	var v externglib.Type
	v = externglib.Type(t.native._type)
	return v
}

// TypeName gets the field inside the struct.
func (t *TypeQuery) TypeName() string {
	var v string
	v = C.GoString(t.native.type_name)
	return v
}

// ClassSize gets the field inside the struct.
func (t *TypeQuery) ClassSize() uint {
	var v uint
	v = (uint)(t.native.class_size)
	return v
}

// InstanceSize gets the field inside the struct.
func (t *TypeQuery) InstanceSize() uint {
	var v uint
	v = (uint)(t.native.instance_size)
	return v
}
