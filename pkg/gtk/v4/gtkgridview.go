// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_grid_view_get_type()), F: marshalGridViewer},
	})
}

// GridView: GtkGridView presents a large dynamic grid of items.
//
// GtkGridView uses its factory to generate one child widget for each visible
// item and shows them in a grid. The orientation of the grid view determines if
// the grid reflows vertically or horizontally.
//
// GtkGridView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// gtk.GridView:enable-rubberband.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
// CSS nodes
//
//    gridview
//    ├── child
//    │
//    ├── child
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkGridView uses a single CSS node with name gridview. Each child uses a
// single CSS node with name child. For rubberband selection, a subnode with
// name rubberband is used.
//
//
// Accessibility
//
// GtkGridView uses the GTK_ACCESSIBLE_ROLE_GRID role, and the items use the
// GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type GridView struct {
	ListBase
}

func wrapGridView(obj *externglib.Object) *GridView {
	return &GridView{
		ListBase: ListBase{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
			Orientable: Orientable{
				Object: obj,
			},
			Scrollable: Scrollable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalGridViewer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGridView(obj), nil
}

// NewGridView creates a new GtkGridView that uses the given factory for mapping
// items to widgets.
//
// The function takes ownership of the arguments, so you can write code like
//
//    grid_view = gtk_grid_view_new (create_model (),
//      gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));
func NewGridView(model SelectionModeller, factory *ListItemFactory) *GridView {
	var _arg1 *C.GtkSelectionModel  // out
	var _arg2 *C.GtkListItemFactory // out
	var _cret *C.GtkWidget          // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref((*C.GObject)(unsafe.Pointer(model.Native())))
	}
	if factory != nil {
		_arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
		factory.Ref()
	}

	_cret = C.gtk_grid_view_new(_arg1, _arg2)

	var _gridView *GridView // out

	_gridView = wrapGridView(externglib.Take(unsafe.Pointer(_cret)))

	return _gridView
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
func (self *GridView) EnableRubberband() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_enable_rubberband(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items.
func (self *GridView) Factory() *ListItemFactory {
	var _arg0 *C.GtkGridView        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_factory(_arg0)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// MaxColumns gets the maximum number of columns that the grid will use.
func (self *GridView) MaxColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_max_columns(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MinColumns gets the minimum number of columns that the grid will use.
func (self *GridView) MinColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_min_columns(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Model gets the model that's currently used to read the items displayed.
func (self *GridView) Model() SelectionModeller {
	var _arg0 *C.GtkGridView       // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_model(_arg0)

	var _selectionModel SelectionModeller // out

	if _cret != nil {
		_selectionModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(SelectionModeller)
	}

	return _selectionModel
}

// SingleClickActivate returns whether items will be activated on single click
// and selected on hover.
func (self *GridView) SingleClickActivate() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_grid_view_get_single_click_activate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
func (self *GridView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_enable_rubberband(_arg0, _arg1)
}

// SetFactory sets the GtkListItemFactory to use for populating list items.
func (self *GridView) SetFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkGridView        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
	}

	C.gtk_grid_view_set_factory(_arg0, _arg1)
}

// SetMaxColumns sets the maximum number of columns to use.
//
// This number must be at least 1.
//
// If max_columns is smaller than the minimum set via
// gtk.GridView.SetMinColumns(), that value is used instead.
func (self *GridView) SetMaxColumns(maxColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(maxColumns)

	C.gtk_grid_view_set_max_columns(_arg0, _arg1)
}

// SetMinColumns sets the minimum number of columns to use.
//
// This number must be at least 1.
//
// If min_columns is smaller than the minimum set via
// gtk.GridView.SetMaxColumns(), that value is ignored.
func (self *GridView) SetMinColumns(minColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(minColumns)

	C.gtk_grid_view_set_min_columns(_arg0, _arg1)
}

// SetModel sets the imodel to use.
//
// This must be a gtk.SelectionModel.
func (self *GridView) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkGridView       // out
	var _arg1 *C.GtkSelectionModel // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_grid_view_set_model(_arg0, _arg1)
}

// SetSingleClickActivate sets whether items should be activated on single click
// and selected on hover.
func (self *GridView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(self.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_single_click_activate(_arg0, _arg1)
}
