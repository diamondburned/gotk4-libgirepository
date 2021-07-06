// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_stream_get_type()), F: marshalMediaStream},
	})
}

// MediaStream: `GtkMediaStream` is the integration point for media playback
// inside GTK.
//
// GTK provides an implementation of the `GtkMediaStream` interface that is
// called [class@Gtk.MediaFile].
//
// Apart from application-facing API for stream playback, `GtkMediaStream` has a
// number of APIs that are only useful for implementations and should not be
// used in applications: [method@Gtk.MediaStream.prepared],
// [method@Gtk.MediaStream.unprepared], [method@Gtk.MediaStream.update],
// [method@Gtk.MediaStream.ended], [method@Gtk.MediaStream.seek_success],
// [method@Gtk.MediaStream.seek_failed], [method@Gtk.MediaStream.gerror],
// [method@Gtk.MediaStream.error], [method@Gtk.MediaStream.error_valist].
type MediaStream interface {
	gextras.Objector

	// Ended pauses the media stream and marks it as ended.
	//
	// This is a hint only, calls to GtkMediaStream.play() may still happen.
	//
	// The media stream must be prepared when this function is called.
	Ended()
	// Gerror sets @self into an error state.
	//
	// This will pause the stream (you can check for an error via
	// [method@Gtk.MediaStream.get_error] in your GtkMediaStream.pause()
	// implementation), abort pending seeks and mark the stream as prepared.
	//
	// if the stream is already in an error state, this call will be ignored and
	// the existing error will be retained.
	//
	// To unset an error, the stream must be reset via a call to
	// [method@Gtk.MediaStream.unprepared].
	Gerror(err error)
	// Duration gets the duration of the stream.
	//
	// If the duration is not known, 0 will be returned.
	Duration() int64
	// GetEnded returns whether the streams playback is finished.
	GetEnded() bool
	// Error: if the stream is in an error state, returns the `GError`
	// explaining that state.
	//
	// Any type of error can be reported here depending on the implementation of
	// the media stream.
	//
	// A media stream in an error cannot be operated on, calls like
	// [method@Gtk.MediaStream.play] or [method@Gtk.MediaStream.seek] will not
	// have any effect.
	//
	// `GtkMediaStream` itself does not provide a way to unset an error, but
	// implementations may provide options. For example, a [class@Gtk.MediaFile]
	// will unset errors when a new source is set, e.g. with
	// [method@Gtk.MediaFile.set_file].
	Error() error
	// Loop returns whether the stream is set to loop.
	//
	// See [method@Gtk.MediaStream.set_loop] for details.
	Loop() bool
	// Muted returns whether the audio for the stream is muted.
	//
	// See [method@Gtk.MediaStream.set_muted] for details.
	Muted() bool
	// Playing: return whether the stream is currently playing.
	Playing() bool
	// Timestamp returns the current presentation timestamp in microseconds.
	Timestamp() int64
	// Volume returns the volume of the audio for the stream.
	//
	// See [method@Gtk.MediaStream.set_volume] for details.
	Volume() float64
	// HasAudio returns whether the stream has audio.
	HasAudio() bool
	// HasVideo returns whether the stream has video.
	HasVideo() bool
	// IsPrepared returns whether the stream has finished initializing.
	//
	// At this point the existence of audio and video is known.
	IsPrepared() bool
	// IsSeekable checks if a stream may be seekable.
	//
	// This is meant to be a hint. Streams may not allow seeking even if this
	// function returns true. However, if this function returns false, streams
	// are guaranteed to not be seekable and user interfaces may hide controls
	// that allow seeking.
	//
	// It is allowed to call [method@Gtk.MediaStream.seek] on a non-seekable
	// stream, though it will not do anything.
	IsSeekable() bool
	// IsSeeking checks if there is currently a seek operation going on.
	IsSeeking() bool
	// Pause pauses playback of the stream.
	//
	// If the stream is not playing, do nothing.
	Pause()
	// Play starts playing the stream.
	//
	// If the stream is in error or already playing, do nothing.
	Play()
	// Prepared: called by `GtkMediaStream` implementations to advertise the
	// stream being ready to play and providing details about the stream.
	//
	// Note that the arguments are hints. If the stream implementation cannot
	// determine the correct values, it is better to err on the side of caution
	// and return true. User interfaces will use those values to determine what
	// controls to show.
	//
	// This function may not be called again until the stream has been reset via
	// [method@Gtk.MediaStream.unprepared].
	Prepared(hasAudio bool, hasVideo bool, seekable bool, duration int64)
	// Realize: called by users to attach the media stream to a `GdkSurface`
	// they manage.
	//
	// The stream can then access the resources of @surface for its rendering
	// purposes. In particular, media streams might want to create a
	// `GdkGLContext` or sync to the `GdkFrameClock`.
	//
	// Whoever calls this function is responsible for calling
	// [method@Gtk.MediaStream.unrealize] before either the stream or @surface
	// get destroyed.
	//
	// Multiple calls to this function may happen from different users of the
	// video, even with the same @surface. Each of these calls must be followed
	// by its own call to [method@Gtk.MediaStream.unrealize].
	//
	// It is not required to call this function to make a media stream work.
	Realize(surface gdk.Surface)
	// Seek: start a seek operation on @self to @timestamp.
	//
	// If @timestamp is out of range, it will be clamped.
	//
	// Seek operations may not finish instantly. While a seek operation is in
	// process, the [property@Gtk.MediaStream:seeking] property will be set.
	//
	// When calling gtk_media_stream_seek() during an ongoing seek operation,
	// the new seek will override any pending seek.
	Seek(timestamp int64)
	// SeekFailed ends a seek operation started via GtkMediaStream.seek() as a
	// failure.
	//
	// This will not cause an error on the stream and will assume that playback
	// continues as if no seek had happened.
	//
	// See [method@Gtk.MediaStream.seek_success] for the other way of ending a
	// seek.
	SeekFailed()
	// SeekSuccess ends a seek operation started via GtkMediaStream.seek()
	// successfully.
	//
	// This function will unset the GtkMediaStream:ended property if it was set.
	//
	// See [method@Gtk.MediaStream.seek_failed] for the other way of ending a
	// seek.
	SeekSuccess()
	// SetLoop sets whether the stream should loop.
	//
	// In this case, it will attempt to restart playback from the beginning
	// instead of stopping at the end.
	//
	// Not all streams may support looping, in particular non-seekable streams.
	// Those streams will ignore the loop setting and just end.
	SetLoop(loop bool)
	// SetMuted sets whether the audio stream should be muted.
	//
	// Muting a stream will cause no audio to be played, but it does not modify
	// the volume. This means that muting and then unmuting the stream will
	// restore the volume settings.
	//
	// If the stream has no audio, calling this function will still work but it
	// will not have an audible effect.
	SetMuted(muted bool)
	// SetPlaying starts or pauses playback of the stream.
	SetPlaying(playing bool)
	// SetVolume sets the volume of the audio stream.
	//
	// This function call will work even if the stream is muted.
	//
	// The given @volume should range from 0.0 for silence to 1.0 for as loud as
	// possible. Values outside of this range will be clamped to the nearest
	// value.
	//
	// If the stream has no audio or is muted, calling this function will still
	// work but it will not have an immediate audible effect. When the stream is
	// unmuted, the new volume setting will take effect.
	SetVolume(volume float64)
	// Unprepared resets a given media stream implementation.
	//
	// [method@Gtk.MediaStream.prepared] can then be called again.
	//
	// This function will also reset any error state the stream was in.
	Unprepared()
	// Unrealize undoes a previous call to gtk_media_stream_realize().
	//
	// This causes the stream to release all resources it had allocated from
	// @surface.
	Unrealize(surface gdk.Surface)
	// Update: media stream implementations should regularly call this function
	// to update the timestamp reported by the stream.
	//
	// It is up to implementations to call this at the frequency they deem
	// appropriate.
	//
	// The media stream must be prepared when this function is called.
	Update(timestamp int64)
}

