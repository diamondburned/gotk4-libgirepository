// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkimcontextsimple.go.
var GTypeIMContextSimple = coreglib.Type(C.gtk_im_context_simple_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeIMContextSimple, F: marshalIMContextSimple},
	})
}

const MAX_COMPOSE_LEN = 7

// IMContextSimpleOverrider contains methods that are overridable.
type IMContextSimpleOverrider interface {
}

// IMContextSimple: GtkIMContextSimple is an input method supporting table-based
// input methods.
//
// GtkIMContextSimple has a built-in table of compose sequences that is derived
// from the X11 Compose files.
//
// GtkIMContextSimple reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-4.0/Compose, ~/.XCompose,
// /usr/share/X11/locale/$locale/Compose (for locales that have a nontrivial
// Compose file). The syntax of these files is described in the Compose(5)
// manual page.
//
//
// Unicode characters
//
// GtkIMContextSimple also supports numeric entry of Unicode characters by
// typing <kbd>Ctrl</kbd>-<kbd>Shift</kbd>-<kbd>u</kbd>, followed by a
// hexadecimal Unicode codepoint.
//
// For example,
//
//    Ctrl-Shift-u 1 2 3 Enter
//
// yields U+0123 LATIN SMALL LETTER G WITH CEDILLA, i.e. ģ.
type IMContextSimple struct {
	_ [0]func() // equal guard
	IMContext
}

var (
	_ IMContexter = (*IMContextSimple)(nil)
)

func classInitIMContextSimpler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIMContextSimple(obj *coreglib.Object) *IMContextSimple {
	return &IMContextSimple{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMContextSimple(p uintptr) (interface{}, error) {
	return wrapIMContextSimple(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMContextSimple creates a new IMContextSimple.
//
// The function returns the following values:
//
//    - imContextSimple: new IMContextSimple.
//
func NewIMContextSimple() *IMContextSimple {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "IMContextSimple").InvokeMethod("new_IMContextSimple", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _imContextSimple *IMContextSimple // out

	_imContextSimple = wrapIMContextSimple(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imContextSimple
}

// AddComposeFile adds an additional table from the X11 compose file.
//
// The function takes the following parameters:
//
//    - composeFile: path of compose file.
//
func (contextSimple *IMContextSimple) AddComposeFile(composeFile string) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(contextSimple).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(composeFile)))
	defer C.free(unsafe.Pointer(_arg1))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "IMContextSimple").InvokeMethod("add_compose_file", _args[:], nil)

	runtime.KeepAlive(contextSimple)
	runtime.KeepAlive(composeFile)
}
