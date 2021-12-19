// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vpaned_get_type()), F: marshalVPanedder},
	})
}

// VPanedOverrider contains methods that are overridable.
type VPanedOverrider interface {
}

// VPaned widget is a container widget with two children arranged vertically.
// The division between the two panes is adjustable by the user by dragging a
// handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned struct {
	_ [0]func() // equal guard
	Paned
}

var (
	_ Containerer         = (*VPaned)(nil)
	_ externglib.Objector = (*VPaned)(nil)
)

func classInitVPanedder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVPaned(obj *externglib.Object) *VPaned {
	return &VPaned{
		Paned: Paned{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalVPanedder(p uintptr) (interface{}, error) {
	return wrapVPaned(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVPaned: create a new VPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function returns the following values:
//
//    - vPaned: new VPaned.
//
func NewVPaned() *VPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vpaned_new()

	var _vPaned *VPaned // out

	_vPaned = wrapVPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _vPaned
}