// mediaStream implements the MediaStream interface.
type mediaStream struct {
	*externglib.Object
}

var _ MediaStream = (*mediaStream)(nil)

// WrapMediaStream wraps a GObject to a type that implements
// interface MediaStream. It is primarily used internally.
func WrapMediaStream(obj *externglib.Object) MediaStream {
	return mediaStream{obj}
}

func marshalMediaStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaStream(obj), nil
}

func (s mediaStream) Ended() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_ended(_arg0)
}

func (s mediaStream) Gerror(err error) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 *C.GError         // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.gtk_media_stream_gerror(_arg0, _arg1)
}

func (s mediaStream) Duration() int64 {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gint64          // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_duration(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (s mediaStream) GetEnded() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_ended(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) Error() error {
	var _arg0 *C.GtkMediaStream // out
	var _cret *C.GError         // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_error(_arg0)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

func (s mediaStream) Loop() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_loop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) Muted() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_muted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) Playing() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_playing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) Timestamp() int64 {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gint64          // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_timestamp(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (s mediaStream) Volume() float64 {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.double          // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_get_volume(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s mediaStream) HasAudio() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_has_audio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) HasVideo() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_has_video(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) IsPrepared() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_is_prepared(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) IsSeekable() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_is_seekable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) IsSeeking() bool {
	var _arg0 *C.GtkMediaStream // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_stream_is_seeking(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s mediaStream) Pause() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_pause(_arg0)
}

