// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.gtk_recent_action_get_type()), F: marshalRecentAction},
	})
}

// RecentAction represents a list of recently used files, which can be shown by
// widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction interface {
	gextras.Objector

	// AsAction casts the class to the Action interface.
	AsAction() Action
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsRecentChooser casts the class to the RecentChooser interface.
	AsRecentChooser() RecentChooser

	// Activate emits the “activate” signal on the specified action, if it isn't
	// insensitive. This gets called by the proxy widgets when they get
	// activated.
	//
	// It can also be used to manually activate an action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	Activate()
	// BlockActivate: disable activation signals from the action
	//
	// This is needed when updating the state of your proxy Activatable widget
	// could result in calling gtk_action_activate(), this is a convenience
	// function to avoid recursing in those cases (updating toggle state for
	// instance).
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
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
	//
	// This method is inherited from Action
	ConnectAccelerator()
	// CreateIcon: this function is intended for use by action implementations
	// to create icons displayed in the proxy widgets.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateIcon(iconSize int) Widget
	// CreateMenu: if @action provides a Menu widget as a submenu for the menu
	// item or the toolbar item it creates, this function returns an instance of
	// that menu.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateMenu() Widget
	// CreateMenuItem creates a menu item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateMenuItem() Widget
	// CreateToolItem creates a toolbar item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateToolItem() Widget
	// DisconnectAccelerator undoes the effect of one call to
	// gtk_action_connect_accelerator().
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	DisconnectAccelerator()
	// GetAccelPath returns the accel path for this action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetAccelPath() string
	// GetAlwaysShowImage returns whether @action's menu item proxies will
	// always show their image, if available.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetAlwaysShowImage() bool
	// GetIconName gets the icon name of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetIconName() string
	// GetIsImportant checks whether @action is important or not
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetIsImportant() bool
	// GetLabel gets the label text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetLabel() string
	// GetName returns the name of the action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetName() string
	// GetSensitive returns whether the action itself is sensitive. Note that
	// this doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetSensitive() bool
	// GetShortLabel gets the short label text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetShortLabel() string
	// GetStockID gets the stock id of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetStockID() string
	// GetTooltip gets the tooltip text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetTooltip() string
	// GetVisible returns whether the action itself is visible. Note that this
	// doesn’t necessarily mean effective visibility. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisible() bool
	// GetVisibleHorizontal checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisibleHorizontal() bool
	// GetVisibleVertical checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisibleVertical() bool
	// IsSensitive returns whether the action is effectively sensitive.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	IsSensitive() bool
	// IsVisible returns whether the action is effectively visible.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	IsVisible() bool
	// SetAccelGroup sets the AccelGroup in which the accelerator for this
	// action will be installed.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
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
	//
	// This method is inherited from Action
	SetAccelPath(accelPath string)
	// SetAlwaysShowImage sets whether @action's menu item proxies will ignore
	// the Settings:gtk-menu-images setting and always show their image, if
	// available.
	//
	// Use this if the menu item would be useless or hard to use without their
	// image.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetAlwaysShowImage(alwaysShow bool)
	// SetIconName sets the icon name on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetIconName(iconName string)
	// SetIsImportant sets whether the action is important, this attribute is
	// used primarily by toolbar items to decide whether to show a label or not.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetIsImportant(isImportant bool)
	// SetLabel sets the label of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetLabel(label string)
	// SetSensitive sets the :sensitive property of the action to @sensitive.
	// Note that this doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetSensitive(sensitive bool)
	// SetShortLabel sets a shorter label text on @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetShortLabel(shortLabel string)
	// SetStockID sets the stock id on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetStockID(stockId string)
	// SetTooltip sets the tooltip text on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetTooltip(tooltip string)
	// SetVisible sets the :visible property of the action to @visible. Note
	// that this doesn’t necessarily mean effective visibility. See
	// gtk_action_is_visible() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisible(visible bool)
	// SetVisibleHorizontal sets whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisibleHorizontal(visibleHorizontal bool)
	// SetVisibleVertical sets whether @action is visible when vertical
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisibleVertical(visibleVertical bool)
	// UnblockActivate: reenable activation signals from the action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	UnblockActivate()
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
	// AddFilter adds @filter to the list of RecentFilter objects held by
	// @chooser.
	//
	// If no previous filter objects were defined, this function will call
	// gtk_recent_chooser_set_filter().
	//
	// This method is inherited from RecentChooser
	AddFilter(filter RecentFilter)
	// GetCurrentItem gets the RecentInfo currently selected by @chooser.
	//
	// This method is inherited from RecentChooser
	GetCurrentItem() *RecentInfo
	// GetCurrentURI gets the URI currently selected by @chooser.
	//
	// This method is inherited from RecentChooser
	GetCurrentURI() string
	// GetFilter gets the RecentFilter object currently used by @chooser to
	// affect the display of the recently used resources.
	//
	// This method is inherited from RecentChooser
	GetFilter() RecentFilter
	// GetLimit gets the number of items returned by
	// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
	//
	// This method is inherited from RecentChooser
	GetLimit() int
	// GetLocalOnly gets whether only local resources should be shown in the
	// recently used resources selector. See gtk_recent_chooser_set_local_only()
	//
	// This method is inherited from RecentChooser
	GetLocalOnly() bool
	// GetSelectMultiple gets whether @chooser can select multiple items.
	//
	// This method is inherited from RecentChooser
	GetSelectMultiple() bool
	// GetShowIcons retrieves whether @chooser should show an icon near the
	// resource.
	//
	// This method is inherited from RecentChooser
	GetShowIcons() bool
	// GetShowNotFound retrieves whether @chooser should show the recently used
	// resources that were not found.
	//
	// This method is inherited from RecentChooser
	GetShowNotFound() bool
	// GetShowPrivate returns whether @chooser should display recently used
	// resources registered as private.
	//
	// This method is inherited from RecentChooser
	GetShowPrivate() bool
	// GetShowTips gets whether @chooser should display tooltips containing the
	// full path of a recently user resource.
	//
	// This method is inherited from RecentChooser
	GetShowTips() bool
	// GetSortType gets the value set by gtk_recent_chooser_set_sort_type().
	//
	// This method is inherited from RecentChooser
	GetSortType() RecentSortType
	// RemoveFilter removes @filter from the list of RecentFilter objects held
	// by @chooser.
	//
	// This method is inherited from RecentChooser
	RemoveFilter(filter RecentFilter)
	// SelectAll selects all the items inside @chooser, if the @chooser supports
	// multiple selection.
	//
	// This method is inherited from RecentChooser
	SelectAll()
	// SelectURI selects @uri inside @chooser.
	//
	// This method is inherited from RecentChooser
	SelectURI(uri string) error
	// SetCurrentURI sets @uri as the current URI for @chooser.
	//
	// This method is inherited from RecentChooser
	SetCurrentURI(uri string) error
	// SetFilter sets @filter as the current RecentFilter object used by
	// @chooser to affect the displayed recently used resources.
	//
	// This method is inherited from RecentChooser
	SetFilter(filter RecentFilter)
	// SetLimit sets the number of items that should be returned by
	// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
	//
	// This method is inherited from RecentChooser
	SetLimit(limit int)
	// SetLocalOnly sets whether only local resources, that is resources using
	// the file:// URI scheme, should be shown in the recently used resources
	// selector. If @local_only is true (the default) then the shown resources
	// are guaranteed to be accessible through the operating system native file
	// system.
	//
	// This method is inherited from RecentChooser
	SetLocalOnly(localOnly bool)
	// SetSelectMultiple sets whether @chooser can select multiple items.
	//
	// This method is inherited from RecentChooser
	SetSelectMultiple(selectMultiple bool)
	// SetShowIcons sets whether @chooser should show an icon near the resource
	// when displaying it.
	//
	// This method is inherited from RecentChooser
	SetShowIcons(showIcons bool)
	// SetShowNotFound sets whether @chooser should display the recently used
	// resources that it didn’t find. This only applies to local resources.
	//
	// This method is inherited from RecentChooser
	SetShowNotFound(showNotFound bool)
	// SetShowPrivate: whether to show recently used resources marked registered
	// as private.
	//
	// This method is inherited from RecentChooser
	SetShowPrivate(showPrivate bool)
	// SetShowTips sets whether to show a tooltips containing the full path of
	// each recently used resource in a RecentChooser widget.
	//
	// This method is inherited from RecentChooser
	SetShowTips(showTips bool)
	// SetSortType changes the sorting order of the recently used resources list
	// displayed by @chooser.
	//
	// This method is inherited from RecentChooser
	SetSortType(sortType RecentSortType)
	// UnselectAll unselects all the items inside @chooser.
	//
	// This method is inherited from RecentChooser
	UnselectAll()
	// UnselectURI unselects @uri inside @chooser.
	//
	// This method is inherited from RecentChooser
	UnselectURI(uri string)

	// ShowNumbers returns the value set by
	// gtk_recent_chooser_menu_set_show_numbers().
	//
	// Deprecated: since version 3.10.
	ShowNumbers() bool
	// SetShowNumbers sets whether a number should be added to the items shown
	// by the widgets representing @action. The numbers are shown to provide a
	// unique character for a mnemonic to be used inside the menu item's label.
	// Only the first ten items get a number to avoid clashes.
	//
	// Deprecated: since version 3.10.
	SetShowNumbers(showNumbers bool)
}

