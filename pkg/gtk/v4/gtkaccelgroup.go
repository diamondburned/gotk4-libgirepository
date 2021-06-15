// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// AcceleratorGetDefaultModMask gets the modifier mask.
//
// The modifier mask determines which modifiers are considered significant for
// keyboard accelerators. This includes all keyboard modifiers except for
// GDK_LOCK_MASK.
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
	var _cret *C.char           // in

	_arg1 = (C.guint)(acceleratorKey)
	_arg2 = (C.GdkModifierType)(acceleratorMods)

	_cret = C.gtk_accelerator_get_label(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorGetLabelWithKeycode converts an accelerator keyval and modifier
// mask into a string that can be displayed to the user.
//
// The string may be translated.
//
// This function is similar to [func@Gtk.accelerator_get_label], but handling
// keycodes.
//
// This is only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
func AcceleratorGetLabelWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.char           // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (C.guint)(acceleratorKey)
	_arg3 = (C.guint)(keycode)
	_arg4 = (C.GdkModifierType)(acceleratorMods)

	_cret = C.gtk_accelerator_get_label_with_keycode(_arg1, _arg2, _arg3, _arg4)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorName converts an accelerator keyval and modifier mask into a
// string parseable by gtk_accelerator_parse().
//
// For example, if you pass in GDK_KEY_q and GDK_CONTROL_MASK, this function
// returns `<Control>q`.
//
// If you need to display accelerators in the user interface, see
// [func@Gtk.accelerator_get_label].
func AcceleratorName(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret *C.char           // in

	_arg1 = (C.guint)(acceleratorKey)
	_arg2 = (C.GdkModifierType)(acceleratorMods)

	_cret = C.gtk_accelerator_name(_arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorNameWithKeycode converts an accelerator keyval and modifier mask
// into a string parseable by gtk_accelerator_parse_with_keycode().
//
// This is similar to [func@Gtk.accelerator_name] but handling keycodes. This is
// only useful for system-level components, applications should use
// gtk_accelerator_parse() instead.
func AcceleratorNameWithKeycode(display gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 *C.GdkDisplay     // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _arg4 C.GdkModifierType // out
	var _cret *C.char           // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (C.guint)(acceleratorKey)
	_arg3 = (C.guint)(keycode)
	_arg4 = (C.GdkModifierType)(acceleratorMods)

	_cret = C.gtk_accelerator_name_with_keycode(_arg1, _arg2, _arg3, _arg4)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AcceleratorParse parses a string representing an accelerator.
//
// The format looks like “<Control>a” or “<Shift><Alt>F1”.
//
// The parser is fairly liberal and allows lower or upper case, and also
// abbreviations such as “<Ctl>” and “<Ctrl>”. Key names are parsed using
// [func@Gdk.keyval_from_name]. For character keys the name is not the symbol,
// but the lowercase name, e.g. one would use “<Ctrl>minus” instead of
// “<Ctrl>-”.
//
// If the parse fails, @accelerator_key and @accelerator_mods will be set to 0
// (zero).
func AcceleratorParse(accelerator string) (uint, gdk.ModifierType, bool) {
	var _arg1 *C.char           // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg1 = (*C.char)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_accelerator_parse(_arg1, &_arg2, &_arg3)

	var _acceleratorKey uint              // out
	var _acceleratorMods gdk.ModifierType // out
	var _ok bool                          // out

	_acceleratorKey = (uint)(_arg2)
	_acceleratorMods = gdk.ModifierType(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _acceleratorKey, _acceleratorMods, _ok
}

// AcceleratorParseWithKeycode parses a string representing an accelerator.
//
// This is similar to [func@Gtk.accelerator_parse] but handles keycodes as well.
// This is only useful for system-level components, applications should use
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
func AcceleratorParseWithKeycode(accelerator string, display gdk.Display) (uint, []uint, gdk.ModifierType, bool) {
	var _arg1 *C.char       // out
	var _arg2 *C.GdkDisplay // out
	var _arg3 C.guint       // in
	var _arg4 *C.guint
	var _arg5 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg1 = (*C.char)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_accelerator_parse_with_keycode(_arg1, _arg2, &_arg3, &_arg4, &_arg5)

	var _acceleratorKey uint // out
	var _acceleratorCodes []uint
	var _acceleratorMods gdk.ModifierType // out
	var _ok bool                          // out

	_acceleratorKey = (uint)(_arg3)
	{
		var i int
		var z C.guint
		for p := _arg4; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg4, i)
		_acceleratorCodes = make([]uint, i)
		for i := range src {
			_acceleratorCodes[i] = (uint)(src[i])
		}
	}
	_acceleratorMods = gdk.ModifierType(_arg5)
	if _cret != 0 {
		_ok = true
	}

	return _acceleratorKey, _acceleratorCodes, _acceleratorMods, _ok
}

// AcceleratorValid determines whether a given keyval and modifier mask
// constitute a valid keyboard accelerator.
//
// For example, the GDK_KEY_a keyval plus GDK_CONTROL_MASK mark is valid, and
// matches the “Ctrl+a” accelerator. But, you can't, for instance, use the
// GDK_KEY_Control_L keyval as an accelerator.
func AcceleratorValid(keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (C.guint)(keyval)
	_arg2 = (C.GdkModifierType)(modifiers)

	_cret = C.gtk_accelerator_valid(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
