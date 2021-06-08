// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_glyph_string_get_type()), F: marshalGlyphString},
	})
}

// GlyphUnit: the `PangoGlyphUnit` type is used to store dimensions within
// Pango.
//
// Dimensions are stored in 1/PANGO_SCALE of a device unit. (A device unit might
// be a pixel for screen display, or a point on a printer.) PANGO_SCALE is
// currently 1024, and may change in the future (unlikely though), but you
// should not depend on its exact value. The PANGO_PIXELS() macro can be used to
// convert from glyph units into device units with correct rounding.
type GlyphUnit int32

// ReorderItems: reorder items from logical order to visual order.
//
// The visual order is determined from the associated directional levels of the
// items. The original list is unmodified.
func ReorderItems(logicalItems *glib.List) *glib.List {
	var arg1 *C.GList

	arg1 = (*C.GList)(unsafe.Pointer(logicalItems.Native()))

	cret := new(C.GList)
	var goret *glib.List

	cret = C.pango_reorder_items(arg1)

	goret = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Shape: convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@itemize], convert the characters into glyphs. You may
// also pass in only a substring of the item from [func@itemize].
//
// It is recommended that you use [func@shape_full] instead, since that API
// allows for shaping interaction happening across text item boundaries.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@itemize] have indices that are relative to the entire paragraph, so you
// need to subtract the item offset from their indices before calling
// [func@shape].
func Shape(text string, length int, analysis *Analysis, glyphs *GlyphString) {
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.PangoAnalysis
	var arg4 *C.PangoGlyphString

	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)
	arg3 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	arg4 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_shape(arg1, arg2, arg3, arg4)
}

// ShapeFull: convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@itemize], convert the characters into glyphs. You may
// also pass in only a substring of the item from [func@itemize].
//
// This is similar to [func@shape], except it also can optionally take the full
// paragraph text as input, which will then be used to perform certain
// cross-item shaping interactions. If you have access to the broader text of
// which @item_text is part of, provide the broader text as @paragraph_text. If
// @paragraph_text is nil, item text is used instead.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@itemize] have indices that are relative to the entire paragraph, so you
// do not pass the full paragraph text as @paragraph_text, you need to subtract
// the item offset from their indices before calling [func@shape_full].
func ShapeFull(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString) {
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.char
	var arg4 C.int
	var arg5 *C.PangoAnalysis
	var arg6 *C.PangoGlyphString

	arg1 = (*C.char)(C.CString(itemText))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(itemLength)
	arg3 = (*C.char)(C.CString(paragraphText))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.int(paragraphLength)
	arg5 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	arg6 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_shape_full(arg1, arg2, arg3, arg4, arg5, arg6)
}

// ShapeWithFlags: convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@itemize], convert the characters into glyphs. You may
// also pass in only a substring of the item from [func@itemize].
//
// This is similar to [func@shape_full], except it also takes flags that can
// influence the shaping process.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@itemize] have indices that are relative to the entire paragraph, so you
// do not pass the full paragraph text as @paragraph_text, you need to subtract
// the item offset from their indices before calling [func@shape_with_flags].
func ShapeWithFlags(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString, flags ShapeFlags) {
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.char
	var arg4 C.int
	var arg5 *C.PangoAnalysis
	var arg6 *C.PangoGlyphString
	var arg7 C.PangoShapeFlags

	arg1 = (*C.char)(C.CString(itemText))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(itemLength)
	arg3 = (*C.char)(C.CString(paragraphText))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.int(paragraphLength)
	arg5 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	arg6 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))
	arg7 = (C.PangoShapeFlags)(flags)

	C.pango_shape_with_flags(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GlyphGeometry: the `PangoGlyphGeometry` structure contains width and
// positioning information for a single glyph.
type GlyphGeometry struct {
	native C.PangoGlyphGeometry
}

// WrapGlyphGeometry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyphGeometry(ptr unsafe.Pointer) *GlyphGeometry {
	if ptr == nil {
		return nil
	}

	return (*GlyphGeometry)(ptr)
}