// recentAction implements the RecentAction interface.
type recentAction struct {
	*externglib.Object
}

var _ RecentAction = (*recentAction)(nil)

// WrapRecentAction wraps a GObject to a type that implements
// interface RecentAction. It is primarily used internally.
func WrapRecentAction(obj *externglib.Object) RecentAction {
	return recentAction{obj}
}

func marshalRecentAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentAction(obj), nil
}

// NewRecentAction creates a new RecentAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRecentAction(name string, label string, tooltip string, stockId string) RecentAction {
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

	_cret = C.gtk_recent_action_new(_arg1, _arg2, _arg3, _arg4)

	var _recentAction RecentAction // out

	_recentAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RecentAction)

	return _recentAction
}

// NewRecentActionForManager creates a new RecentAction object. To add the
// action to a ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRecentActionForManager(name string, label string, tooltip string, stockId string, manager RecentManager) RecentAction {
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _arg4 *C.gchar            // out
	var _arg5 *C.GtkRecentManager // out
	var _cret *C.GtkAction        // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_action_new_for_manager(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _recentAction RecentAction // out

	_recentAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RecentAction)

	return _recentAction
}

func (r recentAction) AsAction() Action {
	return WrapAction(gextras.InternObject(r))
}

func (r recentAction) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(r))
}

