// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_object_factory_get_type()), F: marshalObjectFactorier},
	})
}

// ObjectFactoryOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ObjectFactoryOverrider interface {
	// Invalidate: inform factory that it is no longer being used to create
	// accessibles. When called, factory may need to inform Objects which it has
	// created that they need to be re-instantiated. Note: primarily used for
	// runtime replacement of ObjectFactorys in object registries.
	Invalidate()
}

// ObjectFactory: this class is the base object class for a factory used to
// create an accessible object for a specific GType. The function
// atk_registry_set_factory_type() is normally called to store in the registry
// the factory type to be used to create an accessible of a particular GType.
type ObjectFactory struct {
	*externglib.Object
}

func wrapObjectFactory(obj *externglib.Object) *ObjectFactory {
	return &ObjectFactory{
		Object: obj,
	}
}

func marshalObjectFactorier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapObjectFactory(obj), nil
}

// CreateAccessible provides an Object that implements an accessibility
// interface on behalf of obj.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
func (factory *ObjectFactory) CreateAccessible(obj *externglib.Object) *ObjectClass {
	var _arg0 *C.AtkObjectFactory // out
	var _arg1 *C.GObject          // out
	var _cret *C.AtkObject        // in

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(factory.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_object_factory_create_accessible(_arg0, _arg1)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(obj)

	var _object *ObjectClass // out

	_object = wrapObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _object
}

// AccessibleType gets the GType of the accessible which is created by the
// factory.
func (factory *ObjectFactory) AccessibleType() externglib.Type {
	var _arg0 *C.AtkObjectFactory // out
	var _cret C.GType             // in

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(factory.Native()))

	_cret = C.atk_object_factory_get_accessible_type(_arg0)
	runtime.KeepAlive(factory)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// Invalidate: inform factory that it is no longer being used to create
// accessibles. When called, factory may need to inform Objects which it has
// created that they need to be re-instantiated. Note: primarily used for
// runtime replacement of ObjectFactorys in object registries.
func (factory *ObjectFactory) Invalidate() {
	var _arg0 *C.AtkObjectFactory // out

	_arg0 = (*C.AtkObjectFactory)(unsafe.Pointer(factory.Native()))

	C.atk_object_factory_invalidate(_arg0)
	runtime.KeepAlive(factory)
}
