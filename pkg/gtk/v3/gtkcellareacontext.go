// Code generated by girgen. DO NOT EDIT.

package gtk

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
	GTypeCellAreaContext = coreglib.Type(girepository.MustFind("Gtk", "CellAreaContext").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellAreaContext, F: marshalCellAreaContext},
	})
}

// CellAreaContextOverrides contains methods that are overridable.
type CellAreaContextOverrides struct {
}

func defaultCellAreaContextOverrides(v *CellAreaContext) CellAreaContextOverrides {
	return CellAreaContextOverrides{}
}

// CellAreaContext object is created by a given CellArea implementation via its
// CellAreaClass.create_context() virtual method and is used to store cell sizes
// and alignments for a series of TreeModel rows that are requested and rendered
// in the same context.
//
// CellLayout widgets can create any number of contexts in which to request and
// render groups of data rows. However, it’s important that the same context
// which was used to request sizes for a given TreeModel row also be used for
// the same row when calling other CellArea APIs such as gtk_cell_area_render()
// and gtk_cell_area_event().
type CellAreaContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellAreaContext)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CellAreaContext, *CellAreaContextClass, CellAreaContextOverrides](
		GTypeCellAreaContext,
		initCellAreaContextClass,
		wrapCellAreaContext,
		defaultCellAreaContextOverrides,
	)
}

func initCellAreaContextClass(gclass unsafe.Pointer, overrides CellAreaContextOverrides, classInitFunc func(*CellAreaContextClass)) {
	if classInitFunc != nil {
		class := (*CellAreaContextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellAreaContext(obj *coreglib.Object) *CellAreaContext {
	return &CellAreaContext{
		Object: obj,
	}
}

func marshalCellAreaContext(p uintptr) (interface{}, error) {
	return wrapCellAreaContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CellAreaContextClass: instance of this type is always passed by reference.
type CellAreaContextClass struct {
	*cellAreaContextClass
}

// cellAreaContextClass is the struct that's finalized.
type cellAreaContextClass struct {
	native unsafe.Pointer
}

var GIRInfoCellAreaContextClass = girepository.MustFind("Gtk", "CellAreaContextClass")
