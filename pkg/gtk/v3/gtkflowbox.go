// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_flow_box_get_type()), F: marshalFlowBox},
		{T: externglib.Type(C.gtk_flow_box_child_get_type()), F: marshalFlowBoxChild},
	})
}

// FlowBoxCreateWidgetFunc: called for flow boxes that are bound to a Model with
// gtk_flow_box_bind_model() for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func() (widget Widget)

//export gotk4_FlowBoxCreateWidgetFunc
func gotk4_FlowBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxCreateWidgetFunc)
	widget := fn()

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
}

// FlowBoxFilterFunc: a function that will be called whenrever a child changes
// or is added. It lets you control if the child should be visible or not.
type FlowBoxFilterFunc func() (ok bool)

//export gotk4_FlowBoxFilterFunc
func gotk4_FlowBoxFilterFunc(arg0 *C.GtkFlowBoxChild, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxFilterFunc)
	ok := fn()

	if ok {
		cret = C.gboolean(1)
	}
}

// FlowBoxForeachFunc: a function used by gtk_flow_box_selected_foreach(). It
// will be called on every selected child of the @box.
type FlowBoxForeachFunc func()

//export gotk4_FlowBoxForeachFunc
func gotk4_FlowBoxForeachFunc(arg0 *C.GtkFlowBox, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxForeachFunc)
	fn()
}

// FlowBoxSortFunc: a function to compare two children to determine which should
// come first.
type FlowBoxSortFunc func() (gint int)

//export gotk4_FlowBoxSortFunc
func gotk4_FlowBoxSortFunc(arg0 *C.GtkFlowBoxChild, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) C.gint {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxSortFunc)
	gint := fn()

	cret = C.gint(gint)
}

