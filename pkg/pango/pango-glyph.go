// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeShapeFlags  = coreglib.Type(girepository.MustFind("Pango", "ShapeFlags").RegisteredGType())
	GTypeGlyphString = coreglib.Type(girepository.MustFind("Pango", "GlyphString").RegisteredGType())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShapeFlags, F: marshalShapeFlags},
		coreglib.TypeMarshaler{T: GTypeGlyphString, F: marshalGlyphString},
	})
}

// GlyphUnit: PangoGlyphUnit type is used to store dimensions within Pango.
//
// Dimensions are stored in 1/PANGO_SCALE of a device unit. (A device unit might
// be a pixel for screen display, or a point on a printer.) PANGO_SCALE is
// currently 1024, and may change in the future (unlikely though), but you
// should not depend on its exact value. The PANGO_PIXELS() macro can be used to
// convert from glyph units into device units with correct rounding.
type GlyphUnit = int32

// ShapeFlags flags influencing the shaping process.
//
// PangoShapeFlags can be passed to pango_shape_with_flags().
type ShapeFlags C.guint

const (
	// ShapeNone: default value.
	ShapeNone ShapeFlags = 0b0
	// ShapeRoundPositions: round glyph positions and widths to whole device
	// units. This option should be set if the target renderer can't do subpixel
	// positioning of glyphs.
	ShapeRoundPositions ShapeFlags = 0b1
)

func marshalShapeFlags(p uintptr) (interface{}, error) {
	return ShapeFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ShapeFlags.
func (s ShapeFlags) String() string {
	if s == 0 {
		return "ShapeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(29)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case ShapeNone:
			builder.WriteString("None|")
		case ShapeRoundPositions:
			builder.WriteString("RoundPositions|")
		default:
			builder.WriteString(fmt.Sprintf("ShapeFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s ShapeFlags) Has(other ShapeFlags) bool {
	return (s & other) == other
}

// GlyphGeometry: PangoGlyphGeometry structure contains width and positioning
// information for a single glyph.
//
// An instance of this type is always passed by reference.
type GlyphGeometry struct {
	*glyphGeometry
}

// glyphGeometry is the struct that's finalized.
type glyphGeometry struct {
	native unsafe.Pointer
}

var GIRInfoGlyphGeometry = girepository.MustFind("Pango", "GlyphGeometry")

// Width: logical width to use for the the character.
func (g *GlyphGeometry) Width() GlyphUnit {
	offset := GIRInfoGlyphGeometry.StructFieldOffset("width")
	valptr := (*GlyphUnit)(unsafe.Add(g.native, offset))
	var _v GlyphUnit // out
	_v = int32(*valptr)
	type _ = GlyphUnit
	type _ = int32
	return _v
}

// XOffset: horizontal offset from nominal character position.
func (g *GlyphGeometry) XOffset() GlyphUnit {
	offset := GIRInfoGlyphGeometry.StructFieldOffset("x_offset")
	valptr := (*GlyphUnit)(unsafe.Add(g.native, offset))
	var _v GlyphUnit // out
	_v = int32(*valptr)
	type _ = GlyphUnit
	type _ = int32
	return _v
}

// YOffset: vertical offset from nominal character position.
func (g *GlyphGeometry) YOffset() GlyphUnit {
	offset := GIRInfoGlyphGeometry.StructFieldOffset("y_offset")
	valptr := (*GlyphUnit)(unsafe.Add(g.native, offset))
	var _v GlyphUnit // out
	_v = int32(*valptr)
	type _ = GlyphUnit
	type _ = int32
	return _v
}

// GlyphInfo: PangoGlyphInfo structure represents a single glyph with
// positioning information and visual attributes.
//
// An instance of this type is always passed by reference.
type GlyphInfo struct {
	*glyphInfo
}

// glyphInfo is the struct that's finalized.
type glyphInfo struct {
	native unsafe.Pointer
}

var GIRInfoGlyphInfo = girepository.MustFind("Pango", "GlyphInfo")

// GlyphString: PangoGlyphString is used to store strings of glyphs with
// geometry and visual attribute information.
//
// The storage for the glyph information is owned by the structure which
// simplifies memory management.
//
// An instance of this type is always passed by reference.
type GlyphString struct {
	*glyphString
}

// glyphString is the struct that's finalized.
type glyphString struct {
	native unsafe.Pointer
}

var GIRInfoGlyphString = girepository.MustFind("Pango", "GlyphString")

func marshalGlyphString(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &GlyphString{&glyphString{(unsafe.Pointer)(b)}}, nil
}

// GlyphVisAttr: PangoGlyphVisAttr structure communicates information between
// the shaping and rendering phases.
//
// Currently, it contains only cluster start information. yMore attributes may
// be added in the future.
//
// An instance of this type is always passed by reference.
type GlyphVisAttr struct {
	*glyphVisAttr
}

// glyphVisAttr is the struct that's finalized.
type glyphVisAttr struct {
	native unsafe.Pointer
}

var GIRInfoGlyphVisAttr = girepository.MustFind("Pango", "GlyphVisAttr")
