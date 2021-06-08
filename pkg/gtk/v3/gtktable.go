// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_table_get_type()), F: marshalTable},
	})
}

// Table: the Table functions allow the programmer to arrange widgets in rows
// and columns, making it easy to align many widgets next to each other,
// horizontally and vertically.
//
// Tables are created with a call to gtk_table_new(), the size of which can
// later be changed with gtk_table_resize().
//
// Widgets can be added to a table using gtk_table_attach() or the more
// convenient (but slightly less flexible) gtk_table_attach_defaults().
//
// To alter the space next to a specific row, use gtk_table_set_row_spacing(),
// and for a column, gtk_table_set_col_spacing(). The gaps between all rows or
// columns can be changed by calling gtk_table_set_row_spacings() or
// gtk_table_set_col_spacings() respectively. Note that spacing is added between
// the children, while padding added by gtk_table_attach() is added on either
// side of the widget it belongs to.
//
// gtk_table_set_homogeneous(), can be used to set whether all cells in the
// table will resize themselves to the size of the largest widget in the table.
//
// > Table has been deprecated. Use Grid instead. It provides the same >
// capabilities as GtkTable for arranging widgets in a rectangular grid, but >
// does support height-for-width geometry management.
type Table interface {
	Container
	Buildable

	// Attach adds a widget to a table. The number of “cells” that a widget will
	// occupy is specified by @left_attach, @right_attach, @top_attach and
	// @bottom_attach. These each represent the leftmost, rightmost, uppermost
	// and lowest column and row numbers of the table. (Columns and rows are
	// indexed from zero).
	//
	// To make a button occupy the lower right cell of a 2x2 table, use
	//
	//    gtk_table_attach (table, button,
	//                      1, 2, // left, right attach
	//                      1, 2, // top, bottom attach
	//                      xoptions, yoptions,
	//                      xpadding, ypadding);
	//
	// If you want to make the button span the entire bottom row, use
	// @left_attach == 0 and @right_attach = 2 instead.
	Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint)
	// AttachDefaults as there are many options associated with
	// gtk_table_attach(), this convenience function provides the programmer
	// with a means to add children to a table with identical padding and
	// expansion options. The values used for the AttachOptions are `GTK_EXPAND
	// | GTK_FILL`, and the padding is set to 0.
	AttachDefaults(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint)
	// ColSpacing gets the amount of space between column @col, and column @col
	// + 1. See gtk_table_set_col_spacing().
	ColSpacing(column uint) uint
	// DefaultColSpacing gets the default column spacing for the table. This is
	// the spacing that will be used for newly added columns. (See
	// gtk_table_set_col_spacings())
	DefaultColSpacing() uint
	// DefaultRowSpacing gets the default row spacing for the table. This is the
	// spacing that will be used for newly added rows. (See
	// gtk_table_set_row_spacings())
	DefaultRowSpacing() uint
	// Homogeneous returns whether the table cells are all constrained to the
	// same width and height. (See gtk_table_set_homogeneous ())
	Homogeneous() bool
	// RowSpacing gets the amount of space between row @row, and row @row + 1.
	// See gtk_table_set_row_spacing().
	RowSpacing(row uint) uint
	// Size gets the number of rows and columns in the table.
	Size() (rows uint, columns uint)
	// Resize: if you need to change a table’s size after it has been created,
	// this function allows you to do so.
	Resize(rows uint, columns uint)
	// SetColSpacing alters the amount of space between a given table column and
	// the following column.
	SetColSpacing(column uint, spacing uint)
	// SetColSpacings sets the space between every column in @table equal to
	// @spacing.
	SetColSpacings(spacing uint)
	// SetHomogeneous changes the homogenous property of table cells, ie.
	// whether all cells are an equal size or not.
	SetHomogeneous(homogeneous bool)
	// SetRowSpacing changes the space between a given table row and the
	// subsequent row.
	SetRowSpacing(row uint, spacing uint)
	// SetRowSpacings sets the space between every row in @table equal to
	// @spacing.
	SetRowSpacings(spacing uint)
}

// table implements the Table interface.
type table struct {
	Container
	Buildable
}

var _ Table = (*table)(nil)

