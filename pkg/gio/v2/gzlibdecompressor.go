// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_zlib_decompressor_get_type()), F: marshalZlibDecompressor},
	})
}

// ZlibDecompressor: zlib decompression
type ZlibDecompressor interface {
	gextras.Objector
	Converter

	// FileInfo retrieves the Info constructed from the GZIP header data of
	// compressed data processed by @compressor, or nil if @decompressor's
	// Decompressor:format property is not G_ZLIB_COMPRESSOR_FORMAT_GZIP, or the
	// header data was not fully processed yet, or it not present in the data
	// stream at all.
	FileInfo() FileInfo
}

// zlibDecompressor implements the ZlibDecompressor class.
type zlibDecompressor struct {
	gextras.Objector
	Converter
}

var _ ZlibDecompressor = (*zlibDecompressor)(nil)

// WrapZlibDecompressor wraps a GObject to the right type. It is
// primarily used internally.
func WrapZlibDecompressor(obj *externglib.Object) ZlibDecompressor {
	return zlibDecompressor{
		Objector:  obj,
		Converter: WrapConverter(obj),
	}
}

func marshalZlibDecompressor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapZlibDecompressor(obj), nil
}

// NewZlibDecompressor constructs a class ZlibDecompressor.
func NewZlibDecompressor(format ZlibCompressorFormat) ZlibDecompressor {
	var _arg1 C.GZlibCompressorFormat // out
	var _cret C.GZlibDecompressor     // in

	_arg1 = (C.GZlibCompressorFormat)(format)

	_cret = C.g_zlib_decompressor_new(_arg1)

	var _zlibDecompressor ZlibDecompressor // out

	_zlibDecompressor = WrapZlibDecompressor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _zlibDecompressor
}

// FileInfo retrieves the Info constructed from the GZIP header data of
// compressed data processed by @compressor, or nil if @decompressor's
// Decompressor:format property is not G_ZLIB_COMPRESSOR_FORMAT_GZIP, or the
// header data was not fully processed yet, or it not present in the data
// stream at all.
func (d zlibDecompressor) FileInfo() FileInfo {
	var _arg0 *C.GZlibDecompressor // out
	var _cret *C.GFileInfo         // in

	_arg0 = (*C.GZlibDecompressor)(unsafe.Pointer(d.Native()))

	_cret = C.g_zlib_decompressor_get_file_info(_arg0)

	var _fileInfo FileInfo // out

	_fileInfo = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FileInfo)

	return _fileInfo
}
