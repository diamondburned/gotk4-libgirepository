// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeTLSDatabase = coreglib.Type(girepository.MustFind("Gio", "TlsDatabase").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTLSDatabase, F: marshalTLSDatabase},
	})
}

// TLSDatabaseOverrides contains methods that are overridable.
type TLSDatabaseOverrides struct {
}

func defaultTLSDatabaseOverrides(v *TLSDatabase) TLSDatabaseOverrides {
	return TLSDatabaseOverrides{}
}

// TLSDatabase is used to look up certificates and other information from a
// certificate or key store. It is an abstract base class which TLS library
// specific subtypes override.
//
// A Database may be accessed from multiple threads by the TLS backend. All
// implementations are required to be fully thread-safe.
//
// Most common client applications will not directly interact with Database. It
// is used internally by Connection.
type TLSDatabase struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSDatabase)(nil)
)

// TLSDatabaser describes types inherited from class TLSDatabase.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type TLSDatabaser interface {
	coreglib.Objector
	baseTLSDatabase() *TLSDatabase
}

var _ TLSDatabaser = (*TLSDatabase)(nil)

func init() {
	coreglib.RegisterClassInfo[*TLSDatabase, *TLSDatabaseClass, TLSDatabaseOverrides](
		GTypeTLSDatabase,
		initTLSDatabaseClass,
		wrapTLSDatabase,
		defaultTLSDatabaseOverrides,
	)
}

func initTLSDatabaseClass(gclass unsafe.Pointer, overrides TLSDatabaseOverrides, classInitFunc func(*TLSDatabaseClass)) {
	if classInitFunc != nil {
		class := (*TLSDatabaseClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTLSDatabase(obj *coreglib.Object) *TLSDatabase {
	return &TLSDatabase{
		Object: obj,
	}
}

func marshalTLSDatabase(p uintptr) (interface{}, error) {
	return wrapTLSDatabase(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TLSDatabase) baseTLSDatabase() *TLSDatabase {
	return v
}

// BaseTLSDatabase returns the underlying base object.
func BaseTLSDatabase(obj TLSDatabaser) *TLSDatabase {
	return obj.baseTLSDatabase()
}

// TLSDatabaseClass class for Database. Derived classes should implement the
// various virtual methods. _async and _finish methods have a default
// implementation that runs the corresponding sync method in a thread.
//
// An instance of this type is always passed by reference.
type TLSDatabaseClass struct {
	*tlsDatabaseClass
}

// tlsDatabaseClass is the struct that's finalized.
type tlsDatabaseClass struct {
	native unsafe.Pointer
}

var GIRInfoTLSDatabaseClass = girepository.MustFind("Gio", "TlsDatabaseClass")
