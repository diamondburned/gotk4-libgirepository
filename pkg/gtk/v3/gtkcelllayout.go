// Code generated by girgen. DO NOT EDIT.

package gtk

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
	GTypeCellLayout = coreglib.Type(girepository.MustFind("Gtk", "CellLayout").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellLayout, F: marshalCellLayout},
	})
}

// CellLayoutDataFunc: function which should set the value of cell_layout’s cell
// renderer(s) as appropriate.
type CellLayoutDataFunc func(cellLayout CellLayouter, cell CellRendererer, treeModel TreeModeller, iter *TreeIter)

// CellLayoutOverrider contains methods that are overridable.
type CellLayoutOverrider interface {
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
// out cells (all GtkCellLayouts in GTK+ use a GtkCellArea) [cell
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
//
// CellLayout wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CellLayout struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellLayout)(nil)
)

// CellLayouter describes CellLayout's interface methods.
type CellLayouter interface {
	coreglib.Objector

	baseCellLayout() *CellLayout
}

var _ CellLayouter = (*CellLayout)(nil)

func ifaceInitCellLayouter(gifacePtr, data C.gpointer) {
}

func wrapCellLayout(obj *coreglib.Object) *CellLayout {
	return &CellLayout{
		Object: obj,
	}
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	return wrapCellLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *CellLayout) baseCellLayout() *CellLayout {
	return v
}

// BaseCellLayout returns the underlying base object.
func BaseCellLayout(obj CellLayouter) *CellLayout {
	return obj.baseCellLayout()
}

// CellLayoutIface: instance of this type is always passed by reference.
type CellLayoutIface struct {
	*cellLayoutIface
}

// cellLayoutIface is the struct that's finalized.
type cellLayoutIface struct {
	native unsafe.Pointer
}

var GIRInfoCellLayoutIface = girepository.MustFind("Gtk", "CellLayoutIface")
