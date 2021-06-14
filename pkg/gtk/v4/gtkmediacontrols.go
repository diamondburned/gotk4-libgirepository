// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_controls_get_type()), F: marshalMediaControls},
	})
}

// MediaControls: `GtkMediaControls` is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, `GtkMediaControls` is used as part of [class@Gtk.Video].
type MediaControls interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// SetMediaStream sets the stream that is controlled by @controls.
	SetMediaStream(stream MediaStream)
}

// mediaControls implements the MediaControls class.
type mediaControls struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ MediaControls = (*mediaControls)(nil)

// WrapMediaControls wraps a GObject to the right type. It is
// primarily used internally.
func WrapMediaControls(obj *externglib.Object) MediaControls {
	return mediaControls{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalMediaControls(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaControls(obj), nil
}

// SetMediaStream sets the stream that is controlled by @controls.
func (c mediaControls) SetMediaStream(stream MediaStream) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(stream.Native()))

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
}
