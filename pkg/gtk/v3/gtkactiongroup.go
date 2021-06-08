// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_action_group_get_type()), F: marshalActionGroup},
	})
}

// ActionGroup actions are organised into groups. An action group is essentially
// a map from names to Action objects.
//
// All actions that would make sense to use in a particular context should be in
// a single group. Multiple action groups may be used for a particular user
// interface. In fact, it is expected that most nontrivial applications will
// make use of multiple groups. For example, in an application that can edit
// multiple documents, one group holding global actions (e.g. quit, about, new),
// and one group per document holding actions that act on that document (eg.
// save, cut/copy/paste, etc). Each window’s menus would be constructed from a
// combination of two action groups.
//
//
// Accelerators
//
// Accelerators are handled by the GTK+ accelerator map. All actions are
// assigned an accelerator path (which normally has the form
// `<Actions>/group-name/action-name`) and a shortcut is associated with this
// accelerator path. All menuitems and toolitems take on this accelerator path.
// The GTK+ accelerator map code makes sure that the correct shortcut is
// displayed next to the menu item.
//
//
// GtkActionGroup as GtkBuildable
//
// The ActionGroup implementation of the Buildable interface accepts Action
// objects as <child> elements in UI definitions.
//
// Note that it is probably more common to define actions and action groups in
// the code, since they are directly related to what the code can do.
//
// The GtkActionGroup implementation of the GtkBuildable interface supports a
// custom <accelerator> element, which has attributes named “key“ and
// “modifiers“ and allows to specify accelerators. This is similar to the
// <accelerator> element of Widget, the main difference is that it doesn’t allow
// you to specify a signal.
//
// A Dialog UI definition fragment. ##
//
//    <object class="GtkActionGroup" id="actiongroup">
//      <child>
//          <object class="GtkAction" id="About">
//              <property name="name">About</property>
//              <property name="stock_id">gtk-about</property>
//              <signal handler="about_activate" name="activate"/>
//          </object>
//          <accelerator key="F1" modifiers="GDK_CONTROL_MASK | GDK_SHIFT_MASK"/>
//      </child>
//    </object>
type ActionGroup interface {
	gextras.Objector
	Buildable

	// AddAction adds an action object to the action group. Note that this
	// function does not set up the accel path of the action, which can lead to
	// problems if a user tries to modify the accelerator of a menuitem
	// associated with the action. Therefore you must either set the accel path
	// yourself with gtk_action_set_accel_path(), or use
	// `gtk_action_group_add_action_with_accel (..., NULL)`.
	AddAction(action Action)
	// AddActionWithAccel adds an action object to the action group and sets up
	// the accelerator.
	//
	// If @accelerator is nil, attempts to use the accelerator associated with
	// the stock_id of the action.
	//
	// Accel paths are set to `<Actions>/group-name/action-name`.
	AddActionWithAccel(action Action, accelerator string)
	// AddRadioActions: this is a convenience routine to create a group of radio
	// actions and add them to the action group.
	//
	// The “changed” signal of the first radio action is connected to the
	// @on_change callback and the accel paths of the actions are set to
	// `<Actions>/group-name/action-name`.
	AddRadioActions()
	// AddRadioActionsFull: this variant of gtk_action_group_add_radio_actions()
	// adds a Notify callback for @user_data.
	AddRadioActionsFull()
	// AccelGroup gets the accelerator group.
	AccelGroup() AccelGroup
	// Action looks up an action in the action group by name.
	Action(actionName string) Action
	// Name gets the name of the action group.
	Name() string
	// Sensitive returns true if the group is sensitive. The constituent actions
	// can only be logically sensitive (see gtk_action_is_sensitive()) if they
	// are sensitive (see gtk_action_get_sensitive()) and their group is
	// sensitive.
	Sensitive() bool
	// Visible returns true if the group is visible. The constituent actions can
	// only be logically visible (see gtk_action_is_visible()) if they are
	// visible (see gtk_action_get_visible()) and their group is visible.
	Visible() bool
	// ListActions lists the actions in the action group.
	ListActions() *glib.List
	// RemoveAction removes an action object from the action group.
	RemoveAction(action Action)
	// SetAccelGroup sets the accelerator group to be used by every action in
	// this group.
	SetAccelGroup(accelGroup AccelGroup)
	// SetSensitive changes the sensitivity of @action_group
	SetSensitive(sensitive bool)
	// SetTranslateFunc sets a function to be used for translating the @label
	// and @tooltip of ActionEntrys added by gtk_action_group_add_actions().
	//
	// If you’re using gettext(), it is enough to set the translation domain
	// with gtk_action_group_set_translation_domain().
	SetTranslateFunc()
	// SetTranslationDomain sets the translation domain and uses g_dgettext()
	// for translating the @label and @tooltip of ActionEntrys added by
	// gtk_action_group_add_actions().
	//
	// If you’re not using gettext() for localization, see
	// gtk_action_group_set_translate_func().
	SetTranslationDomain(domain string)
	// SetVisible changes the visible of @action_group.
	SetVisible(visible bool)
	// TranslateString translates a string using the function set with
	// gtk_action_group_set_translate_func(). This is mainly intended for
	// language bindings.
	TranslateString(string string) string
}

