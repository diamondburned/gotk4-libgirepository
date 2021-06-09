// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_tool_item_group_get_type()), F: marshalToolItemGroup},
	})
}

// ToolItemGroup: a ToolItemGroup is used together with ToolPalette to add
// ToolItems to a palette like container with different categories and drag and
// drop support.
//
//
// CSS nodes
//
// GtkToolItemGroup has a single CSS node named toolitemgroup.
type ToolItemGroup interface {
	Container
	Buildable
	ToolShell

	// Collapsed gets whether @group is collapsed or expanded.
	Collapsed() bool
	// DropItem gets the tool item at position (x, y).
	DropItem(x int, y int) ToolItem
	// Ellipsize gets the ellipsization mode of @group.
	Ellipsize() pango.EllipsizeMode
	// HeaderRelief gets the relief mode of the header button of @group.
	HeaderRelief() ReliefStyle
	// ItemPosition gets the position of @item in @group as index.
	ItemPosition(item ToolItem) int
	// Label gets the label of @group.
	Label() string
	// LabelWidget gets the label widget of @group. See
	// gtk_tool_item_group_set_label_widget().
	LabelWidget() Widget
	// NItems gets the number of tool items in @group.
	NItems() uint
	// NthItem gets the tool item at @index in group.
	NthItem(index uint) ToolItem
	// Insert inserts @item at @position in the list of children of @group.
	Insert(item ToolItem, position int)
	// SetCollapsed sets whether the @group should be collapsed or expanded.
	SetCollapsed(collapsed bool)
	// SetEllipsize sets the ellipsization mode which should be used by labels
	// in @group.
	SetEllipsize(ellipsize pango.EllipsizeMode)
	// SetHeaderRelief: set the button relief of the group header. See
	// gtk_button_set_relief() for details.
	SetHeaderRelief(style ReliefStyle)
	// SetItemPosition sets the position of @item in the list of children of
	// @group.
	SetItemPosition(item ToolItem, position int)
	// SetLabel sets the label of the tool item group. The label is displayed in
	// the header of the group.
	SetLabel(label string)
	// SetLabelWidget sets the label of the tool item group. The label widget is
	// displayed in the header of the group, in place of the usual label.
	SetLabelWidget(labelWidget Widget)
}

// toolItemGroup implements the ToolItemGroup interface.
type toolItemGroup struct {
	Container
	Buildable
	ToolShell
}

var _ ToolItemGroup = (*toolItemGroup)(nil)

// WrapToolItemGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolItemGroup(obj *externglib.Object) ToolItemGroup {
	return ToolItemGroup{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
		ToolShell: WrapToolShell(obj),
	}
}

func marshalToolItemGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolItemGroup(obj), nil
}

// NewToolItemGroup constructs a class ToolItemGroup.
func NewToolItemGroup(label string) ToolItemGroup {
	var _arg1 *C.gchar

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkToolItemGroup

	cret = C.gtk_tool_item_group_new(_arg1)

	var _toolItemGroup ToolItemGroup

	_toolItemGroup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToolItemGroup)

	return _toolItemGroup
}

// Collapsed gets whether @group is collapsed or expanded.
func (g toolItemGroup) Collapsed() bool {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean

	cret = C.gtk_tool_item_group_get_collapsed(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// DropItem gets the tool item at position (x, y).
func (g toolItemGroup) DropItem(x int, y int) ToolItem {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 C.gint
	var _arg2 C.gint

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	var _cret *C.GtkToolItem

	cret = C.gtk_tool_item_group_get_drop_item(_arg0, _arg1, _arg2)

	var _toolItem ToolItem

	_toolItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToolItem)

	return _toolItem
}

// Ellipsize gets the ellipsization mode of @group.
func (g toolItemGroup) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret C.PangoEllipsizeMode

	cret = C.gtk_tool_item_group_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// HeaderRelief gets the relief mode of the header button of @group.
func (g toolItemGroup) HeaderRelief() ReliefStyle {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret C.GtkReliefStyle

	cret = C.gtk_tool_item_group_get_header_relief(_arg0)

	var _reliefStyle ReliefStyle

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// ItemPosition gets the position of @item in @group as index.
func (g toolItemGroup) ItemPosition(item ToolItem) int {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 *C.GtkToolItem

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	var _cret C.gint

	cret = C.gtk_tool_item_group_get_item_position(_arg0, _arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Label gets the label of @group.
func (g toolItemGroup) Label() string {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret *C.gchar

	cret = C.gtk_tool_item_group_get_label(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// LabelWidget gets the label widget of @group. See
// gtk_tool_item_group_set_label_widget().
func (g toolItemGroup) LabelWidget() Widget {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_tool_item_group_get_label_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// NItems gets the number of tool items in @group.
func (g toolItemGroup) NItems() uint {
	var _arg0 *C.GtkToolItemGroup

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var _cret C.guint

	cret = C.gtk_tool_item_group_get_n_items(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// NthItem gets the tool item at @index in group.
func (g toolItemGroup) NthItem(index uint) ToolItem {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 C.guint

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(index)

	var _cret *C.GtkToolItem

	cret = C.gtk_tool_item_group_get_nth_item(_arg0, _arg1)

	var _toolItem ToolItem

	_toolItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToolItem)

	return _toolItem
}

// Insert inserts @item at @position in the list of children of @group.
func (g toolItemGroup) Insert(item ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 *C.GtkToolItem
	var _arg2 C.gint

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_insert(_arg0, _arg1, _arg2)
}

// SetCollapsed sets whether the @group should be collapsed or expanded.
func (g toolItemGroup) SetCollapsed(collapsed bool) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 C.gboolean

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	if collapsed {
		_arg1 = C.gboolean(1)
	}

	C.gtk_tool_item_group_set_collapsed(_arg0, _arg1)
}

// SetEllipsize sets the ellipsization mode which should be used by labels
// in @group.
func (g toolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 C.PangoEllipsizeMode

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (C.PangoEllipsizeMode)(ellipsize)

	C.gtk_tool_item_group_set_ellipsize(_arg0, _arg1)
}

// SetHeaderRelief: set the button relief of the group header. See
// gtk_button_set_relief() for details.
func (g toolItemGroup) SetHeaderRelief(style ReliefStyle) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 C.GtkReliefStyle

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (C.GtkReliefStyle)(style)

	C.gtk_tool_item_group_set_header_relief(_arg0, _arg1)
}

// SetItemPosition sets the position of @item in the list of children of
// @group.
func (g toolItemGroup) SetItemPosition(item ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 *C.GtkToolItem
	var _arg2 C.gint

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_set_item_position(_arg0, _arg1, _arg2)
}

// SetLabel sets the label of the tool item group. The label is displayed in
// the header of the group.
func (g toolItemGroup) SetLabel(label string) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 *C.gchar

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_group_set_label(_arg0, _arg1)
}

// SetLabelWidget sets the label of the tool item group. The label widget is
// displayed in the header of the group, in place of the usual label.
func (g toolItemGroup) SetLabelWidget(labelWidget Widget) {
	var _arg0 *C.GtkToolItemGroup
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_item_group_set_label_widget(_arg0, _arg1)
}