// FlowBox: a GtkFlowBox positions child widgets in sequence according to its
// orientation.
//
// For instance, with the horizontal orientation, the widgets will be arranged
// from left to right, starting a new row under the previous row when necessary.
// Reducing the width in this case will require more rows, so a larger height
// will be requested.
//
// Likewise, with the vertical orientation, the widgets will be arranged from
// top to bottom, starting a new column to the right when necessary. Reducing
// the height will require more columns, so a larger width will be requested.
//
// The size request of a GtkFlowBox alone may not be what you expect; if you
// need to be able to shrink it along both axes and dynamically reflow its
// children, you may have to wrap it in a ScrolledWindow to enable that.
//
// The children of a GtkFlowBox can be dynamically sorted and filtered.
//
// Although a GtkFlowBox must have only FlowBoxChild children, you can add any
// kind of widget to it via gtk_container_add(), and a GtkFlowBoxChild widget
// will automatically be inserted between the box and the widget.
//
// Also see ListBox.
//
// GtkFlowBox was added in GTK+ 3.12.
//
// CSS nodes
//
//    flowbox
//    ├── flowboxchild
//    │   ╰── <child>
//    ├── flowboxchild
//    │   ╰── <child>
//    ┊
//    ╰── [rubberband]
//
// GtkFlowBox uses a single CSS node with name flowbox. GtkFlowBoxChild uses a
// single CSS node with name flowboxchild. For rubberband selection, a subnode
// with name rubberband is used.
type FlowBox interface {
	Container
	Buildable
	Orientable

	// BindModel binds @model to @box.
	//
	// If @box was already bound to a model, that previous binding is destroyed.
	//
	// The contents of @box are cleared and then filled with widgets that
	// represent items from @model. @box is updated whenever @model changes. If
	// @model is nil, @box is left empty.
	//
	// It is undefined to add or remove widgets directly (for example, with
	// gtk_flow_box_insert() or gtk_container_add()) while @box is bound to a
	// model.
	//
	// Note that using a model is incompatible with the filtering and sorting
	// functionality in GtkFlowBox. When using a model, filtering and sorting
	// should be implemented by the model.
	BindModel()
	// ActivateOnSingleClick returns whether children activate on single clicks.
	ActivateOnSingleClick() bool
	// ChildAtIndex gets the nth child in the @box.
	ChildAtIndex(idx int) FlowBoxChild
	// ChildAtPos gets the child in the (@x, @y) position.
	ChildAtPos(x int, y int) FlowBoxChild
	// ColumnSpacing gets the horizontal spacing.
	ColumnSpacing() uint
	// Homogeneous returns whether the box is homogeneous (all children are the
	// same size). See gtk_box_set_homogeneous().
	Homogeneous() bool
	// MaxChildrenPerLine gets the maximum number of children per line.
	MaxChildrenPerLine() uint
	// MinChildrenPerLine gets the minimum number of children per line.
	MinChildrenPerLine() uint
	// RowSpacing gets the vertical spacing.
	RowSpacing() uint
	// SelectedChildren creates a list of all selected children.
	SelectedChildren() *glib.List
	// SelectionMode gets the selection mode of @box.
	SelectionMode() SelectionMode
	// Insert inserts the @widget into @box at @position.
	//
	// If a sort function is set, the widget will actually be inserted at the
	// calculated position and this function has the same effect as
	// gtk_container_add().
	//
	// If @position is -1, or larger than the total number of children in the
	// @box, then the @widget will be appended to the end.
	Insert(widget Widget, position int)
	// InvalidateFilter updates the filtering for all children.
	//
	// Call this function when the result of the filter function on the @box is
	// changed due ot an external factor. For instance, this would be used if
	// the filter function just looked for a specific search term, and the entry
	// with the string has changed.
	InvalidateFilter()
	// InvalidateSort updates the sorting for all children.
	//
	// Call this when the result of the sort function on @box is changed due to
	// an external factor.
	InvalidateSort()
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	// SelectChild selects a single child of @box, if the selection mode allows
	// it.
	SelectChild(child FlowBoxChild)
	// SelectedForeach calls a function for each selected child.
	//
	// Note that the selection cannot be modified from within this function.
	SelectedForeach()
	// SetActivateOnSingleClick: if @single is true, children will be activated
	// when you click on them, otherwise you need to double-click.
	SetActivateOnSingleClick(single bool)
	// SetColumnSpacing sets the horizontal space to add between children. See
	// the FlowBox:column-spacing property.
	SetColumnSpacing(spacing uint)
	// SetFilterFunc: by setting a filter function on the @box one can decide
	// dynamically which of the children to show. For instance, to implement a
	// search function that only shows the children matching the search terms.
	//
	// The @filter_func will be called for each child after the call, and it
	// will continue to be called each time a child changes (via
	// gtk_flow_box_child_changed()) or when gtk_flow_box_invalidate_filter() is
	// called.
	//
	// Note that using a filter function is incompatible with using a model (see
	// gtk_flow_box_bind_model()).
	SetFilterFunc()
	// SetHAdjustment hooks up an adjustment to focus handling in @box. The
	// adjustment is also used for autoscrolling during rubberband selection.
	// See gtk_scrolled_window_get_hadjustment() for a typical way of obtaining
	// the adjustment, and gtk_flow_box_set_vadjustment()for setting the
	// vertical adjustment.
	//
	// The adjustments have to be in pixel units and in the same coordinate
	// system as the allocation for immediate children of the box.
	SetHAdjustment(adjustment Adjustment)
	// SetHomogeneous sets the FlowBox:homogeneous property of @box, controlling
	// whether or not all children of @box are given equal space in the box.
	SetHomogeneous(homogeneous bool)
	// SetMaxChildrenPerLine sets the maximum number of children to request and
	// allocate space for in @box’s orientation.
	//
	// Setting the maximum number of children per line limits the overall
	// natural size request to be no more than @n_children children long in the
	// given orientation.
	SetMaxChildrenPerLine(nChildren uint)
	// SetMinChildrenPerLine sets the minimum number of children to line up in
	// @box’s orientation before flowing.
	SetMinChildrenPerLine(nChildren uint)
	// SetRowSpacing sets the vertical space to add between children. See the
	// FlowBox:row-spacing property.
	SetRowSpacing(spacing uint)
	// SetSelectionMode sets how selection works in @box. See SelectionMode for
	// details.
	SetSelectionMode(mode SelectionMode)
	// SetSortFunc: by setting a sort function on the @box, one can dynamically
	// reorder the children of the box, based on the contents of the children.
	//
	// The @sort_func will be called for each child after the call, and will
	// continue to be called each time a child changes (via
	// gtk_flow_box_child_changed()) and when gtk_flow_box_invalidate_sort() is
	// called.
	//
	// Note that using a sort function is incompatible with using a model (see
	// gtk_flow_box_bind_model()).
	SetSortFunc()
	// SetVAdjustment hooks up an adjustment to focus handling in @box. The
	// adjustment is also used for autoscrolling during rubberband selection.
	// See gtk_scrolled_window_get_vadjustment() for a typical way of obtaining
	// the adjustment, and gtk_flow_box_set_hadjustment()for setting the
	// horizontal adjustment.
	//
	// The adjustments have to be in pixel units and in the same coordinate
	// system as the allocation for immediate children of the box.
	SetVAdjustment(adjustment Adjustment)
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
	// UnselectChild unselects a single child of @box, if the selection mode
	// allows it.
	UnselectChild(child FlowBoxChild)
}