func (r recentAction) AsRecentChooser() RecentChooser {
	return WrapRecentChooser(gextras.InternObject(r))
}

func (a recentAction) Activate() {
	WrapAction(gextras.InternObject(a)).Activate()
}

func (a recentAction) BlockActivate() {
	WrapAction(gextras.InternObject(a)).BlockActivate()
}

func (a recentAction) ConnectAccelerator() {
	WrapAction(gextras.InternObject(a)).ConnectAccelerator()
}

func (a recentAction) CreateIcon(iconSize int) Widget {
	return WrapAction(gextras.InternObject(a)).CreateIcon(iconSize)
}

func (a recentAction) CreateMenu() Widget {
	return WrapAction(gextras.InternObject(a)).CreateMenu()
}

func (a recentAction) CreateMenuItem() Widget {
	return WrapAction(gextras.InternObject(a)).CreateMenuItem()
}

func (a recentAction) CreateToolItem() Widget {
	return WrapAction(gextras.InternObject(a)).CreateToolItem()
}

func (a recentAction) DisconnectAccelerator() {
	WrapAction(gextras.InternObject(a)).DisconnectAccelerator()
}

func (a recentAction) GetAccelPath() string {
	return WrapAction(gextras.InternObject(a)).GetAccelPath()
}

func (a recentAction) GetAlwaysShowImage() bool {
	return WrapAction(gextras.InternObject(a)).GetAlwaysShowImage()
}

