// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeAlternativeTrigger = coreglib.Type(C.gtk_alternative_trigger_get_type())
	GTypeKeyvalTrigger      = coreglib.Type(C.gtk_keyval_trigger_get_type())
	GTypeMnemonicTrigger    = coreglib.Type(C.gtk_mnemonic_trigger_get_type())
	GTypeNeverTrigger       = coreglib.Type(C.gtk_never_trigger_get_type())
	GTypeShortcutTrigger    = coreglib.Type(C.gtk_shortcut_trigger_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAlternativeTrigger, F: marshalAlternativeTrigger},
		coreglib.TypeMarshaler{T: GTypeKeyvalTrigger, F: marshalKeyvalTrigger},
		coreglib.TypeMarshaler{T: GTypeMnemonicTrigger, F: marshalMnemonicTrigger},
		coreglib.TypeMarshaler{T: GTypeNeverTrigger, F: marshalNeverTrigger},
		coreglib.TypeMarshaler{T: GTypeShortcutTrigger, F: marshalShortcutTrigger},
	})
}

// AlternativeTriggerOverrider contains methods that are overridable.
type AlternativeTriggerOverrider interface {
}

// AlternativeTrigger: GtkShortcutTrigger that combines two triggers.
//
// The GtkAlternativeTrigger triggers when either of two trigger.
//
// This can be cascaded to combine more than two triggers.
type AlternativeTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*AlternativeTrigger)(nil)
)

func initClassAlternativeTrigger(gclass unsafe.Pointer, goval any) {
}

func wrapAlternativeTrigger(obj *coreglib.Object) *AlternativeTrigger {
	return &AlternativeTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalAlternativeTrigger(p uintptr) (interface{}, error) {
	return wrapAlternativeTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAlternativeTrigger creates a GtkShortcutTrigger that will trigger whenever
// either of the two given triggers gets triggered.
//
// Note that nesting is allowed, so if you want more than two alternative,
// create a new alternative trigger for each option.
//
// The function takes the following parameters:
//
//    - first trigger that may trigger.
//    - second trigger that may trigger.
//
// The function returns the following values:
//
//    - alternativeTrigger: new GtkShortcutTrigger.
//
func NewAlternativeTrigger(first, second ShortcutTriggerer) *AlternativeTrigger {
	var _arg1 *C.GtkShortcutTrigger // out
	var _arg2 *C.GtkShortcutTrigger // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(coreglib.InternObject(first).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(first).Native()))
	_arg2 = (*C.GtkShortcutTrigger)(unsafe.Pointer(coreglib.InternObject(second).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(second).Native()))

	_cret = C.gtk_alternative_trigger_new(_arg1, _arg2)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _alternativeTrigger *AlternativeTrigger // out

	_alternativeTrigger = wrapAlternativeTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _alternativeTrigger
}

// First gets the first of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetSecond() will return the other one.
//
// The function returns the following values:
//
//    - shortcutTrigger: first alternative trigger.
//
func (self *AlternativeTrigger) First() ShortcutTriggerer {
	var _arg0 *C.GtkAlternativeTrigger // out
	var _cret *C.GtkShortcutTrigger    // in

	_arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_alternative_trigger_get_first(_arg0)
	runtime.KeepAlive(self)

	var _shortcutTrigger ShortcutTriggerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.ShortcutTriggerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(ShortcutTriggerer)
			return ok
		})
		rv, ok := casted.(ShortcutTriggerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutTriggerer")
		}
		_shortcutTrigger = rv
	}

	return _shortcutTrigger
}

// Second gets the second of the two alternative triggers that may trigger self.
//
// gtk.AlternativeTrigger.GetFirst() will return the other one.
//
// The function returns the following values:
//
//    - shortcutTrigger: second alternative trigger.
//
func (self *AlternativeTrigger) Second() ShortcutTriggerer {
	var _arg0 *C.GtkAlternativeTrigger // out
	var _cret *C.GtkShortcutTrigger    // in

	_arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_alternative_trigger_get_second(_arg0)
	runtime.KeepAlive(self)

	var _shortcutTrigger ShortcutTriggerer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.ShortcutTriggerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(ShortcutTriggerer)
			return ok
		})
		rv, ok := casted.(ShortcutTriggerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutTriggerer")
		}
		_shortcutTrigger = rv
	}

	return _shortcutTrigger
}

// KeyvalTriggerOverrider contains methods that are overridable.
type KeyvalTriggerOverrider interface {
}

// KeyvalTrigger: GtkShortcutTrigger that triggers when a specific keyval and
// modifiers are pressed.
type KeyvalTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*KeyvalTrigger)(nil)
)

func initClassKeyvalTrigger(gclass unsafe.Pointer, goval any) {
}

