// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.gtk_movement_step_get_type()), F: marshalMovementStep},
		{T: externglib.Type(C.gtk_notebook_tab_get_type()), F: marshalNotebookTab},
		{T: externglib.Type(C.gtk_resize_mode_get_type()), F: marshalResizeMode},
		{T: externglib.Type(C.gtk_scroll_step_get_type()), F: marshalScrollStep},
		{T: externglib.Type(C.gtk_debug_flag_get_type()), F: marshalDebugFlag},
		{T: externglib.Type(C.gtk_entry_icon_accessible_get_type()), F: marshalEntryIconAccessible},
	})
}

type MovementStep int

const (
	// LogicalPositions: move forward or back by graphemes
	MovementLogicalPositions MovementStep = iota
	// VisualPositions: move left or right by graphemes
	MovementVisualPositions
	// Words: move forward or back by words
	MovementWords
	// DisplayLines: move up or down lines (wrapped lines)
	MovementDisplayLines
	// DisplayLineEnds: move to either end of a line
	MovementDisplayLineEnds
	// Paragraphs: move up or down paragraphs (newline-ended lines)
	MovementParagraphs
	// ParagraphEnds: move to either end of a paragraph
	MovementParagraphEnds
	// Pages: move by pages
	MovementPages
	// BufferEnds: move to ends of the buffer
	MovementBufferEnds
	// HorizontalPages: move horizontally by pages
	MovementHorizontalPages
)