// WrapTable wraps a GObject to the right type. It is
// primarily used internally.
func WrapTable(obj *externglib.Object) Table {
	return Table{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTable(obj), nil
}

// NewTable constructs a class Table.
func NewTable(rows uint, columns uint, homogeneous bool) Table {
	var arg1 C.guint
	var arg2 C.guint
	var arg3 C.gboolean

	arg1 = C.guint(rows)
	arg2 = C.guint(columns)
	if homogeneous {
		arg3 = C.gboolean(1)
	}

	var cret C.GtkTable
	var goret Table

	cret = C.gtk_table_new(arg1, arg2, arg3)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Table)

	return goret
}

// Attach adds a widget to a table. The number of “cells” that a widget will
// occupy is specified by @left_attach, @right_attach, @top_attach and
// @bottom_attach. These each represent the leftmost, rightmost, uppermost
// and lowest column and row numbers of the table. (Columns and rows are
// indexed from zero).
//
// To make a button occupy the lower right cell of a 2x2 table, use
//
//    gtk_table_attach (table, button,
//                      1, 2, // left, right attach
//                      1, 2, // top, bottom attach
//                      xoptions, yoptions,
//                      xpadding, ypadding);
//
// If you want to make the button span the entire bottom row, use
// @left_attach == 0 and @right_attach = 2 instead.
func (t table) Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	var arg0 *C.GtkTable
	var arg1 *C.GtkWidget
	var arg2 C.guint
	var arg3 C.guint
	var arg4 C.guint
	var arg5 C.guint
	var arg6 C.GtkAttachOptions
	var arg7 C.GtkAttachOptions
	var arg8 C.guint
	var arg9 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = C.guint(leftAttach)
	arg3 = C.guint(rightAttach)
	arg4 = C.guint(topAttach)
	arg5 = C.guint(bottomAttach)
	arg6 = (C.GtkAttachOptions)(xoptions)
	arg7 = (C.GtkAttachOptions)(yoptions)
	arg8 = C.guint(xpadding)
	arg9 = C.guint(ypadding)

	C.gtk_table_attach(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// AttachDefaults as there are many options associated with
// gtk_table_attach(), this convenience function provides the programmer
// with a means to add children to a table with identical padding and
// expansion options. The values used for the AttachOptions are `GTK_EXPAND
// | GTK_FILL`, and the padding is set to 0.
func (t table) AttachDefaults(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var arg0 *C.GtkTable
	var arg1 *C.GtkWidget
	var arg2 C.guint
	var arg3 C.guint
	var arg4 C.guint
	var arg5 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.guint(leftAttach)
	arg3 = C.guint(rightAttach)
	arg4 = C.guint(topAttach)
	arg5 = C.guint(bottomAttach)

	C.gtk_table_attach_defaults(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ColSpacing gets the amount of space between column @col, and column @col
// + 1. See gtk_table_set_col_spacing().
func (t table) ColSpacing(column uint) uint {
	var arg0 *C.GtkTable
	var arg1 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(column)

	var cret C.guint
	var goret uint

	cret = C.gtk_table_get_col_spacing(arg0, arg1)

	goret = uint(cret)

	return goret
}

// DefaultColSpacing gets the default column spacing for the table. This is
// the spacing that will be used for newly added columns. (See
// gtk_table_set_col_spacings())
func (t table) DefaultColSpacing() uint {
	var arg0 *C.GtkTable

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_table_get_default_col_spacing(arg0)

	goret = uint(cret)

	return goret
}

// DefaultRowSpacing gets the default row spacing for the table. This is the
// spacing that will be used for newly added rows. (See
// gtk_table_set_row_spacings())
func (t table) DefaultRowSpacing() uint {
	var arg0 *C.GtkTable

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_table_get_default_row_spacing(arg0)

	goret = uint(cret)

	return goret
}

// Homogeneous returns whether the table cells are all constrained to the
// same width and height. (See gtk_table_set_homogeneous ())
func (t table) Homogeneous() bool {
	var arg0 *C.GtkTable

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_table_get_homogeneous(arg0)

	if cret {
		goret = true
	}

	return goret
}

// RowSpacing gets the amount of space between row @row, and row @row + 1.
// See gtk_table_set_row_spacing().
func (t table) RowSpacing(row uint) uint {
	var arg0 *C.GtkTable
	var arg1 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(row)

	var cret C.guint
	var goret uint

	cret = C.gtk_table_get_row_spacing(arg0, arg1)

	goret = uint(cret)

	return goret
}

// Size gets the number of rows and columns in the table.
func (t table) Size() (rows uint, columns uint) {
	var arg0 *C.GtkTable

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	arg1 := new(C.guint)
	var ret1 uint
	arg2 := new(C.guint)
	var ret2 uint

	C.gtk_table_get_size(arg0, arg1, arg2)

	ret1 = uint(*arg1)
	ret2 = uint(*arg2)

	return ret1, ret2
}

// Resize: if you need to change a table’s size after it has been created,
// this function allows you to do so.
func (t table) Resize(rows uint, columns uint) {
	var arg0 *C.GtkTable
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(rows)
	arg2 = C.guint(columns)

	C.gtk_table_resize(arg0, arg1, arg2)
}

// SetColSpacing alters the amount of space between a given table column and
// the following column.
func (t table) SetColSpacing(column uint, spacing uint) {
	var arg0 *C.GtkTable
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(column)
	arg2 = C.guint(spacing)

	C.gtk_table_set_col_spacing(arg0, arg1, arg2)
}

// SetColSpacings sets the space between every column in @table equal to
// @spacing.
func (t table) SetColSpacings(spacing uint) {
	var arg0 *C.GtkTable
	var arg1 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(spacing)

	C.gtk_table_set_col_spacings(arg0, arg1)
}

// SetHomogeneous changes the homogenous property of table cells, ie.
// whether all cells are an equal size or not.
func (t table) SetHomogeneous(homogeneous bool) {
	var arg0 *C.GtkTable
	var arg1 C.gboolean

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	if homogeneous {
		arg1 = C.gboolean(1)
	}

	C.gtk_table_set_homogeneous(arg0, arg1)
}

// SetRowSpacing changes the space between a given table row and the
// subsequent row.
func (t table) SetRowSpacing(row uint, spacing uint) {
	var arg0 *C.GtkTable
	var arg1 C.guint
	var arg2 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(row)
	arg2 = C.guint(spacing)

	C.gtk_table_set_row_spacing(arg0, arg1, arg2)
}

// SetRowSpacings sets the space between every row in @table equal to
// @spacing.
func (t table) SetRowSpacings(spacing uint) {
	var arg0 *C.GtkTable
	var arg1 C.guint

	arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	arg1 = C.guint(spacing)

	C.gtk_table_set_row_spacings(arg0, arg1)
}

type TableChild struct {
	native C.GtkTableChild
}

// WrapTableChild wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableChild(ptr unsafe.Pointer) *TableChild {
	if ptr == nil {
		return nil
	}

	return (*TableChild)(ptr)
}

func marshalTableChild(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTableChild(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TableChild) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Widget gets the field inside the struct.
func (t *TableChild) Widget() Widget {
	var v Widget
	v = gextras.CastObject(externglib.Take(unsafe.Pointer(t.native.widget.Native()))).(Widget)
	return v
}

// LeftAttach gets the field inside the struct.
func (t *TableChild) LeftAttach() uint16 {
	var v uint16
	v = uint16(t.native.left_attach)
	return v
}

// RightAttach gets the field inside the struct.
func (t *TableChild) RightAttach() uint16 {
	var v uint16
	v = uint16(t.native.right_attach)
	return v
}

// TopAttach gets the field inside the struct.
func (t *TableChild) TopAttach() uint16 {
	var v uint16
	v = uint16(t.native.top_attach)
	return v
}

// BottomAttach gets the field inside the struct.
func (t *TableChild) BottomAttach() uint16 {
	var v uint16
	v = uint16(t.native.bottom_attach)
	return v
}

// Xpadding gets the field inside the struct.
func (t *TableChild) Xpadding() uint16 {
	var v uint16
	v = uint16(t.native.xpadding)
	return v
}

// Ypadding gets the field inside the struct.
func (t *TableChild) Ypadding() uint16 {
	var v uint16
	v = uint16(t.native.ypadding)
	return v
}

type TableRowCol struct {
	native C.GtkTableRowCol
}

// WrapTableRowCol wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableRowCol(ptr unsafe.Pointer) *TableRowCol {
	if ptr == nil {
		return nil
	}

	return (*TableRowCol)(ptr)
}

func marshalTableRowCol(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTableRowCol(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TableRowCol) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Requisition gets the field inside the struct.
func (t *TableRowCol) Requisition() uint16 {
	var v uint16
	v = uint16(t.native.requisition)
	return v
}

// Allocation gets the field inside the struct.
func (t *TableRowCol) Allocation() uint16 {
	var v uint16
	v = uint16(t.native.allocation)
	return v
}

// Spacing gets the field inside the struct.
func (t *TableRowCol) Spacing() uint16 {
	var v uint16
	v = uint16(t.native.spacing)
	return v
}
