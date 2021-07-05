// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_get_type()), F: marshalComboBox},
	})
}

// ComboBox is a widget that allows the user to choose from a list of valid
// choices. The GtkComboBox displays the selected choice. When activated, the
// GtkComboBox displays a popup which allows the user to make a new choice. The
// style in which the selected value is displayed, and the style of the popup is
// determined by the current theme. It may be similar to a Windows-style combo
// box.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since GtkComboBox implements the CellLayout
// interface. The tree model holding the valid choices is not restricted to a
// flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the “has-entry” property
// allows the GtkComboBox to contain a Entry. This entry can be accessed by
// calling gtk_bin_get_child() on the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox can
// be a bit overwhelming. In this case, ComboBoxText offers a simple
// alternative. Both GtkComboBox and ComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
// A GtkComboBox with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
type ComboBox interface {
	Bin

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellEditable casts the class to the CellEditable interface.
	AsCellEditable() CellEditable
	// AsCellLayout casts the class to the CellLayout interface.
	AsCellLayout() CellLayout

	// Active returns the index of the currently active item, or -1 if there’s
	// no active item. If the model is a non-flat treemodel, and the active item
	// is not an immediate child of the root of the tree, this function returns
	// `gtk_tree_path_get_indices (path)[0]`, where `path` is the TreePath of
	// the active item.
	Active() int
	// ActiveID returns the ID of the active row of @combo_box. This value is
	// taken from the active row and the column specified by the
	// ComboBox:id-column property of @combo_box (see
	// gtk_combo_box_set_id_column()).
	//
	// The returned value is an interned string which means that you can compare
	// the pointer by value to other interned strings and that you must not free
	// it.
	//
	// If the ComboBox:id-column property of @combo_box is not set, or if no row
	// is active, or if the active row has a nil ID value, then nil is returned.
	ActiveID() string
	// ActiveIter sets @iter to point to the currently active item, if any item
	// is active. Otherwise, @iter is left unchanged.
	ActiveIter() (TreeIter, bool)
	// AddTearoffs gets the current value of the :add-tearoffs property.
	//
	// Deprecated: since version 3.10.
	AddTearoffs() bool
	// ButtonSensitivity returns whether the combo box sets the dropdown button
	// sensitive or not when there are no items in the model.
	ButtonSensitivity() SensitivityType
	// ColumnSpanColumn returns the column with column span information for
	// @combo_box.
	ColumnSpanColumn() int
	// EntryTextColumn returns the column which @combo_box is using to get the
	// strings from to display in the internal entry.
	EntryTextColumn() int
	// FocusOnClick returns whether the combo box grabs focus when it is clicked
	// with the mouse. See gtk_combo_box_set_focus_on_click().
	//
	// Deprecated: since version 3.20.
	FocusOnClick() bool
	// HasEntry returns whether the combo box has an entry.
	HasEntry() bool
	// IDColumn returns the column which @combo_box is using to get string IDs
	// for values from.
	IDColumn() int
	// Model returns the TreeModel which is acting as data source for
	// @combo_box.
	Model() TreeModel
	// PopupAccessible gets the accessible object corresponding to the combo
	// box’s popup.
	//
	// This function is mostly intended for use by accessibility technologies;
	// applications should have little use for it.
	PopupAccessible() atk.Object
	// PopupFixedWidth gets whether the popup uses a fixed width matching the
	// allocated width of the combo box.
	PopupFixedWidth() bool
	// RowSpanColumn returns the column with row span information for
	// @combo_box.
	RowSpanColumn() int
	// Title gets the current title of the menu in tearoff mode. See
	// gtk_combo_box_set_add_tearoffs().
	//
	// Deprecated: since version 3.10.
	Title() string
	// WrapWidth returns the wrap width which is used to determine the number of
	// columns for the popup menu. If the wrap width is larger than 1, the combo
	// box is in table mode.
	WrapWidth() int
	// PopdownComboBox hides the menu or dropdown list of @combo_box.
	//
	// This function is mostly intended for use by accessibility technologies;
	// applications should have little use for it.
	PopdownComboBox()
	// PopupComboBox pops up the menu or dropdown list of @combo_box.
	//
	// This function is mostly intended for use by accessibility technologies;
	// applications should have little use for it.
	//
	// Before calling this, @combo_box must be mapped, or nothing will happen.
	PopupComboBox()
	// PopupForDeviceComboBox pops up the menu or dropdown list of @combo_box,
	// the popup window will be grabbed so only @device and its associated
	// pointer/keyboard are the only Devices able to send events to it.
	PopupForDeviceComboBox(device gdk.Device)
	// SetActiveComboBox sets the active item of @combo_box to be the item at
	// @index.
	SetActiveComboBox(index_ int)
	// SetActiveIDComboBox changes the active row of @combo_box to the one that
	// has an ID equal to @active_id, or unsets the active row if @active_id is
	// nil. Rows having a nil ID string cannot be made active by this function.
	//
	// If the ComboBox:id-column property of @combo_box is unset or if no row
	// has the given ID then the function does nothing and returns false.
	SetActiveIDComboBox(activeId string) bool
	// SetActiveIterComboBox sets the current active item to be the one
	// referenced by @iter, or unsets the active item if @iter is nil.
	SetActiveIterComboBox(iter *TreeIter)
	// SetAddTearoffsComboBox sets whether the popup menu should have a tearoff
	// menu item.
	//
	// Deprecated: since version 3.10.
	SetAddTearoffsComboBox(addTearoffs bool)
	// SetButtonSensitivityComboBox sets whether the dropdown button of the
	// combo box should be always sensitive (GTK_SENSITIVITY_ON), never
	// sensitive (GTK_SENSITIVITY_OFF) or only if there is at least one item to
	// display (GTK_SENSITIVITY_AUTO).
	SetButtonSensitivityComboBox(sensitivity SensitivityType)
	// SetColumnSpanColumnComboBox sets the column with column span information
	// for @combo_box to be @column_span. The column span column contains
	// integers which indicate how many columns an item should span.
	SetColumnSpanColumnComboBox(columnSpan int)
	// SetEntryTextColumnComboBox sets the model column which @combo_box should
	// use to get strings from to be @text_column. The column @text_column in
	// the model of @combo_box must be of type G_TYPE_STRING.
	//
	// This is only relevant if @combo_box has been created with
	// ComboBox:has-entry as true.
	SetEntryTextColumnComboBox(textColumn int)
	// SetFocusOnClickComboBox sets whether the combo box will grab focus when
	// it is clicked with the mouse. Making mouse clicks not grab focus is
	// useful in places like toolbars where you don’t want the keyboard focus
	// removed from the main area of the application.
	//
	// Deprecated: since version 3.20.
	SetFocusOnClickComboBox(focusOnClick bool)
	// SetIDColumnComboBox sets the model column which @combo_box should use to
	// get string IDs for values from. The column @id_column in the model of
	// @combo_box must be of type G_TYPE_STRING.
	SetIDColumnComboBox(idColumn int)
	// SetModelComboBox sets the model used by @combo_box to be @model. Will
	// unset a previously set model (if applicable). If model is nil, then it
	// will unset the model.
	//
	// Note that this function does not clear the cell renderers, you have to
	// call gtk_cell_layout_clear() yourself if you need to set up different
	// cell renderers for the new model.
	SetModelComboBox(model TreeModel)
	// SetPopupFixedWidthComboBox specifies whether the popup’s width should be
	// a fixed width matching the allocated width of the combo box.
	SetPopupFixedWidthComboBox(fixed bool)
	// SetRowSpanColumnComboBox sets the column with row span information for
	// @combo_box to be @row_span. The row span column contains integers which
	// indicate how many rows an item should span.
	SetRowSpanColumnComboBox(rowSpan int)
	// SetTitleComboBox sets the menu’s title in tearoff mode.
	//
	// Deprecated: since version 3.10.
	SetTitleComboBox(title string)
	// SetWrapWidthComboBox sets the wrap width of @combo_box to be @width. The
	// wrap width is basically the preferred number of columns when you want the
	// popup to be layed out in a table.
	SetWrapWidthComboBox(width int)
}

