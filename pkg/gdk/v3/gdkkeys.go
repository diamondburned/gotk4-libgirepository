// Code generated by girgen. DO NOT EDIT.

package gdk

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
// extern void _gotk4_gdk3_Keymap_ConnectStateChanged(gpointer, guintptr);
// extern void _gotk4_gdk3_Keymap_ConnectKeysChanged(gpointer, guintptr);
// extern void _gotk4_gdk3_Keymap_ConnectDirectionChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeKeymap = coreglib.Type(girepository.MustFind("Gdk", "Keymap").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeKeymap, F: marshalKeymap},
	})
}

// Keymap defines the translation from keyboard state (including a hardware key,
// a modifier mask, and active keyboard group) to a keyval. This translation has
// two phases. The first phase is to determine the effective keyboard group and
// level for the keyboard state; the second phase is to look up the
// keycode/group/level triplet in the keymap and see what keyval it corresponds
// to.
type Keymap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Keymap)(nil)
)

func wrapKeymap(obj *coreglib.Object) *Keymap {
	return &Keymap{
		Object: obj,
	}
}

func marshalKeymap(p uintptr) (interface{}, error) {
	return wrapKeymap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDirectionChanged signal gets emitted when the direction of the keymap
// changes.
func (v *Keymap) ConnectDirectionChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "direction-changed", false, unsafe.Pointer(C._gotk4_gdk3_Keymap_ConnectDirectionChanged), f)
}

// ConnectKeysChanged signal is emitted when the mapping represented by keymap
// changes.
func (v *Keymap) ConnectKeysChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "keys-changed", false, unsafe.Pointer(C._gotk4_gdk3_Keymap_ConnectKeysChanged), f)
}

// ConnectStateChanged signal is emitted when the state of the keyboard changes,
// e.g when Caps Lock is turned on or off. See gdk_keymap_get_caps_lock_state().
func (v *Keymap) ConnectStateChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "state-changed", false, unsafe.Pointer(C._gotk4_gdk3_Keymap_ConnectStateChanged), f)
}

// KeymapKey is a hardware key that can be mapped to a keyval.
//
// An instance of this type is always passed by reference.
type KeymapKey struct {
	*keymapKey
}

// keymapKey is the struct that's finalized.
type keymapKey struct {
	native unsafe.Pointer
}

var GIRInfoKeymapKey = girepository.MustFind("Gdk", "KeymapKey")

// NewKeymapKey creates a new KeymapKey instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewKeymapKey(keycode uint, group, level int) KeymapKey {
	var f0 C.guint // out
	f0 = C.guint(keycode)
	var f1 C.gint // out
	f1 = C.gint(group)
	var f2 C.gint // out
	f2 = C.gint(level)

	size := GIRInfoKeymapKey.StructSize()
	native := make([]byte, size)
	gextras.Sink(&native[0])

	offset0 := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr0 := (*C.guint)(unsafe.Add(unsafe.Pointer(&native[0]), offset0))
	*valptr0 = f0

	offset1 := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr1 := (*C.gint)(unsafe.Add(unsafe.Pointer(&native[0]), offset1))
	*valptr1 = f1

	offset2 := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr2 := (*C.gint)(unsafe.Add(unsafe.Pointer(&native[0]), offset2))
	*valptr2 = f2

	return *(*KeymapKey)(gextras.NewStructNative(unsafe.Pointer(&native[0])))
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) Keycode() uint {
	offset := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr := (*uint)(unsafe.Add(k.native, offset))
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	offset := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr := (*int)(unsafe.Add(k.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	offset := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr := (*int)(unsafe.Add(k.native, offset))
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) SetKeycode(keycode uint) {
	offset := GIRInfoKeymapKey.StructFieldOffset("keycode")
	valptr := (*C.guint)(unsafe.Add(k.native, offset))
	*valptr = C.guint(keycode)
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) SetGroup(group int) {
	offset := GIRInfoKeymapKey.StructFieldOffset("group")
	valptr := (*C.gint)(unsafe.Add(k.native, offset))
	*valptr = C.gint(group)
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) SetLevel(level int) {
	offset := GIRInfoKeymapKey.StructFieldOffset("level")
	valptr := (*C.gint)(unsafe.Add(k.native, offset))
	*valptr = C.gint(level)
}
