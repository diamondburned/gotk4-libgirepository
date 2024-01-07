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
import "C"

// GType values.
var (
	GTypeBinLayout = coreglib.Type(girepository.MustFind("Gtk", "BinLayout").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBinLayout, F: marshalBinLayout},
	})
}

// BinLayoutOverrides contains methods that are overridable.
type BinLayoutOverrides struct {
}

func defaultBinLayoutOverrides(v *BinLayout) BinLayoutOverrides {
	return BinLayoutOverrides{}
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

func init() {
	coreglib.RegisterClassInfo[*BinLayout, *BinLayoutClass, BinLayoutOverrides](
		GTypeBinLayout,
		initBinLayoutClass,
		wrapBinLayout,
		defaultBinLayoutOverrides,
	)
}

func initBinLayoutClass(gclass unsafe.Pointer, overrides BinLayoutOverrides, classInitFunc func(*BinLayoutClass)) {
	if classInitFunc != nil {
		class := (*BinLayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
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

// BinLayoutClass: instance of this type is always passed by reference.
type BinLayoutClass struct {
	*binLayoutClass
}

// binLayoutClass is the struct that's finalized.
type binLayoutClass struct {
	native unsafe.Pointer
}

var GIRInfoBinLayoutClass = girepository.MustFind("Gtk", "BinLayoutClass")
