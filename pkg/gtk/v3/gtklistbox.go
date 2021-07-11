// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
//
// void gotk4_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBoxxer},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRowwer},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a Model with
// gtk_list_box_bind_model() for each item that gets added to the model.
//
// Versions of GTK+ prior to 3.18 called gtk_widget_show_all() on the rows
// created by the GtkListBoxCreateWidgetFunc, but this forced all widgets inside
// the row to be shown, and is no longer the case. Applications should be
// updated to show the desired row widgets.
type ListBoxCreateWidgetFunc func(item *externglib.Object, userData cgo.Handle) (widget *Widget)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GtkWidget) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData cgo.Handle     // out

	item = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*externglib.Object)
	userData = (cgo.Handle)(arg1)

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item, userData)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// ListBoxFilterFunc: will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row *ListBoxRow, userData cgo.Handle) (ok bool)

//export gotk4_ListBoxFilterFunc
func gotk4_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow     // out
	var userData cgo.Handle // out

	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRow)
	userData = (cgo.Handle)(arg1)

	fn := v.(ListBoxFilterFunc)
	ok := fn(row, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ListBoxForeachFunc: function used by gtk_list_box_selected_foreach(). It will
// be called on every selected child of the @box.
type ListBoxForeachFunc func(box *ListBox, row *ListBoxRow, userData cgo.Handle)

//export gotk4_ListBoxForeachFunc
func gotk4_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box *ListBox        // out
	var row *ListBoxRow     // out
	var userData cgo.Handle // out

	box = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBox)
	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRow)
	userData = (cgo.Handle)(arg2)

	fn := v.(ListBoxForeachFunc)
	fn(box, row, userData)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 *ListBoxRow, row2 *ListBoxRow, userData cgo.Handle) (gint int)

//export gotk4_ListBoxSortFunc
func gotk4_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row1 *ListBoxRow    // out
	var row2 *ListBoxRow    // out
	var userData cgo.Handle // out

	row1 = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRow)
	row2 = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRow)
	userData = (cgo.Handle)(arg2)

	fn := v.(ListBoxSortFunc)
	gint := fn(row1, row2, userData)

	cret = C.gint(gint)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever @row changes or which row is before @row
// changes this is called, which lets you update the header on @row. You may
// remove or set a new one via gtk_list_box_row_set_header() or just change the
// state of the current header widget.
type ListBoxUpdateHeaderFunc func(row *ListBoxRow, before *ListBoxRow, userData cgo.Handle)

//export gotk4_ListBoxUpdateHeaderFunc
func gotk4_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRow     // out
	var before *ListBoxRow  // out
	var userData cgo.Handle // out

	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRow)
	before = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRow)
	userData = (cgo.Handle)(arg2)

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before, userData)
}

// ListBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxOverrider interface {
	//
	ActivateCursorRow()
	//
	RowActivated(row ListBoxRowwer)
	//
	RowSelected(row ListBoxRowwer)
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	//
	SelectedRowsChanged()
	//
	ToggleCursorRow()
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
}

// ListBoxxer describes ListBox's methods.
type ListBoxxer interface {
	// DragHighlightRow: this is a helper function for implementing DnD onto a
	// ListBox.
	DragHighlightRow(row ListBoxRowwer)
	// DragUnhighlightRow: if a row has previously been highlighted via
	// gtk_list_box_drag_highlight_row() it will have the highlight removed.
	DragUnhighlightRow()
	// ActivateOnSingleClick returns whether rows activate on single clicks.
	ActivateOnSingleClick() bool
	// Adjustment gets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	Adjustment() *Adjustment
	// RowAtIndex gets the n-th child in the list (not counting headers).
	RowAtIndex(index_ int) *ListBoxRow
	// RowAtY gets the row at the @y position.
	RowAtY(y int) *ListBoxRow
	// SelectedRow gets the selected row.
	SelectedRow() *ListBoxRow
	// SelectionMode gets the selection mode of the listbox.
	SelectionMode() SelectionMode
	// Insert the @child into the @box at @position.
	Insert(child Widgetter, position int)
	// InvalidateFilter: update the filtering for all rows.
	InvalidateFilter()
	// InvalidateHeaders: update the separators for all rows.
	InvalidateHeaders()
	// InvalidateSort: update the sorting for all rows.
	InvalidateSort()
	// Prepend a widget to the list.
	Prepend(child Widgetter)
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	// SelectRow: make @row the currently selected row.
	SelectRow(row ListBoxRowwer)
	// SelectedForeach calls a function for each selected child.
	SelectedForeach(fn ListBoxForeachFunc)
	// SetActivateOnSingleClick: if @single is true, rows will be activated when
	// you click on them, otherwise you need to double-click.
	SetActivateOnSingleClick(single bool)
	// SetAdjustment sets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	SetAdjustment(adjustment Adjustmenter)
	// SetPlaceholder sets the placeholder widget that is shown in the list when
	// it doesn't display any visible children.
	SetPlaceholder(placeholder Widgetter)
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
	// UnselectRow unselects a single row of @box, if the selection mode allows
	// it.
	UnselectRow(row ListBoxRowwer)
}

