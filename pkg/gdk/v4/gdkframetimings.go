// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_frame_timings_get_type()), F: marshalFrameTimings},
	})
}

// FrameTimings: `GdkFrameTimings` object holds timing information for a single
// frame of the application’s displays.
//
// To retrieve `GdkFrameTimings` objects, use
// [method@Gdk.FrameClock.get_timings] or
// [method@Gdk.FrameClock.get_current_timings]. The information in
// `GdkFrameTimings` is useful for precise synchronization of video with the
// event or audio streams, and for measuring quality metrics for the
// application’s display, such as latency and jitter.
type FrameTimings struct {
	native C.GdkFrameTimings
}

// WrapFrameTimings wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFrameTimings(ptr unsafe.Pointer) *FrameTimings {
	return (*FrameTimings)(ptr)
}

func marshalFrameTimings(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*FrameTimings)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FrameTimings) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// Complete returns whether @timings are complete.
//
// The timing information in a `GdkFrameTimings` is filled in incrementally as
// the frame as drawn and passed off to the window system for processing and
// display to the user. The accessor functions for `GdkFrameTimings` can return
// 0 to indicate an unavailable value for two reasons: either because the
// information is not yet available, or because it isn't available at all.
//
// Once this function returns true for a frame, you can be certain that no
// further values will become available and be stored in the `GdkFrameTimings`.
func (t *FrameTimings) Complete() bool {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_complete(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FrameCounter gets the frame counter value of the `GdkFrameClock` when this
// frame was drawn.
func (t *FrameTimings) FrameCounter() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_frame_counter(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// FrameTime returns the frame time for the frame.
//
// This is the time value that is typically used to time animations for the
// frame. See [method@Gdk.FrameClock.get_frame_time].
func (t *FrameTimings) FrameTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_frame_time(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PredictedPresentationTime gets the predicted time at which this frame will be
// displayed.
//
// Although no predicted time may be available, if one is available, it will be
// available while the frame is being generated, in contrast to
// [method@Gdk.FrameTimings.get_presentation_time], which is only available
// after the frame has been presented.
//
// In general, if you are simply animating, you should use
// [method@Gdk.FrameClock.get_frame_time] rather than this function, but this
// function is useful for applications that want exact control over latency. For
// example, a movie player may want this information for Audio/Video
// synchronization.
func (t *FrameTimings) PredictedPresentationTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_predicted_presentation_time(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PresentationTime reurns the presentation time.
//
// This is the time at which the frame became visible to the user.
func (t *FrameTimings) PresentationTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_presentation_time(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// RefreshInterval gets the natural interval between presentation times for the
// display that this frame was displayed on.
//
// Frame presentation usually happens during the “vertical blanking interval”.
func (t *FrameTimings) RefreshInterval() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_get_refresh_interval(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// Ref increases the reference count of @timings.
func (t *FrameTimings) Ref() *FrameTimings {
	var _arg0 *C.GdkFrameTimings // out
	var _cret *C.GdkFrameTimings // in

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	_cret = C.gdk_frame_timings_ref(_arg0)

	var _frameTimings *FrameTimings // out

	_frameTimings = (*FrameTimings)(unsafe.Pointer(_cret))
	C.gdk_frame_timings_ref(_cret)
	runtime.SetFinalizer(_frameTimings, func(v *FrameTimings) {
		C.gdk_frame_timings_unref((*C.GdkFrameTimings)(unsafe.Pointer(v)))
	})

	return _frameTimings
}

// Unref decreases the reference count of @timings.
//
// If @timings is no longer referenced, it will be freed.
func (t *FrameTimings) Unref() {
	var _arg0 *C.GdkFrameTimings // out

	_arg0 = (*C.GdkFrameTimings)(unsafe.Pointer(t))

	C.gdk_frame_timings_unref(_arg0)
}
