// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFlattenListModel = coreglib.Type(girepository.MustFind("Gtk", "FlattenListModel").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFlattenListModel, F: marshalFlattenListModel},
	})
}

// FlattenListModelOverrides contains methods that are overridable.
type FlattenListModelOverrides struct {
}

func defaultFlattenListModelOverrides(v *FlattenListModel) FlattenListModelOverrides {
	return FlattenListModelOverrides{}
}

// FlattenListModel: GtkFlattenListModel is a list model that concatenates other
// list models.
//
// GtkFlattenListModel takes a list model containing list models, and flattens
// it into a single model.
type FlattenListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*FlattenListModel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FlattenListModel, *FlattenListModelClass, FlattenListModelOverrides](
		GTypeFlattenListModel,
		initFlattenListModelClass,
		wrapFlattenListModel,
		defaultFlattenListModelOverrides,
	)
}

func initFlattenListModelClass(gclass unsafe.Pointer, overrides FlattenListModelOverrides, classInitFunc func(*FlattenListModelClass)) {
	if classInitFunc != nil {
		class := (*FlattenListModelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFlattenListModel(obj *coreglib.Object) *FlattenListModel {
	return &FlattenListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalFlattenListModel(p uintptr) (interface{}, error) {
	return wrapFlattenListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FlattenListModelClass: instance of this type is always passed by reference.
type FlattenListModelClass struct {
	*flattenListModelClass
}

// flattenListModelClass is the struct that's finalized.
type flattenListModelClass struct {
	native unsafe.Pointer
}

var GIRInfoFlattenListModelClass = girepository.MustFind("Gtk", "FlattenListModelClass")
