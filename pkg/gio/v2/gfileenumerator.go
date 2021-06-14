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
		{T: externglib.Type(C.g_file_enumerator_get_type()), F: marshalFileEnumerator},
	})
}

// FileEnumerator allows you to operate on a set of #GFiles, returning a Info
// structure for each file enumerated (e.g. g_file_enumerate_children() will
// return a Enumerator for each of the children within a directory).
//
// To get the next file's information from a Enumerator, use
// g_file_enumerator_next_file() or its asynchronous version,
// g_file_enumerator_next_files_async(). Note that the asynchronous version will
// return a list of Infos, whereas the synchronous will only return the next
// file in the enumerator.
//
// The ordering of returned files is unspecified for non-Unix platforms; for
// more information, see g_dir_read_name(). On Unix, when operating on local
// files, returned files will be sorted by inode number. Effectively you can
// assume that the ordering of returned files will be stable between successive
// calls (and applications) assuming the directory is unchanged.
//
// If your application needs a specific ordering, such as by name or
// modification time, you will have to implement that in your application code.
//
// To close a Enumerator, use g_file_enumerator_close(), or its asynchronous
// version, g_file_enumerator_close_async(). Once a Enumerator is closed, no
// further actions may be performed on it, and it should be freed with
// g_object_unref().
type FileEnumerator interface {
	gextras.Objector

	// HasPending checks if the file enumerator has pending operations.
	HasPending() bool
	// IsClosed checks if the file enumerator has been closed.
	IsClosed() bool
	// SetPending sets the file enumerator as having pending operations.
	SetPending(pending bool)
}

// fileEnumerator implements the FileEnumerator class.
type fileEnumerator struct {
	gextras.Objector
}

var _ FileEnumerator = (*fileEnumerator)(nil)

// WrapFileEnumerator wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileEnumerator(obj *externglib.Object) FileEnumerator {
	return fileEnumerator{
		Objector: obj,
	}
}

func marshalFileEnumerator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileEnumerator(obj), nil
}

// HasPending checks if the file enumerator has pending operations.
func (e fileEnumerator) HasPending() bool {
	var _arg0 *C.GFileEnumerator // out

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	var _cret C.gboolean // in

	_cret = C.g_file_enumerator_has_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if the file enumerator has been closed.
func (e fileEnumerator) IsClosed() bool {
	var _arg0 *C.GFileEnumerator // out

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	var _cret C.gboolean // in

	_cret = C.g_file_enumerator_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPending sets the file enumerator as having pending operations.
func (e fileEnumerator) SetPending(pending bool) {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	if pending {
		_arg1 = C.TRUE
	}

	C.g_file_enumerator_set_pending(_arg0, _arg1)
}
