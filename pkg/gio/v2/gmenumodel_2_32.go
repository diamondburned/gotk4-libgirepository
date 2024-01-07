// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_MenuModel_ConnectItemsChanged(gpointer, gint, gint, gint, guintptr);
import "C"

// GType values.
var (
	GTypeMenuAttributeIter = coreglib.Type(girepository.MustFind("Gio", "MenuAttributeIter").RegisteredGType())
	GTypeMenuLinkIter      = coreglib.Type(girepository.MustFind("Gio", "MenuLinkIter").RegisteredGType())
	GTypeMenuModel         = coreglib.Type(girepository.MustFind("Gio", "MenuModel").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuAttributeIter, F: marshalMenuAttributeIter},
		coreglib.TypeMarshaler{T: GTypeMenuLinkIter, F: marshalMenuLinkIter},
		coreglib.TypeMarshaler{T: GTypeMenuModel, F: marshalMenuModel},
	})
}

// MENU_ATTRIBUTE_ACTION: menu item attribute which holds the action name of the
// item. Action names are namespaced with an identifier for the action group in
// which the action resides. For example, "win." for window-specific actions and
// "app." for application-wide actions.
//
// See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().
const MENU_ATTRIBUTE_ACTION = "action"

// MENU_ATTRIBUTE_LABEL: menu item attribute which holds the label of the item.
const MENU_ATTRIBUTE_LABEL = "label"

// MENU_ATTRIBUTE_TARGET: menu item attribute which holds the target with which
// the item's action will be activated.
//
// See also g_menu_item_set_action_and_target().
const MENU_ATTRIBUTE_TARGET = "target"

// MENU_LINK_SECTION: name of the link that associates a menu item with a
// section. The linked menu will usually be shown in place of the menu item,
// using the item's label as a header.
//
// See also g_menu_item_set_link().
const MENU_LINK_SECTION = "section"

// MENU_LINK_SUBMENU: name of the link that associates a menu item with a
// submenu.
//
// See also g_menu_item_set_link().
const MENU_LINK_SUBMENU = "submenu"

// MenuAttributeIterOverrides contains methods that are overridable.
type MenuAttributeIterOverrides struct {
}

func defaultMenuAttributeIterOverrides(v *MenuAttributeIter) MenuAttributeIterOverrides {
	return MenuAttributeIterOverrides{}
}

// MenuAttributeIter is an opaque structure type. You must access it using the
// functions below.
type MenuAttributeIter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*MenuAttributeIter)(nil)
)

// MenuAttributeIterer describes types inherited from class MenuAttributeIter.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MenuAttributeIterer interface {
	coreglib.Objector
	baseMenuAttributeIter() *MenuAttributeIter
}

var _ MenuAttributeIterer = (*MenuAttributeIter)(nil)

func init() {
	coreglib.RegisterClassInfo[*MenuAttributeIter, *MenuAttributeIterClass, MenuAttributeIterOverrides](
		GTypeMenuAttributeIter,
		initMenuAttributeIterClass,
		wrapMenuAttributeIter,
		defaultMenuAttributeIterOverrides,
	)
}

func initMenuAttributeIterClass(gclass unsafe.Pointer, overrides MenuAttributeIterOverrides, classInitFunc func(*MenuAttributeIterClass)) {
	if classInitFunc != nil {
		class := (*MenuAttributeIterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMenuAttributeIter(obj *coreglib.Object) *MenuAttributeIter {
	return &MenuAttributeIter{
		Object: obj,
	}
}

func marshalMenuAttributeIter(p uintptr) (interface{}, error) {
	return wrapMenuAttributeIter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *MenuAttributeIter) baseMenuAttributeIter() *MenuAttributeIter {
	return v
}

// BaseMenuAttributeIter returns the underlying base object.
func BaseMenuAttributeIter(obj MenuAttributeIterer) *MenuAttributeIter {
	return obj.baseMenuAttributeIter()
}

// MenuLinkIterOverrides contains methods that are overridable.
type MenuLinkIterOverrides struct {
}

func defaultMenuLinkIterOverrides(v *MenuLinkIter) MenuLinkIterOverrides {
	return MenuLinkIterOverrides{}
}

// MenuLinkIter is an opaque structure type. You must access it using the
// functions below.
type MenuLinkIter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*MenuLinkIter)(nil)
)

// MenuLinkIterer describes types inherited from class MenuLinkIter.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MenuLinkIterer interface {
	coreglib.Objector
	baseMenuLinkIter() *MenuLinkIter
}

