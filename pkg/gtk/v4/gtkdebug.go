// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk.h>
import "C"

// GetDebugFlags returns the GTK debug flags that are currently active.
//
// This function is intended for GTK modules that want to adjust their debug
// output based on GTK debug flags.
func GetDebugFlags() DebugFlags {
	var _cret C.GtkDebugFlags // in

	_cret = C.gtk_get_debug_flags()

	var _debugFlags DebugFlags // out

	_debugFlags = DebugFlags(_cret)

	return _debugFlags
}
