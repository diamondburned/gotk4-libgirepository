// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_CellRendererText_ConnectEdited(gpointer, gchar*, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeCellRendererText = coreglib.Type(girepository.MustFind("Gtk", "CellRendererText").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererText, F: marshalCellRendererText},
	})
}

// CellRendererTextOverrides contains methods that are overridable.
type CellRendererTextOverrides struct {
}

func defaultCellRendererTextOverrides(v *CellRendererText) CellRendererTextOverrides {
	return CellRendererTextOverrides{}
}

// CellRendererText renders text in a cell
//
// A CellRendererText renders a given text in its cell, using the font, color
// and style information provided by its properties. The text will be ellipsized
// if it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererText)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CellRendererText, *CellRendererTextClass, CellRendererTextOverrides](
		GTypeCellRendererText,
		initCellRendererTextClass,
		wrapCellRendererText,
		defaultCellRendererTextOverrides,
	)
}

func initCellRendererTextClass(gclass unsafe.Pointer, overrides CellRendererTextOverrides, classInitFunc func(*CellRendererTextClass)) {
	if classInitFunc != nil {
		class := (*CellRendererTextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellRendererText(obj *coreglib.Object) *CellRendererText {
	return &CellRendererText{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	return wrapCellRendererText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEdited: this signal is emitted after renderer has been edited.
//
// It is the responsibility of the application to update the model and store
// new_text at the position indicated by path.
func (v *CellRendererText) ConnectEdited(f func(path, newText string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "edited", false, unsafe.Pointer(C._gotk4_gtk4_CellRendererText_ConnectEdited), f)
}

// CellRendererTextClass: instance of this type is always passed by reference.
type CellRendererTextClass struct {
	*cellRendererTextClass
}

// cellRendererTextClass is the struct that's finalized.
type cellRendererTextClass struct {
	native unsafe.Pointer
}

var GIRInfoCellRendererTextClass = girepository.MustFind("Gtk", "CellRendererTextClass")
