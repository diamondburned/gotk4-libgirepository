// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_controls_get_type()), F: marshalMediaControlser},
	})
}

// MediaControls: GtkMediaControls is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, GtkMediaControls is used as part of gtk.Video.
type MediaControls struct {
	Widget

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Widgetter = (*MediaControls)(nil)
)

func wrapMediaControls(obj *externglib.Object) *MediaControls {
	return &MediaControls{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalMediaControlser(p uintptr) (interface{}, error) {
	return wrapMediaControls(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMediaControls creates a new GtkMediaControls managing the stream passed to
// it.
//
// The function takes the following parameters:
//
//    - stream (optional) to manage or NULL for none.
//
// The function returns the following values:
//
//    - mediaControls: new GtkMediaControls.
//
func NewMediaControls(stream MediaStreamer) *MediaControls {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	if stream != nil {
		_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(stream.Native()))
	}

	_cret = C.gtk_media_controls_new(_arg1)
	runtime.KeepAlive(stream)

	var _mediaControls *MediaControls // out

	_mediaControls = wrapMediaControls(externglib.Take(unsafe.Pointer(_cret)))

	return _mediaControls
}

// MediaStream gets the media stream managed by controls or NULL if none.
//
// The function returns the following values:
//
//    - mediaStream (optional): media stream managed by controls.
//
func (controls *MediaControls) MediaStream() MediaStreamer {
	var _arg0 *C.GtkMediaControls // out
	var _cret *C.GtkMediaStream   // in

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(controls.Native()))

	_cret = C.gtk_media_controls_get_media_stream(_arg0)
	runtime.KeepAlive(controls)

	var _mediaStream MediaStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(MediaStreamer)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.MediaStreamer")
			}
			_mediaStream = rv
		}
	}

	return _mediaStream
}

// SetMediaStream sets the stream that is controlled by controls.
//
// The function takes the following parameters:
//
//    - stream (optional): GtkMediaStream, or NULL.
//
func (controls *MediaControls) SetMediaStream(stream MediaStreamer) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(controls.Native()))
	if stream != nil {
		_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(stream.Native()))
	}

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
	runtime.KeepAlive(controls)
	runtime.KeepAlive(stream)
}
