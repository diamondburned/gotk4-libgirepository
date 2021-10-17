// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_text_cell_accessible_get_type()), F: marshalTextCellAccessibler},
	})
}

type TextCellAccessible struct {
	RendererCellAccessible

	atk.Text
	*externglib.Object
}

func wrapTextCellAccessible(obj *externglib.Object) *TextCellAccessible {
	return &TextCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				TableCell: atk.TableCell{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Object: obj,
			},
		},
		Text: atk.Text{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalTextCellAccessibler(p uintptr) (interface{}, error) {
	return wrapTextCellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (*TextCellAccessible) privateTextCellAccessible() {}
