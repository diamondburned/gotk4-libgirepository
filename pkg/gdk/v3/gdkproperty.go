// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 glib-2.0 gtk+-3.0
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
type PropMode int

const (
	// PropModeReplace: the new data replaces the existing data.
	PropModeReplace PropMode = 0
	// PropModePrepend: the new data is prepended to the existing data.
	PropModePrepend PropMode = 1
	// PropModeAppend: the new data is appended to the existing data.
	PropModeAppend PropMode = 2
)

func marshalPropMode(p uintptr) (interface{}, error) {
	return PropMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PropertyDelete deletes a property from a window.
func PropertyDelete(window Window, property *Atom) {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.GdkAtom    // out

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(property.Native()))

	C.gdk_property_delete(_arg1, _arg2)
}

// PropertyGet retrieves a portion of the contents of a property. If the
// property does not exist, then the function returns false, and GDK_NONE will
// be stored in @actual_property_type.
//
// The XGetWindowProperty() function that gdk_property_get() uses has a very
// confusing and complicated set of semantics. Unfortunately, gdk_property_get()
// makes the situation worse instead of better (the semantics should be
// considered undefined), and also prints warnings to stderr in cases where it
// should return a useful error to the program. You are advised to use
// XGetWindowProperty() directly until a replacement function for
// gdk_property_get() is provided.
func PropertyGet(window Window, property *Atom, typ *Atom, offset uint32, length uint32, pdelete int) (Atom, int, []byte, bool) {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.GdkAtom    // out
	var _arg4 C.gulong     // out
	var _arg5 C.gulong     // out
	var _arg6 C.gint       // out
	var _actualPropertyType Atom
	var _arg8 C.gint // in
	var _arg10 *C.guchar
	var _arg9 C.gint     // in
	var _cret C.gboolean // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(property.Native()))
	_arg3 = (C.GdkAtom)(unsafe.Pointer(typ.Native()))
	_arg4 = (C.gulong)(offset)
	_arg5 = (C.gulong)(length)
	_arg6 = (C.gint)(pdelete)

	_cret = C.gdk_property_get(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, (*C.GdkAtom)(unsafe.Pointer(&_actualPropertyType)), &_arg8, &_arg9, &_arg10)

	var _actualFormat int // out
	var _data []byte
	var _ok bool // out

	_actualFormat = (int)(_arg8)
	_data = unsafe.Slice((*byte)(unsafe.Pointer(_arg10)), _arg9)
	runtime.SetFinalizer(&_data, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _actualPropertyType, _actualFormat, _data, _ok
}

// TextPropertyToUTF8ListForDisplay converts a text property in the given
// encoding to a list of UTF-8 strings.
func TextPropertyToUTF8ListForDisplay(display Display, encoding *Atom, format int, text []byte) ([]string, int) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.GdkAtom     // out
	var _arg3 C.gint        // out
	var _arg4 *C.guchar
	var _arg5 C.gint
	var _arg6 **C.gchar
	var _cret C.gint // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(encoding.Native()))
	_arg3 = (C.gint)(format)
	_arg5 = C.gint(len(text))
	_arg4 = (*C.guchar)(unsafe.Pointer(&text[0]))

	_cret = C.gdk_text_property_to_utf8_list_for_display(_arg1, _arg2, _arg3, _arg4, _arg5, &_arg6)

	var _list []string
	var _gint int // out

	{
		var i int
		var z *C.gchar
		for p := _arg6; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg6, i)
		_list = make([]string, i)
		for i := range src {
			_list[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_gint = (int)(_cret)

	return _list, _gint
}

// UTF8ToStringTarget converts an UTF-8 string into the best possible
// representation as a STRING. The representation of characters not in STRING is
// not specified; it may be as pseudo-escape sequences \x{ABCD}, or it may be in
// some other form of approximation.
func UTF8ToStringTarget(str string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_utf8_to_string_target(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