var _ MenuLinkIterer = (*MenuLinkIter)(nil)

func init() {
	coreglib.RegisterClassInfo[*MenuLinkIter, *MenuLinkIterClass, MenuLinkIterOverrides](
		GTypeMenuLinkIter,
		initMenuLinkIterClass,
		wrapMenuLinkIter,
		defaultMenuLinkIterOverrides,
	)
}

func initMenuLinkIterClass(gclass unsafe.Pointer, overrides MenuLinkIterOverrides, classInitFunc func(*MenuLinkIterClass)) {
	if classInitFunc != nil {
		class := (*MenuLinkIterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMenuLinkIter(obj *coreglib.Object) *MenuLinkIter {
	return &MenuLinkIter{
		Object: obj,
	}
}

func marshalMenuLinkIter(p uintptr) (interface{}, error) {
	return wrapMenuLinkIter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *MenuLinkIter) baseMenuLinkIter() *MenuLinkIter {
	return v
}

// BaseMenuLinkIter returns the underlying base object.
func BaseMenuLinkIter(obj MenuLinkIterer) *MenuLinkIter {
	return obj.baseMenuLinkIter()
}

// MenuModelOverrides contains methods that are overridable.
type MenuModelOverrides struct {
}

func defaultMenuModelOverrides(v *MenuModel) MenuModelOverrides {
	return MenuModelOverrides{}
}

// MenuModel represents the contents of a menu -- an ordered list of menu items.
// The items are associated with actions, which can be activated through them.
// Items can be grouped in sections, and may have submenus associated with them.
// Both items and sections usually have some representation data, such as labels
// or icons. The type of the associated action (ie whether it is stateful, and
// what kind of state it has) can influence the representation of the item.
//
// The conceptual model of menus in Model is hierarchical: sections and submenus
// are again represented by Models. Menus themselves do not define their own
// roles. Rather, the role of a particular Model is defined by the item that
// references it (or, in the case of the 'root' menu, is defined by the context
// in which it is used).
//
// As an example, consider the visible portions of this menu:
//
//
// An example menu
//
// ! (menu-example.png)
//
// There are 8 "menus" visible in the screenshot: one menubar, two submenus and
// 5 sections:
//
// - the toplevel menubar (containing 4 items)
//
// - the View submenu (containing 3 sections)
//
// - the first section of the View submenu (containing 2 items)
//
// - the second section of the View submenu (containing 1 item)
//
// - the final section of the View submenu (containing 1 item)
//
// - the Highlight Mode submenu (containing 2 sections)
//
// - the Sources section (containing 2 items)
//
// - the Markup section (containing 2 items)
//
// The [example][menu-model] illustrates the conceptual connection between these
// 8 menus. Each large block in the figure represents a menu and the smaller
// blocks within the large block represent items in that menu. Some items
// contain references to other menus.
//
//
// A menu example
//
// ! (menu-model.png)
//
// Notice that the separators visible in the [example][menu-example] appear
// nowhere in the [menu model][menu-model]. This is because separators are not
// explicitly represented in the menu model. Instead, a separator is inserted
// between any two non-empty sections of a menu. Section items can have labels
// just like any other item. In that case, a display system may show a section
// header instead of a separator.
//
// The motivation for this abstract model of application controls is that modern
// user interfaces tend to make these controls available outside the
// application. Examples include global menus, jumplists, dash boards, etc. To
// support such uses, it is necessary to 'export' information about actions and
// their representation in menus, which is exactly what the [GActionGroup
// exporter][gio-GActionGroup-exporter] and the [GMenuModel
// exporter][gio-GMenuModel-exporter] do for Group and Model. The client-side
// counterparts to make use of the exported information are BusActionGroup and
// BusMenuModel.
//
// The API of Model is very generic, with iterators for the attributes and links
// of an item, see g_menu_model_iterate_item_attributes() and
// g_menu_model_iterate_item_links(). The 'standard' attributes and link types
// have predefined names: G_MENU_ATTRIBUTE_LABEL, G_MENU_ATTRIBUTE_ACTION,
// G_MENU_ATTRIBUTE_TARGET, G_MENU_LINK_SECTION and G_MENU_LINK_SUBMENU.
//
// Items in a Model represent active controls if they refer to an action that
// can get activated when the user interacts with the menu item. The reference
// to the action is encoded by the string id in the G_MENU_ATTRIBUTE_ACTION
// attribute. An action id uniquely identifies an action in an action group.
// Which action group(s) provide actions depends on the context in which the
// menu model is used. E.g. when the model is exported as the application menu
// of a Application, actions can be application-wide or window-specific (and
// thus come from two different action groups). By convention, the
// application-wide actions have names that start with "app.", while the names
// of window-specific actions start with "win.".
//
// While a wide variety of stateful actions is possible, the following is the
// minimum that is expected to be supported by all users of exported menu
// information:
//
// - an action with no parameter type and no state
//
// - an action with no parameter type and boolean state
//
// - an action with string parameter type and string state
//
//
// Stateless
//
// A stateless action typically corresponds to an ordinary menu item.
//
// Selecting such a menu item will activate the action (with no parameter).
//
//
// Boolean State
//
// An action with a boolean state will most typically be used with a "toggle" or
// "switch" menu item. The state can be set directly, but activating the action
// (with no parameter) results in the state being toggled.
//
// Selecting a toggle menu item will activate the action. The menu item should
// be rendered as "checked" when the state is true.
//
//
// String Parameter and State
//
// Actions with string parameters and state will most typically be used to
// represent an enumerated choice over the items available for a group of radio
// menu items. Activating the action with a string parameter is equivalent to
// setting that parameter as the state.
//
// Radio menu items, in addition to being associated with the action, will have
// a target value. Selecting that menu item will result in activation of the
// action with the target value as the parameter. The menu item should be
// rendered as "selected" when the state of the action is equal to the target
// value of the menu item.
type MenuModel struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*MenuModel)(nil)
)

// MenuModeller describes types inherited from class MenuModel.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MenuModeller interface {
	coreglib.Objector
	baseMenuModel() *MenuModel
}

