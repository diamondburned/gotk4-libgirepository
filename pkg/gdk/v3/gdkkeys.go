// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_keymap_get_type()), F: marshalKeymapper},
	})
}

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// symbol. Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var _arg1 C.guint // out
	var _arg2 C.guint // in
	var _arg3 C.guint // in

	_arg1 = C.guint(symbol)

	C.gdk_keyval_convert_case(_arg1, &_arg2, &_arg3)

	var _lower uint // out
	var _upper uint // out

	_lower = uint(_arg2)
	_upper = uint(_arg3)

	return _lower, _upper
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the gdk/gdkkeysyms.h header file but
// without the leading “GDK_KEY_”.
func KeyvalFromName(keyvalName string) uint {
	var _arg1 *C.gchar // out
	var _cret C.guint  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(keyvalName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_keyval_from_name(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalIsLower returns TRUE if the given key value is in lower case.
func KeyvalIsLower(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_lower(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalIsUpper returns TRUE if the given key value is in upper case.
func KeyvalIsUpper(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_upper(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the gdk/gdkkeysyms.h header file but
// without the leading “GDK_KEY_”.
func KeyvalName(keyval uint) string {
	var _arg1 C.guint  // out
	var _cret *C.gchar // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_name(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// KeyvalToLower converts a key value to lower case, if applicable.
func KeyvalToLower(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_lower(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding ISO10646
// (Unicode) character.
func KeyvalToUnicode(keyval uint) uint32 {
	var _arg1 C.guint   // out
	var _cret C.guint32 // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_unicode(_arg1)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// KeyvalToUpper converts a key value to upper case, if applicable.
func KeyvalToUpper(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_upper(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// UnicodeToKeyval: convert from a ISO10646 character to a key symbol.
func UnicodeToKeyval(wc uint32) uint {
	var _arg1 C.guint32 // out
	var _cret C.guint   // in

	_arg1 = C.guint32(wc)

	_cret = C.gdk_unicode_to_keyval(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Keymap defines the translation from keyboard state (including a hardware key,
// a modifier mask, and active keyboard group) to a keyval. This translation has
// two phases. The first phase is to determine the effective keyboard group and
// level for the keyboard state; the second phase is to look up the
// keycode/group/level triplet in the keymap and see what keyval it corresponds
// to.
type Keymap struct {
	*externglib.Object
}

func wrapKeymap(obj *externglib.Object) *Keymap {
	return &Keymap{
		Object: obj,
	}
}

func marshalKeymapper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapKeymap(obj), nil
}

// CapsLockState returns whether the Caps Lock modifer is locked.
func (keymap *Keymap) CapsLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_get_caps_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Direction returns the direction of effective layout of the keymap.
func (keymap *Keymap) Direction() pango.Direction {
	var _arg0 *C.GdkKeymap     // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_get_direction(_arg0)

	var _direction pango.Direction // out

	_direction = pango.Direction(_cret)

	return _direction
}

// EntriesForKeycode returns the keyvals bound to hardware_keycode. The Nth
// KeymapKey in keys is bound to the Nth keyval in keyvals. Free the returned
// arrays with g_free(). When a keycode is pressed by the user, the keyval from
// this list of entries is selected by considering the effective keyboard group
// and level. See gdk_keymap_translate_keyboard_state().
func (keymap *Keymap) EntriesForKeycode(hardwareKeycode uint) ([]KeymapKey, []uint, bool) {
	var _arg0 *C.GdkKeymap    // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg4 C.gint          // in
	var _arg3 *C.guint        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(hardwareKeycode)

	_cret = C.gdk_keymap_get_entries_for_keycode(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _keys []KeymapKey // out
	var _keyvals []uint   // out
	var _ok bool          // out

	defer C.free(unsafe.Pointer(_arg2))
	_keys = make([]KeymapKey, _arg4)
	copy(_keys, unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg4))
	defer C.free(unsafe.Pointer(_arg3))
	_keyvals = make([]uint, _arg4)
	copy(_keyvals, unsafe.Slice((*uint)(unsafe.Pointer(_arg3)), _arg4))
	if _cret != 0 {
		_ok = true
	}

	return _keys, _keyvals, _ok
}

// EntriesForKeyval obtains a list of keycode/group/level combinations that will
// generate keyval. Groups and levels are two kinds of keyboard mode; in
// general, the level determines whether the top or bottom symbol on a key is
// used, and the group determines whether the left or right symbol is used. On
// US keyboards, the shift key changes the keyboard level, and there are no
// groups. A group switch key might convert a keyboard between Hebrew to English
// modes, for example. EventKey contains a group field that indicates the active
// keyboard group. The level is computed from the modifier mask. The returned
// array should be freed with g_free().
func (keymap *Keymap) EntriesForKeyval(keyval uint) ([]KeymapKey, bool) {
	var _arg0 *C.GdkKeymap    // out
	var _arg1 C.guint         // out
	var _arg2 *C.GdkKeymapKey // in
	var _arg3 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(keyval)

	_cret = C.gdk_keymap_get_entries_for_keyval(_arg0, _arg1, &_arg2, &_arg3)

	var _keys []KeymapKey // out
	var _ok bool          // out

	defer C.free(unsafe.Pointer(_arg2))
	_keys = make([]KeymapKey, _arg3)
	copy(_keys, unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg3))
	if _cret != 0 {
		_ok = true
	}

	return _keys, _ok
}

// ModifierMask returns the modifier mask the keymap’s windowing system backend
// uses for a particular purpose.
//
// Note that this function always returns real hardware modifiers, not virtual
// ones (e.g. it will return K_MOD1_MASK rather than K_META_MASK if the backend
// maps MOD1 to META), so there are use cases where the return value of this
// function has to be transformed by gdk_keymap_add_virtual_modifiers() in order
// to contain the expected result.
func (keymap *Keymap) ModifierMask(intent ModifierIntent) ModifierType {
	var _arg0 *C.GdkKeymap        // out
	var _arg1 C.GdkModifierIntent // out
	var _cret C.GdkModifierType   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.GdkModifierIntent(intent)

	_cret = C.gdk_keymap_get_modifier_mask(_arg0, _arg1)

	var _modifierType ModifierType // out

	_modifierType = ModifierType(_cret)

	return _modifierType
}

// ModifierState returns the current modifier state.
func (keymap *Keymap) ModifierState() uint {
	var _arg0 *C.GdkKeymap // out
	var _cret C.guint      // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_get_modifier_state(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NumLockState returns whether the Num Lock modifer is locked.
func (keymap *Keymap) NumLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_get_num_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScrollLockState returns whether the Scroll Lock modifer is locked.
func (keymap *Keymap) ScrollLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_get_scroll_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HaveBidiLayouts determines if keyboard layouts for both right-to-left and
// left-to-right languages are in use.
func (keymap *Keymap) HaveBidiLayouts() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))

	_cret = C.gdk_keymap_have_bidi_layouts(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupKey looks up the keyval mapped to a keycode/group/level triplet. If no
// keyval is bound to key, returns 0. For normal user input, you want to use
// gdk_keymap_translate_keyboard_state() instead of this function, since the
// effective group/level may not be the same as the current keyboard state.
func (keymap *Keymap) LookupKey(key *KeymapKey) uint {
	var _arg0 *C.GdkKeymap    // out
	var _arg1 *C.GdkKeymapKey // out
	var _cret C.guint         // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = (*C.GdkKeymapKey)(gextras.StructNative(unsafe.Pointer(key)))

	_cret = C.gdk_keymap_lookup_key(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TranslateKeyboardState translates the contents of a EventKey into a keyval,
// effective group, and level. Modifiers that affected the translation and are
// thus unavailable for application use are returned in consumed_modifiers. See
// [Groups][key-group-explanation] for an explanation of groups and levels. The
// effective_group is the group that was actually used for the translation; some
// keys such as Enter are not affected by the active keyboard group. The level
// is derived from state. For convenience, EventKey already contains the
// translated keyval, so this function isn’t as useful as you might think.
//
// consumed_modifiers gives modifiers that should be masked outfrom state when
// comparing this key press to a hot key. For instance, on a US keyboard, the
// plus symbol is shifted, so when comparing a key press to a <Control>plus
// accelerator <Shift> should be masked out.
//
//    // XXX Don’t do this XXX
//    if (keyval == accel_keyval &&
//        (event->state & ~consumed & ALL_ACCELS_MASK) == (accel_mods & ~consumed))
//      // Accelerator was pressed
//
// However, this did not work if multi-modifier combinations were used in the
// keymap, since, for instance, <Control> would be masked out even if only
// <Control><Alt> was used in the keymap. To support this usage as well as well
// as possible, all single modifier combinations that could affect the key for
// any combination of modifiers will be returned in consumed_modifiers;
// multi-modifier combinations are returned only when actually found in state.
// When you store accelerators, you should always store them with consumed
// modifiers removed. Store <Control>plus, not <Control><Shift>plus,
func (keymap *Keymap) TranslateKeyboardState(hardwareKeycode uint, state ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumedModifiers ModifierType, ok bool) {
	var _arg0 *C.GdkKeymap      // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 C.gint            // out
	var _arg4 C.guint           // in
	var _arg5 C.gint            // in
	var _arg6 C.gint            // in
	var _arg7 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(hardwareKeycode)
	_arg2 = C.GdkModifierType(state)
	_arg3 = C.gint(group)

	_cret = C.gdk_keymap_translate_keyboard_state(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)

	var _keyval uint                    // out
	var _effectiveGroup int             // out
	var _level int                      // out
	var _consumedModifiers ModifierType // out
	var _ok bool                        // out

	_keyval = uint(_arg4)
	_effectiveGroup = int(_arg5)
	_level = int(_arg6)
	_consumedModifiers = ModifierType(_arg7)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _effectiveGroup, _level, _consumedModifiers, _ok
}

// KeymapGetDefault returns the Keymap attached to the default display.
//
// Deprecated: Use gdk_keymap_get_for_display() instead.
func KeymapGetDefault() *Keymap {
	var _cret *C.GdkKeymap // in

	_cret = C.gdk_keymap_get_default()

	var _keymap *Keymap // out

	_keymap = wrapKeymap(externglib.Take(unsafe.Pointer(_cret)))

	return _keymap
}

// KeymapGetForDisplay returns the Keymap attached to display.
func KeymapGetForDisplay(display *Display) *Keymap {
	var _arg1 *C.GdkDisplay // out
	var _cret *C.GdkKeymap  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_keymap_get_for_display(_arg1)

	var _keymap *Keymap // out

	_keymap = wrapKeymap(externglib.Take(unsafe.Pointer(_cret)))

	return _keymap
}

// KeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	nocopy gextras.NoCopy
	native *C.GdkKeymapKey
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) Keycode() uint {
	var v uint // out
	v = uint(k.native.keycode)
	return v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	var v int // out
	v = int(k.native.group)
	return v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	var v int // out
	v = int(k.native.level)
	return v
}
