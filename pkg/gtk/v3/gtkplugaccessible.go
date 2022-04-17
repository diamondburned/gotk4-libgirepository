// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkplugaccessible.go.
var GTypePlugAccessible = externglib.Type(C.gtk_plug_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePlugAccessible, F: marshalPlugAccessible},
	})
}

// PlugAccessibleOverrider contains methods that are overridable.
type PlugAccessibleOverrider interface {
	externglib.Objector
}

type PlugAccessible struct {
	_ [0]func() // equal guard
	WindowAccessible
}

var (
	_ externglib.Objector = (*PlugAccessible)(nil)
)

func classInitPlugAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPlugAccessible(obj *externglib.Object) *PlugAccessible {
	return &PlugAccessible{
		WindowAccessible: WindowAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						ObjectClass: atk.ObjectClass{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
			Window: atk.Window{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	return wrapPlugAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func (plug *PlugAccessible) ID() string {
	var _arg0 *C.GtkPlugAccessible // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(externglib.InternObject(plug).Native()))

	_cret = C.gtk_plug_accessible_get_id(_arg0)
	runtime.KeepAlive(plug)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
