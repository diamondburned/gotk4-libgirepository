// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk3_TreeViewRowSeparatorFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_get_type()), F: marshalComboBoxer},
	})
}

// ComboBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ComboBoxOverrider interface {
	Changed()
	FormatEntryText(path string) string
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
type ComboBox struct {
	Bin

	CellEditable
	CellLayout
	*externglib.Object
}

func wrapComboBox(obj *externglib.Object) *ComboBox {
	return &ComboBox{
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
					Object: obj,
				},
			},
		},
		CellEditable: CellEditable{
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
				Object: obj,
			},
		},
		CellLayout: CellLayout{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalComboBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBox(obj), nil
}

// NewComboBox creates a new empty ComboBox.
func NewComboBox() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithArea creates a new empty ComboBox using area to layout cells.
func NewComboBoxWithArea(area CellAreaer) *ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithAreaAndEntry creates a new empty ComboBox with an entry.
//
// The new combo box will use area to layout cells.
func NewComboBoxWithAreaAndEntry(area CellAreaer) *ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_combo_box_new_with_area_and_entry(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty ComboBox with an entry.
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new ComboBox with the model initialized to
// model.
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty ComboBox with an entry and
// with the model initialized to model.
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// Active returns the index of the currently active item, or -1 if there’s no
// active item. If the model is a non-flat treemodel, and the active item is not
// an immediate child of the root of the tree, this function returns
// gtk_tree_path_get_indices (path)[0], where path is the TreePath of the active
// item.
func (comboBox *ComboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of combo_box. This value is taken
// from the active row and the column specified by the ComboBox:id-column
// property of combo_box (see gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the ComboBox:id-column property of combo_box is not set, or if no row is
// active, or if the active row has a NULL ID value, then NULL is returned.
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActiveIter sets iter to point to the currently active item, if any item is
// active. Otherwise, iter is left unchanged.
func (comboBox *ComboBox) ActiveIter() (TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.GtkTreeIter  // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, &_arg1)

	var _iter TreeIter // out
	var _ok bool       // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// AddTearoffs gets the current value of the :add-tearoffs property.
//
// Deprecated: since version 3.10.
func (comboBox *ComboBox) AddTearoffs() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_add_tearoffs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (comboBox *ComboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// ColumnSpanColumn returns the column with column span information for
// combo_box.
func (comboBox *ComboBox) ColumnSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_column_span_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// EntryTextColumn returns the column which combo_box is using to get the
// strings from to display in the internal entry.
func (comboBox *ComboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// FocusOnClick returns whether the combo box grabs focus when it is clicked
// with the mouse. See gtk_combo_box_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
func (combo *ComboBox) FocusOnClick() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(combo.Native()))

	_cret = C.gtk_combo_box_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasEntry returns whether the combo box has an entry.
func (comboBox *ComboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which combo_box is using to get string IDs for
// values from.
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the TreeModel which is acting as data source for combo_box.
func (comboBox *ComboBox) Model() TreeModeller {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)

	var _treeModel TreeModeller // out

	_treeModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(TreeModeller)

	return _treeModel
}

// PopupAccessible gets the accessible object corresponding to the combo box’s
// popup.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) PopupAccessible() *atk.ObjectClass {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.AtkObject   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_popup_accessible(_arg0)

	var _object *atk.ObjectClass // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_object = &atk.ObjectClass{
			Object: obj,
		}
	}

	return _object
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpanColumn returns the column with row span information for combo_box.
func (comboBox *ComboBox) RowSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_row_span_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Title gets the current title of the menu in tearoff mode. See
// gtk_combo_box_set_add_tearoffs().
//
// Deprecated: since version 3.10.
func (comboBox *ComboBox) Title() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WrapWidth returns the wrap width which is used to determine the number of
// columns for the popup menu. If the wrap width is larger than 1, the combo box
// is in table mode.
func (comboBox *ComboBox) WrapWidth() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_wrap_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popdown(_arg0)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popup(_arg0)
}

// PopupForDevice pops up the menu or dropdown list of combo_box, the popup
// window will be grabbed so only device and its associated pointer/keyboard are
// the only Devices able to send events to it.
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
}

// SetActive sets the active item of combo_box to be the item at index.
func (comboBox *ComboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id, or unsets the active row if active_id is NULL. Rows
// having a NULL ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of combo_box is unset or if no row has the
// given ID then the function does nothing and returns FALSE.
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if activeId != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter,
// or unsets the active item if iter is NULL.
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if iter != nil {
		_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
}

// SetAddTearoffs sets whether the popup menu should have a tearoff menu item.
//
// Deprecated: since version 3.10.
func (comboBox *ComboBox) SetAddTearoffs(addTearoffs bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_add_tearoffs(_arg0, _arg1)
}

// SetButtonSensitivity sets whether the dropdown button of the combo box should
// be always sensitive (GTK_SENSITIVITY_ON), never sensitive
// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
// (GTK_SENSITIVITY_AUTO).
func (comboBox *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
}

// SetColumnSpanColumn sets the column with column span information for
// combo_box to be column_span. The column span column contains integers which
// indicate how many columns an item should span.
func (comboBox *ComboBox) SetColumnSpanColumn(columnSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(columnSpan)

	C.gtk_combo_box_set_column_span_column(_arg0, _arg1)
}

// SetEntryTextColumn sets the model column which combo_box should use to get
// strings from to be text_column. The column text_column in the model of
// combo_box must be of type G_TYPE_STRING.
//
// This is only relevant if combo_box has been created with ComboBox:has-entry
// as TRUE.
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
}

// SetFocusOnClick sets whether the combo box will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
func (combo *ComboBox) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(combo.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_focus_on_click(_arg0, _arg1)
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from. The column id_column in the model of combo_box must be
// of type G_TYPE_STRING.
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
}

// SetModel sets the model used by combo_box to be model. Will unset a
// previously set model (if applicable). If model is NULL, then it will unset
// the model.
//
// Note that this function does not clear the cell renderers, you have to call
// gtk_cell_layout_clear() yourself if you need to set up different cell
// renderers for the new model.
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_combo_box_set_model(_arg0, _arg1)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator. If the row separator
// function is NULL, no separators are drawn. This is the default value.
func (comboBox *ComboBox) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var _arg0 *C.GtkComboBox                // out
	var _arg1 C.GtkTreeViewRowSeparatorFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeViewRowSeparatorFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_combo_box_set_row_separator_func(_arg0, _arg1, _arg2, _arg3)
}

// SetRowSpanColumn sets the column with row span information for combo_box to
// be row_span. The row span column contains integers which indicate how many
// rows an item should span.
func (comboBox *ComboBox) SetRowSpanColumn(rowSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(rowSpan)

	C.gtk_combo_box_set_row_span_column(_arg0, _arg1)
}

// SetTitle sets the menu’s title in tearoff mode.
//
// Deprecated: since version 3.10.
func (comboBox *ComboBox) SetTitle(title string) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_set_title(_arg0, _arg1)
}

// SetWrapWidth sets the wrap width of combo_box to be width. The wrap width is
// basically the preferred number of columns when you want the popup to be layed
// out in a table.
func (comboBox *ComboBox) SetWrapWidth(width int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.gint(width)

	C.gtk_combo_box_set_wrap_width(_arg0, _arg1)
}