func marshalGlyphGeometry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyphGeometry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *GlyphGeometry) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// GlyphInfo: a `PangoGlyphInfo` structure represents a single glyph with
// positioning information and visual attributes.
type GlyphInfo struct {
	native C.PangoGlyphInfo
}

// WrapGlyphInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyphInfo(ptr unsafe.Pointer) *GlyphInfo {
	if ptr == nil {
		return nil
	}

	return (*GlyphInfo)(ptr)
}

func marshalGlyphInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyphInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *GlyphInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// Geometry gets the field inside the struct.
func (g *GlyphInfo) Geometry() GlyphGeometry {
	var v GlyphGeometry
	v = *WrapGlyphGeometry(unsafe.Pointer(&g.native.geometry))
	return v
}

// Attr gets the field inside the struct.
func (g *GlyphInfo) Attr() GlyphVisAttr {
	var v GlyphVisAttr
	v = *WrapGlyphVisAttr(unsafe.Pointer(&g.native.attr))
	return v
}

// GlyphString: a `PangoGlyphString` is used to store strings of glyphs with
// geometry and visual attribute information.
//
// The storage for the glyph information is owned by the structure which
// simplifies memory management.
type GlyphString struct {
	native C.PangoGlyphString
}

// WrapGlyphString wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyphString(ptr unsafe.Pointer) *GlyphString {
	if ptr == nil {
		return nil
	}

	return (*GlyphString)(ptr)
}

func marshalGlyphString(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyphString(unsafe.Pointer(b)), nil
}

// NewGlyphString constructs a struct GlyphString.
func NewGlyphString() *GlyphString {
	cret := new(C.PangoGlyphString)
	var goret *GlyphString

	cret = C.pango_glyph_string_new()

	goret = WrapGlyphString(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *GlyphString) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Native returns the underlying C source pointer.
func (g *GlyphString) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// NumGlyphs gets the field inside the struct.
func (g *GlyphString) NumGlyphs() int {
	var v int
	v = int(g.native.num_glyphs)
	return v
}

// LogClusters gets the field inside the struct.
func (g *GlyphString) LogClusters() int {
	var v int
	v = int(g.native.log_clusters)
	return v
}

// Copy: copy a glyph string and associated storage.
func (s *GlyphString) Copy() *GlyphString {
	var arg0 *C.PangoGlyphString

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(s.Native()))

	cret := new(C.PangoGlyphString)
	var goret *GlyphString

	cret = C.pango_glyph_string_copy(arg0)

	goret = WrapGlyphString(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret, func(v *GlyphString) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret
}

// Extents: compute the logical and ink extents of a glyph string.
//
// See the documentation for [method@Pango.Font.get_glyph_extents] for details
// about the interpretation of the rectangles.
//
// Examples of logical (red) and ink (green) rects:
//
// ! (rects1.png) ! (rects2.png)
func (g *GlyphString) Extents(font Font) (inkRect *Rectangle, logicalRect *Rectangle) {
	var arg0 *C.PangoGlyphString
	var arg1 *C.PangoFont

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(g.Native()))
	arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	arg2 := new(C.PangoRectangle)
	var ret2 *Rectangle
	arg3 := new(C.PangoRectangle)
	var ret3 *Rectangle

	C.pango_glyph_string_extents(arg0, arg1, arg2, arg3)

	ret2 = WrapRectangle(unsafe.Pointer(arg2))
	ret3 = WrapRectangle(unsafe.Pointer(arg3))

	return ret2, ret3
}