func (a recentAction) GetIconName() string {
	return WrapAction(gextras.InternObject(a)).GetIconName()
}

func (a recentAction) GetIsImportant() bool {
	return WrapAction(gextras.InternObject(a)).GetIsImportant()
}

func (a recentAction) GetLabel() string {
	return WrapAction(gextras.InternObject(a)).GetLabel()
}

func (a recentAction) GetName() string {
	return WrapAction(gextras.InternObject(a)).GetName()
}

func (a recentAction) GetSensitive() bool {
	return WrapAction(gextras.InternObject(a)).GetSensitive()
}

func (a recentAction) GetShortLabel() string {
	return WrapAction(gextras.InternObject(a)).GetShortLabel()
}

func (a recentAction) GetStockID() string {
	return WrapAction(gextras.InternObject(a)).GetStockID()
}

func (a recentAction) GetTooltip() string {
	return WrapAction(gextras.InternObject(a)).GetTooltip()
}

func (a recentAction) GetVisible() bool {
	return WrapAction(gextras.InternObject(a)).GetVisible()
}

func (a recentAction) GetVisibleHorizontal() bool {
	return WrapAction(gextras.InternObject(a)).GetVisibleHorizontal()
}

func (a recentAction) GetVisibleVertical() bool {
	return WrapAction(gextras.InternObject(a)).GetVisibleVertical()
}

func (a recentAction) IsSensitive() bool {
	return WrapAction(gextras.InternObject(a)).IsSensitive()
}

func (a recentAction) IsVisible() bool {
	return WrapAction(gextras.InternObject(a)).IsVisible()
}

func (a recentAction) SetAccelGroup(accelGroup AccelGroup) {
	WrapAction(gextras.InternObject(a)).SetAccelGroup(accelGroup)
}

func (a recentAction) SetAccelPath(accelPath string) {
	WrapAction(gextras.InternObject(a)).SetAccelPath(accelPath)
}

func (a recentAction) SetAlwaysShowImage(alwaysShow bool) {
	WrapAction(gextras.InternObject(a)).SetAlwaysShowImage(alwaysShow)
}

func (a recentAction) SetIconName(iconName string) {
	WrapAction(gextras.InternObject(a)).SetIconName(iconName)
}

func (a recentAction) SetIsImportant(isImportant bool) {
	WrapAction(gextras.InternObject(a)).SetIsImportant(isImportant)
}

func (a recentAction) SetLabel(label string) {
	WrapAction(gextras.InternObject(a)).SetLabel(label)
}

func (a recentAction) SetSensitive(sensitive bool) {
	WrapAction(gextras.InternObject(a)).SetSensitive(sensitive)
}

func (a recentAction) SetShortLabel(shortLabel string) {
	WrapAction(gextras.InternObject(a)).SetShortLabel(shortLabel)
}

func (a recentAction) SetStockID(stockId string) {
	WrapAction(gextras.InternObject(a)).SetStockID(stockId)
}

func (a recentAction) SetTooltip(tooltip string) {
	WrapAction(gextras.InternObject(a)).SetTooltip(tooltip)
}

func (a recentAction) SetVisible(visible bool) {
	WrapAction(gextras.InternObject(a)).SetVisible(visible)
}

func (a recentAction) SetVisibleHorizontal(visibleHorizontal bool) {
	WrapAction(gextras.InternObject(a)).SetVisibleHorizontal(visibleHorizontal)
}

