// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_column_view_get_type()), F: marshalColumnView},
	})
}

// ColumnView: `GtkColumnView` presents a large dynamic list of items using
// multiple columns with headers.
//
// `GtkColumnView` uses the factories of its columns to generate a cell widget
// for each column, for each visible item and displays them together as the row
// for this item.
//
// The [property@Gtk.ColumnView:show-row-separators] and
// [propertyGtk.ColumnView:show-column-separators] properties offer a simple way
// to display separators between the rows or columns.
//
// `GtkColumnView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// [property@Gtk.ColumnView:enable-rubberband].
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the `GtkSorter` returned by
// [method@Gtk.ColumnView.get_sorter] must be attached to a sort model for the
// data that the view is showing, and the columns must have sorters attached to
// them by calling [method@Gtk.ColumnViewColumn.set_sorter]. The initial sort
// order can be set with [method@Gtk.ColumnView.sort_by_column].
//
// The column view also supports interactive resizing and reordering of columns,
// via Drag-and-Drop of the column headers. This can be enabled or disabled with
// the [property@Gtk.ColumnView:reorderable] and
// [property@Gtk.ColumnViewColumn:resizable] properties.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
//
// CSS nodes
//
// “`
// columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
// ├── header │ ├── <column header> ┊ ┊ │ ╰── <column header> │ ├── listview │ ┊
// ╰── [rubberband] “`
//
// `GtkColumnView` uses a single CSS node named columnview. It may carry the
// .column-separators style class, when
// [property@Gtk.ColumnView:show-column-separators] property is set. Header
// widgets appear below a node with name header. The rows are contained in a
// `GtkListView` widget, so there is a listview node with the same structure as
// for a standalone `GtkListView` widget. If
// [property@Gtk.ColumnView:show-row-separators] is set, it will be passed on to
// the list view, causing its CSS node to carry the .separators style class. For
// rubberband selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select the style of
// list presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkColumnView` uses the GTK_ACCESSIBLE_ROLE_TREE_GRID role, header title
// widgets are using the GTK_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets
// are using the GTK_ACCESSIBLE_ROLE_ROW role, and individual cells are using
// the GTK_ACCESSIBLE_ROLE_GRID_CELL role
type ColumnView interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable

	// AppendColumn appends the @column to the end of the columns in @self.
	AppendColumn(column ColumnViewColumn)
	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband() bool
	// Reorderable returns whether columns are reorderable.
	Reorderable() bool
	// ShowColumnSeparators returns whether the list should show separators
	// between columns.
	ShowColumnSeparators() bool
	// ShowRowSeparators returns whether the list should show separators between
	// rows.
	ShowRowSeparators() bool
	// SingleClickActivate returns whether rows will be activated on single
	// click and selected on hover.
	SingleClickActivate() bool
	// InsertColumn inserts a column at the given position in the columns of
	// @self.
	//
	// If @column is already a column of @self, it will be repositioned.
	InsertColumn(position uint, column ColumnViewColumn)
	// RemoveColumn removes the @column from the list of columns of @self.
	RemoveColumn(column ColumnViewColumn)
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(enableRubberband bool)
	// SetModel sets the model to use.
	//
	// This must be a [iface@Gtk.SelectionModel].
	SetModel(model SelectionModel)
	// SetReorderable sets whether columns should be reorderable by dragging.
	SetReorderable(reorderable bool)
	// SetShowColumnSeparators sets whether the list should show separators
	// between columns.
	SetShowColumnSeparators(showColumnSeparators bool)
	// SetShowRowSeparators sets whether the list should show separators between
	// rows.
	SetShowRowSeparators(showRowSeparators bool)
	// SetSingleClickActivate sets whether rows should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(singleClickActivate bool)
	// SortByColumn sets the sorting of the view.
	//
	// This function should be used to set up the initial sorting. At runtime,
	// users can change the sorting of a column view by clicking on the list
	// headers.
	//
	// This call only has an effect if the sorter returned by
	// [method@Gtk.ColumnView.get_sorter] is set on a sort model, and
	// [method@Gtk.ColumnViewColumn.set_sorter] has been called on @column to
	// associate a sorter with the column.
	//
	// If @column is nil, the view will be unsorted.
	SortByColumn(column ColumnViewColumn, direction SortType)
}

