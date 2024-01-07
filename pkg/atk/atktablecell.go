// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

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
	GTypeTableCell = coreglib.Type(girepository.MustFind("Atk", "TableCell").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTableCell, F: marshalTableCell},
	})
}

// TableCellOverrider contains methods that are overridable.
type TableCellOverrider interface {
}

// TableCell: being Table a component which present elements ordered via rows
// and columns, an TableCell is the interface which each of those elements, so
// "cells" should implement.
//
// See also Table.
//
// TableCell wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TableCell struct {
	_ [0]func() // equal guard
	AtkObject
}

var (
	_ coreglib.Objector = (*TableCell)(nil)
)

// TableCeller describes TableCell's interface methods.
type TableCeller interface {
	coreglib.Objector

	baseTableCell() *TableCell
}

var _ TableCeller = (*TableCell)(nil)

func ifaceInitTableCeller(gifacePtr, data C.gpointer) {
}

func wrapTableCell(obj *coreglib.Object) *TableCell {
	return &TableCell{
		AtkObject: AtkObject{
			Object: obj,
		},
	}
}

func marshalTableCell(p uintptr) (interface{}, error) {
	return wrapTableCell(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TableCell) baseTableCell() *TableCell {
	return v
}

// BaseTableCell returns the underlying base object.
func BaseTableCell(obj TableCeller) *TableCell {
	return obj.baseTableCell()
}
