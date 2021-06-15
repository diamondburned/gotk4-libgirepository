// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hpaned_get_type()), F: marshalHPaned},
	})
}

// HPaned: the HPaned widget is a container widget with two children arranged
// horizontally. The division between the two panes is adjustable by the user by
// dragging a handle. See Paned for details.
//
// GtkHPaned has been deprecated, use Paned instead.
type HPaned interface {
	Paned
	Buildable
	Orientable
}

// hPaned implements the HPaned class.
type hPaned struct {
	Paned
	Buildable
	Orientable
}

var _ HPaned = (*hPaned)(nil)

// WrapHPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapHPaned(obj *externglib.Object) HPaned {
	return hPaned{
		Paned:      WrapPaned(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHPaned(obj), nil
}

// NewHPaned constructs a class HPaned.
func NewHPaned() HPaned {
	var _cret C.GtkHPaned // in

	_cret = C.gtk_hpaned_new()

	var _hPaned HPaned // out

	_hPaned = WrapHPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _hPaned
}
