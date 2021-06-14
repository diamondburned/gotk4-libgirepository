// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: `GtkMountOperation` is an implementation of
// `GMountOperation`.
//
// The functions and objects described here make working with GTK and GIO more
// convenient.
//
// `GtkMountOperation` is needed when mounting volumes: It is an implementation
// of `GMountOperation` that can be used with GIO functions for mounting volumes
// such as g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, `GtkMountOperation` shows dialogs to let the user enter
// passwords, ask questions or show processes blocking unmount.
type MountOperation interface {
	gio.MountOperation

	// IsShowing returns whether the `GtkMountOperation` is currently displaying
	// a window.
	IsShowing() bool
	// SetDisplay sets the display to show windows of the `GtkMountOperation`
	// on.
	SetDisplay(display gdk.Display)
	// SetParent sets the transient parent for windows shown by the
	// `GtkMountOperation`.
	SetParent(parent Window)
}

// mountOperation implements the MountOperation class.
type mountOperation struct {
	gio.MountOperation
}

var _ MountOperation = (*mountOperation)(nil)

// WrapMountOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return mountOperation{
		gio.MountOperation: gio.WrapMountOperation(obj),
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

// IsShowing returns whether the `GtkMountOperation` is currently displaying
// a window.
func (o mountOperation) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_mount_operation_is_showing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDisplay sets the display to show windows of the `GtkMountOperation`
// on.
func (o mountOperation) SetDisplay(display gdk.Display) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_mount_operation_set_display(_arg0, _arg1)
}

// SetParent sets the transient parent for windows shown by the
// `GtkMountOperation`.
func (o mountOperation) SetParent(parent Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}
