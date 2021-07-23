// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

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
// void _gotk4_gtk3_TreeCellDataFunc(GtkTreeViewColumn*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_view_column_sizing_get_type()), F: marshalTreeViewColumnSizing},
		{T: externglib.Type(C.gtk_tree_view_column_get_type()), F: marshalTreeViewColumner},
	})
}

// TreeViewColumnSizing: sizing method the column uses to determine its width.
// Please note that GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for large
// views, and can make columns appear choppy.
type TreeViewColumnSizing int

const (
	// TreeViewColumnGrowOnly columns only get bigger in reaction to changes in
	// the model
	TreeViewColumnGrowOnly TreeViewColumnSizing = iota
	// TreeViewColumnAutosize columns resize to be the optimal size everytime
	// the model changes.
	TreeViewColumnAutosize
	// TreeViewColumnFixed columns are a fixed numbers of pixels wide.
	TreeViewColumnFixed
)

func marshalTreeViewColumnSizing(p uintptr) (interface{}, error) {
	return TreeViewColumnSizing(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for TreeViewColumnSizing.
func (t TreeViewColumnSizing) String() string {
	switch t {
	case TreeViewColumnGrowOnly:
		return "GrowOnly"
	case TreeViewColumnAutosize:
		return "Autosize"
	case TreeViewColumnFixed:
		return "Fixed"
	default:
		return fmt.Sprintf("TreeViewColumnSizing(%d)", t)
	}
}

// TreeCellDataFunc: function to set the properties of a cell instead of just
// using the straight mapping between the cell and the model. This is useful for
// customizing the cell renderer. For example, a function might get an integer
// from the tree_model, and render it to the “text” attribute of “cell” by
// converting it to its written equivalent. This is set by calling
// gtk_tree_view_column_set_cell_data_func()
type TreeCellDataFunc func(treeColumn *TreeViewColumn, cell CellRendererer, treeModel TreeModeller, iter *TreeIter)

//export _gotk4_gtk3_TreeCellDataFunc
func _gotk4_gtk3_TreeCellDataFunc(arg0 *C.GtkTreeViewColumn, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var treeColumn *TreeViewColumn // out
	var cell CellRendererer        // out
	var treeModel TreeModeller     // out
	var iter *TreeIter             // out

	treeColumn = wrapTreeViewColumn(externglib.Take(unsafe.Pointer(arg0)))
	cell = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(CellRendererer)
	treeModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg2)))).(TreeModeller)
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	runtime.SetFinalizer(iter, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(v))))
	})

	fn := v.(TreeCellDataFunc)
	fn(treeColumn, cell, treeModel, iter)
}

// TreeViewColumnOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeViewColumnOverrider interface {
	// Clicked emits the “clicked” signal on the column. This function will only
	// work if tree_column is clickable.
	Clicked()
}

// TreeViewColumn object represents a visible column in a TreeView widget. It
// allows to set properties of the column header, and functions as a holding pen
// for the cell renderers which determine how the data in the column is
// displayed.
//
// Please refer to the [tree widget conceptual overview][TreeWidget] for an
// overview of all the objects and data types related to the tree widget and how
// they work together.
type TreeViewColumn struct {
	externglib.InitiallyUnowned

	Buildable
	CellLayout
	*externglib.Object
}