// flowBox implements the FlowBox interface.
type flowBox struct {
	Container
	Buildable
	Orientable
}

var _ FlowBox = (*flowBox)(nil)

// WrapFlowBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBox(obj *externglib.Object) FlowBox {
	return FlowBox{
		Container:  WrapContainer(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalFlowBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBox(obj), nil
}

// NewFlowBox constructs a class FlowBox.
func NewFlowBox() FlowBox {
	var _cret C.GtkFlowBox

	cret = C.gtk_flow_box_new()

	var _flowBox FlowBox

	_flowBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FlowBox)

	return _flowBox
}

// BindModel binds @model to @box.
//
// If @box was already bound to a model, that previous binding is destroyed.
//
// The contents of @box are cleared and then filled with widgets that
// represent items from @model. @box is updated whenever @model changes. If
// @model is nil, @box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// gtk_flow_box_insert() or gtk_container_add()) while @box is bound to a
// model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in GtkFlowBox. When using a model, filtering and sorting
// should be implemented by the model.
func (b flowBox) BindModel() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_bind_model(_arg0)
}

// ActivateOnSingleClick returns whether children activate on single clicks.
func (b flowBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_flow_box_get_activate_on_single_click(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ChildAtIndex gets the nth child in the @box.
func (b flowBox) ChildAtIndex(idx int) FlowBoxChild {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.gint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(idx)

	var _cret *C.GtkFlowBoxChild

	cret = C.gtk_flow_box_get_child_at_index(_arg0, _arg1)

	var _flowBoxChild FlowBoxChild

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FlowBoxChild)

	return _flowBoxChild
}

// ChildAtPos gets the child in the (@x, @y) position.
func (b flowBox) ChildAtPos(x int, y int) FlowBoxChild {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.gint
	var _arg2 C.gint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	var _cret *C.GtkFlowBoxChild

	cret = C.gtk_flow_box_get_child_at_pos(_arg0, _arg1, _arg2)

	var _flowBoxChild FlowBoxChild

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FlowBoxChild)

	return _flowBoxChild
}

