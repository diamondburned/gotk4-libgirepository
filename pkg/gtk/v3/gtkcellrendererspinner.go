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
	GTypeCellRendererSpinner = coreglib.Type(girepository.MustFind("Gtk", "CellRendererSpinner").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererSpinner, F: marshalCellRendererSpinner},
	})
}

// CellRendererSpinnerOverrides contains methods that are overridable.
type CellRendererSpinnerOverrides struct {
}

func defaultCellRendererSpinnerOverrides(v *CellRendererSpinner) CellRendererSpinnerOverrides {
	return CellRendererSpinnerOverrides{}
}

// CellRendererSpinner renders a spinning animation in a cell, very similar to
// Spinner. It can often be used as an alternative to a CellRendererProgress for
// displaying indefinite activity, instead of actual progress.
//
// To start the animation in a cell, set the CellRendererSpinner:active property
// to TRUE and increment the CellRendererSpinner:pulse property at regular
// intervals. The usual way to set the cell renderer properties for each cell is
// to bind them to columns in your tree model using e.g.
// gtk_tree_view_column_add_attribute().
type CellRendererSpinner struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererSpinner)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CellRendererSpinner, *CellRendererSpinnerClass, CellRendererSpinnerOverrides](
		GTypeCellRendererSpinner,
		initCellRendererSpinnerClass,
		wrapCellRendererSpinner,
		defaultCellRendererSpinnerOverrides,
	)
}

func initCellRendererSpinnerClass(gclass unsafe.Pointer, overrides CellRendererSpinnerOverrides, classInitFunc func(*CellRendererSpinnerClass)) {
	if classInitFunc != nil {
		class := (*CellRendererSpinnerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellRendererSpinner(obj *coreglib.Object) *CellRendererSpinner {
	return &CellRendererSpinner{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	return wrapCellRendererSpinner(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CellRendererSpinnerClass: instance of this type is always passed by
// reference.
type CellRendererSpinnerClass struct {
	*cellRendererSpinnerClass
}

// cellRendererSpinnerClass is the struct that's finalized.
type cellRendererSpinnerClass struct {
	native unsafe.Pointer
}

var GIRInfoCellRendererSpinnerClass = girepository.MustFind("Gtk", "CellRendererSpinnerClass")
