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
// extern void _gotk4_gtk3_EntryBuffer_ConnectInsertedText(gpointer, guint, gchar*, guint, guintptr);
// extern void _gotk4_gtk3_EntryBuffer_ConnectDeletedText(gpointer, guint, guint, guintptr);
import "C"

// GType values.
var (
	GTypeEntryBuffer = coreglib.Type(girepository.MustFind("Gtk", "EntryBuffer").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEntryBuffer, F: marshalEntryBuffer},
	})
}

// EntryBufferOverrides contains methods that are overridable.
type EntryBufferOverrides struct {
}

func defaultEntryBufferOverrides(v *EntryBuffer) EntryBufferOverrides {
	return EntryBufferOverrides{}
}

// EntryBuffer class contains the actual text displayed in a Entry widget.
//
// A single EntryBuffer object can be shared by multiple Entry widgets which
// will then share the same text content, but not the cursor position,
// visibility attributes, icon etc.
//
// EntryBuffer may be derived from. Such a derived class might allow text to be
// stored in an alternate location, such as non-pageable memory, useful in the
// case of important passwords. Or a derived class could integrate with an
// application’s concept of undo/redo.
type EntryBuffer struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EntryBuffer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EntryBuffer, *EntryBufferClass, EntryBufferOverrides](
		GTypeEntryBuffer,
		initEntryBufferClass,
		wrapEntryBuffer,
		defaultEntryBufferOverrides,
	)
}

func initEntryBufferClass(gclass unsafe.Pointer, overrides EntryBufferOverrides, classInitFunc func(*EntryBufferClass)) {
	if classInitFunc != nil {
		class := (*EntryBufferClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEntryBuffer(obj *coreglib.Object) *EntryBuffer {
	return &EntryBuffer{
		Object: obj,
	}
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	return wrapEntryBuffer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDeletedText: this signal is emitted after text is deleted from the
// buffer.
func (v *EntryBuffer) ConnectDeletedText(f func(position, nChars uint)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "deleted-text", false, unsafe.Pointer(C._gotk4_gtk3_EntryBuffer_ConnectDeletedText), f)
}

// ConnectInsertedText: this signal is emitted after text is inserted into the
// buffer.
func (v *EntryBuffer) ConnectInsertedText(f func(position uint, chars string, nChars uint)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "inserted-text", false, unsafe.Pointer(C._gotk4_gtk3_EntryBuffer_ConnectInsertedText), f)
}