// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
import "C"

// SourceSetDummyCallback sets a dummy callback for @source. The callback will
// do nothing, and if the source expects a #gboolean return value, it will
// return true. (If the source expects any other type of return value, it will
// return a 0/nil value; whatever g_value_init() initializes a #GValue to for
// that type.)
//
// If the source is not one of the standard GLib types, the @closure_callback
// and @closure_marshal fields of the Funcs structure must have been filled in
// with pointers to appropriate functions.
func SourceSetDummyCallback(source *glib.Source) {
	var _arg1 *C.GSource

	_arg1 = (*C.GSource)(unsafe.Pointer(source.Native()))

	C.g_source_set_dummy_callback(_arg1)
}