// actionGroup implements the ActionGroup interface.
type actionGroup struct {
	gextras.Objector
	Buildable
}

var _ ActionGroup = (*actionGroup)(nil)

// WrapActionGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapActionGroup(obj *externglib.Object) ActionGroup {
	return ActionGroup{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActionGroup(obj), nil
}

// NewActionGroup constructs a class ActionGroup.
func NewActionGroup(name string) ActionGroup {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	cret := new(C.GtkActionGroup)
	var goret ActionGroup

	cret = C.gtk_action_group_new(arg1)

	goret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ActionGroup)

	return goret
}

// AddAction adds an action object to the action group. Note that this
// function does not set up the accel path of the action, which can lead to
// problems if a user tries to modify the accelerator of a menuitem
// associated with the action. Therefore you must either set the accel path
// yourself with gtk_action_set_accel_path(), or use
// `gtk_action_group_add_action_with_accel (..., NULL)`.
func (a actionGroup) AddAction(action Action) {
	var arg0 *C.GtkActionGroup
	var arg1 *C.GtkAction

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_add_action(arg0, arg1)
}

// AddActionWithAccel adds an action object to the action group and sets up
// the accelerator.
//
// If @accelerator is nil, attempts to use the accelerator associated with
// the stock_id of the action.
//
// Accel paths are set to `<Actions>/group-name/action-name`.
func (a actionGroup) AddActionWithAccel(action Action, accelerator string) {
	var arg0 *C.GtkActionGroup
	var arg1 *C.GtkAction
	var arg2 *C.gchar

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	arg2 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_action_group_add_action_with_accel(arg0, arg1, arg2)
}

