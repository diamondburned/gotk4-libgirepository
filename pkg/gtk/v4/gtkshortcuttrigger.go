// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_alternative_trigger_get_type()), F: marshalAlternativeTriggerer},
		{T: externglib.Type(C.gtk_keyval_trigger_get_type()), F: marshalKeyvalTriggerer},
		{T: externglib.Type(C.gtk_mnemonic_trigger_get_type()), F: marshalMnemonicTriggerer},
		{T: externglib.Type(C.gtk_never_trigger_get_type()), F: marshalNeverTriggerer},
		{T: externglib.Type(C.gtk_shortcut_trigger_get_type()), F: marshalShortcutTriggerer},
	})
}

// AlternativeTrigger: GtkShortcutTrigger that combines two triggers.
//
// The GtkAlternativeTrigger triggers when either of two trigger.
//
// This can be cascaded to combine more than two triggers.
type AlternativeTrigger struct {
	ShortcutTrigger
}

func wrapAlternativeTrigger(obj *externglib.Object) *AlternativeTrigger {
	return &AlternativeTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalAlternativeTriggerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAlternativeTrigger(obj), nil
}

// NewAlternativeTrigger creates a GtkShortcutTrigger that will trigger whenever
// either of the two given triggers gets triggered.
//
// Note that nesting is allowed, so if you want more than two alternative,
// create a new alternative trigger for each option.
func NewAlternativeTrigger(first ShortcutTriggerer, second ShortcutTriggerer) *AlternativeTrigger {
	var _arg1 *C.GtkShortcutTrigger // out
	var _arg2 *C.GtkShortcutTrigger // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(first.Native()))
	_arg2 = (*C.GtkShortcutTrigger)(unsafe.Pointer(second.Native()))

	_cret = C.gtk_alternative_trigger_new(_arg1, _arg2)

	var _alternativeTrigger *AlternativeTrigger // out

	_alternativeTrigger = wrapAlternativeTrigger(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _alternativeTrigger
}

// First gets the first of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetSecond() will return the other one.
func (self *AlternativeTrigger) First() ShortcutTriggerer {
	var _arg0 *C.GtkAlternativeTrigger // out
	var _cret *C.GtkShortcutTrigger    // in

	_arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_alternative_trigger_get_first(_arg0)

	var _shortcutTrigger ShortcutTriggerer // out

	_shortcutTrigger = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(ShortcutTriggerer)

	return _shortcutTrigger
}

// Second gets the second of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetFirst() will return the other one.
func (self *AlternativeTrigger) Second() ShortcutTriggerer {
	var _arg0 *C.GtkAlternativeTrigger // out
	var _cret *C.GtkShortcutTrigger    // in

	_arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_alternative_trigger_get_second(_arg0)

	var _shortcutTrigger ShortcutTriggerer // out

	_shortcutTrigger = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(ShortcutTriggerer)

	return _shortcutTrigger
}

// KeyvalTrigger: GtkShortcutTrigger that triggers when a specific keyval and
// modifiers are pressed.
type KeyvalTrigger struct {
	ShortcutTrigger
}

func wrapKeyvalTrigger(obj *externglib.Object) *KeyvalTrigger {
	return &KeyvalTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalKeyvalTriggerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapKeyvalTrigger(obj), nil
}

// NewKeyvalTrigger creates a GtkShortcutTrigger that will trigger whenever the
// key with the given keyval and modifiers is pressed.
func NewKeyvalTrigger(keyval uint, modifiers gdk.ModifierType) *KeyvalTrigger {
	var _arg1 C.guint               // out
	var _arg2 C.GdkModifierType     // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)

	_cret = C.gtk_keyval_trigger_new(_arg1, _arg2)

	var _keyvalTrigger *KeyvalTrigger // out

	_keyvalTrigger = wrapKeyvalTrigger(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _keyvalTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
func (self *KeyvalTrigger) Keyval() uint {
	var _arg0 *C.GtkKeyvalTrigger // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_keyval_trigger_get_keyval(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Modifiers gets the modifiers that must be present to succeed triggering self.
func (self *KeyvalTrigger) Modifiers() gdk.ModifierType {
	var _arg0 *C.GtkKeyvalTrigger // out
	var _cret C.GdkModifierType   // in

	_arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_keyval_trigger_get_modifiers(_arg0)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// MnemonicTrigger: GtkShortcutTrigger that triggers when a specific mnemonic is
// pressed.
//
// Mnemonics require a *mnemonic modifier* (typically <kbd>Alt</kbd>) to be
// pressed together with the mnemonic key.
type MnemonicTrigger struct {
	ShortcutTrigger
}

func wrapMnemonicTrigger(obj *externglib.Object) *MnemonicTrigger {
	return &MnemonicTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalMnemonicTriggerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMnemonicTrigger(obj), nil
}

// NewMnemonicTrigger creates a GtkShortcutTrigger that will trigger whenever
// the key with the given keyval is pressed and mnemonics have been activated.
//
// Mnemonics are activated by calling code when a key event with the right
// modifiers is detected.
func NewMnemonicTrigger(keyval uint) *MnemonicTrigger {
	var _arg1 C.guint               // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = C.guint(keyval)

	_cret = C.gtk_mnemonic_trigger_new(_arg1)

	var _mnemonicTrigger *MnemonicTrigger // out

	_mnemonicTrigger = wrapMnemonicTrigger(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mnemonicTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
func (self *MnemonicTrigger) Keyval() uint {
	var _arg0 *C.GtkMnemonicTrigger // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkMnemonicTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_mnemonic_trigger_get_keyval(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NeverTrigger: GtkShortcutTrigger that never triggers.
type NeverTrigger struct {
	ShortcutTrigger
}

func wrapNeverTrigger(obj *externglib.Object) *NeverTrigger {
	return &NeverTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalNeverTriggerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNeverTrigger(obj), nil
}

func (*NeverTrigger) privateNeverTrigger() {}

// NeverTriggerGet gets the never trigger.
//
// This is a singleton for a trigger that never triggers. Use this trigger
// instead of NULL because it implements all virtual functions.
func NeverTriggerGet() *NeverTrigger {
	var _cret *C.GtkShortcutTrigger // in

	_cret = C.gtk_never_trigger_get()

	var _neverTrigger *NeverTrigger // out

	_neverTrigger = wrapNeverTrigger(externglib.Take(unsafe.Pointer(_cret)))

	return _neverTrigger
}

// ShortcutTrigger: GtkShortcutTrigger tracks how a GtkShortcut should be
// activated.
//
// To find out if a GtkShortcutTrigger triggers, you can call
// gtk.ShortcutTrigger.Trigger() on a GdkEvent.
//
// GtkShortcutTriggers contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutTriggers are immutable, you can only specify their properties
// during construction. If you want to change a trigger, you have to replace it
// with a new one.
type ShortcutTrigger struct {
	*externglib.Object
}

// ShortcutTriggerer describes ShortcutTrigger's abstract methods.
type ShortcutTriggerer interface {
	gextras.Objector

	// Compare types of trigger1 and trigger2 are #gconstpointer only to allow
	// use of this function as a Func.
	Compare(trigger2 ShortcutTriggerer) int
	// Equal checks if trigger1 and trigger2 trigger under the same conditions.
	Equal(trigger2 ShortcutTriggerer) bool
	// Hash generates a hash value for a GtkShortcutTrigger.
	Hash() uint
	// ToLabel gets textual representation for the given trigger.
	ToLabel(display *gdk.Display) string
	// String prints the given trigger into a human-readable string.
	String() string
	// Trigger checks if the given event triggers self.
	Trigger(event gdk.Eventer, enableMnemonics bool) gdk.KeyMatch
}

var _ ShortcutTriggerer = (*ShortcutTrigger)(nil)

func wrapShortcutTrigger(obj *externglib.Object) *ShortcutTrigger {
	return &ShortcutTrigger{
		Object: obj,
	}
}

func marshalShortcutTriggerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutTrigger(obj), nil
}

// NewShortcutTriggerParseString tries to parse the given string into a trigger.
//
// On success, the parsed trigger is returned. When parsing failed, NULL is
// returned.
//
// The accepted strings are:
//
//    - never, for GtkNeverTrigger
//    - a string parsed by gtk_accelerator_parse(), for a GtkKeyvalTrigger, e.g. <Control>C
//    - underscore, followed by a single character, for GtkMnemonicTrigger, e.g. _l
//    - two valid trigger strings, separated by a | character, for a
//      GtkAlternativeTrigger: <Control>q|<Control>w
//
// Note that you will have to escape the < and > characters when specifying
// triggers in XML files, such as GtkBuilder ui files. Use &lt; instead of < and
// &gt; instead of >.
func NewShortcutTriggerParseString(_string string) *ShortcutTrigger {
	var _arg1 *C.char               // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_shortcut_trigger_parse_string(_arg1)

	var _shortcutTrigger *ShortcutTrigger // out

	if _cret != nil {
		_shortcutTrigger = wrapShortcutTrigger(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _shortcutTrigger
}

// Compare types of trigger1 and trigger2 are #gconstpointer only to allow use
// of this function as a Func.
//
// They must each be a GtkShortcutTrigger.
func (trigger1 *ShortcutTrigger) Compare(trigger2 ShortcutTriggerer) int {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.int           // in

	_arg0 = C.gconstpointer(unsafe.Pointer(trigger1.Native()))
	_arg1 = C.gconstpointer(unsafe.Pointer(trigger2.Native()))

	_cret = C.gtk_shortcut_trigger_compare(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Equal checks if trigger1 and trigger2 trigger under the same conditions.
//
// The types of one and two are #gconstpointer only to allow use of this
// function with Table. They must each be a GtkShortcutTrigger.
func (trigger1 *ShortcutTrigger) Equal(trigger2 ShortcutTriggerer) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(unsafe.Pointer(trigger1.Native()))
	_arg1 = C.gconstpointer(unsafe.Pointer(trigger2.Native()))

	_cret = C.gtk_shortcut_trigger_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Hash generates a hash value for a GtkShortcutTrigger.
//
// The output of this function is guaranteed to be the same for a given value
// only per-process. It may change between different processor architectures or
// even different versions of GTK. Do not use this function as a basis for
// building protocols or file formats.
//
// The types of trigger is #gconstpointer only to allow use of this function
// with Table. They must each be a GtkShortcutTrigger.
func (trigger *ShortcutTrigger) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(unsafe.Pointer(trigger.Native()))

	_cret = C.gtk_shortcut_trigger_hash(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ToLabel gets textual representation for the given trigger.
//
// This function is returning a translated string for presentation to end users
// for example in menu items or in help texts.
//
// The display in use may influence the resulting string in various forms, such
// as resolving hardware keycodes or by causing display-specific modifier names.
//
// The form of the representation may change at any time and is not guaranteed
// to stay identical.
func (self *ShortcutTrigger) ToLabel(display *gdk.Display) string {
	var _arg0 *C.GtkShortcutTrigger // out
	var _arg1 *C.GdkDisplay         // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_shortcut_trigger_to_label(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// String prints the given trigger into a human-readable string.
//
// This is a small wrapper around gtk.ShortcutTrigger.Print() to help when
// debugging.
func (self *ShortcutTrigger) String() string {
	var _arg0 *C.GtkShortcutTrigger // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_trigger_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Trigger checks if the given event triggers self.
func (self *ShortcutTrigger) Trigger(event gdk.Eventer, enableMnemonics bool) gdk.KeyMatch {
	var _arg0 *C.GtkShortcutTrigger // out
	var _arg1 *C.GdkEvent           // out
	var _arg2 C.gboolean            // out
	var _cret C.GdkKeyMatch         // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	if enableMnemonics {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_shortcut_trigger_trigger(_arg0, _arg1, _arg2)

	var _keyMatch gdk.KeyMatch // out

	_keyMatch = gdk.KeyMatch(_cret)

	return _keyMatch
}
