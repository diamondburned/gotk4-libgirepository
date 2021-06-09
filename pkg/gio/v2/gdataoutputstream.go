// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_data_output_stream_get_type()), F: marshalDataOutputStream},
	})
}

// DataOutputStream: data output stream implements Stream and includes functions
// for writing data directly to an output stream.
type DataOutputStream interface {
	FilterOutputStream
	Seekable

	// ByteOrder gets the byte order for the stream.
	ByteOrder() DataStreamByteOrder
	// PutByte puts a byte into the output stream.
	PutByte(data byte, cancellable Cancellable) error
	// PutInt16 puts a signed 16-bit integer into the output stream.
	PutInt16(data int16, cancellable Cancellable) error
	// PutInt32 puts a signed 32-bit integer into the output stream.
	PutInt32(data int32, cancellable Cancellable) error
	// PutInt64 puts a signed 64-bit integer into the stream.
	PutInt64(data int64, cancellable Cancellable) error
	// PutString puts a string into the output stream.
	PutString(str string, cancellable Cancellable) error
	// PutUint16 puts an unsigned 16-bit integer into the output stream.
	PutUint16(data uint16, cancellable Cancellable) error
	// PutUint32 puts an unsigned 32-bit integer into the stream.
	PutUint32(data uint32, cancellable Cancellable) error
	// PutUint64 puts an unsigned 64-bit integer into the stream.
	PutUint64(data uint64, cancellable Cancellable) error
	// SetByteOrder sets the byte order of the data output stream to @order.
	SetByteOrder(order DataStreamByteOrder)
}

// dataOutputStream implements the DataOutputStream interface.
type dataOutputStream struct {
	FilterOutputStream
	Seekable
}

var _ DataOutputStream = (*dataOutputStream)(nil)

// WrapDataOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapDataOutputStream(obj *externglib.Object) DataOutputStream {
	return DataOutputStream{
		FilterOutputStream: WrapFilterOutputStream(obj),
		Seekable:           WrapSeekable(obj),
	}
}

func marshalDataOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDataOutputStream(obj), nil
}

// NewDataOutputStream constructs a class DataOutputStream.
func NewDataOutputStream(baseStream OutputStream) DataOutputStream {
	var _arg1 *C.GOutputStream

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))

	var _cret C.GDataOutputStream

	cret = C.g_data_output_stream_new(_arg1)

	var _dataOutputStream DataOutputStream

	_dataOutputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DataOutputStream)

	return _dataOutputStream
}

// ByteOrder gets the byte order for the stream.
func (s dataOutputStream) ByteOrder() DataStreamByteOrder {
	var _arg0 *C.GDataOutputStream

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))

	var _cret C.GDataStreamByteOrder

	cret = C.g_data_output_stream_get_byte_order(_arg0)

	var _dataStreamByteOrder DataStreamByteOrder

	_dataStreamByteOrder = DataStreamByteOrder(_cret)

	return _dataStreamByteOrder
}

// PutByte puts a byte into the output stream.
func (s dataOutputStream) PutByte(data byte, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.guchar
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.guchar(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_byte(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt16 puts a signed 16-bit integer into the output stream.
func (s dataOutputStream) PutInt16(data int16, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.gint16
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint16(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_int16(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt32 puts a signed 32-bit integer into the output stream.
func (s dataOutputStream) PutInt32(data int32, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.gint32
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint32(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_int32(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt64 puts a signed 64-bit integer into the stream.
func (s dataOutputStream) PutInt64(data int64, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.gint64
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint64(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_int64(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutString puts a string into the output stream.
func (s dataOutputStream) PutString(str string, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 *C.char
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_string(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint16 puts an unsigned 16-bit integer into the output stream.
func (s dataOutputStream) PutUint16(data uint16, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.guint16
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint16(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_uint16(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint32 puts an unsigned 32-bit integer into the stream.
func (s dataOutputStream) PutUint32(data uint32, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.guint32
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint32(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_uint32(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint64 puts an unsigned 64-bit integer into the stream.
func (s dataOutputStream) PutUint64(data uint64, cancellable Cancellable) error {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.guint64
	var _arg2 *C.GCancellable

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint64(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError

	C.g_data_output_stream_put_uint64(_arg0, _arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetByteOrder sets the byte order of the data output stream to @order.
func (s dataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	var _arg0 *C.GDataOutputStream
	var _arg1 C.GDataStreamByteOrder

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GDataStreamByteOrder)(order)

	C.g_data_output_stream_set_byte_order(_arg0, _arg1)
}
