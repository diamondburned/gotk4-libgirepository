// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for atkrelation.go.
var GTypeRelation = coreglib.Type(C.atk_relation_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeRelation, F: marshalRelation},
	})
}

// RelationOverrider contains methods that are overridable.
type RelationOverrider interface {
}

// Relation describes a relation between an object and one or more other
// objects. The actual relations that an object has with other objects are
// defined as an AtkRelationSet, which is a set of AtkRelations.
type Relation struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Relation)(nil)
)

func classInitRelationer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRelation(obj *coreglib.Object) *Relation {
	return &Relation{
		Object: obj,
	}
}

func marshalRelation(p uintptr) (interface{}, error) {
	return wrapRelation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddTarget adds the specified AtkObject to the target for the relation, if it
// is not already present. See also atk_object_add_relationship().
//
// The function takes the following parameters:
//
//    - target: Object.
//
func (relation *Relation) AddTarget(target *ObjectClass) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(relation).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(target).Native()))

	girepository.MustFind("Atk", "Relation").InvokeMethod("add_target", _args[:], nil)

	runtime.KeepAlive(relation)
	runtime.KeepAlive(target)
}

// RemoveTarget: remove the specified AtkObject from the target for the
// relation.
//
// The function takes the following parameters:
//
//    - target: Object.
//
// The function returns the following values:
//
//    - ok: TRUE if the removal is successful.
//
func (relation *Relation) RemoveTarget(target *ObjectClass) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(relation).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(target).Native()))

	_gret := girepository.MustFind("Atk", "Relation").InvokeMethod("remove_target", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(relation)
	runtime.KeepAlive(target)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