// comboBox implements the ComboBox class.
type comboBox struct {
	Bin
}

// WrapComboBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapComboBox(obj *externglib.Object) ComboBox {
	return comboBox{
		Bin: WrapBin(obj),
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComboBox(obj), nil
}

// NewComboBox creates a new empty ComboBox.
func NewComboBox() ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithArea creates a new empty ComboBox using @area to layout cells.
func NewComboBoxWithArea(area CellArea) ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area(_arg1)

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithAreaAndEntry creates a new empty ComboBox with an entry.
//
// The new combo box will use @area to layout cells.
func NewComboBoxWithAreaAndEntry(area CellArea) ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area_and_entry(_arg1)

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty ComboBox with an entry.
func NewComboBoxWithEntry() ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new ComboBox with the model initialized to
// @model.
func NewComboBoxWithModel(model TreeModel) ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty ComboBox with an entry and
// with the model initialized to @model.
func NewComboBoxWithModelAndEntry(model TreeModel) ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)

	var _comboBox ComboBox // out

	_comboBox = WrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

func (c comboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c comboBox) ActiveIter() (TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.GtkTreeIter  // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, &_arg1)

	var _iter TreeIter // out
	var _ok bool       // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

func (c comboBox) AddTearoffs() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_add_tearoffs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c comboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

func (c comboBox) ColumnSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_column_span_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) FocusOnClick() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c comboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c comboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) Model() TreeModel {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

func (c comboBox) PopupAccessible() atk.Object {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.AtkObject   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_popup_accessible(_arg0)

	var _object atk.Object // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(atk.Object)

	return _object
}

func (c comboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c comboBox) RowSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_row_span_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) Title() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c comboBox) WrapWidth() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_get_wrap_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c comboBox) PopdownComboBox() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popdown(_arg0)
}

