// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// TickCallback: callback type for adding a function to update animations. See
// gtk_widget_add_tick_callback().
type TickCallback func(widget Widgetter, frameClock gdk.FrameClocker) (ok bool)
