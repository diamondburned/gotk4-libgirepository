// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hseparator_get_type()), F: marshalHSeparator},
	})
}

// HSeparator: the HSeparator widget is a horizontal separator, used to group
// the widgets within a window. It displays a horizontal line with a shadow to
// make it appear sunken into the interface.
//
// > The HSeparator widget is not used as a separator within menus. > To create
// a separator in a menu create an empty SeparatorMenuItem > widget using
// gtk_separator_menu_item_new() and add it to the menu with >
// gtk_menu_shell_append().
//
// GtkHSeparator has been deprecated, use Separator instead.
type HSeparator interface {
	Separator
	Buildable
	Orientable
}

// hSeparator implements the HSeparator interface.
type hSeparator struct {
	Separator
	Buildable
	Orientable
}

var _ HSeparator = (*hSeparator)(nil)

// WrapHSeparator wraps a GObject to the right type. It is
// primarily used internally.
func WrapHSeparator(obj *externglib.Object) HSeparator {
	return HSeparator{
		Separator:  WrapSeparator(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHSeparator(obj), nil
}

// NewHSeparator constructs a class HSeparator.
func NewHSeparator() HSeparator {
	var _cret C.GtkHSeparator

	cret = C.gtk_hseparator_new()

	var _hSeparator HSeparator

	_hSeparator = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(HSeparator)

	return _hSeparator
}
