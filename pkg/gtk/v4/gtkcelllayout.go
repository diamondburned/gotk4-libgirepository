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
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_layout_get_type()), F: marshalCellLayout},
	})
}

// CellLayoutDataFunc: a function which should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
type CellLayoutDataFunc func()

//export gotk4_CellLayoutDataFunc
func gotk4_CellLayoutDataFunc(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(CellLayoutDataFunc)
	fn()
}

// CellLayoutOverrider contains methods that are overridable. This
// interface is a subset of the interface CellLayout.
type CellLayoutOverrider interface {
	// AddAttribute adds an attribute mapping to the list in @cell_layout.
	//
	// The @column is the column of the model to get a value from, and the
	// @attribute is the parameter on @cell to be set from the value. So for
	// example if column 2 of the model contains strings, you could have the
	// “text” attribute of a CellRendererText get its values from column 2.
	AddAttribute(cell CellRenderer, attribute string, column int)
	// Clear unsets all the mappings on all renderers on @cell_layout and
	// removes all renderers from @cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRenderer)
	// Area returns the underlying CellArea which might be @cell_layout if
	// called on a CellArea or might be nil if no CellArea is used by
	// @cell_layout.
	Area() CellArea
	// Cells returns the cell renderers which have been added to @cell_layout.
	Cells() *glib.List
	// PackEnd adds the @cell to the end of @cell_layout. If @expand is false,
	// then the @cell is allocated no more space than it needs. Any unused space
	// is divided evenly between cells for which @expand is true.
	//
	// Note that reusing the same cell renderer is not supported.
	PackEnd(cell CellRenderer, expand bool)
	// PackStart packs the @cell into the beginning of @cell_layout. If @expand
	// is false, then the @cell is allocated no more space than it needs. Any
	// unused space is divided evenly between cells for which @expand is true.
	//
	// Note that reusing the same cell renderer is not supported.
	PackStart(cell CellRenderer, expand bool)
	// Reorder re-inserts @cell at @position.
	//
	// Note that @cell has already to be packed into @cell_layout for this to
	// function properly.
	Reorder(cell CellRenderer, position int)
	// SetCellDataFunc sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	SetCellDataFunc()
}

// CellLayout is an interface to be implemented by all objects which want to
// provide a TreeViewColumn like API for packing cells, setting attributes and
// data funcs.
//
// One of the notable features provided by implementations of GtkCellLayout are
// attributes. Attributes let you set the properties in flexible ways. They can
// just be set to constant values like regular properties. But they can also be
// mapped to a column of the underlying tree model with
// gtk_cell_layout_set_attributes(), which means that the value of the attribute
// can change from cell to cell as they are rendered by the cell renderer.
// Finally, it is possible to specify a function with
// gtk_cell_layout_set_cell_data_func() that is called to determine the value of
// the attribute for each cell that is rendered.
//
//
// GtkCellLayouts as GtkBuildable
//
// Implementations of GtkCellLayout which also implement the GtkBuildable
// interface (CellView, IconView, ComboBox, EntryCompletion, TreeViewColumn)
// accept GtkCellRenderer objects as <child> elements in UI definitions. They
// support a custom <attributes> element for their children, which can contain
// multiple <attribute> elements. Each <attribute> element has a name attribute
// which specifies a property of the cell renderer; the content of the element
// is the attribute value.
//
// This is an example of a UI definition fragment specifying attributes:
//
//    <object class="GtkCellView">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <attributes>
//          <attribute name="text">0</attribute>
//        </attributes>
//      </child>"
//    </object>
//
// Furthermore for implementations of GtkCellLayout that use a CellArea to lay
// out cells (all GtkCellLayouts in GTK use a GtkCellArea) [cell
// properties][cell-properties] can also be defined in the format by specifying
// the custom <cell-packing> attribute which can contain multiple <property>
// elements defined in the normal way.
//
// Here is a UI definition fragment specifying cell properties:
//
//    <object class="GtkTreeViewColumn">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <cell-packing>
//          <property name="align">True</property>
//          <property name="expand">False</property>
//        </cell-packing>
//      </child>"
//    </object>
//
//
// Subclassing GtkCellLayout implementations
//
// When subclassing a widget that implements CellLayout like IconView or
// ComboBox, there are some considerations related to the fact that these
// widgets internally use a CellArea. The cell area is exposed as a
// construct-only property by these widgets. This means that it is possible to
// e.g. do
//
//    static void
//    my_combo_box_init (MyComboBox *b)
//    {
//      GtkCellRenderer *cell;
//
//      cell = gtk_cell_renderer_pixbuf_new ();
//      // The following call causes the default cell area for combo boxes,
//      // a GtkCellAreaBox, to be instantiated
//      gtk_cell_layout_pack_start (GTK_CELL_LAYOUT (b), cell, FALSE);
//      ...
//    }
//
//    GtkWidget *
//    my_combo_box_new (GtkCellArea *area)
//    {
//      // This call is going to cause a warning about area being ignored
//      return g_object_new (MY_TYPE_COMBO_BOX, "cell-area", area, NULL);
//    }
//
// If supporting alternative cell areas with your derived widget is not
// important, then this does not have to concern you. If you want to support
// alternative cell areas, you can do so by moving the problematic calls out of
// init() and into a constructor() for your class.
type CellLayout interface {
	gextras.Objector
	CellLayoutOverrider
}

