// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_cell_accessible_get_type()), F: marshalImageCellAccessibler},
	})
}

type ImageCellAccessible struct {
	RendererCellAccessible

	*externglib.Object
	atk.Image
	atk.ObjectClass

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ externglib.Objector = (*ImageCellAccessible)(nil)
)

func wrapImageCellAccessible(obj *externglib.Object) *ImageCellAccessible {
	return &ImageCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
				TableCell: atk.TableCell{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		Image: atk.Image{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalImageCellAccessibler(p uintptr) (interface{}, error) {
	return wrapImageCellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
