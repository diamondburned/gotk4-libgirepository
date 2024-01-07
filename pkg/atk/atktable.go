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
// extern void _gotk4_atk1_Table_ConnectRowReordered(gpointer, guintptr);
// extern void _gotk4_atk1_Table_ConnectRowInserted(gpointer, gint, gint, guintptr);
// extern void _gotk4_atk1_Table_ConnectRowDeleted(gpointer, gint, gint, guintptr);
// extern void _gotk4_atk1_Table_ConnectModelChanged(gpointer, guintptr);
// extern void _gotk4_atk1_Table_ConnectColumnReordered(gpointer, guintptr);
// extern void _gotk4_atk1_Table_ConnectColumnInserted(gpointer, gint, gint, guintptr);
// extern void _gotk4_atk1_Table_ConnectColumnDeleted(gpointer, gint, gint, guintptr);
import "C"

// GType values.
var (
	GTypeTable = coreglib.Type(girepository.MustFind("Atk", "Table").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTable, F: marshalTable},
	})
}

// TableOverrider contains methods that are overridable.
type TableOverrider interface {
}

// Table should be implemented by components which present elements ordered via
// rows and columns. It may also be used to present tree-structured information
// if the nodes of the trees can be said to contain multiple "columns".
// Individual elements of an Table are typically referred to as "cells". Those
// cells should implement the interface TableCell, but #Atk doesn't require them
// to be direct children of the current Table. They can be grand-children,
// grand-grand-children etc. Table provides the API needed to get a individual
// cell based on the row and column numbers.
//
// Children of Table are frequently "lightweight" objects, that is, they may not
// have backing widgets in the host UI toolkit. They are therefore often
// transient.
//
// Since tables are often very complex, Table includes provision for offering
// simplified summary information, as well as row and column headers and
// captions. Headers and captions are Objects which may implement other
// interfaces (Text, Image, etc.) as appropriate. Table summaries may themselves
// be (simplified) Tables, etc.
//
// Note for implementors: in the past, Table required that all the cells should
// be direct children of Table, and provided some index based methods to request
// the cells. The practice showed that that forcing made Table implementation
// complex, and hard to expose other kind of children, like rows or captions.
// Right now, index-based methods are deprecated.
//
// Table wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Table struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Table)(nil)
)

// Tabler describes Table's interface methods.
type Tabler interface {
	coreglib.Objector

	baseTable() *Table
}

var _ Tabler = (*Table)(nil)

func ifaceInitTabler(gifacePtr, data C.gpointer) {
}

func wrapTable(obj *coreglib.Object) *Table {
	return &Table{
		Object: obj,
	}
}

func marshalTable(p uintptr) (interface{}, error) {
	return wrapTable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *Table) baseTable() *Table {
	return v
}

// BaseTable returns the underlying base object.
func BaseTable(obj Tabler) *Table {
	return obj.baseTable()
}

// ConnectColumnDeleted: "column-deleted" signal is emitted by an object which
// implements the AtkTable interface when a column is deleted.
func (v *Table) ConnectColumnDeleted(f func(arg1, arg2 int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "column-deleted", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectColumnDeleted), f)
}

// ConnectColumnInserted: "column-inserted" signal is emitted by an object which
// implements the AtkTable interface when a column is inserted.
func (v *Table) ConnectColumnInserted(f func(arg1, arg2 int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "column-inserted", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectColumnInserted), f)
}

// ConnectColumnReordered: "column-reordered" signal is emitted by an object
// which implements the AtkTable interface when the columns are reordered.
func (v *Table) ConnectColumnReordered(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "column-reordered", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectColumnReordered), f)
}

// ConnectModelChanged: "model-changed" signal is emitted by an object which
// implements the AtkTable interface when the model displayed by the table
// changes.
func (v *Table) ConnectModelChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "model-changed", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectModelChanged), f)
}

// ConnectRowDeleted: "row-deleted" signal is emitted by an object which
// implements the AtkTable interface when a row is deleted.
func (v *Table) ConnectRowDeleted(f func(arg1, arg2 int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "row-deleted", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectRowDeleted), f)
}

// ConnectRowInserted: "row-inserted" signal is emitted by an object which
// implements the AtkTable interface when a row is inserted.
func (v *Table) ConnectRowInserted(f func(arg1, arg2 int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "row-inserted", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectRowInserted), f)
}

// ConnectRowReordered: "row-reordered" signal is emitted by an object which
// implements the AtkTable interface when the rows are reordered.
func (v *Table) ConnectRowReordered(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "row-reordered", false, unsafe.Pointer(C._gotk4_atk1_Table_ConnectRowReordered), f)
}

// TableIface: instance of this type is always passed by reference.
type TableIface struct {
	*tableIface
}

// tableIface is the struct that's finalized.
type tableIface struct {
	native unsafe.Pointer
}

var GIRInfoTableIface = girepository.MustFind("Atk", "TableIface")
