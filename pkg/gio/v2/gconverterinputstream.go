// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_converter_input_stream_get_type()), F: marshalConverterInputStream},
	})
}

// ConverterInputStream: converter input stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, InputStream implements InputStream.
type ConverterInputStream interface {
	FilterInputStream
	PollableInputStream
}

// converterInputStream implements the ConverterInputStream class.
type converterInputStream struct {
	FilterInputStream
	PollableInputStream
}

var _ ConverterInputStream = (*converterInputStream)(nil)

// WrapConverterInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapConverterInputStream(obj *externglib.Object) ConverterInputStream {
	return converterInputStream{
		FilterInputStream:   WrapFilterInputStream(obj),
		PollableInputStream: WrapPollableInputStream(obj),
	}
}

func marshalConverterInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConverterInputStream(obj), nil
}
