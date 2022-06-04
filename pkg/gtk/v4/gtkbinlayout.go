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
import "C"

// glib.Type values for gtkbinlayout.go.
var GTypeBinLayout = coreglib.Type(C.gtk_bin_layout_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBinLayout, F: marshalBinLayout},
	})
}

// BinLayoutOverrider contains methods that are overridable.
type BinLayoutOverrider interface {
}

// BinLayout: GtkBinLayout is a GtkLayoutManager subclass useful for create
// "bins" of widgets.
//
// GtkBinLayout will stack each child of a widget on top of each other, using
// the gtk.Widget:hexpand, gtk.Widget:vexpand, gtk.Widget:halign, and
// gtk.Widget:valign properties of each child to determine where they should be
// positioned.
type BinLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*BinLayout)(nil)
)

func classInitBinLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBinLayout(obj *coreglib.Object) *BinLayout {
	return &BinLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalBinLayout(p uintptr) (interface{}, error) {
	return wrapBinLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBinLayout creates a new GtkBinLayout instance.
//
// The function returns the following values:
//
//    - binLayout: newly created GtkBinLayout.
//
func NewBinLayout() *BinLayout {
	_gret := girepository.MustFind("Gtk", "BinLayout").InvokeMethod("new_BinLayout", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _binLayout *BinLayout // out

	_binLayout = wrapBinLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _binLayout
}
