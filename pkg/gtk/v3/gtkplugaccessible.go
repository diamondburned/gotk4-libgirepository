// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_plug_accessible_get_type()), F: marshalPlugAccessibler},
	})
}

type PlugAccessible struct {
	WindowAccessible
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

func marshalPlugAccessibler(p uintptr) (interface{}, error) {
	return wrapPlugAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (plug *PlugAccessible) ID() string {
	var _arg0 *C.GtkPlugAccessible // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(plug.Native()))

	_cret = C.gtk_plug_accessible_get_id(_arg0)
	runtime.KeepAlive(plug)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
