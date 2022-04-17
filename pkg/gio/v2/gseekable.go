// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_SeekableIface_can_seek(GSeekable*);
// extern gboolean _gotk4_gio2_SeekableIface_can_truncate(GSeekable*);
// extern gboolean _gotk4_gio2_SeekableIface_seek(GSeekable*, goffset, GSeekType, GCancellable*, GError**);
// extern gboolean _gotk4_gio2_SeekableIface_truncate_fn(GSeekable*, goffset, GCancellable*, GError**);
// extern goffset _gotk4_gio2_SeekableIface_tell(GSeekable*);
import "C"

// glib.Type values for gseekable.go.
var GTypeSeekable = externglib.Type(C.g_seekable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSeekable, F: marshalSeekable},
	})
}

// SeekableOverrider contains methods that are overridable.
type SeekableOverrider interface {
	externglib.Objector
	// CanSeek tests if the stream supports the Iface.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if seekable can be seeked. FALSE otherwise.
	//
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the stream can be truncated, FALSE otherwise.
	//
	CanTruncate() bool
	// Seek seeks in the stream by the given offset, modified by type.
	//
	// Attempting to seek past the end of the stream will have different results
	// depending on if the stream is fixed-sized or resizable. If the stream is
	// resizable then seeking past the end and then writing will result in zeros
	// filling the empty space. Seeking past the end of a resizable stream and
	// reading will result in EOF. Seeking past the end of a fixed-sized stream
	// will fail.
	//
	// Any operation that would result in a negative offset will fail.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - offset: #goffset.
	//    - typ: Type.
	//
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	// Tell tells the current position within the stream.
	//
	// The function returns the following values:
	//
	//    - gint64: offset from the beginning of the buffer.
	//
	Tell() int64
	// TruncateFn sets the length of the stream to offset. If the stream was
	// previously larger than offset, the extra data is discarded. If the stream
	// was previously shorter than offset, it is extended with NUL ('\0') bytes.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - offset: new length for seekable, in bytes.
	//
	TruncateFn(ctx context.Context, offset int64) error
}

// Seekable is implemented by streams (implementations of Stream or Stream) that
// support seeking.
//
// Seekable streams largely fall into two categories: resizable and fixed-size.
//
// #GSeekable on fixed-sized streams is approximately the same as POSIX lseek()
// on a block device (for example: attempting to seek past the end of the device
// is an error). Fixed streams typically cannot be truncated.
//
// #GSeekable on resizable streams is approximately the same as POSIX lseek() on
// a normal file. Seeking past the end and writing data will usually cause the
// stream to resize by introducing zero bytes.
type Seekable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Seekable)(nil)
)

// Seekabler describes Seekable's interface methods.
type Seekabler interface {
	externglib.Objector

	// CanSeek tests if the stream supports the Iface.
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	CanTruncate() bool
	// Seek seeks in the stream by the given offset, modified by type.
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	// Tell tells the current position within the stream.
	Tell() int64
	// Truncate sets the length of the stream to offset.
	Truncate(ctx context.Context, offset int64) error
}

var _ Seekabler = (*Seekable)(nil)

func ifaceInitSeekabler(gifacePtr, data C.gpointer) {
	iface := (*C.GSeekableIface)(unsafe.Pointer(gifacePtr))
	iface.can_seek = (*[0]byte)(C._gotk4_gio2_SeekableIface_can_seek)
	iface.can_truncate = (*[0]byte)(C._gotk4_gio2_SeekableIface_can_truncate)
	iface.seek = (*[0]byte)(C._gotk4_gio2_SeekableIface_seek)
	iface.tell = (*[0]byte)(C._gotk4_gio2_SeekableIface_tell)
	iface.truncate_fn = (*[0]byte)(C._gotk4_gio2_SeekableIface_truncate_fn)
}

