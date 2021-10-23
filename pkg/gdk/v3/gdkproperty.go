// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_prop_mode_get_type()), F: marshalPropMode},
	})
}

// PropMode describes how existing data is combined with new data when using
// gdk_property_change().
type PropMode C.gint

const (
	// PropModeReplace: new data replaces the existing data.
	PropModeReplace PropMode = iota
	// PropModePrepend: new data is prepended to the existing data.
	PropModePrepend
	// PropModeAppend: new data is appended to the existing data.
	PropModeAppend
)

func marshalPropMode(p uintptr) (interface{}, error) {
	return PropMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PropMode.
func (p PropMode) String() string {
	switch p {
	case PropModeReplace:
		return "Replace"
	case PropModePrepend:
		return "Prepend"
	case PropModeAppend:
		return "Append"
	default:
		return fmt.Sprintf("PropMode(%d)", p)
	}
}

// UTF8ToStringTarget converts an UTF-8 string into the best possible
// representation as a STRING. The representation of characters not in STRING is
// not specified; it may be as pseudo-escape sequences \x{ABCD}, or it may be in
// some other form of approximation.
//
// The function takes the following parameters:
//
//    - str: UTF-8 string.
//
func UTF8ToStringTarget(str string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_utf8_to_string_target(_arg1)
	runtime.KeepAlive(str)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}
