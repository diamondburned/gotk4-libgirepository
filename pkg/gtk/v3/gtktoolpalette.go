// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_tool_palette_drag_targets_get_type()), F: marshalToolPaletteDragTargets},
		{T: externglib.Type(C.gtk_tool_palette_get_type()), F: marshalToolPaletter},
	})
}

// ToolPaletteDragTargets flags used to specify the supported drag targets.
type ToolPaletteDragTargets int

const (
	// ToolPaletteDragTargetsItems: support drag of items.
	ToolPaletteDragTargetsItems ToolPaletteDragTargets = 0b1
	// ToolPaletteDragTargetsGroups: support drag of groups.
	ToolPaletteDragTargetsGroups ToolPaletteDragTargets = 0b10
)

func marshalToolPaletteDragTargets(p uintptr) (interface{}, error) {
	return ToolPaletteDragTargets(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ToolPaletter describes ToolPalette's methods.
type ToolPaletter interface {
	// DragItem: get the dragged item from the selection.
	DragItem(selection *SelectionData) *Widget
	// DropGroup gets the group at position (x, y).
	DropGroup(x int, y int) *ToolItemGroup
	// DropItem gets the item at position (x, y).
	DropItem(x int, y int) *ToolItem
	// Exclusive gets whether @group is exclusive or not.
	Exclusive(group ToolItemGrouper) bool
	// Expand gets whether group should be given extra space.
	Expand(group ToolItemGrouper) bool
	// GroupPosition gets the position of @group in @palette as index.
	GroupPosition(group ToolItemGrouper) int
	// HAdjustment gets the horizontal adjustment of the tool palette.
	HAdjustment() *Adjustment
	// IconSize gets the size of icons in the tool palette.
	IconSize() int
	// Style gets the style (icons, text or both) of items in the tool palette.
	Style() ToolbarStyle
	// VAdjustment gets the vertical adjustment of the tool palette.
	VAdjustment() *Adjustment
	// SetExclusive sets whether the group should be exclusive or not.
	SetExclusive(group ToolItemGrouper, exclusive bool)
	// SetExpand sets whether the group should be given extra space.
	SetExpand(group ToolItemGrouper, expand bool)
	// SetGroupPosition sets the position of the group as an index of the tool
	// palette.
	SetGroupPosition(group ToolItemGrouper, position int)
	// SetIconSize sets the size of icons in the tool palette.
	SetIconSize(iconSize int)
	// UnsetIconSize unsets the tool palette icon size set with
	// gtk_tool_palette_set_icon_size(), so that user preferences will be used
	// to determine the icon size.
	UnsetIconSize()
	// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(),
	// so that user preferences will be used to determine the toolbar style.
	UnsetStyle()
}

// ToolPalette allows you to add ToolItems to a palette-like container with
// different categories and drag and drop support.
//
// A ToolPalette is created with a call to gtk_tool_palette_new().
//
// ToolItems cannot be added directly to a ToolPalette - instead they are added
// to a ToolItemGroup which can than be added to a ToolPalette. To add a
// ToolItemGroup to a ToolPalette, use gtk_container_add().
//
//    static void
//    passive_canvas_drag_data_received (GtkWidget        *widget,
//                                       GdkDragContext   *context,
//                                       gint              x,
//                                       gint              y,
//                                       GtkSelectionData *selection,
//                                       guint             info,
//                                       guint             time,
//                                       gpointer          data)
//    {
//      GtkWidget *palette;
//      GtkWidget *item;
//
//      // Get the dragged item
//      palette = gtk_widget_get_ancestor (gtk_drag_get_source_widget (context),
//                                         GTK_TYPE_TOOL_PALETTE);
//      if (palette != NULL)
//        item = gtk_tool_palette_get_drag_item (GTK_TOOL_PALETTE (palette),
//                                               selection);
//
//      // Do something with item
//    }
//
//    GtkWidget *target, palette;
//
//    palette = gtk_tool_palette_new ();
//    target = gtk_drawing_area_new ();
//
//    g_signal_connect (G_OBJECT (target), "drag-data-received",
//                      G_CALLBACK (passive_canvas_drag_data_received), NULL);
//    gtk_tool_palette_add_drag_dest (GTK_TOOL_PALETTE (palette), target,
//                                    GTK_DEST_DEFAULT_ALL,
//                                    GTK_TOOL_PALETTE_DRAG_ITEMS,
//                                    GDK_ACTION_COPY);
//
//
// CSS nodes
//
// GtkToolPalette has a single CSS node named toolpalette.
type ToolPalette struct {
	Container

	Orientable
	Scrollable
}

var (
	_ ToolPaletter    = (*ToolPalette)(nil)
	_ gextras.Nativer = (*ToolPalette)(nil)
)

func wrapToolPalette(obj *externglib.Object) ToolPaletter {
	return &ToolPalette{
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
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalToolPaletter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToolPalette(obj), nil
}

// NewToolPalette creates a new tool palette.
func NewToolPalette() *ToolPalette {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tool_palette_new()

	var _toolPalette *ToolPalette // out

	_toolPalette = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ToolPalette)

	return _toolPalette
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ToolPalette) Native() uintptr {
	return v.Container.Widget.InitiallyUnowned.Object.Native()
}

// DragItem: get the dragged item from the selection. This could be a ToolItem
// or a ToolItemGroup.
func (palette *ToolPalette) DragItem(selection *SelectionData) *Widget {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkSelectionData // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selection))

	_cret = C.gtk_tool_palette_get_drag_item(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// DropGroup gets the group at position (x, y).
func (palette *ToolPalette) DropGroup(x int, y int) *ToolItemGroup {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 C.gint              // out
	var _arg2 C.gint              // out
	var _cret *C.GtkToolItemGroup // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_palette_get_drop_group(_arg0, _arg1, _arg2)

	var _toolItemGroup *ToolItemGroup // out

	_toolItemGroup = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ToolItemGroup)

	return _toolItemGroup
}

// DropItem gets the item at position (x, y). See
// gtk_tool_palette_get_drop_group().
func (palette *ToolPalette) DropItem(x int, y int) *ToolItem {
	var _arg0 *C.GtkToolPalette // out
	var _arg1 C.gint            // out
	var _arg2 C.gint            // out
	var _cret *C.GtkToolItem    // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_palette_get_drop_item(_arg0, _arg1, _arg2)

	var _toolItem *ToolItem // out

	_toolItem = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ToolItem)

	return _toolItem
}

// Exclusive gets whether @group is exclusive or not. See
// gtk_tool_palette_set_exclusive().
func (palette *ToolPalette) Exclusive(group ToolItemGrouper) bool {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))

	_cret = C.gtk_tool_palette_get_exclusive(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expand gets whether group should be given extra space. See
// gtk_tool_palette_set_expand().
func (palette *ToolPalette) Expand(group ToolItemGrouper) bool {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))

	_cret = C.gtk_tool_palette_get_expand(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GroupPosition gets the position of @group in @palette as index. See
// gtk_tool_palette_set_group_position().
func (palette *ToolPalette) GroupPosition(group ToolItemGrouper) int {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))

	_cret = C.gtk_tool_palette_get_group_position(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HAdjustment gets the horizontal adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
func (palette *ToolPalette) HAdjustment() *Adjustment {
	var _arg0 *C.GtkToolPalette // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_hadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// IconSize gets the size of icons in the tool palette. See
// gtk_tool_palette_set_icon_size().
func (palette *ToolPalette) IconSize() int {
	var _arg0 *C.GtkToolPalette // out
	var _cret C.GtkIconSize     // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_icon_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Style gets the style (icons, text or both) of items in the tool palette.
func (palette *ToolPalette) Style() ToolbarStyle {
	var _arg0 *C.GtkToolPalette // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_style(_arg0)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// VAdjustment gets the vertical adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
func (palette *ToolPalette) VAdjustment() *Adjustment {
	var _arg0 *C.GtkToolPalette // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_vadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// SetExclusive sets whether the group should be exclusive or not. If an
// exclusive group is expanded all other groups are collapsed.
func (palette *ToolPalette) SetExclusive(group ToolItemGrouper, exclusive bool) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))
	if exclusive {
		_arg2 = C.TRUE
	}

	C.gtk_tool_palette_set_exclusive(_arg0, _arg1, _arg2)
}

// SetExpand sets whether the group should be given extra space.
func (palette *ToolPalette) SetExpand(group ToolItemGrouper, expand bool) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tool_palette_set_expand(_arg0, _arg1, _arg2)
}

// SetGroupPosition sets the position of the group as an index of the tool
// palette. If position is 0 the group will become the first child, if position
// is -1 it will become the last child.
func (palette *ToolPalette) SetGroupPosition(group ToolItemGrouper, position int) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_palette_set_group_position(_arg0, _arg1, _arg2)
}

// SetIconSize sets the size of icons in the tool palette.
func (palette *ToolPalette) SetIconSize(iconSize int) {
	var _arg0 *C.GtkToolPalette // out
	var _arg1 C.GtkIconSize     // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_tool_palette_set_icon_size(_arg0, _arg1)
}

// UnsetIconSize unsets the tool palette icon size set with
// gtk_tool_palette_set_icon_size(), so that user preferences will be used to
// determine the icon size.
func (palette *ToolPalette) UnsetIconSize() {
	var _arg0 *C.GtkToolPalette // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	C.gtk_tool_palette_unset_icon_size(_arg0)
}

// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(), so
// that user preferences will be used to determine the toolbar style.
func (palette *ToolPalette) UnsetStyle() {
	var _arg0 *C.GtkToolPalette // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	C.gtk_tool_palette_unset_style(_arg0)
}
