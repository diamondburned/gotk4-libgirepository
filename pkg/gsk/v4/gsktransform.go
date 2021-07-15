// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gsk/gsk.h>
import "C"

// TransformParse parses the given string into a transform and puts it in
// out_transform.
//
// Strings printed via gsk.Transform.ToString() can be read in again
// successfully using this function.
//
// If string does not describe a valid transform, FALSE is returned and NULL is
// put in out_transform.
func TransformParse(_string string) (*Transform, bool) {
	var _arg1 *C.char         // out
	var _arg2 *C.GskTransform // in
	var _cret C.gboolean      // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))

	_cret = C.gsk_transform_parse(_arg1, &_arg2)

	var _outTransform *Transform // out
	var _ok bool                 // out

	_outTransform = (*Transform)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
	C.gsk_transform_ref(_arg2)
	runtime.SetFinalizer(_outTransform, func(v *Transform) {
		C.gsk_transform_unref((*C.GskTransform)(gextras.StructNative(unsafe.Pointer(v))))
	})
	if _cret != 0 {
		_ok = true
	}

	return _outTransform, _ok
}