func marshalMovementStep(p uintptr) (interface{}, error) {
	return MovementStep(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type NotebookTab int

const (
	NotebookTabFirst NotebookTab = iota
	NotebookTabLast
)

func marshalNotebookTab(p uintptr) (interface{}, error) {
	return NotebookTab(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type ResizeMode int

const (
	// Parent pass resize request to the parent
	ResizeParent ResizeMode = iota
	// Queue: queue resizes on this widget
	ResizeQueue
	// Immediate: resize immediately. Deprecated.
	ResizeImmediate
)

func marshalResizeMode(p uintptr) (interface{}, error) {
	return ResizeMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type ScrollStep int

const (
	// Steps: scroll in steps.
	ScrollSteps ScrollStep = iota
	// Pages: scroll by pages.
	ScrollPages
	// Ends: scroll to ends.
	ScrollEnds
	// HorizontalSteps: scroll in horizontal steps.
	ScrollHorizontalSteps
	// HorizontalPages: scroll by horizontal pages.
	ScrollHorizontalPages
	// HorizontalEnds: scroll to the horizontal ends.
	ScrollHorizontalEnds
)

func marshalScrollStep(p uintptr) (interface{}, error) {
	return ScrollStep(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type DebugFlag int

const (
	DebugFlagMisc         DebugFlag = 0b1
	DebugFlagPlugsocket   DebugFlag = 0b10
	DebugFlagText         DebugFlag = 0b100
	DebugFlagTree         DebugFlag = 0b1000
	DebugFlagUpdates      DebugFlag = 0b10000
	DebugFlagKeybindings  DebugFlag = 0b100000
	DebugFlagMultihead    DebugFlag = 0b1000000
	DebugFlagModules      DebugFlag = 0b10000000
	DebugFlagGeometry     DebugFlag = 0b100000000
	DebugFlagIcontheme    DebugFlag = 0b1000000000
	DebugFlagPrinting     DebugFlag = 0b10000000000
	DebugFlagBuilder      DebugFlag = 0b100000000000
	DebugFlagSizeRequest  DebugFlag = 0b1000000000000
	DebugFlagNoCSSCache   DebugFlag = 0b10000000000000
	DebugFlagBaselines    DebugFlag = 0b100000000000000
	DebugFlagPixelCache   DebugFlag = 0b1000000000000000
	DebugFlagNoPixelCache DebugFlag = 0b10000000000000000
	DebugFlagInteractive  DebugFlag = 0b100000000000000000
	DebugFlagTouchscreen  DebugFlag = 0b1000000000000000000
	DebugFlagActions      DebugFlag = 0b10000000000000000000
	DebugFlagResize       DebugFlag = 0b100000000000000000000
	DebugFlagLayout       DebugFlag = 0b1000000000000000000000
)

func marshalDebugFlag(p uintptr) (interface{}, error) {
	return DebugFlag(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type EntryIconAccessible interface {
	gextras.Objector

	// AsObject casts the class to the atk.Object interface.
	AsObject() atk.Object
	// AsAction casts the class to the atk.Action interface.
	AsAction() atk.Action

	// AddRelationship adds a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from atk.Object
	AddRelationship(relationship atk.RelationType, target atk.Object) bool
	// GetAccessibleID gets the accessible id of the accessible.
	//
	// This method is inherited from atk.Object
	GetAccessibleID() string
	// GetDescription gets the accessible description of the accessible.
	//
	// This method is inherited from atk.Object
	GetDescription() string
	// GetIndexInParent gets the 0-based index of this accessible in its parent;
	// returns -1 if the accessible does not have an accessible parent.
	//
	// This method is inherited from atk.Object
	GetIndexInParent() int
	// GetLayer gets the layer of the accessible.
	//
	// Deprecated.
	//
	// This method is inherited from atk.Object
	GetLayer() atk.Layer
	// GetMDIZOrder gets the zorder of the accessible. The value G_MININT will
	// be returned if the layer of the accessible is not ATK_LAYER_MDI.
	//
	// Deprecated.
	//
	// This method is inherited from atk.Object
	GetMDIZOrder() int
	// GetNAccessibleChildren gets the number of accessible children of the
	// accessible.
	//
	// This method is inherited from atk.Object
	GetNAccessibleChildren() int
	// GetName gets the accessible name of the accessible.
	//
	// This method is inherited from atk.Object
	GetName() string
	// GetObjectLocale gets a UTF-8 string indicating the POSIX-style
	// LC_MESSAGES locale of @accessible.
	//
	// This method is inherited from atk.Object
	GetObjectLocale() string
	// GetParent gets the accessible parent of the accessible. By default this
	// is the one assigned with atk_object_set_parent(), but it is assumed that
	// ATK implementors have ways to get the parent of the object without the
	// need of assigning it manually with atk_object_set_parent(), and will
	// return it with this method.
	//
	// If you are only interested on the parent assigned with
	// atk_object_set_parent(), use atk_object_peek_parent().
	//
	// This method is inherited from atk.Object
	GetParent() atk.Object
	// GetRole gets the role of the accessible.
	//
	// This method is inherited from atk.Object
	GetRole() atk.Role
	// Initialize: this function is called when implementing subclasses of
	// Object. It does initialization required for the new object. It is
	// intended that this function should called only in the ..._new() functions
	// used to create an instance of a subclass of Object
	//
	// This method is inherited from atk.Object
	Initialize(data interface{})
	// PeekParent gets the accessible parent of the accessible, if it has been
	// manually assigned with atk_object_set_parent. Otherwise, this function
	// returns nil.
	//
	// This method is intended as an utility for ATK implementors, and not to be
	// exposed to accessible tools. See atk_object_get_parent() for further
	// reference.
	//
	// This method is inherited from atk.Object
	PeekParent() atk.Object
	// RefAccessibleChild gets a reference to the specified accessible child of
	// the object. The accessible children are 0-based so the first accessible
	// child is at index 0, the second at index 1 and so on.
	//
	// This method is inherited from atk.Object
	RefAccessibleChild(i int) atk.Object
	// RefRelationSet gets the RelationSet associated with the object.
	//
	// This method is inherited from atk.Object
	RefRelationSet() atk.RelationSet
	// RefStateSet gets a reference to the state set of the accessible; the
	// caller must unreference it when it is no longer needed.
	//
	// This method is inherited from atk.Object
	RefStateSet() atk.StateSet
	// RemovePropertyChangeHandler removes a property change handler.
	//
	// Deprecated: since version 2.12.
	//
	// This method is inherited from atk.Object
	RemovePropertyChangeHandler(handlerId uint)
	// RemoveRelationship removes a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from atk.Object
	RemoveRelationship(relationship atk.RelationType, target atk.Object) bool
	// SetAccessibleID sets the accessible ID of the accessible. This is not
	// meant to be presented to the user, but to be an ID which is stable over
	// application development. Typically, this is the gtkbuilder ID. Such an ID
	// will be available for instance to identify a given well-known accessible
	// object for tailored screen reading, or for automatic regression testing.
	//
	// This method is inherited from atk.Object
	SetAccessibleID(name string)
	// SetDescription sets the accessible description of the accessible. You
	// can't set the description to NULL. This is reserved for the initial
	// value. In this aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to
	// set the name to a empty value you can use "".
	//
	// This method is inherited from atk.Object
	SetDescription(description string)
	// SetName sets the accessible name of the accessible. You can't set the
	// name to NULL. This is reserved for the initial value. In this aspect NULL
	// is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a empty
	// value you can use "".
	//
	// This method is inherited from atk.Object
	SetName(name string)
	// SetParent sets the accessible parent of the accessible. @parent can be
	// NULL.
	//
	// This method is inherited from atk.Object
	SetParent(parent atk.Object)
	// SetRole sets the role of the accessible.
	//
	// This method is inherited from atk.Object
	SetRole(role atk.Role)
	// DoAction: perform the specified action on the object.
	//
	// This method is inherited from atk.Action
	DoAction(i int) bool
	// GetDescription returns a description of the specified action of the
	// object.
	//
	// This method is inherited from atk.Action
	GetDescription(i int) string
	// GetKeybinding gets the keybinding which can be used to activate this
	// action, if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button. - The sequence is the full list of keys which
	// invoke the action even if the relevant element is not currently shown on
	// screen. For instance, for a menu item the sequence is the keybindings
	// used to open the parent menus before invoking. The sequence string is
	// colon-delimited. Example: "Alt+F:N" in a traditional "New..." menu item.
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	//
	// This method is inherited from atk.Action
	GetKeybinding(i int) string
	// GetLocalizedName returns the localized name of the specified action of
	// the object.
	//
	// This method is inherited from atk.Action
	GetLocalizedName(i int) string
	// GetNActions gets the number of accessible actions available on the
	// object. If there are more than one, the first one is considered the
	// "default" action of the object.
	//
	// This method is inherited from atk.Action
	GetNActions() int
	// GetName returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	//
	// This method is inherited from atk.Action
	GetName(i int) string
	// SetDescription sets a description of the specified action of the object.
	//
	// This method is inherited from atk.Action
	SetDescription(i int, desc string) bool
}

// entryIconAccessible implements the EntryIconAccessible interface.
type entryIconAccessible struct {
	*externglib.Object
}

var _ EntryIconAccessible = (*entryIconAccessible)(nil)

// WrapEntryIconAccessible wraps a GObject to a type that implements
// interface EntryIconAccessible. It is primarily used internally.
func WrapEntryIconAccessible(obj *externglib.Object) EntryIconAccessible {
	return entryIconAccessible{obj}
}

func marshalEntryIconAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntryIconAccessible(obj), nil
}

func (e entryIconAccessible) AsObject() atk.Object {
	return atk.WrapObject(gextras.InternObject(e))
}

func (e entryIconAccessible) AsAction() atk.Action {
	return atk.WrapAction(gextras.InternObject(e))
}

func (o entryIconAccessible) AddRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).AddRelationship(relationship, target)
}

func (a entryIconAccessible) GetAccessibleID() string {
	return atk.WrapObject(gextras.InternObject(a)).GetAccessibleID()
}

func (a entryIconAccessible) GetDescription() string {
	return atk.WrapObject(gextras.InternObject(a)).GetDescription()
}

func (a entryIconAccessible) GetIndexInParent() int {
	return atk.WrapObject(gextras.InternObject(a)).GetIndexInParent()
}

func (a entryIconAccessible) GetLayer() atk.Layer {
	return atk.WrapObject(gextras.InternObject(a)).GetLayer()
}

func (a entryIconAccessible) GetMDIZOrder() int {
	return atk.WrapObject(gextras.InternObject(a)).GetMDIZOrder()
}

func (a entryIconAccessible) GetNAccessibleChildren() int {
	return atk.WrapObject(gextras.InternObject(a)).GetNAccessibleChildren()
}

func (a entryIconAccessible) GetName() string {
	return atk.WrapObject(gextras.InternObject(a)).GetName()
}

func (a entryIconAccessible) GetObjectLocale() string {
	return atk.WrapObject(gextras.InternObject(a)).GetObjectLocale()
}

func (a entryIconAccessible) GetParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).GetParent()
}

func (a entryIconAccessible) GetRole() atk.Role {
	return atk.WrapObject(gextras.InternObject(a)).GetRole()
}

func (a entryIconAccessible) Initialize(data interface{}) {
	atk.WrapObject(gextras.InternObject(a)).Initialize(data)
}

func (a entryIconAccessible) PeekParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).PeekParent()
}

func (a entryIconAccessible) RefAccessibleChild(i int) atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).RefAccessibleChild(i)
}

