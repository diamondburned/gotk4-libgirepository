// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11displaymanager.go.
var GTypeX11DisplayManager = externglib.Type(C.gdk_x11_display_manager_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11DisplayManager, F: marshalX11DisplayManager},
	})
}

// X11DisplayManagerOverrider contains methods that are overridable.
type X11DisplayManagerOverrider interface {
	externglib.Objector
}

type X11DisplayManager struct {
	_ [0]func() // equal guard
	gdk.DisplayManager
}

var (
	_ externglib.Objector = (*X11DisplayManager)(nil)
)

func classInitX11DisplayManagerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11DisplayManager(obj *externglib.Object) *X11DisplayManager {
	return &X11DisplayManager{
		DisplayManager: gdk.DisplayManager{
			Object: obj,
		},
	}
}

func marshalX11DisplayManager(p uintptr) (interface{}, error) {
	return wrapX11DisplayManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