// ListBox is a vertical container that contains GtkListBoxRow children. These
// rows can by dynamically sorted and filtered, and headers can be added
// dynamically depending on the row content. It also allows keyboard and mouse
// navigation and selection like a typical list.
//
// Using GtkListBox is often an alternative to TreeView, especially when the
// list contents has a more complicated layout than what is allowed by a
// CellRenderer, or when the contents is interactive (i.e. has a button in it).
//
// Although a ListBox must have only ListBoxRow children you can add any kind of
// widget to it via gtk_container_add(), and a ListBoxRow widget will
// automatically be inserted between the list and the widget.
//
// ListBoxRows can be marked as activatable or selectable. If a row is
// activatable, ListBox::row-activated will be emitted for it when the user
// tries to activate it. If it is selectable, the row will be marked as selected
// when the user tries to select it.
//
// The GtkListBox widget was added in GTK+ 3.10.
//
//
// GtkListBox as GtkBuildable
//
// The GtkListBox implementation of the Buildable interface supports setting a
// child as the placeholder by specifying “placeholder” as the “type” attribute
// of a <child> element. See gtk_list_box_set_placeholder() for info.
//
// CSS nodes
//
//    list
//    ╰── row[.activatable]
//
// GtkListBox uses a single CSS node named list. Each GtkListBoxRow uses a
// single CSS node named row. The row nodes get the .activatable style class
// added when appropriate.
type ListBox struct {
	Container
}

var (
	_ ListBoxxer      = (*ListBox)(nil)
	_ gextras.Nativer = (*ListBox)(nil)
)

func wrapListBox(obj *externglib.Object) ListBoxxer {
	return &ListBox{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalListBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBox(obj), nil
}

// NewListBox creates a new ListBox container.
func NewListBox() *ListBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_new()

	var _listBox *ListBox // out

	_listBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBox)

	return _listBox
}

// DragHighlightRow: this is a helper function for implementing DnD onto a
// ListBox. The passed in @row will be highlighted via gtk_drag_highlight(), and
// any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave event.
func (box *ListBox) DragHighlightRow(row ListBoxRowwer) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row() it will have the highlight removed.
func (box *ListBox) DragUnhighlightRow() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (box *ListBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Adjustment gets the adjustment (if any) that the widget uses to for vertical
// scrolling.
func (box *ListBox) Adjustment() *Adjustment {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_adjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// RowAtIndex gets the n-th child in the list (not counting headers). If @_index
// is negative or larger than the number of items in the list, nil is returned.
func (box *ListBox) RowAtIndex(index_ int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.gint           // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(index_)

	_cret = C.gtk_list_box_get_row_at_index(_arg0, _arg1)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRow)

	return _listBoxRow
}

// RowAtY gets the row at the @y position.
func (box *ListBox) RowAtY(y int) *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.gint           // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(y)

	_cret = C.gtk_list_box_get_row_at_y(_arg0, _arg1)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRow)

	return _listBoxRow
}

// SelectedRow gets the selected row.
//
// Note that the box may allow multiple selection, in which case you should use
// gtk_list_box_selected_foreach() to find all selected rows.
func (box *ListBox) SelectedRow() *ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selected_row(_arg0)

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRow)

	return _listBoxRow
}

// SelectionMode gets the selection mode of the listbox.
func (box *ListBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkListBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Insert the @child into the @box at @position. If a sort function is set, the
// widget will actually be inserted at the calculated position and this function
// has the same effect of gtk_container_add().
//
// If @position is -1, or larger than the total number of items in the @box,
// then the @child will be appended to the end.
func (box *ListBox) Insert(child Widgetter, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = C.gint(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
}

// InvalidateFilter: update the filtering for all rows. Call this when result of
// the filter function on the @box is changed due to an external factor. For
// instance, this would be used if the filter function just looked for a
// specific search string and the entry with the search string has changed.
func (box *ListBox) InvalidateFilter() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
}

// InvalidateHeaders: update the separators for all rows. Call this when result
// of the header function on the @box is changed due to an external factor.
func (box *ListBox) InvalidateHeaders() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
}

// InvalidateSort: update the sorting for all rows. Call this when result of the
// sort function on the @box is changed due to an external factor.
func (box *ListBox) InvalidateSort() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
}

