// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_menu_get_type()), F: marshalMenu},
	})
}

// MenuPositionFunc: a user function supplied when calling gtk_menu_popup()
// which controls the positioning of the menu when it is displayed. The function
// sets the @x and @y parameters to the coordinates where the menu is to be
// drawn. To make the menu appear on a different monitor than the mouse pointer,
// gtk_menu_set_monitor() must be called.
type MenuPositionFunc func() (pushIn bool)

//export gotk4_MenuPositionFunc
func gotk4_MenuPositionFunc(arg0 *C.GtkMenu, arg1 *C.gint, arg2 *C.gint, arg3 *C.gboolean, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(MenuPositionFunc)
	pushIn := fn()

	if pushIn {
		arg3 = *C.gboolean(1)
	}
}

// Menu: a Menu is a MenuShell that implements a drop down menu consisting of a
// list of MenuItem objects which can be navigated and activated by the user to
// perform application functions.
//
// A Menu is most commonly dropped down by activating a MenuItem in a MenuBar or
// popped up by activating a MenuItem in another Menu.
//
// A Menu can also be popped up by activating a ComboBox. Other composite
// widgets such as the Notebook can pop up a Menu as well.
//
// Applications can display a Menu as a popup menu by calling the
// gtk_menu_popup() function. The example below shows how an application can pop
// up a menu when the 3rd mouse button is pressed.
//
// Connecting the popup signal handler.
//
//    menu
//    ├── arrow.top
//    ├── <child>
//    ┊
//    ├── <child>
//    ╰── arrow.bottom
//
// The main CSS node of GtkMenu has name menu, and there are two subnodes with
// name arrow, for scrolling menu arrows. These subnodes get the .top and
// .bottom style classes.
type Menu interface {
	MenuShell
	Buildable

	// Attach adds a new MenuItem to a (table) menu. The number of “cells” that
	// an item will occupy is specified by @left_attach, @right_attach,
	// @top_attach and @bottom_attach. These each represent the leftmost,
	// rightmost, uppermost and lower column and row numbers of the table.
	// (Columns and rows are indexed from zero).
	//
	// Note that this function is not related to gtk_menu_detach().
	Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint)
	// Detach detaches the menu from the widget to which it had been attached.
	// This function will call the callback function, @detacher, provided when
	// the gtk_menu_attach_to_widget() function was called.
	Detach()
	// AccelGroup gets the AccelGroup which holds global accelerators for the
	// menu. See gtk_menu_set_accel_group().
	AccelGroup() AccelGroup
	// AccelPath retrieves the accelerator path set on the menu.
	AccelPath() string
	// Active returns the selected menu item from the menu. This is used by the
	// ComboBox.
	Active() Widget
	// AttachWidget returns the Widget that the menu is attached to.
	AttachWidget() Widget
	// Monitor retrieves the number of the monitor on which to show the menu.
	Monitor() int
	// ReserveToggleSize returns whether the menu reserves space for toggles and
	// icons, regardless of their actual presence.
	ReserveToggleSize() bool
	// TearoffState returns whether the menu is torn off. See
	// gtk_menu_set_tearoff_state().
	TearoffState() bool
	// Title returns the title of the menu. See gtk_menu_set_title().
	Title() string
	// PlaceOnMonitor places @menu on the given monitor.
	PlaceOnMonitor(monitor gdk.Monitor)
	// Popdown removes the menu from the screen.
	Popdown()
	// PopupAtPointer displays @menu and makes it available for selection.
	//
	// See gtk_menu_popup_at_widget () to pop up a menu at a widget.
	// gtk_menu_popup_at_rect () also allows you to position a menu at an
	// arbitrary rectangle.
	//
	// @menu will be positioned at the pointer associated with @trigger_event.
	//
	// Properties that influence the behaviour of this function are
	// Menu:anchor-hints, Menu:rect-anchor-dx, Menu:rect-anchor-dy, and
	// Menu:menu-type-hint. Connect to the Menu::popped-up signal to find out
	// how it was actually positioned.
	PopupAtPointer(triggerEvent *gdk.Event)
	// PopupAtRect displays @menu and makes it available for selection.
	//
	// See gtk_menu_popup_at_widget () and gtk_menu_popup_at_pointer (), which
	// handle more common cases for popping up menus.
	//
	// @menu will be positioned at @rect, aligning their anchor points. @rect is
	// relative to the top-left corner of @rect_window. @rect_anchor and
	// @menu_anchor determine anchor points on @rect and @menu to pin together.
	// @menu can optionally be offset by Menu:rect-anchor-dx and
	// Menu:rect-anchor-dy.
	//
	// Anchors should be specified under the assumption that the text direction
	// is left-to-right; they will be flipped horizontally automatically if the
	// text direction is right-to-left.
	//
	// Other properties that influence the behaviour of this function are
	// Menu:anchor-hints and Menu:menu-type-hint. Connect to the Menu::popped-up
	// signal to find out how it was actually positioned.
	PopupAtRect(rectWindow gdk.Window, rect *gdk.Rectangle, rectAnchor gdk.Gravity, menuAnchor gdk.Gravity, triggerEvent *gdk.Event)
	// PopupAtWidget displays @menu and makes it available for selection.
	//
	// See gtk_menu_popup_at_pointer () to pop up a menu at the master pointer.
	// gtk_menu_popup_at_rect () also allows you to position a menu at an
	// arbitrary rectangle.
	//
	// ! (popup-anchors.png)
	//
	// @menu will be positioned at @widget, aligning their anchor points.
	// @widget_anchor and @menu_anchor determine anchor points on @widget and
	// @menu to pin together. @menu can optionally be offset by
	// Menu:rect-anchor-dx and Menu:rect-anchor-dy.
	//
	// Anchors should be specified under the assumption that the text direction
	// is left-to-right; they will be flipped horizontally automatically if the
	// text direction is right-to-left.
	//
	// Other properties that influence the behaviour of this function are
	// Menu:anchor-hints and Menu:menu-type-hint. Connect to the Menu::popped-up
	// signal to find out how it was actually positioned.
	PopupAtWidget(widget Widget, widgetAnchor gdk.Gravity, menuAnchor gdk.Gravity, triggerEvent *gdk.Event)
	// ReorderChild moves @child to a new @position in the list of @menu
	// children.
	ReorderChild(child Widget, position int)
	// Reposition repositions the menu according to its position function.
	Reposition()
	// SetAccelGroup: set the AccelGroup which holds global accelerators for the
	// menu. This accelerator group needs to also be added to all windows that
	// this menu is being used in with gtk_window_add_accel_group(), in order
	// for those windows to support all the accelerators contained in this
	// group.
	SetAccelGroup(accelGroup AccelGroup)
	// SetAccelPath sets an accelerator path for this menu from which
	// accelerator paths for its immediate children, its menu items, can be
	// constructed. The main purpose of this function is to spare the programmer
	// the inconvenience of having to call gtk_menu_item_set_accel_path() on
	// each menu item that should support runtime user changable accelerators.
	// Instead, by just calling gtk_menu_set_accel_path() on their parent, each
	// menu item of this menu, that contains a label describing its purpose,
	// automatically gets an accel path assigned.
	//
	// For example, a menu containing menu items “New” and “Exit”, will, after
	// `gtk_menu_set_accel_path (menu, "<Gnumeric-Sheet>/File");` has been
	// called, assign its items the accel paths: `"<Gnumeric-Sheet>/File/New"`
	// and `"<Gnumeric-Sheet>/File/Exit"`.
	//
	// Assigning accel paths to menu items then enables the user to change their
	// accelerators at runtime. More details about accelerator paths and their
	// default setups can be found at gtk_accel_map_add_entry().
	//
	// Note that @accel_path string will be stored in a #GQuark. Therefore, if
	// you pass a static string, you can save some memory by interning it first
	// with g_intern_static_string().
	SetAccelPath(accelPath string)
	// SetActive selects the specified menu item within the menu. This is used
	// by the ComboBox and should not be used by anyone else.
	SetActive(index uint)
	// SetMonitor informs GTK+ on which monitor a menu should be popped up. See
	// gdk_monitor_get_geometry().
	//
	// This function should be called from a MenuPositionFunc if the menu should
	// not appear on the same monitor as the pointer. This information can’t be
	// reliably inferred from the coordinates returned by a MenuPositionFunc,
	// since, for very long menus, these coordinates may extend beyond the
	// monitor boundaries or even the screen boundaries.
	SetMonitor(monitorNum int)
	// SetReserveToggleSize sets whether the menu should reserve space for
	// drawing toggles or icons, regardless of their actual presence.
	SetReserveToggleSize(reserveToggleSize bool)
	// SetScreen sets the Screen on which the menu will be displayed.
	SetScreen(screen gdk.Screen)
	// SetTearoffState changes the tearoff state of the menu. A menu is normally
	// displayed as drop down menu which persists as long as the menu is active.
	// It can also be displayed as a tearoff menu which persists until it is
	// closed or reattached.
	SetTearoffState(tornOff bool)
	// SetTitle sets the title string for the menu.
	//
	// The title is displayed when the menu is shown as a tearoff menu. If
	// @title is nil, the menu will see if it is attached to a parent menu item,
	// and if so it will try to use the same text as that menu item’s label.
	SetTitle(title string)
}