// cellLayout implements the CellLayout interface.
type cellLayout struct {
	gextras.Objector
}

var _ CellLayout = (*cellLayout)(nil)

// WrapCellLayout wraps a GObject to a type that implements interface
// CellLayout. It is primarily used internally.
func WrapCellLayout(obj *externglib.Object) CellLayout {
	return CellLayout{
		Objector: obj,
	}
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellLayout(obj), nil
}

// AddAttribute adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the parameter on @cell to be set from the value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a CellRendererText get its values from column 2.
func (c cellLayout) AddAttribute(cell CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkCellLayout
	var _arg1 *C.GtkCellRenderer
	var _arg2 *C.char
	var _arg3 C.int

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(column)

	C.gtk_cell_layout_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

// Clear unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (c cellLayout) Clear() {
	var _arg0 *C.GtkCellLayout

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	C.gtk_cell_layout_clear(_arg0)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (c cellLayout) ClearAttributes(cell CellRenderer) {
	var _arg0 *C.GtkCellLayout
	var _arg1 *C.GtkCellRenderer

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_layout_clear_attributes(_arg0, _arg1)
}

// Area returns the underlying CellArea which might be @cell_layout if
// called on a CellArea or might be nil if no CellArea is used by
// @cell_layout.
func (c cellLayout) Area() CellArea {
	var _arg0 *C.GtkCellLayout

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	var _cret *C.GtkCellArea

	cret = C.gtk_cell_layout_get_area(_arg0)

	var _cellArea CellArea

	_cellArea = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(CellArea)

	return _cellArea
}

// Cells returns the cell renderers which have been added to @cell_layout.
func (c cellLayout) Cells() *glib.List {
	var _arg0 *C.GtkCellLayout

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	var _cret *C.GList

	cret = C.gtk_cell_layout_get_cells(_arg0)

	var _list *glib.List

	_list = glib.WrapList(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _list
}

// PackEnd adds the @cell to the end of @cell_layout. If @expand is false,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is true.
//
// Note that reusing the same cell renderer is not supported.
func (c cellLayout) PackEnd(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkCellLayout
	var _arg1 *C.GtkCellRenderer
	var _arg2 C.gboolean

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.gboolean(1)
	}

	C.gtk_cell_layout_pack_end(_arg0, _arg1, _arg2)
}

// PackStart packs the @cell into the beginning of @cell_layout. If @expand
// is false, then the @cell is allocated no more space than it needs. Any
// unused space is divided evenly between cells for which @expand is true.
//
// Note that reusing the same cell renderer is not supported.
func (c cellLayout) PackStart(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkCellLayout
	var _arg1 *C.GtkCellRenderer
	var _arg2 C.gboolean

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.gboolean(1)
	}

	C.gtk_cell_layout_pack_start(_arg0, _arg1, _arg2)
}

// Reorder re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout for this to
// function properly.
func (c cellLayout) Reorder(cell CellRenderer, position int) {
	var _arg0 *C.GtkCellLayout
	var _arg1 *C.GtkCellRenderer
	var _arg2 C.int

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = C.int(position)

	C.gtk_cell_layout_reorder(_arg0, _arg1, _arg2)
}

// SetCellDataFunc sets the CellLayoutDataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping for
// setting the column value, and should set the value of @cell_layout’s cell
// renderer(s) as appropriate.
//
// @func may be nil to remove a previously set function.
func (c cellLayout) SetCellDataFunc() {
	var _arg0 *C.GtkCellLayout

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	C.gtk_cell_layout_set_cell_data_func(_arg0)
}