// Prepend a widget to the list. If a sort function is set, the widget will
// actually be inserted at the calculated position and this function has the
// same effect of gtk_container_add().
func (box *ListBox) Prepend(child Widgetter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (box *ListBox) SelectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_select_all(_arg0)
}

// SelectRow: make @row the currently selected row.
func (box *ListBox) SelectRow(row ListBoxRowwer) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_select_row(_arg0, _arg1)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (box *ListBox) SelectedForeach(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ListBoxForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
}

// SetActivateOnSingleClick: if @single is true, rows will be activated when you
// click on them, otherwise you need to double-click.
func (box *ListBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling. For instance, this is used to get the page size for
// PageUp/Down key handling.
//
// In the normal case when the @box is packed inside a ScrolledWindow the
// adjustment from that will be picked up automatically, so there is no need to
// manually do that.
func (box *ListBox) SetAdjustment(adjustment Adjustmenter) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((adjustment).(gextras.Nativer).Native()))

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when it
// doesn't display any visible children.
func (box *ListBox) SetPlaceholder(placeholder Widgetter) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((placeholder).(gextras.Nativer).Native()))

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
}

// UnselectAll: unselect all children of @box, if the selection mode allows it.
func (box *ListBox) UnselectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_unselect_all(_arg0)
}

// UnselectRow unselects a single row of @box, if the selection mode allows it.
func (box *ListBox) UnselectRow(row ListBoxRowwer) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer((row).(gextras.Nativer).Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
}

// ListBoxRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxRowOverrider interface {
	//
	Activate()
}

// ListBoxRowwer describes ListBoxRow's methods.
type ListBoxRowwer interface {
	// Changed marks @row as changed, causing any state that depends on this to
	// be updated.
	Changed()
	// Activatable gets the value of the ListBoxRow:activatable property for
	// this row.
	Activatable() bool
	// Header returns the current header of the @row.
	Header() *Widget
	// Index gets the current index of the @row in its ListBox container.
	Index() int
	// Selectable gets the value of the ListBoxRow:selectable property for this
	// row.
	Selectable() bool
	// IsSelected returns whether the child is currently selected in its ListBox
	// container.
	IsSelected() bool
	// SetActivatable: set the ListBoxRow:activatable property for this row.
	SetActivatable(activatable bool)
	// SetHeader sets the current header of the @row.
	SetHeader(header Widgetter)
	// SetSelectable: set the ListBoxRow:selectable property for this row.
	SetSelectable(selectable bool)
}

//
type ListBoxRow struct {
	Bin

	Actionable
}

var (
	_ ListBoxRowwer   = (*ListBoxRow)(nil)
	_ gextras.Nativer = (*ListBoxRow)(nil)
)

func wrapListBoxRow(obj *externglib.Object) ListBoxRowwer {
	return &ListBoxRow{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalListBoxRowwer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBoxRow(obj), nil
}

// NewListBoxRow creates a new ListBoxRow, to be used as a child of a ListBox.
func NewListBoxRow() *ListBoxRow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_row_new()

	var _listBoxRow *ListBoxRow // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRow)

	return _listBoxRow
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ListBoxRow) Native() uintptr {
	return v.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// Changed marks @row as changed, causing any state that depends on this to be
// updated. This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data used for the row
// functions. For instance, if the list is mirroring some external data set, and
// *two* rows changed in the external data set then when you call
// gtk_list_box_row_changed() on the first row the sort function must only read
// the new data for the first of the two changed rows, otherwise the resorting
// of the rows will be wrong.
//
// This generally means that if you don’t fully control the data model you have
// to duplicate the data that affects the listbox row functions into the row
// widgets themselves. Another alternative is to call
// gtk_list_box_invalidate_sort() on any model change, but that is more
// expensive.
func (row *ListBoxRow) Changed() {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

// Activatable gets the value of the ListBoxRow:activatable property for this
// row.
func (row *ListBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Header returns the current header of the @row. This can be used in a
// ListBoxUpdateHeaderFunc to see if there is a header set already, and if so to
// update the state of it.
func (row *ListBoxRow) Header() *Widget {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_header(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// Index gets the current index of the @row in its ListBox container.
func (row *ListBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Selectable gets the value of the ListBoxRow:selectable property for this row.
func (row *ListBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its ListBox
// container.
func (row *ListBoxRow) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable: set the ListBoxRow:activatable property for this row.
func (row *ListBoxRow) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

// SetHeader sets the current header of the @row. This is only allowed to be
// called from a ListBoxUpdateHeaderFunc. It will replace any existing header in
// the row, and be shown in front of the row in the listbox.
func (row *ListBoxRow) SetHeader(header Widgetter) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((header).(gextras.Nativer).Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

// SetSelectable: set the ListBoxRow:selectable property for this row.
func (row *ListBoxRow) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}