// menu implements the Menu interface.
type menu struct {
	MenuShell
	Buildable
}

var _ Menu = (*menu)(nil)

// WrapMenu wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenu(obj *externglib.Object) Menu {
	return Menu{
		MenuShell: WrapMenuShell(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalMenu(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenu(obj), nil
}

// NewMenu constructs a class Menu.
func NewMenu() Menu {
	var _cret C.GtkMenu

	cret = C.gtk_menu_new()

	var _menu Menu

	_menu = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Menu)

	return _menu
}

// NewMenuFromModel constructs a class Menu.
func NewMenuFromModel(model gio.MenuModel) Menu {
	var _arg1 *C.GMenuModel

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkMenu

	cret = C.gtk_menu_new_from_model(_arg1)

	var _menu Menu

	_menu = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Menu)

	return _menu
}

// Attach adds a new MenuItem to a (table) menu. The number of “cells” that
// an item will occupy is specified by @left_attach, @right_attach,
// @top_attach and @bottom_attach. These each represent the leftmost,
// rightmost, uppermost and lower column and row numbers of the table.
// (Columns and rows are indexed from zero).
//
// Note that this function is not related to gtk_menu_detach().
func (m menu) Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GtkWidget
	var _arg2 C.guint
	var _arg3 C.guint
	var _arg4 C.guint
	var _arg5 C.guint

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_menu_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// Detach detaches the menu from the widget to which it had been attached.
// This function will call the callback function, @detacher, provided when
// the gtk_menu_attach_to_widget() function was called.
func (m menu) Detach() {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	C.gtk_menu_detach(_arg0)
}

