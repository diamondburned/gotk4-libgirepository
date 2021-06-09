// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_menu_attribute_iter_get_type()), F: marshalMenuAttributeIter},
		{T: externglib.Type(C.g_menu_link_iter_get_type()), F: marshalMenuLinkIter},
		{T: externglib.Type(C.g_menu_model_get_type()), F: marshalMenuModel},
	})
}

// MenuAttributeIter is an opaque structure type. You must access it using the
// functions below.
type MenuAttributeIter interface {
	gextras.Objector

	// Name gets the name of the attribute at the current iterator position, as
	// a string.
	//
	// The iterator is not advanced.
	Name() string
	// GetNext: this function combines g_menu_attribute_iter_next() with
	// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
	//
	// First the iterator is advanced to the next (possibly first) attribute. If
	// that fails, then false is returned and there are no other effects.
	//
	// If successful, @name and @value are set to the name and value of the
	// attribute that has just been advanced to. At this point,
	// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value()
	// will return the same values again.
	//
	// The value returned in @name remains valid for as long as the iterator
	// remains at the current position. The value returned in @value must be
	// unreffed using g_variant_unref() when it is no longer in use.
	GetNext() (string, *glib.Variant, bool)
	// Value gets the value of the attribute at the current iterator position.
	//
	// The iterator is not advanced.
	Value() *glib.Variant
	// Next attempts to advance the iterator to the next (possibly first)
	// attribute.
	//
	// true is returned on success, or false if there are no more attributes.
	//
	// You must call this function when you first acquire the iterator to
	// advance it to the first attribute (and determine if the first attribute
	// exists at all).
	Next() bool
}

// menuAttributeIter implements the MenuAttributeIter interface.
type menuAttributeIter struct {
	gextras.Objector
}

var _ MenuAttributeIter = (*menuAttributeIter)(nil)

// WrapMenuAttributeIter wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuAttributeIter(obj *externglib.Object) MenuAttributeIter {
	return MenuAttributeIter{
		Objector: obj,
	}
}

func marshalMenuAttributeIter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuAttributeIter(obj), nil
}

// Name gets the name of the attribute at the current iterator position, as
// a string.
//
// The iterator is not advanced.
func (i menuAttributeIter) Name() string {
	var _arg0 *C.GMenuAttributeIter

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(i.Native()))

	var _cret *C.gchar

	cret = C.g_menu_attribute_iter_get_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// GetNext: this function combines g_menu_attribute_iter_next() with
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) attribute. If
// that fails, then false is returned and there are no other effects.
//
// If successful, @name and @value are set to the name and value of the
// attribute that has just been advanced to. At this point,
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value()
// will return the same values again.
//
// The value returned in @name remains valid for as long as the iterator
// remains at the current position. The value returned in @value must be
// unreffed using g_variant_unref() when it is no longer in use.
func (i menuAttributeIter) GetNext() (string, *glib.Variant, bool) {
	var _arg0 *C.GMenuAttributeIter

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(i.Native()))

	var _arg1 **C.gchar
	var _value *glib.Variant
	var _cret C.gboolean

	cret = C.g_menu_attribute_iter_get_next(_arg0, _arg1, (**C.GVariant)(unsafe.Pointer(&_value)))

	var _outName string

	var _ok bool

	_outName = C.GoString(_arg1)

	if _cret {
		_ok = true
	}

	return _outName, _value, _ok
}

// Value gets the value of the attribute at the current iterator position.
//
// The iterator is not advanced.
func (i menuAttributeIter) Value() *glib.Variant {
	var _arg0 *C.GMenuAttributeIter

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(i.Native()))

	var _cret *C.GVariant

	cret = C.g_menu_attribute_iter_get_value(_arg0)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _variant
}

// Next attempts to advance the iterator to the next (possibly first)
// attribute.
//
// true is returned on success, or false if there are no more attributes.
//
// You must call this function when you first acquire the iterator to
// advance it to the first attribute (and determine if the first attribute
// exists at all).
func (i menuAttributeIter) Next() bool {
	var _arg0 *C.GMenuAttributeIter

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_menu_attribute_iter_next(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// MenuLinkIter is an opaque structure type. You must access it using the
// functions below.
type MenuLinkIter interface {
	gextras.Objector

	// Name gets the name of the link at the current iterator position.
	//
	// The iterator is not advanced.
	Name() string
	// GetNext: this function combines g_menu_link_iter_next() with
	// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
	//
	// First the iterator is advanced to the next (possibly first) link. If that
	// fails, then false is returned and there are no other effects.
	//
	// If successful, @out_link and @value are set to the name and Model of the
	// link that has just been advanced to. At this point,
	// g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return
	// the same values again.
	//
	// The value returned in @out_link remains valid for as long as the iterator
	// remains at the current position. The value returned in @value must be
	// unreffed using g_object_unref() when it is no longer in use.
	GetNext() (string, MenuModel, bool)
	// Value gets the linked Model at the current iterator position.
	//
	// The iterator is not advanced.
	Value() MenuModel
	// Next attempts to advance the iterator to the next (possibly first) link.
	//
	// true is returned on success, or false if there are no more links.
	//
	// You must call this function when you first acquire the iterator to
	// advance it to the first link (and determine if the first link exists at
	// all).
	Next() bool
}

// menuLinkIter implements the MenuLinkIter interface.
type menuLinkIter struct {
	gextras.Objector
}

var _ MenuLinkIter = (*menuLinkIter)(nil)

// WrapMenuLinkIter wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuLinkIter(obj *externglib.Object) MenuLinkIter {
	return MenuLinkIter{
		Objector: obj,
	}
}

func marshalMenuLinkIter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuLinkIter(obj), nil
}

