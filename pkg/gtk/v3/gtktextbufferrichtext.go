// Code generated by girgen. DO NOT EDIT.

package gtk

import ()

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// TextBufferSerializeFunc: function that is called to serialize the content of
// a text buffer. It must return the serialized form of the content.
type TextBufferSerializeFunc func(registerBuffer, contentBuffer *TextBuffer, start, end *TextIter) (length uint, guint8 *byte)
