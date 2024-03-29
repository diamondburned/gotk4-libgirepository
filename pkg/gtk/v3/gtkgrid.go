// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeGrid = coreglib.Type(girepository.MustFind("Gtk", "Grid").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGrid, F: marshalGrid},
	})
}

// GridOverrides contains methods that are overridable.
type GridOverrides struct {
}

func defaultGridOverrides(v *Grid) GridOverrides {
	return GridOverrides{}
}

// Grid is a container which arranges its child widgets in rows and columns,
// with arbitrary positions and horizontal/vertical spans.
//
// Children are added using gtk_grid_attach(). They can span multiple rows or
// columns. It is also possible to add a child next to an existing child, using
// gtk_grid_attach_next_to(). The behaviour of GtkGrid when several children
// occupy the same grid cell is undefined.
//
// GtkGrid can be used like a Box by just using gtk_container_add(), which will
// place children next to each other in the direction determined by the
// Orientable:orientation property. However, if all you want is a single row or
// column, then Box is the preferred widget.
//
//
// CSS nodes
//
// GtkGrid uses a single CSS node with name grid.
type Grid struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Orientable
}

var (
	_ Containerer       = (*Grid)(nil)
	_ coreglib.Objector = (*Grid)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Grid, *GridClass, GridOverrides](
		GTypeGrid,
		initGridClass,
		wrapGrid,
		defaultGridOverrides,
	)
}

func initGridClass(gclass unsafe.Pointer, overrides GridOverrides, classInitFunc func(*GridClass)) {
	if classInitFunc != nil {
		class := (*GridClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGrid(obj *coreglib.Object) *Grid {
	return &Grid{
		Container: Container{
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
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalGrid(p uintptr) (interface{}, error) {
	return wrapGrid(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GridClass: instance of this type is always passed by reference.
type GridClass struct {
	*gridClass
}

// gridClass is the struct that's finalized.
type gridClass struct {
	native unsafe.Pointer
}

var GIRInfoGridClass = girepository.MustFind("Gtk", "GridClass")