func (c comboBox) PopupComboBox() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_popup(_arg0)
}

func (c comboBox) PopupForDeviceComboBox(device gdk.Device) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
}

func (c comboBox) SetActiveComboBox(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
}

func (c comboBox) SetActiveIDComboBox(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(activeId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c comboBox) SetActiveIterComboBox(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
}

func (c comboBox) SetAddTearoffsComboBox(addTearoffs bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_add_tearoffs(_arg0, _arg1)
}

func (c comboBox) SetButtonSensitivityComboBox(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
}

func (c comboBox) SetColumnSpanColumnComboBox(columnSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(columnSpan)

	C.gtk_combo_box_set_column_span_column(_arg0, _arg1)
}

func (c comboBox) SetEntryTextColumnComboBox(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
}

func (c comboBox) SetFocusOnClickComboBox(focusOnClick bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_focus_on_click(_arg0, _arg1)
}

func (c comboBox) SetIDColumnComboBox(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
}

func (c comboBox) SetModelComboBox(model TreeModel) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_combo_box_set_model(_arg0, _arg1)
}

func (c comboBox) SetPopupFixedWidthComboBox(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
}

func (c comboBox) SetRowSpanColumnComboBox(rowSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(rowSpan)

	C.gtk_combo_box_set_row_span_column(_arg0, _arg1)
}

func (c comboBox) SetTitleComboBox(title string) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_set_title(_arg0, _arg1)
}

func (c comboBox) SetWrapWidthComboBox(width int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(width)

	C.gtk_combo_box_set_wrap_width(_arg0, _arg1)
}

func (c comboBox) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(c))
}

func (c comboBox) AsCellEditable() CellEditable {
	return WrapCellEditable(gextras.InternObject(c))
}

func (c comboBox) AsCellLayout() CellLayout {
	return WrapCellLayout(gextras.InternObject(c))
}
