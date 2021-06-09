// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// ShowURI: this function launches the default application for showing a given
// uri, or shows an error dialog if that fails.
func ShowURI(parent Window, uri string, timestamp uint32) {
	var _arg1 *C.GtkWindow
	var _arg2 *C.char
	var _arg3 C.guint32

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint32(timestamp)

	C.gtk_show_uri(_arg1, _arg2, _arg3)
}

// ShowURIFull: this function launches the default application for showing a
// given uri.
//
// The @callback will be called when the launch is completed. It should call
// gtk_show_uri_full_finish() to obtain the result.
//
// This is the recommended call to be used as it passes information necessary
// for sandbox helpers to parent their dialogs properly.
func ShowURIFull() {
	C.gtk_show_uri_full()
}

// ShowURIFullFinish finishes the gtk_show_uri() call and returns the result of
// the operation.
func ShowURIFullFinish(parent Window, result gio.AsyncResult) error {
	var _arg1 *C.GtkWindow
	var _arg2 *C.GAsyncResult

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cerr *C.GError

	C.gtk_show_uri_full_finish(_arg1, _arg2, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}