var _ MenuModeller = (*MenuModel)(nil)

func init() {
	coreglib.RegisterClassInfo[*MenuModel, *MenuModelClass, MenuModelOverrides](
		GTypeMenuModel,
		initMenuModelClass,
		wrapMenuModel,
		defaultMenuModelOverrides,
	)
}

func initMenuModelClass(gclass unsafe.Pointer, overrides MenuModelOverrides, classInitFunc func(*MenuModelClass)) {
	if classInitFunc != nil {
		class := (*MenuModelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMenuModel(obj *coreglib.Object) *MenuModel {
	return &MenuModel{
		Object: obj,
	}
}

func marshalMenuModel(p uintptr) (interface{}, error) {
	return wrapMenuModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *MenuModel) baseMenuModel() *MenuModel {
	return v
}

// BaseMenuModel returns the underlying base object.
func BaseMenuModel(obj MenuModeller) *MenuModel {
	return obj.baseMenuModel()
}

// ConnectItemsChanged is emitted when a change has occurred to the menu.
//
// The only changes that can occur to a menu is that items are removed or added.
// Items may not change (except by being removed and added back in the same
// location). This signal is capable of describing both of those changes (at the
// same time).
//
// The signal means that starting at the index position, removed items were
// removed and added items were added in their place. If removed is zero then
// only items were added. If added is zero then only items were removed.
//
// As an example, if the menu contains items a, b, c, d (in that order) and the
// signal (2, 1, 3) occurs then the new composition of the menu will be a, b, _,
// _, _, d (with each _ representing some new item).
//
// Signal handlers may query the model (particularly the added items) and expect
// to see the results of the modification that is being reported. The signal is
// emitted after the modification.
func (v *MenuModel) ConnectItemsChanged(f func(position, removed, added int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "items-changed", false, unsafe.Pointer(C._gotk4_gio2_MenuModel_ConnectItemsChanged), f)
}
