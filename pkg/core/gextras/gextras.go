// Package gextras contains supplemental types to gotk3.
package gextras

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
import "C"

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// Objector is an interface that describes partially the glib.Object type.
type Objector interface {
	Connect(string, interface{}) glib.SignalHandle
	ConnectAfter(string, interface{}) glib.SignalHandle

	HandlerBlock(glib.SignalHandle)
	HandlerDisconnect(glib.SignalHandle)
	HandlerUnblock(glib.SignalHandle)

	GetProperty(string) (interface{}, error)
	SetProperty(string, interface{}) error

	Native() uintptr
}

var _ Objector = (*glib.Object)(nil)

// MustSet panics if the given property cannot be set into the given object.
func MustSet(obj Objector, k string, v interface{}) {
	if err := obj.SetProperty(k, v); err != nil {
		log.Panicf("cannot set object property %q: %s", k, err)
	}
}

// MustGet panics if the given property cannot be retrieved from the given
// object.
func MustGet(obj Objector, k string) interface{} {
	v, err := obj.GetProperty(k)
	if err != nil {
		log.Panicf("cannot get object property %q: %s", k, err)
	}
	return v
}

// GetInto gets the given property key from the object into the given pointer.
// An error is returned if it cannot get the property or the type is wrong. This
// method is mostly useful for avoiding type assertions.
func GetInto(obj Objector, k string, ptr interface{}) error {
	return getInto(obj, k, ptr, false)
}

// MustGetInto is similar to GetInfo, except it does not do safety checks and
// will panic on an error. Code that uses constants should use this function
// over GetInto.
func MustGetInto(obj Objector, k string, ptr interface{}) {
	getInto(obj, k, ptr, true)
}

func getInto(obj Objector, k string, ptr interface{}, must bool) error {
	dst := reflect.ValueOf(ptr)
	if !must {
		typ := dst.Type()
		if !(typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Interface) {
			return fmt.Errorf("ptr type %s is not a pointer", typ)
		}
	}

	elem := dst.Elem()
	if !must {
		if !elem.CanSet() {
			return errors.New("ptr's value cannot be set")
		}
	}

	v, err := obj.GetProperty(k)
	if err != nil {
		if !must {
			return fmt.Errorf("cannot get object property %q: %s", k, err)
		}

		log.Panicf("cannot get object property %q: %s", k, err)
	}

	val := reflect.ValueOf(v)

	if !must {
		property := val.Type()
		given := elem.Type()

		if !property.AssignableTo(given) {
			return fmt.Errorf("property type %s not assignable to given %s", property, given)
		}
	}

	elem.Set(val)
	return nil
}

// InternObject gets the internal Object type. This is used for calling methods
// not in the Objector.
func InternObject(obj Objector) *glib.Object {
	return &glib.Object{
		GObject: glib.ToGObject(unsafe.Pointer(obj.Native())),
	}
}

// CastObject casts the given object pointer to the class name. The caller is
// responsible for recasting the interface to the wanted type.
func CastObject(obj *glib.Object) interface{} {
	var gvalue C.GValue

	C.g_value_init(&gvalue, C.GType(obj.TypeFromInstance()))
	defer C.g_value_unset(&gvalue)

	v, err := glib.ValueFromNative(unsafe.Pointer(&gvalue)).GoValue()
	if err != nil {
		return Objector(obj)
	}

	return v
}