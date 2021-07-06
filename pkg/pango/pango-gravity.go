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
		{T: externglib.Type(C.pango_gravity_get_type()), F: marshalGravity},
		{T: externglib.Type(C.pango_gravity_hint_get_type()), F: marshalGravityHint},
	})
}

// Gravity: `PangoGravity` represents the orientation of glyphs in a segment of
// text.
//
// This is useful when rendering vertical text layouts. In those situations, the
// layout is rotated using a non-identity [struct@Pango.Matrix], and then glyph
// orientation is controlled using `PangoGravity`.
//
// Not every value in this enumeration makes sense for every usage of
// `PangoGravity`; for example, PANGO_GRAVITY_AUTO only can be passed to
// [method@Pango.Context.set_base_gravity] and can only be returned by
// [method@Pango.Context.get_base_gravity].
//
// See also: [enum@Pango.GravityHint]
type Gravity int

const (
	// South glyphs stand upright (default)
	GravitySouth Gravity = iota
	// East glyphs are rotated 90 degrees clockwise
	GravityEast
	// North glyphs are upside-down
	GravityNorth
	// West glyphs are rotated 90 degrees counter-clockwise
	GravityWest
	// Auto: gravity is resolved from the context matrix
	GravityAuto
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GravityHint: `PangoGravityHint` defines how horizontal scripts should behave
// in a vertical context.
//
// That is, English excerpts in a vertical paragraph for example.
//
// See also [enum@Pango.Gravity]
type GravityHint int

const (
	// Natural scripts will take their natural gravity based on the base gravity
	// and the script. This is the default.
	GravityHintNatural GravityHint = iota
	// Strong always use the base gravity set, regardless of the script.
	GravityHintStrong
	// Line: for scripts not in their natural direction (eg. Latin in East
	// gravity), choose per-script gravity such that every script respects the
	// line progression. This means, Latin and Arabic will take opposite
	// gravities and both flow top-to-bottom for example.
	GravityHintLine
)

func marshalGravityHint(p uintptr) (interface{}, error) {
	return GravityHint(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
