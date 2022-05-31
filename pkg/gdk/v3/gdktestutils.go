// Code generated by girgen. DO NOT EDIT.

package gdk

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

// TestRenderSync retrieves a pixel from window to force the windowing system to
// carry out any pending rendering commands.
//
// This function is intended to be used to synchronize with rendering pipelines,
// to benchmark windowing system rendering operations.
//
// The function takes the following parameters:
//
//    - window: mapped Window.
//
func TestRenderSync(window Windower) {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gdk", "test_render_sync").Invoke(_args[:], nil)

	runtime.KeepAlive(window)
}
