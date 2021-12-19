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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_manager_get_type()), F: marshalX11DisplayManagerer},
	})
}

// X11DisplayManagerOverrider contains methods that are overridable.
type X11DisplayManagerOverrider interface {
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

func marshalX11DisplayManagerer(p uintptr) (interface{}, error) {
	return wrapX11DisplayManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