func (a entryIconAccessible) RefRelationSet() atk.RelationSet {
	return atk.WrapObject(gextras.InternObject(a)).RefRelationSet()
}

func (a entryIconAccessible) RefStateSet() atk.StateSet {
	return atk.WrapObject(gextras.InternObject(a)).RefStateSet()
}

func (a entryIconAccessible) RemovePropertyChangeHandler(handlerId uint) {
	atk.WrapObject(gextras.InternObject(a)).RemovePropertyChangeHandler(handlerId)
}

func (o entryIconAccessible) RemoveRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).RemoveRelationship(relationship, target)
}

func (a entryIconAccessible) SetAccessibleID(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetAccessibleID(name)
}

func (a entryIconAccessible) SetDescription(description string) {
	atk.WrapObject(gextras.InternObject(a)).SetDescription(description)
}

func (a entryIconAccessible) SetName(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetName(name)
}

func (a entryIconAccessible) SetParent(parent atk.Object) {
	atk.WrapObject(gextras.InternObject(a)).SetParent(parent)
}

func (a entryIconAccessible) SetRole(role atk.Role) {
	atk.WrapObject(gextras.InternObject(a)).SetRole(role)
}

func (a entryIconAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a entryIconAccessible) GetDescription(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetDescription(i)
}

func (a entryIconAccessible) GetKeybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetKeybinding(i)
}

func (a entryIconAccessible) GetLocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetLocalizedName(i)
}

func (a entryIconAccessible) GetNActions() int {
	return atk.WrapAction(gextras.InternObject(a)).GetNActions()
}

func (a entryIconAccessible) GetName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetName(i)
}

func (a entryIconAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}
