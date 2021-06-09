// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

func SynthesizeWindowState(window Window, unsetFlags WindowState, setFlags WindowState) {
	var _arg1 *C.GdkWindow
	var _arg2 C.GdkWindowState
	var _arg3 C.GdkWindowState

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (C.GdkWindowState)(unsetFlags)
	_arg3 = (C.GdkWindowState)(setFlags)

	C.gdk_synthesize_window_state(_arg1, _arg2, _arg3)
}
