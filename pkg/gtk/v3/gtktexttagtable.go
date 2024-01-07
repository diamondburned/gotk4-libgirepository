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
// extern void _gotk4_gtk3_TextTagTable_ConnectTagRemoved(gpointer, void*, guintptr);
// extern void _gotk4_gtk3_TextTagTable_ConnectTagChanged(gpointer, void*, gboolean, guintptr);
// extern void _gotk4_gtk3_TextTagTable_ConnectTagAdded(gpointer, void*, guintptr);
import "C"

// GType values.
var (
	GTypeTextTagTable = coreglib.Type(girepository.MustFind("Gtk", "TextTagTable").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextTagTable, F: marshalTextTagTable},
	})
}

type TextTagTableForEach func(tag *TextTag)

// TextTagTableOverrides contains methods that are overridable.
type TextTagTableOverrides struct {
}

func defaultTextTagTableOverrides(v *TextTagTable) TextTagTableOverrides {
	return TextTagTableOverrides{}
}

// TextTagTable: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The GtkTextTagTable implementation of the GtkBuildable interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags:
//
//    <object class="GtkTextTagTable">
//     <child type="tag">
//       <object class="GtkTextTag"/>
//     </child>
//    </object>.
type TextTagTable struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
}

var (
	_ coreglib.Objector = (*TextTagTable)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TextTagTable, *TextTagTableClass, TextTagTableOverrides](
		GTypeTextTagTable,
		initTextTagTableClass,
		wrapTextTagTable,
		defaultTextTagTableOverrides,
	)
}

func initTextTagTableClass(gclass unsafe.Pointer, overrides TextTagTableOverrides, classInitFunc func(*TextTagTableClass)) {
	if classInitFunc != nil {
		class := (*TextTagTableClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTextTagTable(obj *coreglib.Object) *TextTagTable {
	return &TextTagTable{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	return wrapTextTagTable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TextTagTable) ConnectTagAdded(f func(tag *TextTag)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "tag-added", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagAdded), f)
}

func (v *TextTagTable) ConnectTagChanged(f func(tag *TextTag, sizeChanged bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "tag-changed", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagChanged), f)
}

func (v *TextTagTable) ConnectTagRemoved(f func(tag *TextTag)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "tag-removed", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagRemoved), f)
}

// TextTagTableClass: instance of this type is always passed by reference.
type TextTagTableClass struct {
	*textTagTableClass
}

// textTagTableClass is the struct that's finalized.
type textTagTableClass struct {
	native unsafe.Pointer
}

var GIRInfoTextTagTableClass = girepository.MustFind("Gtk", "TextTagTableClass")
