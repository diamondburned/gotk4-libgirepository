// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_level_bar_accessible_get_type()), F: marshalLevelBarAccessibler},
	})
}

// LevelBarAccessibler describes LevelBarAccessible's methods.
type LevelBarAccessibler interface {
	privateLevelBarAccessible()
}

//
type LevelBarAccessible struct {
	WidgetAccessible

	atk.Value
}

var (
	_ LevelBarAccessibler = (*LevelBarAccessible)(nil)
	_ gextras.Nativer     = (*LevelBarAccessible)(nil)
)

func wrapLevelBarAccessible(obj *externglib.Object) LevelBarAccessibler {
	return &LevelBarAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalLevelBarAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLevelBarAccessible(obj), nil
}

func (*LevelBarAccessible) privateLevelBarAccessible() {}
