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

// AcceleratorGetLabel converts an accelerator keyval and modifier mask into a
// string which can be used to represent the accelerator to the user.
func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out

	_arg1 = C.guint(acceleratorKey)
	_arg2 = (C.GdkModifierType)(acceleratorMods)

	var _cret *C.char // in

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

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = (C.GdkModifierType)(acceleratorMods)

	var _cret *C.char // in

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

	_arg1 = C.guint(acceleratorKey)
	_arg2 = (C.GdkModifierType)(acceleratorMods)

	var _cret *C.char // in

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

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = C.guint(acceleratorKey)
	_arg3 = C.guint(keycode)
	_arg4 = (C.GdkModifierType)(acceleratorMods)

	var _cret *C.char // in

	_cret = C.gtk_accelerator_name_with_keycode(_arg1, _arg2, _arg3, _arg4)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
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

	_arg1 = C.guint(keyval)
	_arg2 = (C.GdkModifierType)(modifiers)

	var _cret C.gboolean // in

	_cret = C.gtk_accelerator_valid(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
