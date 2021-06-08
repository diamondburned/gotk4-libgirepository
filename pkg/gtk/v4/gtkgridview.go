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
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_grid_view_get_type()), F: marshalGridView},
	})
}

// GridView: `GtkGridView` presents a large dynamic grid of items.
//
// `GtkGridView` uses its factory to generate one child widget for each visible
// item and shows them in a grid. The orientation of the grid view determines if
// the grid reflows vertically or horizontally.
//
// `GtkGridView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// [property@Gtk.GridView:enable-rubberband].
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
//
// CSS nodes
//
// “` gridview ├── child │ ├── child │ ┊ ╰── [rubberband] “`
//
// `GtkGridView` uses a single CSS node with name gridview. Each child uses a
// single CSS node with name child. For rubberband selection, a subnode with
// name rubberband is used.
//
//
// Accessibility
//
// `GtkGridView` uses the GTK_ACCESSIBLE_ROLE_GRID role, and the items use the
// GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type GridView interface {
	ListBase
	Accessible
	Buildable
	ConstraintTarget
	Orientable
	Scrollable

	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband() bool
	// Factory gets the factory that's currently used to populate list items.
	Factory() ListItemFactory
	// MaxColumns gets the maximum number of columns that the grid will use.
	MaxColumns() uint
	// MinColumns gets the minimum number of columns that the grid will use.
	MinColumns() uint
	// Model gets the model that's currently used to read the items displayed.
	Model() SelectionModel
	// SingleClickActivate returns whether items will be activated on single
	// click and selected on hover.
	SingleClickActivate() bool
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(enableRubberband bool)
	// SetFactory sets the `GtkListItemFactory` to use for populating list
	// items.
	SetFactory(factory ListItemFactory)
	// SetMaxColumns sets the maximum number of columns to use.
	//
	// This number must be at least 1.
	//
	// If @max_columns is smaller than the minimum set via
	// [method@Gtk.GridView.set_min_columns], that value is used instead.
	SetMaxColumns(maxColumns uint)
	// SetMinColumns sets the minimum number of columns to use.
	//
	// This number must be at least 1.
	//
	// If @min_columns is smaller than the minimum set via
	// [method@Gtk.GridView.set_max_columns], that value is ignored.
	SetMinColumns(minColumns uint)
	// SetModel sets the imodel to use.
	//
	// This must be a [iface@Gtk.SelectionModel].
	SetModel(model SelectionModel)
	// SetSingleClickActivate sets whether items should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(singleClickActivate bool)
}

// gridView implements the GridView interface.
type gridView struct {
	ListBase
	Accessible
	Buildable
	ConstraintTarget
	Orientable
	Scrollable
}

var _ GridView = (*gridView)(nil)

// WrapGridView wraps a GObject to the right type. It is
// primarily used internally.
func WrapGridView(obj *externglib.Object) GridView {
	return GridView{
		ListBase:         WrapListBase(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
		Scrollable:       WrapScrollable(obj),
	}
}

func marshalGridView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGridView(obj), nil
}

// NewGridView constructs a class GridView.
func NewGridView(model SelectionModel, factory ListItemFactory) GridView {
	var arg1 *C.GtkSelectionModel
	var arg2 *C.GtkListItemFactory

	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	var cret C.GtkGridView
	var goret GridView

	cret = C.gtk_grid_view_new(arg1, arg2)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(GridView)

	return goret
}

// EnableRubberband returns whether rows can be selected by dragging with
// the mouse.
func (s gridView) EnableRubberband() bool {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_grid_view_get_enable_rubberband(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Factory gets the factory that's currently used to populate list items.
func (s gridView) Factory() ListItemFactory {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret *C.GtkListItemFactory
	var goret ListItemFactory

	cret = C.gtk_grid_view_get_factory(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListItemFactory)

	return goret
}

// MaxColumns gets the maximum number of columns that the grid will use.
func (s gridView) MaxColumns() uint {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_grid_view_get_max_columns(arg0)

	goret = uint(cret)

	return goret
}

// MinColumns gets the minimum number of columns that the grid will use.
func (s gridView) MinColumns() uint {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret uint

	cret = C.gtk_grid_view_get_min_columns(arg0)

	goret = uint(cret)

	return goret
}

// Model gets the model that's currently used to read the items displayed.
func (s gridView) Model() SelectionModel {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret *C.GtkSelectionModel
	var goret SelectionModel

	cret = C.gtk_grid_view_get_model(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(SelectionModel)

	return goret
}

// SingleClickActivate returns whether items will be activated on single
// click and selected on hover.
func (s gridView) SingleClickActivate() bool {
	var arg0 *C.GtkGridView

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_grid_view_get_single_click_activate(arg0)

	if cret {
		goret = true
	}

	return goret
}

// SetEnableRubberband sets whether selections can be changed by dragging
// with the mouse.
func (s gridView) SetEnableRubberband(enableRubberband bool) {
	var arg0 *C.GtkGridView
	var arg1 C.gboolean

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		arg1 = C.gboolean(1)
	}

	C.gtk_grid_view_set_enable_rubberband(arg0, arg1)
}

// SetFactory sets the `GtkListItemFactory` to use for populating list
// items.
func (s gridView) SetFactory(factory ListItemFactory) {
	var arg0 *C.GtkGridView
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_grid_view_set_factory(arg0, arg1)
}

// SetMaxColumns sets the maximum number of columns to use.
//
// This number must be at least 1.
//
// If @max_columns is smaller than the minimum set via
// [method@Gtk.GridView.set_min_columns], that value is used instead.
func (s gridView) SetMaxColumns(maxColumns uint) {
	var arg0 *C.GtkGridView
	var arg1 C.guint

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(maxColumns)

	C.gtk_grid_view_set_max_columns(arg0, arg1)
}

// SetMinColumns sets the minimum number of columns to use.
//
// This number must be at least 1.
//
// If @min_columns is smaller than the minimum set via
// [method@Gtk.GridView.set_max_columns], that value is ignored.
func (s gridView) SetMinColumns(minColumns uint) {
	var arg0 *C.GtkGridView
	var arg1 C.guint

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(minColumns)

	C.gtk_grid_view_set_min_columns(arg0, arg1)
}

// SetModel sets the imodel to use.
//
// This must be a [iface@Gtk.SelectionModel].
func (s gridView) SetModel(model SelectionModel) {
	var arg0 *C.GtkGridView
	var arg1 *C.GtkSelectionModel

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_grid_view_set_model(arg0, arg1)
}

// SetSingleClickActivate sets whether items should be activated on single
// click and selected on hover.
func (s gridView) SetSingleClickActivate(singleClickActivate bool) {
	var arg0 *C.GtkGridView
	var arg1 C.gboolean

	arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		arg1 = C.gboolean(1)
	}

	C.gtk_grid_view_set_single_click_activate(arg0, arg1)
}