func wrapTreeViewColumn(obj *externglib.Object) *TreeViewColumn {
	return &TreeViewColumn{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		CellLayout: CellLayout{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalTreeViewColumner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeViewColumn(obj), nil
}

// NewTreeViewColumn creates a new TreeViewColumn.
func NewTreeViewColumn() *TreeViewColumn {
	var _cret *C.GtkTreeViewColumn // in

	_cret = C.gtk_tree_view_column_new()

	var _treeViewColumn *TreeViewColumn // out

	_treeViewColumn = wrapTreeViewColumn(externglib.Take(unsafe.Pointer(_cret)))

	return _treeViewColumn
}

// NewTreeViewColumnWithArea creates a new TreeViewColumn using area to render
// its cells.
func NewTreeViewColumnWithArea(area CellAreaer) *TreeViewColumn {
	var _arg1 *C.GtkCellArea       // out
	var _cret *C.GtkTreeViewColumn // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_tree_view_column_new_with_area(_arg1)

	var _treeViewColumn *TreeViewColumn // out

	_treeViewColumn = wrapTreeViewColumn(externglib.Take(unsafe.Pointer(_cret)))

	return _treeViewColumn
}

// Native solves the ambiguous selector of this class or interface.
func (treeColumn *TreeViewColumn) Native() uintptr {
	return treeColumn.Object.Native()
}

// AddAttribute adds an attribute mapping to the list in tree_column. The column
// is the column of the model to get a value from, and the attribute is the
// parameter on cell_renderer to be set from the value. So for example if column
// 2 of the model contains strings, you could have the “text” attribute of a
// CellRendererText get its values from column 2.
func (treeColumn *TreeViewColumn) AddAttribute(cellRenderer CellRendererer, attribute string, column int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 *C.gchar             // out
	var _arg3 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(column)

	C.gtk_tree_view_column_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

// CellGetPosition obtains the horizontal position and size of a cell in a
// column. If the cell is not found in the column, start_pos and width are not
// changed and FALSE is returned.
func (treeColumn *TreeViewColumn) CellGetPosition(cellRenderer CellRendererer) (xOffset int, width int, ok bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gint               // in
	var _arg3 C.gint               // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	_cret = C.gtk_tree_view_column_cell_get_position(_arg0, _arg1, &_arg2, &_arg3)

	var _xOffset int // out
	var _width int   // out
	var _ok bool     // out

	_xOffset = int(_arg2)
	_width = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _xOffset, _width, _ok
}

// CellGetSize obtains the width and height needed to render the column. This is
// used primarily by the TreeView.
func (treeColumn *TreeViewColumn) CellGetSize(cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GdkRectangle      // out
	var _arg2 C.gint               // in
	var _arg3 C.gint               // in
	var _arg4 C.gint               // in
	var _arg5 C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if cellArea != nil {
		_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	}

	C.gtk_tree_view_column_cell_get_size(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _xOffset int // out
	var _yOffset int // out
	var _width int   // out
	var _height int  // out

	_xOffset = int(_arg2)
	_yOffset = int(_arg3)
	_width = int(_arg4)
	_height = int(_arg5)

	return _xOffset, _yOffset, _width, _height
}

// CellIsVisible returns TRUE if any of the cells packed into the tree_column
// are visible. For this to be meaningful, you must first initialize the cells
// with gtk_tree_view_column_cell_set_cell_data()
func (treeColumn *TreeViewColumn) CellIsVisible() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_cell_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CellSetCellData sets the cell renderer based on the tree_model and iter. That
// is, for every attribute mapping in tree_column, it will get a value from the
// set column on the iter, and use that value to set the attribute on the cell
// renderer. This is used primarily by the TreeView.
func (treeColumn *TreeViewColumn) CellSetCellData(treeModel TreeModeller, iter *TreeIter, isExpander bool, isExpanded bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkTreeModel      // out
	var _arg2 *C.GtkTreeIter       // out
	var _arg3 C.gboolean           // out
	var _arg4 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if isExpander {
		_arg3 = C.TRUE
	}
	if isExpanded {
		_arg4 = C.TRUE
	}

	C.gtk_tree_view_column_cell_set_cell_data(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Clear unsets all the mappings on all renderers on the tree_column.
func (treeColumn *TreeViewColumn) Clear() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	C.gtk_tree_view_column_clear(_arg0)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_tree_view_column_set_attributes().
func (treeColumn *TreeViewColumn) ClearAttributes(cellRenderer CellRendererer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	C.gtk_tree_view_column_clear_attributes(_arg0, _arg1)
}

// Clicked emits the “clicked” signal on the column. This function will only
// work if tree_column is clickable.
func (treeColumn *TreeViewColumn) Clicked() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	C.gtk_tree_view_column_clicked(_arg0)
}

// FocusCell sets the current keyboard focus to be at cell, if the column
// contains 2 or more editable and activatable cells.
func (treeColumn *TreeViewColumn) FocusCell(cell CellRendererer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_tree_view_column_focus_cell(_arg0, _arg1)
}

// Alignment returns the current x alignment of tree_column. This value can
// range between 0.0 and 1.0.
func (treeColumn *TreeViewColumn) Alignment() float32 {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gfloat             // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Button returns the button used in the treeview column header
func (treeColumn *TreeViewColumn) Button() Widgetter {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_button(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Clickable returns TRUE if the user can click on the header for the column.
func (treeColumn *TreeViewColumn) Clickable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_clickable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expand returns TRUE if the column expands to fill available space.
func (treeColumn *TreeViewColumn) Expand() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_expand(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FixedWidth gets the fixed width of the column. This may not be the actual
// displayed width of the column; for that, use
// gtk_tree_view_column_get_width().
func (treeColumn *TreeViewColumn) FixedWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_fixed_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxWidth returns the maximum width in pixels of the tree_column, or -1 if no
// maximum width is set.
func (treeColumn *TreeViewColumn) MaxWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_max_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MinWidth returns the minimum width in pixels of the tree_column, or -1 if no
// minimum width is set.
func (treeColumn *TreeViewColumn) MinWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_min_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Reorderable returns TRUE if the tree_column can be reordered by the user.
func (treeColumn *TreeViewColumn) Reorderable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Resizable returns TRUE if the tree_column can be resized by the end user.
func (treeColumn *TreeViewColumn) Resizable() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_resizable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sizing returns the current type of tree_column.
func (treeColumn *TreeViewColumn) Sizing() TreeViewColumnSizing {
	var _arg0 *C.GtkTreeViewColumn      // out
	var _cret C.GtkTreeViewColumnSizing // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_sizing(_arg0)

	var _treeViewColumnSizing TreeViewColumnSizing // out

	_treeViewColumnSizing = TreeViewColumnSizing(_cret)

	return _treeViewColumnSizing
}

// SortColumnID gets the logical sort_column_id that the model sorts on when
// this column is selected for sorting. See
// gtk_tree_view_column_set_sort_column_id().
func (treeColumn *TreeViewColumn) SortColumnID() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_sort_column_id(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SortIndicator gets the value set by
// gtk_tree_view_column_set_sort_indicator().
func (treeColumn *TreeViewColumn) SortIndicator() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_sort_indicator(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SortOrder gets the value set by gtk_tree_view_column_set_sort_order().
func (treeColumn *TreeViewColumn) SortOrder() SortType {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.GtkSortType        // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_sort_order(_arg0)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

// Spacing returns the spacing of tree_column.
func (treeColumn *TreeViewColumn) Spacing() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Title returns the title of the widget.
func (treeColumn *TreeViewColumn) Title() string {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TreeView returns the TreeView wherein tree_column has been inserted. If
// column is currently not inserted in any tree view, NULL is returned.
func (treeColumn *TreeViewColumn) TreeView() Widgetter {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_tree_view(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Visible returns TRUE if tree_column is visible.
func (treeColumn *TreeViewColumn) Visible() bool {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Widget returns the Widget in the button on the column header. If a custom
// widget has not been set then NULL is returned.
func (treeColumn *TreeViewColumn) Widget() Widgetter {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_widget(_arg0)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Width returns the current size of tree_column in pixels.
func (treeColumn *TreeViewColumn) Width() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// XOffset returns the current X offset of tree_column in pixels.
func (treeColumn *TreeViewColumn) XOffset() int {
	var _arg0 *C.GtkTreeViewColumn // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	_cret = C.gtk_tree_view_column_get_x_offset(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PackEnd adds the cell to end of the column. If expand is FALSE, then the cell
// is allocated no more space than it needs. Any unused space is divided evenly
// between cells for which expand is TRUE.
func (treeColumn *TreeViewColumn) PackEnd(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_end(_arg0, _arg1, _arg2)
}

// PackStart packs the cell into the beginning of the column. If expand is
// FALSE, then the cell is allocated no more space than it needs. Any unused
// space is divided evenly between cells for which expand is TRUE.
func (treeColumn *TreeViewColumn) PackStart(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_start(_arg0, _arg1, _arg2)
}

// QueueResize flags the column, and the cell renderers added to this column, to
// have their sizes renegotiated.
func (treeColumn *TreeViewColumn) QueueResize() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))

	C.gtk_tree_view_column_queue_resize(_arg0)
}

// SetAlignment sets the alignment of the title or custom widget inside the
// column header. The alignment determines its location inside the button -- 0.0
// for left, 0.5 for center, 1.0 for right.
func (treeColumn *TreeViewColumn) SetAlignment(xalign float32) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gfloat             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_tree_view_column_set_alignment(_arg0, _arg1)
}

// SetCellDataFunc sets the TreeCellDataFunc to use for the column. This
// function is used instead of the standard attributes mapping for setting the
// column value, and should set the value of tree_column's cell renderer as
// appropriate. func may be NULL to remove an older one.
func (treeColumn *TreeViewColumn) SetCellDataFunc(cellRenderer CellRendererer, fn TreeCellDataFunc) {
	var _arg0 *C.GtkTreeViewColumn  // out
	var _arg1 *C.GtkCellRenderer    // out
	var _arg2 C.GtkTreeCellDataFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))
	if fn != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk3_TreeCellDataFunc)
		_arg3 = C.gpointer(gbox.Assign(fn))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_tree_view_column_set_cell_data_func(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetClickable sets the header to be active if clickable is TRUE. When the
// header is active, then it can take keyboard focus, and can be clicked.
func (treeColumn *TreeViewColumn) SetClickable(clickable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if clickable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_clickable(_arg0, _arg1)
}

// SetExpand sets the column to take available extra space. This space is shared
// equally amongst all columns that have the expand set to TRUE. If no column
// has this option set, then the last column gets all extra space. By default,
// every column is created with this FALSE.
//
// Along with “fixed-width”, the “expand” property changes when the column is
// resized by the user.
func (treeColumn *TreeViewColumn) SetExpand(expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_expand(_arg0, _arg1)
}

// SetFixedWidth: if fixed_width is not -1, sets the fixed width of tree_column;
// otherwise unsets it. The effective value of fixed_width is clamped between
// the minimum and maximum width of the column; however, the value stored in the
// “fixed-width” property is not clamped. If the column sizing is
// K_TREE_VIEW_COLUMN_GROW_ONLY or K_TREE_VIEW_COLUMN_AUTOSIZE, setting a fixed
// width overrides the automatically calculated width. Note that fixed_width is
// only a hint to GTK+; the width actually allocated to the column may be
// greater or less than requested.
//
// Along with “expand”, the “fixed-width” property changes when the column is
// resized by the user.
func (treeColumn *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gint(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width(_arg0, _arg1)
}

// SetMaxWidth sets the maximum width of the tree_column. If max_width is -1,
// then the maximum width is unset. Note, the column can actually be wider than
// max width if it’s the last column in a view. In this case, the column expands
// to fill any extra space.
func (treeColumn *TreeViewColumn) SetMaxWidth(maxWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gint(maxWidth)

	C.gtk_tree_view_column_set_max_width(_arg0, _arg1)
}

// SetMinWidth sets the minimum width of the tree_column. If min_width is -1,
// then the minimum width is unset.
func (treeColumn *TreeViewColumn) SetMinWidth(minWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gint(minWidth)

	C.gtk_tree_view_column_set_min_width(_arg0, _arg1)
}

// SetReorderable: if reorderable is TRUE, then the column can be reordered by
// the end user dragging the header.
func (treeColumn *TreeViewColumn) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_reorderable(_arg0, _arg1)
}

// SetResizable: if resizable is TRUE, then the user can explicitly resize the
// column by grabbing the outer edge of the column button. If resizable is TRUE
// and sizing mode of the column is K_TREE_VIEW_COLUMN_AUTOSIZE, then the sizing
// mode is changed to K_TREE_VIEW_COLUMN_GROW_ONLY.
func (treeColumn *TreeViewColumn) SetResizable(resizable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_resizable(_arg0, _arg1)
}

// SetSizing sets the growth behavior of tree_column to type.
func (treeColumn *TreeViewColumn) SetSizing(typ TreeViewColumnSizing) {
	var _arg0 *C.GtkTreeViewColumn      // out
	var _arg1 C.GtkTreeViewColumnSizing // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.GtkTreeViewColumnSizing(typ)

	C.gtk_tree_view_column_set_sizing(_arg0, _arg1)
}

// SetSortColumnID sets the logical sort_column_id that this column sorts on
// when this column is selected for sorting. Doing so makes the column header
// clickable.
func (treeColumn *TreeViewColumn) SetSortColumnID(sortColumnId int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gint(sortColumnId)

	C.gtk_tree_view_column_set_sort_column_id(_arg0, _arg1)
}

// SetSortIndicator: call this function with a setting of TRUE to display an
// arrow in the header button indicating the column is sorted. Call
// gtk_tree_view_column_set_sort_order() to change the direction of the arrow.
func (treeColumn *TreeViewColumn) SetSortIndicator(setting bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_sort_indicator(_arg0, _arg1)
}

// SetSortOrder changes the appearance of the sort indicator.
//
// This does not actually sort the model. Use
// gtk_tree_view_column_set_sort_column_id() if you want automatic sorting
// support. This function is primarily for custom sorting behavior, and should
// be used in conjunction with gtk_tree_sortable_set_sort_column_id() to do
// that. For custom models, the mechanism will vary.
//
// The sort indicator changes direction to indicate normal sort or reverse sort.
// Note that you must have the sort indicator enabled to see anything when
// calling this function; see gtk_tree_view_column_set_sort_indicator().
func (treeColumn *TreeViewColumn) SetSortOrder(order SortType) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.GtkSortType        // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.GtkSortType(order)

	C.gtk_tree_view_column_set_sort_order(_arg0, _arg1)
}

// SetSpacing sets the spacing field of tree_column, which is the number of
// pixels to place between cell renderers packed into it.
func (treeColumn *TreeViewColumn) SetSpacing(spacing int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_tree_view_column_set_spacing(_arg0, _arg1)
}

// SetTitle sets the title of the tree_column. If a custom widget has been set,
// then this value is ignored.
func (treeColumn *TreeViewColumn) SetTitle(title string) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tree_view_column_set_title(_arg0, _arg1)
}

// SetVisible sets the visibility of tree_column.
func (treeColumn *TreeViewColumn) SetVisible(visible bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_visible(_arg0, _arg1)
}

// SetWidget sets the widget in the header to be widget. If widget is NULL, then
// the header button is set with a Label set to the title of tree_column.
func (treeColumn *TreeViewColumn) SetWidget(widget Widgetter) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_tree_view_column_set_widget(_arg0, _arg1)
}
