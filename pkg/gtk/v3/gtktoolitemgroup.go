// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_item_group_get_type()), F: marshalToolItemGrouper},
	})
}

// ToolItemGroup is used together with ToolPalette to add ToolItems to a palette
// like container with different categories and drag and drop support.
//
//
// CSS nodes
//
// GtkToolItemGroup has a single CSS node named toolitemgroup.
type ToolItemGroup struct {
	Container

	ToolShell
	*externglib.Object
}

func wrapToolItemGroup(obj *externglib.Object) *ToolItemGroup {
	return &ToolItemGroup{
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
		ToolShell: ToolShell{
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
		Object: obj,
	}
}

func marshalToolItemGrouper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToolItemGroup(obj), nil
}

// NewToolItemGroup creates a new tool item group with label label.
func NewToolItemGroup(label string) *ToolItemGroup {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_tool_item_group_new(_arg1)

	var _toolItemGroup *ToolItemGroup // out

	_toolItemGroup = wrapToolItemGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItemGroup
}

// Collapsed gets whether group is collapsed or expanded.
func (group *ToolItemGroup) Collapsed() bool {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_collapsed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DropItem gets the tool item at position (x, y).
func (group *ToolItemGroup) DropItem(x int, y int) *ToolItem {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.gint              // out
	var _arg2 C.gint              // out
	var _cret *C.GtkToolItem      // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_item_group_get_drop_item(_arg0, _arg1, _arg2)

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// Ellipsize gets the ellipsization mode of group.
func (group *ToolItemGroup) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkToolItemGroup  // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// HeaderRelief gets the relief mode of the header button of group.
func (group *ToolItemGroup) HeaderRelief() ReliefStyle {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.GtkReliefStyle    // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_header_relief(_arg0)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// ItemPosition gets the position of item in group as index.
func (group *ToolItemGroup) ItemPosition(item *ToolItem) int {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	_cret = C.gtk_tool_item_group_get_item_position(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Label gets the label of group.
func (group *ToolItemGroup) Label() string {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LabelWidget gets the label widget of group. See
// gtk_tool_item_group_set_label_widget().
func (group *ToolItemGroup) LabelWidget() Widgetter {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_label_widget(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// NItems gets the number of tool items in group.
func (group *ToolItemGroup) NItems() uint {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_n_items(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NthItem gets the tool item at index in group.
func (group *ToolItemGroup) NthItem(index uint) *ToolItem {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.guint             // out
	var _cret *C.GtkToolItem      // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.guint(index)

	_cret = C.gtk_tool_item_group_get_nth_item(_arg0, _arg1)

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// Insert inserts item at position in the list of children of group.
func (group *ToolItemGroup) Insert(item *ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_insert(_arg0, _arg1, _arg2)
}

// SetCollapsed sets whether the group should be collapsed or expanded.
func (group *ToolItemGroup) SetCollapsed(collapsed bool) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if collapsed {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_group_set_collapsed(_arg0, _arg1)
}

// SetEllipsize sets the ellipsization mode which should be used by labels in
// group.
func (group *ToolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	var _arg0 *C.GtkToolItemGroup  // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.PangoEllipsizeMode(ellipsize)

	C.gtk_tool_item_group_set_ellipsize(_arg0, _arg1)
}

// SetHeaderRelief: set the button relief of the group header. See
// gtk_button_set_relief() for details.
func (group *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.GtkReliefStyle    // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.GtkReliefStyle(style)

	C.gtk_tool_item_group_set_header_relief(_arg0, _arg1)
}

// SetItemPosition sets the position of item in the list of children of group.
func (group *ToolItemGroup) SetItemPosition(item *ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_set_item_position(_arg0, _arg1, _arg2)
}

// SetLabel sets the label of the tool item group. The label is displayed in the
// header of the group.
func (group *ToolItemGroup) SetLabel(label string) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_group_set_label(_arg0, _arg1)
}

// SetLabelWidget sets the label of the tool item group. The label widget is
// displayed in the header of the group, in place of the usual label.
func (group *ToolItemGroup) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_item_group_set_label_widget(_arg0, _arg1)
}
