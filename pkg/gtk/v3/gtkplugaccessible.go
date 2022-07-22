// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypePlugAccessible = coreglib.Type(C.gtk_plug_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePlugAccessible, F: marshalPlugAccessible},
	})
}

// PlugAccessibleOverrider contains methods that are overridable.
type PlugAccessibleOverrider interface {
}

type PlugAccessible struct {
	_ [0]func() // equal guard
	WindowAccessible
}

var (
	_ coreglib.Objector = (*PlugAccessible)(nil)
)

func initClassPlugAccessible(gclass unsafe.Pointer, goval any) {
}

func wrapPlugAccessible(obj *coreglib.Object) *PlugAccessible {
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
	return wrapPlugAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func (plug *PlugAccessible) ID() string {
	var _arg0 *C.GtkPlugAccessible // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(coreglib.InternObject(plug).Native()))

	_cret = C.gtk_plug_accessible_get_id(_arg0)
	runtime.KeepAlive(plug)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