// columnView implements the ColumnView class.
type columnView struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable
}

var _ ColumnView = (*columnView)(nil)

// WrapColumnView wraps a GObject to the right type. It is
// primarily used internally.
func WrapColumnView(obj *externglib.Object) ColumnView {
	return columnView{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Scrollable:       WrapScrollable(obj),
	}
}

func marshalColumnView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColumnView(obj), nil
}

// AppendColumn appends the @column to the end of the columns in @self.
func (s columnView) AppendColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_append_column(_arg0, _arg1)
}

// EnableRubberband returns whether rows can be selected by dragging with
// the mouse.
func (s columnView) EnableRubberband() bool {
	var _arg0 *C.GtkColumnView // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_get_enable_rubberband(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reorderable returns whether columns are reorderable.
func (s columnView) Reorderable() bool {
	var _arg0 *C.GtkColumnView // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowColumnSeparators returns whether the list should show separators
// between columns.
func (s columnView) ShowColumnSeparators() bool {
	var _arg0 *C.GtkColumnView // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_get_show_column_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRowSeparators returns whether the list should show separators between
// rows.
func (s columnView) ShowRowSeparators() bool {
	var _arg0 *C.GtkColumnView // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_get_show_row_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SingleClickActivate returns whether rows will be activated on single
// click and selected on hover.
func (s columnView) SingleClickActivate() bool {
	var _arg0 *C.GtkColumnView // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_get_single_click_activate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InsertColumn inserts a column at the given position in the columns of
// @self.
//
// If @column is already a column of @self, it will be repositioned.
func (s columnView) InsertColumn(position uint, column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 C.guint                // out
	var _arg2 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_insert_column(_arg0, _arg1, _arg2)
}

// RemoveColumn removes the @column from the list of columns of @self.
func (s columnView) RemoveColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_remove_column(_arg0, _arg1)
}

// SetEnableRubberband sets whether selections can be changed by dragging
// with the mouse.
func (s columnView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_enable_rubberband(_arg0, _arg1)
}

// SetModel sets the model to use.
//
// This must be a [iface@Gtk.SelectionModel].
func (s columnView) SetModel(model SelectionModel) {
	var _arg0 *C.GtkColumnView     // out
	var _arg1 *C.GtkSelectionModel // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_column_view_set_model(_arg0, _arg1)
}

// SetReorderable sets whether columns should be reorderable by dragging.
func (s columnView) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_reorderable(_arg0, _arg1)
}

// SetShowColumnSeparators sets whether the list should show separators
// between columns.
func (s columnView) SetShowColumnSeparators(showColumnSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showColumnSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_column_separators(_arg0, _arg1)
}

// SetShowRowSeparators sets whether the list should show separators between
// rows.
func (s columnView) SetShowRowSeparators(showRowSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showRowSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_row_separators(_arg0, _arg1)
}

// SetSingleClickActivate sets whether rows should be activated on single
// click and selected on hover.
func (s columnView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_single_click_activate(_arg0, _arg1)
}

// SortByColumn sets the sorting of the view.
//
// This function should be used to set up the initial sorting. At runtime,
// users can change the sorting of a column view by clicking on the list
// headers.
//
// This call only has an effect if the sorter returned by
// [method@Gtk.ColumnView.get_sorter] is set on a sort model, and
// [method@Gtk.ColumnViewColumn.set_sorter] has been called on @column to
// associate a sorter with the column.
//
// If @column is nil, the view will be unsorted.
func (s columnView) SortByColumn(column ColumnViewColumn, direction SortType) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out
	var _arg2 C.GtkSortType          // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))
	_arg2 = (C.GtkSortType)(direction)

	C.gtk_column_view_sort_by_column(_arg0, _arg1, _arg2)
}
