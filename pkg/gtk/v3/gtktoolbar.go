// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_toolbar_space_style_get_type()), F: marshalToolbarSpaceStyle},
		{T: externglib.Type(C.gtk_toolbar_get_type()), F: marshalToolbarrer},
	})
}

// ToolbarSpaceStyle: whether spacers are vertical lines or just blank.
//
// Deprecated: since version 3.20.
type ToolbarSpaceStyle int

const (
	// ToolbarSpaceEmpty: use blank spacers.
	ToolbarSpaceEmpty ToolbarSpaceStyle = iota
	// ToolbarSpaceLine: use vertical lines for spacers.
	ToolbarSpaceLine
)

func marshalToolbarSpaceStyle(p uintptr) (interface{}, error) {
	return ToolbarSpaceStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ToolbarSpaceStyle.
func (t ToolbarSpaceStyle) String() string {
	switch t {
	case ToolbarSpaceEmpty:
		return "Empty"
	case ToolbarSpaceLine:
		return "Line"
	default:
		return fmt.Sprintf("ToolbarSpaceStyle(%d)", t)
	}
}

// ToolbarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToolbarOverrider interface {
	OrientationChanged(orientation Orientation)
	PopupContextMenu(x int, y int, buttonNumber int) bool
	StyleChanged(style ToolbarStyle)
}

// Toolbar: toolbar is created with a call to gtk_toolbar_new().
//
// A toolbar can contain instances of a subclass of ToolItem. To add a ToolItem
// to the a toolbar, use gtk_toolbar_insert(). To remove an item from the
// toolbar use gtk_container_remove(). To add a button to the toolbar, add an
// instance of ToolButton.
//
// Toolbar items can be visually grouped by adding instances of
// SeparatorToolItem to the toolbar. If the GtkToolbar child property “expand”
// is UE and the property SeparatorToolItem:draw is set to LSE, the effect is to
// force all following items to the end of the toolbar.
//
// By default, a toolbar can be shrunk, upon which it will add an arrow button
// to show an overflow menu offering access to any ToolItem child that has a
// proxy menu item. To disable this and request enough size for all children,
// call gtk_toolbar_set_show_arrow() to set Toolbar:show-arrow to FALSE.
//
// Creating a context menu for the toolbar can be done by connecting to the
// Toolbar::popup-context-menu signal.
//
//
// CSS nodes
//
// GtkToolbar has a single CSS node with name toolbar.
type Toolbar struct {
	Container

	Orientable
	ToolShell
	*externglib.Object
}