// ExtentsRange computes the extents of a sub-portion of a glyph string.
//
// The extents are relative to the start of the glyph string range (the origin
// of their coordinate system is at the start of the range, not at the start of
// the entire glyph string).
func (g *GlyphString) ExtentsRange(start int, end int, font Font) (inkRect *Rectangle, logicalRect *Rectangle) {
	var arg0 *C.PangoGlyphString
	var arg1 C.int
	var arg2 C.int
	var arg3 *C.PangoFont

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(g.Native()))
	arg1 = C.int(start)
	arg2 = C.int(end)
	arg3 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	arg4 := new(C.PangoRectangle)
	var ret4 *Rectangle
	arg5 := new(C.PangoRectangle)
	var ret5 *Rectangle

	C.pango_glyph_string_extents_range(arg0, arg1, arg2, arg3, arg4, arg5)

	ret4 = WrapRectangle(unsafe.Pointer(arg4))
	ret5 = WrapRectangle(unsafe.Pointer(arg5))

	return ret4, ret5
}

// Free: free a glyph string and associated storage.
func (s *GlyphString) Free() {
	var arg0 *C.PangoGlyphString

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(s.Native()))

	C.pango_glyph_string_free(arg0)
}

// Width computes the logical width of the glyph string.
//
// This can also be computed using [method@Pango.GlyphString.extents]. However,
// since this only computes the width, it's much faster. This is in fact only a
// convenience function that computes the sum of @geometry.width for each glyph
// in the @glyphs.
func (g *GlyphString) Width() int {
	var arg0 *C.PangoGlyphString

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(g.Native()))

	var cret C.int
	var goret int

	cret = C.pango_glyph_string_get_width(arg0)

	goret = int(cret)

	return goret
}

// IndexToX converts from character position to x position.
//
// The X position is measured from the left edge of the run. Character positions
// are computed by dividing up each cluster into equal portions.
func (g *GlyphString) IndexToX(text string, length int, analysis *Analysis, index_ int, trailing bool) int {
	var arg0 *C.PangoGlyphString
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.PangoAnalysis
	var arg4 C.int
	var arg5 C.gboolean

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(g.Native()))
	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)
	arg3 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	arg4 = C.int(index_)
	if trailing {
		arg5 = C.gboolean(1)
	}

	arg6 := new(C.int)
	var ret6 int

	C.pango_glyph_string_index_to_x(arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	ret6 = int(*arg6)

	return ret6
}

// SetSize: resize a glyph string to the given length.
func (s *GlyphString) SetSize(newLen int) {
	var arg0 *C.PangoGlyphString
	var arg1 C.gint

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(s.Native()))
	arg1 = C.gint(newLen)

	C.pango_glyph_string_set_size(arg0, arg1)
}

// XToIndex: convert from x offset to character position.
//
// Character positions are computed by dividing up each cluster into equal
// portions. In scripts where positioning within a cluster is not allowed (such
// as Thai), the returned value may not be a valid cursor position; the caller
// must combine the result with the logical attributes for the text to compute
// the valid cursor position.
func (g *GlyphString) XToIndex(text string, length int, analysis *Analysis, xPos int) (index_ int, trailing int) {
	var arg0 *C.PangoGlyphString
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.PangoAnalysis
	var arg4 C.int

	arg0 = (*C.PangoGlyphString)(unsafe.Pointer(g.Native()))
	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(length)
	arg3 = (*C.PangoAnalysis)(unsafe.Pointer(analysis.Native()))
	arg4 = C.int(xPos)

	arg5 := new(C.int)
	var ret5 int
	arg6 := new(C.int)
	var ret6 int

	C.pango_glyph_string_x_to_index(arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	ret5 = int(*arg5)
	ret6 = int(*arg6)

	return ret5, ret6
}

// GlyphVisAttr: a `PangoGlyphVisAttr` structure communicates information
// between the shaping and rendering phases.
//
// Currently, it contains only cluster start information. yMore attributes may
// be added in the future.
type GlyphVisAttr struct {
	native C.PangoGlyphVisAttr
}

// WrapGlyphVisAttr wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyphVisAttr(ptr unsafe.Pointer) *GlyphVisAttr {
	if ptr == nil {
		return nil
	}

	return (*GlyphVisAttr)(ptr)
}

func marshalGlyphVisAttr(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyphVisAttr(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *GlyphVisAttr) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}
