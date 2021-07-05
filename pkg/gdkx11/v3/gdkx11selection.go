// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
import "C"

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
func X11FreeCompoundText(ctext *byte) {
	var _arg1 *C.guchar // out

	_arg1 = (*C.guchar)(unsafe.Pointer(ctext))

	C.gdk_x11_free_compound_text(_arg1)
}

// X11FreeTextList frees the array of strings created by
// gdk_x11_display_text_property_to_text_list().
func X11FreeTextList(list *string) {
	var _arg1 **C.gchar // out

	{
		var refTmpIn string
		var refTmpOut *C.gchar

		refTmpIn = *list

		refTmpOut = (*C.gchar)(C.CString(refTmpIn))
		defer C.free(unsafe.Pointer(refTmpOut))

		if refTmpOut != nil {
			out0 := &refTmpOut
			_arg1 = out0
		}
	}

	C.gdk_x11_free_text_list(_arg1)
}