func (a recentAction) SetVisibleVertical(visibleVertical bool) {
	WrapAction(gextras.InternObject(a)).SetVisibleVertical(visibleVertical)
}

func (a recentAction) UnblockActivate() {
	WrapAction(gextras.InternObject(a)).UnblockActivate()
}

func (b recentAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b recentAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b recentAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b recentAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b recentAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b recentAction) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b recentAction) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b recentAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b recentAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b recentAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (b recentAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b recentAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b recentAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b recentAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b recentAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b recentAction) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b recentAction) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b recentAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b recentAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b recentAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c recentAction) AddFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).AddFilter(filter)
}

func (c recentAction) GetCurrentItem() *RecentInfo {
	return WrapRecentChooser(gextras.InternObject(c)).GetCurrentItem()
}

func (c recentAction) GetCurrentURI() string {
	return WrapRecentChooser(gextras.InternObject(c)).GetCurrentURI()
}

func (c recentAction) GetFilter() RecentFilter {
	return WrapRecentChooser(gextras.InternObject(c)).GetFilter()
}

func (c recentAction) GetLimit() int {
	return WrapRecentChooser(gextras.InternObject(c)).GetLimit()
}

func (c recentAction) GetLocalOnly() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetLocalOnly()
}

func (c recentAction) GetSelectMultiple() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetSelectMultiple()
}

func (c recentAction) GetShowIcons() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetShowIcons()
}

func (c recentAction) GetShowNotFound() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetShowNotFound()
}

func (c recentAction) GetShowPrivate() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetShowPrivate()
}

func (c recentAction) GetShowTips() bool {
	return WrapRecentChooser(gextras.InternObject(c)).GetShowTips()
}

func (c recentAction) GetSortType() RecentSortType {
	return WrapRecentChooser(gextras.InternObject(c)).GetSortType()
}

func (c recentAction) RemoveFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).RemoveFilter(filter)
}

func (c recentAction) SelectAll() {
	WrapRecentChooser(gextras.InternObject(c)).SelectAll()
}

func (c recentAction) SelectURI(uri string) error {
	return WrapRecentChooser(gextras.InternObject(c)).SelectURI(uri)
}

func (c recentAction) SetCurrentURI(uri string) error {
	return WrapRecentChooser(gextras.InternObject(c)).SetCurrentURI(uri)
}

func (c recentAction) SetFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).SetFilter(filter)
}

func (c recentAction) SetLimit(limit int) {
	WrapRecentChooser(gextras.InternObject(c)).SetLimit(limit)
}

func (c recentAction) SetLocalOnly(localOnly bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetLocalOnly(localOnly)
}

func (c recentAction) SetSelectMultiple(selectMultiple bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetSelectMultiple(selectMultiple)
}

func (c recentAction) SetShowIcons(showIcons bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowIcons(showIcons)
}

func (c recentAction) SetShowNotFound(showNotFound bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowNotFound(showNotFound)
}

func (c recentAction) SetShowPrivate(showPrivate bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowPrivate(showPrivate)
}

func (c recentAction) SetShowTips(showTips bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowTips(showTips)
}

func (c recentAction) SetSortType(sortType RecentSortType) {
	WrapRecentChooser(gextras.InternObject(c)).SetSortType(sortType)
}

func (c recentAction) UnselectAll() {
	WrapRecentChooser(gextras.InternObject(c)).UnselectAll()
}

func (c recentAction) UnselectURI(uri string) {
	WrapRecentChooser(gextras.InternObject(c)).UnselectURI(uri)
}

func (a recentAction) ShowNumbers() bool {
	var _arg0 *C.GtkRecentAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_recent_action_get_show_numbers(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a recentAction) SetShowNumbers(showNumbers bool) {
	var _arg0 *C.GtkRecentAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(a.Native()))
	if showNumbers {
		_arg1 = C.TRUE
	}

	C.gtk_recent_action_set_show_numbers(_arg0, _arg1)
}
