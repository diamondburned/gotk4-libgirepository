// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.pango_coverage_level_get_type()), F: marshalCoverageLevel},
		{T: externglib.Type(C.pango_coverage_get_type()), F: marshalCoverage},
	})
}

// CoverageLevel: `PangoCoverageLevel` is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel int

const (
	// none: the character is not representable with the font.
	CoverageLevelNone CoverageLevel = 0
	// fallback: the character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageLevelFallback CoverageLevel = 1
	// approximate: the character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for the
	// current script.
	CoverageLevelApproximate CoverageLevel = 2
	// exact: the character is represented as the correct graphical form.
	CoverageLevelExact CoverageLevel = 3
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Coverage structure is a map from Unicode characters to CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage interface {
	gextras.Objector

	// CopyCoverage: copy an existing `PangoCoverage`.
	CopyCoverage() Coverage
	// GetCoverage: determine whether a particular index is covered by
	// @coverage.
	GetCoverage(index_ int) CoverageLevel
	// MaxCoverage: set the coverage for each index in @coverage to be the max
	// (better) value of the current coverage for the index and the coverage for
	// the corresponding index in @other.
	//
	// Deprecated: since version 1.44.
	MaxCoverage(other Coverage)
	// RefCoverage: increase the reference count on the `PangoCoverage` by one.
	RefCoverage() Coverage
	// SetCoverage: modify a particular index within @coverage
	SetCoverage(index_ int, level CoverageLevel)
	// ToBytesCoverage: convert a `PangoCoverage` structure into a flat binary
	// format.
	//
	// Deprecated: since version 1.44.
	ToBytesCoverage() []byte
	// UnrefCoverage: decrease the reference count on the `PangoCoverage` by
	// one.
	//
	// If the result is zero, free the coverage and all associated memory.
	UnrefCoverage()
}

// coverage implements the Coverage class.
type coverage struct {
	gextras.Objector
}

// WrapCoverage wraps a GObject to the right type. It is
// primarily used internally.
func WrapCoverage(obj *externglib.Object) Coverage {
	return coverage{
		Objector: obj,
	}
}

func marshalCoverage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCoverage(obj), nil
}

// NewCoverage: create a new `PangoCoverage`
func NewCoverage() Coverage {
	var _cret *C.PangoCoverage // in

	_cret = C.pango_coverage_new()

	var _coverage Coverage // out

	_coverage = WrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _coverage
}

func (c coverage) CopyCoverage() Coverage {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	_cret = C.pango_coverage_copy(_arg0)

	var _ret Coverage // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Coverage)

	return _ret
}

func (c coverage) GetCoverage(index_ int) CoverageLevel {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _cret C.PangoCoverageLevel // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(index_)

	_cret = C.pango_coverage_get(_arg0, _arg1)

	var _coverageLevel CoverageLevel // out

	_coverageLevel = CoverageLevel(_cret)

	return _coverageLevel
}

func (c coverage) MaxCoverage(other Coverage) {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.PangoCoverage)(unsafe.Pointer(other.Native()))

	C.pango_coverage_max(_arg0, _arg1)
}

func (c coverage) RefCoverage() Coverage {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	_cret = C.pango_coverage_ref(_arg0)

	var _ret Coverage // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Coverage)

	return _ret
}

func (c coverage) SetCoverage(index_ int, level CoverageLevel) {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _arg2 C.PangoCoverageLevel // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(index_)
	_arg2 = C.PangoCoverageLevel(level)

	C.pango_coverage_set(_arg0, _arg1, _arg2)
}

func (c coverage) ToBytesCoverage() []byte {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.guchar
	var _arg2 C.int // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	C.pango_coverage_to_bytes(_arg0, &_arg1, &_arg2)

	var _bytes []byte

	_bytes = unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_bytes, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})

	return _bytes
}

func (c coverage) UnrefCoverage() {
	var _arg0 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	C.pango_coverage_unref(_arg0)
}
