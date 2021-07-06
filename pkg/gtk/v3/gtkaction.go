// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_action_get_type()), F: marshalAction},
	})
}

// Action: > In GTK+ 3.10, GtkAction has been deprecated. Use #GAction >
// instead, and associate actions with Actionable widgets. Use > Model for
// creating menus with gtk_menu_new_from_model().
//
// Actions represent operations that the user can be perform, along with some
// information how it should be presented in the interface. Each action provides
// methods to create icons, menu items and toolbar items representing itself.
//
// As well as the callback that is called when the action gets activated, the
// following also gets associated with the action:
//
// - a name (not translated, for path lookup)
//
// - a label (translated, for display)
//
// - an accelerator
//
// - whether label indicates a stock id
//
// - a tooltip (optional, translated)
//
// - a toolbar label (optional, shorter than label)
//
//    The action will also have some state information:
//
// - visible (shown/hidden)
//
// - sensitive (enabled/disabled)
//
// Apart from regular actions, there are [toggle actions][GtkToggleAction],
// which can be toggled between two states and [radio actions][GtkRadioAction],
// of which only one in a group can be in the “active” state. Other actions can
// be implemented as Action subclasses.
//
// Each action can have one or more proxy widgets. To act as an action proxy,
// widget needs to implement Activatable interface. Proxies mirror the state of
// the action and should change when the action’s state changes. Properties that
// are always mirrored by proxies are Action:sensitive and Action:visible.
// Action:gicon, Action:icon-name, Action:label, Action:short-label and
// Action:stock-id properties are only mirorred if proxy widget has
// Activatable:use-action-appearance property set to true.
//
// When the proxy is activated, it should activate its action.
type Action interface {
	gextras.Objector

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	//
	// This method is inherited from Buildable
	AddChild(builder Builder, child gextras.Objector, typ string)
	// ConstructChild constructs a child of @buildable with the name @name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// This method is inherited from Buildable
	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the @buildable.
	//
	// This method is inherited from Buildable
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// This method is inherited from Buildable
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// This method is inherited from Buildable
	CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool)
	// GetInternalChild: get the internal child called @childname of the
	// @buildable object.
	//
	// This method is inherited from Buildable
	GetInternalChild(builder Builder, childname string) gextras.Objector
	// GetName gets the name of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetName() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// This method is inherited from Buildable
	ParserFinished(builder Builder)
	// SetBuildableProperty sets the property name @name to @value on the
	// @buildable object.
	//
	// This method is inherited from Buildable
	SetBuildableProperty(builder Builder, name string, value externglib.Value)
	// SetName sets the name of the @buildable object.
	//
	// This method is inherited from Buildable
	SetName(name string)

	// Activate emits the “activate” signal on the specified action, if it isn't
	// insensitive. This gets called by the proxy widgets when they get
	// activated.
	//
	// It can also be used to manually activate an action.
	//
	// Deprecated: since version 3.10.
	Activate()
	// BlockActivate: disable activation signals from the action
	//
	// This is needed when updating the state of your proxy Activatable widget
	// could result in calling gtk_action_activate(), this is a convenience
	// function to avoid recursing in those cases (updating toggle state for
	// instance).
	//
	// Deprecated: since version 3.10.
	BlockActivate()
	// ConnectAccelerator installs the accelerator for @action if @action has an
	// accel path and group. See gtk_action_set_accel_path() and
	// gtk_action_set_accel_group()
	//
	// Since multiple proxies may independently trigger the installation of the
	// accelerator, the @action counts the number of times this function has
	// been called and doesn’t remove the accelerator until
	// gtk_action_disconnect_accelerator() has been called as many times.
	//
	// Deprecated: since version 3.10.
	ConnectAccelerator()
	// CreateIcon: this function is intended for use by action implementations
	// to create icons displayed in the proxy widgets.
	//
	// Deprecated: since version 3.10.
	CreateIcon(iconSize int) Widget
	// CreateMenu: if @action provides a Menu widget as a submenu for the menu
	// item or the toolbar item it creates, this function returns an instance of
	// that menu.
	//
	// Deprecated: since version 3.10.
	CreateMenu() Widget
	// CreateMenuItem creates a menu item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	CreateMenuItem() Widget
	// CreateToolItem creates a toolbar item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	CreateToolItem() Widget
	// DisconnectAccelerator undoes the effect of one call to
	// gtk_action_connect_accelerator().
	//
	// Deprecated: since version 3.10.
	DisconnectAccelerator()
	// AccelPath returns the accel path for this action.
	//
	// Deprecated: since version 3.10.
	AccelPath() string
	// AlwaysShowImage returns whether @action's menu item proxies will always
	// show their image, if available.
	//
	// Deprecated: since version 3.10.
	AlwaysShowImage() bool
	// IconName gets the icon name of @action.
	//
	// Deprecated: since version 3.10.
	IconName() string
	// IsImportant checks whether @action is important or not
	//
	// Deprecated: since version 3.10.
	IsImportant() bool
	// Label gets the label text of @action.
	//
	// Deprecated: since version 3.10.
	Label() string
	// Name returns the name of the action.
	//
	// Deprecated: since version 3.10.
	Name() string
	// Sensitive returns whether the action itself is sensitive. Note that this
	// doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	Sensitive() bool
	// ShortLabel gets the short label text of @action.
	//
	// Deprecated: since version 3.10.
	ShortLabel() string
	// StockID gets the stock id of @action.
	//
	// Deprecated: since version 3.10.
	StockID() string
	// Tooltip gets the tooltip text of @action.
	//
	// Deprecated: since version 3.10.
	Tooltip() string
	// Visible returns whether the action itself is visible. Note that this
	// doesn’t necessarily mean effective visibility. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	Visible() bool
	// VisibleHorizontal checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	VisibleHorizontal() bool
	// VisibleVertical checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	VisibleVertical() bool
	// IsSensitive returns whether the action is effectively sensitive.
	//
	// Deprecated: since version 3.10.
	IsSensitive() bool
	// IsVisible returns whether the action is effectively visible.
	//
	// Deprecated: since version 3.10.
	IsVisible() bool
	// SetAccelGroup sets the AccelGroup in which the accelerator for this
	// action will be installed.
	//
	// Deprecated: since version 3.10.
	SetAccelGroup(accelGroup AccelGroup)
	// SetAccelPath sets the accel path for this action. All proxy widgets
	// associated with the action will have this accel path, so that their
	// accelerators are consistent.
	//
	// Note that @accel_path string will be stored in a #GQuark. Therefore, if
	// you pass a static string, you can save some memory by interning it first
	// with g_intern_static_string().
	//
	// Deprecated: since version 3.10.
	SetAccelPath(accelPath string)
	// SetAlwaysShowImage sets whether @action's menu item proxies will ignore
	// the Settings:gtk-menu-images setting and always show their image, if
	// available.
	//
	// Use this if the menu item would be useless or hard to use without their
	// image.
	//
	// Deprecated: since version 3.10.
	SetAlwaysShowImage(alwaysShow bool)
	// SetIconName sets the icon name on @action
	//
	// Deprecated: since version 3.10.
	SetIconName(iconName string)
	// SetIsImportant sets whether the action is important, this attribute is
	// used primarily by toolbar items to decide whether to show a label or not.
	//
	// Deprecated: since version 3.10.
	SetIsImportant(isImportant bool)
	// SetLabel sets the label of @action.
	//
	// Deprecated: since version 3.10.
	SetLabel(label string)
	// SetSensitive sets the :sensitive property of the action to @sensitive.
	// Note that this doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	SetSensitive(sensitive bool)
	// SetShortLabel sets a shorter label text on @action.
	//
	// Deprecated: since version 3.10.
	SetShortLabel(shortLabel string)
	// SetStockID sets the stock id on @action
	//
	// Deprecated: since version 3.10.
	SetStockID(stockId string)
	// SetTooltip sets the tooltip text on @action
	//
	// Deprecated: since version 3.10.
	SetTooltip(tooltip string)
	// SetVisible sets the :visible property of the action to @visible. Note
	// that this doesn’t necessarily mean effective visibility. See
	// gtk_action_is_visible() for that.
	//
	// Deprecated: since version 3.10.
	SetVisible(visible bool)
	// SetVisibleHorizontal sets whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	SetVisibleHorizontal(visibleHorizontal bool)
	// SetVisibleVertical sets whether @action is visible when vertical
	//
	// Deprecated: since version 3.10.
	SetVisibleVertical(visibleVertical bool)
	// UnblockActivate: reenable activation signals from the action
	//
	// Deprecated: since version 3.10.
	UnblockActivate()
}

