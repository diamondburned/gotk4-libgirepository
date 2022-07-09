// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeHSeparator returns the GType for the type HSeparator.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeHSeparator() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "HSeparator").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalHSeparator)
	return gtype
}

// HSeparatorOverrider contains methods that are overridable.
type HSeparatorOverrider interface {
}

// HSeparator widget is a horizontal separator, used to group the widgets within
// a window. It displays a horizontal line with a shadow to make it appear
// sunken into the interface.
//
// > The HSeparator widget is not used as a separator within menus. > To create
// a separator in a menu create an empty SeparatorMenuItem > widget using
// gtk_separator_menu_item_new() and add it to the menu with >
// gtk_menu_shell_append().
//
// GtkHSeparator has been deprecated, use Separator instead.
type HSeparator struct {
	_ [0]func() // equal guard
	Separator
}

var (
	_ Widgetter         = (*HSeparator)(nil)
	_ coreglib.Objector = (*HSeparator)(nil)
)

func classInitHSeparatorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHSeparator(obj *coreglib.Object) *HSeparator {
	return &HSeparator{
		Separator: Separator{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalHSeparator(p uintptr) (interface{}, error) {
	return wrapHSeparator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHSeparator creates a new HSeparator.
//
// Deprecated: Use gtk_separator_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function returns the following values:
//
//    - hSeparator: new HSeparator.
//
func NewHSeparator() *HSeparator {
	_gret := girepository.MustFind("Gtk", "HSeparator").InvokeMethod("new_HSeparator", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _hSeparator *HSeparator // out

	_hSeparator = wrapHSeparator(coreglib.Take(unsafe.Pointer(_cret)))

	return _hSeparator
}
