// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
	GTypeCellRendererPixbuf = coreglib.Type(girepository.MustFind("Gtk", "CellRendererPixbuf").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererPixbuf, F: marshalCellRendererPixbuf},
	})
}

// CellRendererPixbuf renders a pixbuf in a cell
//
// A CellRendererPixbuf can be used to render an image in a cell. It allows to
// render either a given Pixbuf (set via the CellRendererPixbuf:pixbuf property)
// or a named icon (set via the CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is TRUE. If
// the CellRenderer:is-expanded property is TRUE and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is FALSE and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererPixbuf)(nil)
)

func wrapCellRendererPixbuf(obj *coreglib.Object) *CellRendererPixbuf {
	return &CellRendererPixbuf{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	return wrapCellRendererPixbuf(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