// Name gets the name of the link at the current iterator position.
//
// The iterator is not advanced.
func (i menuLinkIter) Name() string {
	var _arg0 *C.GMenuLinkIter

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(i.Native()))

	var _cret *C.gchar

	cret = C.g_menu_link_iter_get_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// GetNext: this function combines g_menu_link_iter_next() with
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) link. If that
// fails, then false is returned and there are no other effects.
//
// If successful, @out_link and @value are set to the name and Model of the
// link that has just been advanced to. At this point,
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return
// the same values again.
//
// The value returned in @out_link remains valid for as long as the iterator
// remains at the current position. The value returned in @value must be
// unreffed using g_object_unref() when it is no longer in use.
func (i menuLinkIter) GetNext() (string, MenuModel, bool) {
	var _arg0 *C.GMenuLinkIter

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(i.Native()))

	var _arg1 **C.gchar
	var _arg2 *C.GMenuModel
	var _cret C.gboolean

	cret = C.g_menu_link_iter_get_next(_arg0, _arg1, &_arg2)

	var _outLink string
	var _value MenuModel
	var _ok bool

	_outLink = C.GoString(_arg1)
	_value = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_arg2.Native()))).(MenuModel)
	if _cret {
		_ok = true
	}

	return _outLink, _value, _ok
}

// Value gets the linked Model at the current iterator position.
//
// The iterator is not advanced.
func (i menuLinkIter) Value() MenuModel {
	var _arg0 *C.GMenuLinkIter

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(i.Native()))

	var _cret *C.GMenuModel

	cret = C.g_menu_link_iter_get_value(_arg0)

	var _menuModel MenuModel

	_menuModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MenuModel)

	return _menuModel
}

// Next attempts to advance the iterator to the next (possibly first) link.
//
// true is returned on success, or false if there are no more links.
//
// You must call this function when you first acquire the iterator to
// advance it to the first link (and determine if the first link exists at
// all).
func (i menuLinkIter) Next() bool {
	var _arg0 *C.GMenuLinkIter

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_menu_link_iter_next(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
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
// - the toplevel menubar (containing 4 items) - the View submenu (containing 3
// sections) - the first section of the View submenu (containing 2 items) - the
// second section of the View submenu (containing 1 item) - the final section of
// the View submenu (containing 1 item) - the Highlight Mode submenu (containing
// 2 sections) - the Sources section (containing 2 items) - the Markup section
// (containing 2 items)
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
// information: - an action with no parameter type and no state - an action with
// no parameter type and boolean state - an action with string parameter type
// and string state
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
type MenuModel interface {
	gextras.Objector

	// ItemAttributeValue queries the item at position @item_index in @model for
	// the attribute specified by @attribute.
	//
	// If @expected_type is non-nil then it specifies the expected type of the
	// attribute. If it is nil then any type will be accepted.
	//
	// If the attribute exists and matches @expected_type (or if the expected
	// type is unspecified) then the value is returned.
	//
	// If the attribute does not exist, or does not match the expected type then
	// nil is returned.
	ItemAttributeValue(itemIndex int, attribute string, expectedType *glib.VariantType) *glib.Variant
	// ItemLink queries the item at position @item_index in @model for the link
	// specified by @link.
	//
	// If the link exists, the linked Model is returned. If the link does not
	// exist, nil is returned.
	ItemLink(itemIndex int, link string) MenuModel
	// NItems: query the number of items in @model.
	NItems() int
	// IsMutable queries if @model is mutable.
	//
	// An immutable Model will never emit the Model::items-changed signal.
	// Consumers of the model may make optimisations accordingly.
	IsMutable() bool
	// ItemsChanged requests emission of the Model::items-changed signal on
	// @model.
	//
	// This function should never be called except by Model subclasses. Any
	// other calls to this function will very likely lead to a violation of the
	// interface of the model.
	//
	// The implementation should update its internal representation of the menu
	// before emitting the signal. The implementation should further expect to
	// receive queries about the new state of the menu (and particularly added
	// menu items) while signal handlers are running.
	//
	// The implementation must dispatch this call directly from a mainloop entry
	// and not in response to calls -- particularly those from the Model API.
	// Said another way: the menu must not change while user code is running
	// without returning to the mainloop.
	ItemsChanged(position int, removed int, added int)
	// IterateItemAttributes creates a AttributeIter to iterate over the
	// attributes of the item at position @item_index in @model.
	//
	// You must free the iterator with g_object_unref() when you are done.
	IterateItemAttributes(itemIndex int) MenuAttributeIter
	// IterateItemLinks creates a LinkIter to iterate over the links of the item
	// at position @item_index in @model.
	//
	// You must free the iterator with g_object_unref() when you are done.
	IterateItemLinks(itemIndex int) MenuLinkIter
}

// menuModel implements the MenuModel interface.
type menuModel struct {
	gextras.Objector
}

var _ MenuModel = (*menuModel)(nil)

// WrapMenuModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuModel(obj *externglib.Object) MenuModel {
	return MenuModel{
		Objector: obj,
	}
}

func marshalMenuModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuModel(obj), nil
}