func (s mediaStream) Play() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_play(_arg0)
}

func (s mediaStream) Prepared(hasAudio bool, hasVideo bool, seekable bool, duration int64) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gboolean        // out
	var _arg2 C.gboolean        // out
	var _arg3 C.gboolean        // out
	var _arg4 C.gint64          // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	if hasAudio {
		_arg1 = C.TRUE
	}
	if hasVideo {
		_arg2 = C.TRUE
	}
	if seekable {
		_arg3 = C.TRUE
	}
	_arg4 = C.gint64(duration)

	C.gtk_media_stream_prepared(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (s mediaStream) Realize(surface gdk.Surface) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 *C.GdkSurface     // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gtk_media_stream_realize(_arg0, _arg1)
}

func (s mediaStream) Seek(timestamp int64) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gint64          // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint64(timestamp)

	C.gtk_media_stream_seek(_arg0, _arg1)
}

func (s mediaStream) SeekFailed() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_seek_failed(_arg0)
}

func (s mediaStream) SeekSuccess() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_seek_success(_arg0)
}

func (s mediaStream) SetLoop(loop bool) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	if loop {
		_arg1 = C.TRUE
	}

	C.gtk_media_stream_set_loop(_arg0, _arg1)
}

func (s mediaStream) SetMuted(muted bool) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	if muted {
		_arg1 = C.TRUE
	}

	C.gtk_media_stream_set_muted(_arg0, _arg1)
}

func (s mediaStream) SetPlaying(playing bool) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	if playing {
		_arg1 = C.TRUE
	}

	C.gtk_media_stream_set_playing(_arg0, _arg1)
}

func (s mediaStream) SetVolume(volume float64) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(volume)

	C.gtk_media_stream_set_volume(_arg0, _arg1)
}

func (s mediaStream) Unprepared() {
	var _arg0 *C.GtkMediaStream // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))

	C.gtk_media_stream_unprepared(_arg0)
}

func (s mediaStream) Unrealize(surface gdk.Surface) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 *C.GdkSurface     // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gtk_media_stream_unrealize(_arg0, _arg1)
}

func (s mediaStream) Update(timestamp int64) {
	var _arg0 *C.GtkMediaStream // out
	var _arg1 C.gint64          // out

	_arg0 = (*C.GtkMediaStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint64(timestamp)

	C.gtk_media_stream_update(_arg0, _arg1)
}
