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
	GTypeMultiSelection = coreglib.Type(girepository.MustFind("Gtk", "MultiSelection").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMultiSelection, F: marshalMultiSelection},
	})
}

// MultiSelectionOverrides contains methods that are overridable.
type MultiSelectionOverrides struct {
}

func defaultMultiSelectionOverrides(v *MultiSelection) MultiSelectionOverrides {
	return MultiSelectionOverrides{}
}

// MultiSelection: GtkMultiSelection is a GtkSelectionModel that allows
// selecting multiple elements.
type MultiSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SelectionModel
}

var (
	_ coreglib.Objector = (*MultiSelection)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MultiSelection, *MultiSelectionClass, MultiSelectionOverrides](
		GTypeMultiSelection,
		initMultiSelectionClass,
		wrapMultiSelection,
		defaultMultiSelectionOverrides,
	)
}

func initMultiSelectionClass(gclass unsafe.Pointer, overrides MultiSelectionOverrides, classInitFunc func(*MultiSelectionClass)) {
	if classInitFunc != nil {
		class := (*MultiSelectionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMultiSelection(obj *coreglib.Object) *MultiSelection {
	return &MultiSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	return wrapMultiSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MultiSelectionClass: instance of this type is always passed by reference.
type MultiSelectionClass struct {
	*multiSelectionClass
}

// multiSelectionClass is the struct that's finalized.
type multiSelectionClass struct {
	native unsafe.Pointer
}

var GIRInfoMultiSelectionClass = girepository.MustFind("Gtk", "MultiSelectionClass")