// action implements the Action interface.
type action struct {
	*externglib.Object
}

var _ Action = (*action)(nil)

// WrapAction wraps a GObject to a type that implements
// interface Action. It is primarily used internally.
func WrapAction(obj *externglib.Object) Action {
	return action{obj}
}

func marshalAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAction(obj), nil
}

// NewAction creates a new Action object. To add the action to a ActionGroup and
// set the accelerator for the action, call
// gtk_action_group_add_action_with_accel(). See the [UI Definition
// section][XML-UI] for information on allowed action names.
//
// Deprecated: since version 3.10.
func NewAction(name string, label string, tooltip string, stockId string) Action {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_action_new(_arg1, _arg2, _arg3, _arg4)

	var _action Action // out

	_action = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (a action) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(a))
}

func (b action) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b action) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b action) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b action) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b action) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b action) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b action) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b action) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b action) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b action) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a action) Activate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_activate(_arg0)
}

func (a action) BlockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_block_activate(_arg0)
}

func (a action) ConnectAccelerator() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_connect_accelerator(_arg0)
}

func (a action) CreateIcon(iconSize int) Widget {
	var _arg0 *C.GtkAction  // out
	var _arg1 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	_cret = C.gtk_action_create_icon(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateMenu() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_menu(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateMenuItem() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_menu_item(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateToolItem() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_tool_item(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) DisconnectAccelerator() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_disconnect_accelerator(_arg0)
}

func (a action) AccelPath() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_accel_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) AlwaysShowImage() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_always_show_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IconName() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) IsImportant() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_is_important(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) Label() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Name() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Sensitive() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) ShortLabel() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_short_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) StockID() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_stock_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Tooltip() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_tooltip(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Visible() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) VisibleHorizontal() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible_horizontal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) VisibleVertical() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible_vertical(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IsSensitive() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_is_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IsVisible() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) SetAccelGroup(accelGroup AccelGroup) {
	var _arg0 *C.GtkAction     // out
	var _arg1 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_action_set_accel_group(_arg0, _arg1)
}

func (a action) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_accel_path(_arg0, _arg1)
}

func (a action) SetAlwaysShowImage(alwaysShow bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_always_show_image(_arg0, _arg1)
}

func (a action) SetIconName(iconName string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_icon_name(_arg0, _arg1)
}

func (a action) SetIsImportant(isImportant bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_is_important(_arg0, _arg1)
}

func (a action) SetLabel(label string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_label(_arg0, _arg1)
}

func (a action) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_sensitive(_arg0, _arg1)
}

func (a action) SetShortLabel(shortLabel string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(shortLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_short_label(_arg0, _arg1)
}

func (a action) SetStockID(stockId string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_stock_id(_arg0, _arg1)
}

func (a action) SetTooltip(tooltip string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_tooltip(_arg0, _arg1)
}

func (a action) SetVisible(visible bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible(_arg0, _arg1)
}

func (a action) SetVisibleHorizontal(visibleHorizontal bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_horizontal(_arg0, _arg1)
}

func (a action) SetVisibleVertical(visibleVertical bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_vertical(_arg0, _arg1)
}

func (a action) UnblockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_unblock_activate(_arg0)
}