// AccelGroup gets the AccelGroup which holds global accelerators for the
// menu. See gtk_menu_set_accel_group().
func (m menu) AccelGroup() AccelGroup {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret *C.GtkAccelGroup

	cret = C.gtk_menu_get_accel_group(_arg0)

	var _accelGroup AccelGroup

	_accelGroup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(AccelGroup)

	return _accelGroup
}

// AccelPath retrieves the accelerator path set on the menu.
func (m menu) AccelPath() string {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret *C.gchar

	cret = C.gtk_menu_get_accel_path(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Active returns the selected menu item from the menu. This is used by the
// ComboBox.
func (m menu) Active() Widget {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_menu_get_active(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// AttachWidget returns the Widget that the menu is attached to.
func (m menu) AttachWidget() Widget {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_menu_get_attach_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Monitor retrieves the number of the monitor on which to show the menu.
func (m menu) Monitor() int {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret C.gint

	cret = C.gtk_menu_get_monitor(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// ReserveToggleSize returns whether the menu reserves space for toggles and
// icons, regardless of their actual presence.
func (m menu) ReserveToggleSize() bool {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean

	cret = C.gtk_menu_get_reserve_toggle_size(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// TearoffState returns whether the menu is torn off. See
// gtk_menu_set_tearoff_state().
func (m menu) TearoffState() bool {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean

	cret = C.gtk_menu_get_tearoff_state(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Title returns the title of the menu. See gtk_menu_set_title().
func (m menu) Title() string {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	var _cret *C.gchar

	cret = C.gtk_menu_get_title(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PlaceOnMonitor places @menu on the given monitor.
func (m menu) PlaceOnMonitor(monitor gdk.Monitor) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GdkMonitor

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gtk_menu_place_on_monitor(_arg0, _arg1)
}

// Popdown removes the menu from the screen.
func (m menu) Popdown() {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	C.gtk_menu_popdown(_arg0)
}

// PopupAtPointer displays @menu and makes it available for selection.
//
// See gtk_menu_popup_at_widget () to pop up a menu at a widget.
// gtk_menu_popup_at_rect () also allows you to position a menu at an
// arbitrary rectangle.
//
// @menu will be positioned at the pointer associated with @trigger_event.
//
// Properties that influence the behaviour of this function are
// Menu:anchor-hints, Menu:rect-anchor-dx, Menu:rect-anchor-dy, and
// Menu:menu-type-hint. Connect to the Menu::popped-up signal to find out
// how it was actually positioned.
func (m menu) PopupAtPointer(triggerEvent *gdk.Event) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GdkEvent

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	var _arg1 *C.GdkEvent // unsupported

	C.gtk_menu_popup_at_pointer(_arg0, _arg1)
}

// PopupAtRect displays @menu and makes it available for selection.
//
// See gtk_menu_popup_at_widget () and gtk_menu_popup_at_pointer (), which
// handle more common cases for popping up menus.
//
// @menu will be positioned at @rect, aligning their anchor points. @rect is
// relative to the top-left corner of @rect_window. @rect_anchor and
// @menu_anchor determine anchor points on @rect and @menu to pin together.
// @menu can optionally be offset by Menu:rect-anchor-dx and
// Menu:rect-anchor-dy.
//
// Anchors should be specified under the assumption that the text direction
// is left-to-right; they will be flipped horizontally automatically if the
// text direction is right-to-left.
//
// Other properties that influence the behaviour of this function are
// Menu:anchor-hints and Menu:menu-type-hint. Connect to the Menu::popped-up
// signal to find out how it was actually positioned.
func (m menu) PopupAtRect(rectWindow gdk.Window, rect *gdk.Rectangle, rectAnchor gdk.Gravity, menuAnchor gdk.Gravity, triggerEvent *gdk.Event) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GdkWindow
	var _arg2 *C.GdkRectangle
	var _arg3 C.GdkGravity
	var _arg4 C.GdkGravity
	var _arg5 *C.GdkEvent

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(rectWindow.Native()))
	_arg2 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))
	_arg3 = (C.GdkGravity)(rectAnchor)
	_arg4 = (C.GdkGravity)(menuAnchor)
	var _arg5 *C.GdkEvent // unsupported

	C.gtk_menu_popup_at_rect(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// PopupAtWidget displays @menu and makes it available for selection.
//
// See gtk_menu_popup_at_pointer () to pop up a menu at the master pointer.
// gtk_menu_popup_at_rect () also allows you to position a menu at an
// arbitrary rectangle.
//
// ! (popup-anchors.png)
//
// @menu will be positioned at @widget, aligning their anchor points.
// @widget_anchor and @menu_anchor determine anchor points on @widget and
// @menu to pin together. @menu can optionally be offset by
// Menu:rect-anchor-dx and Menu:rect-anchor-dy.
//
// Anchors should be specified under the assumption that the text direction
// is left-to-right; they will be flipped horizontally automatically if the
// text direction is right-to-left.
//
// Other properties that influence the behaviour of this function are
// Menu:anchor-hints and Menu:menu-type-hint. Connect to the Menu::popped-up
// signal to find out how it was actually positioned.
func (m menu) PopupAtWidget(widget Widget, widgetAnchor gdk.Gravity, menuAnchor gdk.Gravity, triggerEvent *gdk.Event) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GtkWidget
	var _arg2 C.GdkGravity
	var _arg3 C.GdkGravity
	var _arg4 *C.GdkEvent

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (C.GdkGravity)(widgetAnchor)
	_arg3 = (C.GdkGravity)(menuAnchor)
	var _arg4 *C.GdkEvent // unsupported

	C.gtk_menu_popup_at_widget(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// ReorderChild moves @child to a new @position in the list of @menu
// children.
func (m menu) ReorderChild(child Widget, position int) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GtkWidget
	var _arg2 C.gint

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_menu_reorder_child(_arg0, _arg1, _arg2)
}

// Reposition repositions the menu according to its position function.
func (m menu) Reposition() {
	var _arg0 *C.GtkMenu

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))

	C.gtk_menu_reposition(_arg0)
}

// SetAccelGroup: set the AccelGroup which holds global accelerators for the
// menu. This accelerator group needs to also be added to all windows that
// this menu is being used in with gtk_window_add_accel_group(), in order
// for those windows to support all the accelerators contained in this
// group.
func (m menu) SetAccelGroup(accelGroup AccelGroup) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GtkAccelGroup

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_menu_set_accel_group(_arg0, _arg1)
}

// SetAccelPath sets an accelerator path for this menu from which
// accelerator paths for its immediate children, its menu items, can be
// constructed. The main purpose of this function is to spare the programmer
// the inconvenience of having to call gtk_menu_item_set_accel_path() on
// each menu item that should support runtime user changable accelerators.
// Instead, by just calling gtk_menu_set_accel_path() on their parent, each
// menu item of this menu, that contains a label describing its purpose,
// automatically gets an accel path assigned.
//
// For example, a menu containing menu items “New” and “Exit”, will, after
// `gtk_menu_set_accel_path (menu, "<Gnumeric-Sheet>/File");` has been
// called, assign its items the accel paths: `"<Gnumeric-Sheet>/File/New"`
// and `"<Gnumeric-Sheet>/File/Exit"`.
//
// Assigning accel paths to menu items then enables the user to change their
// accelerators at runtime. More details about accelerator paths and their
// default setups can be found at gtk_accel_map_add_entry().
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if
// you pass a static string, you can save some memory by interning it first
// with g_intern_static_string().
func (m menu) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.gchar

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_set_accel_path(_arg0, _arg1)
}

// SetActive selects the specified menu item within the menu. This is used
// by the ComboBox and should not be used by anyone else.
func (m menu) SetActive(index uint) {
	var _arg0 *C.GtkMenu
	var _arg1 C.guint

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = C.guint(index)

	C.gtk_menu_set_active(_arg0, _arg1)
}

// SetMonitor informs GTK+ on which monitor a menu should be popped up. See
// gdk_monitor_get_geometry().
//
// This function should be called from a MenuPositionFunc if the menu should
// not appear on the same monitor as the pointer. This information can’t be
// reliably inferred from the coordinates returned by a MenuPositionFunc,
// since, for very long menus, these coordinates may extend beyond the
// monitor boundaries or even the screen boundaries.
func (m menu) SetMonitor(monitorNum int) {
	var _arg0 *C.GtkMenu
	var _arg1 C.gint

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(monitorNum)

	C.gtk_menu_set_monitor(_arg0, _arg1)
}

// SetReserveToggleSize sets whether the menu should reserve space for
// drawing toggles or icons, regardless of their actual presence.
func (m menu) SetReserveToggleSize(reserveToggleSize bool) {
	var _arg0 *C.GtkMenu
	var _arg1 C.gboolean

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	if reserveToggleSize {
		_arg1 = C.gboolean(1)
	}

	C.gtk_menu_set_reserve_toggle_size(_arg0, _arg1)
}

// SetScreen sets the Screen on which the menu will be displayed.
func (m menu) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.GdkScreen

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_menu_set_screen(_arg0, _arg1)
}

// SetTearoffState changes the tearoff state of the menu. A menu is normally
// displayed as drop down menu which persists as long as the menu is active.
// It can also be displayed as a tearoff menu which persists until it is
// closed or reattached.
func (m menu) SetTearoffState(tornOff bool) {
	var _arg0 *C.GtkMenu
	var _arg1 C.gboolean

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	if tornOff {
		_arg1 = C.gboolean(1)
	}

	C.gtk_menu_set_tearoff_state(_arg0, _arg1)
}

// SetTitle sets the title string for the menu.
//
// The title is displayed when the menu is shown as a tearoff menu. If
// @title is nil, the menu will see if it is attached to a parent menu item,
// and if so it will try to use the same text as that menu item’s label.
func (m menu) SetTitle(title string) {
	var _arg0 *C.GtkMenu
	var _arg1 *C.gchar

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_set_title(_arg0, _arg1)
}