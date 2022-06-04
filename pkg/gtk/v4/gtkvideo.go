// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkvideo.go.
var GTypeVideo = coreglib.Type(C.gtk_video_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeVideo, F: marshalVideo},
	})
}

// VideoOverrider contains methods that are overridable.
type VideoOverrider interface {
}

// Video: GtkVideo is a widget to show a GtkMediaStream with media controls.
//
// !An example GtkVideo (video.png)
//
// The controls are available separately as gtk.MediaControls. If you just want
// to display a video without controls, you can treat it like any other
// paintable and for example put it into a gtk.Picture.
//
// GtkVideo aims to cover use cases such as previews, embedded animations, etc.
// It supports autoplay, looping, and simple media controls. It does not have
// support for video overlays, multichannel audio, device selection, or input.
// If you are writing a full-fledged video player, you may want to use the
// gdk.Paintable API and a media framework such as Gstreamer directly.
type Video struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Video)(nil)
)

func classInitVideoer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVideo(obj *coreglib.Object) *Video {
	return &Video{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalVideo(p uintptr) (interface{}, error) {
	return wrapVideo(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVideo creates a new empty GtkVideo.
//
// The function returns the following values:
//
//    - video: new GtkVideo.
//
func NewVideo() *Video {
	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("new_Video", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _video *Video // out

	_video = wrapVideo(coreglib.Take(unsafe.Pointer(_cret)))

	return _video
}

// NewVideoForFile creates a GtkVideo to play back the given file.
//
// The function takes the following parameters:
//
//    - file (optional): GFile.
//
// The function returns the following values:
//
//    - video: new GtkVideo.
//
func NewVideoForFile(file gio.Filer) *Video {
	var _args [1]girepository.Argument

	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("new_Video_for_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(file)

	var _video *Video // out

	_video = wrapVideo(coreglib.Take(unsafe.Pointer(_cret)))

	return _video
}

// NewVideoForFilename creates a GtkVideo to play back the given filename.
//
// This is a utility function that calls gtk.Video.NewForFile, See that function
// for details.
//
// The function takes the following parameters:
//
//    - filename (optional) to play back.
//
// The function returns the following values:
//
//    - video: new GtkVideo.
//
func NewVideoForFilename(filename string) *Video {
	var _args [1]girepository.Argument

	if filename != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("new_Video_for_filename", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filename)

	var _video *Video // out

	_video = wrapVideo(coreglib.Take(unsafe.Pointer(_cret)))

	return _video
}

// NewVideoForMediaStream creates a GtkVideo to play back the given stream.
//
// The function takes the following parameters:
//
//    - stream (optional): GtkMediaStream.
//
// The function returns the following values:
//
//    - video: new GtkVideo.
//
func NewVideoForMediaStream(stream MediaStreamer) *Video {
	var _args [1]girepository.Argument

	if stream != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	}

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("new_Video_for_media_stream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _video *Video // out

	_video = wrapVideo(coreglib.Take(unsafe.Pointer(_cret)))

	return _video
}

// NewVideoForResource creates a GtkVideo to play back the resource at the given
// resource_path.
//
// This is a utility function that calls gtk.Video.NewForFile.
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource path to play back.
//
// The function returns the following values:
//
//    - video: new GtkVideo.
//
func NewVideoForResource(resourcePath string) *Video {
	var _args [1]girepository.Argument

	if resourcePath != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("new_Video_for_resource", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resourcePath)

	var _video *Video // out

	_video = wrapVideo(coreglib.Take(unsafe.Pointer(_cret)))

	return _video
}

// Autoplay returns TRUE if videos have been set to loop.
//
// The function returns the following values:
//
//    - ok: TRUE if streams should autoplay.
//
func (self *Video) Autoplay() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("get_autoplay", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// File gets the file played by self or NULL if not playing back a file.
//
// The function returns the following values:
//
//    - file (optional) played by self.
//
func (self *Video) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("get_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _file *gio.File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// Loop returns TRUE if videos have been set to loop.
//
// The function returns the following values:
//
//    - ok: TRUE if streams should loop.
//
func (self *Video) Loop() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("get_loop", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MediaStream gets the media stream managed by self or NULL if none.
//
// The function returns the following values:
//
//    - mediaStream (optional): media stream managed by self.
//
func (self *Video) MediaStream() MediaStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "Video").InvokeMethod("get_media_stream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _mediaStream MediaStreamer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(MediaStreamer)
				return ok
			})
			rv, ok := casted.(MediaStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.MediaStreamer")
			}
			_mediaStream = rv
		}
	}

	return _mediaStream
}

// SetAutoplay sets whether self automatically starts playback when it becomes
// visible or when a new file gets loaded.
//
// The function takes the following parameters:
//
//    - autoplay: whether media streams should autoplay.
//
func (self *Video) SetAutoplay(autoplay bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if autoplay {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_autoplay", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(autoplay)
}

// SetFile makes self play the given file.
//
// The function takes the following parameters:
//
//    - file (optional) to play.
//
func (self *Video) SetFile(file gio.Filer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_file", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetFilename makes self play the given filename.
//
// This is a utility function that calls gtk_video_set_file(),.
//
// The function takes the following parameters:
//
//    - filename (optional) to play.
//
func (self *Video) SetFilename(filename string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if filename != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_filename", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(filename)
}

// SetLoop sets whether new files loaded by self should be set to loop.
//
// The function takes the following parameters:
//
//    - loop: whether media streams should loop.
//
func (self *Video) SetLoop(loop bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if loop {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_loop", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(loop)
}

// SetMediaStream sets the media stream to be played back.
//
// self will take full control of managing the media stream. If you want to
// manage a media stream yourself, consider using a gtk.Picture for display.
//
// If you want to display a file, consider using gtk.Video.SetFile() instead.
//
// The function takes the following parameters:
//
//    - stream (optional): media stream to play or NULL to unset.
//
func (self *Video) SetMediaStream(stream MediaStreamer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if stream != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_media_stream", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(stream)
}

// SetResource makes self play the resource at the given resource_path.
//
// This is a utility function that calls gtk.Video.SetFile().
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource to set.
//
func (self *Video) SetResource(resourcePath string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if resourcePath != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "Video").InvokeMethod("set_resource", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(resourcePath)
}