func wrapToolbar(obj *externglib.Object) *Toolbar {
	return &Toolbar{
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
		Orientable: Orientable{
			Object: obj,
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

func marshalToolbarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToolbar(obj), nil
}

// NewToolbar creates a new toolbar.
func NewToolbar() *Toolbar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toolbar_new()

	var _toolbar *Toolbar // out

	_toolbar = wrapToolbar(externglib.Take(unsafe.Pointer(_cret)))

	return _toolbar
}

// DropIndex returns the position corresponding to the indicated point on
// toolbar. This is useful when dragging items to the toolbar: this function
// returns the position a new item should be inserted.
//
// x and y are in toolbar coordinates.
func (toolbar *Toolbar) DropIndex(x int, y int) int {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.gint        // out
	var _arg2 C.gint        // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_toolbar_get_drop_index(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IconSize retrieves the icon size for the toolbar. See
// gtk_toolbar_set_icon_size().
func (toolbar *Toolbar) IconSize() IconSize {
	var _arg0 *C.GtkToolbar // out
	var _cret C.GtkIconSize // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	_cret = C.gtk_toolbar_get_icon_size(_arg0)

	var _iconSize IconSize // out

	_iconSize = IconSize(_cret)

	return _iconSize
}

// ItemIndex returns the position of item on the toolbar, starting from 0. It is
// an error if item is not a child of the toolbar.
func (toolbar *Toolbar) ItemIndex(item *ToolItem) int {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	_cret = C.gtk_toolbar_get_item_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NItems returns the number of items on the toolbar.
func (toolbar *Toolbar) NItems() int {
	var _arg0 *C.GtkToolbar // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	_cret = C.gtk_toolbar_get_n_items(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthItem returns the n'th item on toolbar, or NULL if the toolbar does not
// contain an n'th item.
func (toolbar *Toolbar) NthItem(n int) *ToolItem {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 C.gint         // out
	var _cret *C.GtkToolItem // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = C.gint(n)

	_cret = C.gtk_toolbar_get_nth_item(_arg0, _arg1)

	var _toolItem *ToolItem // out

	if _cret != nil {
		_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItem
}

// ReliefStyle returns the relief style of buttons on toolbar. See
// gtk_button_set_relief().
func (toolbar *Toolbar) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolbar    // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	_cret = C.gtk_toolbar_get_relief_style(_arg0)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// ShowArrow returns whether the toolbar has an overflow menu. See
// gtk_toolbar_set_show_arrow().
func (toolbar *Toolbar) ShowArrow() bool {
	var _arg0 *C.GtkToolbar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	_cret = C.gtk_toolbar_get_show_arrow(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Style retrieves whether the toolbar has text, icons, or both . See
// gtk_toolbar_set_style().
func (toolbar *Toolbar) Style() ToolbarStyle {
	var _arg0 *C.GtkToolbar     // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	_cret = C.gtk_toolbar_get_style(_arg0)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// Insert a ToolItem into the toolbar at position pos. If pos is 0 the item is
// prepended to the start of the toolbar. If pos is negative, the item is
// appended to the end of the toolbar.
func (toolbar *Toolbar) Insert(item *ToolItem, pos int) {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(pos)

	C.gtk_toolbar_insert(_arg0, _arg1, _arg2)
}

// SetDropHighlightItem highlights toolbar to give an idea of what it would look
// like if item was added to toolbar at the position indicated by index_. If
// item is NULL, highlighting is turned off. In that case index_ is ignored.
//
// The tool_item passed to this function must not be part of any widget
// hierarchy. When an item is set as drop highlight item it can not added to any
// widget hierarchy or used as highlight item for another toolbar.
func (toolbar *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index_ int) {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	if toolItem != nil {
		_arg1 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	}
	_arg2 = C.gint(index_)

	C.gtk_toolbar_set_drop_highlight_item(_arg0, _arg1, _arg2)
}

// SetIconSize: this function sets the size of stock icons in the toolbar. You
// can call it both before you add the icons and after they’ve been added. The
// size you set will override user preferences for the default icon size.
//
// This should only be used for special-purpose toolbars, normal application
// toolbars should respect the user preferences for the size of icons.
func (toolbar *Toolbar) SetIconSize(iconSize IconSize) {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.GtkIconSize // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_toolbar_set_icon_size(_arg0, _arg1)
}

// SetShowArrow sets whether to show an overflow menu when toolbar isn’t
// allocated enough size to show all of its items. If TRUE, items which can’t
// fit in toolbar, and which have a proxy menu item set by
// gtk_tool_item_set_proxy_menu_item() or ToolItem::create-menu-proxy, will be
// available in an overflow menu, which can be opened by an added arrow button.
// If FALSE, toolbar will request enough size to fit all of its child items
// without any overflow.
func (toolbar *Toolbar) SetShowArrow(showArrow bool) {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	if showArrow {
		_arg1 = C.TRUE
	}

	C.gtk_toolbar_set_show_arrow(_arg0, _arg1)
}

// SetStyle alters the view of toolbar to display either icons only, text only,
// or both.
func (toolbar *Toolbar) SetStyle(style ToolbarStyle) {
	var _arg0 *C.GtkToolbar     // out
	var _arg1 C.GtkToolbarStyle // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))
	_arg1 = C.GtkToolbarStyle(style)

	C.gtk_toolbar_set_style(_arg0, _arg1)
}

// UnsetIconSize unsets toolbar icon size set with gtk_toolbar_set_icon_size(),
// so that user preferences will be used to determine the icon size.
func (toolbar *Toolbar) UnsetIconSize() {
	var _arg0 *C.GtkToolbar // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	C.gtk_toolbar_unset_icon_size(_arg0)
}

// UnsetStyle unsets a toolbar style set with gtk_toolbar_set_style(), so that
// user preferences will be used to determine the toolbar style.
func (toolbar *Toolbar) UnsetStyle() {
	var _arg0 *C.GtkToolbar // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(toolbar.Native()))

	C.gtk_toolbar_unset_style(_arg0)
}