func wrapKeyvalTrigger(obj *coreglib.Object) *KeyvalTrigger {
	return &KeyvalTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalKeyvalTrigger(p uintptr) (interface{}, error) {
	return wrapKeyvalTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewKeyvalTrigger creates a GtkShortcutTrigger that will trigger whenever the
// key with the given keyval and modifiers is pressed.
//
// The function takes the following parameters:
//
//    - keyval to trigger for.
//    - modifiers that need to be present.
//
// The function returns the following values:
//
//    - keyvalTrigger: new GtkShortcutTrigger.
//
func NewKeyvalTrigger(keyval uint, modifiers gdk.ModifierType) *KeyvalTrigger {
	var _arg1 C.guint               // out
	var _arg2 C.GdkModifierType     // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)

	_cret = C.gtk_keyval_trigger_new(_arg1, _arg2)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)

	var _keyvalTrigger *KeyvalTrigger // out

	_keyvalTrigger = wrapKeyvalTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _keyvalTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
//
// The function returns the following values:
//
//    - guint: keyval.
//
func (self *KeyvalTrigger) Keyval() uint {
	var _arg0 *C.GtkKeyvalTrigger // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_keyval_trigger_get_keyval(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Modifiers gets the modifiers that must be present to succeed triggering self.
//
// The function returns the following values:
//
//    - modifierType: modifiers.
//
func (self *KeyvalTrigger) Modifiers() gdk.ModifierType {
	var _arg0 *C.GtkKeyvalTrigger // out
	var _cret C.GdkModifierType   // in

	_arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_keyval_trigger_get_modifiers(_arg0)
	runtime.KeepAlive(self)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// MnemonicTriggerOverrider contains methods that are overridable.
type MnemonicTriggerOverrider interface {
}

// MnemonicTrigger: GtkShortcutTrigger that triggers when a specific mnemonic is
// pressed.
//
// Mnemonics require a *mnemonic modifier* (typically <kbd>Alt</kbd>) to be
// pressed together with the mnemonic key.
type MnemonicTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*MnemonicTrigger)(nil)
)

func initClassMnemonicTrigger(gclass unsafe.Pointer, goval any) {
}

func wrapMnemonicTrigger(obj *coreglib.Object) *MnemonicTrigger {
	return &MnemonicTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalMnemonicTrigger(p uintptr) (interface{}, error) {
	return wrapMnemonicTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMnemonicTrigger creates a GtkShortcutTrigger that will trigger whenever
// the key with the given keyval is pressed and mnemonics have been activated.
//
// Mnemonics are activated by calling code when a key event with the right
// modifiers is detected.
//
// The function takes the following parameters:
//
//    - keyval to trigger for.
//
// The function returns the following values:
//
//    - mnemonicTrigger: new GtkShortcutTrigger.
//
func NewMnemonicTrigger(keyval uint) *MnemonicTrigger {
	var _arg1 C.guint               // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = C.guint(keyval)

	_cret = C.gtk_mnemonic_trigger_new(_arg1)
	runtime.KeepAlive(keyval)

	var _mnemonicTrigger *MnemonicTrigger // out

	_mnemonicTrigger = wrapMnemonicTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mnemonicTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering self.
//
// The function returns the following values:
//
//    - guint: keyval.
//
func (self *MnemonicTrigger) Keyval() uint {
	var _arg0 *C.GtkMnemonicTrigger // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkMnemonicTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_mnemonic_trigger_get_keyval(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NeverTriggerOverrider contains methods that are overridable.
type NeverTriggerOverrider interface {
}

// NeverTrigger: GtkShortcutTrigger that never triggers.
type NeverTrigger struct {
	_ [0]func() // equal guard
	ShortcutTrigger
}

var (
	_ ShortcutTriggerer = (*NeverTrigger)(nil)
)

func initClassNeverTrigger(gclass unsafe.Pointer, goval any) {
}

func wrapNeverTrigger(obj *coreglib.Object) *NeverTrigger {
	return &NeverTrigger{
		ShortcutTrigger: ShortcutTrigger{
			Object: obj,
		},
	}
}

func marshalNeverTrigger(p uintptr) (interface{}, error) {
	return wrapNeverTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NeverTriggerGet gets the never trigger.
//
// This is a singleton for a trigger that never triggers. Use this trigger
// instead of NULL because it implements all virtual functions.
//
// The function returns the following values:
//
//    - neverTrigger: never trigger.
//
func NeverTriggerGet() *NeverTrigger {
	var _cret *C.GtkShortcutTrigger // in

	_cret = C.gtk_never_trigger_get()

	var _neverTrigger *NeverTrigger // out

	_neverTrigger = wrapNeverTrigger(coreglib.Take(unsafe.Pointer(_cret)))

	return _neverTrigger
}

// ShortcutTriggerOverrider contains methods that are overridable.
type ShortcutTriggerOverrider interface {
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutTrigger)(nil)
)

// ShortcutTriggerer describes types inherited from class ShortcutTrigger.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ShortcutTriggerer interface {
	coreglib.Objector
	baseShortcutTrigger() *ShortcutTrigger
}

var _ ShortcutTriggerer = (*ShortcutTrigger)(nil)

func initClassShortcutTrigger(gclass unsafe.Pointer, goval any) {
}

func wrapShortcutTrigger(obj *coreglib.Object) *ShortcutTrigger {
	return &ShortcutTrigger{
		Object: obj,
	}
}

func marshalShortcutTrigger(p uintptr) (interface{}, error) {
	return wrapShortcutTrigger(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (trigger1 *ShortcutTrigger) baseShortcutTrigger() *ShortcutTrigger {
	return trigger1
}

// BaseShortcutTrigger returns the underlying base object.
func BaseShortcutTrigger(obj ShortcutTriggerer) *ShortcutTrigger {
	return obj.baseShortcutTrigger()
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
//
// The function takes the following parameters:
//
//    - str: string to parse.
//
// The function returns the following values:
//
//    - shortcutTrigger (optional): new GtkShortcutTrigger or NULL on error.
//
func NewShortcutTriggerParseString(str string) *ShortcutTrigger {
	var _arg1 *C.char               // out
	var _cret *C.GtkShortcutTrigger // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_shortcut_trigger_parse_string(_arg1)
	runtime.KeepAlive(str)

	var _shortcutTrigger *ShortcutTrigger // out

	if _cret != nil {
		_shortcutTrigger = wrapShortcutTrigger(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _shortcutTrigger
}

// Compare types of trigger1 and trigger2 are #gconstpointer only to allow use
// of this function as a Func.
//
// They must each be a GtkShortcutTrigger.
//
// The function takes the following parameters:
//
//    - trigger2: GtkShortcutTrigger.
//
// The function returns the following values:
//
//    - gint: integer less than, equal to, or greater than zero if trigger1 is
//      found, respectively, to be less than, to match, or be greater than
//      trigger2.
//
func (trigger1 *ShortcutTrigger) Compare(trigger2 ShortcutTriggerer) int {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.int           // in

	_arg0 = *(*C.gconstpointer)(unsafe.Pointer(coreglib.InternObject(trigger1).Native()))
	_arg1 = *(*C.gconstpointer)(unsafe.Pointer(coreglib.InternObject(trigger2).Native()))

	_cret = C.gtk_shortcut_trigger_compare(_arg0, _arg1)
	runtime.KeepAlive(trigger1)
	runtime.KeepAlive(trigger2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Equal checks if trigger1 and trigger2 trigger under the same conditions.
//
// The types of one and two are #gconstpointer only to allow use of this
// function with Table. They must each be a GtkShortcutTrigger.
//
// The function takes the following parameters:
//
//    - trigger2: GtkShortcutTrigger.
//
// The function returns the following values:
//
//    - ok: TRUE if trigger1 and trigger2 are equal.
//
func (trigger1 *ShortcutTrigger) Equal(trigger2 ShortcutTriggerer) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = *(*C.gconstpointer)(unsafe.Pointer(coreglib.InternObject(trigger1).Native()))
	_arg1 = *(*C.gconstpointer)(unsafe.Pointer(coreglib.InternObject(trigger2).Native()))

	_cret = C.gtk_shortcut_trigger_equal(_arg0, _arg1)
	runtime.KeepAlive(trigger1)
	runtime.KeepAlive(trigger2)

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
//
// The function returns the following values:
//
//    - guint: hash value corresponding to trigger.
//
func (trigger *ShortcutTrigger) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = *(*C.gconstpointer)(unsafe.Pointer(coreglib.InternObject(trigger).Native()))

	_cret = C.gtk_shortcut_trigger_hash(_arg0)
	runtime.KeepAlive(trigger)

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
//
// The function takes the following parameters:
//
//    - display: GdkDisplay to print for.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (self *ShortcutTrigger) ToLabel(display *gdk.Display) string {
	var _arg0 *C.GtkShortcutTrigger // out
	var _arg1 *C.GdkDisplay         // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gtk_shortcut_trigger_to_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// String prints the given trigger into a human-readable string.
//
// This is a small wrapper around gtk.ShortcutTrigger.Print() to help when
// debugging.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (self *ShortcutTrigger) String() string {
	var _arg0 *C.GtkShortcutTrigger // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_shortcut_trigger_to_string(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Trigger checks if the given event triggers self.
//
// The function takes the following parameters:
//
//    - event to check.
//    - enableMnemonics: TRUE if mnemonics should trigger. Usually the value of
//      this property is determined by checking that the passed in event is a Key
//      event and has the right modifiers set.
//
// The function returns the following values:
//
//    - keyMatch: whether the event triggered the shortcut.
//
func (self *ShortcutTrigger) Trigger(event gdk.Eventer, enableMnemonics bool) gdk.KeyMatch {
	var _arg0 *C.GtkShortcutTrigger // out
	var _arg1 *C.GdkEvent           // out
	var _arg2 C.gboolean            // out
	var _cret C.GdkKeyMatch         // in

	_arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(coreglib.InternObject(event).Native()))
	if enableMnemonics {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_shortcut_trigger_trigger(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(event)
	runtime.KeepAlive(enableMnemonics)

	var _keyMatch gdk.KeyMatch // out

	_keyMatch = gdk.KeyMatch(_cret)

	return _keyMatch
}
