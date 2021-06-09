// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// Break determines possible line, word, and character breaks for a string of
// Unicode text with a single analysis.
//
// For most purposes you may want to use pango_get_log_attrs().
func Break() {
	C.pango_break()
}

// DefaultBreak: this is the default break algorithm.
//
// It applies Unicode rules without language-specific tailoring, therefore the
// @analyis argument is unused and can be nil.
//
// See pango_tailor_break() for language-specific breaks.
func DefaultBreak(text string, length int, analysis *Analysis, attrs *LogAttr, attrsLen int) {
	var _arg1 *C.gchar
	var _arg2 C.int
	var _arg3 *C.PangoAnalysis
	var _arg4 *C.PangoLogAttr
	var _arg5 C.int

	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	_arg4 = (*C.PangoLogAttr)(unsafe.Pointer(attrs.Native()))
	_arg5 = C.int(attrsLen)

	C.pango_default_break(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// FindParagraphBoundary locates a paragraph boundary in @text.
//
// A boundary is caused by delimiter characters, such as a newline, carriage
// return, carriage return-newline pair, or Unicode paragraph separator
// character. The index of the run of delimiters is returned in
// @paragraph_delimiter_index. The index of the start of the paragrap (index
// after all delimiters) is stored in @next_paragraph_start.
//
// If no delimiters are found, both @paragraph_delimiter_index and
// @next_paragraph_start are filled with the length of @text (an index one off
// the end).
func FindParagraphBoundary(text string, length int) (paragraphDelimiterIndex int, nextParagraphStart int) {
	var _arg1 *C.gchar
	var _arg2 C.gint

	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	var _arg3 C.gint
	var _arg4 C.gint

	C.pango_find_paragraph_boundary(_arg1, _arg2, &_arg3, &_arg4)

	var _paragraphDelimiterIndex int
	var _nextParagraphStart int

	_paragraphDelimiterIndex = (int)(_arg3)
	_nextParagraphStart = (int)(_arg4)

	return _paragraphDelimiterIndex, _nextParagraphStart
}

// GetLogAttrs computes a `PangoLogAttr` for each character in @text.
//
// The @log_attrs array must have one `PangoLogAttr` for each position in @text;
// if @text contains N characters, it has N+1 positions, including the last
// position at the end of the text. @text should be an entire paragraph; logical
// attributes can't be computed without context (for example you need to see
// spaces on either side of a word to know the word is a word).
func GetLogAttrs() {
	C.pango_get_log_attrs()
}

// TailorBreak: apply language-specific tailoring to the breaks in @log_attrs.
//
// The line breaks are assumed to have been produced by [func@default_break].
//
// If @offset is not -1, it is used to apply attributes from @analysis that are
// relevant to line breaking.
func TailorBreak() {
	C.pango_tailor_break()
}

// LogAttr: the `PangoLogAttr` structure stores information about the attributes
// of a single character.
type LogAttr struct {
	native C.PangoLogAttr
}

// WrapLogAttr wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLogAttr(ptr unsafe.Pointer) *LogAttr {
	if ptr == nil {
		return nil
	}

	return (*LogAttr)(ptr)
}

func marshalLogAttr(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapLogAttr(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *LogAttr) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}