// ColumnSpacing gets the horizontal spacing.
func (b flowBox) ColumnSpacing() uint {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.guint

	cret = C.gtk_flow_box_get_column_spacing(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// Homogeneous returns whether the box is homogeneous (all children are the
// same size). See gtk_box_set_homogeneous().
func (b flowBox) Homogeneous() bool {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean

	cret = C.gtk_flow_box_get_homogeneous(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// MaxChildrenPerLine gets the maximum number of children per line.
func (b flowBox) MaxChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.guint

	cret = C.gtk_flow_box_get_max_children_per_line(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// MinChildrenPerLine gets the minimum number of children per line.
func (b flowBox) MinChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.guint

	cret = C.gtk_flow_box_get_min_children_per_line(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// RowSpacing gets the vertical spacing.
func (b flowBox) RowSpacing() uint {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.guint

	cret = C.gtk_flow_box_get_row_spacing(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// SelectedChildren creates a list of all selected children.
func (b flowBox) SelectedChildren() *glib.List {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret *C.GList

	cret = C.gtk_flow_box_get_selected_children(_arg0)

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _list
}

// SelectionMode gets the selection mode of @box.
func (b flowBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	var _cret C.GtkSelectionMode

	cret = C.gtk_flow_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Insert inserts the @widget into @box at @position.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position and this function has the same effect as
// gtk_container_add().
//
// If @position is -1, or larger than the total number of children in the
// @box, then the @widget will be appended to the end.
func (b flowBox) Insert(widget Widget, position int) {
	var _arg0 *C.GtkFlowBox
	var _arg1 *C.GtkWidget
	var _arg2 C.gint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.gint(position)

	C.gtk_flow_box_insert(_arg0, _arg1, _arg2)
}

// InvalidateFilter updates the filtering for all children.
//
// Call this function when the result of the filter function on the @box is
// changed due ot an external factor. For instance, this would be used if
// the filter function just looked for a specific search term, and the entry
// with the string has changed.
func (b flowBox) InvalidateFilter() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_invalidate_filter(_arg0)
}

// InvalidateSort updates the sorting for all children.
//
// Call this when the result of the sort function on @box is changed due to
// an external factor.
func (b flowBox) InvalidateSort() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_invalidate_sort(_arg0)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (b flowBox) SelectAll() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_select_all(_arg0)
}

// SelectChild selects a single child of @box, if the selection mode allows
// it.
func (b flowBox) SelectChild(child FlowBoxChild) {
	var _arg0 *C.GtkFlowBox
	var _arg1 *C.GtkFlowBoxChild

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_select_child(_arg0, _arg1)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (b flowBox) SelectedForeach() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_selected_foreach(_arg0)
}

// SetActivateOnSingleClick: if @single is true, children will be activated
// when you click on them, otherwise you need to double-click.
func (b flowBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	if single {
		_arg1 = C.gboolean(1)
	}

	C.gtk_flow_box_set_activate_on_single_click(_arg0, _arg1)
}

// SetColumnSpacing sets the horizontal space to add between children. See
// the FlowBox:column-spacing property.
func (b flowBox) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.guint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_column_spacing(_arg0, _arg1)
}

// SetFilterFunc: by setting a filter function on the @box one can decide
// dynamically which of the children to show. For instance, to implement a
// search function that only shows the children matching the search terms.
//
// The @filter_func will be called for each child after the call, and it
// will continue to be called each time a child changes (via
// gtk_flow_box_child_changed()) or when gtk_flow_box_invalidate_filter() is
// called.
//
// Note that using a filter function is incompatible with using a model (see
// gtk_flow_box_bind_model()).
func (b flowBox) SetFilterFunc() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_set_filter_func(_arg0)
}

// SetHAdjustment hooks up an adjustment to focus handling in @box. The
// adjustment is also used for autoscrolling during rubberband selection.
// See gtk_scrolled_window_get_hadjustment() for a typical way of obtaining
// the adjustment, and gtk_flow_box_set_vadjustment()for setting the
// vertical adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate
// system as the allocation for immediate children of the box.
func (b flowBox) SetHAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkFlowBox
	var _arg1 *C.GtkAdjustment

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_hadjustment(_arg0, _arg1)
}

// SetHomogeneous sets the FlowBox:homogeneous property of @box, controlling
// whether or not all children of @box are given equal space in the box.
func (b flowBox) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.gboolean

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.gboolean(1)
	}

	C.gtk_flow_box_set_homogeneous(_arg0, _arg1)
}

// SetMaxChildrenPerLine sets the maximum number of children to request and
// allocate space for in @box’s orientation.
//
// Setting the maximum number of children per line limits the overall
// natural size request to be no more than @n_children children long in the
// given orientation.
func (b flowBox) SetMaxChildrenPerLine(nChildren uint) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.guint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_max_children_per_line(_arg0, _arg1)
}

// SetMinChildrenPerLine sets the minimum number of children to line up in
// @box’s orientation before flowing.
func (b flowBox) SetMinChildrenPerLine(nChildren uint) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.guint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_min_children_per_line(_arg0, _arg1)
}

// SetRowSpacing sets the vertical space to add between children. See the
// FlowBox:row-spacing property.
func (b flowBox) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.guint

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_row_spacing(_arg0, _arg1)
}

// SetSelectionMode sets how selection works in @box. See SelectionMode for
// details.
func (b flowBox) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkFlowBox
	var _arg1 C.GtkSelectionMode

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (C.GtkSelectionMode)(mode)

	C.gtk_flow_box_set_selection_mode(_arg0, _arg1)
}

