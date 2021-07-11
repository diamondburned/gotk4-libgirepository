// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_bidi_type_get_type()), F: marshalBidiType},
	})
}

// BidiType: `PangoBidiType` represents the bidirectional character type of a
// Unicode character as specified by the <ulink
// url="http://www.unicode.org/reports/tr9/">Unicode bidirectional
// algorithm</ulink>.
//
// Deprecated: Use fribidi for this information.
type BidiType int

const (
	// L: left-to-Right
	BidiTypeL BidiType = iota
	// Lre: left-to-Right Embedding
	BidiTypeLre
	// Lro: left-to-Right Override
	BidiTypeLro
	// R: right-to-Left
	BidiTypeR
	// Al: right-to-Left Arabic
	BidiTypeAl
	// Rle: right-to-Left Embedding
	BidiTypeRle
	// Rlo: right-to-Left Override
	BidiTypeRlo
	// PDF: pop Directional Format
	BidiTypePDF
	// En: european Number
	BidiTypeEn
	// ES: european Number Separator
	BidiTypeES
	// Et: european Number Terminator
	BidiTypeEt
	// An: arabic Number
	BidiTypeAn
	// Cs: common Number Separator
	BidiTypeCs
	// Nsm: nonspacing Mark
	BidiTypeNsm
	// Bn: boundary Neutral
	BidiTypeBn
	// B: paragraph Separator
	BidiTypeB
	// S: segment Separator
	BidiTypeS
	// Ws: whitespace
	BidiTypeWs
	// On: other Neutrals
	BidiTypeOn
)

func marshalBidiType(p uintptr) (interface{}, error) {
	return BidiType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FindBaseDir searches a string the first character that has a strong
// direction, according to the Unicode bidirectional algorithm.
func FindBaseDir(text string, length int) Direction {
	var _arg1 *C.gchar         // out
	var _arg2 C.gint           // out
	var _cret C.PangoDirection // in

	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	_cret = C.pango_find_base_dir(_arg1, _arg2)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}

// GetMirrorChar returns the mirrored character of a Unicode character.
//
// Mirror characters are determined by the Unicode mirrored property.
//
// Use g_unichar_get_mirror_char() instead; the docs for that function provide
// full details.
func GetMirrorChar(ch uint32, mirroredCh *uint32) bool {
	var _arg1 C.gunichar  // out
	var _arg2 *C.gunichar // out
	var _cret C.gboolean  // in

	_arg1 = C.gunichar(ch)
	_arg2 = (*C.gunichar)(unsafe.Pointer(mirroredCh))

	_cret = C.pango_get_mirror_char(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnicharDirection determines the inherent direction of a character.
//
// The inherent direction is either PANGO_DIRECTION_LTR, PANGO_DIRECTION_RTL, or
// PANGO_DIRECTION_NEUTRAL.
//
// This function is useful to categorize characters into left-to-right letters,
// right-to-left letters, and everything else. If full Unicode bidirectional
// type of a character is needed, [type_func@Pango.BidiType.for_unichar] can be
// used instead.
func UnicharDirection(ch uint32) Direction {
	var _arg1 C.gunichar       // out
	var _cret C.PangoDirection // in

	_arg1 = C.gunichar(ch)

	_cret = C.pango_unichar_direction(_arg1)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}
