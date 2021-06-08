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
func ShowURI(parent Window, urI string, timestamp uint32) {
	var arg1 *C.GtkWindow
	var arg2 *C.char
	var arg3 C.guint32

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	arg2 = (*C.char)(C.CString(urI))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.guint32(timestamp)

	C.gtk_show_uri(arg1, arg2, arg3)
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
	C.gtk_show_uri_full(arg1, arg2, arg3, arg4, arg5, arg6)
}

// ShowURIFullFinish finishes the gtk_show_uri() call and returns the result of
// the operation.
func ShowURIFullFinish(parent Window, result gio.AsyncResult) error {
	var arg1 *C.GtkWindow
	var arg2 *C.GAsyncResult

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	arg2 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError
	var goerr error

	C.gtk_show_uri_full_finish(arg1, arg2, &cerr)

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}