// SetSortFunc: by setting a sort function on the @box, one can dynamically
// reorder the children of the box, based on the contents of the children.
//
// The @sort_func will be called for each child after the call, and will
// continue to be called each time a child changes (via
// gtk_flow_box_child_changed()) and when gtk_flow_box_invalidate_sort() is
// called.
//
// Note that using a sort function is incompatible with using a model (see
// gtk_flow_box_bind_model()).
func (b flowBox) SetSortFunc() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_set_sort_func(_arg0)
}

// SetVAdjustment hooks up an adjustment to focus handling in @box. The
// adjustment is also used for autoscrolling during rubberband selection.
// See gtk_scrolled_window_get_vadjustment() for a typical way of obtaining
// the adjustment, and gtk_flow_box_set_hadjustment()for setting the
// horizontal adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate
// system as the allocation for immediate children of the box.
func (b flowBox) SetVAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkFlowBox
	var _arg1 *C.GtkAdjustment

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_vadjustment(_arg0, _arg1)
}

// UnselectAll: unselect all children of @box, if the selection mode allows
// it.
func (b flowBox) UnselectAll() {
	var _arg0 *C.GtkFlowBox

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_unselect_all(_arg0)
}

// UnselectChild unselects a single child of @box, if the selection mode
// allows it.
func (b flowBox) UnselectChild(child FlowBoxChild) {
	var _arg0 *C.GtkFlowBox
	var _arg1 *C.GtkFlowBoxChild

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_unselect_child(_arg0, _arg1)
}

type FlowBoxChild interface {
	Bin
	Buildable

	// Changed marks @child as changed, causing any state that depends on this
	// to be updated. This affects sorting and filtering.
	//
	// Note that calls to this method must be in sync with the data used for the
	// sorting and filtering functions. For instance, if the list is mirroring
	// some external data set, and *two* children changed in the external data
	// set when you call gtk_flow_box_child_changed() on the first child, the
	// sort function must only read the new data for the first of the two
	// changed children, otherwise the resorting of the children will be wrong.
	//
	// This generally means that if you don’t fully control the data model, you
	// have to duplicate the data that affects the sorting and filtering
	// functions into the widgets themselves. Another alternative is to call
	// gtk_flow_box_invalidate_sort() on any model change, but that is more
	// expensive.
	Changed()
	// Index gets the current index of the @child in its FlowBox container.
	Index() int
	// IsSelected returns whether the @child is currently selected in its
	// FlowBox container.
	IsSelected() bool
}

// flowBoxChild implements the FlowBoxChild interface.
type flowBoxChild struct {
	Bin
	Buildable
}

var _ FlowBoxChild = (*flowBoxChild)(nil)

// WrapFlowBoxChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBoxChild(obj *externglib.Object) FlowBoxChild {
	return FlowBoxChild{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBoxChild(obj), nil
}

// NewFlowBoxChild constructs a class FlowBoxChild.
func NewFlowBoxChild() FlowBoxChild {
	var _cret C.GtkFlowBoxChild

	cret = C.gtk_flow_box_child_new()

	var _flowBoxChild FlowBoxChild

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FlowBoxChild)

	return _flowBoxChild
}

// Changed marks @child as changed, causing any state that depends on this
// to be updated. This affects sorting and filtering.
//
// Note that calls to this method must be in sync with the data used for the
// sorting and filtering functions. For instance, if the list is mirroring
// some external data set, and *two* children changed in the external data
// set when you call gtk_flow_box_child_changed() on the first child, the
// sort function must only read the new data for the first of the two
// changed children, otherwise the resorting of the children will be wrong.
//
// This generally means that if you don’t fully control the data model, you
// have to duplicate the data that affects the sorting and filtering
// functions into the widgets themselves. Another alternative is to call
// gtk_flow_box_invalidate_sort() on any model change, but that is more
// expensive.
func (c flowBoxChild) Changed() {
	var _arg0 *C.GtkFlowBoxChild

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	C.gtk_flow_box_child_changed(_arg0)
}

// Index gets the current index of the @child in its FlowBox container.
func (c flowBoxChild) Index() int {
	var _arg0 *C.GtkFlowBoxChild

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	var _cret C.gint

	cret = C.gtk_flow_box_child_get_index(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// IsSelected returns whether the @child is currently selected in its
// FlowBox container.
func (c flowBoxChild) IsSelected() bool {
	var _arg0 *C.GtkFlowBoxChild

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean

	cret = C.gtk_flow_box_child_is_selected(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}