// AddRadioActions: this is a convenience routine to create a group of radio
// actions and add them to the action group.
//
// The “changed” signal of the first radio action is connected to the
// @on_change callback and the accel paths of the actions are set to
// `<Actions>/group-name/action-name`.
func (a actionGroup) AddRadioActions() {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	C.gtk_action_group_add_radio_actions(arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddRadioActionsFull: this variant of gtk_action_group_add_radio_actions()
// adds a Notify callback for @user_data.
func (a actionGroup) AddRadioActionsFull() {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	C.gtk_action_group_add_radio_actions_full(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AccelGroup gets the accelerator group.
func (a actionGroup) AccelGroup() AccelGroup {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	var cret *C.GtkAccelGroup
	var goret AccelGroup

	cret = C.gtk_action_group_get_accel_group(arg0)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(AccelGroup)

	return goret
}

// Action looks up an action in the action group by name.
func (a actionGroup) Action(actionName string) Action {
	var arg0 *C.GtkActionGroup
	var arg1 *C.gchar

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GtkAction
	var goret Action

	cret = C.gtk_action_group_get_action(arg0, arg1)

	goret = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Action)

	return goret
}

// Name gets the name of the action group.
func (a actionGroup) Name() string {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	var cret *C.gchar
	var goret string

	cret = C.gtk_action_group_get_name(arg0)

	goret = C.GoString(cret)

	return goret
}

// Sensitive returns true if the group is sensitive. The constituent actions
// can only be logically sensitive (see gtk_action_is_sensitive()) if they
// are sensitive (see gtk_action_get_sensitive()) and their group is
// sensitive.
func (a actionGroup) Sensitive() bool {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_action_group_get_sensitive(arg0)

	if cret {
		goret = true
	}

	return goret
}

// Visible returns true if the group is visible. The constituent actions can
// only be logically visible (see gtk_action_is_visible()) if they are
// visible (see gtk_action_get_visible()) and their group is visible.
func (a actionGroup) Visible() bool {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var goret bool

	cret = C.gtk_action_group_get_visible(arg0)

	if cret {
		goret = true
	}

	return goret
}

// ListActions lists the actions in the action group.
func (a actionGroup) ListActions() *glib.List {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	cret := new(C.GList)
	var goret *glib.List

	cret = C.gtk_action_group_list_actions(arg0)

	goret = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// RemoveAction removes an action object from the action group.
func (a actionGroup) RemoveAction(action Action) {
	var arg0 *C.GtkActionGroup
	var arg1 *C.GtkAction

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_remove_action(arg0, arg1)
}

// SetAccelGroup sets the accelerator group to be used by every action in
// this group.
func (a actionGroup) SetAccelGroup(accelGroup AccelGroup) {
	var arg0 *C.GtkActionGroup
	var arg1 *C.GtkAccelGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_action_group_set_accel_group(arg0, arg1)
}

// SetSensitive changes the sensitivity of @action_group
func (a actionGroup) SetSensitive(sensitive bool) {
	var arg0 *C.GtkActionGroup
	var arg1 C.gboolean

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	if sensitive {
		arg1 = C.gboolean(1)
	}

	C.gtk_action_group_set_sensitive(arg0, arg1)
}

// SetTranslateFunc sets a function to be used for translating the @label
// and @tooltip of ActionEntrys added by gtk_action_group_add_actions().
//
// If you’re using gettext(), it is enough to set the translation domain
// with gtk_action_group_set_translation_domain().
func (a actionGroup) SetTranslateFunc() {
	var arg0 *C.GtkActionGroup

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	C.gtk_action_group_set_translate_func(arg0, arg1, arg2, arg3)
}

// SetTranslationDomain sets the translation domain and uses g_dgettext()
// for translating the @label and @tooltip of ActionEntrys added by
// gtk_action_group_add_actions().
//
// If you’re not using gettext() for localization, see
// gtk_action_group_set_translate_func().
func (a actionGroup) SetTranslationDomain(domain string) {
	var arg0 *C.GtkActionGroup
	var arg1 *C.gchar

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_action_group_set_translation_domain(arg0, arg1)
}

// SetVisible changes the visible of @action_group.
func (a actionGroup) SetVisible(visible bool) {
	var arg0 *C.GtkActionGroup
	var arg1 C.gboolean

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	if visible {
		arg1 = C.gboolean(1)
	}

	C.gtk_action_group_set_visible(arg0, arg1)
}

// TranslateString translates a string using the function set with
// gtk_action_group_set_translate_func(). This is mainly intended for
// language bindings.
func (a actionGroup) TranslateString(string string) string {
	var arg0 *C.GtkActionGroup
	var arg1 *C.gchar

	arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret string

	cret = C.gtk_action_group_translate_string(arg0, arg1)

	goret = C.GoString(cret)

	return goret
}

// ActionEntry structs are used with gtk_action_group_add_actions() to construct
// actions.
type ActionEntry struct {
	native C.GtkActionEntry
}

// WrapActionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapActionEntry(ptr unsafe.Pointer) *ActionEntry {
	if ptr == nil {
		return nil
	}

	return (*ActionEntry)(ptr)
}

func marshalActionEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapActionEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *ActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Name gets the field inside the struct.
func (a *ActionEntry) Name() string {
	var v string
	v = C.GoString(a.native.name)
	return v
}

// StockID gets the field inside the struct.
func (a *ActionEntry) StockID() string {
	var v string
	v = C.GoString(a.native.stock_id)
	return v
}

// Label gets the field inside the struct.
func (a *ActionEntry) Label() string {
	var v string
	v = C.GoString(a.native.label)
	return v
}

// Accelerator gets the field inside the struct.
func (a *ActionEntry) Accelerator() string {
	var v string
	v = C.GoString(a.native.accelerator)
	return v
}

// Tooltip gets the field inside the struct.
func (a *ActionEntry) Tooltip() string {
	var v string
	v = C.GoString(a.native.tooltip)
	return v
}

// RadioActionEntry structs are used with gtk_action_group_add_radio_actions()
// to construct groups of radio actions.
type RadioActionEntry struct {
	native C.GtkRadioActionEntry
}

// WrapRadioActionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRadioActionEntry(ptr unsafe.Pointer) *RadioActionEntry {
	if ptr == nil {
		return nil
	}

	return (*RadioActionEntry)(ptr)
}

func marshalRadioActionEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRadioActionEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RadioActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Name gets the field inside the struct.
func (r *RadioActionEntry) Name() string {
	var v string
	v = C.GoString(r.native.name)
	return v
}

// StockID gets the field inside the struct.
func (r *RadioActionEntry) StockID() string {
	var v string
	v = C.GoString(r.native.stock_id)
	return v
}

// Label gets the field inside the struct.
func (r *RadioActionEntry) Label() string {
	var v string
	v = C.GoString(r.native.label)
	return v
}

// Accelerator gets the field inside the struct.
func (r *RadioActionEntry) Accelerator() string {
	var v string
	v = C.GoString(r.native.accelerator)
	return v
}

// Tooltip gets the field inside the struct.
func (r *RadioActionEntry) Tooltip() string {
	var v string
	v = C.GoString(r.native.tooltip)
	return v
}

// Value gets the field inside the struct.
func (r *RadioActionEntry) Value() int {
	var v int
	v = int(r.native.value)
	return v
}

// ToggleActionEntry structs are used with gtk_action_group_add_toggle_actions()
// to construct toggle actions.
type ToggleActionEntry struct {
	native C.GtkToggleActionEntry
}

// WrapToggleActionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToggleActionEntry(ptr unsafe.Pointer) *ToggleActionEntry {
	if ptr == nil {
		return nil
	}

	return (*ToggleActionEntry)(ptr)
}

func marshalToggleActionEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapToggleActionEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ToggleActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Name gets the field inside the struct.
func (t *ToggleActionEntry) Name() string {
	var v string
	v = C.GoString(t.native.name)
	return v
}

// StockID gets the field inside the struct.
func (t *ToggleActionEntry) StockID() string {
	var v string
	v = C.GoString(t.native.stock_id)
	return v
}

// Label gets the field inside the struct.
func (t *ToggleActionEntry) Label() string {
	var v string
	v = C.GoString(t.native.label)
	return v
}

// Accelerator gets the field inside the struct.
func (t *ToggleActionEntry) Accelerator() string {
	var v string
	v = C.GoString(t.native.accelerator)
	return v
}

// Tooltip gets the field inside the struct.
func (t *ToggleActionEntry) Tooltip() string {
	var v string
	v = C.GoString(t.native.tooltip)
	return v
}

// IsActive gets the field inside the struct.
func (t *ToggleActionEntry) IsActive() bool {
	var v bool
	if t.native.is_active {
		v = true
	}
	return v
}
