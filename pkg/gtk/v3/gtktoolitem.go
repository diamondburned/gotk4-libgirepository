// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_item_get_type()), F: marshalToolItem},
	})
}

// ToolItem ToolItems are widgets that can appear on a toolbar. To create a
// toolbar item that contain something else than a button, use
// gtk_tool_item_new(). Use gtk_container_add() to add a child widget to the
// tool item.
//
// For toolbar items that contain buttons, see the ToolButton, ToggleToolButton
// and RadioToolButton classes.
//
// See the Toolbar class for a description of the toolbar widget, and ToolShell
// for a description of the tool shell interface.
type ToolItem interface {
	Bin
	Activatable
	Buildable

	// Expand returns whether @tool_item is allocated extra space. See
	// gtk_tool_item_set_expand().
	Expand() bool
	// Homogeneous returns whether @tool_item is the same size as other
	// homogeneous items. See gtk_tool_item_set_homogeneous().
	Homogeneous() bool
	// IconSize returns the icon size used for @tool_item. Custom subclasses of
	// ToolItem should call this function to find out what size icons they
	// should use.
	IconSize() int
	// IsImportant returns whether @tool_item is considered important. See
	// gtk_tool_item_set_is_important()
	IsImportant() bool
	// TextAlignment returns the text alignment used for @tool_item. Custom
	// subclasses of ToolItem should call this function to find out how text
	// should be aligned.
	TextAlignment() float32
	// UseDragWindow returns whether @tool_item has a drag window. See
	// gtk_tool_item_set_use_drag_window().
	UseDragWindow() bool
	// VisibleHorizontal returns whether the @tool_item is visible on toolbars
	// that are docked horizontally.
	VisibleHorizontal() bool
	// VisibleVertical returns whether @tool_item is visible when the toolbar is
	// docked vertically. See gtk_tool_item_set_visible_vertical().
	VisibleVertical() bool
	// RebuildMenu: calling this function signals to the toolbar that the
	// overflow menu item for @tool_item has changed. If the overflow menu is
	// visible when this function it called, the menu will be rebuilt.
	//
	// The function must be called when the tool item changes what it will do in
	// response to the ToolItem::create-menu-proxy signal.
	RebuildMenu()
	// SetExpand sets whether @tool_item is allocated extra space when there is
	// more room on the toolbar then needed for the items. The effect is that
	// the item gets bigger when the toolbar gets bigger and smaller when the
	// toolbar gets smaller.
	SetExpand(expand bool)
	// SetHomogeneous sets whether @tool_item is to be allocated the same size
	// as other homogeneous items. The effect is that all homogeneous items will
	// have the same width as the widest of the items.
	SetHomogeneous(homogeneous bool)
	// SetIsImportant sets whether @tool_item should be considered important.
	// The ToolButton class uses this property to determine whether to show or
	// hide its label when the toolbar style is GTK_TOOLBAR_BOTH_HORIZ. The
	// result is that only tool buttons with the “is_important” property set
	// have labels, an effect known as “priority text”
	SetIsImportant(isImportant bool)
	// SetProXYMenuItem sets the MenuItem used in the toolbar overflow menu. The
	// @menu_item_id is used to identify the caller of this function and should
	// also be used with gtk_tool_item_get_proxy_menu_item().
	//
	// See also ToolItem::create-menu-proxy.
	SetProXYMenuItem(menuItemId string, menuItem Widget)
	// SetTooltipMarkup sets the markup text to be displayed as tooltip on the
	// item. See gtk_widget_set_tooltip_markup().
	SetTooltipMarkup(markup string)
	// SetTooltipText sets the text to be displayed as tooltip on the item. See
	// gtk_widget_set_tooltip_text().
	SetTooltipText(text string)
	// SetUseDragWindow sets whether @tool_item has a drag window. When true the
	// toolitem can be used as a drag source through gtk_drag_source_set(). When
	// @tool_item has a drag window it will intercept all events, even those
	// that would otherwise be sent to a child of @tool_item.
	SetUseDragWindow(useDragWindow bool)
	// SetVisibleHorizontal sets whether @tool_item is visible when the toolbar
	// is docked horizontally.
	SetVisibleHorizontal(visibleHorizontal bool)
	// SetVisibleVertical sets whether @tool_item is visible when the toolbar is
	// docked vertically. Some tool items, such as text entries, are too wide to
	// be useful on a vertically docked toolbar. If @visible_vertical is false
	// @tool_item will not appear on toolbars that are docked vertically.
	SetVisibleVertical(visibleVertical bool)
	// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
	// @tool_item. Toolbar and other ToolShell implementations use this function
	// to notify children, when some aspect of their configuration changes.
	ToolbarReconfigured()
}

// toolItem implements the ToolItem class.
type toolItem struct {
	Bin
	Activatable
	Buildable
}

var _ ToolItem = (*toolItem)(nil)

// WrapToolItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolItem(obj *externglib.Object) ToolItem {
	return toolItem{
		Bin:         WrapBin(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalToolItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolItem(obj), nil
}

// Expand returns whether @tool_item is allocated extra space. See
// gtk_tool_item_set_expand().
func (t toolItem) Expand() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_expand(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous returns whether @tool_item is the same size as other
// homogeneous items. See gtk_tool_item_set_homogeneous().
func (t toolItem) Homogeneous() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconSize returns the icon size used for @tool_item. Custom subclasses of
// ToolItem should call this function to find out what size icons they
// should use.
func (t toolItem) IconSize() int {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.GtkIconSize // in

	_cret = C.gtk_tool_item_get_icon_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// IsImportant returns whether @tool_item is considered important. See
// gtk_tool_item_set_is_important()
func (t toolItem) IsImportant() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_is_important(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextAlignment returns the text alignment used for @tool_item. Custom
// subclasses of ToolItem should call this function to find out how text
// should be aligned.
func (t toolItem) TextAlignment() float32 {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gfloat // in

	_cret = C.gtk_tool_item_get_text_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// UseDragWindow returns whether @tool_item has a drag window. See
// gtk_tool_item_set_use_drag_window().
func (t toolItem) UseDragWindow() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_use_drag_window(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleHorizontal returns whether the @tool_item is visible on toolbars
// that are docked horizontally.
func (t toolItem) VisibleHorizontal() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_visible_horizontal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleVertical returns whether @tool_item is visible when the toolbar is
// docked vertically. See gtk_tool_item_set_visible_vertical().
func (t toolItem) VisibleVertical() bool {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tool_item_get_visible_vertical(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RebuildMenu: calling this function signals to the toolbar that the
// overflow menu item for @tool_item has changed. If the overflow menu is
// visible when this function it called, the menu will be rebuilt.
//
// The function must be called when the tool item changes what it will do in
// response to the ToolItem::create-menu-proxy signal.
func (t toolItem) RebuildMenu() {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	C.gtk_tool_item_rebuild_menu(_arg0)
}

// SetExpand sets whether @tool_item is allocated extra space when there is
// more room on the toolbar then needed for the items. The effect is that
// the item gets bigger when the toolbar gets bigger and smaller when the
// toolbar gets smaller.
func (t toolItem) SetExpand(expand bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_expand(_arg0, _arg1)
}

// SetHomogeneous sets whether @tool_item is to be allocated the same size
// as other homogeneous items. The effect is that all homogeneous items will
// have the same width as the widest of the items.
func (t toolItem) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_homogeneous(_arg0, _arg1)
}

// SetIsImportant sets whether @tool_item should be considered important.
// The ToolButton class uses this property to determine whether to show or
// hide its label when the toolbar style is GTK_TOOLBAR_BOTH_HORIZ. The
// result is that only tool buttons with the “is_important” property set
// have labels, an effect known as “priority text”
func (t toolItem) SetIsImportant(isImportant bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_is_important(_arg0, _arg1)
}

// SetProXYMenuItem sets the MenuItem used in the toolbar overflow menu. The
// @menu_item_id is used to identify the caller of this function and should
// also be used with gtk_tool_item_get_proxy_menu_item().
//
// See also ToolItem::create-menu-proxy.
func (t toolItem) SetProXYMenuItem(menuItemId string, menuItem Widget) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(menuItemId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))

	C.gtk_tool_item_set_proxy_menu_item(_arg0, _arg1, _arg2)
}

// SetTooltipMarkup sets the markup text to be displayed as tooltip on the
// item. See gtk_widget_set_tooltip_markup().
func (t toolItem) SetTooltipMarkup(markup string) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_set_tooltip_markup(_arg0, _arg1)
}

// SetTooltipText sets the text to be displayed as tooltip on the item. See
// gtk_widget_set_tooltip_text().
func (t toolItem) SetTooltipText(text string) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_set_tooltip_text(_arg0, _arg1)
}

// SetUseDragWindow sets whether @tool_item has a drag window. When true the
// toolitem can be used as a drag source through gtk_drag_source_set(). When
// @tool_item has a drag window it will intercept all events, even those
// that would otherwise be sent to a child of @tool_item.
func (t toolItem) SetUseDragWindow(useDragWindow bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if useDragWindow {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_use_drag_window(_arg0, _arg1)
}

// SetVisibleHorizontal sets whether @tool_item is visible when the toolbar
// is docked horizontally.
func (t toolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_visible_horizontal(_arg0, _arg1)
}

// SetVisibleVertical sets whether @tool_item is visible when the toolbar is
// docked vertically. Some tool items, such as text entries, are too wide to
// be useful on a vertically docked toolbar. If @visible_vertical is false
// @tool_item will not appear on toolbars that are docked vertically.
func (t toolItem) SetVisibleVertical(visibleVertical bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_visible_vertical(_arg0, _arg1)
}

// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
// @tool_item. Toolbar and other ToolShell implementations use this function
// to notify children, when some aspect of their configuration changes.
func (t toolItem) ToolbarReconfigured() {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(t.Native()))

	C.gtk_tool_item_toolbar_reconfigured(_arg0)
}
