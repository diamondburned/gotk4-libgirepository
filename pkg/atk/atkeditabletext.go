// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_editable_text_get_type()), F: marshalEditableText},
	})
}

// EditableText should be implemented by UI components which contain text which
// the user can edit, via the Object corresponding to that component (see
// Object).
//
// EditableText is a subclass of Text, and as such, an object which implements
// EditableText is by definition an Text implementor as well.
//
// See also: Text
type EditableText interface {
	gextras.Objector

	// CopyText: copy text from @start_pos up to, but not including @end_pos to
	// the clipboard.
	CopyText(startPos int, endPos int)
	// CutText: copy text from @start_pos up to, but not including @end_pos to
	// the clipboard and then delete from the widget.
	CutText(startPos int, endPos int)
	// DeleteText: delete text @start_pos up to, but not including @end_pos.
	DeleteText(startPos int, endPos int)
	// InsertText: insert text at a given position.
	InsertText(_string string, length int, position *int)
	// PasteText: paste text from clipboard to specified @position.
	PasteText(position int)
	// SetTextContents: set text contents of @text.
	SetTextContents(_string string)
}

// editableText implements the EditableText interface.
type editableText struct {
	gextras.Objector
}

var _ EditableText = (*editableText)(nil)

// WrapEditableText wraps a GObject to a type that implements
// interface EditableText. It is primarily used internally.
func WrapEditableText(obj *externglib.Object) EditableText {
	return editableText{
		Objector: obj,
	}
}

func marshalEditableText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEditableText(obj), nil
}

func (t editableText) CopyText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_copy_text(_arg0, _arg1, _arg2)
}

func (t editableText) CutText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_cut_text(_arg0, _arg1, _arg2)
}

func (t editableText) DeleteText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_delete_text(_arg0, _arg1, _arg2)
}

func (t editableText) InsertText(_string string, length int, position *int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 *C.gchar           // out
	var _arg2 C.gint             // out
	var _arg3 *C.gint            // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)
	_arg3 = (*C.gint)(unsafe.Pointer(position))

	C.atk_editable_text_insert_text(_arg0, _arg1, _arg2, _arg3)
}

func (t editableText) PasteText(position int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(position)

	C.atk_editable_text_paste_text(_arg0, _arg1)
}

func (t editableText) SetTextContents(_string string) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_editable_text_set_text_contents(_arg0, _arg1)
}
