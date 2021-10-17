// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_mark_get_type()), F: marshalTextMarker},
	})
}

// TextMark: GtkTextMark is a position in a GtkTextbuffer that is preserved
// across modifications.
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
// A GtkTextMark is like a bookmark in a text buffer; it preserves a position in
// the text. You can convert the mark to an iterator using
// gtk.TextBuffer.GetIterAtMark(). Unlike iterators, marks remain valid across
// buffer mutations, because their behavior is defined when text is inserted or
// deleted. When text containing a mark is deleted, the mark remains in the
// position originally occupied by the deleted text. When text is inserted at a
// mark, a mark with “left gravity” will be moved to the beginning of the
// newly-inserted text, and a mark with “right gravity” will be moved to the
// end.
//
// Note that “left” and “right” here refer to logical direction (left is the
// toward the start of the buffer); in some languages such as Hebrew the
// logically-leftmost text is not actually on the left when displayed.
//
// Marks are reference counted, but the reference count only controls the
// validity of the memory; marks can be deleted from the buffer at any time with
// gtk.TextBuffer.DeleteMark(). Once deleted from the buffer, a mark is
// essentially useless.
//
// Marks optionally have names; these can be convenient to avoid passing the
// GtkTextMark object around.
//
// Marks are typically created using the gtk.TextBuffer.CreateMark() function.
type TextMark struct {
	*externglib.Object
}

func wrapTextMark(obj *externglib.Object) *TextMark {
	return &TextMark{
		Object: obj,
	}
}

func marshalTextMarker(p uintptr) (interface{}, error) {
	return wrapTextMark(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextMark creates a text mark.
//
// Add it to a buffer using gtk.TextBuffer.AddMark(). If name is NULL, the mark
// is anonymous; otherwise, the mark can be retrieved by name using
// gtk.TextBuffer.GetMark(). If a mark has left gravity, and text is inserted at
// the mark’s current location, the mark will be moved to the left of the
// newly-inserted text. If the mark has right gravity (left_gravity = FALSE),
// the mark will end up on the right of newly-inserted text. The standard
// left-to-right cursor is a mark with right gravity (when you type, the cursor
// stays on the right side of the text you’re typing).
//
// The function takes the following parameters:
//
//    - name: mark name or NULL.
//    - leftGravity: whether the mark should have left gravity.
//
func NewTextMark(name string, leftGravity bool) *TextMark {
	var _arg1 *C.char        // out
	var _arg2 C.gboolean     // out
	var _cret *C.GtkTextMark // in

	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if leftGravity {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_text_mark_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(leftGravity)

	var _textMark *TextMark // out

	_textMark = wrapTextMark(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textMark
}

// Buffer gets the buffer this mark is located inside.
//
// Returns NULL if the mark is deleted.
func (mark *TextMark) Buffer() *TextBuffer {
	var _arg0 *C.GtkTextMark   // out
	var _cret *C.GtkTextBuffer // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_buffer(_arg0)
	runtime.KeepAlive(mark)

	var _textBuffer *TextBuffer // out

	_textBuffer = wrapTextBuffer(externglib.Take(unsafe.Pointer(_cret)))

	return _textBuffer
}

// Deleted returns TRUE if the mark has been removed from its buffer.
//
// See gtk.TextBuffer.AddMark() for a way to add it to a buffer again.
func (mark *TextMark) Deleted() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_deleted(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LeftGravity determines whether the mark has left gravity.
func (mark *TextMark) LeftGravity() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_left_gravity(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name returns the mark name.
//
// Returns NULL for anonymous marks.
func (mark *TextMark) Name() string {
	var _arg0 *C.GtkTextMark // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_name(_arg0)
	runtime.KeepAlive(mark)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Visible returns TRUE if the mark is visible.
//
// A cursor is displayed for visible marks.
func (mark *TextMark) Visible() bool {
	var _arg0 *C.GtkTextMark // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_text_mark_get_visible(_arg0)
	runtime.KeepAlive(mark)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

//
// The function takes the following parameters:
//

//
func (mark *TextMark) SetVisible(setting bool) {
	var _arg0 *C.GtkTextMark // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTextMark)(unsafe.Pointer(mark.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_text_mark_set_visible(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(setting)
}