// ItemAttributeValue queries the item at position @item_index in @model for
// the attribute specified by @attribute.
//
// If @expected_type is non-nil then it specifies the expected type of the
// attribute. If it is nil then any type will be accepted.
//
// If the attribute exists and matches @expected_type (or if the expected
// type is unspecified) then the value is returned.
//
// If the attribute does not exist, or does not match the expected type then
// nil is returned.
func (m menuModel) ItemAttributeValue(itemIndex int, attribute string, expectedType *glib.VariantType) *glib.Variant {
	var _arg0 *C.GMenuModel
	var _arg1 C.gint
	var _arg2 *C.gchar
	var _arg3 *C.GVariantType

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(itemIndex)
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GVariantType)(unsafe.Pointer(expectedType.Native()))

	var _cret *C.GVariant

	cret = C.g_menu_model_get_item_attribute_value(_arg0, _arg1, _arg2, _arg3)

	var _variant *glib.Variant

	_variant = glib.WrapVariant(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _variant
}

// ItemLink queries the item at position @item_index in @model for the link
// specified by @link.
//
// If the link exists, the linked Model is returned. If the link does not
// exist, nil is returned.
func (m menuModel) ItemLink(itemIndex int, link string) MenuModel {
	var _arg0 *C.GMenuModel
	var _arg1 C.gint
	var _arg2 *C.gchar

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(itemIndex)
	_arg2 = (*C.gchar)(C.CString(link))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret *C.GMenuModel

	cret = C.g_menu_model_get_item_link(_arg0, _arg1, _arg2)

	var _menuModel MenuModel

	_menuModel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MenuModel)

	return _menuModel
}

// NItems: query the number of items in @model.
func (m menuModel) NItems() int {
	var _arg0 *C.GMenuModel

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))

	var _cret C.gint

	cret = C.g_menu_model_get_n_items(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// IsMutable queries if @model is mutable.
//
// An immutable Model will never emit the Model::items-changed signal.
// Consumers of the model may make optimisations accordingly.
func (m menuModel) IsMutable() bool {
	var _arg0 *C.GMenuModel

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean

	cret = C.g_menu_model_is_mutable(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ItemsChanged requests emission of the Model::items-changed signal on
// @model.
//
// This function should never be called except by Model subclasses. Any
// other calls to this function will very likely lead to a violation of the
// interface of the model.
//
// The implementation should update its internal representation of the menu
// before emitting the signal. The implementation should further expect to
// receive queries about the new state of the menu (and particularly added
// menu items) while signal handlers are running.
//
// The implementation must dispatch this call directly from a mainloop entry
// and not in response to calls -- particularly those from the Model API.
// Said another way: the menu must not change while user code is running
// without returning to the mainloop.
func (m menuModel) ItemsChanged(position int, removed int, added int) {
	var _arg0 *C.GMenuModel
	var _arg1 C.gint
	var _arg2 C.gint
	var _arg3 C.gint

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(position)
	_arg2 = C.gint(removed)
	_arg3 = C.gint(added)

	C.g_menu_model_items_changed(_arg0, _arg1, _arg2, _arg3)
}

// IterateItemAttributes creates a AttributeIter to iterate over the
// attributes of the item at position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
func (m menuModel) IterateItemAttributes(itemIndex int) MenuAttributeIter {
	var _arg0 *C.GMenuModel
	var _arg1 C.gint

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(itemIndex)

	var _cret *C.GMenuAttributeIter

	cret = C.g_menu_model_iterate_item_attributes(_arg0, _arg1)

	var _menuAttributeIter MenuAttributeIter

	_menuAttributeIter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MenuAttributeIter)

	return _menuAttributeIter
}

// IterateItemLinks creates a LinkIter to iterate over the links of the item
// at position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
func (m menuModel) IterateItemLinks(itemIndex int) MenuLinkIter {
	var _arg0 *C.GMenuModel
	var _arg1 C.gint

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(itemIndex)

	var _cret *C.GMenuLinkIter

	cret = C.g_menu_model_iterate_item_links(_arg0, _arg1)

	var _menuLinkIter MenuLinkIter

	_menuLinkIter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MenuLinkIter)

	return _menuLinkIter
}