//export _gotk4_gio2_SeekableIface_can_seek
func _gotk4_gio2_SeekableIface_can_seek(arg0 *C.GSeekable) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	ok := iface.CanSeek()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_SeekableIface_can_truncate
func _gotk4_gio2_SeekableIface_can_truncate(arg0 *C.GSeekable) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	ok := iface.CanTruncate()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_SeekableIface_seek
func _gotk4_gio2_SeekableIface_seek(arg0 *C.GSeekable, arg1 C.goffset, arg2 C.GSeekType, arg3 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	var _cancellable context.Context // out
	var _offset int64                // out
	var _typ glib.SeekType           // out

	if arg3 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg3))
	}
	_offset = int64(arg1)
	_typ = glib.SeekType(arg2)

	_goerr := iface.Seek(_cancellable, _offset, _typ)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_SeekableIface_tell
func _gotk4_gio2_SeekableIface_tell(arg0 *C.GSeekable) (cret C.goffset) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	gint64 := iface.Tell()

	cret = C.goffset(gint64)

	return cret
}

//export _gotk4_gio2_SeekableIface_truncate_fn
func _gotk4_gio2_SeekableIface_truncate_fn(arg0 *C.GSeekable, arg1 C.goffset, arg2 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SeekableOverrider)

	var _cancellable context.Context // out
	var _offset int64                // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_offset = int64(arg1)

	_goerr := iface.TruncateFn(_cancellable, _offset)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

func wrapSeekable(obj *externglib.Object) *Seekable {
	return &Seekable{
		Object: obj,
	}
}

func marshalSeekable(p uintptr) (interface{}, error) {
	return wrapSeekable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanSeek tests if the stream supports the Iface.
//
// The function returns the following values:
//
//    - ok: TRUE if seekable can be seeked. FALSE otherwise.
//
func (seekable *Seekable) CanSeek() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(externglib.InternObject(seekable).Native()))

	_cret = C.g_seekable_can_seek(_arg0)
	runtime.KeepAlive(seekable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanTruncate tests if the length of the stream can be adjusted with
// g_seekable_truncate().
//
// The function returns the following values:
//
//    - ok: TRUE if the stream can be truncated, FALSE otherwise.
//
func (seekable *Seekable) CanTruncate() bool {
	var _arg0 *C.GSeekable // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(externglib.InternObject(seekable).Native()))

	_cret = C.g_seekable_can_truncate(_arg0)
	runtime.KeepAlive(seekable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Seek seeks in the stream by the given offset, modified by type.
//
// Attempting to seek past the end of the stream will have different results
// depending on if the stream is fixed-sized or resizable. If the stream is
// resizable then seeking past the end and then writing will result in zeros
// filling the empty space. Seeking past the end of a resizable stream and
// reading will result in EOF. Seeking past the end of a fixed-sized stream will
// fail.
//
// Any operation that would result in a negative offset will fail.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - offset: #goffset.
//    - typ: Type.
//
func (seekable *Seekable) Seek(ctx context.Context, offset int64, typ glib.SeekType) error {
	var _arg0 *C.GSeekable    // out
	var _arg3 *C.GCancellable // out
	var _arg1 C.goffset       // out
	var _arg2 C.GSeekType     // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(externglib.InternObject(seekable).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.goffset(offset)
	_arg2 = C.GSeekType(typ)

	C.g_seekable_seek(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(seekable)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(typ)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Tell tells the current position within the stream.
//
// The function returns the following values:
//
//    - gint64: offset from the beginning of the buffer.
//
func (seekable *Seekable) Tell() int64 {
	var _arg0 *C.GSeekable // out
	var _cret C.goffset    // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(externglib.InternObject(seekable).Native()))

	_cret = C.g_seekable_tell(_arg0)
	runtime.KeepAlive(seekable)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// Truncate sets the length of the stream to offset. If the stream was
// previously larger than offset, the extra data is discarded. If the stream was
// previously shorter than offset, it is extended with NUL ('\0') bytes.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - offset: new length for seekable, in bytes.
//
func (seekable *Seekable) Truncate(ctx context.Context, offset int64) error {
	var _arg0 *C.GSeekable    // out
	var _arg2 *C.GCancellable // out
	var _arg1 C.goffset       // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GSeekable)(unsafe.Pointer(externglib.InternObject(seekable).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.goffset(offset)

	C.g_seekable_truncate(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(seekable)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(offset)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
