// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TextBufferDeserializeFunc: function that is called to deserialize rich text
// that has been serialized with gtk_text_buffer_serialize(), and insert it at
// @iter.
type TextBufferDeserializeFunc func(registerBuffer TextBuffer, contentBuffer TextBuffer, iter *TextIter, data []byte, createTags bool) (ok bool)

//export gotk4_TextBufferDeserializeFunc
func gotk4_TextBufferDeserializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.guint8, arg4 C.gsize, arg5 C.gboolean, arg6 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var registerBuffer TextBuffer // out
	var contentBuffer TextBuffer  // out
	var iter *TextIter            // out
	var data []byte
	var createTags bool // out

	registerBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TextBuffer)
	contentBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(TextBuffer)
	iter = (*TextIter)(unsafe.Pointer(arg2))
	data = make([]byte, arg4)
	copy(data, unsafe.Slice((*byte)(unsafe.Pointer(arg3)), arg4))
	if arg5 != 0 {
		createTags = true
	}

	fn := v.(TextBufferDeserializeFunc)
	ok := fn(registerBuffer, contentBuffer, iter, data, createTags)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TextBufferSerializeFunc: function that is called to serialize the content of
// a text buffer. It must return the serialized form of the content.
type TextBufferSerializeFunc func(registerBuffer TextBuffer, contentBuffer TextBuffer, start *TextIter, end *TextIter) (length uint, guint8 *byte)

//export gotk4_TextBufferSerializeFunc
func gotk4_TextBufferSerializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.GtkTextIter, arg4 *C.gsize, arg5 C.gpointer) (cret *C.guint8) {
	v := box.Get(uintptr(arg5))
	if v == nil {
		panic(`callback not found`)
	}

	var registerBuffer TextBuffer // out
	var contentBuffer TextBuffer  // out
	var start *TextIter           // out
	var end *TextIter             // out

	registerBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(TextBuffer)
	contentBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(TextBuffer)
	start = (*TextIter)(unsafe.Pointer(arg2))
	end = (*TextIter)(unsafe.Pointer(arg3))

	fn := v.(TextBufferSerializeFunc)
	length, guint8 := fn(registerBuffer, contentBuffer, start, end)

	arg4 = *C.gsize(length)
	cret = (*C.guint8)(unsafe.Pointer(guint8))

	return cret
}
