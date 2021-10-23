// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_pad_action_type_get_type()), F: marshalPadActionType},
		{T: externglib.Type(C.gtk_pad_controller_get_type()), F: marshalPadControllerer},
	})
}

// PadActionType: type of a pad action.
type PadActionType C.gint

const (
	// PadActionButton: action is triggered by a pad button.
	PadActionButton PadActionType = iota
	// PadActionRing: action is triggered by a pad ring.
	PadActionRing
	// PadActionStrip: action is triggered by a pad strip.
	PadActionStrip
)

func marshalPadActionType(p uintptr) (interface{}, error) {
	return PadActionType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PadActionType.
func (p PadActionType) String() string {
	switch p {
	case PadActionButton:
		return "Button"
	case PadActionRing:
		return "Ring"
	case PadActionStrip:
		return "Strip"
	default:
		return fmt.Sprintf("PadActionType(%d)", p)
	}
}

// PadController: GtkPadController is an event controller for the pads found in
// drawing tablets.
//
// Pads are the collection of buttons and tactile sensors often found around the
// stylus-sensitive area.
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action. GtkPadController is provided to map those to GAction
// objects, thus letting the application give them a more semantic meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// GDK_SOURCE_TABLET_PAD devices feature multiple "modes". All these input
// elements have one current mode, which may determine the final action being
// triggered.
//
// Pad devices often divide buttons and sensors into groups. All elements in a
// group share the same current mode, but different groups may have different
// modes. See gdk.DevicePad.GetNGroups() and gdk.DevicePad.GetGroupNModes().
//
// Each of the actions that a given button/strip/ring performs for a given mode
// is defined by a gtk.PadActionEntry. It contains an action name that will be
// looked up in the given GActionGroup and activated whenever the specified
// input element and mode are triggered.
//
// A simple example of GtkPadController usage: Assigning button 1 in all modes
// and pad devices to an "invert-selection" action:
//
//    GtkPadActionEntry *pad_actions[] = {
//      { GTK_PAD_ACTION_BUTTON, 1, -1, "Invert selection", "pad-actions.invert-selection" },
//      …
//    };
//
//    …
//    action_group = g_simple_action_group_new ();
//    action = g_simple_action_new ("pad-actions.invert-selection", NULL);
//    g_signal_connect (action, "activate", on_invert_selection_activated, NULL);
//    g_action_map_add_action (G_ACTION_MAP (action_group), action);
//    …
//    pad_controller = gtk_pad_controller_new (action_group, NULL);
//
//
// The actions belonging to rings/strips will be activated with a parameter of
// type G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it is
// required that those are made stateful and accepting this GVariantType.
type PadController struct {
	EventController
}

func wrapPadController(obj *externglib.Object) *PadController {
	return &PadController{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalPadControllerer(p uintptr) (interface{}, error) {
	return wrapPadController(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPadController creates a new GtkPadController that will associate events
// from pad to actions.
//
// A NULL pad may be provided so the controller manages all pad devices
// generically, it is discouraged to mix GtkPadController objects with NULL and
// non-NULL pad argument on the same toplevel window, as execution order is not
// guaranteed.
//
// The GtkPadController is created with no mapped actions. In order to map pad
// events to actions, use gtk.PadController.SetActionEntries() or
// gtk.PadController.SetAction().
//
// Be aware that pad events will only be delivered to GtkWindows, so adding a
// pad controller to any other type of widget will not have an effect.
//
// The function takes the following parameters:
//
//    - group: GActionGroup to trigger actions from.
//    - pad: GDK_SOURCE_TABLET_PAD device, or NULL to handle all pads.
//
func NewPadController(group gio.ActionGrouper, pad gdk.Devicer) *PadController {
	var _arg1 *C.GActionGroup     // out
	var _arg2 *C.GdkDevice        // out
	var _cret *C.GtkPadController // in

	_arg1 = (*C.GActionGroup)(unsafe.Pointer(group.Native()))
	if pad != nil {
		_arg2 = (*C.GdkDevice)(unsafe.Pointer(pad.Native()))
	}

	_cret = C.gtk_pad_controller_new(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(pad)

	var _padController *PadController // out

	_padController = wrapPadController(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _padController
}

// SetAction adds an individual action to controller.
//
// This action will only be activated if the given button/ring/strip number in
// index is interacted while the current mode is mode. -1 may be used for simple
// cases, so the action is triggered on all modes.
//
// The given label should be considered user-visible, so internationalization
// rules apply. Some windowing systems may be able to use those for user
// feedback.
//
// The function takes the following parameters:
//
//    - typ: type of pad feature that will trigger this action.
//    - index: 0-indexed button/ring/strip number that will trigger this
//    action.
//    - mode that will trigger this action, or -1 for all modes.
//    - label: human readable description of this action, this string should be
//    deemed user-visible.
//    - actionName: action name that will be activated in the Group.
//
func (controller *PadController) SetAction(typ PadActionType, index, mode int, label, actionName string) {
	var _arg0 *C.GtkPadController // out
	var _arg1 C.GtkPadActionType  // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 *C.char             // out
	var _arg5 *C.char             // out

	_arg0 = (*C.GtkPadController)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkPadActionType(typ)
	_arg2 = C.int(index)
	_arg3 = C.int(mode)
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.char)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg5))

	C.gtk_pad_controller_set_action(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(index)
	runtime.KeepAlive(mode)
	runtime.KeepAlive(label)
	runtime.KeepAlive(actionName)
}

// SetActionEntries: convenience function to add a group of action entries on
// controller.
//
// See gtk.PadActionEntry and gtk.PadController.SetAction().
//
// The function takes the following parameters:
//
//    - entries: action entries to set on controller.
//
func (controller *PadController) SetActionEntries(entries []PadActionEntry) {
	var _arg0 *C.GtkPadController  // out
	var _arg1 *C.GtkPadActionEntry // out
	var _arg2 C.int

	_arg0 = (*C.GtkPadController)(unsafe.Pointer(controller.Native()))
	_arg2 = (C.int)(len(entries))
	_arg1 = (*C.GtkPadActionEntry)(C.malloc(C.size_t(uint(len(entries)) * uint(C.sizeof_GtkPadActionEntry))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GtkPadActionEntry)(_arg1), len(entries))
		for i := range entries {
			out[i] = *(*C.GtkPadActionEntry)(gextras.StructNative(unsafe.Pointer((&entries[i]))))
		}
	}

	C.gtk_pad_controller_set_action_entries(_arg0, _arg1, _arg2)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(entries)
}

// PadActionEntry: struct defining a pad action entry.
//
// An instance of this type is always passed by reference.
type PadActionEntry struct {
	*padActionEntry
}

// padActionEntry is the struct that's finalized.
type padActionEntry struct {
	native *C.GtkPadActionEntry
}

// Type: type of pad feature that will trigger this action entry.
func (p *PadActionEntry) Type() PadActionType {
	var v PadActionType // out
	v = PadActionType(p.native._type)
	return v
}

// Index: 0-indexed button/ring/strip number that will trigger this action
// entry.
func (p *PadActionEntry) Index() int {
	var v int // out
	v = int(p.native.index)
	return v
}

// Mode: mode that will trigger this action entry, or -1 for all modes.
func (p *PadActionEntry) Mode() int {
	var v int // out
	v = int(p.native.mode)
	return v
}

// Label: human readable description of this action entry, this string should be
// deemed user-visible.
func (p *PadActionEntry) Label() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(p.native.label)))
	return v
}

// ActionName: action name that will be activated in the Group.
func (p *PadActionEntry) ActionName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(p.native.action_name)))
	return v
}
