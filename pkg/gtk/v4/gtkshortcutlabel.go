// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_label_get_type()), F: marshalShortcutLabeller},
	})
}

// ShortcutLabel: GtkShortcutLabel displays a single keyboard shortcut or
// gesture.
//
// The main use case for GtkShortcutLabel is inside a gtk.ShortcutsWindow.
type ShortcutLabel struct {
	Widget
}

func wrapShortcutLabel(obj *externglib.Object) *ShortcutLabel {
	return &ShortcutLabel{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalShortcutLabeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutLabel(obj), nil
}

// NewShortcutLabel creates a new GtkShortcutLabel with accelerator set.
func NewShortcutLabel(accelerator string) *ShortcutLabel {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_shortcut_label_new(_arg1)

	var _shortcutLabel *ShortcutLabel // out

	_shortcutLabel = wrapShortcutLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _shortcutLabel
}

// Accelerator retrieves the current accelerator of self.
func (self *ShortcutLabel) Accelerator() string {
	var _arg0 *C.GtkShortcutLabel // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_label_get_accelerator(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DisabledText retrieves the text that is displayed when no accelerator is set.
func (self *ShortcutLabel) DisabledText() string {
	var _arg0 *C.GtkShortcutLabel // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_label_get_disabled_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetAccelerator sets the accelerator to be displayed by self.
func (self *ShortcutLabel) SetAccelerator(accelerator string) {
	var _arg0 *C.GtkShortcutLabel // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_shortcut_label_set_accelerator(_arg0, _arg1)
}

// SetDisabledText sets the text to be displayed by self when no accelerator is
// set.
func (self *ShortcutLabel) SetDisabledText(disabledText string) {
	var _arg0 *C.GtkShortcutLabel // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(disabledText)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_shortcut_label_set_disabled_text(_arg0, _arg1)
}
