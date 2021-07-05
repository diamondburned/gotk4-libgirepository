// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_accel_flags_get_type()), F: marshalAccelFlags},
		{T: externglib.Type(C.gtk_accel_group_get_type()), F: marshalAccelGroup},
	})
}

// AccelFlags: accelerator flags used with gtk_accel_group_connect().
type AccelFlags int

const (
	// AccelFlagsVisible: accelerator is visible
	AccelFlagsVisible AccelFlags = 0b1
	// AccelFlagsLocked: accelerator not removable
	AccelFlagsLocked AccelFlags = 0b10
	// AccelFlagsMask: mask
	AccelFlagsMask AccelFlags = 0b111
)

func marshalAccelFlags(p uintptr) (interface{}, error) {
	return AccelFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AccelGroupsActivate finds the first accelerator in any AccelGroup attached to
// @object that matches @accel_key and @accel_mods, and activates that
// accelerator.
func AccelGroupsActivate(object gextras.Objector, accelKey uint, accelMods gdk.ModifierType) bool {
	var _arg1 *C.GObject        // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_groups_activate(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AcceleratorGetDefaultModMask gets the modifier mask.
//
// The modifier mask determines which modifiers are considered significant for
// keyboard accelerators. See gtk_accelerator_set_default_mod_mask().
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	var _cret C.GdkModifierType // in

	_cret = C.gtk_accelerator_get_default_mod_mask()

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// AcceleratorGetLabel converts an accelerator keyval and modifier mask into a
// string which can be used to represent the accelerator to the user.
func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = C.guint(acceleratorKey)
	_arg2 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_get_label(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorGetLabelWithKeycode converts an accelerator keyval and modifier
// mask into a (possibly translated) string that can be displayed to a user,
// similarly to gtk_accelerator_get_label(), but handling keycodes.
//
// This is only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
func AcceleratorGetLabelWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_get_label_with_keycode(_arg1, _arg2, _arg3, _arg4)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorName converts an accelerator keyval and modifier mask into a
// string parseable by gtk_accelerator_parse(). For example, if you pass in
// K_KEY_q and K_CONTROL_MASK, this function returns “<Control>q”.
//
// If you need to display accelerators in the user interface, see
// gtk_accelerator_get_label().
func AcceleratorName(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = C.guint(acceleratorKey)
	_arg2 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_name(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorNameWithKeycode converts an accelerator keyval and modifier mask
// into a string parseable by gtk_accelerator_parse_with_keycode(), similarly to
// gtk_accelerator_name() but handling keycodes. This is only useful for
// system-level components, applications should use gtk_accelerator_parse()
// instead.
func AcceleratorNameWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.gchar          // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = C.GdkModifierType(acceleratorMods)

	_cret = C.gtk_accelerator_name_with_keycode(_arg1, _arg2, _arg3, _arg4)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorParse parses a string representing an accelerator. The format
// looks like “<Control>a” or “<Shift><Alt>F1” or “<Release>z” (the last one is
// for key release).
//
// The parser is fairly liberal and allows lower or upper case, and also
// abbreviations such as “<Ctl>” and “<Ctrl>”. Key names are parsed using
// gdk_keyval_from_name(). For character keys the name is not the symbol, but
// the lowercase name, e.g. one would use “<Ctrl>minus” instead of “<Ctrl>-”.
//
// If the parse fails, @accelerator_key and @accelerator_mods will be set to 0
// (zero).
func AcceleratorParse(accelerator string) (uint, gdk.ModifierType) {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in

	_arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accelerator_parse(_arg1, &_arg2, &_arg3)

	var _acceleratorKey uint              // out
	var _acceleratorMods gdk.ModifierType // out

	_acceleratorKey = uint(_arg2)
	_acceleratorMods = gdk.ModifierType(_arg3)

	return _acceleratorKey, _acceleratorMods
}

// AcceleratorParseWithKeycode parses a string representing an accelerator,
// similarly to gtk_accelerator_parse() but handles keycodes as well. This is
// only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
//
// If @accelerator_codes is given and the result stored in it is non-nil, the
// result must be freed with g_free().
//
// If a keycode is present in the accelerator and no @accelerator_codes is
// given, the parse will fail.
//
// If the parse fails, @accelerator_key, @accelerator_mods and
// @accelerator_codes will be set to 0 (zero).
func AcceleratorParseWithKeycode(accelerator string) (uint, []uint, gdk.ModifierType) {
	var _arg1 *C.gchar // out
	var _arg2 C.guint  // in
	var _arg3 *C.guint
	var _arg4 C.GdkModifierType // in

	_arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accelerator_parse_with_keycode(_arg1, &_arg2, &_arg3, &_arg4)

	var _acceleratorKey uint // out
	var _acceleratorCodes []uint
	var _acceleratorMods gdk.ModifierType // out

	_acceleratorKey = uint(_arg2)
	{
		var i int
		var z C.guint
		for p := _arg3; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg3, i)
		_acceleratorCodes = make([]uint, i)
		for i := range src {
			_acceleratorCodes[i] = uint(src[i])
		}
	}
	_acceleratorMods = gdk.ModifierType(_arg4)

	return _acceleratorKey, _acceleratorCodes, _acceleratorMods
}

// AcceleratorSetDefaultModMask sets the modifiers that will be considered
// significant for keyboard accelerators. The default mod mask depends on the
// GDK backend in use, but will typically include K_CONTROL_MASK | K_SHIFT_MASK
// | K_MOD1_MASK | K_SUPER_MASK | K_HYPER_MASK | K_META_MASK. In other words,
// Control, Shift, Alt, Super, Hyper and Meta. Other modifiers will by default
// be ignored by AccelGroup.
//
// You must include at least the three modifiers Control, Shift and Alt in any
// value you pass to this function.
//
// The default mod mask should be changed on application startup, before using
// any accelerator groups.
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	var _arg1 C.GdkModifierType // out

	_arg1 = C.GdkModifierType(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(_arg1)
}

// AcceleratorValid determines whether a given keyval and modifier mask
// constitute a valid keyboard accelerator. For example, the K_KEY_a keyval plus
// K_CONTROL_MASK is valid - this is a “Ctrl+a” accelerator. But, you can't, for
// instance, use the K_KEY_Control_L keyval as an accelerator.
func AcceleratorValid(keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)

	_cret = C.gtk_accelerator_valid(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AccelGroup represents a group of keyboard accelerators, typically attached to
// a toplevel Window (with gtk_window_add_accel_group()). Usually you won’t need
// to create a AccelGroup directly; instead, when using UIManager, GTK+
// automatically sets up the accelerators for your menus in the ui manager’s
// AccelGroup.
//
// Note that “accelerators” are different from “mnemonics”. Accelerators are
// shortcuts for activating a menu item; they appear alongside the menu item
// they’re a shortcut for. For example “Ctrl+Q” might appear alongside the
// “Quit” menu item. Mnemonics are shortcuts for GUI elements such as text
// entries or buttons; they appear as underlined characters. See
// gtk_label_new_with_mnemonic(). Menu items can have both accelerators and
// mnemonics, of course.
type AccelGroup interface {
	gextras.Objector

	// DisconnectKeyAccelGroup removes an accelerator previously installed
	// through gtk_accel_group_connect().
	DisconnectKeyAccelGroup(accelKey uint, accelMods gdk.ModifierType) bool
	// IsLocked locks are added and removed using gtk_accel_group_lock() and
	// gtk_accel_group_unlock().
	IsLocked() bool
	// ModifierMask gets a ModifierType representing the mask for this
	// @accel_group. For example, K_CONTROL_MASK, K_SHIFT_MASK, etc.
	ModifierMask() gdk.ModifierType
	// LockAccelGroup locks the given accelerator group.
	//
	// Locking an acelerator group prevents the accelerators contained within it
	// to be changed during runtime. Refer to gtk_accel_map_change_entry() about
	// runtime accelerator changes.
	//
	// If called more than once, @accel_group remains locked until
	// gtk_accel_group_unlock() has been called an equivalent number of times.
	LockAccelGroup()
	// UnlockAccelGroup undoes the last call to gtk_accel_group_lock() on this
	// @accel_group.
	UnlockAccelGroup()
}

// accelGroup implements the AccelGroup class.
type accelGroup struct {
	gextras.Objector
}

// WrapAccelGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapAccelGroup(obj *externglib.Object) AccelGroup {
	return accelGroup{
		Objector: obj,
	}
}

func marshalAccelGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAccelGroup(obj), nil
}

// NewAccelGroup creates a new AccelGroup.
func NewAccelGroup() AccelGroup {
	var _cret *C.GtkAccelGroup // in

	_cret = C.gtk_accel_group_new()

	var _accelGroup AccelGroup // out

	_accelGroup = WrapAccelGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _accelGroup
}

func (a accelGroup) DisconnectKeyAccelGroup(accelKey uint, accelMods gdk.ModifierType) bool {
	var _arg0 *C.GtkAccelGroup  // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(accelKey)
	_arg2 = C.GdkModifierType(accelMods)

	_cret = C.gtk_accel_group_disconnect_key(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a accelGroup) IsLocked() bool {
	var _arg0 *C.GtkAccelGroup // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_accel_group_get_is_locked(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a accelGroup) ModifierMask() gdk.ModifierType {
	var _arg0 *C.GtkAccelGroup  // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_accel_group_get_modifier_mask(_arg0)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

func (a accelGroup) LockAccelGroup() {
	var _arg0 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	C.gtk_accel_group_lock(_arg0)
}

func (a accelGroup) UnlockAccelGroup() {
	var _arg0 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAccelGroup)(unsafe.Pointer(a.Native()))

	C.gtk_accel_group_unlock(_arg0)
}

type AccelKey struct {
	native C.GtkAccelKey
}

// WrapAccelKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAccelKey(ptr unsafe.Pointer) *AccelKey {
	return (*AccelKey)(ptr)
}

// Native returns the underlying C source pointer.
func (a *AccelKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}
