// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeDataOutputStream returns the GType for the type DataOutputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDataOutputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DataOutputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDataOutputStream)
	return gtype
}

// DataOutputStreamOverrider contains methods that are overridable.
type DataOutputStreamOverrider interface {
}

// DataOutputStream: data output stream implements Stream and includes functions
// for writing data directly to an output stream.
type DataOutputStream struct {
	_ [0]func() // equal guard
	FilterOutputStream

	Seekable
}

var (
	_ FilterOutputStreamer = (*DataOutputStream)(nil)
)

func classInitDataOutputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDataOutputStream(obj *coreglib.Object) *DataOutputStream {
	return &DataOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalDataOutputStream(p uintptr) (interface{}, error) {
	return wrapDataOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDataOutputStream creates a new data output stream for base_stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//
// The function returns the following values:
//
//    - dataOutputStream: OutputStream.
//
func NewDataOutputStream(baseStream OutputStreamer) *DataOutputStream {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))

	_gret := girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("new_DataOutputStream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)

	var _dataOutputStream *DataOutputStream // out

	_dataOutputStream = wrapDataOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dataOutputStream
}

// PutByte puts a byte into the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #guchar.
//
func (stream *DataOutputStream) PutByte(ctx context.Context, data byte) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.guchar)(unsafe.Pointer(&_args[1])) = C.guchar(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_byte", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt16 puts a signed 16-bit integer into the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #gint16.
//
func (stream *DataOutputStream) PutInt16(ctx context.Context, data int16) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gint16)(unsafe.Pointer(&_args[1])) = C.gint16(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_int16", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt32 puts a signed 32-bit integer into the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #gint32.
//
func (stream *DataOutputStream) PutInt32(ctx context.Context, data int32) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gint32)(unsafe.Pointer(&_args[1])) = C.gint32(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_int32", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt64 puts a signed 64-bit integer into the stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #gint64.
//
func (stream *DataOutputStream) PutInt64(ctx context.Context, data int64) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gint64)(unsafe.Pointer(&_args[1])) = C.gint64(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_int64", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutString puts a string into the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - str: string.
//
func (stream *DataOutputStream) PutString(ctx context.Context, str string) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_string", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(str)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint16 puts an unsigned 16-bit integer into the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #guint16.
//
func (stream *DataOutputStream) PutUint16(ctx context.Context, data uint16) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.guint16)(unsafe.Pointer(&_args[1])) = C.guint16(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_uint16", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint32 puts an unsigned 32-bit integer into the stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #guint32.
//
func (stream *DataOutputStream) PutUint32(ctx context.Context, data uint32) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.guint32)(unsafe.Pointer(&_args[1])) = C.guint32(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_uint32", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint64 puts an unsigned 64-bit integer into the stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - data: #guint64.
//
func (stream *DataOutputStream) PutUint64(ctx context.Context, data uint64) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.guint64)(unsafe.Pointer(&_args[1])) = C.guint64(data)

	girepository.MustFind("Gio", "DataOutputStream").InvokeMethod("put_uint64", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(data)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
